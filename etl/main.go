package main

import (
	"github.com/ryantangit/sjsubark/etl/config"
	"github.com/ryantangit/sjsubark/etl/extract"
	"github.com/ryantangit/sjsubark/etl/loader"
	"github.com/ryantangit/sjsubark/etl/sjsu"
	"github.com/ryantangit/sjsubark/etl/transform"
)

func main() {

	webX := extract.NewWebpageExtractor(config.WebpageUrl(), config.WebpageDir())
	gr := webX.FetchRecords()
	sjsu := sjsu.SanJoseCampus{YeartoCloseCampusMap: make(map[int][]sjsu.CloseCampusInstance)}
	sjsu.SanJoseCampusInit(config.CampusClosePath())
	cgr := []transform.CompleteGarageRecord{}
	for _, r := range gr {
		cgr = append(cgr, transform.TransformRecord(r, sjsu))
	}
	csv := loader.NewCSVLoader(config.CSVPath())
	for _, r := range cgr {
		csv.Upload(r)
	}
}
