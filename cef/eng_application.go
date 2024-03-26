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

// ICefApplication Parent: ICefApplicationCore
//
//	Main class used to simplify the CEF initialization and destruction.
type ICefApplication interface {
	ICefApplicationCore
	DestroyApplicationObject() bool          // property
	SetDestroyApplicationObject(AValue bool) // property
	DestroyAppWindows() bool                 // property
	SetDestroyAppWindows(AValue bool)        // property
}

// TCefApplication Parent: TCefApplicationCore
//
//	Main class used to simplify the CEF initialization and destruction.
type TCefApplication struct {
	TCefApplicationCore
}

func NewCefApplication() ICefApplication {
	r1 := CEF().SysCallN(581)
	return AsCefApplication(r1)
}

func (m *TCefApplication) DestroyApplicationObject() bool {
	r1 := CEF().SysCallN(583, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplication) SetDestroyApplicationObject(AValue bool) {
	CEF().SysCallN(583, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCefApplication) DestroyAppWindows() bool {
	r1 := CEF().SysCallN(582, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCefApplication) SetDestroyAppWindows(AValue bool) {
	CEF().SysCallN(582, 1, m.Instance(), PascalBool(AValue))
}

func CefApplicationClass() TClass {
	ret := CEF().SysCallN(580)
	return TClass(ret)
}
