//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows && cgo
// +build !windows,cgo

package lcl

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.10 -DMACOSX_DEPLOYMENT_TARGET=10.10
// #cgo darwin CFLAGS: -mmacosx-version-min=10.10
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.10
// #include <stdint.h>
//
// extern void* doEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doGetEventCallbackAddr() {
//    return &doEventCallbackProc;
// }
//
// extern void* doMessageCallbackProc(uintptr_t f, void* msg);
// static void* doGetMessageCallbackAddr() {
//    return &doMessageCallbackProc;
// }
//
// extern void* doThreadSyncCallbackProc();
// static void* doGetThreadSyncCallbackAddr() {
//    return &doThreadSyncCallbackProc;
// }
//
// extern void* doRequestCallCreateParamsCallbackProc(uintptr_t ptr, void* sender, void* params);
// static void* doRequestCallCreateParamsCallbackAddr() {
//    return &doRequestCallCreateParamsCallbackProc;
// }
//
// extern void* doRequestCallFormCreateCallbackProc(uintptr_t ptr, void* sender);
// static void* doRequestCallFormCreateCallbackAddr() {
//    return &doRequestCallFormCreateCallbackProc;
// }
//
// extern void* doRemoveEventCallbackProc(uintptr_t ptr);
// static void* doRemoveEventCallbackAddr() {
//    return &doRemoveEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doEventCallbackProc
func doEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nil
}

//export doMessageCallbackProc
func doMessageCallbackProc(f C.uintptr_t, msg unsafe.Pointer) unsafe.Pointer {
	messageCallbackProc(uintptr(f), uintptr(msg))
	return nil
}

//export doThreadSyncCallbackProc
func doThreadSyncCallbackProc() unsafe.Pointer {
	threadSyncCallbackProc()
	return nil
}

//export doRequestCallCreateParamsCallbackProc
func doRequestCallCreateParamsCallbackProc(ptr C.uintptr_t, sender, params unsafe.Pointer) unsafe.Pointer {
	requestCallCreateParamsCallbackProc(uintptr(ptr), uintptr(sender), uintptr(params))
	return nil
}

//export doRequestCallFormCreateCallbackProc
func doRequestCallFormCreateCallbackProc(ptr C.uintptr_t, sender unsafe.Pointer) unsafe.Pointer {
	requestCallFormCreateCallbackProc(uintptr(ptr), uintptr(sender))
	return nil
}

//export doRemoveEventCallbackProc
func doRemoveEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nil
}

var (
	eventCallback                   = uintptr(C.doGetEventCallbackAddr())
	messageCallback                 = uintptr(C.doGetMessageCallbackAddr())
	threadSyncCallback              = uintptr(C.doGetThreadSyncCallbackAddr())
	requestCallCreateParamsCallback = uintptr(C.doRequestCallCreateParamsCallbackAddr())
	requestCallFormCreateCallback   = uintptr(C.doRequestCallFormCreateCallbackAddr())
	removeEventCallback             = uintptr(C.doRemoveEventCallbackAddr())
)
