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

// IWVBrowserBase Is Abstract Class Parent: IComponent
//
//	Parent class of TWVBrowser and TWVFMXBrowser that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type IWVBrowserBase interface {
	IComponent
	// Initialized
	// Custom properties
	Initialized() bool // property
	// CoreWebView2PrintSettings
	//  Settings used for printing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2printsettings">See the ICoreWebView2PrintSettings article.</a>
	CoreWebView2PrintSettings() ICoreWebView2PrintSettings // property
	// CoreWebView2Settings
	//  CoreWebView2Settings contains various modifiable settings for the running WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings">See the ICoreWebView2Settings article.</a>
	CoreWebView2Settings() ICoreWebView2Settings // property
	// CoreWebView2Environment
	//  Represents the WebView2 Environment.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment">See the ICoreWebView2Environment article.</a>
	CoreWebView2Environment() ICoreWebView2Environment // property
	// CoreWebView2Controller
	//  The owner of the `CoreWebView2` object that provides support for resizing,
	//  showing and hiding, focusing, and other functionality related to
	//  windowing and composition. The `CoreWebView2Controller` owns the
	//  `CoreWebView2`, and if all references to the `CoreWebView2Controller` go
	//  away, the WebView is closed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller">See the ICoreWebView2Controller article.</a>
	CoreWebView2Controller() ICoreWebView2Controller // property
	// CoreWebView2CompositionController
	//  ICoreWebView2CompositionController wrapper used by this browser.
	//  This interface is an extension of the ICoreWebView2Controller interface to
	//  support visual hosting. An object implementing the
	//  ICoreWebView2CompositionController interface will also implement
	//  ICoreWebView2Controller. Callers are expected to use
	//  ICoreWebView2Controller for resizing, visibility, focus, and so on, and
	//  then use ICoreWebView2CompositionController to connect to a composition
	//  tree and provide input meant for the WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller">See the ICoreWebView2CompositionController article.</a>
	CoreWebView2CompositionController() ICoreWebView2CompositionController // property
	// CoreWebView2
	//  ICoreWebView2 wrapper used by this browser.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2">See the ICoreWebView2 article.</a>
	CoreWebView2() ICoreWebView2 // property
	// Profile
	//  The associated `ICoreWebView2Profile` object. If this CoreWebView2 was created with a
	//  CoreWebView2ControllerOptions, the CoreWebView2Profile will match those specified options.
	//  Otherwise if this CoreWebView2 was created without a CoreWebView2ControllerOptions, then
	//  this will be the default CoreWebView2Profile for the corresponding CoreWebView2Environment.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_13#get_profile">See the ICoreWebView2_13 article.</a>
	Profile() ICoreWebView2Profile // property
	// DefaultURL
	//  First URL loaded by the browser after its creation.
	DefaultURL() string // property
	// SetDefaultURL Set DefaultURL
	SetDefaultURL(AValue string) // property
	// IsNavigating
	//  Returns true after OnNavigationStarting and before OnNavigationCompleted.
	IsNavigating() bool // property
	// ZoomPct
	//  Returns the current zoom value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	ZoomPct() (resultDouble float64) // property
	// SetZoomPct Set ZoomPct
	SetZoomPct(AValue float64) // property
	// ZoomStep
	//  Returns the current zoom value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	ZoomStep() byte // property
	// SetZoomStep Set ZoomStep
	SetZoomStep(AValue byte) // property
	// Widget0CompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	Widget0CompHWND() THandle // property
	// Widget1CompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	Widget1CompHWND() THandle // property
	// RenderCompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	RenderCompHWND() THandle // property
	// D3DWindowCompHWND
	//  Handle of one to the child controls created automatically by the browser to show the web contents.
	D3DWindowCompHWND() THandle // property
	// ScreenScale
	//  Returns the GlobalWebView2Loader.DeviceScaleFactor value.
	ScreenScale() float32 // property
	// Offline
	//  Uses the Network.emulateNetworkConditions DevTool method to set the browser in offline mode.
	//  The TWVBrowserBase.OnOfflineCompleted event is triggered asynchronously after setting this property.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions">See the Network Domain article.</a>
	Offline() bool // property
	// SetOffline Set Offline
	SetOffline(AValue bool) // property
	// IgnoreCertificateErrors
	//  Uses the Security.setIgnoreCertificateErrors DevTool method to enable/disable whether all certificate errors should be ignored.
	//  The TWVBrowserBase.OnIgnoreCertificateErrorsCompleted event is triggered asynchronously after setting this property.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors">See the Security Domain article.</a>
	IgnoreCertificateErrors() bool // property
	// SetIgnoreCertificateErrors Set IgnoreCertificateErrors
	SetIgnoreCertificateErrors(AValue bool) // property
	// BrowserExecPath
	// Properties used in the ICoreWebView2Environment creation
	BrowserExecPath() string // property
	// SetBrowserExecPath Set BrowserExecPath
	SetBrowserExecPath(AValue string) // property
	// UserDataFolder
	//  Returns the user data folder that all CoreWebView2's created from this
	//  environment are using.
	//  This could be either the value passed in by the developer when creating
	//  the environment object or the calculated one for default handling. It
	//  will always be an absolute path.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment7#get_userdatafolder">See the ICoreWebView2Environment7 article.</a>
	UserDataFolder() string // property
	// SetUserDataFolder Set UserDataFolder
	SetUserDataFolder(AValue string) // property
	// AdditionalBrowserArguments
	//  Additional command line switches.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AdditionalBrowserArguments.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	AdditionalBrowserArguments() string // property
	// SetAdditionalBrowserArguments Set AdditionalBrowserArguments
	SetAdditionalBrowserArguments(AValue string) // property
	// Language
	//  The default display language for WebView. It applies to browser UI such as
	//  context menu and dialogs. It also applies to the `accept-languages` HTTP
	//  header that WebView sends to websites.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_Language.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	Language() string // property
	// SetLanguage Set Language
	SetLanguage(AValue string) // property
	// TargetCompatibleBrowserVersion
	//  Specifies the version of the WebView2 Runtime binaries required to be
	//  compatible with your app.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_TargetCompatibleBrowserVersion.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	TargetCompatibleBrowserVersion() string // property
	// SetTargetCompatibleBrowserVersion Set TargetCompatibleBrowserVersion
	SetTargetCompatibleBrowserVersion(AValue string) // property
	// AllowSingleSignOnUsingOSPrimaryAccount
	//  Used to enable single sign on with Azure Active Directory(AAD) and personal Microsoft
	//  Account(MSA) resources inside WebView.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions.get_AllowSingleSignOnUsingOSPrimaryAccount.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions">See the ICoreWebView2EnvironmentOptions article.</a>
	AllowSingleSignOnUsingOSPrimaryAccount() bool // property
	// SetAllowSingleSignOnUsingOSPrimaryAccount Set AllowSingleSignOnUsingOSPrimaryAccount
	SetAllowSingleSignOnUsingOSPrimaryAccount(AValue bool) // property
	// ExclusiveUserDataFolderAccess
	//  Whether other processes can create WebView2 from WebView2Environment created with the
	//  same user data folder and therefore sharing the same WebView browser process instance.
	//  Default is FALSE.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions2.Get_ExclusiveUserDataFolderAccess.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions2">See the ICoreWebView2EnvironmentOptions2 article.</a>
	ExclusiveUserDataFolderAccess() bool // property
	// SetExclusiveUserDataFolderAccess Set ExclusiveUserDataFolderAccess
	SetExclusiveUserDataFolderAccess(AValue bool) // property
	// CustomCrashReportingEnabled
	//  When `CustomCrashReportingEnabled` is set to `TRUE`, Windows won't send crash data to Microsoft endpoint.
	//  `CustomCrashReportingEnabled` is default to be `FALSE`, in this case, WebView will respect OS consent.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions3.Get_IsCustomCrashReportingEnabled.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions3">See the ICoreWebView2EnvironmentOptions3 article.</a>
	CustomCrashReportingEnabled() bool // property
	// SetCustomCrashReportingEnabled Set CustomCrashReportingEnabled
	SetCustomCrashReportingEnabled(AValue bool) // property
	// EnableTrackingPrevention
	//  The `EnableTrackingPrevention` property is used to enable/disable tracking prevention
	//  feature in WebView2. This property enable/disable tracking prevention for all the
	//  WebView2's created in the same environment.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions5.Get_EnableTrackingPrevention.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions5">See the ICoreWebView2EnvironmentOptions5 article.</a>
	EnableTrackingPrevention() bool // property
	// SetEnableTrackingPrevention Set EnableTrackingPrevention
	SetEnableTrackingPrevention(AValue bool) // property
	// AreBrowserExtensionsEnabled
	//  When `AreBrowserExtensionsEnabled` is set to `TRUE`, new extensions can be added to user
	//  profile and used. `AreBrowserExtensionsEnabled` is default to be `FALSE`, in this case,
	//  new extensions can't be installed, and already installed extension won't be
	//  available to use in user profile.
	//  If connecting to an already running environment with a different value for `AreBrowserExtensionsEnabled`
	//  property, it will fail with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Property used to create the environment. Used as ICoreWebView2EnvironmentOptions6.Get_AreBrowserExtensionsEnabled.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions6">See the ICoreWebView2EnvironmentOptions6 article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension">See the ICoreWebView2BrowserExtension article for Extensions API details.</a>
	AreBrowserExtensionsEnabled() bool // property
	// SetAreBrowserExtensionsEnabled Set AreBrowserExtensionsEnabled
	SetAreBrowserExtensionsEnabled(AValue bool) // property
	// BrowserVersionInfo
	//  The browser version info of the current `ICoreWebView2Environment`,
	//  including channel name if it is not the WebView2 Runtime. It matches the
	//  format of the `GetAvailableCoreWebView2BrowserVersionString` API.
	//  Channel names are `beta`, `dev`, and `canary`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#get_browserversionstring">See the ICoreWebView2Environment article.</a>
	BrowserVersionInfo() string // property
	// BrowserProcessID
	//  The process ID of the browser process that hosts the WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_browserprocessid">See the ICoreWebView2 article.</a>
	BrowserProcessID() uint32 // property
	// CanGoBack
	//  `TRUE` if the WebView is able to navigate to a previous page in the
	//  navigation history. If `CanGoBack` changes value, the `HistoryChanged`
	//  event runs.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2##get_cangoback">See the ICoreWebView2 article.</a>
	CanGoBack() bool // property
	// CanGoForward
	//  `TRUE` if the WebView is able to navigate to a next page in the
	//  navigation history. If `CanGoForward` changes value, the
	//  `HistoryChanged` event runs.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_cangoforward">See the ICoreWebView2 article.</a>
	CanGoForward() bool // property
	// ContainsFullScreenElement
	//  Indicates if the WebView contains a fullscreen HTML element.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_containsfullscreenelement">See the ICoreWebView2 article.</a>
	ContainsFullScreenElement() bool // property
	// DocumentTitle
	//  The title for the current top-level document. If the document has no
	//  explicit title or is otherwise empty, a default that may or may not match
	//  the URI of the document is used.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_documenttitle">See the ICoreWebView2 article.</a>
	DocumentTitle() string // property
	// Source
	//  The URI of the current top level document. This value potentially
	//  changes as a part of the `SourceChanged` event that runs for some cases
	//  such as navigating to a different site or fragment navigations. It
	//  remains the same for other types of navigations such as page refreshes
	//  or `history.pushState` with the same URL as the current page.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_source">See the ICoreWebView2 article.</a>
	Source() string // property
	// CookieManager
	//  Gets the cookie manager object associated with this ICoreWebView2.
	//  See ICoreWebView2CookieManager.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#get_cookiemanager">See the ICoreWebView2_2 article.</a>
	CookieManager() ICoreWebView2CookieManager // property
	// IsSuspended
	//  Whether WebView is suspended.
	//  `TRUE` when WebView is suspended, from the time when TrySuspend has completed
	//  successfully until WebView is resumed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#get_issuspended">See the ICoreWebView2_3 article.</a>
	IsSuspended() bool // property
	// IsDocumentPlayingAudio
	//  Indicates whether any audio output from this CoreWebView2 is playing.
	//  This property will be true if audio is playing even if IsMuted is true.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_isdocumentplayingaudio">See the ICoreWebView2_8 article.</a>
	IsDocumentPlayingAudio() bool // property
	// IsMuted
	//  Indicates whether all audio output from this CoreWebView2 is muted or not.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_ismuted">See the ICoreWebView2_8 article.</a>
	IsMuted() bool // property
	// SetIsMuted Set IsMuted
	SetIsMuted(AValue bool) // property
	// DefaultDownloadDialogCornerAlignment
	//  Get the default download dialog corner alignment.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_defaultdownloaddialogcorneralignment">See the ICoreWebView2_9 article.</a>
	DefaultDownloadDialogCornerAlignment() TWVDefaultDownloadDialogCornerAlignment // property
	// SetDefaultDownloadDialogCornerAlignment Set DefaultDownloadDialogCornerAlignment
	SetDefaultDownloadDialogCornerAlignment(AValue TWVDefaultDownloadDialogCornerAlignment) // property
	// DefaultDownloadDialogMargin
	//  Get the default download dialog margin.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_defaultdownloaddialogmargin">See the ICoreWebView2_9 article.</a>
	DefaultDownloadDialogMargin() (resultPoint TPoint) // property
	// SetDefaultDownloadDialogMargin Set DefaultDownloadDialogMargin
	SetDefaultDownloadDialogMargin(AValue *TPoint) // property
	// IsDefaultDownloadDialogOpen
	//  `TRUE` if the default download dialog is currently open. The value of this
	//  property changes only when the default download dialog is explicitly
	//  opened or closed. Hiding the WebView implicitly hides the dialog, but does
	//  not change the value of this property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_isdefaultdownloaddialogopen">See the ICoreWebView2_9 article.</a>
	IsDefaultDownloadDialogOpen() bool // property
	// StatusBarText
	//  The status message text.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_12#get_statusbartext">See the ICoreWebView2_12 article.</a>
	StatusBarText() string // property
	// FaviconURI
	//  Get the current Uri of the favicon as a string.
	//  If the value is null, then the return value is `E_POINTER`, otherwise it is `S_OK`.
	//  If a page has no favicon then the value is an empty string.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_15#get_faviconuri">See the ICoreWebView2_15 article.</a>
	FaviconURI() string // property
	// MemoryUsageTargetLevel
	//  `MemoryUsageTargetLevel` indicates desired memory consumption level of
	//  WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_19#get_memoryusagetargetlevel">See the ICoreWebView2_19 article.</a>
	MemoryUsageTargetLevel() TWVMemoryUsageTargetLevel // property
	// SetMemoryUsageTargetLevel Set MemoryUsageTargetLevel
	SetMemoryUsageTargetLevel(AValue TWVMemoryUsageTargetLevel) // property
	// Bounds
	//  The WebView bounds. Bounds are relative to the parent `HWND`. The app
	//  has two ways to position a WebView.
	//  * Create a child `HWND` that is the WebView parent `HWND`. Position
	//  the window where the WebView should be. Use `(0, 0)` for the
	//  top-left corner(the offset) of the `Bounds` of the WebView.
	//  * Use the top-most window of the app as the WebView parent HWND. For
	//  example, to position WebView correctly in the app, set the top-left
	//  corner of the Bound of the WebView.
	//  The values of `Bounds` are limited by the coordinate space of the host.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_bounds">See the ICoreWebView2Controller article.</a>
	Bounds() (resultRect TRect) // property
	// SetBounds Set Bounds
	SetBounds(AValue *TRect) // property
	// IsVisible
	//  The `IsVisible` property determines whether to show or hide the WebView2.
	//  If `IsVisible` is set to `FALSE`, the WebView2 is transparent and is
	//  not rendered. However, this does not affect the window containing the
	//  WebView2(the `HWND` parameter that was passed to
	//  `CreateCoreWebView2Controller`). If you want that window to disappear
	//  too, run `ShowWindow` on it directly in addition to modifying the
	//  `IsVisible` property. WebView2 as a child window does not get window
	//  messages when the top window is minimized or restored. For performance
	//  reasons, developers should set the `IsVisible` property of the WebView to
	//  `FALSE` when the app window is minimized and back to `TRUE` when the app
	//  window is restored. The app window does this by handling
	//  `SIZE_MINIMIZED and SIZE_RESTORED` command upon receiving `WM_SIZE`
	//  message.
	//  There are CPU and memory benefits when the page is hidden. For instance,
	//  Chromium has code that throttles activities on the page like animations
	//  and some tasks are run less frequently. Similarly, WebView2 will
	//  purge some caches to reduce memory usage.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_isvisible">See the ICoreWebView2Controller article.</a>
	IsVisible() bool // property
	// SetIsVisible Set IsVisible
	SetIsVisible(AValue bool) // property
	// ParentWindow
	//  The parent window provided by the app that this WebView is using to
	//  render content. This API initially returns the window passed into
	//  `CreateCoreWebView2Controller`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_parentwindow">See the ICoreWebView2Controller article.</a>
	//  <a href="https://github.com/salvadordf/WebView4Delphi/issues/13">See the WebView4Delphi issue #13 to know how to reparent a browser.</a>
	ParentWindow() THandle // property
	// SetParentWindow Set ParentWindow
	SetParentWindow(AValue THandle) // property
	// ZoomFactor
	//  The zoom factor for the WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	ZoomFactor() (resultDouble float64) // property
	// SetZoomFactor Set ZoomFactor
	SetZoomFactor(AValue float64) // property
	// DefaultBackgroundColor
	//  The `DefaultBackgroundColor` property is the color WebView renders
	//  underneath all web content. This means WebView renders this color when
	//  there is no web content loaded such as before the initial navigation or
	//  between navigations. This also means web pages with undefined css
	//  background properties or background properties containing transparent
	//  pixels will render their contents over this color. Web pages with defined
	//  and opaque background properties that span the page will obscure the
	//  `DefaultBackgroundColor` and display normally. The default value for this
	//  property is white to resemble the native browser experience.
	//  The Color is specified by the COREWEBVIEW2_COLOR that represents an RGBA
	//  value. The `A` represents an Alpha value, meaning
	//  `DefaultBackgroundColor` can be transparent. In the case of a transparent
	//  `DefaultBackgroundColor` WebView will render hosting app content as the
	//  background. This Alpha value is not supported on Windows 7. Any `A` value
	//  other than 255 will result in E_INVALIDARG on Windows 7.
	//  It is supported on all other WebView compatible platforms.
	//  Semi-transparent colors are not currently supported by this API and
	//  setting `DefaultBackgroundColor` to a semi-transparent color will fail
	//  with E_INVALIDARG. The only supported alpha values are 0 and 255, all
	//  other values will result in E_INVALIDARG.
	//  `DefaultBackgroundColor` can only be an opaque color or transparent.
	//  This value may also be set by using the
	//  `WEBVIEW2_DEFAULT_BACKGROUND_COLOR` environment variable. There is a
	//  known issue with background color where setting the color by API can
	//  still leave the app with a white flicker before the
	//  `DefaultBackgroundColor` takes effect. Setting the color via environment
	//  variable solves this issue. The value must be a hex value that can
	//  optionally prepend a 0x. The value must account for the alpha value
	//  which is represented by the first 2 digits. So any hex value fewer than 8
	//  digits will assume a prepended 00 to the hex value and result in a
	//  transparent color.
	//  `get_DefaultBackgroundColor` will return the result of this environment
	//  variable if used. This environment variable can only set the
	//  `DefaultBackgroundColor` once. Subsequent updates to background color
	//  must be done through API call.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller2#get_defaultbackgroundcolor">See the ICoreWebView2Controller2 article.</a>
	DefaultBackgroundColor() TColor // property
	// SetDefaultBackgroundColor Set DefaultBackgroundColor
	SetDefaultBackgroundColor(AValue TColor) // property
	// BoundsMode
	//  BoundsMode affects how setting the Bounds and RasterizationScale
	//  properties work. Bounds mode can either be in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS
	//  mode or COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE mode.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS, setting the bounds
	//  property will set the size of the WebView in raw screen pixels. Changing
	//  the rasterization scale in this mode won't change the raw pixel size of
	//  the WebView and will only change the rasterization scale.
	//  When the mode is in COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE, setting the
	//  bounds property will change the logical size of the WebView which can be
	//  described by the following equation: Logical size * rasterization scale = Raw Pixel size
	//  In this case, changing the rasterization scale will keep the logical size
	//  the same and change the raw pixel size.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_boundsmode">See the ICoreWebView2Controller3 article.</a>
	BoundsMode() TWVBoundsMode // property
	// SetBoundsMode Set BoundsMode
	SetBoundsMode(AValue TWVBoundsMode) // property
	// RasterizationScale
	//  The rasterization scale for the WebView. The rasterization scale is the
	//  combination of the monitor DPI scale and text scaling set by the user.
	//  This value should be updated when the DPI scale of the app's top level
	//  window changes(i.e. monitor DPI scale changes or window changes monitor)
	//  or when the text scale factor of the system changes.
	//  Rasterization scale applies to the WebView content, as well as
	//  popups, context menus, scroll bars, and so on. Normal app scaling
	//  scenarios should use the ZoomFactor property or SetBoundsAndZoomFactor
	//  API which only scale the rendered HTML content and not popups, context
	//  menus, scroll bars, and so on.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_rasterizationscale">See the ICoreWebView2Controller3 article.</a>
	RasterizationScale() (resultDouble float64) // property
	// SetRasterizationScale Set RasterizationScale
	SetRasterizationScale(AValue float64) // property
	// ShouldDetectMonitorScaleChanges
	//  ShouldDetectMonitorScaleChanges property determines whether the WebView
	//  attempts to track monitor DPI scale changes. When true, the WebView will
	//  track monitor DPI scale changes, update the RasterizationScale property,
	//  and raises RasterizationScaleChanged event. When false, the WebView will
	//  not track monitor DPI scale changes, and the app must update the
	//  RasterizationScale property itself. RasterizationScaleChanged event will
	//  never raise when ShouldDetectMonitorScaleChanges is false. Apps that want
	//  to set their own rasterization scale should set this property to false to
	//  avoid the WebView2 updating the RasterizationScale property to match the
	//  monitor DPI scale.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#get_shoulddetectmonitorscalechanges">See the ICoreWebView2Controller3 article.</a>
	ShouldDetectMonitorScaleChanges() bool // property
	// SetShouldDetectMonitorScaleChanges Set ShouldDetectMonitorScaleChanges
	SetShouldDetectMonitorScaleChanges(AValue bool) // property
	// AllowExternalDrop
	//  Gets the `AllowExternalDrop` property which is used to configure the
	//  capability that dragging objects from outside the bounds of webview2 and
	//  dropping into webview2 is allowed or disallowed. The default value is
	//  TRUE.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller4#get_allowexternaldrop">See the ICoreWebView2Controller4 article.</a>
	AllowExternalDrop() bool // property
	// SetAllowExternalDrop Set AllowExternalDrop
	SetAllowExternalDrop(AValue bool) // property
	// DefaultContextMenusEnabled
	//  The `DefaultContextMenusEnabled` property is used to prevent default
	//  context menus from being shown to user in WebView.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultcontextmenusenabled">See the ICoreWebView2Settings article.</a>
	DefaultContextMenusEnabled() bool // property
	// SetDefaultContextMenusEnabled Set DefaultContextMenusEnabled
	SetDefaultContextMenusEnabled(AValue bool) // property
	// DefaultScriptDialogsEnabled
	//  `DefaultScriptDialogsEnabled` is used when loading a new HTML
	//  document. If set to `FALSE`, WebView2 does not render the default JavaScript
	//  dialog box(Specifically those displayed by the JavaScript alert,
	//  confirm, prompt functions and `beforeunload` event). Instead, if an
	//  event handler is set using `add_ScriptDialogOpening`, WebView sends an
	//  event that contains all of the information for the dialog and allow the
	//  host app to show a custom UI. The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredefaultscriptdialogsenabled">See the ICoreWebView2Settings article.</a>
	DefaultScriptDialogsEnabled() bool // property
	// SetDefaultScriptDialogsEnabled Set DefaultScriptDialogsEnabled
	SetDefaultScriptDialogsEnabled(AValue bool) // property
	// DevToolsEnabled
	//  `DevToolsEnabled` controls whether the user is able to use the context
	//  menu or keyboard shortcuts to open the DevTools window.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_aredevtoolsenabled">See the ICoreWebView2Settings article.</a>
	DevToolsEnabled() bool // property
	// SetDevToolsEnabled Set DevToolsEnabled
	SetDevToolsEnabled(AValue bool) // property
	// AreHostObjectsAllowed
	//  The `AreHostObjectsAllowed` property is used to control whether host
	//  objects are accessible from the page in WebView.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_arehostobjectsallowed">See the ICoreWebView2Settings article.</a>
	AreHostObjectsAllowed() bool // property
	// SetAreHostObjectsAllowed Set AreHostObjectsAllowed
	SetAreHostObjectsAllowed(AValue bool) // property
	// BuiltInErrorPageEnabled
	//  The `BuiltInErrorPageEnabled` property is used to disable built in
	//  error page for navigation failure and render process failure. When
	//  disabled, a blank page is displayed when the related error happens.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isbuiltinerrorpageenabled">See the ICoreWebView2Settings article.</a>
	BuiltInErrorPageEnabled() bool // property
	// SetBuiltInErrorPageEnabled Set BuiltInErrorPageEnabled
	SetBuiltInErrorPageEnabled(AValue bool) // property
	// ScriptEnabled
	//  Controls if running JavaScript is enabled in all future navigations in
	//  the WebView. This only affects scripts in the document. Scripts
	//  injected with `ExecuteScript` runs even if script is disabled.
	//  The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isscriptenabled">See the ICoreWebView2Settings article.</a>
	ScriptEnabled() bool // property
	// SetScriptEnabled Set ScriptEnabled
	SetScriptEnabled(AValue bool) // property
	// StatusBarEnabled
	//  `StatusBarEnabled` controls whether the status bar is displayed. The
	//  status bar is usually displayed in the lower left of the WebView and
	//  shows things such as the URI of a link when the user hovers over it and
	//  other information. The default value is `TRUE`. The status bar UI can be
	//  altered by web content and should not be considered secure.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_isstatusbarenabled">See the ICoreWebView2Settings article.</a>
	StatusBarEnabled() bool // property
	// SetStatusBarEnabled Set StatusBarEnabled
	SetStatusBarEnabled(AValue bool) // property
	// WebMessageEnabled
	//  The `WebMessageEnabled` property is used when loading a new HTML
	//  document. If set to `TRUE`, communication from the host to the top-level
	//  HTML document of the WebView is allowed using `PostWebMessageAsJson`,
	//  `PostWebMessageAsString`, and message event of `window.chrome.webview`.
	//  For more information, navigate to PostWebMessageAsJson. Communication
	//  from the top-level HTML document of the WebView to the host is allowed
	//  using the postMessage function of `window.chrome.webview` and
	//  `add_WebMessageReceived` method.
	//  If set to false, then communication is disallowed. `PostWebMessageAsJson`
	//  and `PostWebMessageAsString` fails with `E_ACCESSDENIED` and
	//  `window.chrome.webview.postMessage` fails by throwing an instance of an
	//  `Error` object. The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived">See the add_WebMessageReceived method article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iswebmessageenabled">See the ICoreWebView2Settings article.</a>
	WebMessageEnabled() bool // property
	// SetWebMessageEnabled Set WebMessageEnabled
	SetWebMessageEnabled(AValue bool) // property
	// ZoomControlEnabled
	//  The `ZoomControlEnabled` property is used to prevent the user from
	//  impacting the zoom of the WebView. When disabled, the user is not able
	//  to zoom using Ctrl++, Ctrl+-, or Ctrl+mouse wheel, but the zoom
	//  is set using `ZoomFactor` API. The default value is `TRUE`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings#get_iszoomcontrolenabled">See the ICoreWebView2Settings article.</a>
	ZoomControlEnabled() bool // property
	// SetZoomControlEnabled Set ZoomControlEnabled
	SetZoomControlEnabled(AValue bool) // property
	// UserAgent
	//  Returns the User Agent. The default value is the default User Agent of the
	//  Microsoft Edge browser.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings2#get_useragent">See the ICoreWebView2Settings2 article.</a>
	UserAgent() string // property
	// SetUserAgent Set UserAgent
	SetUserAgent(AValue string) // property
	// AreBrowserAcceleratorKeysEnabled
	//  When this setting is set to FALSE, it disables all accelerator keys that
	//  access features specific to a web browser.
	//  The default value for `AreBrowserAcceleratorKeysEnabled` is TRUE.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings3#get_arebrowseracceleratorkeysenabled">See the ICoreWebView2Settings3 article.</a>
	AreBrowserAcceleratorKeysEnabled() bool // property
	// SetAreBrowserAcceleratorKeysEnabled Set AreBrowserAcceleratorKeysEnabled
	SetAreBrowserAcceleratorKeysEnabled(AValue bool) // property
	// IsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields. It will take effect
	//  immediately after setting. The default value is `TRUE`.
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
	//  It will take effect immediately after setting. The default value is `FALSE`.
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
	// Cursor
	//  The current cursor that WebView thinks it should be. The cursor should be
	//  set in WM_SETCURSOR through \::SetCursor or set on the corresponding
	//  parent/ancestor HWND of the WebView through \::SetClassLongPtr. The HCURSOR
	//  can be freed so CopyCursor/DestroyCursor is recommended to keep your own
	//  copy if you are doing more than immediately setting the cursor.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_cursor">See the ICoreWebView2CompositionController article.</a>
	Cursor() HCURSOR // property
	// RootVisualTarget
	//  The RootVisualTarget is a visual in the hosting app's visual tree. This
	//  visual is where the WebView will connect its visual tree. The app uses
	//  this visual to position the WebView within the app. The app still needs
	//  to use the Bounds property to size the WebView. The RootVisualTarget
	//  property can be an IDCompositionVisual or a
	//  Windows::UI::Composition::ContainerVisual. WebView will connect its visual
	//  tree to the provided visual before returning from the property setter. The
	//  app needs to commit on its device setting the RootVisualTarget property.
	//  The RootVisualTarget property supports being set to nullptr to disconnect
	//  the WebView from the app's visual tree.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_rootvisualtarget">See the ICoreWebView2CompositionController article.</a>
	RootVisualTarget() IUnknown // property
	// SetRootVisualTarget Set RootVisualTarget
	SetRootVisualTarget(AValue IUnknown) // property
	// SystemCursorID
	//  The current system cursor ID reported by the underlying rendering engine
	//  for WebView. For example, most of the time, when the cursor is over text,
	//  this will return the int value for IDC_IBEAM. The systemCursorId is only
	//  valid if the rendering engine reports a default Windows cursor resource
	//  value. Navigate to
	//  [LoadCursorW](/windows/win32/api/winuser/nf-winuser-loadcursorw) for more
	//  details. Otherwise, if custom CSS cursors are being used, this will return
	//  0. To actually use systemCursorId in LoadCursor or LoadImage,
	//  MAKEINTRESOURCE must be called on it first.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_systemcursorid">See the ICoreWebView2CompositionController article.</a>
	SystemCursorID() uint32 // property
	// AutomationProvider
	//  Returns the Automation Provider for the WebView. This object implements
	//  IRawElementProviderSimple.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller2#get_automationprovider">See the ICoreWebView2CompositionController2 article.</a>
	AutomationProvider() IUnknown // property
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`. Provide a list of all
	//  process using same user data folder except for crashpad process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8#getprocessinfos">See the ICoreWebView2Environment8 article.</a>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property
	// ProfileName
	//  `ProfileName` property is to specify a profile name, which is only allowed to contain
	//  the following ASCII characters. It has a maximum length of 64 characters excluding the null-terminator.
	//  It is ASCII case insensitive.
	//  * alphabet characters: a-z and A-Z
	//  * digit characters: 0-9
	//  * and '#', '@', '$', '(', ')', '+', '-', '_', '~', '.', ' '(space).
	//  Note: the text must not end with a period '.' or ' '(space). And, although upper-case letters are
	//  allowed, they're treated just as lower-case counterparts because the profile name will be mapped to
	//  the real profile directory path on disk and Windows file system handles path names in a case-insensitive way.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_profilename">See the ICoreWebView2ControllerOptions article.</a>
	ProfileName() string // property
	// SetProfileName Set ProfileName
	SetProfileName(AValue string) // property
	// IsInPrivateModeEnabled
	//  `IsInPrivateModeEnabled` property is to enable/disable InPrivate mode.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions#get_isinprivatemodeenabled">See the ICoreWebView2ControllerOptions article.</a>
	IsInPrivateModeEnabled() bool // property
	// SetIsInPrivateModeEnabled Set IsInPrivateModeEnabled
	SetIsInPrivateModeEnabled(AValue bool) // property
	// ScriptLocale
	//  The default locale for the WebView2. It sets the default locale for all
	//  Intl JavaScript APIs and other JavaScript APIs that depend on it, namely
	//  `Intl.DateTimeFormat()` which affects string formatting like
	//  in the time/date formats. Example: `Intl.DateTimeFormat().format(new Date())`
	//  The intended locale value is in the format of
	//  BCP 47 Language Tags. More information can be found from
	//  [IETF BCP47](https://www.ietf.org/rfc/bcp/bcp47.html).
	//  This property sets the locale for a CoreWebView2Environment used to create the
	//  WebView2ControllerOptions object, which is passed as a parameter in
	//  `CreateCoreWebView2ControllerWithOptions`.
	//  Changes to the ScriptLocale property apply to renderer processes created after
	//  the change. Any existing renderer processes will continue to use the previous
	//  ScriptLocale value. To ensure changes are applied to all renderer process,
	//  close and restart the CoreWebView2Environment and all associated WebView2 objects.
	//  The default value for ScriptLocale will depend on the WebView2 language
	//  and OS region. If the language portions of the WebView2 language and OS region
	//  match, then it will use the OS region. Otherwise, it will use the WebView2
	//  language.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controlleroptions2#get_scriptlocale">See the ICoreWebView2ControllerOptions2 article.</a>
	ScriptLocale() string // property
	// SetScriptLocale Set ScriptLocale
	SetScriptLocale(AValue string) // property
	// ProfilePath
	//  Full path of the profile directory.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_profilepath">See the ICoreWebView2Profile article.</a>
	ProfilePath() string // property
	// DefaultDownloadFolderPath
	//  Gets the `DefaultDownloadFolderPath` property. The default value is the
	//  system default download folder path for the user.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_defaultdownloadfolderpath">See the ICoreWebView2Profile article.</a>
	DefaultDownloadFolderPath() string // property
	// SetDefaultDownloadFolderPath Set DefaultDownloadFolderPath
	SetDefaultDownloadFolderPath(AValue string) // property
	// PreferredColorScheme
	//  The PreferredColorScheme property sets the overall color scheme of the
	//  WebView2s associated with this profile. This sets the color scheme for
	//  WebView2 UI like dialogs, prompts, and context menus by setting the
	//  media feature `prefers-color-scheme` for websites to respond to.
	//  The default value for this is COREWEBVIEW2_PREFERRED_COLOR_AUTO,
	//  which will follow whatever theme the OS is currently set to.
	//  Returns the value of the `PreferredColorScheme` property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_preferredcolorscheme">See the ICoreWebView2Profile article.</a>
	PreferredColorScheme() TWVPreferredColorScheme // property
	// SetPreferredColorScheme Set PreferredColorScheme
	SetPreferredColorScheme(AValue TWVPreferredColorScheme) // property
	// PreferredTrackingPreventionLevel
	//  The `PreferredTrackingPreventionLevel` property allows you to control levels of tracking prevention for WebView2
	//  which are associated with a profile. This level would apply to the context of the profile. That is, all WebView2s
	//  sharing the same profile will be affected and also the value is persisted in the user data folder.
	//  See `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL` for descriptions of levels.
	//  If tracking prevention feature is enabled when creating the WebView2 environment, you can also disable tracking
	//  prevention later using this property and `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE` value but that doesn't
	//  improves runtime performance.
	//  There is `ICoreWebView2EnvironmentOptions5.EnableTrackingPrevention` property to enable/disable tracking prevention feature
	//  for all the WebView2's created in the same environment. If enabled, `PreferredTrackingPreventionLevel` is set to
	//  `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED` by default for all the WebView2's and profiles created in the same
	//  environment or is set to the level whatever value was last changed/persisted to the profile. If disabled
	//  `PreferredTrackingPreventionLevel` is not respected by WebView2. If `PreferredTrackingPreventionLevel` is set when the
	//  feature is disabled, the property value get changed and persisted but it will takes effect only if
	//  `ICoreWebView2EnvironmentOptions5.EnableTrackingPrevention` is true.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile3#get_preferredtrackingpreventionlevel">See the ICoreWebView2Profile3 article.</a>
	PreferredTrackingPreventionLevel() TWVTrackingPreventionLevel // property
	// SetPreferredTrackingPreventionLevel Set PreferredTrackingPreventionLevel
	SetPreferredTrackingPreventionLevel(AValue TWVTrackingPreventionLevel) // property
	// ProfileCookieManager
	//  Get the cookie manager for the profile. All CoreWebView2s associated with this
	//  profile share the same cookie values. Changes to cookies in this cookie manager apply to all
	//  CoreWebView2s associated with this profile. See ICoreWebView2CookieManager.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile5#get_cookiemanager">See the ICoreWebView2Profile5 article.</a>
	ProfileCookieManager() ICoreWebView2CookieManager // property
	// ProfileIsPasswordAutosaveEnabled
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
	//  It will take effect immediately after setting. The default value is `FALSE`.
	//  This property has the same value as
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled` and
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled` will always have the same
	//  value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_ispasswordautosaveenabled">See the ICoreWebView2Profile6 article.</a>
	ProfileIsPasswordAutosaveEnabled() bool // property
	// SetProfileIsPasswordAutosaveEnabled Set ProfileIsPasswordAutosaveEnabled
	SetProfileIsPasswordAutosaveEnabled(AValue bool) // property
	// ProfileIsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields.
	//  It will take effect immediately after setting. The default value is `TRUE`.
	//  This property has the same value as
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled` and
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled` will always have the same
	//  value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_isgeneralautofillenabled">See the ICoreWebView2Profile6 article.</a>
	ProfileIsGeneralAutofillEnabled() bool // property
	// SetProfileIsGeneralAutofillEnabled Set ProfileIsGeneralAutofillEnabled
	SetProfileIsGeneralAutofillEnabled(AValue bool) // property
	// FrameId
	//  The unique identifier of the main frame. It's the same kind of ID as
	//  with the `FrameId` in `ICoreWebView2Frame` and via `ICoreWebView2FrameInfo`.
	//  Note that `FrameId` may not be valid if `ICoreWebView2` has not done
	//  any navigation. It's safe to get this value during or after the first
	//  `ContentLoading` event. Otherwise, it could return the invalid frame Id 0.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_20#get_frameid">See the ICoreWebView2_20 article.</a>
	FrameId() uint32 // property
	// CreateBrowser
	//  Used to create the browser using the global environment by default.
	//  The browser will be fully initialized when the TWVBrowserBase.OnAfterCreated
	//  event is triggered.
	//  <param name="aHandle">The TWVWindowParent handle.</param>
	//  <param name="aUseDefaultEnvironment">Use the global environment or create a new one for this browser.</param>
	CreateBrowser(aHandle THandle, aUseDefaultEnvironment bool) bool // function
	// CreateBrowser1
	//  Used to create the browser using a custom environment. The browser will be
	//  fully initialized when the TWVBrowserBase.OnAfterCreated event is triggered.
	//  <param name="aHandle">The TWVWindowParent handle.</param>
	//  <param name="aEnvironment">Custom environment to be used by this browser.</param>
	CreateBrowser1(aHandle THandle, aEnvironment ICoreWebView2Environment) bool // function
	// CreateWindowlessBrowser
	//  Used to create a windowless browser using the global environment by default.
	//  The browser will be fully initialized when the TWVBrowserBase.OnAfterCreated
	//  event is triggered.
	//  <param name="aHandle">The TWVDirectCompositionHost handle.</param>
	//  <param name="aUseDefaultEnvironment">Use the global environment or create a new one for this browser.</param>
	CreateWindowlessBrowser(aHandle THandle, aUseDefaultEnvironment bool) bool // function
	// CreateWindowlessBrowser1
	//  Used to create a windowless browser using a custom environment. The browser will be
	//  fully initialized when the TWVBrowserBase.OnAfterCreated event is triggered.
	//  <param name="aHandle">The TWVDirectCompositionHost handle.</param>
	//  <param name="aEnvironment">Custom environment to be used by this browser.</param>
	CreateWindowlessBrowser1(aHandle THandle, aEnvironment ICoreWebView2Environment) bool // function
	// GoBack
	//  Navigates the WebView to the previous page in the navigation history.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#goback">See the ICoreWebView2 article.</a>
	GoBack() bool // function
	// GoForward
	//  Navigates the WebView to the next page in the navigation history.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#goforward">See the ICoreWebView2 article.</a>
	GoForward() bool // function
	// Refresh
	//  Reload the current page. This is similar to navigating to the URI of
	//  current top level document including all navigation events firing and
	//  respecting any entries in the HTTP cache. But, the back or forward
	//  history are not modified.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#reload">See the ICoreWebView2 article.</a>
	Refresh() bool // function
	// RefreshIgnoreCache
	//  Reload the current page. Browser cache is ignored as if the user pressed Shift+refresh.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnRefreshIgnoreCacheCompleted event when it finishes executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload">See the Page Domain article.</a>
	RefreshIgnoreCache() bool // function
	// Stop
	//  Stop all navigations and pending resource fetches. Does not stop scripts.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#stop">See the ICoreWebView2 article.</a>
	Stop() bool // function
	// Navigate
	//  Cause a navigation of the top-level document to run to the specified URI.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#navigate">See the ICoreWebView2 article.</a>
	Navigate(aURI string) bool // function
	// NavigateToString
	//  Initiates a navigation to aHTMLContent as source HTML of a new document.
	//  The origin of the new page is `about:blank`.
	//  <param name="aHTMLContent">Source HTML. It may not be larger than 2 MB(2 * 1024 * 1024 bytes) in total size.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#navigatetostring">See the ICoreWebView2 article.</a>
	NavigateToString(aHTMLContent string) bool // function
	// NavigateWithWebResourceRequest
	//  Navigates using a constructed ICoreWebView2WebResourceRequest object. This lets you
	//  provide post data or additional request headers during navigation.
	//  The headers in aRequest override headers added by WebView2 runtime except for Cookie headers.
	//  Method can only be either "GET" or "POST". Provided post data will only
	//  be sent only if the method is "POST" and the uri scheme is HTTP(S).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#navigatewithwebresourcerequest">See the ICoreWebView2_2 article.</a>
	NavigateWithWebResourceRequest(aRequest ICoreWebView2WebResourceRequestRef) bool // function
	// SubscribeToDevToolsProtocolEvent
	//  Subscribe to a DevTools protocol event. The TWVBrowserBase.OnDevToolsProtocolEventReceived
	//  event will be triggered on each DevTools event.
	//  <param name="aEventName">The DevTools protocol event name.</param>
	//  <param name="aEventID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceiver#add_devtoolsprotocoleventreceived">See the ICoreWebView2DevToolsProtocolEventReceiver article.</a>
	SubscribeToDevToolsProtocolEvent(aEventName string, aEventID int32) bool // function
	// CallDevToolsProtocolMethod
	//  Runs an asynchronous `DevToolsProtocol` method.
	//  The TWVBrowserBase.OnCallDevToolsProtocolMethodCompleted event is triggered
	//  when it finishes executing. This function returns E_INVALIDARG if the `aMethodName` is
	//  unknown or the `aParametersAsJson` has an error. In the case of such an error, the
	//  `aReturnObjectAsJson` parameter of the event will include information
	//  about the error.
	//  Note even though WebView2 dispatches the CDP messages in the order called,
	//  CDP method calls may be processed out of order.
	//  If you require CDP methods to run in a particular order, you should wait
	//  for the previous method is finished before calling the next method.
	//  <param name="aMethodName">The DevTools protocol full method name.</param>
	//  <param name="aParametersAsJson">JSON formatted string containing the parameters for the corresponding method.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot">See the Chrome DevTools Protocol web page.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#calldevtoolsprotocolmethod">See the ICoreWebView2 article.</a>
	CallDevToolsProtocolMethod(aMethodName, aParametersAsJson string, aExecutionID int32) bool // function
	// CallDevToolsProtocolMethodForSession
	//  Runs an asynchronous `DevToolsProtocol` method for a specific session of
	//  an attached target.
	//  There could be multiple `DevToolsProtocol` targets in a WebView.
	//  Besides the top level page, iframes from different origin and web workers
	//  are also separate targets. Attaching to these targets allows interaction with them.
	//  When the DevToolsProtocol is attached to a target, the connection is identified
	//  by a sessionId.
	//  To use this API, you must set the `flatten` parameter to true when calling
	//  `Target.attachToTarget` or `Target.setAutoAttach` `DevToolsProtocol` method.
	//  Using `Target.setAutoAttach` is recommended as that would allow you to attach
	//  to dedicated worker targets, which are not discoverable via other APIs like
	//  `Target.getTargets`.
	//  The TWVBrowserBase.OnCallDevToolsProtocolMethodCompleted event is triggered
	//  when it finishes executing. This function returns E_INVALIDARG if the `aMethodName` is
	//  unknown or the `aParametersAsJson` has an error. In the case of such an error, the
	//  `aReturnObjectAsJson` parameter of the event will include information
	//  about the error.
	//  Note even though WebView2 dispatches the CDP messages in the order called,
	//  CDP method calls may be processed out of order.
	//  If you require CDP methods to run in a particular order, you should wait
	//  for the previous method is finished before calling the next method.
	//  <param name="aSessionId">The sessionId for an attached target. An empty string is treated as the session for the default target for the top page.</param>
	//  <param name="aMethodName">The DevTools protocol full method name.</param>
	//  <param name="aParametersAsJson">JSON formatted string containing the parameters for the corresponding method.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot">See the Chrome DevTools Protocol web page.</a>
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Target">Information about targets and sessions.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_11#calldevtoolsprotocolmethodforsession">See the ICoreWebView2_11 article.</a>
	CallDevToolsProtocolMethodForSession(aSessionId, aMethodName, aParametersAsJson string, aExecutionID int32) bool // function
	// SetFocus
	//  Moves focus into WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#movefocus">See the ICoreWebView2Controller article.</a>
	SetFocus() bool // function
	// FocusNext
	//  Moves the focus to the next element.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#movefocus">See the ICoreWebView2Controller article.</a>
	FocusNext() bool // function
	// FocusPrevious
	//  Moves the focus to the previous element.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#movefocus">See the ICoreWebView2Controller article.</a>
	FocusPrevious() bool // function
	// ExecuteScriptWithResult
	//  Run JavaScript code from the JavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The TWVBrowserBase.OnExecuteScriptWithResultCompleted event is triggered
	//  when it finishes executing.
	//  The result of the execution is returned asynchronously in the ICoreWebView2ExecuteScriptResult object
	//  which has methods and properties to obtain the successful result of script execution as well as any
	//  unhandled JavaScript exceptions.
	//  If this method is run after the NavigationStarting event during a navigation, the script
	//  runs in the new document when loading it, around the time
	//  ContentLoading is run. This operation executes the script even if
	//  ICoreWebView2Settings.IsScriptEnabled is set to FALSE.
	//  <param name="aJavaScript">The JavaScript code.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_21#executescriptwithresult">See the icorewebview2_21 article.</a>
	ExecuteScriptWithResult(aJavaScript string, aExecutionID int32) bool // function
	// ExecuteScript
	//  Run JavaScript code from the aJavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The TWVBrowserBase.OnExecuteScriptCompleted event is triggered
	//  when it finishes executing.
	//  The result of evaluating the provided JavaScript is available in the
	//  aResultObjectAsJson parameter of the TWVBrowserBase.OnExecuteScriptCompleted
	//  event as a JSON encoded string. If the result is undefined, contains a reference
	//  cycle, or otherwise is not able to be encoded into JSON, then the result
	//  is considered to be null, which is encoded in JSON as the string "null".
	//  If the script that was run throws an unhandled exception, then the result is
	//  also "null".
	//  If the method is run after the `NavigationStarting` event during a navigation,
	//  the script runs in the new document when loading it, around the time
	//  `ContentLoading` is run. This operation executes the script even if
	//  `ICoreWebView2Settings.IsScriptEnabled` is set to `FALSE`.
	//  <param name="aJavaScript">The JavaScript code.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#executescript">See the ICoreWebView2 article.</a>
	ExecuteScript(aJavaScript string, aExecutionID int32) bool // function
	// CapturePreview
	//  Capture an image of what WebView is displaying. Specify the format of
	//  the image with the aImageFormat parameter. The resulting image binary
	//  data is written to the provided aImageStream parameter. This method fails if called
	//  before the first ContentLoading event. For example if this is called in
	//  the NavigationStarting event for the first navigation it will fail.
	//  For subsequent navigations, the method may not fail, but will not capture
	//  an image of a given webpage until the ContentLoading event has been fired
	//  for it. Any call to this method prior to that will result in a capture of
	//  the page being navigated away from. When this function finishes writing to the stream,
	//  the TWVBrowserBase.OnCapturePreviewCompleted event is triggered.
	//  <param name="aImageFormat">The format of the image.</param>
	//  <param name="aImageStream">The resulting image binary data is written to this stream.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#capturepreview">See the ICoreWebView2 article.</a>
	CapturePreview(aImageFormat TWVCapturePreviewImageFormat, aImageStream IStream) bool // function
	// NotifyParentWindowPositionChanged
	//  This is a notification that tells WebView that the main WebView parent
	// (or any ancestor) `HWND` moved. This is needed for accessibility and
	//  certain dialogs in WebView to work correctly.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#notifyparentwindowpositionchanged">See the ICoreWebView2Controller article.</a>
	NotifyParentWindowPositionChanged() bool // function
	// SetPermissionState
	//  Sets permission state for the given permission kind and origin
	//  asynchronously. The change persists across sessions until it is changed by
	//  another call to `SetPermissionState`, or by setting the `State` property
	//  in `PermissionRequestedEventArgs`.
	//  Setting the state to `COREWEBVIEW2_PERMISSION_STATE_DEFAULT` will
	//  erase any state saved in the profile and restore the default behavior.
	//  The origin should have a valid scheme and host(e.g. "https://www.example.com"),
	//  otherwise the method fails with `E_INVALIDARG`. Additional URI parts like
	//  path and fragment are ignored. For example, "https://wwww.example.com/app1/index.html/"
	//  is treated the same as "https://wwww.example.com".
	//  This function triggers the TWVBrowserBase.OnSetPermissionStateCompleted event.
	//  <a href="https://developer.mozilla.org/en-US/docs/Glossary/Origin">See the MDN origin definition.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#setpermissionstate">See the ICoreWebView2Profile4 article.</a>
	SetPermissionState(aPermissionKind TWVPermissionKind, aOrigin string, aState TWVPermissionState) bool // function
	// GetNonDefaultPermissionSettings
	//  Invokes the handler with a collection of all nondefault permission settings.
	//  Use this method to get the permission state set in the current and previous
	//  sessions.
	//  This function triggers the TWVBrowserBase.OnGetNonDefaultPermissionSettingsCompleted event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#getnondefaultpermissionsettings">See the ICoreWebView2Profile4 article.</a>
	GetNonDefaultPermissionSettings() bool // function
	// AddBrowserExtension
	//  Adds the [browser extension](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions)
	//  using the extension path for unpacked extensions from the local device. Extension is
	//  running right after installation.
	//  The extension folder path is the topmost folder of an unpacked browser extension and
	//  contains the browser extension manifest file.
	//  If the `extensionFolderPath` is an invalid path or doesn't contain the extension manifest.json
	//  file, this function will return `ERROR_FILE_NOT_FOUND` to callers.
	//  Installed extension will default `IsEnabled` to true.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `AddBrowserExtension` will fail and return
	//  HRESULT `ERROR_NOT_SUPPORTED`.
	//  During installation, the content of the extension is not copied to the user data folder.
	//  Once the extension is installed, changing the content of the extension will cause the
	//  extension to be removed from the installed profile.
	//  When an extension is added the extension is persisted in the corresponding profile. The
	//  extension will still be installed the next time you use this profile.
	//  When an extension is installed from a folder path, adding the same extension from the same
	//  folder path means reinstalling this extension. When two extensions with the same Id are
	//  installed, only the later installed extension will be kept.
	//  Extensions that are designed to include any UI interactions(e.g. icon, badge, pop up, etc.)
	//  can be loaded and used but will have missing UI entry points due to the lack of browser
	//  UI elements to host these entry points in WebView2.
	//  The following summarizes the possible error values that can be returned from
	//  `AddBrowserExtension` and a description of why these errors occur.
	//  <code>
	//  Error value | Description
	//  ----------------------------------------------- | --------------------------
	//  `HRESULT_FROM_WIN32(ERROR_NOT_SUPPORTED)` | Extensions are disabled.
	//  `HRESULT_FROM_WIN32(ERROR_FILE_NOT_FOUND)` | Cannot find `manifest.json` file or it is not a valid extension manifest.
	//  `E_ACCESSDENIED` | Cannot load extension with file or directory name starting with \"_\", reserved for use by the system.
	//  `E_FAIL` | Extension failed to install with other unknown reasons.
	//  </code>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#addbrowserextension">See the ICoreWebView2Profile7 article.</a>
	//  The TWVBrowserBase.OnProfileAddBrowserExtensionCompleted event is triggered when TWVBrowserBase.AddBrowserExtension finishes executing.
	AddBrowserExtension(extensionFolderPath string) bool // function
	// GetBrowserExtensions
	//  Gets a snapshot of the set of extensions installed at the time `GetBrowserExtensions` is
	//  called. If an extension is installed or uninstalled after `GetBrowserExtensions` completes,
	//  the list returned by `GetBrowserExtensions` remains the same.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `GetBrowserExtensions` won't return any
	//  extensions on current user profile.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#getbrowserextensions">See the ICoreWebView2Profile7 article.</a>
	//  The TWVBrowserBase.OnProfileGetBrowserExtensionsCompleted event is triggered when TWVBrowserBase.GetBrowserExtensions finishes executing.
	GetBrowserExtensions() bool // function
	// DeleteProfile
	//  After the API is called, the profile will be marked for deletion. The
	//  local profile's directory will be deleted at browser process exit. If it
	//  fails to delete, because something else is holding the files open,
	//  WebView2 will try to delete the profile at all future browser process
	//  starts until successful.
	//  The corresponding CoreWebView2s will be closed and the
	//  CoreWebView2Profile.Deleted event will be raised. See
	//  `CoreWebView2Profile.Deleted` for more information.
	//  If you try to create a new profile with the same name as an existing
	//  profile that has been marked as deleted but hasn't yet been deleted,
	//  profile creation will fail with HRESULT_FROM_WIN32(ERROR_DELETE_PENDING).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile8#delete">See the ICoreWebView2Profile8 article.</a>
	DeleteProfile() bool // function
	// TrySuspend
	//  An app may call the `TrySuspend` API to have the WebView2 consume less memory.
	//  This is useful when a Win32 app becomes invisible, or when a Universal Windows
	//  Platform app is being suspended, during the suspended event handler before completing
	//  the suspended event.
	//  The IsVisible property must be false when the API is called.
	//  Otherwise, the API fails with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Suspending is similar to putting a tab to sleep in the Edge browser. Suspending pauses
	//  WebView script timers and animations, minimizes CPU usage for the associated
	//  browser renderer process and allows the operating system to reuse the memory that was
	//  used by the renderer process for other processes.
	//  Note that Suspend is best effort and considered completed successfully once the request
	//  is sent to browser renderer process. If there is a running script, the script will continue
	//  to run and the renderer process will be suspended after that script is done.
	//  for conditions that might prevent WebView from being suspended. In those situations,
	//  the completed handler will be invoked with isSuccessful as false and errorCode as S_OK.
	//  The WebView will be automatically resumed when it becomes visible. Therefore, the
	//  app normally does not have to call `Resume` explicitly.
	//  The app can call `Resume` and then `TrySuspend` periodically for an invisible WebView so that
	//  the invisible WebView can sync up with latest data and the page ready to show fresh content
	//  when it becomes visible.
	//  All WebView APIs can still be accessed when a WebView is suspended. Some APIs like Navigate
	//  will auto resume the WebView. To avoid unexpected auto resume, check `IsSuspended` property
	//  before calling APIs that might change WebView state.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnTrySuspendCompleted event when it finishes executing.
	//  <a href="https://techcommunity.microsoft.com/t5/articles/sleeping-tabs-faq/m-p/1705434">See the sleeping Tabs FAQ.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#trysuspend">See the ICoreWebView2_3 article.</a>
	TrySuspend() bool // function
	// Resume
	//  Resumes the WebView so that it resumes activities on the web page.
	//  This API can be called while the WebView2 controller is invisible.
	//  The app can interact with the WebView immediately after `Resume`.
	//  WebView will be automatically resumed when it becomes visible.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#resume">See the ICoreWebView2_3 article.</a>
	Resume() bool // function
	// SetVirtualHostNameToFolderMapping
	//  Sets a mapping between a virtual host name and a folder path to make available to web sites
	//  via that host name.
	//  After setting the mapping, documents loaded in the WebView can use HTTP or HTTPS URLs at
	//  the specified host name specified by hostName to access files in the local folder specified
	//  by folderPath.
	//  <param name="aHostName">Host name to access files in the local folder specified by aFolderPath.</param>
	//  <param name="aFolderPath">The path to the local files. Both absolute and relative paths are supported. Relative paths are interpreted as relative to the folder where the exe of the app is in. Note that the aFolderPath length must not exceed the Windows MAX_PATH limit.</param>
	//  <param name="aAccessKind">aAccessKind specifies the level of access to resources under the virtual host from other sites.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#setvirtualhostnametofoldermapping">See the ICoreWebView2_3 article.</a>
	SetVirtualHostNameToFolderMapping(aHostName, aFolderPath string, aAccessKind TWVHostResourceAcccessKind) bool // function
	// ClearVirtualHostNameToFolderMapping
	//  Clears a host name mapping for local folder that was added by `SetVirtualHostNameToFolderMapping`.
	//  <param name="aHostName">Host name used previously with SetVirtualHostNameToFolderMapping to access files in the local folder.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#clearvirtualhostnametofoldermapping">See the ICoreWebView2_3 article.</a>
	ClearVirtualHostNameToFolderMapping(aHostName string) bool // function
	// RetrieveHTML
	//  Retrieve the HTML contents. The TWVBrowserBase.OnRetrieveHTMLCompleted event is triggered asynchronously with the HTML contents.
	RetrieveHTML() bool // function
	// RetrieveText
	//  Retrieve the text contents. The TWVBrowserBase.OnRetrieveTextCompleted event is triggered asynchronously with the text contents.
	RetrieveText() bool // function
	// RetrieveMHTML
	//  Retrieve the web page contents in MHTML format. The TWVBrowserBase.OnRetrieveMHTMLCompleted event is triggered asynchronously with the MHTML contents.
	RetrieveMHTML() bool // function
	// Print
	//  Print the current web page asynchronously to the specified printer with the TWVBrowserBase.CoreWebView2PrintSettings settings.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnPrintCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#print">See the ICoreWebView2_16 article.</a>
	Print() bool // function
	// ShowPrintUI
	//  Opens the print dialog to print the current web page using the system print dialog or the browser print preview dialog.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#showprintui">See the ICoreWebView2_16 article.</a>
	ShowPrintUI(aUseSystemPrintDialog bool) bool // function
	// PrintToPdf
	//  Print the current page to PDF asynchronously with the TWVBrowserBase.CoreWebView2PrintSettings settings.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnPrintToPdfCompleted event when it finishes executing.
	//  <param name="aResultFilePath">The path to the PDF file.</param>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_7#printtopdf">See the ICoreWebView2_7 article.</a>
	PrintToPdf(aResultFilePath string) bool // function
	// PrintToPdfStream
	//  Provides the Pdf data of current web page asynchronously for the TWVBrowserBase.CoreWebView2PrintSettings settings.
	//  Stream will be rewound to the start of the pdf data.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnPrintToPdfStreamCompleted event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#printtopdfstream">See the ICoreWebView2_16 article.</a>
	PrintToPdfStream() bool // function
	// OpenDevToolsWindow
	//  Opens the DevTools window for the current document in the WebView. Does
	//  nothing if run when the DevTools window is already open.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#opendevtoolswindow">See the ICoreWebView2 article.</a>
	OpenDevToolsWindow() bool // function
	// OpenTaskManagerWindow
	//  Opens the Browser Task Manager view as a new window in the foreground.
	//  If the Browser Task Manager is already open, this will bring it into
	//  the foreground. WebView2 currently blocks the Shift+Esc shortcut for
	//  opening the task manager. An end user can open the browser task manager
	//  manually via the `Browser task manager` entry of the DevTools window's
	//  title bar's context menu.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_6#opentaskmanagerwindow">See the ICoreWebView2_6 article.</a>
	OpenTaskManagerWindow() bool // function
	// PostWebMessageAsJson
	//  Post the specified webMessage to the top level document in this WebView.
	//  The main page receives the message by subscribing to the `message` event of the
	//  `window.chrome.webview` of the page document.
	//  <code>
	//  ```cpp
	//  window.chrome.webview.addEventListener('message', handler)
	//  window.chrome.webview.removeEventListener('message', handler)
	//  ```
	//  </code>
	//  The event args is an instance of `MessageEvent`. The
	//  `ICoreWebView2Settings.IsWebMessageEnabled` setting must be `TRUE` or
	//  the web message will not be sent. The `data` property of the event
	//  arg is the `webMessage` string parameter parsed as a JSON string into a
	//  JavaScript object. The `source` property of the event arg is a reference
	//  to the `window.chrome.webview` object. For information about sending
	//  messages from the HTML document in the WebView to the host, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  The message is delivered asynchronously. If a navigation occurs before
	//  the message is posted to the page, the message is discarded.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#postwebmessageasjson">See the ICoreWebView2 article.</a>
	PostWebMessageAsJson(aWebMessageAsJson string) bool // function
	// PostWebMessageAsString
	//  Posts a message that is a simple string rather than a JSON string
	//  representation of a JavaScript object. This behaves in exactly the same
	//  manner as `PostWebMessageAsJson`, but the `data` property of the event
	//  arg of the `window.chrome.webview` message is a string with the same
	//  value as `aWebMessageAsString`. Use this instead of
	//  `PostWebMessageAsJson` if you want to communicate using simple strings
	//  rather than JSON objects.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#postwebmessageasstring">See the ICoreWebView2 article.</a>
	PostWebMessageAsString(aWebMessageAsString string) bool // function
	// AddWebResourceRequestedFilter
	//  Adds a URI and resource context filter for the `WebResourceRequested`
	//  event. A web resource request with a resource context that matches this
	//  filter's resource context and a URI that matches this filter's URI
	//  wildcard string will be raised via the `WebResourceRequested` event.
	//  The `aURI` parameter value is a wildcard string matched against the URI
	//  of the web resource request. This is a glob style
	//  wildcard string in which a `*` matches zero or more characters and a `?`
	//  matches exactly one character.
	//  These wildcard characters can be escaped using a backslash just before
	//  the wildcard character in order to represent the literal `*` or `?`.
	//  The matching occurs over the URI as a whole string and not limiting
	//  wildcard matches to particular parts of the URI.
	//  The wildcard filter is compared to the URI after the URI has been
	//  normalized, any URI fragment has been removed, and non-ASCII hostnames
	//  have been converted to punycode.
	//  Specifying an empty string for aURI matches no URIs.
	//  For more information about resource context filters, navigate to
	//  [COREWEBVIEW2_WEB_RESOURCE_CONTEXT](/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_web_resource_context).
	//  <code>
	//  | URI Filter String | Request URI | Match | Notes |
	//  | ---- | ---- | ---- | ---- |
	//  | `*` | `https://contoso.com/a/b/c` | Yes | A single * will match all URIs |
	//  | `*://contoso.com/*` | `https://contoso.com/a/b/c` | Yes | Matches everything in contoso.com across all schemes |
	//  | `*://contoso.com/*` | `https://example.com/?https://contoso.com/` | Yes | But also matches a URI with just the same text anywhere in the URI |
	//  | `example` | `https://contoso.com/example` | No | The filter does not perform partial matches |
	//  | `*example` | `https://contoso.com/example` | Yes | The filter matches across URI parts |
	//  | `*example` | `https://contoso.com/path/?example` | Yes | The filter matches across URI parts |
	//  | `*example` | `https://contoso.com/path/?query#example` | No | The filter is matched against the URI with no fragment |
	//  | `*example` | `https://example` | No | The URI is normalized before filter matching so the actual URI used for comparison is `https://example/` |
	//  | `*example/` | `https://example` | Yes | Just like above, but this time the filter ends with a / just like the normalized URI |
	//  | `https://xn--qei.example/` | `https://&#x2764;.example/` | Yes | Non-ASCII hostnames are normalized to punycode before wildcard comparison |
	//  | `https://&#x2764;.example/` | `https://xn--qei.example/` | No | Non-ASCII hostnames are normalized to punycode before wildcard comparison |
	//  </code>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addwebresourcerequestedfilter">See the ICoreWebView2 article.</a>
	AddWebResourceRequestedFilter(aURI string, aResourceContext TWVWebResourceContext) bool // function
	// RemoveWebResourceRequestedFilter
	//  Removes a matching WebResource filter that was previously added for the
	//  `WebResourceRequested` event. If the same filter was added multiple
	//  times, then it must be removed as many times as it was added for the
	//  removal to be effective. Returns false for a filter that was
	//  never added.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#removewebresourcerequestedfilter">See the ICoreWebView2 article.</a>
	RemoveWebResourceRequestedFilter(aURI string, aResourceContext TWVWebResourceContext) bool // function
	// RemoveHostObjectFromScript
	//  Remove the host object specified by the name so that it is no longer
	//  accessible from JavaScript code in the WebView. While new access
	//  attempts are denied, if the object is already obtained by JavaScript code
	//  in the WebView, the JavaScript code continues to have access to that
	//  object. Run this method for a name that is already removed or never
	//  added fails.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#removehostobjectfromscript">See the ICoreWebView2 article.</a>
	RemoveHostObjectFromScript(aName string) bool // function
	// AddScriptToExecuteOnDocumentCreated
	//  Add the provided JavaScript to a list of scripts that should be run after
	//  the global object has been created, but before the HTML document has
	//  been parsed and before any other script included by the HTML document is
	//  run. This method injects a script that runs on all top-level document
	//  and child frame page navigations. This method runs asynchronously, and
	//  you must wait for the completion handler to finish before the injected
	//  script is ready to run. When this method completes, the `Invoke` method
	//  of the handler is run with the `id` of the injected script. `id` is a
	//  string. To remove the injected script, use
	//  `RemoveScriptToExecuteOnDocumentCreated`.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnAddScriptToExecuteOnDocumentCreatedCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addscripttoexecuteondocumentcreated">See the ICoreWebView2 article.</a>
	AddScriptToExecuteOnDocumentCreated(JavaScript string) bool // function
	// RemoveScriptToExecuteOnDocumentCreated
	//  Remove the corresponding JavaScript added using
	//  `AddScriptToExecuteOnDocumentCreated` with the specified script ID. The
	//  script ID should be the one returned by the `AddScriptToExecuteOnDocumentCreated`.
	//  Both use `AddScriptToExecuteOnDocumentCreated` and this method in `NewWindowRequested`
	//  event handler at the same time sometimes causes trouble. Since invalid scripts will
	//  be ignored, the script IDs you got may not be valid anymore.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#removescripttoexecuteondocumentcreated">See the ICoreWebView2 article.</a>
	RemoveScriptToExecuteOnDocumentCreated(aID string) bool // function
	// CreateCookie
	//  Create a cookie object with a specified name, value, domain, and path.
	//  One can set other optional properties after cookie creation.
	//  This only creates a cookie object and it is not added to the cookie
	//  manager until you call AddOrUpdateCookie.
	//  Leading or trailing whitespace(s), empty string, and special characters
	//  are not allowed for name.
	//  See ICoreWebView2Cookie for more details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#createcookie">See the ICoreWebView2CookieManager article.</a>
	CreateCookie(aName, aValue, aDomain, aPath string) ICoreWebView2Cookie // function
	// CopyCookie
	//  Creates a cookie whose params matches those of the specified cookie.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#copycookie">See the ICoreWebView2CookieManager article.</a>
	CopyCookie(aCookie ICoreWebView2Cookie) ICoreWebView2Cookie // function
	// GetCookies
	//  Gets a list of cookies matching the specific URI.
	//  If uri is empty string or null, all cookies under the same profile are
	//  returned.
	//  You can modify the cookie objects by calling
	//  ICoreWebView2CookieManager.AddOrUpdateCookie, and the changes
	//  will be applied to the webview.
	//  The TWVBrowserBase.OnGetCookiesCompleted event is triggered asynchronously with the cookies.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#getcookies">See the ICoreWebView2CookieManager article.</a>
	GetCookies(aURI string) bool // function
	// AddOrUpdateCookie
	//  Adds or updates a cookie with the given cookie data; may overwrite
	//  cookies with matching name, domain, and path if they exist.
	//  This method will fail if the domain of the given cookie is not specified.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#addorupdatecookie">See the ICoreWebView2CookieManager article.</a>
	AddOrUpdateCookie(aCookie ICoreWebView2Cookie) bool // function
	// DeleteCookie
	//  Deletes a cookie whose name and domain/path pair
	//  match those of the specified cookie.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deletecookie">See the ICoreWebView2CookieManager article.</a>
	DeleteCookie(aCookie ICoreWebView2Cookie) bool // function
	// DeleteCookies
	//  Deletes cookies with matching name and uri.
	//  Cookie name is required.
	//  All cookies with the given name where domain
	//  and path match provided URI are deleted.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deletecookies">See the ICoreWebView2CookieManager article.</a>
	DeleteCookies(aName, aURI string) bool // function
	// DeleteCookiesWithDomainAndPath
	//  Deletes cookies with matching name and domain/path pair.
	//  Cookie name is required.
	//  If domain is specified, deletes only cookies with the exact domain.
	//  If path is specified, deletes only cookies with the exact path.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deletecookieswithdomainandpath">See the ICoreWebView2CookieManager article.</a>
	DeleteCookiesWithDomainAndPath(aName, aDomain, aPath string) bool // function
	// DeleteAllCookies
	//  Deletes all cookies under the same profile.
	//  This could affect other WebViews under the same user profile.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#deleteallcookies">See the ICoreWebView2CookieManager article.</a>
	DeleteAllCookies() bool // function
	// SetBoundsAndZoomFactor
	//  Updates `Bounds` and `ZoomFactor` properties at the same time. This
	//  operation is atomic from the perspective of the host. After returning
	//  from this function, the `Bounds` and `ZoomFactor` properties are both
	//  updated if the function is successful, or neither is updated if the
	//  function fails. If `Bounds` and `ZoomFactor` are both updated by the
	//  same scale(for example, `Bounds` and `ZoomFactor` are both doubled),
	//  then the page does not display a change in `window.innerWidth` or
	//  `window.innerHeight` and the WebView renders the content at the new size
	//  and zoom without intermediate renderings. This function also updates
	//  just one of `ZoomFactor` or `Bounds` by passing in the new value for one
	//  and the current value for the other.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#setboundsandzoomfactor">See the ICoreWebView2Controller article.</a>
	SetBoundsAndZoomFactor(aBounds *TRect, aZoomFactor float64) bool // function
	// OpenDefaultDownloadDialog
	//  Open the default download dialog. If the dialog is opened before there
	//  are recent downloads, the dialog shows all past downloads for the
	//  current profile. Otherwise, the dialog shows only the recent downloads
	//  with a "See more" button for past downloads. Calling this method raises
	//  the `IsDefaultDownloadDialogOpenChanged` event if the dialog was closed.
	//  No effect if the dialog is already open.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#opendefaultdownloaddialog">See the ICoreWebView2_9 article.</a>
	OpenDefaultDownloadDialog() bool // function
	// CloseDefaultDownloadDialog
	//  Close the default download dialog. Calling this method raises the
	//  IsDefaultDownloadDialogOpenChanged event if the dialog was open.
	//  No effect if the dialog is already closed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#opendefaultdownloaddialog">See the ICoreWebView2_9 article.</a>
	CloseDefaultDownloadDialog() bool // function
	// SimulateEditingCommand
	//  Simulate editing commands using the "Input.dispatchKeyEvent" DevTools method.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</a>
	//  <a href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.</a>
	SimulateEditingCommand(aEditingCommand TWV2EditingCommand) bool // function
	// SimulateKeyEvent
	//  Dispatches a key event to the page using the "Input.dispatchKeyEvent"
	//  DevTools method. The browser has to be focused before simulating any
	//  key event. This function is asynchronous and it triggers the
	//  TWVBrowserBase.OnSimulateKeyEventCompleted event when it finishes executing.
	//  <param name="type_">Type of the key event.</param>
	//  <param name="modifiers">Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8.</param>
	//  <param name="windowsVirtualKeyCode">Windows virtual key code.</param>
	//  <param name="nativeVirtualKeyCode">Native virtual key code.</param>
	//  <param name="timestamp">Time at which the event occurred.</param>
	//  <param name="location">Whether the event was from the left or right side of the keyboard. 1=Left, 2=Right.</param>
	//  <param name="autoRepeat">Whether the event was generated from auto repeat.</param>
	//  <param name="isKeypad">Whether the event was generated from the keypad.</param>
	//  <param name="isSystemKey">Whether the event was a system key event.</param>
	//  <param name="text">Text as generated by processing a virtual key code with a keyboard layout. Not needed for for keyUp and rawKeyDown events.</param>
	//  <param name="unmodifiedtext">Text that would have been generated by the keyboard if no modifiers were pressed(except for shift). Useful for shortcut(accelerator) key handling.</param>
	//  <param name="keyIdentifier">Unique key identifier(e.g., 'U+0041').</param>
	//  <param name="code">Unique DOM defined string value for each physical key(e.g., 'KeyA').</param>
	//  <param name="key">Unique DOM defined string value describing the meaning of the key in the context of active modifiers, keyboard layout, etc(e.g., 'AltGr').</param>
	//  <a href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</a>
	SimulateKeyEvent(type_ TWV2KeyEventType, modifiers, windowsVirtualKeyCode, nativeVirtualKeyCode int32, timestamp int32, location int32, autoRepeat bool, isKeypad bool, isSystemKey bool, text string, unmodifiedtext string, keyIdentifier string, code string, key string) bool // function
	// KeyboardShortcutSearch
	//  Simulate that the F3 key was pressed and released.
	//  The browser has to be focused before simulating any key event.
	//  This key information was logged using a Spanish keyboard. It might not work with different keyboard layouts.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnSimulateKeyEventCompleted event several times.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</a>
	KeyboardShortcutSearch() bool // function
	// KeyboardShortcutRefreshIgnoreCache
	//  Simulate that SHIFT + F5 keys were pressed and released.
	//  The browser has to be focused before simulating any key event.
	//  This key information was logged using a Spanish keyboard. It might not work with different keyboard layouts.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnSimulateKeyEventCompleted event several times.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</a>
	KeyboardShortcutRefreshIgnoreCache() bool // function
	// SendMouseInput
	//  This function is only available in "Windowless mode" and provides mouse input meant for the WebView.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL or
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL, then mouseData specifies the amount of
	//  wheel movement. A positive value indicates that the wheel was rotated
	//  forward, away from the user; a negative value indicates that the wheel was
	//  rotated backward, toward the user. One wheel click is defined as
	//  WHEEL_DELTA, which is 120.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN, or
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP, then mouseData specifies which X
	//  buttons were pressed or released. This value should be 1 if the first X
	//  button is pressed/released and 2 if the second X button is
	//  pressed/released.
	//  If eventKind is COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE, then virtualKeys,
	//  mouseData, and point should all be zero.
	//  If eventKind is any other value, then mouseData should be zero.
	//  Point is expected to be in the client coordinate space of the WebView.
	//  To track mouse events that start in the WebView and can potentially move
	//  outside of the WebView and host application, calling SetCapture and
	//  ReleaseCapture is recommended.
	//  To dismiss hover popups, it is also recommended to send
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE messages.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#sendmouseinput">See the ICoreWebView2CompositionController article.</a>
	SendMouseInput(aEventKind TWVMouseEventKind, aVirtualKeys TWVMouseEventVirtualKeys, aMouseData uint32, aPoint *TPoint) bool // function
	// SendPointerInput
	//  This function is only available in "Windowless mode" and provides pointer input meant for the WebView.
	//  SendPointerInput accepts touch or pen pointer input of types defined in
	//  COREWEBVIEW2_POINTER_EVENT_KIND. Any pointer input from the system must be
	//  converted into an ICoreWebView2PointerInfo first.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#sendpointerinput">See the ICoreWebView2CompositionController article.</a>
	SendPointerInput(aEventKind TWVPointerEventKind, aPointerInfo ICoreWebView2PointerInfo) bool // function
	// DragEnter
	//  This function is only available in "Windowless mode" and corresponds to
	//  [IDropTarget::DragEnter](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragenter).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragEnter calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates(Similar to how SendMouseInput works).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3#dragenter">See the ICoreWebView2CompositionController3 article.</a>
	DragEnter(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 // function
	// DragLeave
	//  This function is only available in "Windowless mode" and corresponds to
	//  [IDropTarget::DragLeave](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragleave).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragLeave calls to this function.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3#dragleave">See the ICoreWebView2CompositionController3 article.</a>
	DragLeave() int32 // function
	// DragOver
	//  This function is only available in "Windowless mode" and corresponds to
	//  [IDropTarget::DragOver](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragover).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragOver calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates(Similar to how SendMouseInput works).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3#dragover">See the ICoreWebView2CompositionController3 article.</a>
	DragOver(keyState uint32, point *TPoint, OutEffect *uint32) int32 // function
	// Drop
	//  This function is only available in "Windowless mode" and corresponds to
	//  [IDropTarget::Drop](/windows/win32/api/oleidl/nf-oleidl-idroptarget-drop).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward Drop calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates(Similar to how SendMouseInput works).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller3#drop">See the ICoreWebView2CompositionController3 article.</a>
	Drop(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 // function
	// ClearCache
	//  Clears the browser cache. This function is asynchronous and it triggers the TWVBrowserBase.OnClearCacheCompleted event when it finishes executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache">See the Chrome DevTools Protocol page about the Network.clearBrowserCache method.</a>
	ClearCache() bool // function
	// ClearDataForOrigin
	//  Clears the storage for origin. This function is asynchronous and it triggers the TWVBrowserBase.OnClearDataForOriginCompleted event when it finishes executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin">See the Chrome DevTools Protocol page about the Storage.clearDataForOrigin method.</a>
	ClearDataForOrigin(aOrigin string, aStorageTypes TWVClearDataStorageTypes) bool // function
	// ClearBrowsingData
	//  Clear browsing data based on a data type. This method takes two parameters,
	//  the first being a mask of one or more `COREWEBVIEW2_BROWSING_DATA_KINDS`. OR
	//  operation(s) can be applied to multiple `COREWEBVIEW2_BROWSING_DATA_KINDS` to
	//  create a mask representing those data types.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnClearBrowsingDataCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdata">See the ICoreWebView2Profile2 article.</a>
	ClearBrowsingData(dataKinds TWVBrowsingDataKinds) bool // function
	// ClearBrowsingDataInTimeRange
	//  ClearBrowsingDataInTimeRange behaves like ClearBrowsingData except that it
	//  takes in two additional parameters for the start and end time for which it
	//  should clear the data between. The `startTime` and `endTime`
	//  parameters correspond to the number of seconds since the UNIX epoch.
	//  `startTime` is inclusive while `endTime` is exclusive, therefore the data will
	//  be cleared between startTime and endTime.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnClearBrowsingDataCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataintimerange">See the ICoreWebView2Profile2 article.</a>
	ClearBrowsingDataInTimeRange(dataKinds TWVBrowsingDataKinds, startTime, endTime TDateTime) bool // function
	// ClearBrowsingDataAll
	//  ClearBrowsingDataAll behaves like ClearBrowsingData except that it
	//  clears the entirety of the data associated with the profile it is called on.
	//  It clears the data regardless of timestamp.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnClearBrowsingDataCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataall">See the ICoreWebView2Profile2 article.</a>
	ClearBrowsingDataAll() bool // function
	// ClearServerCertificateErrorActions
	//  Clears all cached decisions to proceed with TLS certificate errors from the
	//  OnServerCertificateErrorDetected event for all WebView2's sharing the same session.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnServerCertificateErrorActionsCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_14#clearservercertificateerroractions">See the ICoreWebView2_14 article.</a>
	ClearServerCertificateErrorActions() bool // function
	// GetFavicon
	//  Async function for getting the actual image data of the favicon.
	//  This function is asynchronous and it triggers the TWVBrowserBase.OnGetFaviconCompleted event when it finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_15#getfavicon">See the ICoreWebView2_15 article.</a>
	GetFavicon(aFormat TWVFaviconImageFormat) bool // function
	// CreateSharedBuffer
	//  Create a shared memory based buffer with the specified size in bytes.
	//  The buffer can be shared with web contents in WebView by calling
	//  `PostSharedBufferToScript` on `CoreWebView2` or `CoreWebView2Frame` object.
	//  Once shared, the same content of the buffer will be accessible from both
	//  the app process and script in WebView. Modification to the content will be visible
	//  to all parties that have access to the buffer.
	//  The shared buffer is presented to the script as ArrayBuffer. All JavaScript APIs
	//  that work for ArrayBuffer including Atomics APIs can be used on it.
	//  There is currently a limitation that only size less than 2GB is supported.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment12#createsharedbuffer">See the ICoreWebView2Environment12 article.</a>
	CreateSharedBuffer(aSize int64, aSharedBuffer *ICoreWebView2SharedBuffer) bool // function
	// PostSharedBufferToScript
	//  Share a shared buffer object with script of the main frame in the WebView.
	//  The script will receive a `sharedbufferreceived` event from chrome.webview.
	//  Read the linked article for all the details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_17#postsharedbuffertoscript">See the ICoreWebView2_17 article.</a>
	PostSharedBufferToScript(aSharedBuffer ICoreWebView2SharedBuffer, aAccess TWVSharedBufferAccess, aAdditionalDataAsJson string) bool // function
	// GetProcessExtendedInfos
	//  Gets a snapshot collection of `ProcessExtendedInfo`s corresponding to all
	//  currently running processes associated with this `ICoreWebView2Environment`
	//  excludes crashpad process.
	//  This provides the same list of `ProcessInfo`s as what's provided in
	//  `GetProcessInfos`, but additionally provides a list of associated `FrameInfo`s
	//  which are actively running(showing or hiding UI elements) in the renderer
	//  process. See `AssociatedFrameInfos` for more information.
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment13#getprocessextendedinfos">See the ICoreWebView2Environment13 article.</a>
	//  This function triggers the TWVBrowserBase.OnGetProcessExtendedInfosCompleted event.
	GetProcessExtendedInfos() bool // function
	// MoveFormTo
	//  Move the parent form to the x and y coordinates.
	MoveFormTo(x, y int32) // procedure Is Abstract
	// MoveFormBy
	//  Move the parent form adding x and y to the coordinates.
	MoveFormBy(x, y int32) // procedure Is Abstract
	// ResizeFormWidthTo
	//  Add x to the parent form width.
	ResizeFormWidthTo(x int32) // procedure Is Abstract
	// ResizeFormHeightTo
	//  Add y to the parent form height.
	ResizeFormHeightTo(y int32) // procedure Is Abstract
	// SetFormLeftTo
	//  Set the parent form left property to x.
	SetFormLeftTo(x int32) // procedure Is Abstract
	// SetFormTopTo
	//  Set the parent form top property to y.
	SetFormTopTo(y int32) // procedure Is Abstract
	// IncZoomStep
	//  Increments the browser zoom.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	IncZoomStep() // procedure
	// DecZoomStep
	//  Decrements the browser zoom.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	DecZoomStep() // procedure
	// ResetZoom
	//  Sets the browser zoom to 100%.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#get_zoomfactor">See the ICoreWebView2Controller article.</a>
	ResetZoom() // procedure
	// ToggleMuteState
	//  Enables or disables all audio output from this browser.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_ismuted">See the ICoreWebView2_8 article.</a>
	ToggleMuteState() // procedure
	// SetOnBrowserProcessExited
	//  The OnBrowserProcessExited event is triggered when the collection of WebView2
	//  Runtime processes for the browser process of this environment terminate
	//  due to browser process failure or normal shutdown(for example, when all
	//  associated WebViews are closed), after all resources have been released
	// (including the user data folder).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment5#add_browserprocessexited">See the ICoreWebView2Environment5 article.</a>
	SetOnBrowserProcessExited(fn TOnBrowserProcessExitedEvent) // property event
	// SetOnProcessInfosChanged
	//  OnProcessInfosChanged is triggered when the ProcessInfos property has changed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8#add_processinfoschanged">See the ICoreWebView2Environment8 article.</a>
	SetOnProcessInfosChanged(fn TOnProcessInfosChangedEvent) // property event
	// SetOnContainsFullScreenElementChanged
	//  `OnContainsFullScreenElementChanged` triggers when the
	//  `ContainsFullScreenElement` property changes. An HTML element inside the
	//  WebView may enter fullscreen to the size of the WebView or leave
	//  fullscreen. This event is useful when, for example, a video element
	//  requests to go fullscreen. The listener of
	//  `ContainsFullScreenElementChanged` may resize the WebView in response.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_containsfullscreenelementchanged">See the ICoreWebView2 article.</a>
	SetOnContainsFullScreenElementChanged(fn TNotifyEvent) // property event
	// SetOnContentLoading
	//  `OnContentLoading` triggers before any content is loaded, including scripts added with
	//  `AddScriptToExecuteOnDocumentCreated`. `ContentLoading` does not trigger
	//  if a same page navigation occurs(such as through `fragment`
	//  navigations or `history.pushState` navigations). This operation
	//  follows the `NavigationStarting` and `SourceChanged` events and precedes
	//  the `HistoryChanged` and `NavigationCompleted` events.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_contentloading">See the ICoreWebView2 article.</a>
	SetOnContentLoading(fn TOnContentLoadingEvent) // property event
	// SetOnDocumentTitleChanged
	//  `OnDocumentTitleChanged` runs when the `DocumentTitle` property of the
	//  WebView changes and may run before or after the `NavigationCompleted`
	//  event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_documenttitlechanged">See the ICoreWebView2 article.</a>
	SetOnDocumentTitleChanged(fn TNotifyEvent) // property event
	// SetOnFrameNavigationCompleted
	//  `OnFrameNavigationCompleted` triggers when a child frame has completely
	//  loaded(concurrently when `body.onload` has triggered) or loading stopped with error.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_framenavigationcompleted">See the ICoreWebView2 article.</a>
	SetOnFrameNavigationCompleted(fn TOnNavigationCompletedEvent) // property event
	// SetOnFrameNavigationStarting
	//  `OnFrameNavigationStarting` triggers when a child frame in the WebView
	//  requests permission to navigate to a different URI. Redirects trigger
	//  this operation as well, and the navigation id is the same as the original
	//  one. Navigations will be blocked until all `FrameNavigationStarting` event
	//  handlers return.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_framenavigationstarting">See the ICoreWebView2 article.</a>
	SetOnFrameNavigationStarting(fn TOnNavigationStartingEvent) // property event
	// SetOnHistoryChanged
	//  `OnHistoryChanged` is raised for changes to joint session history, which consists of top-level
	//  and manual frame navigations. Use `HistoryChanged` to verify that the
	//  `CanGoBack` or `CanGoForward` value has changed. `HistoryChanged` also
	//  runs for using `GoBack` or `GoForward`. `HistoryChanged` runs after
	//  `SourceChanged` and `ContentLoading`. `CanGoBack` is false for
	//  navigations initiated through ICoreWebView2Frame APIs if there has not yet
	//  been a user gesture.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_historychanged">See the ICoreWebView2 article.</a>
	SetOnHistoryChanged(fn TNotifyEvent) // property event
	// SetOnNavigationCompleted
	//  `OnNavigationCompleted` runs when the WebView has completely loaded
	// (concurrently when `body.onload` runs) or loading stopped with error.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_navigationcompleted">See the ICoreWebView2 article.</a>
	SetOnNavigationCompleted(fn TOnNavigationCompletedEvent) // property event
	// SetOnNavigationStarting
	//  `OnNavigationStarting` runs when the WebView main frame is requesting
	//  permission to navigate to a different URI. Redirects trigger this
	//  operation as well, and the navigation id is the same as the original
	//  one. Navigations will be blocked until all `NavigationStarting` event handlers
	//  return.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_navigationstarting">See the ICoreWebView2 article.</a>
	SetOnNavigationStarting(fn TOnNavigationStartingEvent) // property event
	// SetOnNewWindowRequested
	//  `OnNewWindowRequested` runs when content inside the WebView requests to
	//  open a new window, such as through `window.open`. The app can pass a
	//  target WebView that is considered the opened window or mark the event as
	//  `Handled`, in which case WebView2 does not open a window.
	//  If either `Handled` or `NewWindow` properties are not set, the target
	//  content will be opened on a popup window.
	//  If a deferral is not taken on the event args, scripts that resulted in the
	//  new window that are requested are blocked until the event handler returns.
	//  If a deferral is taken, then scripts are blocked until the deferral is
	//  completed or new window is set.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_newwindowrequested">See the ICoreWebView2 article.</a>
	SetOnNewWindowRequested(fn TOnNewWindowRequestedEvent) // property event
	// SetOnPermissionRequested
	//  `OnPermissionRequested` runs when content in a WebView requests permission
	//  to access some privileged resources. If a deferral is not taken on the event
	//  args, the subsequent scripts are blocked until the event handler returns.
	//  If a deferral is taken, the scripts are blocked until the deferral is completed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_permissionrequested">See the ICoreWebView2 article.</a>
	SetOnPermissionRequested(fn TOnPermissionRequestedEvent) // property event
	// SetOnProcessFailed
	//  `OnProcessFailed` runs when any of the processes in the
	//  [WebView2 Process Group](https://learn.microsoft.com/microsoft-edge/webview2/concepts/process-model?tabs=csharp#processes-in-the-webview2-runtime)
	//  encounters one of the following conditions:
	//  <code>
	//  Condition | Details
	//  ---|---
	//  Unexpected exit | The process indicated by the event args has exited unexpectedly(usually due to a crash). The failure might or might not be recoverable and some failures are auto-recoverable.
	//  Unresponsiveness | The process indicated by the event args has become unresponsive to user input. This is only reported for renderer processes, and will run every few seconds until the process becomes responsive again.
	//  </code>
	//  NOTE: When the failing process is the browser process, a
	//  `ICoreWebView2Environment5.BrowserProcessExited` event will run too.
	//  Your application can use `ICoreWebView2ProcessFailedEventArgs` and
	//  `ICoreWebView2ProcessFailedEventArgs2` to identify which condition and
	//  process the event is for, and to collect diagnostics and handle recovery
	//  if necessary. For more details about which cases need to be handled by
	//  your application, see `COREWEBVIEW2_PROCESS_FAILED_KIND`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_processfailed">See the ICoreWebView2 article.</a>
	SetOnProcessFailed(fn TOnProcessFailedEvent) // property event
	// SetOnScriptDialogOpening
	//  `OnScriptDialogOpening` runs when a JavaScript dialog(`alert`, `confirm`,
	//  `prompt`, or `beforeunload`) displays for the webview. This event only
	//  triggers if the `ICoreWebView2Settings.AreDefaultScriptDialogsEnabled`
	//  property is set to `FALSE`. The `ScriptDialogOpening` event suppresses
	//  dialogs or replaces default dialogs with custom dialogs.
	//  If a deferral is not taken on the event args, the subsequent scripts are
	//  blocked until the event handler returns. If a deferral is taken, the
	//  scripts are blocked until the deferral is completed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_scriptdialogopening">See the ICoreWebView2 article.</a>
	SetOnScriptDialogOpening(fn TOnScriptDialogOpeningEvent) // property event
	// SetOnSourceChanged
	//  `OnSourceChanged` triggers when the `Source` property changes. `SourceChanged` runs when
	//  navigating to a different site or fragment navigations. It does not
	//  trigger for other types of navigations such as page refreshes or
	//  `history.pushState` with the same URL as the current page.
	//  `SourceChanged` runs before `ContentLoading` for navigation to a new
	//  document.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_sourcechanged">See the ICoreWebView2 article.</a>
	SetOnSourceChanged(fn TOnSourceChangedEvent) // property event
	// SetOnWebMessageReceived
	//  `OnWebMessageReceived` runs when the `ICoreWebView2Settings.IsWebMessageEnabled`
	//  setting is set and the top-level document of the WebView runs
	//  `window.chrome.webview.postMessage`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived">See the ICoreWebView2 article.</a>
	SetOnWebMessageReceived(fn TOnWebMessageReceivedEvent) // property event
	// SetOnWebResourceRequested
	//  `OnWebResourceRequested` runs when the WebView is performing a URL request
	//  to a matching URL and resource context filter that was added with
	//  `AddWebResourceRequestedFilter`. At least one filter must be added for
	//  the event to run.
	//  The web resource requested may be blocked until the event handler returns
	//  if a deferral is not taken on the event args. If a deferral is taken,
	//  then the web resource requested is blocked until the deferral is
	//  completed.
	//  If this event is subscribed in the add_NewWindowRequested handler it should be called
	//  after the new window is set. For more details see `ICoreWebView2NewWindowRequestedEventArgs.put_NewWindow`.
	//  This event is by default raised for file, http, and https URI schemes.
	//  This is also raised for registered custom URI schemes. For more details
	//  see `ICoreWebView2CustomSchemeRegistration`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_webresourcerequested">See the ICoreWebView2 article.</a>
	SetOnWebResourceRequested(fn TOnWebResourceRequestedEvent) // property event
	// SetOnWindowCloseRequested
	//  `OnWindowCloseRequested` triggers when content inside the WebView
	//  requested to close the window, such as after `window.close` is run. The
	//  app should close the WebView and related app window if that makes sense
	//  to the app.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#add_windowcloserequested">See the ICoreWebView2 article.</a>
	SetOnWindowCloseRequested(fn TNotifyEvent) // property event
	// SetOnDOMContentLoaded
	//  OnDOMContentLoaded is raised when the initial html document has been parsed.
	//  This aligns with the document's DOMContentLoaded event in html.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#add_domcontentloaded">See the ICoreWebView2_2 article.</a>
	SetOnDOMContentLoaded(fn TOnDOMContentLoadedEvent) // property event
	// SetOnWebResourceResponseReceived
	//  OnWebResourceResponseReceived is raised when the WebView receives the
	//  response for a request for a web resource(any URI resolution performed by
	//  the WebView; such as HTTP/HTTPS, file and data requests from redirects,
	//  navigations, declarations in HTML, implicit favicon lookups, and fetch API
	//  usage in the document). The host app can use this event to view the actual
	//  request and response for a web resource. There is no guarantee about the
	//  order in which the WebView processes the response and the host app's
	//  handler runs. The app's handler will not block the WebView from processing
	//  the response.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#add_webresourceresponsereceived">See the ICoreWebView2_2 article.</a>
	SetOnWebResourceResponseReceived(fn TOnWebResourceResponseReceivedEvent) // property event
	// SetOnDownloadStarting
	//  This event is raised when a download has begun, blocking the default download dialog,
	//  but not blocking the progress of the download.
	//  The host can choose to cancel a download, change the result file path,
	//  and hide the default download dialog.
	//  If the host chooses to cancel the download, the download is not saved, no
	//  dialog is shown, and the state is changed to
	//  COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED with interrupt reason
	//  COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED. Otherwise, the
	//  download is saved to the default path after the event completes,
	//  and default download dialog is shown if the host did not choose to hide it.
	//  The host can change the visibility of the download dialog using the
	//  `Handled` property. If the event is not handled, downloads complete
	//  normally with the default dialog shown.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_4#add_downloadstarting">See the ICoreWebView2_4 article.</a>
	SetOnDownloadStarting(fn TOnDownloadStartingEvent) // property event
	// SetOnFrameCreated
	//  Raised when a new iframe is created.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_4#add_framecreated">See the ICoreWebView2_4 article.</a>
	SetOnFrameCreated(fn TOnFrameCreatedEvent) // property event
	// SetOnClientCertificateRequested
	//  The OnClientCertificateRequested event is raised when the WebView2
	//  is making a request to an HTTP server that needs a client certificate
	//  for HTTP authentication.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_5#add_clientcertificaterequested">See the ICoreWebView2_5 article.</a>
	SetOnClientCertificateRequested(fn TOnClientCertificateRequestedEvent) // property event
	// SetOnIsDocumentPlayingAudioChanged
	//  `OnIsDocumentPlayingAudioChanged` is raised when the IsDocumentPlayingAudio property changes value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#add_isdocumentplayingaudiochanged">See the ICoreWebView2_8 article.</a>
	SetOnIsDocumentPlayingAudioChanged(fn TOnIsDocumentPlayingAudioChangedEvent) // property event
	// SetOnIsMutedChanged
	//  `OnIsMutedChanged` is raised when the IsMuted property changes value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#add_ismutedchanged">See the ICoreWebView2_8 article.</a>
	SetOnIsMutedChanged(fn TOnIsMutedChangedEvent) // property event
	// SetOnIsDefaultDownloadDialogOpenChanged
	//  Raised when the `IsDefaultDownloadDialogOpen` property changes. This event
	//  comes after the `DownloadStarting` event. Setting the `Handled` property
	//  on the `DownloadStartingEventArgs` disables the default download dialog
	//  and ensures that this event is never raised.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#add_isdefaultdownloaddialogopenchanged">See the ICoreWebView2_9 article.</a>
	SetOnIsDefaultDownloadDialogOpenChanged(fn TOnIsDefaultDownloadDialogOpenChangedEvent) // property event
	// SetOnBasicAuthenticationRequested
	//  Add an event handler for the BasicAuthenticationRequested event.
	//  BasicAuthenticationRequested event is raised when WebView encounters a
	//  Basic HTTP Authentication request as described in
	//  https://developer.mozilla.org/docs/Web/HTTP/Authentication, a Digest
	//  HTTP Authentication request as described in
	//  https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization#digest,
	//  an NTLM authentication or a Proxy Authentication request.
	//  The host can provide a response with credentials for the authentication or
	//  cancel the request. If the host sets the Cancel property to false but does not
	//  provide either UserName or Password properties on the Response property, then
	//  WebView2 will show the default authentication challenge dialog prompt to
	//  the user.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_10#add_basicauthenticationrequested">See the ICoreWebView2_10 article.</a>
	SetOnBasicAuthenticationRequested(fn TOnBasicAuthenticationRequestedEvent) // property event
	// SetOnContextMenuRequested
	//  `OnContextMenuRequested` event is raised when a context menu is requested by the user
	//  and the content inside WebView hasn't disabled context menus.
	//  The host has the option to create their own context menu with the information provided in
	//  the event or can add items to or remove items from WebView context menu.
	//  If the host doesn't handle the event, WebView will display the default context menu.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_11#add_contextmenurequested">See the ICoreWebView2_11 article.</a>
	SetOnContextMenuRequested(fn TOnContextMenuRequestedEvent) // property event
	// SetOnStatusBarTextChanged
	//  `OnStatusBarTextChanged` fires when the WebView is showing a status message,
	//  a URL, or an empty string(an indication to hide the status bar).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_12#add_statusbartextchanged">See the ICoreWebView2_12 article.</a>
	SetOnStatusBarTextChanged(fn TOnStatusBarTextChangedEvent) // property event
	// SetOnServerCertificateErrorActionsCompleted
	//  Event triggered when TWVBrowserBase.ClearServerCertificateErrorActions finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_14#clearservercertificateerroractions">See the ICoreWebView2_14 article.</a>
	SetOnServerCertificateErrorActionsCompleted(fn TOnServerCertificateErrorActionsCompletedEvent) // property event
	// SetOnServerCertificateErrorDetected
	//  The OnServerCertificateErrorDetected event is raised when the WebView2
	//  cannot verify server's digital certificate while loading a web page.
	//  This event will raise for all web resources and follows the `WebResourceRequested` event.
	//  If you don't handle the event, WebView2 will show the default TLS interstitial error page to the user
	//  for navigations, and for non-navigations the web request is cancelled.
	//  Note that WebView2 before raising `OnServerCertificateErrorDetected` raises a `OnNavigationCompleted` event
	//  with `IsSuccess` as FALSE and any of the below WebErrorStatuses that indicate a certificate failure.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_14#add_servercertificateerrordetected">See the ICoreWebView2_14 article.</a>
	SetOnServerCertificateErrorDetected(fn TOnServerCertificateErrorDetectedEvent) // property event
	// SetOnFaviconChanged
	//  The `OnFaviconChanged` event is raised when the
	//  [favicon](https://developer.mozilla.org/docs/Glossary/Favicon)
	//  had a different URL then the previous URL.
	//  The OnFaviconChanged event will be raised for first navigating to a new
	//  document, whether or not a document declares a Favicon in HTML if the
	//  favicon is different from the previous fav icon. The event will
	//  be raised again if a favicon is declared in its HTML or has script
	//  to set its favicon. The favicon information can then be retrieved with
	//  `GetFavicon` and `FaviconUri`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_15#add_faviconchanged">See the ICoreWebView2_15 article.</a>
	SetOnFaviconChanged(fn TOnFaviconChangedEvent) // property event
	// SetOnGetFaviconCompleted
	//  The TWVBrowserBase.OnGetFaviconCompleted event is triggered when the TWVBrowserBase.GetFavicon call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_15#getfavicon">See the ICoreWebView2_15 article.</a>
	SetOnGetFaviconCompleted(fn TOnGetFaviconCompletedEvent) // property event
	// SetOnPrintCompleted
	//  The TWVBrowserBase.OnPrintCompleted event is triggered when the TWVBrowserBase.Print call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#print">See the ICoreWebView2_16 article.</a>
	SetOnPrintCompleted(fn TOnPrintCompletedEvent) // property event
	// SetOnPrintToPdfStreamCompleted
	//  The TWVBrowserBase.OnPrintCompleted event is triggered when the TWVBrowserBase.PrintToPdfStream call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_16#printtopdfstream">See the ICoreWebView2_16 article.</a>
	SetOnPrintToPdfStreamCompleted(fn TOnPrintToPdfStreamCompletedEvent) // property event
	// SetOnAcceleratorKeyPressed
	//  `OnAcceleratorKeyPressed` runs when an accelerator key or key combo is
	//  pressed or released while the WebView is focused. A key is considered an
	//  accelerator if either of the following conditions are true.
	//  * Ctrl or Alt is currently being held.
	//  * The pressed key does not map to a character.
	//  A few specific keys are never considered accelerators, such as Shift.
	//  The `Escape` key is always considered an accelerator.
	//  Auto-repeated key events caused by holding the key down also triggers
	//  this event. Filter out the auto-repeated key events by verifying the
	//  `KeyEventLParam` or `PhysicalKeyStatus` event args.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#add_acceleratorkeypressed">See the ICoreWebView2Controller article.</a>
	SetOnAcceleratorKeyPressed(fn TOnAcceleratorKeyPressedEvent) // property event
	// SetOnGotFocus
	//  `OnGotFocus` runs when WebView has focus.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#add_gotfocus">See the ICoreWebView2Controller article.</a>
	SetOnGotFocus(fn TNotifyEvent) // property event
	// SetOnLostFocus
	//  `OnLostFocus` runs when WebView loses focus. In the case where `OnMoveFocusRequested` event is
	//  run, the focus is still on WebView when `OnMoveFocusRequested` event runs.
	//  `LostFocus` only runs afterwards when code of the app or default action
	//  of `OnMoveFocusRequested` event set focus away from WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#add_lostfocus">See the ICoreWebView2Controller article.</a>
	SetOnLostFocus(fn TNotifyEvent) // property event
	// SetOnMoveFocusRequested
	//  `OnMoveFocusRequested` runs when user tries to tab out of the WebView. The
	//  focus of the WebView has not changed when this event is run.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#add_movefocusrequested">See the ICoreWebView2Controller article.</a>
	SetOnMoveFocusRequested(fn TOnMoveFocusRequestedEvent) // property event
	// SetOnZoomFactorChanged
	//  `OnZoomFactorChanged` runs when the `ZoomFactor` property of the WebView
	//  changes. The event may run because the `ZoomFactor` property was
	//  modified, or due to the user manually modifying the zoom. When it is
	//  modified using the `ZoomFactor` property, the internal zoom factor is
	//  updated immediately and no `OnZoomFactorChanged` event is triggered.
	//  WebView associates the last used zoom factor for each site. It is
	//  possible for the zoom factor to change when navigating to a different
	//  page. When the zoom factor changes due to a navigation change, the
	//  `OnZoomFactorChanged` event runs right after the `ContentLoading` event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller#add_zoomfactorchanged">See the ICoreWebView2Controller article.</a>
	SetOnZoomFactorChanged(fn TNotifyEvent) // property event
	// SetOnRasterizationScaleChanged
	//  The event is raised when the WebView detects that the monitor DPI scale
	//  has changed, ShouldDetectMonitorScaleChanges is true, and the WebView has
	//  changed the RasterizationScale property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2controller3#add_rasterizationscalechanged">See the ICoreWebView2Controller3 article.</a>
	SetOnRasterizationScaleChanged(fn TNotifyEvent) // property event
	// SetOnCursorChanged
	//  The event is raised when WebView thinks the cursor should be changed. For
	//  example, when the mouse cursor is currently the default cursor but is then
	//  moved over text, it may try to change to the IBeam cursor.
	//  It is expected for the developer to send
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE messages(in addition to
	//  COREWEBVIEW2_MOUSE_EVENT_KIND_MOVE messages) through the SendMouseInput
	//  API. This is to ensure that the mouse is actually within the WebView that
	//  sends out CursorChanged events.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#add_cursorchanged">See the ICoreWebView2CompositionController article.</a>
	SetOnCursorChanged(fn TNotifyEvent) // property event
	// SetOnBytesReceivedChanged
	//  The event is raised when the number of received bytes for a download operation changes.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#add_bytesreceivedchanged">See the ICoreWebView2DownloadOperation article.</a>
	SetOnBytesReceivedChanged(fn TOnBytesReceivedChangedEvent) // property event
	// SetOnEstimatedEndTimeChanged
	//  The event is raised when the estimated end time for a download operation changes.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#add_estimatedendtimechanged">See the ICoreWebView2DownloadOperation article.</a>
	SetOnEstimatedEndTimeChanged(fn TOnEstimatedEndTimeChangedEvent) // property event
	// SetOnDownloadStateChanged
	//  The event is raised when the download operation state changes.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2downloadoperation#add_statechanged">See the ICoreWebView2DownloadOperation article.</a>
	SetOnDownloadStateChanged(fn TOnDownloadStateChangedEvent) // property event
	// SetOnFrameDestroyed
	//  The OnFrameDestroyed event is raised when the iframe corresponding
	//  to this CoreWebView2Frame object is removed or the document
	//  containing that iframe is destroyed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#add_destroyed">See the ICoreWebView2Frame article.</a>
	SetOnFrameDestroyed(fn TOnFrameDestroyedEvent) // property event
	// SetOnFrameNameChanged
	//  Raised when the iframe changes its window.name property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame#add_namechanged">See the ICoreWebView2Frame article.</a>
	SetOnFrameNameChanged(fn TOnFrameNameChangedEvent) // property event
	// SetOnFrameNavigationStarting2
	//  A frame navigation will raise a `OnFrameNavigationStarting2` event and
	//  a `OnFrameNavigationStarting` event. All of the
	//  `FrameNavigationStarting` event handlers for the current frame will be
	//  run before the `OnFrameNavigationStarting2` event handlers. All of the event handlers
	//  share a common `NavigationStartingEventArgs` object. Whichever event handler is
	//  last to change the `NavigationStartingEventArgs.Cancel` property will
	//  decide if the frame navigation will be cancelled. Redirects raise this
	//  event as well, and the navigation id is the same as the original one.
	//  Navigations will be blocked until all `OnFrameNavigationStarting2` and
	//  `OnFrameNavigationStarting` event handlers return.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#add_navigationstarting">See the ICoreWebView2Frame2 article.</a>
	SetOnFrameNavigationStarting2(fn TOnFrameNavigationStartingEvent) // property event
	// SetOnFrameNavigationCompleted2
	//  `OnFrameNavigationCompleted2` runs when the CoreWebView2Frame has completely
	//  loaded(concurrently when `body.onload` runs) or loading stopped with error.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#add_navigationcompleted">See the ICoreWebView2Frame2 article.</a>
	SetOnFrameNavigationCompleted2(fn TOnFrameNavigationCompletedEvent) // property event
	// SetOnFrameContentLoading
	//  `OnFrameContentLoading` triggers before any content is loaded, including scripts added with
	//  `AddScriptToExecuteOnDocumentCreated`. `OnFrameContentLoading` does not trigger
	//  if a same page navigation occurs(such as through `fragment`
	//  navigations or `history.pushState` navigations). This operation
	//  follows the `OnFrameNavigationStarting2` and precedes `OnFrameNavigationCompleted2` events.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#add_contentloading">See the ICoreWebView2Frame2 article.</a>
	SetOnFrameContentLoading(fn TOnFrameContentLoadingEvent) // property event
	// SetOnFrameDOMContentLoaded
	//  OnFrameDOMContentLoaded is raised when the iframe html document has been parsed.
	//  This aligns with the document's DOMContentLoaded event in html.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#add_domcontentloaded">See the ICoreWebView2Frame2 article.</a>
	SetOnFrameDOMContentLoaded(fn TOnFrameDOMContentLoadedEvent) // property event
	// SetOnFrameWebMessageReceived
	//  `OnFrameWebMessageReceived` runs when the
	//  `ICoreWebView2Settings.IsWebMessageEnabled` setting is set and the
	//  frame document runs `window.chrome.webview.postMessage`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame2#add_webmessagereceived">See the ICoreWebView2Frame2 article.</a>
	SetOnFrameWebMessageReceived(fn TOnFrameWebMessageReceivedEvent) // property event
	// SetOnFramePermissionRequested
	//  `OnFramePermissionRequested` is raised when content in an iframe any of its
	//  descendant iframes requests permission to privileged resources.
	//  This relates to the `OnPermissionRequested` event on the `CoreWebView2`.
	//  Both these events will be raised in the case of an iframe requesting
	//  permission. The `CoreWebView2Frame`'s event handlers will be invoked
	//  before the event handlers on the `CoreWebView2`. If the `Handled` property
	//  of the `PermissionRequestedEventArgs` is set to TRUE within the
	//  `CoreWebView2Frame` event handler, then the event will not be
	//  raised on the `CoreWebView2`, and it's event handlers will not be invoked.
	//  In the case of nested iframes, the 'OnFramePermissionRequested' event will
	//  be raised from the top level iframe.
	//  If a deferral is not taken on the event args, the subsequent scripts are
	//  blocked until the event handler returns. If a deferral is taken, the
	//  scripts are blocked until the deferral is completed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2frame3#add_permissionrequested">See the ICoreWebView2Frame3 article.</a>
	SetOnFramePermissionRequested(fn TOnFramePermissionRequestedEvent) // property event
	// SetOnDevToolsProtocolEventReceived
	//  OnDevToolsProtocolEventReceived is triggered when a DevTools protocol
	//  event runs. It's necessary to subscribe to that event with a
	//  SubscribeToDevToolsProtocolEvent call.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2devtoolsprotocoleventreceiver#add_devtoolsprotocoleventreceived">See the ICoreWebView2DevToolsProtocolEventReceiver article.</a>
	SetOnDevToolsProtocolEventReceived(fn TOnDevToolsProtocolEventReceivedEvent) // property event
	// SetOnCustomItemSelected
	//  `OnCustomItemSelected` event is raised when the user selects a custom `ContextMenuItem`.
	//  Will only be raised for end developer created context menu items.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2contextmenuitem#add_customitemselected">See the ICoreWebView2ContextMenuItem article.</a>
	SetOnCustomItemSelected(fn TOnCustomItemSelectedEvent) // property event
	// SetOnClearBrowsingDataCompleted
	//  This event is triggered when the TWVBrowserBase.ClearBrowsingData call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdata">See the ICoreWebView2Profile2 article.</a>
	SetOnClearBrowsingDataCompleted(fn TOnClearBrowsingDataCompletedEvent) // property event
	// SetOnInitializationError
	//  Called if any of the browser initialization steps fail.
	SetOnInitializationError(fn TOnInitializationErrorEvent) // property event
	// SetOnEnvironmentCompleted
	//  Called when the environment was created successfully.
	SetOnEnvironmentCompleted(fn TNotifyEvent) // property event
	// SetOnControllerCompleted
	//  Called when the controller was created successfully.
	SetOnControllerCompleted(fn TNotifyEvent) // property event
	// SetOnAfterCreated
	//  Called after a new browser is created and it's ready to navigate to the default URL.
	SetOnAfterCreated(fn TNotifyEvent) // property event
	// SetOnExecuteScriptCompleted
	//  Triggered when a TWVBrowserBase.ExecuteScript call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#executescript">See the ICoreWebView2 article.</a>
	SetOnExecuteScriptCompleted(fn TOnExecuteScriptCompletedEvent) // property event
	// SetOnCapturePreviewCompleted
	//  Triggered when a TWVBrowserBase.CapturePreview call finishes writting the image to the stream.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#capturepreview">See the ICoreWebView2 article.</a>
	SetOnCapturePreviewCompleted(fn TOnCapturePreviewCompletedEvent) // property event
	// SetOnGetCookiesCompleted
	//  Triggered when a TWVBrowserBase.GetCookies call finishes executing. This event includes the requested cookies.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2cookiemanager#getcookies">See the ICoreWebView2CookieManager article.</a>
	SetOnGetCookiesCompleted(fn TOnGetCookiesCompletedEvent) // property event
	// SetOnTrySuspendCompleted
	//  The TWVBrowserBase.OnTrySuspendCompleted event is triggered when a TWVBrowserBase.TrySuspend call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#trysuspend">See the ICoreWebView2_3 article.</a>
	SetOnTrySuspendCompleted(fn TOnTrySuspendCompletedEvent) // property event
	// SetOnPrintToPdfCompleted
	//  The TWVBrowserBase.OnPrintToPdfCompleted event is triggered when the TWVBrowserBase.PrintToPdf call finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_7#printtopdf">See the ICoreWebView2_7 article.</a>
	SetOnPrintToPdfCompleted(fn TOnPrintToPdfCompletedEvent) // property event
	// SetOnCompositionControllerCompleted
	//  Called when the composition controller was created successfully.
	SetOnCompositionControllerCompleted(fn TNotifyEvent) // property event
	// SetOnCallDevToolsProtocolMethodCompleted
	//  The TWVBrowserBase.OnCallDevToolsProtocolMethodCompleted event is triggered
	//  when TWVBrowserBase.CallDevToolsProtocolMethod finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#calldevtoolsprotocolmethod">See the ICoreWebView2 article.</a>
	SetOnCallDevToolsProtocolMethodCompleted(fn TOnCallDevToolsProtocolMethodCompletedEvent) // property event
	// SetOnAddScriptToExecuteOnDocumentCreatedCompleted
	//  The TWVBrowserBase.OnAddScriptToExecuteOnDocumentCreatedCompleted event is triggered
	//  when TWVBrowserBase.AddScriptToExecuteOnDocumentCreated finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addscripttoexecuteondocumentcreated">See the ICoreWebView2 article.</a>
	SetOnAddScriptToExecuteOnDocumentCreatedCompleted(fn TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent) // property event
	// SetOnWebResourceResponseViewGetContentCompleted
	//  The TWVBrowserBase.OnWebResourceResponseViewGetContentCompleted event is triggered
	//  when TCoreWebView2WebResourceResponseView.GetContent finishes executing. This event includes the resource contents.
	//  See the MiniBrowser demo for an example.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2webresourceresponseview#getcontent">See the ICoreWebView2WebResourceResponseView article.</a>
	SetOnWebResourceResponseViewGetContentCompleted(fn TOnWebResourceResponseViewGetContentCompletedEvent) // property event
	// SetOnWidget0CompMsg
	//  Event triggered when the window called 'Chrome_WidgetWin_0' receives a message.
	SetOnWidget0CompMsg(fn TOnCompMsgEvent) // property event
	// SetOnWidget1CompMsg
	//  Event triggered when the window called 'Chrome_WidgetWin_1' receives a message.
	SetOnWidget1CompMsg(fn TOnCompMsgEvent) // property event
	// SetOnRenderCompMsg
	//  Event triggered when the window called 'Chrome_RenderWidgetHostHWND' receives a message.
	SetOnRenderCompMsg(fn TOnCompMsgEvent) // property event
	// SetOnD3DWindowCompMsg
	//  Event triggered when the window called 'Intermediate D3D Window' receives a message.
	SetOnD3DWindowCompMsg(fn TOnCompMsgEvent) // property event
	// SetOnRetrieveHTMLCompleted
	//  The TWVBrowserBase.OnRetrieveHTMLCompleted event is triggered when TWVBrowserBase.RetrieveHTML finishes executing. It includes the HTML contents.
	SetOnRetrieveHTMLCompleted(fn TOnRetrieveHTMLCompletedEvent) // property event
	// SetOnRetrieveTextCompleted
	//  The TWVBrowserBase.OnRetrieveTextCompleted event is triggered when TWVBrowserBase.RetrieveText finishes executing. It includes the text contents.
	SetOnRetrieveTextCompleted(fn TOnRetrieveTextCompletedEvent) // property event
	// SetOnRetrieveMHTMLCompleted
	//  The TWVBrowserBase.OnRetrieveMHTMLCompleted event is triggered when TWVBrowserBase.RetrieveMHTML finishes executing. It includes the MHTML contents.
	SetOnRetrieveMHTMLCompleted(fn TOnRetrieveMHTMLCompletedEvent) // property event
	// SetOnClearCacheCompleted
	//  The TWVBrowserBase.OnClearCacheCompleted event is triggered when TWVBrowserBase.ClearCache finishes executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-clearBrowserCache">See the Chrome DevTools Protocol page about the Network.clearBrowserCache method.</a>
	SetOnClearCacheCompleted(fn TOnClearCacheCompletedEvent) // property event
	// SetOnClearDataForOriginCompleted
	//  The TWVBrowserBase.OnClearDataForOriginCompleted event is triggered when TWVBrowserBase.ClearDataForOrigin finishes executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Storage/#method-clearDataForOrigin">See the Chrome DevTools Protocol page about the Storage.clearDataForOrigin method.</a>
	SetOnClearDataForOriginCompleted(fn TOnClearDataForOriginCompletedEvent) // property event
	// SetOnOfflineCompleted
	//  The TWVBrowserBase.OnOfflineCompleted event is triggered after setting the TWVBrowserBase.Offline property.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Network/#method-emulateNetworkConditions">See the Network Domain article.</a>
	SetOnOfflineCompleted(fn TOnOfflineCompletedEvent) // property event
	// SetOnIgnoreCertificateErrorsCompleted
	//  The TWVBrowserBase.OnIgnoreCertificateErrorsCompleted event is triggered after setting the TWVBrowserBase.IgnoreCertificateErrors property.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Security/#method-setIgnoreCertificateErrors">See the Security Domain article.</a>
	SetOnIgnoreCertificateErrorsCompleted(fn TOnIgnoreCertificateErrorsCompletedEvent) // property event
	// SetOnRefreshIgnoreCacheCompleted
	//  The TWVBrowserBase.OnRefreshIgnoreCacheCompleted event is triggered when TWVBrowserBase.RefreshIgnoreCache finishes executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-reload">See the Page Domain article.</a>
	SetOnRefreshIgnoreCacheCompleted(fn TOnRefreshIgnoreCacheCompletedEvent) // property event
	// SetOnSimulateKeyEventCompleted
	//  The TWVBrowserBase.OnSimulateKeyEventCompleted event is triggered when TWVBrowserBase.SimulateKeyEvent or TWVBrowserBase.SimulateEditingCommand finish executing.
	//  <a href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.</a>
	SetOnSimulateKeyEventCompleted(fn TOnSimulateKeyEventCompletedEvent) // property event
	// SetOnGetCustomSchemes
	//  OnGetCustomSchemes is triggered automatically before creaing the environment to register custom schemes.
	//  Fill the aCustomSchemes event parameter with all the information to create one or more
	//  ICoreWebView2CustomSchemeRegistration instances that will be used during the creation of the Environment.
	//  <a cref="uWVTypes|TWVCustomSchemeInfo">See TWVCustomSchemeInfo.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2customschemeregistration">See the ICoreWebView2CustomSchemeRegistration article.</a>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environmentoptions4">See the ICoreWebView2EnvironmentOptions4 article.</a>
	SetOnGetCustomSchemes(fn TOnGetCustomSchemesEvent) // property event
	// SetOnGetNonDefaultPermissionSettingsCompleted
	//  The TWVBrowserBase.OnGetNonDefaultPermissionSettingsCompleted event is triggered when TWVBrowserBase.GetNonDefaultPermissionSettings finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#getnondefaultpermissionsettings">See the ICoreWebView2Profile4 article.</a>
	SetOnGetNonDefaultPermissionSettingsCompleted(fn TOnGetNonDefaultPermissionSettingsCompletedEvent) // property event
	// SetOnSetPermissionStateCompleted
	//  The TWVBrowserBase.OnSetPermissionStateCompleted event is triggered when TWVBrowserBase.SetPermissionState finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#setpermissionstate">See the ICoreWebView2Profile4 article.</a>
	SetOnSetPermissionStateCompleted(fn TOnSetPermissionStateCompletedEvent) // property event
	// SetOnLaunchingExternalUriScheme
	//  The `OnLaunchingExternalUriScheme` event is raised when a navigation request is made to
	//  a URI scheme that is registered with the OS.
	//  The `OnLaunchingExternalUriScheme` event handler may suppress the default dialog
	//  or replace the default dialog with a custom dialog.
	//  If a deferral is not taken on the event args, the external URI scheme launch is
	//  blocked until the event handler returns. If a deferral is taken, the
	//  external URI scheme launch is blocked until the deferral is completed.
	//  The host also has the option to cancel the URI scheme launch.
	//  The `NavigationStarting` and `NavigationCompleted` events will be raised,
	//  regardless of whether the `Cancel` property is set to `TRUE` or
	//  `FALSE`. The `NavigationCompleted` event will be raised with the `IsSuccess` property
	//  set to `FALSE` and the `WebErrorStatus` property set to `ConnectionAborted` regardless of
	//  whether the host sets the `Cancel` property on the
	//  `ICoreWebView2LaunchingExternalUriSchemeEventArgs`. The `SourceChanged`, `ContentLoading`,
	//  and `HistoryChanged` events will not be raised for this navigation to the external URI
	//  scheme regardless of the `Cancel` property.
	//  The `OnLaunchingExternalUriScheme` event will be raised after the
	//  `NavigationStarting` event and before the `NavigationCompleted` event.
	//  The default `CoreWebView2Settings` will also be updated upon navigation to an external
	//  URI scheme. If a setting on the `CoreWebView2Settings` interface has been changed,
	//  navigating to an external URI scheme will trigger the `CoreWebView2Settings` to update.
	//  The WebView2 may not display the default dialog based on user settings, browser settings,
	//  and whether the origin is determined as a
	//  [trustworthy origin](https://w3c.github.io/webappsec-secure-contexts#
	//  potentially-trustworthy-origin); however, the event will still be raised.
	//  If the request is initiated by a cross-origin frame without a user gesture,
	//  the request will be blocked and the `OnLaunchingExternalUriScheme` event will not
	//  be raised.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_18#add_launchingexternalurischeme">See the ICoreWebView2_18 article.</a>
	SetOnLaunchingExternalUriScheme(fn TOnLaunchingExternalUriSchemeEvent) // property event
	// SetOnGetProcessExtendedInfosCompleted
	//  The TWVBrowserBase.OnGetProcessExtendedInfosCompleted event is triggered when TWVBrowserBase.GetProcessExtendedInfos finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment13#getprocessextendedinfos">See the ICoreWebView2Environment13 article.</a>
	SetOnGetProcessExtendedInfosCompleted(fn TOnGetProcessExtendedInfosCompletedEvent) // property event
	// SetOnBrowserExtensionRemoveCompleted
	//  The TWVBrowserBase.OnBrowserExtensionRemoveCompleted event is triggered when TCoreWebView2BrowserExtension.Remove finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension#remove">See the ICoreWebView2BrowserExtension article.</a>
	SetOnBrowserExtensionRemoveCompleted(fn TOnBrowserExtensionRemoveCompletedEvent) // property event
	// SetOnBrowserExtensionEnableCompleted
	//  The TWVBrowserBase.OnBrowserExtensionEnableCompleted event is triggered when TCoreWebView2BrowserExtension.Enable finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2browserextension#enable">See the ICoreWebView2BrowserExtension article.</a>
	SetOnBrowserExtensionEnableCompleted(fn TOnBrowserExtensionEnableCompletedEvent) // property event
	// SetOnProfileAddBrowserExtensionCompleted
	//  The TWVBrowserBase.OnProfileAddBrowserExtensionCompleted event is triggered when TWVBrowserBase.AddBrowserExtension finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#addbrowserextension">See the ICoreWebView2Profile7 article.</a>
	SetOnProfileAddBrowserExtensionCompleted(fn TOnProfileAddBrowserExtensionCompletedEvent) // property event
	// SetOnProfileGetBrowserExtensionsCompleted
	//  The TWVBrowserBase.OnProfileGetBrowserExtensionsCompleted event is triggered when TWVBrowserBase.GetBrowserExtensions finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#getbrowserextensions">See the ICoreWebView2Profile7 article.</a>
	SetOnProfileGetBrowserExtensionsCompleted(fn TOnProfileGetBrowserExtensionsCompletedEvent) // property event
	// SetOnProfileDeleted
	//  The TWVBrowserBase.OnProfileDeleted event is triggered when TWVBrowserBase.Delete finishes executing.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile8#delete">See the ICoreWebView2Profile8 article.</a>
	SetOnProfileDeleted(fn TOnProfileDeletedEvent) // property event
	// SetOnExecuteScriptWithResultCompleted
	//  Provides the result of ExecuteScriptWithResult.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_21#executescriptwithresult">See the ICoreWebView2_21 article.</a>
	SetOnExecuteScriptWithResultCompleted(fn TOnExecuteScriptWithResultCompletedEvent) // property event
}

// TWVBrowserBase Is Abstract Class Parent: TComponent
//
//	Parent class of TWVBrowser and TWVFMXBrowser that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type TWVBrowserBase struct {
	TComponent
	browserProcessExitedPtr                         uintptr
	processInfosChangedPtr                          uintptr
	containsFullScreenElementChangedPtr             uintptr
	contentLoadingPtr                               uintptr
	documentTitleChangedPtr                         uintptr
	frameNavigationCompletedPtr                     uintptr
	frameNavigationStartingPtr                      uintptr
	historyChangedPtr                               uintptr
	navigationCompletedPtr                          uintptr
	navigationStartingPtr                           uintptr
	newWindowRequestedPtr                           uintptr
	permissionRequestedPtr                          uintptr
	processFailedPtr                                uintptr
	scriptDialogOpeningPtr                          uintptr
	sourceChangedPtr                                uintptr
	webMessageReceivedPtr                           uintptr
	webResourceRequestedPtr                         uintptr
	windowCloseRequestedPtr                         uintptr
	dOMContentLoadedPtr                             uintptr
	webResourceResponseReceivedPtr                  uintptr
	downloadStartingPtr                             uintptr
	frameCreatedPtr                                 uintptr
	clientCertificateRequestedPtr                   uintptr
	isDocumentPlayingAudioChangedPtr                uintptr
	isMutedChangedPtr                               uintptr
	isDefaultDownloadDialogOpenChangedPtr           uintptr
	basicAuthenticationRequestedPtr                 uintptr
	contextMenuRequestedPtr                         uintptr
	statusBarTextChangedPtr                         uintptr
	serverCertificateErrorActionsCompletedPtr       uintptr
	serverCertificateErrorDetectedPtr               uintptr
	faviconChangedPtr                               uintptr
	getFaviconCompletedPtr                          uintptr
	printCompletedPtr                               uintptr
	printToPdfStreamCompletedPtr                    uintptr
	acceleratorKeyPressedPtr                        uintptr
	gotFocusPtr                                     uintptr
	lostFocusPtr                                    uintptr
	moveFocusRequestedPtr                           uintptr
	zoomFactorChangedPtr                            uintptr
	rasterizationScaleChangedPtr                    uintptr
	cursorChangedPtr                                uintptr
	bytesReceivedChangedPtr                         uintptr
	estimatedEndTimeChangedPtr                      uintptr
	downloadStateChangedPtr                         uintptr
	frameDestroyedPtr                               uintptr
	frameNameChangedPtr                             uintptr
	frameNavigationStarting2Ptr                     uintptr
	frameNavigationCompleted2Ptr                    uintptr
	frameContentLoadingPtr                          uintptr
	frameDOMContentLoadedPtr                        uintptr
	frameWebMessageReceivedPtr                      uintptr
	framePermissionRequestedPtr                     uintptr
	devToolsProtocolEventReceivedPtr                uintptr
	customItemSelectedPtr                           uintptr
	clearBrowsingDataCompletedPtr                   uintptr
	initializationErrorPtr                          uintptr
	environmentCompletedPtr                         uintptr
	controllerCompletedPtr                          uintptr
	afterCreatedPtr                                 uintptr
	executeScriptCompletedPtr                       uintptr
	capturePreviewCompletedPtr                      uintptr
	getCookiesCompletedPtr                          uintptr
	trySuspendCompletedPtr                          uintptr
	printToPdfCompletedPtr                          uintptr
	compositionControllerCompletedPtr               uintptr
	callDevToolsProtocolMethodCompletedPtr          uintptr
	addScriptToExecuteOnDocumentCreatedCompletedPtr uintptr
	webResourceResponseViewGetContentCompletedPtr   uintptr
	widget0CompMsgPtr                               uintptr
	widget1CompMsgPtr                               uintptr
	renderCompMsgPtr                                uintptr
	d3DWindowCompMsgPtr                             uintptr
	retrieveHTMLCompletedPtr                        uintptr
	retrieveTextCompletedPtr                        uintptr
	retrieveMHTMLCompletedPtr                       uintptr
	clearCacheCompletedPtr                          uintptr
	clearDataForOriginCompletedPtr                  uintptr
	offlineCompletedPtr                             uintptr
	ignoreCertificateErrorsCompletedPtr             uintptr
	refreshIgnoreCacheCompletedPtr                  uintptr
	simulateKeyEventCompletedPtr                    uintptr
	getCustomSchemesPtr                             uintptr
	getNonDefaultPermissionSettingsCompletedPtr     uintptr
	setPermissionStateCompletedPtr                  uintptr
	launchingExternalUriSchemePtr                   uintptr
	getProcessExtendedInfosCompletedPtr             uintptr
	browserExtensionRemoveCompletedPtr              uintptr
	browserExtensionEnableCompletedPtr              uintptr
	profileAddBrowserExtensionCompletedPtr          uintptr
	profileGetBrowserExtensionsCompletedPtr         uintptr
	profileDeletedPtr                               uintptr
	executeScriptWithResultCompletedPtr             uintptr
}

func (m *TWVBrowserBase) Initialized() bool {
	r1 := WV().SysCallN(866, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) CoreWebView2PrintSettings() ICoreWebView2PrintSettings {
	var resultCoreWebView2PrintSettings uintptr
	WV().SysCallN(818, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PrintSettings)))
	return AsCoreWebView2PrintSettings(resultCoreWebView2PrintSettings)
}

func (m *TWVBrowserBase) CoreWebView2Settings() ICoreWebView2Settings {
	var resultCoreWebView2Settings uintptr
	WV().SysCallN(819, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Settings)))
	return AsCoreWebView2Settings(resultCoreWebView2Settings)
}

func (m *TWVBrowserBase) CoreWebView2Environment() ICoreWebView2Environment {
	var resultCoreWebView2Environment uintptr
	WV().SysCallN(817, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Environment)))
	return AsCoreWebView2Environment(resultCoreWebView2Environment)
}

func (m *TWVBrowserBase) CoreWebView2Controller() ICoreWebView2Controller {
	var resultCoreWebView2Controller uintptr
	WV().SysCallN(816, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Controller)))
	return AsCoreWebView2Controller(resultCoreWebView2Controller)
}

func (m *TWVBrowserBase) CoreWebView2CompositionController() ICoreWebView2CompositionController {
	var resultCoreWebView2CompositionController uintptr
	WV().SysCallN(815, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CompositionController)))
	return AsCoreWebView2CompositionController(resultCoreWebView2CompositionController)
}

func (m *TWVBrowserBase) CoreWebView2() ICoreWebView2 {
	var resultCoreWebView2 uintptr
	WV().SysCallN(814, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2)))
	return AsCoreWebView2(resultCoreWebView2)
}

func (m *TWVBrowserBase) Profile() ICoreWebView2Profile {
	var resultCoreWebView2Profile uintptr
	WV().SysCallN(903, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Profile)))
	return AsCoreWebView2Profile(resultCoreWebView2Profile)
}

func (m *TWVBrowserBase) DefaultURL() string {
	r1 := WV().SysCallN(836, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetDefaultURL(AValue string) {
	WV().SysCallN(836, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) IsNavigating() bool {
	r1 := WV().SysCallN(872, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) ZoomPct() (resultDouble float64) {
	WV().SysCallN(1048, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TWVBrowserBase) SetZoomPct(AValue float64) {
	WV().SysCallN(1048, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TWVBrowserBase) ZoomStep() byte {
	r1 := WV().SysCallN(1049, 0, m.Instance(), 0)
	return byte(r1)
}

func (m *TWVBrowserBase) SetZoomStep(AValue byte) {
	WV().SysCallN(1049, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) Widget0CompHWND() THandle {
	r1 := WV().SysCallN(1044, m.Instance())
	return THandle(r1)
}

func (m *TWVBrowserBase) Widget1CompHWND() THandle {
	r1 := WV().SysCallN(1045, m.Instance())
	return THandle(r1)
}

func (m *TWVBrowserBase) RenderCompHWND() THandle {
	r1 := WV().SysCallN(915, m.Instance())
	return THandle(r1)
}

func (m *TWVBrowserBase) D3DWindowCompHWND() THandle {
	r1 := WV().SysCallN(828, m.Instance())
	return THandle(r1)
}

func (m *TWVBrowserBase) ScreenScale() float32 {
	r1 := WV().SysCallN(924, m.Instance())
	return float32(r1)
}

func (m *TWVBrowserBase) Offline() bool {
	r1 := WV().SysCallN(889, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetOffline(AValue bool) {
	WV().SysCallN(889, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) IgnoreCertificateErrors() bool {
	r1 := WV().SysCallN(864, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIgnoreCertificateErrors(AValue bool) {
	WV().SysCallN(864, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) BrowserExecPath() string {
	r1 := WV().SysCallN(793, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetBrowserExecPath(AValue string) {
	WV().SysCallN(793, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) UserDataFolder() string {
	r1 := WV().SysCallN(1042, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetUserDataFolder(AValue string) {
	WV().SysCallN(1042, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) AdditionalBrowserArguments() string {
	r1 := WV().SysCallN(784, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetAdditionalBrowserArguments(AValue string) {
	WV().SysCallN(784, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) Language() string {
	r1 := WV().SysCallN(881, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetLanguage(AValue string) {
	WV().SysCallN(881, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) TargetCompatibleBrowserVersion() string {
	r1 := WV().SysCallN(1038, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetTargetCompatibleBrowserVersion(AValue string) {
	WV().SysCallN(1038, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) AllowSingleSignOnUsingOSPrimaryAccount() bool {
	r1 := WV().SysCallN(786, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetAllowSingleSignOnUsingOSPrimaryAccount(AValue bool) {
	WV().SysCallN(786, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) ExclusiveUserDataFolderAccess() bool {
	r1 := WV().SysCallN(849, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetExclusiveUserDataFolderAccess(AValue bool) {
	WV().SysCallN(849, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) CustomCrashReportingEnabled() bool {
	r1 := WV().SysCallN(827, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetCustomCrashReportingEnabled(AValue bool) {
	WV().SysCallN(827, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) EnableTrackingPrevention() bool {
	r1 := WV().SysCallN(848, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetEnableTrackingPrevention(AValue bool) {
	WV().SysCallN(848, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) AreBrowserExtensionsEnabled() bool {
	r1 := WV().SysCallN(788, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetAreBrowserExtensionsEnabled(AValue bool) {
	WV().SysCallN(788, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) BrowserVersionInfo() string {
	r1 := WV().SysCallN(795, m.Instance())
	return GoStr(r1)
}

func (m *TWVBrowserBase) BrowserProcessID() uint32 {
	r1 := WV().SysCallN(794, m.Instance())
	return uint32(r1)
}

func (m *TWVBrowserBase) CanGoBack() bool {
	r1 := WV().SysCallN(799, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) CanGoForward() bool {
	r1 := WV().SysCallN(800, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) ContainsFullScreenElement() bool {
	r1 := WV().SysCallN(811, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) DocumentTitle() string {
	r1 := WV().SysCallN(843, m.Instance())
	return GoStr(r1)
}

func (m *TWVBrowserBase) Source() string {
	r1 := WV().SysCallN(1032, m.Instance())
	return GoStr(r1)
}

func (m *TWVBrowserBase) CookieManager() ICoreWebView2CookieManager {
	var resultCoreWebView2CookieManager uintptr
	WV().SysCallN(812, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CookieManager)))
	return AsCoreWebView2CookieManager(resultCoreWebView2CookieManager)
}

func (m *TWVBrowserBase) IsSuspended() bool {
	r1 := WV().SysCallN(876, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) IsDocumentPlayingAudio() bool {
	r1 := WV().SysCallN(868, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) IsMuted() bool {
	r1 := WV().SysCallN(871, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsMuted(AValue bool) {
	WV().SysCallN(871, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) DefaultDownloadDialogCornerAlignment() TWVDefaultDownloadDialogCornerAlignment {
	r1 := WV().SysCallN(832, 0, m.Instance(), 0)
	return TWVDefaultDownloadDialogCornerAlignment(r1)
}

func (m *TWVBrowserBase) SetDefaultDownloadDialogCornerAlignment(AValue TWVDefaultDownloadDialogCornerAlignment) {
	WV().SysCallN(832, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) DefaultDownloadDialogMargin() (resultPoint TPoint) {
	WV().SysCallN(833, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TWVBrowserBase) SetDefaultDownloadDialogMargin(AValue *TPoint) {
	WV().SysCallN(833, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TWVBrowserBase) IsDefaultDownloadDialogOpen() bool {
	r1 := WV().SysCallN(867, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) StatusBarText() string {
	r1 := WV().SysCallN(1034, m.Instance())
	return GoStr(r1)
}

func (m *TWVBrowserBase) FaviconURI() string {
	r1 := WV().SysCallN(852, m.Instance())
	return GoStr(r1)
}

func (m *TWVBrowserBase) MemoryUsageTargetLevel() TWVMemoryUsageTargetLevel {
	r1 := WV().SysCallN(882, 0, m.Instance(), 0)
	return TWVMemoryUsageTargetLevel(r1)
}

func (m *TWVBrowserBase) SetMemoryUsageTargetLevel(AValue TWVMemoryUsageTargetLevel) {
	WV().SysCallN(882, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) Bounds() (resultRect TRect) {
	WV().SysCallN(791, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TWVBrowserBase) SetBounds(AValue *TRect) {
	WV().SysCallN(791, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TWVBrowserBase) IsVisible() bool {
	r1 := WV().SysCallN(878, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsVisible(AValue bool) {
	WV().SysCallN(878, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) ParentWindow() THandle {
	r1 := WV().SysCallN(893, 0, m.Instance(), 0)
	return THandle(r1)
}

func (m *TWVBrowserBase) SetParentWindow(AValue THandle) {
	WV().SysCallN(893, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) ZoomFactor() (resultDouble float64) {
	WV().SysCallN(1047, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TWVBrowserBase) SetZoomFactor(AValue float64) {
	WV().SysCallN(1047, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TWVBrowserBase) DefaultBackgroundColor() TColor {
	r1 := WV().SysCallN(830, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TWVBrowserBase) SetDefaultBackgroundColor(AValue TColor) {
	WV().SysCallN(830, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) BoundsMode() TWVBoundsMode {
	r1 := WV().SysCallN(792, 0, m.Instance(), 0)
	return TWVBoundsMode(r1)
}

func (m *TWVBrowserBase) SetBoundsMode(AValue TWVBoundsMode) {
	WV().SysCallN(792, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) RasterizationScale() (resultDouble float64) {
	WV().SysCallN(909, 0, m.Instance(), uintptr(unsafePointer(&resultDouble)), uintptr(unsafePointer(&resultDouble)))
	return
}

func (m *TWVBrowserBase) SetRasterizationScale(AValue float64) {
	WV().SysCallN(909, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TWVBrowserBase) ShouldDetectMonitorScaleChanges() bool {
	r1 := WV().SysCallN(1028, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetShouldDetectMonitorScaleChanges(AValue bool) {
	WV().SysCallN(1028, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) AllowExternalDrop() bool {
	r1 := WV().SysCallN(785, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetAllowExternalDrop(AValue bool) {
	WV().SysCallN(785, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) DefaultContextMenusEnabled() bool {
	r1 := WV().SysCallN(831, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetDefaultContextMenusEnabled(AValue bool) {
	WV().SysCallN(831, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) DefaultScriptDialogsEnabled() bool {
	r1 := WV().SysCallN(835, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetDefaultScriptDialogsEnabled(AValue bool) {
	WV().SysCallN(835, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) DevToolsEnabled() bool {
	r1 := WV().SysCallN(842, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetDevToolsEnabled(AValue bool) {
	WV().SysCallN(842, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) AreHostObjectsAllowed() bool {
	r1 := WV().SysCallN(789, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetAreHostObjectsAllowed(AValue bool) {
	WV().SysCallN(789, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) BuiltInErrorPageEnabled() bool {
	r1 := WV().SysCallN(796, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetBuiltInErrorPageEnabled(AValue bool) {
	WV().SysCallN(796, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) ScriptEnabled() bool {
	r1 := WV().SysCallN(925, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetScriptEnabled(AValue bool) {
	WV().SysCallN(925, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) StatusBarEnabled() bool {
	r1 := WV().SysCallN(1033, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetStatusBarEnabled(AValue bool) {
	WV().SysCallN(1033, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) WebMessageEnabled() bool {
	r1 := WV().SysCallN(1043, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetWebMessageEnabled(AValue bool) {
	WV().SysCallN(1043, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) ZoomControlEnabled() bool {
	r1 := WV().SysCallN(1046, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetZoomControlEnabled(AValue bool) {
	WV().SysCallN(1046, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) UserAgent() string {
	r1 := WV().SysCallN(1041, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetUserAgent(AValue string) {
	WV().SysCallN(1041, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) AreBrowserAcceleratorKeysEnabled() bool {
	r1 := WV().SysCallN(787, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetAreBrowserAcceleratorKeysEnabled(AValue bool) {
	WV().SysCallN(787, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) IsGeneralAutofillEnabled() bool {
	r1 := WV().SysCallN(869, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsGeneralAutofillEnabled(AValue bool) {
	WV().SysCallN(869, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) IsPasswordAutosaveEnabled() bool {
	r1 := WV().SysCallN(873, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsPasswordAutosaveEnabled(AValue bool) {
	WV().SysCallN(873, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) IsPinchZoomEnabled() bool {
	r1 := WV().SysCallN(874, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsPinchZoomEnabled(AValue bool) {
	WV().SysCallN(874, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) IsSwipeNavigationEnabled() bool {
	r1 := WV().SysCallN(877, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsSwipeNavigationEnabled(AValue bool) {
	WV().SysCallN(877, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) HiddenPdfToolbarItems() TWVPDFToolbarItems {
	r1 := WV().SysCallN(863, 0, m.Instance(), 0)
	return TWVPDFToolbarItems(r1)
}

func (m *TWVBrowserBase) SetHiddenPdfToolbarItems(AValue TWVPDFToolbarItems) {
	WV().SysCallN(863, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) IsReputationCheckingRequired() bool {
	r1 := WV().SysCallN(875, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsReputationCheckingRequired(AValue bool) {
	WV().SysCallN(875, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) Cursor() HCURSOR {
	r1 := WV().SysCallN(826, m.Instance())
	return HCURSOR(r1)
}

func (m *TWVBrowserBase) RootVisualTarget() IUnknown {
	var resultUnknown uintptr
	WV().SysCallN(923, 0, m.Instance(), 0, uintptr(unsafePointer(&resultUnknown)))
	return AsUnknown(resultUnknown)
}

func (m *TWVBrowserBase) SetRootVisualTarget(AValue IUnknown) {
	WV().SysCallN(923, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TWVBrowserBase) SystemCursorID() uint32 {
	r1 := WV().SysCallN(1037, m.Instance())
	return uint32(r1)
}

func (m *TWVBrowserBase) AutomationProvider() IUnknown {
	var resultUnknown uintptr
	WV().SysCallN(790, m.Instance(), uintptr(unsafePointer(&resultUnknown)))
	return AsUnknown(resultUnknown)
}

func (m *TWVBrowserBase) ProcessInfos() ICoreWebView2ProcessInfoCollection {
	var resultCoreWebView2ProcessInfoCollection uintptr
	WV().SysCallN(902, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessInfoCollection)))
	return AsCoreWebView2ProcessInfoCollection(resultCoreWebView2ProcessInfoCollection)
}

func (m *TWVBrowserBase) ProfileName() string {
	r1 := WV().SysCallN(907, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetProfileName(AValue string) {
	WV().SysCallN(907, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) IsInPrivateModeEnabled() bool {
	r1 := WV().SysCallN(870, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetIsInPrivateModeEnabled(AValue bool) {
	WV().SysCallN(870, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) ScriptLocale() string {
	r1 := WV().SysCallN(926, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetScriptLocale(AValue string) {
	WV().SysCallN(926, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) ProfilePath() string {
	r1 := WV().SysCallN(908, m.Instance())
	return GoStr(r1)
}

func (m *TWVBrowserBase) DefaultDownloadFolderPath() string {
	r1 := WV().SysCallN(834, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVBrowserBase) SetDefaultDownloadFolderPath(AValue string) {
	WV().SysCallN(834, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVBrowserBase) PreferredColorScheme() TWVPreferredColorScheme {
	r1 := WV().SysCallN(897, 0, m.Instance(), 0)
	return TWVPreferredColorScheme(r1)
}

func (m *TWVBrowserBase) SetPreferredColorScheme(AValue TWVPreferredColorScheme) {
	WV().SysCallN(897, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) PreferredTrackingPreventionLevel() TWVTrackingPreventionLevel {
	r1 := WV().SysCallN(898, 0, m.Instance(), 0)
	return TWVTrackingPreventionLevel(r1)
}

func (m *TWVBrowserBase) SetPreferredTrackingPreventionLevel(AValue TWVTrackingPreventionLevel) {
	WV().SysCallN(898, 1, m.Instance(), uintptr(AValue))
}

func (m *TWVBrowserBase) ProfileCookieManager() ICoreWebView2CookieManager {
	var resultCoreWebView2CookieManager uintptr
	WV().SysCallN(904, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CookieManager)))
	return AsCoreWebView2CookieManager(resultCoreWebView2CookieManager)
}

func (m *TWVBrowserBase) ProfileIsPasswordAutosaveEnabled() bool {
	r1 := WV().SysCallN(906, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetProfileIsPasswordAutosaveEnabled(AValue bool) {
	WV().SysCallN(906, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) ProfileIsGeneralAutofillEnabled() bool {
	r1 := WV().SysCallN(905, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetProfileIsGeneralAutofillEnabled(AValue bool) {
	WV().SysCallN(905, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVBrowserBase) FrameId() uint32 {
	r1 := WV().SysCallN(855, m.Instance())
	return uint32(r1)
}

func (m *TWVBrowserBase) CreateBrowser(aHandle THandle, aUseDefaultEnvironment bool) bool {
	r1 := WV().SysCallN(820, m.Instance(), uintptr(aHandle), PascalBool(aUseDefaultEnvironment))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CreateBrowser1(aHandle THandle, aEnvironment ICoreWebView2Environment) bool {
	r1 := WV().SysCallN(821, m.Instance(), uintptr(aHandle), GetObjectUintptr(aEnvironment))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CreateWindowlessBrowser(aHandle THandle, aUseDefaultEnvironment bool) bool {
	r1 := WV().SysCallN(824, m.Instance(), uintptr(aHandle), PascalBool(aUseDefaultEnvironment))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CreateWindowlessBrowser1(aHandle THandle, aEnvironment ICoreWebView2Environment) bool {
	r1 := WV().SysCallN(825, m.Instance(), uintptr(aHandle), GetObjectUintptr(aEnvironment))
	return GoBool(r1)
}

func (m *TWVBrowserBase) GoBack() bool {
	r1 := WV().SysCallN(861, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) GoForward() bool {
	r1 := WV().SysCallN(862, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) Refresh() bool {
	r1 := WV().SysCallN(910, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) RefreshIgnoreCache() bool {
	r1 := WV().SysCallN(911, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) Stop() bool {
	r1 := WV().SysCallN(1035, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) Navigate(aURI string) bool {
	r1 := WV().SysCallN(885, m.Instance(), PascalStr(aURI))
	return GoBool(r1)
}

func (m *TWVBrowserBase) NavigateToString(aHTMLContent string) bool {
	r1 := WV().SysCallN(886, m.Instance(), PascalStr(aHTMLContent))
	return GoBool(r1)
}

func (m *TWVBrowserBase) NavigateWithWebResourceRequest(aRequest ICoreWebView2WebResourceRequestRef) bool {
	r1 := WV().SysCallN(887, m.Instance(), GetObjectUintptr(aRequest))
	return GoBool(r1)
}

func (m *TWVBrowserBase) SubscribeToDevToolsProtocolEvent(aEventName string, aEventID int32) bool {
	r1 := WV().SysCallN(1036, m.Instance(), PascalStr(aEventName), uintptr(aEventID))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CallDevToolsProtocolMethod(aMethodName, aParametersAsJson string, aExecutionID int32) bool {
	r1 := WV().SysCallN(797, m.Instance(), PascalStr(aMethodName), PascalStr(aParametersAsJson), uintptr(aExecutionID))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CallDevToolsProtocolMethodForSession(aSessionId, aMethodName, aParametersAsJson string, aExecutionID int32) bool {
	r1 := WV().SysCallN(798, m.Instance(), PascalStr(aSessionId), PascalStr(aMethodName), PascalStr(aParametersAsJson), uintptr(aExecutionID))
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetFocus() bool {
	r1 := WV().SysCallN(930, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) FocusNext() bool {
	r1 := WV().SysCallN(853, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) FocusPrevious() bool {
	r1 := WV().SysCallN(854, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) ExecuteScriptWithResult(aJavaScript string, aExecutionID int32) bool {
	r1 := WV().SysCallN(851, m.Instance(), PascalStr(aJavaScript), uintptr(aExecutionID))
	return GoBool(r1)
}

func (m *TWVBrowserBase) ExecuteScript(aJavaScript string, aExecutionID int32) bool {
	r1 := WV().SysCallN(850, m.Instance(), PascalStr(aJavaScript), uintptr(aExecutionID))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CapturePreview(aImageFormat TWVCapturePreviewImageFormat, aImageStream IStream) bool {
	r1 := WV().SysCallN(801, m.Instance(), uintptr(aImageFormat), GetObjectUintptr(aImageStream))
	return GoBool(r1)
}

func (m *TWVBrowserBase) NotifyParentWindowPositionChanged() bool {
	r1 := WV().SysCallN(888, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetPermissionState(aPermissionKind TWVPermissionKind, aOrigin string, aState TWVPermissionState) bool {
	r1 := WV().SysCallN(1026, m.Instance(), uintptr(aPermissionKind), PascalStr(aOrigin), uintptr(aState))
	return GoBool(r1)
}

func (m *TWVBrowserBase) GetNonDefaultPermissionSettings() bool {
	r1 := WV().SysCallN(859, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) AddBrowserExtension(extensionFolderPath string) bool {
	r1 := WV().SysCallN(780, m.Instance(), PascalStr(extensionFolderPath))
	return GoBool(r1)
}

func (m *TWVBrowserBase) GetBrowserExtensions() bool {
	r1 := WV().SysCallN(856, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) DeleteProfile() bool {
	r1 := WV().SysCallN(841, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) TrySuspend() bool {
	r1 := WV().SysCallN(1040, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) Resume() bool {
	r1 := WV().SysCallN(919, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetVirtualHostNameToFolderMapping(aHostName, aFolderPath string, aAccessKind TWVHostResourceAcccessKind) bool {
	r1 := WV().SysCallN(1027, m.Instance(), PascalStr(aHostName), PascalStr(aFolderPath), uintptr(aAccessKind))
	return GoBool(r1)
}

func (m *TWVBrowserBase) ClearVirtualHostNameToFolderMapping(aHostName string) bool {
	r1 := WV().SysCallN(809, m.Instance(), PascalStr(aHostName))
	return GoBool(r1)
}

func (m *TWVBrowserBase) RetrieveHTML() bool {
	r1 := WV().SysCallN(920, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) RetrieveText() bool {
	r1 := WV().SysCallN(922, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) RetrieveMHTML() bool {
	r1 := WV().SysCallN(921, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) Print() bool {
	r1 := WV().SysCallN(899, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) ShowPrintUI(aUseSystemPrintDialog bool) bool {
	r1 := WV().SysCallN(1029, m.Instance(), PascalBool(aUseSystemPrintDialog))
	return GoBool(r1)
}

func (m *TWVBrowserBase) PrintToPdf(aResultFilePath string) bool {
	r1 := WV().SysCallN(900, m.Instance(), PascalStr(aResultFilePath))
	return GoBool(r1)
}

func (m *TWVBrowserBase) PrintToPdfStream() bool {
	r1 := WV().SysCallN(901, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) OpenDevToolsWindow() bool {
	r1 := WV().SysCallN(891, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) OpenTaskManagerWindow() bool {
	r1 := WV().SysCallN(892, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) PostWebMessageAsJson(aWebMessageAsJson string) bool {
	r1 := WV().SysCallN(895, m.Instance(), PascalStr(aWebMessageAsJson))
	return GoBool(r1)
}

func (m *TWVBrowserBase) PostWebMessageAsString(aWebMessageAsString string) bool {
	r1 := WV().SysCallN(896, m.Instance(), PascalStr(aWebMessageAsString))
	return GoBool(r1)
}

func (m *TWVBrowserBase) AddWebResourceRequestedFilter(aURI string, aResourceContext TWVWebResourceContext) bool {
	r1 := WV().SysCallN(783, m.Instance(), PascalStr(aURI), uintptr(aResourceContext))
	return GoBool(r1)
}

func (m *TWVBrowserBase) RemoveWebResourceRequestedFilter(aURI string, aResourceContext TWVWebResourceContext) bool {
	r1 := WV().SysCallN(914, m.Instance(), PascalStr(aURI), uintptr(aResourceContext))
	return GoBool(r1)
}

func (m *TWVBrowserBase) RemoveHostObjectFromScript(aName string) bool {
	r1 := WV().SysCallN(912, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func (m *TWVBrowserBase) AddScriptToExecuteOnDocumentCreated(JavaScript string) bool {
	r1 := WV().SysCallN(782, m.Instance(), PascalStr(JavaScript))
	return GoBool(r1)
}

func (m *TWVBrowserBase) RemoveScriptToExecuteOnDocumentCreated(aID string) bool {
	r1 := WV().SysCallN(913, m.Instance(), PascalStr(aID))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CreateCookie(aName, aValue, aDomain, aPath string) ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	WV().SysCallN(822, m.Instance(), PascalStr(aName), PascalStr(aValue), PascalStr(aDomain), PascalStr(aPath), uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func (m *TWVBrowserBase) CopyCookie(aCookie ICoreWebView2Cookie) ICoreWebView2Cookie {
	var resultCoreWebView2Cookie uintptr
	WV().SysCallN(813, m.Instance(), GetObjectUintptr(aCookie), uintptr(unsafePointer(&resultCoreWebView2Cookie)))
	return AsCoreWebView2Cookie(resultCoreWebView2Cookie)
}

func (m *TWVBrowserBase) GetCookies(aURI string) bool {
	r1 := WV().SysCallN(857, m.Instance(), PascalStr(aURI))
	return GoBool(r1)
}

func (m *TWVBrowserBase) AddOrUpdateCookie(aCookie ICoreWebView2Cookie) bool {
	r1 := WV().SysCallN(781, m.Instance(), GetObjectUintptr(aCookie))
	return GoBool(r1)
}

func (m *TWVBrowserBase) DeleteCookie(aCookie ICoreWebView2Cookie) bool {
	r1 := WV().SysCallN(838, m.Instance(), GetObjectUintptr(aCookie))
	return GoBool(r1)
}

func (m *TWVBrowserBase) DeleteCookies(aName, aURI string) bool {
	r1 := WV().SysCallN(839, m.Instance(), PascalStr(aName), PascalStr(aURI))
	return GoBool(r1)
}

func (m *TWVBrowserBase) DeleteCookiesWithDomainAndPath(aName, aDomain, aPath string) bool {
	r1 := WV().SysCallN(840, m.Instance(), PascalStr(aName), PascalStr(aDomain), PascalStr(aPath))
	return GoBool(r1)
}

func (m *TWVBrowserBase) DeleteAllCookies() bool {
	r1 := WV().SysCallN(837, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) SetBoundsAndZoomFactor(aBounds *TRect, aZoomFactor float64) bool {
	r1 := WV().SysCallN(929, m.Instance(), uintptr(unsafePointer(aBounds)), uintptr(unsafePointer(&aZoomFactor)))
	return GoBool(r1)
}

func (m *TWVBrowserBase) OpenDefaultDownloadDialog() bool {
	r1 := WV().SysCallN(890, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) CloseDefaultDownloadDialog() bool {
	r1 := WV().SysCallN(810, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) SimulateEditingCommand(aEditingCommand TWV2EditingCommand) bool {
	r1 := WV().SysCallN(1030, m.Instance(), uintptr(aEditingCommand))
	return GoBool(r1)
}

func (m *TWVBrowserBase) SimulateKeyEvent(type_ TWV2KeyEventType, modifiers, windowsVirtualKeyCode, nativeVirtualKeyCode int32, timestamp int32, location int32, autoRepeat bool, isKeypad bool, isSystemKey bool, text string, unmodifiedtext string, keyIdentifier string, code string, key string) bool {
	r1 := WV().SysCallN(1031, m.Instance(), uintptr(type_), uintptr(modifiers), uintptr(windowsVirtualKeyCode), uintptr(nativeVirtualKeyCode), uintptr(timestamp), uintptr(location), PascalBool(autoRepeat), PascalBool(isKeypad), PascalBool(isSystemKey), PascalStr(text), PascalStr(unmodifiedtext), PascalStr(keyIdentifier), PascalStr(code), PascalStr(key))
	return GoBool(r1)
}

func (m *TWVBrowserBase) KeyboardShortcutSearch() bool {
	r1 := WV().SysCallN(880, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) KeyboardShortcutRefreshIgnoreCache() bool {
	r1 := WV().SysCallN(879, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) SendMouseInput(aEventKind TWVMouseEventKind, aVirtualKeys TWVMouseEventVirtualKeys, aMouseData uint32, aPoint *TPoint) bool {
	r1 := WV().SysCallN(927, m.Instance(), uintptr(aEventKind), uintptr(aVirtualKeys), uintptr(aMouseData), uintptr(unsafePointer(aPoint)))
	return GoBool(r1)
}

func (m *TWVBrowserBase) SendPointerInput(aEventKind TWVPointerEventKind, aPointerInfo ICoreWebView2PointerInfo) bool {
	r1 := WV().SysCallN(928, m.Instance(), uintptr(aEventKind), GetObjectUintptr(aPointerInfo))
	return GoBool(r1)
}

func (m *TWVBrowserBase) DragEnter(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 {
	var result3 uintptr
	r1 := WV().SysCallN(844, m.Instance(), GetObjectUintptr(dataObject), uintptr(keyState), uintptr(unsafePointer(point)), uintptr(unsafePointer(&result3)))
	*OutEffect = uint32(result3)
	return int32(r1)
}

func (m *TWVBrowserBase) DragLeave() int32 {
	r1 := WV().SysCallN(845, m.Instance())
	return int32(r1)
}

func (m *TWVBrowserBase) DragOver(keyState uint32, point *TPoint, OutEffect *uint32) int32 {
	var result2 uintptr
	r1 := WV().SysCallN(846, m.Instance(), uintptr(keyState), uintptr(unsafePointer(point)), uintptr(unsafePointer(&result2)))
	*OutEffect = uint32(result2)
	return int32(r1)
}

func (m *TWVBrowserBase) Drop(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 {
	var result3 uintptr
	r1 := WV().SysCallN(847, m.Instance(), GetObjectUintptr(dataObject), uintptr(keyState), uintptr(unsafePointer(point)), uintptr(unsafePointer(&result3)))
	*OutEffect = uint32(result3)
	return int32(r1)
}

func (m *TWVBrowserBase) ClearCache() bool {
	r1 := WV().SysCallN(806, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) ClearDataForOrigin(aOrigin string, aStorageTypes TWVClearDataStorageTypes) bool {
	r1 := WV().SysCallN(807, m.Instance(), PascalStr(aOrigin), uintptr(aStorageTypes))
	return GoBool(r1)
}

func (m *TWVBrowserBase) ClearBrowsingData(dataKinds TWVBrowsingDataKinds) bool {
	r1 := WV().SysCallN(803, m.Instance(), uintptr(dataKinds))
	return GoBool(r1)
}

func (m *TWVBrowserBase) ClearBrowsingDataInTimeRange(dataKinds TWVBrowsingDataKinds, startTime, endTime TDateTime) bool {
	r1 := WV().SysCallN(805, m.Instance(), uintptr(dataKinds), uintptr(startTime), uintptr(endTime))
	return GoBool(r1)
}

func (m *TWVBrowserBase) ClearBrowsingDataAll() bool {
	r1 := WV().SysCallN(804, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) ClearServerCertificateErrorActions() bool {
	r1 := WV().SysCallN(808, m.Instance())
	return GoBool(r1)
}

func (m *TWVBrowserBase) GetFavicon(aFormat TWVFaviconImageFormat) bool {
	r1 := WV().SysCallN(858, m.Instance(), uintptr(aFormat))
	return GoBool(r1)
}

func (m *TWVBrowserBase) CreateSharedBuffer(aSize int64, aSharedBuffer *ICoreWebView2SharedBuffer) bool {
	var result1 uintptr
	r1 := WV().SysCallN(823, m.Instance(), uintptr(unsafePointer(&aSize)), uintptr(unsafePointer(&result1)))
	*aSharedBuffer = AsCoreWebView2SharedBuffer(result1)
	return GoBool(r1)
}

func (m *TWVBrowserBase) PostSharedBufferToScript(aSharedBuffer ICoreWebView2SharedBuffer, aAccess TWVSharedBufferAccess, aAdditionalDataAsJson string) bool {
	r1 := WV().SysCallN(894, m.Instance(), GetObjectUintptr(aSharedBuffer), uintptr(aAccess), PascalStr(aAdditionalDataAsJson))
	return GoBool(r1)
}

func (m *TWVBrowserBase) GetProcessExtendedInfos() bool {
	r1 := WV().SysCallN(860, m.Instance())
	return GoBool(r1)
}

func WVBrowserBaseClass() TClass {
	ret := WV().SysCallN(802)
	return TClass(ret)
}

func (m *TWVBrowserBase) MoveFormTo(x, y int32) {
	WV().SysCallN(884, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TWVBrowserBase) MoveFormBy(x, y int32) {
	WV().SysCallN(883, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TWVBrowserBase) ResizeFormWidthTo(x int32) {
	WV().SysCallN(918, m.Instance(), uintptr(x))
}

func (m *TWVBrowserBase) ResizeFormHeightTo(y int32) {
	WV().SysCallN(917, m.Instance(), uintptr(y))
}

func (m *TWVBrowserBase) SetFormLeftTo(x int32) {
	WV().SysCallN(931, m.Instance(), uintptr(x))
}

func (m *TWVBrowserBase) SetFormTopTo(y int32) {
	WV().SysCallN(932, m.Instance(), uintptr(y))
}

func (m *TWVBrowserBase) IncZoomStep() {
	WV().SysCallN(865, m.Instance())
}

func (m *TWVBrowserBase) DecZoomStep() {
	WV().SysCallN(829, m.Instance())
}

func (m *TWVBrowserBase) ResetZoom() {
	WV().SysCallN(916, m.Instance())
}

func (m *TWVBrowserBase) ToggleMuteState() {
	WV().SysCallN(1039, m.Instance())
}

func (m *TWVBrowserBase) SetOnBrowserProcessExited(fn TOnBrowserProcessExitedEvent) {
	if m.browserProcessExitedPtr != 0 {
		RemoveEventElement(m.browserProcessExitedPtr)
	}
	m.browserProcessExitedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(939, m.Instance(), m.browserProcessExitedPtr)
}

func (m *TWVBrowserBase) SetOnProcessInfosChanged(fn TOnProcessInfosChangedEvent) {
	if m.processInfosChangedPtr != 0 {
		RemoveEventElement(m.processInfosChangedPtr)
	}
	m.processInfosChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1000, m.Instance(), m.processInfosChangedPtr)
}

func (m *TWVBrowserBase) SetOnContainsFullScreenElementChanged(fn TNotifyEvent) {
	if m.containsFullScreenElementChangedPtr != 0 {
		RemoveEventElement(m.containsFullScreenElementChangedPtr)
	}
	m.containsFullScreenElementChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(948, m.Instance(), m.containsFullScreenElementChangedPtr)
}

func (m *TWVBrowserBase) SetOnContentLoading(fn TOnContentLoadingEvent) {
	if m.contentLoadingPtr != 0 {
		RemoveEventElement(m.contentLoadingPtr)
	}
	m.contentLoadingPtr = MakeEventDataPtr(fn)
	WV().SysCallN(949, m.Instance(), m.contentLoadingPtr)
}

func (m *TWVBrowserBase) SetOnDocumentTitleChanged(fn TNotifyEvent) {
	if m.documentTitleChangedPtr != 0 {
		RemoveEventElement(m.documentTitleChangedPtr)
	}
	m.documentTitleChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(957, m.Instance(), m.documentTitleChangedPtr)
}

func (m *TWVBrowserBase) SetOnFrameNavigationCompleted(fn TOnNavigationCompletedEvent) {
	if m.frameNavigationCompletedPtr != 0 {
		RemoveEventElement(m.frameNavigationCompletedPtr)
	}
	m.frameNavigationCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(970, m.Instance(), m.frameNavigationCompletedPtr)
}

func (m *TWVBrowserBase) SetOnFrameNavigationStarting(fn TOnNavigationStartingEvent) {
	if m.frameNavigationStartingPtr != 0 {
		RemoveEventElement(m.frameNavigationStartingPtr)
	}
	m.frameNavigationStartingPtr = MakeEventDataPtr(fn)
	WV().SysCallN(972, m.Instance(), m.frameNavigationStartingPtr)
}

func (m *TWVBrowserBase) SetOnHistoryChanged(fn TNotifyEvent) {
	if m.historyChangedPtr != 0 {
		RemoveEventElement(m.historyChangedPtr)
	}
	m.historyChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(982, m.Instance(), m.historyChangedPtr)
}

func (m *TWVBrowserBase) SetOnNavigationCompleted(fn TOnNavigationCompletedEvent) {
	if m.navigationCompletedPtr != 0 {
		RemoveEventElement(m.navigationCompletedPtr)
	}
	m.navigationCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(991, m.Instance(), m.navigationCompletedPtr)
}

func (m *TWVBrowserBase) SetOnNavigationStarting(fn TOnNavigationStartingEvent) {
	if m.navigationStartingPtr != 0 {
		RemoveEventElement(m.navigationStartingPtr)
	}
	m.navigationStartingPtr = MakeEventDataPtr(fn)
	WV().SysCallN(992, m.Instance(), m.navigationStartingPtr)
}

func (m *TWVBrowserBase) SetOnNewWindowRequested(fn TOnNewWindowRequestedEvent) {
	if m.newWindowRequestedPtr != 0 {
		RemoveEventElement(m.newWindowRequestedPtr)
	}
	m.newWindowRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(993, m.Instance(), m.newWindowRequestedPtr)
}

func (m *TWVBrowserBase) SetOnPermissionRequested(fn TOnPermissionRequestedEvent) {
	if m.permissionRequestedPtr != 0 {
		RemoveEventElement(m.permissionRequestedPtr)
	}
	m.permissionRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(995, m.Instance(), m.permissionRequestedPtr)
}

func (m *TWVBrowserBase) SetOnProcessFailed(fn TOnProcessFailedEvent) {
	if m.processFailedPtr != 0 {
		RemoveEventElement(m.processFailedPtr)
	}
	m.processFailedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(999, m.Instance(), m.processFailedPtr)
}

func (m *TWVBrowserBase) SetOnScriptDialogOpening(fn TOnScriptDialogOpeningEvent) {
	if m.scriptDialogOpeningPtr != 0 {
		RemoveEventElement(m.scriptDialogOpeningPtr)
	}
	m.scriptDialogOpeningPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1010, m.Instance(), m.scriptDialogOpeningPtr)
}

func (m *TWVBrowserBase) SetOnSourceChanged(fn TOnSourceChangedEvent) {
	if m.sourceChangedPtr != 0 {
		RemoveEventElement(m.sourceChangedPtr)
	}
	m.sourceChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1015, m.Instance(), m.sourceChangedPtr)
}

func (m *TWVBrowserBase) SetOnWebMessageReceived(fn TOnWebMessageReceivedEvent) {
	if m.webMessageReceivedPtr != 0 {
		RemoveEventElement(m.webMessageReceivedPtr)
	}
	m.webMessageReceivedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1018, m.Instance(), m.webMessageReceivedPtr)
}

func (m *TWVBrowserBase) SetOnWebResourceRequested(fn TOnWebResourceRequestedEvent) {
	if m.webResourceRequestedPtr != 0 {
		RemoveEventElement(m.webResourceRequestedPtr)
	}
	m.webResourceRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1019, m.Instance(), m.webResourceRequestedPtr)
}

func (m *TWVBrowserBase) SetOnWindowCloseRequested(fn TNotifyEvent) {
	if m.windowCloseRequestedPtr != 0 {
		RemoveEventElement(m.windowCloseRequestedPtr)
	}
	m.windowCloseRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1024, m.Instance(), m.windowCloseRequestedPtr)
}

func (m *TWVBrowserBase) SetOnDOMContentLoaded(fn TOnDOMContentLoadedEvent) {
	if m.dOMContentLoadedPtr != 0 {
		RemoveEventElement(m.dOMContentLoadedPtr)
	}
	m.dOMContentLoadedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(955, m.Instance(), m.dOMContentLoadedPtr)
}

func (m *TWVBrowserBase) SetOnWebResourceResponseReceived(fn TOnWebResourceResponseReceivedEvent) {
	if m.webResourceResponseReceivedPtr != 0 {
		RemoveEventElement(m.webResourceResponseReceivedPtr)
	}
	m.webResourceResponseReceivedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1020, m.Instance(), m.webResourceResponseReceivedPtr)
}

func (m *TWVBrowserBase) SetOnDownloadStarting(fn TOnDownloadStartingEvent) {
	if m.downloadStartingPtr != 0 {
		RemoveEventElement(m.downloadStartingPtr)
	}
	m.downloadStartingPtr = MakeEventDataPtr(fn)
	WV().SysCallN(958, m.Instance(), m.downloadStartingPtr)
}

func (m *TWVBrowserBase) SetOnFrameCreated(fn TOnFrameCreatedEvent) {
	if m.frameCreatedPtr != 0 {
		RemoveEventElement(m.frameCreatedPtr)
	}
	m.frameCreatedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(966, m.Instance(), m.frameCreatedPtr)
}

func (m *TWVBrowserBase) SetOnClientCertificateRequested(fn TOnClientCertificateRequestedEvent) {
	if m.clientCertificateRequestedPtr != 0 {
		RemoveEventElement(m.clientCertificateRequestedPtr)
	}
	m.clientCertificateRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(946, m.Instance(), m.clientCertificateRequestedPtr)
}

func (m *TWVBrowserBase) SetOnIsDocumentPlayingAudioChanged(fn TOnIsDocumentPlayingAudioChangedEvent) {
	if m.isDocumentPlayingAudioChangedPtr != 0 {
		RemoveEventElement(m.isDocumentPlayingAudioChangedPtr)
	}
	m.isDocumentPlayingAudioChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(986, m.Instance(), m.isDocumentPlayingAudioChangedPtr)
}

func (m *TWVBrowserBase) SetOnIsMutedChanged(fn TOnIsMutedChangedEvent) {
	if m.isMutedChangedPtr != 0 {
		RemoveEventElement(m.isMutedChangedPtr)
	}
	m.isMutedChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(987, m.Instance(), m.isMutedChangedPtr)
}

func (m *TWVBrowserBase) SetOnIsDefaultDownloadDialogOpenChanged(fn TOnIsDefaultDownloadDialogOpenChangedEvent) {
	if m.isDefaultDownloadDialogOpenChangedPtr != 0 {
		RemoveEventElement(m.isDefaultDownloadDialogOpenChangedPtr)
	}
	m.isDefaultDownloadDialogOpenChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(985, m.Instance(), m.isDefaultDownloadDialogOpenChangedPtr)
}

func (m *TWVBrowserBase) SetOnBasicAuthenticationRequested(fn TOnBasicAuthenticationRequestedEvent) {
	if m.basicAuthenticationRequestedPtr != 0 {
		RemoveEventElement(m.basicAuthenticationRequestedPtr)
	}
	m.basicAuthenticationRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(936, m.Instance(), m.basicAuthenticationRequestedPtr)
}

func (m *TWVBrowserBase) SetOnContextMenuRequested(fn TOnContextMenuRequestedEvent) {
	if m.contextMenuRequestedPtr != 0 {
		RemoveEventElement(m.contextMenuRequestedPtr)
	}
	m.contextMenuRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(950, m.Instance(), m.contextMenuRequestedPtr)
}

func (m *TWVBrowserBase) SetOnStatusBarTextChanged(fn TOnStatusBarTextChangedEvent) {
	if m.statusBarTextChangedPtr != 0 {
		RemoveEventElement(m.statusBarTextChangedPtr)
	}
	m.statusBarTextChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1016, m.Instance(), m.statusBarTextChangedPtr)
}

func (m *TWVBrowserBase) SetOnServerCertificateErrorActionsCompleted(fn TOnServerCertificateErrorActionsCompletedEvent) {
	if m.serverCertificateErrorActionsCompletedPtr != 0 {
		RemoveEventElement(m.serverCertificateErrorActionsCompletedPtr)
	}
	m.serverCertificateErrorActionsCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1011, m.Instance(), m.serverCertificateErrorActionsCompletedPtr)
}

func (m *TWVBrowserBase) SetOnServerCertificateErrorDetected(fn TOnServerCertificateErrorDetectedEvent) {
	if m.serverCertificateErrorDetectedPtr != 0 {
		RemoveEventElement(m.serverCertificateErrorDetectedPtr)
	}
	m.serverCertificateErrorDetectedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1012, m.Instance(), m.serverCertificateErrorDetectedPtr)
}

func (m *TWVBrowserBase) SetOnFaviconChanged(fn TOnFaviconChangedEvent) {
	if m.faviconChangedPtr != 0 {
		RemoveEventElement(m.faviconChangedPtr)
	}
	m.faviconChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(964, m.Instance(), m.faviconChangedPtr)
}

func (m *TWVBrowserBase) SetOnGetFaviconCompleted(fn TOnGetFaviconCompletedEvent) {
	if m.getFaviconCompletedPtr != 0 {
		RemoveEventElement(m.getFaviconCompletedPtr)
	}
	m.getFaviconCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(978, m.Instance(), m.getFaviconCompletedPtr)
}

func (m *TWVBrowserBase) SetOnPrintCompleted(fn TOnPrintCompletedEvent) {
	if m.printCompletedPtr != 0 {
		RemoveEventElement(m.printCompletedPtr)
	}
	m.printCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(996, m.Instance(), m.printCompletedPtr)
}

func (m *TWVBrowserBase) SetOnPrintToPdfStreamCompleted(fn TOnPrintToPdfStreamCompletedEvent) {
	if m.printToPdfStreamCompletedPtr != 0 {
		RemoveEventElement(m.printToPdfStreamCompletedPtr)
	}
	m.printToPdfStreamCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(998, m.Instance(), m.printToPdfStreamCompletedPtr)
}

func (m *TWVBrowserBase) SetOnAcceleratorKeyPressed(fn TOnAcceleratorKeyPressedEvent) {
	if m.acceleratorKeyPressedPtr != 0 {
		RemoveEventElement(m.acceleratorKeyPressedPtr)
	}
	m.acceleratorKeyPressedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(933, m.Instance(), m.acceleratorKeyPressedPtr)
}

func (m *TWVBrowserBase) SetOnGotFocus(fn TNotifyEvent) {
	if m.gotFocusPtr != 0 {
		RemoveEventElement(m.gotFocusPtr)
	}
	m.gotFocusPtr = MakeEventDataPtr(fn)
	WV().SysCallN(981, m.Instance(), m.gotFocusPtr)
}

func (m *TWVBrowserBase) SetOnLostFocus(fn TNotifyEvent) {
	if m.lostFocusPtr != 0 {
		RemoveEventElement(m.lostFocusPtr)
	}
	m.lostFocusPtr = MakeEventDataPtr(fn)
	WV().SysCallN(989, m.Instance(), m.lostFocusPtr)
}

func (m *TWVBrowserBase) SetOnMoveFocusRequested(fn TOnMoveFocusRequestedEvent) {
	if m.moveFocusRequestedPtr != 0 {
		RemoveEventElement(m.moveFocusRequestedPtr)
	}
	m.moveFocusRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(990, m.Instance(), m.moveFocusRequestedPtr)
}

func (m *TWVBrowserBase) SetOnZoomFactorChanged(fn TNotifyEvent) {
	if m.zoomFactorChangedPtr != 0 {
		RemoveEventElement(m.zoomFactorChangedPtr)
	}
	m.zoomFactorChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1025, m.Instance(), m.zoomFactorChangedPtr)
}

func (m *TWVBrowserBase) SetOnRasterizationScaleChanged(fn TNotifyEvent) {
	if m.rasterizationScaleChangedPtr != 0 {
		RemoveEventElement(m.rasterizationScaleChangedPtr)
	}
	m.rasterizationScaleChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1004, m.Instance(), m.rasterizationScaleChangedPtr)
}

func (m *TWVBrowserBase) SetOnCursorChanged(fn TNotifyEvent) {
	if m.cursorChangedPtr != 0 {
		RemoveEventElement(m.cursorChangedPtr)
	}
	m.cursorChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(952, m.Instance(), m.cursorChangedPtr)
}

func (m *TWVBrowserBase) SetOnBytesReceivedChanged(fn TOnBytesReceivedChangedEvent) {
	if m.bytesReceivedChangedPtr != 0 {
		RemoveEventElement(m.bytesReceivedChangedPtr)
	}
	m.bytesReceivedChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(940, m.Instance(), m.bytesReceivedChangedPtr)
}

func (m *TWVBrowserBase) SetOnEstimatedEndTimeChanged(fn TOnEstimatedEndTimeChangedEvent) {
	if m.estimatedEndTimeChangedPtr != 0 {
		RemoveEventElement(m.estimatedEndTimeChangedPtr)
	}
	m.estimatedEndTimeChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(961, m.Instance(), m.estimatedEndTimeChangedPtr)
}

func (m *TWVBrowserBase) SetOnDownloadStateChanged(fn TOnDownloadStateChangedEvent) {
	if m.downloadStateChangedPtr != 0 {
		RemoveEventElement(m.downloadStateChangedPtr)
	}
	m.downloadStateChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(959, m.Instance(), m.downloadStateChangedPtr)
}

func (m *TWVBrowserBase) SetOnFrameDestroyed(fn TOnFrameDestroyedEvent) {
	if m.frameDestroyedPtr != 0 {
		RemoveEventElement(m.frameDestroyedPtr)
	}
	m.frameDestroyedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(968, m.Instance(), m.frameDestroyedPtr)
}

func (m *TWVBrowserBase) SetOnFrameNameChanged(fn TOnFrameNameChangedEvent) {
	if m.frameNameChangedPtr != 0 {
		RemoveEventElement(m.frameNameChangedPtr)
	}
	m.frameNameChangedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(969, m.Instance(), m.frameNameChangedPtr)
}

func (m *TWVBrowserBase) SetOnFrameNavigationStarting2(fn TOnFrameNavigationStartingEvent) {
	if m.frameNavigationStarting2Ptr != 0 {
		RemoveEventElement(m.frameNavigationStarting2Ptr)
	}
	m.frameNavigationStarting2Ptr = MakeEventDataPtr(fn)
	WV().SysCallN(973, m.Instance(), m.frameNavigationStarting2Ptr)
}

func (m *TWVBrowserBase) SetOnFrameNavigationCompleted2(fn TOnFrameNavigationCompletedEvent) {
	if m.frameNavigationCompleted2Ptr != 0 {
		RemoveEventElement(m.frameNavigationCompleted2Ptr)
	}
	m.frameNavigationCompleted2Ptr = MakeEventDataPtr(fn)
	WV().SysCallN(971, m.Instance(), m.frameNavigationCompleted2Ptr)
}

func (m *TWVBrowserBase) SetOnFrameContentLoading(fn TOnFrameContentLoadingEvent) {
	if m.frameContentLoadingPtr != 0 {
		RemoveEventElement(m.frameContentLoadingPtr)
	}
	m.frameContentLoadingPtr = MakeEventDataPtr(fn)
	WV().SysCallN(965, m.Instance(), m.frameContentLoadingPtr)
}

func (m *TWVBrowserBase) SetOnFrameDOMContentLoaded(fn TOnFrameDOMContentLoadedEvent) {
	if m.frameDOMContentLoadedPtr != 0 {
		RemoveEventElement(m.frameDOMContentLoadedPtr)
	}
	m.frameDOMContentLoadedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(967, m.Instance(), m.frameDOMContentLoadedPtr)
}

func (m *TWVBrowserBase) SetOnFrameWebMessageReceived(fn TOnFrameWebMessageReceivedEvent) {
	if m.frameWebMessageReceivedPtr != 0 {
		RemoveEventElement(m.frameWebMessageReceivedPtr)
	}
	m.frameWebMessageReceivedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(975, m.Instance(), m.frameWebMessageReceivedPtr)
}

func (m *TWVBrowserBase) SetOnFramePermissionRequested(fn TOnFramePermissionRequestedEvent) {
	if m.framePermissionRequestedPtr != 0 {
		RemoveEventElement(m.framePermissionRequestedPtr)
	}
	m.framePermissionRequestedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(974, m.Instance(), m.framePermissionRequestedPtr)
}

func (m *TWVBrowserBase) SetOnDevToolsProtocolEventReceived(fn TOnDevToolsProtocolEventReceivedEvent) {
	if m.devToolsProtocolEventReceivedPtr != 0 {
		RemoveEventElement(m.devToolsProtocolEventReceivedPtr)
	}
	m.devToolsProtocolEventReceivedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(956, m.Instance(), m.devToolsProtocolEventReceivedPtr)
}

func (m *TWVBrowserBase) SetOnCustomItemSelected(fn TOnCustomItemSelectedEvent) {
	if m.customItemSelectedPtr != 0 {
		RemoveEventElement(m.customItemSelectedPtr)
	}
	m.customItemSelectedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(953, m.Instance(), m.customItemSelectedPtr)
}

func (m *TWVBrowserBase) SetOnClearBrowsingDataCompleted(fn TOnClearBrowsingDataCompletedEvent) {
	if m.clearBrowsingDataCompletedPtr != 0 {
		RemoveEventElement(m.clearBrowsingDataCompletedPtr)
	}
	m.clearBrowsingDataCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(943, m.Instance(), m.clearBrowsingDataCompletedPtr)
}

func (m *TWVBrowserBase) SetOnInitializationError(fn TOnInitializationErrorEvent) {
	if m.initializationErrorPtr != 0 {
		RemoveEventElement(m.initializationErrorPtr)
	}
	m.initializationErrorPtr = MakeEventDataPtr(fn)
	WV().SysCallN(984, m.Instance(), m.initializationErrorPtr)
}

func (m *TWVBrowserBase) SetOnEnvironmentCompleted(fn TNotifyEvent) {
	if m.environmentCompletedPtr != 0 {
		RemoveEventElement(m.environmentCompletedPtr)
	}
	m.environmentCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(960, m.Instance(), m.environmentCompletedPtr)
}

func (m *TWVBrowserBase) SetOnControllerCompleted(fn TNotifyEvent) {
	if m.controllerCompletedPtr != 0 {
		RemoveEventElement(m.controllerCompletedPtr)
	}
	m.controllerCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(951, m.Instance(), m.controllerCompletedPtr)
}

func (m *TWVBrowserBase) SetOnAfterCreated(fn TNotifyEvent) {
	if m.afterCreatedPtr != 0 {
		RemoveEventElement(m.afterCreatedPtr)
	}
	m.afterCreatedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(935, m.Instance(), m.afterCreatedPtr)
}

func (m *TWVBrowserBase) SetOnExecuteScriptCompleted(fn TOnExecuteScriptCompletedEvent) {
	if m.executeScriptCompletedPtr != 0 {
		RemoveEventElement(m.executeScriptCompletedPtr)
	}
	m.executeScriptCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(962, m.Instance(), m.executeScriptCompletedPtr)
}

func (m *TWVBrowserBase) SetOnCapturePreviewCompleted(fn TOnCapturePreviewCompletedEvent) {
	if m.capturePreviewCompletedPtr != 0 {
		RemoveEventElement(m.capturePreviewCompletedPtr)
	}
	m.capturePreviewCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(942, m.Instance(), m.capturePreviewCompletedPtr)
}

func (m *TWVBrowserBase) SetOnGetCookiesCompleted(fn TOnGetCookiesCompletedEvent) {
	if m.getCookiesCompletedPtr != 0 {
		RemoveEventElement(m.getCookiesCompletedPtr)
	}
	m.getCookiesCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(976, m.Instance(), m.getCookiesCompletedPtr)
}

func (m *TWVBrowserBase) SetOnTrySuspendCompleted(fn TOnTrySuspendCompletedEvent) {
	if m.trySuspendCompletedPtr != 0 {
		RemoveEventElement(m.trySuspendCompletedPtr)
	}
	m.trySuspendCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1017, m.Instance(), m.trySuspendCompletedPtr)
}

func (m *TWVBrowserBase) SetOnPrintToPdfCompleted(fn TOnPrintToPdfCompletedEvent) {
	if m.printToPdfCompletedPtr != 0 {
		RemoveEventElement(m.printToPdfCompletedPtr)
	}
	m.printToPdfCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(997, m.Instance(), m.printToPdfCompletedPtr)
}

func (m *TWVBrowserBase) SetOnCompositionControllerCompleted(fn TNotifyEvent) {
	if m.compositionControllerCompletedPtr != 0 {
		RemoveEventElement(m.compositionControllerCompletedPtr)
	}
	m.compositionControllerCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(947, m.Instance(), m.compositionControllerCompletedPtr)
}

func (m *TWVBrowserBase) SetOnCallDevToolsProtocolMethodCompleted(fn TOnCallDevToolsProtocolMethodCompletedEvent) {
	if m.callDevToolsProtocolMethodCompletedPtr != 0 {
		RemoveEventElement(m.callDevToolsProtocolMethodCompletedPtr)
	}
	m.callDevToolsProtocolMethodCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(941, m.Instance(), m.callDevToolsProtocolMethodCompletedPtr)
}

func (m *TWVBrowserBase) SetOnAddScriptToExecuteOnDocumentCreatedCompleted(fn TOnAddScriptToExecuteOnDocumentCreatedCompletedEvent) {
	if m.addScriptToExecuteOnDocumentCreatedCompletedPtr != 0 {
		RemoveEventElement(m.addScriptToExecuteOnDocumentCreatedCompletedPtr)
	}
	m.addScriptToExecuteOnDocumentCreatedCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(934, m.Instance(), m.addScriptToExecuteOnDocumentCreatedCompletedPtr)
}

func (m *TWVBrowserBase) SetOnWebResourceResponseViewGetContentCompleted(fn TOnWebResourceResponseViewGetContentCompletedEvent) {
	if m.webResourceResponseViewGetContentCompletedPtr != 0 {
		RemoveEventElement(m.webResourceResponseViewGetContentCompletedPtr)
	}
	m.webResourceResponseViewGetContentCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1021, m.Instance(), m.webResourceResponseViewGetContentCompletedPtr)
}

func (m *TWVBrowserBase) SetOnWidget0CompMsg(fn TOnCompMsgEvent) {
	if m.widget0CompMsgPtr != 0 {
		RemoveEventElement(m.widget0CompMsgPtr)
	}
	m.widget0CompMsgPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1022, m.Instance(), m.widget0CompMsgPtr)
}

func (m *TWVBrowserBase) SetOnWidget1CompMsg(fn TOnCompMsgEvent) {
	if m.widget1CompMsgPtr != 0 {
		RemoveEventElement(m.widget1CompMsgPtr)
	}
	m.widget1CompMsgPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1023, m.Instance(), m.widget1CompMsgPtr)
}

func (m *TWVBrowserBase) SetOnRenderCompMsg(fn TOnCompMsgEvent) {
	if m.renderCompMsgPtr != 0 {
		RemoveEventElement(m.renderCompMsgPtr)
	}
	m.renderCompMsgPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1006, m.Instance(), m.renderCompMsgPtr)
}

func (m *TWVBrowserBase) SetOnD3DWindowCompMsg(fn TOnCompMsgEvent) {
	if m.d3DWindowCompMsgPtr != 0 {
		RemoveEventElement(m.d3DWindowCompMsgPtr)
	}
	m.d3DWindowCompMsgPtr = MakeEventDataPtr(fn)
	WV().SysCallN(954, m.Instance(), m.d3DWindowCompMsgPtr)
}

func (m *TWVBrowserBase) SetOnRetrieveHTMLCompleted(fn TOnRetrieveHTMLCompletedEvent) {
	if m.retrieveHTMLCompletedPtr != 0 {
		RemoveEventElement(m.retrieveHTMLCompletedPtr)
	}
	m.retrieveHTMLCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1007, m.Instance(), m.retrieveHTMLCompletedPtr)
}

func (m *TWVBrowserBase) SetOnRetrieveTextCompleted(fn TOnRetrieveTextCompletedEvent) {
	if m.retrieveTextCompletedPtr != 0 {
		RemoveEventElement(m.retrieveTextCompletedPtr)
	}
	m.retrieveTextCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1009, m.Instance(), m.retrieveTextCompletedPtr)
}

func (m *TWVBrowserBase) SetOnRetrieveMHTMLCompleted(fn TOnRetrieveMHTMLCompletedEvent) {
	if m.retrieveMHTMLCompletedPtr != 0 {
		RemoveEventElement(m.retrieveMHTMLCompletedPtr)
	}
	m.retrieveMHTMLCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1008, m.Instance(), m.retrieveMHTMLCompletedPtr)
}

func (m *TWVBrowserBase) SetOnClearCacheCompleted(fn TOnClearCacheCompletedEvent) {
	if m.clearCacheCompletedPtr != 0 {
		RemoveEventElement(m.clearCacheCompletedPtr)
	}
	m.clearCacheCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(944, m.Instance(), m.clearCacheCompletedPtr)
}

func (m *TWVBrowserBase) SetOnClearDataForOriginCompleted(fn TOnClearDataForOriginCompletedEvent) {
	if m.clearDataForOriginCompletedPtr != 0 {
		RemoveEventElement(m.clearDataForOriginCompletedPtr)
	}
	m.clearDataForOriginCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(945, m.Instance(), m.clearDataForOriginCompletedPtr)
}

func (m *TWVBrowserBase) SetOnOfflineCompleted(fn TOnOfflineCompletedEvent) {
	if m.offlineCompletedPtr != 0 {
		RemoveEventElement(m.offlineCompletedPtr)
	}
	m.offlineCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(994, m.Instance(), m.offlineCompletedPtr)
}

func (m *TWVBrowserBase) SetOnIgnoreCertificateErrorsCompleted(fn TOnIgnoreCertificateErrorsCompletedEvent) {
	if m.ignoreCertificateErrorsCompletedPtr != 0 {
		RemoveEventElement(m.ignoreCertificateErrorsCompletedPtr)
	}
	m.ignoreCertificateErrorsCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(983, m.Instance(), m.ignoreCertificateErrorsCompletedPtr)
}

func (m *TWVBrowserBase) SetOnRefreshIgnoreCacheCompleted(fn TOnRefreshIgnoreCacheCompletedEvent) {
	if m.refreshIgnoreCacheCompletedPtr != 0 {
		RemoveEventElement(m.refreshIgnoreCacheCompletedPtr)
	}
	m.refreshIgnoreCacheCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1005, m.Instance(), m.refreshIgnoreCacheCompletedPtr)
}

func (m *TWVBrowserBase) SetOnSimulateKeyEventCompleted(fn TOnSimulateKeyEventCompletedEvent) {
	if m.simulateKeyEventCompletedPtr != 0 {
		RemoveEventElement(m.simulateKeyEventCompletedPtr)
	}
	m.simulateKeyEventCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1014, m.Instance(), m.simulateKeyEventCompletedPtr)
}

func (m *TWVBrowserBase) SetOnGetCustomSchemes(fn TOnGetCustomSchemesEvent) {
	if m.getCustomSchemesPtr != 0 {
		RemoveEventElement(m.getCustomSchemesPtr)
	}
	m.getCustomSchemesPtr = MakeEventDataPtr(fn)
	WV().SysCallN(977, m.Instance(), m.getCustomSchemesPtr)
}

func (m *TWVBrowserBase) SetOnGetNonDefaultPermissionSettingsCompleted(fn TOnGetNonDefaultPermissionSettingsCompletedEvent) {
	if m.getNonDefaultPermissionSettingsCompletedPtr != 0 {
		RemoveEventElement(m.getNonDefaultPermissionSettingsCompletedPtr)
	}
	m.getNonDefaultPermissionSettingsCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(979, m.Instance(), m.getNonDefaultPermissionSettingsCompletedPtr)
}

func (m *TWVBrowserBase) SetOnSetPermissionStateCompleted(fn TOnSetPermissionStateCompletedEvent) {
	if m.setPermissionStateCompletedPtr != 0 {
		RemoveEventElement(m.setPermissionStateCompletedPtr)
	}
	m.setPermissionStateCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1013, m.Instance(), m.setPermissionStateCompletedPtr)
}

func (m *TWVBrowserBase) SetOnLaunchingExternalUriScheme(fn TOnLaunchingExternalUriSchemeEvent) {
	if m.launchingExternalUriSchemePtr != 0 {
		RemoveEventElement(m.launchingExternalUriSchemePtr)
	}
	m.launchingExternalUriSchemePtr = MakeEventDataPtr(fn)
	WV().SysCallN(988, m.Instance(), m.launchingExternalUriSchemePtr)
}

func (m *TWVBrowserBase) SetOnGetProcessExtendedInfosCompleted(fn TOnGetProcessExtendedInfosCompletedEvent) {
	if m.getProcessExtendedInfosCompletedPtr != 0 {
		RemoveEventElement(m.getProcessExtendedInfosCompletedPtr)
	}
	m.getProcessExtendedInfosCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(980, m.Instance(), m.getProcessExtendedInfosCompletedPtr)
}

func (m *TWVBrowserBase) SetOnBrowserExtensionRemoveCompleted(fn TOnBrowserExtensionRemoveCompletedEvent) {
	if m.browserExtensionRemoveCompletedPtr != 0 {
		RemoveEventElement(m.browserExtensionRemoveCompletedPtr)
	}
	m.browserExtensionRemoveCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(938, m.Instance(), m.browserExtensionRemoveCompletedPtr)
}

func (m *TWVBrowserBase) SetOnBrowserExtensionEnableCompleted(fn TOnBrowserExtensionEnableCompletedEvent) {
	if m.browserExtensionEnableCompletedPtr != 0 {
		RemoveEventElement(m.browserExtensionEnableCompletedPtr)
	}
	m.browserExtensionEnableCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(937, m.Instance(), m.browserExtensionEnableCompletedPtr)
}

func (m *TWVBrowserBase) SetOnProfileAddBrowserExtensionCompleted(fn TOnProfileAddBrowserExtensionCompletedEvent) {
	if m.profileAddBrowserExtensionCompletedPtr != 0 {
		RemoveEventElement(m.profileAddBrowserExtensionCompletedPtr)
	}
	m.profileAddBrowserExtensionCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1001, m.Instance(), m.profileAddBrowserExtensionCompletedPtr)
}

func (m *TWVBrowserBase) SetOnProfileGetBrowserExtensionsCompleted(fn TOnProfileGetBrowserExtensionsCompletedEvent) {
	if m.profileGetBrowserExtensionsCompletedPtr != 0 {
		RemoveEventElement(m.profileGetBrowserExtensionsCompletedPtr)
	}
	m.profileGetBrowserExtensionsCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1003, m.Instance(), m.profileGetBrowserExtensionsCompletedPtr)
}

func (m *TWVBrowserBase) SetOnProfileDeleted(fn TOnProfileDeletedEvent) {
	if m.profileDeletedPtr != 0 {
		RemoveEventElement(m.profileDeletedPtr)
	}
	m.profileDeletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(1002, m.Instance(), m.profileDeletedPtr)
}

func (m *TWVBrowserBase) SetOnExecuteScriptWithResultCompleted(fn TOnExecuteScriptWithResultCompletedEvent) {
	if m.executeScriptWithResultCompletedPtr != 0 {
		RemoveEventElement(m.executeScriptWithResultCompletedPtr)
	}
	m.executeScriptWithResultCompletedPtr = MakeEventDataPtr(fn)
	WV().SysCallN(963, m.Instance(), m.executeScriptWithResultCompletedPtr)
}
