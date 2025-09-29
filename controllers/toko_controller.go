package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetMyToko(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var userID uint
	switch v := userIDVal.(type) {
	case float64:
		userID = uint(v)
	case int:
		userID = uint(v)
	case uint:
		userID = v
	default:
		return c.Status(401).JSON(fiber.Map{"error": "Invalid user_id type"})
	}

	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko not found"})
	}
	return c.JSON(toko)
}

func UpdateMyToko(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var userID uint
	switch v := userIDVal.(type) {
	case float64:
		userID = uint(v)
	case int:
		userID = uint(v)
	case uint:
		userID = v
	default:
		return c.Status(401).JSON(fiber.Map{"error": "Invalid user_id type"})
	}

	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko not found"})
	}

	var input struct {
		NamaToko string `json:"nama_toko"`
		URLFoto  string `json:"url_foto"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.NamaToko != "" {
		toko.NamaToko = input.NamaToko
	}
	if input.URLFoto != "" {
		toko.URLFoto = input.URLFoto
	}
	toko.UpdatedAt = time.Now()
	config.DB.Save(&toko)
	return c.JSON(fiber.Map{"message": "Toko updated", "toko": toko})
}

func UploadTokoPhoto(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	if userIDVal == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}
	var userID uint
	switch v := userIDVal.(type) {
	case float64:
		userID = uint(v)
	case int:
		userID = uint(v)
	case uint:
		userID = v
	default:
		return c.Status(401).JSON(fiber.Map{"error": "Invalid user_id type"})
	}

	file, err := c.FormFile("foto")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}
	path := "./upload/toko/" + file.Filename
	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Update URLFoto toko di database
	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko not found"})
	}
	toko.URLFoto = path
	toko.UpdatedAt = time.Now()
	config.DB.Save(&toko)

	return c.JSON(fiber.Map{"message": "Foto toko berhasil diupload", "url_foto": path})
}
