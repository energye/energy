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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefWindowDelegate
// include/capi/views/cef_window_delegate_capi.h (cef_window_delegate_t)
type ICefWindowDelegate struct {
	*ICefPanelDelegate
}

// WindowDelegateRef -> ICefWindowDelegate
var WindowDelegateRef windowDelegateDelegate

type windowDelegateDelegate uintptr

func (*windowDelegateDelegate) New() *ICefWindowDelegate {
	var result uintptr
	imports.Proc(def.WindowDelegate_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindowDelegate{&ICefPanelDelegate{
			&ICefViewDelegate{
				instance: getInstance(result),
			},
		}}
	}
	return nil
}

func (*windowDelegateDelegate) NewForCustom(window *TCEFWindowComponent) *ICefWindowDelegate {
	if !window.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.WindowDelegate_CreateForCustom).Call(window.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindowDelegate{&ICefPanelDelegate{
			&ICefViewDelegate{
				instance: getInstance(result),
				ct:       consts.CtOther,
			},
		}}
	}
	return nil
}

func (m *ICefWindowDelegate) SetOnWindowCreated(fn windowOnWindowCreated) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowClosing(fn windowOnWindowClosing) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowClosing).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowDestroyed(fn windowOnWindowDestroyed) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowActivationChanged(fn windowOnWindowActivationChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowActivationChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowBoundsChanged(fn windowOnWindowBoundsChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowBoundsChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetParentWindow(fn windowOnGetParentWindow) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetParentWindow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnIsWindowModalDialog(fn windowOnIsWindowModalDialog) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnIsWindowModalDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetInitialBounds(fn windowOnGetInitialBounds) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetInitialBounds).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetInitialShowState(fn windowOnGetInitialShowState) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetInitialShowState).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnIsFrameless(fn windowOnIsFrameless) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnIsFrameless).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWithStandardWindowButtons(fn windowOnWithStandardWindowButtons) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWithStandardWindowButtons).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetTitleBarHeight(fn windowOnGetTitleBarHeight) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetTitlebarHeight).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnAcceptsFirstMouse(fn windowOnAcceptsFirstMouse) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnAcceptsFirstMouse).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanResize(fn windowOnCanResize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanResize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanMaximize(fn windowOnCanMaximize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanMaximize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanMinimize(fn windowOnCanMinimize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanMinimize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanClose(fn windowOnCanClose) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnAccelerator(fn windowOnAccelerator) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnAccelerator).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnKeyEvent(fn windowOnKey) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowFullscreenTransition(fn windowOnWindowFullscreenTransition) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowFullscreenTransition).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnThemeColorsChanged(fn windowOnThemeColorsChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnThemeColorsChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetWindowRuntimeStyle(fn windowOnGetWindowRuntimeStyle) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetWindowRuntimeStyle).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetLinuxWindowProperties(fn windowOnGetLinuxWindowProperties) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetLinuxWindowProperties).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
