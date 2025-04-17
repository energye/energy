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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCefSchemeRegistrarRef
type TCefSchemeRegistrarRef struct {
	instance unsafe.Pointer
}

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

// AddCustomScheme 注册自定义方案。不应为内置的HTTP、HTTPS、FILE、FTP、ABOUT和DATA方案调用此函数。
func (m *TCefSchemeRegistrarRef) AddCustomScheme(schemeName string, options consts.CefSchemeOption) bool {
	if !m.IsValid() {
		return false
	}
	if application.Is49() {
		// CEF49
		var (
			isStandard, isLocal, isDisplayIsolated bool
		)
		if options&consts.CEF_SCHEME_OPTION_STANDARD == consts.CEF_SCHEME_OPTION_STANDARD {
			isStandard = true
		}
		if options&consts.CEF_SCHEME_OPTION_LOCAL == consts.CEF_SCHEME_OPTION_LOCAL {
			isLocal = true
		}
		if options&consts.CEF_SCHEME_OPTION_DISPLAY_ISOLATED == consts.CEF_SCHEME_OPTION_DISPLAY_ISOLATED {
			isDisplayIsolated = true
		}
		r1, _, _ := imports.Proc(def.SchemeRegistrarRef_AddCustomScheme).Call(m.Instance(), api.PascalStr(schemeName), api.PascalBool(isStandard), api.PascalBool(isLocal), api.PascalBool(isDisplayIsolated))
		return api.GoBool(r1)
	} else {
		r1, _, _ := imports.Proc(def.SchemeRegistrarRef_AddCustomScheme).Call(m.Instance(), api.PascalStr(schemeName), uintptr(options))
		return api.GoBool(r1)
	}
}
