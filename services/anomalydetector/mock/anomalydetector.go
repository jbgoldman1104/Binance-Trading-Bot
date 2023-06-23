// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	decimal "github.com/shopspring/decimal"
	mock "github.com/stretchr/testify/mock"
)

// AnomalyDetector is an autogenerated mock type for the AnomalyDetector type
type AnomalyDetector struct {
	mock.Mock
}

// IsAnomaly provides a mock function with given fields: price
func (_m *AnomalyDetector) IsAnomaly(price decimal.Decimal) bool {
	ret := _m.Called(price)

	var r0 bool
	if rf, ok := ret.Get(0).(func(decimal.Decimal) bool); ok {
		r0 = rf(price)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewAnomalyDetector interface {
	mock.TestingT
	Cleanup(func())
}

// NewAnomalyDetector creates a new instance of AnomalyDetector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAnomalyDetector(t mockConstructorTestingTNewAnomalyDetector) *AnomalyDetector {
	mock := &AnomalyDetector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
