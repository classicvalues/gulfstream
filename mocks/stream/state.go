// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/xxx/projects/gulfstream/pkg/stream/state.go

// Package mockstream is a generated GoMock package.
package mockstream

import (
	event "github.com/go-gulfstream/gulfstream/pkg/event"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockState is a mock of State interface
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Mutate mocks base method
func (m *MockState) Mutate(arg0 *event.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Mutate", arg0)
}

// Mutate indicates an expected call of Mutate
func (mr *MockStateMockRecorder) Mutate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mutate", reflect.TypeOf((*MockState)(nil).Mutate), arg0)
}

// UnmarshalBinary mocks base method
func (m *MockState) UnmarshalBinary(data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalBinary", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalBinary indicates an expected call of UnmarshalBinary
func (mr *MockStateMockRecorder) UnmarshalBinary(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalBinary", reflect.TypeOf((*MockState)(nil).UnmarshalBinary), data)
}

// MarshalBinary mocks base method
func (m *MockState) MarshalBinary() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalBinary")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalBinary indicates an expected call of MarshalBinary
func (mr *MockStateMockRecorder) MarshalBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalBinary", reflect.TypeOf((*MockState)(nil).MarshalBinary))
}
