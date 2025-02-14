// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/resource/service (interfaces: State,ResourceStoreGetter)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination package_mock_test.go github.com/juju/juju/domain/resource/service State,ResourceStoreGetter
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	application "github.com/juju/juju/core/application"
	resource "github.com/juju/juju/core/resource"
	store "github.com/juju/juju/core/resource/store"
	unit "github.com/juju/juju/core/unit"
	resource0 "github.com/juju/juju/domain/resource"
	resource1 "github.com/juju/juju/internal/charm/resource"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// AddResourcesBeforeApplication mocks base method.
func (m *MockState) AddResourcesBeforeApplication(arg0 context.Context, arg1 resource0.AddResourcesBeforeApplicationArgs) ([]resource.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResourcesBeforeApplication", arg0, arg1)
	ret0, _ := ret[0].([]resource.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddResourcesBeforeApplication indicates an expected call of AddResourcesBeforeApplication.
func (mr *MockStateMockRecorder) AddResourcesBeforeApplication(arg0, arg1 any) *MockStateAddResourcesBeforeApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResourcesBeforeApplication", reflect.TypeOf((*MockState)(nil).AddResourcesBeforeApplication), arg0, arg1)
	return &MockStateAddResourcesBeforeApplicationCall{Call: call}
}

// MockStateAddResourcesBeforeApplicationCall wrap *gomock.Call
type MockStateAddResourcesBeforeApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateAddResourcesBeforeApplicationCall) Return(arg0 []resource.UUID, arg1 error) *MockStateAddResourcesBeforeApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateAddResourcesBeforeApplicationCall) Do(f func(context.Context, resource0.AddResourcesBeforeApplicationArgs) ([]resource.UUID, error)) *MockStateAddResourcesBeforeApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateAddResourcesBeforeApplicationCall) DoAndReturn(f func(context.Context, resource0.AddResourcesBeforeApplicationArgs) ([]resource.UUID, error)) *MockStateAddResourcesBeforeApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteApplicationResources mocks base method.
func (m *MockState) DeleteApplicationResources(arg0 context.Context, arg1 application.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplicationResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplicationResources indicates an expected call of DeleteApplicationResources.
func (mr *MockStateMockRecorder) DeleteApplicationResources(arg0, arg1 any) *MockStateDeleteApplicationResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplicationResources", reflect.TypeOf((*MockState)(nil).DeleteApplicationResources), arg0, arg1)
	return &MockStateDeleteApplicationResourcesCall{Call: call}
}

// MockStateDeleteApplicationResourcesCall wrap *gomock.Call
type MockStateDeleteApplicationResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateDeleteApplicationResourcesCall) Return(arg0 error) *MockStateDeleteApplicationResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateDeleteApplicationResourcesCall) Do(f func(context.Context, application.ID) error) *MockStateDeleteApplicationResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateDeleteApplicationResourcesCall) DoAndReturn(f func(context.Context, application.ID) error) *MockStateDeleteApplicationResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteImportedResources mocks base method.
func (m *MockState) DeleteImportedResources(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImportedResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImportedResources indicates an expected call of DeleteImportedResources.
func (mr *MockStateMockRecorder) DeleteImportedResources(arg0, arg1 any) *MockStateDeleteImportedResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImportedResources", reflect.TypeOf((*MockState)(nil).DeleteImportedResources), arg0, arg1)
	return &MockStateDeleteImportedResourcesCall{Call: call}
}

// MockStateDeleteImportedResourcesCall wrap *gomock.Call
type MockStateDeleteImportedResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateDeleteImportedResourcesCall) Return(arg0 error) *MockStateDeleteImportedResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateDeleteImportedResourcesCall) Do(f func(context.Context, []string) error) *MockStateDeleteImportedResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateDeleteImportedResourcesCall) DoAndReturn(f func(context.Context, []string) error) *MockStateDeleteImportedResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteResourcesAddedBeforeApplication mocks base method.
func (m *MockState) DeleteResourcesAddedBeforeApplication(arg0 context.Context, arg1 []resource.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResourcesAddedBeforeApplication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResourcesAddedBeforeApplication indicates an expected call of DeleteResourcesAddedBeforeApplication.
func (mr *MockStateMockRecorder) DeleteResourcesAddedBeforeApplication(arg0, arg1 any) *MockStateDeleteResourcesAddedBeforeApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourcesAddedBeforeApplication", reflect.TypeOf((*MockState)(nil).DeleteResourcesAddedBeforeApplication), arg0, arg1)
	return &MockStateDeleteResourcesAddedBeforeApplicationCall{Call: call}
}

// MockStateDeleteResourcesAddedBeforeApplicationCall wrap *gomock.Call
type MockStateDeleteResourcesAddedBeforeApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateDeleteResourcesAddedBeforeApplicationCall) Return(arg0 error) *MockStateDeleteResourcesAddedBeforeApplicationCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateDeleteResourcesAddedBeforeApplicationCall) Do(f func(context.Context, []resource.UUID) error) *MockStateDeleteResourcesAddedBeforeApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateDeleteResourcesAddedBeforeApplicationCall) DoAndReturn(f func(context.Context, []resource.UUID) error) *MockStateDeleteResourcesAddedBeforeApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteUnitResources mocks base method.
func (m *MockState) DeleteUnitResources(arg0 context.Context, arg1 unit.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUnitResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnitResources indicates an expected call of DeleteUnitResources.
func (mr *MockStateMockRecorder) DeleteUnitResources(arg0, arg1 any) *MockStateDeleteUnitResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnitResources", reflect.TypeOf((*MockState)(nil).DeleteUnitResources), arg0, arg1)
	return &MockStateDeleteUnitResourcesCall{Call: call}
}

// MockStateDeleteUnitResourcesCall wrap *gomock.Call
type MockStateDeleteUnitResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateDeleteUnitResourcesCall) Return(arg0 error) *MockStateDeleteUnitResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateDeleteUnitResourcesCall) Do(f func(context.Context, unit.UUID) error) *MockStateDeleteUnitResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateDeleteUnitResourcesCall) DoAndReturn(f func(context.Context, unit.UUID) error) *MockStateDeleteUnitResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetApplicationResourceID mocks base method.
func (m *MockState) GetApplicationResourceID(arg0 context.Context, arg1 resource0.GetApplicationResourceIDArgs) (resource.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationResourceID", arg0, arg1)
	ret0, _ := ret[0].(resource.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationResourceID indicates an expected call of GetApplicationResourceID.
func (mr *MockStateMockRecorder) GetApplicationResourceID(arg0, arg1 any) *MockStateGetApplicationResourceIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationResourceID", reflect.TypeOf((*MockState)(nil).GetApplicationResourceID), arg0, arg1)
	return &MockStateGetApplicationResourceIDCall{Call: call}
}

// MockStateGetApplicationResourceIDCall wrap *gomock.Call
type MockStateGetApplicationResourceIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetApplicationResourceIDCall) Return(arg0 resource.UUID, arg1 error) *MockStateGetApplicationResourceIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetApplicationResourceIDCall) Do(f func(context.Context, resource0.GetApplicationResourceIDArgs) (resource.UUID, error)) *MockStateGetApplicationResourceIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetApplicationResourceIDCall) DoAndReturn(f func(context.Context, resource0.GetApplicationResourceIDArgs) (resource.UUID, error)) *MockStateGetApplicationResourceIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetResource mocks base method.
func (m *MockState) GetResource(arg0 context.Context, arg1 resource.UUID) (resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", arg0, arg1)
	ret0, _ := ret[0].(resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource.
func (mr *MockStateMockRecorder) GetResource(arg0, arg1 any) *MockStateGetResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockState)(nil).GetResource), arg0, arg1)
	return &MockStateGetResourceCall{Call: call}
}

// MockStateGetResourceCall wrap *gomock.Call
type MockStateGetResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetResourceCall) Return(arg0 resource.Resource, arg1 error) *MockStateGetResourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetResourceCall) Do(f func(context.Context, resource.UUID) (resource.Resource, error)) *MockStateGetResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetResourceCall) DoAndReturn(f func(context.Context, resource.UUID) (resource.Resource, error)) *MockStateGetResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetResourceType mocks base method.
func (m *MockState) GetResourceType(arg0 context.Context, arg1 resource.UUID) (resource1.Type, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceType", arg0, arg1)
	ret0, _ := ret[0].(resource1.Type)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceType indicates an expected call of GetResourceType.
func (mr *MockStateMockRecorder) GetResourceType(arg0, arg1 any) *MockStateGetResourceTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceType", reflect.TypeOf((*MockState)(nil).GetResourceType), arg0, arg1)
	return &MockStateGetResourceTypeCall{Call: call}
}

// MockStateGetResourceTypeCall wrap *gomock.Call
type MockStateGetResourceTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetResourceTypeCall) Return(arg0 resource1.Type, arg1 error) *MockStateGetResourceTypeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetResourceTypeCall) Do(f func(context.Context, resource.UUID) (resource1.Type, error)) *MockStateGetResourceTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetResourceTypeCall) DoAndReturn(f func(context.Context, resource.UUID) (resource1.Type, error)) *MockStateGetResourceTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetResourceUUIDByApplicationAndResourceName mocks base method.
func (m *MockState) GetResourceUUIDByApplicationAndResourceName(arg0 context.Context, arg1, arg2 string) (resource.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceUUIDByApplicationAndResourceName", arg0, arg1, arg2)
	ret0, _ := ret[0].(resource.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceUUIDByApplicationAndResourceName indicates an expected call of GetResourceUUIDByApplicationAndResourceName.
func (mr *MockStateMockRecorder) GetResourceUUIDByApplicationAndResourceName(arg0, arg1, arg2 any) *MockStateGetResourceUUIDByApplicationAndResourceNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceUUIDByApplicationAndResourceName", reflect.TypeOf((*MockState)(nil).GetResourceUUIDByApplicationAndResourceName), arg0, arg1, arg2)
	return &MockStateGetResourceUUIDByApplicationAndResourceNameCall{Call: call}
}

// MockStateGetResourceUUIDByApplicationAndResourceNameCall wrap *gomock.Call
type MockStateGetResourceUUIDByApplicationAndResourceNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetResourceUUIDByApplicationAndResourceNameCall) Return(arg0 resource.UUID, arg1 error) *MockStateGetResourceUUIDByApplicationAndResourceNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetResourceUUIDByApplicationAndResourceNameCall) Do(f func(context.Context, string, string) (resource.UUID, error)) *MockStateGetResourceUUIDByApplicationAndResourceNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetResourceUUIDByApplicationAndResourceNameCall) DoAndReturn(f func(context.Context, string, string) (resource.UUID, error)) *MockStateGetResourceUUIDByApplicationAndResourceNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetResourcesByApplicationID mocks base method.
func (m *MockState) GetResourcesByApplicationID(arg0 context.Context, arg1 application.ID) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourcesByApplicationID", arg0, arg1)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcesByApplicationID indicates an expected call of GetResourcesByApplicationID.
func (mr *MockStateMockRecorder) GetResourcesByApplicationID(arg0, arg1 any) *MockStateGetResourcesByApplicationIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcesByApplicationID", reflect.TypeOf((*MockState)(nil).GetResourcesByApplicationID), arg0, arg1)
	return &MockStateGetResourcesByApplicationIDCall{Call: call}
}

// MockStateGetResourcesByApplicationIDCall wrap *gomock.Call
type MockStateGetResourcesByApplicationIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetResourcesByApplicationIDCall) Return(arg0 []resource.Resource, arg1 error) *MockStateGetResourcesByApplicationIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetResourcesByApplicationIDCall) Do(f func(context.Context, application.ID) ([]resource.Resource, error)) *MockStateGetResourcesByApplicationIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetResourcesByApplicationIDCall) DoAndReturn(f func(context.Context, application.ID) ([]resource.Resource, error)) *MockStateGetResourcesByApplicationIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ImportResources mocks base method.
func (m *MockState) ImportResources(arg0 context.Context, arg1 resource0.ImportResourcesArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportResources indicates an expected call of ImportResources.
func (mr *MockStateMockRecorder) ImportResources(arg0, arg1 any) *MockStateImportResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportResources", reflect.TypeOf((*MockState)(nil).ImportResources), arg0, arg1)
	return &MockStateImportResourcesCall{Call: call}
}

// MockStateImportResourcesCall wrap *gomock.Call
type MockStateImportResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateImportResourcesCall) Return(arg0 error) *MockStateImportResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateImportResourcesCall) Do(f func(context.Context, resource0.ImportResourcesArgs) error) *MockStateImportResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateImportResourcesCall) DoAndReturn(f func(context.Context, resource0.ImportResourcesArgs) error) *MockStateImportResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListResources mocks base method.
func (m *MockState) ListResources(arg0 context.Context, arg1 application.ID) (resource.ApplicationResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0, arg1)
	ret0, _ := ret[0].(resource.ApplicationResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockStateMockRecorder) ListResources(arg0, arg1 any) *MockStateListResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockState)(nil).ListResources), arg0, arg1)
	return &MockStateListResourcesCall{Call: call}
}

// MockStateListResourcesCall wrap *gomock.Call
type MockStateListResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateListResourcesCall) Return(arg0 resource.ApplicationResources, arg1 error) *MockStateListResourcesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateListResourcesCall) Do(f func(context.Context, application.ID) (resource.ApplicationResources, error)) *MockStateListResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateListResourcesCall) DoAndReturn(f func(context.Context, application.ID) (resource.ApplicationResources, error)) *MockStateListResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RecordStoredResource mocks base method.
func (m *MockState) RecordStoredResource(arg0 context.Context, arg1 resource0.RecordStoredResourceArgs) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordStoredResource", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordStoredResource indicates an expected call of RecordStoredResource.
func (mr *MockStateMockRecorder) RecordStoredResource(arg0, arg1 any) *MockStateRecordStoredResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordStoredResource", reflect.TypeOf((*MockState)(nil).RecordStoredResource), arg0, arg1)
	return &MockStateRecordStoredResourceCall{Call: call}
}

// MockStateRecordStoredResourceCall wrap *gomock.Call
type MockStateRecordStoredResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateRecordStoredResourceCall) Return(arg0 string, arg1 error) *MockStateRecordStoredResourceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateRecordStoredResourceCall) Do(f func(context.Context, resource0.RecordStoredResourceArgs) (string, error)) *MockStateRecordStoredResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateRecordStoredResourceCall) DoAndReturn(f func(context.Context, resource0.RecordStoredResourceArgs) (string, error)) *MockStateRecordStoredResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetApplicationResource mocks base method.
func (m *MockState) SetApplicationResource(arg0 context.Context, arg1 resource.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetApplicationResource", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetApplicationResource indicates an expected call of SetApplicationResource.
func (mr *MockStateMockRecorder) SetApplicationResource(arg0, arg1 any) *MockStateSetApplicationResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApplicationResource", reflect.TypeOf((*MockState)(nil).SetApplicationResource), arg0, arg1)
	return &MockStateSetApplicationResourceCall{Call: call}
}

// MockStateSetApplicationResourceCall wrap *gomock.Call
type MockStateSetApplicationResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetApplicationResourceCall) Return(arg0 error) *MockStateSetApplicationResourceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetApplicationResourceCall) Do(f func(context.Context, resource.UUID) error) *MockStateSetApplicationResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetApplicationResourceCall) DoAndReturn(f func(context.Context, resource.UUID) error) *MockStateSetApplicationResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetRepositoryResources mocks base method.
func (m *MockState) SetRepositoryResources(arg0 context.Context, arg1 resource0.SetRepositoryResourcesArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRepositoryResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRepositoryResources indicates an expected call of SetRepositoryResources.
func (mr *MockStateMockRecorder) SetRepositoryResources(arg0, arg1 any) *MockStateSetRepositoryResourcesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRepositoryResources", reflect.TypeOf((*MockState)(nil).SetRepositoryResources), arg0, arg1)
	return &MockStateSetRepositoryResourcesCall{Call: call}
}

// MockStateSetRepositoryResourcesCall wrap *gomock.Call
type MockStateSetRepositoryResourcesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetRepositoryResourcesCall) Return(arg0 error) *MockStateSetRepositoryResourcesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetRepositoryResourcesCall) Do(f func(context.Context, resource0.SetRepositoryResourcesArgs) error) *MockStateSetRepositoryResourcesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetRepositoryResourcesCall) DoAndReturn(f func(context.Context, resource0.SetRepositoryResourcesArgs) error) *MockStateSetRepositoryResourcesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetUnitResource mocks base method.
func (m *MockState) SetUnitResource(arg0 context.Context, arg1 resource.UUID, arg2 unit.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnitResource", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUnitResource indicates an expected call of SetUnitResource.
func (mr *MockStateMockRecorder) SetUnitResource(arg0, arg1, arg2 any) *MockStateSetUnitResourceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnitResource", reflect.TypeOf((*MockState)(nil).SetUnitResource), arg0, arg1, arg2)
	return &MockStateSetUnitResourceCall{Call: call}
}

// MockStateSetUnitResourceCall wrap *gomock.Call
type MockStateSetUnitResourceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetUnitResourceCall) Return(arg0 error) *MockStateSetUnitResourceCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetUnitResourceCall) Do(f func(context.Context, resource.UUID, unit.UUID) error) *MockStateSetUnitResourceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetUnitResourceCall) DoAndReturn(f func(context.Context, resource.UUID, unit.UUID) error) *MockStateSetUnitResourceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateResourceRevisionAndDeletePriorVersion mocks base method.
func (m *MockState) UpdateResourceRevisionAndDeletePriorVersion(arg0 context.Context, arg1 resource0.UpdateResourceRevisionArgs, arg2 resource1.Type) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResourceRevisionAndDeletePriorVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResourceRevisionAndDeletePriorVersion indicates an expected call of UpdateResourceRevisionAndDeletePriorVersion.
func (mr *MockStateMockRecorder) UpdateResourceRevisionAndDeletePriorVersion(arg0, arg1, arg2 any) *MockStateUpdateResourceRevisionAndDeletePriorVersionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceRevisionAndDeletePriorVersion", reflect.TypeOf((*MockState)(nil).UpdateResourceRevisionAndDeletePriorVersion), arg0, arg1, arg2)
	return &MockStateUpdateResourceRevisionAndDeletePriorVersionCall{Call: call}
}

// MockStateUpdateResourceRevisionAndDeletePriorVersionCall wrap *gomock.Call
type MockStateUpdateResourceRevisionAndDeletePriorVersionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateResourceRevisionAndDeletePriorVersionCall) Return(arg0 string, arg1 error) *MockStateUpdateResourceRevisionAndDeletePriorVersionCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateResourceRevisionAndDeletePriorVersionCall) Do(f func(context.Context, resource0.UpdateResourceRevisionArgs, resource1.Type) (string, error)) *MockStateUpdateResourceRevisionAndDeletePriorVersionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateResourceRevisionAndDeletePriorVersionCall) DoAndReturn(f func(context.Context, resource0.UpdateResourceRevisionArgs, resource1.Type) (string, error)) *MockStateUpdateResourceRevisionAndDeletePriorVersionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockResourceStoreGetter is a mock of ResourceStoreGetter interface.
type MockResourceStoreGetter struct {
	ctrl     *gomock.Controller
	recorder *MockResourceStoreGetterMockRecorder
}

// MockResourceStoreGetterMockRecorder is the mock recorder for MockResourceStoreGetter.
type MockResourceStoreGetterMockRecorder struct {
	mock *MockResourceStoreGetter
}

// NewMockResourceStoreGetter creates a new mock instance.
func NewMockResourceStoreGetter(ctrl *gomock.Controller) *MockResourceStoreGetter {
	mock := &MockResourceStoreGetter{ctrl: ctrl}
	mock.recorder = &MockResourceStoreGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceStoreGetter) EXPECT() *MockResourceStoreGetterMockRecorder {
	return m.recorder
}

// GetResourceStore mocks base method.
func (m *MockResourceStoreGetter) GetResourceStore(arg0 context.Context, arg1 resource1.Type) (store.ResourceStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceStore", arg0, arg1)
	ret0, _ := ret[0].(store.ResourceStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceStore indicates an expected call of GetResourceStore.
func (mr *MockResourceStoreGetterMockRecorder) GetResourceStore(arg0, arg1 any) *MockResourceStoreGetterGetResourceStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceStore", reflect.TypeOf((*MockResourceStoreGetter)(nil).GetResourceStore), arg0, arg1)
	return &MockResourceStoreGetterGetResourceStoreCall{Call: call}
}

// MockResourceStoreGetterGetResourceStoreCall wrap *gomock.Call
type MockResourceStoreGetterGetResourceStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockResourceStoreGetterGetResourceStoreCall) Return(arg0 store.ResourceStore, arg1 error) *MockResourceStoreGetterGetResourceStoreCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockResourceStoreGetterGetResourceStoreCall) Do(f func(context.Context, resource1.Type) (store.ResourceStore, error)) *MockResourceStoreGetterGetResourceStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockResourceStoreGetterGetResourceStoreCall) DoAndReturn(f func(context.Context, resource1.Type) (store.ResourceStore, error)) *MockResourceStoreGetterGetResourceStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
