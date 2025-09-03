package transform

import (
	"github.com/ryantangit/sjsubark/extract"
)

// The transformer is mainly responsible for mining time information from the timestamps.

type Transformer interface {
	TransformRecord(extract.GarageRecord) CompleteGarageRecord
}

type CompleteGarageRecord struct {
	name      string
	addr      string
	fullness  int
	second    int
	minute    int
	hour      int
	month     Month
	year      int
	weekday   Weekday
	isWeeknd  bool
	isHoliday bool
	semester  Semester
}

type Weekday int

const (
	Mon Weekday = iota + 1
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

type Month int

const (
	Jan Month = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

type Semester int

const (
	Fall Semester = iota
	WinterBreak
	Spring
	SummerBreak
)

type TimestampTransformer interface {
	ToSecond() int
	ToMinute() int
	ToHour() int
	ToMonth() Month
	ToYear() int
	ToWeekday() Weekday
	IsWeekend() bool
	isHoliday() bool
	ToSemster() Semester
}
