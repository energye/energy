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

// TreeSelection is a representation of GTK's GtkTreeSelection.
type TreeSelection struct {
	Object
}

func AsTreeSelection(ptr unsafe.Pointer) *TreeSelection {
	if ptr == nil {
		return nil
	}
	m := new(TreeSelection)
	m.instance = ptr
	return m
}

// SetMode is a wrapper around gtk_tree_selection_set_mode().
func (m *TreeSelection) SetMode(mode SelectionMode) {
	gtk3.SysCall("gtk_tree_selection_set_mode", m.Instance(), uintptr(mode))
}

// GetMode is a wrapper around gtk_tree_selection_get_mode().
func (m *TreeSelection) GetMode() SelectionMode {
	r := gtk3.SysCall("gtk_tree_selection_get_mode", m.Instance())
	return SelectionMode(r)
}

func (m *TreeSelection) SetOnChanged(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnChanged, callback.C_trampoline_2_void, fn, 0)
}
