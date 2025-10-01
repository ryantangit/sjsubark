package transform

import (
	"fmt"
	"time"

	"github.com/ryantangit/sjsubark/etl/extract"
	"github.com/ryantangit/sjsubark/etl/sjsu"
)

// The transformer is mainly responsible for mining time information from the timestamps.
type Transformer interface {
	TransformRecord(gr extract.GarageRecord) (cgr CompleteGarageRecord)
}

type CompleteGarageRecord struct {
	Name           string
	Addr           string
	Fullness       int
	Second         int
	Minute         int
	Hour           int
	Day            int
	Month          time.Month
	Year           int
	Weekday        time.Weekday
	IsWeekend      bool //Not the Singer
	IsCampusClosed bool
	Semester       sjsu.Semester
}

func (cgr CompleteGarageRecord) String() string {
	return fmt.Sprintf(" Name: %s\n Fullness: %d\n Month: %d \n Day: %d\n Year: %d\n Weekday: %s\n IsWeekend: %t\n IsCampusClosed: %t\n", cgr.Name, cgr.Fullness, cgr.Month, cgr.Day, cgr.Year, cgr.Weekday, cgr.IsWeekend, cgr.IsCampusClosed)
}

func (cgr CompleteGarageRecord) CSVRecord() string {
	return fmt.Sprintf("%s, %s, %d, %d, %d, %d, %d, %d, %d, %d, %t, %t, %d\n", cgr.Name, cgr.Addr, cgr.Fullness, cgr.Second, cgr.Minute, cgr.Hour, cgr.Day, cgr.Month, cgr.Year, cgr.Weekday, cgr.IsWeekend, cgr.IsCampusClosed, cgr.Semester)
}

func TransformRecord(gr extract.GarageRecord) CompleteGarageRecord {
	record := CompleteGarageRecord{Name: gr.Name, Addr: gr.Addr, Fullness: gr.Fullness}
	timeConverter := StdTimeConverter{time: gr.Timestamp}
	record.Second = timeConverter.Second()
	record.Minute = timeConverter.Minute()
	record.Hour = timeConverter.Hour()
	record.Day = timeConverter.Day()
	record.Month = timeConverter.Month()
	record.Year = timeConverter.Year()
	record.Weekday = timeConverter.Weekday()
	record.IsWeekend = timeConverter.IsWeekend()
	record.IsCampusClosed = timeConverter.IsCampusClosed()
	record.Semester = timeConverter.ToSemster()
	return record
}

type TimeConverter interface {
	Second() int
	Minute() int
	Hour() int
	Day() int
	Month() time.Month
	Year() int
	Weekday() time.Weekday
	IsWeekend() bool
	IsCampusClosed() bool
	ToSemster() sjsu.Semester
}

type StdTimeConverter struct {
	time time.Time
}

func (t StdTimeConverter) Second() int {
	return t.time.Second()
}

func (t StdTimeConverter) Minute() int {
	return t.time.Minute()
}

func (t StdTimeConverter) Hour() int {
	return t.time.Hour()
}

func (t StdTimeConverter) Day() int {
	return t.time.Day()
}

func (t StdTimeConverter) Month() time.Month {
	return t.time.Month()
}

func (t StdTimeConverter) Year() int {
	return t.time.Year()
}

func (t StdTimeConverter) Weekday() time.Weekday {
	return t.time.Weekday()
}

func (t StdTimeConverter) IsWeekend() bool {
	weekday := t.time.Weekday()
	return (weekday == time.Friday || weekday == time.Saturday || weekday == time.Sunday)
}

func (t StdTimeConverter) IsCampusClosed() bool {
	return sjsu.IsCampusClosed(t.time)
}

func (t StdTimeConverter) ToSemster() sjsu.Semester {
	return sjsu.SchoolSemester(t.time)
}
