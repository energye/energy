//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux
// +build linux

package platform

/*
#cgo linux pkg-config: x11

#include <X11/Xutil.h>
#include <stdlib.h>

void setWMClass(Display *display, Window win, const char *res_name, const char *res_class) {
   XClassHint class_hint;
   class_hint.res_name = (char *)res_name;
   class_hint.res_class = (char *)res_class;
   XSetClassHint(display, win, &class_hint);
}
*/
import "C"
import (
	"unsafe"
)

// linux window properties
func SetWMClass(className, classClass string, windowHandle uintptr) {
	xDisplay := C.XOpenDisplay(nil)
	if xDisplay == nil {
		return
	}
	defer C.XCloseDisplay(xDisplay)
	xWindow := C.Window(windowHandle)

	classNamePtr := C.CString(className)
	classClassPtr := C.CString(classClass)

	//classHint := C.XClassHint{
	//	res_name:  classNamePtr,
	//	res_class: classClassPtr,
	//}
	//C.XSetClassHint(xDisplay, xWindow, &classHint)

	C.setWMClass(xDisplay, xWindow, classNamePtr, classClassPtr)

	C.free(unsafe.Pointer(classNamePtr))
	C.free(unsafe.Pointer(classClassPtr))
}
