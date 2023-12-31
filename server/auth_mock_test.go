// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mk6i/retro-aim-server/oscar"
	mock "github.com/stretchr/testify/mock"

	state "github.com/mk6i/retro-aim-server/state"

	uuid "github.com/google/uuid"
)

// mockAuthHandler is an autogenerated mock type for the AuthHandler type
type mockAuthHandler struct {
	mock.Mock
}

type mockAuthHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockAuthHandler) EXPECT() *mockAuthHandler_Expecter {
	return &mockAuthHandler_Expecter{mock: &_m.Mock}
}

// BUCPChallengeRequestHandler provides a mock function with given fields: bodyIn, newUUID
func (_m *mockAuthHandler) BUCPChallengeRequestHandler(bodyIn oscar.SNAC_0x17_0x06_BUCPChallengeRequest, newUUID func() uuid.UUID) (oscar.SNACMessage, error) {
	ret := _m.Called(bodyIn, newUUID)

	if len(ret) == 0 {
		panic("no return value specified for BUCPChallengeRequestHandler")
	}

	var r0 oscar.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(oscar.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) (oscar.SNACMessage, error)); ok {
		return rf(bodyIn, newUUID)
	}
	if rf, ok := ret.Get(0).(func(oscar.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) oscar.SNACMessage); ok {
		r0 = rf(bodyIn, newUUID)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(oscar.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) error); ok {
		r1 = rf(bodyIn, newUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthHandler_BUCPChallengeRequestHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BUCPChallengeRequestHandler'
type mockAuthHandler_BUCPChallengeRequestHandler_Call struct {
	*mock.Call
}

// BUCPChallengeRequestHandler is a helper method to define mock.On call
//   - bodyIn oscar.SNAC_0x17_0x06_BUCPChallengeRequest
//   - newUUID func() uuid.UUID
func (_e *mockAuthHandler_Expecter) BUCPChallengeRequestHandler(bodyIn interface{}, newUUID interface{}) *mockAuthHandler_BUCPChallengeRequestHandler_Call {
	return &mockAuthHandler_BUCPChallengeRequestHandler_Call{Call: _e.mock.On("BUCPChallengeRequestHandler", bodyIn, newUUID)}
}

func (_c *mockAuthHandler_BUCPChallengeRequestHandler_Call) Run(run func(bodyIn oscar.SNAC_0x17_0x06_BUCPChallengeRequest, newUUID func() uuid.UUID)) *mockAuthHandler_BUCPChallengeRequestHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(oscar.SNAC_0x17_0x06_BUCPChallengeRequest), args[1].(func() uuid.UUID))
	})
	return _c
}

func (_c *mockAuthHandler_BUCPChallengeRequestHandler_Call) Return(_a0 oscar.SNACMessage, _a1 error) *mockAuthHandler_BUCPChallengeRequestHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthHandler_BUCPChallengeRequestHandler_Call) RunAndReturn(run func(oscar.SNAC_0x17_0x06_BUCPChallengeRequest, func() uuid.UUID) (oscar.SNACMessage, error)) *mockAuthHandler_BUCPChallengeRequestHandler_Call {
	_c.Call.Return(run)
	return _c
}

// BUCPLoginRequestHandler provides a mock function with given fields: bodyIn, newUUID, fn
func (_m *mockAuthHandler) BUCPLoginRequestHandler(bodyIn oscar.SNAC_0x17_0x02_BUCPLoginRequest, newUUID func() uuid.UUID, fn func(string) (state.User, error)) (oscar.SNACMessage, error) {
	ret := _m.Called(bodyIn, newUUID, fn)

	if len(ret) == 0 {
		panic("no return value specified for BUCPLoginRequestHandler")
	}

	var r0 oscar.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(oscar.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) (oscar.SNACMessage, error)); ok {
		return rf(bodyIn, newUUID, fn)
	}
	if rf, ok := ret.Get(0).(func(oscar.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) oscar.SNACMessage); ok {
		r0 = rf(bodyIn, newUUID, fn)
	} else {
		r0 = ret.Get(0).(oscar.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(oscar.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) error); ok {
		r1 = rf(bodyIn, newUUID, fn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthHandler_BUCPLoginRequestHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BUCPLoginRequestHandler'
type mockAuthHandler_BUCPLoginRequestHandler_Call struct {
	*mock.Call
}

// BUCPLoginRequestHandler is a helper method to define mock.On call
//   - bodyIn oscar.SNAC_0x17_0x02_BUCPLoginRequest
//   - newUUID func() uuid.UUID
//   - fn func(string)(state.User , error)
func (_e *mockAuthHandler_Expecter) BUCPLoginRequestHandler(bodyIn interface{}, newUUID interface{}, fn interface{}) *mockAuthHandler_BUCPLoginRequestHandler_Call {
	return &mockAuthHandler_BUCPLoginRequestHandler_Call{Call: _e.mock.On("BUCPLoginRequestHandler", bodyIn, newUUID, fn)}
}

func (_c *mockAuthHandler_BUCPLoginRequestHandler_Call) Run(run func(bodyIn oscar.SNAC_0x17_0x02_BUCPLoginRequest, newUUID func() uuid.UUID, fn func(string) (state.User, error))) *mockAuthHandler_BUCPLoginRequestHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(oscar.SNAC_0x17_0x02_BUCPLoginRequest), args[1].(func() uuid.UUID), args[2].(func(string) (state.User, error)))
	})
	return _c
}

func (_c *mockAuthHandler_BUCPLoginRequestHandler_Call) Return(_a0 oscar.SNACMessage, _a1 error) *mockAuthHandler_BUCPLoginRequestHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthHandler_BUCPLoginRequestHandler_Call) RunAndReturn(run func(oscar.SNAC_0x17_0x02_BUCPLoginRequest, func() uuid.UUID, func(string) (state.User, error)) (oscar.SNACMessage, error)) *mockAuthHandler_BUCPLoginRequestHandler_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveBOSSession provides a mock function with given fields: sessionID
func (_m *mockAuthHandler) RetrieveBOSSession(sessionID string) (*state.Session, error) {
	ret := _m.Called(sessionID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBOSSession")
	}

	var r0 *state.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*state.Session, error)); ok {
		return rf(sessionID)
	}
	if rf, ok := ret.Get(0).(func(string) *state.Session); ok {
		r0 = rf(sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthHandler_RetrieveBOSSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveBOSSession'
type mockAuthHandler_RetrieveBOSSession_Call struct {
	*mock.Call
}

// RetrieveBOSSession is a helper method to define mock.On call
//   - sessionID string
func (_e *mockAuthHandler_Expecter) RetrieveBOSSession(sessionID interface{}) *mockAuthHandler_RetrieveBOSSession_Call {
	return &mockAuthHandler_RetrieveBOSSession_Call{Call: _e.mock.On("RetrieveBOSSession", sessionID)}
}

func (_c *mockAuthHandler_RetrieveBOSSession_Call) Run(run func(sessionID string)) *mockAuthHandler_RetrieveBOSSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockAuthHandler_RetrieveBOSSession_Call) Return(_a0 *state.Session, _a1 error) *mockAuthHandler_RetrieveBOSSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthHandler_RetrieveBOSSession_Call) RunAndReturn(run func(string) (*state.Session, error)) *mockAuthHandler_RetrieveBOSSession_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveChatSession provides a mock function with given fields: chatID, sessionID
func (_m *mockAuthHandler) RetrieveChatSession(chatID string, sessionID string) (*state.Session, error) {
	ret := _m.Called(chatID, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveChatSession")
	}

	var r0 *state.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*state.Session, error)); ok {
		return rf(chatID, sessionID)
	}
	if rf, ok := ret.Get(0).(func(string, string) *state.Session); ok {
		r0 = rf(chatID, sessionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(chatID, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAuthHandler_RetrieveChatSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveChatSession'
type mockAuthHandler_RetrieveChatSession_Call struct {
	*mock.Call
}

// RetrieveChatSession is a helper method to define mock.On call
//   - chatID string
//   - sessionID string
func (_e *mockAuthHandler_Expecter) RetrieveChatSession(chatID interface{}, sessionID interface{}) *mockAuthHandler_RetrieveChatSession_Call {
	return &mockAuthHandler_RetrieveChatSession_Call{Call: _e.mock.On("RetrieveChatSession", chatID, sessionID)}
}

func (_c *mockAuthHandler_RetrieveChatSession_Call) Run(run func(chatID string, sessionID string)) *mockAuthHandler_RetrieveChatSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *mockAuthHandler_RetrieveChatSession_Call) Return(_a0 *state.Session, _a1 error) *mockAuthHandler_RetrieveChatSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAuthHandler_RetrieveChatSession_Call) RunAndReturn(run func(string, string) (*state.Session, error)) *mockAuthHandler_RetrieveChatSession_Call {
	_c.Call.Return(run)
	return _c
}

// Signout provides a mock function with given fields: ctx, sess
func (_m *mockAuthHandler) Signout(ctx context.Context, sess *state.Session) error {
	ret := _m.Called(ctx, sess)

	if len(ret) == 0 {
		panic("no return value specified for Signout")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAuthHandler_Signout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Signout'
type mockAuthHandler_Signout_Call struct {
	*mock.Call
}

// Signout is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
func (_e *mockAuthHandler_Expecter) Signout(ctx interface{}, sess interface{}) *mockAuthHandler_Signout_Call {
	return &mockAuthHandler_Signout_Call{Call: _e.mock.On("Signout", ctx, sess)}
}

func (_c *mockAuthHandler_Signout_Call) Run(run func(ctx context.Context, sess *state.Session)) *mockAuthHandler_Signout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session))
	})
	return _c
}

func (_c *mockAuthHandler_Signout_Call) Return(_a0 error) *mockAuthHandler_Signout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAuthHandler_Signout_Call) RunAndReturn(run func(context.Context, *state.Session) error) *mockAuthHandler_Signout_Call {
	_c.Call.Return(run)
	return _c
}

// SignoutChat provides a mock function with given fields: ctx, sess, chatID
func (_m *mockAuthHandler) SignoutChat(ctx context.Context, sess *state.Session, chatID string) error {
	ret := _m.Called(ctx, sess, chatID)

	if len(ret) == 0 {
		panic("no return value specified for SignoutChat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, string) error); ok {
		r0 = rf(ctx, sess, chatID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAuthHandler_SignoutChat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignoutChat'
type mockAuthHandler_SignoutChat_Call struct {
	*mock.Call
}

// SignoutChat is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - chatID string
func (_e *mockAuthHandler_Expecter) SignoutChat(ctx interface{}, sess interface{}, chatID interface{}) *mockAuthHandler_SignoutChat_Call {
	return &mockAuthHandler_SignoutChat_Call{Call: _e.mock.On("SignoutChat", ctx, sess, chatID)}
}

func (_c *mockAuthHandler_SignoutChat_Call) Run(run func(ctx context.Context, sess *state.Session, chatID string)) *mockAuthHandler_SignoutChat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(string))
	})
	return _c
}

func (_c *mockAuthHandler_SignoutChat_Call) Return(_a0 error) *mockAuthHandler_SignoutChat_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAuthHandler_SignoutChat_Call) RunAndReturn(run func(context.Context, *state.Session, string) error) *mockAuthHandler_SignoutChat_Call {
	_c.Call.Return(run)
	return _c
}

// newMockAuthHandler creates a new instance of mockAuthHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAuthHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAuthHandler {
	mock := &mockAuthHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
