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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

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

//func (m *ICefWindowDelegate) SetOnGetPreferredSize(fn onGetPreferredSize) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnGetPreferredSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnGetMinimumSize(fn onGetMinimumSize) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnGetMinimumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnGetMaximumSize(fn onGetMaximumSize) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnGetMaximumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnGetHeightForWidth(fn onGetHeightForWidth) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnGetHeightForWidth).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnParentViewChanged(fn onParentViewChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnParentViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnChildViewChanged(fn onChildViewChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnChildViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnWindowChanged(fn onWindowChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnWindowChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnLayoutChanged(fn onLayoutChanged) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnLayoutChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnFocus(fn onFocus) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}
//
//func (m *ICefWindowDelegate) SetOnBlur(fn onBlur) {
//	if !m.IsValid() || m.IsOtherEvent() {
//		return
//	}
//	imports.Proc(def.WindowDelegate_SetOnBlur).Call(m.Instance(), api.MakeEventDataPtr(fn))
//}

func (m *ICefWindowDelegate) SetOnWindowCreated(fn WindowComponentOnWindowCreated) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowClosing(fn WindowOnWindowClosing) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowClosing).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowDestroyed(fn WindowComponentOnWindowDestroyed) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowActivationChanged(fn WindowComponentOnWindowActivationChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowActivationChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowBoundsChanged(fn WindowOnWindowBoundsChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowBoundsChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetParentWindow(fn WindowComponentOnGetParentWindow) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetParentWindow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnIsWindowModalDialog(fn WindowComponentOnIsWindowModalDialog) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnIsWindowModalDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetInitialBounds(fn WindowComponentOnGetInitialBounds) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetInitialBounds).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetInitialShowState(fn WindowComponentOnGetInitialShowState) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetInitialShowState).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnIsFrameless(fn WindowComponentOnIsFrameless) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnIsFrameless).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWithStandardWindowButtons(fn WindowComponentOnWithStandardWindowButtons) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWithStandardWindowButtons).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnGetTitleBarHeight(fn WindowComponentOnGetTitleBarHeight) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnGetTitlebarHeight).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanResize(fn WindowComponentOnCanResize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanResize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanMaximize(fn WindowComponentOnCanMaximize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanMaximize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanMinimize(fn WindowComponentOnCanMinimize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanMinimize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnCanClose(fn WindowComponentOnCanClose) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnCanClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnAccelerator(fn WindowComponentOnAccelerator) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnAccelerator).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnKeyEvent(fn WindowComponentOnKey) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefWindowDelegate) SetOnWindowFullscreenTransition(fn WindowComponentOnWindowFullscreenTransition) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.WindowDelegate_SetOnWindowFullscreenTransition).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type WindowOnWindowClosing func(window *ICefWindow)
type WindowOnWindowBoundsChanged func(window *ICefWindow, newBounds *TCefRect)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case WindowOnWindowClosing:
			window := &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(getPtr(0))}}}
			fn.(WindowOnWindowClosing)(window)
		case WindowOnWindowBoundsChanged:
			window := &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(getPtr(0))}}}
			newBounds := (*TCefRect)(getPtr(1))
			fn.(WindowOnWindowBoundsChanged)(window, newBounds)
		default:
			return false
		}
		return true
	})
}
