package main

import (
	"github.com/ryantangit/sjsubark/etl/extract"
	"github.com/ryantangit/sjsubark/etl/transform"
)

func main() {

	webX := extract.NewWebpageExtractor()
	gr := webX.FetchRecords()
	cgr := []transform.CompleteGarageRecord{}
	for _, r := range gr {
		cgr = append(cgr, transform.TransformRecord(r))
	}
	for _, r := range cgr {
		println(r.String())
	}
}
