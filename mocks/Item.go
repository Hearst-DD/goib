package mocks

import "github.com/Hearst-DD/goib"
import "github.com/stretchr/testify/mock"

type Item struct {
	mock.Mock
}

// GetType provides a mock function with given fields:
func (_m *Item) GetType() goib.ItemType {
	ret := _m.Called()

	var r0 goib.ItemType
	if rf, ok := ret.Get(0).(func() goib.ItemType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(goib.ItemType)
	}

	return r0
}

// GetContentID provides a mock function with given fields:
func (_m *Item) GetContentID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetTeaserTitle provides a mock function with given fields:
func (_m *Item) GetTeaserTitle() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTeaserText provides a mock function with given fields:
func (_m *Item) GetTeaserText() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetPublicationDate provides a mock function with given fields:
func (_m *Item) GetPublicationDate() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}
