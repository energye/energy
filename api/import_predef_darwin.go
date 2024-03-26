//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package api

import (
	"github.com/energye/energy/v2/api/internal/lcl"
)

func NSWindow_FromForm(obj uintptr) uintptr {
	return defSyscallN(lcl.NSWINDOW_FROMFORM, obj)
}

func NSWindow_titleVisibility(n uintptr) uintptr {
	return defSyscallN(lcl.NSWINDOW_TITLEVISIBILITY, n)
}

func NSWindow_setTitleVisibility(n uintptr, flag uintptr) {
	defSyscallN(lcl.NSWINDOW_SETTITLEVISIBILITY, n, flag)
}

func NSWindow_titlebarAppearsTransparent(n uintptr) bool {
	return GoBool(defSyscallN(lcl.NSWINDOW_TITLEBARAPPEARSTRANSPARENT, n))
}

func NSWindow_setTitlebarAppearsTransparent(n uintptr, flag bool) {
	defSyscallN(lcl.NSWINDOW_SETTITLEBARAPPEARSTRANSPARENT, n, PascalBool(flag))
}

func NSWindow_setRepresentedURL(n uintptr, url uintptr) {
	defSyscallN(lcl.NSWINDOW_SETREPRESENTEDURL, n, url)
}

func NSWindow_styleMask(n uintptr) uint {
	return uint(defSyscallN(lcl.NSWINDOW_STYLEMASK, n))
}

func NSWindow_setStyleMask(n uintptr, mask uint) {
	defSyscallN(lcl.NSWINDOW_SETSTYLEMASK, n, uintptr(mask))
}
