package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func deleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rec := Record{}
	phoneNumber := chi.URLParam(r, "phoneNumber")

	database, err := sql.Open("sqlite3", "pin_records.db")
	defer database.Close()
	if err != nil {
		log.Printf("Error opening database. Cannot update record: %v", err)
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
		statement, err := database.Prepare("DELETE from records WHERE phone_number=?;")
		if err != nil {
			log.Printf("Error preparing database: %v", err)
		}

		_, err = statement.Exec(phoneNumber)
		if err != nil {
			log.Printf("Error executing database statement: %v", err)
		} else {
			log.Println("Deleted record")
			w.Write([]byte(`{"status" : "record deleted successfully"}`))
		}

	}

}
