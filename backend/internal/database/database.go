package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"ploykong-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes MySQL connection with connection pooling
func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// Connection pool settings (high performance)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return db, nil
}

// Migrate runs schema creation (idempotent)
func Migrate(db *sql.DB) error {
	log.Println("🔄 Running database migrations...")

	queries := []string{
		// ─── Users ──────────────────────────────────────────────────────────
		`CREATE TABLE IF NOT EXISTS users (
			id          VARCHAR(26)  NOT NULL PRIMARY KEY,
			email       VARCHAR(255) NOT NULL UNIQUE,
			password    VARCHAR(255) NOT NULL,
			name        VARCHAR(255) NOT NULL DEFAULT '',
			avatar_url  VARCHAR(500),
			plan        ENUM('free','pro','team') NOT NULL DEFAULT 'free',
			is_verified BOOLEAN NOT NULL DEFAULT FALSE,
			created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at  TIMESTAMP NULL,
			INDEX idx_email (email),
			INDEX idx_plan (plan)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,

		// ─── Portfolios ──────────────────────────────────────────────────────
		`CREATE TABLE IF NOT EXISTS portfolios (
			id            VARCHAR(26)  NOT NULL PRIMARY KEY,
			user_id       VARCHAR(26)  NOT NULL,
			slug          VARCHAR(100) NOT NULL UNIQUE,
			title         VARCHAR(255) NOT NULL DEFAULT 'My Portfolio',
			description   TEXT,
			theme         JSON,
			seo_title     VARCHAR(255),
			seo_desc      VARCHAR(500),
			og_image_url  VARCHAR(500),
			is_published  BOOLEAN NOT NULL DEFAULT FALSE,
			password_hash VARCHAR(255),
			expires_at    TIMESTAMP NULL,
			view_count    INT NOT NULL DEFAULT 0,
			created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at    TIMESTAMP NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			INDEX idx_user_id (user_id),
			INDEX idx_slug (slug),
			INDEX idx_published (is_published)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,

		// ─── Sections (dynamic blocks) ────────────────────────────────────
		`CREATE TABLE IF NOT EXISTS sections (
			id           VARCHAR(26) NOT NULL PRIMARY KEY,
			portfolio_id VARCHAR(26) NOT NULL,
			type         VARCHAR(50) NOT NULL,
			position     INT NOT NULL DEFAULT 0,
			data         JSON NOT NULL,
			is_visible   BOOLEAN NOT NULL DEFAULT TRUE,
			created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			FOREIGN KEY (portfolio_id) REFERENCES portfolios(id) ON DELETE CASCADE,
			INDEX idx_portfolio_pos (portfolio_id, position)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,

		// (Column span migration logic will run programmatically below to support older MySQL 8 versions)

		// ─── Analytics ────────────────────────────────────────────────────
		`CREATE TABLE IF NOT EXISTS analytics (
			id           BIGINT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
			portfolio_id VARCHAR(26)  NOT NULL,
			event_type   ENUM('view','pdf_download','hire_click','ai_chat','link_click') NOT NULL,
			ip_hash      VARCHAR(64),
			country_code CHAR(2),
			city         VARCHAR(100),
			source       VARCHAR(100),
			referrer     VARCHAR(500),
			user_agent   VARCHAR(500),
			created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			INDEX idx_portfolio_date (portfolio_id, created_at),
			INDEX idx_event_type (event_type),
			INDEX idx_country (country_code)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,

		// ─── Refresh Tokens ────────────────────────────────────────────────
		`CREATE TABLE IF NOT EXISTS refresh_tokens (
			id         VARCHAR(26)  NOT NULL PRIMARY KEY,
			user_id    VARCHAR(26)  NOT NULL,
			token_hash VARCHAR(255) NOT NULL UNIQUE,
			expires_at TIMESTAMP    NOT NULL,
			created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			INDEX idx_token_hash (token_hash),
			INDEX idx_user_expires (user_id, expires_at)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,

		// ─── AI Chat Logs (RAG) ────────────────────────────────────────────
		`CREATE TABLE IF NOT EXISTS ai_chat_logs (
			id           BIGINT      NOT NULL AUTO_INCREMENT PRIMARY KEY,
			portfolio_id VARCHAR(26) NOT NULL,
			session_id   VARCHAR(64),
			role         ENUM('user','assistant') NOT NULL,
			content      TEXT NOT NULL,
			created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			INDEX idx_portfolio_session (portfolio_id, session_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return fmt.Errorf("migration failed: %w\nQuery: %s", err, q[:50])
		}
	}

	// ─── Programmatic Migrations for older MySQL (< 8.0.29) ─────────
	var colCount int
	err := db.QueryRow(`
		SELECT COUNT(*) 
		FROM INFORMATION_SCHEMA.COLUMNS 
		WHERE TABLE_SCHEMA = DATABASE() 
		  AND TABLE_NAME = 'sections' 
		  AND COLUMN_NAME = 'column_span'
	`).Scan(&colCount)
	if err != nil {
		return fmt.Errorf("failed to check column_span existence: %w", err)
	}

	if colCount == 0 {
		if _, err := db.Exec(`ALTER TABLE sections ADD COLUMN column_span VARCHAR(10) NOT NULL DEFAULT 'full'`); err != nil {
			return fmt.Errorf("failed to add column_span: %w", err)
		}
	}

	log.Println("✅ Migrations completed successfully")
	return nil
}
