// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CLISelectCommandSupporter is an autogenerated mock type for the CLISelectCommandSupporter type
type CLISelectCommandSupporter struct {
	mock.Mock
}

type CLISelectCommandSupporter_Expecter struct {
	mock *mock.Mock
}

func (_m *CLISelectCommandSupporter) EXPECT() *CLISelectCommandSupporter_Expecter {
	return &CLISelectCommandSupporter_Expecter{mock: &_m.Mock}
}

// SelectCommand provides a mock function with given fields: ctx, options
func (_m *CLISelectCommandSupporter) SelectCommand(ctx context.Context, options []string) (string, error) {
	ret := _m.Called(ctx, options)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, []string) string); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CLISelectCommandSupporter_SelectCommand_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectCommand'
type CLISelectCommandSupporter_SelectCommand_Call struct {
	*mock.Call
}

// SelectCommand is a helper method to define mock.On call
//  - ctx context.Context
//  - options []string
func (_e *CLISelectCommandSupporter_Expecter) SelectCommand(ctx interface{}, options interface{}) *CLISelectCommandSupporter_SelectCommand_Call {
	return &CLISelectCommandSupporter_SelectCommand_Call{Call: _e.mock.On("SelectCommand", ctx, options)}
}

func (_c *CLISelectCommandSupporter_SelectCommand_Call) Run(run func(ctx context.Context, options []string)) *CLISelectCommandSupporter_SelectCommand_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *CLISelectCommandSupporter_SelectCommand_Call) Return(_a0 string, _a1 error) *CLISelectCommandSupporter_SelectCommand_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewCLISelectCommandSupporter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCLISelectCommandSupporter creates a new instance of CLISelectCommandSupporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCLISelectCommandSupporter(t mockConstructorTestingTNewCLISelectCommandSupporter) *CLISelectCommandSupporter {
	mock := &CLISelectCommandSupporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
