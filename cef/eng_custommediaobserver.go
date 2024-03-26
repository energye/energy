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

// ICustomMediaObserver Parent: ICefMediaObserver
type ICustomMediaObserver interface {
	ICefMediaObserver
}

// TCustomMediaObserver Parent: TCefMediaObserver
type TCustomMediaObserver struct {
	TCefMediaObserver
}

func NewCustomMediaObserver(events IChromiumEvents) ICustomMediaObserver {
	r1 := CEF().SysCallN(2150, GetObjectUintptr(events))
	return AsCustomMediaObserver(r1)
}
