//go:generate minimock -i .userCreator -s _mock.go -o ./mock -g
package user_register_post

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

type userCreator interface {
	CreateUser(ctx context.Context, user models.User) error
}

type Handler struct {
	userCreator userCreator
}

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(userCreator userCreator) *Handler {
	return &Handler{
		userCreator: userCreator,
	}
}

// Handle godoc
// @Summary create user
// @Tags user
// @Description create user in database
// @ID user-register-post
// @Param user body createUserRequest true "user for create"
// @Success 200 "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /users/register [post]
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

	var user createUserRequest

	err = json.Unmarshal(data, &user)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	if user.Username == "" {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(errors.New("username is empty")),
		)

		return
	}

	if user.Password == "" {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(errors.New("password is empty")),
		)

		return
	}

	err = h.userCreator.CreateUser(ctx, models.User{
		Username: user.Username,
		Password: user.Password,
		Role:     "user",
	})
	if err != nil {
		if errors.Is(err, models.ErrUserAlreadyExists) {
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
	return "/users/register"
}
