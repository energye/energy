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

// ISchemeHandlerFactory Parent: ICefSchemeHandlerFactory
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type ISchemeHandlerFactory interface {
	ICefSchemeHandlerFactory
	// SetOnNew
	//  Return a new resource handler instance to handle the request or an NULL
	//  reference to allow default handling of the request. |browser| and |frame|
	//  will be the browser window and frame respectively that originated the
	//  request or NULL if the request did not originate from a browser window
	// (for example, if the request came from ICefUrlRequest). The |request|
	//  object passed to this function cannot be modified.
	SetOnNew(fn TOnSchemeHandlerFactoryNew) // property event
}

// TSchemeHandlerFactory Parent: TCefSchemeHandlerFactory
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type TSchemeHandlerFactory struct {
	TCefSchemeHandlerFactory
	newPtr uintptr
}

func NewSchemeHandlerFactory(aClass TCefResourceHandlerClass) ISchemeHandlerFactory {
	r1 := CEF().SysCallN(2227, uintptr(aClass))
	return AsSchemeHandlerFactory(r1)
}

func SchemeHandlerFactoryClass() TClass {
	ret := CEF().SysCallN(2226)
	return TClass(ret)
}

func (m *TSchemeHandlerFactory) SetOnNew(fn TOnSchemeHandlerFactoryNew) {
	if m.newPtr != 0 {
		RemoveEventElement(m.newPtr)
	}
	m.newPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2228, m.Instance(), m.newPtr)
}
