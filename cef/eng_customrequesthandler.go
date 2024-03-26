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

// ICustomRequestHandler Parent: ICefRequestHandler
type ICustomRequestHandler interface {
	ICefRequestHandler
	RemoveReferences() // procedure
}

// TCustomRequestHandler Parent: TCefRequestHandler
type TCustomRequestHandler struct {
	TCefRequestHandler
}

func NewCustomRequestHandler(events IChromiumEvents) ICustomRequestHandler {
	r1 := CEF().SysCallN(2159, GetObjectUintptr(events))
	return AsCustomRequestHandler(r1)
}

func (m *TCustomRequestHandler) RemoveReferences() {
	CEF().SysCallN(2160, m.Instance())
}
