package main

import (
	"auth/controller"
	"auth/models"
	"auth/models/test_database"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	db := test_database.TestDb()
	db.AutoMigrate(&models.Users{})
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	router.GET("/tampil_data_user", controller.TampilDataUsers)
	router.POST("/tambah_data_user", controller.TambahUsers)

	router.Run(":8080")
}
