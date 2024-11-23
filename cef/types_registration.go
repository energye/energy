//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import "unsafe"

// / Generic callback interface used for managing the lifespan of a registration.
// / <para><see cref="uCEFTypes|TCefRegistration">Implements TCefRegistration</see></para>
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_registration_capi.h">CEF source file: /include/capi/cef_registration_capi.h (cef_registration_t)</see></para>
type ICefRegistration struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefRegistration) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRegistration) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRegistration) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}
