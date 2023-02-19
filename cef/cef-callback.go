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
