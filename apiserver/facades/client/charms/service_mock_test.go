// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/charms (interfaces: ModelConfigService,ApplicationService,MachineService)
//
// Generated by this command:
//
//	mockgen -typed -package charms_test -destination service_mock_test.go github.com/juju/juju/apiserver/facades/client/charms ModelConfigService,ApplicationService,MachineService
//

// Package charms_test is a generated GoMock package.
package charms_test

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/juju/core/charm"
	instance "github.com/juju/juju/core/instance"
	machine "github.com/juju/juju/core/machine"
	charm0 "github.com/juju/juju/domain/application/charm"
	config "github.com/juju/juju/environs/config"
	gomock "go.uber.org/mock/gomock"
)

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *MockModelConfigServiceModelConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
	return &MockModelConfigServiceModelConfigCall{Call: call}
}

// MockModelConfigServiceModelConfigCall wrap *gomock.Call
type MockModelConfigServiceModelConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelConfigServiceModelConfigCall) Return(arg0 *config.Config, arg1 error) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelConfigServiceModelConfigCall) Do(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelConfigServiceModelConfigCall) DoAndReturn(f func(context.Context) (*config.Config, error)) *MockModelConfigServiceModelConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// ListCharmsWithOriginByNames mocks base method.
func (m *MockApplicationService) ListCharmsWithOriginByNames(arg0 context.Context, arg1 ...string) ([]charm0.CharmWithOrigin, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCharmsWithOriginByNames", varargs...)
	ret0, _ := ret[0].([]charm0.CharmWithOrigin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCharmsWithOriginByNames indicates an expected call of ListCharmsWithOriginByNames.
func (mr *MockApplicationServiceMockRecorder) ListCharmsWithOriginByNames(arg0 any, arg1 ...any) *MockApplicationServiceListCharmsWithOriginByNamesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmsWithOriginByNames", reflect.TypeOf((*MockApplicationService)(nil).ListCharmsWithOriginByNames), varargs...)
	return &MockApplicationServiceListCharmsWithOriginByNamesCall{Call: call}
}

// MockApplicationServiceListCharmsWithOriginByNamesCall wrap *gomock.Call
type MockApplicationServiceListCharmsWithOriginByNamesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceListCharmsWithOriginByNamesCall) Return(arg0 []charm0.CharmWithOrigin, arg1 error) *MockApplicationServiceListCharmsWithOriginByNamesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceListCharmsWithOriginByNamesCall) Do(f func(context.Context, ...string) ([]charm0.CharmWithOrigin, error)) *MockApplicationServiceListCharmsWithOriginByNamesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceListCharmsWithOriginByNamesCall) DoAndReturn(f func(context.Context, ...string) ([]charm0.CharmWithOrigin, error)) *MockApplicationServiceListCharmsWithOriginByNamesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetCharm mocks base method.
func (m *MockApplicationService) SetCharm(arg0 context.Context, arg1 charm0.SetCharmArgs) (charm.ID, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", arg0, arg1)
	ret0, _ := ret[0].(charm.ID)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetCharm indicates an expected call of SetCharm.
func (mr *MockApplicationServiceMockRecorder) SetCharm(arg0, arg1 any) *MockApplicationServiceSetCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockApplicationService)(nil).SetCharm), arg0, arg1)
	return &MockApplicationServiceSetCharmCall{Call: call}
}

// MockApplicationServiceSetCharmCall wrap *gomock.Call
type MockApplicationServiceSetCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationServiceSetCharmCall) Return(arg0 charm.ID, arg1 []string, arg2 error) *MockApplicationServiceSetCharmCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationServiceSetCharmCall) Do(f func(context.Context, charm0.SetCharmArgs) (charm.ID, []string, error)) *MockApplicationServiceSetCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationServiceSetCharmCall) DoAndReturn(f func(context.Context, charm0.SetCharmArgs) (charm.ID, []string, error)) *MockApplicationServiceSetCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// GetMachineUUID mocks base method.
func (m *MockMachineService) GetMachineUUID(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineUUID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineUUID indicates an expected call of GetMachineUUID.
func (mr *MockMachineServiceMockRecorder) GetMachineUUID(arg0, arg1 any) *MockMachineServiceGetMachineUUIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineUUID", reflect.TypeOf((*MockMachineService)(nil).GetMachineUUID), arg0, arg1)
	return &MockMachineServiceGetMachineUUIDCall{Call: call}
}

// MockMachineServiceGetMachineUUIDCall wrap *gomock.Call
type MockMachineServiceGetMachineUUIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceGetMachineUUIDCall) Return(arg0 string, arg1 error) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceGetMachineUUIDCall) Do(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceGetMachineUUIDCall) DoAndReturn(f func(context.Context, machine.Name) (string, error)) *MockMachineServiceGetMachineUUIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HardwareCharacteristics mocks base method.
func (m *MockMachineService) HardwareCharacteristics(arg0 context.Context, arg1 string) (*instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareCharacteristics", arg0, arg1)
	ret0, _ := ret[0].(*instance.HardwareCharacteristics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareCharacteristics indicates an expected call of HardwareCharacteristics.
func (mr *MockMachineServiceMockRecorder) HardwareCharacteristics(arg0, arg1 any) *MockMachineServiceHardwareCharacteristicsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareCharacteristics", reflect.TypeOf((*MockMachineService)(nil).HardwareCharacteristics), arg0, arg1)
	return &MockMachineServiceHardwareCharacteristicsCall{Call: call}
}

// MockMachineServiceHardwareCharacteristicsCall wrap *gomock.Call
type MockMachineServiceHardwareCharacteristicsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineServiceHardwareCharacteristicsCall) Return(arg0 *instance.HardwareCharacteristics, arg1 error) *MockMachineServiceHardwareCharacteristicsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineServiceHardwareCharacteristicsCall) Do(f func(context.Context, string) (*instance.HardwareCharacteristics, error)) *MockMachineServiceHardwareCharacteristicsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineServiceHardwareCharacteristicsCall) DoAndReturn(f func(context.Context, string) (*instance.HardwareCharacteristics, error)) *MockMachineServiceHardwareCharacteristicsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
