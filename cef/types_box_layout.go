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
	"unsafe"
)

// BoxLayoutRef -> ICefBoxLayout
var BoxLayoutRef boxLayout

type boxLayout uintptr

func (*boxLayout) UnWrap(data *ICefBoxLayout) *ICefBoxLayout {
	var result uintptr
	imports.Proc(def.BoxLayoutRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return nil
}

func (m *ICefBoxLayout) SetFlexForView(view *ICefView, flex int32) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.BoxLayout_SetFlexForView).Call(m.Instance(), view.Instance(), uintptr(flex))
}

func (m *ICefBoxLayout) ClearFlexForView(view *ICefView) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.BoxLayout_ClearFlexForView).Call(m.Instance(), view.Instance())
}
