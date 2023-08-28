package api

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func RenderUploadForm(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func HandleUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("media")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Upload Error")
	}

	mediaDir := "./../../media"
	if _, err := os.Stat(mediaDir); os.IsNotExist(err) {
		os.Mkdir(mediaDir, 0755)
	}

	fileExtension := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + fileExtension

	filePath := mediaDir + "/" + newFilename
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save locally")
	}

	return c.SendString("File uploaded successfully!")
}
