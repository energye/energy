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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
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

type windowOnWindowCreated func(window *ICefWindow)
type windowOnWindowDestroyed func(window *ICefWindow)
type windowOnWindowActivationChanged func(window *ICefWindow, active bool)
type windowOnGetParentWindow func(window *ICefWindow, isMenu, canActivateMenu *bool) *ICefWindow
type windowOnIsWindowModalDialog func(window *ICefWindow, result *bool)
type windowOnGetInitialBounds func(window *ICefWindow, result *TCefRect)
type windowOnGetInitialShowState func(window *ICefWindow, result *consts.TCefShowState)
type windowOnIsFrameless func(window *ICefWindow, result *bool)
type windowOnWithStandardWindowButtons func(window *ICefWindow, result *bool)
type windowOnGetTitleBarHeight func(window *ICefWindow, titleBarHeight *float32, result *bool)
type windowOnCanResize func(window *ICefWindow, result *bool)
type windowOnCanMaximize func(window *ICefWindow, result *bool)
type windowOnCanMinimize func(window *ICefWindow, result *bool)
type windowOnCanClose func(window *ICefWindow, result *bool)
type windowOnCanCloseEx func(cefWindow *ICefWindow, window IBrowserWindow, canClose *bool) bool
type windowOnAccelerator func(window *ICefWindow, commandId int32, result *bool)
type windowOnKey func(window *ICefWindow, event *TCefKeyEvent, result *bool)
type windowOnWindowFullscreenTransition func(window *ICefWindow, isCompleted bool)
type windowOnThemeColorsChanged func(window *ICefWindow, chromeTheme int32)
type windowOnGetWindowRuntimeStyle func(result *consts.TCefRuntimeStyle)
type windowOnGetLinuxWindowProperties func(window *ICefWindow, properties *TLinuxWindowProperties, result *bool)
type windowOnWindowClosing func(window *ICefWindow)
type windowOnWindowBoundsChanged func(window *ICefWindow, newBounds TCefRect)
type windowOnAcceptsFirstMouse func(window *ICefWindow, result *consts.TCefState)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		var getWindow = func(index int) *ICefWindow {
			return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(getPtr(index))}}}
		}
		switch fn.(type) {
		case windowOnWindowCreated:
			window := getWindow(0)
			fn.(windowOnWindowCreated)(window)
		case windowOnWindowDestroyed:
			window := getWindow(0)
			fn.(windowOnWindowDestroyed)(window)
		case windowOnWindowActivationChanged:
			window := getWindow(0)
			fn.(windowOnWindowActivationChanged)(window, api.GoBool(getVal(1)))
		case windowOnGetParentWindow:
			window := getWindow(0)
			resultWindowPtr := (*uintptr)(getPtr(3))
			parentWindow := fn.(windowOnGetParentWindow)(window, (*bool)(getPtr(1)), (*bool)(getPtr(2)))
			if window != nil {
				*resultWindowPtr = parentWindow.Instance()
			}
		case windowOnIsWindowModalDialog:
			window := getWindow(0)
			fn.(windowOnIsWindowModalDialog)(window, (*bool)(getPtr(1)))
		case windowOnGetInitialBounds:
			window := getWindow(0)
			resultRectPtr := (*TCefRect)(getPtr(1))
			resultRect := new(TCefRect)
			resultRect.X = 0
			resultRect.Y = 0
			resultRect.Width = 600
			resultRect.Height = 400
			fn.(windowOnGetInitialBounds)(window, resultRect)
			*resultRectPtr = *resultRect
		case windowOnGetInitialShowState:
			window := getWindow(0)
			resultShowState := (*consts.TCefShowState)(getPtr(1))
			fn.(windowOnGetInitialShowState)(window, resultShowState)
		case windowOnIsFrameless:
			window := getWindow(0)
			fn.(windowOnIsFrameless)(window, (*bool)(getPtr(1)))
		case windowOnWithStandardWindowButtons:
			window := getWindow(0)
			fn.(windowOnWithStandardWindowButtons)(window, (*bool)(getPtr(1)))
		case windowOnGetTitleBarHeight:
			window := getWindow(0)
			titleBarHeight := (*float32)(getPtr(1))
			fn.(windowOnGetTitleBarHeight)(window, titleBarHeight, (*bool)(getPtr(2)))
		case windowOnCanResize:
			window := getWindow(0)
			fn.(windowOnCanResize)(window, (*bool)(getPtr(1)))
		case windowOnCanMaximize:
			window := getWindow(0)
			fn.(windowOnCanMaximize)(window, (*bool)(getPtr(1)))
		case windowOnCanMinimize:
			window := getWindow(0)
			fn.(windowOnCanMinimize)(window, (*bool)(getPtr(1)))
		case windowOnCanClose:
			window := getWindow(0)
			fn.(windowOnCanClose)(window, (*bool)(getPtr(1)))
		case windowOnAccelerator:
			window := getWindow(0)
			fn.(windowOnAccelerator)(window, int32(getVal(1)), (*bool)(getPtr(2)))
		case windowOnKey:
			window := getWindow(0)
			keyEvent := (*TCefKeyEvent)(getPtr(1))
			fn.(windowOnKey)(window, keyEvent, (*bool)(getPtr(2)))
		case windowOnWindowFullscreenTransition:
			window := getWindow(0)
			isCompleted := api.GoBool(getVal(1))
			fn.(windowOnWindowFullscreenTransition)(window, isCompleted)
		case windowOnThemeColorsChanged:
			window := getWindow(0)
			fn.(windowOnThemeColorsChanged)(window, int32(getVal(1)))
		case windowOnGetWindowRuntimeStyle:
			fn.(windowOnGetWindowRuntimeStyle)((*consts.TCefRuntimeStyle)(getPtr(0)))
		case windowOnGetLinuxWindowProperties:
			window := getWindow(0)
			propertiesPtr := (*tLinuxWindowPropertiesPtr)(getPtr(1))
			properties := propertiesPtr.convert()
			fn.(windowOnGetLinuxWindowProperties)(window, properties, (*bool)(getPtr(2)))
			properties.setInstanceValue()
		case windowOnWindowClosing:
			window := getWindow(0)
			fn.(windowOnWindowClosing)(window)
		case windowOnWindowBoundsChanged:
			window := getWindow(0)
			newBounds := *(*TCefRect)(getPtr(1))
			fn.(windowOnWindowBoundsChanged)(window, newBounds)
		case windowOnAcceptsFirstMouse:
			window := getWindow(0)
			result := (*consts.TCefState)(getPtr(1))
			fn.(windowOnAcceptsFirstMouse)(window, result)
		default:
			return false
		}
		return true
	})
}
