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

// ICoreWebView2Settings Parent: IObject
//
//	Defines properties that enable, disable, or modify WebView features.
//	Changes to IsGeneralAutofillEnabled and IsPasswordAutosaveEnabled
//	apply immediately, while other setting changes made after NavigationStarting
//	event do not apply until the next top-level navigation.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings">See the ICoreWebView2Settings article.</a>
type ICoreWebView2Settings interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Settings // property
	// IsBuiltInErrorPageEnabled
	//  The `IsBuiltInErrorPageEnabled` property is used to disable built in
	//  error page for navigation failure and render process failure. When
	//  disabled, a blank page is displayed when the related error happens.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isbuiltinerrorpageenabled">See the ICoreWebView2Settings article.</a>
	IsBuiltInErrorPageEnabled() bool // property
	// SetIsBuiltInErrorPageEnabled Set IsBuiltInErrorPageEnabled
	SetIsBuiltInErrorPageEnabled(AValue bool) // property
	// AreDefaultContextMenusEnabled
	//  The `AreDefaultContextMenusEnabled` property is used to prevent default
	//  context menus from being shown to user in WebView.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultcontextmenusenabled">See the ICoreWebView2Settings article.</a>
	AreDefaultContextMenusEnabled() bool // property
	// SetAreDefaultContextMenusEnabled Set AreDefaultContextMenusEnabled
	SetAreDefaultContextMenusEnabled(AValue bool) // property
	// AreDefaultScriptDialogsEnabled
	//  `AreDefaultScriptDialogsEnabled` is used when loading a new HTML
	//  document. If set to `FALSE`, WebView2 does not render the default JavaScript
	//  dialog box(Specifically those displayed by the JavaScript alert,
	//  confirm, prompt functions and `beforeunload` event). Instead, if an
	//  event handler is set using `add_ScriptDialogOpening`, WebView sends an
	//  event that contains all of the information for the dialog and allow the
	//  host app to show a custom UI.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultscriptdialogsenabled">See the ICoreWebView2Settings article.</a>
	AreDefaultScriptDialogsEnabled() bool // property
	// SetAreDefaultScriptDialogsEnabled Set AreDefaultScriptDialogsEnabled
	SetAreDefaultScriptDialogsEnabled(AValue bool) // property
	// AreDevToolsEnabled
	//  `AreDevToolsEnabled` controls whether the user is able to use the context
	//  menu or keyboard shortcuts to open the DevTools window.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredevtoolsenabled">See the ICoreWebView2Settings article.</a>
	AreDevToolsEnabled() bool // property
	// SetAreDevToolsEnabled Set AreDevToolsEnabled
	SetAreDevToolsEnabled(AValue bool) // property
	// IsScriptEnabled
	//  Controls if running JavaScript is enabled in all future navigations in
	//  the WebView. This only affects scripts in the document. Scripts
	//  injected with `ExecuteScript` runs even if script is disabled.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isscriptenabled">See the ICoreWebView2Settings article.</a>
	IsScriptEnabled() bool // property
	// SetIsScriptEnabled Set IsScriptEnabled
	SetIsScriptEnabled(AValue bool) // property
	// IsStatusBarEnabled
	//  `IsStatusBarEnabled` controls whether the status bar is displayed. The
	//  status bar is usually displayed in the lower left of the WebView and
	//  shows things such as the URI of a link when the user hovers over it and
	//  other information. The default value is `TRUE`. The status bar UI can be
	//  altered by web content and should not be considered secure.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isstatusbarenabled">See the ICoreWebView2Settings article.</a>
	IsStatusBarEnabled() bool // property
	// SetIsStatusBarEnabled Set IsStatusBarEnabled
	SetIsStatusBarEnabled(AValue bool) // property
	// IsWebMessageEnabled
	//  The `IsWebMessageEnabled` property is used when loading a new HTML
	//  document. If set to `TRUE`, communication from the host to the top-level
	//  HTML document of the WebView is allowed using `PostWebMessageAsJson`,
	//  `PostWebMessageAsString`, and message event of `window.chrome.webview`.
	//  For more information, navigate to PostWebMessageAsJson. Communication
	//  from the top-level HTML document of the WebView to the host is allowed
	//  using the postMessage function of `window.chrome.webview` and
	//  `add_WebMessageReceived` method. For more information, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  If set to false, then communication is disallowed. `PostWebMessageAsJson`
	//  and `PostWebMessageAsString` fails with `E_ACCESSDENIED` and
	//  `window.chrome.webview.postMessage` fails by throwing an instance of an
	//  `Error` object. The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iswebmessageenabled">See the ICoreWebView2Settings article.</a>
	IsWebMessageEnabled() bool // property
	// SetIsWebMessageEnabled Set IsWebMessageEnabled
	SetIsWebMessageEnabled(AValue bool) // property
	// IsZoomControlEnabled
	//  The `IsZoomControlEnabled` property is used to prevent the user from
	//  impacting the zoom of the WebView. When disabled, the user is not able
	//  to zoom using Ctrl++, Ctrl+-, or Ctrl+mouse wheel, but the zoom
	//  is set using `ZoomFactor` API. The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iszoomcontrolenabled">See the ICoreWebView2Settings article.</a>
	IsZoomControlEnabled() bool // property
	// SetIsZoomControlEnabled Set IsZoomControlEnabled
	SetIsZoomControlEnabled(AValue bool) // property
	// AreHostObjectsAllowed
	//  The `AreHostObjectsAllowed` property is used to control whether host
	//  objects are accessible from the page in WebView.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_arehostobjectsallowed">See the ICoreWebView2Settings article.</a>
	AreHostObjectsAllowed() bool // property
	// SetAreHostObjectsAllowed Set AreHostObjectsAllowed
	SetAreHostObjectsAllowed(AValue bool) // property
	// UserAgent
	//  Returns the User Agent. The default value is the default User Agent of the
	//  Microsoft Edge browser.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings2#get_useragent">See the ICoreWebView2Settings2 article.</a>
	UserAgent() string // property
	// SetUserAgent Set UserAgent
	SetUserAgent(AValue string) // property
	// AreBrowserAcceleratorKeysEnabled
	//  When this setting is set to FALSE, it disables all accelerator keys that
	//  access features specific to a web browser, including but not limited to:
	//  - Ctrl-F and F3 for Find on Page
	//  - Ctrl-P for Print
	//  - Ctrl-R and F5 for Reload
	//  - Ctrl-Plus and Ctrl-Minus for zooming
	//  - Ctrl-Shift-C and F12 for DevTools
	//  - Special keys for browser functions, such as Back, Forward, and Search
	//  It does not disable accelerator keys related to movement and text editing,
	//  such as:
	//  - Home, End, Page Up, and Page Down
	//  - Ctrl-X, Ctrl-C, Ctrl-V
	//  - Ctrl-A for Select All
	//  - Ctrl-Z for Undo
	//  Those accelerator keys will always be enabled unless they are handled in
	//  the `AcceleratorKeyPressed` event.
	//  This setting has no effect on the `AcceleratorKeyPressed` event. The event
	//  will be fired for all accelerator keys, whether they are enabled or not.
	//  The default value for `AreBrowserAcceleratorKeysEnabled` is TRUE.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings3#get_arebrowseracceleratorkeysenabled">See the ICoreWebView2Settings3 article.</a>
	AreBrowserAcceleratorKeysEnabled() bool // property
	// SetAreBrowserAcceleratorKeysEnabled Set AreBrowserAcceleratorKeysEnabled
	SetAreBrowserAcceleratorKeysEnabled(AValue bool) // property
	// IsPasswordAutosaveEnabled
	//  IsPasswordAutosaveEnabled controls whether autosave for password
	//  information is enabled. The IsPasswordAutosaveEnabled property behaves
	//  independently of the IsGeneralAutofillEnabled property. When IsPasswordAutosaveEnabled is
	//  false, no new password data is saved and no Save/Update Password prompts are displayed.
	//  However, if there was password data already saved before disabling this setting,
	//  then that password information is auto-populated, suggestions are shown and clicking on
	//  one will populate the fields.
	//  When IsPasswordAutosaveEnabled is true, password information is auto-populated,
	//  suggestions are shown and clicking on one will populate the fields, new data
	//  is saved, and a Save/Update Password prompt is displayed.
	//  It will take effect immediately after setting.
	//  The default value is `FALSE`.
	//  This property has the same value as
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled` and
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled` will always have the same
	//  value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings4#get_ispasswordautosaveenabled">See the ICoreWebView2Settings4 article.</a>
	IsPasswordAutosaveEnabled() bool // property
	// SetIsPasswordAutosaveEnabled Set IsPasswordAutosaveEnabled
	SetIsPasswordAutosaveEnabled(AValue bool) // property
	// IsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields.
	//  It will take effect immediately after setting.
	//  The default value is `TRUE`.
	//  This property has the same value as
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled` and
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled` will always have the same
	//  value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings4#get_isgeneralautofillenabled">See the ICoreWebView2Settings4 article.</a>
	IsGeneralAutofillEnabled() bool // property
	// SetIsGeneralAutofillEnabled Set IsGeneralAutofillEnabled
	SetIsGeneralAutofillEnabled(AValue bool) // property
	// IsPinchZoomEnabled
	//  Pinch-zoom, referred to as "Page Scale" zoom, is performed as a post-rendering step,
	//  it changes the page scale factor property and scales the surface the web page is
	//  rendered onto when user performs a pinch zooming action. It does not change the layout
	//  but rather changes the viewport and clips the web content, the content outside of the
	//  viewport isn't visible onscreen and users can't reach this content using mouse.
	//  The `IsPinchZoomEnabled` property enables or disables the ability of
	//  the end user to use a pinching motion on touch input enabled devices
	//  to scale the web content in the WebView2. It defaults to `TRUE`.
	//  When set to `FALSE`, the end user cannot pinch zoom after the next navigation.
	//  Disabling/Enabling `IsPinchZoomEnabled` only affects the end user's ability to use
	//  pinch motions and does not change the page scale factor.
	//  This API only affects the Page Scale zoom and has no effect on the
	//  existing browser zoom properties(`IsZoomControlEnabled` and `ZoomFactor`)
	//  or other end user mechanisms for zooming.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings5#get_ispinchzoomenabled">See the ICoreWebView2Settings5 article.</a>
	IsPinchZoomEnabled() bool // property
	// SetIsPinchZoomEnabled Set IsPinchZoomEnabled
	SetIsPinchZoomEnabled(AValue bool) // property
	// IsSwipeNavigationEnabled
	//  The `IsSwipeNavigationEnabled` property enables or disables the ability of the
	//  end user to use swiping gesture on touch input enabled devices to
	//  navigate in WebView2. It defaults to `TRUE`.
	//  When this property is `TRUE`, then all configured navigation gestures are enabled:
	//  1. Swiping left and right to navigate forward and backward is always configured.
	//  2. Swiping down to refresh is off by default and not exposed via our API currently,
	//  it requires the "--pull-to-refresh" option to be included in the additional browser
	//  arguments to be configured.(See put_AdditionalBrowserArguments.)
	//  When set to `FALSE`, the end user cannot swipe to navigate or pull to refresh.
	//  This API only affects the overscrolling navigation functionality and has no
	//  effect on the scrolling interaction used to explore the web content shown
	//  in WebView2.
	//  Disabling/Enabling IsSwipeNavigationEnabled takes effect after the
	//  next navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings6#get_isswipenavigationenabled">See the ICoreWebView2Settings6 article.</a>
	IsSwipeNavigationEnabled() bool // property
	// SetIsSwipeNavigationEnabled Set IsSwipeNavigationEnabled
	SetIsSwipeNavigationEnabled(AValue bool) // property
	// HiddenPdfToolbarItems
	//  `HiddenPdfToolbarItems` is used to customize the PDF toolbar items. By default, it is COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE and so it displays all of the items.
	//  Changes to this property apply to all CoreWebView2s in the same environment and using the same profile.
	//  Changes to this setting apply only after the next navigation.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings7#get_hiddenpdftoolbaritems">See the ICoreWebView2Settings7 article.</a>
	HiddenPdfToolbarItems() TWVPDFToolbarItems // property
	// SetHiddenPdfToolbarItems Set HiddenPdfToolbarItems
	SetHiddenPdfToolbarItems(AValue TWVPDFToolbarItems) // property
	// IsReputationCheckingRequired
	//  SmartScreen helps webviews identify reported phishing and malware websites
	//  and also helps users make informed decisions about downloads.
	//  `IsReputationCheckingRequired` is used to control whether SmartScreen
	//  enabled or not. SmartScreen is enabled or disabled for all CoreWebView2s
	//  using the same user data folder. If
	//  CoreWebView2Setting.IsReputationCheckingRequired is true for any
	//  CoreWebView2 using the same user data folder, then SmartScreen is enabled.
	//  If CoreWebView2Setting.IsReputationCheckingRequired is false for all
	//  CoreWebView2 using the same user data folder, then SmartScreen is
	//  disabled. When it is changed, the change will be applied to all WebViews
	//  using the same user data folder on the next navigation or download. The
	//  default value for `IsReputationCheckingRequired` is true. If the newly
	//  created CoreWebview2 does not set SmartScreen to false, when
	//  navigating(Such as Navigate(), LoadDataUrl(), ExecuteScript(), etc.), the
	//  default value will be applied to all CoreWebview2 using the same user data
	//  folder.
	//  SmartScreen of WebView2 apps can be controlled by Windows system setting
	//  "SmartScreen for Microsoft Edge", specially, for WebView2 in Windows
	//  Store apps, SmartScreen is controlled by another Windows system setting
	//  "SmartScreen for Microsoft Store apps". When the Windows setting is enabled, the
	//  SmartScreen operates under the control of the `IsReputationCheckingRequired`.
	//  When the Windows setting is disabled, the SmartScreen will be disabled
	//  regardless of the `IsReputationCheckingRequired` value set in WebView2 apps.
	//  In other words, under this circumstance the value of
	//  `IsReputationCheckingRequired` will be saved but overridden by system setting.
	//  Upon re-enabling the Windows setting, the CoreWebview2 will reference the
	//  `IsReputationCheckingRequired` to determine the SmartScreen status.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings8#get_isreputationcheckingrequired">See the ICoreWebView2Settings8 article.</a>
	IsReputationCheckingRequired() bool // property
	// SetIsReputationCheckingRequired Set IsReputationCheckingRequired
	SetIsReputationCheckingRequired(AValue bool) // property
}

// TCoreWebView2Settings Parent: TObject
//
//	Defines properties that enable, disable, or modify WebView features.
//	Changes to IsGeneralAutofillEnabled and IsPasswordAutosaveEnabled
//	apply immediately, while other setting changes made after NavigationStarting
//	event do not apply until the next top-level navigation.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings">See the ICoreWebView2Settings article.</a>
type TCoreWebView2Settings struct {
	TObject
}

func NewCoreWebView2Settings(aBaseIntf ICoreWebView2Settings) ICoreWebView2Settings {
	r1 := WV().SysCallN(620, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Settings(r1)
}

func (m *TCoreWebView2Settings) Initialized() bool {
	r1 := WV().SysCallN(622, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) BaseIntf() ICoreWebView2Settings {
	var resultCoreWebView2Settings uintptr
	WV().SysCallN(618, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Settings)))
	return AsCoreWebView2Settings(resultCoreWebView2Settings)
}

func (m *TCoreWebView2Settings) IsBuiltInErrorPageEnabled() bool {
	r1 := WV().SysCallN(623, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsBuiltInErrorPageEnabled(AValue bool) {
	WV().SysCallN(623, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) AreDefaultContextMenusEnabled() bool {
	r1 := WV().SysCallN(614, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetAreDefaultContextMenusEnabled(AValue bool) {
	WV().SysCallN(614, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) AreDefaultScriptDialogsEnabled() bool {
	r1 := WV().SysCallN(615, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetAreDefaultScriptDialogsEnabled(AValue bool) {
	WV().SysCallN(615, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) AreDevToolsEnabled() bool {
	r1 := WV().SysCallN(616, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetAreDevToolsEnabled(AValue bool) {
	WV().SysCallN(616, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsScriptEnabled() bool {
	r1 := WV().SysCallN(628, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsScriptEnabled(AValue bool) {
	WV().SysCallN(628, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsStatusBarEnabled() bool {
	r1 := WV().SysCallN(629, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsStatusBarEnabled(AValue bool) {
	WV().SysCallN(629, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsWebMessageEnabled() bool {
	r1 := WV().SysCallN(631, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsWebMessageEnabled(AValue bool) {
	WV().SysCallN(631, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsZoomControlEnabled() bool {
	r1 := WV().SysCallN(632, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsZoomControlEnabled(AValue bool) {
	WV().SysCallN(632, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) AreHostObjectsAllowed() bool {
	r1 := WV().SysCallN(617, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetAreHostObjectsAllowed(AValue bool) {
	WV().SysCallN(617, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) UserAgent() string {
	r1 := WV().SysCallN(633, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2Settings) SetUserAgent(AValue string) {
	WV().SysCallN(633, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2Settings) AreBrowserAcceleratorKeysEnabled() bool {
	r1 := WV().SysCallN(613, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetAreBrowserAcceleratorKeysEnabled(AValue bool) {
	WV().SysCallN(613, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsPasswordAutosaveEnabled() bool {
	r1 := WV().SysCallN(625, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsPasswordAutosaveEnabled(AValue bool) {
	WV().SysCallN(625, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsGeneralAutofillEnabled() bool {
	r1 := WV().SysCallN(624, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsGeneralAutofillEnabled(AValue bool) {
	WV().SysCallN(624, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsPinchZoomEnabled() bool {
	r1 := WV().SysCallN(626, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsPinchZoomEnabled(AValue bool) {
	WV().SysCallN(626, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) IsSwipeNavigationEnabled() bool {
	r1 := WV().SysCallN(630, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsSwipeNavigationEnabled(AValue bool) {
	WV().SysCallN(630, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Settings) HiddenPdfToolbarItems() TWVPDFToolbarItems {
	r1 := WV().SysCallN(621, 0, m.Instance(), 0)
	return TWVPDFToolbarItems(r1)
}

func (m *TCoreWebView2Settings) SetHiddenPdfToolbarItems(AValue TWVPDFToolbarItems) {
	WV().SysCallN(621, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Settings) IsReputationCheckingRequired() bool {
	r1 := WV().SysCallN(627, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Settings) SetIsReputationCheckingRequired(AValue bool) {
	WV().SysCallN(627, 1, m.Instance(), PascalBool(AValue))
}

func CoreWebView2SettingsClass() TClass {
	ret := WV().SysCallN(619)
	return TClass(ret)
}
