package models

import "time"

type DetailTrx struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	IDTrx       uint // foreign key trx
	IDLogProduk uint // foreign key log produk
	IDToko      uint // foreign key toko
	Kuantitas   uint
	HargaTotal  uint
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
