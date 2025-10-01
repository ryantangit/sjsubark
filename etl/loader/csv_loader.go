package loader

import (
	"os"
	"log"

	"github.com/ryantangit/sjsubark/etl/transform"
)

type CSVLoader struct {
	Filepath string
}

func NewCSVLoader(filepath string) CSVLoader {
	return CSVLoader{Filepath: filepath}
}

func (csv CSVLoader) Upload (cgr transform.CompleteGarageRecord) {
	f, err := os.OpenFile(csv.Filepath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(cgr.CSVRecord())); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
