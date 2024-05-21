// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tenderly/avalanchego/snow/engine/avalanche/vertex (interfaces: LinearizableVM)
//
// Generated by this command:
//
//	mockgen -package=vertex -destination=snow/engine/avalanche/vertex/mock_vm.go github.com/tenderly/avalanchego/snow/engine/avalanche/vertex LinearizableVM
//

// Package vertex is a generated GoMock package.
package vertex

import (
	context "context"
	http "net/http"
	reflect "reflect"
	time "time"

	database "github.com/tenderly/avalanchego/database"
	ids "github.com/tenderly/avalanchego/ids"
	snow "github.com/tenderly/avalanchego/snow"
	snowman "github.com/tenderly/avalanchego/snow/consensus/snowman"
	snowstorm "github.com/tenderly/avalanchego/snow/consensus/snowstorm"
	common "github.com/tenderly/avalanchego/snow/engine/common"
	version "github.com/tenderly/avalanchego/version"
	gomock "go.uber.org/mock/gomock"
)

// MockLinearizableVM is a mock of LinearizableVM interface.
type MockLinearizableVM struct {
	ctrl     *gomock.Controller
	recorder *MockLinearizableVMMockRecorder
}

// MockLinearizableVMMockRecorder is the mock recorder for MockLinearizableVM.
type MockLinearizableVMMockRecorder struct {
	mock *MockLinearizableVM
}

// NewMockLinearizableVM creates a new mock instance.
func NewMockLinearizableVM(ctrl *gomock.Controller) *MockLinearizableVM {
	mock := &MockLinearizableVM{ctrl: ctrl}
	mock.recorder = &MockLinearizableVMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLinearizableVM) EXPECT() *MockLinearizableVMMockRecorder {
	return m.recorder
}

// AppGossip mocks base method.
func (m *MockLinearizableVM) AppGossip(arg0 context.Context, arg1 ids.NodeID, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGossip", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppGossip indicates an expected call of AppGossip.
func (mr *MockLinearizableVMMockRecorder) AppGossip(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGossip", reflect.TypeOf((*MockLinearizableVM)(nil).AppGossip), arg0, arg1, arg2)
}

// AppRequest mocks base method.
func (m *MockLinearizableVM) AppRequest(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 time.Time, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequest indicates an expected call of AppRequest.
func (mr *MockLinearizableVMMockRecorder) AppRequest(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequest", reflect.TypeOf((*MockLinearizableVM)(nil).AppRequest), arg0, arg1, arg2, arg3, arg4)
}

// AppRequestFailed mocks base method.
func (m *MockLinearizableVM) AppRequestFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 *common.AppError) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequestFailed", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequestFailed indicates an expected call of AppRequestFailed.
func (mr *MockLinearizableVMMockRecorder) AppRequestFailed(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequestFailed", reflect.TypeOf((*MockLinearizableVM)(nil).AppRequestFailed), arg0, arg1, arg2, arg3)
}

// AppResponse mocks base method.
func (m *MockLinearizableVM) AppResponse(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppResponse indicates an expected call of AppResponse.
func (mr *MockLinearizableVMMockRecorder) AppResponse(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*MockLinearizableVM)(nil).AppResponse), arg0, arg1, arg2, arg3)
}

// BuildBlock mocks base method.
func (m *MockLinearizableVM) BuildBlock(arg0 context.Context) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildBlock indicates an expected call of BuildBlock.
func (mr *MockLinearizableVMMockRecorder) BuildBlock(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildBlock", reflect.TypeOf((*MockLinearizableVM)(nil).BuildBlock), arg0)
}

// Connected mocks base method.
func (m *MockLinearizableVM) Connected(arg0 context.Context, arg1 ids.NodeID, arg2 *version.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connected indicates an expected call of Connected.
func (mr *MockLinearizableVMMockRecorder) Connected(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockLinearizableVM)(nil).Connected), arg0, arg1, arg2)
}

// CreateHandlers mocks base method.
func (m *MockLinearizableVM) CreateHandlers(arg0 context.Context) (map[string]http.Handler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHandlers", arg0)
	ret0, _ := ret[0].(map[string]http.Handler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHandlers indicates an expected call of CreateHandlers.
func (mr *MockLinearizableVMMockRecorder) CreateHandlers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHandlers", reflect.TypeOf((*MockLinearizableVM)(nil).CreateHandlers), arg0)
}

// CrossChainAppRequest mocks base method.
func (m *MockLinearizableVM) CrossChainAppRequest(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 time.Time, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppRequest indicates an expected call of CrossChainAppRequest.
func (mr *MockLinearizableVMMockRecorder) CrossChainAppRequest(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppRequest", reflect.TypeOf((*MockLinearizableVM)(nil).CrossChainAppRequest), arg0, arg1, arg2, arg3, arg4)
}

// CrossChainAppRequestFailed mocks base method.
func (m *MockLinearizableVM) CrossChainAppRequestFailed(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 *common.AppError) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppRequestFailed", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppRequestFailed indicates an expected call of CrossChainAppRequestFailed.
func (mr *MockLinearizableVMMockRecorder) CrossChainAppRequestFailed(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppRequestFailed", reflect.TypeOf((*MockLinearizableVM)(nil).CrossChainAppRequestFailed), arg0, arg1, arg2, arg3)
}

// CrossChainAppResponse mocks base method.
func (m *MockLinearizableVM) CrossChainAppResponse(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppResponse indicates an expected call of CrossChainAppResponse.
func (mr *MockLinearizableVMMockRecorder) CrossChainAppResponse(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppResponse", reflect.TypeOf((*MockLinearizableVM)(nil).CrossChainAppResponse), arg0, arg1, arg2, arg3)
}

// Disconnected mocks base method.
func (m *MockLinearizableVM) Disconnected(arg0 context.Context, arg1 ids.NodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnected", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockLinearizableVMMockRecorder) Disconnected(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockLinearizableVM)(nil).Disconnected), arg0, arg1)
}

// GetBlock mocks base method.
func (m *MockLinearizableVM) GetBlock(arg0 context.Context, arg1 ids.ID) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0, arg1)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockLinearizableVMMockRecorder) GetBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockLinearizableVM)(nil).GetBlock), arg0, arg1)
}

// GetBlockIDAtHeight mocks base method.
func (m *MockLinearizableVM) GetBlockIDAtHeight(arg0 context.Context, arg1 uint64) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockIDAtHeight", arg0, arg1)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockIDAtHeight indicates an expected call of GetBlockIDAtHeight.
func (mr *MockLinearizableVMMockRecorder) GetBlockIDAtHeight(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockIDAtHeight", reflect.TypeOf((*MockLinearizableVM)(nil).GetBlockIDAtHeight), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockLinearizableVM) HealthCheck(arg0 context.Context) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockLinearizableVMMockRecorder) HealthCheck(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockLinearizableVM)(nil).HealthCheck), arg0)
}

// Initialize mocks base method.
func (m *MockLinearizableVM) Initialize(arg0 context.Context, arg1 *snow.Context, arg2 database.Database, arg3, arg4, arg5 []byte, arg6 chan<- common.Message, arg7 []*common.Fx, arg8 common.AppSender) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockLinearizableVMMockRecorder) Initialize(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockLinearizableVM)(nil).Initialize), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// LastAccepted mocks base method.
func (m *MockLinearizableVM) LastAccepted(arg0 context.Context) (ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastAccepted", arg0)
	ret0, _ := ret[0].(ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastAccepted indicates an expected call of LastAccepted.
func (mr *MockLinearizableVMMockRecorder) LastAccepted(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastAccepted", reflect.TypeOf((*MockLinearizableVM)(nil).LastAccepted), arg0)
}

// Linearize mocks base method.
func (m *MockLinearizableVM) Linearize(arg0 context.Context, arg1 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Linearize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Linearize indicates an expected call of Linearize.
func (mr *MockLinearizableVMMockRecorder) Linearize(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Linearize", reflect.TypeOf((*MockLinearizableVM)(nil).Linearize), arg0, arg1)
}

// ParseBlock mocks base method.
func (m *MockLinearizableVM) ParseBlock(arg0 context.Context, arg1 []byte) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBlock", arg0, arg1)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseBlock indicates an expected call of ParseBlock.
func (mr *MockLinearizableVMMockRecorder) ParseBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBlock", reflect.TypeOf((*MockLinearizableVM)(nil).ParseBlock), arg0, arg1)
}

// ParseTx mocks base method.
func (m *MockLinearizableVM) ParseTx(arg0 context.Context, arg1 []byte) (snowstorm.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseTx", arg0, arg1)
	ret0, _ := ret[0].(snowstorm.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseTx indicates an expected call of ParseTx.
func (mr *MockLinearizableVMMockRecorder) ParseTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseTx", reflect.TypeOf((*MockLinearizableVM)(nil).ParseTx), arg0, arg1)
}

// SetPreference mocks base method.
func (m *MockLinearizableVM) SetPreference(arg0 context.Context, arg1 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPreference", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPreference indicates an expected call of SetPreference.
func (mr *MockLinearizableVMMockRecorder) SetPreference(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPreference", reflect.TypeOf((*MockLinearizableVM)(nil).SetPreference), arg0, arg1)
}

// SetState mocks base method.
func (m *MockLinearizableVM) SetState(arg0 context.Context, arg1 snow.State) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState.
func (mr *MockLinearizableVMMockRecorder) SetState(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockLinearizableVM)(nil).SetState), arg0, arg1)
}

// Shutdown mocks base method.
func (m *MockLinearizableVM) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockLinearizableVMMockRecorder) Shutdown(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockLinearizableVM)(nil).Shutdown), arg0)
}

// Version mocks base method.
func (m *MockLinearizableVM) Version(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockLinearizableVMMockRecorder) Version(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockLinearizableVM)(nil).Version), arg0)
}
