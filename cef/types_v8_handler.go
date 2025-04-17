//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//CEF v8 执行函数拦截

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefV8Handler
type ICefV8Handler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// V8HandlerRef -> ICefV8Handler
var V8HandlerRef cefV8Handler

// cefV8Handler
type cefV8Handler uintptr

// New 创建v8执行函数拦截
func (*cefV8Handler) New() *ICefV8Handler {
	var result uintptr
	imports.Proc(def.CefV8Handler_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Handler{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8Handler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefV8Handler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefV8Handler) Free() {
	if m != nil && m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// Execute 执行拦截函数
func (m *ICefV8Handler) Execute(fn onV8HandlerExecute) {
	imports.Proc(def.CefV8Handler_Execute).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// name 函数名
// object 当前对象
// arguments 函数传入的参数
// retVal 函数执行完返回结果
// exception 返回的异常信息
// return true 执行成功-返回值有效
type onV8HandlerExecute func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onV8HandlerExecute:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			argumentsPtr := getVal(2)
			argumentsLength := int32(getVal(3))
			arguments := &TCefV8ValueArray{instance: unsafe.Pointer(argumentsPtr), arguments: argumentsPtr, argumentsLength: int(argumentsLength), argumentsCollect: make([]*ICefV8Value, int(argumentsLength))}
			retValPtr := (*uintptr)(getPtr(4))
			retVal := &ResultV8Value{}
			exceptionPtr := (*uintptr)(getPtr(5))
			exception := &ResultString{}
			resultPtr := (*bool)(getPtr(6))
			result := fn.(onV8HandlerExecute)(name, object, arguments, retVal, exception)
			if retVal.v8value != nil {
				*retValPtr = retVal.v8value.Instance()
			}
			if exception.Value() != "" {
				*exceptionPtr = api.PascalStr(exception.Value())
			} else {
				*exceptionPtr = 0
			}
			*resultPtr = result
			arguments.Free()
			object.Free()
		default:
			return false
		}
		return true
	})
}
