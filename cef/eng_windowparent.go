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

// ICEFWindowParent Parent: ICEFWinControl
type ICEFWindowParent interface {
	ICEFWinControl
}

// TCEFWindowParent Parent: TCEFWinControl
type TCEFWindowParent struct {
	TCEFWinControl
}

func NewCEFWindowParent(theOwner IComponent) ICEFWindowParent {
	r1 := CEF().SysCallN(404, GetObjectUintptr(theOwner))
	return AsCEFWindowParent(r1)
}

func CEFWindowParentClass() TClass {
	ret := CEF().SysCallN(403)
	return TClass(ret)
}
