// Code generated by MockGen. DO NOT EDIT.
// Source: cosmos.go

// Package mock_gremcos is a generated GoMock package.
package mock_gremcos

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/supplyon/gremcos/interfaces"
)

// MockCosmos is a mock of Cosmos interface.
type MockCosmos struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosMockRecorder
}

// MockCosmosMockRecorder is the mock recorder for MockCosmos.
type MockCosmosMockRecorder struct {
	mock *MockCosmos
}

// NewMockCosmos creates a new mock instance.
func NewMockCosmos(ctrl *gomock.Controller) *MockCosmos {
	mock := &MockCosmos{ctrl: ctrl}
	mock.recorder = &MockCosmosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCosmos) EXPECT() *MockCosmosMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCosmos) Execute(query string) ([]interfaces.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", query)
	ret0, _ := ret[0].([]interfaces.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCosmosMockRecorder) Execute(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCosmos)(nil).Execute), query)
}

// ExecuteAsync mocks base method.
func (m *MockCosmos) ExecuteAsync(query string, responseChannel chan interfaces.AsyncResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteAsync", query, responseChannel)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteAsync indicates an expected call of ExecuteAsync.
func (mr *MockCosmosMockRecorder) ExecuteAsync(query, responseChannel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAsync", reflect.TypeOf((*MockCosmos)(nil).ExecuteAsync), query, responseChannel)
}

// ExecuteQuery mocks base method.
func (m *MockCosmos) ExecuteQuery(query interfaces.QueryBuilder) ([]interfaces.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteQuery", query)
	ret0, _ := ret[0].([]interfaces.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteQuery indicates an expected call of ExecuteQuery.
func (mr *MockCosmosMockRecorder) ExecuteQuery(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteQuery", reflect.TypeOf((*MockCosmos)(nil).ExecuteQuery), query)
}

// IsConnected mocks base method.
func (m *MockCosmos) IsConnected() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnected")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConnected indicates an expected call of IsConnected.
func (mr *MockCosmosMockRecorder) IsConnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnected", reflect.TypeOf((*MockCosmos)(nil).IsConnected))
}

// IsHealthy mocks base method.
func (m *MockCosmos) IsHealthy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHealthy")
	ret0, _ := ret[0].(error)
	return ret0
}

// IsHealthy indicates an expected call of IsHealthy.
func (mr *MockCosmosMockRecorder) IsHealthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealthy", reflect.TypeOf((*MockCosmos)(nil).IsHealthy))
}

// Stop mocks base method.
func (m *MockCosmos) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockCosmosMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCosmos)(nil).Stop))
}

// String mocks base method.
func (m *MockCosmos) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockCosmosMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockCosmos)(nil).String))
}