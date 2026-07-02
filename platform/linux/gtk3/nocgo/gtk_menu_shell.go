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

// MenuShell is a representation of GTK's GtkMenuShell.
type MenuShell struct {
	Container
}

func AsMenuShell(ptr unsafe.Pointer) IMenuShell {
	if ptr == nil {
		return nil
	}
	m := new(MenuShell)
	m.instance = ptr
	return m
}

// Append is a wrapper around gtk_menu_shell_append().
func (m *MenuShell) Append(child IWidget) {
	gtk3.SysCall("gtk_menu_shell_append", m.Instance(), child.Instance())
}

// Prepend is a wrapper around gtk_menu_shell_prepend().
func (m *MenuShell) Prepend(child IWidget) {
	gtk3.SysCall("gtk_menu_shell_prepend", m.Instance(), child.Instance())
}

// Insert is a wrapper around gtk_menu_shell_insert().
func (m *MenuShell) Insert(child IWidget, position int) {
	gtk3.SysCall("gtk_menu_shell_insert", m.Instance(), child.Instance(), uintptr(position))
}

// Deactivate is a wrapper around gtk_menu_shell_deactivate().
func (m *MenuShell) Deactivate() {
	gtk3.SysCall("gtk_menu_shell_deactivate", m.Instance())
}

// SelectItem is a wrapper around gtk_menu_shell_select_item().
func (m *MenuShell) SelectItem(child IWidget) {
	gtk3.SysCall("gtk_menu_shell_select_item", m.Instance(), child.Instance())
}

// SelectFirst is a wrapper around gtk_menu_shell_select_first().
func (m *MenuShell) SelectFirst(searchSensitive bool) {
	gtk3.SysCall("gtk_menu_shell_select_first", m.Instance(), ToCBool(searchSensitive))
}

// Deselect is a wrapper around gtk_menu_shell_deselect().
func (m *MenuShell) Deselect() {
	gtk3.SysCall("gtk_menu_shell_deselect", m.Instance())
}

// ActivateItem is a wrapper around gtk_menu_shell_activate_item().
func (m *MenuShell) ActivateItem(child IWidget, forceDeactivate bool) {
	gtk3.SysCall("gtk_menu_shell_activate_item", m.Instance(), child.Instance(), ToCBool(forceDeactivate))
}

// Cancel is a wrapper around gtk_menu_shell_cancel().
func (m *MenuShell) Cancel() {
	gtk3.SysCall("gtk_menu_shell_cancel", m.Instance())
}

// SetTakeFocus is a wrapper around gtk_menu_shell_set_take_focus().
func (m *MenuShell) SetTakeFocus(takeFocus bool) {
	gtk3.SysCall("gtk_menu_shell_set_take_focus", m.Instance(), ToCBool(takeFocus))
}

// GetTakeFocus is a wrapper around gtk_menu_shell_get_take_focus().
func (m *MenuShell) GetTakeFocus() bool {
	r := gtk3.SysCall("gtk_menu_shell_get_take_focus", m.Instance())
	return ToGoBool(r)
}

// GetSelectedItem is a wrapper around gtk_menu_shell_get_selected_item().
func (m *MenuShell) GetSelectedItem() (IWidget, error) {
	r := gtk3.SysCall("gtk_menu_shell_get_selected_item", m.Instance())
	if r == 0 {
		return nil, errNilPtr
	}
	return AsWidget(unsafe.Pointer(r)), nil
}

// GetParentShell is a wrapper around gtk_menu_shell_get_parent_shell().
func (m *MenuShell) GetParentShell() (IMenuShell, error) {
	r := gtk3.SysCall("gtk_menu_shell_get_parent_shell", m.Instance())
	if r == 0 {
		return nil, errNilPtr
	}
	return AsMenuShell(unsafe.Pointer(r)), nil
}

// BindModel is a wrapper around gtk_menu_shell_bind_model().
func (m *MenuShell) BindModel(model *GMenuModel, actionNamespace string, withSeparators bool) {
	var modelPtr uintptr
	if model != nil {
		modelPtr = model.Instance()
	}
	cstr := CStr(actionNamespace)
	gtk3.SysCall("gtk_menu_shell_bind_model", m.Instance(), modelPtr, cstr, ToCBool(withSeparators))
}
