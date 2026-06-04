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

type Context struct {
	instance unsafe.Pointer
}

func AsContext(ptr unsafe.Pointer) IContext {
	if ptr == nil {
		return nil
	}
	m := new(Context)
	m.instance = ptr
	return m
}

func (m *Context) Instance() uintptr {
	return uintptr(m.instance)
}

// Status is a wrapper around cairo_status().
func (m *Context) Status() Status {
	r := cairo.SysCall("cairo_status", m.Instance())
	return Status(r)
}
