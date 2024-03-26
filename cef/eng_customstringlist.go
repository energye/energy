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

// ICefCustomStringList Parent: IObject
//
//	CEF string maps are a set of key/value string pairs.
type ICefCustomStringList interface {
	IObject
}

// TCefCustomStringList Parent: TObject
//
//	CEF string maps are a set of key/value string pairs.
type TCefCustomStringList struct {
	TObject
}

func NewCefCustomStringList() ICefCustomStringList {
	r1 := CEF().SysCallN(788)
	return AsCefCustomStringList(r1)
}
