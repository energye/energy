//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// v8 缓存释放时回调 V8ArrayBufferReleaseCallbackRef.New()
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

//V8ArrayBufferReleaseCallback 释放时回调函数
// array buffer 缓存
// 返回 true:释放buffer, false:不释放
type V8ArrayBufferReleaseCallback func(buffer uintptr) bool

//V8ArrayBufferReleaseCallbackRef -> ICefV8ArrayBufferReleaseCallback
var V8ArrayBufferReleaseCallbackRef cefV8ArrayBufferReleaseCallback

//cefV8ArrayBufferReleaseCallback
type cefV8ArrayBufferReleaseCallback uintptr

// New 创建V8 ArrayBuffer 释放回调函数
//
// 默认自动释放 buffer
func (*cefV8ArrayBufferReleaseCallback) New() *ICefV8ArrayBufferReleaseCallback {
	var result uintptr
	imports.Proc(internale_CefV8ArrayBufferReleaseCallback_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefV8ArrayBufferReleaseCallback{
		instance: unsafe.Pointer(result),
	}
}

// Instance 实例
func (m *ICefV8ArrayBufferReleaseCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

//ReleaseBuffer 释放时回调函数, 默认自动释放
//
// array buffer 缓存
//
// 返回 true:释放buffer, false:不释放buffer
func (m *ICefV8ArrayBufferReleaseCallback) ReleaseBuffer(fn V8ArrayBufferReleaseCallback) {
	imports.Proc(internale_CefV8ArrayBufferReleaseCallback_ReleaseBuffer).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case V8ArrayBufferReleaseCallback:
			returnIsReleaseBuf := (*bool)(getPtr(0))
			buffPtr := getVal(1)
			*returnIsReleaseBuf = fn.(V8ArrayBufferReleaseCallback)(buffPtr)
		default:
			return false
		}
		return true
	})
}
