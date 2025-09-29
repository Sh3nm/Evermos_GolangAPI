package models

import "time"

type FotoProduk struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	IDProduk  uint   // foreign key produk
	Url       string `gorm:"size:255"`
	UpdatedAt time.Time
	CreatedAt time.Time

	// relasi
	// LogProduk []LogProduk
}
