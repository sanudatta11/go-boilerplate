// Code generated by MockGen. DO NOT EDIT.
// Source: ./LoginService.go
//
// Generated by this command:
//
//	mockgen -source=./LoginService.go -destination=../mocks/LoginService_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	models "boilerplate/models"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockLoginService is a mock of LoginService interface.
type MockLoginService struct {
	ctrl     *gomock.Controller
	recorder *MockLoginServiceMockRecorder
}

// MockLoginServiceMockRecorder is the mock recorder for MockLoginService.
type MockLoginServiceMockRecorder struct {
	mock *MockLoginService
}

// NewMockLoginService creates a new mock instance.
func NewMockLoginService(ctrl *gomock.Controller) *MockLoginService {
	mock := &MockLoginService{ctrl: ctrl}
	mock.recorder = &MockLoginServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginService) EXPECT() *MockLoginServiceMockRecorder {
	return m.recorder
}

// GetLoggedInUser mocks base method.
func (m *MockLoginService) GetLoggedInUser(ctx *gin.Context, userEmail, userPass string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoggedInUser", ctx, userEmail, userPass)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoggedInUser indicates an expected call of GetLoggedInUser.
func (mr *MockLoginServiceMockRecorder) GetLoggedInUser(ctx, userEmail, userPass any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoggedInUser", reflect.TypeOf((*MockLoginService)(nil).GetLoggedInUser), ctx, userEmail, userPass)
}
