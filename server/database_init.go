package main


import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)



func createdatabase() {

	database, err := sql.Open("sqlite3", "pin_records.db")
	if err != nil {
		log.Fatalf("sql.Open error: %v", err)
	}

	sqlCreate, err := ioutil.ReadFile("records_table_creation.sql")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	// if tables haven't been created, then create them

	_, err = database.Exec(string(sqlCreate))
	if err != nil {
		log.Fatalf("database.Exec table creation error: %v", err)
	}

	database.Close()

}