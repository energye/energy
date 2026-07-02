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

// CheckButton is a representation of GTK's GtkCheckButton.
type CheckButton struct {
	Button
}

func AsCheckButton(ptr unsafe.Pointer) ICheckButton {
	if ptr == nil {
		return nil
	}
	m := new(CheckButton)
	m.instance = ptr
	return m
}

// NewCheckButton is a wrapper around gtk_check_button_new().
func NewCheckButton() ICheckButton {
	r := gtk3.SysCall("gtk_check_button_new")
	if r == 0 {
		return nil
	}
	return AsCheckButton(unsafe.Pointer(r))
}

// GetActive is a wrapper around gtk_toggle_button_get_active().
func (m *CheckButton) GetActive() bool {
	r := gtk3.SysCall("gtk_toggle_button_get_active", m.Instance())
	return ToGoBool(r)
}

// SetActive is a wrapper around gtk_toggle_button_set_active().
func (m *CheckButton) SetActive(isActive bool) {
	gtk3.SysCall("gtk_toggle_button_set_active", m.Instance(), ToCBool(isActive))
}
