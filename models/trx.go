package models

import "time"

type Trx struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	IDUser           uint // foreign key user
	AlamatPengiriman uint // foreign key alamat
	TotalHarga       uint
	KodeInvoice      string `gorm:"size:255"`
	MethodBayar      string `gorm:"size:255"`
	UpdatedAt        time.Time
	CreatedAt        time.Time

	// relasi
	DetailTrx DetailTrx `gorm:"foreignKey:IDTrx"`
}
