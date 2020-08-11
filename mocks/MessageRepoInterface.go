// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	models "warpin/models"

	mock "github.com/stretchr/testify/mock"
)

// MessageRepoInterface is an autogenerated mock type for the MessageRepoInterface type
type MessageRepoInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MessageRepoInterface) Create(_a0 models.Message) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Message) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *MessageRepoInterface) FindAll() []models.Message {
	ret := _m.Called()

	var r0 []models.Message
	if rf, ok := ret.Get(0).(func() []models.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Message)
		}
	}

	return r0
}

// FindOne provides a mock function with given fields: _a0
func (_m *MessageRepoInterface) FindOne(_a0 string) models.Message {
	ret := _m.Called(_a0)

	var r0 models.Message
	if rf, ok := ret.Get(0).(func(string) models.Message); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Message)
	}

	return r0
}
