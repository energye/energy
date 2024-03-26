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

// ICustomClientHandler Parent: ICefClient
type ICustomClientHandler interface {
	ICefClient
	RemoveReferences() // procedure
}

// TCustomClientHandler Parent: TCefClient
type TCustomClientHandler struct {
	TCefClient
}

func NewCustomClientHandler(events IChromiumEvents, aDevToolsClient bool) ICustomClientHandler {
	r1 := CEF().SysCallN(2133, GetObjectUintptr(events), PascalBool(aDevToolsClient))
	return AsCustomClientHandler(r1)
}

func (m *TCustomClientHandler) RemoveReferences() {
	CEF().SysCallN(2134, m.Instance())
}
