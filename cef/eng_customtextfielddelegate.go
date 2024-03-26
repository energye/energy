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

// ICustomTextfieldDelegate Parent: ICefTextfieldDelegate
//
//	This class handles all the ICefTextfieldDelegate and ICefViewDelegate methods which call the ICefTextfieldDelegateEvents methods.
//	ICefTextfieldDelegateEvents will be implemented by the control receiving the ICefTextfieldDelegate events.
type ICustomTextfieldDelegate interface {
	ICefTextfieldDelegate
}

// TCustomTextfieldDelegate Parent: TCefTextfieldDelegate
//
//	This class handles all the ICefTextfieldDelegate and ICefViewDelegate methods which call the ICefTextfieldDelegateEvents methods.
//	ICefTextfieldDelegateEvents will be implemented by the control receiving the ICefTextfieldDelegate events.
type TCustomTextfieldDelegate struct {
	TCefTextfieldDelegate
}

func NewCustomTextfieldDelegate(events ICefTextfieldDelegateEvents) ICustomTextfieldDelegate {
	r1 := CEF().SysCallN(2167, GetObjectUintptr(events))
	return AsCustomTextfieldDelegate(r1)
}
