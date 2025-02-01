// Code generated by MockGen. DO NOT EDIT.
// Source: payment-service/gapi/helpers (interfaces: Helper)

// Package helpmck is a generated GoMock package.
package helpmck

import (
	context "context"
	payment "payment-service/payment/payment-service"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	stripe "github.com/stripe/stripe-go/v81"
)

// MockHelper is a mock of Helper interface.
type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

// MockHelperMockRecorder is the mock recorder for MockHelper.
type MockHelperMockRecorder struct {
	mock *MockHelper
}

// NewMockHelper creates a new mock instance.
func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelper) EXPECT() *MockHelperMockRecorder {
	return m.recorder
}

// GetOrCreateCustomer mocks base method.
func (m *MockHelper) GetOrCreateCustomer(arg0 context.Context, arg1, arg2, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreateCustomer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreateCustomer indicates an expected call of GetOrCreateCustomer.
func (mr *MockHelperMockRecorder) GetOrCreateCustomer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreateCustomer", reflect.TypeOf((*MockHelper)(nil).GetOrCreateCustomer), arg0, arg1, arg2, arg3)
}

// HandlePaymentIfSuccesful mocks base method.
func (m *MockHelper) HandlePaymentIfSuccesful(arg0 context.Context, arg1 *stripe.PaymentIntent) (*payment.WebhookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandlePaymentIfSuccesful", arg0, arg1)
	ret0, _ := ret[0].(*payment.WebhookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandlePaymentIfSuccesful indicates an expected call of HandlePaymentIfSuccesful.
func (mr *MockHelperMockRecorder) HandlePaymentIfSuccesful(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandlePaymentIfSuccesful", reflect.TypeOf((*MockHelper)(nil).HandlePaymentIfSuccesful), arg0, arg1)
}
