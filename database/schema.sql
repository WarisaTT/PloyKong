-- ============================================================
-- PloyKong Database Schema
-- MySQL 8.0+  |  utf8mb4_unicode_ci
-- ============================================================

CREATE DATABASE IF NOT EXISTS ploykong
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE ploykong;

-- ─────────────────────────────────────────────────────────────
-- TABLE: users
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS users (
  id          VARCHAR(26)  NOT NULL,
  email       VARCHAR(255) NOT NULL,
  password    VARCHAR(255) NOT NULL,
  name        VARCHAR(255) NOT NULL DEFAULT '',
  avatar_url  VARCHAR(500) DEFAULT NULL,
  bio         TEXT         DEFAULT NULL,
  plan        ENUM('free','pro','team') NOT NULL DEFAULT 'free',
  plan_expires_at TIMESTAMP NULL,
  is_verified BOOLEAN NOT NULL DEFAULT FALSE,
  last_login  TIMESTAMP NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at  TIMESTAMP NULL,

  PRIMARY KEY (id),
  UNIQUE KEY uq_email (email),
  INDEX idx_plan (plan),
  INDEX idx_deleted (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='User accounts';

-- ─────────────────────────────────────────────────────────────
-- TABLE: portfolios
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS portfolios (
  id            VARCHAR(26)  NOT NULL,
  user_id       VARCHAR(26)  NOT NULL,
  slug          VARCHAR(100) NOT NULL,
  title         VARCHAR(255) NOT NULL DEFAULT 'My Portfolio',
  description   TEXT         DEFAULT NULL,

  -- Theme stored as JSON (mode, primary_color, font, layout)
  theme         JSON         DEFAULT NULL,

  -- SEO
  seo_title     VARCHAR(255) DEFAULT NULL,
  seo_desc      VARCHAR(500) DEFAULT NULL,
  og_image_url  VARCHAR(500) DEFAULT NULL,

  -- Visibility & Security
  is_published  BOOLEAN NOT NULL DEFAULT FALSE,
  password_hash VARCHAR(255) DEFAULT NULL,   -- bcrypt if set
  expires_at    TIMESTAMP NULL,              -- self-destruct link

  -- Stats (denormalized for quick reads)
  view_count    INT UNSIGNED NOT NULL DEFAULT 0,

  created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at    TIMESTAMP NULL,

  PRIMARY KEY (id),
  UNIQUE KEY uq_slug (slug),
  INDEX idx_user_id (user_id),
  INDEX idx_published (is_published),
  INDEX idx_expires (expires_at),
  CONSTRAINT fk_portfolio_user FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='Portfolio / resume pages';

-- ─────────────────────────────────────────────────────────────
-- TABLE: sections  (dynamic blocks — core of the builder)
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS sections (
  id           VARCHAR(26) NOT NULL,
  portfolio_id VARCHAR(26) NOT NULL,

  -- Type: hero | experience | skills | projects | education |
  --        contact | ai_chat | custom_text | stats | social
  type         VARCHAR(50) NOT NULL,
  position     TINYINT UNSIGNED NOT NULL DEFAULT 0,

  -- All block content lives here as flexible JSON
  -- Schema varies per type (see docs/section-schemas.md)
  data         JSON NOT NULL DEFAULT (JSON_OBJECT()),

  is_visible   BOOLEAN NOT NULL DEFAULT TRUE,

  created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX idx_portfolio_pos (portfolio_id, position),
  INDEX idx_type (type),
  CONSTRAINT fk_section_portfolio FOREIGN KEY (portfolio_id)
    REFERENCES portfolios(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='Dynamic drag-and-drop sections';

-- ─────────────────────────────────────────────────────────────
-- TABLE: analytics
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS analytics (
  id           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  portfolio_id VARCHAR(26)     NOT NULL,

  event_type   ENUM(
    'view',
    'pdf_download',
    'hire_click',
    'ai_chat',
    'link_click',
    'contact_click'
  ) NOT NULL,

  -- Visitor info (privacy-safe: IP is hashed)
  ip_hash      VARCHAR(64)  DEFAULT NULL,
  country_code CHAR(2)      DEFAULT NULL,
  city         VARCHAR(100) DEFAULT NULL,
  latitude     DECIMAL(9,6) DEFAULT NULL,
  longitude    DECIMAL(9,6) DEFAULT NULL,

  -- Traffic source
  source       VARCHAR(100) DEFAULT NULL,  -- linkedin | facebook | direct | email | ...
  referrer     VARCHAR(500) DEFAULT NULL,
  utm_source   VARCHAR(100) DEFAULT NULL,
  utm_medium   VARCHAR(100) DEFAULT NULL,
  utm_campaign VARCHAR(100) DEFAULT NULL,

  user_agent   VARCHAR(500) DEFAULT NULL,

  created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX idx_portfolio_date (portfolio_id, created_at),
  INDEX idx_event_type (event_type),
  INDEX idx_country (country_code),
  INDEX idx_source (source)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='Visitor events and analytics';

-- ─────────────────────────────────────────────────────────────
-- TABLE: refresh_tokens
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS refresh_tokens (
  id         VARCHAR(26)  NOT NULL,
  user_id    VARCHAR(26)  NOT NULL,
  token_hash VARCHAR(255) NOT NULL,
  expires_at TIMESTAMP    NOT NULL,
  created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  UNIQUE KEY uq_token_hash (token_hash),
  INDEX idx_user_expires (user_id, expires_at),
  CONSTRAINT fk_rtoken_user FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='JWT refresh tokens (hashed)';

-- ─────────────────────────────────────────────────────────────
-- TABLE: ai_chat_logs
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS ai_chat_logs (
  id           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  portfolio_id VARCHAR(26)     NOT NULL,
  session_id   VARCHAR(64)     DEFAULT NULL,
  role         ENUM('user','assistant') NOT NULL,
  content      TEXT NOT NULL,
  tokens_used  SMALLINT UNSIGNED DEFAULT 0,
  created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX idx_portfolio_session (portfolio_id, session_id),
  INDEX idx_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='AI interview coach conversation logs';

-- ─────────────────────────────────────────────────────────────
-- TABLE: files  (uploaded images / PDFs)
-- ─────────────────────────────────────────────────────────────
CREATE TABLE IF NOT EXISTS files (
  id           VARCHAR(26)  NOT NULL,
  user_id      VARCHAR(26)  NOT NULL,
  original_name VARCHAR(255) NOT NULL,
  stored_name  VARCHAR(255) NOT NULL,
  file_path    VARCHAR(500) NOT NULL,
  mime_type    VARCHAR(100) NOT NULL,
  size_bytes   INT UNSIGNED NOT NULL DEFAULT 0,
  is_public    BOOLEAN NOT NULL DEFAULT TRUE,
  created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  INDEX idx_user_id (user_id),
  CONSTRAINT fk_file_user FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
  COMMENT='Uploaded media files';

-- ─────────────────────────────────────────────────────────────
-- VIEWS
-- ─────────────────────────────────────────────────────────────

-- View: portfolio_stats  (quick summary per portfolio)
CREATE OR REPLACE VIEW portfolio_stats AS
SELECT
  p.id             AS portfolio_id,
  p.slug,
  p.user_id,
  p.view_count,
  COUNT(DISTINCT CASE WHEN a.event_type = 'pdf_download' THEN a.id END) AS pdf_downloads,
  COUNT(DISTINCT CASE WHEN a.event_type = 'hire_click'   THEN a.id END) AS hire_clicks,
  COUNT(DISTINCT CASE WHEN a.event_type = 'ai_chat'      THEN a.id END) AS ai_chats,
  MAX(a.created_at) AS last_visit_at
FROM portfolios p
LEFT JOIN analytics a ON a.portfolio_id = p.id
WHERE p.deleted_at IS NULL
GROUP BY p.id;

-- ─────────────────────────────────────────────────────────────
-- STORED PROCEDURES
-- ─────────────────────────────────────────────────────────────

DELIMITER $$

-- Proc: cleanup_expired_tokens
-- Run via cron: CALL cleanup_expired_tokens()
CREATE PROCEDURE IF NOT EXISTS cleanup_expired_tokens()
BEGIN
  DELETE FROM refresh_tokens WHERE expires_at < NOW();
  DELETE FROM portfolios WHERE expires_at IS NOT NULL AND expires_at < NOW() AND deleted_at IS NULL;
END$$

-- Proc: get_portfolio_analytics_summary
CREATE PROCEDURE IF NOT EXISTS get_portfolio_analytics_summary(IN p_id VARCHAR(26))
BEGIN
  SELECT
    SUM(CASE WHEN event_type = 'view' THEN 1 ELSE 0 END) AS total_views,
    SUM(CASE WHEN event_type = 'view' AND DATE(created_at) = CURDATE() THEN 1 ELSE 0 END) AS today_views,
    SUM(CASE WHEN event_type = 'view' AND created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) THEN 1 ELSE 0 END) AS week_views,
    SUM(CASE WHEN event_type = 'pdf_download' THEN 1 ELSE 0 END) AS pdf_downloads,
    SUM(CASE WHEN event_type = 'hire_click' THEN 1 ELSE 0 END) AS hire_clicks
  FROM analytics
  WHERE portfolio_id = p_id;
END$$

DELIMITER ;

-- ─────────────────────────────────────────────────────────────
-- EVENTS (Auto-maintenance)
-- ─────────────────────────────────────────────────────────────

SET GLOBAL event_scheduler = ON;

CREATE EVENT IF NOT EXISTS evt_cleanup_expired_tokens
  ON SCHEDULE EVERY 1 HOUR
  DO CALL cleanup_expired_tokens();

-- ─────────────────────────────────────────────────────────────
-- SEED DATA (development only)
-- ─────────────────────────────────────────────────────────────

-- Insert demo user (password: Password123!)
-- bcrypt hash: $2a$10$...
INSERT IGNORE INTO users (id, email, password, name, plan, is_verified) VALUES
('01J5EXAMPLE000000000000001',
 'demo@ploykong.com',
 '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
 'Demo User',
 'pro',
 TRUE);

-- Insert demo portfolio
INSERT IGNORE INTO portfolios (id, user_id, slug, title, theme, is_published) VALUES
('01J5PORTFOLIO000000000001',
 '01J5EXAMPLE000000000000001',
 'demo',
 'Demo Portfolio',
 '{"mode":"dark","primary_color":"#4F46E5","font":"Syne","layout":"centered"}',
 TRUE);

-- Insert starter sections for demo portfolio
INSERT IGNORE INTO sections (id, portfolio_id, type, position, data) VALUES
('01J5SECTION0000000000001', '01J5PORTFOLIO000000000001', 'hero', 0,
 '{"name":"Demo User","role":"Full-Stack Developer","tagline":"Building cool things with Go, Vue.js & Flutter","avatar_url":"","show_hire_me":true}'),

('01J5SECTION0000000000002', '01J5PORTFOLIO000000000001', 'skills', 1,
 '{"items":[{"name":"Go","level":90,"category":"Backend"},{"name":"Vue.js 3","level":85,"category":"Frontend"},{"name":"Flutter","level":75,"category":"Mobile"},{"name":"MySQL","level":80,"category":"Database"},{"name":"Docker","level":70,"category":"DevOps"}]}'),

('01J5SECTION0000000000003', '01J5PORTFOLIO000000000001', 'projects', 2,
 '{"items":[{"title":"PloyKong","description":"No-Code Portfolio Builder for Gen-Z","image_url":"","live_url":"https://ploykong.com","github_url":"","tags":["Go","Vue.js","Flutter","MySQL"]}]}'),

('01J5SECTION0000000000004', '01J5PORTFOLIO000000000001', 'contact', 3,
 '{"email":"demo@ploykong.com","linkedin":"","github":"","location":"Bangkok, Thailand"}');
