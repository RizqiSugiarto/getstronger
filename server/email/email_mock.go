// Code generated by MockGen. DO NOT EDIT.
// Source: email.go
//
// Generated by this command:
//
//	mockgen -package email -source=email.go -destination=email_mock.go Email
//

// Package email is a generated GoMock package.
package email

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEmail is a mock of Email interface.
type MockEmail struct {
	ctrl     *gomock.Controller
	recorder *MockEmailMockRecorder
	isgomock struct{}
}

// MockEmailMockRecorder is the mock recorder for MockEmail.
type MockEmailMockRecorder struct {
	mock *MockEmail
}

// NewMockEmail creates a new mock instance.
func NewMockEmail(ctrl *gomock.Controller) *MockEmail {
	mock := &MockEmail{ctrl: ctrl}
	mock.recorder = &MockEmailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmail) EXPECT() *MockEmailMockRecorder {
	return m.recorder
}

// SendPasswordReset mocks base method.
func (m *MockEmail) SendPasswordReset(ctx context.Context, req SendPasswordReset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPasswordReset", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPasswordReset indicates an expected call of SendPasswordReset.
func (mr *MockEmailMockRecorder) SendPasswordReset(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPasswordReset", reflect.TypeOf((*MockEmail)(nil).SendPasswordReset), ctx, req)
}

// SendVerification mocks base method.
func (m *MockEmail) SendVerification(ctx context.Context, req SendVerification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerification", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerification indicates an expected call of SendVerification.
func (mr *MockEmailMockRecorder) SendVerification(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerification", reflect.TypeOf((*MockEmail)(nil).SendVerification), ctx, req)
}
