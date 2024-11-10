//go:generate minimock -i .store,.tokenGenerator -s _mock.go -o ./mock -g
package user_code_confirm_post

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"web_lab/internal/models"
	"web_lab/internal/view"
)

type store interface {
	GetUserByID(ctx context.Context, userID int64) (models.User, error)
	DeleteState(ctx context.Context, stateID string, code int64) (int64, error)
}

type tokenGenerator interface {
	GenerateToken(user models.User) (string, error)
}

type Handler struct {
	store          store
	tokenGenerator tokenGenerator
}

type confirmCodeRequest struct {
	State string `json:"state"`
	Code  int64  `json:"code"`
}

func New(store store, tokenGenerator tokenGenerator) *Handler {
	return &Handler{
		store:          store,
		tokenGenerator: tokenGenerator,
	}
}

// Handle godoc
// @Summary confirm code by user
// @Tags user
// @Description get token after success login user
// @ID user-code-post
// @Param user body confirmCodeRequest  true "code and state for login"
// @Produce json
// @Success 200 {object} view.Token "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /users/code [post]
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

	var request confirmCodeRequest

	err = json.Unmarshal(data, &request)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	userID, err := h.store.DeleteState(ctx, request.State, request.Code)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	user, err := h.store.GetUserByID(ctx, userID)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	token, err := h.tokenGenerator.GenerateToken(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.JSON(http.StatusOK, view.Token{
		Token: token,
		Role:  user.Role,
	})
}

func (h *Handler) GetMethod() string {
	return http.MethodPost
}

func (h *Handler) GetPath() string {
	return "/users/code"
}
