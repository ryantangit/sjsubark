package sjsucalendar

import (
	"log"
	"time"
)

// This file keeps track of instances where school campus would be closed.
// Note campus closed does not mean the same as winter breaks or summer breaks as the campus would still be opened during those times.
// The information is extracted from SJSU's academic calendar
// Example link: https://www.sjsu.edu/provost/docs/2025-26%20Calendar%20revised%207-11-25.pdf

type CloseCampusInstance struct {
	Reason     string //Most likely holidays, but whatever is placed on the calendar is there.
	StartMonth time.Month
	StartDay   int // The time it starts is at midnight
	EndMonth   time.Month
	EndDay     int // The time it ends is at midnight
}

// Hard Coded Campus Close Dates Starts Here
// -----------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------

var year2025Instances = []CloseCampusInstance{
	{
		Reason:     "Independence Day",
		StartMonth: time.July,
		StartDay:   4,
		EndMonth:   time.July,
		EndDay:     5,
	},
	{
		Reason:     "Labor Day",
		StartMonth: time.September,
		StartDay:   1,
		EndMonth:   time.September,
		EndDay:     2,
	},
	{
		Reason:     "Veteran's Day",
		StartMonth: time.November,
		StartDay:   11,
		EndMonth:   time.November,
		EndDay:     12,
	},
	{
		Reason:     "Thanksgiving",
		StartMonth: time.November,
		StartDay:   26,
		EndMonth:   time.November,
		EndDay:     28,
	},
	{
		Reason:     "Christmas",
		StartMonth: time.December,
		StartDay:   25,
		EndMonth:   time.December,
		EndDay:     26,
	},
}

var year2026Instances = []CloseCampusInstance{
	{
		Reason:     "New Year's Day",
		StartMonth: time.January,
		StartDay:   1,
		EndMonth:   time.January,
		EndDay:     2,
	},
	{
		Reason:     "Dr. Martin Luther King Jr's Day",
		StartMonth: time.January,
		StartDay:   19,
		EndMonth:   time.January,
		EndDay:     20,
	},
	{
		Reason:     "Cesar Chavez Day",
		StartMonth: time.March,
		StartDay:   31,
		EndMonth:   time.April,
		EndDay:     1,
	},
	{
		Reason:     "Memorial Day",
		StartMonth: time.May,
		StartDay:   25,
		EndMonth:   time.May,
		EndDay:     26,
	},
	{
		Reason:     "Juneteenth",
		StartMonth: time.June,
		StartDay:   19,
		EndMonth:   time.June,
		EndDay:     20,
	},
}

// Hard Coded Campus Close Dates End Here
// -----------------------------------------------------------------------------------------
// -----------------------------------------------------------------------------------------

var YeartoCloseCampusMap = map[int][]CloseCampusInstance{
	2025: year2025Instances,
	2026: year2026Instances,
}

// If no information is found in regards to the time, the default answer will always be No.
func isCampusClosed(timestamp time.Time) bool {
	year := timestamp.Year()
	instances, ok := YeartoCloseCampusMap[year]
	if !ok {
		log.Printf("Closed Campus Mapping for year %d was not found. Defaulting isCampusClosed inqury to false.", year)
		return false
	}
	timezone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	for _, instance := range instances {
		startTime := time.Date(year, instance.StartMonth, instance.StartDay, 0, 0, 0, 0, timezone)
		endTime := time.Date(year, instance.EndMonth, instance.EndDay, 0, 0, 0, 0, timezone)
		if startTime.Before(timestamp) && endTime.After(timestamp) {
			return true
		}
	}
	return false
}
