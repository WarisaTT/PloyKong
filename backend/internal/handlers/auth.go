package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"time"

	"ploykong-api/internal/config"
	"ploykong-api/internal/middleware"
	"ploykong-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db  *sql.DB
	cfg *config.Config
}

func NewAuthHandler(db *sql.DB, cfg *config.Config) *AuthHandler {
	return &AuthHandler{db: db, cfg: cfg}
}

// ─── Register ────────────────────────────────────────────────────────────────

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	// Validate
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email, password, and name are required")
	}
	if len(req.Password) < 8 {
		return fiber.NewError(fiber.StatusBadRequest, "password must be at least 8 characters")
	}

	// Check email exists
	var count int
	if err := h.db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", req.Email).Scan(&count); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database error")
	}
	if count > 0 {
		return fiber.NewError(fiber.StatusConflict, "email already registered")
	}

	// Hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to hash password")
	}

	// Create user
	userID := ulid.Make().String()
	_, err = h.db.Exec(
		"INSERT INTO users (id, email, password, name) VALUES (?, ?, ?, ?)",
		userID, req.Email, string(hashedPwd), req.Name,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create user")
	}

	// Generate tokens
	claims := &models.Claims{UserID: userID, Email: req.Email, Plan: "free"}
	accessToken, refreshToken, err := middleware.GenerateTokenPair(claims, h.cfg)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate tokens")
	}

	// Store refresh token hash
	if err := h.storeRefreshToken(userID, refreshToken); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to store refresh token")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "registered successfully",
		"data": fiber.Map{
			"user": fiber.Map{
				"id":    userID,
				"email": req.Email,
				"name":  req.Name,
				"plan":  "free",
			},
			"tokens": fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
				"expires_in":    900, // 15 minutes
			},
		},
	})
}

// ─── Login ────────────────────────────────────────────────────────────────────

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if req.Email == "" || req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "email and password are required")
	}

	// Find user
	var user models.User
	err := h.db.QueryRow(
		"SELECT id, email, password, name, plan FROM users WHERE email = ? AND deleted_at IS NULL",
		req.Email,
	).Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.Plan)

	if err == sql.ErrNoRows {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid email or password")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database error")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid email or password")
	}

	// Generate tokens
	claims := &models.Claims{UserID: user.ID, Email: user.Email, Plan: user.Plan}
	accessToken, refreshToken, err := middleware.GenerateTokenPair(claims, h.cfg)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate tokens")
	}

	// Store refresh token
	if err := h.storeRefreshToken(user.ID, refreshToken); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to store refresh token")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": fiber.Map{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.Name,
				"plan":  user.Plan,
			},
			"tokens": fiber.Map{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
				"expires_in":    900,
			},
		},
	})
}

// ─── Refresh Token ────────────────────────────────────────────────────────────

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	// Hash the incoming token and look it up
	hash := hashToken(req.RefreshToken)

	var userID string
	var expiresAt time.Time
	err := h.db.QueryRow(
		"SELECT user_id, expires_at FROM refresh_tokens WHERE token_hash = ?",
		hash,
	).Scan(&userID, &expiresAt)

	if err == sql.ErrNoRows {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid refresh token")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "database error")
	}

	if time.Now().After(expiresAt) {
		return fiber.NewError(fiber.StatusUnauthorized, "refresh token expired")
	}

	// Get user
	var user models.User
	if err := h.db.QueryRow(
		"SELECT id, email, plan FROM users WHERE id = ? AND deleted_at IS NULL", userID,
	).Scan(&user.ID, &user.Email, &user.Plan); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user not found")
	}

	// Rotate tokens (invalidate old refresh token)
	h.db.Exec("DELETE FROM refresh_tokens WHERE token_hash = ?", hash)

	claims := &models.Claims{UserID: user.ID, Email: user.Email, Plan: user.Plan}
	accessToken, newRefreshToken, err := middleware.GenerateTokenPair(claims, h.cfg)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to generate tokens")
	}

	if err := h.storeRefreshToken(user.ID, newRefreshToken); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to store refresh token")
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"access_token":  accessToken,
			"refresh_token": newRefreshToken,
			"expires_in":    900,
		},
	})
}

// ─── Logout ───────────────────────────────────────────────────────────────────

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	var req LogoutRequest
	c.BodyParser(&req)

	if req.RefreshToken != "" {
		hash := hashToken(req.RefreshToken)
		h.db.Exec("DELETE FROM refresh_tokens WHERE token_hash = ?", hash)
	}

	return c.JSON(fiber.Map{"success": true, "message": "logged out successfully"})
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

func (h *AuthHandler) storeRefreshToken(userID, token string) error {
	id := ulid.Make().String()
	hash := hashToken(token)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	_, err := h.db.Exec(
		"INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at) VALUES (?, ?, ?, ?)",
		id, userID, hash, expiresAt,
	)
	return err
}

func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}

// generate slug from name (utility)
func generateSlug(name string) string {
	return fmt.Sprintf("%s-%d", name, time.Now().Unix()%10000)
}
