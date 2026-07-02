//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
// Licensed under Apache License Version 2.0, January 2004
// https://www.apache.org/licenses/LICENSE-2.0
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// ColorChooserDialog is a representation of GTK's GtkColorChooserDialog.
type ColorChooserDialog struct {
	Dialog
}

func AsColorChooserDialog(ptr unsafe.Pointer) *ColorChooserDialog {
	if ptr == nil {
		return nil
	}
	m := new(ColorChooserDialog)
	m.instance = ptr
	return m
}

// NewColorChooserDialog is a wrapper around gtk_color_chooser_dialog_new().
func NewColorChooserDialog(title string, parent IWindow) *ColorChooserDialog {
	cstr := CStr(title)
	var p uintptr
	if parent != nil {
		p = parent.Instance()
	}
	r := gtk3.SysCall("gtk_color_chooser_dialog_new", cstr, p)
	if r == 0 {
		return nil
	}
	return AsColorChooserDialog(unsafe.Pointer(r))
}

// GetUseAlpha is a wrapper around gtk_color_chooser_get_use_alpha().
func (m *ColorChooserDialog) GetUseAlpha() bool {
	r := gtk3.SysCall("gtk_color_chooser_get_use_alpha", m.Instance())
	return ToGoBool(r)
}

// SetUseAlpha is a wrapper around gtk_color_chooser_set_use_alpha().
func (m *ColorChooserDialog) SetUseAlpha(useAlpha bool) {
	gtk3.SysCall("gtk_color_chooser_set_use_alpha", m.Instance(), ToCBool(useAlpha))
}
