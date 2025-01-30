// Code generated by MockGen. DO NOT EDIT.
// Source: payment-service/stripe (interfaces: Stripe)

// Package stripemck is a generated GoMock package.
package stripemck

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	stripe "github.com/stripe/stripe-go/v81"
)

// MockStripe is a mock of Stripe interface.
type MockStripe struct {
	ctrl     *gomock.Controller
	recorder *MockStripeMockRecorder
}

// MockStripeMockRecorder is the mock recorder for MockStripe.
type MockStripeMockRecorder struct {
	mock *MockStripe
}

// NewMockStripe creates a new mock instance.
func NewMockStripe(ctrl *gomock.Controller) *MockStripe {
	mock := &MockStripe{ctrl: ctrl}
	mock.recorder = &MockStripeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStripe) EXPECT() *MockStripeMockRecorder {
	return m.recorder
}

// NewPaymentIntent mocks base method.
func (m *MockStripe) NewPaymentIntent(arg0 *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPaymentIntent", arg0)
	ret0, _ := ret[0].(*stripe.PaymentIntent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPaymentIntent indicates an expected call of NewPaymentIntent.
func (mr *MockStripeMockRecorder) NewPaymentIntent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPaymentIntent", reflect.TypeOf((*MockStripe)(nil).NewPaymentIntent), arg0)
}
