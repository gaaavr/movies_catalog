package models

import sq "github.com/Masterminds/squirrel"

// Функция для формирования запроса на создание нового пользователя
func buildCreateUserQuery(user User) (string, []interface{}, error) {
	return qb.Insert(tableUsers).Columns("username", "password", "role").
		Values(user.Username, user.Password, user.Role).ToSql()
}

// Функция для формирования запроса для получения информации о пользователе
func buildGetUserQuery(username, paswword string) (string, []interface{}, error) {
	return qb.
		Select(
			"id", "username", "role").
		From(tableUsers).
		Where(sq.And{sq.Eq{"username": username}, sq.Eq{"password": paswword}}).
		ToSql()
}

// Функция для формирования запроса для получения информации о пользователе по id
func buildGetUserByIDQuery(userID int64) (string, []interface{}, error) {
	return qb.
		Select(
			"id", "username", "role").
		From(tableUsers).
		Where(sq.And{sq.Eq{"id": userID}}).
		ToSql()
}
