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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
)

// Instance 实例
func (m *ICefMediaAccessCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMediaAccessCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMediaAccessCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefMediaAccessCallback) Cont(allowedPermissions consts.TCefMediaAccessPermissionTypes) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefMediaAccessCallback_Cont).Call(m.Instance(), allowedPermissions.ToPtr())
}

func (m *ICefMediaAccessCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefMediaAccessCallback_Cancel).Call(m.Instance())
}
