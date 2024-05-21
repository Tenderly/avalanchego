// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tenderly/avalanchego/vms (interfaces: Factory,Manager)
//
// Generated by this command:
//
//	mockgen -package=vms -destination=vms/mock_manager.go github.com/tenderly/avalanchego/vms Factory,Manager
//

// Package vms is a generated GoMock package.
package vms

import (
	context "context"
	reflect "reflect"

	ids "github.com/tenderly/avalanchego/ids"
	logging "github.com/tenderly/avalanchego/utils/logging"
	gomock "go.uber.org/mock/gomock"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// New mocks base method.
func (m *MockFactory) New(arg0 logging.Logger) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockFactoryMockRecorder) New(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockFactory)(nil).New), arg0)
}

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Alias mocks base method.
func (m *MockManager) Alias(arg0 ids.ID, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alias", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Alias indicates an expected call of Alias.
func (mr *MockManagerMockRecorder) Alias(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alias", reflect.TypeOf((*MockManager)(nil).Alias), arg0, arg1)
}

// Aliases mocks base method.
func (m *MockManager) Aliases(arg0 ids.ID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aliases", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aliases indicates an expected call of Aliases.
func (mr *MockManagerMockRecorder) Aliases(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aliases", reflect.TypeOf((*MockManager)(nil).Aliases), arg0)
}

// GetFactory mocks base method.
func (m *MockManager) GetFactory(arg0 ids.ID) (Factory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFactory", arg0)
	ret0, _ := ret[0].(Factory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFactory indicates an expected call of GetFactory.
func (mr *MockManagerMockRecorder) GetFactory(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFactory", reflect.TypeOf((*MockManager)(nil).GetFactory), arg0)
}

// ListFactories mocks base method.
func (m *MockManager) ListFactories() ([]ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFactories")
	ret0, _ := ret[0].([]ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFactories indicates an expected call of ListFactories.
func (mr *MockManagerMockRecorder) ListFactories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFactories", reflect.TypeOf((*MockManager)(nil).ListFactories))
}

// Lookup mocks base method.
func (m *MockManager) Lookup(arg0 string) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup.
func (mr *MockManagerMockRecorder) Lookup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockManager)(nil).Lookup), arg0)
}

// PrimaryAlias mocks base method.
func (m *MockManager) PrimaryAlias(arg0 ids.ID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrimaryAlias", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrimaryAlias indicates an expected call of PrimaryAlias.
func (mr *MockManagerMockRecorder) PrimaryAlias(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrimaryAlias", reflect.TypeOf((*MockManager)(nil).PrimaryAlias), arg0)
}

// PrimaryAliasOrDefault mocks base method.
func (m *MockManager) PrimaryAliasOrDefault(arg0 ids.ID) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrimaryAliasOrDefault", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// PrimaryAliasOrDefault indicates an expected call of PrimaryAliasOrDefault.
func (mr *MockManagerMockRecorder) PrimaryAliasOrDefault(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrimaryAliasOrDefault", reflect.TypeOf((*MockManager)(nil).PrimaryAliasOrDefault), arg0)
}

// RegisterFactory mocks base method.
func (m *MockManager) RegisterFactory(arg0 context.Context, arg1 ids.ID, arg2 Factory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFactory", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterFactory indicates an expected call of RegisterFactory.
func (mr *MockManagerMockRecorder) RegisterFactory(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFactory", reflect.TypeOf((*MockManager)(nil).RegisterFactory), arg0, arg1, arg2)
}

// RemoveAliases mocks base method.
func (m *MockManager) RemoveAliases(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAliases", arg0)
}

// RemoveAliases indicates an expected call of RemoveAliases.
func (mr *MockManagerMockRecorder) RemoveAliases(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAliases", reflect.TypeOf((*MockManager)(nil).RemoveAliases), arg0)
}

// Versions mocks base method.
func (m *MockManager) Versions() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Versions")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Versions indicates an expected call of Versions.
func (mr *MockManagerMockRecorder) Versions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Versions", reflect.TypeOf((*MockManager)(nil).Versions))
}
