//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	"unsafe"
)

type ICefCallback struct {
	instance unsafe.Pointer
}

func (m *ICefCallback) Cont() {
	Proc(internale_cefCallback_Cont).Call(uintptr(m.instance))
}

func (m *ICefCallback) Cancel() {
	Proc(internale_cefCallback_Cancel).Call(uintptr(m.instance))
}
