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

// ICustomResourceRequestHandler Parent: ICefResourceRequestHandler
type ICustomResourceRequestHandler interface {
	ICefResourceRequestHandler
	RemoveReferences() // procedure
}

// TCustomResourceRequestHandler Parent: TCefResourceRequestHandler
type TCustomResourceRequestHandler struct {
	TCefResourceRequestHandler
}

func NewCustomResourceRequestHandler(events IChromiumEvents) ICustomResourceRequestHandler {
	r1 := CEF().SysCallN(2161, GetObjectUintptr(events))
	return AsCustomResourceRequestHandler(r1)
}

func (m *TCustomResourceRequestHandler) RemoveReferences() {
	CEF().SysCallN(2162, m.Instance())
}
