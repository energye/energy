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

// HeaderBar is a representation of GTK's GtkHeaderBar.
type HeaderBar struct {
	Container
}

func AsHeaderBar(ptr unsafe.Pointer) IHeaderBar {
	if ptr == nil {
		return nil
	}
	m := new(HeaderBar)
	m.instance = ptr
	return m
}

// NewHeaderBar is a wrapper around gtk_header_bar_new().
func NewHeaderBar() IHeaderBar {
	r := gtk3.SysCall("gtk_header_bar_new")
	if r == 0 {
		return nil
	}
	return AsHeaderBar(unsafe.Pointer(r))
}

// SetTitle is a wrapper around gtk_header_bar_set_title().
func (m *HeaderBar) SetTitle(title string) {
	gtk3.SysCall("gtk_header_bar_set_title", m.Instance(), CStr(title))
}

// GetTitle is a wrapper around gtk_header_bar_get_title().
func (m *HeaderBar) GetTitle() string {
	r := gtk3.SysCall("gtk_header_bar_get_title", m.Instance())
	return GoStr(r)
}

// SetSubtitle is a wrapper around gtk_header_bar_set_subtitle().
func (m *HeaderBar) SetSubtitle(subtitle string) {
	gtk3.SysCall("gtk_header_bar_set_subtitle", m.Instance(), CStr(subtitle))
}

// GetSubtitle is a wrapper around gtk_header_bar_get_subtitle().
func (m *HeaderBar) GetSubtitle() string {
	r := gtk3.SysCall("gtk_header_bar_get_subtitle", m.Instance())
	return GoStr(r)
}

// SetCustomTitle is a wrapper around gtk_header_bar_set_custom_title().
func (m *HeaderBar) SetCustomTitle(titleWidget IWidget) {
	var widget uintptr
	if titleWidget != nil {
		widget = titleWidget.Instance()
	}
	gtk3.SysCall("gtk_header_bar_set_custom_title", m.Instance(), widget)
}

// SetShowCloseButton is a wrapper around gtk_header_bar_set_show_close_button().
func (m *HeaderBar) SetShowCloseButton(setting bool) {
	gtk3.SysCall("gtk_header_bar_set_show_close_button", m.Instance(), ToCBool(setting))
}

// GetShowCloseButton is a wrapper around gtk_header_bar_get_show_close_button().
func (m *HeaderBar) GetShowCloseButton() bool {
	r := gtk3.SysCall("gtk_header_bar_get_show_close_button", m.Instance())
	return ToGoBool(r)
}

// PackStart is a wrapper around gtk_header_bar_pack_start().
func (m *HeaderBar) PackStart(child IWidget) {
	gtk3.SysCall("gtk_header_bar_pack_start", m.Instance(), child.Instance())
}

// PackEnd is a wrapper around gtk_header_bar_pack_end().
func (m *HeaderBar) PackEnd(child IWidget) {
	gtk3.SysCall("gtk_header_bar_pack_end", m.Instance(), child.Instance())
}
