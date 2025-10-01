package main

import (
	"github.com/ryantangit/sjsubark/etl/extract"
	"github.com/ryantangit/sjsubark/etl/transform"
	"github.com/ryantangit/sjsubark/etl/loader"
)

func main() {
	webX := extract.NewWebpageExtractor()
	gr := webX.FetchRecords()
	cgr := []transform.CompleteGarageRecord{}
	for _, r := range gr {
		cgr = append(cgr, transform.TransformRecord(r))
	}
	csv := loader.NewCSVLoader("/home/ryantan/.sjsubark/master.csv")
	for _, r := range cgr {
		csv.Upload(r)
	}
}
