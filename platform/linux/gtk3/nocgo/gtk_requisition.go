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

// Requisition is a representation of GTK's GtkRequisition.
type Requisition struct {
	Width  int32
	Height int32
}

func AsRequisition(ptr unsafe.Pointer) IRequisition {
	if ptr == nil {
		return nil
	}
	return (*Requisition)(ptr)
}

// NewRequisition is a wrapper around gtk_requisition_new().
func NewRequisition() IRequisition {
	r := gtk3.SysCall("gtk_requisition_new")
	if r == 0 {
		return nil
	}
	return AsRequisition(unsafe.Pointer(r))
}

// GetWidth returns the width of the requisition.
func (m *Requisition) GetWidth() int {
	return int(m.Width)
}

// GetHeight returns the height of the requisition.
func (m *Requisition) GetHeight() int {
	return int(m.Height)
}

// Free is a wrapper around gtk_requisition_free().
func (m *Requisition) Free() {
	gtk3.SysCall("gtk_requisition_free", uintptr(unsafe.Pointer(m)))
}

// Native returns the underlying pointer as uintptr.
func (m *Requisition) Native() uintptr {
	return uintptr(unsafe.Pointer(m))
}

// Copy is a wrapper around gtk_requisition_copy().
func (m *Requisition) Copy() *Requisition {
	r := gtk3.SysCall("gtk_requisition_copy", uintptr(unsafe.Pointer(m)))
	if r == 0 {
		return nil
	}
	return (*Requisition)(unsafe.Pointer(r))
}
