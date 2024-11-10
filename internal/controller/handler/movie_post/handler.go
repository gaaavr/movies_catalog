//go:generate minimock -i .movieCreator -s _mock.go -o ./mock -g
package movie_post

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

type movieCreator interface {
	CreateMovie(ctx context.Context, movie models.Movie) error
}

type Handler struct {
	movieCreator movieCreator
}

type createMovieRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func New(movieCreator movieCreator) *Handler {
	return &Handler{
		movieCreator: movieCreator,
	}
}

// Handle godoc
// @Summary create movie
// @Tags movie
// @Description create movie in catalog
// @ID movie-post
// @Param Authorization header string true "jwt token"
// @Param movie body createMovieRequest true "movie for create"
// @Success 200 "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 401 {object} view.ErrorResponse "unauthorized"
// @Failure 403 {object} view.ErrorResponse "forbidden"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies [post]
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

	var movie createMovieRequest

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

	err = h.movieCreator.CreateMovie(ctx, models.Movie{
		Title:       movie.Title,
		Description: movie.Description,
		Image:       movie.Image,
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
	return "/movies"
}
