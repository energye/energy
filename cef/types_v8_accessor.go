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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefV8Accessor
//
//	Interface that should be implemented to handle V8 accessor calls. Accessor
//	identifiers are registered by calling ICefV8value.SetValue(). The
//	functions of this interface will be called on the thread associated with the
//	V8 accessor.
//	<para><see cref="uCEFTypes|TCefV8Accessor">Implements TCefV8Accessor</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8accessor_t)</see></para>
type ICefV8Accessor struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// V8AccessorRef -> ICefV8Accessor
var V8AccessorRef cefV8Accessor

// cefV8Accessor
type cefV8Accessor uintptr

// New 创建一个v8对象访问器
func (*cefV8Accessor) New() *ICefV8Accessor {
	var result uintptr
	imports.Proc(def.CefV8Accessor_Create).Call(uintptr(unsafe.Pointer(&result)))
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
func (m *ICefV8Accessor) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

func (m *ICefV8Accessor) Free() {
	if m != nil && m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// OnGet
//
//	Handle retrieval the accessor value identified by |name|. |object| is the
//	receiver ('this' object) of the accessor. If retrieval succeeds set
//	|retval| to the return value. If retrieval fails set |exception| to the
//	exception that will be thrown. Return true (1) if accessor retrieval was
//	handled.
func (m *ICefV8Accessor) OnGet(fn onV8AccessorGet) {
	imports.Proc(def.CefV8Accessor_Get).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// OnSet
//
//	Handle assignment of the accessor value identified by |name|. |object| is
//	the receiver ('this' object) of the accessor. |value| is the new value
//	being assigned to the accessor. If assignment fails set |exception| to the
//	exception that will be thrown. Return true (1) if accessor assignment was
//	handled.
func (m *ICefV8Accessor) OnSet(fn onV8AccessorSet) {
	imports.Proc(def.CefV8Accessor_Set).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onV8AccessorGet func(name string, object *ICefV8Value, retVal *ResultV8Value, exception *ResultString) bool
type onV8AccessorSet func(name string, object *ICefV8Value, value *ICefV8Value, exception *ResultString) bool

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onV8AccessorGet:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			retValPtr := (*uintptr)(getPtr(2))
			retVal := &ResultV8Value{}
			exceptionPtr := (*uintptr)(getPtr(3))
			exception := &ResultString{}
			resultPtr := (*bool)(getPtr(4))
			result := fn.(onV8AccessorGet)(name, object, retVal, exception)
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
		case onV8AccessorSet:
			name := api.GoStr(getVal(0))
			object := &ICefV8Value{instance: getPtr(1)}
			value := &ICefV8Value{instance: getPtr(2)}
			exceptionPtr := (*uintptr)(getPtr(3))
			exception := &ResultString{}
			resultPtr := (*bool)(getPtr(4))
			result := fn.(onV8AccessorSet)(name, object, value, exception)
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
