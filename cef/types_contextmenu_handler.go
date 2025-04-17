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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefContextMenuHandler
type ICefContextMenuHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// ContextMenuHandlerRef -> ICefContextMenuHandler
var ContextMenuHandlerRef contextMenuHandler

type contextMenuHandler uintptr

func (*contextMenuHandler) New() *ICefContextMenuHandler {
	var result uintptr
	imports.Proc(def.CefContextMenuHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefContextMenuHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefContextMenuHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefContextMenuHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefContextMenuHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefContextMenuHandler) SetOnBeforeContextMenu(fn onBeforeContextMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_OnBeforeContextMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefContextMenuHandler) SetRunContextMenu(fn runContextMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_RunContextMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefContextMenuHandler) SetOnContextMenuCommand(fn onContextMenuCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_OnContextMenuCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefContextMenuHandler) SetOnContextMenuDismissed(fn onContextMenuDismissed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_OnContextMenuDismissed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefContextMenuHandler) SetRunQuickMenu(fn runQuickMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_RunQuickMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefContextMenuHandler) SetOnQuickMenuCommand(fn onQuickMenuCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_OnQuickMenuCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefContextMenuHandler) SetOnQuickMenuDismissed(fn onQuickMenuDismissed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefContextMenuHandler_OnQuickMenuDismissed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onBeforeContextMenu func(browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel)
type runContextMenu func(browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel, callback *ICefRunContextMenuCallback) bool
type onContextMenuCommand func(browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId int32, eventFlags consts.TCefEventFlags) bool
type onContextMenuDismissed func(browser *ICefBrowser, frame *ICefFrame)
type runQuickMenu func(browser *ICefBrowser, frame *ICefFrame, location *TCefPoint, size *TCefSize, editStateFlags consts.TCefQuickMenuEditStateFlags, callback *ICefRunQuickMenuCallback) bool
type onQuickMenuCommand func(browser *ICefBrowser, frame *ICefFrame, commandId int32, eventFlags consts.TCefEventFlags) bool
type onQuickMenuDismissed func(browser *ICefBrowser, frame *ICefFrame)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onBeforeContextMenu:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			params := &ICefContextMenuParams{instance: getPtr(2)}
			model := &ICefMenuModel{instance: getPtr(3)}
			fn.(onBeforeContextMenu)(browse, frame, params, model)
			params.Free()
		case runContextMenu:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			params := &ICefContextMenuParams{instance: getPtr(2)}
			model := &ICefMenuModel{instance: getPtr(3)}
			callback := &ICefRunContextMenuCallback{instance: getPtr(4)}
			result := (*bool)(getPtr(5))
			*result = fn.(runContextMenu)(browse, frame, params, model, callback)
			params.Free()
		case onContextMenuCommand:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			params := &ICefContextMenuParams{instance: getPtr(2)}
			commandId := int32(getVal(3))
			eventFlags := consts.TCefEventFlags(getVal(4))
			result := (*bool)(getPtr(5))
			*result = fn.(onContextMenuCommand)(browse, frame, params, commandId, eventFlags)
			params.Free()
		case onContextMenuDismissed:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(onContextMenuDismissed)(browse, frame)
		case runQuickMenu:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			location := (*TCefPoint)(getPtr(2))
			size := (*TCefSize)(getPtr(3))
			editStateFlags := consts.TCefQuickMenuEditStateFlags(getVal(4))
			callback := &ICefRunQuickMenuCallback{instance: getPtr(5)}
			result := (*bool)(getPtr(6))
			*result = fn.(runQuickMenu)(browse, frame, location, size, editStateFlags, callback)
		case onQuickMenuCommand:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			commandId := int32(getVal(2))
			eventFlags := consts.TCefEventFlags(getVal(3))
			result := (*bool)(getPtr(4))
			*result = fn.(onQuickMenuCommand)(browse, frame, commandId, eventFlags)
		case onQuickMenuDismissed:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			fn.(onQuickMenuDismissed)(browse, frame)
		default:
			return false
		}
		return true
	})
}
