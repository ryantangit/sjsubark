package extractor

import "time"

// The GarageExtractor is expected to find all parking garages and their related fullness 0 - 100%.
// The timestamp will be extracted from the moment the record is generated.

type GarageExtractor interface {
	FetchRecords() (gr []GarageRecord)
}

type GarageRecord struct {
	Timestamp time.Time
	Name      string
	Fullness  int
	Addr      string
}
