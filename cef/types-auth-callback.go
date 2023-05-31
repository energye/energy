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
	"github.com/energye/golcl/lcl/api"
)

// Instance 实例
func (m *ICefAuthCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefAuthCallback) Free() {
	m.base.Free(m.Instance())
}

func (m *ICefAuthCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefAuthCallback) Cont(username, password string) {
	if m.instance != nil {
		imports.Proc(internale_CefAuthCallback_Cont).Call(m.Instance(), api.PascalStr(username), api.PascalStr(password))
	}
}

func (m *ICefAuthCallback) Cancel() {
	if m.instance != nil {
		imports.Proc(internale_CefAuthCallback_Cancel).Call(m.Instance())
	}
}
