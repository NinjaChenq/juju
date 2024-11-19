// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/charms/interfaces (interfaces: BackendState,Application,Machine,Unit,Downloader)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/state_mock.go github.com/juju/juju/apiserver/facades/client/charms/interfaces BackendState,Application,Machine,Unit,Downloader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	interfaces "github.com/juju/juju/apiserver/facades/client/charms/interfaces"
	charm "github.com/juju/juju/core/charm"
	constraints "github.com/juju/juju/core/constraints"
	charm0 "github.com/juju/juju/internal/charm"
	services "github.com/juju/juju/internal/charm/services"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockBackendState is a mock of BackendState interface.
type MockBackendState struct {
	ctrl     *gomock.Controller
	recorder *MockBackendStateMockRecorder
}

// MockBackendStateMockRecorder is the mock recorder for MockBackendState.
type MockBackendStateMockRecorder struct {
	mock *MockBackendState
}

// NewMockBackendState creates a new mock instance.
func NewMockBackendState(ctrl *gomock.Controller) *MockBackendState {
	mock := &MockBackendState{ctrl: ctrl}
	mock.recorder = &MockBackendStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackendState) EXPECT() *MockBackendStateMockRecorder {
	return m.recorder
}

// AddCharmMetadata mocks base method.
func (m *MockBackendState) AddCharmMetadata(arg0 state.CharmInfo) (state.CharmRefFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharmMetadata", arg0)
	ret0, _ := ret[0].(state.CharmRefFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharmMetadata indicates an expected call of AddCharmMetadata.
func (mr *MockBackendStateMockRecorder) AddCharmMetadata(arg0 any) *MockBackendStateAddCharmMetadataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharmMetadata", reflect.TypeOf((*MockBackendState)(nil).AddCharmMetadata), arg0)
	return &MockBackendStateAddCharmMetadataCall{Call: call}
}

// MockBackendStateAddCharmMetadataCall wrap *gomock.Call
type MockBackendStateAddCharmMetadataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateAddCharmMetadataCall) Return(arg0 state.CharmRefFull, arg1 error) *MockBackendStateAddCharmMetadataCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateAddCharmMetadataCall) Do(f func(state.CharmInfo) (state.CharmRefFull, error)) *MockBackendStateAddCharmMetadataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateAddCharmMetadataCall) DoAndReturn(f func(state.CharmInfo) (state.CharmRefFull, error)) *MockBackendStateAddCharmMetadataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Application mocks base method.
func (m *MockBackendState) Application(arg0 string) (interfaces.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(interfaces.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockBackendStateMockRecorder) Application(arg0 any) *MockBackendStateApplicationCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockBackendState)(nil).Application), arg0)
	return &MockBackendStateApplicationCall{Call: call}
}

// MockBackendStateApplicationCall wrap *gomock.Call
type MockBackendStateApplicationCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateApplicationCall) Return(arg0 interfaces.Application, arg1 error) *MockBackendStateApplicationCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateApplicationCall) Do(f func(string) (interfaces.Application, error)) *MockBackendStateApplicationCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateApplicationCall) DoAndReturn(f func(string) (interfaces.Application, error)) *MockBackendStateApplicationCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Charm mocks base method.
func (m *MockBackendState) Charm(arg0 string) (state.CharmRefFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(state.CharmRefFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm.
func (mr *MockBackendStateMockRecorder) Charm(arg0 any) *MockBackendStateCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockBackendState)(nil).Charm), arg0)
	return &MockBackendStateCharmCall{Call: call}
}

// MockBackendStateCharmCall wrap *gomock.Call
type MockBackendStateCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateCharmCall) Return(arg0 state.CharmRefFull, arg1 error) *MockBackendStateCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateCharmCall) Do(f func(string) (state.CharmRefFull, error)) *MockBackendStateCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateCharmCall) DoAndReturn(f func(string) (state.CharmRefFull, error)) *MockBackendStateCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerTag mocks base method.
func (m *MockBackendState) ControllerTag() names.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerTag")
	ret0, _ := ret[0].(names.ControllerTag)
	return ret0
}

// ControllerTag indicates an expected call of ControllerTag.
func (mr *MockBackendStateMockRecorder) ControllerTag() *MockBackendStateControllerTagCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerTag", reflect.TypeOf((*MockBackendState)(nil).ControllerTag))
	return &MockBackendStateControllerTagCall{Call: call}
}

// MockBackendStateControllerTagCall wrap *gomock.Call
type MockBackendStateControllerTagCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateControllerTagCall) Return(arg0 names.ControllerTag) *MockBackendStateControllerTagCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateControllerTagCall) Do(f func() names.ControllerTag) *MockBackendStateControllerTagCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateControllerTagCall) DoAndReturn(f func() names.ControllerTag) *MockBackendStateControllerTagCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Machine mocks base method.
func (m *MockBackendState) Machine(arg0 string) (interfaces.Machine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(interfaces.Machine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockBackendStateMockRecorder) Machine(arg0 any) *MockBackendStateMachineCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockBackendState)(nil).Machine), arg0)
	return &MockBackendStateMachineCall{Call: call}
}

// MockBackendStateMachineCall wrap *gomock.Call
type MockBackendStateMachineCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateMachineCall) Return(arg0 interfaces.Machine, arg1 error) *MockBackendStateMachineCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateMachineCall) Do(f func(string) (interfaces.Machine, error)) *MockBackendStateMachineCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateMachineCall) DoAndReturn(f func(string) (interfaces.Machine, error)) *MockBackendStateMachineCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelConstraints mocks base method.
func (m *MockBackendState) ModelConstraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConstraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConstraints indicates an expected call of ModelConstraints.
func (mr *MockBackendStateMockRecorder) ModelConstraints() *MockBackendStateModelConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConstraints", reflect.TypeOf((*MockBackendState)(nil).ModelConstraints))
	return &MockBackendStateModelConstraintsCall{Call: call}
}

// MockBackendStateModelConstraintsCall wrap *gomock.Call
type MockBackendStateModelConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateModelConstraintsCall) Return(arg0 constraints.Value, arg1 error) *MockBackendStateModelConstraintsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateModelConstraintsCall) Do(f func() (constraints.Value, error)) *MockBackendStateModelConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateModelConstraintsCall) DoAndReturn(f func() (constraints.Value, error)) *MockBackendStateModelConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PrepareCharmUpload mocks base method.
func (m *MockBackendState) PrepareCharmUpload(arg0 string) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockBackendStateMockRecorder) PrepareCharmUpload(arg0 any) *MockBackendStatePrepareCharmUploadCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockBackendState)(nil).PrepareCharmUpload), arg0)
	return &MockBackendStatePrepareCharmUploadCall{Call: call}
}

// MockBackendStatePrepareCharmUploadCall wrap *gomock.Call
type MockBackendStatePrepareCharmUploadCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStatePrepareCharmUploadCall) Return(arg0 services.UploadedCharm, arg1 error) *MockBackendStatePrepareCharmUploadCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStatePrepareCharmUploadCall) Do(f func(string) (services.UploadedCharm, error)) *MockBackendStatePrepareCharmUploadCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStatePrepareCharmUploadCall) DoAndReturn(f func(string) (services.UploadedCharm, error)) *MockBackendStatePrepareCharmUploadCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUploadedCharm mocks base method.
func (m *MockBackendState) UpdateUploadedCharm(arg0 state.CharmInfo) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockBackendStateMockRecorder) UpdateUploadedCharm(arg0 any) *MockBackendStateUpdateUploadedCharmCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockBackendState)(nil).UpdateUploadedCharm), arg0)
	return &MockBackendStateUpdateUploadedCharmCall{Call: call}
}

// MockBackendStateUpdateUploadedCharmCall wrap *gomock.Call
type MockBackendStateUpdateUploadedCharmCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockBackendStateUpdateUploadedCharmCall) Return(arg0 services.UploadedCharm, arg1 error) *MockBackendStateUpdateUploadedCharmCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockBackendStateUpdateUploadedCharmCall) Do(f func(state.CharmInfo) (services.UploadedCharm, error)) *MockBackendStateUpdateUploadedCharmCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockBackendStateUpdateUploadedCharmCall) DoAndReturn(f func(state.CharmInfo) (services.UploadedCharm, error)) *MockBackendStateUpdateUploadedCharmCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// AllUnits mocks base method.
func (m *MockApplication) AllUnits() ([]interfaces.Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllUnits")
	ret0, _ := ret[0].([]interfaces.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllUnits indicates an expected call of AllUnits.
func (mr *MockApplicationMockRecorder) AllUnits() *MockApplicationAllUnitsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllUnits", reflect.TypeOf((*MockApplication)(nil).AllUnits))
	return &MockApplicationAllUnitsCall{Call: call}
}

// MockApplicationAllUnitsCall wrap *gomock.Call
type MockApplicationAllUnitsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationAllUnitsCall) Return(arg0 []interfaces.Unit, arg1 error) *MockApplicationAllUnitsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationAllUnitsCall) Do(f func() ([]interfaces.Unit, error)) *MockApplicationAllUnitsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationAllUnitsCall) DoAndReturn(f func() ([]interfaces.Unit, error)) *MockApplicationAllUnitsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Constraints mocks base method.
func (m *MockApplication) Constraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints.
func (mr *MockApplicationMockRecorder) Constraints() *MockApplicationConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockApplication)(nil).Constraints))
	return &MockApplicationConstraintsCall{Call: call}
}

// MockApplicationConstraintsCall wrap *gomock.Call
type MockApplicationConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationConstraintsCall) Return(arg0 constraints.Value, arg1 error) *MockApplicationConstraintsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationConstraintsCall) Do(f func() (constraints.Value, error)) *MockApplicationConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationConstraintsCall) DoAndReturn(f func() (constraints.Value, error)) *MockApplicationConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsPrincipal mocks base method.
func (m *MockApplication) IsPrincipal() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrincipal")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrincipal indicates an expected call of IsPrincipal.
func (mr *MockApplicationMockRecorder) IsPrincipal() *MockApplicationIsPrincipalCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrincipal", reflect.TypeOf((*MockApplication)(nil).IsPrincipal))
	return &MockApplicationIsPrincipalCall{Call: call}
}

// MockApplicationIsPrincipalCall wrap *gomock.Call
type MockApplicationIsPrincipalCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockApplicationIsPrincipalCall) Return(arg0 bool) *MockApplicationIsPrincipalCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockApplicationIsPrincipalCall) Do(f func() bool) *MockApplicationIsPrincipalCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockApplicationIsPrincipalCall) DoAndReturn(f func() bool) *MockApplicationIsPrincipalCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMachine is a mock of Machine interface.
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine.
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance.
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// Constraints mocks base method.
func (m *MockMachine) Constraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Constraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Constraints indicates an expected call of Constraints.
func (mr *MockMachineMockRecorder) Constraints() *MockMachineConstraintsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Constraints", reflect.TypeOf((*MockMachine)(nil).Constraints))
	return &MockMachineConstraintsCall{Call: call}
}

// MockMachineConstraintsCall wrap *gomock.Call
type MockMachineConstraintsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMachineConstraintsCall) Return(arg0 constraints.Value, arg1 error) *MockMachineConstraintsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMachineConstraintsCall) Do(f func() (constraints.Value, error)) *MockMachineConstraintsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMachineConstraintsCall) DoAndReturn(f func() (constraints.Value, error)) *MockMachineConstraintsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockUnit is a mock of Unit interface.
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit.
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance.
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// AssignedMachineId mocks base method.
func (m *MockUnit) AssignedMachineId() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignedMachineId")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignedMachineId indicates an expected call of AssignedMachineId.
func (mr *MockUnitMockRecorder) AssignedMachineId() *MockUnitAssignedMachineIdCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignedMachineId", reflect.TypeOf((*MockUnit)(nil).AssignedMachineId))
	return &MockUnitAssignedMachineIdCall{Call: call}
}

// MockUnitAssignedMachineIdCall wrap *gomock.Call
type MockUnitAssignedMachineIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUnitAssignedMachineIdCall) Return(arg0 string, arg1 error) *MockUnitAssignedMachineIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUnitAssignedMachineIdCall) Do(f func() (string, error)) *MockUnitAssignedMachineIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUnitAssignedMachineIdCall) DoAndReturn(f func() (string, error)) *MockUnitAssignedMachineIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockDownloader is a mock of Downloader interface.
type MockDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockDownloaderMockRecorder
}

// MockDownloaderMockRecorder is the mock recorder for MockDownloader.
type MockDownloaderMockRecorder struct {
	mock *MockDownloader
}

// NewMockDownloader creates a new mock instance.
func NewMockDownloader(ctrl *gomock.Controller) *MockDownloader {
	mock := &MockDownloader{ctrl: ctrl}
	mock.recorder = &MockDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloader) EXPECT() *MockDownloaderMockRecorder {
	return m.recorder
}

// DownloadAndStore mocks base method.
func (m *MockDownloader) DownloadAndStore(arg0 context.Context, arg1 *charm0.URL, arg2 charm.Origin, arg3 bool) (charm.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadAndStore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadAndStore indicates an expected call of DownloadAndStore.
func (mr *MockDownloaderMockRecorder) DownloadAndStore(arg0, arg1, arg2, arg3 any) *MockDownloaderDownloadAndStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAndStore", reflect.TypeOf((*MockDownloader)(nil).DownloadAndStore), arg0, arg1, arg2, arg3)
	return &MockDownloaderDownloadAndStoreCall{Call: call}
}

// MockDownloaderDownloadAndStoreCall wrap *gomock.Call
type MockDownloaderDownloadAndStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockDownloaderDownloadAndStoreCall) Return(arg0 charm.Origin, arg1 error) *MockDownloaderDownloadAndStoreCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockDownloaderDownloadAndStoreCall) Do(f func(context.Context, *charm0.URL, charm.Origin, bool) (charm.Origin, error)) *MockDownloaderDownloadAndStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockDownloaderDownloadAndStoreCall) DoAndReturn(f func(context.Context, *charm0.URL, charm.Origin, bool) (charm.Origin, error)) *MockDownloaderDownloadAndStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
