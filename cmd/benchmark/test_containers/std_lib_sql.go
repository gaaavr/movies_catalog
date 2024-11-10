package test_containers

import (
	"context"
	"fmt"
	"os"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"web_lab/internal/config"
	"web_lab/internal/models"
	"web_lab/internal/storage/pg"
)

func SetupTestDatabaseStdLib() (testcontainers.Container, *models.PGStd, error) {
	ctx := context.Background()

	cfg := config.PostgresConfig{
		Username: "local",
		Password: "local",
		Database: "dev",
		SslMode:  "disable",
	}

	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:15.3-alpine3.18",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       cfg.Database,
			"POSTGRES_PASSWORD": cfg.Password,
			"POSTGRES_USER":     cfg.Username,
			"POSTGRES_SSL_MODE": cfg.SslMode,
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate container: %w", err)
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		return nil, nil, err
	}

	natPort, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, err
	}

	cfg.Host = host
	cfg.Port = natPort.Port()

	stdStorage, err := pg.NewStorageStd(cfg)
	if err != nil {
		return nil, nil, err
	}

	text, err := os.ReadFile("./db/dumps/init.sql")
	if err != nil {
		return nil, nil, err
	}

	if _, err = stdStorage.Exec(string(text)); err != nil {
		return nil, nil, err
	}

	store := models.NewPGStd(stdStorage)

	return dbContainer, store, nil
}
