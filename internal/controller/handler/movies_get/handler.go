//go:generate minimock -i .moviesGetter -s _mock.go -o ./mock -g
package movies_get

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"web_lab/internal/models"
	"web_lab/internal/view"
)

type moviesGetter interface {
	GetMovies(ctx context.Context, opts models.MoviesOpts) ([]models.Movie, int64, error)
}

type Handler struct {
	moviesGetter moviesGetter
}

func New(moviesGetter moviesGetter) *Handler {
	return &Handler{
		moviesGetter: moviesGetter,
	}
}

// Handle godoc
// @Summary get movies
// @Tags movie
// @Description get movies from catalog
// @ID movies-get
// @Param search query string false "movie title or description, may be incomplete"
// @Param limit query int false "element count limit in movies list"
// @Param offset query int false "offset in movies list"
// @Produce json
// @Success 200 {object} view.MoviesResponse "success"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /movies [get]
func (h *Handler) Handle(c *gin.Context) {
	ctx := context.Background()
	opts := getClientsOpts(c)

	movies, count, err := h.moviesGetter.GetMovies(ctx, opts)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.JSON(http.StatusOK, view.ConvertMoviesResponse(movies, count))
}

func getClientsOpts(c *gin.Context) models.MoviesOpts {
	var opts models.MoviesOpts

	fillFilters(c, &opts)

	if limit, ok := c.GetQuery("limit"); ok {
		intLimit, err := strconv.ParseInt(limit, 10, 64)
		if err == nil {
			opts.Pagination.Limit = &intLimit
		}
	}

	if offset, ok := c.GetQuery("offset"); ok {
		intOffset, err := strconv.ParseInt(offset, 10, 64)
		if err == nil {
			opts.Pagination.Offset = &intOffset
		}
	}

	return opts
}

func fillFilters(c *gin.Context, opts *models.MoviesOpts) {
	if search, ok := c.GetQuery("search"); ok && search != "" {
		opts.Search = &search
	}
}

func (h *Handler) GetMethod() string {
	return http.MethodGet
}

func (h *Handler) GetPath() string {
	return "/movies"
}
