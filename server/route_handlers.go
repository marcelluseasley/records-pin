package main


import(
	"fmt"
	"log"
	"net/http"
)


func createRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Header.Get("Content-type"))
	log.Println(r.Header.Get("Content-Type"))
}


func readRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "read record")
}

func updateRecordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update record")
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
