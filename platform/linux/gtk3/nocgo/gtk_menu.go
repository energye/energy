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

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

func AsMenu(ptr unsafe.Pointer) IMenu {
	if ptr == nil {
		return nil
	}
	m := new(Menu)
	m.instance = ptr
	return m
}

// NewMenu is a wrapper around gtk_menu_new().
func NewMenu() IMenu {
	r := gtk3.SysCall("gtk_menu_new")
	if r == 0 {
		return nil
	}
	return AsMenu(unsafe.Pointer(r))
}

// GtkMenuNewFromModel is a wrapper around gtk_menu_new_from_model().
func GtkMenuNewFromModel(model *GMenuModel) IMenu {
	r := gtk3.SysCall("gtk_menu_new_from_model", model.Instance())
	if r == 0 {
		return nil
	}
	return AsMenu(unsafe.Pointer(r))
}

// SetScreen is a wrapper around gtk_menu_set_screen().
func (m *Menu) SetScreen(screen *Screen) {
	gtk3.SysCall("gtk_menu_set_screen", m.Instance(), screen.Instance())
}

// Attach is a wrapper around gtk_menu_attach().
func (m *Menu) Attach(child IWidget, l, r, t, b uint) {
	gtk3.SysCall("gtk_menu_attach", m.Instance(), child.Instance(),
		uintptr(l), uintptr(r), uintptr(t), uintptr(b))
}

// SetMonitor is a wrapper around gtk_menu_set_monitor().
func (m *Menu) SetMonitor(monitorNum int) {
	gtk3.SysCall("gtk_menu_set_monitor", m.Instance(), uintptr(monitorNum))
}

// GetMonitor is a wrapper around gtk_menu_get_monitor().
func (m *Menu) GetMonitor() int {
	r := gtk3.SysCall("gtk_menu_get_monitor", m.Instance())
	return int(r)
}

// ReorderChild is a wrapper around gtk_menu_reorder_child().
func (m *Menu) ReorderChild(child IWidget, position int) {
	gtk3.SysCall("gtk_menu_reorder_child", m.Instance(), child.Instance(), uintptr(position))
}

// SetReserveToggleSize is a wrapper around gtk_menu_set_reserve_toggle_size().
func (m *Menu) SetReserveToggleSize(reserve bool) {
	gtk3.SysCall("gtk_menu_set_reserve_toggle_size", m.Instance(), ToCBool(reserve))
}

// GetReserveToggleSize is a wrapper around gtk_menu_get_reserve_toggle_size().
func (m *Menu) GetReserveToggleSize() bool {
	r := gtk3.SysCall("gtk_menu_get_reserve_toggle_size", m.Instance())
	return ToGoBool(r)
}

// Popdown is a wrapper around gtk_menu_popdown().
func (m *Menu) Popdown() {
	gtk3.SysCall("gtk_menu_popdown", m.Instance())
}

// GetActive is a wrapper around gtk_menu_get_active().
func (m *Menu) GetActive() *Menu {
	r := gtk3.SysCall("gtk_menu_get_active", m.Instance())
	if r == 0 {
		return nil
	}
	menu := new(Menu)
	menu.instance = unsafe.Pointer(r)
	return menu
}

// SetActive is a wrapper around gtk_menu_set_active().
func (m *Menu) SetActive(index uint) {
	gtk3.SysCall("gtk_menu_set_active", m.Instance(), uintptr(index))
}

// GetAttachWidget is a wrapper around gtk_menu_get_attach_widget().
func (m *Menu) GetAttachWidget() IWidget {
	r := gtk3.SysCall("gtk_menu_get_attach_widget", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}
