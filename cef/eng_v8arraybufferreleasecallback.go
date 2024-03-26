//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// IV8ArrayBufferReleaseCallback Parent: ICefv8ArrayBufferReleaseCallback
//
//	Callback interface that is passed to ICefV8value.CreateArrayBuffer.
//	<a cref="uCEFTypes|TCefv8ArrayBufferReleaseCallback">Implements TCefv8ArrayBufferReleaseCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8array_buffer_release_callback_t)</a>
type IV8ArrayBufferReleaseCallback interface {
	ICefv8ArrayBufferReleaseCallback
	// SetOnReleaseBuffer
	//  Called to release |buffer| when the ArrayBuffer JS object is garbage
	//  collected. |buffer| is the value that was passed to CreateArrayBuffer
	//  along with this object.
	SetOnReleaseBuffer(fn TOnV8ArrayBufferReleaseBuffer) // property event
}

// TV8ArrayBufferReleaseCallback Parent: TCefv8ArrayBufferReleaseCallback
//
//	Callback interface that is passed to ICefV8value.CreateArrayBuffer.
//	<a cref="uCEFTypes|TCefv8ArrayBufferReleaseCallback">Implements TCefv8ArrayBufferReleaseCallback</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8array_buffer_release_callback_t)</a>
type TV8ArrayBufferReleaseCallback struct {
	TCefv8ArrayBufferReleaseCallback
	releaseBufferPtr uintptr
}

func NewV8ArrayBufferReleaseCallback() IV8ArrayBufferReleaseCallback {
	r1 := CEF().SysCallN(2234)
	return AsV8ArrayBufferReleaseCallback(r1)
}

func V8ArrayBufferReleaseCallbackClass() TClass {
	ret := CEF().SysCallN(2233)
	return TClass(ret)
}

func (m *TV8ArrayBufferReleaseCallback) SetOnReleaseBuffer(fn TOnV8ArrayBufferReleaseBuffer) {
	if m.releaseBufferPtr != 0 {
		RemoveEventElement(m.releaseBufferPtr)
	}
	m.releaseBufferPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2235, m.Instance(), m.releaseBufferPtr)
}
