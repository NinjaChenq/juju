// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facade (interfaces: Context,Authorizer)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/facade_mock.go github.com/juju/juju/apiserver/facade Context,Authorizer
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	facade "github.com/juju/juju/apiserver/facade"
	changestream "github.com/juju/juju/core/changestream"
	leadership "github.com/juju/juju/core/leadership"
	lease "github.com/juju/juju/core/lease"
	multiwatcher "github.com/juju/juju/core/multiwatcher"
	objectstore "github.com/juju/juju/core/objectstore"
	permission "github.com/juju/juju/core/permission"
	servicefactory "github.com/juju/juju/internal/servicefactory"
	state "github.com/juju/juju/state"
	loggo "github.com/juju/loggo"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockContext is a mock of Context interface.
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext.
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance.
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MockContext) Auth() facade.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth")
	ret0, _ := ret[0].(facade.Authorizer)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MockContextMockRecorder) Auth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MockContext)(nil).Auth))
}

// ControllerDB mocks base method.
func (m *MockContext) ControllerDB() (changestream.WatchableDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerDB")
	ret0, _ := ret[0].(changestream.WatchableDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerDB indicates an expected call of ControllerDB.
func (mr *MockContextMockRecorder) ControllerDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerDB", reflect.TypeOf((*MockContext)(nil).ControllerDB))
}

// ControllerObjectStore mocks base method.
func (m *MockContext) ControllerObjectStore() objectstore.ObjectStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerObjectStore")
	ret0, _ := ret[0].(objectstore.ObjectStore)
	return ret0
}

// ControllerObjectStore indicates an expected call of ControllerObjectStore.
func (mr *MockContextMockRecorder) ControllerObjectStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerObjectStore", reflect.TypeOf((*MockContext)(nil).ControllerObjectStore))
}

// DataDir mocks base method.
func (m *MockContext) DataDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DataDir indicates an expected call of DataDir.
func (mr *MockContextMockRecorder) DataDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataDir", reflect.TypeOf((*MockContext)(nil).DataDir))
}

// Dispose mocks base method.
func (m *MockContext) Dispose() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Dispose")
}

// Dispose indicates an expected call of Dispose.
func (mr *MockContextMockRecorder) Dispose() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispose", reflect.TypeOf((*MockContext)(nil).Dispose))
}

// HTTPClient mocks base method.
func (m *MockContext) HTTPClient(arg0 facade.HTTPClientPurpose) facade.HTTPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient", arg0)
	ret0, _ := ret[0].(facade.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockContextMockRecorder) HTTPClient(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockContext)(nil).HTTPClient), arg0)
}

// Hub mocks base method.
func (m *MockContext) Hub() facade.Hub {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hub")
	ret0, _ := ret[0].(facade.Hub)
	return ret0
}

// Hub indicates an expected call of Hub.
func (mr *MockContextMockRecorder) Hub() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hub", reflect.TypeOf((*MockContext)(nil).Hub))
}

// ID mocks base method.
func (m *MockContext) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockContextMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockContext)(nil).ID))
}

// LeadershipChecker mocks base method.
func (m *MockContext) LeadershipChecker() (leadership.Checker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipChecker")
	ret0, _ := ret[0].(leadership.Checker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeadershipChecker indicates an expected call of LeadershipChecker.
func (mr *MockContextMockRecorder) LeadershipChecker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipChecker", reflect.TypeOf((*MockContext)(nil).LeadershipChecker))
}

// LeadershipClaimer mocks base method.
func (m *MockContext) LeadershipClaimer(arg0 string) (leadership.Claimer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipClaimer", arg0)
	ret0, _ := ret[0].(leadership.Claimer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeadershipClaimer indicates an expected call of LeadershipClaimer.
func (mr *MockContextMockRecorder) LeadershipClaimer(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipClaimer", reflect.TypeOf((*MockContext)(nil).LeadershipClaimer), arg0)
}

// LeadershipPinner mocks base method.
func (m *MockContext) LeadershipPinner(arg0 string) (leadership.Pinner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipPinner", arg0)
	ret0, _ := ret[0].(leadership.Pinner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeadershipPinner indicates an expected call of LeadershipPinner.
func (mr *MockContextMockRecorder) LeadershipPinner(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipPinner", reflect.TypeOf((*MockContext)(nil).LeadershipPinner), arg0)
}

// LeadershipReader mocks base method.
func (m *MockContext) LeadershipReader(arg0 string) (leadership.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipReader", arg0)
	ret0, _ := ret[0].(leadership.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeadershipReader indicates an expected call of LeadershipReader.
func (mr *MockContextMockRecorder) LeadershipReader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipReader", reflect.TypeOf((*MockContext)(nil).LeadershipReader), arg0)
}

// LeadershipRevoker mocks base method.
func (m *MockContext) LeadershipRevoker(arg0 string) (leadership.Revoker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeadershipRevoker", arg0)
	ret0, _ := ret[0].(leadership.Revoker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeadershipRevoker indicates an expected call of LeadershipRevoker.
func (mr *MockContextMockRecorder) LeadershipRevoker(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeadershipRevoker", reflect.TypeOf((*MockContext)(nil).LeadershipRevoker), arg0)
}

// LogDir mocks base method.
func (m *MockContext) LogDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// LogDir indicates an expected call of LogDir.
func (mr *MockContextMockRecorder) LogDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogDir", reflect.TypeOf((*MockContext)(nil).LogDir))
}

// Logger mocks base method.
func (m *MockContext) Logger() loggo.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(loggo.Logger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockContextMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockContext)(nil).Logger))
}

// MachineTag mocks base method.
func (m *MockContext) MachineTag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MachineTag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// MachineTag indicates an expected call of MachineTag.
func (mr *MockContextMockRecorder) MachineTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MachineTag", reflect.TypeOf((*MockContext)(nil).MachineTag))
}

// MultiwatcherFactory mocks base method.
func (m *MockContext) MultiwatcherFactory() multiwatcher.Factory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiwatcherFactory")
	ret0, _ := ret[0].(multiwatcher.Factory)
	return ret0
}

// MultiwatcherFactory indicates an expected call of MultiwatcherFactory.
func (mr *MockContextMockRecorder) MultiwatcherFactory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiwatcherFactory", reflect.TypeOf((*MockContext)(nil).MultiwatcherFactory))
}

// ObjectStore mocks base method.
func (m *MockContext) ObjectStore() objectstore.ObjectStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStore")
	ret0, _ := ret[0].(objectstore.ObjectStore)
	return ret0
}

// ObjectStore indicates an expected call of ObjectStore.
func (mr *MockContextMockRecorder) ObjectStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStore", reflect.TypeOf((*MockContext)(nil).ObjectStore))
}

// Presence mocks base method.
func (m *MockContext) Presence() facade.Presence {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Presence")
	ret0, _ := ret[0].(facade.Presence)
	return ret0
}

// Presence indicates an expected call of Presence.
func (mr *MockContextMockRecorder) Presence() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Presence", reflect.TypeOf((*MockContext)(nil).Presence))
}

// RequestRecorder mocks base method.
func (m *MockContext) RequestRecorder() facade.RequestRecorder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestRecorder")
	ret0, _ := ret[0].(facade.RequestRecorder)
	return ret0
}

// RequestRecorder indicates an expected call of RequestRecorder.
func (mr *MockContextMockRecorder) RequestRecorder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestRecorder", reflect.TypeOf((*MockContext)(nil).RequestRecorder))
}

// Resources mocks base method.
func (m *MockContext) Resources() facade.Resources {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resources")
	ret0, _ := ret[0].(facade.Resources)
	return ret0
}

// Resources indicates an expected call of Resources.
func (mr *MockContextMockRecorder) Resources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resources", reflect.TypeOf((*MockContext)(nil).Resources))
}

// ServiceFactory mocks base method.
func (m *MockContext) ServiceFactory() servicefactory.ServiceFactory {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceFactory")
	ret0, _ := ret[0].(servicefactory.ServiceFactory)
	return ret0
}

// ServiceFactory indicates an expected call of ServiceFactory.
func (mr *MockContextMockRecorder) ServiceFactory() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceFactory", reflect.TypeOf((*MockContext)(nil).ServiceFactory))
}

// SingularClaimer mocks base method.
func (m *MockContext) SingularClaimer() (lease.Claimer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SingularClaimer")
	ret0, _ := ret[0].(lease.Claimer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SingularClaimer indicates an expected call of SingularClaimer.
func (mr *MockContextMockRecorder) SingularClaimer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SingularClaimer", reflect.TypeOf((*MockContext)(nil).SingularClaimer))
}

// State mocks base method.
func (m *MockContext) State() *state.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(*state.State)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockContextMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockContext)(nil).State))
}

// StatePool mocks base method.
func (m *MockContext) StatePool() *state.StatePool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatePool")
	ret0, _ := ret[0].(*state.StatePool)
	return ret0
}

// StatePool indicates an expected call of StatePool.
func (mr *MockContextMockRecorder) StatePool() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatePool", reflect.TypeOf((*MockContext)(nil).StatePool))
}

// WatcherRegistry mocks base method.
func (m *MockContext) WatcherRegistry() facade.WatcherRegistry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatcherRegistry")
	ret0, _ := ret[0].(facade.WatcherRegistry)
	return ret0
}

// WatcherRegistry indicates an expected call of WatcherRegistry.
func (mr *MockContextMockRecorder) WatcherRegistry() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatcherRegistry", reflect.TypeOf((*MockContext)(nil).WatcherRegistry))
}

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// AuthApplicationAgent mocks base method.
func (m *MockAuthorizer) AuthApplicationAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthApplicationAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthApplicationAgent indicates an expected call of AuthApplicationAgent.
func (mr *MockAuthorizerMockRecorder) AuthApplicationAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthApplicationAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthApplicationAgent))
}

// AuthClient mocks base method.
func (m *MockAuthorizer) AuthClient() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthClient")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthClient indicates an expected call of AuthClient.
func (mr *MockAuthorizerMockRecorder) AuthClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthClient", reflect.TypeOf((*MockAuthorizer)(nil).AuthClient))
}

// AuthController mocks base method.
func (m *MockAuthorizer) AuthController() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthController")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthController indicates an expected call of AuthController.
func (mr *MockAuthorizerMockRecorder) AuthController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthController", reflect.TypeOf((*MockAuthorizer)(nil).AuthController))
}

// AuthMachineAgent mocks base method.
func (m *MockAuthorizer) AuthMachineAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMachineAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthMachineAgent indicates an expected call of AuthMachineAgent.
func (mr *MockAuthorizerMockRecorder) AuthMachineAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMachineAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthMachineAgent))
}

// AuthModelAgent mocks base method.
func (m *MockAuthorizer) AuthModelAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthModelAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthModelAgent indicates an expected call of AuthModelAgent.
func (mr *MockAuthorizerMockRecorder) AuthModelAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthModelAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthModelAgent))
}

// AuthOwner mocks base method.
func (m *MockAuthorizer) AuthOwner(arg0 names.Tag) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthOwner", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthOwner indicates an expected call of AuthOwner.
func (mr *MockAuthorizerMockRecorder) AuthOwner(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthOwner", reflect.TypeOf((*MockAuthorizer)(nil).AuthOwner), arg0)
}

// AuthUnitAgent mocks base method.
func (m *MockAuthorizer) AuthUnitAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUnitAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthUnitAgent indicates an expected call of AuthUnitAgent.
func (mr *MockAuthorizerMockRecorder) AuthUnitAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUnitAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthUnitAgent))
}

// ConnectedModel mocks base method.
func (m *MockAuthorizer) ConnectedModel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedModel")
	ret0, _ := ret[0].(string)
	return ret0
}

// ConnectedModel indicates an expected call of ConnectedModel.
func (mr *MockAuthorizerMockRecorder) ConnectedModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedModel", reflect.TypeOf((*MockAuthorizer)(nil).ConnectedModel))
}

// EntityHasPermission mocks base method.
func (m *MockAuthorizer) EntityHasPermission(arg0 names.Tag, arg1 permission.Access, arg2 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntityHasPermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EntityHasPermission indicates an expected call of EntityHasPermission.
func (mr *MockAuthorizerMockRecorder) EntityHasPermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityHasPermission", reflect.TypeOf((*MockAuthorizer)(nil).EntityHasPermission), arg0, arg1, arg2)
}

// GetAuthTag mocks base method.
func (m *MockAuthorizer) GetAuthTag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthTag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// GetAuthTag indicates an expected call of GetAuthTag.
func (mr *MockAuthorizerMockRecorder) GetAuthTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthTag", reflect.TypeOf((*MockAuthorizer)(nil).GetAuthTag))
}

// HasPermission mocks base method.
func (m *MockAuthorizer) HasPermission(arg0 permission.Access, arg1 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HasPermission indicates an expected call of HasPermission.
func (mr *MockAuthorizerMockRecorder) HasPermission(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermission", reflect.TypeOf((*MockAuthorizer)(nil).HasPermission), arg0, arg1)
}
