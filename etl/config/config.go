package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CampusClosePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".sjsubark", "etl", "campus_close.json")
}

func CSVPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".sjsubark", "etl", "master.csv")
}

// For PostgresTable, configure the following environment variables
func PostgresURL() string {
	password := os.Getenv("SJSUBARK_PSQL_PASSWORD")
	user := os.Getenv("SJSUBARK_PSQL_USER")
	db := os.Getenv("SJSUBARK_PSQL_DB")
	port := os.Getenv("SJSUBARK_PSQL_PORT")
	host := os.Getenv("SJSUBARK_PSQL_HOST")

	//Ex) "postgresql://postgres:password@localhost:5432/db"
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, db)
	return connStr
}

// School time zone, really shouldn't change unless the tectonic plates are moving
func Timezone() *time.Location {
	timezone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	return timezone
}

// This will be where snapshots of the scrapes will be located.
func WebpageDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".sjsubark", "etl", "webpages")
}

// This is where the web scrape will look for.
func WebpageUrl() string {
	const ParkingStatusUrl = "https://sjsuparkingstatus.sjsu.edu/"
	return ParkingStatusUrl
}
