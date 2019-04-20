package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func createRecordHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Fprintf(w, "this api call only accepts application/json requests")
	} else {
		rec := Record{}
		json.NewDecoder(r.Body).Decode(&rec)

		database, err := sql.Open("sqlite3", "pin_records.db")
		defer database.Close()
		if err != nil {
			log.Printf("Error opening database. Cannot store created record: %v", err)
		}
		statement, err := database.Prepare("INSERT INTO records (phone_number, carrier, score) VALUES (?,?,?);")
		if err != nil {
			log.Printf("Error preparing database: %v", err)
		}

		_, err = statement.Exec(rec.PhoneNumber, rec.Carrier, rec.Score)
		if err != nil {
			log.Printf("Error executing database statement: %v", err)
		} else {
			log.Printf("Added record to database: %v", rec)
			w.Write([]byte(`{"status" : "record written successfully"}`))
		}
	}

}
