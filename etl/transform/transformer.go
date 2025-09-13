package transform

import (
	"log"
	"time"

	"github.com/ryantangit/sjsubark/etl/extract"
	"github.com/ryantangit/sjsubark/etl/sjsu"
)

// The transformer is mainly responsible for mining time information from the timestamps.
type Transformer interface {
	TransformRecord(gr extract.GarageRecord) (cgr CompleteGarageRecord)
}

type CompleteGarageRecord struct {
	Name      string
	Addr      string
	Fullness  int
	Second    int
	Minute    int
	Hour      int
	Month     time.Month
	Year      int
	Weekday   time.Weekday
	IsWeeknd  bool
	IsHoliday bool
	Semester  sjsu.Semester
}



type DefaultTransformer struct {
}

func (t DefaultTransformer) TransformRecord(gr extract.GarageRecord) CompleteGarageRecord {
	record := CompleteGarageRecord{Name: gr.Name, Addr: gr.Addr, Fullness: gr.Fullness}
	timeConverter := StdTimeConverter{time: gr.Timestamp}
	record.Second = timeConverter.Second()
	record.Minute = timeConverter.Minute()
	record.Hour = timeConverter.Hour()
	record.Month = timeConverter.Month()
	record.Year = timeConverter.Year()
	record.Weekday = timeConverter.Weekday()
	record.IsWeeknd = timeConverter.IsWeekend()
	record.IsHoliday = timeConverter.IsHoliday()
	record.Semester = timeConverter.ToSemster()
	return record
}

type TimeConverter interface {
	Second() int
	Minute() int
	Hour() int
	Month() time.Month
	Year() int
	Weekday() time.Weekday
	IsWeekend() bool
	IsHoliday() bool
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

func (t StdTimeConverter) IsHoliday() bool {
	return false
}

func (t StdTimeConverter) ToSemster() sjsu.Semester {
	return -1
}
