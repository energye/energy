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

// ICefFrameHandler
type ICefFrameHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// FrameHandlerRef -> ICefFrameHandler
var FrameHandlerRef frameHandler

type frameHandler uintptr

func (*frameHandler) New() *ICefFrameHandler {
	var result uintptr
	imports.Proc(def.CefFrameHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFrameHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefFrameHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefFrameHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefFrameHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefFrameHandler) SetOnFrameCreated(fn onFrameCreated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFrameHandler_OnFrameCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefFrameHandler) SetOnFrameAttached(fn onFrameAttached) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFrameHandler_OnFrameAttached).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefFrameHandler) SetOnFrameDetached(fn onFrameDetached) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFrameHandler_OnFrameDetached).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefFrameHandler) SetOnMainFrameChanged(fn onMainFrameChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefFrameHandler_OnMainFrameChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onFrameCreated func(browser *ICefBrowser, frame *ICefFrame)
type onFrameAttached func(browser *ICefBrowser, frame *ICefFrame, reattached bool)
type onFrameDetached func(browser *ICefBrowser, frame *ICefFrame)
type onMainFrameChanged func(browser *ICefBrowser, oldFrame, newFrame *ICefFrame)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onFrameCreated:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(onFrameCreated)(browse, frame)
		case onFrameAttached:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			reattached := api.GoBool(getVal(2))
			fn.(onFrameAttached)(browse, frame, reattached)
		case onFrameDetached:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(onFrameDetached)(browse, frame)
		case onMainFrameChanged:
			browse := &ICefBrowser{instance: getPtr(0)}
			oldFrame := &ICefFrame{instance: getPtr(1)}
			newFrame := &ICefFrame{instance: getPtr(2)}
			fn.(onMainFrameChanged)(browse, oldFrame, newFrame)
		default:
			return false
		}
		return true
	})
}
