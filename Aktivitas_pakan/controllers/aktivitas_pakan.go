package controllers

import (
	"aktivitas_pakan/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TambahAktivitasPakan struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Nama_kolam        string    `gorm:"size:255; not_null;" json:"nama_kolam"`
	Merk_pakan        string    `gorm:"size:255; not_null;" json:"merk_pakan"`
	Jenis_pakan       string    `gorm:"size:255; not_null;" json:"jenis_pakan"`
	Jumlah_pakan      string    `gorm:"size:255; not_null;" json:"jumlah_pakan"`
	Presentase        int32     `gorm:"size:100; not_null;" json:"presentase"`
	Tanggal_aktivitas time.Time `json:"tanggal_aktivitas"`
	Jam_aktivitas     time.Time `json:"jam_aktivitas"`
}

func TampilDataPakan(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var DataPakan []models.AktivitasPakan
	db.Find(&DataPakan)

	c.JSON(http.StatusOK, gin.H{"data": DataPakan})
}
