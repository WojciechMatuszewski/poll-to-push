// Code generated by MockGen. DO NOT EDIT.
// Source: ../handler.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStepFunctionStarter is a mock of StepFunctionStarter interface
type MockStepFunctionStarter struct {
	ctrl     *gomock.Controller
	recorder *MockStepFunctionStarterMockRecorder
}

// MockStepFunctionStarterMockRecorder is the mock recorder for MockStepFunctionStarter
type MockStepFunctionStarterMockRecorder struct {
	mock *MockStepFunctionStarter
}

// NewMockStepFunctionStarter creates a new mock instance
func NewMockStepFunctionStarter(ctrl *gomock.Controller) *MockStepFunctionStarter {
	mock := &MockStepFunctionStarter{ctrl: ctrl}
	mock.recorder = &MockStepFunctionStarterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStepFunctionStarter) EXPECT() *MockStepFunctionStarterMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockStepFunctionStarter) Start(ctx context.Context, input, id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx, input, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start
func (mr *MockStepFunctionStarterMockRecorder) Start(ctx, input, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStepFunctionStarter)(nil).Start), ctx, input, id)
}