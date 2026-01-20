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

func (m *TWindow) Toolbar() {
	nsWindow := m.NSWindow()
	_ = nsWindow
	mask := nsWindow.StyleMask()
	mask |= C.NSWindowStyleMaskFullSizeContentView
	//mask ^= C.NSWindowStyleMaskTitled
	//nsWindow.SetStyleMask(C.NSWindowStyleMaskTitled | C.NSWindowStyleMaskFullSizeContentView | C.NSWindowStyleMaskResizable)
	nsWindow.SetStyleMask(mask)
	nsWindow.SetTitleBarAppearsTransparent(true)
	nsWindow.SetTitleVisibility(1)
	cocoa.NewToolBar(m.NSInstance(), cocoa.ToolbarConfiguration{ShowSeparator: false})
	//m.Frameless()
	//[self.window setMovableByWindowBackground:YES];
	//C.SetWindowRadius(m.NSInstance())
}

func (m *TWindow) Frameless() {
	nsWindow := m.NSWindow()
	_ = nsWindow
	mask := nsWindow.StyleMask()
	mask |= C.NSWindowStyleMaskFullSizeContentView
	nsWindow.SetTitleBarAppearsTransparent(true)
	nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	options := application.GApplication.Options
	if options.DisableResize {
		mask ^= C.NSWindowStyleMaskResizable
	}
	//if options.MacOS.HideTitleBar {
	//	mask ^= C.NSWindowStyleMaskTitled
	//}
	//mask = C.NSWindowStyleMaskClosable
	//mask = C.NSWindowStyleMaskMiniaturizable
	//mask = C.NSWindowStyleMaskResizable
	nsWindow.SetStyleMask(mask)
	//C.SetWindowRadius(m.NSInstance())
}
