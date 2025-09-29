package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement"`
	Nama         string    `gorm:"column:nama;size:255;not null"`          // VARCHAR(255)
	KataSandi    string    `gorm:"column:kata_sandi;size:255;not null"`    // VARCHAR(255)
	NoTelp       string    `gorm:"column:notelp;size:255;unique;not null"` // VARCHAR(50)
	TanggalLahir time.Time `gorm:"column:tanggal_lahir"`                   // DATE
	JenisKelamin string    `gorm:"column:jenis_kelamin;size:255"`
	Tentang      string    `gorm:"column;tentangtype:text"`               // TEXT
	Pekerjaan    string    `gorm:"column:pekerjaan;size:255"`             // VARCHAR(100)
	Email        string    `gorm:"column:email;size:255;unique;not null"` // VARCHAR(100)
	IDProvinsi   string    `gorm:"column:id_provinsi;size:255"`           // VARCHAR(100)
	IDKota       string    `gorm:"column:id_kota;size:255"`               // VARCHAR(100)
	IsAdmin      bool      `gorm:"column:isAdmin;default:false"`
	CreatedAt    time.Time `gorm:"column:createdAt"`
	UpdatedAt    time.Time `gorm:"column:updatedAt"`

	//relasi
	Toko   Toko     `gorm:"foreignKey:IDUser"`
	Alamat []Alamat `gorm:"foreignKey:IDUser"`
	Trx    Trx      `gorm:"foreignKey:IDUser"`
}
