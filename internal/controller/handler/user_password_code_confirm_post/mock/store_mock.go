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

// StoreMock implements user_password_code_confirm_post.store
type StoreMock struct {
	t minimock.Tester

	funcDeleteState          func(ctx context.Context, stateID string, code int64) (i1 int64, s1 string, err error)
	inspectFuncDeleteState   func(ctx context.Context, stateID string, code int64)
	afterDeleteStateCounter  uint64
	beforeDeleteStateCounter uint64
	DeleteStateMock          mStoreMockDeleteState

	funcUpdateUser          func(ctx context.Context, user models.User) (err error)
	inspectFuncUpdateUser   func(ctx context.Context, user models.User)
	afterUpdateUserCounter  uint64
	beforeUpdateUserCounter uint64
	UpdateUserMock          mStoreMockUpdateUser
}

// NewStoreMock returns a mock for user_password_code_confirm_post.store
func NewStoreMock(t minimock.Tester) *StoreMock {
	m := &StoreMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteStateMock = mStoreMockDeleteState{mock: m}
	m.DeleteStateMock.callArgs = []*StoreMockDeleteStateParams{}

	m.UpdateUserMock = mStoreMockUpdateUser{mock: m}
	m.UpdateUserMock.callArgs = []*StoreMockUpdateUserParams{}

	return m
}

type mStoreMockDeleteState struct {
	mock               *StoreMock
	defaultExpectation *StoreMockDeleteStateExpectation
	expectations       []*StoreMockDeleteStateExpectation

	callArgs []*StoreMockDeleteStateParams
	mutex    sync.RWMutex
}

// StoreMockDeleteStateExpectation specifies expectation struct of the store.DeleteState
type StoreMockDeleteStateExpectation struct {
	mock    *StoreMock
	params  *StoreMockDeleteStateParams
	results *StoreMockDeleteStateResults
	Counter uint64
}

// StoreMockDeleteStateParams contains parameters of the store.DeleteState
type StoreMockDeleteStateParams struct {
	ctx     context.Context
	stateID string
	code    int64
}

// StoreMockDeleteStateResults contains results of the store.DeleteState
type StoreMockDeleteStateResults struct {
	i1  int64
	s1  string
	err error
}

// Expect sets up expected params for store.DeleteState
func (mmDeleteState *mStoreMockDeleteState) Expect(ctx context.Context, stateID string, code int64) *mStoreMockDeleteState {
	if mmDeleteState.mock.funcDeleteState != nil {
		mmDeleteState.mock.t.Fatalf("StoreMock.DeleteState mock is already set by Set")
	}

	if mmDeleteState.defaultExpectation == nil {
		mmDeleteState.defaultExpectation = &StoreMockDeleteStateExpectation{}
	}

	mmDeleteState.defaultExpectation.params = &StoreMockDeleteStateParams{ctx, stateID, code}
	for _, e := range mmDeleteState.expectations {
		if minimock.Equal(e.params, mmDeleteState.defaultExpectation.params) {
			mmDeleteState.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteState.defaultExpectation.params)
		}
	}

	return mmDeleteState
}

// Inspect accepts an inspector function that has same arguments as the store.DeleteState
func (mmDeleteState *mStoreMockDeleteState) Inspect(f func(ctx context.Context, stateID string, code int64)) *mStoreMockDeleteState {
	if mmDeleteState.mock.inspectFuncDeleteState != nil {
		mmDeleteState.mock.t.Fatalf("Inspect function is already set for StoreMock.DeleteState")
	}

	mmDeleteState.mock.inspectFuncDeleteState = f

	return mmDeleteState
}

// Return sets up results that will be returned by store.DeleteState
func (mmDeleteState *mStoreMockDeleteState) Return(i1 int64, s1 string, err error) *StoreMock {
	if mmDeleteState.mock.funcDeleteState != nil {
		mmDeleteState.mock.t.Fatalf("StoreMock.DeleteState mock is already set by Set")
	}

	if mmDeleteState.defaultExpectation == nil {
		mmDeleteState.defaultExpectation = &StoreMockDeleteStateExpectation{mock: mmDeleteState.mock}
	}
	mmDeleteState.defaultExpectation.results = &StoreMockDeleteStateResults{i1, s1, err}
	return mmDeleteState.mock
}

// Set uses given function f to mock the store.DeleteState method
func (mmDeleteState *mStoreMockDeleteState) Set(f func(ctx context.Context, stateID string, code int64) (i1 int64, s1 string, err error)) *StoreMock {
	if mmDeleteState.defaultExpectation != nil {
		mmDeleteState.mock.t.Fatalf("Default expectation is already set for the store.DeleteState method")
	}

	if len(mmDeleteState.expectations) > 0 {
		mmDeleteState.mock.t.Fatalf("Some expectations are already set for the store.DeleteState method")
	}

	mmDeleteState.mock.funcDeleteState = f
	return mmDeleteState.mock
}

// When sets expectation for the store.DeleteState which will trigger the result defined by the following
// Then helper
func (mmDeleteState *mStoreMockDeleteState) When(ctx context.Context, stateID string, code int64) *StoreMockDeleteStateExpectation {
	if mmDeleteState.mock.funcDeleteState != nil {
		mmDeleteState.mock.t.Fatalf("StoreMock.DeleteState mock is already set by Set")
	}

	expectation := &StoreMockDeleteStateExpectation{
		mock:   mmDeleteState.mock,
		params: &StoreMockDeleteStateParams{ctx, stateID, code},
	}
	mmDeleteState.expectations = append(mmDeleteState.expectations, expectation)
	return expectation
}

// Then sets up store.DeleteState return parameters for the expectation previously defined by the When method
func (e *StoreMockDeleteStateExpectation) Then(i1 int64, s1 string, err error) *StoreMock {
	e.results = &StoreMockDeleteStateResults{i1, s1, err}
	return e.mock
}

// DeleteState implements user_password_code_confirm_post.store
func (mmDeleteState *StoreMock) DeleteState(ctx context.Context, stateID string, code int64) (i1 int64, s1 string, err error) {
	mm_atomic.AddUint64(&mmDeleteState.beforeDeleteStateCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteState.afterDeleteStateCounter, 1)

	if mmDeleteState.inspectFuncDeleteState != nil {
		mmDeleteState.inspectFuncDeleteState(ctx, stateID, code)
	}

	mm_params := &StoreMockDeleteStateParams{ctx, stateID, code}

	// Record call args
	mmDeleteState.DeleteStateMock.mutex.Lock()
	mmDeleteState.DeleteStateMock.callArgs = append(mmDeleteState.DeleteStateMock.callArgs, mm_params)
	mmDeleteState.DeleteStateMock.mutex.Unlock()

	for _, e := range mmDeleteState.DeleteStateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.s1, e.results.err
		}
	}

	if mmDeleteState.DeleteStateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteState.DeleteStateMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteState.DeleteStateMock.defaultExpectation.params
		mm_got := StoreMockDeleteStateParams{ctx, stateID, code}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteState.t.Errorf("StoreMock.DeleteState got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteState.DeleteStateMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteState.t.Fatal("No results are set for the StoreMock.DeleteState")
		}
		return (*mm_results).i1, (*mm_results).s1, (*mm_results).err
	}
	if mmDeleteState.funcDeleteState != nil {
		return mmDeleteState.funcDeleteState(ctx, stateID, code)
	}
	mmDeleteState.t.Fatalf("Unexpected call to StoreMock.DeleteState. %v %v %v", ctx, stateID, code)
	return
}

// DeleteStateAfterCounter returns a count of finished StoreMock.DeleteState invocations
func (mmDeleteState *StoreMock) DeleteStateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteState.afterDeleteStateCounter)
}

// DeleteStateBeforeCounter returns a count of StoreMock.DeleteState invocations
func (mmDeleteState *StoreMock) DeleteStateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteState.beforeDeleteStateCounter)
}

// Calls returns a list of arguments used in each call to StoreMock.DeleteState.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteState *mStoreMockDeleteState) Calls() []*StoreMockDeleteStateParams {
	mmDeleteState.mutex.RLock()

	argCopy := make([]*StoreMockDeleteStateParams, len(mmDeleteState.callArgs))
	copy(argCopy, mmDeleteState.callArgs)

	mmDeleteState.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteStateDone returns true if the count of the DeleteState invocations corresponds
// the number of defined expectations
func (m *StoreMock) MinimockDeleteStateDone() bool {
	for _, e := range m.DeleteStateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteStateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteStateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteState != nil && mm_atomic.LoadUint64(&m.afterDeleteStateCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteStateInspect logs each unmet expectation
func (m *StoreMock) MinimockDeleteStateInspect() {
	for _, e := range m.DeleteStateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StoreMock.DeleteState with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteStateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteStateCounter) < 1 {
		if m.DeleteStateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StoreMock.DeleteState")
		} else {
			m.t.Errorf("Expected call to StoreMock.DeleteState with params: %#v", *m.DeleteStateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteState != nil && mm_atomic.LoadUint64(&m.afterDeleteStateCounter) < 1 {
		m.t.Error("Expected call to StoreMock.DeleteState")
	}
}

type mStoreMockUpdateUser struct {
	mock               *StoreMock
	defaultExpectation *StoreMockUpdateUserExpectation
	expectations       []*StoreMockUpdateUserExpectation

	callArgs []*StoreMockUpdateUserParams
	mutex    sync.RWMutex
}

// StoreMockUpdateUserExpectation specifies expectation struct of the store.UpdateUser
type StoreMockUpdateUserExpectation struct {
	mock    *StoreMock
	params  *StoreMockUpdateUserParams
	results *StoreMockUpdateUserResults
	Counter uint64
}

// StoreMockUpdateUserParams contains parameters of the store.UpdateUser
type StoreMockUpdateUserParams struct {
	ctx  context.Context
	user models.User
}

// StoreMockUpdateUserResults contains results of the store.UpdateUser
type StoreMockUpdateUserResults struct {
	err error
}

// Expect sets up expected params for store.UpdateUser
func (mmUpdateUser *mStoreMockUpdateUser) Expect(ctx context.Context, user models.User) *mStoreMockUpdateUser {
	if mmUpdateUser.mock.funcUpdateUser != nil {
		mmUpdateUser.mock.t.Fatalf("StoreMock.UpdateUser mock is already set by Set")
	}

	if mmUpdateUser.defaultExpectation == nil {
		mmUpdateUser.defaultExpectation = &StoreMockUpdateUserExpectation{}
	}

	mmUpdateUser.defaultExpectation.params = &StoreMockUpdateUserParams{ctx, user}
	for _, e := range mmUpdateUser.expectations {
		if minimock.Equal(e.params, mmUpdateUser.defaultExpectation.params) {
			mmUpdateUser.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateUser.defaultExpectation.params)
		}
	}

	return mmUpdateUser
}

// Inspect accepts an inspector function that has same arguments as the store.UpdateUser
func (mmUpdateUser *mStoreMockUpdateUser) Inspect(f func(ctx context.Context, user models.User)) *mStoreMockUpdateUser {
	if mmUpdateUser.mock.inspectFuncUpdateUser != nil {
		mmUpdateUser.mock.t.Fatalf("Inspect function is already set for StoreMock.UpdateUser")
	}

	mmUpdateUser.mock.inspectFuncUpdateUser = f

	return mmUpdateUser
}

// Return sets up results that will be returned by store.UpdateUser
func (mmUpdateUser *mStoreMockUpdateUser) Return(err error) *StoreMock {
	if mmUpdateUser.mock.funcUpdateUser != nil {
		mmUpdateUser.mock.t.Fatalf("StoreMock.UpdateUser mock is already set by Set")
	}

	if mmUpdateUser.defaultExpectation == nil {
		mmUpdateUser.defaultExpectation = &StoreMockUpdateUserExpectation{mock: mmUpdateUser.mock}
	}
	mmUpdateUser.defaultExpectation.results = &StoreMockUpdateUserResults{err}
	return mmUpdateUser.mock
}

// Set uses given function f to mock the store.UpdateUser method
func (mmUpdateUser *mStoreMockUpdateUser) Set(f func(ctx context.Context, user models.User) (err error)) *StoreMock {
	if mmUpdateUser.defaultExpectation != nil {
		mmUpdateUser.mock.t.Fatalf("Default expectation is already set for the store.UpdateUser method")
	}

	if len(mmUpdateUser.expectations) > 0 {
		mmUpdateUser.mock.t.Fatalf("Some expectations are already set for the store.UpdateUser method")
	}

	mmUpdateUser.mock.funcUpdateUser = f
	return mmUpdateUser.mock
}

// When sets expectation for the store.UpdateUser which will trigger the result defined by the following
// Then helper
func (mmUpdateUser *mStoreMockUpdateUser) When(ctx context.Context, user models.User) *StoreMockUpdateUserExpectation {
	if mmUpdateUser.mock.funcUpdateUser != nil {
		mmUpdateUser.mock.t.Fatalf("StoreMock.UpdateUser mock is already set by Set")
	}

	expectation := &StoreMockUpdateUserExpectation{
		mock:   mmUpdateUser.mock,
		params: &StoreMockUpdateUserParams{ctx, user},
	}
	mmUpdateUser.expectations = append(mmUpdateUser.expectations, expectation)
	return expectation
}

// Then sets up store.UpdateUser return parameters for the expectation previously defined by the When method
func (e *StoreMockUpdateUserExpectation) Then(err error) *StoreMock {
	e.results = &StoreMockUpdateUserResults{err}
	return e.mock
}

// UpdateUser implements user_password_code_confirm_post.store
func (mmUpdateUser *StoreMock) UpdateUser(ctx context.Context, user models.User) (err error) {
	mm_atomic.AddUint64(&mmUpdateUser.beforeUpdateUserCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateUser.afterUpdateUserCounter, 1)

	if mmUpdateUser.inspectFuncUpdateUser != nil {
		mmUpdateUser.inspectFuncUpdateUser(ctx, user)
	}

	mm_params := &StoreMockUpdateUserParams{ctx, user}

	// Record call args
	mmUpdateUser.UpdateUserMock.mutex.Lock()
	mmUpdateUser.UpdateUserMock.callArgs = append(mmUpdateUser.UpdateUserMock.callArgs, mm_params)
	mmUpdateUser.UpdateUserMock.mutex.Unlock()

	for _, e := range mmUpdateUser.UpdateUserMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateUser.UpdateUserMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateUser.UpdateUserMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateUser.UpdateUserMock.defaultExpectation.params
		mm_got := StoreMockUpdateUserParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateUser.t.Errorf("StoreMock.UpdateUser got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateUser.UpdateUserMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateUser.t.Fatal("No results are set for the StoreMock.UpdateUser")
		}
		return (*mm_results).err
	}
	if mmUpdateUser.funcUpdateUser != nil {
		return mmUpdateUser.funcUpdateUser(ctx, user)
	}
	mmUpdateUser.t.Fatalf("Unexpected call to StoreMock.UpdateUser. %v %v", ctx, user)
	return
}

// UpdateUserAfterCounter returns a count of finished StoreMock.UpdateUser invocations
func (mmUpdateUser *StoreMock) UpdateUserAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateUser.afterUpdateUserCounter)
}

// UpdateUserBeforeCounter returns a count of StoreMock.UpdateUser invocations
func (mmUpdateUser *StoreMock) UpdateUserBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateUser.beforeUpdateUserCounter)
}

// Calls returns a list of arguments used in each call to StoreMock.UpdateUser.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateUser *mStoreMockUpdateUser) Calls() []*StoreMockUpdateUserParams {
	mmUpdateUser.mutex.RLock()

	argCopy := make([]*StoreMockUpdateUserParams, len(mmUpdateUser.callArgs))
	copy(argCopy, mmUpdateUser.callArgs)

	mmUpdateUser.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateUserDone returns true if the count of the UpdateUser invocations corresponds
// the number of defined expectations
func (m *StoreMock) MinimockUpdateUserDone() bool {
	for _, e := range m.UpdateUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateUserCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateUser != nil && mm_atomic.LoadUint64(&m.afterUpdateUserCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateUserInspect logs each unmet expectation
func (m *StoreMock) MinimockUpdateUserInspect() {
	for _, e := range m.UpdateUserMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StoreMock.UpdateUser with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateUserMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateUserCounter) < 1 {
		if m.UpdateUserMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StoreMock.UpdateUser")
		} else {
			m.t.Errorf("Expected call to StoreMock.UpdateUser with params: %#v", *m.UpdateUserMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateUser != nil && mm_atomic.LoadUint64(&m.afterUpdateUserCounter) < 1 {
		m.t.Error("Expected call to StoreMock.UpdateUser")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StoreMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteStateInspect()

		m.MinimockUpdateUserInspect()
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
		m.MinimockDeleteStateDone() &&
		m.MinimockUpdateUserDone()
}