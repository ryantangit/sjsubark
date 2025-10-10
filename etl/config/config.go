package config

import (
	"log"
	"os"
	"path/filepath"
)

func CSVLocation() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, ".sjsubarker", "etl", "master.csv")
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
