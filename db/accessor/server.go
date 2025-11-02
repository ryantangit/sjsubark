package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/ Requested")
	fmt.Fprintf(w, "Hello there.")
}

//TODO:
//1. Establish ConnectionPool with Postgres
//2. Handle for most recent garage data.

func main() {
	log.Printf("Launching Server Instance")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":5431", nil)
}
