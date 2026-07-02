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

// EntryBuffer is a representation of GTK's GtkEntryBuffer.
type EntryBuffer struct {
	Object
}

func AsEntryBuffer(ptr unsafe.Pointer) IEntryBuffer {
	if ptr == nil {
		return nil
	}
	m := new(EntryBuffer)
	m.instance = ptr
	return m
}

// NewEntryBuffer is a wrapper around gtk_entry_buffer_new().
func NewEntryBuffer(initialChars string) IEntryBuffer {
	// -1 means use strlen, represented as max uintptr
	nInit := ^uintptr(0)
	var r uintptr
	if initialChars == "" {
		r = gtk3.SysCall("gtk_entry_buffer_new", 0, nInit)
	} else {
		r = gtk3.SysCall("gtk_entry_buffer_new", CStr(initialChars), nInit)
	}
	if r == 0 {
		return nil
	}
	return AsEntryBuffer(unsafe.Pointer(r))
}

// GetText is a wrapper around gtk_entry_buffer_get_text().
func (m *EntryBuffer) GetText() (string, error) {
	r := gtk3.SysCall("gtk_entry_buffer_get_text", m.Instance())
	return GoStr(r), nil
}

// SetText is a wrapper around gtk_entry_buffer_set_text().
func (m *EntryBuffer) SetText(text string) {
	gtk3.SysCall("gtk_entry_buffer_set_text", m.Instance(), CStr(text), ^uintptr(0))
}

// GetBytes is a wrapper around gtk_entry_buffer_get_bytes().
func (m *EntryBuffer) GetBytes() uint {
	r := gtk3.SysCall("gtk_entry_buffer_get_bytes", m.Instance())
	return uint(r)
}

// GetLength is a wrapper around gtk_entry_buffer_get_length().
func (m *EntryBuffer) GetLength() uint {
	r := gtk3.SysCall("gtk_entry_buffer_get_length", m.Instance())
	return uint(r)
}

// GetMaxLength is a wrapper around gtk_entry_buffer_get_max_length().
func (m *EntryBuffer) GetMaxLength() int {
	r := gtk3.SysCall("gtk_entry_buffer_get_max_length", m.Instance())
	return int(r)
}

// SetMaxLength is a wrapper around gtk_entry_buffer_set_max_length().
func (m *EntryBuffer) SetMaxLength(maxLength int) {
	gtk3.SysCall("gtk_entry_buffer_set_max_length", m.Instance(), uintptr(maxLength))
}

// InsertText is a wrapper around gtk_entry_buffer_insert_text().
func (m *EntryBuffer) InsertText(position uint, text string) uint {
	r := gtk3.SysCall("gtk_entry_buffer_insert_text", m.Instance(), uintptr(position), CStr(text), uintptr(len(text)))
	return uint(r)
}

// DeleteText is a wrapper around gtk_entry_buffer_delete_text().
func (m *EntryBuffer) DeleteText(position uint, nChars int) uint {
	r := gtk3.SysCall("gtk_entry_buffer_delete_text", m.Instance(), uintptr(position), uintptr(nChars))
	return uint(r)
}

// EmitDeletedText is a wrapper around gtk_entry_buffer_emit_deleted_text().
func (m *EntryBuffer) EmitDeletedText(pos, nChars uint) {
	gtk3.SysCall("gtk_entry_buffer_emit_deleted_text", m.Instance(), uintptr(pos), uintptr(nChars))
}

// EmitInsertedText is a wrapper around gtk_entry_buffer_emit_inserted_text().
func (m *EntryBuffer) EmitInsertedText(pos uint, text string) {
	gtk3.SysCall("gtk_entry_buffer_emit_inserted_text", m.Instance(), uintptr(pos), CStr(text), uintptr(len(text)))
}
