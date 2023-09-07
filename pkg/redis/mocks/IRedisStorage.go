// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IRedisStorage is an autogenerated mock type for the IRedisStorage type
type IRedisStorage struct {
	mock.Mock
}

// Del provides a mock function with given fields: key
func (_m *IRedisStorage) Del(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByKey provides a mock function with given fields: key, convertValue
func (_m *IRedisStorage) GetByKey(key string, convertValue interface{}) error {
	ret := _m.Called(key, convertValue)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, convertValue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: key, dataset
func (_m *IRedisStorage) Set(key string, dataset interface{}) error {
	ret := _m.Called(key, dataset)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, dataset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIRedisStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRedisStorage creates a new instance of IRedisStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRedisStorage(t mockConstructorTestingTNewIRedisStorage) *IRedisStorage {
	mock := &IRedisStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}