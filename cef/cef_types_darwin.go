//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

// cef -> energy 所有结构类型定义 macos

package cef

import . "github.com/energye/energy/v2/api"

// TCefWindowInfo
//
//	Structure representing window information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_win.h">CEF source file: /include/internal/cef_types_win.h (cef_window_info_t)</a>
type TCefWindowInfo struct {
	instance   *tCefWindowInfo
	WindowName string
	X          int32            // Initial window x.
	Y          int32            // Initial window y.
	Width      int32            // Initial window width.
	Height     int32            // Initial window height.
	Hidden     int32            // Set to true (1) to create the view initially hidden.
	ParentView TCefWindowHandle // NSView pointer for the parent view.
	// Set to true (1) to create the browser using windowless (off-screen)
	// rendering. No view will be created for the browser and all rendering will
	// occur via the CefRenderHandler interface. The |parent_view| value will be
	// used to identify monitor info and to act as the parent view for dialogs,
	// context menus, etc. If |parent_view| is not provided then the main screen
	// monitor will be used and some functionality that requires a parent view
	// may not function correctly. In order to create windowless browsers the
	// TCefSettings.windowless_rendering_enabled value must be set to true.
	// Transparent painting is enabled by default but can be disabled by setting
	// TCefBrowserSettings.background_color to an opaque value.
	WindowlessRenderingEnabled int32
	// Set to true (1) to enable shared textures for windowless rendering. Only
	// valid if windowless_rendering_enabled above is also set to true. Currently
	// only supported on Windows (D3D11).
	SharedTextureEnabled      int32
	ExternalBeginFrameEnabled int32            // Set to true (1) to enable the ability to issue BeginFrame from the client application.
	View                      TCefWindowHandle // NSView pointer for the new browser view. Only used with windowed rendering.
}

type tCefWindowInfo struct {
	WindowName                 uintptr //string
	X                          uintptr //int32
	Y                          uintptr //int32
	Width                      uintptr //int32
	Height                     uintptr //int32
	Hidden                     uintptr //int32
	ParentView                 uintptr //TCefWindowHandle
	WindowlessRenderingEnabled uintptr //int32
	SharedTextureEnabled       uintptr //int32
	ExternalBeginFrameEnabled  uintptr //int32
	View                       uintptr //TCefWindowHandle
}

func (m *TCefWindowInfo) Pointer() *tCefWindowInfo {
	if m == nil {
		return nil
	}
	return &tCefWindowInfo{
		WindowName:                 PascalStr(m.WindowName),
		X:                          uintptr(unsafePointer(&m.X)),
		Y:                          uintptr(unsafePointer(&m.Y)),
		Width:                      uintptr(unsafePointer(&m.Width)),
		Height:                     uintptr(unsafePointer(&m.Height)),
		ParentView:                 uintptr(unsafePointer(&m.ParentView)),
		WindowlessRenderingEnabled: uintptr(unsafePointer(&m.WindowlessRenderingEnabled)),
		SharedTextureEnabled:       uintptr(unsafePointer(&m.SharedTextureEnabled)),
		ExternalBeginFrameEnabled:  uintptr(unsafePointer(&m.ExternalBeginFrameEnabled)),
		View:                       uintptr(unsafePointer(&m.View)),
	}
}

// SetInstanceValue 实例指针设置值
func (m *TCefWindowInfo) SetInstanceValue() {
	if m.instance == nil {
		return
	}
	m.instance.WindowName = PascalStr(m.WindowName)
	*(*int32)(unsafePointer(m.instance.X)) = m.X
	*(*int32)(unsafePointer(m.instance.Y)) = m.Y
	*(*int32)(unsafePointer(m.instance.Width)) = m.Width
	*(*int32)(unsafePointer(m.instance.Height)) = m.Height
	*(*uintptr)(unsafePointer(m.instance.ParentView)) = m.ParentView
	*(*int32)(unsafePointer(m.instance.WindowlessRenderingEnabled)) = m.WindowlessRenderingEnabled
	*(*int32)(unsafePointer(m.instance.SharedTextureEnabled)) = m.SharedTextureEnabled
	*(*int32)(unsafePointer(m.instance.ExternalBeginFrameEnabled)) = m.ExternalBeginFrameEnabled
	*(*uintptr)(unsafePointer(m.instance.View)) = m.View
}

func (m *tCefWindowInfo) Convert() *TCefWindowInfo {
	if m == nil {
		return nil
	}
	getInt32 := func(ptr uintptr) int32 {
		if ptr == 0 {
			return 0
		}
		return *(*int32)(unsafePointer(ptr))
	}
	return &TCefWindowInfo{
		instance:                   m,
		WindowName:                 GoStr(m.WindowName),
		X:                          getInt32(m.X),
		Y:                          getInt32(m.Y),
		Width:                      getInt32(m.Width),
		Height:                     getInt32(m.Height),
		ParentView:                 *(*TCefWindowHandle)(unsafePointer(m.ParentView)),
		WindowlessRenderingEnabled: getInt32(m.WindowlessRenderingEnabled),
		SharedTextureEnabled:       getInt32(m.SharedTextureEnabled),
		ExternalBeginFrameEnabled:  getInt32(m.ExternalBeginFrameEnabled),
		View:                       *(*TCefWindowHandle)(unsafePointer(m.View)),
	}
}

// TCefEventHandle
//
//	Native event handle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_win.h">CEF source file: /include/internal/cef_types_win.h (cef_event_handle_t)</a>
type TCefEventHandle = uintptr
