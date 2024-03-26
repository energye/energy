//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import "github.com/energye/energy/v2/api"

// GetAlls
//
//	Returns all Displays. Mirrored displays are excluded; this function is
//	intended to return distinct, usable displays.
func (m *display) GetAlls(aDisplayArray *ICefDisplayArray) bool {
	var result uintptr
	var size int32
	r1 := api.CEFPreDef().SysCallN(7, uintptr(unsafePointer(&result)), uintptr(unsafePointer(&result)))
	*aDisplayArray = DisplayArrayRef.New(int(size), result)
	return api.GoBool(r1)
}
