package auth_middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"web_lab/internal/view"
)

const (
	authHeader = "Authorization"
	userRole   = "user"
	adminRole  = "admin"
	userIDKey  = "userID"
)

var (
	forbiddenErr    = errors.New("user has no rights")
	emptyTokenErr   = errors.New("token is empty")
	invalidTokenErr = errors.New("token is invalid")
)

type tokenChecker interface {
	TokenIsValid(token string) (bool, string, int64, error)
}

type Handler struct {
	tokenChecker tokenChecker
}

func New(tokenChecker tokenChecker) *Handler {
	return &Handler{
		tokenChecker: tokenChecker,
	}
}

func (h *Handler) AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(authHeader)
		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				view.NewErrorResponse(emptyTokenErr),
			)

			return
		}

		isValid, role, userID, err := h.tokenChecker.TokenIsValid(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				view.NewErrorResponse(err),
			)

			return
		}

		if !isValid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				view.NewErrorResponse(invalidTokenErr),
			)

			return
		}

		if role != adminRole {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				view.NewErrorResponse(forbiddenErr))

			return
		}

		c.Set(userIDKey, userID)

		c.Next()
	}
}

func (h *Handler) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(authHeader)
		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				view.NewErrorResponse(emptyTokenErr),
			)

			return
		}

		isValid, role, userID, err := h.tokenChecker.TokenIsValid(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				view.NewErrorResponse(emptyTokenErr),
			)

			return
		}

		if !isValid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				view.NewErrorResponse(invalidTokenErr),
			)

			return
		}

		if role != userRole && role != adminRole {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				view.NewErrorResponse(forbiddenErr))

			return
		}

		c.Set(userIDKey, userID)

		c.Next()
	}
}
