package extract

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryantangit/sjsubark/etl/config"
)

// I have a thinkpad also running a previous version of this entire pipeline.
// This is a CSV Extractor that transforms a legacy record format to the current new one.
// Serves no additional purpose beyond this.

type LegacyCSVExtractor struct {
	Filepath string
}

func (ex LegacyCSVExtractor) FetchRecords() (gr []GarageRecord) {
	f, err := os.Open(ex.Filepath)
	if err != nil {
		log.Panic(err)
	}
	s := bufio.NewScanner(f)
	result := []GarageRecord{}

	for s.Scan() {
		line := s.Text()
		line = strings.ReplaceAll(line, " ", "")
		fmt.Println(line)
		parts := strings.Split(line, ",")
		fullness, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		timestamp := parts[2]
		const layout = "2006-01-02::15:04:05"
		laLocation := config.Timezone()
		time, err := time.ParseInLocation(layout, timestamp, laLocation)
		if err != nil {
			panic(err)
		}
		record := GarageRecord{Name: parts[0], Fullness: fullness, Timestamp: time}
		fmt.Printf("%+v\n", record)
		result = append(result, record)
	}
	return result
}
