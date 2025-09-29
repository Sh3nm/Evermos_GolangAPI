package controllers

import (
	"Project_Evermos/config"
	"Project_Evermos/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetMyProduk(c *fiber.Ctx) error {
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

	// Ambil toko user
	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	// Pagination & filtering
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	var produk []models.Produk
	query := config.DB.Where("id_toko = ?", toko.ID)

	// Filtering by nama_produk (opsional)
	if nama := c.Query("nama_produk"); nama != "" {
		query = query.Where("nama_produk LIKE ?", "%"+nama+"%")
	}

	if err := query.Offset(offset).Limit(limit).Find(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil produk"})
	}
	return c.JSON(produk)
}

func CreateProduk(c *fiber.Ctx) error {
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

	// Ambil toko user
	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	var input struct {
		NamaProduk    string `json:"nama_produk"`
		Slug          string `json:"slug"`
		HargaReseller string `json:"harga_reseller"`
		HargaKonsumen string `json:"harga_konsumen"`
		Stok          int    `json:"stok"`
		Deskripsi     string `json:"deskripsi"`
		IDCategory    uint   `json:"id_category"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	produk := models.Produk{
		NamaProduk:    input.NamaProduk,
		Slug:          input.Slug,
		HargaReseller: input.HargaReseller,
		HargaKonsumen: input.HargaKonsumen,
		Stok:          input.Stok,
		Deskripsi:     input.Deskripsi,
		IDToko:        toko.ID,
		IDCategory:    input.IDCategory,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := config.DB.Create(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat produk"})
	}
	return c.JSON(fiber.Map{"message": "Produk berhasil dibuat", "produk": produk})
}

func UpdateProduk(c *fiber.Ctx) error {
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

	// Ambil toko user
	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	id := c.Params("id")
	var produk models.Produk
	if err := config.DB.Where("id = ? AND id_toko = ?", id, toko.ID).First(&produk).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	var input struct {
		NamaProduk    string `json:"nama_produk"`
		Slug          string `json:"slug"`
		HargaReseller string `json:"harga_reseller"`
		HargaKonsumen string `json:"harga_konsumen"`
		Stok          int    `json:"stok"`
		Deskripsi     string `json:"deskripsi"`
		IDCategory    uint   `json:"id_category"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.NamaProduk != "" {
		produk.NamaProduk = input.NamaProduk
	}
	if input.Slug != "" {
		produk.Slug = input.Slug
	}
	if input.HargaReseller != "" {
		produk.HargaReseller = input.HargaReseller
	}
	if input.HargaKonsumen != "" {
		produk.HargaKonsumen = input.HargaKonsumen
	}
	if input.Stok != 0 {
		produk.Stok = input.Stok
	}
	if input.Deskripsi != "" {
		produk.Deskripsi = input.Deskripsi
	}
	if input.IDCategory != 0 {
		produk.IDCategory = input.IDCategory
	}
	produk.UpdatedAt = time.Now()
	config.DB.Save(&produk)
	return c.JSON(fiber.Map{"message": "Produk berhasil diupdate", "produk": produk})
}

func DeleteProduk(c *fiber.Ctx) error {
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

	// Ambil toko user
	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	id := c.Params("id")
	var produk models.Produk
	if err := config.DB.Where("id = ? AND id_toko = ?", id, toko.ID).First(&produk).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	if err := config.DB.Delete(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus produk"})
	}
	return c.JSON(fiber.Map{"message": "Produk berhasil dihapus"})
}

// Upload foto produk
func UploadProdukPhoto(c *fiber.Ctx) error {
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

	// Ambil toko user
	var toko models.Toko
	if err := config.DB.Where("id_user = ?", userID).First(&toko).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Toko tidak ditemukan"})
	}

	id := c.Params("id")
	var produk models.Produk
	if err := config.DB.Where("id = ? AND id_toko = ?", id, toko.ID).First(&produk).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	file, err := c.FormFile("foto")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}
	// Simpan file ke folder lokal (misal: uploads/produk/)
	path := "./upload/produk/" + file.Filename
	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Tambahkan foto ke slice FotoProduk
	produk.FotoProduk = append(produk.FotoProduk, models.FotoProduk{
		IDProduk: produk.ID,
		Url:      path,
	})
	produk.UpdatedAt = time.Now()
	config.DB.Save(&produk)

	return c.JSON(fiber.Map{"message": "Foto produk berhasil diupload", "url_foto": path})
}
