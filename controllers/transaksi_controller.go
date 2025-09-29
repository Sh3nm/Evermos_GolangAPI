package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaksi(c *fiber.Ctx) error {
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
		AlamatPengiriman uint `json:"alamat_pengiriman"`
		ProdukList       []struct {
			IDProduk  uint `json:"id_produk"`
			Kuantitas uint `json:"kuantitas"`
		} `json:"produk_list"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validasi alamat milik user
	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND id_user = ?", input.AlamatPengiriman, userID).First(&alamat).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Alamat tidak ditemukan"})
	}

	// Buat transaksi
	trx := models.Trx{
		IDUser:           userID,
		AlamatPengiriman: input.AlamatPengiriman,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	if err := config.DB.Create(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat transaksi"})
	}

	// Untuk setiap produk, buat log produk & detail transaksi
	for _, p := range input.ProdukList {
		var produk models.Produk
		if err := config.DB.First(&produk, p.IDProduk).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
		}

		// Buat log produk (copy data produk ke log)
		logProduk := models.LogProduk{
			IDProduk:      produk.ID,
			NamaProduk:    produk.NamaProduk,
			Slug:          produk.Slug,
			HargaReseller: parseUint(produk.HargaReseller),
			HargaKonsumen: parseUint(produk.HargaKonsumen),
			Deskripsi:     produk.Deskripsi,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			IDToko:        produk.IDToko,
			IDCategory:    produk.IDCategory,
		}
		if err := config.DB.Create(&logProduk).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat log produk"})
		}

		// Buat detail transaksi
		detail := models.DetailTrx{
			IDTrx:       trx.ID,
			IDLogProduk: logProduk.ID,
			IDToko:      produk.IDToko,
			Kuantitas:   p.Kuantitas,
			HargaTotal:  parseUint(produk.HargaKonsumen) * uint(p.Kuantitas),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		if err := config.DB.Create(&detail).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat detail transaksi"})
		}
	}

	return c.JSON(fiber.Map{"message": "Transaksi berhasil dibuat", "trx_id": trx.ID})
}

// Helper untuk konversi string ke uint
func parseUint(s string) uint {
	u, _ := strconv.ParseUint(s, 10, 64)
	return uint(u)
}

func GetMyTransaksi(c *fiber.Ctx) error {
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

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	var trxs []models.Trx
	query := config.DB.Where("id_user = ?", userID)
	if err := query.Offset(offset).Limit(limit).Find(&trxs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil transaksi"})
	}
	return c.JSON(trxs)
}
