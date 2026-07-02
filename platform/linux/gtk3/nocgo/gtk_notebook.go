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

// Notebook is a representation of GTK's GtkNotebook.
type Notebook struct {
	Container
}

func AsNotebook(ptr unsafe.Pointer) *Notebook {
	if ptr == nil {
		return nil
	}
	m := new(Notebook)
	m.instance = ptr
	return m
}

// NewNotebook is a wrapper around gtk_notebook_new().
func NewNotebook() *Notebook {
	r := gtk3.SysCall("gtk_notebook_new")
	if r == 0 {
		return nil
	}
	return AsNotebook(unsafe.Pointer(r))
}

// AppendPage is a wrapper around gtk_notebook_append_page().
func (m *Notebook) AppendPage(child IWidget, tabLabel IWidget) int {
	var tab uintptr
	if tabLabel != nil {
		tab = tabLabel.Instance()
	}
	r := gtk3.SysCall("gtk_notebook_append_page", m.Instance(), child.Instance(), tab)
	return int(r)
}

// RemovePage is a wrapper around gtk_notebook_remove_page().
func (m *Notebook) RemovePage(pageNum int) {
	gtk3.SysCall("gtk_notebook_remove_page", m.Instance(), uintptr(pageNum))
}

// GetCurrentPage is a wrapper around gtk_notebook_get_current_page().
func (m *Notebook) GetCurrentPage() int {
	r := gtk3.SysCall("gtk_notebook_get_current_page", m.Instance())
	return int(r)
}

// SetCurrentPage is a wrapper around gtk_notebook_set_current_page().
func (m *Notebook) SetCurrentPage(pageNum int) {
	gtk3.SysCall("gtk_notebook_set_current_page", m.Instance(), uintptr(pageNum))
}

// GetNPages is a wrapper around gtk_notebook_get_n_pages().
func (m *Notebook) GetNPages() int {
	r := gtk3.SysCall("gtk_notebook_get_n_pages", m.Instance())
	return int(r)
}

// SetShowTabs is a wrapper around gtk_notebook_set_show_tabs().
func (m *Notebook) SetShowTabs(showTabs bool) {
	gtk3.SysCall("gtk_notebook_set_show_tabs", m.Instance(), ToCBool(showTabs))
}

// SetShowBorder is a wrapper around gtk_notebook_set_show_border().
func (m *Notebook) SetShowBorder(showBorder bool) {
	gtk3.SysCall("gtk_notebook_set_show_border", m.Instance(), ToCBool(showBorder))
}

// SetTabPos is a wrapper around gtk_notebook_set_tab_pos().
func (m *Notebook) SetTabPos(pos PositionType) {
	gtk3.SysCall("gtk_notebook_set_tab_pos", m.Instance(), uintptr(pos))
}
