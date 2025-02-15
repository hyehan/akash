// Code generated by mockery v2.5.1. DO NOT EDIT.

package kubernetes_mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta1"
)

// FlowSchemasGetter is an autogenerated mock type for the FlowSchemasGetter type
type FlowSchemasGetter struct {
	mock.Mock
}

// FlowSchemas provides a mock function with given fields:
func (_m *FlowSchemasGetter) FlowSchemas() v1beta1.FlowSchemaInterface {
	ret := _m.Called()

	var r0 v1beta1.FlowSchemaInterface
	if rf, ok := ret.Get(0).(func() v1beta1.FlowSchemaInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1beta1.FlowSchemaInterface)
		}
	}

	return r0
}
