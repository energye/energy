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

// ICefStringMultimapRef Parent: ICefStringMultimap
type ICefStringMultimapRef interface {
	ICefStringMultimap
}

// TCefStringMultimapRef Parent: TCefStringMultimap
type TCefStringMultimapRef struct {
	TCefStringMultimap
}

func NewCefStringMultimapRef(aHandle TCefStringMultimapHandle) ICefStringMultimapRef {
	r1 := CEF().SysCallN(1406, uintptr(aHandle))
	return AsCefStringMultimapRef(r1)
}
