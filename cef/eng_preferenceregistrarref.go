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

// ICefPreferenceRegistrarRef Parent: ICEFBaseScopedWrapperRef
//
//	Class that manages custom preference registrations.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_capi.h">CEF source file: /include/capi/cef_preference_capi.h (cef_preference_registrar_t)</a>
type ICefPreferenceRegistrarRef interface {
	ICEFBaseScopedWrapperRef
	// AddPreference
	//  Register a preference with the specified |name| and |default_value|. To
	//  avoid conflicts with built-in preferences the |name| value should contain
	//  an application-specific prefix followed by a period(e.g. "myapp.value").
	//  The contents of |default_value| will be copied. The data type for the
	//  preference will be inferred from |default_value|'s type and cannot be
	//  changed after registration. Returns true(1) on success. Returns false(0)
	//  if |name| is already registered or if |default_value| has an invalid type.
	//  This function must be called from within the scope of the
	//  ICefBrowserProcessHandler.OnRegisterCustomPreferences callback.
	AddPreference(name string, defaultvalue ICefValue) bool // function
}

// TCefPreferenceRegistrarRef Parent: TCEFBaseScopedWrapperRef
//
//	Class that manages custom preference registrations.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_capi.h">CEF source file: /include/capi/cef_preference_capi.h (cef_preference_registrar_t)</a>
type TCefPreferenceRegistrarRef struct {
	TCEFBaseScopedWrapperRef
}

func (m *TCefPreferenceRegistrarRef) AddPreference(name string, defaultvalue ICefValue) bool {
	r1 := CEF().SysCallN(1213, m.Instance(), PascalStr(name), GetObjectUintptr(defaultvalue))
	return GoBool(r1)
}
