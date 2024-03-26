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

// ICoreWebView2SharedBuffer Parent: IObject
//
//	The shared buffer object that is created by CreateSharedBuffer.
//	The object is presented to script as ArrayBuffer when posted to script with
//	PostSharedBufferToScript.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer">See the ICoreWebView2SharedBuffer article.</a>
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment12#createsharedbuffer">See the CreateSharedBuffer article.</a>
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_17#postsharedbuffertoscript">See the PostSharedBufferToScript article.</a>
type ICoreWebView2SharedBuffer interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2SharedBuffer // property
	// Size
	//  The size of the shared buffer in bytes.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#get_size">See the ICoreWebView2SharedBuffer article.</a>
	Size() (resultInt64 int64) // property
	// Buffer
	//  The memory address of the shared buffer.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#get_buffer">See the ICoreWebView2SharedBuffer article.</a>
	Buffer() PByte // property
	// OpenStream
	//  Get an IStream object that can be used to access the shared buffer.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#openstream">See the ICoreWebView2SharedBuffer article.</a>
	OpenStream() IStream // property
	// FileMappingHandle
	//  Returns a handle to the file mapping object that backs this shared buffer.
	//  The returned handle is owned by the shared buffer object. You should not
	//  call CloseHandle on it.
	//  Normal app should use `Buffer` or `OpenStream` to get memory address
	//  or IStream object to access the buffer.
	//  For advanced scenarios, you could use file mapping APIs to obtain other views
	//  or duplicate this handle to another application process and create a view from
	//  the duplicated handle in that process to access the buffer from that separate process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#get_filemappinghandle">See the ICoreWebView2SharedBuffer article.</a>
	FileMappingHandle() HANDLE // property
	// Close
	//  Release the backing shared memory. The application should call this API when no
	//  access to the buffer is needed any more, to ensure that the underlying resources
	//  are released timely even if the shared buffer object itself is not released due to
	//  some leaked reference.
	//  After the shared buffer is closed, the buffer address and file mapping handle previously
	//  obtained becomes invalid and cannot be used anymore. Accessing properties of the object
	//  will fail with `RO_E_CLOSED`. Operations like Read or Write on the IStream objects returned
	//  from `OpenStream` will fail with `RO_E_CLOSED`. `PostSharedBufferToScript` will also
	//  fail with `RO_E_CLOSED`.
	//  The script code should call `chrome.webview.releaseBuffer` with
	//  the shared buffer as the parameter to release underlying resources as soon
	//  as it does not need access the shared buffer any more.
	//  When script tries to access the buffer after calling `chrome.webview.releaseBuffer`,
	//  JavaScript `TypeError` exception will be raised complaining about accessing a
	//  detached ArrayBuffer, the same exception when trying to access a transferred ArrayBuffer.
	//  Closing the buffer object on native side doesn't impact access from Script and releasing
	//  the buffer from script doesn't impact access to the buffer from native side.
	//  The underlying shared memory will be released by the OS when both native and script side
	//  release the buffer.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer#close">See the ICoreWebView2SharedBuffer article.</a>
	Close() bool // function
}

// TCoreWebView2SharedBuffer Parent: TObject
//
//	The shared buffer object that is created by CreateSharedBuffer.
//	The object is presented to script as ArrayBuffer when posted to script with
//	PostSharedBufferToScript.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2sharedbuffer">See the ICoreWebView2SharedBuffer article.</a>
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment12#createsharedbuffer">See the CreateSharedBuffer article.</a>
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_17#postsharedbuffertoscript">See the PostSharedBufferToScript article.</a>
type TCoreWebView2SharedBuffer struct {
	TObject
}

func NewCoreWebView2SharedBuffer(aBaseIntf ICoreWebView2SharedBuffer) ICoreWebView2SharedBuffer {
	r1 := WV().SysCallN(638, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2SharedBuffer(r1)
}

func (m *TCoreWebView2SharedBuffer) Initialized() bool {
	r1 := WV().SysCallN(640, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2SharedBuffer) BaseIntf() ICoreWebView2SharedBuffer {
	var resultCoreWebView2SharedBuffer uintptr
	WV().SysCallN(634, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2SharedBuffer)))
	return AsCoreWebView2SharedBuffer(resultCoreWebView2SharedBuffer)
}

func (m *TCoreWebView2SharedBuffer) Size() (resultInt64 int64) {
	WV().SysCallN(642, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCoreWebView2SharedBuffer) Buffer() PByte {
	r1 := WV().SysCallN(635, m.Instance())
	return PByte(r1)
}

func (m *TCoreWebView2SharedBuffer) OpenStream() IStream {
	var resultStream uintptr
	WV().SysCallN(641, m.Instance(), uintptr(unsafePointer(&resultStream)))
	return AsStream(resultStream)
}

func (m *TCoreWebView2SharedBuffer) FileMappingHandle() HANDLE {
	r1 := WV().SysCallN(639, m.Instance())
	return HANDLE(r1)
}

func (m *TCoreWebView2SharedBuffer) Close() bool {
	r1 := WV().SysCallN(637, m.Instance())
	return GoBool(r1)
}

func CoreWebView2SharedBufferClass() TClass {
	ret := WV().SysCallN(636)
	return TClass(ret)
}
