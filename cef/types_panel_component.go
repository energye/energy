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
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"unsafe"
)

type TCEFPanelComponent struct {
	*TCEFViewComponent
}

// PanelComponentRef -> TCEFPanelComponent
var PanelComponentRef panelComponent

type panelComponent uintptr

func (*panelComponent) New(AOwner lcl.IComponent) *TCEFPanelComponent {
	var result uintptr
	imports.Proc(def.PanelComponent_Create).Call(lcl.CheckPtr(AOwner), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFPanelComponent{&TCEFViewComponent{instance: getInstance(result)}}
	}
	return nil
}

func (m *TCEFPanelComponent) CreatePanel() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_CreatePanel).Call(m.Instance())
}

func (m *TCEFPanelComponent) SetToFillLayout() *ICefFillLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.PanelComponent_SetToFillLayout).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefFillLayout{&ICefLayout{
			instance: getInstance(result),
		}}
	}
	return nil
}

func (m *TCEFPanelComponent) SetToBoxLayout(settings TCefBoxLayoutSettings) *ICefBoxLayout {
	if !m.IsValid() {
		return nil
	}
	var (
		result      uintptr
		settingsPtr = settings.ToPtr()
	)
	imports.Proc(def.PanelComponent_SetToBoxLayout).Call(m.Instance(), uintptr(unsafe.Pointer(settingsPtr)), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBoxLayout{&ICefLayout{
			instance: getInstance(result),
		}}
	}
	return nil
}

func (m *TCEFPanelComponent) GetLayout() *ICefLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.PanelComponent_GetLayout).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLayout{
			instance: getInstance(result),
		}
	}
	return nil
}

func (m *TCEFPanelComponent) Layout() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_Layout).Call(m.Instance())
}

func (m *TCEFPanelComponent) AddChildView(view *ICefView) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_AddChildView).Call(m.Instance(), view.Instance())
}

func (m *TCEFPanelComponent) AddChildViewAt(view *ICefView, index int32) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_AddChildViewAt).Call(m.Instance(), view.Instance(), uintptr(index))
}

func (m *TCEFPanelComponent) ReorderChildView(view *ICefView, index int32) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_ReorderChildView).Call(m.Instance(), view.Instance(), uintptr(index))
}

func (m *TCEFPanelComponent) RemoveChildView(view *ICefView) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_RemoveChildView).Call(m.Instance(), view.Instance())
}

func (m *TCEFPanelComponent) RemoveAllChildViews() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.PanelComponent_RemoveAllChildViews).Call(m.Instance())
}

func (m *TCEFPanelComponent) GetChildViewCount() uint32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.PanelComponent_GetChildViewCount).Call(m.Instance())
	return uint32(r1)
}

func (m *TCEFPanelComponent) GetChildViewAt(index int32) *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.PanelComponent_GetChildViewAt).Call(m.Instance(), uintptr(index), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{
			instance: getInstance(result),
		}
	}
	return nil
}

func (m *TCEFPanelComponent) AsWindow() *ICefWindow {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.PanelComponent_AsWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}
