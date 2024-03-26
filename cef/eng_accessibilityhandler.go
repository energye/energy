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

// IAccessibilityHandler Parent: ICEFAccessibilityHandler
//
//	Implement this interface to receive accessibility notification when
//	accessibility events have been registered. The functions of this interface
//	will be called on the UI thread.
//	<a cref="uCEFTypes|TCefAccessibilityHandler">Implements TCefAccessibilityHandler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_accessibility_handler_capi.h">CEF source file: /include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t)</a>
type IAccessibilityHandler interface {
	ICEFAccessibilityHandler
	// SetOnTreeChange
	//  Called after renderer process sends accessibility tree changes to the
	//  browser process.
	SetOnTreeChange(fn TOnAccessibility) // property event
	// SetOnLocationChange
	//  Called after renderer process sends accessibility location changes to the
	//  browser process.
	SetOnLocationChange(fn TOnAccessibility) // property event
}

// TAccessibilityHandler Parent: TCEFAccessibilityHandler
//
//	Implement this interface to receive accessibility notification when
//	accessibility events have been registered. The functions of this interface
//	will be called on the UI thread.
//	<a cref="uCEFTypes|TCefAccessibilityHandler">Implements TCefAccessibilityHandler</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_accessibility_handler_capi.h">CEF source file: /include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t)</a>
type TAccessibilityHandler struct {
	TCEFAccessibilityHandler
	treeChangePtr     uintptr
	locationChangePtr uintptr
}

func NewAccessibilityHandler() IAccessibilityHandler {
	r1 := CEF().SysCallN(1)
	return AsAccessibilityHandler(r1)
}

func AccessibilityHandlerClass() TClass {
	ret := CEF().SysCallN(0)
	return TClass(ret)
}

func (m *TAccessibilityHandler) SetOnTreeChange(fn TOnAccessibility) {
	if m.treeChangePtr != 0 {
		RemoveEventElement(m.treeChangePtr)
	}
	m.treeChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(3, m.Instance(), m.treeChangePtr)
}

func (m *TAccessibilityHandler) SetOnLocationChange(fn TOnAccessibility) {
	if m.locationChangePtr != 0 {
		RemoveEventElement(m.locationChangePtr)
	}
	m.locationChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2, m.Instance(), m.locationChangePtr)
}
