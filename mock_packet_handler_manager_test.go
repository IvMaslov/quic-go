// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: PacketHandlerManager)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/IvMaslov/quic-go/internal/protocol"
)

// MockPacketHandlerManager is a mock of PacketHandlerManager interface.
type MockPacketHandlerManager struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerManagerMockRecorder
}

// MockPacketHandlerManagerMockRecorder is the mock recorder for MockPacketHandlerManager.
type MockPacketHandlerManagerMockRecorder struct {
	mock *MockPacketHandlerManager
}

// NewMockPacketHandlerManager creates a new mock instance.
func NewMockPacketHandlerManager(ctrl *gomock.Controller) *MockPacketHandlerManager {
	mock := &MockPacketHandlerManager{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketHandlerManager) EXPECT() *MockPacketHandlerManagerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPacketHandlerManager) Add(arg0 protocol.ConnectionID, arg1 packetHandler) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockPacketHandlerManagerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPacketHandlerManager)(nil).Add), arg0, arg1)
}

// AddResetToken mocks base method.
func (m *MockPacketHandlerManager) AddResetToken(arg0 protocol.StatelessResetToken, arg1 packetHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResetToken", arg0, arg1)
}

// AddResetToken indicates an expected call of AddResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) AddResetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).AddResetToken), arg0, arg1)
}

// AddWithConnID mocks base method.
func (m *MockPacketHandlerManager) AddWithConnID(arg0, arg1 protocol.ConnectionID, arg2 func() (packetHandler, bool)) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWithConnID", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddWithConnID indicates an expected call of AddWithConnID.
func (mr *MockPacketHandlerManagerMockRecorder) AddWithConnID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWithConnID", reflect.TypeOf((*MockPacketHandlerManager)(nil).AddWithConnID), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockPacketHandlerManager) Close(arg0 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close", arg0)
}

// Close indicates an expected call of Close.
func (mr *MockPacketHandlerManagerMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPacketHandlerManager)(nil).Close), arg0)
}

// CloseServer mocks base method.
func (m *MockPacketHandlerManager) CloseServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseServer")
}

// CloseServer indicates an expected call of CloseServer.
func (mr *MockPacketHandlerManagerMockRecorder) CloseServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseServer", reflect.TypeOf((*MockPacketHandlerManager)(nil).CloseServer))
}

// Get mocks base method.
func (m *MockPacketHandlerManager) Get(arg0 protocol.ConnectionID) (packetHandler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(packetHandler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPacketHandlerManagerMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPacketHandlerManager)(nil).Get), arg0)
}

// GetByResetToken mocks base method.
func (m *MockPacketHandlerManager) GetByResetToken(arg0 protocol.StatelessResetToken) (packetHandler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByResetToken", arg0)
	ret0, _ := ret[0].(packetHandler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetByResetToken indicates an expected call of GetByResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) GetByResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).GetByResetToken), arg0)
}

// GetStatelessResetToken mocks base method.
func (m *MockPacketHandlerManager) GetStatelessResetToken(arg0 protocol.ConnectionID) protocol.StatelessResetToken {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessResetToken", arg0)
	ret0, _ := ret[0].(protocol.StatelessResetToken)
	return ret0
}

// GetStatelessResetToken indicates an expected call of GetStatelessResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) GetStatelessResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).GetStatelessResetToken), arg0)
}

// Remove mocks base method.
func (m *MockPacketHandlerManager) Remove(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockPacketHandlerManagerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPacketHandlerManager)(nil).Remove), arg0)
}

// RemoveResetToken mocks base method.
func (m *MockPacketHandlerManager) RemoveResetToken(arg0 protocol.StatelessResetToken) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveResetToken", arg0)
}

// RemoveResetToken indicates an expected call of RemoveResetToken.
func (mr *MockPacketHandlerManagerMockRecorder) RemoveResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResetToken", reflect.TypeOf((*MockPacketHandlerManager)(nil).RemoveResetToken), arg0)
}

// ReplaceWithClosed mocks base method.
func (m *MockPacketHandlerManager) ReplaceWithClosed(arg0 []protocol.ConnectionID, arg1 protocol.Perspective, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReplaceWithClosed", arg0, arg1, arg2)
}

// ReplaceWithClosed indicates an expected call of ReplaceWithClosed.
func (mr *MockPacketHandlerManagerMockRecorder) ReplaceWithClosed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceWithClosed", reflect.TypeOf((*MockPacketHandlerManager)(nil).ReplaceWithClosed), arg0, arg1, arg2)
}

// Retire mocks base method.
func (m *MockPacketHandlerManager) Retire(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Retire", arg0)
}

// Retire indicates an expected call of Retire.
func (mr *MockPacketHandlerManagerMockRecorder) Retire(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retire", reflect.TypeOf((*MockPacketHandlerManager)(nil).Retire), arg0)
}