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

// ICoreWebView2NewWindowRequestedEventArgs Parent: IObject
//
//	Event args for the NewWindowRequested event.  The event is run when
//	content inside webview requested to a open a new window (through
//	window.open() and so on).
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
type ICoreWebView2NewWindowRequestedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2NewWindowRequestedEventArgs // property
	// URI
	//  The target uri of the new window requested.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_uri">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
	URI() string // property
	// NewWindow
	//  Gets the new window.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_newwindow">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
	NewWindow() ICoreWebView2 // property
	// SetNewWindow Set NewWindow
	SetNewWindow(AValue ICoreWebView2) // property
	// Handled
	//  Gets whether the `NewWindowRequested` event is handled by host.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_handled">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// IsUserInitiated
	//  `TRUE` when the new window request was initiated through a user gesture.
	//  Examples of user initiated requests are:
	//  - Selecting an anchor tag with target
	//  - Programmatic window open from a script that directly run as a result of
	//  user interaction such as via onclick handlers.
	//  Non-user initiated requests are programmatic window opens from a script
	//  that are not directly triggered by user interaction, such as those that
	//  run while loading a new page or via timers.
	//  The Microsoft Edge popup blocker is disabled for WebView so the app is
	//  able to use this flag to block non-user initiated popups.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_isuserinitiated">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
	IsUserInitiated() bool // property
	// Deferral
	//  Obtain an `ICoreWebView2Deferral` object and put the event into a
	//  deferred state. Use the `ICoreWebView2Deferral` object to complete the
	//  window open request at a later time. While this event is deferred the
	//  opener window returns a `WindowProxy` to an un-navigated window, which
	//  navigates when the deferral is complete.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#getdeferral">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
	Deferral() ICoreWebView2Deferral // property
	// WindowFeatures
	//  Window features specified by the `window.open`. The features should be
	//  considered for positioning and sizing of new webview windows.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs#get_windowfeatures">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
	WindowFeatures() ICoreWebView2WindowFeatures // property
	// Name
	//  Gets the name of the new window. This window can be created via `window.open(url, windowName)`,
	//  where the windowName parameter corresponds to `Name` property.
	//  If no windowName is passed to `window.open`, then the `Name` property
	//  will be set to an empty string. Additionally, if window is opened through other means,
	//  such as `<a target="windowName">...</a>` or `<iframe name="windowName">...</iframe>`,
	//  then the `Name` property will be set accordingly. In the case of target=_blank,
	//  the `Name` property will be an empty string.
	//  Opening a window via ctrl+clicking a link would result in the `Name` property
	//  being set to an empty string.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs2#get_name">See the ICoreWebView2NewWindowRequestedEventArgs2 article.</a>
	Name() string // property
	// OriginalSourceFrameInfo
	//  The frame info of the frame where the new window request originated. The
	//  `OriginalSourceFrameInfo` is a snapshot of frame information at the time when the
	//  new window was requested. See `ICoreWebView2FrameInfo` for details on frame
	//  properties.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs3#get_originalsourceframeinfo">See the ICoreWebView2NewWindowRequestedEventArgs3 article.</a>
	OriginalSourceFrameInfo() ICoreWebView2FrameInfo // property
}

// TCoreWebView2NewWindowRequestedEventArgs Parent: TObject
//
//	Event args for the NewWindowRequested event.  The event is run when
//	content inside webview requested to a open a new window (through
//	window.open() and so on).
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2newwindowrequestedeventargs">See the ICoreWebView2NewWindowRequestedEventArgs article.</a>
type TCoreWebView2NewWindowRequestedEventArgs struct {
	TObject
}

func NewCoreWebView2NewWindowRequestedEventArgs(aArgs ICoreWebView2NewWindowRequestedEventArgs) ICoreWebView2NewWindowRequestedEventArgs {
	r1 := WV().SysCallN(419, GetObjectUintptr(aArgs))
	return AsCoreWebView2NewWindowRequestedEventArgs(r1)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(422, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) BaseIntf() ICoreWebView2NewWindowRequestedEventArgs {
	var resultCoreWebView2NewWindowRequestedEventArgs uintptr
	WV().SysCallN(417, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2NewWindowRequestedEventArgs)))
	return AsCoreWebView2NewWindowRequestedEventArgs(resultCoreWebView2NewWindowRequestedEventArgs)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) URI() string {
	r1 := WV().SysCallN(427, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) NewWindow() ICoreWebView2 {
	var resultCoreWebView2 uintptr
	WV().SysCallN(425, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2)))
	return AsCoreWebView2(resultCoreWebView2)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) SetNewWindow(AValue ICoreWebView2) {
	WV().SysCallN(425, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Handled() bool {
	r1 := WV().SysCallN(421, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) SetHandled(AValue bool) {
	WV().SysCallN(421, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) IsUserInitiated() bool {
	r1 := WV().SysCallN(423, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Deferral() ICoreWebView2Deferral {
	var resultCoreWebView2Deferral uintptr
	WV().SysCallN(420, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Deferral)))
	return AsCoreWebView2Deferral(resultCoreWebView2Deferral)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) WindowFeatures() ICoreWebView2WindowFeatures {
	var resultCoreWebView2WindowFeatures uintptr
	WV().SysCallN(428, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2WindowFeatures)))
	return AsCoreWebView2WindowFeatures(resultCoreWebView2WindowFeatures)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) Name() string {
	r1 := WV().SysCallN(424, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2NewWindowRequestedEventArgs) OriginalSourceFrameInfo() ICoreWebView2FrameInfo {
	var resultCoreWebView2FrameInfo uintptr
	WV().SysCallN(426, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2FrameInfo)))
	return AsCoreWebView2FrameInfo(resultCoreWebView2FrameInfo)
}

func CoreWebView2NewWindowRequestedEventArgsClass() TClass {
	ret := WV().SysCallN(418)
	return TClass(ret)
}
