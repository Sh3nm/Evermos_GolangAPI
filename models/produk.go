package models

import "time"

type Produk struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	NamaProduk    string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;unique;not null"`
	HargaReseller string `gorm:"size:255;not null"`
	HargaKonsumen string `gorm:"size:255;not null"`
	Stok          int    `gorm:"not null"`
	Deskripsi     string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IDToko        uint // FK ke Toko
	IDCategory    uint // FK ke Category

	// Relasi
	FotoProduk []FotoProduk `gorm:"foreignKey:IDProduk"`
	LogProduk  LogProduk    `gorm:"foreignKey:IDProduk"`
}
