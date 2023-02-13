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
	"unsafe"
)

type ICefCallback struct {
	instance unsafe.Pointer
}

func (m *ICefCallback) Cont() {
	imports.Proc(internale_cefCallback_Cont).Call(uintptr(m.instance))
}

func (m *ICefCallback) Cancel() {
	imports.Proc(internale_cefCallback_Cancel).Call(uintptr(m.instance))
}
