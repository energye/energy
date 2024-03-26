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

// ICEFBaseScopedWrapperRef Parent: IObject
type ICEFBaseScopedWrapperRef interface {
	IObject
	Wrap() uintptr // function
}

// TCEFBaseScopedWrapperRef Parent: TObject
type TCEFBaseScopedWrapperRef struct {
	TObject
}

func NewCEFBaseScopedWrapperRef(data uintptr) ICEFBaseScopedWrapperRef {
	r1 := CEF().SysCallN(73, uintptr(data))
	return AsCEFBaseScopedWrapperRef(r1)
}

func (m *TCEFBaseScopedWrapperRef) Wrap() uintptr {
	r1 := CEF().SysCallN(74, m.Instance())
	return uintptr(r1)
}
