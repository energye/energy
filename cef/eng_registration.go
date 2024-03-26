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

// ICefRegistration Parent: ICefBaseRefCounted
//
//	Generic callback interface used for managing the lifespan of a registration.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_registration_capi.h">CEF source file: /include/capi/cef_registration_capi.h (cef_registration_t))
type ICefRegistration interface {
	ICefBaseRefCounted
}

// TCefRegistration Parent: TCefBaseRefCounted
//
//	Generic callback interface used for managing the lifespan of a registration.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_registration_capi.h">CEF source file: /include/capi/cef_registration_capi.h (cef_registration_t))
type TCefRegistration struct {
	TCefBaseRefCounted
}

// RegistrationRef -> ICefRegistration
var RegistrationRef registration

// registration TCefRegistration Ref
type registration uintptr

func (m *registration) UnWrap(data uintptr) ICefRegistration {
	var resultCefRegistration uintptr
	CEF().SysCallN(1251, uintptr(data), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}
