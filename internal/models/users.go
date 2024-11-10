package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

const tableUsers = "users"

var (
	// ErrUserAlreadyExists возвращается, когда в таблице users уже существует запись
	// с таким именем
	ErrUserAlreadyExists = errors.New(" already exists")
)

type User struct {
	ID        int
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
}

func (s *PG) CreateUser(ctx context.Context, user User) error {
	query, args, err := buildCreateUserQuery(user)
	if err != nil {
		return fmt.Errorf("[users_pg] failed to build create user query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.SQLState() == uniqueViolationCode {
			return ErrUserAlreadyExists
		}

		return fmt.Errorf("[users_pg] failed to create user: %w", err)
	}

	return nil
}

func (s *PG) UpdateUser(ctx context.Context, user User) error {
	query, args, err := buildUpdateUserQuery(user)
	if err != nil {
		return fmt.Errorf("[users_pg] failed to build update user query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("[users_pg] failed to update user: %w", err)
	}

	return nil
}

func (s *PG) GetUser(ctx context.Context, username, password string) (User, error) {
	query, args, err := buildGetUserQuery(username, password)
	if err != nil {
		return User{}, fmt.Errorf("[users_pg] failed to build get user query: %w", err)
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var u User

	err = row.Scan(&u.ID, &u.Username, &u.Role)
	if err != nil {
		return User{}, fmt.Errorf("[users_pg] failed to scan user: %w", err)
	}

	return u, nil
}

func (s *PG) GetUserByID(ctx context.Context, id int64) (User, error) {
	query, args, err := buildGetUserByIDQuery(id)
	if err != nil {
		return User{}, fmt.Errorf("[users_pg] failed to build get user query: %w", err)
	}

	row := s.db.QueryRowContext(ctx, query, args...)

	var u User

	err = row.Scan(&u.ID, &u.Username, &u.Role)
	if err != nil {
		return User{}, fmt.Errorf("[users_pg] failed to scan user: %w", err)
	}

	return u, nil
}
