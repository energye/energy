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

// ButtonDelegateRef -> TCEFButtonDelegate
var ButtonDelegateRef buttonDelegate

type buttonDelegate uintptr

func (*buttonDelegate) New() *ICefButtonDelegate {
	var result uintptr
	imports.Proc(def.ButtonDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefButtonDelegate{&ICefViewDelegate{instance: getInstance(result)}}
	}
	return nil
}

func (*buttonDelegate) NewForCustom(buttonComponent *TCEFButtonComponent) *ICefButtonDelegate {
	if !buttonComponent.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ButtonDelegateRef_CreateForCustom).Call(buttonComponent.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefButtonDelegate{&ICefViewDelegate{instance: getInstance(result), ct: consts.CtOther}}
	}
	return nil
}

func (m *ICefButtonDelegate) SetOnButtonPressed(fn onButtonPressed) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ButtonDelegate_SetOnButtonPressed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefButtonDelegate) SetOnButtonStateChanged(fn onButtonStateChanged) {
	if !m.IsValid() || m.IsOtherEvent() {
		return
	}
	imports.Proc(def.ButtonDelegate_SetOnButtonStateChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
