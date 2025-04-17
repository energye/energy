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

// ICefRunContextMenuCallback
//
//	/include/capi/cef_context_menu_handler_capi.h (cef_run_context_menu_callback_t)
type ICefRunContextMenuCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefRunContextMenuCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRunContextMenuCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRunContextMenuCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRunContextMenuCallback) Cont(commandId int32, eventFlags consts.TCefEventFlags) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RunQuickMenuCallback_Cont).Call(m.Instance(), uintptr(commandId), eventFlags.ToPtr())
}

func (m *ICefRunContextMenuCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RunQuickMenuCallback_Cancel).Call(m.Instance())
}
