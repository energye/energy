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

import "unsafe"

// Source is a representation of GLib's GSource.
type Source struct {
	instance unsafe.Pointer
}

func AsSource(ptr unsafe.Pointer) *Source {
	if ptr == nil {
		return nil
	}
	return &Source{instance: ptr}
}

// Native returns the underlying GSource pointer as uintptr.
func (v *Source) Native() uintptr {
	if v == nil {
		return 0
	}
	return uintptr(v.instance)
}

// MainCurrentSource is a wrapper around g_main_current_source().
func MainCurrentSource() *Source {
	r := glib2_0.SysCall("g_main_current_source")
	if r == 0 {
		return nil
	}
	return AsSource(unsafe.Pointer(r))
}
