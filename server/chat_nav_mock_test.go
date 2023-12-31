// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mk6i/retro-aim-server/oscar"
	mock "github.com/stretchr/testify/mock"

	state "github.com/mk6i/retro-aim-server/state"
)

// mockChatNavHandler is an autogenerated mock type for the ChatNavHandler type
type mockChatNavHandler struct {
	mock.Mock
}

type mockChatNavHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockChatNavHandler) EXPECT() *mockChatNavHandler_Expecter {
	return &mockChatNavHandler_Expecter{mock: &_m.Mock}
}

// CreateRoomHandler provides a mock function with given fields: ctx, sess, inFrame, inBody
func (_m *mockChatNavHandler) CreateRoomHandler(ctx context.Context, sess *state.Session, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (oscar.SNACMessage, error) {
	ret := _m.Called(ctx, sess, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for CreateRoomHandler")
	}

	var r0 oscar.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (oscar.SNACMessage, error)); ok {
		return rf(ctx, sess, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) oscar.SNACMessage); ok {
		r0 = rf(ctx, sess, inFrame, inBody)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) error); ok {
		r1 = rf(ctx, sess, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockChatNavHandler_CreateRoomHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRoomHandler'
type mockChatNavHandler_CreateRoomHandler_Call struct {
	*mock.Call
}

// CreateRoomHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame oscar.SNACFrame
//   - inBody oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate
func (_e *mockChatNavHandler_Expecter) CreateRoomHandler(ctx interface{}, sess interface{}, inFrame interface{}, inBody interface{}) *mockChatNavHandler_CreateRoomHandler_Call {
	return &mockChatNavHandler_CreateRoomHandler_Call{Call: _e.mock.On("CreateRoomHandler", ctx, sess, inFrame, inBody)}
}

func (_c *mockChatNavHandler_CreateRoomHandler_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate)) *mockChatNavHandler_CreateRoomHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(oscar.SNACFrame), args[3].(oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate))
	})
	return _c
}

func (_c *mockChatNavHandler_CreateRoomHandler_Call) Return(_a0 oscar.SNACMessage, _a1 error) *mockChatNavHandler_CreateRoomHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockChatNavHandler_CreateRoomHandler_Call) RunAndReturn(run func(context.Context, *state.Session, oscar.SNACFrame, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (oscar.SNACMessage, error)) *mockChatNavHandler_CreateRoomHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RequestChatRightsHandler provides a mock function with given fields: ctx, inFrame
func (_m *mockChatNavHandler) RequestChatRightsHandler(ctx context.Context, inFrame oscar.SNACFrame) oscar.SNACMessage {
	ret := _m.Called(ctx, inFrame)

	if len(ret) == 0 {
		panic("no return value specified for RequestChatRightsHandler")
	}

	var r0 oscar.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame) oscar.SNACMessage); ok {
		r0 = rf(ctx, inFrame)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	return r0
}

// mockChatNavHandler_RequestChatRightsHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestChatRightsHandler'
type mockChatNavHandler_RequestChatRightsHandler_Call struct {
	*mock.Call
}

// RequestChatRightsHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame oscar.SNACFrame
func (_e *mockChatNavHandler_Expecter) RequestChatRightsHandler(ctx interface{}, inFrame interface{}) *mockChatNavHandler_RequestChatRightsHandler_Call {
	return &mockChatNavHandler_RequestChatRightsHandler_Call{Call: _e.mock.On("RequestChatRightsHandler", ctx, inFrame)}
}

func (_c *mockChatNavHandler_RequestChatRightsHandler_Call) Run(run func(ctx context.Context, inFrame oscar.SNACFrame)) *mockChatNavHandler_RequestChatRightsHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNACFrame))
	})
	return _c
}

func (_c *mockChatNavHandler_RequestChatRightsHandler_Call) Return(_a0 oscar.SNACMessage) *mockChatNavHandler_RequestChatRightsHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockChatNavHandler_RequestChatRightsHandler_Call) RunAndReturn(run func(context.Context, oscar.SNACFrame) oscar.SNACMessage) *mockChatNavHandler_RequestChatRightsHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RequestRoomInfoHandler provides a mock function with given fields: ctx, inFrame, inBody
func (_m *mockChatNavHandler) RequestRoomInfoHandler(ctx context.Context, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (oscar.SNACMessage, error) {
	ret := _m.Called(ctx, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for RequestRoomInfoHandler")
	}

	var r0 oscar.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame, oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (oscar.SNACMessage, error)); ok {
		return rf(ctx, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, oscar.SNACFrame, oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) oscar.SNACMessage); ok {
		r0 = rf(ctx, inFrame, inBody)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, oscar.SNACFrame, oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) error); ok {
		r1 = rf(ctx, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockChatNavHandler_RequestRoomInfoHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequestRoomInfoHandler'
type mockChatNavHandler_RequestRoomInfoHandler_Call struct {
	*mock.Call
}

// RequestRoomInfoHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame oscar.SNACFrame
//   - inBody oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo
func (_e *mockChatNavHandler_Expecter) RequestRoomInfoHandler(ctx interface{}, inFrame interface{}, inBody interface{}) *mockChatNavHandler_RequestRoomInfoHandler_Call {
	return &mockChatNavHandler_RequestRoomInfoHandler_Call{Call: _e.mock.On("RequestRoomInfoHandler", ctx, inFrame, inBody)}
}

func (_c *mockChatNavHandler_RequestRoomInfoHandler_Call) Run(run func(ctx context.Context, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo)) *mockChatNavHandler_RequestRoomInfoHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(oscar.SNACFrame), args[2].(oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo))
	})
	return _c
}

func (_c *mockChatNavHandler_RequestRoomInfoHandler_Call) Return(_a0 oscar.SNACMessage, _a1 error) *mockChatNavHandler_RequestRoomInfoHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockChatNavHandler_RequestRoomInfoHandler_Call) RunAndReturn(run func(context.Context, oscar.SNACFrame, oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (oscar.SNACMessage, error)) *mockChatNavHandler_RequestRoomInfoHandler_Call {
	_c.Call.Return(run)
	return _c
}

// newMockChatNavHandler creates a new instance of mockChatNavHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockChatNavHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockChatNavHandler {
	mock := &mockChatNavHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
