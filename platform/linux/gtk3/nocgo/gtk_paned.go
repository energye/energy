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

type Paned struct {
	Container
}

func AsPaned(ptr unsafe.Pointer) *Paned {
	if ptr == nil {
		return nil
	}
	m := new(Paned)
	m.instance = ptr
	return m
}

func NewPaned(orientation Orientation) *Paned {
	r := gtk3.SysCall("gtk_paned_new", uintptr(orientation))
	if r == 0 {
		return nil
	}
	return AsPaned(unsafe.Pointer(r))
}

func (m *Paned) Add1(child IWidget) {
	gtk3.SysCall("gtk_paned_add1", m.Instance(), child.Instance())
}

func (m *Paned) Add2(child IWidget) {
	gtk3.SysCall("gtk_paned_add2", m.Instance(), child.Instance())
}

func (m *Paned) SetPosition(position int) {
	gtk3.SysCall("gtk_paned_set_position", m.Instance(), uintptr(position))
}

func (m *Paned) GetPosition() int {
	r := gtk3.SysCall("gtk_paned_get_position", m.Instance())
	return int(r)
}

func (m *Paned) SetWideHandle(wide bool) {
	gtk3.SysCall("gtk_paned_set_wide_handle", m.Instance(), ToCBool(wide))
}

func (m *Paned) GetWideHandle() bool {
	r := gtk3.SysCall("gtk_paned_get_wide_handle", m.Instance())
	return ToGoBool(r)
}
