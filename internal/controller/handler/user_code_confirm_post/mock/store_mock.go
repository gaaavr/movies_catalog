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

// StoreMock implements user_code_confirm_post.store
type StoreMock struct {
	t minimock.Tester

	funcDeleteState          func(ctx context.Context, stateID string, code int64) (i1 int64, err error)
	inspectFuncDeleteState   func(ctx context.Context, stateID string, code int64)
	afterDeleteStateCounter  uint64
	beforeDeleteStateCounter uint64
	DeleteStateMock          mStoreMockDeleteState

	funcGetUserByID          func(ctx context.Context, userID int64) (u1 models.User, err error)
	inspectFuncGetUserByID   func(ctx context.Context, userID int64)
	afterGetUserByIDCounter  uint64
	beforeGetUserByIDCounter uint64
	GetUserByIDMock          mStoreMockGetUserByID
}

// NewStoreMock returns a mock for user_code_confirm_post.store
func NewStoreMock(t minimock.Tester) *StoreMock {
	m := &StoreMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteStateMock = mStoreMockDeleteState{mock: m}
	m.DeleteStateMock.callArgs = []*StoreMockDeleteStateParams{}

	m.GetUserByIDMock = mStoreMockGetUserByID{mock: m}
	m.GetUserByIDMock.callArgs = []*StoreMockGetUserByIDParams{}

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
func (mmDeleteState *mStoreMockDeleteState) Return(i1 int64, err error) *StoreMock {
	if mmDeleteState.mock.funcDeleteState != nil {
		mmDeleteState.mock.t.Fatalf("StoreMock.DeleteState mock is already set by Set")
	}

	if mmDeleteState.defaultExpectation == nil {
		mmDeleteState.defaultExpectation = &StoreMockDeleteStateExpectation{mock: mmDeleteState.mock}
	}
	mmDeleteState.defaultExpectation.results = &StoreMockDeleteStateResults{i1, err}
	return mmDeleteState.mock
}

// Set uses given function f to mock the store.DeleteState method
func (mmDeleteState *mStoreMockDeleteState) Set(f func(ctx context.Context, stateID string, code int64) (i1 int64, err error)) *StoreMock {
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
func (e *StoreMockDeleteStateExpectation) Then(i1 int64, err error) *StoreMock {
	e.results = &StoreMockDeleteStateResults{i1, err}
	return e.mock
}

// DeleteState implements user_code_confirm_post.store
func (mmDeleteState *StoreMock) DeleteState(ctx context.Context, stateID string, code int64) (i1 int64, err error) {
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
			return e.results.i1, e.results.err
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
		return (*mm_results).i1, (*mm_results).err
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

type mStoreMockGetUserByID struct {
	mock               *StoreMock
	defaultExpectation *StoreMockGetUserByIDExpectation
	expectations       []*StoreMockGetUserByIDExpectation

	callArgs []*StoreMockGetUserByIDParams
	mutex    sync.RWMutex
}

// StoreMockGetUserByIDExpectation specifies expectation struct of the store.GetUserByID
type StoreMockGetUserByIDExpectation struct {
	mock    *StoreMock
	params  *StoreMockGetUserByIDParams
	results *StoreMockGetUserByIDResults
	Counter uint64
}

// StoreMockGetUserByIDParams contains parameters of the store.GetUserByID
type StoreMockGetUserByIDParams struct {
	ctx    context.Context
	userID int64
}

// StoreMockGetUserByIDResults contains results of the store.GetUserByID
type StoreMockGetUserByIDResults struct {
	u1  models.User
	err error
}

// Expect sets up expected params for store.GetUserByID
func (mmGetUserByID *mStoreMockGetUserByID) Expect(ctx context.Context, userID int64) *mStoreMockGetUserByID {
	if mmGetUserByID.mock.funcGetUserByID != nil {
		mmGetUserByID.mock.t.Fatalf("StoreMock.GetUserByID mock is already set by Set")
	}

	if mmGetUserByID.defaultExpectation == nil {
		mmGetUserByID.defaultExpectation = &StoreMockGetUserByIDExpectation{}
	}

	mmGetUserByID.defaultExpectation.params = &StoreMockGetUserByIDParams{ctx, userID}
	for _, e := range mmGetUserByID.expectations {
		if minimock.Equal(e.params, mmGetUserByID.defaultExpectation.params) {
			mmGetUserByID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetUserByID.defaultExpectation.params)
		}
	}

	return mmGetUserByID
}

// Inspect accepts an inspector function that has same arguments as the store.GetUserByID
func (mmGetUserByID *mStoreMockGetUserByID) Inspect(f func(ctx context.Context, userID int64)) *mStoreMockGetUserByID {
	if mmGetUserByID.mock.inspectFuncGetUserByID != nil {
		mmGetUserByID.mock.t.Fatalf("Inspect function is already set for StoreMock.GetUserByID")
	}

	mmGetUserByID.mock.inspectFuncGetUserByID = f

	return mmGetUserByID
}

// Return sets up results that will be returned by store.GetUserByID
func (mmGetUserByID *mStoreMockGetUserByID) Return(u1 models.User, err error) *StoreMock {
	if mmGetUserByID.mock.funcGetUserByID != nil {
		mmGetUserByID.mock.t.Fatalf("StoreMock.GetUserByID mock is already set by Set")
	}

	if mmGetUserByID.defaultExpectation == nil {
		mmGetUserByID.defaultExpectation = &StoreMockGetUserByIDExpectation{mock: mmGetUserByID.mock}
	}
	mmGetUserByID.defaultExpectation.results = &StoreMockGetUserByIDResults{u1, err}
	return mmGetUserByID.mock
}

// Set uses given function f to mock the store.GetUserByID method
func (mmGetUserByID *mStoreMockGetUserByID) Set(f func(ctx context.Context, userID int64) (u1 models.User, err error)) *StoreMock {
	if mmGetUserByID.defaultExpectation != nil {
		mmGetUserByID.mock.t.Fatalf("Default expectation is already set for the store.GetUserByID method")
	}

	if len(mmGetUserByID.expectations) > 0 {
		mmGetUserByID.mock.t.Fatalf("Some expectations are already set for the store.GetUserByID method")
	}

	mmGetUserByID.mock.funcGetUserByID = f
	return mmGetUserByID.mock
}

// When sets expectation for the store.GetUserByID which will trigger the result defined by the following
// Then helper
func (mmGetUserByID *mStoreMockGetUserByID) When(ctx context.Context, userID int64) *StoreMockGetUserByIDExpectation {
	if mmGetUserByID.mock.funcGetUserByID != nil {
		mmGetUserByID.mock.t.Fatalf("StoreMock.GetUserByID mock is already set by Set")
	}

	expectation := &StoreMockGetUserByIDExpectation{
		mock:   mmGetUserByID.mock,
		params: &StoreMockGetUserByIDParams{ctx, userID},
	}
	mmGetUserByID.expectations = append(mmGetUserByID.expectations, expectation)
	return expectation
}

// Then sets up store.GetUserByID return parameters for the expectation previously defined by the When method
func (e *StoreMockGetUserByIDExpectation) Then(u1 models.User, err error) *StoreMock {
	e.results = &StoreMockGetUserByIDResults{u1, err}
	return e.mock
}

// GetUserByID implements user_code_confirm_post.store
func (mmGetUserByID *StoreMock) GetUserByID(ctx context.Context, userID int64) (u1 models.User, err error) {
	mm_atomic.AddUint64(&mmGetUserByID.beforeGetUserByIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetUserByID.afterGetUserByIDCounter, 1)

	if mmGetUserByID.inspectFuncGetUserByID != nil {
		mmGetUserByID.inspectFuncGetUserByID(ctx, userID)
	}

	mm_params := &StoreMockGetUserByIDParams{ctx, userID}

	// Record call args
	mmGetUserByID.GetUserByIDMock.mutex.Lock()
	mmGetUserByID.GetUserByIDMock.callArgs = append(mmGetUserByID.GetUserByIDMock.callArgs, mm_params)
	mmGetUserByID.GetUserByIDMock.mutex.Unlock()

	for _, e := range mmGetUserByID.GetUserByIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetUserByID.GetUserByIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetUserByID.GetUserByIDMock.defaultExpectation.Counter, 1)
		mm_want := mmGetUserByID.GetUserByIDMock.defaultExpectation.params
		mm_got := StoreMockGetUserByIDParams{ctx, userID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetUserByID.t.Errorf("StoreMock.GetUserByID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetUserByID.GetUserByIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetUserByID.t.Fatal("No results are set for the StoreMock.GetUserByID")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetUserByID.funcGetUserByID != nil {
		return mmGetUserByID.funcGetUserByID(ctx, userID)
	}
	mmGetUserByID.t.Fatalf("Unexpected call to StoreMock.GetUserByID. %v %v", ctx, userID)
	return
}

// GetUserByIDAfterCounter returns a count of finished StoreMock.GetUserByID invocations
func (mmGetUserByID *StoreMock) GetUserByIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetUserByID.afterGetUserByIDCounter)
}

// GetUserByIDBeforeCounter returns a count of StoreMock.GetUserByID invocations
func (mmGetUserByID *StoreMock) GetUserByIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetUserByID.beforeGetUserByIDCounter)
}

// Calls returns a list of arguments used in each call to StoreMock.GetUserByID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetUserByID *mStoreMockGetUserByID) Calls() []*StoreMockGetUserByIDParams {
	mmGetUserByID.mutex.RLock()

	argCopy := make([]*StoreMockGetUserByIDParams, len(mmGetUserByID.callArgs))
	copy(argCopy, mmGetUserByID.callArgs)

	mmGetUserByID.mutex.RUnlock()

	return argCopy
}

// MinimockGetUserByIDDone returns true if the count of the GetUserByID invocations corresponds
// the number of defined expectations
func (m *StoreMock) MinimockGetUserByIDDone() bool {
	for _, e := range m.GetUserByIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetUserByIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetUserByIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetUserByID != nil && mm_atomic.LoadUint64(&m.afterGetUserByIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetUserByIDInspect logs each unmet expectation
func (m *StoreMock) MinimockGetUserByIDInspect() {
	for _, e := range m.GetUserByIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StoreMock.GetUserByID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetUserByIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetUserByIDCounter) < 1 {
		if m.GetUserByIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StoreMock.GetUserByID")
		} else {
			m.t.Errorf("Expected call to StoreMock.GetUserByID with params: %#v", *m.GetUserByIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetUserByID != nil && mm_atomic.LoadUint64(&m.afterGetUserByIDCounter) < 1 {
		m.t.Error("Expected call to StoreMock.GetUserByID")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StoreMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteStateInspect()

		m.MinimockGetUserByIDInspect()
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
		m.MinimockGetUserByIDDone()
}
