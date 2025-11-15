package main

import (
	"context"
	"fmt"
	"log"
	"net/http"


	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ryantangit/sjsubark/config"
	"github.com/ryantangit/sjsubark/db/datastore"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/ Requested")
	fmt.Fprintf(w, "Hello there.")
}

func main() {
	pool, err := pgxpool.New(context.Background(), config.PostgresURL())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	ds := datastore.NewDataStore(pool)
	log.Printf("Launching Server Instance")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/latest", ds.GetLatestStatus)
	http.ListenAndServe(":5431", nil)
}
