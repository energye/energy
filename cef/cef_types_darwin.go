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

func (m *TCefWindowInfo) pointer() *tCefWindowInfo {
	if m == nil {
		return nil
	}
	return &tCefWindowInfo{
		WindowName:                 PascalStr(m.WindowName),
		X:                          uintptr(m.X),
		Y:                          uintptr(m.Y),
		Width:                      uintptr(m.Width),
		Height:                     uintptr(m.Height),
		Hidden:                     uintptr(m.Hidden),
		ParentView:                 uintptr(unsafePointer(&m.ParentView)),
		WindowlessRenderingEnabled: uintptr(m.WindowlessRenderingEnabled),
		SharedTextureEnabled:       uintptr(m.SharedTextureEnabled),
		ExternalBeginFrameEnabled:  uintptr(m.ExternalBeginFrameEnabled),
		View:                       uintptr(unsafePointer(&m.View)),
	}
}
func (m *tCefWindowInfo) convert() *TCefWindowInfo {
	if m == nil {
		return nil
	}
	return &TCefWindowInfo{
		WindowName:                 GoStr(m.WindowName),
		X:                          int32(m.X),
		Y:                          int32(m.Y),
		Width:                      int32(m.Width),
		Height:                     int32(m.Height),
		Hidden:                     int32(m.Hidden),
		ParentView:                 *(*TCefWindowHandle)(unsafePointer(m.ParentView)),
		Menu:                       *(*HMENU)(unsafePointer(m.Menu)),
		WindowlessRenderingEnabled: int32(m.WindowlessRenderingEnabled),
		SharedTextureEnabled:       int32(m.SharedTextureEnabled),
		ExternalBeginFrameEnabled:  int32(m.ExternalBeginFrameEnabled),
		View:                       *(*TCefWindowHandle)(unsafePointer(m.View)),
	}
}

// TCefEventHandle
//
//	Native event handle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_win.h">CEF source file: /include/internal/cef_types_win.h (cef_event_handle_t)</a>
type TCefEventHandle = uintptr
