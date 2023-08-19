// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/http3 (interfaces: QUICEarlyListener)

// Package http3 is a generated GoMock package.
package http3

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	quic "github.com/IvMaslov/quic-go"
)

// MockQUICEarlyListener is a mock of QUICEarlyListener interface.
type MockQUICEarlyListener struct {
	ctrl     *gomock.Controller
	recorder *MockQUICEarlyListenerMockRecorder
}

// MockQUICEarlyListenerMockRecorder is the mock recorder for MockQUICEarlyListener.
type MockQUICEarlyListenerMockRecorder struct {
	mock *MockQUICEarlyListener
}

// NewMockQUICEarlyListener creates a new mock instance.
func NewMockQUICEarlyListener(ctrl *gomock.Controller) *MockQUICEarlyListener {
	mock := &MockQUICEarlyListener{ctrl: ctrl}
	mock.recorder = &MockQUICEarlyListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQUICEarlyListener) EXPECT() *MockQUICEarlyListenerMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockQUICEarlyListener) Accept(arg0 context.Context) (quic.EarlyConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(quic.EarlyConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept.
func (mr *MockQUICEarlyListenerMockRecorder) Accept(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockQUICEarlyListener)(nil).Accept), arg0)
}

// Addr mocks base method.
func (m *MockQUICEarlyListener) Addr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Addr indicates an expected call of Addr.
func (mr *MockQUICEarlyListenerMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockQUICEarlyListener)(nil).Addr))
}

// Close mocks base method.
func (m *MockQUICEarlyListener) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockQUICEarlyListenerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQUICEarlyListener)(nil).Close))
}
