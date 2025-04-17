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

// ICefDisplayHandler
type ICefDisplayHandler struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DisplayHandlerRef -> ICefDisplayHandler
var DisplayHandlerRef displayHandler

type displayHandler uintptr

func (*displayHandler) New() *ICefDisplayHandler {
	var result uintptr
	imports.Proc(def.CefDisplayHandlerRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDisplayHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefDisplayHandler) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDisplayHandler) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDisplayHandler) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDisplayHandler) SetOnAddressChange(fn onAddressChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnAddressChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnTitleChange(fn onTitleChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnTitleChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnFaviconUrlChange(fn onFaviconUrlChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnFaviconUrlChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnFullScreenModeChange(fn onFullScreenModeChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnFullScreenModeChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnTooltip(fn onTooltip) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnTooltip).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnStatusMessage(fn onStatusMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnStatusMessage).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnConsoleMessage(fn onConsoleMessage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnConsoleMessage).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnAutoResize(fn onAutoResize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnAutoResize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnLoadingProgressChange(fn onLoadingProgressChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnLoadingProgressChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnCursorChange(fn onCursorChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnCursorChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefDisplayHandler) SetOnMediaAccessChange(fn onMediaAccessChange) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDisplayHandler_OnMediaAccessChange).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onAddressChange func(browser *ICefBrowser, frame *ICefFrame, url string)
type onTitleChange func(browser *ICefBrowser, title string)
type onFaviconUrlChange func(browser *ICefBrowser, iconUrls []string) // TStrings => []string
type onFullScreenModeChange func(browser *ICefBrowser, fullscreen bool)
type onTooltip func(browser *ICefBrowser) (text string, result bool)
type onStatusMessage func(browser *ICefBrowser, value string)
type onConsoleMessage func(browser *ICefBrowser, level consts.TCefLogSeverity, message, source string, line int32) bool
type onAutoResize func(browser *ICefBrowser, newSize *TCefSize) bool
type onLoadingProgressChange func(browser *ICefBrowser, progress float64)
type onCursorChange func(browser *ICefBrowser, cursor consts.TCefCursorHandle, cursorType consts.TCefCursorType, customCursorInfo *TCefCursorInfo) bool
type onMediaAccessChange func(browser *ICefBrowser, hasVideoAccess, hasAudioAccess bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onAddressChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			frame := &ICefFrame{instance: getPtr(1)}
			url := api.GoStr(getVal(2))
			fn.(onAddressChange)(browse, frame, url)
		case onTitleChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			title := api.GoStr(getVal(1))
			fn.(onTitleChange)(browse, title)
		case onFaviconUrlChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			iconUrlsList := lcl.AsStrings(getVal(1))
			var iconUrls []string
			if iconUrlsList.IsValid() {
				count := int(iconUrlsList.Count())
				iconUrls = make([]string, count, count)
				for i := 0; i < count; i++ {
					iconUrls[i] = iconUrlsList.Strings(int32(i))
				}
				iconUrlsList.Free()
			}
			fn.(onFaviconUrlChange)(browse, iconUrls)
		case onFullScreenModeChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			fullscreen := api.GoBool(getVal(1))
			fn.(onFullScreenModeChange)(browse, fullscreen)
		case onTooltip:
			browse := &ICefBrowser{instance: getPtr(0)}
			textPtr := (*uintptr)(getPtr(1))
			result := (*bool)(getPtr(2))
			text, ok := fn.(onTooltip)(browse)
			*textPtr = api.PascalStr(text)
			*result = ok
		case onStatusMessage:
			browse := &ICefBrowser{instance: getPtr(0)}
			value := api.GoStr(getVal(1))
			fn.(onStatusMessage)(browse, value)
		case onConsoleMessage:
			browse := &ICefBrowser{instance: getPtr(0)}
			level := consts.TCefLogSeverity(getVal(1))
			message, source := api.GoStr(2), api.GoStr(3)
			line := int32(getVal(4))
			result := (*bool)(getPtr(5))
			*result = fn.(onConsoleMessage)(browse, level, message, source, line)
		case onAutoResize:
			browse := &ICefBrowser{instance: getPtr(0)}
			newSize := (*TCefSize)(getPtr(1))
			result := (*bool)(getPtr(2))
			*result = fn.(onAutoResize)(browse, newSize)
		case onLoadingProgressChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			progress := *(*float64)(getPtr(1))
			fn.(onLoadingProgressChange)(browse, progress)
		case onCursorChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			cursor := consts.TCefCursorHandle(getVal(1))
			cursorType := consts.TCefCursorType(getVal(2))
			customCursorInfo := (*TCefCursorInfo)(getPtr(3))
			fn.(onCursorChange)(browse, cursor, cursorType, customCursorInfo)
		case onMediaAccessChange:
			browse := &ICefBrowser{instance: getPtr(0)}
			hasVideoAccess, hasAudioAccess := api.GoBool(getVal(1)), api.GoBool(getVal(2))
			fn.(onMediaAccessChange)(browse, hasVideoAccess, hasAudioAccess)
		default:
			return false
		}
		return true
	})
}
