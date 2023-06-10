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
	"github.com/energye/golcl/lcl/api"
)

// Instance 实例
func (m *ICefJsDialogCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefJsDialogCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefJsDialogCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefJsDialogCallback) Cont(success bool, userInput string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefJsDialogCallback_Cont).Call(m.Instance(), api.PascalBool(success), api.PascalStr(userInput))
}
