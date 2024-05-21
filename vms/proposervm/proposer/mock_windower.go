// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tenderly/avalanchego/vms/proposervm/proposer (interfaces: Windower)
//
// Generated by this command:
//
//	mockgen -package=proposer -destination=vms/proposervm/proposer/mock_windower.go github.com/tenderly/avalanchego/vms/proposervm/proposer Windower
//

// Package proposer is a generated GoMock package.
package proposer

import (
	context "context"
	reflect "reflect"
	time "time"

	ids "github.com/tenderly/avalanchego/ids"
	gomock "go.uber.org/mock/gomock"
)

// MockWindower is a mock of Windower interface.
type MockWindower struct {
	ctrl     *gomock.Controller
	recorder *MockWindowerMockRecorder
}

// MockWindowerMockRecorder is the mock recorder for MockWindower.
type MockWindowerMockRecorder struct {
	mock *MockWindower
}

// NewMockWindower creates a new mock instance.
func NewMockWindower(ctrl *gomock.Controller) *MockWindower {
	mock := &MockWindower{ctrl: ctrl}
	mock.recorder = &MockWindowerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWindower) EXPECT() *MockWindowerMockRecorder {
	return m.recorder
}

// Delay mocks base method.
func (m *MockWindower) Delay(arg0 context.Context, arg1, arg2 uint64, arg3 ids.NodeID, arg4 int) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delay indicates an expected call of Delay.
func (mr *MockWindowerMockRecorder) Delay(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockWindower)(nil).Delay), arg0, arg1, arg2, arg3, arg4)
}

// ExpectedProposer mocks base method.
func (m *MockWindower) ExpectedProposer(arg0 context.Context, arg1, arg2, arg3 uint64) (ids.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpectedProposer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(ids.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpectedProposer indicates an expected call of ExpectedProposer.
func (mr *MockWindowerMockRecorder) ExpectedProposer(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpectedProposer", reflect.TypeOf((*MockWindower)(nil).ExpectedProposer), arg0, arg1, arg2, arg3)
}

// MinDelayForProposer mocks base method.
func (m *MockWindower) MinDelayForProposer(arg0 context.Context, arg1, arg2 uint64, arg3 ids.NodeID, arg4 uint64) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinDelayForProposer", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MinDelayForProposer indicates an expected call of MinDelayForProposer.
func (mr *MockWindowerMockRecorder) MinDelayForProposer(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinDelayForProposer", reflect.TypeOf((*MockWindower)(nil).MinDelayForProposer), arg0, arg1, arg2, arg3, arg4)
}

// Proposers mocks base method.
func (m *MockWindower) Proposers(arg0 context.Context, arg1, arg2 uint64, arg3 int) ([]ids.NodeID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proposers", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]ids.NodeID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Proposers indicates an expected call of Proposers.
func (mr *MockWindowerMockRecorder) Proposers(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proposers", reflect.TypeOf((*MockWindower)(nil).Proposers), arg0, arg1, arg2, arg3)
}
