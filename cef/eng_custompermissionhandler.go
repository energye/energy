//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICustomPermissionHandler Parent: ICefPermissionHandler
type ICustomPermissionHandler interface {
	ICefPermissionHandler
}

// TCustomPermissionHandler Parent: TCefPermissionHandler
type TCustomPermissionHandler struct {
	TCefPermissionHandler
}

func NewCustomPermissionHandler(events IChromiumEvents) ICustomPermissionHandler {
	r1 := CEF().SysCallN(2153, GetObjectUintptr(events))
	return AsCustomPermissionHandler(r1)
}
