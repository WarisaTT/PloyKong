package handlers

import (
	"database/sql"
	"encoding/json"

	"ploykong-api/internal/middleware"
	"ploykong-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

type PortfolioHandler struct {
	db *sql.DB
}

func NewPortfolioHandler(db *sql.DB) *PortfolioHandler {
	return &PortfolioHandler{db: db}
}

// ─── List My Portfolios ───────────────────────────────────────────────────────

func (h *PortfolioHandler) ListMyPortfolios(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)

	rows, err := h.db.Query(
		`SELECT id, slug, title, is_published, view_count, created_at, updated_at
		 FROM portfolios WHERE user_id = ? AND deleted_at IS NULL ORDER BY updated_at DESC`,
		claims.UserID,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch portfolios")
	}
	defer rows.Close()

	portfolios := []fiber.Map{}
	for rows.Next() {
		var p models.Portfolio
		if err := rows.Scan(&p.ID, &p.Slug, &p.Title, &p.IsPublished, &p.ViewCount, &p.CreatedAt, &p.UpdatedAt); err != nil {
			continue
		}
		portfolios = append(portfolios, fiber.Map{
			"id":           p.ID,
			"slug":         p.Slug,
			"title":        p.Title,
			"is_published": p.IsPublished,
			"view_count":   p.ViewCount,
			"url":          "https://" + p.Slug + ".ploykong.com",
			"created_at":   p.CreatedAt,
			"updated_at":   p.UpdatedAt,
		})
	}

	return c.JSON(fiber.Map{"success": true, "data": portfolios})
}

// ─── Create Portfolio ─────────────────────────────────────────────────────────

type CreatePortfolioRequest struct {
	Slug  string              `json:"slug"`
	Title string              `json:"title"`
	Theme *models.ThemeConfig `json:"theme"`
}

func (h *PortfolioHandler) Create(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)

	var req CreatePortfolioRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.Slug == "" {
		return fiber.NewError(fiber.StatusBadRequest, "slug is required")
	}

	// Check slug availability
	var count int
	h.db.QueryRow("SELECT COUNT(*) FROM portfolios WHERE slug = ?", req.Slug).Scan(&count)
	if count > 0 {
		return fiber.NewError(fiber.StatusConflict, "slug already taken")
	}

	// Default theme
	if req.Theme == nil {
		req.Theme = &models.ThemeConfig{
			Mode:         "dark",
			PrimaryColor: "#4F46E5",
			Font:         "Syne",
			Layout:       "centered",
		}
	}

	themeJSON, _ := json.Marshal(req.Theme)
	id := ulid.Make().String()
	title := req.Title
	if title == "" {
		title = "My Portfolio"
	}

	_, err := h.db.Exec(
		"INSERT INTO portfolios (id, user_id, slug, title, theme) VALUES (?, ?, ?, ?, ?)",
		id, claims.UserID, req.Slug, title, string(themeJSON),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create portfolio")
	}

	// Auto-create starter sections
	h.createStarterSections(id)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":    id,
			"slug":  req.Slug,
			"title": title,
			"theme": req.Theme,
			"url":   "https://" + req.Slug + ".ploykong.com",
		},
	})
}

// ─── Get Portfolio by ID ──────────────────────────────────────────────────────

func (h *PortfolioHandler) GetByID(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	id := c.Params("id")

	var p models.Portfolio
	var themeRaw []byte
	var descNull, seoTitleNull, seoDescNull sql.NullString

	err := h.db.QueryRow(
		`SELECT id, user_id, slug, title, description, theme, seo_title, seo_desc,
		 is_published, view_count, expires_at, created_at, updated_at
		 FROM portfolios WHERE id = ? AND user_id = ? AND deleted_at IS NULL`,
		id, claims.UserID,
	).Scan(
		&p.ID, &p.UserID, &p.Slug, &p.Title, &descNull, &themeRaw,
		&seoTitleNull, &seoDescNull, &p.IsPublished, &p.ViewCount,
		&p.ExpiresAt, &p.CreatedAt, &p.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return fiber.NewError(fiber.StatusNotFound, "portfolio not found")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database error")
	}

	p.ScanTheme(themeRaw)
	p.Description = descNull
	p.SEOTitle = seoTitleNull
	p.SEODesc = seoDescNull

	// Load sections
	sections, _ := h.loadSections(id)
	p.Sections = sections

	return c.JSON(fiber.Map{"success": true, "data": p})
}

// ─── Update Portfolio ─────────────────────────────────────────────────────────

type UpdatePortfolioRequest struct {
	Title       *string             `json:"title"`
	Description *string             `json:"description"`
	Theme       *models.ThemeConfig `json:"theme"`
	SEOTitle    *string             `json:"seo_title"`
	SEODesc     *string             `json:"seo_desc"`
}

func (h *PortfolioHandler) Update(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	id := c.Params("id")

	var req UpdatePortfolioRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	// Verify ownership
	var count int
	h.db.QueryRow("SELECT COUNT(*) FROM portfolios WHERE id = ? AND user_id = ? AND deleted_at IS NULL", id, claims.UserID).Scan(&count)
	if count == 0 {
		return fiber.NewError(fiber.StatusNotFound, "portfolio not found")
	}

	var themeParam interface{}
	if req.Theme != nil {
		themeJSON, _ := json.Marshal(req.Theme)
		themeParam = string(themeJSON)
	}

	_, err := h.db.Exec(
		`UPDATE portfolios SET title = COALESCE(?, title),
		 description = COALESCE(?, description),
		 theme = COALESCE(?, theme),
		 seo_title = COALESCE(?, seo_title),
		 seo_desc = COALESCE(?, seo_desc)
		 WHERE id = ?`,
		req.Title, req.Description, themeParam, req.SEOTitle, req.SEODesc, id,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update portfolio")
	}

	return c.JSON(fiber.Map{"success": true, "message": "portfolio updated"})
}

// ─── Delete Portfolio ─────────────────────────────────────────────────────────

func (h *PortfolioHandler) Delete(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	id := c.Params("id")

	result, err := h.db.Exec(
		"UPDATE portfolios SET deleted_at = NOW() WHERE id = ? AND user_id = ? AND deleted_at IS NULL",
		id, claims.UserID,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete portfolio")
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "portfolio not found")
	}

	return c.JSON(fiber.Map{"success": true, "message": "portfolio deleted"})
}

// ─── Publish / Unpublish ──────────────────────────────────────────────────────

func (h *PortfolioHandler) Publish(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	id := c.Params("id")

	_, err := h.db.Exec(
		"UPDATE portfolios SET is_published = TRUE WHERE id = ? AND user_id = ? AND deleted_at IS NULL",
		id, claims.UserID,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to publish")
	}
	return c.JSON(fiber.Map{"success": true, "message": "portfolio published"})
}

func (h *PortfolioHandler) Unpublish(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	id := c.Params("id")

	_, err := h.db.Exec(
		"UPDATE portfolios SET is_published = FALSE WHERE id = ? AND user_id = ? AND deleted_at IS NULL",
		id, claims.UserID,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to unpublish")
	}
	return c.JSON(fiber.Map{"success": true, "message": "portfolio unpublished"})
}

// ─── Export PDF (placeholder — integrate with headless browser) ───────────────

func (h *PortfolioHandler) ExportPDF(c *fiber.Ctx) error {
	// In production: use chromedp or wkhtmltopdf to render the portfolio URL
	// and stream PDF back
	return c.JSON(fiber.Map{
		"success": true,
		"message": "PDF generation queued",
		"data": fiber.Map{
			"download_url": "https://api.ploykong.com/exports/" + c.Params("id") + ".pdf",
		},
	})
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func (h *PortfolioHandler) loadSections(portfolioID string) ([]models.Section, error) {
	rows, err := h.db.Query(
		"SELECT id, portfolio_id, type, position, data, is_visible, created_at, updated_at FROM sections WHERE portfolio_id = ? ORDER BY position ASC",
		portfolioID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sections []models.Section
	for rows.Next() {
		var s models.Section
		var dataRaw []byte
		if err := rows.Scan(&s.ID, &s.PortfolioID, &s.Type, &s.Position, &dataRaw, &s.IsVisible, &s.CreatedAt, &s.UpdatedAt); err != nil {
			continue
		}
		s.Data = json.RawMessage(dataRaw)
		sections = append(sections, s)
	}
	return sections, nil
}

func (h *PortfolioHandler) createStarterSections(portfolioID string) {
	starters := []struct {
		sType string
		pos   int
		data  interface{}
	}{
		{"hero", 0, models.HeroData{Name: "Your Name", Role: "Your Role", Tagline: "Welcome to my portfolio", ShowHireMe: true}},
		{"skills", 1, models.SkillsData{Items: []models.SkillItem{{Name: "Go", Level: 80, Category: "Backend"}, {Name: "Vue.js", Level: 75, Category: "Frontend"}}}},
		{"projects", 2, models.ProjectsData{Items: []models.ProjectItem{{Title: "My Project", Description: "Project description", Tags: []string{"Go", "Vue.js"}}}}},
		{"contact", 3, map[string]string{"email": "you@example.com"}},
	}

	for _, s := range starters {
		dataJSON, _ := json.Marshal(s.data)
		h.db.Exec(
			"INSERT INTO sections (id, portfolio_id, type, position, data) VALUES (?, ?, ?, ?, ?)",
			ulid.Make().String(), portfolioID, s.sType, s.pos, string(dataJSON),
		)
	}
}
