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

// Box is a representation of GTK's GtkBox.
type Box struct {
	Container
}

func AsBox(ptr unsafe.Pointer) IBox {
	if ptr == nil {
		return nil
	}
	m := new(Box)
	m.instance = ptr
	return m
}

// PackStart is a wrapper around gtk_box_pack_start().
func (m *Box) PackStart(child IWidget, expand, fill bool, padding uint) {
	gtk3.SysCall("gtk_box_pack_start", m.Instance(), child.Instance(), ToCBool(expand), ToCBool(fill), uintptr(padding))
}

// PackEnd is a wrapper around gtk_box_pack_end().
func (m *Box) PackEnd(child IWidget, expand, fill bool, padding uint) {
	gtk3.SysCall("gtk_box_pack_end", m.Instance(), child.Instance(), ToCBool(expand), ToCBool(fill), uintptr(padding))
}

// NewBox is a wrapper around gtk_box_new().
func NewBox(orientation Orientation, spacing int) *Box {
	r := gtk3.SysCall("gtk_box_new", uintptr(orientation), uintptr(spacing))
	if r == 0 {
		return nil
	}
	return &Box{Container{Widget{Object{instance: unsafe.Pointer(r)}}}}
}

// GetOrientation is a wrapper around gtk_orientable_get_orientation().
func (m *Box) GetOrientation() Orientation {
	r := gtk3.SysCall("gtk_orientable_get_orientation", m.Instance())
	return Orientation(r)
}

// SetOrientation is a wrapper around gtk_orientable_set_orientation().
func (m *Box) SetOrientation(o Orientation) {
	gtk3.SysCall("gtk_orientable_set_orientation", m.Instance(), uintptr(o))
}

// GetHomogeneous is a wrapper around gtk_box_get_homogeneous().
func (m *Box) GetHomogeneous() bool {
	r := gtk3.SysCall("gtk_box_get_homogeneous", m.Instance())
	return ToGoBool(r)
}

// SetHomogeneous is a wrapper around gtk_box_set_homogeneous().
func (m *Box) SetHomogeneous(homogeneous bool) {
	gtk3.SysCall("gtk_box_set_homogeneous", m.Instance(), ToCBool(homogeneous))
}

// GetSpacing is a wrapper around gtk_box_get_spacing().
func (m *Box) GetSpacing() int {
	r := gtk3.SysCall("gtk_box_get_spacing", m.Instance())
	return int(r)
}

// SetSpacing is a wrapper around gtk_box_set_spacing().
func (m *Box) SetSpacing(spacing int) {
	gtk3.SysCall("gtk_box_set_spacing", m.Instance(), uintptr(spacing))
}

// ReorderChild is a wrapper around gtk_box_reorder_child().
func (m *Box) ReorderChild(child IWidget, position int) {
	gtk3.SysCall("gtk_box_reorder_child", m.Instance(), child.Instance(), uintptr(position))
}

// QueryChildPacking is a wrapper around gtk_box_query_child_packing().
func (m *Box) QueryChildPacking(child IWidget) (expand, fill bool, padding uint, packType PackType) {
	var cexpand, cfill, cpadding, cpackType uintptr
	gtk3.SysCall("gtk_box_query_child_packing", m.Instance(), child.Instance(),
		uintptr(unsafe.Pointer(&cexpand)), uintptr(unsafe.Pointer(&cfill)),
		uintptr(unsafe.Pointer(&cpadding)), uintptr(unsafe.Pointer(&cpackType)))
	return ToGoBool(cexpand), ToGoBool(cfill), uint(cpadding), PackType(cpackType)
}

// SetChildPacking is a wrapper around gtk_box_set_child_packing().
func (m *Box) SetChildPacking(child IWidget, expand, fill bool, padding uint, packType PackType) {
	gtk3.SysCall("gtk_box_set_child_packing", m.Instance(), child.Instance(),
		ToCBool(expand), ToCBool(fill), uintptr(padding), uintptr(packType))
}

// SetCenterWidget is a wrapper around gtk_box_set_center_widget().
func (m *Box) SetCenterWidget(child IWidget) {
	gtk3.SysCall("gtk_box_set_center_widget", m.Instance(), child.Instance())
}

// GetCenterWidget is a wrapper around gtk_box_get_center_widget().
func (m *Box) GetCenterWidget() IWidget {
	r := gtk3.SysCall("gtk_box_get_center_widget", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}
