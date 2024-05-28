//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows && cgo
// +build !windows,cgo

package cef

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.10 -DMACOSX_DEPLOYMENT_TARGET=10.10
// #cgo darwin CFLAGS: -mmacosx-version-min=10.10
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.10
// #include <stdint.h>
//
// extern void* doCEFEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doGetCEFEventCallbackAddr() {
//    return &doCEFEventCallbackProc;
// }
//
// extern void* doRemoveCEFEventCallbackProc(uintptr_t ptr);
// static void* doRemoveCEFEventCallbackAddr() {
//    return &doRemoveCEFEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doCEFEventCallbackProc
func doCEFEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nil
}

//export doRemoveCEFEventCallbackProc
func doRemoveCEFEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nil
}

var (
	eventCallback       = uintptr(C.doGetCEFEventCallbackAddr())
	removeEventCallback = uintptr(C.doRemoveCEFEventCallbackAddr())
)
