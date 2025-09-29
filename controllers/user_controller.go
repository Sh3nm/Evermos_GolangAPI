package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetMyProfile(c *fiber.Ctx) error {
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

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func UpdateMyProfile(c *fiber.Ctx) error {
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

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	var updateData struct {
		Nama         string `json:"nama"`
		KataSandi    string `json:"kata_sandi"`
		NoTelp       string `json:"notelp"`
		TanggalLahir string `json:"tanggal_lahir"`
		JenisKelamin string `json:"jenis_kelamin"`
		Tentang      string `json:"tentang"`
		Pekerjaan    string `json:"pekerjaan"`
		Email        string `json:"email"`
		IDProvinsi   string `json:"id_provinsi"`
		IDKota       string `json:"id_kota"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Cek email unik jika diupdate
	if updateData.Email != "" && updateData.Email != user.Email {
		var count int64
		config.DB.Model(&models.User{}).Where("email = ? AND id != ?", updateData.Email, userID).Count(&count)
		if count > 0 {
			return c.Status(400).JSON(fiber.Map{"error": "Email sudah digunakan user lain"})
		}
		user.Email = updateData.Email
	}

	// Cek notelp unik jika diupdate
	if updateData.NoTelp != "" && updateData.NoTelp != user.NoTelp {
		var count int64
		config.DB.Model(&models.User{}).Where("notelp = ? AND id != ?", updateData.NoTelp, userID).Count(&count)
		if count > 0 {
			return c.Status(400).JSON(fiber.Map{"error": "No Telp sudah digunakan user lain"})
		}
		user.NoTelp = updateData.NoTelp
	}

	if updateData.Nama != "" {
		user.Nama = updateData.Nama
	}
	if updateData.KataSandi != "" {
		user.KataSandi = updateData.KataSandi
	}
	if updateData.TanggalLahir != "" {
		parsed, err := time.Parse("2006-01-02", updateData.TanggalLahir)
		if err == nil {
			user.TanggalLahir = parsed
		} else {
			return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_lahir harus YYYY-MM-DD"})
		}
	}
	if updateData.JenisKelamin != "" {
		user.JenisKelamin = updateData.JenisKelamin
	}
	if updateData.Tentang != "" {
		user.Tentang = updateData.Tentang
	}
	if updateData.Pekerjaan != "" {
		user.Pekerjaan = updateData.Pekerjaan
	}
	if updateData.IDProvinsi != "" {
		user.IDProvinsi = updateData.IDProvinsi
	}
	if updateData.IDKota != "" {
		user.IDKota = updateData.IDKota
	}

	config.DB.Save(&user)
	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
		"user":    user,
	})
}
