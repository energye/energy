//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo && webkit2_4_1

package cgo

/*
#cgo pkg-config: webkit2gtk-4.1

#include <webkit2/webkit2.h>
#include "webkit2gtk4_x.go.h"
*/
import "C"
import (
	"github.com/energye/lcl/types/colors"
	wvTypes "github.com/energye/wv/types/linux"
	"unsafe"
)

const Wkv = wvTypes.Wkv4_1

func WebkitOpenDevTools(instance uintptr) {
	if instance == 0 {
		return
	}
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	C.WebkitOpenDevTools(webview)
}

func WebkitSetBackgroundColor(instance uintptr, color *colors.TARGB) {
	if instance == 0 || color == nil {
		return
	}
	cR := C.gdouble(float64(color.R) / 255.0)
	cG := C.gdouble(float64(color.G) / 255.0)
	cB := C.gdouble(float64(color.B) / 255.0)
	cA := C.gdouble(float64(color.A) / 255.0)
	C.WebkitSetBackgroundColor((*C.WebKitWebView)(unsafe.Pointer(instance)), cR, cG, cB, cA)
}

func WebkitUndo(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("Undo")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}

func WebkitCanUndo(instance uintptr) bool {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	return C.WebkitCanUndo(webview) != 0
}

func WebkitRedo(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("Redo")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}

func WebkitCanRedo(instance uintptr) bool {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	return C.WebkitCanRedo(webview) != 0
}

func WebkitCut(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("Cut")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}

func WebkitCanCut(instance uintptr) bool {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	return C.WebkitCanCut(webview) != 0
}

func WebkitCopy(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("Copy")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}

func WebkitCanCopy(instance uintptr) bool {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	return C.WebkitCanCopy(webview) != 0
}

func WebkitPaste(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("Paste")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}

func WebkitCanPaste(instance uintptr) bool {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	return C.WebkitCanPaste(webview) != 0
}

func WebkitDelete(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("Delete")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}

func WebkitCanDelete(instance uintptr) bool {
	return WebkitCanCut(instance)
}

func WebkitSelectAll(instance uintptr) {
	webview := (*C.WebKitWebView)(unsafe.Pointer(instance))
	command := C.CString("SelectAll")
	defer C.free(unsafe.Pointer(command))
	C.WebkitExecuteEditingCommand(webview, command)
}
