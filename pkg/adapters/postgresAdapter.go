package adapters

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	DB *sql.DB
}

func NewPostgresAdapter(connStr string) (*PostgresAdapter, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresAdapter{DB: db}, nil
}

func (p *PostgresAdapter) PlaceholderAdapter(ctx context.Context, placeholder string) error {
	//TODO implement me
	panic("implement me")
}
