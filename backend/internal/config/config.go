package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port           string
	AllowedOrigins string
	Env            string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT
	JWTSecret        string
	JWTAccessExpiry  string
	JWTRefreshExpiry string

	// OpenAI
	OpenAIKey   string
	OpenAIModel string

	// Storage
	UploadDir     string
	MaxUploadMB   int64
	CloudinaryURL string

	// S3 Storage
	S3Bucket    string
	S3Region    string
	S3AccessKey string
	S3SecretKey string
	S3Endpoint  string // For S3-compatible services like DigitalOcean Spaces or MinIO

	// Google Auth
	GoogleClientID string

	// URLs (for SEO/Redirects)
	BaseURL string // API Base URL
	FEURL   string // Frontend Base URL
}

func Load() *Config {
	// Load .env file (ignore error in production, env vars will be set by Docker)
	if err := godotenv.Load(); err != nil {
		log.Println("ℹ️  No .env file found, using system environment variables")
	}

	return &Config{
		// Server
		Port:           getEnv("PORT", "3000"),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:5173,https://ploy-kong.vercel.app"),
		Env:            getEnv("ENV", "development"),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "secret"),
		DBName:     getEnv("DB_NAME", "ploykong"),

		// JWT
		JWTSecret:        getEnv("JWT_SECRET", "change-this-super-secret-key-in-production"),
		JWTAccessExpiry:  getEnv("JWT_ACCESS_EXPIRY", "15m"),
		JWTRefreshExpiry: getEnv("JWT_REFRESH_EXPIRY", "7d"),

		// OpenAI
		OpenAIKey:   getEnv("OPENAI_API_KEY", ""),
		OpenAIModel: getEnv("OPENAI_MODEL", "gpt-4o-mini"),

		// Storage
		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		MaxUploadMB:   10,
		CloudinaryURL: getEnv("CLOUDINARY_URL", ""),

		// S3 Storage
		S3Bucket:       getEnv("S3_BUCKET", ""),
		S3Region:       getEnv("S3_REGION", "ap-southeast-1"),
		S3AccessKey:    getEnv("S3_ACCESS_KEY", ""),
		S3SecretKey:    getEnv("S3_SECRET_KEY", ""),
		S3Endpoint:     getEnv("S3_ENDPOINT", ""),
		GoogleClientID: getEnv("GOOGLE_CLIENT_ID", ""),
		BaseURL:        strings.TrimRight(getEnv("BASE_URL", "http://localhost:3000"), "/) "),
		FEURL:          strings.TrimRight(getEnv("FE_URL", "https://ploy-kong.vercel.app"), "/) "),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
