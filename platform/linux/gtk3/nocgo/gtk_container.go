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

type Container struct {
	Widget
}

func AsContainer(ptr unsafe.Pointer) IContainer {
	if ptr == nil {
		return nil
	}
	m := new(Container)
	m.instance = ptr
	return m
}

// Add is a wrapper around gtk_container_add().
func (m *Container) Add(w IWidget) {
	gtk3.SysCall("gtk_container_add", m.Instance(), w.Instance())
}

// Remove is a wrapper around gtk_container_remove().
func (m *Container) Remove(w IWidget) {
	gtk3.SysCall("gtk_container_remove", m.Instance(), w.Instance())
}

// CheckResize is a wrapper around gtk_container_check_resize().
func (m *Container) CheckResize() {
	gtk3.SysCall("gtk_container_check_resize", m.Instance())
}

// GetChildren is a wrapper around gtk_container_get_children().
func (m *Container) GetChildren() IList {
	gList := gtk3.SysCall("gtk_container_get_children", m.Instance())
	if gList == 0 {
		return nil
	}
	return AsList(unsafe.Pointer(gList))
}

// GetFocusChild is a wrapper around gtk_container_get_focus_child().
func (m *Container) GetFocusChild() IWidget {
	r := gtk3.SysCall("gtk_container_get_focus_child", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

// SetFocusChild is a wrapper around gtk_container_set_focus_child().
func (m *Container) SetFocusChild(child IWidget) {
	gtk3.SysCall("gtk_container_set_focus_child", m.Instance(), child.Instance())
}

// GetFocusVAdjustment is a wrapper around gtk_container_get_focus_vadjustment().
func (m *Container) GetFocusVAdjustment() IAdjustment {
	r := gtk3.SysCall("gtk_container_get_focus_vadjustment", m.Instance())
	if r == 0 {
		return nil
	}
	return AsAdjustment(unsafe.Pointer(r))
}

// SetFocusVAdjustment is a wrapper around gtk_container_set_focus_vadjustment().
func (m *Container) SetFocusVAdjustment(adjustment *Adjustment) {
	gtk3.SysCall("gtk_container_set_focus_vadjustment", m.Instance(), adjustment.Instance())
}

// GetFocusHAdjustment is a wrapper around gtk_container_get_focus_hadjustment().
func (m *Container) GetFocusHAdjustment() IAdjustment {
	r := gtk3.SysCall("gtk_container_get_focus_hadjustment", m.Instance())
	if r == 0 {
		return nil
	}
	return AsAdjustment(unsafe.Pointer(r))
}

// SetFocusHAdjustment is a wrapper around gtk_container_set_focus_hadjustment().
func (m *Container) SetFocusHAdjustment(adjustment *Adjustment) {
	gtk3.SysCall("gtk_container_set_focus_hadjustment", m.Instance(), adjustment.Instance())
}

// ChildType is a wrapper around gtk_container_child_type().
func (m *Container) ChildType() uintptr {
	return gtk3.SysCall("gtk_container_child_type", m.Instance())
}

// ChildNotify is a wrapper around gtk_container_child_notify().
func (m *Container) ChildNotify(child IWidget, childProperty string) {
	gtk3.SysCall("gtk_container_child_notify", m.Instance(), child.Instance(), CStr(childProperty))
}

// GetBorderWidth is a wrapper around gtk_container_get_border_width().
func (m *Container) GetBorderWidth() uint {
	return uint(gtk3.SysCall("gtk_container_get_border_width", m.Instance()))
}

// SetBorderWidth is a wrapper around gtk_container_set_border_width().
func (m *Container) SetBorderWidth(borderWidth uint) {
	gtk3.SysCall("gtk_container_set_border_width", m.Instance(), uintptr(borderWidth))
}

// PropagateDraw is a wrapper around gtk_container_propagate_draw().
func (m *Container) PropagateDraw(child IWidget, cr *Context) {
	gtk3.SysCall("gtk_container_propagate_draw", m.Instance(), child.Instance(), cr.Instance())
}
