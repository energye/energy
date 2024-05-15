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

// IV8Handler Parent: ICefv8Handler
//
//	Interface that should be implemented to handle V8 function calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 function.
//	<a cref="uCEFTypes|TCefv8Handler">Implements TCefv8Handler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8handler_t)</a>
type IV8Handler interface {
	ICefv8Handler
	// SetOnExecute
	//  Handle execution of the function identified by |name|. |object| is the
	//  receiver('this' object) of the function. |arguments| is the list of
	//  arguments passed to the function. If execution succeeds set |retval| to
	//  the function return value. If execution fails set |exception| to the
	//  exception that will be thrown. Return true(1) if execution was handled.
	SetOnExecute(fn TOnV8HandlerExecute) // property event
}

// TV8Handler Parent: TCefv8Handler
//
//	Interface that should be implemented to handle V8 function calls. The
//	functions of this interface will be called on the thread associated with the
//	V8 function.
//	<a cref="uCEFTypes|TCefv8Handler">Implements TCefv8Handler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8handler_t)</a>
type TV8Handler struct {
	TCefv8Handler
	executePtr uintptr
}

func NewV8Handler() IV8Handler {
	r1 := CEF().SysCallN(2242)
	return AsV8Handler(r1)
}

func V8HandlerClass() TClass {
	ret := CEF().SysCallN(2241)
	return TClass(ret)
}

func (m *TV8Handler) SetOnExecute(fn TOnV8HandlerExecute) {
	if m.executePtr != 0 {
		RemoveEventElement(m.executePtr)
	}
	m.executePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2243, m.Instance(), m.executePtr)
}
