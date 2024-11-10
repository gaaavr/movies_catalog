package view

import (
	"time"

	"web_lab/internal/models"
)

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	MovieID   int64  `json:"movie_id"`
	CreatedAt string `json:"created_at"`
}

func ConvertCommentsResponse(comments []models.Comment) CommentsResponse {
	result := make([]Comment, 0, len(comments))

	loc, _ := time.LoadLocation("Europe/Moscow")

	for _, comment := range comments {
		result = append(result, Comment{
			ID:        comment.ID,
			Content:   comment.Content,
			UserID:    comment.UserID,
			Username:  comment.Username,
			MovieID:   comment.MovieID,
			CreatedAt: comment.CreatedAt.In(loc).Format(time.RFC822),
		})
	}

	return CommentsResponse{
		Comments: result,
	}
}
