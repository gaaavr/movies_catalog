//go:generate minimock -i .store -s _mock.go -o ./mock -g
package user_login_post

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

	"web_lab/internal/models"
	"web_lab/internal/view"
)

type store interface {
	GetUser(ctx context.Context, username, password string) (models.User, error)
	AddState(ctx context.Context, stateID string, code, userID int64) error
}

type Handler struct {
	store store
}

type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func New(store store) *Handler {
	return &Handler{
		store: store,
	}
}

// Handle godoc
// @Summary login user
// @Tags user
// @Description get code after enter login and password
// @ID user-login-post
// @Param user body loginUserRequest  true "username and password for login"
// @Produce json
// @Success 200 {object} view.State "success"
// @Failure 400 {object} view.ErrorResponse "bad request"
// @Failure 500 {object} view.ErrorResponse "internal error"
// @Router /users/login [post]
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

	var request loginUserRequest

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

	err = h.store.AddState(ctx, state, randomNumber, int64(user.ID))
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			view.NewErrorResponse(err),
		)

		return
	}

	err = sendEmail(request.Username, "Код подтверждения от каталога фильмов", randomNumber)
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

func sendEmail(to string, subject string, code int64) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "ssg0808@yandex.ru")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", strconv.FormatInt(code, 10))

	d := gomail.NewDialer("smtp.yandex.ru", 587, "ssg0808@yandex.ru", "prvhxbylrwrmgnxo")

	return d.DialAndSend(m)
}

func (h *Handler) GetMethod() string {
	return http.MethodPost
}

func (h *Handler) GetPath() string {
	return "/users/login"
}
