//go:generate minimock -i .movieDeleter -s _mock.go -o ./mock -g
package movie_delete

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"web_lab/internal/view"
)

type movieDeleter interface {
	DeleteMovie(ctx context.Context, id int64) error
}

type Handler struct {
	movieDeleter movieDeleter
}

func New(movieDeleter movieDeleter) *Handler {
	return &Handler{
		movieDeleter: movieDeleter,
	}
}

// Handle godoc
// @Summary delete movie
// @Tags movie
// @Description delete movie from catalog
// @ID movie-delete
// @Param Authorization header string true "jwt token"
// @Param id path string true "movie id"
// @Produce json
// @Success 200 "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 401 {object} view.ErrorResponse "unauthorized"
// @Failure 403 {object} view.ErrorResponse "forbidden"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies/{id} [delete]
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

	err = h.movieDeleter.DeleteMovie(ctx, movieID)
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
	return "/movies/:id"
}
