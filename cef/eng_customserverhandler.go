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

// ICustomServerHandler Parent: ICEFServerHandler
type ICustomServerHandler interface {
	ICEFServerHandler
}

// TCustomServerHandler Parent: TCEFServerHandler
type TCustomServerHandler struct {
	TCEFServerHandler
}

func NewCustomServerHandler(events IServerEvents) ICustomServerHandler {
	r1 := CEF().SysCallN(2166, GetObjectUintptr(events))
	return AsCustomServerHandler(r1)
}
