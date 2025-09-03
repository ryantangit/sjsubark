package extract

import "time"

// The GarageExtractor is expected to find all parking garages and their related fullness 0 - 100%.
// The timestamp will be extracted from the moment the record is generated.

type GarageExtractor interface {
	FetchRecord() GarageRecord
}

type GarageRecord struct {
	timestamp time.Time
	garages   []Garage
}

type Garage struct {
	name     string
	fullness int
	addr     string
}
