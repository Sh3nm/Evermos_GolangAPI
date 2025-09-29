package models

import "time"

type LogProduk struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	IDProduk      uint   // foreign key produk
	NamaProduk    string `gorm:"size:255"`
	Slug          string `gorm:"size:255"`
	HargaReseller uint   `gorm:"size:255"`
	HargaKonsumen uint   `gorm:"size:255"`
	Deskripsi     string `gorm:"type:text"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IDToko        uint // foreign key toko
	IDCategory    uint // foreign key category

	// relasi
	// FotoProduk []FotoProduk
	DetailTrx []DetailTrx `gorm:"foreignKey:IDLogProduk"`
}
