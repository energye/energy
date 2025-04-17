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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCefRunFileDialogCallback
// include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t)
type ICefRunFileDialogCallback struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// RunFileDialogCallbackRef -> ICefRunFileDialogCallback
var RunFileDialogCallbackRef runFileDialogCallback

type runFileDialogCallback uintptr

func (*runFileDialogCallback) New() *ICefRunFileDialogCallback {
	var result uintptr
	imports.Proc(def.RunFileDialogCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRunFileDialogCallback{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefRunFileDialogCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefRunFileDialogCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefRunFileDialogCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRunFileDialogCallback) SetOnFileDialogDismissed(fn onFileDialogDismissed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.RunFileDialogCallback_OnFileDialogDismissed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onFileDialogDismissed func(filePaths *lcl.TStrings)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onFileDialogDismissed:
			fn.(onFileDialogDismissed)(lcl.AsStrings(getPtr(0)))
		default:
			return false
		}
		return true
	})
}
