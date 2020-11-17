// Code generated by mockery v1.1.2. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	rest "k8s.io/client-go/rest"

	v1beta1 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
)

// DiscoveryV1beta1Interface is an autogenerated mock type for the DiscoveryV1beta1Interface type
type DiscoveryV1beta1Interface struct {
	mock.Mock
}

// EndpointSlices provides a mock function with given fields: namespace
func (_m *DiscoveryV1beta1Interface) EndpointSlices(namespace string) v1beta1.EndpointSliceInterface {
	ret := _m.Called(namespace)

	var r0 v1beta1.EndpointSliceInterface
	if rf, ok := ret.Get(0).(func(string) v1beta1.EndpointSliceInterface); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.EndpointSliceInterface)
		}
	}

	return r0
}

// RESTClient provides a mock function with given fields:
func (_m *DiscoveryV1beta1Interface) RESTClient() rest.Interface {
	ret := _m.Called()

	var r0 rest.Interface
	if rf, ok := ret.Get(0).(func() rest.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rest.Interface)
		}
	}

	return r0
}
