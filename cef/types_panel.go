//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https//www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"unsafe"
)

// ICefPanel
// include/capi/views/cef_panel_capi.h (cef_panel_t)
type ICefPanel struct {
	*ICefView
}

// PanelRef -> ICefPanel
var PanelRef panel

type panel uintptr

func (*panel) New(delegate *ICefPanelDelegate) *ICefPanel {
	var result uintptr
	imports.Proc(def.CEFPanelRef_New).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanel{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefPanel) GetAsWindow() *ICefWindow {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFPanel_GetAsWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

func (m *ICefPanel) SetToFillLayout() *ICefFillLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFPanel_SetToFillLayout).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFillLayout{&ICefLayout{
			instance: getInstance(result),
		}}
	}
	return nil
}

func (m *ICefPanel) SetToBoxLayout(settings TCefBoxLayoutSettings) *ICefBoxLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	settingsPtr := settings.ToPtr()
	imports.Proc(def.CEFPanel_SetToBoxLayout).Call(m.Instance(), uintptr(unsafe.Pointer(settingsPtr)), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBoxLayout{&ICefLayout{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefPanel) GetLayout() *ICefLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFPanel_GetLayout).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLayout{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefPanel) Layout() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFPanel_Layout).Call(m.Instance())
}

func (m *ICefPanel) AddChildView(view *ICefView) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFPanel_AddChildView).Call(m.Instance(), view.Instance())
}

func (m *ICefPanel) AddChildViewAt(view *ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFPanel_AddChildViewAt).Call(m.Instance(), view.Instance(), uintptr(index))
}

func (m *ICefPanel) ReorderChildView(view *ICefView, index int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFPanel_ReorderChildView).Call(m.Instance(), view.Instance(), uintptr(index))
}

func (m *ICefPanel) RemoveChildView(view *ICefView) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFPanel_RemoveChildView).Call(m.Instance(), view.Instance())
}

func (m *ICefPanel) RemoveAllChildViews() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFPanel_RemoveAllChildViews).Call(m.Instance())
}

func (m *ICefPanel) GetChildViewCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFPanel_GetChildViewCount).Call(m.Instance())
	return uint32(r1)
}

func (m *ICefPanel) GetChildViewAt(index int32) *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFPanel_GetChildViewAt).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}
