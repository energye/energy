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
	"sync"
	"github.com/ebitengine/purego"
)

// CellRendererText is a representation of GTK's GtkCellRendererText.
type CellRendererText struct {
	Object
}

var (
	cellRendererOnce sync.Once
	gObjectSetStr    func(uintptr, uintptr, uintptr, uintptr)
	gObjectSetInt    func(uintptr, uintptr, uintptr, uintptr)
	gObjectSetFloat  func(uintptr, uintptr, float32, uintptr)
)

func registerCellRendererFuncs() {
	cellRendererOnce.Do(func() {
		if gtk3 == nil || gtk3.Dll == 0 {
			return
		}
		lib := uintptr(gtk3.Dll)
		purego.RegisterLibFunc(&gObjectSetStr, lib, "g_object_set")
		purego.RegisterLibFunc(&gObjectSetInt, lib, "g_object_set")
		purego.RegisterLibFunc(&gObjectSetFloat, lib, "g_object_set")
	})
}

func AsCellRendererText(ptr unsafe.Pointer) *CellRendererText {
	if ptr == nil {
		return nil
	}
	m := new(CellRendererText)
	m.instance = ptr
	return m
}

// NewCellRendererText is a wrapper around gtk_cell_renderer_text_new().
func NewCellRendererText() *CellRendererText {
	r := gtk3.SysCall("gtk_cell_renderer_text_new")
	if r == 0 {
		return nil
	}
	return AsCellRendererText(unsafe.Pointer(r))
}

func cellPropName(s string) uintptr {
	return CStr(s)
}

// SetText sets the "text" property.
func (m *CellRendererText) SetText(text string) {
	registerCellRendererFuncs()
	gObjectSetStr(m.Instance(), cellPropName("text"), CStr(text), 0)
}

// SetMarkup sets the "markup" property.
func (m *CellRendererText) SetMarkup(markup string) {
	registerCellRendererFuncs()
	gObjectSetStr(m.Instance(), cellPropName("markup"), CStr(markup), 0)
}

// SetFont sets the "font" property (e.g. "Sans 12").
func (m *CellRendererText) SetFont(font string) {
	registerCellRendererFuncs()
	gObjectSetStr(m.Instance(), cellPropName("font"), CStr(font), 0)
}

// SetForeground sets the "foreground" color property.
func (m *CellRendererText) SetForeground(color string) {
	registerCellRendererFuncs()
	gObjectSetStr(m.Instance(), cellPropName("foreground"), CStr(color), 0)
}

// SetBackground sets the "background" color property.
func (m *CellRendererText) SetBackground(color string) {
	registerCellRendererFuncs()
	gObjectSetStr(m.Instance(), cellPropName("background"), CStr(color), 0)
}

// SetAlignment sets the "alignment" property (0.0=left, 0.5=center, 1.0=right).
func (m *CellRendererText) SetAlignment(align float32) {
	registerCellRendererFuncs()
	gObjectSetFloat(m.Instance(), cellPropName("alignment"), align, 0)
}

// SetEllipsize sets the "ellipsize" property.
func (m *CellRendererText) SetEllipsize(mode EllipsizeMode) {
	registerCellRendererFuncs()
	gObjectSetInt(m.Instance(), cellPropName("ellipsize"), uintptr(mode), 0)
}

// SetWidthChars sets the "width-chars" property.
func (m *CellRendererText) SetWidthChars(n int) {
	registerCellRendererFuncs()
	gObjectSetInt(m.Instance(), cellPropName("width-chars"), uintptr(n), 0)
}
