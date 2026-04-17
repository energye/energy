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

package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "ns_window.h"
*/
import "C"
import (
	"fmt"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/lcl/types"
	"unsafe"
)

type NSWindow struct {
	NSResponder
	frostedView INSVisualEffectView
}

func AsNSWindow(ptr unsafe.Pointer) INSWindow {
	if ptr == nil {
		return nil
	}
	m := new(NSWindow)
	m.SetInstance(ptr)
	m.registerEvents()
	return m
}

func (m *NSWindow) registerEvents() {
	nsWindow := unsafe.Pointer(m.Instance())
	windowResizeEventId := fmt.Sprintf("%d_%v", TWindowEventDidResize, nsWindow)
	RegisterEvent(windowResizeEventId, MakeNotifyEvent(func(identifier string, owner Pointer, sender Pointer) *GoArguments {
		m.doWindowDidResie()
		return nil
	}))
	windowEnterFullScreenEventId := fmt.Sprintf("%d_%v", TWindowEventEnterFullScreen, nsWindow)
	RegisterEvent(windowEnterFullScreenEventId, MakeNotifyEvent(func(identifier string, owner Pointer, sender Pointer) *GoArguments {
		m.doWindowWillEnterFullScreen()
		return nil
	}))
	windowExitFullScreenEventId := fmt.Sprintf("%d_%v", TWindowEventExitFullScreen, nsWindow)
	RegisterEvent(windowExitFullScreenEventId, MakeNotifyEvent(func(identifier string, owner Pointer, sender Pointer) *GoArguments {
		m.doWindowDidExitFullScreen()
		return nil
	}))
	windowWillUseFullScreenPresentationOptionsEventId := fmt.Sprintf("%d_%v", TWindowEventWillUseFullScreenPresentationOptions, nsWindow)
	RegisterEvent(windowWillUseFullScreenPresentationOptionsEventId, MakeDelegateEvent(func(arguments *OCGoArguments, owner Pointer, sender Pointer) *GoArguments {
		defaultOptions := NSApplicationPresentationAutoHideToolbar | NSApplicationPresentationAutoHideMenuBar | NSApplicationPresentationFullScreen
		if arguments != nil {
			defaultOptions = NSApplicationPresentationOptions(arguments.GetInt(0))
		}
		defaultOptions = m.doWindowWillUseFullScreenPresentationOptions(defaultOptions)
		result := new(GoArguments)
		result.Add(defaultOptions)
		return result
	}))
}

func (m *NSWindow) Restore() {
	if m == nil {
		return
	}
	C.WindowRestore(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) Minimized() {
	if m == nil {
		return
	}
	C.WindowMinimized(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) ExitMinimized() {
	if m == nil {
		return
	}
	C.WindowExitMinimized(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) Maximize() {
	if m == nil {
		return
	}
	C.WindowMaximize(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) EnterFullScreen() {
	if m == nil {
		return
	}
	C.WindowEnterFullScreen(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) ExitFullScreen() {
	if m == nil {
		return
	}
	C.WindowExitFullScreen(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) Drag() {
	if m == nil {
		return
	}
	C.DragWindow(unsafe.Pointer(m.Instance()))
}

func (m *NSWindow) SetBackgroundColor(red, green, blue, alpha uint8) {
	if m == nil {
		return
	}
	C.SetWindowBackgroundColor(unsafe.Pointer(m.Instance()), C.int(red), C.int(green), C.int(blue), C.int(alpha))
}

func (m *NSWindow) SetRadius(value float32) {
	if m == nil {
		return
	}
	C.SetWindowRadius(unsafe.Pointer(m.Instance()), C.float(value))
}

func (m *NSWindow) SetTransparent() INSVisualEffectView {
	if m.frostedView != nil {
		return m.frostedView
	}
	frostedView := C.SetWindowTransparent(unsafe.Pointer(m.Instance()))
	nsFrostedView := AsNSVisualEffectView(unsafe.Pointer(frostedView))
	m.frostedView = nsFrostedView
	return nsFrostedView
}

func (m *NSWindow) SwitchFrostedMaterial(appearanceName string) {
	C.SwitchFrostedMaterial(unsafe.Pointer(m.frostedView.Instance()), unsafe.Pointer(m.Instance()), C.CString(appearanceName))
}

func (m *NSWindow) AddSubview(view INSView, x, y, width, height float32) {
	C.WindowAddSubview(unsafe.Pointer(m.Instance()), unsafe.Pointer(view.Instance()), C.float(x), C.float(y), C.float(width), C.float(height))
}

func (m *NSWindow) ContentViewFrame() (rect types.TRect) {
	nsRect := C.WindowContentViewFrame(unsafe.Pointer(m.Instance()))
	rect.Left = int32(nsRect.Left)
	rect.Top = int32(nsRect.Top)
	rect.Right = int32(nsRect.Right)
	rect.Bottom = int32(nsRect.Bottom)
	return
}

func (m *NSWindow) UpdateFrostedViewBounds() {
	if m.frostedView != nil && m.frostedView.Instance() != 0 {
		C.UpdateFrostedViewBounds(unsafe.Pointer(m.frostedView.Instance()), unsafe.Pointer(m.Instance()))
	}
}

func (m *NSWindow) doWindowDidResie() {
	m.UpdateFrostedViewBounds()
}

func (m *NSWindow) doWindowWillEnterFullScreen() {

}

func (m *NSWindow) doWindowDidExitFullScreen() {

}

func (m *NSWindow) doWindowWillUseFullScreenPresentationOptions(options NSApplicationPresentationOptions) NSApplicationPresentationOptions {
	//println("[DEBUG] WindowWillUseFullScreenPresentationOptions options:", options)
	return options
}
