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

type Revealer struct {
	Bin
}

func AsRevealer(ptr unsafe.Pointer) *Revealer {
	if ptr == nil {
		return nil
	}
	m := new(Revealer)
	m.instance = ptr
	return m
}

func NewRevealer() *Revealer {
	r := gtk3.SysCall("gtk_revealer_new")
	if r == 0 {
		return nil
	}
	return AsRevealer(unsafe.Pointer(r))
}

func (m *Revealer) SetRevealChild(reveal bool) {
	gtk3.SysCall("gtk_revealer_set_reveal_child", m.Instance(), ToCBool(reveal))
}

func (m *Revealer) GetRevealChild() bool {
	r := gtk3.SysCall("gtk_revealer_get_reveal_child", m.Instance())
	return ToGoBool(r)
}

func (m *Revealer) SetTransitionDuration(duration uint) {
	gtk3.SysCall("gtk_revealer_set_transition_duration", m.Instance(), uintptr(duration))
}

func (m *Revealer) GetTransitionDuration() uint {
	r := gtk3.SysCall("gtk_revealer_get_transition_duration", m.Instance())
	return uint(r)
}

func (m *Revealer) SetTransitionType(t RevealerTransitionType) {
	gtk3.SysCall("gtk_revealer_set_transition_type", m.Instance(), uintptr(t))
}

func (m *Revealer) GetTransitionType() RevealerTransitionType {
	r := gtk3.SysCall("gtk_revealer_get_transition_type", m.Instance())
	return RevealerTransitionType(r)
}
