package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"
	"web_lab/internal/models"

	"github.com/gojuno/minimock/v3"
)

// StoreMock implements user_change_password_put.store
type StoreMock struct {
	t minimock.Tester

	funcAddState          func(ctx context.Context, stateID string, password string, code int64, userID int64) (err error)
	inspectFuncAddState   func(ctx context.Context, stateID string, password string, code int64, userID int64)
	afterAddStateCounter  uint64
	beforeAddStateCounter uint64
	AddStateMock          mStoreMockAddState

	funcGetUser          func(ctx context.Context, username string, password string) (u1 models.User, err error)
	inspectFuncGetUser   func(ctx context.Context, username string, password string)
	afterGetUserCounter  uint64
	beforeGetUserCounter uint64
	GetUserMock          mStoreMockGetUser
}

// NewStoreMock returns a mock for user_change_password_put.store
func NewStoreMock(t minimock.Tester) *StoreMock {
	m := &StoreMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddStateMock = mStoreMockAddState{mock: m}
	m.AddStateMock.callArgs = []*StoreMockAddStateParams{}

	m.GetUserMock = mStoreMockGetUser{mock: m}
	m.GetUserMock.callArgs = []*StoreMockGetUserParams{}

	return m
}

type mStoreMockAddState struct {
	mock               *StoreMock
	defaultExpectation *StoreMockAddStateExpectation
	expectations       []*StoreMockAddStateExpectation

	callArgs []*StoreMockAddStateParams
	mutex    sync.RWMutex
}

// StoreMockAddStateExpectation specifies expectation struct of the store.AddState
type StoreMockAddStateExpectation struct {
	mock    *StoreMock
	params  *StoreMockAddStateParams
	results *StoreMockAddStateResults
	Counter uint64
}

// StoreMockAddStateParams contains parameters of the store.AddState
type StoreMockAddStateParams struct {
	ctx      context.Context
	stateID  string
	password string
	code     int64
	userID   int64
}

// StoreMockAddStateResults contains results of the store.AddState
type StoreMockAddStateResults struct {
	err error
}

// Expect sets up expected params for store.AddState
func (mmAddState *mStoreMockAddState) Expect(ctx context.Context, stateID string, password string, code int64, userID int64) *mStoreMockAddState {
	if mmAddState.mock.funcAddState != nil {
		mmAddState.mock.t.Fatalf("StoreMock.AddState mock is already set by Set")
	}

	if mmAddState.defaultExpectation == nil {
		mmAddState.defaultExpectation = &StoreMockAddStateExpectation{}
	}

	mmAddState.defaultExpectation.params = &StoreMockAddStateParams{ctx, stateID, password, code, userID}
	for _, e := range mmAddState.expectations {
		if minimock.Equal(e.params, mmAddState.defaultExpectation.params) {
			mmAddState.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddState.defaultExpectation.params)
		}
	}

	return mmAddState
}

// Inspect accepts an inspector function that has same arguments as the store.AddState
func (mmAddState *mStoreMockAddState) Inspect(f func(ctx context.Context, stateID string, password string, code int64, userID int64)) *mStoreMockAddState {
	if mmAddState.mock.inspectFuncAddState != nil {
		mmAddState.mock.t.Fatalf("Inspect function is already set for StoreMock.AddState")
	}

	mmAddState.mock.inspectFuncAddState = f

	return mmAddState
}

// Return sets up results that will be returned by store.AddState
func (mmAddState *mStoreMockAddState) Return(err error) *StoreMock {
	if mmAddState.mock.funcAddState != nil {
		mmAddState.mock.t.Fatalf("StoreMock.AddState mock is already set by Set")
	}

	if mmAddState.defaultExpectation == nil {
		mmAddState.defaultExpectation = &StoreMockAddStateExpectation{mock: mmAddState.mock}
	}
	mmAddState.defaultExpectation.results = &StoreMockAddStateResults{err}
	return mmAddState.mock
}

// Set uses given function f to mock the store.AddState method
func (mmAddState *mStoreMockAddState) Set(f func(ctx context.Context, stateID string, password string, code int64, userID int64) (err error)) *StoreMock {
	if mmAddState.defaultExpectation != nil {
		mmAddState.mock.t.Fatalf("Default expectation is already set for the store.AddState method")
	}

	if len(mmAddState.expectations) > 0 {
		mmAddState.mock.t.Fatalf("Some expectations are already set for the store.AddState method")
	}

	mmAddState.mock.funcAddState = f
	return mmAddState.mock
}

// When sets expectation for the store.AddState which will trigger the result defined by the following
// Then helper
func (mmAddState *mStoreMockAddState) When(ctx context.Context, stateID string, password string, code int64, userID int64) *StoreMockAddStateExpectation {
	if mmAddState.mock.funcAddState != nil {
		mmAddState.mock.t.Fatalf("StoreMock.AddState mock is already set by Set")
	}

	expectation := &StoreMockAddStateExpectation{
		mock:   mmAddState.mock,
		params: &StoreMockAddStateParams{ctx, stateID, password, code, userID},
	}
	mmAddState.expectations = append(mmAddState.expectations, expectation)
	return expectation
}

// Then sets up store.AddState return parameters for the expectation previously defined by the When method
func (e *StoreMockAddStateExpectation) Then(err error) *StoreMock {
	e.results = &StoreMockAddStateResults{err}
	return e.mock
}

// AddState implements user_change_password_put.store
func (mmAddState *StoreMock) AddState(ctx context.Context, stateID string, password string, code int64, userID int64) (err error) {
	mm_atomic.AddUint64(&mmAddState.beforeAddStateCounter, 1)
	defer mm_atomic.AddUint64(&mmAddState.afterAddStateCounter, 1)

	if mmAddState.inspectFuncAddState != nil {
		mmAddState.inspectFuncAddState(ctx, stateID, password, code, userID)
	}

	mm_params := &StoreMockAddStateParams{ctx, stateID, password, code, userID}

	// Record call args
	mmAddState.AddStateMock.mutex.Lock()
	mmAddState.AddStateMock.callArgs = append(mmAddState.AddStateMock.callArgs, mm_params)
	mmAddState.AddStateMock.mutex.Unlock()

	for _, e := range mmAddState.AddStateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddState.AddStateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddState.AddStateMock.defaultExpectation.Counter, 1)
		mm_want := mmAddState.AddStateMock.defaultExpectation.params
		mm_got := StoreMockAddStateParams{ctx, stateID, password, code, userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddState.t.Errorf("StoreMock.AddState got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddState.AddStateMock.defaultExpectation.results
		if mm_results == nil {
			mmAddState.t.Fatal("No results are set for the StoreMock.AddState")
		}
		return (*mm_results).err
	}
	if mmAddState.funcAddState != nil {
		return mmAddState.funcAddState(ctx, stateID, password, code, userID)
	}
	mmAddState.t.Fatalf("Unexpected call to StoreMock.AddState. %v %v %v %v %v", ctx, stateID, password, code, userID)
	return
}

// AddStateAfterCounter returns a count of finished StoreMock.AddState invocations
func (mmAddState *StoreMock) AddStateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddState.afterAddStateCounter)
}

// AddStateBeforeCounter returns a count of StoreMock.AddState invocations
func (mmAddState *StoreMock) AddStateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddState.beforeAddStateCounter)
}

// Calls returns a list of arguments used in each call to StoreMock.AddState.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddState *mStoreMockAddState) Calls() []*StoreMockAddStateParams {
	mmAddState.mutex.RLock()

	argCopy := make([]*StoreMockAddStateParams, len(mmAddState.callArgs))
	copy(argCopy, mmAddState.callArgs)

	mmAddState.mutex.RUnlock()

	return argCopy
}

// MinimockAddStateDone returns true if the count of the AddState invocations corresponds
// the number of defined expectations
func (m *StoreMock) MinimockAddStateDone() bool {
	for _, e := range m.AddStateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddStateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddStateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddState != nil && mm_atomic.LoadUint64(&m.afterAddStateCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddStateInspect logs each unmet expectation
func (m *StoreMock) MinimockAddStateInspect() {
	for _, e := range m.AddStateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StoreMock.AddState with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddStateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddStateCounter) < 1 {
		if m.AddStateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StoreMock.AddState")
		} else {
			m.t.Errorf("Expected call to StoreMock.AddState with params: %#v", *m.AddStateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddState != nil && mm_atomic.LoadUint64(&m.afterAddStateCounter) < 1 {
		m.t.Error("Expected call to StoreMock.AddState")
	}
}

type mStoreMockGetUser struct {
	mock               *StoreMock
	defaultExpectation *StoreMockGetUserExpectation
	expectations       []*StoreMockGetUserExpectation

	callArgs []*StoreMockGetUserParams
	mutex    sync.RWMutex
}

// StoreMockGetUserExpectation specifies expectation struct of the store.GetUser
type StoreMockGetUserExpectation struct {
	mock    *StoreMock
	params  *StoreMockGetUserParams
	results *StoreMockGetUserResults
	Counter uint64
}

// StoreMockGetUserParams contains parameters of the store.GetUser
type StoreMockGetUserParams struct {
	ctx      context.Context
	username string
	password string
}

// StoreMockGetUserResults contains results of the store.GetUser
type StoreMockGetUserResults struct {
	u1  models.User
	err error
}

// Expect sets up expected params for store.GetUser
func (mmGetUser *mStoreMockGetUser) Expect(ctx context.Context, username string, password string) *mStoreMockGetUser {
	if mmGetUser.mock.funcGetUser != nil {
		mmGetUser.mock.t.Fatalf("StoreMock.GetUser mock is already set by Set")
	}

	if mmGetUser.defaultExpectation == nil {
		mmGetUser.defaultExpectation = &StoreMockGetUserExpectation{}
	}

	mmGetUser.defaultExpectation.params = &StoreMockGetUserParams{ctx, username, password}
	for _, e := range mmGetUser.expectations {
		if minimock.Equal(e.params, mmGetUser.defaultExpectation.params) {
			mmGetUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetUser.defaultExpectation.params)
		}
	}

	return mmGetUser
}

// Inspect accepts an inspector function that has same arguments as the store.GetUser
func (mmGetUser *mStoreMockGetUser) Inspect(f func(ctx context.Context, username string, password string)) *mStoreMockGetUser {
	if mmGetUser.mock.inspectFuncGetUser != nil {
		mmGetUser.mock.t.Fatalf("Inspect function is already set for StoreMock.GetUser")
	}

	mmGetUser.mock.inspectFuncGetUser = f

	return mmGetUser
}

// Return sets up results that will be returned by store.GetUser
func (mmGetUser *mStoreMockGetUser) Return(u1 models.User, err error) *StoreMock {
	if mmGetUser.mock.funcGetUser != nil {
		mmGetUser.mock.t.Fatalf("StoreMock.GetUser mock is already set by Set")
	}

	if mmGetUser.defaultExpectation == nil {
		mmGetUser.defaultExpectation = &StoreMockGetUserExpectation{mock: mmGetUser.mock}
	}
	mmGetUser.defaultExpectation.results = &StoreMockGetUserResults{u1, err}
	return mmGetUser.mock
}

// Set uses given function f to mock the store.GetUser method
func (mmGetUser *mStoreMockGetUser) Set(f func(ctx context.Context, username string, password string) (u1 models.User, err error)) *StoreMock {
	if mmGetUser.defaultExpectation != nil {
		mmGetUser.mock.t.Fatalf("Default expectation is already set for the store.GetUser method")
	}

	if len(mmGetUser.expectations) > 0 {
		mmGetUser.mock.t.Fatalf("Some expectations are already set for the store.GetUser method")
	}

	mmGetUser.mock.funcGetUser = f
	return mmGetUser.mock
}

// When sets expectation for the store.GetUser which will trigger the result defined by the following
// Then helper
func (mmGetUser *mStoreMockGetUser) When(ctx context.Context, username string, password string) *StoreMockGetUserExpectation {
	if mmGetUser.mock.funcGetUser != nil {
		mmGetUser.mock.t.Fatalf("StoreMock.GetUser mock is already set by Set")
	}

	expectation := &StoreMockGetUserExpectation{
		mock:   mmGetUser.mock,
		params: &StoreMockGetUserParams{ctx, username, password},
	}
	mmGetUser.expectations = append(mmGetUser.expectations, expectation)
	return expectation
}

// Then sets up store.GetUser return parameters for the expectation previously defined by the When method
func (e *StoreMockGetUserExpectation) Then(u1 models.User, err error) *StoreMock {
	e.results = &StoreMockGetUserResults{u1, err}
	return e.mock
}

// GetUser implements user_change_password_put.store
func (mmGetUser *StoreMock) GetUser(ctx context.Context, username string, password string) (u1 models.User, err error) {
	mm_atomic.AddUint64(&mmGetUser.beforeGetUserCounter, 1)
	defer mm_atomic.AddUint64(&mmGetUser.afterGetUserCounter, 1)

	if mmGetUser.inspectFuncGetUser != nil {
		mmGetUser.inspectFuncGetUser(ctx, username, password)
	}

	mm_params := &StoreMockGetUserParams{ctx, username, password}

	// Record call args
	mmGetUser.GetUserMock.mutex.Lock()
	mmGetUser.GetUserMock.callArgs = append(mmGetUser.GetUserMock.callArgs, mm_params)
	mmGetUser.GetUserMock.mutex.Unlock()

	for _, e := range mmGetUser.GetUserMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetUser.GetUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetUser.GetUserMock.defaultExpectation.Counter, 1)
		mm_want := mmGetUser.GetUserMock.defaultExpectation.params
		mm_got := StoreMockGetUserParams{ctx, username, password}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetUser.t.Errorf("StoreMock.GetUser got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetUser.GetUserMock.defaultExpectation.results
		if mm_results == nil {
			mmGetUser.t.Fatal("No results are set for the StoreMock.GetUser")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetUser.funcGetUser != nil {
		return mmGetUser.funcGetUser(ctx, username, password)
	}
	mmGetUser.t.Fatalf("Unexpected call to StoreMock.GetUser. %v %v %v", ctx, username, password)
	return
}

// GetUserAfterCounter returns a count of finished StoreMock.GetUser invocations
func (mmGetUser *StoreMock) GetUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetUser.afterGetUserCounter)
}

// GetUserBeforeCounter returns a count of StoreMock.GetUser invocations
func (mmGetUser *StoreMock) GetUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetUser.beforeGetUserCounter)
}

// Calls returns a list of arguments used in each call to StoreMock.GetUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetUser *mStoreMockGetUser) Calls() []*StoreMockGetUserParams {
	mmGetUser.mutex.RLock()

	argCopy := make([]*StoreMockGetUserParams, len(mmGetUser.callArgs))
	copy(argCopy, mmGetUser.callArgs)

	mmGetUser.mutex.RUnlock()

	return argCopy
}

// MinimockGetUserDone returns true if the count of the GetUser invocations corresponds
// the number of defined expectations
func (m *StoreMock) MinimockGetUserDone() bool {
	for _, e := range m.GetUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetUserCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetUser != nil && mm_atomic.LoadUint64(&m.afterGetUserCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetUserInspect logs each unmet expectation
func (m *StoreMock) MinimockGetUserInspect() {
	for _, e := range m.GetUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StoreMock.GetUser with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetUserCounter) < 1 {
		if m.GetUserMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StoreMock.GetUser")
		} else {
			m.t.Errorf("Expected call to StoreMock.GetUser with params: %#v", *m.GetUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetUser != nil && mm_atomic.LoadUint64(&m.afterGetUserCounter) < 1 {
		m.t.Error("Expected call to StoreMock.GetUser")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StoreMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddStateInspect()

		m.MinimockGetUserInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StoreMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *StoreMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddStateDone() &&
		m.MinimockGetUserDone()
}
