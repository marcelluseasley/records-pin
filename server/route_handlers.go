package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func createRecordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Fprintf(w, "this api call only accepts application/json requests")
	}

	rec := Record{}
	json.NewDecoder(r.Body).Decode(&rec)

	database, err := sql.Open("sqlite3", "pin_records.db")
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status" : "record written successfully"}`))
	}

}

func readRecordHandler(w http.ResponseWriter, r *http.Request) {
	rec := Record{}
	phoneNumber := chi.URLParam(r, "phoneNumber")

	log.Printf("Retrieving phone number: %v", phoneNumber)
	database, err := sql.Open("sqlite3", "pin_records.db")
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status" : "phone number not found"}`))
	} else {
		json.NewEncoder(w).Encode(rec)
	}

}

func updateRecordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Fprintf(w, "this api call only accepts application/json requests")
	}

	rec := Record{}
	
	phoneNumber := chi.URLParam(r, "phoneNumber")

	database, err := sql.Open("sqlite3", "pin_records.db")
	if err != nil {
		log.Printf("Error opening database. Cannot update record: %v", err)
	}

	recordRow := database.QueryRow(fmt.Sprintf(`
SELECT phone_number, carrier, score 
FROM records
WHERE phone_number = '%s';`, phoneNumber))

	// we don't really care about these values being put into the struct
	// at this point. But Scan requires the same number of arguments
	// as number of columns returned.
	err = recordRow.Scan(&rec.PhoneNumber, &rec.Carrier, &rec.Score)
	if err == sql.ErrNoRows {
		log.Printf("Phone number %v not found.", phoneNumber)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status" : "phone number not found"}`))
	} else {

		statement, err := database.Prepare("UPDATE records SET phone_number=?, carrier=?, score=? WHERE phone_number=?;")
		if err != nil {
			log.Printf("Error preparing database: %v", err)
		}
		json.NewDecoder(r.Body).Decode(&rec)
		_, err = statement.Exec(rec.PhoneNumber, rec.Carrier, rec.Score, rec.PhoneNumber)
		if err != nil {
			log.Printf("Error executing database statement: %v", err)
		}

		json.NewEncoder(w).Encode(rec)
	}

}

func deleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete record")
}

func listAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list all records")
}

func deleteAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete all records")
}
