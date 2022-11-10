// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CLIRegistrationSupporter is an autogenerated mock type for the CLIRegistrationSupporter type
type CLIRegistrationSupporter struct {
	mock.Mock
}

type CLIRegistrationSupporter_Expecter struct {
	mock *mock.Mock
}

func (_m *CLIRegistrationSupporter) EXPECT() *CLIRegistrationSupporter_Expecter {
	return &CLIRegistrationSupporter_Expecter{mock: &_m.Mock}
}

// Registration provides a mock function with given fields: ctx, data
func (_m *CLIRegistrationSupporter) Registration(ctx context.Context, data interface{}) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CLIRegistrationSupporter_Registration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Registration'
type CLIRegistrationSupporter_Registration_Call struct {
	*mock.Call
}

// Registration is a helper method to define mock.On call
//  - ctx context.Context
//  - data interface{}
func (_e *CLIRegistrationSupporter_Expecter) Registration(ctx interface{}, data interface{}) *CLIRegistrationSupporter_Registration_Call {
	return &CLIRegistrationSupporter_Registration_Call{Call: _e.mock.On("Registration", ctx, data)}
}

func (_c *CLIRegistrationSupporter_Registration_Call) Run(run func(ctx context.Context, data interface{})) *CLIRegistrationSupporter_Registration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *CLIRegistrationSupporter_Registration_Call) Return(_a0 error) *CLIRegistrationSupporter_Registration_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewCLIRegistrationSupporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCLIRegistrationSupporter creates a new instance of CLIRegistrationSupporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCLIRegistrationSupporter(t mockConstructorTestingTNewCLIRegistrationSupporter) *CLIRegistrationSupporter {
	mock := &CLIRegistrationSupporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
