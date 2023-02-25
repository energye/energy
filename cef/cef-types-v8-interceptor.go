//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//CEF v8 对象拦截器 V8InterceptorRef.New
//
// ICefV8Value
package cef

import (
	"fmt"
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// V8InterceptorGetByName 拦截函数 根据name获取一个值
type V8InterceptorGetByName func(name string, object, retVal *ICefV8Value, exception *Exception)

// V8InterceptorGetByIndex 拦截函数 根据index获取一个值
type V8InterceptorGetByIndex func(index int32, object, retVal *ICefV8Value, exception *Exception)

// V8InterceptorSetByName 拦截函数 根据name设置一个值
type V8InterceptorSetByName func(name string, object, value *ICefV8Value, exception *Exception)

// V8InterceptorSetByIndex 拦截函数 根据index设置一个值
type V8InterceptorSetByIndex func(index int32, object, value *ICefV8Value, exception *Exception)

//V8InterceptorRef -> ICefV8Interceptor
var V8InterceptorRef cefV8Interceptor

//cefV8Interceptor
type cefV8Interceptor uintptr

func (*cefV8Interceptor) New() *ICefV8Interceptor {
	var result uintptr
	imports.Proc(internale_CefV8InterceptorRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8Interceptor{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8Interceptor) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

// GetByName 拦截函数 根据name获取一个值
func (m *ICefV8Interceptor) GetByName(fn V8InterceptorGetByName) {
	imports.Proc(internale_CefV8InterceptorRef_GetByName).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// GetByIndex 拦截函数 根据index获取一个值
func (m *ICefV8Interceptor) GetByIndex(fn V8InterceptorGetByIndex) {
	imports.Proc(internale_CefV8InterceptorRef_GetByIndex).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetByName 拦截函数 根据name设置一个值
func (m *ICefV8Interceptor) SetByName(fn V8InterceptorSetByName) {
	imports.Proc(internale_CefV8InterceptorRef_SetByName).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetByIndex 拦截函数 根据index设置一个值
func (m *ICefV8Interceptor) SetByIndex(fn V8InterceptorSetByIndex) {
	imports.Proc(internale_CefV8InterceptorRef_SetByIndex).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Destroy 销毁这个拦截器
func (m *ICefV8Interceptor) Destroy() {
	imports.Proc(internale_CefV8InterceptorRef_Destroy).Call(m.Instance())
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8InterceptorGetByName:
			fmt.Println("----V8InterceptorGetByName", getPtr)
		case V8InterceptorGetByIndex:
			fmt.Println("----V8InterceptorGetByIndex", getPtr)
		case V8InterceptorSetByName:
			fmt.Println("----V8InterceptorSetByName", getPtr)
		case V8InterceptorSetByIndex:
			fmt.Println("----V8InterceptorSetByIndex", getPtr)
		default:
			return false
		}
		return true
	})
}
