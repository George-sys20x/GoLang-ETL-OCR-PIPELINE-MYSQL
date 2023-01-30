package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read the data from the OCR tool (in this example, a text file)
	data, err := readOCRData("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the SQL statement to insert the data into the database
	stmt, err := db.Prepare("INSERT INTO transactions (location, date, price) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Process the data to extract the location, date, and price
	for _, line := range data {
		// Split the line into fields
		fields := strings.Split(line, ":")
		if len(fields) != 3 {
			log.Printf("Invalid data format: %s", line)
			continue
		}

		// Extract the location, date, and price
		location := fields[0]
		dateString := fields[1]
		priceString := fields[2]

		// Parse the date
		date, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			log.Printf("Invalid date: %s", dateString)
			continue
		}

		// Parse the price
		price, err := strconv.ParseFloat(priceString, 64)
		if err != nil {
			log.Printf("Invalid price: %s", priceString)
			continue
		}

		// Insert the data into the database
		_, err = stmt.Exec(location, date, price)
		if err != nil {
			log.Printf("Error inserting data into the database: %s", err)
			continue
		}
	}
}


