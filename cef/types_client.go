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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefClient
type ICefClient struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// ClientRef -> ICefClient
var ClientRef cefClient

type cefClient uintptr

// New
//
//	 创建 Client
//		自定义处理器事件
func (*cefClient) New() *ICefClient {
	var result uintptr
	imports.Proc(def.CefClientRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefClient{instance: unsafe.Pointer(result)}
	}
	return nil
}

// NewForChromium
//
//	 创建 Client 事件传递至 Chromium Event
//		获取处理器对象返回nil, 因为事件已传递至 Chromium Event 对应的事件中
func (*cefClient) NewForChromium(chromium IChromium, aDevToolsClient bool) *ICefClient {
	if chromium == nil || chromium.Instance() == 0 {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClientRef_CreateForChromium).Call(chromium.Instance(), api.PascalBool(aDevToolsClient), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefClient{instance: unsafe.Pointer(result), ct: consts.CtOther}
	}
	return nil
}

// Instance 实例
func (m *ICefClient) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefClient) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// IsSelfEvent
//
//	当前client对象是自己事件代理
func (m *ICefClient) IsSelfEvent() bool {
	return m.ct == consts.CtSelfOwn
}

// IsOtherEvent
//
//	当前client对象是其他对象事件代理
//	例如chromium events
func (m *ICefClient) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

func (m *ICefClient) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefClient) GetAudioHandler() *ICefAudioHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetAudioHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefAudioHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetCommandHandler() *ICefCommandHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetCommandHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefCommandHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetContextMenuHandler() *ICefContextMenuHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetContextMenuHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefContextMenuHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetDialogHandler() *ICefDialogHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetDialogHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDialogHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetDisplayHandler() *ICefDisplayHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetDisplayHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDisplayHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetDownloadHandler() *ICefDownloadHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetDownloadHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDownloadHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetDragHandler() *ICefDragHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetDragHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDragHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetFindHandler() *ICefFindHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetFindHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFindHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetFocusHandler() *ICefFocusHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetFocusHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFocusHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetFrameHandler() *ICefFrameHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetFrameHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFrameHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetPermissionHandler() *ICefPermissionHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetPermissionHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPermissionHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetJsdialogHandler() *ICefJsDialogHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetJsdialogHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefJsDialogHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetKeyboardHandler() *ICefKeyboardHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetKeyboardHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefKeyboardHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetLifeSpanHandler() *ICefLifeSpanHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetLifeSpanHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLifeSpanHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetLoadHandler() *ICefLoadHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetLoadHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLoadHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetPrintHandler() *ICefPrintHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetPrintHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPrintHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetRenderHandler() *ICefRenderHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetRenderHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRenderHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}

func (m *ICefClient) GetRequestHandler() *ICefRequestHandler {
	if !m.IsValid() || m.IsOtherEvent() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CefClient_GetRequestHandler).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefRequestHandler{instance: unsafe.Pointer(result)}
	}
	return nil
}
