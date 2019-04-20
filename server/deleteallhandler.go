package main

import (
	"database/sql"
	"log"
	"net/http"
)

func deleteAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	database, err := sql.Open("sqlite3", "pin_records.db")
	defer database.Close()
	if err != nil {
		log.Printf("Error opening database. Cannot store created record: %v", err)
	}
	statement, err := database.Prepare("DELETE from records;")
	if err != nil {
		log.Printf("Error preparing database: %v", err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Printf("Error executing database statement: %v", err)
	} else {
		log.Println("Deleted all records")
		w.Write([]byte(`{"status" : "all records deleted successfully"}`))
	}
}
