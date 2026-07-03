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

type Scale struct {
	Range
}

func AsScale(ptr unsafe.Pointer) *Scale {
	if ptr == nil {
		return nil
	}
	m := new(Scale)
	m.instance = ptr
	return m
}

func NewHScale(adjustment IAdjustment) *Scale {
	var adj uintptr
	if adjustment != nil {
		adj = adjustment.Instance()
	}
	r := gtk3.SysCall("gtk_scale_new", uintptr(ORIENTATION_HORIZONTAL), adj)
	if r == 0 {
		return nil
	}
	return AsScale(unsafe.Pointer(r))
}

func NewVScale(adjustment IAdjustment) *Scale {
	var adj uintptr
	if adjustment != nil {
		adj = adjustment.Instance()
	}
	r := gtk3.SysCall("gtk_scale_new", uintptr(ORIENTATION_VERTICAL), adj)
	if r == 0 {
		return nil
	}
	return AsScale(unsafe.Pointer(r))
}

func (m *Scale) SetDigits(digits int) {
	gtk3.SysCall("gtk_scale_set_digits", m.Instance(), uintptr(digits))
}

func (m *Scale) GetDigits() int {
	r := gtk3.SysCall("gtk_scale_get_digits", m.Instance())
	return int(r)
}

func (m *Scale) SetDrawValue(drawValue bool) {
	gtk3.SysCall("gtk_scale_set_draw_value", m.Instance(), ToCBool(drawValue))
}

func (m *Scale) GetDrawValue() bool {
	r := gtk3.SysCall("gtk_scale_get_draw_value", m.Instance())
	return ToGoBool(r)
}

func (m *Scale) SetValuePos(pos PositionType) {
	gtk3.SysCall("gtk_scale_set_value_pos", m.Instance(), uintptr(pos))
}

func (m *Scale) GetValuePos() PositionType {
	r := gtk3.SysCall("gtk_scale_get_value_pos", m.Instance())
	return PositionType(r)
}
