package routes

import (
	"Project_Evermos/controllers"
	"Project_Evermos/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root Check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(("API is running"))
	})

	// Auth
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)

	// User
	user := app.Group("/users", middlewares.JWTProtected())
	user.Get("/me", controllers.GetMyProfile)
	user.Put("/me", controllers.UpdateMyProfile)

	// Toko
	toko := app.Group("/tokos", middlewares.JWTProtected())
	toko.Get("/me", controllers.GetMyToko)
	toko.Put("/me", controllers.UpdateMyToko)
	toko.Post("/upload_foto", controllers.UploadTokoPhoto)

	// Alamat
	alamat := app.Group("/alamat", middlewares.JWTProtected())
	alamat.Post("/", controllers.CreateAlamat)
	alamat.Get("/", controllers.GetMyAlamat)
	alamat.Put("/:id", controllers.UpdateAlamat)
	alamat.Delete("/:id", controllers.DeleteAlamat)

	// Kategori
	kategori := app.Group("/kategori", middlewares.JWTProtected(), middlewares.Admin())
	kategori.Post("/", controllers.CreateKategori)
	kategori.Get("/", controllers.GetAllKategori)
	kategori.Put("/:id", controllers.UpdateKategori)
	kategori.Delete("/:id", controllers.DeleteKategori)

	// Produk
	produk := app.Group("/produk", middlewares.JWTProtected())
	produk.Post("/", controllers.CreateProduk)
	produk.Get("/myproduk", controllers.GetMyProduk)
	produk.Put("/:id", controllers.UpdateProduk)
	produk.Delete("/:id", controllers.DeleteProduk)
	produk.Post("/upload_foto/:id", controllers.UploadProdukPhoto)

	// Trx
	trx := app.Group("/transaksi", middlewares.JWTProtected())
	trx.Get("/", controllers.GetMyTransaksi)   // List transaksi sendiri (pagination/filtering)
	trx.Post("/", controllers.CreateTransaksi) // Buat transaksi
}
