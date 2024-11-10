//go:generate minimock -i .movieGetter -s _mock.go -o ./mock -g
package movie_get

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"web_lab/internal/models"
	"web_lab/internal/view"
)

type movieGetter interface {
	GetMovie(ctx context.Context, id int64) (models.Movie, error)
}

type Handler struct {
	movieGetter movieGetter
}

func New(movieGetter movieGetter) *Handler {
	return &Handler{
		movieGetter: movieGetter,
	}
}

// Handle godoc
// @Summary get movie
// @Tags movie
// @Description get movie by ID
// @ID movie-get
// @Param id path string true "movie id"
// @Produce json
// @Success 200 {object} view.Movie "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies/{id} [get]
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

	movie, err := h.movieGetter.GetMovie(ctx, movieID)
	if err != nil {
		if errors.Is(err, models.ErrMovieNotFound) {
			c.JSON(
				http.StatusNotFound,
				view.NewErrorResponse(err))

			return
		}
		
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.JSON(http.StatusOK, view.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Image:       movie.Image,
	})
}

func (h *Handler) GetMethod() string {
	return http.MethodGet
}

func (h *Handler) GetPath() string {
	return "/movies/:id"
}
