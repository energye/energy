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
//
// TODO 未使用

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefV8Interceptor
//
//	Interface that should be implemented to handle V8 interceptor calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 interceptor. Interceptor's named property handlers (with first argument
//	of type CefString) are called when object is indexed by string. Indexed
//	property handlers (with first argument of type int) are called when object
//	is indexed by integer.
//	<para><see cref="uCEFTypes|TCefV8Interceptor">Implements TCefV8Interceptor</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8interceptor_t)</see></para>
type ICefV8Interceptor struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// V8InterceptorRef -> ICefV8Interceptor
var V8InterceptorRef cefV8Interceptor

// cefV8Interceptor
type cefV8Interceptor uintptr

func (*cefV8Interceptor) New() *ICefV8Interceptor {
	var result uintptr
	imports.Proc(def.CefV8InterceptorRef_Create).Call(uintptr(unsafe.Pointer(&result)))
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

func (m *ICefV8Interceptor) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefV8Interceptor) Free() {
	if m != nil && m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// SetOnGetByName
//
//	Handle retrieval of the interceptor value identified by |name|. |object|
//	is the receiver ('this' object) of the interceptor. If retrieval succeeds,
//	set |retval| to the return value. If the requested value does not exist,
//	don't set either |retval| or |exception|. If retrieval fails, set
//	|exception| to the exception that will be thrown. If the property has an
//	associated accessor, it will be called only if you don't set |retval|.
//	Return true (1) if interceptor retrieval was handled, false (0) otherwise.
func (m *ICefV8Interceptor) SetOnGetByName(fn onV8InterceptorGetByName) {
	imports.Proc(def.CefV8Interceptor_GetByName).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetByIndex
//
//	Handle retrieval of the interceptor value identified by |index|. |object|
//	is the receiver ('this' object) of the interceptor. If retrieval succeeds,
//	set |retval| to the return value. If the requested value does not exist,
//	don't set either |retval| or |exception|. If retrieval fails, set
//	|exception| to the exception that will be thrown. Return true (1) if
//	interceptor retrieval was handled, false (0) otherwise.
func (m *ICefV8Interceptor) SetOnGetByIndex(fn onV8InterceptorGetByIndex) {
	imports.Proc(def.CefV8Interceptor_GetByIndex).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnSetByName
//
//	Handle assignment of the interceptor value identified by |name|. |object|
//	is the receiver ('this' object) of the interceptor. |value| is the new
//	value being assigned to the interceptor. If assignment fails, set
//	|exception| to the exception that will be thrown. This setter will always
//	be called, even when the property has an associated accessor. Return true
//	(1) if interceptor assignment was handled, false (0) otherwise.
func (m *ICefV8Interceptor) SetOnSetByName(fn onV8InterceptorSetByName) {
	imports.Proc(def.CefV8Interceptor_SetByName).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnSetByIndex
//
//	Handle assignment of the interceptor value identified by |index|. |object|
//	is the receiver ('this' object) of the interceptor. |value| is the new
//	value being assigned to the interceptor. If assignment fails, set
//	|exception| to the exception that will be thrown. Return true (1) if
//	interceptor assignment was handled, false (0) otherwise.
func (m *ICefV8Interceptor) SetOnSetByIndex(fn onV8InterceptorSetByIndex) {
	imports.Proc(def.CefV8Interceptor_SetByIndex).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onV8InterceptorGetByName func(name string, object *ICefV8Value) (resultValue *ICefV8Value, exception string, ok bool)
type onV8InterceptorGetByIndex func(index int32, object *ICefV8Value) (resultValue *ICefV8Value, exception string, ok bool)
type onV8InterceptorSetByName func(name string, object, value *ICefV8Value) (exception string, ok bool)
type onV8InterceptorSetByIndex func(index int32, object, value *ICefV8Value) (exception string, ok bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onV8InterceptorGetByName:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			retVal := (*uintptr)(getPtr(2))
			retException := (*uintptr)(getPtr(3))
			retOk := (*bool)(getPtr(4))
			resultValue, exception, ok := fn.(onV8InterceptorGetByName)(name, object)
			if resultValue != nil {
				*retVal = resultValue.Instance()
			}
			if exception != "" {
				*retException = api.PascalStr(exception)
			} else {
				*retException = 0
			}
			*retOk = ok
		case onV8InterceptorGetByIndex:
			index := int32(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			retVal := (*uintptr)(getPtr(2))
			retException := (*uintptr)(getPtr(3))
			retOk := (*bool)(getPtr(4))
			resultValue, exception, ok := fn.(onV8InterceptorGetByIndex)(index, object)
			if resultValue != nil {
				*retVal = resultValue.Instance()
			}
			if exception != "" {
				*retException = api.PascalStr(exception)
			} else {
				*retException = 0
			}
			*retOk = ok
		case onV8InterceptorSetByName:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			value := &ICefV8Value{instance: getPtr(2)}
			retException := (*uintptr)(getPtr(3))
			retOk := (*bool)(getPtr(4))
			exception, ok := fn.(onV8InterceptorSetByName)(name, object, value)
			if exception != "" {
				*retException = api.PascalStr(exception)
			} else {
				*retException = 0
			}
			*retOk = ok
		case onV8InterceptorSetByIndex:
			index := int32(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			value := &ICefV8Value{instance: getPtr(2)}
			retException := (*uintptr)(getPtr(3))
			retOk := (*bool)(getPtr(4))
			exception, ok := fn.(onV8InterceptorSetByIndex)(index, object, value)
			if exception != "" {
				*retException = api.PascalStr(exception)
			} else {
				*retException = 0
			}
			*retOk = ok
		default:
			return false
		}
		return true
	})
}
