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

package cocoa

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa

#include "ns_window.h"

extern GoArguments* doOnWindowDelegateEvent(TCallbackContext *CContext);

*/
import "C"
import (
	"unsafe"
)

//export doOnWindowDelegateEvent
func doOnWindowDelegateEvent(CContext *C.TCallbackContext) *C.GoArguments {
	ctx := TCallbackContext{
		Identifier: C.GoString(CContext.identifier),
		Value:      C.GoString(CContext.value),
		Index:      int(CContext.index),
		Owner:      CContext.owner,
		Sender:     CContext.sender,
	}
	cArguments := CContext.arguments
	if cArguments != nil {
		ctx.Arguments = &OCGoArguments{arguments: Pointer(cArguments), count: int(cArguments.Count)}
	}
	eventId := ctx.Identifier
	cb := eventList[eventId]
	if cb == nil {
		return nil
	}
	if result := cb.cb(&ctx); result != nil {
		return result.ToOC()
	} else {
		return nil
	}
}

func _CreateEventCallback() C.TEventCallback {
	return (C.TEventCallback)(C.doOnWindowDelegateEvent)
}

func CreateWindowDelegate(window unsafe.Pointer) unsafe.Pointer {
	if window == nil {
		return nil
	}
	windowEventCallback := _CreateEventCallback()
	windowDelegate := C.CreateWindowDelegate(window, windowEventCallback)
	return unsafe.Pointer(windowDelegate)
}

func WindowRestore(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowRestore(window)
}

func WindowMinimized(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowMinimized(window)
}

func WindowExitMinimized(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowExitMinimized(window)
}

func WindowMaximize(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowMaximize(window)
}

func WindowEnterFullScreen(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowEnterFullScreen(window)
}

func WindowExitFullScreen(window unsafe.Pointer) {
	if window == nil {
		return
	}
	C.WindowExitFullScreen(window)
}

func DragWindow(window unsafe.Pointer) {
	C.DragWindow(window)
}

func SetWindowBackgroundColor(window unsafe.Pointer, red, green, blue, alpha uint8) {
	C.SetWindowBackgroundColor(window, C.int(red), C.int(green), C.int(blue), C.int(alpha))
}

func SetWindowTransparent(windowInstance unsafe.Pointer) unsafe.Pointer {
	frostedView := C.SetWindowTransparent(windowInstance)
	return unsafe.Pointer(frostedView)
}

func WindowSwitchFrostedMaterial(frostedView, windowInstance unsafe.Pointer, appearanceName string) {
	C.SwitchFrostedMaterial(frostedView, windowInstance, C.CString(appearanceName))
}

func SetWindowRadius(window unsafe.Pointer, value float32) {
	C.SetWindowRadius(window, C.float(value))
}

func WindowAddSubview(window, view unsafe.Pointer, x, y, width, height float32) {
	C.WindowAddSubview(window, view, C.float(x), C.float(y), C.float(width), C.float(height))
}
