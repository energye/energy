//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows && cgo
// +build !windows,cgo

package exception

// #cgo darwin CFLAGS: -mmacosx-version-min=10.12
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.12
//
// extern void* doExceptionHandlerProc(void* f);
// static void* doExceptionHandlerProcEventAddr() {
//    return &doExceptionHandlerProc;
// }
import "C"

import (
	"unsafe"
)

//export doExceptionHandlerProc
func doExceptionHandlerProc(f unsafe.Pointer) unsafe.Pointer {
	exceptionHandlerProc(uintptr(f))
	return nil
}

var (
	exceptionHandlerProcEventAddr = uintptr(C.doExceptionHandlerProcEventAddr())
)
