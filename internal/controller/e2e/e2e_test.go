//go:build e2e
// +build e2e

package e2e

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/rs/zerolog/log"

	"web_lab/internal/view"
)

const (
	contentType = "application/json"
)

type responseCode struct{}
type responseBody struct{}
type state struct{}
type tokenKey struct{}
type timeKey struct{}

func TestRegisterLoginAndWriteComment(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeWriteCommentScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/write_comment.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestDeleteCommentErr(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeDeleteCommentFailedScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/delete_comment.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeWriteCommentScenario(sc *godog.ScenarioContext) {
	sc.Step(`^user registers with "([^"]+)" and "([^"]+)"$`, register)
	sc.Step(`^should be register successfully and receive code (\d+)$`, checkResponseCode)
	sc.Step(`^the user tries to log in with "([^"]+)" and "([^"]+)"$`, login)
	sc.Step(`^he receives response code (\d+), non-empty state and send confirmation code by email$`, checkResponseCodeAndState)
	sc.Step(`^the user enters the verification code from "([^"]+)" and state$`, setConfirmCode)
	sc.Step(`^he receive response code (\d+) and a non-empty token$`, checkResponseAndCodeAfterConfirmCode)
	sc.Step(`^user with a token leaves a comment "([^"]+)" on a movie with ID (\d+)$`, writeComment)
	sc.Step(`^he receives a response code of (\d+) in response$`, checkResponseCode)
}

func InitializeDeleteCommentFailedScenario(sc *godog.ScenarioContext) {
	sc.Step(`^user with a "([^"]+)" tries to delete a comment with ID (\d+)$`, deleteComment)
	sc.Step(`^he will get response code (\d+)$`, checkResponseCode)
}

func deleteComment(ctx context.Context, token string, commentID int64) (context.Context, error) {
	port := os.Getenv("SERVICE_PORT")

	if port == "" {
		return ctx, errors.New("SERVICE_PORT is empty")
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:%s/movies/comments/%d", port, commentID), nil)
	if err != nil {
		return ctx, fmt.Errorf("failed build delete request: %w", err)
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return ctx, fmt.Errorf("failed to make deleterequest: %w", err)
	}

	defer resp.Body.Close()

	ctx = context.WithValue(ctx, responseCode{}, resp.StatusCode)

	return ctx, nil
}

func register(ctx context.Context, emailKey, passwordKey string) (context.Context, error) {
	email := os.Getenv(emailKey)
	password := os.Getenv(passwordKey)

	if email == "" {
		return ctx, errors.New("email env is empty")
	}

	if password == "" {
		return ctx, errors.New("password env is empty")
	}

	port := os.Getenv("SERVICE_PORT")

	if port == "" {
		return ctx, errors.New("SERVICE_PORT is empty")
	}

	type createUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := createUserRequest{
		Username: email,
		Password: password,
	}

	jsonStr, err := json.Marshal(req)
	if err != nil {
		return ctx, err
	}

	resp, err := http.Post(fmt.Sprintf("http://localhost:%s/users/register", port),
		contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ctx, fmt.Errorf("failed to make post request: %w", err)
	}

	defer resp.Body.Close()

	ctx = context.WithValue(ctx, responseCode{}, resp.StatusCode)

	return ctx, nil
}

func checkResponseCode(ctx context.Context, code int) error {
	respCode, ok := ctx.Value(responseCode{}).(int)
	if !ok {
		return errors.New("response code not found")
	}

	if respCode != code {
		return fmt.Errorf("expected response code - 200, actual - %d", respCode)
	}

	return nil
}

func login(ctx context.Context, emailKey, passwordKey string) (context.Context, error) {
	email := os.Getenv(emailKey)
	password := os.Getenv(passwordKey)

	if email == "" {
		return ctx, errors.New("email env is empty")
	}

	if password == "" {
		return ctx, errors.New("password env is empty")
	}

	port := os.Getenv("SERVICE_PORT")

	if port == "" {
		return ctx, errors.New("SERVICE_PORT is empty")
	}

	type createUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := createUserRequest{
		Username: email,
		Password: password,
	}

	jsonStr, err := json.Marshal(req)
	if err != nil {
		return ctx, err
	}

	resp, err := http.Post(fmt.Sprintf("http://localhost:%s/users/login", port),
		contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ctx, fmt.Errorf("failed to make post request: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx, fmt.Errorf("failed to read body: %w", err)
	}

	ctx = context.WithValue(ctx, responseCode{}, resp.StatusCode)
	ctx = context.WithValue(ctx, responseBody{}, string(data))
	ctx = context.WithValue(ctx, timeKey{}, time.Now())

	return ctx, nil
}

func checkResponseCodeAndState(ctx context.Context, code int) (context.Context, error) {
	respCode, ok := ctx.Value(responseCode{}).(int)
	if !ok {
		return ctx, errors.New("response code not found")
	}

	if respCode != code {
		return ctx, fmt.Errorf("expected response code after create role - 200, actual - %d", respCode)
	}

	respBody, ok := ctx.Value(responseBody{}).(string)
	if !ok {
		return ctx, errors.New("response body not found")
	}

	var st view.State

	err := json.Unmarshal([]byte(respBody), &st)
	if err != nil {
		return ctx, fmt.Errorf("failed to unmarshal state body: %w", err)
	}

	ctx = context.WithValue(ctx, state{}, st.State)

	return ctx, nil
}

func setConfirmCode(ctx context.Context, emailKey string) (context.Context, error) {
	port := os.Getenv("SERVICE_PORT")

	if port == "" {
		return ctx, errors.New("SERVICE_PORT is empty")
	}

	st, ok := ctx.Value(state{}).(string)
	if !ok {
		return ctx, errors.New("state not found")
	}

	sendTimeCode, ok := ctx.Value(timeKey{}).(time.Time)
	if !ok {
		return ctx, errors.New("send code time not found")
	}

	var (
		code int64
		err  error
	)

	for time.Since(sendTimeCode) < 1*time.Minute {
		code, err = readEmail(emailKey)
		if err != nil {
			log.Printf("failed to get confirm code from email: %v", err)
			continue
		}

		break
	}

	type confirmCodeRequest struct {
		State string `json:"state"`
		Code  int64  `json:"code"`
	}

	req := confirmCodeRequest{
		State: st,
		Code:  code,
	}

	jsonStr, err := json.Marshal(req)
	if err != nil {
		return ctx, err
	}

	resp, err := http.Post(fmt.Sprintf("http://localhost:%s/users/code", port),
		contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ctx, fmt.Errorf("failed to make post request: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ctx, fmt.Errorf("failed to read body: %w", err)
	}

	ctx = context.WithValue(ctx, responseCode{}, resp.StatusCode)
	ctx = context.WithValue(ctx, responseBody{}, string(data))

	return ctx, nil
}

func checkResponseAndCodeAfterConfirmCode(ctx context.Context, code int) (context.Context, error) {
	respCode, ok := ctx.Value(responseCode{}).(int)
	if !ok {
		return ctx, errors.New("response code not found")
	}

	if respCode != code {
		return ctx, fmt.Errorf("expected response code after create role - 200, actual - %d", respCode)
	}

	respBody, ok := ctx.Value(responseBody{}).(string)
	if !ok {
		return ctx, errors.New("response body not found")
	}

	var token view.Token

	err := json.Unmarshal([]byte(respBody), &token)
	if err != nil {
		return ctx, fmt.Errorf("failed to unmarshal state body: %w", err)
	}

	if token.Token == "" {
		return ctx, errors.New("token is empty")
	}

	ctx = context.WithValue(ctx, tokenKey{}, token.Token)

	return ctx, nil
}

func writeComment(ctx context.Context, content string, movieID int64) (context.Context, error) {
	token, ok := ctx.Value(tokenKey{}).(string)
	if !ok {
		return ctx, errors.New("token not found")
	}

	port := os.Getenv("SERVICE_PORT")

	if port == "" {
		return ctx, errors.New("SERVICE_PORT is empty")
	}

	type createCommentRequest struct {
		Content string `json:"content"`
		MovieID int64  `json:"movie_id"`
	}

	commentReq := createCommentRequest{
		Content: content,
		MovieID: movieID,
	}

	jsonStr, err := json.Marshal(commentReq)
	if err != nil {
		return ctx, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%s/movies/comments", port), bytes.NewBuffer(jsonStr))
	if err != nil {
		return ctx, fmt.Errorf("failed build post request: %w", err)
	}

	req.Header.Set("Authorization", token)

	client := &http.Client{}

	// Выполнение запроса
	resp, err := client.Do(req)
	if err != nil {
		return ctx, fmt.Errorf("failed to make post request: %w", err)
	}

	defer resp.Body.Close()

	ctx = context.WithValue(ctx, responseCode{}, resp.StatusCode)

	return ctx, nil
}

func readEmail(emailKey string) (int64, error) {
	email := os.Getenv(emailKey)
	password := os.Getenv("EMAIL_APP_PASS")
	server := os.Getenv("EMAIL_SERVER")
	serverName := os.Getenv("EMAIL_SERVER_NAME")

	if email == "" {
		return 0, errors.New("email env is empty")
	}

	if password == "" {
		return 0, errors.New("email app password env is empty")
	}

	if server == "" {
		return 0, errors.New("server is empty")
	}

	if serverName == "" {
		return 0, errors.New("server name is empty")
	}

	// Connect to the IMAP server
	c, err := client.DialTLS(server, &tls.Config{ServerName: serverName})
	if err != nil {
		return 0, fmt.Errorf("dial error: %w", err)
	}
	defer c.Logout()

	// Login
	err = c.Login(email, password)
	if err != nil {
		return 0, fmt.Errorf("login error: %w", err)
	}

	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return 0, fmt.Errorf("select error: %w", err)
	}

	// Get the last message
	seqset := new(imap.SeqSet)
	seqset.AddNum(mbox.Messages)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()

	for msg := range messages {
		if msg == nil {
			continue
		}

		// Process the message
		r := msg.GetBody(section)
		if r == nil {
			continue
		}

		mr, err := mail.CreateReader(r)
		if err != nil {
			return 0, fmt.Errorf("failed to create message reader: %w", err)
		}

		header := mr.Header
		subject, err := header.Subject()
		if err != nil {
			return 0, fmt.Errorf("failed to get subject: %w", err)
		}

		if subject != "Код подтверждения от каталога фильмов" {
			return 0, fmt.Errorf("subject got: %s, want: %s", subject, "Код подтверждения от каталога фильмов")
		}

		// Process each message part
		for {
			p, err := mr.NextPart()
			if err != nil {
				break
			}

			switch p.Header.(type) {
			case *mail.InlineHeader:
				b, err := io.ReadAll(p.Body)
				if err != nil {
					return 0, fmt.Errorf("failed to read message body: %w", err)
				}

				confCode, err := strconv.ParseInt(strings.TrimSuffix(string(b), "\r\n"), 10, 64)
				if err != nil {
					return 0, fmt.Errorf("failed to parse code: %w", err)
				}

				return confCode, nil
			}
		}
	}

	if err := <-done; err != nil {
		return 0, fmt.Errorf("failed to get confirmation code: %w", err)
	}

	return 0, errors.New("unknown err")
}
