//go:generate minimock -i .store -s _mock.go -o ./mock -g
package user_change_password_put

import (
	"context"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/gomail.v2"

	"web_lab/internal/config"
	"web_lab/internal/models"
	"web_lab/internal/view"
)

type store interface {
	GetUser(ctx context.Context, username, password string) (models.User, error)
	AddState(ctx context.Context, stateID, password string, code, userID int64) error
}

type Handler struct {
	store    store
	server   string
	port     int
	username string
	password string
}

type userChangePasswordRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func New(store store, emailCfg config.EmailConfig) *Handler {
	return &Handler{
		store:    store,
		server:   emailCfg.Server,
		port:     emailCfg.Port,
		username: emailCfg.Username,
		password: emailCfg.AppPass,
	}
}

// Handle godoc
// @Summary user change password
// @Tags user
// @Description get code after change password
// @ID user-pass-put
// @Param user body userChangePasswordRequest true "username, password  and new password"
// @Produce json
// @Success 200 {object} view.State "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /users/password [put]
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

	var request userChangePasswordRequest

	err = json.Unmarshal(data, &request)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			view.NewErrorResponse(err),
		)

		return
	}

	user, err := h.store.GetUser(ctx, request.Username, request.Password)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := randomizer.Int63n(9000) + 1000
	state := uuid.NewString()

	err = h.store.AddState(ctx, state, request.NewPassword, randomNumber, int64(user.ID))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	err = h.sendEmail(request.Username, "Код подтверждения от каталога фильмов", randomNumber)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	c.JSON(http.StatusOK, view.State{
		State: state,
	})
}

func (h *Handler) sendEmail(to string, subject string, code int64) error {
	m := gomail.NewMessage()
	m.SetHeader("From", h.username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", strconv.FormatInt(code, 10))

	d := gomail.NewDialer(h.server, h.port, h.username, h.password)

	return d.DialAndSend(m)
}

func (h *Handler) GetMethod() string {
	return http.MethodPut
}

func (h *Handler) GetPath() string {
	return "/users/password"
}
