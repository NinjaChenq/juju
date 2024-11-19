// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/charm/services (interfaces: StateBackend,Storage,UploadedCharm)
//
// Generated by this command:
//
//	mockgen -typed -package services -destination interface_mocks_test.go github.com/juju/juju/internal/charm/services StateBackend,Storage,UploadedCharm
//

// Package services is a generated GoMock package.
package services

import (
	context "context"
	io "io"
	reflect "reflect"

	objectstore "github.com/juju/juju/core/objectstore"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockStateBackend is a mock of StateBackend interface.
type MockStateBackend struct {
	ctrl     *gomock.Controller
	recorder *MockStateBackendMockRecorder
}

// MockStateBackendMockRecorder is the mock recorder for MockStateBackend.
type MockStateBackendMockRecorder struct {
	mock *MockStateBackend
}

// NewMockStateBackend creates a new mock instance.
func NewMockStateBackend(ctrl *gomock.Controller) *MockStateBackend {
	mock := &MockStateBackend{ctrl: ctrl}
	mock.recorder = &MockStateBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateBackend) EXPECT() *MockStateBackendMockRecorder {
	return m.recorder
}

// ModelUUID mocks base method.
func (m *MockStateBackend) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockStateBackendMockRecorder) ModelUUID() *MockStateBackendModelUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockStateBackend)(nil).ModelUUID))
	return &MockStateBackendModelUUIDCall{Call: call}
}

// MockStateBackendModelUUIDCall wrap *gomock.Call
type MockStateBackendModelUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateBackendModelUUIDCall) Return(arg0 string) *MockStateBackendModelUUIDCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateBackendModelUUIDCall) Do(f func() string) *MockStateBackendModelUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateBackendModelUUIDCall) DoAndReturn(f func() string) *MockStateBackendModelUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PrepareCharmUpload mocks base method.
func (m *MockStateBackend) PrepareCharmUpload(arg0 string) (UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockStateBackendMockRecorder) PrepareCharmUpload(arg0 any) *MockStateBackendPrepareCharmUploadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockStateBackend)(nil).PrepareCharmUpload), arg0)
	return &MockStateBackendPrepareCharmUploadCall{Call: call}
}

// MockStateBackendPrepareCharmUploadCall wrap *gomock.Call
type MockStateBackendPrepareCharmUploadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateBackendPrepareCharmUploadCall) Return(arg0 UploadedCharm, arg1 error) *MockStateBackendPrepareCharmUploadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateBackendPrepareCharmUploadCall) Do(f func(string) (UploadedCharm, error)) *MockStateBackendPrepareCharmUploadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateBackendPrepareCharmUploadCall) DoAndReturn(f func(string) (UploadedCharm, error)) *MockStateBackendPrepareCharmUploadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUploadedCharm mocks base method.
func (m *MockStateBackend) UpdateUploadedCharm(arg0 state.CharmInfo) (UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockStateBackendMockRecorder) UpdateUploadedCharm(arg0 any) *MockStateBackendUpdateUploadedCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockStateBackend)(nil).UpdateUploadedCharm), arg0)
	return &MockStateBackendUpdateUploadedCharmCall{Call: call}
}

// MockStateBackendUpdateUploadedCharmCall wrap *gomock.Call
type MockStateBackendUpdateUploadedCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateBackendUpdateUploadedCharmCall) Return(arg0 UploadedCharm, arg1 error) *MockStateBackendUpdateUploadedCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateBackendUpdateUploadedCharmCall) Do(f func(state.CharmInfo) (UploadedCharm, error)) *MockStateBackendUpdateUploadedCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateBackendUpdateUploadedCharmCall) DoAndReturn(f func(state.CharmInfo) (UploadedCharm, error)) *MockStateBackendUpdateUploadedCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Put mocks base method.
func (m *MockStorage) Put(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 int64) (objectstore.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(objectstore.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockStorageMockRecorder) Put(arg0, arg1, arg2, arg3 any) *MockStoragePutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStorage)(nil).Put), arg0, arg1, arg2, arg3)
	return &MockStoragePutCall{Call: call}
}

// MockStoragePutCall wrap *gomock.Call
type MockStoragePutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStoragePutCall) Return(arg0 objectstore.UUID, arg1 error) *MockStoragePutCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStoragePutCall) Do(f func(context.Context, string, io.Reader, int64) (objectstore.UUID, error)) *MockStoragePutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStoragePutCall) DoAndReturn(f func(context.Context, string, io.Reader, int64) (objectstore.UUID, error)) *MockStoragePutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Remove mocks base method.
func (m *MockStorage) Remove(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStorageMockRecorder) Remove(arg0, arg1 any) *MockStorageRemoveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStorage)(nil).Remove), arg0, arg1)
	return &MockStorageRemoveCall{Call: call}
}

// MockStorageRemoveCall wrap *gomock.Call
type MockStorageRemoveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStorageRemoveCall) Return(arg0 error) *MockStorageRemoveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStorageRemoveCall) Do(f func(context.Context, string) error) *MockStorageRemoveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStorageRemoveCall) DoAndReturn(f func(context.Context, string) error) *MockStorageRemoveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUploadedCharm is a mock of UploadedCharm interface.
type MockUploadedCharm struct {
	ctrl     *gomock.Controller
	recorder *MockUploadedCharmMockRecorder
}

// MockUploadedCharmMockRecorder is the mock recorder for MockUploadedCharm.
type MockUploadedCharmMockRecorder struct {
	mock *MockUploadedCharm
}

// NewMockUploadedCharm creates a new mock instance.
func NewMockUploadedCharm(ctrl *gomock.Controller) *MockUploadedCharm {
	mock := &MockUploadedCharm{ctrl: ctrl}
	mock.recorder = &MockUploadedCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadedCharm) EXPECT() *MockUploadedCharmMockRecorder {
	return m.recorder
}
