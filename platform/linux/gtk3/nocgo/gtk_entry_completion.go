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

// EntryCompletion is a representation of GTK's GtkEntryCompletion.
type EntryCompletion struct {
	Object
}

func AsEntryCompletion(ptr unsafe.Pointer) IEntryCompletion {
	if ptr == nil {
		return nil
	}
	m := new(EntryCompletion)
	m.instance = ptr
	return m
}

// NewEntryCompletion is a wrapper around gtk_entry_completion_new().
func NewEntryCompletion() IEntryCompletion {
	r := gtk3.SysCall("gtk_entry_completion_new")
	if r == 0 {
		return nil
	}
	return AsEntryCompletion(unsafe.Pointer(r))
}

// SetTextColumn is a wrapper around gtk_entry_completion_set_text_column().
func (m *EntryCompletion) SetTextColumn(column int) {
	gtk3.SysCall("gtk_entry_completion_set_text_column", m.Instance(), uintptr(column))
}

// GetTextColumn is a wrapper around gtk_entry_completion_get_text_column().
func (m *EntryCompletion) GetTextColumn() int {
	r := gtk3.SysCall("gtk_entry_completion_get_text_column", m.Instance())
	return int(r)
}

// SetMinimumKeyLength is a wrapper around gtk_entry_completion_set_minimum_key_length().
func (m *EntryCompletion) SetMinimumKeyLength(length int) {
	gtk3.SysCall("gtk_entry_completion_set_minimum_key_length", m.Instance(), uintptr(length))
}

// GetMinimumKeyLength is a wrapper around gtk_entry_completion_get_minimum_key_length().
func (m *EntryCompletion) GetMinimumKeyLength() int {
	r := gtk3.SysCall("gtk_entry_completion_get_minimum_key_length", m.Instance())
	return int(r)
}

// SetInlineCompletion is a wrapper around gtk_entry_completion_set_inline_completion().
func (m *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	gtk3.SysCall("gtk_entry_completion_set_inline_completion", m.Instance(), ToCBool(inlineCompletion))
}

// GetInlineCompletion is a wrapper around gtk_entry_completion_get_inline_completion().
func (m *EntryCompletion) GetInlineCompletion() bool {
	r := gtk3.SysCall("gtk_entry_completion_get_inline_completion", m.Instance())
	return ToGoBool(r)
}

// SetPopupCompletion is a wrapper around gtk_entry_completion_set_popup_completion().
func (m *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	gtk3.SysCall("gtk_entry_completion_set_popup_completion", m.Instance(), ToCBool(popupCompletion))
}

// GetPopupCompletion is a wrapper around gtk_entry_completion_get_popup_completion().
func (m *EntryCompletion) GetPopupCompletion() bool {
	r := gtk3.SysCall("gtk_entry_completion_get_popup_completion", m.Instance())
	return ToGoBool(r)
}

// SetPopupSetWidth is a wrapper around gtk_entry_completion_set_popup_set_width().
func (m *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	gtk3.SysCall("gtk_entry_completion_set_popup_set_width", m.Instance(), ToCBool(popupSetWidth))
}

// GetPopupSetWidth is a wrapper around gtk_entry_completion_get_popup_set_width().
func (m *EntryCompletion) GetPopupSetWidth() bool {
	r := gtk3.SysCall("gtk_entry_completion_get_popup_set_width", m.Instance())
	return ToGoBool(r)
}
