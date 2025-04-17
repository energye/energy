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
	"github.com/energye/golcl/lcl"
	"unsafe"
)

// ICefFileDialogCallback
//
//	/include/capi/cef_dialog_handler_capi.h (cef_file_dialog_callback_t)
type ICefFileDialogCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefFileDialogCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefFileDialogCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefFileDialogCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefFileDialogCallback) Cont(filePaths []string) {
	if !m.IsValid() {
		return
	}
	fps := lcl.NewStringList()
	if filePaths != nil {
		for _, fp := range filePaths {
			fps.Add(fp)
		}
	}
	imports.Proc(def.FileDialogCallback_Cont).Call(m.Instance(), fps.Instance())
	fps.Free()
}

func (m *ICefFileDialogCallback) Cancel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.FileDialogCallback_Cancel).Call(m.Instance())
}
