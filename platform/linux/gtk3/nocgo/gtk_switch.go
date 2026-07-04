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
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Switch is a representation of GTK's GtkSwitch.
type Switch struct {
	Widget
}

func AsSwitch(ptr unsafe.Pointer) *Switch {
	if ptr == nil {
		return nil
	}
	m := new(Switch)
	m.instance = ptr
	return m
}

func NewSwitch() *Switch {
	r := gtk3.SysCall("gtk_switch_new")
	if r == 0 {
		return nil
	}
	return AsSwitch(unsafe.Pointer(r))
}

func (m *Switch) GetActive() bool {
	r := gtk3.SysCall("gtk_switch_get_active", m.Instance())
	return ToGoBool(r)
}

func (m *Switch) SetActive(isActive bool) {
	gtk3.SysCall("gtk_switch_set_active", m.Instance(), ToCBool(isActive))
}

func (m *Switch) GetState() bool {
	r := gtk3.SysCall("gtk_switch_get_state", m.Instance())
	return ToGoBool(r)
}

func (m *Switch) SetState(state bool) {
	gtk3.SysCall("gtk_switch_set_state", m.Instance(), ToCBool(state))
}

func (m *Switch) SetOnActiveNotify(fn TNotifyActiveEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnNotifyActive, callback.C_trampoline_3_void, fn, 0)
}
