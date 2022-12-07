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
// extern void* doApplicationQueueAsyncCallEventProc(void* f);
// static void* doGetApplicationQueueAsyncCallEventAddr() {
//    return &doApplicationQueueAsyncCallEventProc;
// }
import "C"

import (
	"unsafe"
)

//export doApplicationQueueAsyncCallEventProc
func doApplicationQueueAsyncCallEventProc(f unsafe.Pointer) unsafe.Pointer {
	applicationQueueAsyncCallProc(uintptr(f))
	return nullptr
}

var (
	applicationQueueAsyncCallEvent = uintptr(C.doGetApplicationQueueAsyncCallEventAddr())
)
