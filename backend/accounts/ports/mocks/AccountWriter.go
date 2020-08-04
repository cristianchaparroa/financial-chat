// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	accounts "chat/accounts"

	mock "github.com/stretchr/testify/mock"
)

// AccountWriter is an autogenerated mock type for the AccountWriter type
type AccountWriter struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *AccountWriter) Create(account *accounts.Account) (*accounts.Account, error) {
	ret := _m.Called(account)

	var r0 *accounts.Account
	if rf, ok := ret.Get(0).(func(*accounts.Account) *accounts.Account); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*accounts.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*accounts.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}