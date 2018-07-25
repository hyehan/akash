// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"
import types "github.com/ovrclk/akash/types"

// Reservation is an autogenerated mock type for the Reservation type
type Reservation struct {
	mock.Mock
}

// OrderID provides a mock function with given fields:
func (_m *Reservation) OrderID() types.OrderID {
	ret := _m.Called()

	var r0 types.OrderID
	if rf, ok := ret.Get(0).(func() types.OrderID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.OrderID)
	}

	return r0
}

// Resources provides a mock function with given fields:
func (_m *Reservation) Resources() types.ResourceList {
	ret := _m.Called()

	var r0 types.ResourceList
	if rf, ok := ret.Get(0).(func() types.ResourceList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.ResourceList)
		}
	}

	return r0
}
