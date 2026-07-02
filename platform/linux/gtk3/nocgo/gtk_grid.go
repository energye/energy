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

// Grid is a representation of GTK's GtkGrid.
type Grid struct {
	Container
}

func AsGrid(ptr unsafe.Pointer) *Grid {
	if ptr == nil {
		return nil
	}
	m := new(Grid)
	m.instance = ptr
	return m
}

// NewGrid is a wrapper around gtk_grid_new().
func NewGrid() *Grid {
	r := gtk3.SysCall("gtk_grid_new")
	if r == 0 {
		return nil
	}
	return AsGrid(unsafe.Pointer(r))
}

// Attach is a wrapper around gtk_grid_attach().
func (m *Grid) Attach(child IWidget, left, top, width, height int) {
	gtk3.SysCall("gtk_grid_attach", m.Instance(), child.Instance(),
		uintptr(left), uintptr(top), uintptr(width), uintptr(height))
}

// AttachNextTo is a wrapper around gtk_grid_attach_next_to().
func (m *Grid) AttachNextTo(child, sibling IWidget, side PositionType, width, height int) {
	gtk3.SysCall("gtk_grid_attach_next_to", m.Instance(), child.Instance(),
		sibling.Instance(), uintptr(side), uintptr(width), uintptr(height))
}

// SetRowSpacing is a wrapper around gtk_grid_set_row_spacing().
func (m *Grid) SetRowSpacing(spacing uint) {
	gtk3.SysCall("gtk_grid_set_row_spacing", m.Instance(), uintptr(spacing))
}

// GetRowSpacing is a wrapper around gtk_grid_get_row_spacing().
func (m *Grid) GetRowSpacing() uint {
	r := gtk3.SysCall("gtk_grid_get_row_spacing", m.Instance())
	return uint(r)
}

// SetColumnSpacing is a wrapper around gtk_grid_set_column_spacing().
func (m *Grid) SetColumnSpacing(spacing uint) {
	gtk3.SysCall("gtk_grid_set_column_spacing", m.Instance(), uintptr(spacing))
}

// GetColumnSpacing is a wrapper around gtk_grid_get_column_spacing().
func (m *Grid) GetColumnSpacing() uint {
	r := gtk3.SysCall("gtk_grid_get_column_spacing", m.Instance())
	return uint(r)
}

// SetRowHomogeneous is a wrapper around gtk_grid_set_row_homogeneous().
func (m *Grid) SetRowHomogeneous(homogeneous bool) {
	gtk3.SysCall("gtk_grid_set_row_homogeneous", m.Instance(), ToCBool(homogeneous))
}

// GetRowHomogeneous is a wrapper around gtk_grid_get_row_homogeneous().
func (m *Grid) GetRowHomogeneous() bool {
	r := gtk3.SysCall("gtk_grid_get_row_homogeneous", m.Instance())
	return ToGoBool(r)
}

// SetColumnHomogeneous is a wrapper around gtk_grid_set_column_homogeneous().
func (m *Grid) SetColumnHomogeneous(homogeneous bool) {
	gtk3.SysCall("gtk_grid_set_column_homogeneous", m.Instance(), ToCBool(homogeneous))
}

// GetColumnHomogeneous is a wrapper around gtk_grid_get_column_homogeneous().
func (m *Grid) GetColumnHomogeneous() bool {
	r := gtk3.SysCall("gtk_grid_get_column_homogeneous", m.Instance())
	return ToGoBool(r)
}
