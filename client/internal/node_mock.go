// Code generated by MockGen. DO NOT EDIT.
// Source: client/internal/node.go

// Package internal is a generated GoMock package.
package internal

import (
	gomock "github.com/golang/mock/gomock"
	protobuf "github.com/vacp2p/mvds/protobuf"
	state "github.com/vacp2p/mvds/state"
	reflect "reflect"
)

// MockDataSyncNode is a mock of DataSyncNode interface
type MockDataSyncNode struct {
	ctrl     *gomock.Controller
	recorder *MockDataSyncNodeMockRecorder
}

// MockDataSyncNodeMockRecorder is the mock recorder for MockDataSyncNode
type MockDataSyncNodeMockRecorder struct {
	mock *MockDataSyncNode
}

// NewMockDataSyncNode creates a new mock instance
func NewMockDataSyncNode(ctrl *gomock.Controller) *MockDataSyncNode {
	mock := &MockDataSyncNode{ctrl: ctrl}
	mock.recorder = &MockDataSyncNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataSyncNode) EXPECT() *MockDataSyncNodeMockRecorder {
	return m.recorder
}

// AppendMessage mocks base method
func (m *MockDataSyncNode) AppendMessage(groupID state.GroupID, data []byte) (state.MessageID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendMessage", groupID, data)
	ret0, _ := ret[0].(state.MessageID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppendMessage indicates an expected call of AppendMessage
func (mr *MockDataSyncNodeMockRecorder) AppendMessage(groupID, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendMessage", reflect.TypeOf((*MockDataSyncNode)(nil).AppendMessage), groupID, data)
}

// Subscribe mocks base method
func (m *MockDataSyncNode) Subscribe() chan protobuf.Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe")
	ret0, _ := ret[0].(chan protobuf.Message)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockDataSyncNodeMockRecorder) Subscribe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockDataSyncNode)(nil).Subscribe))
}

// RequestMessage mocks base method
func (m *MockDataSyncNode) RequestMessage(group state.GroupID, id state.MessageID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestMessage", group, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestMessage indicates an expected call of RequestMessage
func (mr *MockDataSyncNodeMockRecorder) RequestMessage(group, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestMessage", reflect.TypeOf((*MockDataSyncNode)(nil).RequestMessage), group, id)
}
