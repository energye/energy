//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

// cef -> energy 所有结构类型定义 windows

package cef

import (
	"github.com/cyber-xxm/energy/v2/consts"
	. "github.com/cyber-xxm/energy/v2/types"
)

// TCefWindowInfo /include/internal/cef_types_win.h (cef_window_info_t)
type TCefWindowInfo struct {
	instance                   *tCefWindowInfoPtr
	ExStyle                    DWORD
	WindowName                 TCefString
	Style                      DWORD
	X                          Integer
	Y                          Integer
	Width                      Integer
	Height                     Integer
	ParentWindow               TCefWindowHandle
	Menu                       HMENU
	WindowlessRenderingEnabled Integer
	TransparentPaintingEnabled Integer
	SharedTextureEnabled       Integer
	ExternalBeginFrameEnabled  Integer
	Window                     TCefWindowHandle
}

// SetInstanceValue 实例指针设置值
func (m *TCefWindowInfo) setInstanceValue() {
	if m.instance == nil {
		return
	}
	// 字段指针引用赋值, 如果是字符串类型需直接赋值
	m.instance.ExStyle.SetValue(uint32(m.ExStyle))                                      // DWORD
	m.instance.WindowName = UIntptr(m.WindowName.ToPtr())                               // TCefString
	m.instance.Style.SetValue(uint32(m.Style))                                          // DWORD
	m.instance.X.SetValue(int32(m.X))                                                   // Integer
	m.instance.Y.SetValue(int32(m.Y))                                                   // Integer
	m.instance.Width.SetValue(int32(m.Width))                                           // Integer
	m.instance.Height.SetValue(int32(m.Height))                                         // Integer
	m.instance.ParentWindow.SetValue(uintptr(m.ParentWindow))                           // TCefWindowHandle
	m.instance.Menu.SetValue(uintptr(m.Menu))                                           // HMENU
	m.instance.WindowlessRenderingEnabled.SetValue(int32(m.WindowlessRenderingEnabled)) // Integer
	m.instance.TransparentPaintingEnabled.SetValue(int32(m.TransparentPaintingEnabled)) // Integer
	m.instance.SharedTextureEnabled.SetValue(int32(m.SharedTextureEnabled))             // Integer
	m.instance.ExternalBeginFrameEnabled.SetValue(int32(m.ExternalBeginFrameEnabled))   // Integer
	m.instance.Window.SetValue(uintptr(m.Window))                                       // TCefWindowHandle
}

type TCefAcceleratedPaintInfo struct {
	/// Handle for the shared texture. The shared texture is instantiated
	/// without a keyed mutex.
	shared_texture_handle TCefSharedTextureHandle
	/// The pixel format of the texture.
	format consts.TCefColorType
}
