//go:generate minimock -i .commentDeleter -s _mock.go -o ./mock -g
package comment_delete

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"web_lab/internal/view"
)

type commentDeleter interface {
	DeleteComment(ctx context.Context, id int64) error
}

type Handler struct {
	commentDeleter commentDeleter
}

func New(commentDeleter commentDeleter) *Handler {
	return &Handler{
		commentDeleter: commentDeleter,
	}
}

// Handle godoc
// @Summary delete comment
// @Tags comment
// @Description delete comment for movie
// @ID comment-delete
// @Param Authorization header string true "jwt token"
// @Param id path string true "comment id"
// @Produce json
// @Success 200 "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 401 {object} view.ErrorResponse "unauthorized"
// @Failure 403 {object} view.ErrorResponse "forbidden"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies/comments/{id} [delete]
func (h *Handler) Handle(c *gin.Context) {
	ctx := context.Background()

	id := c.Param("id")

	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	err = h.commentDeleter.DeleteComment(ctx, commentID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetMethod() string {
	return http.MethodDelete
}

func (h *Handler) GetPath() string {
	return "/movies/comments/:id"
}
