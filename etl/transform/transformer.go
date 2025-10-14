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
	Fullness       int
	UTCTime		   string
	Second         int
	Minute         int
	Hour           int
	Day            int
	Month          time.Month
	Year           int
	Weekday        time.Weekday
	IsWeekend      bool
	IsCampusClosed bool
}

func (cgr CompleteGarageRecord) String() string {
	return fmt.Sprintf(" Name: %s\n Fullness: %d\n Month: %d \n Day: %d\n Year: %d\n Weekday: %s\n IsWeekend: %t\n", cgr.Name, cgr.Fullness, cgr.Month, cgr.Day, cgr.Year, cgr.Weekday, cgr.IsWeekend)
}

func (cgr CompleteGarageRecord) CSVRecord() string {
	return fmt.Sprintf("%s, %d, %s, %d, %d, %d, %d, %d, %d, %d, %t, %t\n", cgr.Name, cgr.Fullness, cgr.UTCTime, cgr.Hour, cgr.Minute, cgr.Second, cgr.Year, cgr.Month, cgr.Day, cgr.Weekday, cgr.IsWeekend, cgr.IsCampusClosed)
}

func TransformRecord(gr extract.GarageRecord, sjsu sjsu.SanJoseCampus) CompleteGarageRecord {
	record := CompleteGarageRecord{Name: gr.Name, Fullness: gr.Fullness}
	timeConverter := StdTimeConverter{time: gr.Timestamp}
	record.UTCTime = gr.Timestamp.UTC().String()
	record.Second = timeConverter.Second()
	record.Minute = timeConverter.Minute()
	record.Hour = timeConverter.Hour()
	record.Day = timeConverter.Day()
	record.Month = timeConverter.Month()
	record.Year = timeConverter.Year()
	record.Weekday = timeConverter.Weekday()
	record.IsWeekend = timeConverter.IsWeekend()
	record.IsCampusClosed = timeConverter.IsCampusClosed(sjsu)
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

func (t StdTimeConverter) IsCampusClosed(sjsu sjsu.SanJoseCampus) bool {
	return sjsu.IsCampusClosed(t.time)
}
