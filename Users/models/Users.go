package models

import (
	"time"
)

type Users struct {
	Id            uint      `json:"id" gorm:"primary_key"`
	Email         string    `json:"email"`
	Nama_lengkap  string    `json:"nama_lengkap"`
	Tanggal_lahir time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"tanggal_lahir"`
	NomorTelp     int16     `json:"NomorTelp"`
	Password      string    `json:"password"`
	Upload_file   string    `json:"upload_file"`
	Created_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
