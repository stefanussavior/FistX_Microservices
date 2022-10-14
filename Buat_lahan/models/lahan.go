package models

import "gorm.io/gorm"

type Lahan struct {
	gorm.Model
	Nama_lahan string `gorm:"size:150;not null;" json:"nama_lahan"`
	Luas_lahan string `gorm:"size:100;not null;" json:"luas_lahan"`
}

func (Lahan) TableName() string {
	return "Lahan"
}

type Lokasi_lahan struct {
	Lahan      Lahan  `gorm:"foreignKey:ID_Lahan;association_foreignkey:id"`
	ID_lahan   uint   `gorm:"column:u_id" json:"-"`
	Provinsi   string `gorm:"size:150;" json:"provinsi"`
	Kota       string `gorm:"size:150;" json:"kota"`
	Kecamatan  string `gorm:"size:150;" json:"kecamatan"`
	Kelurahan  string `gorm:"size:150;" json:"kelurahan"`
	Sumber_air string `gorm:"size:150;" json:"sumber_air"`
}

func (Lokasi_lahan) TableName() string {
	return "lokasi_lahan"
}
