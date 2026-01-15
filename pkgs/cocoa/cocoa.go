//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package cocoa

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "cocoa.h"

extern void onRunOnMainThread(long id);

*/
import "C"
import (
	"sync"
	"unsafe"
)

//export onRunOnMainThread
func onRunOnMainThread(id C.long) {
	doRunOnMainThread(int(id))
}

type runOnMainThreadFn func()

var (
	callbackFuncList     = make(map[int]runOnMainThreadFn)
	callbackFuncListLock = sync.Mutex{}
	isRROMTC             bool
)

func doRunOnMainThread(id int) {
	fn, ok := callbackFuncList[id]
	if ok {
		delete(callbackFuncList, id)
		fn()
	}
}

func RegisterRunOnMainThreadCallback() {
	if isRROMTC {
		return
	}
	isRROMTC = true
	C.RegisterRunOnMainThreadCallback(C.RunOnMainThreadCallback(C.onRunOnMainThread))
}

func RunOnMainThread(fn runOnMainThreadFn) {
	callbackFuncListLock.Lock()
	defer callbackFuncListLock.Unlock()
	id := int(uintptr(unsafe.Pointer(&fn)))
	callbackFuncList[id] = fn
	C.ExecuteRunOnMainThread(C.long(id))
}
