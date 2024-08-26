//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

package cef

// #cgo darwin CFLAGS: -mmacosx-version-min=10.12
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.12
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
	return nil
}

var (
	applicationQueueAsyncCallEvent = uintptr(C.doGetApplicationQueueAsyncCallEventAddr())
)
