//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build linux
// +build linux

package api

import (
	"github.com/energye/energy/v2/api/internal/lcl"
	"unsafe"
)

func GdkWindow_GetXId(g uintptr) (res uintptr) {
	defSyscallN(lcl.GDKWINDOW_GETXID, g, uintptr(unsafe.Pointer(&res)))
	return
}

func GdkWindow_FromForm(obj uintptr) uintptr {
	return defSyscallN(lcl.GDKWINDOW_FROMFORM, obj)
}

func GtkWidget_GetGtkFixed(g uintptr) uintptr {
	return defSyscallN(lcl.GTKWIDGET_GETGTKFIXED, g)
}

func GtkWidget_Window(g uintptr) uintptr {
	return defSyscallN(lcl.GTKWIDGET_WINDOW, g)
}
