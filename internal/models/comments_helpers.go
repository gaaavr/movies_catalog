package models

import sq "github.com/Masterminds/squirrel"

// Функция для формирования запроса на создание нового комментария
func buildCreateCommentQuery(comment Comment) (string, []interface{}, error) {
	return qb.Insert(tableComments).Columns("content", "user_id", "movie_id").
		Values(comment.Content, comment.UserID, comment.MovieID).ToSql()
}

// Функция для формирования запроса для получения комментариев
func buildGetCommentsQuery(movieID int64) (string, []interface{}, error) {
	query := qb.Select("c.id", "c.content", "c.user_id", "u.username",
		"c.movie_id", "c.created_at").
		From("comments c").
		Join("users u on u.id = c.user_id").
		Join("movies m on m.id = c.movie_id").
		Where(sq.Eq{"c.movie_id": movieID}).
		OrderBy("created_at desc")

	return query.ToSql()
}

// Функция для формирования запроса на удаление комментарий
func buildDeleteCommentQuery(id int64) (string, []interface{}, error) {
	return qb.Delete(tableComments).Where(sq.Eq{"id": id}).ToSql()
}
