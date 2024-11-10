package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"web_lab/internal/config"
	"web_lab/internal/controller/middleware/cors_middleware"
)

type authenticator interface {
	AuthUser() gin.HandlerFunc
	AuthAdmin() gin.HandlerFunc
}

type Server struct {
	Address        string
	router         *gin.Engine
	authUserGroup  *gin.RouterGroup
	authAdminGroup *gin.RouterGroup
}

func NewServer(
	config config.ServerConfig,
	authenticator authenticator,
) *Server {
	router := gin.New()

	// Определение функций шаблонов
	funcMap := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
	}
	// Загрузка шаблонов с функциями
	router.SetHTMLTemplate(template.Must(template.New("").Funcs(funcMap).ParseGlob("./internal/view/*")))

	router.Use(gin.Recovery(), cors_middleware.Cors())

	// Middleware для обработки метода PUT
	router.Use(func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			switch c.PostForm("_method") {
			case http.MethodPut:
				c.Request.Method = http.MethodPut
			case http.MethodDelete:
				c.Request.Method = http.MethodDelete
			}
		}
	})

	authUserGroup := router.Group("")
	authAdminGroup := router.Group("")

	authAdminGroup.Use(authenticator.AuthAdmin())
	authUserGroup.Use(authenticator.AuthUser())

	return &Server{
		Address:        fmt.Sprintf("%s:%s", config.Host, config.Port),
		router:         router,
		authUserGroup:  authUserGroup,
		authAdminGroup: authAdminGroup,
	}
}

func (s *Server) RegisterHandler(method, path string, handler func(*gin.Context)) {
	s.router.Handle(method, path, handler)
}

func (s *Server) RegisterUserHandler(method, path string, handler func(*gin.Context)) {
	s.authUserGroup.Handle(method, path, handler)
}

func (s *Server) RegisterAdminHandler(method, path string, handler func(*gin.Context)) {
	s.authAdminGroup.Handle(method, path, handler)
}

func (s *Server) Run() error {
	err := s.router.Run(s.Address)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}
