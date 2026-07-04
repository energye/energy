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

// TreeViewColumn is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	Object
}

func AsTreeViewColumn(ptr unsafe.Pointer) *TreeViewColumn {
	if ptr == nil {
		return nil
	}
	m := new(TreeViewColumn)
	m.instance = ptr
	return m
}

// NewTreeViewColumn is a wrapper around gtk_tree_view_column_new().
func NewTreeViewColumn() *TreeViewColumn {
	r := gtk3.SysCall("gtk_tree_view_column_new")
	if r == 0 {
		return nil
	}
	return AsTreeViewColumn(unsafe.Pointer(r))
}

// SetTitle is a wrapper around gtk_tree_view_column_set_title().
func (m *TreeViewColumn) SetTitle(title string) {
	gtk3.SysCall("gtk_tree_view_column_set_title", m.Instance(), CStr(title))
}

// GetTitle is a wrapper around gtk_tree_view_column_get_title().
func (m *TreeViewColumn) GetTitle() string {
	r := gtk3.SysCall("gtk_tree_view_column_get_title", m.Instance())
	return GoStr(r)
}

// PackStart is a wrapper around gtk_tree_view_column_pack_start().
func (m *TreeViewColumn) PackStart(cell ICellRenderer, expand bool) {
	gtk3.SysCall("gtk_tree_view_column_pack_start", m.Instance(), cell.Instance(), ToCBool(expand))
}

// AddAttribute is a wrapper around gtk_tree_view_column_add_attribute().
func (m *TreeViewColumn) AddAttribute(renderer ICellRenderer, attribute string, column int) {
	gtk3.SysCall("gtk_tree_view_column_add_attribute", m.Instance(),
		renderer.Instance(), CStr(attribute), uintptr(column))
}

func (m *TreeViewColumn) SetOnClicked(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnClicked, callback.C_trampoline_2_void, fn, 0)
}

// PackEnd is a wrapper around gtk_tree_view_column_pack_end().
func (m *TreeViewColumn) PackEnd(cell ICellRenderer, expand bool) {
	gtk3.SysCall("gtk_tree_view_column_pack_end", m.Instance(), cell.Instance(), ToCBool(expand))
}

// SetResizable is a wrapper around gtk_tree_view_column_set_resizable().
func (m *TreeViewColumn) SetResizable(resizable bool) {
	gtk3.SysCall("gtk_tree_view_column_set_resizable", m.Instance(), ToCBool(resizable))
}

// GetResizable is a wrapper around gtk_tree_view_column_get_resizable().
func (m *TreeViewColumn) GetResizable() bool {
	return ToGoBool(gtk3.SysCall("gtk_tree_view_column_get_resizable", m.Instance()))
}

// SetSizing is a wrapper around gtk_tree_view_column_set_sizing().
func (m *TreeViewColumn) SetSizing(sizing TreeViewColumnSizing) {
	gtk3.SysCall("gtk_tree_view_column_set_sizing", m.Instance(), uintptr(sizing))
}

// GetSizing is a wrapper around gtk_tree_view_column_get_sizing().
func (m *TreeViewColumn) GetSizing() TreeViewColumnSizing {
	return TreeViewColumnSizing(gtk3.SysCall("gtk_tree_view_column_get_sizing", m.Instance()))
}

// SetFixedWidth is a wrapper around gtk_tree_view_column_set_fixed_width().
func (m *TreeViewColumn) SetFixedWidth(fixedWidth int) {
	gtk3.SysCall("gtk_tree_view_column_set_fixed_width", m.Instance(), uintptr(fixedWidth))
}

// GetFixedWidth is a wrapper around gtk_tree_view_column_get_fixed_width().
func (m *TreeViewColumn) GetFixedWidth() int {
	return int(gtk3.SysCall("gtk_tree_view_column_get_fixed_width", m.Instance()))
}

// SetMinWidth is a wrapper around gtk_tree_view_column_set_min_width().
func (m *TreeViewColumn) SetMinWidth(minWidth int) {
	gtk3.SysCall("gtk_tree_view_column_set_min_width", m.Instance(), uintptr(minWidth))
}

// GetMinWidth is a wrapper around gtk_tree_view_column_get_min_width().
func (m *TreeViewColumn) GetMinWidth() int {
	return int(gtk3.SysCall("gtk_tree_view_column_get_min_width", m.Instance()))
}

// SetMaxWidth is a wrapper around gtk_tree_view_column_set_max_width().
func (m *TreeViewColumn) SetMaxWidth(maxWidth int) {
	gtk3.SysCall("gtk_tree_view_column_set_max_width", m.Instance(), uintptr(maxWidth))
}

// GetMaxWidth is a wrapper around gtk_tree_view_column_get_max_width().
func (m *TreeViewColumn) GetMaxWidth() int {
	return int(gtk3.SysCall("gtk_tree_view_column_get_max_width", m.Instance()))
}

// SetExpand is a wrapper around gtk_tree_view_column_set_expand().
func (m *TreeViewColumn) SetExpand(expand bool) {
	gtk3.SysCall("gtk_tree_view_column_set_expand", m.Instance(), ToCBool(expand))
}

// GetExpand is a wrapper around gtk_tree_view_column_get_expand().
func (m *TreeViewColumn) GetExpand() bool {
	return ToGoBool(gtk3.SysCall("gtk_tree_view_column_get_expand", m.Instance()))
}

// SetSortColumnId is a wrapper around gtk_tree_view_column_set_sort_column_id().
func (m *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	gtk3.SysCall("gtk_tree_view_column_set_sort_column_id", m.Instance(), uintptr(sortColumnId))
}

// GetSortColumnId is a wrapper around gtk_tree_view_column_get_sort_column_id().
func (m *TreeViewColumn) GetSortColumnId() int {
	return int(gtk3.SysCall("gtk_tree_view_column_get_sort_column_id", m.Instance()))
}

// SetSortIndicator is a wrapper around gtk_tree_view_column_set_sort_indicator().
func (m *TreeViewColumn) SetSortIndicator(setting bool) {
	gtk3.SysCall("gtk_tree_view_column_set_sort_indicator", m.Instance(), ToCBool(setting))
}

// GetSortIndicator is a wrapper around gtk_tree_view_column_get_sort_indicator().
func (m *TreeViewColumn) GetSortIndicator() bool {
	return ToGoBool(gtk3.SysCall("gtk_tree_view_column_get_sort_indicator", m.Instance()))
}

// SetReorderable is a wrapper around gtk_tree_view_column_set_reorderable().
func (m *TreeViewColumn) SetReorderable(reorderable bool) {
	gtk3.SysCall("gtk_tree_view_column_set_reorderable", m.Instance(), ToCBool(reorderable))
}

// GetReorderable is a wrapper around gtk_tree_view_column_get_reorderable().
func (m *TreeViewColumn) GetReorderable() bool {
	return ToGoBool(gtk3.SysCall("gtk_tree_view_column_get_reorderable", m.Instance()))
}

// SetAlignment is a wrapper around gtk_tree_view_column_set_alignment().
func (m *TreeViewColumn) SetAlignment(xalign float32) {
	registerGtkFloatFuncs()
	gtkTreeViewColumnSetAlignment(m.Instance(), xalign)
}

// GetAlignment is a wrapper around gtk_tree_view_column_get_alignment().
func (m *TreeViewColumn) GetAlignment() float32 {
	registerGtkFloatFuncs()
	return gtkTreeViewColumnGetAlignment(m.Instance())
}

// GetWidth is a wrapper around gtk_tree_view_column_get_width().
func (m *TreeViewColumn) GetWidth() int {
	return int(gtk3.SysCall("gtk_tree_view_column_get_width", m.Instance()))
}

// SetSpacing is a wrapper around gtk_tree_view_column_set_spacing().
func (m *TreeViewColumn) SetSpacing(spacing int) {
	gtk3.SysCall("gtk_tree_view_column_set_spacing", m.Instance(), uintptr(spacing))
}

// GetSpacing is a wrapper around gtk_tree_view_column_get_spacing().
func (m *TreeViewColumn) GetSpacing() int {
	return int(gtk3.SysCall("gtk_tree_view_column_get_spacing", m.Instance()))
}
