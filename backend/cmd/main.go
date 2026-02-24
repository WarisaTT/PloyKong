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
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))
	app.Use(limiter.New(limiter.Config{
		Max:        100,
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
	auth.Post("/refresh", authHandler.RefreshToken)
	auth.Post("/logout", middleware.Protected(cfg), authHandler.Logout)

	// User (protected)
	userHandler := handlers.NewUserHandler(db)
	users := api.Group("/users", middleware.Protected(cfg))
	users.Get("/me", userHandler.GetMe)
	users.Put("/me", userHandler.UpdateMe)
	users.Delete("/me", userHandler.DeleteMe)

	// Portfolio (protected)
	portfolioHandler := handlers.NewPortfolioHandler(db)
	portfolios := api.Group("/portfolios", middleware.Protected(cfg))
	portfolios.Get("/", portfolioHandler.ListMyPortfolios)
	portfolios.Post("/", portfolioHandler.Create)
	portfolios.Get("/:id", portfolioHandler.GetByID)
	portfolios.Put("/:id", portfolioHandler.Update)
	portfolios.Delete("/:id", portfolioHandler.Delete)
	portfolios.Post("/:id/publish", portfolioHandler.Publish)
	portfolios.Post("/:id/unpublish", portfolioHandler.Unpublish)
	portfolios.Get("/:id/export-pdf", portfolioHandler.ExportPDF)

	// Sections (protected)
	sectionHandler := handlers.NewSectionHandler(db)
	sections := api.Group("/portfolios/:portfolioId/sections", middleware.Protected(cfg))
	sections.Get("/", sectionHandler.List)
	sections.Post("/", sectionHandler.Create)
	sections.Put("/:id", sectionHandler.Update)
	sections.Delete("/:id", sectionHandler.Delete)
	sections.Post("/reorder", sectionHandler.Reorder)

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

	// Public (no auth)
	publicHandler := handlers.NewPublicHandler(db)
	public := api.Group("/public")
	public.Get("/p/:slug", publicHandler.ViewPortfolio)
	public.Post("/p/:slug/track", publicHandler.TrackEvent)
	public.Post("/p/:slug/chat", publicHandler.AIChat)

	// Static uploads
	app.Static("/uploads", "./uploads")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("🚀 PloyKong API running on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
