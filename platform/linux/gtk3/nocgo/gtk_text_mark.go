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

// TextMark is a representation of GTK's GtkTextMark.
type TextMark struct {
	Object
}

func AsTextMark(ptr unsafe.Pointer) *TextMark {
	if ptr == nil {
		return nil
	}
	m := new(TextMark)
	m.instance = ptr
	return m
}

// SetVisible is a wrapper around gtk_text_mark_set_visible().
func (m *TextMark) SetVisible(setting bool) {
	gtk3.SysCall("gtk_text_mark_set_visible", m.Instance(), ToCBool(setting))
}

// GetVisible is a wrapper around gtk_text_mark_get_visible().
func (m *TextMark) GetVisible() bool {
	r := gtk3.SysCall("gtk_text_mark_get_visible", m.Instance())
	return ToGoBool(r)
}

// GetDeleted is a wrapper around gtk_text_mark_get_deleted().
func (m *TextMark) GetDeleted() bool {
	r := gtk3.SysCall("gtk_text_mark_get_deleted", m.Instance())
	return ToGoBool(r)
}

// GetName is a wrapper around gtk_text_mark_get_name().
func (m *TextMark) GetName() string {
	r := gtk3.SysCall("gtk_text_mark_get_name", m.Instance())
	return GoStr(r)
}

// GetBuffer is a wrapper around gtk_text_mark_get_buffer().
func (m *TextMark) GetBuffer() ITextBuffer {
	r := gtk3.SysCall("gtk_text_mark_get_buffer", m.Instance())
	if r == 0 {
		return nil
	}
	return AsTextBuffer(unsafe.Pointer(r))
}
