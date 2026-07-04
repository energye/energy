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

// FontChooserDialog is a representation of GTK's GtkFontChooserDialog.
type FontChooserDialog struct {
	Dialog
}

func AsFontChooserDialog(ptr unsafe.Pointer) *FontChooserDialog {
	if ptr == nil {
		return nil
	}
	m := new(FontChooserDialog)
	m.instance = ptr
	return m
}

// NewFontChooserDialog is a wrapper around gtk_font_chooser_dialog_new().
func NewFontChooserDialog(title string, parent IWindow) *FontChooserDialog {
	cstr := CStr(title)
	var p uintptr
	if parent != nil {
		p = parent.Instance()
	}
	r := gtk3.SysCall("gtk_font_chooser_dialog_new", cstr, p)
	if r == 0 {
		return nil
	}
	return AsFontChooserDialog(unsafe.Pointer(r))
}

// GetFont is a wrapper around gtk_font_chooser_get_font().
func (m *FontChooserDialog) GetFont() string {
	r := gtk3.SysCall("gtk_font_chooser_get_font", m.Instance())
	if r == 0 {
		return ""
	}
	s := GoStr(r)
	glib2_0.SysCall("g_free", r)
	return s
}

// SetFont is a wrapper around gtk_font_chooser_set_font().
func (m *FontChooserDialog) SetFont(font string) {
	gtk3.SysCall("gtk_font_chooser_set_font", m.Instance(), CStr(font))
}

// GetPreviewText is a wrapper around gtk_font_chooser_get_preview_text().
func (m *FontChooserDialog) GetPreviewText() string {
	r := gtk3.SysCall("gtk_font_chooser_get_preview_text", m.Instance())
	if r == 0 {
		return ""
	}
	s := GoStr(r)
	glib2_0.SysCall("g_free", r)
	return s
}

// SetPreviewText is a wrapper around gtk_font_chooser_set_preview_text().
func (m *FontChooserDialog) SetPreviewText(text string) {
	gtk3.SysCall("gtk_font_chooser_set_preview_text", m.Instance(), CStr(text))
}

// GetFontFamily is a wrapper around gtk_font_chooser_get_font_family().
func (m *FontChooserDialog) GetFontFamily() string {
	r := gtk3.SysCall("gtk_font_chooser_get_font_family", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// GetFontFace is a wrapper around gtk_font_chooser_get_font_face().
func (m *FontChooserDialog) GetFontFace() string {
	r := gtk3.SysCall("gtk_font_chooser_get_font_face", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// GetFontSize is a wrapper around gtk_font_chooser_get_font_size().
func (m *FontChooserDialog) GetFontSize() int {
	return int(gtk3.SysCall("gtk_font_chooser_get_font_size", m.Instance()))
}

// SetShowPreviewEntry is a wrapper around gtk_font_chooser_set_show_preview_entry().
func (m *FontChooserDialog) SetShowPreviewEntry(show bool) {
	gtk3.SysCall("gtk_font_chooser_set_show_preview_entry", m.Instance(), ToCBool(show))
}

// GetShowPreviewEntry is a wrapper around gtk_font_chooser_get_show_preview_entry().
func (m *FontChooserDialog) GetShowPreviewEntry() bool {
	r := gtk3.SysCall("gtk_font_chooser_get_show_preview_entry", m.Instance())
	return ToGoBool(r)
}
