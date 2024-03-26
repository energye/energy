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

// ICoreWebView2AcceleratorKeyPressedEventArgs Parent: IObject
//
//	Event args for the AcceleratorKeyPressed event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
type ICoreWebView2AcceleratorKeyPressedEventArgs interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2AcceleratorKeyPressedEventArgs // property
	// KeyEventKind
	//  The key event type that caused the event to run.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_keyeventkind">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	KeyEventKind() TWVKeyEventKind // property
	// VirtualKey
	//  The Win32 virtual key code of the key that was pressed or released. It
	//  is one of the Win32 virtual key constants such as `VK_RETURN` or an
	// (uppercase) ASCII value such as `A`. Verify whether Ctrl or Alt
	//  are pressed by running `GetKeyState(VK_CONTROL)` or
	//  `GetKeyState(VK_MENU)`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_virtualkey">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	VirtualKey() uint32 // property
	// KeyEventLParam
	//  The `LPARAM` value that accompanied the window message. For more
	//  information, navigate to [WM_KEYDOWN](/windows/win32/inputdev/wm-keydown)
	//  and [WM_KEYUP](/windows/win32/inputdev/wm-keyup).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_keyeventlparam">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	KeyEventLParam() int32 // property
	// RepeatCount
	//  Specifies the repeat count for the current message.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</a>
	RepeatCount() uint32 // property
	// ScanCode
	//  Specifies the scan code.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</a>
	ScanCode() uint32 // property
	// IsExtendedKey
	//  Indicates that the key is an extended key.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</a>
	IsExtendedKey() bool // property
	// IsMenuKeyDown
	//  Indicates that a menu key is held down(context code).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</a>
	IsMenuKeyDown() bool // property
	// WasKeyDown
	//  Indicates that the key was held down.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</a>
	WasKeyDown() bool // property
	// IsKeyReleased
	//  Indicates that the key was released.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_physicalkeystatus">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/winrt/microsoft_web_webview2_core/corewebview2physicalkeystatus">See the CoreWebView2PhysicalKeyStatus Struct article.</a>
	IsKeyReleased() bool // property
	// Handled
	//  During `AcceleratorKeyPressedEvent` handler invocation the WebView is
	//  blocked waiting for the decision of if the accelerator is handled by the
	//  host(or not). If the `Handled` property is set to `TRUE` then this
	//  prevents the WebView from performing the default action for this
	//  accelerator key. Otherwise the WebView performs the default action for
	//  the accelerator key.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs#get_handled">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
	Handled() bool // property
	// SetHandled Set Handled
	SetHandled(AValue bool) // property
	// IsBrowserAcceleratorKeyEnabled
	//  This property allows developers to enable or disable the browser from handling a specific
	//  browser accelerator key such as Ctrl+P or F3, etc.
	//  Browser accelerator keys are the keys/key combinations that access features specific to
	//  a web browser, including but not limited to:
	//  <code>
	//  - Ctrl-F and F3 for Find on Page
	//  - Ctrl-P for Print
	//  - Ctrl-R and F5 for Reload
	//  - Ctrl-Plus and Ctrl-Minus for zooming
	//  - Ctrl-Shift-C and F12 for DevTools
	//  - Special keys for browser functions, such as Back, Forward, and Search
	//  </code>
	//  This property does not disable accelerator keys related to movement and text editing,
	//  such as:
	//  <code>
	//  - Home, End, Page Up, and Page Down
	//  - Ctrl-X, Ctrl-C, Ctrl-V
	//  - Ctrl-A for Select All
	//  - Ctrl-Z for Undo
	//  </code>
	//  The `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` API is a convenient setting
	//  for developers to disable all the browser accelerator keys together, and sets the default
	//  value for the `IsBrowserAcceleratorKeyEnabled` property.
	//  By default, `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` is `TRUE` and
	//  `IsBrowserAcceleratorKeyEnabled` is `TRUE`.
	//  When developers change `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` setting to `FALSE`,
	//  this will change default value for `IsBrowserAcceleratorKeyEnabled` to `FALSE`.
	//  If developers want specific keys to be handled by the browser after changing the
	//  `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` setting to `FALSE`, they need to enable
	//  these keys by setting `IsBrowserAcceleratorKeyEnabled` to `TRUE`.
	//  This API will give the event arg higher priority over the
	//  `ICoreWebView2Settings.AreBrowserAcceleratorKeysEnabled` setting when we handle the keys.
	//  For browser accelerator keys, when an accelerator key is pressed, the propagation and
	//  processing order is:
	//  <code>
	//  1. A ICoreWebView2Controller.AcceleratorKeyPressed event is raised
	//  2. WebView2 browser feature accelerator key handling
	//  3. Web Content Handling: If the key combination isn't reserved for browser actions,
	//  the key event propagates to the web content, where JavaScript event listeners can
	//  capture and respond to it.
	//  </code>
	//  `ICoreWebView2AcceleratorKeyPressedEventArgs` has a `Handled` property, that developers
	//  can use to mark a key as handled. When the key is marked as handled anywhere along
	//  the path, the event propagation stops, and web content will not receive the key.
	//  With `IsBrowserAcceleratorKeyEnabled` property, if developers mark
	//  `IsBrowserAcceleratorKeyEnabled` as `FALSE`, the browser will skip the WebView2
	//  browser feature accelerator key handling process, but the event propagation
	//  continues, and web content will receive the key combination.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs2#get_isbrowseracceleratorkeyenabled">See the ICoreWebView2AcceleratorKeyPressedEventArgs2 article.</a>
	IsBrowserAcceleratorKeyEnabled() bool // property
	// SetIsBrowserAcceleratorKeyEnabled Set IsBrowserAcceleratorKeyEnabled
	SetIsBrowserAcceleratorKeyEnabled(AValue bool) // property
}

// TCoreWebView2AcceleratorKeyPressedEventArgs Parent: TObject
//
//	Event args for the AcceleratorKeyPressed event.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2acceleratorkeypressedeventargs">See the ICoreWebView2AcceleratorKeyPressedEventArgs article.</a>
type TCoreWebView2AcceleratorKeyPressedEventArgs struct {
	TObject
}

func NewCoreWebView2AcceleratorKeyPressedEventArgs(aArgs ICoreWebView2AcceleratorKeyPressedEventArgs) ICoreWebView2AcceleratorKeyPressedEventArgs {
	r1 := WV().SysCallN(2, GetObjectUintptr(aArgs))
	return AsCoreWebView2AcceleratorKeyPressedEventArgs(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) Initialized() bool {
	r1 := WV().SysCallN(4, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) BaseIntf() ICoreWebView2AcceleratorKeyPressedEventArgs {
	var resultCoreWebView2AcceleratorKeyPressedEventArgs uintptr
	WV().SysCallN(0, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2AcceleratorKeyPressedEventArgs)))
	return AsCoreWebView2AcceleratorKeyPressedEventArgs(resultCoreWebView2AcceleratorKeyPressedEventArgs)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) KeyEventKind() TWVKeyEventKind {
	r1 := WV().SysCallN(9, m.Instance())
	return TWVKeyEventKind(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) VirtualKey() uint32 {
	r1 := WV().SysCallN(13, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) KeyEventLParam() int32 {
	r1 := WV().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) RepeatCount() uint32 {
	r1 := WV().SysCallN(11, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) ScanCode() uint32 {
	r1 := WV().SysCallN(12, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsExtendedKey() bool {
	r1 := WV().SysCallN(6, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsMenuKeyDown() bool {
	r1 := WV().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) WasKeyDown() bool {
	r1 := WV().SysCallN(14, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsKeyReleased() bool {
	r1 := WV().SysCallN(7, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) Handled() bool {
	r1 := WV().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) SetHandled(AValue bool) {
	WV().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) IsBrowserAcceleratorKeyEnabled() bool {
	r1 := WV().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2AcceleratorKeyPressedEventArgs) SetIsBrowserAcceleratorKeyEnabled(AValue bool) {
	WV().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func CoreWebView2AcceleratorKeyPressedEventArgsClass() TClass {
	ret := WV().SysCallN(1)
	return TClass(ret)
}
