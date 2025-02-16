package utils

import (
	"fmt"
	"golang-assignment/config"
	"golang-assignment/models"
)

// InsertRecords inserts a slice of records into the MySQL database
func InsertRecords(records []models.Record) error {
	query := `
		INSERT INTO records (first_name, last_name, company_name, address, city, country, postal, phone, email, web)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	for _, record := range records {
		_, err := config.DB.Exec(query, record.FirstName, record.LastName, record.CompanyName, record.Address, record.City, record.Country, record.Postal, record.Phone, record.Email, record.Web)
		if err != nil {
			fmt.Println("Failed to insert record:", err)
			return err
		}
	}

	fmt.Printf("Successfully inserted %d records\n", len(records))
	return nil
}
