//go:generate minimock -i .commentsGetter -s _mock.go -o ./mock -g
package comments_get

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"web_lab/internal/models"
	"web_lab/internal/view"
)

type commentsGetter interface {
	GetComments(ctx context.Context, movieID int64) ([]models.Comment, error)
}

type Handler struct {
	commentsGetter commentsGetter
}

func New(commentsGetter commentsGetter) *Handler {
	return &Handler{
		commentsGetter: commentsGetter,
	}
}

// Handle godoc
// @Summary get comments
// @Tags comment
// @Description get comments for movie
// @ID comments-get
// @Produce json
// @Success 200 {object} view.CommentsResponse "success"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies/{id}/comments [get]
func (h *Handler) Handle(c *gin.Context) {
	ctx := context.Background()

	id := c.Param("id")

	movieID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	comments, err := h.commentsGetter.GetComments(ctx, movieID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.JSON(http.StatusOK, view.ConvertCommentsResponse(comments))
}

func (h *Handler) GetMethod() string {
	return http.MethodGet
}

func (h *Handler) GetPath() string {
	return "/movies/:id/comments"
}
