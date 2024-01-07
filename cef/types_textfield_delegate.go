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

// TextFieldDelegateRef -> ICefTextFieldDelegate
var TextFieldDelegateRef textFieldDelegate

type textFieldDelegate uintptr

func (*textFieldDelegate) New() *ICefTextFieldDelegate {
	var result uintptr
	imports.Proc(def.TextfieldDelegate_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefTextFieldDelegate{&ICefViewDelegate{
			instance: getInstance(result),
		}}
	}
	return nil
}

func (*textFieldDelegate) NewForCustom(textField *TCEFTextFieldComponent) *ICefTextFieldDelegate {
	if !textField.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.TextfieldDelegate_CreateForCustom).Call(textField.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefTextFieldDelegate{&ICefViewDelegate{
			instance: getInstance(result),
			ct:       consts.CtOther,
		}}
	}
	return nil
}

func (m *ICefTextFieldDelegate) SetOnKeyEvent(fn onTextFieldKeyEvent) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.TextfieldDelegate_SetOnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefTextFieldDelegate) SetOnAfterUserAction(fn onAfterUserAction) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.TextfieldDelegate_SetOnAfterUserAction).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
