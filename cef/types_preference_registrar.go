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

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCefPreferenceRegistrarRef
// Class that manages custom preference registrations.
// <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_preference_capi.h">CEF source file: /include/capi/cef_preference_capi.h (cef_preference_registrar_t)</see></para>
type TCefPreferenceRegistrarRef struct {
	instance unsafe.Pointer
}

// AddPreference
// Register a preference with the specified |name| and |default_value|. To
// avoid conflicts with built-in preferences the |name| value should contain
// an application-specific prefix followed by a period (e.g. "myapp.value").
// The contents of |default_value| will be copied. The data type for the
// preference will be inferred from |default_value|'s type and cannot be
// changed after registration. Returns true (1) on success. Returns false (0)
// if |name| is already registered or if |default_value| has an invalid type.
// This function must be called from within the scope of the
// ICefBrowserProcessHandler.OnRegisterCustomPreferences callback.
func (m *TCefPreferenceRegistrarRef) AddPreference(name string, defaultValue *ICefValue) bool {
	r1, _, _ := imports.Proc(def.PreferenceRegistrarRef_AddPreference).Call(m.Instance(), api.PascalStr(name), defaultValue.Instance())
	return api.GoBool(r1)
}

// Instance 实例
func (m *TCefPreferenceRegistrarRef) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}
