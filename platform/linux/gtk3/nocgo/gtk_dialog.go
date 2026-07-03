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

// Dialog is a representation of GTK's GtkDialog.
type Dialog struct {
	Window
}

func AsDialog(ptr unsafe.Pointer) *Dialog {
	if ptr == nil {
		return nil
	}
	m := new(Dialog)
	m.instance = ptr
	return m
}

// NewDialog is a wrapper around gtk_dialog_new().
func NewDialog() *Dialog {
	r := gtk3.SysCall("gtk_dialog_new")
	if r == 0 {
		return nil
	}
	return AsDialog(unsafe.Pointer(r))
}

// Run is a wrapper around gtk_dialog_run().
func (m *Dialog) Run() int {
	r := gtk3.SysCall("gtk_dialog_run", m.Instance())
	return int(r)
}

// Response is a wrapper around gtk_dialog_response().
func (m *Dialog) Response(responseId int) {
	gtk3.SysCall("gtk_dialog_response", m.Instance(), uintptr(responseId))
}

// AddButton is a wrapper around gtk_dialog_add_button().
func (m *Dialog) AddButton(buttonText string, responseId int) IButton {
	r := gtk3.SysCall("gtk_dialog_add_button", m.Instance(), CStr(buttonText), uintptr(responseId))
	if r == 0 {
		return nil
	}
	return AsButton(unsafe.Pointer(r))
}

// SetDefaultResponse is a wrapper around gtk_dialog_set_default_response().
func (m *Dialog) SetDefaultResponse(responseId int) {
	gtk3.SysCall("gtk_dialog_set_default_response", m.Instance(), uintptr(responseId))
}

// GetContentArea is a wrapper around gtk_dialog_get_content_area().
func (m *Dialog) GetContentArea() IBox {
	r := gtk3.SysCall("gtk_dialog_get_content_area", m.Instance())
	if r == 0 {
		return nil
	}
	box := new(Box)
	box.instance = unsafe.Pointer(r)
	return box
}
