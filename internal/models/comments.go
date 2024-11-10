package models

import (
	"context"
	"fmt"
	"time"
)

const tableComments = "comments"

type Comment struct {
	ID        int64
	Content   string
	UserID    int64
	Username  string
	MovieID   int64
	CreatedAt time.Time
}

// CreateComment добавляет новую запись в таблицу comments
func (s *PG) CreateComment(ctx context.Context, comment Comment) error {
	query, args, err := buildCreateCommentQuery(comment)
	if err != nil {
		return fmt.Errorf("[comments_pg] failed to build create comment query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("[comments_pg] failed to create comment: %w", err)
	}

	return nil
}

// DeleteComment удаляет запись из таблицы comments
func (s *PG) DeleteComment(ctx context.Context, id int64) error {
	query, args, err := buildDeleteCommentQuery(id)
	if err != nil {
		return fmt.Errorf("[comments_pg] failed to build delete comment query: %w", err)
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("[comments_pg] failed to delete comment: %w", err)
	}

	return nil
}

// GetComments получает все фильмы из таблицы comments
func (s *PG) GetComments(ctx context.Context, movieID int64) ([]Comment, error) {
	query, args, err := buildGetCommentsQuery(movieID)
	if err != nil {
		return nil, fmt.Errorf("[comments_pg] failed to build get comments query: %w", err)
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("[comments_pg] failed to get commets: %w", err)
	}

	defer rows.Close()

	comments := make([]Comment, 0)

	for rows.Next() {
		var comment Comment

		err = rows.Scan(&comment.ID, &comment.Content, &comment.UserID,
			&comment.Username, &comment.MovieID, &comment.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("[comments_pg] failed to scan data into struct: %w", err)
		}

		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[comments_pg] failed to get comments: %w", err)
	}

	return comments, nil
}
