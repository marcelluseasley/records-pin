package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func readRecordHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rec := Record{}
	phoneNumber := chi.URLParam(r, "phoneNumber")

	log.Printf("Retrieving phone number: %v", phoneNumber)
	database, err := sql.Open("sqlite3", "pin_records.db")
	defer database.Close()
	if err != nil {
		log.Printf("Error opening database. Cannot store created record: %v", err)
	}

	recordRow := database.QueryRow(fmt.Sprintf(`
SELECT phone_number, carrier, score 
FROM records
WHERE phone_number = '%s';`, phoneNumber))

	err = recordRow.Scan(&rec.PhoneNumber, &rec.Carrier, &rec.Score)
	if err == sql.ErrNoRows {
		log.Printf("Phone number %v not found.", phoneNumber)
		w.Write([]byte(`{"status" : "phone number not found"}`))
	} else {

		prettyJSON, err := json.MarshalIndent(rec, "", "    ")
		if err != nil {
			log.Printf("Error marshal indent: %v", err)
		}
		fmt.Fprintf(w, string(prettyJSON))
	}

}
