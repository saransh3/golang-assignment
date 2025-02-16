package utils

import (
	"fmt"
	"golang-assignment/models"

	"github.com/xuri/excelize/v2"
)

func ParseExcel(filePath string) ([]models.Record, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("Failed to open Excel file:", err)
		return nil, err
	}
	defer f.Close()

	rows, err := f.GetRows("uk-500")
	if err != nil {
		fmt.Println("Failed to get rows from Excel:", err)
		return nil, err
	}

	fmt.Printf("Found %d rows in Excel file\n", len(rows))
	if len(rows) < 2 {
		return nil, fmt.Errorf("Excel file does not contain enough data")
	}

	var records []models.Record
	for i, row := range rows {
		if i == 0 {
			continue // Skip header row
		}

		fmt.Printf("Processing row %d: %v\n", i+1, row)
		if len(row) < 10 {
			fmt.Printf("Row %d skipped: insufficient columns\n", i+1)
			continue
		}

		record := models.Record{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			Country:     row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		}
		records = append(records, record)
	}

	fmt.Printf("Successfully parsed %d records\n", len(records))
	return records, nil
}
