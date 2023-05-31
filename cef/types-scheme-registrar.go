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
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
)

// Instance 实例
func (m *TCefSchemeRegistrarRef) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *TCefSchemeRegistrarRef) Free() {
	if m.instance != nil {
		m.instance = nil
	}
}

func (m *TCefSchemeRegistrarRef) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *TCefSchemeRegistrarRef) AddCustomScheme(schemeName string, options consts.CefSchemeOption) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(internale_SchemeRegistrarRef_AddCustomScheme).Call(m.Instance(), api.PascalStr(schemeName), uintptr(options))
	return api.GoBool(r1)
}
