package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

// ─── User ─────────────────────────────────────────────────────────────────────

type User struct {
	ID         string         `json:"id"`
	Email      string         `json:"email"`
	GoogleID   sql.NullString `json:"-"`
	Password   string         `json:"-"` // never expose
	Name       string         `json:"name"`
	AvatarURL  sql.NullString `json:"avatar_url"`
	Plan       string         `json:"plan"`
	IsVerified bool           `json:"is_verified"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type UserResponse struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	AvatarURL  *string   `json:"avatar_url"`
	Plan       string    `json:"plan"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
}

func (u *User) ToResponse() *UserResponse {
	r := &UserResponse{
		ID:         u.ID,
		Email:      u.Email,
		Name:       u.Name,
		Plan:       u.Plan,
		IsVerified: u.IsVerified,
		CreatedAt:  u.CreatedAt,
	}
	if u.AvatarURL.Valid {
		r.AvatarURL = &u.AvatarURL.String
	}
	return r
}

// ─── Portfolio ────────────────────────────────────────────────────────────────

type ThemeConfig struct {
	Mode           string `json:"mode"` // dark | light
	PrimaryColor   string `json:"primary_color"`
	SecondaryColor string `json:"secondary_color,omitempty"`
	Palette        string `json:"palette,omitempty"`
	Template       string `json:"template,omitempty"`
	BgColor        string `json:"bg_color,omitempty"`
	BorderColor    string `json:"border_color,omitempty"`
	Font           string `json:"font"`
	Layout         string `json:"layout"` // centered | left | split
	ShowDivider    bool   `json:"show_divider"`
}

type Portfolio struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	Slug         string         `json:"slug"`
	Title        string         `json:"title"`
	Description  sql.NullString `json:"description"`
	Theme        *ThemeConfig   `json:"theme"`
	SEOTitle     sql.NullString `json:"seo_title"`
	SEODesc      sql.NullString `json:"seo_desc"`
	OGImageURL   sql.NullString `json:"og_image_url"`
	IsPublished  bool           `json:"is_published"`
	PasswordHash sql.NullString `json:"-"`
	HasPassword  bool           `json:"has_password"`
	ExpiresAt    *time.Time     `json:"expires_at"`
	ViewCount    int            `json:"view_count"`
	Sections     []Section      `json:"sections,omitempty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (p *Portfolio) ScanTheme(raw []byte) error {
	if raw == nil {
		p.Theme = &ThemeConfig{
			Mode:         "dark",
			PrimaryColor: "#4F46E5",
			Font:         "Syne",
			Layout:       "centered",
		}
		return nil
	}
	p.Theme = &ThemeConfig{}
	return json.Unmarshal(raw, p.Theme)
}

// ─── Section ──────────────────────────────────────────────────────────────────

type SectionType string

const (
	SectionHero         SectionType = "hero"
	SectionExperience   SectionType = "experience"
	SectionSkills       SectionType = "skills"
	SectionProjects     SectionType = "projects"
	SectionEducation    SectionType = "education"
	SectionContact      SectionType = "contact"
	SectionAIChat       SectionType = "ai_chat"
	SectionCustomText   SectionType = "custom_text"
	SectionStats        SectionType = "stats"
	SectionSocial       SectionType = "social"
	SectionCertificates SectionType = "certificates"
)

type Section struct {
	ID          string          `json:"id"`
	PortfolioID string          `json:"portfolio_id"`
	Type        SectionType     `json:"type"`
	Position    int             `json:"position"`
	Data        json.RawMessage `json:"data"` // flexible JSON
	IsVisible   bool            `json:"is_visible"`
	HideTitle   bool            `json:"hide_title"`
	HideDivider bool            `json:"hide_divider"`
	ColumnSpan  string          `json:"column_span"` // "full" | "half"
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// ─── Hero Section Data ────────────────────────────────────────────────────────

type HeroData struct {
	Name       string `json:"name"`
	Role       string `json:"role"`
	Tagline    string `json:"tagline"`
	AvatarURL  string `json:"avatar_url"`
	ShowHireMe bool   `json:"show_hire_me"`
	HireMeLink string `json:"hire_me_link"`
}

// ─── Experience Section Data ──────────────────────────────────────────────────

type ExperienceItem struct {
	Company     string   `json:"company"`
	Position    string   `json:"position"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	IsCurrent   bool     `json:"is_current"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	ImageURLs   []string `json:"image_urls,omitempty"`
	Skills      []string `json:"skills,omitempty"`
}

type ExperienceData struct {
	Items []ExperienceItem `json:"items"`
}

// ─── Skills Section Data ──────────────────────────────────────────────────────

type SkillItem struct {
	Name     string `json:"name"`
	Level    int    `json:"level"` // 0-100
	Category string `json:"category"`
}

type SkillsData struct {
	HidePercentage bool        `json:"hide_percentage"`
	Items          []SkillItem `json:"items"`
}

// ─── Projects Section Data ────────────────────────────────────────────────────

type ProjectItem struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	LiveURL     string   `json:"live_url"`
	GithubURL   string   `json:"github_url"`
	Tags        []string `json:"tags"`
}

type ProjectsData struct {
	Items []ProjectItem `json:"items"`
}

// ─── Analytics ────────────────────────────────────────────────────────────────

type AnalyticsSummary struct {
	TotalViews   int           `json:"total_views"`
	TodayViews   int           `json:"today_views"`
	WeekViews    int           `json:"week_views"`
	MonthViews   int           `json:"month_views"`
	PDFDownloads int           `json:"pdf_downloads"`
	HireClicks   int           `json:"hire_clicks"`
	AIChatCount  int           `json:"ai_chat_count"`
	CountryStats []CountryStat `json:"country_stats"`
	DailyViews   []DailyView   `json:"daily_views"`
}

type CountryStat struct {
	CountryCode string `json:"country_code"`
	Count       int    `json:"count"`
}

type DailyView struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type SourceStat struct {
	Source string `json:"source"`
	Count  int    `json:"count"`
}

// ─── Auth ─────────────────────────────────────────────────────────────────────

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // seconds
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Plan   string `json:"plan"`
}
