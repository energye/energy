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
	"unsafe"
)

// PanelDelegateRef -> ICefPanelDelegate
var PanelDelegateRef panelDelegate

type panelDelegate uintptr

func (*panelDelegate) New() *ICefPanelDelegate {
	var result uintptr
	imports.Proc(def.PanelDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanelDelegate{&ICefViewDelegate{instance: getInstance(result)}}
	}
	return nil
}

func (*panelDelegate) NewForCustom() *ICefPanelDelegate {
	var result uintptr
	imports.Proc(def.PanelDelegateRef_CreateForCustom).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanelDelegate{&ICefViewDelegate{instance: getInstance(result), ct: consts.CtOther}}
	}
	return nil
}
