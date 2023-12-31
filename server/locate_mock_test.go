// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mk6i/retro-aim-server/oscar"
	mock "github.com/stretchr/testify/mock"

	state "github.com/mk6i/retro-aim-server/state"
)

// mockLocateHandler is an autogenerated mock type for the LocateHandler type
type mockLocateHandler struct {
	mock.Mock
}

type mockLocateHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockLocateHandler) EXPECT() *mockLocateHandler_Expecter {
	return &mockLocateHandler_Expecter{mock: &_m.Mock}
}

// RightsQueryHandler provides a mock function with given fields: ctx, inFrame
func (_m *mockLocateHandler) RightsQueryHandler(ctx context.Context, inFrame oscar.SNACFrame) oscar.SNACMessage {
	ret := _m.Called(ctx, inFrame)

	if len(ret) == 0 {
		panic("no return value specified for RightsQueryHandler")
	}

	var r0 oscar.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame) oscar.SNACMessage); ok {
		r0 = rf(ctx, inFrame)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	return r0
}

// mockLocateHandler_RightsQueryHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RightsQueryHandler'
type mockLocateHandler_RightsQueryHandler_Call struct {
	*mock.Call
}

// RightsQueryHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame oscar.SNACFrame
func (_e *mockLocateHandler_Expecter) RightsQueryHandler(ctx interface{}, inFrame interface{}) *mockLocateHandler_RightsQueryHandler_Call {
	return &mockLocateHandler_RightsQueryHandler_Call{Call: _e.mock.On("RightsQueryHandler", ctx, inFrame)}
}

func (_c *mockLocateHandler_RightsQueryHandler_Call) Run(run func(ctx context.Context, inFrame oscar.SNACFrame)) *mockLocateHandler_RightsQueryHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNACFrame))
	})
	return _c
}

func (_c *mockLocateHandler_RightsQueryHandler_Call) Return(_a0 oscar.SNACMessage) *mockLocateHandler_RightsQueryHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocateHandler_RightsQueryHandler_Call) RunAndReturn(run func(context.Context, oscar.SNACFrame) oscar.SNACMessage) *mockLocateHandler_RightsQueryHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SetDirInfoHandler provides a mock function with given fields: ctx, frame
func (_m *mockLocateHandler) SetDirInfoHandler(ctx context.Context, frame oscar.SNACFrame) oscar.SNACMessage {
	ret := _m.Called(ctx, frame)

	if len(ret) == 0 {
		panic("no return value specified for SetDirInfoHandler")
	}

	var r0 oscar.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame) oscar.SNACMessage); ok {
		r0 = rf(ctx, frame)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	return r0
}

// mockLocateHandler_SetDirInfoHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDirInfoHandler'
type mockLocateHandler_SetDirInfoHandler_Call struct {
	*mock.Call
}

// SetDirInfoHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - frame oscar.SNACFrame
func (_e *mockLocateHandler_Expecter) SetDirInfoHandler(ctx interface{}, frame interface{}) *mockLocateHandler_SetDirInfoHandler_Call {
	return &mockLocateHandler_SetDirInfoHandler_Call{Call: _e.mock.On("SetDirInfoHandler", ctx, frame)}
}

func (_c *mockLocateHandler_SetDirInfoHandler_Call) Run(run func(ctx context.Context, frame oscar.SNACFrame)) *mockLocateHandler_SetDirInfoHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNACFrame))
	})
	return _c
}

func (_c *mockLocateHandler_SetDirInfoHandler_Call) Return(_a0 oscar.SNACMessage) *mockLocateHandler_SetDirInfoHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocateHandler_SetDirInfoHandler_Call) RunAndReturn(run func(context.Context, oscar.SNACFrame) oscar.SNACMessage) *mockLocateHandler_SetDirInfoHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SetInfoHandler provides a mock function with given fields: ctx, sess, inBody
func (_m *mockLocateHandler) SetInfoHandler(ctx context.Context, sess *state.Session, inBody oscar.SNAC_0x02_0x04_LocateSetInfo) error {
	ret := _m.Called(ctx, sess, inBody)

	if len(ret) == 0 {
		panic("no return value specified for SetInfoHandler")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNAC_0x02_0x04_LocateSetInfo) error); ok {
		r0 = rf(ctx, sess, inBody)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocateHandler_SetInfoHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetInfoHandler'
type mockLocateHandler_SetInfoHandler_Call struct {
	*mock.Call
}

// SetInfoHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inBody oscar.SNAC_0x02_0x04_LocateSetInfo
func (_e *mockLocateHandler_Expecter) SetInfoHandler(ctx interface{}, sess interface{}, inBody interface{}) *mockLocateHandler_SetInfoHandler_Call {
	return &mockLocateHandler_SetInfoHandler_Call{Call: _e.mock.On("SetInfoHandler", ctx, sess, inBody)}
}

func (_c *mockLocateHandler_SetInfoHandler_Call) Run(run func(ctx context.Context, sess *state.Session, inBody oscar.SNAC_0x02_0x04_LocateSetInfo)) *mockLocateHandler_SetInfoHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(oscar.SNAC_0x02_0x04_LocateSetInfo))
	})
	return _c
}

func (_c *mockLocateHandler_SetInfoHandler_Call) Return(_a0 error) *mockLocateHandler_SetInfoHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocateHandler_SetInfoHandler_Call) RunAndReturn(run func(context.Context, *state.Session, oscar.SNAC_0x02_0x04_LocateSetInfo) error) *mockLocateHandler_SetInfoHandler_Call {
	_c.Call.Return(run)
	return _c
}

// SetKeywordInfoHandler provides a mock function with given fields: ctx, inFrame
func (_m *mockLocateHandler) SetKeywordInfoHandler(ctx context.Context, inFrame oscar.SNACFrame) oscar.SNACMessage {
	ret := _m.Called(ctx, inFrame)

	if len(ret) == 0 {
		panic("no return value specified for SetKeywordInfoHandler")
	}

	var r0 oscar.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame) oscar.SNACMessage); ok {
		r0 = rf(ctx, inFrame)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	return r0
}

// mockLocateHandler_SetKeywordInfoHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetKeywordInfoHandler'
type mockLocateHandler_SetKeywordInfoHandler_Call struct {
	*mock.Call
}

// SetKeywordInfoHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame oscar.SNACFrame
func (_e *mockLocateHandler_Expecter) SetKeywordInfoHandler(ctx interface{}, inFrame interface{}) *mockLocateHandler_SetKeywordInfoHandler_Call {
	return &mockLocateHandler_SetKeywordInfoHandler_Call{Call: _e.mock.On("SetKeywordInfoHandler", ctx, inFrame)}
}

func (_c *mockLocateHandler_SetKeywordInfoHandler_Call) Run(run func(ctx context.Context, inFrame oscar.SNACFrame)) *mockLocateHandler_SetKeywordInfoHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNACFrame))
	})
	return _c
}

func (_c *mockLocateHandler_SetKeywordInfoHandler_Call) Return(_a0 oscar.SNACMessage) *mockLocateHandler_SetKeywordInfoHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocateHandler_SetKeywordInfoHandler_Call) RunAndReturn(run func(context.Context, oscar.SNACFrame) oscar.SNACMessage) *mockLocateHandler_SetKeywordInfoHandler_Call {
	_c.Call.Return(run)
	return _c
}

// UserInfoQuery2Handler provides a mock function with given fields: ctx, sess, inFrame, inBody
func (_m *mockLocateHandler) UserInfoQuery2Handler(ctx context.Context, sess *state.Session, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x02_0x15_LocateUserInfoQuery2) (oscar.SNACMessage, error) {
	ret := _m.Called(ctx, sess, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for UserInfoQuery2Handler")
	}

	var r0 oscar.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x02_0x15_LocateUserInfoQuery2) (oscar.SNACMessage, error)); ok {
		return rf(ctx, sess, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x02_0x15_LocateUserInfoQuery2) oscar.SNACMessage); ok {
		r0 = rf(ctx, sess, inFrame, inBody)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x02_0x15_LocateUserInfoQuery2) error); ok {
		r1 = rf(ctx, sess, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockLocateHandler_UserInfoQuery2Handler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfoQuery2Handler'
type mockLocateHandler_UserInfoQuery2Handler_Call struct {
	*mock.Call
}

// UserInfoQuery2Handler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame oscar.SNACFrame
//   - inBody oscar.SNAC_0x02_0x15_LocateUserInfoQuery2
func (_e *mockLocateHandler_Expecter) UserInfoQuery2Handler(ctx interface{}, sess interface{}, inFrame interface{}, inBody interface{}) *mockLocateHandler_UserInfoQuery2Handler_Call {
	return &mockLocateHandler_UserInfoQuery2Handler_Call{Call: _e.mock.On("UserInfoQuery2Handler", ctx, sess, inFrame, inBody)}
}

func (_c *mockLocateHandler_UserInfoQuery2Handler_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x02_0x15_LocateUserInfoQuery2)) *mockLocateHandler_UserInfoQuery2Handler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(oscar.SNACFrame), args[3].(oscar.SNAC_0x02_0x15_LocateUserInfoQuery2))
	})
	return _c
}

func (_c *mockLocateHandler_UserInfoQuery2Handler_Call) Return(_a0 oscar.SNACMessage, _a1 error) *mockLocateHandler_UserInfoQuery2Handler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockLocateHandler_UserInfoQuery2Handler_Call) RunAndReturn(run func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x02_0x15_LocateUserInfoQuery2) (oscar.SNACMessage, error)) *mockLocateHandler_UserInfoQuery2Handler_Call {
	_c.Call.Return(run)
	return _c
}

// newMockLocateHandler creates a new instance of mockLocateHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockLocateHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockLocateHandler {
	mock := &mockLocateHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
