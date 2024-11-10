package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const tableMovies = "movies"

var (
	ErrMovieAlreadyExists = errors.New("movie already exists")
	ErrMovieNotFound      = errors.New("movie nor found")
)

type Movie struct {
	ID          int64
	Title       string
	Description string
	Image       string
}

type MoviesOpts struct {
	Search     *string
	Pagination Pagination
}

// Pagination - параметры пагинации
type Pagination struct {
	Limit  *int64
	Offset *int64
}

// CreateMovie добавляет новую запись в таблицу movies
func (s *PG) CreateMovie(ctx context.Context, movie Movie) error {
	query, args, err := buildCreateMovieQuery(movie)
	if err != nil {
		return fmt.Errorf("[movies_pg] failed to build create movie query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.SQLState() == uniqueViolationCode {
			return ErrMovieAlreadyExists
		}

		return fmt.Errorf("[movies_pg] failed to create new movie: %w", err)
	}

	return nil
}

// UpdateMovie обновляет запись в таблице movies
func (s *PG) UpdateMovie(ctx context.Context, movie Movie) error {
	query, args, err := buildUpdateMovieQuery(movie)
	if err != nil {
		return fmt.Errorf("[movies_pg] failed to build update movie query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("[movies_pg] failed to update movie: %w", err)
	}

	return nil
}

// DeleteMovie удаляет запись из таблицы movies
func (s *PG) DeleteMovie(ctx context.Context, id int64) error {
	query, args, err := buildDeleteMovieQuery(id)
	if err != nil {
		return fmt.Errorf("[movies_pg] failed to build delete movie query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("[movies_pg] failed to delete movie: %w", err)
	}

	return nil
}

// GetMovies получает все фильмы из таблицы movies
func (s *PG) GetMovies(ctx context.Context, opts MoviesOpts) ([]Movie, int64, error) {
	query, args, err := buildGetMoviesQuery(opts)
	if err != nil {
		return nil, 0, fmt.Errorf("[movies_pg] failed to build get movies query: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("[movies_pg] failed to get movies: %w", err)
	}

	defer rows.Close()

	movies := make([]Movie, 0)
	var count int64

	for rows.Next() {
		var movie Movie

		err = rows.Scan(&count, &movie.ID, &movie.Title, &movie.Description, &movie.Image)
		if err != nil {
			return nil, 0, fmt.Errorf("[movies_pg] failed to scan data into struct: %w", err)
		}

		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("[movies_pg] failed to get movies: %w", err)
	}

	return movies, count, nil
}

// GetMovie получает фильм по id из таблицы movies
func (s *PG) GetMovie(ctx context.Context, id int64) (Movie, error) {
	query, args, err := buildGetMovieQuery(id)
	if err != nil {
		return Movie{}, fmt.Errorf("[movies_pg] failed to build get movie query: %w", err)
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var movie Movie

	err = row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Image)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Movie{}, ErrMovieNotFound
		}
		return Movie{}, fmt.Errorf("[movies_pg] failed to scan movie: %w", err)
	}

	return movie, nil
}
