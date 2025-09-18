package sjsu

import (
	"log"
	"time"
)

type Semester int

const (
	Fall Semester = iota
	WinterBreak
	Spring
	SummerBreak
)

type SchoolYear struct {
	StartYear int
	EndYear   int
}

type SchoolSemsterSchedule struct {
	FallStart   time.Time
	FallEnd     time.Time
	WinterStart time.Time
	WinterEnd   time.Time
	SpringStart time.Time
	SpringEnd   time.Time
	SummerStart time.Time
	SummerEnd   time.Time
}

func fall2025fall2026() SchoolSemsterSchedule {
	timezone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	return SchoolSemsterSchedule{
		FallStart:   time.Date(2025, time.August, 20, 0, 0, 0, 0, timezone),
		FallEnd:     time.Date(2025, time.December, 26, 0, 0, 0, 0, timezone),
		WinterStart: time.Date(2025, time.December, 26, 0, 0, 0, 0, timezone),
		WinterEnd:   time.Date(2026, time.January, 20, 0, 0, 0, 0, timezone),
		SpringStart: time.Date(2026, time.January, 20, 0, 0, 0, 0, timezone),
		SpringEnd:   time.Date(2026, time.May, 22, 0, 0, 0, 0, timezone),
		SummerStart: time.Date(2026, time.May, 22, 0, 0, 0, 0, timezone),
		SummerEnd:   time.Date(2026, time.August, 19, 0, 0, 0, 0, timezone),
	}
}

var SchoolYearSchedule = map[SchoolYear]SchoolSemsterSchedule{
	{StartYear: 2025, EndYear: 2026}: fall2025fall2026(),
}

func determineSchoolYear(timestamp time.Time) SchoolYear {
	sy := SchoolYear{}
	switch {
	case (timestamp.After(fall2025fall2026().FallStart) && timestamp.Before(fall2025fall2026().SummerEnd)):
		sy.StartYear = 2025
		sy.EndYear = 2026
	}
	return sy
}

func SchoolSemester(timestamp time.Time) Semester {
	schoolYear := determineSchoolYear(timestamp)
	schedule, ok := SchoolYearSchedule[schoolYear]
	var semester Semester
	if !ok {
		log.Printf("Schedule for this schoolyear is not found: Fall %d, Spring %d. Returning Fall by default", schoolYear.StartYear, schoolYear.EndYear)
		return semester

	}
	switch {
	case (timestamp.After(schedule.FallStart) && timestamp.Before(schedule.FallEnd)):
		semester = Fall
	case (timestamp.After(schedule.WinterStart) && timestamp.Before(schedule.WinterEnd)):
		semester = WinterBreak
	case (timestamp.After(schedule.SpringStart) && timestamp.Before(schedule.SpringEnd)):
		semester = Spring
	case (timestamp.After(schedule.SummerStart) && timestamp.Before(schedule.SummerEnd)):
		semester = SummerBreak
	}
	return semester
}
