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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefButton
// include/capi/views/cef_button_capi.h (cef_button_t)
type ICefButton struct {
	*ICefView
}

// Returns this Button as a LabelButton or NULL if this is not a LabelButton.
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

// Sets the current display state of the Button.
func (m *ICefButton) SetState(state consts.TCefButtonState) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetState).Call(m.Instance(), uintptr(state))
}

// Returns the current display state of the Button.
func (m *ICefButton) GetState() consts.TCefButtonState {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.Button_GetState).Call(m.Instance())
	return consts.TCefButtonState(r1)
}

// Sets the Button will use an ink drop effect for displaying state changes.
func (m *ICefButton) SetInkDropEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetInkDropEnabled).Call(m.Instance(), api.PascalBool(enabled))
}

// Sets the tooltip text that will be displayed when the user hovers the
// mouse cursor over the Button.
func (m *ICefButton) SetTooltipText(tooltipText string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetTooltipText).Call(m.Instance(), api.PascalStr(tooltipText))
}

// Sets the accessible name that will be exposed to assistive technology
// (AT).
func (m *ICefButton) SetAccessibleName(name string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.Button_SetAccessibleName).Call(m.Instance(), api.PascalStr(name))
}
