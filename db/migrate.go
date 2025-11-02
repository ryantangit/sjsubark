package db

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryantangit/sjsubark/etl/config"
	"github.com/ryantangit/sjsubark/etl/extract"
	"github.com/ryantangit/sjsubark/etl/loader"
	"github.com/ryantangit/sjsubark/etl/sjsu"
	"github.com/ryantangit/sjsubark/etl/transform"
)

// Copies everything from CSV into Postgres
func migrateCSVtoPSQL() {
	f, err := os.Open(config.CSVPath())
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	result := []extract.GarageRecord{}
	//SouthGarage, 40, 2025-08-26 04:10:01 +0000 UTC, 21, 10, 1, 2025, 8, 25, 1, false, false
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, ",")
		for i, s := range parts {
			parts[i] = strings.TrimLeft(s, " ")
		}
		fmt.Println(line)
		fullness, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		timestamp := parts[2]
		const layout = "2006-01-02 15:04:05 -0700 MST"
		laLocation := config.Timezone()
		time, err := time.ParseInLocation(layout, timestamp, laLocation)
		if err != nil {
			panic(err)
		}
		record := extract.GarageRecord{Name: parts[0], Fullness: fullness, Timestamp: time}
		fmt.Printf("%+v\n", record)
		result = append(result, record)
	}

	sjsu := sjsu.SanJoseCampus{YeartoCloseCampusMap: make(map[int][]sjsu.CloseCampusInstance)}
	sjsu.SanJoseCampusInit(config.CampusClosePath())
	cgr := []transform.CompleteGarageRecord{}
	for _, r := range result {
		cgr = append(cgr, transform.TransformRecord(r, sjsu))
	}

	psql := loader.NewPostgresLoader(config.PostgresURL())
	if psql != nil {
		defer psql.Close(context.Background())
		for _, r := range cgr {
			psql.Upload(r)
		}
	}

}
