package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func listAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	records := []Record{}
	record := Record{}
	database, err := sql.Open("sqlite3", "pin_records.db")
	defer database.Close()
	if err != nil {
		log.Printf("Error opening database: %v", err)
	}

	recordRows, err := database.Query(`
SELECT *  
FROM records;`)
	if err != nil {
		log.Println("Problem getting all rows")
	} else {
		for recordRows.Next() {
			err = recordRows.Scan(&record.PhoneNumber, &record.Carrier, &record.Score)
			if err != nil {
				log.Printf("Error getting row: %v", err)
			}
			records = append(records, record)
		}
		prettyJSON, err := json.MarshalIndent(records, "", "    ")
		if err != nil {
			log.Printf("Error marshal indent: %v", err)
		}
		fmt.Fprintf(w, string(prettyJSON))
	}
}
