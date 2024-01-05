// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/migrationmaster (interfaces: Backend,ControllerState,ModelExporter,UpgradeService)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/backend.go github.com/juju/juju/apiserver/facades/controller/migrationmaster Backend,ControllerState,ModelExporter,UpgradeService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	description "github.com/juju/description/v5"
	controller "github.com/juju/juju/controller"
	network "github.com/juju/juju/core/network"
	objectstore "github.com/juju/juju/core/objectstore"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v5"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// AgentVersion mocks base method.
func (m *MockBackend) AgentVersion(arg0 context.Context) (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentVersion", arg0)
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentVersion indicates an expected call of AgentVersion.
func (mr *MockBackendMockRecorder) AgentVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentVersion", reflect.TypeOf((*MockBackend)(nil).AgentVersion), arg0)
}

// AllLocalRelatedModels mocks base method.
func (m *MockBackend) AllLocalRelatedModels() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllLocalRelatedModels")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllLocalRelatedModels indicates an expected call of AllLocalRelatedModels.
func (mr *MockBackendMockRecorder) AllLocalRelatedModels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllLocalRelatedModels", reflect.TypeOf((*MockBackend)(nil).AllLocalRelatedModels))
}

// ControllerConfig mocks base method.
func (m *MockBackend) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockBackendMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockBackend)(nil).ControllerConfig))
}

// Export mocks base method.
func (m *MockBackend) Export(arg0 map[string]string, arg1 objectstore.ObjectStore) (description.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", arg0, arg1)
	ret0, _ := ret[0].(description.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Export indicates an expected call of Export.
func (mr *MockBackendMockRecorder) Export(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockBackend)(nil).Export), arg0, arg1)
}

// ExportPartial mocks base method.
func (m *MockBackend) ExportPartial(arg0 state.ExportConfig, arg1 objectstore.ObjectStore) (description.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportPartial", arg0, arg1)
	ret0, _ := ret[0].(description.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportPartial indicates an expected call of ExportPartial.
func (mr *MockBackendMockRecorder) ExportPartial(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportPartial", reflect.TypeOf((*MockBackend)(nil).ExportPartial), arg0, arg1)
}

// LatestMigration mocks base method.
func (m *MockBackend) LatestMigration() (state.ModelMigration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestMigration")
	ret0, _ := ret[0].(state.ModelMigration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestMigration indicates an expected call of LatestMigration.
func (mr *MockBackendMockRecorder) LatestMigration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestMigration", reflect.TypeOf((*MockBackend)(nil).LatestMigration))
}

// ModelName mocks base method.
func (m *MockBackend) ModelName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelName indicates an expected call of ModelName.
func (mr *MockBackendMockRecorder) ModelName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelName", reflect.TypeOf((*MockBackend)(nil).ModelName))
}

// ModelOwner mocks base method.
func (m *MockBackend) ModelOwner() (names.UserTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelOwner")
	ret0, _ := ret[0].(names.UserTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelOwner indicates an expected call of ModelOwner.
func (mr *MockBackendMockRecorder) ModelOwner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelOwner", reflect.TypeOf((*MockBackend)(nil).ModelOwner))
}

// ModelUUID mocks base method.
func (m *MockBackend) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockBackendMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockBackend)(nil).ModelUUID))
}

// RemoveExportingModelDocs mocks base method.
func (m *MockBackend) RemoveExportingModelDocs() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveExportingModelDocs")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveExportingModelDocs indicates an expected call of RemoveExportingModelDocs.
func (mr *MockBackendMockRecorder) RemoveExportingModelDocs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveExportingModelDocs", reflect.TypeOf((*MockBackend)(nil).RemoveExportingModelDocs))
}

// WatchForMigration mocks base method.
func (m *MockBackend) WatchForMigration() state.NotifyWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchForMigration")
	ret0, _ := ret[0].(state.NotifyWatcher)
	return ret0
}

// WatchForMigration indicates an expected call of WatchForMigration.
func (mr *MockBackendMockRecorder) WatchForMigration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchForMigration", reflect.TypeOf((*MockBackend)(nil).WatchForMigration))
}

// MockControllerState is a mock of ControllerState interface.
type MockControllerState struct {
	ctrl     *gomock.Controller
	recorder *MockControllerStateMockRecorder
}

// MockControllerStateMockRecorder is the mock recorder for MockControllerState.
type MockControllerStateMockRecorder struct {
	mock *MockControllerState
}

// NewMockControllerState creates a new mock instance.
func NewMockControllerState(ctrl *gomock.Controller) *MockControllerState {
	mock := &MockControllerState{ctrl: ctrl}
	mock.recorder = &MockControllerStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerState) EXPECT() *MockControllerStateMockRecorder {
	return m.recorder
}

// APIHostPortsForClients mocks base method.
func (m *MockControllerState) APIHostPortsForClients(arg0 controller.Config) ([]network.SpaceHostPorts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIHostPortsForClients", arg0)
	ret0, _ := ret[0].([]network.SpaceHostPorts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIHostPortsForClients indicates an expected call of APIHostPortsForClients.
func (mr *MockControllerStateMockRecorder) APIHostPortsForClients(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIHostPortsForClients", reflect.TypeOf((*MockControllerState)(nil).APIHostPortsForClients), arg0)
}

// MockModelExporter is a mock of ModelExporter interface.
type MockModelExporter struct {
	ctrl     *gomock.Controller
	recorder *MockModelExporterMockRecorder
}

// MockModelExporterMockRecorder is the mock recorder for MockModelExporter.
type MockModelExporterMockRecorder struct {
	mock *MockModelExporter
}

// NewMockModelExporter creates a new mock instance.
func NewMockModelExporter(ctrl *gomock.Controller) *MockModelExporter {
	mock := &MockModelExporter{ctrl: ctrl}
	mock.recorder = &MockModelExporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelExporter) EXPECT() *MockModelExporterMockRecorder {
	return m.recorder
}

// ExportModel mocks base method.
func (m *MockModelExporter) ExportModel(arg0 context.Context, arg1 map[string]string, arg2 objectstore.ObjectStore) (description.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportModel", arg0, arg1, arg2)
	ret0, _ := ret[0].(description.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportModel indicates an expected call of ExportModel.
func (mr *MockModelExporterMockRecorder) ExportModel(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportModel", reflect.TypeOf((*MockModelExporter)(nil).ExportModel), arg0, arg1, arg2)
}

// MockUpgradeService is a mock of UpgradeService interface.
type MockUpgradeService struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeServiceMockRecorder
}

// MockUpgradeServiceMockRecorder is the mock recorder for MockUpgradeService.
type MockUpgradeServiceMockRecorder struct {
	mock *MockUpgradeService
}

// NewMockUpgradeService creates a new mock instance.
func NewMockUpgradeService(ctrl *gomock.Controller) *MockUpgradeService {
	mock := &MockUpgradeService{ctrl: ctrl}
	mock.recorder = &MockUpgradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeService) EXPECT() *MockUpgradeServiceMockRecorder {
	return m.recorder
}

// IsUpgrading mocks base method.
func (m *MockUpgradeService) IsUpgrading(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpgrading", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpgrading indicates an expected call of IsUpgrading.
func (mr *MockUpgradeServiceMockRecorder) IsUpgrading(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpgrading", reflect.TypeOf((*MockUpgradeService)(nil).IsUpgrading), arg0)
}
