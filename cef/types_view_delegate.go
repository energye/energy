//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
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

/*
*********************************
************* Delegate **********
*********************************

(*) Has CEF creation function
(d) Has delegate

------------------------          ------------------------------
| ICefViewDelegate (d) | -------> | ICefTextfieldDelegate (*d) |
------------------------    |     ------------------------------
					        |
					        |     --------------------------          ---------------------------
					        |---> | ICefPanelDelegate (*d) | -------> | ICefWindowDelegate (*d) |
					        |     --------------------------          ---------------------------
					        |
					        |     --------------------------------
					        |---> | ICefBrowserViewDelegate (*d) |
					        |     --------------------------------
					        |
					        |     ---------------------------          --------------------------------
					        |---> | ICefButtonDelegate (*d) | -------> | ICefMenuButtonDelegate (*d)  |
                                  ---------------------------          --------------------------------
*/

// ICefViewDelegate
// include/capi/views/cef_view_delegate_capi.h (cef_view_delegate_t)
type ICefViewDelegate struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
	ct       consts.CefCreateType
}

// ViewDelegateRef -> ICefViewDelegate
var ViewDelegateRef viewDelegate

type viewDelegate uintptr

func (*viewDelegate) New() *ICefViewDelegate {
	var result uintptr
	imports.Proc(def.ViewDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefViewDelegate{instance: getInstance(result)}
	}
	return nil
}

func (*viewDelegate) NewForCustom(viewComponent *TCEFViewComponent) *ICefViewDelegate {
	if !viewComponent.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ViewDelegateRef_CreateForCustom).Call(viewComponent.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefViewDelegate{instance: getInstance(result), ct: consts.CtOther}
	}
	return nil
}

func (m *ICefViewDelegate) IsSelfOwnEvent() bool {
	return m.ct == consts.CtSelfOwn
}

func (m *ICefViewDelegate) IsOtherEvent() bool {
	return m.ct == consts.CtOther
}

// Instance 实例
func (m *ICefViewDelegate) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefViewDelegate) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefViewDelegate) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefViewDelegate) SetOnGetPreferredSize(fn viewOnGetPreferredSize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetPreferredSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnGetMinimumSize(fn viewOnGetMinimumSize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetMinimumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnGetMaximumSize(fn viewOnGetMaximumSize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetMaximumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnGetHeightForWidth(fn viewOnGetHeightForWidth) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetHeightForWidth).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnParentViewChanged(fn viewOnParentViewChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnParentViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnChildViewChanged(fn viewOnChildViewChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnChildViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnWindowChanged(fn viewOnWindowChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnWindowChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnLayoutChanged(fn viewOnLayoutChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnLayoutChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnFocus(fn viewOnFocus) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnBlur(fn viewOnBlur) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnBlur).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
func (m *ICefViewDelegate) SetOnThemeChanged(fn viewOnThemeChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnThemeChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type viewOnGetPreferredSize func(view *ICefView, result *TCefSize)
type viewOnGetMinimumSize func(view *ICefView, result *TCefSize)
type viewOnGetMaximumSize func(view *ICefView, result *TCefSize)
type viewOnGetHeightForWidth func(view *ICefView, width int32) int32
type viewOnParentViewChanged func(view *ICefView, added bool, parent *ICefView)
type viewOnChildViewChanged func(view *ICefView, added bool, child *ICefView)
type viewOnWindowChanged func(view *ICefView, added bool)
type viewOnLayoutChanged func(view *ICefView, newBounds *TCefRect)
type viewOnFocus func(view *ICefView)
type viewOnBlur func(view *ICefView)
type viewOnThemeChanged func(view *ICefView)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case viewOnGetPreferredSize:
			view := &ICefView{instance: getPtr(0)}
			resultSize := (*TCefSize)(getPtr(1))
			fn.(viewOnGetPreferredSize)(view, resultSize)
		case viewOnGetMinimumSize:
			view := &ICefView{instance: getPtr(0)}
			resultSize := (*TCefSize)(getPtr(1))
			fn.(viewOnGetMinimumSize)(view, resultSize)
		case viewOnGetMaximumSize:
			view := &ICefView{instance: getPtr(0)}
			resultSize := (*TCefSize)(getPtr(1))
			fn.(viewOnGetMaximumSize)(view, resultSize)
		case viewOnGetHeightForWidth:
			view := &ICefView{instance: getPtr(0)}
			width := int32(getVal(1))
			resultPtr := (*int32)(getPtr(2))
			*resultPtr = fn.(viewOnGetHeightForWidth)(view, width)
		case viewOnParentViewChanged:
			view := &ICefView{instance: getPtr(0)}
			added := api.GoBool(getVal(1))
			parent := &ICefView{instance: getPtr(2)}
			fn.(viewOnParentViewChanged)(view, added, parent)
		case viewOnChildViewChanged:
			view := &ICefView{instance: getPtr(0)}
			added := api.GoBool(getVal(1))
			child := &ICefView{instance: getPtr(2)}
			fn.(viewOnChildViewChanged)(view, added, child)
		case viewOnWindowChanged:
			view := &ICefView{instance: getPtr(0)}
			added := api.GoBool(getVal(1))
			fn.(viewOnWindowChanged)(view, added)
		case viewOnLayoutChanged:
			view := &ICefView{instance: getPtr(0)}
			newBounds := (*TCefRect)(getPtr(1))
			fn.(viewOnLayoutChanged)(view, newBounds)
		case viewOnFocus:
			view := &ICefView{instance: getPtr(0)}
			fn.(viewOnFocus)(view)
		case viewOnBlur:
			view := &ICefView{instance: getPtr(0)}
			fn.(viewOnBlur)(view)
		case viewOnThemeChanged:
			view := &ICefView{instance: getPtr(0)}
			fn.(viewOnThemeChanged)(view)
		default:
			return false
		}
		return true
	})
}
