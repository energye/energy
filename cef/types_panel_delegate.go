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
	"github.com/energye/energy/v2/consts"
	"unsafe"
)

// PanelDelegateRef -> ICefPanelDelegate
var PanelDelegateRef panelDelegate

type panelDelegate uintptr

func (*panelDelegate) New() *ICefPanelDelegate {
	var result uintptr
	imports.Proc(def.PanelDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanelDelegate{&ICefViewDelegate{instance: getInstance(result)}}
	}
	return nil
}

func (*panelDelegate) NewForCustom(panel *TCEFPanelComponent) *ICefPanelDelegate {
	if !panel.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.PanelDelegateRef_CreateForCustom).Call(panel.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanelDelegate{&ICefViewDelegate{instance: getInstance(result), ct: consts.CtOther}}
	}
	return nil
}

//func (m *ICefPanelDelegate) SetOnGetPreferredSize(fn onGetPreferredSize) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnGetPreferredSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnGetMinimumSize(fn onGetMinimumSize) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnGetMinimumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnGetMaximumSize(fn onGetMaximumSize) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnGetMaximumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnGetHeightForWidth(fn onGetHeightForWidth) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnGetHeightForWidth).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnParentViewChanged(fn onParentViewChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnParentViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnChildViewChanged(fn onChildViewChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnChildViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnWindowChanged(fn onWindowChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnWindowChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnLayoutChanged(fn onLayoutChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnLayoutChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnFocus(fn onFocus) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefPanelDelegate) SetOnBlur(fn onBlur) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.PanelDelegate_SetOnBlur).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
