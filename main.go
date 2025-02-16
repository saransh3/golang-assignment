package main

import (
	"golang-assignment/config"
	"golang-assignment/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MySQL and Redis connections
	config.InitMySQL()
	config.InitRedis()

	// Create a new Gin router
	r := gin.Default()

	// API endpoints
	r.POST("/upload", handlers.UploadExcel)
	r.GET("/records", handlers.GetRecords)
	r.PUT("/records/:id", handlers.UpdateRecord)

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
