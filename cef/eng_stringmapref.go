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

// ICefStringMapRef Parent: ICefStringMap
type ICefStringMapRef interface {
	ICefStringMap
}

// TCefStringMapRef Parent: TCefStringMap
type TCefStringMapRef struct {
	TCefStringMap
}

func NewCefStringMapRef(aHandle TCefStringMapHandle) ICefStringMapRef {
	r1 := CEF().SysCallN(1398, uintptr(aHandle))
	return AsCefStringMapRef(r1)
}
