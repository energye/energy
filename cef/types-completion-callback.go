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
	"github.com/energye/energy/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// CompletionCallbackRef -> ICefCompletionCallback
var CompletionCallbackRef completionCallback

type completionCallback uintptr

func (*completionCallback) New() *ICefCompletionCallback {
	var result uintptr
	imports.Proc(internale_CefCompletionCallbackRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	return &ICefCompletionCallback{instance: unsafe.Pointer(result)}
}

// Instance 实例
func (m *ICefCompletionCallback) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefCompletionCallback) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefCompletionCallback) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefCompletionCallback) SetOnComplete(fn CompletionOnComplete) {
	if !m.IsValid() {
		return
	}
	imports.Proc(internale_CefCompletionCallback_OnComplete).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type CompletionOnComplete func()

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		switch fn.(type) {
		case CompletionOnComplete:
			fn.(CompletionOnComplete)()
		default:
			return false
		}
		return true
	})
}
