package models

import (
	"time"
)

type Toko struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	IDUser    uint      // foreign key ke User
	NamaToko  string    `gorm:"size:255"`
	URLFoto   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// relasi
	Produk    []Produk    `gorm:"foreignKey:IDToko"`
	DetailTrx []DetailTrx `gorm:"foreignKey:IDToko"`
	LogProduk []LogProduk `gorm:"foreignKey:IDToko"`
}
