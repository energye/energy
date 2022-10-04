//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

//go:build !windows && cgo
// +build !windows,cgo

package cef

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.8 -DMACOSX_DEPLOYMENT_TARGET=10.8
// #cgo darwin CFLAGS: -mmacosx-version-min=10.8
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.8
//
// extern void* doCefIPCCallbackFuncEventProc(void* f, void* args, long argcount);
// static void* doGetCefIPCCallbackFuncEventAddr() {
//    return &doCefIPCCallbackFuncEventProc;
// }
//
// extern void* doCefWindowBindEventProc(void* f, void* args, long argcount);
// static void* doGetCefWindowBindEventAddr() {
//    return &doCefWindowBindEventProc;
// }
//
// extern void* doApplicationQueueAsyncCallEventProc(void* f);
// static void* doGetApplicationQueueAsyncCallEventAddr() {
//    return &doApplicationQueueAsyncCallEventProc;
// }
import "C"

import (
	"unsafe"
)

//export doCefIPCCallbackFuncEventProc
func doCefIPCCallbackFuncEventProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	cefIPCEventProc(uintptr(f), uintptr(args), int(argcount))
	return nullptr
}

//export doCefWindowBindEventProc
func doCefWindowBindEventProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	cefWindowBindCallbackEventProc(uintptr(f), uintptr(args), int(argcount))
	return nullptr
}

//export doApplicationQueueAsyncCallEventProc
func doApplicationQueueAsyncCallEventProc(f unsafe.Pointer) unsafe.Pointer {
	applicationQueueAsyncCallProc(uintptr(f))
	return nullptr
}

var (
	cefIPCCallbackFuncEvent        = uintptr(C.doGetCefIPCCallbackFuncEventAddr())
	cefWindowBindEvent             = uintptr(C.doGetCefWindowBindEventAddr())
	applicationQueueAsyncCallEvent = uintptr(C.doGetApplicationQueueAsyncCallEventAddr())
)
