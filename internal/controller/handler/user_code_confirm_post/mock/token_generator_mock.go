package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"
	"web_lab/internal/models"

	"github.com/gojuno/minimock/v3"
)

// TokenGeneratorMock implements user_code_confirm_post.tokenGenerator
type TokenGeneratorMock struct {
	t minimock.Tester

	funcGenerateToken          func(user models.User) (s1 string, err error)
	inspectFuncGenerateToken   func(user models.User)
	afterGenerateTokenCounter  uint64
	beforeGenerateTokenCounter uint64
	GenerateTokenMock          mTokenGeneratorMockGenerateToken
}

// NewTokenGeneratorMock returns a mock for user_code_confirm_post.tokenGenerator
func NewTokenGeneratorMock(t minimock.Tester) *TokenGeneratorMock {
	m := &TokenGeneratorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GenerateTokenMock = mTokenGeneratorMockGenerateToken{mock: m}
	m.GenerateTokenMock.callArgs = []*TokenGeneratorMockGenerateTokenParams{}

	return m
}

type mTokenGeneratorMockGenerateToken struct {
	mock               *TokenGeneratorMock
	defaultExpectation *TokenGeneratorMockGenerateTokenExpectation
	expectations       []*TokenGeneratorMockGenerateTokenExpectation

	callArgs []*TokenGeneratorMockGenerateTokenParams
	mutex    sync.RWMutex
}

// TokenGeneratorMockGenerateTokenExpectation specifies expectation struct of the tokenGenerator.GenerateToken
type TokenGeneratorMockGenerateTokenExpectation struct {
	mock    *TokenGeneratorMock
	params  *TokenGeneratorMockGenerateTokenParams
	results *TokenGeneratorMockGenerateTokenResults
	Counter uint64
}

// TokenGeneratorMockGenerateTokenParams contains parameters of the tokenGenerator.GenerateToken
type TokenGeneratorMockGenerateTokenParams struct {
	user models.User
}

// TokenGeneratorMockGenerateTokenResults contains results of the tokenGenerator.GenerateToken
type TokenGeneratorMockGenerateTokenResults struct {
	s1  string
	err error
}

// Expect sets up expected params for tokenGenerator.GenerateToken
func (mmGenerateToken *mTokenGeneratorMockGenerateToken) Expect(user models.User) *mTokenGeneratorMockGenerateToken {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("TokenGeneratorMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &TokenGeneratorMockGenerateTokenExpectation{}
	}

	mmGenerateToken.defaultExpectation.params = &TokenGeneratorMockGenerateTokenParams{user}
	for _, e := range mmGenerateToken.expectations {
		if minimock.Equal(e.params, mmGenerateToken.defaultExpectation.params) {
			mmGenerateToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGenerateToken.defaultExpectation.params)
		}
	}

	return mmGenerateToken
}

// Inspect accepts an inspector function that has same arguments as the tokenGenerator.GenerateToken
func (mmGenerateToken *mTokenGeneratorMockGenerateToken) Inspect(f func(user models.User)) *mTokenGeneratorMockGenerateToken {
	if mmGenerateToken.mock.inspectFuncGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("Inspect function is already set for TokenGeneratorMock.GenerateToken")
	}

	mmGenerateToken.mock.inspectFuncGenerateToken = f

	return mmGenerateToken
}

// Return sets up results that will be returned by tokenGenerator.GenerateToken
func (mmGenerateToken *mTokenGeneratorMockGenerateToken) Return(s1 string, err error) *TokenGeneratorMock {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("TokenGeneratorMock.GenerateToken mock is already set by Set")
	}

	if mmGenerateToken.defaultExpectation == nil {
		mmGenerateToken.defaultExpectation = &TokenGeneratorMockGenerateTokenExpectation{mock: mmGenerateToken.mock}
	}
	mmGenerateToken.defaultExpectation.results = &TokenGeneratorMockGenerateTokenResults{s1, err}
	return mmGenerateToken.mock
}

// Set uses given function f to mock the tokenGenerator.GenerateToken method
func (mmGenerateToken *mTokenGeneratorMockGenerateToken) Set(f func(user models.User) (s1 string, err error)) *TokenGeneratorMock {
	if mmGenerateToken.defaultExpectation != nil {
		mmGenerateToken.mock.t.Fatalf("Default expectation is already set for the tokenGenerator.GenerateToken method")
	}

	if len(mmGenerateToken.expectations) > 0 {
		mmGenerateToken.mock.t.Fatalf("Some expectations are already set for the tokenGenerator.GenerateToken method")
	}

	mmGenerateToken.mock.funcGenerateToken = f
	return mmGenerateToken.mock
}

// When sets expectation for the tokenGenerator.GenerateToken which will trigger the result defined by the following
// Then helper
func (mmGenerateToken *mTokenGeneratorMockGenerateToken) When(user models.User) *TokenGeneratorMockGenerateTokenExpectation {
	if mmGenerateToken.mock.funcGenerateToken != nil {
		mmGenerateToken.mock.t.Fatalf("TokenGeneratorMock.GenerateToken mock is already set by Set")
	}

	expectation := &TokenGeneratorMockGenerateTokenExpectation{
		mock:   mmGenerateToken.mock,
		params: &TokenGeneratorMockGenerateTokenParams{user},
	}
	mmGenerateToken.expectations = append(mmGenerateToken.expectations, expectation)
	return expectation
}

// Then sets up tokenGenerator.GenerateToken return parameters for the expectation previously defined by the When method
func (e *TokenGeneratorMockGenerateTokenExpectation) Then(s1 string, err error) *TokenGeneratorMock {
	e.results = &TokenGeneratorMockGenerateTokenResults{s1, err}
	return e.mock
}

// GenerateToken implements user_code_confirm_post.tokenGenerator
func (mmGenerateToken *TokenGeneratorMock) GenerateToken(user models.User) (s1 string, err error) {
	mm_atomic.AddUint64(&mmGenerateToken.beforeGenerateTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmGenerateToken.afterGenerateTokenCounter, 1)

	if mmGenerateToken.inspectFuncGenerateToken != nil {
		mmGenerateToken.inspectFuncGenerateToken(user)
	}

	mm_params := &TokenGeneratorMockGenerateTokenParams{user}

	// Record call args
	mmGenerateToken.GenerateTokenMock.mutex.Lock()
	mmGenerateToken.GenerateTokenMock.callArgs = append(mmGenerateToken.GenerateTokenMock.callArgs, mm_params)
	mmGenerateToken.GenerateTokenMock.mutex.Unlock()

	for _, e := range mmGenerateToken.GenerateTokenMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmGenerateToken.GenerateTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGenerateToken.GenerateTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmGenerateToken.GenerateTokenMock.defaultExpectation.params
		mm_got := TokenGeneratorMockGenerateTokenParams{user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGenerateToken.t.Errorf("TokenGeneratorMock.GenerateToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGenerateToken.GenerateTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmGenerateToken.t.Fatal("No results are set for the TokenGeneratorMock.GenerateToken")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmGenerateToken.funcGenerateToken != nil {
		return mmGenerateToken.funcGenerateToken(user)
	}
	mmGenerateToken.t.Fatalf("Unexpected call to TokenGeneratorMock.GenerateToken. %v", user)
	return
}

// GenerateTokenAfterCounter returns a count of finished TokenGeneratorMock.GenerateToken invocations
func (mmGenerateToken *TokenGeneratorMock) GenerateTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGenerateToken.afterGenerateTokenCounter)
}

// GenerateTokenBeforeCounter returns a count of TokenGeneratorMock.GenerateToken invocations
func (mmGenerateToken *TokenGeneratorMock) GenerateTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGenerateToken.beforeGenerateTokenCounter)
}

// Calls returns a list of arguments used in each call to TokenGeneratorMock.GenerateToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGenerateToken *mTokenGeneratorMockGenerateToken) Calls() []*TokenGeneratorMockGenerateTokenParams {
	mmGenerateToken.mutex.RLock()

	argCopy := make([]*TokenGeneratorMockGenerateTokenParams, len(mmGenerateToken.callArgs))
	copy(argCopy, mmGenerateToken.callArgs)

	mmGenerateToken.mutex.RUnlock()

	return argCopy
}

// MinimockGenerateTokenDone returns true if the count of the GenerateToken invocations corresponds
// the number of defined expectations
func (m *TokenGeneratorMock) MinimockGenerateTokenDone() bool {
	for _, e := range m.GenerateTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GenerateTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGenerateTokenCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGenerateToken != nil && mm_atomic.LoadUint64(&m.afterGenerateTokenCounter) < 1 {
		return false
	}
	return true
}

// MinimockGenerateTokenInspect logs each unmet expectation
func (m *TokenGeneratorMock) MinimockGenerateTokenInspect() {
	for _, e := range m.GenerateTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TokenGeneratorMock.GenerateToken with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GenerateTokenMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGenerateTokenCounter) < 1 {
		if m.GenerateTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TokenGeneratorMock.GenerateToken")
		} else {
			m.t.Errorf("Expected call to TokenGeneratorMock.GenerateToken with params: %#v", *m.GenerateTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGenerateToken != nil && mm_atomic.LoadUint64(&m.afterGenerateTokenCounter) < 1 {
		m.t.Error("Expected call to TokenGeneratorMock.GenerateToken")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TokenGeneratorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGenerateTokenInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TokenGeneratorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TokenGeneratorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGenerateTokenDone()
}
