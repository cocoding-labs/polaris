// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	state "github.com/berachain/polaris/eth/core/state"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

type Plugin_Expecter struct {
	mock *mock.Mock
}

func (_m *Plugin) EXPECT() *Plugin_Expecter {
	return &Plugin_Expecter{mock: &_m.Mock}
}

// AddBalance provides a mock function with given fields: _a0, _a1
func (_m *Plugin) AddBalance(_a0 common.Address, _a1 *big.Int) {
	_m.Called(_a0, _a1)
}

// Plugin_AddBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBalance'
type Plugin_AddBalance_Call struct {
	*mock.Call
}

// AddBalance is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 *big.Int
func (_e *Plugin_Expecter) AddBalance(_a0 interface{}, _a1 interface{}) *Plugin_AddBalance_Call {
	return &Plugin_AddBalance_Call{Call: _e.mock.On("AddBalance", _a0, _a1)}
}

func (_c *Plugin_AddBalance_Call) Run(run func(_a0 common.Address, _a1 *big.Int)) *Plugin_AddBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(*big.Int))
	})
	return _c
}

func (_c *Plugin_AddBalance_Call) Return() *Plugin_AddBalance_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_AddBalance_Call) RunAndReturn(run func(common.Address, *big.Int)) *Plugin_AddBalance_Call {
	_c.Call.Return(run)
	return _c
}

// Clone provides a mock function with given fields:
func (_m *Plugin) Clone() state.Plugin {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clone")
	}

	var r0 state.Plugin
	if rf, ok := ret.Get(0).(func() state.Plugin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(state.Plugin)
		}
	}

	return r0
}

// Plugin_Clone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clone'
type Plugin_Clone_Call struct {
	*mock.Call
}

// Clone is a helper method to define mock.On call
func (_e *Plugin_Expecter) Clone() *Plugin_Clone_Call {
	return &Plugin_Clone_Call{Call: _e.mock.On("Clone")}
}

func (_c *Plugin_Clone_Call) Run(run func()) *Plugin_Clone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_Clone_Call) Return(_a0 state.Plugin) *Plugin_Clone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_Clone_Call) RunAndReturn(run func() state.Plugin) *Plugin_Clone_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAccount provides a mock function with given fields: _a0
func (_m *Plugin) CreateAccount(_a0 common.Address) {
	_m.Called(_a0)
}

// Plugin_CreateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccount'
type Plugin_CreateAccount_Call struct {
	*mock.Call
}

// CreateAccount is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) CreateAccount(_a0 interface{}) *Plugin_CreateAccount_Call {
	return &Plugin_CreateAccount_Call{Call: _e.mock.On("CreateAccount", _a0)}
}

func (_c *Plugin_CreateAccount_Call) Run(run func(_a0 common.Address)) *Plugin_CreateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_CreateAccount_Call) Return() *Plugin_CreateAccount_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_CreateAccount_Call) RunAndReturn(run func(common.Address)) *Plugin_CreateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAccounts provides a mock function with given fields: _a0
func (_m *Plugin) DeleteAccounts(_a0 []common.Address) {
	_m.Called(_a0)
}

// Plugin_DeleteAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccounts'
type Plugin_DeleteAccounts_Call struct {
	*mock.Call
}

// DeleteAccounts is a helper method to define mock.On call
//   - _a0 []common.Address
func (_e *Plugin_Expecter) DeleteAccounts(_a0 interface{}) *Plugin_DeleteAccounts_Call {
	return &Plugin_DeleteAccounts_Call{Call: _e.mock.On("DeleteAccounts", _a0)}
}

func (_c *Plugin_DeleteAccounts_Call) Run(run func(_a0 []common.Address)) *Plugin_DeleteAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]common.Address))
	})
	return _c
}

func (_c *Plugin_DeleteAccounts_Call) Return() *Plugin_DeleteAccounts_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_DeleteAccounts_Call) RunAndReturn(run func([]common.Address)) *Plugin_DeleteAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// Empty provides a mock function with given fields: _a0
func (_m *Plugin) Empty(_a0 common.Address) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Empty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Plugin_Empty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Empty'
type Plugin_Empty_Call struct {
	*mock.Call
}

// Empty is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) Empty(_a0 interface{}) *Plugin_Empty_Call {
	return &Plugin_Empty_Call{Call: _e.mock.On("Empty", _a0)}
}

func (_c *Plugin_Empty_Call) Run(run func(_a0 common.Address)) *Plugin_Empty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_Empty_Call) Return(_a0 bool) *Plugin_Empty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_Empty_Call) RunAndReturn(run func(common.Address) bool) *Plugin_Empty_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields:
func (_m *Plugin) Error() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Error")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Plugin_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Plugin_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
func (_e *Plugin_Expecter) Error() *Plugin_Error_Call {
	return &Plugin_Error_Call{Call: _e.mock.On("Error")}
}

func (_c *Plugin_Error_Call) Run(run func()) *Plugin_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_Error_Call) Return(_a0 error) *Plugin_Error_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_Error_Call) RunAndReturn(run func() error) *Plugin_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Exist provides a mock function with given fields: _a0
func (_m *Plugin) Exist(_a0 common.Address) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Exist")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Plugin_Exist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exist'
type Plugin_Exist_Call struct {
	*mock.Call
}

// Exist is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) Exist(_a0 interface{}) *Plugin_Exist_Call {
	return &Plugin_Exist_Call{Call: _e.mock.On("Exist", _a0)}
}

func (_c *Plugin_Exist_Call) Run(run func(_a0 common.Address)) *Plugin_Exist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_Exist_Call) Return(_a0 bool) *Plugin_Exist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_Exist_Call) RunAndReturn(run func(common.Address) bool) *Plugin_Exist_Call {
	_c.Call.Return(run)
	return _c
}

// Finalize provides a mock function with given fields:
func (_m *Plugin) Finalize() {
	_m.Called()
}

// Plugin_Finalize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finalize'
type Plugin_Finalize_Call struct {
	*mock.Call
}

// Finalize is a helper method to define mock.On call
func (_e *Plugin_Expecter) Finalize() *Plugin_Finalize_Call {
	return &Plugin_Finalize_Call{Call: _e.mock.On("Finalize")}
}

func (_c *Plugin_Finalize_Call) Run(run func()) *Plugin_Finalize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_Finalize_Call) Return() *Plugin_Finalize_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_Finalize_Call) RunAndReturn(run func()) *Plugin_Finalize_Call {
	_c.Call.Return(run)
	return _c
}

// ForEachStorage provides a mock function with given fields: _a0, _a1
func (_m *Plugin) ForEachStorage(_a0 common.Address, _a1 func(common.Hash, common.Hash) bool) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ForEachStorage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Address, func(common.Hash, common.Hash) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Plugin_ForEachStorage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ForEachStorage'
type Plugin_ForEachStorage_Call struct {
	*mock.Call
}

// ForEachStorage is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 func(common.Hash , common.Hash) bool
func (_e *Plugin_Expecter) ForEachStorage(_a0 interface{}, _a1 interface{}) *Plugin_ForEachStorage_Call {
	return &Plugin_ForEachStorage_Call{Call: _e.mock.On("ForEachStorage", _a0, _a1)}
}

func (_c *Plugin_ForEachStorage_Call) Run(run func(_a0 common.Address, _a1 func(common.Hash, common.Hash) bool)) *Plugin_ForEachStorage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(func(common.Hash, common.Hash) bool))
	})
	return _c
}

func (_c *Plugin_ForEachStorage_Call) Return(_a0 error) *Plugin_ForEachStorage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_ForEachStorage_Call) RunAndReturn(run func(common.Address, func(common.Hash, common.Hash) bool) error) *Plugin_ForEachStorage_Call {
	_c.Call.Return(run)
	return _c
}

// GetBalance provides a mock function with given fields: _a0
func (_m *Plugin) GetBalance(_a0 common.Address) *big.Int {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// Plugin_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type Plugin_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) GetBalance(_a0 interface{}) *Plugin_GetBalance_Call {
	return &Plugin_GetBalance_Call{Call: _e.mock.On("GetBalance", _a0)}
}

func (_c *Plugin_GetBalance_Call) Run(run func(_a0 common.Address)) *Plugin_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_GetBalance_Call) Return(_a0 *big.Int) *Plugin_GetBalance_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetBalance_Call) RunAndReturn(run func(common.Address) *big.Int) *Plugin_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetCode provides a mock function with given fields: _a0
func (_m *Plugin) GetCode(_a0 common.Address) []byte {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCode")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func(common.Address) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Plugin_GetCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCode'
type Plugin_GetCode_Call struct {
	*mock.Call
}

// GetCode is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) GetCode(_a0 interface{}) *Plugin_GetCode_Call {
	return &Plugin_GetCode_Call{Call: _e.mock.On("GetCode", _a0)}
}

func (_c *Plugin_GetCode_Call) Run(run func(_a0 common.Address)) *Plugin_GetCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_GetCode_Call) Return(_a0 []byte) *Plugin_GetCode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetCode_Call) RunAndReturn(run func(common.Address) []byte) *Plugin_GetCode_Call {
	_c.Call.Return(run)
	return _c
}

// GetCodeHash provides a mock function with given fields: _a0
func (_m *Plugin) GetCodeHash(_a0 common.Address) common.Hash {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetCodeHash")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Address) common.Hash); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// Plugin_GetCodeHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCodeHash'
type Plugin_GetCodeHash_Call struct {
	*mock.Call
}

// GetCodeHash is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) GetCodeHash(_a0 interface{}) *Plugin_GetCodeHash_Call {
	return &Plugin_GetCodeHash_Call{Call: _e.mock.On("GetCodeHash", _a0)}
}

func (_c *Plugin_GetCodeHash_Call) Run(run func(_a0 common.Address)) *Plugin_GetCodeHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_GetCodeHash_Call) Return(_a0 common.Hash) *Plugin_GetCodeHash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetCodeHash_Call) RunAndReturn(run func(common.Address) common.Hash) *Plugin_GetCodeHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetCommittedState provides a mock function with given fields: _a0, _a1
func (_m *Plugin) GetCommittedState(_a0 common.Address, _a1 common.Hash) common.Hash {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetCommittedState")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) common.Hash); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// Plugin_GetCommittedState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCommittedState'
type Plugin_GetCommittedState_Call struct {
	*mock.Call
}

// GetCommittedState is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 common.Hash
func (_e *Plugin_Expecter) GetCommittedState(_a0 interface{}, _a1 interface{}) *Plugin_GetCommittedState_Call {
	return &Plugin_GetCommittedState_Call{Call: _e.mock.On("GetCommittedState", _a0, _a1)}
}

func (_c *Plugin_GetCommittedState_Call) Run(run func(_a0 common.Address, _a1 common.Hash)) *Plugin_GetCommittedState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *Plugin_GetCommittedState_Call) Return(_a0 common.Hash) *Plugin_GetCommittedState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetCommittedState_Call) RunAndReturn(run func(common.Address, common.Hash) common.Hash) *Plugin_GetCommittedState_Call {
	_c.Call.Return(run)
	return _c
}

// GetContext provides a mock function with given fields:
func (_m *Plugin) GetContext() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Plugin_GetContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContext'
type Plugin_GetContext_Call struct {
	*mock.Call
}

// GetContext is a helper method to define mock.On call
func (_e *Plugin_Expecter) GetContext() *Plugin_GetContext_Call {
	return &Plugin_GetContext_Call{Call: _e.mock.On("GetContext")}
}

func (_c *Plugin_GetContext_Call) Run(run func()) *Plugin_GetContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_GetContext_Call) Return(_a0 context.Context) *Plugin_GetContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetContext_Call) RunAndReturn(run func() context.Context) *Plugin_GetContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetNonce provides a mock function with given fields: _a0
func (_m *Plugin) GetNonce(_a0 common.Address) uint64 {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetNonce")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Plugin_GetNonce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNonce'
type Plugin_GetNonce_Call struct {
	*mock.Call
}

// GetNonce is a helper method to define mock.On call
//   - _a0 common.Address
func (_e *Plugin_Expecter) GetNonce(_a0 interface{}) *Plugin_GetNonce_Call {
	return &Plugin_GetNonce_Call{Call: _e.mock.On("GetNonce", _a0)}
}

func (_c *Plugin_GetNonce_Call) Run(run func(_a0 common.Address)) *Plugin_GetNonce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address))
	})
	return _c
}

func (_c *Plugin_GetNonce_Call) Return(_a0 uint64) *Plugin_GetNonce_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetNonce_Call) RunAndReturn(run func(common.Address) uint64) *Plugin_GetNonce_Call {
	_c.Call.Return(run)
	return _c
}

// GetState provides a mock function with given fields: _a0, _a1
func (_m *Plugin) GetState(_a0 common.Address, _a1 common.Hash) common.Hash {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func(common.Address, common.Hash) common.Hash); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// Plugin_GetState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetState'
type Plugin_GetState_Call struct {
	*mock.Call
}

// GetState is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 common.Hash
func (_e *Plugin_Expecter) GetState(_a0 interface{}, _a1 interface{}) *Plugin_GetState_Call {
	return &Plugin_GetState_Call{Call: _e.mock.On("GetState", _a0, _a1)}
}

func (_c *Plugin_GetState_Call) Run(run func(_a0 common.Address, _a1 common.Hash)) *Plugin_GetState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash))
	})
	return _c
}

func (_c *Plugin_GetState_Call) Return(_a0 common.Hash) *Plugin_GetState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetState_Call) RunAndReturn(run func(common.Address, common.Hash) common.Hash) *Plugin_GetState_Call {
	_c.Call.Return(run)
	return _c
}

// RegistryKey provides a mock function with given fields:
func (_m *Plugin) RegistryKey() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RegistryKey")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Plugin_RegistryKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegistryKey'
type Plugin_RegistryKey_Call struct {
	*mock.Call
}

// RegistryKey is a helper method to define mock.On call
func (_e *Plugin_Expecter) RegistryKey() *Plugin_RegistryKey_Call {
	return &Plugin_RegistryKey_Call{Call: _e.mock.On("RegistryKey")}
}

func (_c *Plugin_RegistryKey_Call) Run(run func()) *Plugin_RegistryKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_RegistryKey_Call) Return(_a0 string) *Plugin_RegistryKey_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_RegistryKey_Call) RunAndReturn(run func() string) *Plugin_RegistryKey_Call {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields: _a0
func (_m *Plugin) Reset(_a0 context.Context) {
	_m.Called(_a0)
}

// Plugin_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type Plugin_Reset_Call struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *Plugin_Expecter) Reset(_a0 interface{}) *Plugin_Reset_Call {
	return &Plugin_Reset_Call{Call: _e.mock.On("Reset", _a0)}
}

func (_c *Plugin_Reset_Call) Run(run func(_a0 context.Context)) *Plugin_Reset_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Plugin_Reset_Call) Return() *Plugin_Reset_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_Reset_Call) RunAndReturn(run func(context.Context)) *Plugin_Reset_Call {
	_c.Call.Return(run)
	return _c
}

// RevertToSnapshot provides a mock function with given fields: _a0
func (_m *Plugin) RevertToSnapshot(_a0 int) {
	_m.Called(_a0)
}

// Plugin_RevertToSnapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevertToSnapshot'
type Plugin_RevertToSnapshot_Call struct {
	*mock.Call
}

// RevertToSnapshot is a helper method to define mock.On call
//   - _a0 int
func (_e *Plugin_Expecter) RevertToSnapshot(_a0 interface{}) *Plugin_RevertToSnapshot_Call {
	return &Plugin_RevertToSnapshot_Call{Call: _e.mock.On("RevertToSnapshot", _a0)}
}

func (_c *Plugin_RevertToSnapshot_Call) Run(run func(_a0 int)) *Plugin_RevertToSnapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Plugin_RevertToSnapshot_Call) Return() *Plugin_RevertToSnapshot_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_RevertToSnapshot_Call) RunAndReturn(run func(int)) *Plugin_RevertToSnapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SetBalance provides a mock function with given fields: _a0, _a1
func (_m *Plugin) SetBalance(_a0 common.Address, _a1 *big.Int) {
	_m.Called(_a0, _a1)
}

// Plugin_SetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBalance'
type Plugin_SetBalance_Call struct {
	*mock.Call
}

// SetBalance is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 *big.Int
func (_e *Plugin_Expecter) SetBalance(_a0 interface{}, _a1 interface{}) *Plugin_SetBalance_Call {
	return &Plugin_SetBalance_Call{Call: _e.mock.On("SetBalance", _a0, _a1)}
}

func (_c *Plugin_SetBalance_Call) Run(run func(_a0 common.Address, _a1 *big.Int)) *Plugin_SetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(*big.Int))
	})
	return _c
}

func (_c *Plugin_SetBalance_Call) Return() *Plugin_SetBalance_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_SetBalance_Call) RunAndReturn(run func(common.Address, *big.Int)) *Plugin_SetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// SetCode provides a mock function with given fields: _a0, _a1
func (_m *Plugin) SetCode(_a0 common.Address, _a1 []byte) {
	_m.Called(_a0, _a1)
}

// Plugin_SetCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCode'
type Plugin_SetCode_Call struct {
	*mock.Call
}

// SetCode is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 []byte
func (_e *Plugin_Expecter) SetCode(_a0 interface{}, _a1 interface{}) *Plugin_SetCode_Call {
	return &Plugin_SetCode_Call{Call: _e.mock.On("SetCode", _a0, _a1)}
}

func (_c *Plugin_SetCode_Call) Run(run func(_a0 common.Address, _a1 []byte)) *Plugin_SetCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].([]byte))
	})
	return _c
}

func (_c *Plugin_SetCode_Call) Return() *Plugin_SetCode_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_SetCode_Call) RunAndReturn(run func(common.Address, []byte)) *Plugin_SetCode_Call {
	_c.Call.Return(run)
	return _c
}

// SetNonce provides a mock function with given fields: _a0, _a1
func (_m *Plugin) SetNonce(_a0 common.Address, _a1 uint64) {
	_m.Called(_a0, _a1)
}

// Plugin_SetNonce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNonce'
type Plugin_SetNonce_Call struct {
	*mock.Call
}

// SetNonce is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 uint64
func (_e *Plugin_Expecter) SetNonce(_a0 interface{}, _a1 interface{}) *Plugin_SetNonce_Call {
	return &Plugin_SetNonce_Call{Call: _e.mock.On("SetNonce", _a0, _a1)}
}

func (_c *Plugin_SetNonce_Call) Run(run func(_a0 common.Address, _a1 uint64)) *Plugin_SetNonce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(uint64))
	})
	return _c
}

func (_c *Plugin_SetNonce_Call) Return() *Plugin_SetNonce_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_SetNonce_Call) RunAndReturn(run func(common.Address, uint64)) *Plugin_SetNonce_Call {
	_c.Call.Return(run)
	return _c
}

// SetState provides a mock function with given fields: _a0, _a1, _a2
func (_m *Plugin) SetState(_a0 common.Address, _a1 common.Hash, _a2 common.Hash) {
	_m.Called(_a0, _a1, _a2)
}

// Plugin_SetState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetState'
type Plugin_SetState_Call struct {
	*mock.Call
}

// SetState is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 common.Hash
//   - _a2 common.Hash
func (_e *Plugin_Expecter) SetState(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Plugin_SetState_Call {
	return &Plugin_SetState_Call{Call: _e.mock.On("SetState", _a0, _a1, _a2)}
}

func (_c *Plugin_SetState_Call) Run(run func(_a0 common.Address, _a1 common.Hash, _a2 common.Hash)) *Plugin_SetState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(common.Hash), args[2].(common.Hash))
	})
	return _c
}

func (_c *Plugin_SetState_Call) Return() *Plugin_SetState_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_SetState_Call) RunAndReturn(run func(common.Address, common.Hash, common.Hash)) *Plugin_SetState_Call {
	_c.Call.Return(run)
	return _c
}

// SetStorage provides a mock function with given fields: addr, storage
func (_m *Plugin) SetStorage(addr common.Address, storage map[common.Hash]common.Hash) {
	_m.Called(addr, storage)
}

// Plugin_SetStorage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStorage'
type Plugin_SetStorage_Call struct {
	*mock.Call
}

// SetStorage is a helper method to define mock.On call
//   - addr common.Address
//   - storage map[common.Hash]common.Hash
func (_e *Plugin_Expecter) SetStorage(addr interface{}, storage interface{}) *Plugin_SetStorage_Call {
	return &Plugin_SetStorage_Call{Call: _e.mock.On("SetStorage", addr, storage)}
}

func (_c *Plugin_SetStorage_Call) Run(run func(addr common.Address, storage map[common.Hash]common.Hash)) *Plugin_SetStorage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(map[common.Hash]common.Hash))
	})
	return _c
}

func (_c *Plugin_SetStorage_Call) Return() *Plugin_SetStorage_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_SetStorage_Call) RunAndReturn(run func(common.Address, map[common.Hash]common.Hash)) *Plugin_SetStorage_Call {
	_c.Call.Return(run)
	return _c
}

// Snapshot provides a mock function with given fields:
func (_m *Plugin) Snapshot() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Snapshot")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Plugin_Snapshot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Snapshot'
type Plugin_Snapshot_Call struct {
	*mock.Call
}

// Snapshot is a helper method to define mock.On call
func (_e *Plugin_Expecter) Snapshot() *Plugin_Snapshot_Call {
	return &Plugin_Snapshot_Call{Call: _e.mock.On("Snapshot")}
}

func (_c *Plugin_Snapshot_Call) Run(run func()) *Plugin_Snapshot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_Snapshot_Call) Return(_a0 int) *Plugin_Snapshot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_Snapshot_Call) RunAndReturn(run func() int) *Plugin_Snapshot_Call {
	_c.Call.Return(run)
	return _c
}

// SubBalance provides a mock function with given fields: _a0, _a1
func (_m *Plugin) SubBalance(_a0 common.Address, _a1 *big.Int) {
	_m.Called(_a0, _a1)
}

// Plugin_SubBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubBalance'
type Plugin_SubBalance_Call struct {
	*mock.Call
}

// SubBalance is a helper method to define mock.On call
//   - _a0 common.Address
//   - _a1 *big.Int
func (_e *Plugin_Expecter) SubBalance(_a0 interface{}, _a1 interface{}) *Plugin_SubBalance_Call {
	return &Plugin_SubBalance_Call{Call: _e.mock.On("SubBalance", _a0, _a1)}
}

func (_c *Plugin_SubBalance_Call) Run(run func(_a0 common.Address, _a1 *big.Int)) *Plugin_SubBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Address), args[1].(*big.Int))
	})
	return _c
}

func (_c *Plugin_SubBalance_Call) Return() *Plugin_SubBalance_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_SubBalance_Call) RunAndReturn(run func(common.Address, *big.Int)) *Plugin_SubBalance_Call {
	_c.Call.Return(run)
	return _c
}

// NewPlugin creates a new instance of Plugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *Plugin {
	mock := &Plugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
