package datastore

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DataStore struct {
	dbpool *pgxpool.Pool
}

func NewDataStore(dbpool *pgxpool.Pool) *DataStore{
	return &DataStore{dbpool: dbpool}
}

type GarageStatus struct {
	Garage_name string
	Fullness int
	Utc_timestamp time.Time
}

func (ds *DataStore) GetLatestStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("/latest requested")
	ctx := r.Context()	
	query := "SELECT name, fullness, utc_timestamp FROM garage_fullness ORDER BY utc_timestamp DESC LIMIT 4;"	
	conn, err := ds.dbpool.Acquire(r.Context())
	defer conn.Release()
	if err != nil {
		log.Println(err)
	}
	tx, err := conn.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		log.Println(err)
	}
	rows, err := tx.Query(ctx, query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	statusList := []GarageStatus{}
	//Retrieve the utc_timestamp, name, fullness is sufficient
	for rows.Next() {
		var garageStatus GarageStatus
		err := rows.Scan(&garageStatus.Garage_name, &garageStatus.Fullness, &garageStatus.Utc_timestamp)
		if err != nil {
			log.Println(err)
		}
		statusList = append(statusList, garageStatus)
	}
	tx.Commit(ctx)
	jsonData, err := json.Marshal(statusList)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

