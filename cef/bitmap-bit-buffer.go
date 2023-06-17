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
	"unsafe"
)

type TCEFBitmapBitBuffer struct {
	instance unsafe.Pointer
}

func (m *TCEFBitmapBitBuffer) Free() {
	if m.instance != nil {
		//imports.SysCallN(def.BufferPanel_Free, m.Instance())
		m.instance = nil
	}
}

func (m *TCEFBitmapBitBuffer) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *TCEFBitmapBitBuffer) IsValid() bool {
	return m.instance != nil
}
