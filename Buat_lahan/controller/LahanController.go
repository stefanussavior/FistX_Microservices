package controller

import (
	"buat_lahan/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InputDataLahan struct {
	Nama_lahan string `json:"nama_lahan"`
	Luas_lahan string `json:"luas_lahan"`
}

type UpdateLahan struct {
	Nama_lahan string `json:"nama_lahan"`
	Luas_lahan string `json:"luas_lahan"`
}

func TampilDataLahan(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var DataLahan []models.Lahan
	db.Raw("SELECT * FROM lahan").Scan(&DataLahan)

	c.JSON(http.StatusOK, gin.H{"Data": DataLahan})
}

func TambahDataLahan(c *gin.Context) {

	var Input InputDataLahan

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Lahan := models.Lahan{Nama_lahan: Input.Nama_lahan, Luas_lahan: Input.Luas_lahan}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Lahan)
	c.JSON(http.StatusOK, gin.H{"data": Lahan})
}

func CariDataLahan(c *gin.Context) {
	var Lahan models.Lahan
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&Lahan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": Lahan})
}

func UpdateDataLahan(c *gin.Context) {
	var Lahan models.Lahan
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&Lahan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input UpdateLahan

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&Lahan).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"Data": input})
}

func HapusDataLahan(c *gin.Context) {
	var Lahan models.Lahan

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&Lahan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Delete(&Lahan)

	c.JSON(http.StatusOK, gin.H{"Message": "data berhasil terhapus"})
}
