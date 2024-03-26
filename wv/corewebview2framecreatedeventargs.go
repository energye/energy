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

// ICoreWebView2FrameCreatedEventArgs Parent: IObject
//
//	Event args for the FrameCreated events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs">See the ICoreWebView2FrameCreatedEventArgs article.</a>
type ICoreWebView2FrameCreatedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2FrameCreatedEventArgs // property
	// Frame
	//  The frame which was created.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs#get_frame">See the ICoreWebView2FrameCreatedEventArgs article.</a>
	Frame() ICoreWebView2Frame // property
}

// TCoreWebView2FrameCreatedEventArgs Parent: TObject
//
//	Event args for the FrameCreated events.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2framecreatedeventargs">See the ICoreWebView2FrameCreatedEventArgs article.</a>
type TCoreWebView2FrameCreatedEventArgs struct {
	TObject
}

func NewCoreWebView2FrameCreatedEventArgs(aArgs ICoreWebView2FrameCreatedEventArgs) ICoreWebView2FrameCreatedEventArgs {
	r1 := WV().SysCallN(314, GetObjectUintptr(aArgs))
	return AsCoreWebView2FrameCreatedEventArgs(r1)
}

func (m *TCoreWebView2FrameCreatedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(316, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2FrameCreatedEventArgs) BaseIntf() ICoreWebView2FrameCreatedEventArgs {
	var resultCoreWebView2FrameCreatedEventArgs uintptr
	WV().SysCallN(312, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameCreatedEventArgs)))
	return AsCoreWebView2FrameCreatedEventArgs(resultCoreWebView2FrameCreatedEventArgs)
}

func (m *TCoreWebView2FrameCreatedEventArgs) Frame() ICoreWebView2Frame {
	var resultCoreWebView2Frame uintptr
	WV().SysCallN(315, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Frame)))
	return AsCoreWebView2Frame(resultCoreWebView2Frame)
}

func CoreWebView2FrameCreatedEventArgsClass() TClass {
	ret := WV().SysCallN(313)
	return TClass(ret)
}
