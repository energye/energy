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

// MessageDialog is a representation of GTK's GtkMessageDialog.
type MessageDialog struct {
	Dialog
}

func AsMessageDialog(ptr unsafe.Pointer) IMessageDialog {
	if ptr == nil {
		return nil
	}
	m := new(MessageDialog)
	m.instance = ptr
	return m
}

// MessageDialogNew is a wrapper around gtk_message_dialog_new().
func MessageDialogNew(parent IWindow, flags DialogFlags, mType MessageType, buttons ButtonsType, message string) IMessageDialog {
	cstr := CStr(message)
	var p uintptr
	if parent != nil {
		p = parent.Instance()
	}
	r := gtk3.SysCall("gtk_message_dialog_new", p,
		uintptr(flags), uintptr(mType), uintptr(buttons), cstr)
	if r == 0 {
		return nil
	}
	return AsMessageDialog(unsafe.Pointer(r))
}

// FormatSecondaryText is a wrapper around gtk_message_dialog_format_secondary_text().
func (m *MessageDialog) FormatSecondaryText(message string) {
	gtk3.SysCall("gtk_message_dialog_format_secondary_text", m.Instance(), CStr(message))
}
