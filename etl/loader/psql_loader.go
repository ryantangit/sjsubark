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
		cgr.Fullness,
	}
	query := `
	INSERT INTO garage_fullness (utc_timestamp, fullness, garage_id)
	SELECT $2, $3, gi.garage_id 
	FROM garage_info gi 
	WHERE gi.garage_name = $1;
	`
	_, err = tx.Exec(context.Background(), query, args...)
	
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
