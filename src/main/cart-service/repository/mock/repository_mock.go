// Code generated by MockGen. DO NOT EDIT.
// Source: main/cart-service/repository (interfaces: IRepository)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	structs "main/cart-service/structs"
	reflect "reflect"
)

// MockIRepository is a mock of IRepository interface
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// AddToCart mocks base method
func (m *MockIRepository) AddToCart(arg0 int64, arg1 []structs.CartItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCart indicates an expected call of AddToCart
func (mr *MockIRepositoryMockRecorder) AddToCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockIRepository)(nil).AddToCart), arg0, arg1)
}

// ClearCart mocks base method
func (m *MockIRepository) ClearCart(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearCart", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearCart indicates an expected call of ClearCart
func (mr *MockIRepositoryMockRecorder) ClearCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearCart", reflect.TypeOf((*MockIRepository)(nil).ClearCart), arg0)
}

// GetCart mocks base method
func (m *MockIRepository) GetCart(arg0 int64) ([]structs.CartItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0)
	ret0, _ := ret[0].([]structs.CartItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart
func (mr *MockIRepositoryMockRecorder) GetCart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockIRepository)(nil).GetCart), arg0)
}
