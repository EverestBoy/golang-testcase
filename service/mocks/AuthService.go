// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	model "cf-service/model"

	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// UserEmailLoginService provides a mock function with given fields: credential
func (_m *AuthService) UserEmailLoginService(credential model.Credential) (*model.UserView, error) {
	ret := _m.Called(credential)

	var r0 *model.UserView
	if rf, ok := ret.Get(0).(func(model.Credential) *model.UserView); ok {
		r0 = rf(credential)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserView)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Credential) error); ok {
		r1 = rf(credential)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewAuthServiceT interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthService creates a new instance of AuthService. It also registers a testing interface on the mock and a cleanup function to assert the dbMocks expectations.
func NewAuthService(t NewAuthServiceT) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
