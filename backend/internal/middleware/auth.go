package middleware

import (
	"strings"
	"time"

	"ploykong-api/internal/config"
	"ploykong-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const ContextUserKey = "user_claims"

// Protected validates JWT access token from Authorization header
func Protected(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from header: "Bearer <token>"
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "missing authorization header")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid authorization format")
		}

		tokenStr := parts[1]

		// Parse and validate token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired token")
		}

		mapClaims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token claims")
		}

		// Check expiry
		exp, ok := mapClaims["exp"].(float64)
		if !ok || time.Now().Unix() > int64(exp) {
			return fiber.NewError(fiber.StatusUnauthorized, "token expired")
		}

		// Attach claims to context
		claims := &models.Claims{
			UserID: mapClaims["user_id"].(string),
			Email:  mapClaims["email"].(string),
			Plan:   mapClaims["plan"].(string),
		}
		c.Locals(ContextUserKey, claims)

		return c.Next()
	}
}

// GetUserClaims extracts user claims from context
func GetUserClaims(c *fiber.Ctx) *models.Claims {
	claims, _ := c.Locals(ContextUserKey).(*models.Claims)
	return claims
}

// GenerateTokenPair creates access + refresh token pair
func GenerateTokenPair(claims *models.Claims, cfg *config.Config) (string, string, error) {
	// Access token (15min)
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims.UserID,
		"email":   claims.Email,
		"plan":    claims.Plan,
		"type":    "access",
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
	})
	accessStr, err := accessToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	// Refresh token (7 days)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": claims.UserID,
		"type":    "refresh",
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	})
	refreshStr, err := refreshToken.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return accessStr, refreshStr, nil
}
