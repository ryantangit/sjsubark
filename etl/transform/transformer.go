package transform

import (
	"fmt"
	"time"

	"github.com/ryantangit/sjsubark/config"
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
	UTCTime        time.Time
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

func TransformRecord(gr extract.GarageRecord, sjsu sjsu.SanJoseCampus) CompleteGarageRecord {
	record := CompleteGarageRecord{Name: gr.Name, Fullness: gr.Fullness}
	timeConverter := StdTimeConverter{time: gr.Timestamp}
	record.UTCTime = gr.Timestamp.UTC()
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
	return t.time.In(config.Timezone()).Second()
}

func (t StdTimeConverter) Minute() int {
	return t.time.In(config.Timezone()).Minute()
}

func (t StdTimeConverter) Hour() int {
	return t.time.In(config.Timezone()).Hour()
}

func (t StdTimeConverter) Day() int {
	return t.time.In(config.Timezone()).Day()
}

func (t StdTimeConverter) Month() time.Month {
	return t.time.In(config.Timezone()).Month()
}

func (t StdTimeConverter) Year() int {
	return t.time.In(config.Timezone()).Year()
}

func (t StdTimeConverter) Weekday() time.Weekday {
	return t.time.In(config.Timezone()).Weekday()
}

func (t StdTimeConverter) IsWeekend() bool {
	weekday := t.time.In(config.Timezone()).Weekday()
	return (weekday == time.Saturday || weekday == time.Sunday)
}

func (t StdTimeConverter) IsCampusClosed(sjsu sjsu.SanJoseCampus) bool {
	return sjsu.IsCampusClosed(t.time.In(config.Timezone()))
}
