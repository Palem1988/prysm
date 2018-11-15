// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1 (interfaces: ValidatorServiceServer)

package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/prysmaticlabs/prysm/proto/beacon/rpc/v1"
)

// MockValidatorServiceServer is a mock of ValidatorServiceServer interface
type MockValidatorServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorServiceServerMockRecorder
}

// MockValidatorServiceServerMockRecorder is the mock recorder for MockValidatorServiceServer
type MockValidatorServiceServerMockRecorder struct {
	mock *MockValidatorServiceServer
}

// NewMockValidatorServiceServer creates a new mock instance
func NewMockValidatorServiceServer(ctrl *gomock.Controller) *MockValidatorServiceServer {
	mock := &MockValidatorServiceServer{ctrl: ctrl}
	mock.recorder = &MockValidatorServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorServiceServer) EXPECT() *MockValidatorServiceServerMockRecorder {
	return m.recorder
}

// ValidatorIndex mocks base method
func (m *MockValidatorServiceServer) ValidatorIndex(arg0 context.Context, arg1 *v1.PublicKey) (*v1.IndexResponse, error) {
	ret := m.ctrl.Call(m, "ValidatorIndex", arg0, arg1)
	ret0, _ := ret[0].(*v1.IndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorIndex indicates an expected call of ValidatorIndex
func (mr *MockValidatorServiceServerMockRecorder) ValidatorIndex(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorIndex", reflect.TypeOf((*MockValidatorServiceServer)(nil).ValidatorIndex), arg0, arg1)
}

// ValidatorShardID mocks base method
func (m *MockValidatorServiceServer) ValidatorShardID(arg0 context.Context, arg1 *v1.PublicKey) (*v1.ShardIDResponse, error) {
	ret := m.ctrl.Call(m, "ValidatorShardID", arg0, arg1)
	ret0, _ := ret[0].(*v1.ShardIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorShardID indicates an expected call of ValidatorShardID
func (mr *MockValidatorServiceServerMockRecorder) ValidatorShardID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorShardID", reflect.TypeOf((*MockValidatorServiceServer)(nil).ValidatorShardID), arg0, arg1)
}

// ValidatorSlotAndResponsibility mocks base method
func (m *MockValidatorServiceServer) ValidatorSlotAndResponsibility(arg0 context.Context, arg1 *v1.PublicKey) (*v1.SlotResponsibilityResponse, error) {
	ret := m.ctrl.Call(m, "ValidatorSlotAndResponsibility", arg0, arg1)
	ret0, _ := ret[0].(*v1.SlotResponsibilityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorSlotAndResponsibility indicates an expected call of ValidatorSlotAndResponsibility
func (mr *MockValidatorServiceServerMockRecorder) ValidatorSlotAndResponsibility(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorSlotAndResponsibility", reflect.TypeOf((*MockValidatorServiceServer)(nil).ValidatorSlotAndResponsibility), arg0, arg1)
}