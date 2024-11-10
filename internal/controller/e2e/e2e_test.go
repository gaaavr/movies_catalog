package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/cucumber/godog"
)

const (
	contentType = "application/json"
)

type emailCtxKey struct{}
type passwordCtxKey struct{}

func register(ctx context.Context, emailKey, passwordKey string) (context.Context, error) {
	type createUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	email := os.Getenv(emailKey)
	password := os.Getenv(passwordKey)

	ctx = context.WithValue(ctx, emailCtxKey{}, email)
	ctx = context.WithValue(ctx, passwordCtxKey{}, password)

	req := createUserRequest{
		Username: email,
		Password: password,
	}

	jsonStr, err := json.Marshal(req)
	if err != nil {
		return ctx, err
	}

	// Создание новой роли
	resp, err := http.Post(fmt.Sprintf("http://localhost:8080/users/register"),
		contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ctx, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx, err
	}

	if resp.StatusCode != http.StatusOK {
		return ctx, fmt.Errorf("expected response code after create role - 200, actual - %d, error - %s", resp.StatusCode, string(data))
	}

	return ctx, nil
}

func getParams(ctx context.Context, expectedEmail, expectedPassword string) error {
	valueEmail, ok := ctx.Value(emailCtxKey{}).(string)
	if !ok {
		return fmt.Errorf("значение переменной окружения не найдено")
	}

	valuePassword, ok := ctx.Value(passwordCtxKey{}).(string)
	if !ok {
		return fmt.Errorf("значение переменной окружения не найдено")
	}

	if valueEmail != expectedEmail {
		return fmt.Errorf("bad email: %s", valueEmail)
	}

	if valuePassword != expectedPassword {
		return fmt.Errorf("bad pass: %s", valuePassword)
	}

	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^get "([^"]+)" and "([^"]+)"$`, register)
	sc.Step(`^there should be "([^"]+)" and "([^"]+)" remaining$`, getParams)
}
