package handlers

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ploykong-api/internal/config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

type UploadHandler struct {
	cfg *config.Config
}

func NewUploadHandler(cfg *config.Config) *UploadHandler {
	return &UploadHandler{cfg: cfg}
}

// UploadImage handles multipart form file uploads
func (h *UploadHandler) UploadImage(c *fiber.Ctx) error {
	// 1. Get the file from form
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "No image found in request")
	}

	// 2. Validate file type / extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".webp" && ext != ".gif" {
		return fiber.NewError(fiber.StatusBadRequest, "Only images allowed (.png, .jpg, .jpeg, .webp, .gif)")
	}

	// 3. Size limit
	if fileHeader.Size > h.cfg.MaxUploadMB*1024*1024 {
		return fiber.NewError(fiber.StatusRequestEntityTooLarge, fmt.Sprintf("File size must be strictly less than %dMB", h.cfg.MaxUploadMB))
	}

	// 4. Generate safe, unique filename
	filename := fmt.Sprintf("%s_%d%s", ulid.Make().String(), time.Now().Unix(), ext)

	// 5. Check if Cloudinary is configured
	if h.cfg.CloudinaryURL != "" {
		cld, err := cloudinary.NewFromURL(h.cfg.CloudinaryURL)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Cloudinary configuration error")
		}

		file, err := fileHeader.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to open upload file")
		}
		defer file.Close()

		ctx := context.Background()
		resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
			Folder:   "ploykong_uploads",
			PublicID: strings.TrimSuffix(filename, ext),
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload to Cloudinary")
		}

		return c.JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"url": resp.SecureURL,
			},
		})
	}

	// 6. Fallback: Local file system (works but volatile on stateless deployments)
	uploadDir := h.cfg.UploadDir
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create upload directory")
	}

	savePath := filepath.Join(uploadDir, filename)
	if err := c.SaveFile(fileHeader, savePath); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file locally")
	}

	fileURL := fmt.Sprintf("/uploads/%s", filename)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"url": fileURL,
		},
	})
}
