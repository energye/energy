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

// ICefSchemeRegistrarRef Parent: ICEFBaseScopedWrapperRef
//
//	Class that manages custom scheme registrations.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_registrar_t)</a>
type ICefSchemeRegistrarRef interface {
	ICEFBaseScopedWrapperRef
	// AddCustomScheme
	//  Register a custom scheme. This function should not be called for the
	//  built-in HTTP, HTTPS, FILE, FTP, ABOUT and DATA schemes.
	//  This function may be called on any thread. It should only be called once
	//  per unique |scheme_name| value. If |scheme_name| is already registered or
	//  if an error occurs this function will return false(0).
	//  <a>See the CEF_SCHEME_OPTION_* constants in the uCEFConstants unit for possible values for |options|.</a>
	AddCustomScheme(schemeName string, options TCefSchemeOptions) bool // function
}

// TCefSchemeRegistrarRef Parent: TCEFBaseScopedWrapperRef
//
//	Class that manages custom scheme registrations.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_registrar_t)</a>
type TCefSchemeRegistrarRef struct {
	TCEFBaseScopedWrapperRef
}

func (m *TCefSchemeRegistrarRef) AddCustomScheme(schemeName string, options TCefSchemeOptions) bool {
	r1 := CEF().SysCallN(1344, m.Instance(), PascalStr(schemeName), uintptr(options))
	return GoBool(r1)
}
