package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

// UploadImage handles multipart form file uploads
func UploadImage(c *fiber.Ctx) error {
	// 1. Get the file from form
	file, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "No image found in request")
	}

	// 2. Validate file type / extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".webp" && ext != ".gif" {
		return fiber.NewError(fiber.StatusBadRequest, "Only images allowed (.png, .jpg, .jpeg, .webp, .gif)")
	}

	// 3. Size limit (e.g. 25MB)
	if file.Size > 25*1024*1024 {
		return fiber.NewError(fiber.StatusRequestEntityTooLarge, "File size must be strictly less than 25MB")
	}

	// 4. Ensure uploads directory exists
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create upload directory")
	}

	// 5. Generate safe, unique filename
	filename := fmt.Sprintf("%s_%d%s", ulid.Make().String(), time.Now().Unix(), ext)
	savePath := filepath.Join(uploadDir, filename)

	// 6. Save file
	if err := c.SaveFile(file, savePath); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
	}

	// 7. Return URL
	// We use the absolute path root /uploads/
	fileURL := fmt.Sprintf("/uploads/%s", filename)

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"url": fileURL,
		},
	})
}
