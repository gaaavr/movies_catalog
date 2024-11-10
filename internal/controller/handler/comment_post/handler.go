//go:generate minimock -i .commentCreator -s _mock.go -o ./mock -g
package comment_post

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"web_lab/internal/models"
	"web_lab/internal/view"
)

const (
	userIDKey = "userID"
)

type commentCreator interface {
	CreateComment(ctx context.Context, comment models.Comment) error
}

type Handler struct {
	commentCreator commentCreator
}

type createCommentRequest struct {
	Content string `json:"content"`
	MovieID int64  `json:"movie_id"`
}

func New(commentCreator commentCreator) *Handler {
	return &Handler{
		commentCreator: commentCreator,
	}
}

// Handle godoc
// @Summary create comment
// @Tags comment
// @Description create comment for movie
// @ID comment-post
// @Param Authorization header string true "jwt token"
// @Param comment body createCommentRequest true "comment for create"
// @Success 200 "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 401 {object} view.ErrorResponse "unauthorized"
// @Failure 403 {object} view.ErrorResponse "forbidden"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies/comments [post]
func (h *Handler) Handle(c *gin.Context) {
	ctx := context.Background()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	defer c.Request.Body.Close()

	var commentReq createCommentRequest

	err = json.Unmarshal(data, &commentReq)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	if commentReq.Content == "" {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(errors.New("content is empty")),
		)

		return
	}

	uID, err := getUserID(c)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	err = h.commentCreator.CreateComment(ctx, models.Comment{
		Content: commentReq.Content,
		MovieID: commentReq.MovieID,
		UserID:  uID,
	})
	if err != nil {
		if errors.Is(err, models.ErrMovieAlreadyExists) {
			c.JSON(
				http.StatusBadRequest,
				view.NewErrorResponse(err),
			)

			return
		}

		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetMethod() string {
	return http.MethodPost
}

func (h *Handler) GetPath() string {
	return "/movies/comments"
}

func getUserID(c *gin.Context) (int64, error) {
	uID, exist := c.Get(userIDKey)
	if !exist {
		return 0, errors.New("user ID not found")
	}

	uIDInt, ok := uID.(int64)
	if !ok {
		return 0, errors.New("user ID has invalid format")
	}

	return uIDInt, nil
}
