package main

import (
	"log"

	"Project_Evermos/config"
	"Project_Evermos/models"
	"Project_Evermos/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init database
	config.InitDatabase()

	// init fiber
	app := fiber.New()

	// AutoMigrate models
	config.DB.AutoMigrate(
		&models.User{},
		&models.Toko{},
		&models.Category{},
		&models.Produk{},
		&models.FotoProduk{},
		&models.LogProduk{},
		&models.Trx{},
		&models.DetailTrx{},
	)

	// routes
	routes.SetupRoutes(app)

	// test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Halo, Selamat datang di Evermos!")
	})

	log.Fatal(app.Listen(":8080"))
}
