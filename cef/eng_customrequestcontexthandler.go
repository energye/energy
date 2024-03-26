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

// ICustomRequestContextHandler Parent: ICefRequestContextHandler
type ICustomRequestContextHandler interface {
	ICefRequestContextHandler
	RemoveReferences() // procedure
}

// TCustomRequestContextHandler Parent: TCefRequestContextHandler
type TCustomRequestContextHandler struct {
	TCefRequestContextHandler
}

func NewCustomRequestContextHandler(events IChromiumEvents) ICustomRequestContextHandler {
	r1 := CEF().SysCallN(2157, GetObjectUintptr(events))
	return AsCustomRequestContextHandler(r1)
}

func (m *TCustomRequestContextHandler) RemoveReferences() {
	CEF().SysCallN(2158, m.Instance())
}
