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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefLabelButton
// include/capi/views/cef_label_button_capi.h (cef_label_button_t)
type ICefLabelButton struct {
	*ICefButton
}

// LabelButtonRef -> ICefLabelButton
var LabelButtonRef labelButton

type labelButton uintptr

func (*labelButton) New(delegate *ICefButtonDelegate, text string) *ICefLabelButton {
	var result uintptr
	imports.Proc(def.LabelButtonRef_CreateLabelButton).Call(delegate.Instance(), api.PascalStr(text), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefLabelButton{&ICefButton{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

func (m *ICefLabelButton) AsMenuButton() *ICefMenuButton {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.LabelButton_AsMenuButton).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuButton{&ICefLabelButton{&ICefButton{&ICefView{instance: getInstance(result)}}}}
	}
	return nil
}

func (m *ICefLabelButton) SetText(text string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetText).Call(m.Instance(), api.PascalStr(text))
}

func (m *ICefLabelButton) GetText() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.LabelButton_GetText).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefLabelButton) SetImage(buttonState consts.TCefButtonState, image *ICefImage) {
	if !m.IsValid() || !image.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetImage).Call(m.Instance(), uintptr(buttonState), image.Instance())
}

func (m *ICefLabelButton) GetImage() *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.LabelButton_GetImage).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefImage{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefLabelButton) SetTextColor(forState consts.TCefButtonState, color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetTextColor).Call(m.Instance(), uintptr(forState), uintptr(color))
}

func (m *ICefLabelButton) SetEnabledTextColors(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetEnabledTextColors).Call(m.Instance(), uintptr(color))
}

func (m *ICefLabelButton) SetFontList(fontList string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetFontList).Call(m.Instance(), api.PascalStr(fontList))
}

func (m *ICefLabelButton) SetHorizontalAlignment(alignment consts.TCefHorizontalAlignment) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetHorizontalAlignment).Call(m.Instance(), uintptr(alignment))
}

func (m *ICefLabelButton) SetMinimumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetMinimumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

func (m *ICefLabelButton) SetMaximumSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.LabelButton_SetMaximumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}
