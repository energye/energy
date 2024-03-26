//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/v2/api"

// ChangeCompositionRange windows
func (m *TBufferPanel) ChangeCompositionRange(selectionrange *TCefRange, characterbounds TCefRectDynArray) {
	api.CEFPreDef().SysCallN(3, m.Instance(), uintptr(unsafePointer(selectionrange)), uintptr(unsafePointer(&characterbounds[0])), uintptr(int32(len(characterbounds))))
}
