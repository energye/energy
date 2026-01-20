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

package window

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

#include "window_darwin.h"
*/
import "C"
import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/pkgs/cocoa"
	"github.com/energye/lcl/types"
	"unsafe"
)

//export GoLog
func GoLog(message *C.char) {
	msg := C.GoString(message)
	println(msg)
}

func (m *TWindow) DragWindow() {
	C.DragWindow(m.NSInstance())
}

func (m *TWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	C.SetWindowBackgroundColor(m.NSInstance(), C.int(red), C.int(green), C.int(blue), C.int(alpha))
}

func (m *TWindow) SetWindowTransparent() {
	frostedView := C.SetWindowTransparent(m.NSWindowInstance())
	m.frostedView = unsafe.Pointer(frostedView)
}

func (m *TWindow) SwitchFrostedMaterial(appearanceName application.AppearanceNamed) {
	C.SwitchFrostedMaterial(m.frostedView, m.NSWindowInstance(), C.CString(string(appearanceName)))
}

func (m *TWindow) SetWindowRadius() {
	options := application.GApplication.Options
	if options.Frameless {
		if options.MacOS.WindowRadius > 0.0 {
			C.SetWindowRadius(m.NSInstance(), C.float(options.MacOS.WindowRadius))
		}
	}
}

func (m *TWindow) TitleBar() {
	options := application.GApplication.Options
	nsWindow := m.NSWindow()
	mask := nsWindow.StyleMask()
	if options.DisableSystemMenu {
		mask ^= C.NSWindowStyleMaskClosable
	}
	if options.DisableMinimize {
		mask ^= C.NSWindowStyleMaskMiniaturizable
	}
	if options.DisableResize || options.DisableMaximize {
		mask ^= C.NSWindowStyleMaskResizable
	}
	if options.MacOS.FullSizeContent {
		mask |= C.NSWindowStyleMaskFullSizeContentView
	}
	nsWindow.SetStyleMask(mask)
	nsWindow.SetTitleBarAppearsTransparent(options.MacOS.TitleTransparent)
	if options.MacOS.TitleHideText {
		nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	}
	toolBar := options.MacOS.ToolBar
	if toolBar != nil {
		cocoa.NewToolBar(m.NSInstance(), cocoa.ToolbarConfiguration{ShowSeparator: toolBar.ShowSeparator})
	}
}

func (m *TWindow) Frameless() {
	options := application.GApplication.Options
	if options.Frameless {
		nsWindow := m.NSWindow()
		mask := nsWindow.StyleMask()
		mask ^= C.NSWindowStyleMaskTitled
		nsWindow.SetStyleMask(mask)
	}
}
