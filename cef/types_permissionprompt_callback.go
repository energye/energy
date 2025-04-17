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
	"unsafe"
)

// ICefPermissionPromptCallback
//
//	/include/capi/cef_permission_handler_capi.h (cef_permission_prompt_callback_t)
type ICefPermissionPromptCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefPermissionPromptCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefPermissionPromptCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefPermissionPromptCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefPermissionPromptCallback) Cont(result consts.TCefPermissionRequestResult) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefPermissionPromptCallback_Cont).Call(m.Instance(), result.ToPtr())
}
