package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-assignment/config"
	"golang-assignment/models"
	"golang-assignment/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadExcel(c *gin.Context) {
	// Get the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is uploaded or invalid request"})
		return
	}

	// Save the uploaded file to a temporary location
	filePath := "./uploads/" + file.Filename
	os.MkdirAll("./uploads", os.ModePerm)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Respond to the client immediately
	c.JSON(http.StatusOK, gin.H{"message": "File is uploaded successfully"})

	// Process the file asynchronously in a Goroutine
	go func() {
		processExcelFile(filePath)
	}()
}

func processExcelFile(filePath string) {
	defer os.Remove(filePath) // Clean up the file after processing

	fmt.Println("Starting file processing:", filePath)
	records, err := utils.ParseExcel(filePath)
	if err != nil {
		fmt.Println("Failed to parse Excel file:", err)
		return
	}

	// Simulate processing time for large files (remove this in production)
	time.Sleep(2 * time.Second)

	// Insert parsed data into the database
	if err := utils.InsertRecords(records); err != nil {
		fmt.Println("Failed to insert records into MySQL:", err)
		return
	}

	fmt.Println("File processing completed successfully.")
}

func GetRecords(c *gin.Context) {
	ctx := context.Background()

	// Check for cached data
	cachedData, err := config.Redis.Get(ctx, "records_cache").Result()
	if err == nil {
		var records []models.Record
		if err := json.Unmarshal([]byte(cachedData), &records); err == nil {
			fmt.Printf("Returned %d records from cache\n", len(records))
			c.JSON(http.StatusOK, records)
			return
		}
	}

	// Fetch data from MySQL
	records, err := models.FetchRecordsFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records from MySQL"})
		return
	}

	if len(records) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No records found"})
		return
	}

	// Cache the data for 5 minutes
	data, _ := json.Marshal(records)
	config.Redis.Set(ctx, "records_cache", data, 5*time.Minute)

	fmt.Printf("Fetched %d records from MySQL\n", len(records))
	c.JSON(http.StatusOK, records)
}

func UpdateRecord(c *gin.Context) {
	id := c.Param("id")
	var record models.Record

	// Bind the incoming JSON to the record struct
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Update the record in MySQL
	if err := models.UpdateRecordInDB(id, record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	// Refresh Redis cache after updating
	records, err := models.FetchRecordsFromDB()
	if err == nil {
		models.CacheRecords(records)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
