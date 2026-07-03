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

type Popover struct {
	Bin
}

func AsPopover(ptr unsafe.Pointer) *Popover {
	if ptr == nil {
		return nil
	}
	m := new(Popover)
	m.instance = ptr
	return m
}

func NewPopover() *Popover {
	r := gtk3.SysCall("gtk_popover_new", uintptr(0))
	if r == 0 {
		return nil
	}
	return AsPopover(unsafe.Pointer(r))
}

func (m *Popover) SetRelativeTo(widget IWidget) {
	gtk3.SysCall("gtk_popover_set_relative_to", m.Instance(), widget.Instance())
}

func (m *Popover) GetRelativeTo() IWidget {
	r := gtk3.SysCall("gtk_popover_get_relative_to", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

func (m *Popover) SetPosition(position PositionType) {
	gtk3.SysCall("gtk_popover_set_position", m.Instance(), uintptr(position))
}

func (m *Popover) GetPosition() PositionType {
	r := gtk3.SysCall("gtk_popover_get_position", m.Instance())
	return PositionType(r)
}

func (m *Popover) SetModal(modal bool) {
	gtk3.SysCall("gtk_popover_set_modal", m.Instance(), ToCBool(modal))
}

func (m *Popover) GetModal() bool {
	r := gtk3.SysCall("gtk_popover_get_modal", m.Instance())
	return ToGoBool(r)
}

func (m *Popover) Popdown() {
	gtk3.SysCall("gtk_popover_popdown", m.Instance())
}

func (m *Popover) Popup() {
	gtk3.SysCall("gtk_popover_popup", m.Instance())
}
