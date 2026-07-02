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

// TextTag is a representation of GTK's GtkTextTag.
type TextTag struct {
	Object
}

func AsTextTag(ptr unsafe.Pointer) *TextTag {
	if ptr == nil {
		return nil
	}
	m := new(TextTag)
	m.instance = ptr
	return m
}

// GetPriority is a wrapper around gtk_text_tag_get_priority().
func (m *TextTag) GetPriority() int {
	r := gtk3.SysCall("gtk_text_tag_get_priority", m.Instance())
	return int(r)
}

// SetPriority is a wrapper around gtk_text_tag_set_priority().
func (m *TextTag) SetPriority(priority int) {
	gtk3.SysCall("gtk_text_tag_set_priority", m.Instance(), uintptr(priority))
}

// TextTagTable is a representation of GTK's GtkTextTagTable.
type TextTagTable struct {
	Object
}

func AsTextTagTable(ptr unsafe.Pointer) *TextTagTable {
	if ptr == nil {
		return nil
	}
	m := new(TextTagTable)
	m.instance = ptr
	return m
}

// NewTextTagTable is a wrapper around gtk_text_tag_table_new().
func NewTextTagTable() *TextTagTable {
	r := gtk3.SysCall("gtk_text_tag_table_new")
	if r == 0 {
		return nil
	}
	return AsTextTagTable(unsafe.Pointer(r))
}

// Add is a wrapper around gtk_text_tag_table_add().
func (m *TextTagTable) Add(tag ITextTag) bool {
	r := gtk3.SysCall("gtk_text_tag_table_add", m.Instance(), tag.(*TextTag).Instance())
	return ToGoBool(r)
}

// Lookup is a wrapper around gtk_text_tag_table_lookup().
func (m *TextTagTable) Lookup(name string) ITextTag {
	r := gtk3.SysCall("gtk_text_tag_table_lookup", m.Instance(), CStr(name))
	if r == 0 {
		return nil
	}
	return AsTextTag(unsafe.Pointer(r))
}

// Remove is a wrapper around gtk_text_tag_table_remove().
func (m *TextTagTable) Remove(tag ITextTag) {
	gtk3.SysCall("gtk_text_tag_table_remove", m.Instance(), tag.(*TextTag).Instance())
}
