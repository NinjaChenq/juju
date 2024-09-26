// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain (interfaces: LeaseChecker)
//
// Generated by this command:
//
//	mockgen -typed -package domain -destination leasechecker_mock_test.go github.com/juju/juju/domain LeaseChecker
//

// Package domain is a generated GoMock package.
package domain

import (
	context "context"
	reflect "reflect"

	lease "github.com/juju/juju/core/lease"
	gomock "go.uber.org/mock/gomock"
)

// MockLeaseChecker is a mock of LeaseChecker interface.
type MockLeaseChecker struct {
	ctrl     *gomock.Controller
	recorder *MockLeaseCheckerMockRecorder
}

// MockLeaseCheckerMockRecorder is the mock recorder for MockLeaseChecker.
type MockLeaseCheckerMockRecorder struct {
	mock *MockLeaseChecker
}

// NewMockLeaseChecker creates a new mock instance.
func NewMockLeaseChecker(ctrl *gomock.Controller) *MockLeaseChecker {
	mock := &MockLeaseChecker{ctrl: ctrl}
	mock.recorder = &MockLeaseCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeaseChecker) EXPECT() *MockLeaseCheckerMockRecorder {
	return m.recorder
}

// Token mocks base method.
func (m *MockLeaseChecker) Token(arg0, arg1 string) lease.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", arg0, arg1)
	ret0, _ := ret[0].(lease.Token)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockLeaseCheckerMockRecorder) Token(arg0, arg1 any) *MockLeaseCheckerTokenCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockLeaseChecker)(nil).Token), arg0, arg1)
	return &MockLeaseCheckerTokenCall{Call: call}
}

// MockLeaseCheckerTokenCall wrap *gomock.Call
type MockLeaseCheckerTokenCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseCheckerTokenCall) Return(arg0 lease.Token) *MockLeaseCheckerTokenCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseCheckerTokenCall) Do(f func(string, string) lease.Token) *MockLeaseCheckerTokenCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseCheckerTokenCall) DoAndReturn(f func(string, string) lease.Token) *MockLeaseCheckerTokenCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WaitUntilExpired mocks base method.
func (m *MockLeaseChecker) WaitUntilExpired(arg0 context.Context, arg1 string, arg2 chan<- struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilExpired", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilExpired indicates an expected call of WaitUntilExpired.
func (mr *MockLeaseCheckerMockRecorder) WaitUntilExpired(arg0, arg1, arg2 any) *MockLeaseCheckerWaitUntilExpiredCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilExpired", reflect.TypeOf((*MockLeaseChecker)(nil).WaitUntilExpired), arg0, arg1, arg2)
	return &MockLeaseCheckerWaitUntilExpiredCall{Call: call}
}

// MockLeaseCheckerWaitUntilExpiredCall wrap *gomock.Call
type MockLeaseCheckerWaitUntilExpiredCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockLeaseCheckerWaitUntilExpiredCall) Return(arg0 error) *MockLeaseCheckerWaitUntilExpiredCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockLeaseCheckerWaitUntilExpiredCall) Do(f func(context.Context, string, chan<- struct{}) error) *MockLeaseCheckerWaitUntilExpiredCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockLeaseCheckerWaitUntilExpiredCall) DoAndReturn(f func(context.Context, string, chan<- struct{}) error) *MockLeaseCheckerWaitUntilExpiredCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
