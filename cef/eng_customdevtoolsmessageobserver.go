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

// ICustomDevToolsMessageObserver Parent: ICefDevToolsMessageObserver
type ICustomDevToolsMessageObserver interface {
	ICefDevToolsMessageObserver
}

// TCustomDevToolsMessageObserver Parent: TCefDevToolsMessageObserver
type TCustomDevToolsMessageObserver struct {
	TCefDevToolsMessageObserver
}

func NewCustomDevToolsMessageObserver(events IChromiumEvents) ICustomDevToolsMessageObserver {
	r1 := CEF().SysCallN(2137, GetObjectUintptr(events))
	return AsCustomDevToolsMessageObserver(r1)
}
