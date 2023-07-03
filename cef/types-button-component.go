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

// ButtonComponentRef -> TCEFButtonComponent
var ButtonComponentRef buttonComponent

type buttonComponent uintptr

func (*buttonComponent) New(AOwner lcl.IComponent) *TCEFButtonComponent {
	var result uintptr
	imports.Proc(def.ButtonComponent_Create).Call(AOwner.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFButtonComponent{&TCEFViewComponent{instance: getInstance(result)}}
	}
	return nil
}

func (m *TCEFButtonComponent) SetInkDropEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ButtonComponent_SetInkDropEnabled).Call(m.Instance(), api.PascalBool(enabled))
}

func (m *TCEFButtonComponent) SetTooltipText(tooltipText string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ButtonComponent_SetTooltipText).Call(m.Instance(), api.PascalStr(tooltipText))
}

func (m *TCEFButtonComponent) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ButtonComponent_SetAccessibleName).Call(m.Instance(), api.PascalStr(name))
}

func (m *TCEFButtonComponent) AsLabelButton() *ICefLabelButton {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ButtonComponent_AsLabelButton).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLabelButton{&ICefButton{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

func (m *TCEFButtonComponent) GetState() consts.TCefButtonState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ButtonComponent_GetState).Call(m.Instance())
	return consts.TCefButtonState(r1)
}

func (m *TCEFButtonComponent) SetState(state consts.TCefButtonState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ButtonComponent_SetState).Call(m.Instance(), uintptr(state))
}

func (m *TCEFButtonComponent) SetOnButtonPressed(fn onButtonPressed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ButtonComponent_SetOnButtonPressed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFButtonComponent) SetOnButtonStateChanged(fn onButtonStateChanged) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ButtonComponent_SetOnButtonStateChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type onButtonPressed func(sender lcl.IObject, button *ICefButton)
type onButtonStateChanged func(sender lcl.IObject, button *ICefButton)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onButtonPressed:
			button := (*ICefButton)(getPtr(1))
			fn.(onButtonPressed)(lcl.AsObject(getPtr(0)), button)
		case onButtonStateChanged:
			button := (*ICefButton)(getPtr(1))
			fn.(onButtonStateChanged)(lcl.AsObject(getPtr(0)), button)
		default:
			return false
		}
		return true
	})
}
