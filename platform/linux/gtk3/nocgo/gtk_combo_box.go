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

// ComboBoxText is a representation of GTK's GtkComboBoxText.
type ComboBoxText struct {
	Widget
}

func AsComboBoxText(ptr unsafe.Pointer) *ComboBoxText {
	if ptr == nil {
		return nil
	}
	m := new(ComboBoxText)
	m.instance = ptr
	return m
}

// NewComboBoxText is a wrapper around gtk_combo_box_text_new().
func NewComboBoxText() *ComboBoxText {
	r := gtk3.SysCall("gtk_combo_box_text_new")
	if r == 0 {
		return nil
	}
	return AsComboBoxText(unsafe.Pointer(r))
}

// NewComboBoxTextWithEntry is a wrapper around gtk_combo_box_text_new_with_entry().
func NewComboBoxTextWithEntry() *ComboBoxText {
	r := gtk3.SysCall("gtk_combo_box_text_new_with_entry")
	if r == 0 {
		return nil
	}
	return AsComboBoxText(unsafe.Pointer(r))
}

// Append is a wrapper around gtk_combo_box_text_append().
func (m *ComboBoxText) Append(id string, text string) {
	gtk3.SysCall("gtk_combo_box_text_append", m.Instance(), CStr(id), CStr(text))
}

// AppendText is a wrapper around gtk_combo_box_text_append_text().
func (m *ComboBoxText) AppendText(text string) {
	gtk3.SysCall("gtk_combo_box_text_append_text", m.Instance(), CStr(text))
}

// GetActiveText is a wrapper around gtk_combo_box_text_get_active_text().
func (m *ComboBoxText) GetActiveText() string {
	r := gtk3.SysCall("gtk_combo_box_text_get_active_text", m.Instance())
	if r == 0 {
		return ""
	}
	s := GoStr(r)
	glib2_0.SysCall("g_free", r)
	return s
}

// RemoveAll is a wrapper around gtk_combo_box_text_remove_all().
func (m *ComboBoxText) RemoveAll() {
	gtk3.SysCall("gtk_combo_box_text_remove_all", m.Instance())
}

// GetActive is a wrapper around gtk_combo_box_get_active().
func (m *ComboBoxText) GetActive() int {
	r := gtk3.SysCall("gtk_combo_box_get_active", m.Instance())
	return int(r)
}

// SetActive is a wrapper around gtk_combo_box_set_active().
func (m *ComboBoxText) SetActive(index int) {
	gtk3.SysCall("gtk_combo_box_set_active", m.Instance(), uintptr(index))
}

func (m *ComboBoxText) SetOnChanged(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnChanged, callback.C_trampoline_2_void, fn, 0)
}

// Insert is a wrapper around gtk_combo_box_text_insert().
func (m *ComboBoxText) Insert(position int, id string, text string) {
	gtk3.SysCall("gtk_combo_box_text_insert", m.Instance(), uintptr(position), CStr(id), CStr(text))
}

// InsertText is a wrapper around gtk_combo_box_text_insert_text().
func (m *ComboBoxText) InsertText(position int, text string) {
	gtk3.SysCall("gtk_combo_box_text_insert_text", m.Instance(), uintptr(position), CStr(text))
}

// Prepend is a wrapper around gtk_combo_box_text_prepend().
func (m *ComboBoxText) Prepend(id string, text string) {
	gtk3.SysCall("gtk_combo_box_text_prepend", m.Instance(), CStr(id), CStr(text))
}

// PrependText is a wrapper around gtk_combo_box_text_prepend_text().
func (m *ComboBoxText) PrependText(text string) {
	gtk3.SysCall("gtk_combo_box_text_prepend_text", m.Instance(), CStr(text))
}

// Remove is a wrapper around gtk_combo_box_text_remove().
func (m *ComboBoxText) Remove(position int) {
	gtk3.SysCall("gtk_combo_box_text_remove", m.Instance(), uintptr(position))
}
