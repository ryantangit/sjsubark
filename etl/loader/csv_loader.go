package loader

import (
	"fmt"
	"log"
	"os"

	"github.com/ryantangit/sjsubark/etl/transform"
)

type CSVLoader struct {
	Filepath string
}

func NewCSVLoader(filepath string) CSVLoader {
	return CSVLoader{Filepath: filepath}
}

func CSVRecord(cgr transform.CompleteGarageRecord) string {
	return fmt.Sprintf("%s, %d, %s, %d, %d, %d, %d, %d, %d, %d, %t, %t\n", cgr.Name, cgr.Fullness, cgr.UTCTime.String(), cgr.Hour, cgr.Minute, cgr.Second, cgr.Year, cgr.Month, cgr.Day, cgr.Weekday, cgr.IsWeekend, cgr.IsCampusClosed)
}

func (csv CSVLoader) Upload(cgr transform.CompleteGarageRecord) {
	f, err := os.OpenFile(csv.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(CSVRecord(cgr))); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
