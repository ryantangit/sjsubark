package sjsu

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/ryantangit/sjsubark/config"
)

// This file keeps track of instances where school campus would be closed.
// Note campus closed does not mean the same as winter breaks or summer breaks as the campus would still be opened during those times.
// The information is extracted from SJSU's academic calendar into campus_close.json
// Example link: https://www.sjsu.edu/provost/docs/2025-26%20Calendar%20revised%207-11-25.pdf

type CloseCampusInstance struct {
	//Most likely holidays, but whatever is placed on the calendar is there.
	Reason     string
	StartMonth time.Month
	StartDay   int // The time it starts is at midnight
	EndMonth   time.Month
	EndDay     int // The time it ends is at midnight
}

type SanJoseCampus struct {
	YeartoCloseCampusMap map[int][]CloseCampusInstance
}

type SJSUJson struct {
	Year        int                   `json:"year"`
	CampusClose []CloseCampusInstance `json:"instances"`
}

func (sjsu SanJoseCampus) SanJoseCampusInit(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var sjsuJson []SJSUJson
	err = json.Unmarshal(data, &sjsuJson)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sjsuJson {
		sjsu.YeartoCloseCampusMap[entry.Year] = entry.CampusClose
	}
}

// If no information is found in regards to the time, the default answer will always be No.
func (sjsu SanJoseCampus) IsCampusClosed(timestamp time.Time) bool {
	year := timestamp.Year()
	instances, ok := sjsu.YeartoCloseCampusMap[year]
	if !ok {
		log.Printf("Closed Campus Mapping for year %d was not found. Defaulting isCampusClosed inqury to false.", year)
		return false
	}
	timezone := config.Timezone()
	for _, instance := range instances {
		startTime := time.Date(year, instance.StartMonth, instance.StartDay, 0, 0, 0, 0, timezone)
		endTime := time.Date(year, instance.EndMonth, instance.EndDay, 23, 59, 59, 0, timezone)
		if startTime.Before(timestamp) && endTime.After(timestamp) {
			return true
		}
	}
	return false
}
