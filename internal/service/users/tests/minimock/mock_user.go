// Code generated by http://github.com/gojuno/minimock (v3.3.13). DO NOT EDIT.

package minimock

//go:generate minimock -i server/internal/repo.User -o mock_user.go -n UserMock -p minimock

import (
	"server/internal/entities"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// UserMock implements repo.User
type UserMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAdd          func(user entities.User) (err error)
	inspectFuncAdd   func(user entities.User)
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mUserMockAdd

	funcGet          func(mail string) (up1 *entities.User, err error)
	inspectFuncGet   func(mail string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUserMockGet
}

// NewUserMock returns a mock for repo.User
func NewUserMock(t minimock.Tester) *UserMock {
	m := &UserMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mUserMockAdd{mock: m}
	m.AddMock.callArgs = []*UserMockAddParams{}

	m.GetMock = mUserMockGet{mock: m}
	m.GetMock.callArgs = []*UserMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserMockAdd struct {
	optional           bool
	mock               *UserMock
	defaultExpectation *UserMockAddExpectation
	expectations       []*UserMockAddExpectation

	callArgs []*UserMockAddParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserMockAddExpectation specifies expectation struct of the User.Add
type UserMockAddExpectation struct {
	mock      *UserMock
	params    *UserMockAddParams
	paramPtrs *UserMockAddParamPtrs
	results   *UserMockAddResults
	Counter   uint64
}

// UserMockAddParams contains parameters of the User.Add
type UserMockAddParams struct {
	user entities.User
}

// UserMockAddParamPtrs contains pointers to parameters of the User.Add
type UserMockAddParamPtrs struct {
	user *entities.User
}

// UserMockAddResults contains results of the User.Add
type UserMockAddResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmAdd *mUserMockAdd) Optional() *mUserMockAdd {
	mmAdd.optional = true
	return mmAdd
}

// Expect sets up expected params for User.Add
func (mmAdd *mUserMockAdd) Expect(user entities.User) *mUserMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("UserMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &UserMockAddExpectation{}
	}

	if mmAdd.defaultExpectation.paramPtrs != nil {
		mmAdd.mock.t.Fatalf("UserMock.Add mock is already set by ExpectParams functions")
	}

	mmAdd.defaultExpectation.params = &UserMockAddParams{user}
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// ExpectUserParam1 sets up expected param user for User.Add
func (mmAdd *mUserMockAdd) ExpectUserParam1(user entities.User) *mUserMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("UserMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &UserMockAddExpectation{}
	}

	if mmAdd.defaultExpectation.params != nil {
		mmAdd.mock.t.Fatalf("UserMock.Add mock is already set by Expect")
	}

	if mmAdd.defaultExpectation.paramPtrs == nil {
		mmAdd.defaultExpectation.paramPtrs = &UserMockAddParamPtrs{}
	}
	mmAdd.defaultExpectation.paramPtrs.user = &user

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the User.Add
func (mmAdd *mUserMockAdd) Inspect(f func(user entities.User)) *mUserMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for UserMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by User.Add
func (mmAdd *mUserMockAdd) Return(err error) *UserMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("UserMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &UserMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &UserMockAddResults{err}
	return mmAdd.mock
}

// Set uses given function f to mock the User.Add method
func (mmAdd *mUserMockAdd) Set(f func(user entities.User) (err error)) *UserMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the User.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the User.Add method")
	}

	mmAdd.mock.funcAdd = f
	return mmAdd.mock
}

// When sets expectation for the User.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mUserMockAdd) When(user entities.User) *UserMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("UserMock.Add mock is already set by Set")
	}

	expectation := &UserMockAddExpectation{
		mock:   mmAdd.mock,
		params: &UserMockAddParams{user},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up User.Add return parameters for the expectation previously defined by the When method
func (e *UserMockAddExpectation) Then(err error) *UserMock {
	e.results = &UserMockAddResults{err}
	return e.mock
}

// Times sets number of times User.Add should be invoked
func (mmAdd *mUserMockAdd) Times(n uint64) *mUserMockAdd {
	if n == 0 {
		mmAdd.mock.t.Fatalf("Times of UserMock.Add mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmAdd.expectedInvocations, n)
	return mmAdd
}

func (mmAdd *mUserMockAdd) invocationsDone() bool {
	if len(mmAdd.expectations) == 0 && mmAdd.defaultExpectation == nil && mmAdd.mock.funcAdd == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmAdd.mock.afterAddCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmAdd.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Add implements repo.User
func (mmAdd *UserMock) Add(user entities.User) (err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(user)
	}

	mm_params := UserMockAddParams{user}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, &mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_want_ptrs := mmAdd.AddMock.defaultExpectation.paramPtrs

		mm_got := UserMockAddParams{user}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmAdd.t.Errorf("UserMock.Add got unexpected parameter user, want: %#v, got: %#v%s\n", *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("UserMock.Add got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the UserMock.Add")
		}
		return (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(user)
	}
	mmAdd.t.Fatalf("Unexpected call to UserMock.Add. %v", user)
	return
}

// AddAfterCounter returns a count of finished UserMock.Add invocations
func (mmAdd *UserMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of UserMock.Add invocations
func (mmAdd *UserMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to UserMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mUserMockAdd) Calls() []*UserMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*UserMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *UserMock) MinimockAddDone() bool {
	if m.AddMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.AddMock.invocationsDone()
}

// MinimockAddInspect logs each unmet expectation
func (m *UserMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserMock.Add with params: %#v", *e.params)
		}
	}

	afterAddCounter := mm_atomic.LoadUint64(&m.afterAddCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && afterAddCounter < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserMock.Add")
		} else {
			m.t.Errorf("Expected call to UserMock.Add with params: %#v", *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && afterAddCounter < 1 {
		m.t.Error("Expected call to UserMock.Add")
	}

	if !m.AddMock.invocationsDone() && afterAddCounter > 0 {
		m.t.Errorf("Expected %d calls to UserMock.Add but found %d calls",
			mm_atomic.LoadUint64(&m.AddMock.expectedInvocations), afterAddCounter)
	}
}

type mUserMockGet struct {
	optional           bool
	mock               *UserMock
	defaultExpectation *UserMockGetExpectation
	expectations       []*UserMockGetExpectation

	callArgs []*UserMockGetParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UserMockGetExpectation specifies expectation struct of the User.Get
type UserMockGetExpectation struct {
	mock      *UserMock
	params    *UserMockGetParams
	paramPtrs *UserMockGetParamPtrs
	results   *UserMockGetResults
	Counter   uint64
}

// UserMockGetParams contains parameters of the User.Get
type UserMockGetParams struct {
	mail string
}

// UserMockGetParamPtrs contains pointers to parameters of the User.Get
type UserMockGetParamPtrs struct {
	mail *string
}

// UserMockGetResults contains results of the User.Get
type UserMockGetResults struct {
	up1 *entities.User
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mUserMockGet) Optional() *mUserMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for User.Get
func (mmGet *mUserMockGet) Expect(mail string) *mUserMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("UserMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &UserMockGetParams{mail}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectMailParam1 sets up expected param mail for User.Get
func (mmGet *mUserMockGet) ExpectMailParam1(mail string) *mUserMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UserMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UserMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.mail = &mail

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the User.Get
func (mmGet *mUserMockGet) Inspect(f func(mail string)) *mUserMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UserMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by User.Get
func (mmGet *mUserMockGet) Return(up1 *entities.User, err error) *UserMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UserMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UserMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the User.Get method
func (mmGet *mUserMockGet) Set(f func(mail string) (up1 *entities.User, err error)) *UserMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the User.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the User.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the User.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUserMockGet) When(mail string) *UserMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UserMock.Get mock is already set by Set")
	}

	expectation := &UserMockGetExpectation{
		mock:   mmGet.mock,
		params: &UserMockGetParams{mail},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up User.Get return parameters for the expectation previously defined by the When method
func (e *UserMockGetExpectation) Then(up1 *entities.User, err error) *UserMock {
	e.results = &UserMockGetResults{up1, err}
	return e.mock
}

// Times sets number of times User.Get should be invoked
func (mmGet *mUserMockGet) Times(n uint64) *mUserMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of UserMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	return mmGet
}

func (mmGet *mUserMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements repo.User
func (mmGet *UserMock) Get(mail string) (up1 *entities.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(mail)
	}

	mm_params := UserMockGetParams{mail}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := UserMockGetParams{mail}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.mail != nil && !minimock.Equal(*mm_want_ptrs.mail, mm_got.mail) {
				mmGet.t.Errorf("UserMock.Get got unexpected parameter mail, want: %#v, got: %#v%s\n", *mm_want_ptrs.mail, mm_got.mail, minimock.Diff(*mm_want_ptrs.mail, mm_got.mail))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UserMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UserMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(mail)
	}
	mmGet.t.Fatalf("Unexpected call to UserMock.Get. %v", mail)
	return
}

// GetAfterCounter returns a count of finished UserMock.Get invocations
func (mmGet *UserMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UserMock.Get invocations
func (mmGet *UserMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UserMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUserMockGet) Calls() []*UserMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UserMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UserMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *UserMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserMock.Get with params: %#v", *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserMock.Get")
		} else {
			m.t.Errorf("Expected call to UserMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Error("Expected call to UserMock.Get")
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to UserMock.Get but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), afterGetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddInspect()

			m.MinimockGetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone() &&
		m.MinimockGetDone()
}
