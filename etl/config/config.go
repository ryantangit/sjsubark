package config

import (
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
	return filepath.Join(home, ".sjsubarker", "etl", "campus_close.json")
}

func CSVPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".sjsubarker", "etl", "master.csv")
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
	return filepath.Join(home, ".sjsubarker", "etl", "webpages")
}

// This is where the web scrape will look for.
func WebpageUrl() string {
	const ParkingStatusUrl = "https://sjsuparkingstatus.sjsu.edu/"
	return ParkingStatusUrl
}
