package main

import (
	"aktivitas_pakan/controllers"
	"aktivitas_pakan/database"
	"aktivitas_pakan/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.ConnectDatabase()
	db.AutoMigrate(&models.AktivitasPakan{})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		//CORS Setting
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Data telah masuk"})
	})

	r.GET("/api/lihat_data_pakan", controllers.TampilDataPakan)

	r.Use(cors.New(config))
	r.Run(":8080")
}
