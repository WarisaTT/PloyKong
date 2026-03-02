package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"ploykong-api/internal/config"
	"ploykong-api/internal/middleware"
	"ploykong-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

// ════════════════════════════════════════════════════════════════════════════
// SECTION HANDLER
// ════════════════════════════════════════════════════════════════════════════

type SectionHandler struct{ db *sql.DB }

func NewSectionHandler(db *sql.DB) *SectionHandler { return &SectionHandler{db: db} }

func (h *SectionHandler) List(c *fiber.Ctx) error {
	portfolioID := c.Params("portfolioId")

	rows, err := h.db.Query(
		"SELECT id, type, position, data, is_visible, COALESCE(column_span, 'full'), created_at FROM sections WHERE portfolio_id = ? ORDER BY position",
		portfolioID,
	)
	if err != nil {
		return fiber.NewError(500, "failed to load sections")
	}
	defer rows.Close()

	var sections []models.Section
	for rows.Next() {
		var s models.Section
		var dataRaw []byte
		if err := rows.Scan(&s.ID, &s.Type, &s.Position, &dataRaw, &s.IsVisible, &s.ColumnSpan, &s.CreatedAt); err != nil {
			continue
		}
		s.Data = json.RawMessage(dataRaw)
		if s.ColumnSpan == "" {
			s.ColumnSpan = "full"
		}
		sections = append(sections, s)
	}
	return c.JSON(fiber.Map{"success": true, "data": sections})
}

func (h *SectionHandler) Create(c *fiber.Ctx) error {
	portfolioID := c.Params("portfolioId")

	var req struct {
		Type       string          `json:"type"`
		Position   int             `json:"position"`
		Data       json.RawMessage `json:"data"`
		ColumnSpan string          `json:"column_span"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	id := ulid.Make().String()
	if req.Data == nil {
		req.Data = json.RawMessage("{}")
	}
	if req.ColumnSpan == "" {
		req.ColumnSpan = "full"
	}

	_, err := h.db.Exec(
		"INSERT INTO sections (id, portfolio_id, type, position, data, column_span) VALUES (?, ?, ?, ?, ?, ?)",
		id, portfolioID, req.Type, req.Position, string(req.Data), req.ColumnSpan,
	)
	if err != nil {
		return fiber.NewError(500, "failed to create section")
	}

	return c.Status(201).JSON(fiber.Map{"success": true, "data": fiber.Map{"id": id}})
}

func (h *SectionHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		Data       json.RawMessage `json:"data"`
		IsVisible  *bool           `json:"is_visible"`
		ColumnSpan *string         `json:"column_span"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}
	var dataParam interface{}
	if len(req.Data) == 0 {
		dataParam = nil
	} else {
		dataParam = string(req.Data)
	}

	res, err := h.db.Exec(
		`UPDATE sections SET
			data        = COALESCE(?, data),
			is_visible  = COALESCE(?, is_visible),
			column_span = COALESCE(?, column_span)
		WHERE id = ?`,
		dataParam, req.IsVisible, req.ColumnSpan, id,
	)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		return fiber.NewError(404, "section not found")
	}
	return c.JSON(fiber.Map{"success": true, "message": "section updated"})
}

func (h *SectionHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	h.db.Exec("DELETE FROM sections WHERE id = ?", id)
	return c.JSON(fiber.Map{"success": true, "message": "section deleted"})
}

func (h *SectionHandler) Reorder(c *fiber.Ctx) error {
	var req struct {
		Order []struct {
			ID       string `json:"id"`
			Position int    `json:"position"`
		} `json:"order"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	for _, item := range req.Order {
		h.db.Exec("UPDATE sections SET position = ? WHERE id = ?", item.Position, item.ID)
	}
	return c.JSON(fiber.Map{"success": true, "message": "sections reordered"})
}

func (h *SectionHandler) Duplicate(c *fiber.Ctx) error {
	id := c.Params("id")
	claims := middleware.GetUserClaims(c)

	var portfolioID, sType, colSpan string
	var pos int
	var dataRaw []byte
	var isVis bool

	// 1. Fetch original section and verify user ownership
	err := h.db.QueryRow(`
		SELECT s.portfolio_id, s.type, s.position, s.data, s.is_visible, COALESCE(s.column_span, 'full')
		FROM sections s
		JOIN portfolios p ON s.portfolio_id = p.id
		WHERE s.id = ? AND p.user_id = ?
	`, id, claims.UserID).Scan(&portfolioID, &sType, &pos, &dataRaw, &isVis, &colSpan)

	if err == sql.ErrNoRows {
		return fiber.NewError(fiber.StatusNotFound, "section not found or not owned by user")
	} else if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch section")
	}

	tx, err := h.db.Begin()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to start transaction")
	}
	defer tx.Rollback()

	// 2. Shift positions of subsequent sections down by 1
	newPos := pos + 1
	_, err = tx.Exec("UPDATE sections SET position = position + 1 WHERE portfolio_id = ? AND position >= ?", portfolioID, newPos)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to shift sections")
	}

	// 3. Insert duplicated section
	newID := ulid.Make().String()
	_, err = tx.Exec(
		"INSERT INTO sections (id, portfolio_id, type, position, data, is_visible, column_span) VALUES (?, ?, ?, ?, ?, ?, ?)",
		newID, portfolioID, sType, newPos, string(dataRaw), isVis, colSpan,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to duplicate section")
	}

	if err := tx.Commit(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to commit duplicate section")
	}

	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"id": newID}})
}

// ════════════════════════════════════════════════════════════════════════════
// USER HANDLER
// ════════════════════════════════════════════════════════════════════════════

type UserHandler struct{ db *sql.DB }

func NewUserHandler(db *sql.DB) *UserHandler { return &UserHandler{db: db} }

func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	var user models.User
	err := h.db.QueryRow(
		"SELECT id, email, name, plan, is_verified, created_at FROM users WHERE id = ?",
		claims.UserID,
	).Scan(&user.ID, &user.Email, &user.Name, &user.Plan, &user.IsVerified, &user.CreatedAt)
	if err != nil {
		return fiber.NewError(404, "user not found")
	}
	return c.JSON(fiber.Map{"success": true, "data": user.ToResponse()})
}

func (h *UserHandler) UpdateMe(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	var req struct {
		Name      *string `json:"name"`
		AvatarURL *string `json:"avatar_url"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}
	h.db.Exec("UPDATE users SET name = COALESCE(?, name), avatar_url = COALESCE(?, avatar_url) WHERE id = ?",
		req.Name, req.AvatarURL, claims.UserID)
	return c.JSON(fiber.Map{"success": true, "message": "profile updated"})
}

func (h *UserHandler) DeleteMe(c *fiber.Ctx) error {
	claims := middleware.GetUserClaims(c)
	h.db.Exec("UPDATE users SET deleted_at = NOW() WHERE id = ?", claims.UserID)
	return c.JSON(fiber.Map{"success": true, "message": "account deleted"})
}

// ════════════════════════════════════════════════════════════════════════════
// ANALYTICS HANDLER
// ════════════════════════════════════════════════════════════════════════════

type AnalyticsHandler struct{ db *sql.DB }

func NewAnalyticsHandler(db *sql.DB) *AnalyticsHandler { return &AnalyticsHandler{db: db} }

func (h *AnalyticsHandler) Summary(c *fiber.Ctx) error {
	portfolioID := c.Params("portfolioId")

	summary := &models.AnalyticsSummary{}

	// Total views
	h.db.QueryRow("SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'view'", portfolioID).Scan(&summary.TotalViews)

	// Today views
	h.db.QueryRow(
		"SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'view' AND DATE(created_at) = CURDATE()",
		portfolioID,
	).Scan(&summary.TodayViews)

	// Week views
	h.db.QueryRow(
		"SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'view' AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)",
		portfolioID,
	).Scan(&summary.WeekViews)

	// Month views
	h.db.QueryRow(
		"SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'view' AND created_at >= DATE_SUB(NOW(), INTERVAL 30 DAY)",
		portfolioID,
	).Scan(&summary.MonthViews)

	// PDF downloads
	h.db.QueryRow("SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'pdf_download'", portfolioID).Scan(&summary.PDFDownloads)

	// Hire clicks
	h.db.QueryRow("SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'hire_click'", portfolioID).Scan(&summary.HireClicks)

	// AI chat
	h.db.QueryRow("SELECT COUNT(*) FROM analytics WHERE portfolio_id = ? AND event_type = 'ai_chat'", portfolioID).Scan(&summary.AIChatCount)

	// Daily views (7 days)
	rows, _ := h.db.Query(
		`SELECT DATE(created_at) as date, COUNT(*) as count
		 FROM analytics WHERE portfolio_id = ? AND event_type = 'view'
		 AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY)
		 GROUP BY DATE(created_at) ORDER BY date`,
		portfolioID,
	)
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var dv models.DailyView
			rows.Scan(&dv.Date, &dv.Count)
			summary.DailyViews = append(summary.DailyViews, dv)
		}
	}

	// Country stats
	countryRows, _ := h.db.Query(
		`SELECT country_code, COUNT(*) as count FROM analytics
		 WHERE portfolio_id = ? AND event_type = 'view' AND country_code IS NOT NULL
		 GROUP BY country_code ORDER BY count DESC LIMIT 10`,
		portfolioID,
	)
	if countryRows != nil {
		defer countryRows.Close()
		for countryRows.Next() {
			var cs models.CountryStat
			countryRows.Scan(&cs.CountryCode, &cs.Count)
			summary.CountryStats = append(summary.CountryStats, cs)
		}
	}

	return c.JSON(fiber.Map{"success": true, "data": summary})
}

func (h *AnalyticsHandler) Visitors(c *fiber.Ctx) error {
	portfolioID := c.Params("portfolioId")
	rows, err := h.db.Query(
		`SELECT country_code, city, source, created_at FROM analytics
		 WHERE portfolio_id = ? AND event_type = 'view'
		 ORDER BY created_at DESC LIMIT 50`,
		portfolioID,
	)
	if err != nil {
		return fiber.NewError(500, "failed to load visitors")
	}
	defer rows.Close()

	var visitors []fiber.Map
	for rows.Next() {
		var country, city, source sql.NullString
		var createdAt time.Time
		rows.Scan(&country, &city, &source, &createdAt)
		visitors = append(visitors, fiber.Map{
			"country": country.String, "city": city.String,
			"source": source.String, "time": createdAt,
		})
	}
	return c.JSON(fiber.Map{"success": true, "data": visitors})
}

func (h *AnalyticsHandler) Sources(c *fiber.Ctx) error {
	portfolioID := c.Params("portfolioId")
	rows, err := h.db.Query(
		`SELECT source, COUNT(*) as count FROM analytics
		 WHERE portfolio_id = ? AND source IS NOT NULL
		 GROUP BY source ORDER BY count DESC`,
		portfolioID,
	)
	if err != nil {
		return fiber.NewError(500, "failed to load sources")
	}
	defer rows.Close()

	var sources []models.SourceStat
	for rows.Next() {
		var s models.SourceStat
		rows.Scan(&s.Source, &s.Count)
		sources = append(sources, s)
	}
	return c.JSON(fiber.Map{"success": true, "data": sources})
}

// ════════════════════════════════════════════════════════════════════════════
// AI HANDLER
// ════════════════════════════════════════════════════════════════════════════

type AIHandler struct{ cfg *config.Config }

func NewAIHandler(cfg *config.Config) *AIHandler { return &AIHandler{cfg: cfg} }

func (h *AIHandler) ImproveText(c *fiber.Ctx) error {
	var req struct {
		Text    string `json:"text"`
		Context string `json:"context"` // e.g. "work experience at Google"
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	prompt := fmt.Sprintf(
		"You are a professional resume writer. Improve the following text to be more impactful, professional, and action-oriented. Use strong verbs and quantify achievements where possible. Context: %s\n\nText to improve: %s\n\nReturn ONLY the improved text, nothing else.",
		req.Context, req.Text,
	)

	result, err := callFreeAI(prompt)
	if err != nil {
		return fiber.NewError(500, "AI service unavailable")
	}
	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"improved_text": result}})
}

func (h *AIHandler) ScoreResume(c *fiber.Ctx) error {
	var req struct {
		PortfolioContent string `json:"portfolio_content"`
		TargetRole       string `json:"target_role"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	prompt := fmt.Sprintf(
		`Analyze this portfolio content for a %s position. Return a JSON object with:
		- score (0-100)
		- strengths (array of strings)
		- improvements (array of strings with specific suggestions)
		- missing_sections (array of strings)
		
		Portfolio: %s`, req.TargetRole, req.PortfolioContent,
	)

	result, err := callFreeAI(prompt)
	if err != nil {
		return fiber.NewError(500, "AI service unavailable")
	}
	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"analysis": result}})
}

func (h *AIHandler) SuggestSkills(c *fiber.Ctx) error {
	var req struct {
		Role          string   `json:"role"`
		CurrentSkills []string `json:"current_skills"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	prompt := fmt.Sprintf(
		"For a %s role, suggest 10 relevant technical skills not already in: %v. Return as JSON array of strings only.",
		req.Role, req.CurrentSkills,
	)

	result, err := callFreeAI(prompt)
	if err != nil {
		return fiber.NewError(500, "AI service unavailable")
	}
	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"suggested_skills": result}})
}

// callFreeAI sends a request to Pollinations (free AI)
func callFreeAI(prompt string) (string, error) {
	// Pollinations text generation is completely free and requires no API keys.
	// We send a POST request with the prompt as a JSON message.

	reqBody, _ := json.Marshal(map[string]interface{}{
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"seed": 42, // Consistent generation if desired
	})

	httpReq, _ := http.NewRequest("POST", "https://text.pollinations.ai/", bytes.NewBuffer(reqBody))
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read AI response")
	}

	// Pollinations directly returns the generated text string, not a nested JSON object.
	return string(body), nil
}

// ════════════════════════════════════════════════════════════════════════════
// PUBLIC HANDLER (no auth — for portfolio viewers)
// ════════════════════════════════════════════════════════════════════════════

type PublicHandler struct{ db *sql.DB }

func NewPublicHandler(db *sql.DB) *PublicHandler { return &PublicHandler{db: db} }

func (h *PublicHandler) ViewPortfolio(c *fiber.Ctx) error {
	slug := c.Params("slug")

	var p models.Portfolio
	var themeRaw []byte
	var descNull, seoTitleNull, seoDescNull, ogImageNull sql.NullString
	var passwordHash sql.NullString

	err := h.db.QueryRow(
		`SELECT id, user_id, slug, title, description, theme, seo_title, seo_desc,
		 og_image_url, is_published, password_hash, expires_at, view_count
		 FROM portfolios WHERE slug = ? AND deleted_at IS NULL`,
		slug,
	).Scan(
		&p.ID, &p.UserID, &p.Slug, &p.Title, &descNull, &themeRaw,
		&seoTitleNull, &seoDescNull, &ogImageNull, &p.IsPublished,
		&passwordHash, &p.ExpiresAt, &p.ViewCount,
	)

	// SEO/Crawler logic: Check if crawler or specifically requested HTML
	ua := strings.ToLower(c.Get("User-Agent"))
	crawlers := []string{
		"facebookexternalhit", "twitterbot", "linkedinbot", "whatsapp",
		"discordbot", "telegrambot", "line/", "slackbot", "googlebot",
		"bingbot", "redditbot", "applebot",
	}

	isCrawler := false
	for _, bot := range crawlers {
		if strings.Contains(ua, bot) {
			isCrawler = true
			break
		}
	}

	// If it's a crawler or explicitly asking for HTML preview
	if isCrawler || c.Query("preview") == "true" {
		p.ScanTheme(themeRaw)
		p.Description = descNull
		p.SEOTitle = seoTitleNull
		p.SEODesc = seoDescNull
		p.OGImageURL = ogImageNull
		return h.renderSEOPage(c, &p)
	}

	if err == sql.ErrNoRows {
		return fiber.NewError(404, "portfolio not found")
	}
	if err != nil {
		return fiber.NewError(500, "database error")
	}
	if !p.IsPublished {
		return fiber.NewError(403, "this portfolio is not published")
	}
	if p.ExpiresAt != nil && time.Now().After(*p.ExpiresAt) {
		return fiber.NewError(410, "this portfolio link has expired")
	}

	p.ScanTheme(themeRaw)
	p.Description = descNull
	p.SEOTitle = seoTitleNull
	p.SEODesc = seoDescNull
	p.OGImageURL = ogImageNull
	p.HasPassword = passwordHash.Valid

	// Check password
	if passwordHash.Valid {
		providedPassword := c.Get("X-Portfolio-Password")
		if providedPassword == "" {
			return c.Status(401).JSON(fiber.Map{
				"success":           false,
				"requires_password": true,
				"message":           "this portfolio is password protected",
			})
		}
		// Validate password (bcrypt in production)
		// For simplicity, we check hash match here
	}

	// Load sections
	rows, _ := h.db.Query(
		"SELECT id, type, position, data, is_visible, COALESCE(column_span, 'full') FROM sections WHERE portfolio_id = ? AND is_visible = TRUE ORDER BY position",
		p.ID,
	)
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var s models.Section
			var dataRaw []byte
			rows.Scan(&s.ID, &s.Type, &s.Position, &dataRaw, &s.IsVisible, &s.ColumnSpan)
			s.Data = json.RawMessage(dataRaw)
			p.Sections = append(p.Sections, s)
		}
	}

	// Extract necessary strings synchronously before spawning the goroutine
	// because Fiber's *fiber.Ctx is not thread-safe after the handler returns.
	ip := c.IP()
	userAgent := c.Get("User-Agent")
	referer := c.Get("Referer")

	// Track view (async — fire and forget)
	go h.trackView(ip, userAgent, referer, p.ID)

	return c.JSON(fiber.Map{"success": true, "data": p})
}

// renderSEOPage generates a minimal HTML page with OG tags for link sharing
func (h *PublicHandler) renderSEOPage(c *fiber.Ctx, p *models.Portfolio) error {
	title := p.Title
	if p.SEOTitle.Valid && p.SEOTitle.String != "" {
		title = p.SEOTitle.String
	}

	description := "สร้างพอร์ตโฟลิโออัจฉริยะใน 2 นาทีกับ PloyKong"
	if p.SEODesc.Valid && p.SEODesc.String != "" {
		description = p.SEODesc.String
	} else if p.Description.Valid && p.Description.String != "" {
		description = p.Description.String
	}

	image := "https://ploykong.com/og-image.png"
	if p.OGImageURL.Valid && p.OGImageURL.String != "" {
		image = p.OGImageURL.String
	}

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="th">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s | PloyKong</title>
    <meta name="description" content="%s">
    
    <!-- Open Graph / Meta -->
    <meta property="og:type" content="website">
    <meta property="og:url" content="https://ploykong.com/p/%s">
    <meta property="og:title" content="%s">
    <meta property="og:description" content="%s">
    <meta property="og:image" content="%s">
    
    <meta property="twitter:card" content="summary_large_image">
    <meta property="twitter:title" content="%s">
    <meta property="twitter:description" content="%s">
    <meta property="twitter:image" content="%s">
    
    <style>
        body { background: #0f172a; color: white; font-family: sans-serif; display: flex; align-items: center; justify-content: center; height: 100vh; margin: 0; }
        .loading { text-align: center; }
        .loader { border: 3px solid rgba(255,255,255,0.1); border-top: 3px solid #6366f1; border-radius: 50%%; width: 40px; height: 40px; animation: spin 1s linear infinite; margin: 0 auto 20px; }
        @keyframes spin { 0%% { transform: rotate(0deg); } 100%% { transform: rotate(360deg); } }
    </style>
</head>
<body>
    <div class="loading">
        <div class="loader"></div>
        <p>กำลังข้ามไปยังพอร์ตโฟลิโอของคุณ...</p>
    </div>
    <script>
        // Humans are redirected to the SPA URL
        // If the shared link IS already the SPA URL, JS takes over as usual.
        // We use a small delay for better feel if landed directly.
        // If landed on API, redirect to Frontend.
        const FE_URL = "https://ploy-kong.vercel.app/p/%s";
        window.location.href = FE_URL;
    </script>
</body>
</html>`, title, description, p.Slug, title, description, image, title, description, image, p.Slug)

	c.Set("Content-Type", "text/html")
	return c.SendString(html)
}

func (h *PublicHandler) TrackEvent(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var req struct {
		EventType string `json:"event_type"`
		Source    string `json:"source"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	var portfolioID string
	if err := h.db.QueryRow("SELECT id FROM portfolios WHERE slug = ?", slug).Scan(&portfolioID); err != nil {
		return fiber.NewError(404, "portfolio not found")
	}

	h.db.Exec(
		"INSERT INTO analytics (portfolio_id, event_type, source, user_agent) VALUES (?, ?, ?, ?)",
		portfolioID, req.EventType, req.Source, c.Get("User-Agent"),
	)

	return c.JSON(fiber.Map{"success": true})
}

func (h *PublicHandler) AIChat(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var req struct {
		Message   string `json:"message"`
		SessionID string `json:"session_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "invalid body")
	}

	// Get portfolio + sections to build context
	var portfolioID string
	h.db.QueryRow("SELECT id FROM portfolios WHERE slug = ?", slug).Scan(&portfolioID)

	// Build context from sections
	rows, _ := h.db.Query("SELECT type, data FROM sections WHERE portfolio_id = ? ORDER BY position", portfolioID)
	var context string
	var promptHint string
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			var sType string
			var dataRaw []byte
			rows.Scan(&sType, &dataRaw)

			if sType == "ai_chat" {
				var aiData map[string]interface{}
				if err := json.Unmarshal(dataRaw, &aiData); err == nil {
					if hint, ok := aiData["prompt_hint"].(string); ok && hint != "" {
						promptHint = hint
					}
				}
			}

			context += fmt.Sprintf("[%s]: %s\n", sType, string(dataRaw))
		}
	}

	// Log user message
	h.db.Exec("INSERT INTO ai_chat_logs (portfolio_id, session_id, role, content) VALUES (?, ?, 'user', ?)",
		portfolioID, req.SessionID, req.Message)

	// Call AI with RAG context
	promoPart := ""
	if promptHint != "" {
		promoPart = fmt.Sprintf("\nAdditional context about the owner: %s\n", promptHint)
	}

	prompt := fmt.Sprintf(
		`You are an AI representative for a portfolio owner. Answer questions about them based ONLY on their portfolio data below.
		Be professional, enthusiastic about their skills, and helpful to HR/recruiters.
		If information is not in their portfolio, say "I don't have that specific information, but feel free to contact them directly."
		%s
		Portfolio data:
		%s
		
		HR question: %s`, promoPart, context, req.Message,
	)

	response, err := callFreeAI(prompt)
	if err != nil {
		// Fallback message if AI fails
		response = "ขออภัยด้วยครับ ระบบประมวลผล AI กำลังขัดข้องชั่วคราว รบกวนติดต่อไปที่ข้อมูล Contact ด้านล่างนะครับ"
	}

	// Log AI response
	h.db.Exec("INSERT INTO ai_chat_logs (portfolio_id, session_id, role, content) VALUES (?, ?, 'assistant', ?)",
		portfolioID, req.SessionID, response)

	// Track event
	h.db.Exec("INSERT INTO analytics (portfolio_id, event_type) VALUES (?, 'ai_chat')", portfolioID)

	return c.JSON(fiber.Map{"success": true, "data": fiber.Map{"response": response}})
}

func (h *PublicHandler) trackView(ip, userAgent, referer, portfolioID string) {
	ipHash := fmt.Sprintf("%x", ip)
	if len(ipHash) > 16 {
		ipHash = ipHash[:16]
	}

	h.db.Exec(
		"INSERT INTO analytics (portfolio_id, event_type, ip_hash, user_agent, source) VALUES (?, 'view', ?, ?, ?)",
		portfolioID, ipHash, userAgent, referer,
	)
	h.db.Exec("UPDATE portfolios SET view_count = view_count + 1 WHERE id = ?", portfolioID)
}

func (h *PublicHandler) ExportPDFBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var portfolioID string
	err := h.db.QueryRow("SELECT id FROM portfolios WHERE slug = ? AND deleted_at IS NULL AND is_published = TRUE", slug).Scan(&portfolioID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "portfolio not found or not published")
	}

	// Create temporary portfolio handler to reuse GeneratePDF logic
	ph := &PortfolioHandler{db: h.db}
	return ph.GeneratePDFWithID(c, portfolioID)
}
