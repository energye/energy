//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/types"
	"unsafe"
)

func (m *ICefDisplay) ID() (result int64) {
	imports.Proc(internale_CEFDisplay_ID).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) DeviceScaleFactor() (result types.Single) {
	imports.Proc(internale_CEFDisplay_DeviceScaleFactor).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) Rotation() int32 {
	r1, _, _ := imports.Proc(internale_CEFDisplay_Rotation).Call(uintptr(m.instance))
	return int32(r1)
}

func (m *ICefDisplay) Bounds() (result TCefRect) {
	imports.Proc(internale_CEFDisplay_Bounds).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) WorkArea() (result TCefRect) {
	imports.Proc(internale_CEFDisplay_WorkArea).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}
