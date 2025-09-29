package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetMyAlamat(c *fiber.Ctx) error {
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

	var alamat []models.Alamat
	if err := config.DB.Where("id_user = ?", userID).Find(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil alamat"})
	}
	return c.JSON(alamat)
}

func CreateAlamat(c *fiber.Ctx) error {
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

	var input struct {
		JudulAlamat  string `json:"judul_alamat"`
		NamaPenerima string `json:"nama_penerima"`
		NoTelp       string `json:"notelp"`
		DetailAlamat string `json:"detail_alamat"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	alamat := models.Alamat{
		IDUser:       userID,
		JudulAlamat:  input.JudulAlamat,
		NamaPenerima: input.NamaPenerima,
		NoTelp:       input.NoTelp,
		DetailAlamat: input.DetailAlamat,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := config.DB.Create(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat alamat"})
	}
	return c.JSON(fiber.Map{"message": "Alamat berhasil dibuat", "alamat": alamat})
}

func UpdateAlamat(c *fiber.Ctx) error {
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

	id := c.Params("id")
	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND id_user = ?", id, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	var input struct {
		JudulAlamat  string `json:"judul_alamat"`
		NamaPenerima string `json:"nama_penerima"`
		NoTelp       string `json:"notelp"`
		DetailAlamat string `json:"detail_alamat"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.JudulAlamat != "" {
		alamat.JudulAlamat = input.JudulAlamat
	}
	if input.NamaPenerima != "" {
		alamat.NamaPenerima = input.NamaPenerima
	}
	if input.NoTelp != "" {
		alamat.NoTelp = input.NoTelp
	}
	if input.DetailAlamat != "" {
		alamat.DetailAlamat = input.DetailAlamat
	}
	alamat.UpdatedAt = time.Now()
	config.DB.Save(&alamat)
	return c.JSON(fiber.Map{"message": "Alamat berhasil diupdate", "alamat": alamat})
}

func DeleteAlamat(c *fiber.Ctx) error {
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

	id := c.Params("id")
	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND id_user = ?", id, userID).First(&alamat).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	if err := config.DB.Delete(&alamat).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus alamat"})
	}
	return c.JSON(fiber.Map{"message": "Alamat berhasil dihapus"})
}
