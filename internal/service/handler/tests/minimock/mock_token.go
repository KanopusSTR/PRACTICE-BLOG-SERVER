// Code generated by http://github.com/gojuno/minimock (v3.3.13). DO NOT EDIT.

package minimock

//go:generate minimock -i server/internal/service/token.Service -o mock_token.go -n TokenMock -p minimock

import (
	"server/internal/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/golang-jwt/jwt/v5"
)

// TokenMock implements token.Service
type TokenMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateToken          func(mail string) (s1 string, err error)
	inspectFuncCreateToken   func(mail string)
	afterCreateTokenCounter  uint64
	beforeCreateTokenCounter uint64
	CreateTokenMock          mTokenMockCreateToken

	funcParseToken          func(stringToken string) (t1 models.Token, tp1 *jwt.Token, err error)
	inspectFuncParseToken   func(stringToken string)
	afterParseTokenCounter  uint64
	beforeParseTokenCounter uint64
	ParseTokenMock          mTokenMockParseToken
}

// NewTokenMock returns a mock for token.Service
func NewTokenMock(t minimock.Tester) *TokenMock {
	m := &TokenMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateTokenMock = mTokenMockCreateToken{mock: m}
	m.CreateTokenMock.callArgs = []*TokenMockCreateTokenParams{}

	m.ParseTokenMock = mTokenMockParseToken{mock: m}
	m.ParseTokenMock.callArgs = []*TokenMockParseTokenParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mTokenMockCreateToken struct {
	optional           bool
	mock               *TokenMock
	defaultExpectation *TokenMockCreateTokenExpectation
	expectations       []*TokenMockCreateTokenExpectation

	callArgs []*TokenMockCreateTokenParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// TokenMockCreateTokenExpectation specifies expectation struct of the Service.CreateToken
type TokenMockCreateTokenExpectation struct {
	mock      *TokenMock
	params    *TokenMockCreateTokenParams
	paramPtrs *TokenMockCreateTokenParamPtrs
	results   *TokenMockCreateTokenResults
	Counter   uint64
}

// TokenMockCreateTokenParams contains parameters of the Service.CreateToken
type TokenMockCreateTokenParams struct {
	mail string
}

// TokenMockCreateTokenParamPtrs contains pointers to parameters of the Service.CreateToken
type TokenMockCreateTokenParamPtrs struct {
	mail *string
}

// TokenMockCreateTokenResults contains results of the Service.CreateToken
type TokenMockCreateTokenResults struct {
	s1  string
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateToken *mTokenMockCreateToken) Optional() *mTokenMockCreateToken {
	mmCreateToken.optional = true
	return mmCreateToken
}

// Expect sets up expected params for Service.CreateToken
func (mmCreateToken *mTokenMockCreateToken) Expect(mail string) *mTokenMockCreateToken {
	if mmCreateToken.mock.funcCreateToken != nil {
		mmCreateToken.mock.t.Fatalf("TokenMock.CreateToken mock is already set by Set")
	}

	if mmCreateToken.defaultExpectation == nil {
		mmCreateToken.defaultExpectation = &TokenMockCreateTokenExpectation{}
	}

	if mmCreateToken.defaultExpectation.paramPtrs != nil {
		mmCreateToken.mock.t.Fatalf("TokenMock.CreateToken mock is already set by ExpectParams functions")
	}

	mmCreateToken.defaultExpectation.params = &TokenMockCreateTokenParams{mail}
	for _, e := range mmCreateToken.expectations {
		if minimock.Equal(e.params, mmCreateToken.defaultExpectation.params) {
			mmCreateToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateToken.defaultExpectation.params)
		}
	}

	return mmCreateToken
}

// ExpectMailParam1 sets up expected param mail for Service.CreateToken
func (mmCreateToken *mTokenMockCreateToken) ExpectMailParam1(mail string) *mTokenMockCreateToken {
	if mmCreateToken.mock.funcCreateToken != nil {
		mmCreateToken.mock.t.Fatalf("TokenMock.CreateToken mock is already set by Set")
	}

	if mmCreateToken.defaultExpectation == nil {
		mmCreateToken.defaultExpectation = &TokenMockCreateTokenExpectation{}
	}

	if mmCreateToken.defaultExpectation.params != nil {
		mmCreateToken.mock.t.Fatalf("TokenMock.CreateToken mock is already set by Expect")
	}

	if mmCreateToken.defaultExpectation.paramPtrs == nil {
		mmCreateToken.defaultExpectation.paramPtrs = &TokenMockCreateTokenParamPtrs{}
	}
	mmCreateToken.defaultExpectation.paramPtrs.mail = &mail

	return mmCreateToken
}

// Inspect accepts an inspector function that has same arguments as the Service.CreateToken
func (mmCreateToken *mTokenMockCreateToken) Inspect(f func(mail string)) *mTokenMockCreateToken {
	if mmCreateToken.mock.inspectFuncCreateToken != nil {
		mmCreateToken.mock.t.Fatalf("Inspect function is already set for TokenMock.CreateToken")
	}

	mmCreateToken.mock.inspectFuncCreateToken = f

	return mmCreateToken
}

// Return sets up results that will be returned by Service.CreateToken
func (mmCreateToken *mTokenMockCreateToken) Return(s1 string, err error) *TokenMock {
	if mmCreateToken.mock.funcCreateToken != nil {
		mmCreateToken.mock.t.Fatalf("TokenMock.CreateToken mock is already set by Set")
	}

	if mmCreateToken.defaultExpectation == nil {
		mmCreateToken.defaultExpectation = &TokenMockCreateTokenExpectation{mock: mmCreateToken.mock}
	}
	mmCreateToken.defaultExpectation.results = &TokenMockCreateTokenResults{s1, err}
	return mmCreateToken.mock
}

// Set uses given function f to mock the Service.CreateToken method
func (mmCreateToken *mTokenMockCreateToken) Set(f func(mail string) (s1 string, err error)) *TokenMock {
	if mmCreateToken.defaultExpectation != nil {
		mmCreateToken.mock.t.Fatalf("Default expectation is already set for the Service.CreateToken method")
	}

	if len(mmCreateToken.expectations) > 0 {
		mmCreateToken.mock.t.Fatalf("Some expectations are already set for the Service.CreateToken method")
	}

	mmCreateToken.mock.funcCreateToken = f
	return mmCreateToken.mock
}

// When sets expectation for the Service.CreateToken which will trigger the result defined by the following
// Then helper
func (mmCreateToken *mTokenMockCreateToken) When(mail string) *TokenMockCreateTokenExpectation {
	if mmCreateToken.mock.funcCreateToken != nil {
		mmCreateToken.mock.t.Fatalf("TokenMock.CreateToken mock is already set by Set")
	}

	expectation := &TokenMockCreateTokenExpectation{
		mock:   mmCreateToken.mock,
		params: &TokenMockCreateTokenParams{mail},
	}
	mmCreateToken.expectations = append(mmCreateToken.expectations, expectation)
	return expectation
}

// Then sets up Service.CreateToken return parameters for the expectation previously defined by the When method
func (e *TokenMockCreateTokenExpectation) Then(s1 string, err error) *TokenMock {
	e.results = &TokenMockCreateTokenResults{s1, err}
	return e.mock
}

// Times sets number of times Service.CreateToken should be invoked
func (mmCreateToken *mTokenMockCreateToken) Times(n uint64) *mTokenMockCreateToken {
	if n == 0 {
		mmCreateToken.mock.t.Fatalf("Times of TokenMock.CreateToken mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateToken.expectedInvocations, n)
	return mmCreateToken
}

func (mmCreateToken *mTokenMockCreateToken) invocationsDone() bool {
	if len(mmCreateToken.expectations) == 0 && mmCreateToken.defaultExpectation == nil && mmCreateToken.mock.funcCreateToken == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateToken.mock.afterCreateTokenCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateToken.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateToken implements token.Service
func (mmCreateToken *TokenMock) CreateToken(mail string) (s1 string, err error) {
	mm_atomic.AddUint64(&mmCreateToken.beforeCreateTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateToken.afterCreateTokenCounter, 1)

	if mmCreateToken.inspectFuncCreateToken != nil {
		mmCreateToken.inspectFuncCreateToken(mail)
	}

	mm_params := TokenMockCreateTokenParams{mail}

	// Record call args
	mmCreateToken.CreateTokenMock.mutex.Lock()
	mmCreateToken.CreateTokenMock.callArgs = append(mmCreateToken.CreateTokenMock.callArgs, &mm_params)
	mmCreateToken.CreateTokenMock.mutex.Unlock()

	for _, e := range mmCreateToken.CreateTokenMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmCreateToken.CreateTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateToken.CreateTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateToken.CreateTokenMock.defaultExpectation.params
		mm_want_ptrs := mmCreateToken.CreateTokenMock.defaultExpectation.paramPtrs

		mm_got := TokenMockCreateTokenParams{mail}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.mail != nil && !minimock.Equal(*mm_want_ptrs.mail, mm_got.mail) {
				mmCreateToken.t.Errorf("TokenMock.CreateToken got unexpected parameter mail, want: %#v, got: %#v%s\n", *mm_want_ptrs.mail, mm_got.mail, minimock.Diff(*mm_want_ptrs.mail, mm_got.mail))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateToken.t.Errorf("TokenMock.CreateToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateToken.CreateTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateToken.t.Fatal("No results are set for the TokenMock.CreateToken")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmCreateToken.funcCreateToken != nil {
		return mmCreateToken.funcCreateToken(mail)
	}
	mmCreateToken.t.Fatalf("Unexpected call to TokenMock.CreateToken. %v", mail)
	return
}

// CreateTokenAfterCounter returns a count of finished TokenMock.CreateToken invocations
func (mmCreateToken *TokenMock) CreateTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateToken.afterCreateTokenCounter)
}

// CreateTokenBeforeCounter returns a count of TokenMock.CreateToken invocations
func (mmCreateToken *TokenMock) CreateTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateToken.beforeCreateTokenCounter)
}

// Calls returns a list of arguments used in each call to TokenMock.CreateToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateToken *mTokenMockCreateToken) Calls() []*TokenMockCreateTokenParams {
	mmCreateToken.mutex.RLock()

	argCopy := make([]*TokenMockCreateTokenParams, len(mmCreateToken.callArgs))
	copy(argCopy, mmCreateToken.callArgs)

	mmCreateToken.mutex.RUnlock()

	return argCopy
}

// MinimockCreateTokenDone returns true if the count of the CreateToken invocations corresponds
// the number of defined expectations
func (m *TokenMock) MinimockCreateTokenDone() bool {
	if m.CreateTokenMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateTokenMock.invocationsDone()
}

// MinimockCreateTokenInspect logs each unmet expectation
func (m *TokenMock) MinimockCreateTokenInspect() {
	for _, e := range m.CreateTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TokenMock.CreateToken with params: %#v", *e.params)
		}
	}

	afterCreateTokenCounter := mm_atomic.LoadUint64(&m.afterCreateTokenCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateTokenMock.defaultExpectation != nil && afterCreateTokenCounter < 1 {
		if m.CreateTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TokenMock.CreateToken")
		} else {
			m.t.Errorf("Expected call to TokenMock.CreateToken with params: %#v", *m.CreateTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateToken != nil && afterCreateTokenCounter < 1 {
		m.t.Error("Expected call to TokenMock.CreateToken")
	}

	if !m.CreateTokenMock.invocationsDone() && afterCreateTokenCounter > 0 {
		m.t.Errorf("Expected %d calls to TokenMock.CreateToken but found %d calls",
			mm_atomic.LoadUint64(&m.CreateTokenMock.expectedInvocations), afterCreateTokenCounter)
	}
}

type mTokenMockParseToken struct {
	optional           bool
	mock               *TokenMock
	defaultExpectation *TokenMockParseTokenExpectation
	expectations       []*TokenMockParseTokenExpectation

	callArgs []*TokenMockParseTokenParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// TokenMockParseTokenExpectation specifies expectation struct of the Service.ParseToken
type TokenMockParseTokenExpectation struct {
	mock      *TokenMock
	params    *TokenMockParseTokenParams
	paramPtrs *TokenMockParseTokenParamPtrs
	results   *TokenMockParseTokenResults
	Counter   uint64
}

// TokenMockParseTokenParams contains parameters of the Service.ParseToken
type TokenMockParseTokenParams struct {
	stringToken string
}

// TokenMockParseTokenParamPtrs contains pointers to parameters of the Service.ParseToken
type TokenMockParseTokenParamPtrs struct {
	stringToken *string
}

// TokenMockParseTokenResults contains results of the Service.ParseToken
type TokenMockParseTokenResults struct {
	t1  models.Token
	tp1 *jwt.Token
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmParseToken *mTokenMockParseToken) Optional() *mTokenMockParseToken {
	mmParseToken.optional = true
	return mmParseToken
}

// Expect sets up expected params for Service.ParseToken
func (mmParseToken *mTokenMockParseToken) Expect(stringToken string) *mTokenMockParseToken {
	if mmParseToken.mock.funcParseToken != nil {
		mmParseToken.mock.t.Fatalf("TokenMock.ParseToken mock is already set by Set")
	}

	if mmParseToken.defaultExpectation == nil {
		mmParseToken.defaultExpectation = &TokenMockParseTokenExpectation{}
	}

	if mmParseToken.defaultExpectation.paramPtrs != nil {
		mmParseToken.mock.t.Fatalf("TokenMock.ParseToken mock is already set by ExpectParams functions")
	}

	mmParseToken.defaultExpectation.params = &TokenMockParseTokenParams{stringToken}
	for _, e := range mmParseToken.expectations {
		if minimock.Equal(e.params, mmParseToken.defaultExpectation.params) {
			mmParseToken.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmParseToken.defaultExpectation.params)
		}
	}

	return mmParseToken
}

// ExpectStringTokenParam1 sets up expected param stringToken for Service.ParseToken
func (mmParseToken *mTokenMockParseToken) ExpectStringTokenParam1(stringToken string) *mTokenMockParseToken {
	if mmParseToken.mock.funcParseToken != nil {
		mmParseToken.mock.t.Fatalf("TokenMock.ParseToken mock is already set by Set")
	}

	if mmParseToken.defaultExpectation == nil {
		mmParseToken.defaultExpectation = &TokenMockParseTokenExpectation{}
	}

	if mmParseToken.defaultExpectation.params != nil {
		mmParseToken.mock.t.Fatalf("TokenMock.ParseToken mock is already set by Expect")
	}

	if mmParseToken.defaultExpectation.paramPtrs == nil {
		mmParseToken.defaultExpectation.paramPtrs = &TokenMockParseTokenParamPtrs{}
	}
	mmParseToken.defaultExpectation.paramPtrs.stringToken = &stringToken

	return mmParseToken
}

// Inspect accepts an inspector function that has same arguments as the Service.ParseToken
func (mmParseToken *mTokenMockParseToken) Inspect(f func(stringToken string)) *mTokenMockParseToken {
	if mmParseToken.mock.inspectFuncParseToken != nil {
		mmParseToken.mock.t.Fatalf("Inspect function is already set for TokenMock.ParseToken")
	}

	mmParseToken.mock.inspectFuncParseToken = f

	return mmParseToken
}

// Return sets up results that will be returned by Service.ParseToken
func (mmParseToken *mTokenMockParseToken) Return(t1 models.Token, tp1 *jwt.Token, err error) *TokenMock {
	if mmParseToken.mock.funcParseToken != nil {
		mmParseToken.mock.t.Fatalf("TokenMock.ParseToken mock is already set by Set")
	}

	if mmParseToken.defaultExpectation == nil {
		mmParseToken.defaultExpectation = &TokenMockParseTokenExpectation{mock: mmParseToken.mock}
	}
	mmParseToken.defaultExpectation.results = &TokenMockParseTokenResults{t1, tp1, err}
	return mmParseToken.mock
}

// Set uses given function f to mock the Service.ParseToken method
func (mmParseToken *mTokenMockParseToken) Set(f func(stringToken string) (t1 models.Token, tp1 *jwt.Token, err error)) *TokenMock {
	if mmParseToken.defaultExpectation != nil {
		mmParseToken.mock.t.Fatalf("Default expectation is already set for the Service.ParseToken method")
	}

	if len(mmParseToken.expectations) > 0 {
		mmParseToken.mock.t.Fatalf("Some expectations are already set for the Service.ParseToken method")
	}

	mmParseToken.mock.funcParseToken = f
	return mmParseToken.mock
}

// When sets expectation for the Service.ParseToken which will trigger the result defined by the following
// Then helper
func (mmParseToken *mTokenMockParseToken) When(stringToken string) *TokenMockParseTokenExpectation {
	if mmParseToken.mock.funcParseToken != nil {
		mmParseToken.mock.t.Fatalf("TokenMock.ParseToken mock is already set by Set")
	}

	expectation := &TokenMockParseTokenExpectation{
		mock:   mmParseToken.mock,
		params: &TokenMockParseTokenParams{stringToken},
	}
	mmParseToken.expectations = append(mmParseToken.expectations, expectation)
	return expectation
}

// Then sets up Service.ParseToken return parameters for the expectation previously defined by the When method
func (e *TokenMockParseTokenExpectation) Then(t1 models.Token, tp1 *jwt.Token, err error) *TokenMock {
	e.results = &TokenMockParseTokenResults{t1, tp1, err}
	return e.mock
}

// Times sets number of times Service.ParseToken should be invoked
func (mmParseToken *mTokenMockParseToken) Times(n uint64) *mTokenMockParseToken {
	if n == 0 {
		mmParseToken.mock.t.Fatalf("Times of TokenMock.ParseToken mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmParseToken.expectedInvocations, n)
	return mmParseToken
}

func (mmParseToken *mTokenMockParseToken) invocationsDone() bool {
	if len(mmParseToken.expectations) == 0 && mmParseToken.defaultExpectation == nil && mmParseToken.mock.funcParseToken == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmParseToken.mock.afterParseTokenCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmParseToken.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// ParseToken implements token.Service
func (mmParseToken *TokenMock) ParseToken(stringToken string) (t1 models.Token, tp1 *jwt.Token, err error) {
	mm_atomic.AddUint64(&mmParseToken.beforeParseTokenCounter, 1)
	defer mm_atomic.AddUint64(&mmParseToken.afterParseTokenCounter, 1)

	if mmParseToken.inspectFuncParseToken != nil {
		mmParseToken.inspectFuncParseToken(stringToken)
	}

	mm_params := TokenMockParseTokenParams{stringToken}

	// Record call args
	mmParseToken.ParseTokenMock.mutex.Lock()
	mmParseToken.ParseTokenMock.callArgs = append(mmParseToken.ParseTokenMock.callArgs, &mm_params)
	mmParseToken.ParseTokenMock.mutex.Unlock()

	for _, e := range mmParseToken.ParseTokenMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.t1, e.results.tp1, e.results.err
		}
	}

	if mmParseToken.ParseTokenMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmParseToken.ParseTokenMock.defaultExpectation.Counter, 1)
		mm_want := mmParseToken.ParseTokenMock.defaultExpectation.params
		mm_want_ptrs := mmParseToken.ParseTokenMock.defaultExpectation.paramPtrs

		mm_got := TokenMockParseTokenParams{stringToken}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.stringToken != nil && !minimock.Equal(*mm_want_ptrs.stringToken, mm_got.stringToken) {
				mmParseToken.t.Errorf("TokenMock.ParseToken got unexpected parameter stringToken, want: %#v, got: %#v%s\n", *mm_want_ptrs.stringToken, mm_got.stringToken, minimock.Diff(*mm_want_ptrs.stringToken, mm_got.stringToken))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmParseToken.t.Errorf("TokenMock.ParseToken got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmParseToken.ParseTokenMock.defaultExpectation.results
		if mm_results == nil {
			mmParseToken.t.Fatal("No results are set for the TokenMock.ParseToken")
		}
		return (*mm_results).t1, (*mm_results).tp1, (*mm_results).err
	}
	if mmParseToken.funcParseToken != nil {
		return mmParseToken.funcParseToken(stringToken)
	}
	mmParseToken.t.Fatalf("Unexpected call to TokenMock.ParseToken. %v", stringToken)
	return
}

// ParseTokenAfterCounter returns a count of finished TokenMock.ParseToken invocations
func (mmParseToken *TokenMock) ParseTokenAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmParseToken.afterParseTokenCounter)
}

// ParseTokenBeforeCounter returns a count of TokenMock.ParseToken invocations
func (mmParseToken *TokenMock) ParseTokenBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmParseToken.beforeParseTokenCounter)
}

// Calls returns a list of arguments used in each call to TokenMock.ParseToken.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmParseToken *mTokenMockParseToken) Calls() []*TokenMockParseTokenParams {
	mmParseToken.mutex.RLock()

	argCopy := make([]*TokenMockParseTokenParams, len(mmParseToken.callArgs))
	copy(argCopy, mmParseToken.callArgs)

	mmParseToken.mutex.RUnlock()

	return argCopy
}

// MinimockParseTokenDone returns true if the count of the ParseToken invocations corresponds
// the number of defined expectations
func (m *TokenMock) MinimockParseTokenDone() bool {
	if m.ParseTokenMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.ParseTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.ParseTokenMock.invocationsDone()
}

// MinimockParseTokenInspect logs each unmet expectation
func (m *TokenMock) MinimockParseTokenInspect() {
	for _, e := range m.ParseTokenMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TokenMock.ParseToken with params: %#v", *e.params)
		}
	}

	afterParseTokenCounter := mm_atomic.LoadUint64(&m.afterParseTokenCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.ParseTokenMock.defaultExpectation != nil && afterParseTokenCounter < 1 {
		if m.ParseTokenMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TokenMock.ParseToken")
		} else {
			m.t.Errorf("Expected call to TokenMock.ParseToken with params: %#v", *m.ParseTokenMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcParseToken != nil && afterParseTokenCounter < 1 {
		m.t.Error("Expected call to TokenMock.ParseToken")
	}

	if !m.ParseTokenMock.invocationsDone() && afterParseTokenCounter > 0 {
		m.t.Errorf("Expected %d calls to TokenMock.ParseToken but found %d calls",
			mm_atomic.LoadUint64(&m.ParseTokenMock.expectedInvocations), afterParseTokenCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TokenMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateTokenInspect()

			m.MinimockParseTokenInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TokenMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TokenMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateTokenDone() &&
		m.MinimockParseTokenDone()
}