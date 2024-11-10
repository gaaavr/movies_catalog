package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"web_lab/internal/config"
)

type StorageStd struct {
	db *sql.DB
}

func NewStorageStd(cfg config.PostgresConfig) (*StorageStd, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SslMode))
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping postgres connection: %w", err)
	}

	db.SetConnMaxIdleTime(0)
	db.SetMaxOpenConns(10)
	return &StorageStd{
		db: db,
	}, nil
}

func (s *StorageStd) Close() {
	s.db.Close()
}

func (s *StorageStd) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.db.QueryRow(query, args...)
}

func (s *StorageStd) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

func (s *StorageStd) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s *StorageStd) Tx(ctx context.Context) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, &sql.TxOptions{})
}
