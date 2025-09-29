package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Register
func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Nama      string `json:"nama"`
		Email     string `json:"email"`
		NoTelp    string `json:"notelp"`
		KataSandi string `json:"kata_sandi"`
	}

	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi input tidak kosong
	if input.Nama == "" || input.Email == "" || input.NoTelp == "" || input.KataSandi == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Semua field wajib diisi"})
	}

	// Cek email dan no telp unik
	var count int64
	config.DB.Model(&models.User{}).Where("email = ? OR notelp = ?", input.Email, input.NoTelp).Count(&count)
	if count > 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Email atau No Telp sudah digunakan"})
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.KataSandi), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal hash password"})
	}

	// Buat user
	user := models.User{
		Nama:      input.Nama,
		Email:     input.Email,
		NoTelp:    input.NoTelp,
		KataSandi: string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat user"})
	}

	// Buat toko otomatis
	toko := models.Toko{
		IDUser:    user.ID,
		NamaToko:  user.Nama + " Store",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := config.DB.Create(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat toko"})
	}

	return c.JSON(fiber.Map{"message": "Register berhasil", "user_id": user.ID, "toko_id": toko.ID})
}

// Login
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		EmailOrTelp string `json:"email_or_telp"`
		KataSandi   string `json:"kata_sandi"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.EmailOrTelp == "" || input.KataSandi == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email/No Telp dan Kata Sandi wajib diisi"})
	}

	var user models.User
	if err := config.DB.Where("email = ? OR notelp = ?", input.EmailOrTelp, input.EmailOrTelp).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.KataSandi), []byte(input.KataSandi)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Password salah"})
	}

	// Generate JWT
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.Status(500).JSON(fiber.Map{"error": "JWT_SECRET belum diset di .env"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal generate token"})
	}

	return c.JSON(fiber.Map{"token": jwtStr})
}
