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

func (m *ICefButton) AsLabelButton() *ICefLabelButton {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.Button_AsLabelButton).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLabelButton{&ICefButton{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

func (m *ICefButton) SetState(state consts.TCefButtonState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetState).Call(m.Instance(), uintptr(state))
}

func (m *ICefButton) GetState() consts.TCefButtonState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.Button_GetState).Call(m.Instance())
	return consts.TCefButtonState(r1)
}

func (m *ICefButton) SetInkDropEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetInkDropEnabled).Call(m.Instance(), api.PascalBool(enabled))
}

func (m *ICefButton) SetTooltipText(tooltipText string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetTooltipText).Call(m.Instance(), api.PascalStr(tooltipText))
}

func (m *ICefButton) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetAccessibleName).Call(m.Instance(), api.PascalStr(name))
}
