package mocks

import "github.com/Hearst-DD/goib"
import "github.com/stretchr/testify/mock"

import "net/url"

type API struct {
	mock.Mock
}

// Entry provides a mock function with given fields: channel, entrytype, params
func (_m *API) Entry(channel string, entrytype string, params url.Values) (goib.Item, error) {
	ret := _m.Called(channel, entrytype, params)

	var r0 goib.Item
	if rf, ok := ret.Get(0).(func(string, string, url.Values) goib.Item); ok {
		r0 = rf(channel, entrytype, params)
	} else {
		r0 = ret.Get(0).(goib.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, url.Values) error); ok {
		r1 = rf(channel, entrytype, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: channel, query, params
func (_m *API) Search(channel string, query string, params url.Values) (*goib.Collection, error) {
	ret := _m.Called(channel, query, params)

	var r0 *goib.Collection
	if rf, ok := ret.Get(0).(func(string, string, url.Values) *goib.Collection); ok {
		r0 = rf(channel, query, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*goib.Collection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, url.Values) error); ok {
		r1 = rf(channel, query, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Content provides a mock function with given fields: channel, contentID, params
func (_m *API) Content(channel string, contentID int, params url.Values) (goib.Item, error) {
	ret := _m.Called(channel, contentID, params)

	var r0 goib.Item
	if rf, ok := ret.Get(0).(func(string, int, url.Values) goib.Item); ok {
		r0 = rf(channel, contentID, params)
	} else {
		r0 = ret.Get(0).(goib.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, url.Values) error); ok {
		r1 = rf(channel, contentID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContentMedia provides a mock function with given fields: channel, contentID, params
func (_m *API) ContentMedia(channel string, contentID int, params url.Values) ([]goib.Item, error) {
	ret := _m.Called(channel, contentID, params)

	var r0 []goib.Item
	if rf, ok := ret.Get(0).(func(string, int, url.Values) []goib.Item); ok {
		r0 = rf(channel, contentID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]goib.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, url.Values) error); ok {
		r1 = rf(channel, contentID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContentItems provides a mock function with given fields: channel, contentID, params
func (_m *API) ContentItems(channel string, contentID int, params url.Values) ([]goib.Item, error) {
	ret := _m.Called(channel, contentID, params)

	var r0 []goib.Item
	if rf, ok := ret.Get(0).(func(string, int, url.Values) []goib.Item); ok {
		r0 = rf(channel, contentID, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]goib.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, url.Values) error); ok {
		r1 = rf(channel, contentID, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Closings provides a mock function with given fields: channel, filter, providerID
func (_m *API) Closings(channel string, filter goib.ClosingsFilter, providerID ...string) (goib.ClosingsResponse, error) {
	ret := _m.Called(channel, filter, providerID)

	var r0 goib.ClosingsResponse
	if rf, ok := ret.Get(0).(func(string, goib.ClosingsFilter, ...string) goib.ClosingsResponse); ok {
		r0 = rf(channel, filter, providerID...)
	} else {
		r0 = ret.Get(0).(goib.ClosingsResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, goib.ClosingsFilter, ...string) error); ok {
		r1 = rf(channel, filter, providerID...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnmarshalReceiver provides a mock function with given fields: r
func (_m *API) UnmarshalReceiver(r goib.Receiver) (goib.Item, error) {
	ret := _m.Called(r)

	var r0 goib.Item
	if rf, ok := ret.Get(0).(func(goib.Receiver) goib.Item); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(goib.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(goib.Receiver) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
