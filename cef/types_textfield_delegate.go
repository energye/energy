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

// ICefTextFieldDelegate
// include/capi/views/cef_textfield_delegate_capi.h (cef_textfield_delegate_t)
type ICefTextFieldDelegate struct {
	*ICefViewDelegate
}

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

func (m *ICefTextFieldDelegate) SetOnKeyEvent(fn textFieldOnTextFieldKeyEvent) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.TextfieldDelegate_SetOnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefTextFieldDelegate) SetOnAfterUserAction(fn textFieldOnAfterUserAction) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.TextfieldDelegate_SetOnAfterUserAction).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type textFieldOnTextFieldKeyEvent func(textField *ICefTextfield, event *TCefKeyEvent) bool
type textFieldOnAfterUserAction func(textField *ICefTextfield)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case textFieldOnTextFieldKeyEvent:
			textField := &ICefTextfield{&ICefView{instance: getPtr(0)}}
			event := (*TCefKeyEvent)(getPtr(1))
			resultPtr := (*bool)(getPtr(2))
			*resultPtr = fn.(textFieldOnTextFieldKeyEvent)(textField, event)
		case textFieldOnAfterUserAction:
			textField := &ICefTextfield{&ICefView{instance: getPtr(0)}}
			fn.(textFieldOnAfterUserAction)(textField)
		default:
			return false
		}
		return true
	})
}
