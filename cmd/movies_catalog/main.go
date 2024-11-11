package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	telemetryLog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"

	_ "web_lab/docs"
	"web_lab/internal/config"
	"web_lab/internal/controller"
	"web_lab/internal/controller/handler/comment_delete"
	"web_lab/internal/controller/handler/comment_post"
	"web_lab/internal/controller/handler/comments_get"
	"web_lab/internal/controller/handler/movie_delete"
	"web_lab/internal/controller/handler/movie_get"
	"web_lab/internal/controller/handler/movie_post"
	"web_lab/internal/controller/handler/movie_put"
	"web_lab/internal/controller/handler/movies_get"
	"web_lab/internal/controller/handler/user_change_password_put"
	"web_lab/internal/controller/handler/user_code_confirm_post"
	"web_lab/internal/controller/handler/user_login_post"
	"web_lab/internal/controller/handler/user_password_code_confirm_post"
	"web_lab/internal/controller/handler/user_register_post"
	"web_lab/internal/controller/middleware/auth_middleware"
	"web_lab/internal/models"
	"web_lab/internal/service/jwt"
	"web_lab/internal/storage/pg"
)

// @title movies catalog
// @version 1.0
// @description web app for movies catalog
// @host localhost:8080
// @BasePath /
// @Schemes http https
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create config")
	}

	pgStorage, err := pg.NewStorage(cfg.Postgres)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create pg connection")
	}

	defer pgStorage.Close()

	otelShutdown, err := setupOTelSDK(context.Background())
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	storePG := models.NewPG(pgStorage)

	tokenService := jwt.New()

	createMovieHandler := movie_post.New(storePG)
	updateMovieHandler := movie_put.New(storePG)
	deleteMovieHandler := movie_delete.New(storePG)
	getMoviesHandler := movies_get.New(storePG)
	getMovieHandler := movie_get.New(storePG)

	createCommentHandler := comment_post.New(storePG)
	getCommentsHandler := comments_get.New(storePG)
	deleteCommentHandler := comment_delete.New(storePG)

	createUserHandler := user_register_post.New(storePG)
	loginUserHandler := user_login_post.New(storePG, cfg.Email)
	changePassUserHandler := user_change_password_put.New(storePG, cfg.Email)
	confirmCodeChangePassHandler := user_password_code_confirm_post.New(storePG)
	confirmCodeUserHandler := user_code_confirm_post.New(storePG, tokenService)

	server := controller.NewServer(cfg.Server, auth_middleware.New(tokenService))

	// Регистрация администраторских ручек
	server.RegisterAdminHandler(createMovieHandler.GetMethod(),
		createMovieHandler.GetPath(), createMovieHandler.Handle)
	server.RegisterAdminHandler(updateMovieHandler.GetMethod(),
		updateMovieHandler.GetPath(), updateMovieHandler.Handle)
	server.RegisterAdminHandler(deleteMovieHandler.GetMethod(),
		deleteMovieHandler.GetPath(), deleteMovieHandler.Handle)
	server.RegisterAdminHandler(deleteCommentHandler.GetMethod(),
		deleteCommentHandler.GetPath(), deleteCommentHandler.Handle)

	// Регистрация ручек для залогиненных пользователей
	server.RegisterUserHandler(createCommentHandler.GetMethod(),
		createCommentHandler.GetPath(), createCommentHandler.Handle)

	// Регистрация обычных ручек
	server.RegisterHandler(http.MethodGet,
		"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.RegisterHandler(createUserHandler.GetMethod(),
		createUserHandler.GetPath(), createUserHandler.Handle)
	server.RegisterHandler(loginUserHandler.GetMethod(),
		loginUserHandler.GetPath(), loginUserHandler.Handle)
	server.RegisterHandler(changePassUserHandler.GetMethod(),
		changePassUserHandler.GetPath(), changePassUserHandler.Handle)
	server.RegisterHandler(confirmCodeUserHandler.GetMethod(),
		confirmCodeUserHandler.GetPath(), confirmCodeUserHandler.Handle)
	server.RegisterHandler(confirmCodeChangePassHandler.GetMethod(),
		confirmCodeChangePassHandler.GetPath(), confirmCodeChangePassHandler.Handle)
	server.RegisterHandler(getMoviesHandler.GetMethod(),
		getMoviesHandler.GetPath(), getMoviesHandler.Handle)
	server.RegisterHandler(getMovieHandler.GetMethod(),
		getMovieHandler.GetPath(), getMovieHandler.Handle)
	server.RegisterHandler(getCommentsHandler.GetMethod(),
		getCommentsHandler.GetPath(), getCommentsHandler.Handle)

	// Регистрация ручек с шаблонами
	server.RegisterHandler(http.MethodGet,
		"/catalog", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "movies.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/catalog/:id", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "movie.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/movie-update", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "movie_update.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/movie-add", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "movie_add.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/register", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "register.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/login", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "login.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/code", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "confirm_code.html", gin.H{})
		},
	)
	server.RegisterHandler(http.MethodGet,
		"/error", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "error.html", gin.H{})
		},
	)

	server.RegisterHandler(http.MethodGet,
		"/not-found", func(context *gin.Context) {
			otelgin.HTML(context, http.StatusOK, "error_not_found.html", gin.H{})
		},
	)

	log.Info().Msg("Service is running")

	err = server.Run()
	if err != nil {
		log.Fatal().Err(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	log.Info().Msg("Gracefully stopped")
}

// setupOTelSDK bootstraps the OpenTelemetry pipeline.
// If it does not return an error, make sure to call shutdown for proper cleanup.
func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// handleErr calls shutdown for cleanup and makes sure that all errors are returned.
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	// Set up propagator.
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// Set up trace provider.
	tracerProvider, err := newTraceProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up meter provider.
	meterProvider, err := newMeterProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	// Set up logger provider.
	loggerProvider, err := newLoggerProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider() (*trace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT"))))
	if err != nil {
		return nil, err
	}

	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithBatcher(traceExporter))
	return traceProvider, nil
}

func newMeterProvider() (*metric.MeterProvider, error) {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter,
			// Default is 1m. Set to 3s for demonstrative purposes.
			metric.WithInterval(3*time.Second))),
	)
	return meterProvider, nil
}

func newLoggerProvider() (*telemetryLog.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := telemetryLog.NewLoggerProvider(
		telemetryLog.WithProcessor(telemetryLog.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}
