//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF v8 对象字段&值访问器(读取/写入) V8AccessorRef.New()
//
// ICefV8Value
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// V8AccessorGet 读取时函数
// name 字段或对象名
// retVal 读取返回的值
// exception 返回的异常信息
// return true 读取成功-返回值有效
type V8AccessorGet func(name string, object *ICefV8Value, retVal *ResultV8Value, exception *ResultString) bool

// V8AccessorSet 写入时函数
// name 字段或对象名
// value 将要写入的新值
// exception 返回的异常信息
// return true 写入成功
type V8AccessorSet func(name string, object *ICefV8Value, value *ICefV8Value, exception *ResultString) bool

//V8AccessorRef -> ICefV8Accessor
var V8AccessorRef cefV8Accessor

// cefV8Accessor
type cefV8Accessor uintptr

// New 创建一个v8对象访问器
func (*cefV8Accessor) New() *ICefV8Accessor {
	var result uintptr
	imports.Proc(internale_CefV8Accessor_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Accessor{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8Accessor) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

//Get 读取函数
func (m *ICefV8Accessor) Get(fn V8AccessorGet) {
	imports.Proc(internale_CefV8Accessor_Get).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

//Set 写入访问函数
func (m *ICefV8Accessor) Set(fn V8AccessorSet) {
	imports.Proc(internale_CefV8Accessor_Set).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefV8Accessor) Destroy() {
	imports.Proc(internale_CefV8Accessor_Destroy).Call(m.Instance())
}

func (m *ICefV8Accessor) Free() {
	m.Destroy()
	m.instance = nil
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8AccessorGet:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			retValPtr := (*uintptr)(getPtr(2))
			retVal := &ResultV8Value{}
			exceptionPtr := (*uintptr)(getPtr(3))
			exception := &ResultString{}
			resultPtr := (*bool)(getPtr(4))
			result := fn.(V8AccessorGet)(name, object, retVal, exception)
			if retVal.v8value != nil {
				*retValPtr = retVal.v8value.Instance()
			}
			if exception.Value() != "" {
				*exceptionPtr = api.PascalStr(exception.Value())
			} else {
				*exceptionPtr = 0
			}
			*resultPtr = result
			//object.Free()
		case V8AccessorSet:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			value := &ICefV8Value{instance: getPtr(2)}
			exceptionPtr := (*uintptr)(getPtr(3))
			exception := &ResultString{}
			resultPtr := (*bool)(getPtr(4))
			result := fn.(V8AccessorSet)(name, object, value, exception)
			if exception.Value() != "" {
				*exceptionPtr = api.PascalStr(exception.Value())
			} else {
				*exceptionPtr = 0
			}
			*resultPtr = result
			//object.Free()
			//value.Free()
		default:
			return false
		}
		return true
	})
}
