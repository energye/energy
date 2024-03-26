//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2Deferral Parent: IObject
//
//	This interface is used to complete deferrals on event args that support
//	getting deferrals using the GetDeferral method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2deferral">See the ICoreWebView2Deferral article.</a>
type ICoreWebView2Deferral interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Deferral // property
	// Complete
	//  Completes the associated deferred event. Complete should only be run
	//  once for each deferral taken.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2deferral#complete">See the ICoreWebView2Deferral article.</a>
	Complete() bool // function
}

// TCoreWebView2Deferral Parent: TObject
//
//	This interface is used to complete deferrals on event args that support
//	getting deferrals using the GetDeferral method.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2deferral">See the ICoreWebView2Deferral article.</a>
type TCoreWebView2Deferral struct {
	TObject
}

func NewCoreWebView2Deferral(aBaseIntf ICoreWebView2Deferral) ICoreWebView2Deferral {
	r1 := WV().SysCallN(243, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Deferral(r1)
}

func (m *TCoreWebView2Deferral) Initialized() bool {
	r1 := WV().SysCallN(244, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Deferral) BaseIntf() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(240, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func (m *TCoreWebView2Deferral) Complete() bool {
	r1 := WV().SysCallN(242, m.Instance())
	return GoBool(r1)
}

func CoreWebView2DeferralClass() TClass {
	ret := WV().SysCallN(241)
	return TClass(ret)
}
