package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"web_lab/internal/config"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewStorage(cfg config.PostgresConfig) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SslMode))
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping postgres connection pool: %w", err)
	}

	return &Storage{
		pool: pool,
	}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}

func (s *Storage) QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return s.pool.QueryRow(ctx, query, args...)
}

func (s *Storage) QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return s.pool.Query(ctx, query, args...)
}

func (s *Storage) ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return s.pool.Exec(ctx, query, args...)
}

func (s *Storage) Tx(ctx context.Context) (pgx.Tx, error) {
	return s.pool.BeginTx(ctx, pgx.TxOptions{})
}
