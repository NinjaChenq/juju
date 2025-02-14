// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/resource/modelmigration (interfaces: Coordinator,ImportService)
//
// Generated by this command:
//
//	mockgen -typed -package modelmigration -destination migrations_mock_test.go github.com/juju/juju/domain/resource/modelmigration Coordinator,ImportService
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	reflect "reflect"

	modelmigration "github.com/juju/juju/core/modelmigration"
	resource "github.com/juju/juju/domain/resource"
	gomock "go.uber.org/mock/gomock"
)

// MockCoordinator is a mock of Coordinator interface.
type MockCoordinator struct {
	ctrl     *gomock.Controller
	recorder *MockCoordinatorMockRecorder
}

// MockCoordinatorMockRecorder is the mock recorder for MockCoordinator.
type MockCoordinatorMockRecorder struct {
	mock *MockCoordinator
}

// NewMockCoordinator creates a new mock instance.
func NewMockCoordinator(ctrl *gomock.Controller) *MockCoordinator {
	mock := &MockCoordinator{ctrl: ctrl}
	mock.recorder = &MockCoordinatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoordinator) EXPECT() *MockCoordinatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCoordinator) Add(arg0 modelmigration.Operation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockCoordinatorMockRecorder) Add(arg0 any) *MockCoordinatorAddCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCoordinator)(nil).Add), arg0)
	return &MockCoordinatorAddCall{Call: call}
}

// MockCoordinatorAddCall wrap *gomock.Call
type MockCoordinatorAddCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCoordinatorAddCall) Return() *MockCoordinatorAddCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCoordinatorAddCall) Do(f func(modelmigration.Operation)) *MockCoordinatorAddCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCoordinatorAddCall) DoAndReturn(f func(modelmigration.Operation)) *MockCoordinatorAddCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockImportService is a mock of ImportService interface.
type MockImportService struct {
	ctrl     *gomock.Controller
	recorder *MockImportServiceMockRecorder
}

// MockImportServiceMockRecorder is the mock recorder for MockImportService.
type MockImportServiceMockRecorder struct {
	mock *MockImportService
}

// NewMockImportService creates a new mock instance.
func NewMockImportService(ctrl *gomock.Controller) *MockImportService {
	mock := &MockImportService{ctrl: ctrl}
	mock.recorder = &MockImportServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImportService) EXPECT() *MockImportServiceMockRecorder {
	return m.recorder
}

// DeleteImportedResources mocks base method.
func (m *MockImportService) DeleteImportedResources(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImportedResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImportedResources indicates an expected call of DeleteImportedResources.
func (mr *MockImportServiceMockRecorder) DeleteImportedResources(arg0, arg1 any) *MockImportServiceDeleteImportedResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImportedResources", reflect.TypeOf((*MockImportService)(nil).DeleteImportedResources), arg0, arg1)
	return &MockImportServiceDeleteImportedResourcesCall{Call: call}
}

// MockImportServiceDeleteImportedResourcesCall wrap *gomock.Call
type MockImportServiceDeleteImportedResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceDeleteImportedResourcesCall) Return(arg0 error) *MockImportServiceDeleteImportedResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceDeleteImportedResourcesCall) Do(f func(context.Context, []string) error) *MockImportServiceDeleteImportedResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceDeleteImportedResourcesCall) DoAndReturn(f func(context.Context, []string) error) *MockImportServiceDeleteImportedResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportResources mocks base method.
func (m *MockImportService) ImportResources(arg0 context.Context, arg1 resource.ImportResourcesArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportResources indicates an expected call of ImportResources.
func (mr *MockImportServiceMockRecorder) ImportResources(arg0, arg1 any) *MockImportServiceImportResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportResources", reflect.TypeOf((*MockImportService)(nil).ImportResources), arg0, arg1)
	return &MockImportServiceImportResourcesCall{Call: call}
}

// MockImportServiceImportResourcesCall wrap *gomock.Call
type MockImportServiceImportResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockImportServiceImportResourcesCall) Return(arg0 error) *MockImportServiceImportResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockImportServiceImportResourcesCall) Do(f func(context.Context, resource.ImportResourcesArgs) error) *MockImportServiceImportResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockImportServiceImportResourcesCall) DoAndReturn(f func(context.Context, resource.ImportResourcesArgs) error) *MockImportServiceImportResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
