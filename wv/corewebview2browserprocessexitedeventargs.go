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

// ICoreWebView2BrowserProcessExitedEventArgs Parent: IObject
//
//	Event args for the BrowserProcessExited event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
type ICoreWebView2BrowserProcessExitedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2BrowserProcessExitedEventArgs // property
	// BrowserProcessExitKind
	//  The kind of browser process exit that has occurred.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs#get_browserprocessexitkind">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
	BrowserProcessExitKind() TWVBrowserProcessExitKind // property
	// BrowserProcessId
	//  The process ID of the browser process that has exited.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs#get_browserprocessid">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
	BrowserProcessId() uint32 // property
}

// TCoreWebView2BrowserProcessExitedEventArgs Parent: TObject
//
//	Event args for the BrowserProcessExited event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserprocessexitedeventargs">See the ICoreWebView2BrowserProcessExitedEventArgs article.</a>
type TCoreWebView2BrowserProcessExitedEventArgs struct {
	TObject
}

func NewCoreWebView2BrowserProcessExitedEventArgs(aArgs ICoreWebView2BrowserProcessExitedEventArgs) ICoreWebView2BrowserProcessExitedEventArgs {
	r1 := WV().SysCallN(53, GetObjectUintptr(aArgs))
	return AsCoreWebView2BrowserProcessExitedEventArgs(r1)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(54, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BaseIntf() ICoreWebView2BrowserProcessExitedEventArgs {
	var resultCoreWebView2BrowserProcessExitedEventArgs uintptr
	WV().SysCallN(49, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2BrowserProcessExitedEventArgs)))
	return AsCoreWebView2BrowserProcessExitedEventArgs(resultCoreWebView2BrowserProcessExitedEventArgs)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BrowserProcessExitKind() TWVBrowserProcessExitKind {
	r1 := WV().SysCallN(50, m.Instance())
	return TWVBrowserProcessExitKind(r1)
}

func (m *TCoreWebView2BrowserProcessExitedEventArgs) BrowserProcessId() uint32 {
	r1 := WV().SysCallN(51, m.Instance())
	return uint32(r1)
}

func CoreWebView2BrowserProcessExitedEventArgsClass() TClass {
	ret := WV().SysCallN(52)
	return TClass(ret)
}
