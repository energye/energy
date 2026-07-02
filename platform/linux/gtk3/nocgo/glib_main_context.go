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
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// MainContext is a representation of GLib's GMainContext.
type MainContext struct {
	instance unsafe.Pointer
}

func AsMainContext(ptr unsafe.Pointer) *MainContext {
	if ptr == nil {
		return nil
	}
	return &MainContext{instance: ptr}
}

// Native returns the underlying GMainContext pointer as uintptr.
func (v *MainContext) Native() uintptr {
	if v == nil {
		return 0
	}
	return uintptr(v.instance)
}

// MainContextDefault is a wrapper around g_main_context_default().
func MainContextDefault() *MainContext {
	r := glib2_0.SysCall("g_main_context_default")
	if r == 0 {
		return nil
	}
	return AsMainContext(unsafe.Pointer(r))
}

// Iteration is a wrapper around g_main_context_iteration().
func (v *MainContext) Iteration(mayBlock bool) bool {
	r := glib2_0.SysCall("g_main_context_iteration", uintptr(v.instance), ToCBool(mayBlock))
	return ToGoBool(r)
}

// Pending is a wrapper around g_main_context_pending().
func (v *MainContext) Pending() bool {
	r := glib2_0.SysCall("g_main_context_pending", uintptr(v.instance))
	return ToGoBool(r)
}

// MainDepth is a wrapper around g_main_depth().
func MainDepth() int {
	r := glib2_0.SysCall("g_main_depth")
	return int(r)
}

// FindSourceById is a wrapper around g_main_context_find_source_by_id().
func (v *MainContext) FindSourceById(hdlSrc SourceHandle) *Source {
	r := glib2_0.SysCall("g_main_context_find_source_by_id", uintptr(v.instance), uintptr(hdlSrc))
	if r == 0 {
		return nil
	}
	return AsSource(unsafe.Pointer(r))
}

// Acquire is a wrapper around g_main_context_acquire().
func (v *MainContext) Acquire() bool {
	r := glib2_0.SysCall("g_main_context_acquire", uintptr(v.instance))
	return ToGoBool(r)
}

// Release is a wrapper around g_main_context_release().
func (v *MainContext) Release() {
	glib2_0.SysCall("g_main_context_release", uintptr(v.instance))
}

// IsOwner is a wrapper around g_main_context_is_owner().
func (v *MainContext) IsOwner() bool {
	r := glib2_0.SysCall("g_main_context_is_owner", uintptr(v.instance))
	return ToGoBool(r)
}
