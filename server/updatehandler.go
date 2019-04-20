package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func updateRecordHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Fprintf(w, "this api call only accepts application/json requests")
	}

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

	// we don't really care about these values being put into the struct
	// at this point. But Scan requires the same number of arguments
	// as number of columns returned.
	err = recordRow.Scan(&rec.PhoneNumber, &rec.Carrier, &rec.Score)
	if err == sql.ErrNoRows {
		log.Printf("Phone number %v not found.", phoneNumber)

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

		prettyJSON, err := json.MarshalIndent(rec, "", "    ")
		if err != nil {
			log.Printf("Error marshal indent: %v", err)
		}
		fmt.Fprintf(w, string(prettyJSON))
	}

}
