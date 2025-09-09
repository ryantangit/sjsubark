package transform

import (
	"log"
	"time"
)

type SchoolCalendar struct {
	StartYear int
	EndYear   int
	timezone  *time.Location
}

func (sc SchoolCalendar) StartofFall() time.Time {
	day := 1
	switch sc.StartYear {
	case 2025:
		day = 20
	default:
		log.Fatal("Unsupported FirstDayOfFall")
	}
	return time.Date(sc.StartYear, time.August, day, 0, 0, 0, 0, sc.timezone)
}

func (sc SchoolCalendar) EndofFall() time.Time {
	day := 1
	switch sc.StartYear {
	case 2025:
		day = 19
	default:
		log.Fatal("Unsupported LastDayofFall")
	}
	return time.Date(sc.StartYear, time.December, day, 0, 0, 0, 0, sc.timezone)
}

func (sc SchoolCalendar) StartofSpring() time.Time {
	day := 1
	switch sc.EndYear {
	case 2026:
		day = 20
	default:
		log.Fatal("Unsupported FirstDayOfSpring")
	}
	return time.Date(sc.EndYear, time.January, day, 0, 0, 0, 0, sc.timezone)
}

func (sc SchoolCalendar) EndofSpring() time.Time {
	day := 1
	switch sc.EndYear {
	case 2026:
		day = 22
	default:
		log.Fatal("Unsupported FirstDayOfSpring")
	}
	return time.Date(sc.EndYear, time.May, day, 0, 0, 0, 0, sc.timezone)
}
