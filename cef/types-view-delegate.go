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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

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

func (*viewDelegate) NewForCustom() *ICefViewDelegate {
	var result uintptr
	imports.Proc(def.ViewDelegateRef_CreateForCustom).Call(uintptr(unsafe.Pointer(&result)))
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

func (m *ICefViewDelegate) SetOnGetPreferredSize(fn onGetPreferredSize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetPreferredSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnGetMinimumSize(fn onGetMinimumSize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetMinimumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnGetMaximumSize(fn onGetMaximumSize) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetMaximumSize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnGetHeightForWidth(fn onGetHeightForWidth) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnGetHeightForWidth).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnParentViewChanged(fn onParentViewChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnParentViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnChildViewChanged(fn onChildViewChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnChildViewChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnWindowChanged(fn onWindowChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnWindowChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnLayoutChanged(fn onLayoutChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnLayoutChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnFocus(fn onFocus) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnFocus).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefViewDelegate) SetOnBlur(fn onBlur) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ViewDelegate_SetOnBlur).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onGetPreferredSize func(view *ICefView, result *TCefSize)
type onGetMinimumSize func(view *ICefView, result *TCefSize)
type onGetMaximumSize func(view *ICefView, result *TCefSize)
type onGetHeightForWidth func(view *ICefView, width int32) int32
type onParentViewChanged func(view *ICefView, added bool, parent *ICefView)
type onChildViewChanged func(view *ICefView, added bool, child *ICefView)
type onWindowChanged func(view *ICefView, added bool)
type onLayoutChanged func(view *ICefView, newBounds *TCefRect)
type onFocus func(view *ICefView)
type onBlur func(view *ICefView)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onGetPreferredSize:
			view := &ICefView{instance: getPtr(1)}
			resultSize := (*TCefSize)(getPtr(2))
			fn.(onGetPreferredSize)(view, resultSize)
		case onGetMinimumSize:
			view := &ICefView{instance: getPtr(1)}
			resultSize := (*TCefSize)(getPtr(2))
			fn.(onGetMinimumSize)(view, resultSize)
		case onGetMaximumSize:
			view := &ICefView{instance: getPtr(1)}
			resultSize := (*TCefSize)(getPtr(2))
			fn.(onGetMaximumSize)(view, resultSize)
		case onGetHeightForWidth:
			view := &ICefView{instance: getPtr(1)}
			width := int32(getVal(2))
			resultPtr := (*int32)(getPtr(3))
			*resultPtr = fn.(onGetHeightForWidth)(view, width)
		case onParentViewChanged:
			view := &ICefView{instance: getPtr(1)}
			added := api.GoBool(getVal(2))
			parent := &ICefView{instance: getPtr(3)}
			fn.(onParentViewChanged)(view, added, parent)
		case onChildViewChanged:
			view := &ICefView{instance: getPtr(1)}
			added := api.GoBool(getVal(2))
			child := &ICefView{instance: getPtr(3)}
			fn.(onChildViewChanged)(view, added, child)
		case onWindowChanged:
			view := &ICefView{instance: getPtr(1)}
			added := api.GoBool(getVal(2))
			fn.(onWindowChanged)(view, added)
		case onLayoutChanged:
			view := &ICefView{instance: getPtr(1)}
			newBounds := (*TCefRect)(getPtr(2))
			fn.(onLayoutChanged)(view, newBounds)
		case onFocus:
			view := &ICefView{instance: getPtr(1)}
			fn.(onFocus)(view)
		case onBlur:
			view := &ICefView{instance: getPtr(1)}
			fn.(onBlur)(view)
		default:
			return false
		}
		return true
	})
}
