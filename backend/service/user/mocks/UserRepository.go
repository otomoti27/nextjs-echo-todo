// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	domain "echo-api/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *UserRepository) Create(_a0 *domain.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *UserRepository) Delete(_a0 *domain.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: _a0, email
func (_m *UserRepository) GetByEmail(_a0 *domain.User, email string) error {
	ret := _m.Called(_a0, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User, string) error); ok {
		r0 = rf(_a0, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0, id
func (_m *UserRepository) GetByID(_a0 *domain.User, id uint) error {
	ret := _m.Called(_a0, id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User, uint) error); ok {
		r0 = rf(_a0, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *UserRepository) Update(_a0 *domain.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
