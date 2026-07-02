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

// FileChooserDialog is a representation of GTK's GtkFileChooserDialog.
type FileChooserDialog struct {
	Dialog
}

func AsFileChooserDialog(ptr unsafe.Pointer) *FileChooserDialog {
	if ptr == nil {
		return nil
	}
	m := new(FileChooserDialog)
	m.instance = ptr
	return m
}

// NewFileChooserDialog is a wrapper around gtk_file_chooser_dialog_new().
func NewFileChooserDialog(title string, parent IWindow, action FileChooserAction) *FileChooserDialog {
	cstr := CStr(title)
	var p uintptr
	if parent != nil {
		p = parent.Instance()
	}
	// gtk_file_chooser_dialog_new is variadic, pass NULL-terminated button list
	r := gtk3.SysCall("gtk_file_chooser_dialog_new", cstr, p, uintptr(action), 0)
	if r == 0 {
		return nil
	}
	return AsFileChooserDialog(unsafe.Pointer(r))
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (m *FileChooserDialog) GetFilename() string {
	r := gtk3.SysCall("gtk_file_chooser_get_filename", m.Instance())
	if r == 0 {
		return ""
	}
	s := GoStr(r)
	glib2_0.SysCall("g_free", r)
	return s
}

// SetFilename is a wrapper around gtk_file_chooser_set_filename().
func (m *FileChooserDialog) SetFilename(filename string) bool {
	r := gtk3.SysCall("gtk_file_chooser_set_filename", m.Instance(), CStr(filename))
	return ToGoBool(r)
}

// SetCurrentFolder is a wrapper around gtk_file_chooser_set_current_folder().
func (m *FileChooserDialog) SetCurrentFolder(folder string) bool {
	r := gtk3.SysCall("gtk_file_chooser_set_current_folder", m.Instance(), CStr(folder))
	return ToGoBool(r)
}

// SetAction is a wrapper around gtk_file_chooser_set_action().
func (m *FileChooserDialog) SetAction(action FileChooserAction) {
	gtk3.SysCall("gtk_file_chooser_set_action", m.Instance(), uintptr(action))
}

// GetAction is a wrapper around gtk_file_chooser_get_action().
func (m *FileChooserDialog) GetAction() FileChooserAction {
	r := gtk3.SysCall("gtk_file_chooser_get_action", m.Instance())
	return FileChooserAction(r)
}

// SetSelectMultiple is a wrapper around gtk_file_chooser_set_select_multiple().
func (m *FileChooserDialog) SetSelectMultiple(selectMultiple bool) {
	gtk3.SysCall("gtk_file_chooser_set_select_multiple", m.Instance(), ToCBool(selectMultiple))
}

// GetSelectMultiple is a wrapper around gtk_file_chooser_get_select_multiple().
func (m *FileChooserDialog) GetSelectMultiple() bool {
	r := gtk3.SysCall("gtk_file_chooser_get_select_multiple", m.Instance())
	return ToGoBool(r)
}
