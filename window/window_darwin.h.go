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
	"github.com/energye/lcl/types"
	"unsafe"
)

//export GoLog
func GoLog(message *C.char) {
	msg := C.GoString(message)
	println(msg)
}

// NSWindowAbove = 1;
// NSWindowBelow = -1;
// NSWindowOut = 0;

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

func (m *TWindow) SwitchFrostedMaterial(appearanceName string) {
	C.SwitchFrostedMaterial(m.frostedView, m.NSWindowInstance(), C.CString(appearanceName))
}

func (m *TWindow) Frameless() {
	nsWindow := m.NSWindow()
	nsWindow.SetTitleBarAppearsTransparent(true)
	nsWindow.SetTitleVisibility(types.NSWindowTitleHidden)
	mask := uint(NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable | NSWindowStyleMaskResizable)
	options := application.GApplication.Options
	if options.DisableResize {
		mask ^= NSWindowStyleMaskResizable
	}
	nsWindow.SetStyleMask(mask)
	C.SetFrameless(m.NSInstance())
}
