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

// ICefBrowser Parent: ICefBaseRefCounted
//
//	Interface used to represent a browser. When used in the browser process the functions of this interface may be called on any thread unless otherwise indicated in the comments. When used in the render process the functions of this interface may only be called on the main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_browser_t))</a>
type ICefBrowser interface {
	ICefBaseRefCounted
	// IsValid
	//  True if this object is currently valid. This will return false (0) after ICefLifeSpanHandler.OnBeforeClose is called.
	IsValid() bool // function
	// GetHost
	//  Returns the browser host object. This function can only be called in the browser process.
	GetHost() ICefBrowserHost // function
	// CanGoBack
	//  Returns true (1) if the browser can navigate backwards.
	CanGoBack() bool // function
	// CanGoForward
	//  Returns true (1) if the browser can navigate forwards.
	CanGoForward() bool // function
	// IsLoading
	//  Returns true (1) if the browser is currently loading.
	IsLoading() bool // function
	// GetIdentifier
	//  Returns the globally unique identifier for this browser. This value is also used as the tabId for extension APIs.
	GetIdentifier() int32 // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefBrowser) bool // function
	// IsPopup
	//  Returns true (1) if the browser is a popup.
	IsPopup() bool // function
	// HasDocument
	//  Returns true (1) if a document has been loaded in the browser.
	HasDocument() bool // function
	// GetMainFrame
	//  Returns the main (top-level) frame for the browser. In the browser process this will return a valid object until after ICefLifeSpanHandler.OnBeforeClose is called. In the renderer process this will return NULL if the main frame is hosted in a different renderer process (e.g. for cross-origin sub-frames). The main frame object will change during cross-origin navigation or re-navigation after renderer process termination (due to crashes, etc).
	GetMainFrame() ICefFrame // function
	// GetFocusedFrame
	//  Returns the focused frame for the browser.
	GetFocusedFrame() ICefFrame // function
	// GetFrameByident
	//  Returns the frame with the specified identifier, or NULL if not found.
	GetFrameByident(identifier int64) ICefFrame // function
	// GetFrame
	//  Returns the frame with the specified name, or NULL if not found.
	GetFrame(name string) ICefFrame // function
	// GetFrameCount
	//  Returns the number of frames that currently exist.
	GetFrameCount() NativeUInt // function
	// GetFrameIdentifiers
	//  Returns the identifiers of all existing frames.
	GetFrameIdentifiers(aFrameCount *NativeUInt, aFrameIdentifierArray *ICefFrameIdentifierArray) bool // function
	// GetFrameNames
	//  Returns the names of all existing frames.
	GetFrameNames(aFrameNames *IStrings) bool // function
	// GoBack
	//  Navigate backwards.
	GoBack() // procedure
	// GoForward
	//  Navigate forwards.
	GoForward() // procedure
	// Reload
	//  Reload the current page.
	Reload() // procedure
	// ReloadIgnoreCache
	//  Reload the current page ignoring any cached data.
	ReloadIgnoreCache() // procedure
	// StopLoad
	//  Stop loading the page.
	StopLoad() // procedure
}

// TCefBrowser Parent: TCefBaseRefCounted
//
//	Interface used to represent a browser. When used in the browser process the functions of this interface may be called on any thread unless otherwise indicated in the comments. When used in the render process the functions of this interface may only be called on the main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_browser_t))</a>
type TCefBrowser struct {
	TCefBaseRefCounted
}

// BrowserRef -> ICefBrowser
var BrowserRef browser

// browser TCefBrowser Ref
type browser uintptr

func (m *browser) UnWrap(data uintptr) ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(702, uintptr(data), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefBrowser) IsValid() bool {
	r1 := CEF().SysCallN(698, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) GetHost() ICefBrowserHost {
	var resultCefBrowserHost uintptr
	CEF().SysCallN(689, m.Instance(), uintptr(unsafePointer(&resultCefBrowserHost)))
	return AsCefBrowserHost(resultCefBrowserHost)
}

func (m *TCefBrowser) CanGoBack() bool {
	r1 := CEF().SysCallN(681, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) CanGoForward() bool {
	r1 := CEF().SysCallN(682, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) IsLoading() bool {
	r1 := CEF().SysCallN(695, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) GetIdentifier() int32 {
	r1 := CEF().SysCallN(690, m.Instance())
	return int32(r1)
}

func (m *TCefBrowser) IsSame(that ICefBrowser) bool {
	r1 := CEF().SysCallN(697, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefBrowser) IsPopup() bool {
	r1 := CEF().SysCallN(696, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) HasDocument() bool {
	r1 := CEF().SysCallN(694, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowser) GetMainFrame() ICefFrame {
	var resultCefFrame uintptr
	CEF().SysCallN(691, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFocusedFrame() ICefFrame {
	var resultCefFrame uintptr
	CEF().SysCallN(683, m.Instance(), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFrameByident(identifier int64) ICefFrame {
	var resultCefFrame uintptr
	CEF().SysCallN(685, m.Instance(), uintptr(unsafePointer(&identifier)), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFrame(name string) ICefFrame {
	var resultCefFrame uintptr
	CEF().SysCallN(684, m.Instance(), PascalStr(name), uintptr(unsafePointer(&resultCefFrame)))
	return AsCefFrame(resultCefFrame)
}

func (m *TCefBrowser) GetFrameCount() NativeUInt {
	r1 := CEF().SysCallN(686, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefBrowser) GetFrameIdentifiers(aFrameCount *NativeUInt, aFrameIdentifierArray *ICefFrameIdentifierArray) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := CEF().SysCallN(687, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*aFrameCount = NativeUInt(result0)
	*aFrameIdentifierArray = FrameIdentifierArrayRef.New(int(*aFrameCount), result1)
	return GoBool(r1)
}

func (m *TCefBrowser) GetFrameNames(aFrameNames *IStrings) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(688, m.Instance(), uintptr(unsafePointer(&result0)))
	*aFrameNames = AsStrings(result0)
	return GoBool(r1)
}

func (m *TCefBrowser) GoBack() {
	CEF().SysCallN(692, m.Instance())
}

func (m *TCefBrowser) GoForward() {
	CEF().SysCallN(693, m.Instance())
}

func (m *TCefBrowser) Reload() {
	CEF().SysCallN(699, m.Instance())
}

func (m *TCefBrowser) ReloadIgnoreCache() {
	CEF().SysCallN(700, m.Instance())
}

func (m *TCefBrowser) StopLoad() {
	CEF().SysCallN(701, m.Instance())
}
