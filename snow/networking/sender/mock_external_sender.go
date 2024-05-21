// Code generated by MockGen. DO NOT EDIT.
// Source: snow/networking/sender/external_sender.go
//
// Generated by this command:
//
//	mockgen -source=snow/networking/sender/external_sender.go -destination=snow/networking/sender/mock_external_sender.go -package=sender -exclude_interfaces=
//

// Package sender is a generated GoMock package.
package sender

import (
	reflect "reflect"

	ids "github.com/tenderly/avalanchego/ids"
	message "github.com/tenderly/avalanchego/message"
	common "github.com/tenderly/avalanchego/snow/engine/common"
	subnets "github.com/tenderly/avalanchego/subnets"
	set "github.com/tenderly/avalanchego/utils/set"
	gomock "go.uber.org/mock/gomock"
)

// MockExternalSender is a mock of ExternalSender interface.
type MockExternalSender struct {
	ctrl     *gomock.Controller
	recorder *MockExternalSenderMockRecorder
}

// MockExternalSenderMockRecorder is the mock recorder for MockExternalSender.
type MockExternalSenderMockRecorder struct {
	mock *MockExternalSender
}

// NewMockExternalSender creates a new mock instance.
func NewMockExternalSender(ctrl *gomock.Controller) *MockExternalSender {
	mock := &MockExternalSender{ctrl: ctrl}
	mock.recorder = &MockExternalSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalSender) EXPECT() *MockExternalSenderMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockExternalSender) Send(msg message.OutboundMessage, config common.SendConfig, subnetID ids.ID, allower subnets.Allower) set.Set[ids.NodeID] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", msg, config, subnetID, allower)
	ret0, _ := ret[0].(set.Set[ids.NodeID])
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockExternalSenderMockRecorder) Send(msg, config, subnetID, allower any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockExternalSender)(nil).Send), msg, config, subnetID, allower)
}
