package models

import "time"

type Alamat struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	IDUser       uint   `gorm:"column:id_user"` // foreign key ke User
	JudulAlamat  string `gorm:"size:255"`
	NamaPenerima string `gorm:"size:255"`
	NoTelp       string `gorm:"size:255"`
	DetailAlamat string `gorm:"size:255"`
	UpdatedAt    time.Time
	CreatedAt    time.Time

	// relasi
	Trx Trx `gorm:"foreignKey:AlamatPengiriman"`
}
