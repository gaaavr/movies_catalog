//go:generate minimock -i .movieUpdater -s _mock.go -o ./mock -g
package movie_put

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

type movieUpdater interface {
	UpdateMovie(ctx context.Context, movie models.Movie) error
}

type Handler struct {
	movieUpdater movieUpdater
}

type updateMovieRequest struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func New(movieUpdater movieUpdater) *Handler {
	return &Handler{
		movieUpdater: movieUpdater,
	}
}

// Handle godoc
// @Summary update movie
// @Tags movie
// @Description create movie in catalog
// @ID movie-put
// @Param Authorization header string true "jwt token"
// @Param movie body updateMovieRequest true "movie for update"
// @Success 200 "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 401 {object} view.ErrorResponse "unauthorized"
// @Failure 403 {object} view.ErrorResponse "forbidden"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies [put]
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

	var movie updateMovieRequest

	err = json.Unmarshal(data, &movie)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	if movie.Title == "" {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(errors.New("title is empty")),
		)

		return
	}

	err = h.movieUpdater.UpdateMovie(ctx, models.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Image:       movie.Image,
	})
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
	return http.MethodPut
}

func (h *Handler) GetPath() string {
	return "/movies"
}
