package main

import (
	"log"
	"os"
	"time"

	"ploykong-api/internal/config"
	"ploykong-api/internal/database"
	"ploykong-api/internal/handlers"
	"ploykong-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer db.Close()

	if err := database.Migrate(db); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName:      "PloyKong API v1.0",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		BodyLimit:    50 * 1024 * 1024, // 50MB
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	})

	// ─── Global Middleware ───────────────────────────────────────────────────
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} | ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	app.Use(limiter.New(limiter.Config{
		Max:        1000,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	}))

	// ─── Health Check ────────────────────────────────────────────────────────
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "version": "1.0.0"})
	})

	// ─── API Routes ──────────────────────────────────────────────────────────
	api := app.Group("/api/v1")

	// Auth (public)
	authHandler := handlers.NewAuthHandler(db, cfg)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/google", authHandler.GoogleLogin)
	auth.Post("/refresh", authHandler.RefreshToken)
	auth.Post("/logout", middleware.Protected(cfg), authHandler.Logout)

	// User (protected)
	userHandler := handlers.NewUserHandler(db)
	users := api.Group("/users", middleware.Protected(cfg))
	users.Get("/me", userHandler.GetMe)
	users.Patch("/me", userHandler.UpdateMe)
	users.Delete("/me", userHandler.DeleteMe)

	// Portfolio (protected)
	portfolioHandler := handlers.NewPortfolioHandler(db)
	portfolios := api.Group("/portfolios", middleware.Protected(cfg))
	portfolios.Get("/", portfolioHandler.ListMyPortfolios)
	portfolios.Post("/", portfolioHandler.Create)
	portfolios.Get("/:id", portfolioHandler.GetByID)
	portfolios.Patch("/:id", portfolioHandler.Update)
	portfolios.Delete("/:id", portfolioHandler.Delete)
	portfolios.Post("/:id/publish", portfolioHandler.Publish)
	portfolios.Post("/:id/unpublish", portfolioHandler.Unpublish)
	portfolios.Post("/:id/duplicate", portfolioHandler.Duplicate)

	// Sections (protected)
	sectionHandler := handlers.NewSectionHandler(db)
	portSections := api.Group("/portfolios/:portfolioId/sections", middleware.Protected(cfg))
	portSections.Get("/", sectionHandler.List)
	portSections.Post("/", sectionHandler.Create)

	globalSections := api.Group("/sections", middleware.Protected(cfg))
	globalSections.Patch("/:id", sectionHandler.Update)
	globalSections.Delete("/:id", sectionHandler.Delete)
	globalSections.Post("/:id/duplicate", sectionHandler.Duplicate)
	globalSections.Post("/reorder", sectionHandler.Reorder)

	// Analytics (protected)
	analyticsHandler := handlers.NewAnalyticsHandler(db)
	analytics := api.Group("/analytics", middleware.Protected(cfg))
	analytics.Get("/:portfolioId/summary", analyticsHandler.Summary)
	analytics.Get("/:portfolioId/visitors", analyticsHandler.Visitors)
	analytics.Get("/:portfolioId/sources", analyticsHandler.Sources)

	// AI (protected)
	aiHandler := handlers.NewAIHandler(cfg)
	ai := api.Group("/ai", middleware.Protected(cfg))
	ai.Post("/improve-text", aiHandler.ImproveText)
	ai.Post("/score-resume", aiHandler.ScoreResume)
	ai.Post("/suggest-skills", aiHandler.SuggestSkills)

	// Upload (protected)
	uploadHandler := handlers.NewUploadHandler(cfg)
	api.Post("/upload", middleware.Protected(cfg), uploadHandler.UploadImage)

	// Public (no auth)
	publicHandler := handlers.NewPublicHandler(db, cfg)
	public := api.Group("/public")
	public.Get("/p/:slug", publicHandler.ViewPortfolio)
	public.Post("/p/:slug/track", publicHandler.TrackEvent)
	public.Post("/p/:slug/chat", publicHandler.AIChat)
	public.Get("/p/:slug/pdf", publicHandler.ExportPDFBySlug)
	public.Get("/id/:id/pdf", portfolioHandler.GeneratePDF)

	// Static uploads
	app.Static("/uploads", cfg.UploadDir)

	port := os.Getenv("PORT")

	// If running on Vercel, we force port 3000 (standard Vercel port forwarding)
	if port == "" {
		port = "3000" // local dev default
	}

	log.Printf("🚀 PloyKong API running on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
