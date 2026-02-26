package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"ploykong-api/internal/middleware"
	"ploykong-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"github.com/signintech/gopdf"
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
		return fiber.NewError(fiber.StatusConflict, "URL already exists")
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

// ─── Helpers ──────────────────────────────────────────────────────────────────

func (h *PortfolioHandler) loadSections(portfolioID string) ([]models.Section, error) {
	rows, err := h.db.Query(
		`SELECT id, portfolio_id, type, position, data, is_visible 
		 FROM sections WHERE portfolio_id = ? ORDER BY position ASC`,
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
		if err := rows.Scan(&s.ID, &s.PortfolioID, &s.Type, &s.Position, &dataRaw, &s.IsVisible); err != nil {
			continue
		}
		s.Data = json.RawMessage(dataRaw)
		sections = append(sections, s)
	}
	return sections, nil
}

// ─── Generate PDF ─────────────────────────────────────────────────────────────

// ─── Color Palette ────────────────────────────────────────────────────────────
// White/light theme — professional, works for all industries
var (
	cBlack   = [3]uint8{17, 24, 39}    // #111827
	cGray600 = [3]uint8{75, 85, 99}    // #4b5563
	cGray400 = [3]uint8{156, 163, 175} // #9ca3af
	cGray200 = [3]uint8{229, 231, 235} // #e5e7eb
	cGray50  = [3]uint8{249, 250, 251} // #f9fafb
	cAccent  = [3]uint8{114, 17, 218}  // #7211daff  — single blue accent
	cAccentL = [3]uint8{239, 246, 255} // #eff6ff  — accent light bg
	cAccentB = [3]uint8{191, 219, 254} // #bfdbfe  — accent border
	cWhite   = [3]uint8{255, 255, 255}
)

// ─── PDF Canvas wrapper ───────────────────────────────────────────────────────

type cv struct {
	p    *gopdf.GoPdf
	w    float64 // page width
	h    float64 // page height
	ml   float64 // margin left
	mr   float64 // margin right
	y    float64 // current y cursor (top-down: 0 = top)
	page int
}

func newCV() *cv {
	return &cv{p: &gopdf.GoPdf{}, w: 595.28, h: 841.89, ml: 50, mr: 50, page: 0}
}

func (c *cv) bodyW() float64 { return c.w - c.ml - c.mr }

// gopdf uses top-left origin; y increases downward
func (c *cv) gy(y float64) float64 { return y } // alias — gopdf IS top-down

func (c *cv) fill(rgb [3]uint8) { c.p.SetFillColor(rgb[0], rgb[1], rgb[2]) }
func (c *cv) stroke(rgb [3]uint8, lw float64) {
	c.p.SetStrokeColor(rgb[0], rgb[1], rgb[2])
	c.p.SetLineWidth(lw)
}

// rect draws a filled rectangle (gopdf top-left origin)
func (c *cv) rect(x, y, w, h float64, fill [3]uint8) {
	c.fill(fill)
	c.p.RectFromUpperLeftWithStyle(x, y, w, h, "F")
}

// rectStroke draws filled + stroked rectangle
func (c *cv) rectS(x, y, w, h float64, fill, strokeC [3]uint8, lw float64) {
	c.fill(fill)
	c.stroke(strokeC, lw)
	c.p.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
}

// hline draws a horizontal rule
func (c *cv) hline(x1, y, x2 float64, col [3]uint8, lw float64) {
	c.stroke(col, lw)
	c.p.Line(x1, y, x2, y)
}

// text draws a string at (x,y) top-left of line
func (c *cv) text(x, y float64, col [3]uint8, font string, size int, s string) {
	c.fill(col)
	_ = c.p.SetFont(font, "", size)
	c.p.SetXY(x, y)
	_ = c.p.Cell(nil, s)
}

// textR draws string right-aligned ending at x
func (c *cv) textR(x, y float64, col [3]uint8, font string, size int, s string) {
	_ = c.p.SetFont(font, "", size)
	w, _ := c.p.MeasureTextWidth(s)
	c.text(x-w, y, col, font, size, s)
}

// measureW returns string width
func (c *cv) measureW(font string, size int, s string) float64 {
	_ = c.p.SetFont(font, "", size)
	w, _ := c.p.MeasureTextWidth(s)
	return w
}

// wrapText draws word-wrapped text, returns final y after last line
func (c *cv) wrapText(x, y float64, col [3]uint8, font string, size int, text string, maxW, lineH float64, maxLines int) float64 {
	_ = c.p.SetFont(font, "", size)
	words := strings.Fields(text)
	line := ""
	drawn := 0
	for _, w := range words {
		test := strings.TrimSpace(line + " " + w)
		tw, _ := c.p.MeasureTextWidth(test)
		if tw > maxW && line != "" {
			if maxLines > 0 && drawn >= maxLines {
				break
			}
			c.text(x, y, col, font, size, line)
			y += lineH
			drawn++
			line = w
		} else {
			line = test
		}
	}
	if line != "" && (maxLines <= 0 || drawn < maxLines) {
		c.text(x, y, col, font, size, line)
		y += lineH
	}
	return y
}

// addPage adds a new page with white background
func (c *cv) addPage() {
	c.p.AddPage()
	c.page++
	c.rect(0, 0, c.w, c.h, cWhite)
	c.y = 24
}

// needsNewPage returns true if content of height h won't fit on the current page
func (c *cv) needsNewPage(h float64) bool {
	return c.y+h > c.h-50
}

// ─── Section title with blue underline ───────────────────────────────────────

func (c *cv) sectionTitle(x float64, title string) {
	c.text(x, c.y, cBlack, "THSarabun", 11, strings.ToUpper(title))
	tw := c.measureW("THSarabun", 11, strings.ToUpper(title))
	c.hline(x, c.y+14, x+tw, cAccent, 1.5)
	c.y += 20
}

// ─── Skill bar ────────────────────────────────────────────────────────────────

func (c *cv) skillBar(x float64, colW float64, name string, pct int, hidePct bool) {
	c.text(x, c.y, cBlack, "THSarabun", 10, name)
	c.y += 13

	if !hidePct {
		// Track
		c.rect(x, c.y, colW, 5, cGray200)
		// Fill
		fillW := colW * float64(pct) / 100
		c.rect(x, c.y, fillW, 5, cAccent)
		c.y += 16
	} else {
		// Just a bit of spacing if there's no progress bar
		c.y += 4
	}
}

// ─── Role pill (header) ───────────────────────────────────────────────────────

func (c *cv) rolePill(x, y float64, role string) {
	tw := c.measureW("THSarabun", 10, role) + 16
	c.rectS(x, y, tw, 20, cAccentL, cAccentB, 0.5)   // Increased height from 16 to 20
	c.text(x+8, y+3, cAccent, "THSarabun", 10, role) // Shifted down from y+3 to y+5
}

// ─── Inline tag pill ──────────────────────────────────────────────────────────

func (c *cv) tagPill(x, y float64, tag string) float64 {
	tw := c.measureW("THSarabun", 8, tag) + 10
	c.rectS(x, y, tw, 14, cGray50, cGray200, 0.5)   // Increased height from 13 to 17
	c.text(x+5, y+1, cGray600, "THSarabun", 8, tag) // Shifted down from y+2 to y+4
	return x + tw + 5
}


func (c *cv) skillPill(x, y float64, tag string) float64 {
	tw := c.measureW("THSarabun", 8, tag) + 10
	c.rectS(x, y, tw, 14, cAccentL, cAccentL, 0.5)   // Increased height from 13 to 17
	c.text(x+5, y+1, cGray600, "THSarabun", 8, tag) // Shifted down from y+2 to y+4
	return x + tw + 5
}

// ─── Font loading ─────────────────────────────────────────────────────────────

func loadFont(p *gopdf.GoPdf) error {
	paths := []string{
		"assets/fonts/THSarabunNew.ttf",
		"./assets/fonts/THSarabunNew.ttf",
		"/app/assets/fonts/THSarabunNew.ttf",
	}
	for _, path := range paths {
		if err := p.AddTTFFont("THSarabun", path); err == nil {
			return nil
		}
	}
	return fmt.Errorf("THSarabunNew.ttf not found")
}

// ─── JSON helpers ─────────────────────────────────────────────────────────────

type jmap = map[string]interface{}

func jstr(m jmap, key string) string {
	if m == nil {
		return ""
	}
	s, _ := m[key].(string)
	return s
}
func jfloat(m jmap, key string, def float64) float64 {
	if m == nil {
		return def
	}
	switch v := m[key].(type) {
	case float64:
		return v
	case int:
		return float64(v)
	}
	return def
}
func jbool(m jmap, key string) bool {
	if m == nil {
		return false
	}
	b, _ := m[key].(bool)
	return b
}
func jslice(m jmap, key string) []interface{} {
	if m == nil {
		return nil
	}
	s, _ := m[key].([]interface{})
	return s
}
func jstrslice(m jmap, key string) []string {
	var out []string
	for _, v := range jslice(m, key) {
		if s, ok := v.(string); ok && s != "" {
			out = append(out, s)
		}
	}
	return out
}
func toJMap(v interface{}) jmap {
	m, _ := v.(jmap)
	return m
}
func firstOf(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

// ─── Handler ──────────────────────────────────────────────────────────────────

func (h *PortfolioHandler) GeneratePDF(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("🔥 PANIC in GeneratePDF: %v\n%s", r, debug.Stack())
			os.WriteFile("panic.log", []byte(fmt.Sprintf("%v\n%s", r, debug.Stack())), 0644)
			c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("panic: %v", r)})
		}
	}()

	claims := middleware.GetUserClaims(c)
	id := c.Params("id")

	// 1. Fetch portfolio ───────────────────────────────────────────────────────
	var port models.Portfolio
	err := h.db.QueryRow(
		`SELECT id, title FROM portfolios WHERE id = ? AND user_id = ? AND deleted_at IS NULL`,
		id, claims.UserID,
	).Scan(&port.ID, &port.Title)
	if err == sql.ErrNoRows {
		return fiber.NewError(fiber.StatusNotFound, "portfolio not found")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database error")
	}

	// 2. Fetch sections ────────────────────────────────────────────────────────
	sections, err := h.loadSections(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to load sections")
	}

	// 3. Parse section data ────────────────────────────────────────────────────
	data := make([]jmap, len(sections))
	for i, sec := range sections {
		if sec.Data != nil {
			json.Unmarshal(sec.Data, &data[i]) //nolint:errcheck
		}
	}

	// 4. Init PDF ──────────────────────────────────────────────────────────────
	cv := newCV()
	cv.p.Start(gopdf.Config{PageSize: gopdf.Rect{W: cv.w, H: cv.h}})
	if err := loadFont(cv.p); err != nil {
		log.Printf("❌ GeneratePDF: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	cv.addPage()

	// ── HEADER ────────────────────────────────────────────────────────────────
	// Extract hero + contact data first
	var heroName, heroRole, heroTagline string
	var contactEmail, contactPhone, contactLocation, contactLinkedin, contactGithub string

	for i, sec := range sections {
		if !sec.IsVisible {
			continue
		}
		switch sec.Type {
		case "hero":
			heroName = firstOf(jstr(data[i], "name"), port.Title)
			heroRole = jstr(data[i], "role")
			heroTagline = jstr(data[i], "tagline")
		case "contact":
			contactEmail = jstr(data[i], "email")
			contactPhone = jstr(data[i], "phone")
			contactLocation = jstr(data[i], "location")
			contactLinkedin = jstr(data[i], "linkedin")
			contactGithub = jstr(data[i], "github")
		}
	}
	if heroName == "" {
		heroName = port.Title
	}

	// Top accent bar (4px)
	cv.rect(0, 0, cv.w, 4, cAccent)
	cv.y = 30 // Moved down for more top spacing

	contactStartY := cv.y
	// Name (large)
	cv.text(cv.ml, cv.y, cBlack, "THSarabun", 24, heroName)
	cv.y += 26

	cv.rect(0, 0, cv.w, 4, cAccent)
	// Role pill
	if heroRole != "" {
		cv.y += 3
		cv.rolePill(cv.ml, cv.y, heroRole)
		cv.y += 20
	} else {
		cv.y += 10
	}

	// Contact info — right-aligned, stacked
	contactLines := []string{}
	if contactEmail != "" {
		contactLines = append(contactLines, contactEmail)
	}
	if contactPhone != "" {
		contactLines = append(contactLines, contactPhone)
	}
	if contactLinkedin != "" {
		contactLines = append(contactLines, contactLinkedin)
	}
	if contactGithub != "" {
		contactLines = append(contactLines, contactGithub)
	}
	if contactLocation != "" {
		contactLines = append(contactLines, contactLocation)
	}

	cy := contactStartY
	for _, cl := range contactLines {
		cv.textR(cv.w-cv.mr, cy, cGray600, "THSarabun", 9, cl)
		cy += 12
	}

	// Tagline
	if heroTagline != "" {
		// Restrict tagline width to prevent overlap with contacts
		taglineMaxW := cv.bodyW() * 0.65
		cv.y = cv.wrapText(cv.ml, cv.y, cGray600, "THSarabun", 10, heroTagline, taglineMaxW, 13, 3)
	}

	// Dynamic spacing: ensure the next section starts below both tagline and contact info
	if cv.y < cy {
		cv.y = cy
	}
	cv.y += 10

	// Divider
	cv.hline(cv.ml, cv.y, cv.w-cv.mr, cGray200, 0.8)
	cv.y += 14

	// ── TWO-COLUMN BODY ───────────────────────────────────────────────────────
	// Left 58% — Experience, Projects, Education, Custom Text
	// Right 38% — Skills, then remaining
	lw := cv.bodyW() * 0.58
	rw := cv.bodyW() * 0.38
	rx := cv.ml + lw + cv.bodyW()*0.04

	bodyTopY := cv.y
	rightY := bodyTopY

	// ── RIGHT: Skills ──────────────────────────────────────────────────────────
	for i, sec := range sections {
		if !sec.IsVisible || sec.Type != "skills" {
			continue
		}
		items := jslice(data[i], "items")
		if len(items) == 0 {
			continue
		}

		hideTitle := false
		if b, ok := data[i]["hide_title"].(bool); ok && b {
			hideTitle = true
		}

		cv.y = rightY
		if !hideTitle {
			cv.sectionTitle(rx, "Skills")
		}

		hidePct := false
		if b, ok := data[i]["hide_percentage"].(bool); ok && b {
			hidePct = true
		}

		// Group by category
		cats := make(map[string][]map[string]interface{})
		var catNames []string
		for _, item := range items {
			im := toJMap(item)
			cat := jstr(im, "category")
			if cat == "" {
				cat = "Other Skills"
			}
			if _, exists := cats[cat]; !exists {
				catNames = append(catNames, cat)
			}
			cats[cat] = append(cats[cat], im)
		}

		for _, cat := range catNames {
			// Category Title
			cv.y += 2
			cv.text(rx, cv.y, cAccent, "THSarabun", 10, cat)
			cv.y += 16

			if hidePct {
				tx := rx
				for _, im := range cats[cat] {
					name := jstr(im, "name")
					if name == "" {
						continue
					}
					// Measure pill width to handle word wrapping
					tw := cv.measureW("THSarabun", 8, name) + 10
					if tx+tw > rx+rw {
						cv.y += 20 // line height for wrapped pills
						tx = rx
					}
					tx = cv.skillPill(tx, cv.y, name) + 4 // 4px extra spacing
				}
				cv.y += 22 // Bottom padding after a category's pills
			} else {
				for _, im := range cats[cat] {
					name := jstr(im, "name")
					pct := int(jfloat(im, "level", 70))
					if name != "" {
						cv.skillBar(rx, rw, name, pct, hidePct)
					}
				}
			}
		}
		rightY = cv.y + 8
	}

	// ── RIGHT: Education ───────────────────────────────────────────────────────
	for i, sec := range sections {
		if !sec.IsVisible || sec.Type != "education" {
			continue
		}
		items := jslice(data[i], "items")
		if len(items) == 0 {
			continue
		}

		hideTitle := false
		if b, ok := data[i]["hide_title"].(bool); ok && b {
			hideTitle = true
		}

		cv.y = rightY
		if !hideTitle {
			cv.sectionTitle(rx, "Education")
		}
		for _, item := range items {
			im := toJMap(item)
			degree := jstr(im, "degree")
			school := jstr(im, "school")
			startD := jstr(im, "start_date")
			endD := jstr(im, "end_date")
			dates := startD
			if endD != "" {
				dates += " – " + endD
			}

			if degree != "" || school != "" {
				cv.text(rx, cv.y, cBlack, "THSarabun", 10, degree)
				cv.y += 13
				cv.text(rx, cv.y, cAccent, "THSarabun", 9, school)
				cv.y += 12
				if dates != "" {
					cv.text(rx, cv.y, cGray400, "THSarabun", 9, dates)
					cv.y += 12
				}
				gpa := jstr(im, "gpa")
				if gpa != "" {
					cv.text(rx, cv.y, cGray400, "THSarabun", 9, "GPA: "+gpa)
					cv.y += 12
				}
				cv.y += 6
			}
		}
		rightY = cv.y
	}

	// ── LEFT: Experience ───────────────────────────────────────────────────────
	cv.y = bodyTopY
	for i, sec := range sections {
		if !sec.IsVisible || sec.Type != "experience" {
			continue
		}
		items := jslice(data[i], "items")
		if len(items) == 0 {
			continue
		}

		hideTitle := false
		if b, ok := data[i]["hide_title"].(bool); ok && b {
			hideTitle = true
		}

		if !hideTitle {
			cv.sectionTitle(cv.ml, "Experience")
		}
		for _, item := range items {
			im := toJMap(item)
			pos := firstOf(jstr(im, "position"), jstr(im, "title"))
			company := jstr(im, "company")
			startD := jstr(im, "start_date")
			endD := jstr(im, "end_date")
			if jbool(im, "is_current") {
				endD = "Present"
			}
			dates := startD
			if endD != "" {
				dates += " – " + endD
			}
			desc := jstr(im, "description")

			if cv.needsNewPage(72) {
				cv.drawFooter()
				cv.addPage()
			}

			// Position + dates on same row
			cv.text(cv.ml, cv.y, cBlack, "THSarabun", 11, pos)
			if dates != "" {
				cv.textR(cv.ml+lw, cv.y, cGray400, "THSarabun", 9, dates)
			}
			cv.y += 14
			// Company
			if company != "" {
				cv.text(cv.ml, cv.y, cAccent, "THSarabun", 9, company)
				cv.y += 13
			}
			// Description
			if desc != "" {
				cv.y = cv.wrapText(cv.ml, cv.y, cGray600, "THSarabun", 9, desc, lw, 12, 3)
			}
			cv.y += 12
		}
	}

	// ── LEFT: Projects ─────────────────────────────────────────────────────────
	for i, sec := range sections {
		if !sec.IsVisible || sec.Type != "projects" {
			continue
		}
		items := jslice(data[i], "items")
		if len(items) == 0 {
			continue
		}

		hideTitle := false
		if b, ok := data[i]["hide_title"].(bool); ok && b {
			hideTitle = true
		}

		if !hideTitle {
			cv.sectionTitle(cv.ml, "Projects")
		}
		for _, item := range items {
			im := toJMap(item)
			title := jstr(im, "title")
			desc := jstr(im, "description")
			tags := jstrslice(im, "tags")

			if cv.needsNewPage(60) {
				cv.drawFooter()
				cv.addPage()
			}

			cv.text(cv.ml, cv.y, cBlack, "THSarabun", 11, title)
			cv.y += 13

			// Tags
			cv.y += 3
			if len(tags) > 0 {
				tx := cv.ml
				for _, tag := range tags {
					tx = cv.tagPill(tx, cv.y, tag)
				}
				cv.y += 18
			}
			// Description
			if desc != "" {
				cv.y = cv.wrapText(cv.ml, cv.y, cGray600, "THSarabun", 9, desc, lw, 12, 2)
			}
			cv.y += 10
		}
	}

	// ── LEFT: Custom Text ──────────────────────────────────────────────────────
	for i, sec := range sections {
		if !sec.IsVisible || sec.Type != "custom_text" {
			continue
		}
		content := jstr(data[i], "content")
		if content == "" {
			continue
		}

		if cv.needsNewPage(60) {
			cv.drawFooter()
			cv.addPage()
		}
		hideTitle := false
		if b, ok := data[i]["hide_title"].(bool); ok && b {
			hideTitle = true
		}

		if !hideTitle {
			cv.sectionTitle(cv.ml, "Note")
		}
		cv.y = cv.wrapText(cv.ml, cv.y, cGray600, "THSarabun", 10, content, lw, 13, 6)
		cv.y += 10
	}

	// ── Footer ────────────────────────────────────────────────────────────────
	cv.drawFooter()

	// 5. Output ────────────────────────────────────────────────────────────────
	buf := cv.p.GetBytesPdf()
	if len(buf) == 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate PDF")
	}
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", `attachment; filename="resume.pdf"`)
	return c.SendStream(bytes.NewReader(buf))
}

func (cv *cv) drawFooter() {
	cv.hline(cv.ml, cv.h-40, cv.w-cv.mr, cGray200, 0.5)
	cv.text(cv.ml, cv.h-28, cGray400, "THSarabun", 8, "Generated by PloyKong • ploykong.io")
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
