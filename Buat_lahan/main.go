package main

import (
	"buat_lahan/controller"
	"buat_lahan/database"
	"buat_lahan/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.ConnectDatabase()
	db.AutoMigrate(&models.Lahan{}, &models.Lokasi_lahan{}).AddForeignKey("u_id", "lahan(id)", "RESTRICT", "RESTRICT")

	r.Use(func(c *gin.Context) {
		c.Set("db", db)

		//CORS Setting
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Tersambung"})
	})

	//route lahan
	r.GET("/api/lihat_lahan", controller.TampilDataLahan)
	r.POST("/api/tambah_data_lahan", controller.TambahDataLahan)
	r.GET("/api/ambil_data_lahan/:id", controller.CariDataLahan)
	r.PATCH("/api/ambil_data_lahan/:id", controller.UpdateDataLahan)
	r.DELETE("/api/hapus_data_lahan/:id", controller.HapusDataLahan)

	r.Run(":8080")
}
