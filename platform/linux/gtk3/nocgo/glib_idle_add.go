//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"sync"
	"unsafe"

	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/ebitengine/purego"
)

type idleFunc func() bool

var (
	idleRegistry sync.Map
	// sourceFuncPtr is the C-callable function pointer for GSourceFunc.
	sourceFuncPtr uintptr
	// removeFuncPtr is the C-callable function pointer for GDestroyNotify.
	removeFuncPtr uintptr
)

func init() {
	// Create C-callable callbacks via purego.NewCallback.
	// GSourceFunc signature: gboolean func(gpointer data)
	sourceFuncPtr = purego.NewCallback(func(data uintptr) uintptr {
		fn := idleGet(data)
		if fn != nil {
			if fn() {
				return 1
			}
		}
		return 0
	})
	// GDestroyNotify signature: void func(gpointer data)
	removeFuncPtr = purego.NewCallback(func(data uintptr) {
		idleDelete(data)
	})
}

func idlePut(callback idleFunc) uintptr {
	id := uintptr(unsafe.Pointer(&callback))
	idleRegistry.Store(id, callback)
	return id
}

func idleGet(id uintptr) idleFunc {
	if fn, ok := idleRegistry.Load(id); ok {
		return fn.(idleFunc)
	}
	return nil
}

func idleDelete(id uintptr) {
	idleRegistry.Delete(id)
}

// IdleAdd adds an idle source to the default main event loop context with the
// DefaultIdle priority.
func IdleAdd(fn idleFunc) SourceHandle {
	return idleAdd(PRIORITY_DEFAULT_IDLE, fn)
}

// IdleAddPriority adds an idle source to the default main event loop context
// with the given priority.
func IdleAddPriority(priority Priority, fn idleFunc) SourceHandle {
	return idleAdd(priority, fn)
}

func idleAdd(priority Priority, fn idleFunc) SourceHandle {
	id := idlePut(fn)
	h := glib2_0.SysCall("g_idle_add_full",
		uintptr(priority), sourceFuncPtr, id, removeFuncPtr)
	return SourceHandle(h)
}
