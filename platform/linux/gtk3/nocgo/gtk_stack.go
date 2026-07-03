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

// ─────────────────────────────────────────────
// Stack
// ─────────────────────────────────────────────

type Stack struct {
	Container
}

func AsStack(ptr unsafe.Pointer) *Stack {
	if ptr == nil {
		return nil
	}
	m := new(Stack)
	m.instance = ptr
	return m
}

func NewStack() *Stack {
	r := gtk3.SysCall("gtk_stack_new")
	if r == 0 {
		return nil
	}
	return AsStack(unsafe.Pointer(r))
}

func (m *Stack) AddNamed(child IWidget, name string) {
	gtk3.SysCall("gtk_stack_add_named", m.Instance(), child.Instance(), CStr(name))
}

func (m *Stack) AddTitled(child IWidget, name, title string) {
	gtk3.SysCall("gtk_stack_add_titled", m.Instance(), child.Instance(), CStr(name), CStr(title))
}

func (m *Stack) SetVisibleChild(child IWidget) {
	gtk3.SysCall("gtk_stack_set_visible_child", m.Instance(), child.Instance())
}

func (m *Stack) GetVisibleChild() IWidget {
	r := gtk3.SysCall("gtk_stack_get_visible_child", m.Instance())
	if r == 0 {
		return nil
	}
	return AsWidget(unsafe.Pointer(r))
}

func (m *Stack) SetVisibleChildName(name string) {
	gtk3.SysCall("gtk_stack_set_visible_child_name", m.Instance(), CStr(name))
}

func (m *Stack) GetVisibleChildName() string {
	r := gtk3.SysCall("gtk_stack_get_visible_child_name", m.Instance())
	return GoStr(r)
}

func (m *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	gtk3.SysCall("gtk_stack_set_visible_child_full", m.Instance(), CStr(name), uintptr(transition))
}

func (m *Stack) SetHomogeneous(homogeneous bool) {
	gtk3.SysCall("gtk_stack_set_homogeneous", m.Instance(), ToCBool(homogeneous))
}

func (m *Stack) GetHomogeneous() bool {
	r := gtk3.SysCall("gtk_stack_get_homogeneous", m.Instance())
	return ToGoBool(r)
}

func (m *Stack) SetTransitionDuration(duration uint) {
	gtk3.SysCall("gtk_stack_set_transition_duration", m.Instance(), uintptr(duration))
}

func (m *Stack) GetTransitionDuration() uint {
	r := gtk3.SysCall("gtk_stack_get_transition_duration", m.Instance())
	return uint(r)
}

func (m *Stack) SetTransitionType(transition StackTransitionType) {
	gtk3.SysCall("gtk_stack_set_transition_type", m.Instance(), uintptr(transition))
}

func (m *Stack) GetTransitionType() StackTransitionType {
	r := gtk3.SysCall("gtk_stack_get_transition_type", m.Instance())
	return StackTransitionType(r)
}

// ─────────────────────────────────────────────
// StackSwitcher
// ─────────────────────────────────────────────

type StackSwitcher struct {
	Box
}

func AsStackSwitcher(ptr unsafe.Pointer) *StackSwitcher {
	if ptr == nil {
		return nil
	}
	m := new(StackSwitcher)
	m.instance = ptr
	return m
}

func NewStackSwitcher() *StackSwitcher {
	r := gtk3.SysCall("gtk_stack_switcher_new")
	if r == 0 {
		return nil
	}
	return AsStackSwitcher(unsafe.Pointer(r))
}

func (m *StackSwitcher) SetStack(stack IStack) {
	gtk3.SysCall("gtk_stack_switcher_set_stack", m.Instance(), stack.Instance())
}

func (m *StackSwitcher) GetStack() IStack {
	r := gtk3.SysCall("gtk_stack_switcher_get_stack", m.Instance())
	if r == 0 {
		return nil
	}
	return AsStack(unsafe.Pointer(r))
}
