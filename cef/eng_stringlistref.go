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

// ICefStringListRef Parent: ICefCustomStringList
type ICefStringListRef interface {
	ICefCustomStringList
}

// TCefStringListRef Parent: TCefCustomStringList
type TCefStringListRef struct {
	TCefCustomStringList
}

func NewCefStringListRef(aHandle TCefStringList) ICefStringListRef {
	r1 := CEF().SysCallN(1394, uintptr(aHandle))
	return AsCefStringListRef(r1)
}
