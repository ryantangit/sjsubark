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
	EndYear int
}

type SchoolSemsterSchedule struct{
	FallStart time.Time
	FallEnd time.Time
	WinterStart time.Time
	WinterEnd time.Time
	SpringStart time.Time
	SpringEnd time.Time
	SummerStart time.Time
	SummerEnd time.Time
}

func fall2025fall2026() SchoolSemsterSchedule {
	timezone, err:= time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Fatal(err)
	}
	return SchoolSemsterSchedule {
		FallStart: time.Date(2025, time.August, 20, 0, 0, 0, 0, timezone),
		FallEnd: time.Date(2025, time.December, 26, 0, 0, 0, 0, timezone),
		WinterStart: time.Date(2025, time.December, 26, 0, 0, 0, 0, timezone),
		WinterEnd: time.Date(2026, time.January, 20, 0, 0, 0, 0, timezone),
		SpringStart: time.Date(2026, time.January, 20, 0, 0, 0, 0, timezone),
		SpringEnd: time.Date(2026, time.May, 22, 0, 0, 0, 0, timezone),
		SummerStart: time.Date(2026, time.May, 22, 0, 0, 0, 0, timezone),
		SummerEnd: time.Date(2026, time.August, 19, 0, 0, 0, 0, timezone),
	}
}

var SchoolYearSchedule = map[SchoolYear]SchoolSemsterSchedule{
	SchoolYear{StartYear: 2025, EndYear: 2026} : fall2025fall2026(),
}


//TODO: Based off timestamp figure out what semester this is 
func SchoolSemester(timestamp time.Time) Semester {
	return -1 
}
