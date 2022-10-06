package controller

import (
	"auth/models"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TambahDataUsers struct {
	Id           uint      `json:"id" gorm:"primary_key"`
	Email        string    `json:"email"`
	Nama_lengkap string    `json:"nama_lengkap"`
	NomorTelp    int16     `json:"NomorTelp"`
	Password     string    `json:"password"`
	Upload_file  string    `json:"upload_file"`
	Created_at   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func TampilDataUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Users []models.Users
	db.Find(&Users)
	c.JSON(http.StatusOK, gin.H{"data": Users})
}

func TambahUsers(c *gin.Context) {
	var input TambahDataUsers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//validasi input ke database
	Users := models.Users{
		Email:        input.Email,
		Nama_lengkap: input.Nama_lengkap,
		NomorTelp:    input.NomorTelp,
		Password:     input.Password,
		Upload_file:  input.Password,
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Users)

	c.JSON(http.StatusOK, gin.H{"data": Users})
}
