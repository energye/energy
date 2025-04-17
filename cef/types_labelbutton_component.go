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
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCEFLabelButtonComponent
type TCEFLabelButtonComponent struct {
	*TCEFButtonComponent
}

// LabelButtonComponentRef -> TCEFLabelButtonComponent
var LabelButtonComponentRef labelButtonComponent

type labelButtonComponent uintptr

func (m *labelButtonComponent) New(AOwner lcl.IComponent) *TCEFLabelButtonComponent {
	var result uintptr
	imports.Proc(def.LabelButtonComponent_Create).Call(AOwner.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFLabelButtonComponent{&TCEFButtonComponent{&TCEFViewComponent{instance: getInstance(result)}}}
	}
	return nil
}

func (m *TCEFLabelButtonComponent) CreateLabelButton(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_CreateLabelButton).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFLabelButtonComponent) SetTextColor(forState consts.TCefButtonState, color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetTextColor).Call(m.Instance(), uintptr(forState), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetEnabledTextColors(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetEnabledTextColors).Call(m.Instance(), uintptr(color))
}

func (m *TCEFLabelButtonComponent) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetFontList).Call(m.Instance(), api.PascalStr(fontList))
}

func (m *TCEFLabelButtonComponent) SetHorizontalAlignment(alignment consts.TCefHorizontalAlignment) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetHorizontalAlignment).Call(m.Instance(), uintptr(alignment))
}

func (m *TCEFLabelButtonComponent) SetMinimumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetMinimumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

func (m *TCEFLabelButtonComponent) SetMaximumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetMaximumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

func (m *TCEFLabelButtonComponent) GetText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.LabelButtonComponent_GetText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *TCEFLabelButtonComponent) SetText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetText).Call(m.Instance(), api.PascalStr(text))
}

func (m *TCEFLabelButtonComponent) GetImage(buttonState consts.TCefButtonState) *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.LabelButtonComponent_GetImage).Call(m.Instance(), uintptr(buttonState), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefImage{instance: getInstance(result)}
	}
	return nil
}

func (m *TCEFLabelButtonComponent) SetImage(buttonState consts.TCefButtonState, image *ICefImage) {
	if !m.IsValid() || !image.IsValid() {
		return
	}
	imports.Proc(def.LabelButtonComponent_SetImage).Call(m.Instance(), uintptr(buttonState), image.Instance())
}

func (m *TCEFLabelButtonComponent) AsMenuButton() *ICefMenuButton {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.LabelButtonComponent_AsMenuButton).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuButton{&ICefLabelButton{&ICefButton{&ICefView{instance: getInstance(result)}}}}
	}
	return nil
}
