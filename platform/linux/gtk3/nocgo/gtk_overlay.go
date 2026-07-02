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

// Overlay is a representation of GTK's GtkOverlay.
type Overlay struct {
	Bin
}

func AsOverlay(ptr unsafe.Pointer) IOverlay {
	if ptr == nil {
		return nil
	}
	m := new(Overlay)
	m.instance = ptr
	return m
}

// NewOverlay is a wrapper around gtk_overlay_new().
func NewOverlay() IOverlay {
	r := gtk3.SysCall("gtk_overlay_new")
	if r == 0 {
		return nil
	}
	return AsOverlay(unsafe.Pointer(r))
}

// AddOverlay is a wrapper around gtk_overlay_add_overlay().
func (m *Overlay) AddOverlay(widget IWidget) {
	gtk3.SysCall("gtk_overlay_add_overlay", m.Instance(), widget.Instance())
}

// ReorderOverlay is a wrapper around gtk_overlay_reorder_overlay().
func (m *Overlay) ReorderOverlay(child IWidget, position int) {
	gtk3.SysCall("gtk_overlay_reorder_overlay", m.Instance(), child.Instance(), uintptr(position))
}

// GetOverlayPassThrough is a wrapper around gtk_overlay_get_overlay_pass_through().
func (m *Overlay) GetOverlayPassThrough(widget IWidget) bool {
	r := gtk3.SysCall("gtk_overlay_get_overlay_pass_through", m.Instance(), widget.Instance())
	return ToGoBool(r)
}

// SetOverlayPassThrough is a wrapper around gtk_overlay_set_overlay_pass_through().
func (m *Overlay) SetOverlayPassThrough(widget IWidget, passThrough bool) {
	gtk3.SysCall("gtk_overlay_set_overlay_pass_through", m.Instance(), widget.Instance(), ToCBool(passThrough))
}
