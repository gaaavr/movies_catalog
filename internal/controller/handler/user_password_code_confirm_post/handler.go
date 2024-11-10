//go:generate minimock -i .store -s _mock.go -o ./mock -g
package user_password_code_confirm_post

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
	UpdateUser(ctx context.Context, user models.User) error
	DeleteState(ctx context.Context, stateID string, code int64) (int64, string, error)
}

type Handler struct {
	store store
}

type confirmCodeRequest struct {
	State string `json:"state"`
	Code  int64  `json:"code"`
}

func New(store store) *Handler {
	return &Handler{
		store: store,
	}
}

// Handle godoc
// @Summary confirm code by user
// @Tags user
// @Description change password after confirm by user
// @ID user-pass-code-post
// @Param user body confirmCodeRequest  true "code and state for change pass"
// @Success 200
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /users/password/code [post]
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

	userID, newPass, err := h.store.DeleteState(ctx, request.State, request.Code)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	err = h.store.UpdateUser(ctx, models.User{
		ID:       int(userID),
		Password: newPass,
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
	return http.MethodPost
}

func (h *Handler) GetPath() string {
	return "/users/password/code"
}
