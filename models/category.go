package models

import "time"

type Category struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	NamaCategory string `gorm:"size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Produk    Produk    `gorm:"foreignKey:IDCategory;references:ID"`
	LogProduk LogProduk `gorm:"foreignKey:IDCategory;references:ID"`
}
