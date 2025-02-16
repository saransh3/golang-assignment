package models

import (
	"context"
	"encoding/json"
	"golang-assignment/config"
	"time"
)

type Record struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	Postal      string `json:"postal"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}

func InsertRecords(records []Record) error {
	query := "INSERT INTO records (first_name, last_name, company_name, city, address, country, postal, phone, email, web) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	for _, record := range records {
		_, err := config.DB.Exec(query, record.FirstName, record.LastName, record.CompanyName, record.City, record.Address, record.Country, record.Postal, record.Phone, record.Email, record.Web)
		if err != nil {
			return err
		}
	}
	return CacheRecords(records)
}

func FetchRecordsFromDB() ([]Record, error) {
	query := "SELECT id, first_name, last_name, company_name, city, address, country, postal, phone, email, web FROM records"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.ID, &record.FirstName, &record.LastName, &record.CompanyName, &record.City, &record.Address, &record.Country, &record.Postal, &record.Phone, &record.Email, &record.Web)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func CacheRecords(records []Record) error {
	ctx := context.Background()
	jsonData, err := json.Marshal(records)
	if err != nil {
		return err
	}
	return config.Redis.Set(ctx, "records_cache", jsonData, 5*time.Minute).Err()
}

func UpdateRecordInDB(id string, record Record) error {
	query := `
		UPDATE records 
		SET first_name=?, last_name=?, company_name=?, city=?, address=?, country=?, postal=?, phone=?, email=?, web=? 
		WHERE id=?
	`
	_, err := config.DB.Exec(query, record.FirstName, record.LastName, record.CompanyName, record.City, record.Address, record.Country, record.Postal, record.Phone, record.Email, record.Web, id)
	return err
}
