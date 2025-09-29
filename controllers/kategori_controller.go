package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllKategori(c *fiber.Ctx) error {
	var kategori []models.Category
	if err := config.DB.Find(&kategori).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil kategori"})
	}
	return c.JSON(kategori)
}

func CreateKategori(c *fiber.Ctx) error {
	var input struct {
		NamaCategory string `json:"nama_category"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	kategori := models.Category{
		NamaCategory: input.NamaCategory,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := config.DB.Create(&kategori).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat kategori"})
	}
	return c.JSON(fiber.Map{"message": "Kategori berhasil dibuat", "kategori": kategori})
}

func UpdateKategori(c *fiber.Ctx) error {
	id := c.Params("id")
	var kategori models.Category
	if err := config.DB.First(&kategori, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}
	var input struct {
		NamaCategory string `json:"nama_category"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if input.NamaCategory != "" {
		kategori.NamaCategory = input.NamaCategory
	}
	kategori.UpdatedAt = time.Now()
	config.DB.Save(&kategori)
	return c.JSON(fiber.Map{"message": "Kategori berhasil diupdate", "kategori": kategori})
}

func DeleteKategori(c *fiber.Ctx) error {
	id := c.Params("id")
	var kategori models.Category
	if err := config.DB.First(&kategori, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}
	if err := config.DB.Delete(&kategori).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus kategori"})
	}
	return c.JSON(fiber.Map{"message": "Kategori berhasil dihapus"})
}
