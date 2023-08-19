// Code generated by MockGen. DO NOT EDIT.
// Source: sys_conn_oob.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ipv4 "golang.org/x/net/ipv4"
)

// MockBatchConn is a mock of batchConn interface.
type MockBatchConn struct {
	ctrl     *gomock.Controller
	recorder *MockBatchConnMockRecorder
}

// MockBatchConnMockRecorder is the mock recorder for MockBatchConn.
type MockBatchConnMockRecorder struct {
	mock *MockBatchConn
}

// NewMockBatchConn creates a new mock instance.
func NewMockBatchConn(ctrl *gomock.Controller) *MockBatchConn {
	mock := &MockBatchConn{ctrl: ctrl}
	mock.recorder = &MockBatchConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatchConn) EXPECT() *MockBatchConnMockRecorder {
	return m.recorder
}

// ReadBatch mocks base method.
func (m *MockBatchConn) ReadBatch(ms []ipv4.Message, flags int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBatch", ms, flags)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBatch indicates an expected call of ReadBatch.
func (mr *MockBatchConnMockRecorder) ReadBatch(ms, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBatch", reflect.TypeOf((*MockBatchConn)(nil).ReadBatch), ms, flags)
}