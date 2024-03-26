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

// ICoreWebView2Frame Parent: IObject
//
//	ICoreWebView2Frame provides direct access to the iframes information.
//	You can get an ICoreWebView2Frame by handling the ICoreWebView2_4.add_FrameCreated event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame">See the ICoreWebView2Frame article.</a>
type ICoreWebView2Frame interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Frame // property
	// FrameID
	//  The unique identifier of the current frame. It's the same kind of ID as
	//  with the `FrameId` in `ICoreWebView2` and via `CoreWebView2FrameInfo`.
	FrameID() uint32 // property
	// Name
	//  The name of the iframe from the iframe html tag declaring it.
	//  You can access this property even if the iframe is destroyed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#get_name">See the ICoreWebView2Frame article.</a>
	Name() string // property
	// IsDestroyed
	//  Check whether a frame is destroyed. Returns true during
	//  the Destroyed event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#isdestroyed">See the ICoreWebView2Frame article.</a>
	IsDestroyed() bool // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// RemoveHostObjectFromScript
	//  Remove the host object specified by the name so that it is no longer
	//  accessible from JavaScript code in the iframe. While new access
	//  attempts are denied, if the object is already obtained by JavaScript code
	//  in the iframe, the JavaScript code continues to have access to that
	//  object. Calling this method for a name that is already removed or was
	//  never added fails. If the iframe is destroyed this method will return fail
	//  also.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#removehostobjectfromscript">See the ICoreWebView2Frame article.</a>
	RemoveHostObjectFromScript(aName string) bool // function
	// ExecuteScript
	//  Run JavaScript code from the javascript parameter in the current frame.
	//  The result of evaluating the provided JavaScript is passed to the completion handler.
	//  The result value is a JSON encoded string. If the result is undefined,
	//  contains a reference cycle, or otherwise is not able to be encoded into
	//  JSON, then the result is considered to be null, which is encoded
	//  in JSON as the string "null".
	//  NOTE: A function that has no explicit return value returns undefined. If the
	//  script that was run throws an unhandled exception, then the result is
	//  also "null". This method is applied asynchronously. If the method is
	//  run before `ContentLoading`, the script will not be executed
	//  and the string "null" will be returned.
	//  This operation executes the script even if `ICoreWebView2Settings::IsScriptEnabled` is
	//  set to `FALSE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#executescript">See the ICoreWebView2Frame2 article.</a>
	ExecuteScript(JavaScript string, aExecutionID int32, aBrowserComponent IComponent) bool // function
	// PostWebMessageAsJson
	//  Posts the specified webMessage to the frame.
	//  The frame receives the message by subscribing to the `message` event of
	//  the `window.chrome.webview` of the frame document.
	//
	//  <code>
	//  ```cpp
	//  window.chrome.webview.addEventListener('message', handler)
	//  window.chrome.webview.removeEventListener('message', handler)
	//  ```</code>
	//
	//  The event args is an instances of `MessageEvent`. The
	//  `ICoreWebView2Settings::IsWebMessageEnabled` setting must be `TRUE` or
	//  the message will not be sent. The `data` property of the event
	//  args is the `webMessage` string parameter parsed as a JSON string into a
	//  JavaScript object. The `source` property of the event args is a reference
	//  to the `window.chrome.webview` object. For information about sending
	//  messages from the HTML document in the WebView to the host, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  The message is delivered asynchronously. If a navigation occurs before the
	//  message is posted to the page, the message is discarded.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#postwebmessageasjson">See the ICoreWebView2Frame2 article.</a>
	PostWebMessageAsJson(aWebMessageAsJson string) bool // function
	// PostWebMessageAsString
	//  Posts a message that is a simple string rather than a JSON string
	//  representation of a JavaScript object. This behaves in exactly the same
	//  manner as `PostWebMessageAsJson`, but the `data` property of the event
	//  args of the `window.chrome.webview` message is a string with the same
	//  value as `webMessageAsString`. Use this instead of
	//  `PostWebMessageAsJson` if you want to communicate using simple strings
	//  rather than JSON objects.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#postwebmessageasstring">See the ICoreWebView2Frame2 article.</a>
	PostWebMessageAsString(aWebMessageAsString string) bool // function
	// PostSharedBufferToScript
	//  Share a shared buffer object with script of the iframe in the WebView.
	//  The script will receive a `sharedbufferreceived` event from chrome.webview.
	//  The event arg for that event will have the following methods and properties:
	//   `getBuffer()`: return an ArrayBuffer object with the backing content from the shared buffer.
	//   `additionalData`: an object as the result of parsing `additionalDataAsJson` as JSON string.
	//  This property will be `undefined` if `additionalDataAsJson` is nullptr or empty string.
	//   `source`: with a value set as `chrome.webview` object.
	//  If a string is provided as `additionalDataAsJson` but it is not a valid JSON string,
	//  the API will fail with `E_INVALIDARG`.
	//  If `access` is COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY, the script will only have read access to the buffer.
	//  If the script tries to modify the content in a read only buffer, it will cause an access
	//  violation in WebView renderer process and crash the renderer process.
	//  If the shared buffer is already closed, the API will fail with `RO_E_CLOSED`.
	//  The script code should call `chrome.webview.releaseBuffer` with
	//  the shared buffer as the parameter to release underlying resources as soon
	//  as it does not need access to the shared buffer any more.
	//  The application can post the same shared buffer object to multiple web pages or iframes, or
	//  post to the same web page or iframe multiple times. Each `PostSharedBufferToScript` will
	//  create a separate ArrayBuffer object with its own view of the memory and is separately
	//  released. The underlying shared memory will be released when all the views are released.
	//  For example, if we want to send data to script for one time read only consumption.
	//  Sharing a buffer to script has security risk. You should only share buffer with trusted site.
	//  If a buffer is shared to a untrusted site, possible sensitive information could be leaked.
	//  If a buffer is shared as modifiable by the script and the script modifies it in an unexpected way,
	//  it could result in corrupted data that might even crash the application.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame4#postsharedbuffertoscript">See the ICoreWebView2Frame4 article.</a>
	PostSharedBufferToScript(aSharedBuffer ICoreWebView2SharedBuffer, aAccess TWVSharedBufferAccess, aAdditionalDataAsJson string) bool // function
}

// TCoreWebView2Frame Parent: TObject
//
//	ICoreWebView2Frame provides direct access to the iframes information.
//	You can get an ICoreWebView2Frame by handling the ICoreWebView2_4.add_FrameCreated event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame">See the ICoreWebView2Frame article.</a>
type TCoreWebView2Frame struct {
	TObject
}

func NewCoreWebView2Frame(aBaseIntf ICoreWebView2Frame) ICoreWebView2Frame {
	r1 := WV().SysCallN(342, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Frame(r1)
}

func (m *TCoreWebView2Frame) Initialized() bool {
	r1 := WV().SysCallN(345, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) BaseIntf() ICoreWebView2Frame {
	var resultCoreWebView2Frame uintptr
	WV().SysCallN(340, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Frame)))
	return AsCoreWebView2Frame(resultCoreWebView2Frame)
}

func (m *TCoreWebView2Frame) FrameID() uint32 {
	r1 := WV().SysCallN(344, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2Frame) Name() string {
	r1 := WV().SysCallN(347, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Frame) IsDestroyed() bool {
	r1 := WV().SysCallN(346, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(339, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) RemoveHostObjectFromScript(aName string) bool {
	r1 := WV().SysCallN(351, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) ExecuteScript(JavaScript string, aExecutionID int32, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(343, m.Instance(), PascalStr(JavaScript), uintptr(aExecutionID), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) PostWebMessageAsJson(aWebMessageAsJson string) bool {
	r1 := WV().SysCallN(349, m.Instance(), PascalStr(aWebMessageAsJson))
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) PostWebMessageAsString(aWebMessageAsString string) bool {
	r1 := WV().SysCallN(350, m.Instance(), PascalStr(aWebMessageAsString))
	return GoBool(r1)
}

func (m *TCoreWebView2Frame) PostSharedBufferToScript(aSharedBuffer ICoreWebView2SharedBuffer, aAccess TWVSharedBufferAccess, aAdditionalDataAsJson string) bool {
	r1 := WV().SysCallN(348, m.Instance(), GetObjectUintptr(aSharedBuffer), uintptr(aAccess), PascalStr(aAdditionalDataAsJson))
	return GoBool(r1)
}

func CoreWebView2FrameClass() TClass {
	ret := WV().SysCallN(341)
	return TClass(ret)
}
