//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/commons"
	"unsafe"
)

type ICefCallback struct {
	instance uintptr
	ptr      unsafe.Pointer
}

func (m *ICefCallback) Cont() {
	Proc("cefCallback_Cont").Call(m.instance)
}

func (m *ICefCallback) Cancel() {
	Proc("cefCallback_Cancel").Call(m.instance)
}
