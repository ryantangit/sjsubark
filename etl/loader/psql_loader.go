package loader

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/ryantangit/sjsubark/etl/transform"
)

type PostgresLoader struct {
	conn *pgx.Conn
}

func NewPostgresLoader(postgresURL string) *PostgresLoader {
	conn, err := pgx.Connect(context.Background(), postgresURL)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &PostgresLoader{conn: conn}
}

func (pg *PostgresLoader) Close(ctx context.Context) error {
	return pg.conn.Close(ctx)
}

func (pg *PostgresLoader) Upload(cgr transform.CompleteGarageRecord) {
	tx, err := pg.conn.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback(context.Background())

	args := []any{
		cgr.Name,
		cgr.UTCTime,
		cgr.Second,
		cgr.Minute,
		cgr.Hour,
		cgr.Day,
		cgr.Month,
		cgr.Year,
		cgr.Weekday,
		cgr.IsWeekend,
		cgr.IsCampusClosed,
		cgr.Fullness,
	}
	_, err = tx.Exec(context.Background(), `INSERT INTO garage_fullness (name, utc_timestamp, second, minute, hour, day, month, year, weekday, is_weekend, is_campus_closed, fullness) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`, args...)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
