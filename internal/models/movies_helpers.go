package models

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// Функция для формирования запроса на создание нового фильма в каталоге
func buildCreateMovieQuery(movie Movie) (string, []interface{}, error) {
	return qb.Insert(tableMovies).Columns("title", "description", "image").
		Values(movie.Title, movie.Description, movie.Image).ToSql()
}

// Функция для формирования запроса на обновление фильма в каталоге
func buildUpdateMovieQuery(movie Movie) (string, []interface{}, error) {
	return qb.Update(tableMovies).Set("title", movie.Title).
		Set("description", movie.Description).
		Set("image", movie.Image).Where(sq.Eq{"id": movie.ID}).
		Set("updated_at", time.Now()).ToSql()
}

// Функция для формирования запроса на удаление фильма из каталога
func buildDeleteMovieQuery(id int64) (string, []interface{}, error) {
	return qb.Delete(tableMovies).Where(sq.Eq{"id": id}).ToSql()
}

// Функция для формирования запроса для получения списка фильмов
func buildGetMoviesQuery(opts MoviesOpts) (string, []interface{}, error) {
	query := qb.Select("count(*) OVER()", "id", "title", "description", "image").
		From(tableMovies)

	conditions := sq.And{}

	if opts.Search != nil {
		searchParam := fmt.Sprintf("%%%s%%", *opts.Search)

		conditions = append(conditions,
			sq.Or{
				sq.Like{"title": searchParam},
				sq.Like{"description": searchParam},
			})
	}

	if len(conditions) > 0 {
		query = query.Where(conditions)
	}

	if opts.Pagination.Limit != nil {
		query = query.Limit(uint64(*opts.Pagination.Limit))
	}
	if opts.Pagination.Offset != nil {
		query = query.Offset(uint64(*opts.Pagination.Offset))
	}

	query = query.OrderBy("updated_at DESC")

	return query.ToSql()
}

// Функция для формирования запроса для получения фильма
func buildGetMovieQuery(id int64) (string, []interface{}, error) {
	query := qb.Select("id", "title", "description", "image").
		From(tableMovies).Where(sq.Eq{"id": id})

	return query.ToSql()
}
