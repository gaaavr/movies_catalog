package models

import (
	"context"
	"fmt"

	"github.com/lib/pq"
)

// CreateMovie добавляет новую запись в таблицу movies
func (s *PGStd) CreateMovie(ctx context.Context, movie Movie) error {
	query, args, err := buildCreateMovieQuery(movie)
	if err != nil {
		return fmt.Errorf("[movies_pg] failed to build create movie query: %w", err)
	}

	_, err = s.db.Exec(query, args...)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.SQLState() == uniqueViolationCode {
			return ErrMovieAlreadyExists
		}

		return fmt.Errorf("[movies_pg] failed to create new movie: %w", err)
	}

	return nil
}

// GetMovie получает фильм по id из таблицы movies
func (s *PGStd) GetMovie(ctx context.Context, id int64) (Movie, error) {
	query, args, err := buildGetMovieQuery(id)
	if err != nil {
		return Movie{}, fmt.Errorf("[movies_pg] failed to build get movie query: %w", err)
	}

	row := s.db.QueryRow(query, args...)

	var movie Movie

	err = row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Image)
	if err != nil {
		return Movie{}, fmt.Errorf("[movies_pg] failed to scan movie: %w", err)
	}

	return movie, nil
}
