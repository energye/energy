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

// IChromiumCore Parent: IComponent
//
//	Parent class of TChromium and TFMXChromium that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type IChromiumCore interface {
	IComponent
	// IMESetComposition
	//  Begins a new composition or updates the existing composition. Blink has a
	//  special node (a composition node) that allows the input function to change
	//  text without affecting other DOM nodes. |text| is the optional text that
	//  will be inserted into the composition node. |underlines| is an optional
	//  set of ranges that will be underlined in the resulting text.
	//  |replacement_range| is an optional range of the existing text that will be
	//  replaced. |selection_range| is an optional range of the resulting text
	//  that will be selected after insertion or replacement. The
	//  |replacement_range| value is only used on OS X.
	//  This function may be called multiple times as the composition changes.
	//  When the client is done making changes the composition should either be
	//  canceled or completed. To cancel the composition call
	//  ImeCancelComposition. To complete the composition call either
	//  ImeCommitText or ImeFinishComposingText. Completion is usually signaled
	//  when:
	//  1. The client receives a WM_IME_COMPOSITION message with a GCS_RESULTSTR
	//  flag (on Windows), or;
	//  2. The client receives a "commit" signal of GtkIMContext (on Linux), or;
	//  3. insertText of NSTextInput is called (on Mac).
	//  This function is only used when window rendering is disabled.
	IMESetComposition(text string, underlines TCefCompositionUnderlineDynArray, replacementrange, selectionrange *TCefRange)
	// SetCookie
	//  TChromiumCore.SetCookie triggers the TChromiumCore.OnCookieSet event when the cookie has been set
	//  aID is an optional parameter to identify which SetCookie call has triggered the
	//  OnCookieSet event.
	SetCookie(url string, aSetImmediately bool, aID int32, cookie TCookie) bool
	// DefaultUrl
	//  First URL loaded by the browser after its creation.
	DefaultUrl() string // property
	// SetDefaultUrl Set DefaultUrl
	SetDefaultUrl(AValue string) // property
	// Options
	//  Properties used to fill the TCefBrowserSettings record which is used during the browser creation.
	Options() IChromiumOptions // property
	// SetOptions Set Options
	SetOptions(AValue IChromiumOptions) // property
	// FontOptions
	//  Font properties used to fill the TCefBrowserSettings record which is used during the browser creation.
	FontOptions() IChromiumFontOptions // property
	// SetFontOptions Set FontOptions
	SetFontOptions(AValue IChromiumFontOptions) // property
	// PDFPrintOptions
	//  Properties used to fill the TCefPdfPrintSettings record which is used in the TChromiumCore.PrintToPDF call.
	PDFPrintOptions() IPDFPrintOptions // property
	// SetPDFPrintOptions Set PDFPrintOptions
	SetPDFPrintOptions(AValue IPDFPrintOptions) // property
	// BrowserId
	//  Default encoding for Web content. If empty "ISO-8859-1" will be used. Also
	//  configurable using the "default-encoding" command-line switch. It's used during the browser creation.
	//  Globally unique identifier for the seleted browser.
	BrowserId() int32 // property
	// Browser
	//  Returns a ICefBrowser instance of the selected browser.
	Browser() ICefBrowser // property
	// BrowserById
	//  Returns a ICefBrowser instance of the browser with the specified id.
	BrowserById(id int32) ICefBrowser // property
	// BrowserCount
	//  Returns the number of browsers in the browser array when the multi-browser mode is enabled.
	BrowserCount() int32 // property
	// BrowserIdByIndex
	//  Returns the identifier of the browser in the specified array position when the multi-browser mode is enabled.
	BrowserIdByIndex(i int32) int32 // property
	// CefClient
	//  Returns the ICefClient instance used by the selected browser.
	CefClient() ICefClient // property
	// ReqContextHandler
	//  Returns the ICefRequestContextHandler instance used in this browser.
	ReqContextHandler() ICefRequestContextHandler // property
	// ResourceRequestHandler
	//  Returns the ICefResourceRequestHandler instance used in this browser.
	ResourceRequestHandler() ICefResourceRequestHandler // property
	// CefWindowInfo
	//  Returns the TCefWindowInfo record used when the browser was created.
	CefWindowInfo() (resultCefWindowInfo TCefWindowInfo) // property
	// VisibleNavigationEntry
	//  Returns the current visible navigation entry for this browser. This
	//  property can only be read on the UI thread.
	VisibleNavigationEntry() ICefNavigationEntry // property
	// RequestContext
	//  Returns a ICefRequestContext instance used by the selected browser.
	RequestContext() ICefRequestContext // property
	// MediaRouter
	//  Returns a ICefMediaRouter instance used by the selected browser.
	MediaRouter() ICefMediaRouter // property
	// MediaObserver
	//  Returns a ICefMediaObserver instance used by the selected browser.
	MediaObserver() ICefMediaObserver // property
	// MediaObserverReg
	//  Returns the ICefRegistration instance obtained when the default MediaObserver was added.
	MediaObserverReg() ICefRegistration // property
	// DevToolsMsgObserver
	//  Returns a ICefDevToolsMessageObserver instance used by the selected browser.
	DevToolsMsgObserver() ICefDevToolsMessageObserver // property
	// DevToolsMsgObserverReg
	//  Returns the ICefRegistration instance obtained when the default DevToolsMessageObserver was added.
	DevToolsMsgObserverReg() ICefRegistration // property
	// ExtensionHandler
	//  Returns a ICefExtensionHandler instance used by the selected browser.
	ExtensionHandler() ICefExtensionHandler // property
	// MultithreadApp
	//  Returns the value of GlobalCEFApp.MultiThreadedMessageLoop.
	MultithreadApp() bool // property
	// IsLoading
	//  Calls ICefBrowser.IsLoading and returns true if the browser is currently loading.
	IsLoading() bool // property
	// HasDocument
	//  Calls ICefBrowser.HasDocument and returns true if a document has been loaded in the browser.
	HasDocument() bool // property
	// HasView
	//  Calls ICefBrowserHost.HasView and returns true if this browser is wrapped in a ICefBrowserView.
	HasView() bool // property
	// HasDevTools
	//  Calls ICefBrowserHost.HasDevTools and returns true if this browser currently has an associated DevTools browser.
	HasDevTools() bool // property
	// HasClientHandler
	//  Returns true if ICefClient has a valid value.
	HasClientHandler() bool // property
	// HasBrowser
	//  Returns true if this component has a valid selected browser.
	HasBrowser() bool // property
	// CanGoBack
	//  Calls ICefBrowser.CanGoBack and returns true if the browser can navigate back.
	CanGoBack() bool // property
	// CanGoForward
	//  Calls ICefBrowser.CanGoForward and returns true if the browser can navigate forward.
	CanGoForward() bool // property
	// IsPopUp
	//  Calls ICefBrowser.IsPopUp and returns true if the window is a popup window.
	IsPopUp() bool // property
	// WindowHandle
	//  Calls ICefBrowserHost.GetWindowHandle and returns the window handle for this browser.
	WindowHandle() TCefWindowHandle // property
	// OpenerWindowHandle
	//  Calls ICefBrowserHost.GetOpenerWindowHandle and returns the window handle of the browser that opened this browser.
	OpenerWindowHandle() TCefWindowHandle // property
	// BrowserHandle
	//  Handle of one to the child controls created automatically by CEF to show the web contents.
	BrowserHandle() THandle // property
	// WidgetHandle
	//  Handle of one to the child controls created automatically by CEF to show the web contents.
	WidgetHandle() THandle // property
	// RenderHandle
	//  Handle of one to the child controls created automatically by CEF to show the web contents.
	RenderHandle() THandle // property
	// FrameIsFocused
	//  Returns true if ICefBrowser.FocusedFrame has a valid value.
	FrameIsFocused() bool // property
	// Initialized
	//  Returns true when the browser is fully initialized and it's not being closed.
	Initialized() bool // property
	// RequestContextCache
	//  Returns the cache value in ICefRequestContext.CachePath.
	RequestContextCache() string // property
	// RequestContextIsGlobal
	//  Calls ICefRequestContext.IsGlobal to check if the request context is the global context or it's independent.
	RequestContextIsGlobal() bool // property
	// DocumentURL
	//  Returns the URL of the main frame.
	DocumentURL() string // property
	// ZoomLevel
	//  Returns the current zoom value. This property is based on the CefBrowserHost.ZoomLevel value which can only be read in the CEF UI thread.
	ZoomLevel() (resultFloat64 float64) // property
	// SetZoomLevel Set ZoomLevel
	SetZoomLevel(AValue float64) // property
	// DefaultZoomLevel
	//  Get the default zoom level. This value will be 0.0 by default but can be
	//  configured with the Chrome runtime. This function can only be called on
	//  the CEF UI thread.
	DefaultZoomLevel() (resultFloat64 float64) // property
	// CanIncZoom
	//  Returns true(1) if this browser can execute the zoom IN command.
	//  This function can only be called on the CEF UI thread.
	CanIncZoom() bool // property
	// CanDecZoom
	//  Returns true(1) if this browser can execute the zoom OUT command.
	//  This function can only be called on the CEF UI thread.
	CanDecZoom() bool // property
	// CanResetZoom
	//  Returns true(1) if this browser can execute the zoom RESET command.
	//  This function can only be called on the CEF UI thread.
	CanResetZoom() bool // property
	// ZoomPct
	//  Returns the current zoom value. This property is based on the CefBrowserHost.ZoomLevel value which can only be read in the CEF UI thread.
	ZoomPct() (resultFloat64 float64) // property
	// SetZoomPct Set ZoomPct
	SetZoomPct(AValue float64) // property
	// ZoomStep
	//  Returns the current zoom value. This property is based on the CefBrowserHost.ZoomLevel value which can only be read in the CEF UI thread.
	ZoomStep() byte // property
	// SetZoomStep Set ZoomStep
	SetZoomStep(AValue byte) // property
	// WindowlessFrameRate
	//  Returns the maximum rate in frames per second(fps) that OnPaint will be called for a browser in OSR mode.
	WindowlessFrameRate() int32 // property
	// SetWindowlessFrameRate Set WindowlessFrameRate
	SetWindowlessFrameRate(AValue int32) // property
	// CustomHeaderName
	//  Custom HTTP header name added to all requests.
	CustomHeaderName() string // property
	// SetCustomHeaderName Set CustomHeaderName
	SetCustomHeaderName(AValue string) // property
	// CustomHeaderValue
	//  Custom HTTP header value added to all requests.
	CustomHeaderValue() string // property
	// SetCustomHeaderValue Set CustomHeaderValue
	SetCustomHeaderValue(AValue string) // property
	// DoNotTrack
	//  Set to True if you want to send the DNT header.
	DoNotTrack() bool // property
	// SetDoNotTrack Set DoNotTrack
	SetDoNotTrack(AValue bool) // property
	// SendReferrer
	//  Set to True if you want to send the referer header.
	SendReferrer() bool // property
	// SetSendReferrer Set SendReferrer
	SetSendReferrer(AValue bool) // property
	// HyperlinkAuditing
	//  Enable hyperlink auditing.
	HyperlinkAuditing() bool // property
	// SetHyperlinkAuditing Set HyperlinkAuditing
	SetHyperlinkAuditing(AValue bool) // property
	// AllowOutdatedPlugins
	//  Allow using outdated plugins.
	AllowOutdatedPlugins() bool // property
	// SetAllowOutdatedPlugins Set AllowOutdatedPlugins
	SetAllowOutdatedPlugins(AValue bool) // property
	// AlwaysAuthorizePlugins
	//  Always authorize plugins.
	AlwaysAuthorizePlugins() bool // property
	// SetAlwaysAuthorizePlugins Set AlwaysAuthorizePlugins
	SetAlwaysAuthorizePlugins(AValue bool) // property
	// AlwaysOpenPDFExternally
	//  Always open PDF files externally.
	AlwaysOpenPDFExternally() bool // property
	// SetAlwaysOpenPDFExternally Set AlwaysOpenPDFExternally
	SetAlwaysOpenPDFExternally(AValue bool) // property
	// SpellChecking
	//  Set to True if you want to enable the spell checker.
	SpellChecking() bool // property
	// SetSpellChecking Set SpellChecking
	SetSpellChecking(AValue bool) // property
	// SpellCheckerDicts
	//  Comma delimited list of language codes used by the spell checker, for example "es-ES,en-US,fr-FR,de-DE,it-IT".
	SpellCheckerDicts() string // property
	// SetSpellCheckerDicts Set SpellCheckerDicts
	SetSpellCheckerDicts(AValue string) // property
	// HasValidMainFrame
	//  Returns true if the main frame exists and it's valid.
	HasValidMainFrame() bool // property
	// FrameCount
	//  Returns the number of frames that currently exist.
	FrameCount() NativeUInt // property
	// DragOperations
	//  Returns the TcefDragOperation value used during drag and drop.
	DragOperations() TCefDragOperations // property
	// SetDragOperations Set DragOperations
	SetDragOperations(AValue TCefDragOperations) // property
	// AudioMuted
	//  Returns true if the browser's audio is muted.
	AudioMuted() bool // property
	// SetAudioMuted Set AudioMuted
	SetAudioMuted(AValue bool) // property
	// Fullscreen
	//  Returns true(1) if the renderer is currently in browser fullscreen. This
	//  differs from window fullscreen in that browser fullscreen is entered using
	//  the JavaScript Fullscreen API and modifies CSS attributes such as the
	//  ::backdrop pseudo-element and:fullscreen pseudo-structure. This property
	//  can only be called on the UI thread.
	Fullscreen() bool // property
	// SafeSearch
	//  Forces the Google safesearch in the browser preferences.
	SafeSearch() bool // property
	// SetSafeSearch Set SafeSearch
	SetSafeSearch(AValue bool) // property
	// YouTubeRestrict
	//  Forces the YouTube restrictions in the browser preferences.
	YouTubeRestrict() int32 // property
	// SetYouTubeRestrict Set YouTubeRestrict
	SetYouTubeRestrict(AValue int32) // property
	// PrintingEnabled
	//  Enables printing in the browser preferences.
	PrintingEnabled() bool // property
	// SetPrintingEnabled Set PrintingEnabled
	SetPrintingEnabled(AValue bool) // property
	// AcceptLanguageList
	//  Set the accept language list in the browser preferences.
	AcceptLanguageList() string // property
	// SetAcceptLanguageList Set AcceptLanguageList
	SetAcceptLanguageList(AValue string) // property
	// AcceptCookies
	//  Sets the cookies policy value in the browser preferences.
	AcceptCookies() TCefCookiePref // property
	// SetAcceptCookies Set AcceptCookies
	SetAcceptCookies(AValue TCefCookiePref) // property
	// Block3rdPartyCookies
	//  Blocks third party cookies in the browser preferences.
	Block3rdPartyCookies() bool // property
	// SetBlock3rdPartyCookies Set Block3rdPartyCookies
	SetBlock3rdPartyCookies(AValue bool) // property
	// MultiBrowserMode
	//  Enables the multi-browser mode that allows TChromiumCore to handle several browsers with one component. These browsers are usually the main browser, popup windows and new tabs.
	MultiBrowserMode() bool // property
	// SetMultiBrowserMode Set MultiBrowserMode
	SetMultiBrowserMode(AValue bool) // property
	// DefaultWindowInfoExStyle
	//  Default ExStyle value used to initialize the browser. A value of WS_EX_NOACTIVATE can be used as a workaround for some focus issues in CEF.
	DefaultWindowInfoExStyle() DWORD // property
	// SetDefaultWindowInfoExStyle Set DefaultWindowInfoExStyle
	SetDefaultWindowInfoExStyle(AValue DWORD) // property
	// Offline
	//  Uses the Network.emulateNetworkConditions DevTool method to set the browser in offline mode.
	Offline() bool // property
	// SetOffline Set Offline
	SetOffline(AValue bool) // property
	// QuicAllowed
	//  Enables the Quic protocol in the browser preferences.
	QuicAllowed() bool // property
	// SetQuicAllowed Set QuicAllowed
	SetQuicAllowed(AValue bool) // property
	// JavascriptEnabled
	//  Enables JavaScript in the browser preferences.
	JavascriptEnabled() bool // property
	// SetJavascriptEnabled Set JavascriptEnabled
	SetJavascriptEnabled(AValue bool) // property
	// LoadImagesAutomatically
	//  Enables automatic image loading in the browser preferences.
	LoadImagesAutomatically() bool // property
	// SetLoadImagesAutomatically Set LoadImagesAutomatically
	SetLoadImagesAutomatically(AValue bool) // property
	// BatterySaverModeState
	//  Battery saver mode state.
	BatterySaverModeState() TCefBatterySaverModeState // property
	// SetBatterySaverModeState Set BatterySaverModeState
	SetBatterySaverModeState(AValue TCefBatterySaverModeState) // property
	// HighEfficiencyModeState
	//  High efficiency mode state.
	HighEfficiencyModeState() TCefHighEfficiencyModeState // property
	// SetHighEfficiencyModeState Set HighEfficiencyModeState
	SetHighEfficiencyModeState(AValue TCefHighEfficiencyModeState) // property
	// CanFocus
	//  Indicates whether the browser can receive focus.
	CanFocus() bool // property
	// EnableFocusDelayMs
	//  Delay in milliseconds to enable browser focus.
	EnableFocusDelayMs() uint32 // property
	// SetEnableFocusDelayMs Set EnableFocusDelayMs
	SetEnableFocusDelayMs(AValue uint32) // property
	// WebRTCIPHandlingPolicy
	//  WebRTC handling policy setting in the browser preferences.
	WebRTCIPHandlingPolicy() TCefWebRTCHandlingPolicy // property
	// SetWebRTCIPHandlingPolicy Set WebRTCIPHandlingPolicy
	SetWebRTCIPHandlingPolicy(AValue TCefWebRTCHandlingPolicy) // property
	// WebRTCMultipleRoutes
	//  WebRTC multiple routes setting in the browser preferences.
	WebRTCMultipleRoutes() TCefState // property
	// SetWebRTCMultipleRoutes Set WebRTCMultipleRoutes
	SetWebRTCMultipleRoutes(AValue TCefState) // property
	// WebRTCNonproxiedUDP
	//  WebRTC nonproxied UDP setting in the browser preferences.
	WebRTCNonproxiedUDP() TCefState // property
	// SetWebRTCNonproxiedUDP Set WebRTCNonproxiedUDP
	SetWebRTCNonproxiedUDP(AValue TCefState) // property
	// ProxyType
	//  Proxy type: CEF_PROXYTYPE_DIRECT, CEF_PROXYTYPE_AUTODETECT, CEF_PROXYTYPE_SYSTEM, CEF_PROXYTYPE_FIXED_SERVERS or CEF_PROXYTYPE_PAC_SCRIPT.
	ProxyType() int32 // property
	// SetProxyType Set ProxyType
	SetProxyType(AValue int32) // property
	// ProxyScheme
	//  Proxy scheme
	ProxyScheme() TCefProxyScheme // property
	// SetProxyScheme Set ProxyScheme
	SetProxyScheme(AValue TCefProxyScheme) // property
	// ProxyServer
	//  Proxy server address
	ProxyServer() string // property
	// SetProxyServer Set ProxyServer
	SetProxyServer(AValue string) // property
	// ProxyPort
	//  Proxy server port
	ProxyPort() int32 // property
	// SetProxyPort Set ProxyPort
	SetProxyPort(AValue int32) // property
	// ProxyUsername
	//  Proxy username
	ProxyUsername() string // property
	// SetProxyUsername Set ProxyUsername
	SetProxyUsername(AValue string) // property
	// ProxyPassword
	//  Proxy password
	ProxyPassword() string // property
	// SetProxyPassword Set ProxyPassword
	SetProxyPassword(AValue string) // property
	// ProxyScriptURL
	//  URL of the PAC script file.
	ProxyScriptURL() string // property
	// SetProxyScriptURL Set ProxyScriptURL
	SetProxyScriptURL(AValue string) // property
	// ProxyByPassList
	//  This tells chromium to bypass any specified proxy for the given semi-colon-separated list of hosts.
	ProxyByPassList() string // property
	// SetProxyByPassList Set ProxyByPassList
	SetProxyByPassList(AValue string) // property
	// MaxConnectionsPerProxy
	//  Sets the maximum connections per proxy value in the browser preferences(experimental).
	MaxConnectionsPerProxy() int32 // property
	// SetMaxConnectionsPerProxy Set MaxConnectionsPerProxy
	SetMaxConnectionsPerProxy(AValue int32) // property
	// CreateClientHandler
	//  Used to create the client handler which will also create most of the browser handlers needed for the browser.
	CreateClientHandler(aIsOSR bool) bool // function
	// CreateClientHandler1
	//  Used to create the client handler when a browser requests a new browser in a popup window or tab in the TChromiumCore.OnBeforePopup event.
	CreateClientHandler1(aClient *ICefClient, aIsOSR bool) bool // function
	// TryCloseBrowser
	//  Helper for closing a browser. Call this function from the top-level window
	//  close handler(if any). Internally this calls CloseBrowser(false(0)) if
	//  the close has not yet been initiated. This function returns false(0)
	//  while the close is pending and true(1) after the close has completed. See
	//  CloseBrowser() and ICefLifeSpanHandler.DoClose() documentation for
	//  additional usage information. This function must be called on the browser
	//  process UI thread.
	TryCloseBrowser() bool // function
	// SelectBrowser
	//  Select the browser with the aID identifier when TChromiumCore uses the
	//  multi-browser mode.
	SelectBrowser(aID int32) bool // function
	// IndexOfBrowserID
	//  Returns the index in the browsers array of the browser with the aID
	//  identifier when TChromiumCore uses the multi-browser mode.
	IndexOfBrowserID(aID int32) int32 // function
	// ShareRequestContext
	//  Creates a new request context in the aContext parameter that shares
	//  storage with the request context of the current browser and uses an
	//  optional handler.
	ShareRequestContext(aContext *ICefRequestContext, aHandler ICefRequestContextHandler) bool // function
	// SetNewBrowserParent
	//  Used to reparent the browser to a different TCEFWindowParent.
	SetNewBrowserParent(aNewParentHwnd HWND) bool // function
	// CreateBrowser
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowser(aParentHandle TCefWindowHandle, aParentRect *TRect, aWindowName string, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue, aForceAsPopup bool) bool // function
	// CreateBrowser1
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowser1(aURL string, aBrowserViewComp ICEFBrowserViewComponent, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue) bool // function
	// ClearCertificateExceptions
	//  Clears all certificate exceptions that were added as part of handling
	//  OnCertificateError. If you call this it is recommended that you also call
	//  CloseAllConnections() or you risk not being prompted again for server
	//  certificates if you reconnect quickly.
	//  If aClearImmediately is false then OnCertificateExceptionsCleared is
	//  triggered when the exceptions are cleared.
	ClearCertificateExceptions(aClearImmediately bool) bool // function
	// ClearHttpAuthCredentials
	//  Clears all HTTP authentication credentials that were added as part of
	//  handling GetAuthCredentials. If |callback| is non-NULL it will be executed
	//  on the UI thread after completion.
	//  If aClearImmediately is false then OnHttpAuthCredentialsCleared is triggered
	//  when the credeintials are cleared.
	ClearHttpAuthCredentials(aClearImmediately bool) bool // function
	// CloseAllConnections
	//  Clears all active and idle connections that Chromium currently has. This
	//  is only recommended if you have released all other CEF objects but don't
	//  yet want to call cef_shutdown().
	CloseAllConnections(aCloseImmediately bool) bool // function
	// GetFrameNames
	//  Returns the names of all existing frames.
	GetFrameNames(aFrameNames *IStrings) bool // function
	// GetFrameIdentifiers
	//  Returns the identifiers of all existing frames.
	GetFrameIdentifiers(aFrameCount *NativeUInt, aFrameIdentifierArray *ICefFrameIdentifierArray) bool // function
	// IsSameBrowser
	//  Used to check if the browser parameter is the same as the selected browser in TChromiumCore.
	IsSameBrowser(aBrowser ICefBrowser) bool // function
	// ExecuteTaskOnCefThread
	//  Calling ExecuteTaskOnCefThread function will trigger the TChromiumCore.OnExecuteTaskOnCefThread event.
	//  <param name="aCefThreadId">Indicates the CEF thread on which TChromiumCore.OnExecuteTaskOnCefThread will be executed.</param>
	//  <param name="aTaskID">Custom ID used to identify the task that triggered the TChromiumCore.OnExecuteTaskOnCefThread event.</param>
	//  <param name="aDelayMs">Optional delay in milliseconds to trigger the TChromiumCore.OnExecuteTaskOnCefThread event.</param>
	ExecuteTaskOnCefThread(aCefThreadId TCefThreadId, aTaskID uint32, aDelayMs int64) bool // function
	// DeleteCookies
	//  Used to delete cookies immediately or asynchronously. If aDeleteImmediately is false TChromiumCore.DeleteCookies triggers
	//  the TChromiumCore.OnCookiesDeleted event when the cookies are deleted.
	DeleteCookies(url string, cookieName string, aDeleteImmediately bool) bool // function
	// VisitAllCookies
	//  TChromiumCore.VisitAllCookies triggers the TChromiumCore.OnCookiesVisited event for each cookie
	//  aID is an optional parameter to identify which VisitAllCookies call has triggered the
	//  OnCookiesVisited event.
	//  TChromiumCore.OnCookiesVisited may not be triggered if the cookie store is empty but the
	//  TChromium.OnCookieVisitorDestroyed event will always be triggered to signal when the browser
	//  when the visit is over.
	VisitAllCookies(aID int32) bool // function
	// VisitURLCookies
	//  TChromiumCore.VisitURLCookies triggers the TChromiumCore.OnCookiesVisited event for each cookie
	//  aID is an optional parameter to identify which VisitURLCookies call has triggered the
	//  OnCookiesVisited event.
	//  TChromiumCore.OnCookiesVisited may not be triggered if the cookie store is empty but the
	//  TChromium.OnCookieVisitorDestroyed event will always be triggered to signal when the browser
	//  when the visit is over.
	VisitURLCookies(url string, includeHttpOnly bool, aID int32) bool // function
	// FlushCookieStore
	//  Flush the backing store(if any) to disk.
	//  <param name="aFlushImmediately">If aFlushImmediately is false the cookies will be flushed on the CEF UI thread and the OnCookiesFlushed event will be triggered.</param>
	//  <returns>Returns false(0) if cookies cannot be accessed.</returns>
	FlushCookieStore(aFlushImmediately bool) bool // function
	// SendDevToolsMessage
	//  Send a function call message over the DevTools protocol. |message_| must be
	//  a UTF8-encoded JSON dictionary that contains "id"(int), "function"
	//  (string) and "params"(dictionary, optional) values. See the DevTools
	//  protocol documentation at https://chromedevtools.github.io/devtools-
	//  protocol/ for details of supported functions and the expected "params"
	//  dictionary contents. |message_| will be copied if necessary. This function
	//  will return true(1) if called on the UI thread and the message was
	//  successfully submitted for validation, otherwise false(0). Validation
	//  will be applied asynchronously and any messages that fail due to
	//  formatting errors or missing parameters may be discarded without
	//  notification. Prefer ExecuteDevToolsMethod if a more structured approach
	//  to message formatting is desired.
	//  Every valid function call will result in an asynchronous function result
	//  or error message that references the sent message "id". Event messages are
	//  received while notifications are enabled(for example, between function
	//  calls for "Page.enable" and "Page.disable"). All received messages will be
	//  delivered to the observer(s) registered with AddDevToolsMessageObserver.
	//  See ICefDevToolsMessageObserver.OnDevToolsMessage documentation for
	//  details of received message contents.
	//  Usage of the SendDevToolsMessage, ExecuteDevToolsMethod and
	//  AddDevToolsMessageObserver functions does not require an active DevTools
	//  front-end or remote-debugging session. Other active DevTools sessions will
	//  continue to function independently. However, any modification of global
	//  browser state by one session may not be reflected in the UI of other
	//  sessions.
	//  Communication with the DevTools front-end(when displayed) can be logged
	//  for development purposes by passing the `--devtools-protocol-log-
	//  file=<path>` command-line flag.
	SendDevToolsMessage(message string) bool // function
	// ExecuteDevToolsMethod
	//  Execute a function call over the DevTools protocol. This is a more
	//  structured version of SendDevToolsMessage. |message_id| is an incremental
	//  number that uniquely identifies the message(pass 0 to have the next
	//  number assigned automatically based on previous values). |function| is the
	//  function name. |params| are the function parameters, which may be NULL.
	//  See the DevTools protocol documentation(linked above) for details of
	//  supported functions and the expected |params| dictionary contents. This
	//  function will return the assigned message ID if called on the UI thread
	//  and the message was successfully submitted for validation, otherwise 0.
	//  See the SendDevToolsMessage documentation for additional usage
	//  information.
	ExecuteDevToolsMethod(messageid int32, method string, params ICefDictionaryValue) int32 // function
	// AddDevToolsMessageObserver
	//  Add an observer for DevTools protocol messages(function results and
	//  events). The observer will remain registered until the returned
	//  Registration object is destroyed. See the SendDevToolsMessage
	//  documentation for additional usage information.
	AddDevToolsMessageObserver(observer ICefDevToolsMessageObserver) ICefRegistration // function
	// CreateUrlRequest
	//  Create a new URL request that will be treated as originating from this
	//  frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if
	//  you do not want the request to have this association, in which case it may
	//  be handled differently(see documentation on that function). A request
	//  created with this function may only originate from the browser process,
	//  and will behave as follows:
	//  </code>
	//  - It may be intercepted by the client via CefResourceRequestHandler or
	//  CefSchemeHandlerFactory.
	//  - POST data may only contain a single element of type PDE_TYPE_FILE or
	//  PDE_TYPE_BYTES.
	//  </code>
	//  The |request| object will be marked as read-only after calling this
	//  function.
	CreateUrlRequest(request ICefRequest, client ICefUrlRequestClient, aFrameName string) ICefUrlRequest // function
	// CreateUrlRequest1
	//  Create a new URL request that will be treated as originating from this
	//  frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if
	//  you do not want the request to have this association, in which case it may
	//  be handled differently(see documentation on that function). A request
	//  created with this function may only originate from the browser process,
	//  and will behave as follows:
	//  <code>
	//  - It may be intercepted by the client via CefResourceRequestHandler or
	//  CefSchemeHandlerFactory.
	//  - POST data may only contain a single element of type PDE_TYPE_FILE or
	//  PDE_TYPE_BYTES.
	//  </code>
	//  The |request| object will be marked as read-only after calling this
	//  function.
	CreateUrlRequest1(request ICefRequest, client ICefUrlRequestClient, aFrame ICefFrame) ICefUrlRequest // function
	// CreateUrlRequest2
	//  Create a new URL request that will be treated as originating from this
	//  frame and the associated browser. Use TCustomCefUrlrequestClient.Create instead if
	//  you do not want the request to have this association, in which case it may
	//  be handled differently(see documentation on that function). A request
	//  created with this function may only originate from the browser process,
	//  and will behave as follows:
	//  <code>
	//  - It may be intercepted by the client via CefResourceRequestHandler or
	//  CefSchemeHandlerFactory.
	//  - POST data may only contain a single element of type PDE_TYPE_FILE or
	//  PDE_TYPE_BYTES.
	//  </code>
	//  The |request| object will be marked as read-only after calling this
	//  function.
	CreateUrlRequest2(request ICefRequest, client ICefUrlRequestClient, aFrameIdentifier int64) ICefUrlRequest // function
	// AddObserver
	//  Add an observer for MediaRouter events. The observer will remain
	//  registered until the returned Registration object is destroyed.
	AddObserver(observer ICefMediaObserver) ICefRegistration // function
	// GetSource
	//  Returns a MediaSource object for the specified media source URN. Supported
	//  URN schemes include "cast:" and "dial:", and will be already known by the
	//  client application(e.g. "cast:<appId>?clientId=<clientId>").
	GetSource(urn string) ICefMediaSource // function
	// LoadExtension
	//  Load an extension.
	//  If extension resources will be read from disk using the default load
	//  implementation then |root_directory| should be the absolute path to the
	//  extension resources directory and |manifest| should be NULL. If extension
	//  resources will be provided by the client(e.g. via cef_request_handler_t
	//  and/or cef_extension_handler_t) then |root_directory| should be a path
	//  component unique to the extension(if not absolute this will be internally
	//  prefixed with the PK_DIR_RESOURCES path) and |manifest| should contain the
	//  contents that would otherwise be read from the "manifest.json" file on
	//  disk.
	//  The loaded extension will be accessible in all contexts sharing the same
	//  storage(HasExtension returns true(1)). However, only the context on
	//  which this function was called is considered the loader(DidLoadExtension
	//  returns true(1)) and only the loader will receive
	//  TCustomRequestContextHandler callbacks for the extension.
	//  TCustomExtensionHandler.OnExtensionLoaded will be called on load success
	//  or TCustomExtensionHandler.OnExtensionLoadFailed will be called on load
	//  failure.
	//  If the extension specifies a background script via the "background"
	//  manifest key then TCustomExtensionHandler.OnBeforeBackgroundBrowser will
	//  be called to create the background browser. See that function for
	//  additional information about background scripts.
	//  For visible extension views the client application should evaluate the
	//  manifest to determine the correct extension URL to load and then pass that
	//  URL to the ICefBrowserHost.CreateBrowser* function after the extension
	//  has loaded. For example, the client can look for the "browser_action"
	//  manifest key as documented at
	//  https://developer.chrome.com/extensions/browserAction. Extension URLs take
	//  the form "chrome-extension://<extension_id>/<path>".
	//  Browsers that host extensions differ from normal browsers as follows:
	//  <code>
	//  - Can access chrome.* JavaScript APIs if allowed by the manifest. Visit
	//  chrome://extensions-support for the list of extension APIs currently
	//  supported by CEF.
	//  - Main frame navigation to non-extension content is blocked.
	//  - Pinch-zooming is disabled.
	//  - CefBrowserHost::GetExtension returns the hosted extension.
	//  - CefBrowserHost::IsBackgroundHost returns true for background hosts.
	//  </code>
	//  See https://developer.chrome.com/extensions for extension implementation
	//  and usage documentation.
	LoadExtension(rootdirectory string, manifest ICefDictionaryValue, handler ICefExtensionHandler, requestContext ICefRequestContext) bool // function
	// DidLoadExtension
	//  Returns true(1) if this context was used to load the extension identified
	//  by |extension_id|. Other contexts sharing the same storage will also have
	//  access to the extension(see HasExtension). This function must be called
	//  on the browser process UI thread.
	DidLoadExtension(extensionid string) bool // function
	// HasExtension
	//  Returns true(1) if this context has access to the extension identified by
	//  |extension_id|. This may not be the context that was used to load the
	//  extension(see DidLoadExtension). This function must be called on the
	//  browser process UI thread.
	HasExtension(extensionid string) bool // function
	// GetExtensions
	//  Retrieve the list of all extensions that this context has access to(see
	//  HasExtension). |extension_ids| will be populated with the list of
	//  extension ID values. Returns true(1) on success. This function must be
	//  called on the browser process UI thread.
	GetExtensions(extensionids IStringList) bool // function
	// GetExtension
	//  Returns the extension matching |extension_id| or NULL if no matching
	//  extension is accessible in this context(see HasExtension). This function
	//  must be called on the browser process UI thread.
	GetExtension(extensionid string) ICefExtension // function
	// GetWebsiteSetting
	//  Returns the current value for |content_type| that applies for the
	//  specified URLs. If both URLs are NULL the default value will be returned.
	//  Returns nullptr if no value is configured. Must be called on the browser
	//  process UI thread.
	GetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) ICefValue // function
	// GetContentSetting
	//  Returns the current value for |content_type| that applies for the
	//  specified URLs. If both URLs are NULL the default value will be returned.
	//  Returns CEF_CONTENT_SETTING_VALUE_DEFAULT if no value is configured. Must
	//  be called on the browser process UI thread.
	GetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) TCefContentSettingValues // function
	// CloseBrowser
	//  Request that the browser close. The JavaScript 'onbeforeunload' event will
	//  be fired. If |aForceClose| is false(0) the event handler, if any, will be
	//  allowed to prompt the user and the user can optionally cancel the close.
	//  If |aForceClose| is true(1) the prompt will not be displayed and the
	//  close will proceed. Results in a call to
	//  ICefLifeSpanHandler.DoClose() if the event handler allows the close
	//  or if |aForceClose| is true(1). See ICefLifeSpanHandler.DoClose()
	//  documentation for additional usage information.
	CloseBrowser(aForceClose bool) // procedure
	// CloseAllBrowsers
	//  Calls CloseBrowser for all the browsers handled by this TChromiumCore instance.
	CloseAllBrowsers() // procedure
	// InitializeDragAndDrop
	//  Used with browsers in OSR mode to initialize drag and drop in Windows.
	InitializeDragAndDrop(aDropTargetWnd HWND) // procedure
	// ShutdownDragAndDrop
	//  Used with browsers in OSR mode to shutdown drag and drop in Windows.
	ShutdownDragAndDrop() // procedure
	// LoadURL
	//  Used to navigate to a URL in the specified frame or the main frame.
	LoadURL(aURL string, aFrameName string) // procedure
	// LoadURL1
	//  Used to navigate to a URL in the specified frame or the main frame.
	LoadURL1(aURL string, aFrame ICefFrame) // procedure
	// LoadURL2
	//  Used to navigate to a URL in the specified frame or the main frame.
	LoadURL2(aURL string, aFrameIdentifier int64) // procedure
	// LoadString
	//  Used to load a DATA URI with the HTML string contents in the specified frame or the main frame.
	LoadString(aHTML string, aFrameName string) // procedure
	// LoadString1
	//  Used to load a DATA URI with the HTML string contents in the specified frame or the main frame.
	LoadString1(aHTML string, aFrame ICefFrame) // procedure
	// LoadString2
	//  Used to load a DATA URI with the HTML string contents in the specified frame or the main frame.
	LoadString2(aHTML string, aFrameIdentifier int64) // procedure
	// LoadResource
	//  Used to load a DATA URI with the stream contents in the specified frame or the main frame.
	//  The DATA URI will be configured with the mime type and charset specified in the parameters.
	LoadResource(aStream ICustomMemoryStream, aMimeType, aCharset string, aFrameName string) // procedure
	// LoadResource1
	//  Used to load a DATA URI with the stream contents in the specified frame or the main frame.
	//  The DATA URI will be configured with the mime type and charset specified in the parameters.
	LoadResource1(aStream ICustomMemoryStream, aMimeType, aCharset string, aFrame ICefFrame) // procedure
	// LoadResource2
	//  Used to load a DATA URI with the stream contents in the specified frame or the main frame.
	//  The DATA URI will be configured with the mime type and charset specified in the parameters.
	LoadResource2(aStream ICustomMemoryStream, aMimeType, aCharset string, aFrameIdentifier int64) // procedure
	// LoadRequest
	//  Load the request represented by the aRequest object.
	//  WARNING: This function will fail with bad IPC message reason
	//  INVALID_INITIATOR_ORIGIN(213) unless you first navigate to the request
	//  origin using some other mechanism(LoadURL, link click, etc).
	LoadRequest(aRequest ICefRequest) // procedure
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
	// StartDownload
	//  Starts downloading a file in the specified URL.
	StartDownload(aURL string) // procedure
	// DownloadImage
	//  Starts downloading an image in the specified URL.
	//  Use the TChromiumCore.OnDownloadImageFinished event to receive the image.
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool) // procedure
	// SimulateMouseWheel
	//  Calls ICefBrowserHost.SendMouseWheelEvent to simulate a simple mouse wheel event.
	//  Use TChromiumCore.SendMouseWheelEvent if you need to specify the mouse coordinates or the event flags.
	SimulateMouseWheel(aDeltaX, aDeltaY int32) // procedure
	// RetrieveHTML
	//  Retrieve all the HTML content from the specified frame or the main frame.
	//  Leave aFrameName empty to get the HTML source from the main frame.
	//  It uses a CefStringVisitor to get the HTML content asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveHTML(aFrameName string) // procedure
	// RetrieveHTML1
	//  Retrieve all the HTML content from the specified frame or the main frame.
	//  Set aFrame to nil to get the HTML source from the main frame.
	//  It uses a CefStringVisitor to get the HTML content asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveHTML1(aFrame ICefFrame) // procedure
	// RetrieveHTML2
	//  Retrieve all the HTML content from the specified frame or the main frame.
	//  Set aFrameIdentifier to zero to get the HTML source from the main frame.
	//  It uses a CefStringVisitor to get the HTML content asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveHTML2(aFrameIdentifier int64) // procedure
	// RetrieveText
	//  Retrieve all the text content from the specified frame or the main frame.
	//  Leave aFrameName empty to get the text from the main frame.
	//  It uses a CefStringVisitor to get the text asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveText(aFrameName string) // procedure
	// RetrieveText1
	//  Retrieve all the text content from the specified frame or the main frame.
	//  Set aFrame to nil to get the text from the main frame.
	//  It uses a CefStringVisitor to get the text asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveText1(aFrame ICefFrame) // procedure
	// RetrieveText2
	//  Retrieve all the text content from the specified frame or the main frame.
	//  Set aFrameIdentifier to zero to get the text from the main frame.
	//  It uses a CefStringVisitor to get the text asynchronously and the
	//  result will be received in the TChromiumCore.OnTextResultAvailable event.
	RetrieveText2(aFrameIdentifier int64) // procedure
	// GetNavigationEntries
	//  Retrieve a snapshot of current navigation entries asynchronously. The
	//  TChromiumCore.OnNavigationVisitorResultAvailable event will be triggered
	//  for each navigation entry.
	GetNavigationEntries(currentOnly bool) // procedure
	// ExecuteJavaScript
	//  Execute a string of JavaScript code in this frame.
	//  <param name="aCode">JavaScript code.</param>
	//  <param name="aScriptURL">The URL where the script in question can be found, if any. The renderer may request this URL to show the developer the source of the error.</param>
	//  <param name="aFrameName">Name of the frame where the JavaScript code will be executed. This name is generated automatically by Chromium. See ICefBrowser.GetFrameNames.</param>
	//  <param name="aStartLine">The base line number to use for error reporting.</param>
	ExecuteJavaScript(aCode, aScriptURL string, aFrameName string, aStartLine int32) // procedure
	// ExecuteJavaScript1
	//  Execute a string of JavaScript code in this frame.
	//  <param name="aCode">JavaScript code.</param>
	//  <param name="aScriptURL">The URL where the script in question can be found, if any. The renderer may request this URL to show the developer the source of the error.</param>
	//  <param name="aFrame">Frame where the JavaScript code will be executed.</param>
	//  <param name="aStartLine">The base line number to use for error reporting.</param>
	ExecuteJavaScript1(aCode, aScriptURL string, aFrame ICefFrame, aStartLine int32) // procedure
	// ExecuteJavaScript2
	//  Execute a string of JavaScript code in this frame.
	//  <param name="aCode">JavaScript code.</param>
	//  <param name="aScriptURL">The URL where the script in question can be found, if any. The renderer may request this URL to show the developer the source of the error.</param>
	//  <param name="aFrameIdentifier">Frame where the JavaScript code will be executed.</param>
	//  <param name="aStartLine">The base line number to use for error reporting.</param>
	ExecuteJavaScript2(aCode, aScriptURL string, aFrameIdentifier int64, aStartLine int32) // procedure
	// UpdatePreferences
	//  Used to update the browser preferences using the TChromiumCore property values asynchronously.
	UpdatePreferences() // procedure
	// SavePreferences
	//  Save the browser preferences as a text file.
	SavePreferences(aFileName string) // procedure
	// ResolveHost
	//  Calls CefRequestContext.ResolveHost to resolve the domain in the URL parameter
	//  to a list of IP addresses.
	//  The result will be received in the TChromiumCore.OnResolvedHostAvailable event.
	ResolveHost(aURL string) // procedure
	// SetUserAgentOverride
	//  This procedure calls the Emulation.setUserAgentOverride DevTools method to override the user agent string.
	SetUserAgentOverride(aUserAgent string, aAcceptLanguage string, aPlatform string) // procedure
	// ClearDataForOrigin
	//  This procedure calls the Storage.clearDataForOrigin DevTools method to clear the storage data for a given origin.
	ClearDataForOrigin(aOrigin string, aStorageTypes TCefClearDataStorageTypes) // procedure
	// ClearCache
	//  This procedure calls the Network.clearBrowserCache DevTools method to clear the cache data.
	ClearCache() // procedure
	// ToggleAudioMuted
	//  Enable or disable the browser's audio.
	ToggleAudioMuted() // procedure
	// ShowDevTools
	//  Open developer tools(DevTools) in its own browser. If inspectElementAt has a valid point
	//  with coordinates different than low(integer) then the element at the specified location
	//  will be inspected. If the DevTools browser is already open then it will be focused.
	ShowDevTools(inspectElementAt *TPoint, aWindowInfo *TCefWindowInfo) // procedure
	// CloseDevTools
	//  close the developer tools.
	CloseDevTools() // procedure
	// CloseDevTools1
	//  close the developer tools.
	CloseDevTools1(aDevToolsWnd TCefWindowHandle) // procedure
	// Find
	//  Search for |searchText|. |forward| indicates whether to search forward or
	//  backward within the page. |matchCase| indicates whether the search should
	//  be case-sensitive. |findNext| indicates whether this is the first request
	//  or a follow-up. The search will be restarted if |searchText| or
	//  |matchCase| change. The search will be stopped if |searchText| is NULL.
	//  OnFindResult will be triggered to report find results.
	Find(aSearchText string, aForward, aMatchCase, aFindNext bool) // procedure
	// StopFinding
	//  Cancel all searches that are currently going on.
	StopFinding(aClearSelection bool) // procedure
	// Print
	//  Print the current browser contents.
	Print() // procedure
	// PrintToPDF
	//  Print the current browser contents to the PDF file specified by |path| and
	//  execute |callback| on completion. The caller is responsible for deleting
	//  |path| when done. For PDF printing to work on Linux you must implement the
	//  ICefPrintHandler.GetPdfPaperSize function.
	//  The TChromiumCore.OnPdfPrintFinished event will be triggered when the PDF
	//  file is created.
	PrintToPDF(aFilePath string) // procedure
	// ClipboardCopy
	//  Execute copy on the focused frame.
	ClipboardCopy() // procedure
	// ClipboardPaste
	//  Execute paste on the focused frame.
	ClipboardPaste() // procedure
	// ClipboardCut
	//  Execute cut on the focused frame.
	ClipboardCut() // procedure
	// ClipboardUndo
	//  Execute undo on the focused frame.
	ClipboardUndo() // procedure
	// ClipboardRedo
	//  Execute redo on the focused frame.
	ClipboardRedo() // procedure
	// ClipboardDel
	//  Execute delete on the focused frame.
	ClipboardDel() // procedure
	// SelectAll
	//  Execute select all on the focused frame.
	SelectAll() // procedure
	// IncZoomStep
	//  Increase the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	IncZoomStep() // procedure
	// DecZoomStep
	//  Decrease the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	DecZoomStep() // procedure
	// IncZoomPct
	//  Increase the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	IncZoomPct() // procedure
	// DecZoomPct
	//  Decrease the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	DecZoomPct() // procedure
	// ResetZoomStep
	//  Reset the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	ResetZoomStep() // procedure
	// ResetZoomLevel
	//  Reset the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	ResetZoomLevel() // procedure
	// ResetZoomPct
	//  Reset the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the new zoom value.
	ResetZoomPct() // procedure
	// ReadZoom
	//  Read the zoom value. This procedure triggers the TChromium.OnZoomPctAvailable event with the zoom value.
	ReadZoom() // procedure
	// IncZoomCommand
	//  Execute a zoom IN command in this browser. If called on the CEF UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the CEF UI thread.
	IncZoomCommand() // procedure
	// DecZoomCommand
	//  Execute a zoom OUT command in this browser. If called on the CEF UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the CEF UI thread.
	DecZoomCommand() // procedure
	// ResetZoomCommand
	//  Execute a zoom RESET command in this browser. If called on the CEF UI thread the
	//  change will be applied immediately. Otherwise, the change will be applied
	//  asynchronously on the CEF UI thread.
	ResetZoomCommand() // procedure
	// WasResized
	//  Notify the browser that the widget has been resized. The browser will
	//  first call ICefRenderHandler.GetViewRect to get the new size and then
	//  call ICefRenderHandler.OnPaint asynchronously with the updated
	//  regions. This function is only used when window rendering is disabled.
	WasResized() // procedure
	// WasHidden
	//  Notify the browser that it has been hidden or shown. Layouting and
	//  ICefRenderHandler.OnPaint notification will stop when the browser is
	//  hidden. This function is only used when window rendering is disabled.
	WasHidden(hidden bool) // procedure
	// NotifyScreenInfoChanged
	//  Send a notification to the browser that the screen info has changed. The
	//  browser will then call ICefRenderHandler.GetScreenInfo to update the
	//  screen information with the new values. This simulates moving the webview
	//  window from one display to another, or changing the properties of the
	//  current display. This function is only used when window rendering is
	//  disabled.
	NotifyScreenInfoChanged() // procedure
	// NotifyMoveOrResizeStarted
	//  Notify the browser that the window hosting it is about to be moved or
	//  resized. This function is only used on Windows and Linux.
	NotifyMoveOrResizeStarted() // procedure
	// Invalidate
	//  Invalidate the view. The browser will call ICefRenderHandler.OnPaint
	//  asynchronously. This function is only used when window rendering is
	//  disabled.
	Invalidate(type_ TCefPaintElementType) // procedure
	// ExitFullscreen
	//  Requests the renderer to exit browser fullscreen. In most cases exiting
	//  window fullscreen should also exit browser fullscreen. With the Alloy
	//  runtime this function should be called in response to a user action such
	//  as clicking the green traffic light button on MacOS
	//  (ICefWindowDelegate.OnWindowFullscreenTransition callback) or pressing
	//  the "ESC" key(ICefKeyboardHandler.OnPreKeyEvent callback). With the
	//  Chrome runtime these standard exit actions are handled internally but
	//  new/additional user actions can use this function. Set |will_cause_resize|
	//  to true(1) if exiting browser fullscreen will cause a view resize.
	ExitFullscreen(willcauseresize bool) // procedure
	// SendExternalBeginFrame
	//  Issue a BeginFrame request to Chromium. Only valid when
	//  TCefWindowInfo.external_begin_frame_enabled is set to true(1).
	SendExternalBeginFrame() // procedure
	// SendKeyEvent
	//  Send a key event to the browser.
	SendKeyEvent(event *TCefKeyEvent) // procedure
	// SendMouseClickEvent
	//  Send a mouse click event to the browser. The |x| and |y| coordinates are
	//  relative to the upper-left corner of the view.
	SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) // procedure
	// SendMouseMoveEvent
	//  Send a mouse move event to the browser. The |x| and |y| coordinates are
	//  relative to the upper-left corner of the view.
	SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) // procedure
	// SendMouseWheelEvent
	//  Send a mouse wheel event to the browser. The |x| and |y| coordinates are
	//  relative to the upper-left corner of the view. The |deltaX| and |deltaY|
	//  values represent the movement delta in the X and Y directions
	//  respectively. In order to scroll inside select popups with window
	//  rendering disabled ICefRenderHandler.GetScreenPoint should be
	//  implemented properly.
	SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) // procedure
	// SendTouchEvent
	//  Send a touch event to the browser for a windowless browser.
	SendTouchEvent(event *TCefTouchEvent) // procedure
	// SendCaptureLostEvent
	//  Send a capture lost event to the browser.
	SendCaptureLostEvent() // procedure
	// SendProcessMessage
	//  Send a message to the specified |targetProcess|. Ownership of the message
	//  contents will be transferred and the |ProcMessage| reference will be
	//  invalidated. Message delivery is not guaranteed in all cases(for example,
	//  if the browser is closing, navigating, or if the target process crashes).
	//  Send an ACK message back from the target process if confirmation is
	//  required.
	SendProcessMessage(targetProcess TCefProcessId, procMessage ICefProcessMessage, aFrameName string) // procedure
	// SendProcessMessage1
	//  Send a message to the specified |targetProcess|. Ownership of the message
	//  contents will be transferred and the |ProcMessage| reference will be
	//  invalidated. Message delivery is not guaranteed in all cases(for example,
	//  if the browser is closing, navigating, or if the target process crashes).
	//  Send an ACK message back from the target process if confirmation is
	//  required.
	SendProcessMessage1(targetProcess TCefProcessId, procMessage ICefProcessMessage, aFrame ICefFrame) // procedure
	// SendProcessMessage2
	//  Send a message to the specified |targetProcess|. Ownership of the message
	//  contents will be transferred and the |ProcMessage| reference will be
	//  invalidated. Message delivery is not guaranteed in all cases(for example,
	//  if the browser is closing, navigating, or if the target process crashes).
	//  Send an ACK message back from the target process if confirmation is
	//  required.
	SendProcessMessage2(targetProcess TCefProcessId, procMessage ICefProcessMessage, aFrameIdentifier int64) // procedure
	// SetFocus
	//  Set whether the browser is focused.
	SetFocus(focus bool) // procedure
	// SetAccessibilityState
	//  Set accessibility state for all frames. |accessibility_state| may be
	//  default, enabled or disabled. If |accessibility_state| is STATE_DEFAULT
	//  then accessibility will be disabled by default and the state may be
	//  further controlled with the "force-renderer-accessibility" and "disable-
	//  renderer-accessibility" command-line switches. If |accessibility_state| is
	//  STATE_ENABLED then accessibility will be enabled. If |accessibility_state|
	//  is STATE_DISABLED then accessibility will be completely disabled.
	//  For windowed browsers accessibility will be enabled in Complete mode
	//  (which corresponds to kAccessibilityModeComplete in Chromium). In this
	//  mode all platform accessibility objects will be created and managed by
	//  Chromium's internal implementation. The client needs only to detect the
	//  screen reader and call this function appropriately. For example, on macOS
	//  the client can handle the @"AXEnhancedUserStructure" accessibility
	//  attribute to detect VoiceOver state changes and on Windows the client can
	//  handle WM_GETOBJECT with OBJID_CLIENT to detect accessibility readers.
	//  For windowless browsers accessibility will be enabled in TreeOnly mode
	//  (which corresponds to kAccessibilityModeWebContentsOnly in Chromium). In
	//  this mode renderer accessibility is enabled, the full tree is computed,
	//  and events are passed to CefAccessibiltyHandler, but platform
	//  accessibility objects are not created. The client may implement platform
	//  accessibility objects using CefAccessibiltyHandler callbacks if desired.
	SetAccessibilityState(accessibilityState TCefState) // procedure
	// DragTargetDragEnter
	//  Call this function when the user drags the mouse into the web view(before
	//  calling DragTargetDragOver/DragTargetLeave/DragTargetDrop). |drag_data|
	//  should not contain file contents as this type of data is not allowed to be
	//  dragged into the web view. File contents can be removed using
	//  ICefDragData.ResetFileContents(for example, if |drag_data| comes from
	//  ICefRenderHandler.StartDragging). This function is only used when
	//  window rendering is disabled.
	DragTargetDragEnter(dragData ICefDragData, event *TCefMouseEvent, allowedOps TCefDragOperations) // procedure
	// DragTargetDragOver
	//  Call this function each time the mouse is moved across the web view during
	//  a drag operation(after calling DragTargetDragEnter and before calling
	//  DragTargetDragLeave/DragTargetDrop). This function is only used when
	//  window rendering is disabled.
	DragTargetDragOver(event *TCefMouseEvent, allowedOps TCefDragOperations) // procedure
	// DragTargetDragLeave
	//  Call this function when the user drags the mouse out of the web view
	//  (after calling DragTargetDragEnter). This function is only used when
	//  window rendering is disabled.
	DragTargetDragLeave() // procedure
	// DragTargetDrop
	//  Call this function when the user completes the drag operation by dropping
	//  the object onto the web view(after calling DragTargetDragEnter). The
	//  object being dropped is |drag_data|, given as an argument to the previous
	//  DragTargetDragEnter call. This function is only used when window rendering
	//  is disabled.
	DragTargetDrop(event *TCefMouseEvent) // procedure
	// DragSourceEndedAt
	//  Call this function when the drag operation started by a
	//  ICefRenderHandler.StartDragging call has ended either in a drop or by
	//  being cancelled. |x| and |y| are mouse coordinates relative to the upper-
	//  left corner of the view. If the web view is both the drag source and the
	//  drag target then all DragTarget* functions should be called before
	//  DragSource* mthods. This function is only used when window rendering is
	//  disabled.
	DragSourceEndedAt(x, y int32, op TCefDragOperation) // procedure
	// DragSourceSystemDragEnded
	//  Call this function when the drag operation started by a
	//  ICefRenderHandler.StartDragging call has completed. This function may
	//  be called immediately without first calling DragSourceEndedAt to cancel a
	//  drag operation. If the web view is both the drag source and the drag
	//  target then all DragTarget* functions should be called before DragSource*
	//  mthods. This function is only used when window rendering is disabled.
	DragSourceSystemDragEnded() // procedure
	// IMECommitText
	//  Completes the existing composition by optionally inserting the specified
	//  |text| into the composition node. |replacement_range| is an optional range
	//  of the existing text that will be replaced. |relative_cursor_pos| is where
	//  the cursor will be positioned relative to the current cursor position. See
	//  comments on ImeSetComposition for usage. The |replacement_range| and
	//  |relative_cursor_pos| values are only used on OS X. This function is only
	//  used when window rendering is disabled.
	IMECommitText(text string, replacementrange *TCefRange, relativecursorpos int32) // procedure
	// IMEFinishComposingText
	//  Completes the existing composition by applying the current composition
	//  node contents. If |keep_selection| is false(0) the current selection, if
	//  any, will be discarded. See comments on ImeSetComposition for usage. This
	//  function is only used when window rendering is disabled.
	IMEFinishComposingText(keepselection bool) // procedure
	// IMECancelComposition
	//  Cancels the existing composition and discards the composition node
	//  contents without applying them. See comments on ImeSetComposition for
	//  usage. This function is only used when window rendering is disabled.
	IMECancelComposition() // procedure
	// ReplaceMisspelling
	//  If a misspelled word is currently selected in an editable node calling
	//  this function will replace it with the specified |word|.
	ReplaceMisspelling(aWord string) // procedure
	// AddWordToDictionary
	//  Add the specified |word| to the spelling dictionary.
	AddWordToDictionary(aWord string) // procedure
	// UpdateBrowserSize
	//  Used in Linux to resize the browser contents.
	UpdateBrowserSize(aLeft, aTop, aWidth, aHeight int32) // procedure
	// UpdateXWindowVisibility
	//  Used in Linux to update the browser visibility.
	UpdateXWindowVisibility(aVisible bool) // procedure
	// NotifyCurrentSinks
	//  Trigger an asynchronous call to ICefMediaObserver.OnSinks on all
	//  registered observers.
	NotifyCurrentSinks() // procedure
	// NotifyCurrentRoutes
	//  Trigger an asynchronous call to ICefMediaObserver.OnRoutes on all
	//  registered observers.
	NotifyCurrentRoutes() // procedure
	// CreateRoute
	//  Create a new route between |source| and |sink|. Source and sink must be
	//  valid, compatible(as reported by ICefMediaSink.IsCompatibleWith), and
	//  a route between them must not already exist. |callback| will be executed
	//  on success or failure. If route creation succeeds it will also trigger an
	//  asynchronous call to ICefMediaObserver.OnRoutes on all registered
	//  observers.
	//  This procedure is asynchronous and the result, ICefMediaRoute and the error
	//  message will be available in the TChromium.OnMediaRouteCreateFinished event.
	CreateRoute(source ICefMediaSource, sink ICefMediaSink) // procedure
	// GetDeviceInfo
	//  Asynchronously retrieves device info.
	//  This procedure will trigger OnMediaSinkDeviceInfo with the device info.
	GetDeviceInfo(aMediaSink ICefMediaSink) // procedure
	// SetWebsiteSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass nullptr for |value| to remove the
	//  default value for this content type.
	//  WARNING: Incorrect usage of this function may cause instability or
	//  security issues in Chromium. Make sure that you first understand the
	//  potential impact of any changes to |content_type| by reviewing the related
	//  source code in Chromium. For example, if you plan to modify
	//  CEF_CONTENT_SETTING_TYPE_POPUPS, first review and understand the usage of
	//  ContentSettingsType::POPUPS in Chromium:
	//  https://source.chromium.org/search?q=ContentSettingsType::POPUPS
	SetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value ICefValue) // procedure
	// SetContentSetting
	//  Sets the current value for |content_type| for the specified URLs in the
	//  default scope. If both URLs are NULL, and the context is not incognito,
	//  the default value will be set. Pass CEF_CONTENT_SETTING_VALUE_DEFAULT for
	//  |value| to use the default value for this content type.
	//  WARNING: Incorrect usage of this function may cause instability or
	//  security issues in Chromium. Make sure that you first understand the
	//  potential impact of any changes to |content_type| by reviewing the related
	//  source code in Chromium. For example, if you plan to modify
	//  CEF_CONTENT_SETTING_TYPE_POPUPS, first review and understand the usage of
	//  ContentSettingsType::POPUPS in Chromium:
	//  https://source.chromium.org/search?q=ContentSettingsType::POPUPS
	SetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value TCefContentSettingValues) // procedure
	// SetOnTextResultAvailable
	//  Triggered after a TChromiumCore.RetrieveHTML or TChromiumCore.RetrieveText call with the HTML or text results.
	SetOnTextResultAvailable(fn TOnTextResultAvailable) // property event
	// SetOnPdfPrintFinished
	//  Triggered after a TChromiumCore.PrintToPDF call when the PDF has been created.
	SetOnPdfPrintFinished(fn TOnPdfPrintFinished) // property event
	// SetOnPrefsAvailable
	//  Triggered after a TChromiumCore.SavePreferences call when the preferences have been saved.
	SetOnPrefsAvailable(fn TOnPrefsAvailable) // property event
	// SetOnPrefsUpdated
	//  Triggered when the browser preferences have been updated.
	SetOnPrefsUpdated(fn TNotify) // property event
	// SetOnCookiesDeleted
	//  Triggered after a TChromiumCore.DeleteCookies call when the cookies have been deleted.
	SetOnCookiesDeleted(fn TOnCookiesDeleted) // property event
	// SetOnResolvedHostAvailable
	//  Triggered after a TChromiumCore.ResolveHost call with the host information.
	SetOnResolvedHostAvailable(fn TOnResolvedIPsAvailable) // property event
	// SetOnNavigationVisitorResultAvailable
	//  Triggered after a TChromiumCore.GetNavigationEntries call with a navigation entry.
	SetOnNavigationVisitorResultAvailable(fn TOnNavigationVisitorResultAvailable) // property event
	// SetOnDownloadImageFinished
	//  Triggered after a TChromiumCore.DownloadImage call when the download is complete.
	SetOnDownloadImageFinished(fn TOnDownloadImageFinished) // property event
	// SetOnCookiesFlushed
	//  Triggered after a TChromiumCore.FlushCookieStore call when the cookies are flushed.
	SetOnCookiesFlushed(fn TNotify) // property event
	// SetOnCertificateExceptionsCleared
	//  Triggered after a TChromiumCore.ClearCertificateExceptions call when the exceptions are cleared.
	SetOnCertificateExceptionsCleared(fn TNotify) // property event
	// SetOnHttpAuthCredentialsCleared
	//  Triggered after a TChromiumCore.ClearHttpAuthCredentials call when the credentials are cleared.
	SetOnHttpAuthCredentialsCleared(fn TNotify) // property event
	// SetOnAllConnectionsClosed
	//  Triggered after a TChromiumCore.CloseAllConnections call when the connections are closed.
	SetOnAllConnectionsClosed(fn TNotify) // property event
	// SetOnExecuteTaskOnCefThread
	//  Triggered after a TChromiumCore.ExecuteTaskOnCefThread call in the context of the specified CEF thread.
	SetOnExecuteTaskOnCefThread(fn TOnExecuteTaskOnCefThread) // property event
	// SetOnCookiesVisited
	//  Triggered after a TChromiumCore.VisitAllCookies call with cookie information.
	SetOnCookiesVisited(fn TOnCookiesVisited) // property event
	// SetOnCookieVisitorDestroyed
	//  Triggered after a TChromiumCore.VisitAllCookies call when the IcefCookieVisitor has been destroyed.
	SetOnCookieVisitorDestroyed(fn TOnCookieVisitorDestroyed) // property event
	// SetOnCookieSet
	//  Triggered after a TChromiumCore.SetCookie call when the cookie has been set.
	SetOnCookieSet(fn TOnCookieSet) // property event
	// SetOnZoomPctAvailable
	//  Triggered after a call to any of the procedures to increase, decrease or reset the zoom with the new zoom value.
	SetOnZoomPctAvailable(fn TOnZoomPctAvailable) // property event
	// SetOnMediaRouteCreateFinished
	//  Triggered after a TChromiumCore.CreateRoute call when the route is created.
	SetOnMediaRouteCreateFinished(fn TOnMediaRouteCreateFinished) // property event
	// SetOnMediaSinkDeviceInfo
	//  Triggered after a TChromiumCore.GetDeviceInfo call with the device info.
	SetOnMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfo) // property event
	// SetOnCanFocus
	//  Triggered when the browser is capable of being focused.
	SetOnCanFocus(fn TNotify) // property event
	// SetOnBrowserCompMsg
	//  Triggered for all messages sent to the child controls created by CEF to show the web contents.
	SetOnBrowserCompMsg(fn TOnCompMsg) // property event
	// SetOnWidgetCompMsg
	//  Triggered for all messages sent to the child controls created by CEF to show the web contents.
	SetOnWidgetCompMsg(fn TOnCompMsg) // property event
	// SetOnRenderCompMsg
	//  Triggered for all messages sent to the child controls created by CEF to show the web contents.
	SetOnRenderCompMsg(fn TOnCompMsg) // property event
	// SetOnProcessMessageReceived
	//  Called when a new message is received from a different process. Return
	//  true(1) if the message was handled or false(0) otherwise. It is safe to
	//  keep a reference to |message| outside of this callback.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_client_capi.h">CEF source file: /include/capi/cef_client_capi.h(cef_client_t)</a>
	SetOnProcessMessageReceived(fn TOnProcessMessageReceived) // property event
	// SetOnLoadStart
	//  Called after a navigation has been committed and before the browser begins
	//  loading contents in the frame. The |frame| value will never be NULL --
	//  call the IsMain() function to check if this frame is the main frame.
	//  |transition_type| provides information about the source of the navigation
	//  and an accurate value is only available in the browser process. Multiple
	//  frames may be loading at the same time. Sub-frames may start or continue
	//  loading after the main frame load has ended. This function will not be
	//  called for same page navigations(fragments, history state, etc.) or for
	//  navigations that fail or are canceled before commit. For notification of
	//  overall browser load status use OnLoadingStateChange instead.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadStart(fn TOnLoadStart) // property event
	// SetOnLoadEnd
	//  Called when the browser is done loading a frame. The |frame| value will
	//  never be NULL -- call the IsMain() function to check if this frame is the
	//  main frame. Multiple frames may be loading at the same time. Sub-frames
	//  may start or continue loading after the main frame load has ended. This
	//  function will not be called for same page navigations(fragments, history
	//  state, etc.) or for navigations that fail or are canceled before commit.
	//  For notification of overall browser load status use OnLoadingStateChange
	//  instead.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadEnd(fn TOnLoadEnd) // property event
	// SetOnLoadError
	//  Called when a navigation fails or is canceled. This function may be called
	//  by itself if before commit or in combination with OnLoadStart/OnLoadEnd if
	//  after commit. |errorCode| is the error code number, |errorText| is the
	//  error text and |failedUrl| is the URL that failed to load. See
	//  net\base\net_error_list.h for complete descriptions of the error codes.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadError(fn TOnLoadError) // property event
	// SetOnLoadingStateChange
	//  Called when the loading state has changed. This callback will be executed
	//  twice -- once when loading is initiated either programmatically or by user
	//  action, and once when loading is terminated due to completion,
	//  cancellation of failure. It will be called before any calls to OnLoadStart
	//  and after all calls to OnLoadError and/or OnLoadEnd.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h(cef_load_handler_t)</a>
	SetOnLoadingStateChange(fn TOnLoadingStateChange) // property event
	// SetOnTakeFocus
	//  Called when the browser component is about to loose focus. For instance,
	//  if focus was on the last HTML element and the user pressed the TAB key.
	//  |next| will be true(1) if the browser is giving focus to the next
	//  component and false(0) if the browser is giving focus to the previous
	//  component.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_focus_handler_capi.h">CEF source file: /include/capi/cef_focus_handler_capi.h(cef_focus_handler_t)</a>
	SetOnTakeFocus(fn TOnTakeFocus) // property event
	// SetOnSetFocus
	//  Called when the browser component is requesting focus. |source| indicates
	//  where the focus request is originating from. Return false(0) to allow the
	//  focus to be set or true(1) to cancel setting the focus.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_focus_handler_capi.h">CEF source file: /include/capi/cef_focus_handler_capi.h(cef_focus_handler_t)</a>
	SetOnSetFocus(fn TOnSetFocus) // property event
	// SetOnGotFocus
	//  Called when the browser component has received focus.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_focus_handler_capi.h">CEF source file: /include/capi/cef_focus_handler_capi.h(cef_focus_handler_t)</a>
	SetOnGotFocus(fn TOnGotFocus) // property event
	// SetOnBeforeContextMenu
	//  Called before a context menu is displayed. |params| provides information
	//  about the context menu state. |model| initially contains the default
	//  context menu. The |model| can be cleared to show no context menu or
	//  modified to show a custom menu. Do not keep references to |params| or
	//  |model| outside of this callback.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnBeforeContextMenu(fn TOnBeforeContextMenu) // property event
	// SetOnRunContextMenu
	//  Called to allow custom display of the context menu. |params| provides
	//  information about the context menu state. |model| contains the context
	//  menu model resulting from OnBeforeContextMenu. For custom display return
	//  true(1) and execute |callback| either synchronously or asynchronously
	//  with the selected command ID. For default display return false(0). Do not
	//  keep references to |params| or |model| outside of this callback.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnRunContextMenu(fn TOnRunContextMenu) // property event
	// SetOnContextMenuCommand
	//  Called to execute a command selected from the context menu. Return true
	//  (1) if the command was handled or false(0) for the default
	//  implementation. See cef_menu_id_t for the command ids that have default
	//  implementations. All user-defined command ids should be between
	//  MENU_ID_USER_FIRST and MENU_ID_USER_LAST. |params| will have the same
	//  values as what was passed to on_before_context_menu(). Do not keep a
	//  reference to |params| outside of this callback.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnContextMenuCommand(fn TOnContextMenuCommand) // property event
	// SetOnContextMenuDismissed
	//  Called when the context menu is dismissed irregardless of whether the menu
	//  was canceled or a command was selected.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnContextMenuDismissed(fn TOnContextMenuDismissed) // property event
	// SetOnRunQuickMenu
	//  Called to allow custom display of the quick menu for a windowless browser.
	//  |location| is the top left corner of the selected region. |size| is the
	//  size of the selected region. |edit_state_flags| is a combination of flags
	//  that represent the state of the quick menu. Return true(1) if the menu
	//  will be handled and execute |callback| either synchronously or
	//  asynchronously with the selected command ID. Return false(0) to cancel
	//  the menu.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnRunQuickMenu(fn TOnRunQuickMenu) // property event
	// SetOnQuickMenuCommand
	//  Called to execute a command selected from the quick menu for a windowless
	//  browser. Return true(1) if the command was handled or false(0) for the
	//  default implementation. See cef_menu_id_t for command IDs that have
	//  default implementations.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnQuickMenuCommand(fn TOnQuickMenuCommand) // property event
	// SetOnQuickMenuDismissed
	//  Called when the quick menu for a windowless browser is dismissed
	//  irregardless of whether the menu was canceled or a command was selected.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h(cef_context_menu_handler_t)</a>
	SetOnQuickMenuDismissed(fn TOnQuickMenuDismissed) // property event
	// SetOnPreKeyEvent
	//  Called before a keyboard event is sent to the renderer. |event| contains
	//  information about the keyboard event. |os_event| is the operating system
	//  event message, if any. Return true(1) if the event was handled or false
	//  (0) otherwise. If the event will be handled in on_key_event() as a
	//  keyboard shortcut set |is_keyboard_shortcut| to true(1) and return false
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_keyboard_handler_capi.h">CEF source file: /include/capi/cef_keyboard_handler_capi.h(cef_keyboard_handler_t)</a>
	SetOnPreKeyEvent(fn TOnPreKey) // property event
	// SetOnKeyEvent
	//  Called after the renderer and JavaScript in the page has had a chance to
	//  handle the event. |event| contains information about the keyboard event.
	//  |os_event| is the operating system event message, if any. Return true(1)
	//  if the keyboard event was handled or false(0) otherwise.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_keyboard_handler_capi.h">CEF source file: /include/capi/cef_keyboard_handler_capi.h(cef_keyboard_handler_t)</a>
	SetOnKeyEvent(fn TOnKey) // property event
	// SetOnAddressChange
	//  Called when a frame's address has changed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnAddressChange(fn TOnAddressChange) // property event
	// SetOnTitleChange
	//  Called when the page title changes.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnTitleChange(fn TOnTitleChange) // property event
	// SetOnFavIconUrlChange
	//  Called when the page icon changes.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnFavIconUrlChange(fn TOnFavIconUrlChange) // property event
	// SetOnFullScreenModeChange
	//  Called when web content in the page has toggled fullscreen mode. If
	//  |fullscreen| is true(1) the content will automatically be sized to fill
	//  the browser content area. If |fullscreen| is false(0) the content will
	//  automatically return to its original size and position. With the Alloy
	//  runtime the client is responsible for triggering the fullscreen transition
	//  (for example, by calling ICefWindow.SetFullscreen when using Views).
	//  With the Chrome runtime the fullscreen transition will be triggered
	//  automatically. The ICefWindowDelegate.OnWindowFullscreenTransition
	//  function will be called during the fullscreen transition for notification
	//  purposes.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnFullScreenModeChange(fn TOnFullScreenModeChange) // property event
	// SetOnTooltip
	//  Called when the browser is about to display a tooltip. |text| contains the
	//  text that will be displayed in the tooltip. To handle the display of the
	//  tooltip yourself return true(1). Otherwise, you can optionally modify
	//  |text| and then return false(0) to allow the browser to display the
	//  tooltip. When window rendering is disabled the application is responsible
	//  for drawing tooltips and the return value is ignored.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnTooltip(fn TOnTooltip) // property event
	// SetOnStatusMessage
	//  Called when the browser receives a status message. |value| contains the
	//  text that will be displayed in the status message.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnStatusMessage(fn TOnStatusMessage) // property event
	// SetOnConsoleMessage
	//  Called to display a console message. Return true(1) to stop the message
	//  from being output to the console.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnConsoleMessage(fn TOnConsoleMessage) // property event
	// SetOnAutoResize
	//  Called when auto-resize is enabled via
	//  cef_browser_host_t::SetAutoResizeEnabled and the contents have auto-
	//  resized. |new_size| will be the desired size in view coordinates. Return
	//  true(1) if the resize was handled or false(0) for default handling.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnAutoResize(fn TOnAutoResize) // property event
	// SetOnLoadingProgressChange
	//  Called when the overall page loading progress has changed. |progress|
	//  ranges from 0.0 to 1.0.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnLoadingProgressChange(fn TOnLoadingProgressChange) // property event
	// SetOnCursorChange
	//  Called when the browser's cursor has changed. If |type| is CT_CUSTOM then
	//  |custom_cursor_info| will be populated with the custom cursor information.
	//  Return true(1) if the cursor change was handled or false(0) for default
	//  handling.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnCursorChange(fn TOnCursorChange) // property event
	// SetOnMediaAccessChange
	//  Called when the browser's access to an audio and/or video source has
	//  changed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h(cef_display_handler_t)</a>
	SetOnMediaAccessChange(fn TOnMediaAccessChange) // property event
	// SetOnCanDownload
	//  Called before a download begins in response to a user-initiated action
	//  (e.g. alt + link click or link click that returns a `Content-Disposition:
	//  attachment` response from the server). |url| is the target download URL
	//  and |request_function| is the target function(GET, POST, etc). Return
	//  true(1) to proceed with the download or false(0) to cancel the download.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h(cef_download_handler_t)</a>
	SetOnCanDownload(fn TOnCanDownload) // property event
	// SetOnBeforeDownload
	//  Called before a download begins. |suggested_name| is the suggested name
	//  for the download file. By default the download will be canceled. Execute
	//  |callback| either asynchronously or in this function to continue the
	//  download if desired. Do not keep a reference to |download_item| outside of
	//  this function.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h(cef_download_handler_t)</a>
	SetOnBeforeDownload(fn TOnBeforeDownload) // property event
	// SetOnDownloadUpdated
	//  Called when a download's status or progress information has been updated.
	//  This may be called multiple times before and after OnBeforeDownload.
	//  Execute |callback| either asynchronously or in this function to cancel the
	//  download if desired. Do not keep a reference to |download_item| outside of
	//  this function.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h(cef_download_handler_t)</a>
	SetOnDownloadUpdated(fn TOnDownloadUpdated) // property event
	// SetOnJsdialog
	//  Called to run a JavaScript dialog. If |origin_url| is non-NULL it can be
	//  passed to the CefFormatUrlForSecurityDisplay function to retrieve a secure
	//  and user-friendly display string. The |default_prompt_text| value will be
	//  specified for prompt dialogs only. Set |suppress_message| to true(1) and
	//  return false(0) to suppress the message(suppressing messages is
	//  preferable to immediately executing the callback as this is used to detect
	//  presumably malicious behavior like spamming alert messages in
	//  onbeforeunload). Set |suppress_message| to false(0) and return false(0)
	//  to use the default implementation(the default implementation will show
	//  one modal dialog at a time and suppress any additional dialog requests
	//  until the displayed dialog is dismissed). Return true(1) if the
	//  application will use a custom dialog or if the callback has been executed
	//  immediately. Custom dialogs may be either modal or modeless. If a custom
	//  dialog is used the application must execute |callback| once the custom
	//  dialog is dismissed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h(cef_jsdialog_handler_t)</a>
	SetOnJsdialog(fn TOnJsdialog) // property event
	// SetOnBeforeUnloadDialog
	//  Called to run a dialog asking the user if they want to leave a page.
	//  Return false(0) to use the default dialog implementation. Return true(1)
	//  if the application will use a custom dialog or if the callback has been
	//  executed immediately. Custom dialogs may be either modal or modeless. If a
	//  custom dialog is used the application must execute |callback| once the
	//  custom dialog is dismissed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h(cef_jsdialog_handler_t)</a>
	SetOnBeforeUnloadDialog(fn TOnBeforeUnloadDialog) // property event
	// SetOnResetDialogState
	//  Called to cancel any pending dialogs and reset any saved dialog state.
	//  Will be called due to events like page navigation irregardless of whether
	//  any dialogs are currently pending.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h(cef_jsdialog_handler_t)</a>
	SetOnResetDialogState(fn TOnResetDialogState) // property event
	// SetOnDialogClosed
	//  Called when the dialog is closed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h(cef_jsdialog_handler_t)</a>
	SetOnDialogClosed(fn TOnDialogClosed) // property event
	// SetOnBeforePopup
	//  Called on the UI thread before a new popup browser is created. The
	//  |browser| and |frame| values represent the source of the popup request.
	//  The |target_url| and |target_frame_name| values indicate where the popup
	//  browser should navigate and may be NULL if not specified with the request.
	//  The |target_disposition| value indicates where the user intended to open
	//  the popup(e.g. current tab, new tab, etc). The |user_gesture| value will
	//  be true(1) if the popup was opened via explicit user gesture(e.g.
	//  clicking a link) or false(0) if the popup opened automatically(e.g. via
	//  the DomContentLoaded event). The |popupFeatures| structure contains
	//  additional information about the requested popup window. To allow creation
	//  of the popup browser optionally modify |windowInfo|, |client|, |settings|
	//  and |no_javascript_access| and return false(0). To cancel creation of the
	//  popup browser return true(1). The |client| and |settings| values will
	//  default to the source browser's values. If the |no_javascript_access|
	//  value is set to false(0) the new browser will not be scriptable and may
	//  not be hosted in the same renderer process as the source browser. Any
	//  modifications to |windowInfo| will be ignored if the parent browser is
	//  wrapped in a ICefBrowserView. Popup browser creation will be canceled
	//  if the parent browser is destroyed before the popup browser creation
	//  completes(indicated by a call to OnAfterCreated for the popup browser).
	//  The |extra_info| parameter provides an opportunity to specify extra
	//  information specific to the created popup browser that will be passed to
	//  ICefRenderProcessHandler.OnBrowserCreated in the render process.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_life_span_handler_capi.h">CEF source file: /include/capi/cef_life_span_handler_capi.h(cef_life_span_handler_t)</a>
	SetOnBeforePopup(fn TOnBeforePopup) // property event
	// SetOnAfterCreated
	//  Called after a new browser is created. It is now safe to begin performing
	//  actions with |browser|. ICefFrameHandler callbacks related to initial
	//  main frame creation will arrive before this callback. See
	//  ICefFrameHandler documentation for additional usage information.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_life_span_handler_capi.h">CEF source file: /include/capi/cef_life_span_handler_capi.h(cef_life_span_handler_t)</a>
	SetOnAfterCreated(fn TOnAfterCreated) // property event
	// SetOnBeforeClose
	//  Called just before a browser is destroyed. Release all references to the
	//  browser object and do not attempt to execute any functions on the browser
	//  object(other than IsValid, GetIdentifier or IsSame) after this callback
	//  returns. ICefFrameHandler callbacks related to final main frame
	//  destruction will arrive after this callback and ICefBrowser.IsValid
	//  will return false(0) at that time. Any in-progress network requests
	//  associated with |browser| will be aborted when the browser is destroyed,
	//  and ICefResourceRequestHandler callbacks related to those requests may
	//  still arrive on the IO thread after this callback. See ICefFrameHandler
	//  and OnClose() documentation for additional usage information.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_life_span_handler_capi.h">CEF source file: /include/capi/cef_life_span_handler_capi.h(cef_life_span_handler_t)</a>
	SetOnBeforeClose(fn TOnBeforeClose) // property event
	// SetOnClose
	//  Called when a browser has recieved a request to close. This may result
	//  directly from a call to ICefBrowserHost.*CloseBrowser or indirectly
	//  if the browser is parented to a top-level window created by CEF and the
	//  user attempts to close that window(by clicking the 'X', for example). The
	//  OnClose function will be called after the JavaScript 'onunload' event
	//  has been fired.
	//  An application should handle top-level owner window close notifications by
	//  calling ICefBrowserHost.TryCloseBrowser or
	//  ICefBrowserHost.CloseBrowser(false) instead of allowing the window
	//  to close immediately(see the examples below). This gives CEF an
	//  opportunity to process the 'onbeforeunload' event and optionally cancel
	//  the close before OnClose is called.
	//  When windowed rendering is enabled CEF will internally create a window or
	//  view to host the browser. In that case returning false(0) from OnClose()
	//  will send the standard close notification to the browser's top-level owner
	//  window(e.g. WM_CLOSE on Windows, performClose: on OS X, "delete_event" on
	//  Linux or ICefWindowDelegate.CanClose callback from Views). If the
	//  browser's host window/view has already been destroyed(via view hierarchy
	//  tear-down, for example) then OnClose() will not be called for that
	//  browser since is no longer possible to cancel the close.
	//  When windowed rendering is disabled returning false(0) from OnClose()
	//  will cause the browser object to be destroyed immediately.
	//  If the browser's top-level owner window requires a non-standard close
	//  notification then send that notification from OnClose() and return true.
	//  The ICefLifeSpanHandler.OnBeforeClose function will be called
	//  after OnClose()(if OnClose() is called) and immediately before the
	//  browser object is destroyed. The application should only exit after
	//  OnBeforeClose() has been called for all existing browsers.
	//  The below examples describe what should happen during window close when
	//  the browser is parented to an application-provided top-level window.
	//  Example 1: Using ICefBrowserHost.TryCloseBrowser(). This is
	//  recommended for clients using standard close handling and windows created
	//  on the browser process UI thread.
	//  <code>
	//  1. User clicks the window close button which sends a close notification
	//  to the application's top-level window.
	//  2. Application's top-level window receives the close notification and
	//  calls TryCloseBrowser()(which internally calls CloseBrowser(false)).
	//  TryCloseBrowser() returns false so the client cancels the window
	//  close.
	//  3. JavaScript 'onbeforeunload' handler executes and shows the close
	//  confirmation dialog(which can be overridden via
	//  ICefJSDialogHandler.OnBeforeUnloadDialog()).
	//  4. User approves the close.
	//  5. JavaScript 'onunload' handler executes.
	//  6. CEF sends a close notification to the application's top-level window
	//  (because OnClose() returned false by default).
	//  7. Application's top-level window receives the close notification and
	//  calls TryCloseBrowser(). TryCloseBrowser() returns true so the client
	//  allows the window close.
	//  8. Application's top-level window is destroyed.
	//  9. Application's OnBeforeClose() handler is called and the browser object is destroyed.
	//  10. Application exits by calling cef_quit_message_loop() if no other browsers exist.
	//  </code>
	//  Example 2: Using ICefBrowserHost::CloseBrowser(false) and
	//  implementing the OnClose() callback. This is recommended for clients
	//  using non-standard close handling or windows that were not created on the
	//  browser process UI thread.
	//  <code>
	//  1. User clicks the window close button which sends a close notification
	//  to the application's top-level window.
	//  2. Application's top-level window receives the close notification and:
	//  A. Calls ICefBrowserHost.CloseBrowser(false).
	//  B. Cancels the window close.
	//  3. JavaScript 'onbeforeunload' handler executes and shows the close
	//  confirmation dialog(which can be overridden via
	//  ICefJSDialogHandler.OnBeforeUnloadDialog()).
	//  4. User approves the close.
	//  5. JavaScript 'onunload' handler executes.
	//  6. Application's OnClose() handler is called. Application will:
	//  A. Set a flag to indicate that the next close attempt will be allowed.
	//  B. Return false.
	//  7. CEF sends an close notification to the application's top-level window.
	//  8. Application's top-level window receives the close notification and
	//  allows the window to close based on the flag from #6B.
	//  9. Application's top-level window is destroyed.
	//  10. Application's OnBeforeClose() handler is called and the browser object is destroyed.
	//  11. Application exits by calling cef_quit_message_loop() if no other browsers exist.
	//  </code>
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_life_span_handler_capi.h">CEF source file: /include/capi/cef_life_span_handler_capi.h(cef_life_span_handler_t)</a>
	SetOnClose(fn TOnClose) // property event
	// SetOnBeforeBrowse
	//  Called on the UI thread before browser navigation. Return true(1) to
	//  cancel the navigation or false(0) to allow the navigation to proceed. The
	//  |request| object cannot be modified in this callback.
	//  ICefLoadHandler.OnLoadingStateChange will be called twice in all
	//  cases. If the navigation is allowed ICefLoadHandler.OnLoadStart and
	//  ICefLoadHandler.OnLoadEnd will be called. If the navigation is
	//  canceled ICefLoadHandler.OnLoadError will be called with an
	//  |errorCode| value of ERR_ABORTED. The |user_gesture| value will be true
	//  (1) if the browser navigated via explicit user gesture(e.g. clicking a
	//  link) or false(0) if it navigated automatically(e.g. via the
	//  DomContentLoaded event).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnBeforeBrowse(fn TOnBeforeBrowse) // property event
	// SetOnOpenUrlFromTab
	//  Called on the UI thread before OnBeforeBrowse in certain limited cases
	//  where navigating a new or different browser might be desirable. This
	//  includes user-initiated navigation that might open in a special way(e.g.
	//  links clicked via middle-click or ctrl + left-click) and certain types of
	//  cross-origin navigation initiated from the renderer process(e.g.
	//  navigating the top-level frame to/from a file URL). The |browser| and
	//  |frame| values represent the source of the navigation. The
	//  |target_disposition| value indicates where the user intended to navigate
	//  the browser based on standard Chromium behaviors(e.g. current tab, new
	//  tab, etc). The |user_gesture| value will be true(1) if the browser
	//  navigated via explicit user gesture(e.g. clicking a link) or false(0) if
	//  it navigated automatically(e.g. via the DomContentLoaded event). Return
	//  true(1) to cancel the navigation or false(0) to allow the navigation to
	//  proceed in the source browser's top-level frame.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnOpenUrlFromTab(fn TOnOpenUrlFromTab) // property event
	// SetOnGetAuthCredentials
	//  Called on the IO thread when the browser needs credentials from the user.
	//  |origin_url| is the origin making this authentication request. |isProxy|
	//  indicates whether the host is a proxy server. |host| contains the hostname
	//  and |port| contains the port number. |realm| is the realm of the challenge
	//  and may be NULL. |scheme| is the authentication scheme used, such as
	//  "basic" or "digest", and will be NULL if the source of the request is an
	//  FTP server. Return true(1) to continue the request and call
	//  ICefAuthCallback.cont() either in this function or at a later time
	//  when the authentication information is available. Return false(0) to
	//  cancel the request immediately.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnGetAuthCredentials(fn TOnGetAuthCredentials) // property event
	// SetOnCertificateError
	//  Called on the UI thread to handle requests for URLs with an invalid SSL
	//  certificate. Return true(1) and call ICefCallback functions either in
	//  this function or at a later time to continue or cancel the request. Return
	//  false(0) to cancel the request immediately. If
	//  TCefSettings.ignore_certificate_errors is set all invalid certificates
	//  will be accepted without calling this function.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnCertificateError(fn TOnCertificateError) // property event
	// SetOnSelectClientCertificate
	//  Called on the UI thread when a client certificate is being requested for
	//  authentication. Return false(0) to use the default behavior and
	//  automatically select the first certificate available. Return true(1) and
	//  call ICefSelectClientCertificateCallback.Select either in this
	//  function or at a later time to select a certificate. Do not call Select or
	//  call it with NULL to continue without using any certificate. |isProxy|
	//  indicates whether the host is an HTTPS proxy or the origin server. |host|
	//  and |port| contains the hostname and port of the SSL server.
	//  |certificates| is the list of certificates to choose from; this list has
	//  already been pruned by Chromium so that it only contains certificates from
	//  issuers that the server trusts.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnSelectClientCertificate(fn TOnSelectClientCertificate) // property event
	// SetOnRenderViewReady
	//  Called on the browser process UI thread when the render view associated
	//  with |browser| is ready to receive/handle IPC messages in the render
	//  process.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnRenderViewReady(fn TOnRenderViewReady) // property event
	// SetOnRenderProcessTerminated
	//  Called on the browser process UI thread when the render process terminates
	//  unexpectedly. |status| indicates how the process terminated.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnRenderProcessTerminated(fn TOnRenderProcessTerminated) // property event
	// SetOnGetResourceRequestHandler_ReqHdlr
	//  Called on the browser process IO thread before a resource request is
	//  initiated. The |browser| and |frame| values represent the source of the
	//  request. |request| represents the request contents and cannot be modified
	//  in this callback. |is_navigation| will be true(1) if the resource request
	//  is a navigation. |is_download| will be true(1) if the resource request is
	//  a download. |request_initiator| is the origin(scheme + domain) of the
	//  page that initiated the request. Set |disable_default_handling| to true
	//  (1) to disable default handling of the request, in which case it will need
	//  to be handled via ICefResourceRequestHandler.GetResourceHandler or it
	//  will be canceled. To allow the resource load to proceed with default
	//  handling return NULL. To specify a handler for the resource return a
	//  ICefResourceRequestHandler object. If this callback returns NULL the
	//  same function will be called on the associated
	//  ICefRequestContextHandler, if any.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnGetResourceRequestHandler_ReqHdlr(fn TOnGetResourceRequestHandler) // property event
	// SetOnDocumentAvailableInMainFrame
	//  Called on the browser process UI thread when the window.document object of
	//  the main frame has been created.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h(cef_request_handler_t)</a>
	SetOnDocumentAvailableInMainFrame(fn TOnDocumentAvailableInMainFrame) // property event
	// SetOnBeforeResourceLoad
	//  Called on the IO thread before a resource request is loaded. The |browser|
	//  and |frame| values represent the source of the request, and may be NULL
	//  for requests originating from service workers or ICefUrlRequest. To
	//  redirect or change the resource load optionally modify |request|.
	//  Modification of the request URL will be treated as a redirect. Return
	//  RV_CONTINUE to continue the request immediately. Return RV_CONTINUE_ASYNC
	//  and call ICefCallback functions at a later time to continue or cancel
	//  the request asynchronously. Return RV_CANCEL to cancel the request
	//  immediately.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnBeforeResourceLoad(fn TOnBeforeResourceLoad) // property event
	// SetOnGetResourceHandler
	//  Called on the IO thread before a resource is loaded. The |browser| and
	//  |frame| values represent the source of the request, and may be NULL for
	//  requests originating from service workers or ICefUrlRequest. To allow
	//  the resource to load using the default network loader return NULL. To
	//  specify a handler for the resource return a ICefResourceHandler object.
	//  The |request| object cannot not be modified in this callback.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnGetResourceHandler(fn TOnGetResourceHandler) // property event
	// SetOnResourceRedirect
	//  Called on the IO thread when a resource load is redirected. The |browser|
	//  and |frame| values represent the source of the request, and may be NULL
	//  for requests originating from service workers or ICefUrlRequest. The
	//  |request| parameter will contain the old URL and other request-related
	//  information. The |response| parameter will contain the response that
	//  resulted in the redirect. The |new_url| parameter will contain the new URL
	//  and can be changed if desired. The |request| and |response| objects cannot
	//  be modified in this callback.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnResourceRedirect(fn TOnResourceRedirect) // property event
	// SetOnResourceResponse
	//  Called on the IO thread when a resource response is received. The
	//  |browser| and |frame| values represent the source of the request, and may
	//  be NULL for requests originating from service workers or ICefUrlRequest.
	//  To allow the resource load to proceed without modification return false
	//  (0). To redirect or retry the resource load optionally modify |request|
	//  and return true(1). Modification of the request URL will be treated as a
	//  redirect. Requests handled using the default network loader cannot be
	//  redirected in this callback. The |response| object cannot be modified in
	//  this callback.
	//  WARNING: Redirecting using this function is deprecated. Use
	//  OnBeforeResourceLoad or GetResourceHandler to perform redirects.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnResourceResponse(fn TOnResourceResponse) // property event
	// SetOnGetResourceResponseFilter
	//  Called on the IO thread to optionally filter resource response content.
	//  The |browser| and |frame| values represent the source of the request, and
	//  may be NULL for requests originating from service workers or
	//  ICefUrlRequest. |request| and |response| represent the request and
	//  response respectively and cannot be modified in this callback.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnGetResourceResponseFilter(fn TOnGetResourceResponseFilter) // property event
	// SetOnResourceLoadComplete
	//  Called on the IO thread when a resource load has completed. The |browser|
	//  and |frame| values represent the source of the request, and may be NULL
	//  for requests originating from service workers or ICefUrlRequest.
	//  |request| and |response| represent the request and response respectively
	//  and cannot be modified in this callback. |status| indicates the load
	//  completion status. |received_content_length| is the number of response
	//  bytes actually read. This function will be called for all requests,
	//  including requests that are aborted due to CEF shutdown or destruction of
	//  the associated browser. In cases where the associated browser is destroyed
	//  this callback may arrive after the ICefLifeSpanHandler.OnBeforeClose
	//  callback for that browser. The ICefFrame.IsValid function can be used
	//  to test for this situation, and care should be taken not to call |browser|
	//  or |frame| functions that modify state(like LoadURL, SendProcessMessage,
	//  etc.) if the frame is invalid.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnResourceLoadComplete(fn TOnResourceLoadComplete) // property event
	// SetOnProtocolExecution
	//  Called on the IO thread to handle requests for URLs with an unknown
	//  protocol component. The |browser| and |frame| values represent the source
	//  of the request, and may be NULL for requests originating from service
	//  workers or ICefUrlRequest. |request| cannot be modified in this
	//  callback. Set |allow_os_execution| to true(1) to attempt execution via
	//  the registered OS protocol handler, if any. SECURITY WARNING: YOU SHOULD
	//  USE THIS METHOD TO ENFORCE RESTRICTIONS BASED ON SCHEME, HOST OR OTHER URL
	//  ANALYSIS BEFORE ALLOWING OS EXECUTION.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_resource_request_handler_t)</a>
	SetOnProtocolExecution(fn TOnProtocolExecution) // property event
	// SetOnCanSendCookie
	//  Called on the IO thread before a resource request is sent. The |browser|
	//  and |frame| values represent the source of the request, and may be NULL
	//  for requests originating from service workers or ICefUrlRequest.
	//  |request| cannot be modified in this callback. Return true(1) if the
	//  specified cookie can be sent with the request or false(0) otherwise.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_cookie_access_filter_t)</a>
	SetOnCanSendCookie(fn TOnCanSendCookie) // property event
	// SetOnCanSaveCookie
	//  Called on the IO thread after a resource response is received. The
	//  |browser| and |frame| values represent the source of the request, and may
	//  be NULL for requests originating from service workers or ICefUrlRequest.
	//  |request| cannot be modified in this callback. Return true(1) if the
	//  specified cookie returned with the response can be saved or false(0)
	//  otherwise.
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h(cef_cookie_access_filter_t)</a>
	SetOnCanSaveCookie(fn TOnCanSaveCookie) // property event
	// SetOnFileDialog
	//  Called to run a file chooser dialog. |mode| represents the type of dialog
	//  to display. |title| to the title to be used for the dialog and may be NULL
	//  to show the default title("Open" or "Save" depending on the mode).
	//  |default_file_path| is the path with optional directory and/or file name
	//  component that should be initially selected in the dialog.
	//  |accept_filters| are used to restrict the selectable file types and may
	//  any combination of(a) valid lower-cased MIME types(e.g. "text/*" or
	//  "image/*"),(b) individual file extensions(e.g. ".txt" or ".png"), or(c)
	//  combined description and file extension delimited using "|" and ";"(e.g.
	//  "Image Types|.png;.gif;.jpg"). To display a custom dialog return true(1)
	//  and execute |callback| either inline or at a later time. To display the
	//  default dialog return false(0).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h(cef_dialog_handler_t)</a>
	SetOnFileDialog(fn TOnFileDialog) // property event
	// SetOnGetAccessibilityHandler
	//  Return the handler for accessibility notifications. If no handler is
	//  provided the default implementation will be used.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnGetAccessibilityHandler(fn TOnGetAccessibilityHandler) // property event
	// SetOnGetRootScreenRect
	//  Called to retrieve the root window rectangle in screen DIP coordinates.
	//  Return true(1) if the rectangle was provided. If this function returns
	//  false(0) the rectangle from OnGetViewRect will be used.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnGetRootScreenRect(fn TOnGetRootScreenRect) // property event
	// SetOnGetViewRect
	//  Called to retrieve the view rectangle in screen DIP coordinates. This
	//  function must always provide a non-NULL rectangle.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnGetViewRect(fn TOnGetViewRect) // property event
	// SetOnGetScreenPoint
	//  Called to retrieve the translation from view DIP coordinates to screen
	//  coordinates. Windows/Linux should provide screen device(pixel)
	//  coordinates and MacOS should provide screen DIP coordinates. Return true
	//  (1) if the requested coordinates were provided.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnGetScreenPoint(fn TOnGetScreenPoint) // property event
	// SetOnGetScreenInfo
	//  Called to allow the client to fill in the TCefScreenInfo object with
	//  appropriate values. Return true(1) if the |screen_info| structure has
	//  been modified.
	//  If the screen info rectangle is left NULL the rectangle from OnGetViewRect
	//  will be used. If the rectangle is still NULL or invalid popups may not be
	//  drawn correctly.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnGetScreenInfo(fn TOnGetScreenInfo) // property event
	// SetOnPopupShow
	//  Called when the browser wants to show or hide the popup widget. The popup
	//  should be shown if |show| is true(1) and hidden if |show| is false(0).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnPopupShow(fn TOnPopupShow) // property event
	// SetOnPopupSize
	//  Called when the browser wants to move or resize the popup widget. |rect|
	//  contains the new location and size in view coordinates.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnPopupSize(fn TOnPopupSize) // property event
	// SetOnPaint
	//  Called when an element should be painted. Pixel values passed to this
	//  function are scaled relative to view coordinates based on the value of
	//  TCefScreenInfo.device_scale_factor returned from OnGetScreenInfo. |type|
	//  indicates whether the element is the view or the popup widget. |buffer|
	//  contains the pixel data for the whole image. |dirtyRects| contains the set
	//  of rectangles in pixel coordinates that need to be repainted. |buffer|
	//  will be |width|*|height|*4 bytes in size and represents a BGRA image with
	//  an upper-left origin. This function is only called when
	//  TCefWindowInfo.shared_texture_enabled is set to false(0).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnPaint(fn TOnPaint) // property event
	// SetOnAcceleratedPaint
	//  Called when an element has been rendered to the shared texture handle.
	//  |type| indicates whether the element is the view or the popup widget.
	//  |dirtyRects| contains the set of rectangles in pixel coordinates that need
	//  to be repainted. |shared_handle| is the handle for a D3D11 Texture2D that
	//  can be accessed via ID3D11Device using the OpenSharedResource function.
	//  This function is only called when TCefWindowInfo.shared_texture_enabled
	//  is set to true(1), and is currently only supported on Windows.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnAcceleratedPaint(fn TOnAcceleratedPaint) // property event
	// SetOnGetTouchHandleSize
	//  Called to retrieve the size of the touch handle for the specified
	//  |orientation|.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnGetTouchHandleSize(fn TOnGetTouchHandleSize) // property event
	// SetOnTouchHandleStateChanged
	//  Called when touch handle state is updated. The client is responsible for
	//  rendering the touch handles.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnTouchHandleStateChanged(fn TOnTouchHandleStateChanged) // property event
	// SetOnStartDragging
	//  Called when the user starts dragging content in the web view. Contextual
	//  information about the dragged content is supplied by |drag_data|.(|x|,
	//  |y|) is the drag start location in screen coordinates. OS APIs that run a
	//  system message loop may be used within the StartDragging call.
	//  Return false(0) to abort the drag operation. Don't call any of
	//  ICefBrowserHost.DragSource*Ended* functions after returning false(0).
	//  Return true(1) to handle the drag operation. Call
	//  ICefBrowserHost.DragSourceEndedAt and DragSourceSystemDragEnded either
	//  synchronously or asynchronously to inform the web view that the drag
	//  operation has ended.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnStartDragging(fn TOnStartDragging) // property event
	// SetOnUpdateDragCursor
	//  Called when the web view wants to update the mouse cursor during a drag &
	//  drop operation. |operation| describes the allowed operation(none, move,
	//  copy, link).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnUpdateDragCursor(fn TOnUpdateDragCursor) // property event
	// SetOnScrollOffsetChanged
	//  Called when the scroll offset has changed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnScrollOffsetChanged(fn TOnScrollOffsetChanged) // property event
	// SetOnIMECompositionRangeChanged
	//  Called when the IME composition range has changed. |selected_range| is the
	//  range of characters that have been selected. |character_bounds| is the
	//  bounds of each character in view coordinates.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnIMECompositionRangeChanged(fn TOnIMECompositionRangeChanged) // property event
	// SetOnTextSelectionChanged
	//  Called when text selection has changed for the specified |browser|.
	//  |selected_text| is the currently selected text and |selected_range| is the
	//  character range.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnTextSelectionChanged(fn TOnTextSelectionChanged) // property event
	// SetOnVirtualKeyboardRequested
	//  Called when an on-screen keyboard should be shown or hidden for the
	//  specified |browser|. |input_mode| specifies what kind of keyboard should
	//  be opened. If |input_mode| is CEF_TEXT_INPUT_MODE_NONE, any existing
	//  keyboard for this browser should be hidden.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h(cef_render_handler_t)</a>
	SetOnVirtualKeyboardRequested(fn TOnVirtualKeyboardRequested) // property event
	// SetOnDragEnter
	//  Called when an external drag event enters the browser window. |dragData|
	//  contains the drag event data and |mask| represents the type of drag
	//  operation. Return false(0) for default drag handling behavior or true(1)
	//  to cancel the drag event.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_handler_capi.h">CEF source file: /include/capi/cef_drag_handler_capi.h(cef_drag_handler_t)</a>
	SetOnDragEnter(fn TOnDragEnter) // property event
	// SetOnDraggableRegionsChanged
	//  Called whenever draggable regions for the browser window change. These can
	//  be specified using the '-webkit-app-region: drag/no-drag' CSS-property. If
	//  draggable regions are never defined in a document this function will also
	//  never be called. If the last draggable region is removed from a document
	//  this function will be called with an NULL vector.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_handler_capi.h">CEF source file: /include/capi/cef_drag_handler_capi.h(cef_drag_handler_t)</a>
	SetOnDraggableRegionsChanged(fn TOnDraggableRegionsChanged) // property event
	// SetOnFindResult
	//  Called to report find results returned by ICefBrowserHost.find().
	//  |identifer| is a unique incremental identifier for the currently active
	//  search, |count| is the number of matches currently identified,
	//  |selectionRect| is the location of where the match was found(in window
	//  coordinates), |activeMatchOrdinal| is the current position in the search
	//  results, and |finalUpdate| is true(1) if this is the last find
	//  notification.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_find_handler_capi.h">CEF source file: /include/capi/cef_find_handler_capi.h(cef_find_handler_t)</a>
	SetOnFindResult(fn TOnFindResult) // property event
	// SetOnRequestContextInitialized
	//  Called on the browser process UI thread immediately after the request
	//  context has been initialized.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_context_handler_capi.h">CEF source file: /include/capi/cef_request_context_handler_capi.h(cef_request_context_handler_t)</a>
	SetOnRequestContextInitialized(fn TOnRequestContextInitialized) // property event
	// SetOnGetResourceRequestHandler_ReqCtxHdlr
	//  Called on the browser process IO thread before a resource request is
	//  initiated. The |browser| and |frame| values represent the source of the
	//  request, and may be NULL for requests originating from service workers or
	//  ICefUrlRequest. |request| represents the request contents and cannot be
	//  modified in this callback. |is_navigation| will be true(1) if the
	//  resource request is a navigation. |is_download| will be true(1) if the
	//  resource request is a download. |request_initiator| is the origin(scheme
	//  + domain) of the page that initiated the request. Set
	//  |disable_default_handling| to true(1) to disable default handling of the
	//  request, in which case it will need to be handled via
	//  ICefResourceRequestHandler.GetResourceHandler or it will be canceled.
	//  To allow the resource load to proceed with default handling return NULL.
	//  To specify a handler for the resource return a
	//  ICefResourceRequestHandler object. This function will not be called if
	//  the client associated with |browser| returns a non-NULL value from
	//  ICefRequestHandler.GetResourceRequestHandler for the same request
	//  (identified by ICefRequest.GetIdentifier).
	//  This event will be called on the browser process CEF IO thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_context_handler_capi.h">CEF source file: /include/capi/cef_request_context_handler_capi.h(cef_request_context_handler_t)</a>
	SetOnGetResourceRequestHandler_ReqCtxHdlr(fn TOnGetResourceRequestHandler) // property event
	// SetOnSinks
	//  The list of available media sinks has changed or
	//  ICefMediaRouter.NotifyCurrentSinks was called.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h(cef_media_observer_t)</a>
	SetOnSinks(fn TOnSinks) // property event
	// SetOnRoutes
	//  The list of available media routes has changed or
	//  ICefMediaRouter.NotifyCurrentRoutes was called.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h(cef_media_observer_t)</a>
	SetOnRoutes(fn TOnRoutes) // property event
	// SetOnRouteStateChanged
	//  The connection state of |route| has changed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h(cef_media_observer_t)</a>
	SetOnRouteStateChanged(fn TOnRouteStateChanged) // property event
	// SetOnRouteMessageReceived
	//  A message was recieved over |route|. |message| is only valid for the scope
	//  of this callback and should be copied if necessary.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h(cef_media_observer_t)</a>
	SetOnRouteMessageReceived(fn TOnRouteMessageReceived) // property event
	// SetOnGetAudioParameters
	//  Called on the UI thread to allow configuration of audio stream parameters.
	//  Return true(1) to proceed with audio stream capture, or false(0) to
	//  cancel it. All members of |params| can optionally be configured here, but
	//  they are also pre-filled with some sensible defaults.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h(cef_audio_handler_t)</a>
	SetOnGetAudioParameters(fn TOnGetAudioParameters) // property event
	// SetOnAudioStreamStarted
	//  Called on a browser audio capture thread when the browser starts streaming
	//  audio. OnAudioStreamStopped will always be called after
	//  OnAudioStreamStarted; both functions may be called multiple times for the
	//  same browser. |params| contains the audio parameters like sample rate and
	//  channel layout. |channels| is the number of channels.
	//  This event will be called on a browser audio capture thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h(cef_audio_handler_t)</a>
	SetOnAudioStreamStarted(fn TOnAudioStreamStarted) // property event
	// SetOnAudioStreamPacket
	//  Called on the audio stream thread when a PCM packet is received for the
	//  stream. |data| is an array representing the raw PCM data as a floating
	//  point type, i.e. 4-byte value(s). |frames| is the number of frames in the
	//  PCM packet. |pts| is the presentation timestamp(in milliseconds since the
	//  Unix Epoch) and represents the time at which the decompressed packet
	//  should be presented to the user. Based on |frames| and the
	//  |channel_layout| value passed to OnAudioStreamStarted you can calculate
	//  the size of the |data| array in bytes.
	//  This event will be called on a browser audio capture thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h(cef_audio_handler_t)</a>
	SetOnAudioStreamPacket(fn TOnAudioStreamPacket) // property event
	// SetOnAudioStreamStopped
	//  Called on the UI thread when the stream has stopped. OnAudioSteamStopped
	//  will always be called after OnAudioStreamStarted; both functions may be
	//  called multiple times for the same stream.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h(cef_audio_handler_t)</a>
	SetOnAudioStreamStopped(fn TOnAudioStreamStopped) // property event
	// SetOnAudioStreamError
	//  Called on the UI or audio stream thread when an error occurred. During the
	//  stream creation phase this callback will be called on the UI thread while
	//  in the capturing phase it will be called on the audio stream thread. The
	//  stream will be stopped immediately.
	//  This event will be called on the browser process CEF UI thread or a browser audio capture thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h(cef_audio_handler_t)</a>
	SetOnAudioStreamError(fn TOnAudioStreamError) // property event
	// SetOnDevToolsMessage
	//  Method that will be called on receipt of a DevTools protocol message.
	//  |browser| is the originating browser instance. |message| is a UTF8-encoded
	//  JSON dictionary representing either a function result or an event.
	//  |message| is only valid for the scope of this callback and should be
	//  copied if necessary. Return true(1) if the message was handled or false
	//  (0) if the message should be further processed and passed to the
	//  OnDevToolsMethodResult or OnDevToolsEvent functions as appropriate.
	//  Method result dictionaries include an "id"(int) value that identifies the
	//  orginating function call sent from
	//  ICefBrowserHost.SendDevToolsMessage, and optionally either a "result"
	//  (dictionary) or "error"(dictionary) value. The "error" dictionary will
	//  contain "code"(int) and "message"(string) values. Event dictionaries
	//  include a "function"(string) value and optionally a "params"(dictionary)
	//  value. See the DevTools protocol documentation at
	//  https://chromedevtools.github.io/devtools-protocol/ for details of
	//  supported function calls and the expected "result" or "params" dictionary
	//  contents. JSON dictionaries can be parsed using the CefParseJSON function
	//  if desired, however be aware of performance considerations when parsing
	//  large messages(some of which may exceed 1MB in size).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsMessage(fn TOnDevToolsMessage) // property event
	// SetOnDevToolsRawMessage
	//  Method that will be called on receipt of a DevTools protocol message.
	//  |browser| is the originating browser instance. |message| is a UTF8-encoded
	//  JSON dictionary representing either a function result or an event.
	//  |message| is only valid for the scope of this callback and should be
	//  copied if necessary. Return true(1) if the message was handled or false
	//  (0) if the message should be further processed and passed to the
	//  OnDevToolsMethodResult or OnDevToolsEvent functions as appropriate.
	//  Method result dictionaries include an "id"(int) value that identifies the
	//  orginating function call sent from
	//  ICefBrowserHost.SendDevToolsMessage, and optionally either a "result"
	//  (dictionary) or "error"(dictionary) value. The "error" dictionary will
	//  contain "code"(int) and "message"(string) values. Event dictionaries
	//  include a "function"(string) value and optionally a "params"(dictionary)
	//  value. See the DevTools protocol documentation at
	//  https://chromedevtools.github.io/devtools-protocol/ for details of
	//  supported function calls and the expected "result" or "params" dictionary
	//  contents. JSON dictionaries can be parsed using the CefParseJSON function
	//  if desired, however be aware of performance considerations when parsing
	//  large messages(some of which may exceed 1MB in size).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsRawMessage(fn TOnDevToolsRawMessage) // property event
	// SetOnDevToolsMethodResult
	//  Method that will be called after attempted execution of a DevTools
	//  protocol function. |browser| is the originating browser instance.
	//  |message_id| is the "id" value that identifies the originating function
	//  call message. If the function succeeded |success| will be true(1) and
	//  |result| will be the UTF8-encoded JSON "result" dictionary value(which
	//  may be NULL). If the function failed |success| will be false(0) and
	//  |result| will be the UTF8-encoded JSON "error" dictionary value. |result|
	//  is only valid for the scope of this callback and should be copied if
	//  necessary. See the OnDevToolsMessage documentation for additional details
	//  on |result| contents.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsMethodResult(fn TOnDevToolsMethodResult) // property event
	// SetOnDevToolsMethodRawResult
	//  Method that will be called after attempted execution of a DevTools
	//  protocol function. |browser| is the originating browser instance.
	//  |message_id| is the "id" value that identifies the originating function
	//  call message. If the function succeeded |success| will be true(1) and
	//  |result| will be the UTF8-encoded JSON "result" dictionary value(which
	//  may be NULL). If the function failed |success| will be false(0) and
	//  |result| will be the UTF8-encoded JSON "error" dictionary value. |result|
	//  is only valid for the scope of this callback and should be copied if
	//  necessary. See the OnDevToolsMessage documentation for additional details
	//  on |result| contents.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsMethodRawResult(fn TOnDevToolsMethodRawResult) // property event
	// SetOnDevToolsEvent
	//  Method that will be called on receipt of a DevTools protocol event.
	//  |browser| is the originating browser instance. |function| is the
	//  "function" value. |params| is the UTF8-encoded JSON "params" dictionary
	//  value(which may be NULL). |params| is only valid for the scope of this
	//  callback and should be copied if necessary. See the OnDevToolsMessage
	//  documentation for additional details on |params| contents.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsEvent(fn TOnDevToolsEvent) // property event
	// SetOnDevToolsRawEvent
	//  Method that will be called on receipt of a DevTools protocol event.
	//  |browser| is the originating browser instance. |function| is the
	//  "function" value. |params| is the UTF8-encoded JSON "params" dictionary
	//  value(which may be NULL). |params| is only valid for the scope of this
	//  callback and should be copied if necessary. See the OnDevToolsMessage
	//  documentation for additional details on |params| contents.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsRawEvent(fn TOnDevToolsRawEvent) // property event
	// SetOnDevToolsAgentAttached
	//  Method that will be called when the DevTools agent has attached. |browser|
	//  is the originating browser instance. This will generally occur in response
	//  to the first message sent while the agent is detached.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsAgentAttached(fn TOnDevToolsAgentAttached) // property event
	// SetOnDevToolsAgentDetached
	//  Method that will be called when the DevTools agent has detached. |browser|
	//  is the originating browser instance. Any function results that were
	//  pending before the agent became detached will not be delivered, and any
	//  active event subscriptions will be canceled.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h(cef_dev_tools_message_observer_t)</a>
	SetOnDevToolsAgentDetached(fn TOnDevToolsAgentDetached) // property event
	// SetOnExtensionLoadFailed
	//  Called if the ICefRequestContext.LoadExtension request fails. |result|
	//  will be the error code.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionLoadFailed(fn TOnExtensionLoadFailed) // property event
	// SetOnExtensionLoaded
	//  Called if the ICefRequestContext.LoadExtension request succeeds.
	//  |extension| is the loaded extension.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionLoaded(fn TOnExtensionLoaded) // property event
	// SetOnExtensionUnloaded
	//  Called after the ICefExtension.Unload request has completed.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionUnloaded(fn TOnExtensionUnloaded) // property event
	// SetOnExtensionBeforeBackgroundBrowser
	//  Called when an extension needs a browser to host a background script
	//  specified via the "background" manifest key. The browser will have no
	//  visible window and cannot be displayed. |extension| is the extension that
	//  is loading the background script. |url| is an internally generated
	//  reference to an HTML page that will be used to load the background script
	//  via a "<script>" src attribute. To allow creation of the browser
	//  optionally modify |client| and |settings| and return false(0). To cancel
	//  creation of the browser(and consequently cancel load of the background
	//  script) return true(1). Successful creation will be indicated by a call
	//  to ICefLifeSpanHandler.OnAfterCreated, and
	//  ICefBrowserHost.IsBackgroundHost will return true(1) for the
	//  resulting browser. See https://developer.chrome.com/extensions/event_pages
	//  for more information about extension background script usage.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionBeforeBackgroundBrowser(fn TOnBeforeBackgroundBrowser) // property event
	// SetOnExtensionBeforeBrowser
	//  Called when an extension API(e.g. chrome.tabs.create) requests creation
	//  of a new browser. |extension| and |browser| are the source of the API
	//  call. |active_browser| may optionally be specified via the windowId
	//  property or returned via the get_active_browser() callback and provides
	//  the default |client| and |settings| values for the new browser. |index| is
	//  the position value optionally specified via the index property. |url| is
	//  the URL that will be loaded in the browser. |active| is true(1) if the
	//  new browser should be active when opened. To allow creation of the
	//  browser optionally modify |windowInfo|, |client| and |settings| and return
	//  false(0). To cancel creation of the browser return true(1). Successful
	//  creation will be indicated by a call to
	//  ICefLifeSpanHandler.OnAfterCreated. Any modifications to |windowInfo|
	//  will be ignored if |active_browser| is wrapped in a ICefBrowserView.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionBeforeBrowser(fn TOnBeforeBrowser) // property event
	// SetOnExtensionGetActiveBrowser
	//  Called when no tabId is specified to an extension API call that accepts a
	//  tabId parameter(e.g. chrome.tabs.*). |extension| and |browser| are the
	//  source of the API call. Return the browser that will be acted on by the
	//  API call or return NULL to act on |browser|. The returned browser must
	//  share the same ICefRequestContext as |browser|. Incognito browsers
	//  should not be considered unless the source extension has incognito access
	//  enabled, in which case |include_incognito| will be true(1).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionGetActiveBrowser(fn TOnGetActiveBrowser) // property event
	// SetOnExtensionCanAccessBrowser
	//  Called when the tabId associated with |target_browser| is specified to an
	//  extension API call that accepts a tabId parameter(e.g. chrome.tabs.*).
	//  |extension| and |browser| are the source of the API call. Return true(1)
	//  to allow access of false(0) to deny access. Access to incognito browsers
	//  should not be allowed unless the source extension has incognito access
	//  enabled, in which case |include_incognito| will be true(1).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionCanAccessBrowser(fn TOnCanAccessBrowser) // property event
	// SetOnExtensionGetExtensionResource
	//  Called to retrieve an extension resource that would normally be loaded
	//  from disk(e.g. if a file parameter is specified to
	//  chrome.tabs.executeScript). |extension| and |browser| are the source of
	//  the resource request. |file| is the requested relative file path. To
	//  handle the resource request return true(1) and execute |callback| either
	//  synchronously or asynchronously. For the default behavior which reads the
	//  resource from the extension directory on disk return false(0).
	//  Localization substitutions will not be applied to resources handled via
	//  this function.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h(cef_extension_handler_t)</a>
	SetOnExtensionGetExtensionResource(fn TOnGetExtensionResource) // property event
	// SetOnPrintStart
	//  {$IFDEF LINUX}
	//  Called when printing has started for the specified |browser|. This
	//  function will be called before the other OnPrint*() functions and
	//  irrespective of how printing was initiated(e.g.
	//  ICefBrowserHost.print(), JavaScript window.print() or PDF extension
	//  print button).
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h(cef_print_handler_t)</a>
	SetOnPrintStart(fn TOnPrintStart) // property event
	// SetOnPrintSettings
	//  Synchronize |settings| with client state. If |get_defaults| is true(1)
	//  then populate |settings| with the default print settings. Do not keep a
	//  reference to |settings| outside of this callback.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h(cef_print_handler_t)</a>
	SetOnPrintSettings(fn TOnPrintSettings) // property event
	// SetOnPrintDialog
	//  Show the print dialog. Execute |callback| once the dialog is dismissed.
	//  Return true(1) if the dialog will be displayed or false(0) to cancel the
	//  printing immediately.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h(cef_print_handler_t)</a>
	SetOnPrintDialog(fn TOnPrintDialog) // property event
	// SetOnPrintJob
	//  Send the print job to the printer. Execute |callback| once the job is
	//  completed. Return true(1) if the job will proceed or false(0) to cancel
	//  the job immediately.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h(cef_print_handler_t)</a>
	SetOnPrintJob(fn TOnPrintJob) // property event
	// SetOnPrintReset
	//  Reset client state related to printing.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h(cef_print_handler_t)</a>
	SetOnPrintReset(fn TOnPrintReset) // property event
	// SetOnGetPDFPaperSize
	//  Return the PDF paper size in device units. Used in combination with
	//  ICefBrowserHost.PrintToPdf().
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h(cef_print_handler_t)</a>
	SetOnGetPDFPaperSize(fn TOnGetPDFPaperSize) // property event
	// SetOnFrameCreated
	//  Called when a new frame is created. This will be the first notification
	//  that references |frame|. Any commands that require transport to the
	//  associated renderer process(LoadRequest, SendProcessMessage, GetSource,
	//  etc.) will be queued until OnFrameAttached is called for |frame|.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_handler_capi.h">CEF source file: /include/capi/cef_frame_handler_capi.h(cef_frame_handler_t)</a>
	SetOnFrameCreated(fn TOnFrameCreated) // property event
	// SetOnFrameAttached
	//  Called when a frame can begin routing commands to/from the associated
	//  renderer process. |reattached| will be true(1) if the frame was re-
	//  attached after exiting the BackForwardCache. Any commands that were queued
	//  have now been dispatched.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_handler_capi.h">CEF source file: /include/capi/cef_frame_handler_capi.h(cef_frame_handler_t)</a>
	SetOnFrameAttached(fn TOnFrameAttached) // property event
	// SetOnFrameDetached
	//  Called when a frame loses its connection to the renderer process and will
	//  be destroyed. Any pending or future commands will be discarded and
	//  ICefFrame.IsValid() will now return false(0) for |frame|. If called
	//  after ICefLifeSpanHandler.OnBeforeClose() during browser
	//  destruction then ICefBrowser.IsValid() will return false(0) for
	//  |browser|.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_handler_capi.h">CEF source file: /include/capi/cef_frame_handler_capi.h(cef_frame_handler_t)</a>
	SetOnFrameDetached(fn TOnFrameDetached) // property event
	// SetOnMainFrameChanged
	//  Called when the main frame changes due to(a) initial browser creation,
	//  (b) final browser destruction,(c) cross-origin navigation or(d) re-
	//  navigation after renderer process termination(due to crashes, etc).
	//  |old_frame| will be NULL and |new_frame| will be non-NULL when a main
	//  frame is assigned to |browser| for the first time. |old_frame| will be
	//  non-NULL and |new_frame| will be NULL and when a main frame is removed
	//  from |browser| for the last time. Both |old_frame| and |new_frame| will be
	//  non-NULL for cross-origin navigations or re-navigation after renderer
	//  process termination. This function will be called after on_frame_created()
	//  for |new_frame| and/or after OnFrameDetached() for |old_frame|. If
	//  called after ICefLifeSpanHandler.OnBeforeClose() during browser
	//  destruction then ICefBrowser.IsValid() will return false(0) for
	//  |browser|.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_handler_capi.h">CEF source file: /include/capi/cef_frame_handler_capi.h(cef_frame_handler_t)</a>
	SetOnMainFrameChanged(fn TOnMainFrameChanged) // property event
	// SetOnChromeCommand
	//  Called to execute a Chrome command triggered via menu selection or
	//  keyboard shortcut. Values for |command_id| can be found in the
	//  cef_command_ids.h file. |disposition| provides information about the
	//  intended command target. Return true(1) if the command was handled or
	//  false(0) for the default implementation. For context menu commands this
	//  will be called after ICefContextMenuHandler.OnContextMenuCommand.
	//  Only used with the Chrome runtime.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h(cef_command_handler_t)</a>
	SetOnChromeCommand(fn TOnChromeCommand) // property event
	// SetOnIsChromeAppMenuItemVisible
	//  Called to check if a Chrome app menu item should be visible. Values for
	//  |command_id| can be found in the cef_command_ids.h file. Only called for
	//  menu items that would be visible by default.
	//  Only used with the Chrome runtime.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h(cef_command_handler_t)</a>
	SetOnIsChromeAppMenuItemVisible(fn TOnIsChromeAppMenuItemVisible) // property event
	// SetOnIsChromeAppMenuItemEnabled
	//  Called to check if a Chrome app menu item should be enabled. Values for
	//  |command_id| can be found in the cef_command_ids.h file. Only called for
	//  menu items that would be enabled by default.
	//  Only used with the Chrome runtime.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h(cef_command_handler_t)</a>
	SetOnIsChromeAppMenuItemEnabled(fn TOnIsChromeAppMenuItemEnabled) // property event
	// SetOnIsChromePageActionIconVisible
	//  Called during browser creation to check if a Chrome page action icon
	//  should be visible. Only called for icons that would be visible by default.
	//  Only used with the Chrome runtime.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h(cef_command_handler_t)</a>
	SetOnIsChromePageActionIconVisible(fn TOnIsChromePageActionIconVisible) // property event
	// SetOnIsChromeToolbarButtonVisible
	//  Called during browser creation to check if a Chrome toolbar button should
	//  be visible. Only called for buttons that would be visible by default.
	//  Only used with the Chrome runtime.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h(cef_command_handler_t)</a>
	SetOnIsChromeToolbarButtonVisible(fn TOnIsChromeToolbarButtonVisible) // property event
	// SetOnRequestMediaAccessPermission
	//  Called when a page requests permission to access media.
	//  |requesting_origin| is the URL origin requesting permission.
	//  |requested_permissions| is a combination of values from
	//  TCefMediaAccessPermissionTypes that represent the requested
	//  permissions. Return true(1) and call ICefMediaAccessCallback
	//  functions either in this function or at a later time to continue or cancel
	//  the request. Return false(0) to proceed with default handling. With the
	//  Chrome runtime, default handling will display the permission request UI.
	//  With the Alloy runtime, default handling will deny the request. This
	//  function will not be called if the "--enable-media-stream" command-line
	//  switch is used to grant all permissions.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h(cef_permission_handler_t)</a>
	SetOnRequestMediaAccessPermission(fn TOnRequestMediaAccessPermission) // property event
	// SetOnShowPermissionPrompt
	//  Called when a page should show a permission prompt. |prompt_id| uniquely
	//  identifies the prompt. |requesting_origin| is the URL origin requesting
	//  permission. |requested_permissions| is a combination of values from
	//  TCefPermissionRequestTypes that represent the requested permissions.
	//  Return true(1) and call ICefPermissionPromptCallback.Continue either
	//  in this function or at a later time to continue or cancel the request.
	//  Return false(0) to proceed with default handling. With the Chrome
	//  runtime, default handling will display the permission prompt UI. With the
	//  Alloy runtime, default handling is CEF_PERMISSION_RESULT_IGNORE.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h(cef_permission_handler_t)</a>
	SetOnShowPermissionPrompt(fn TOnShowPermissionPrompt) // property event
	// SetOnDismissPermissionPrompt
	//  Called when a permission prompt handled via OnShowPermissionPrompt is
	//  dismissed. |prompt_id| will match the value that was passed to
	//  OnShowPermissionPrompt. |result| will be the value passed to
	//  ICefPermissionPromptCallback.Continue or CEF_PERMISSION_RESULT_IGNORE
	//  if the dialog was dismissed for other reasons such as navigation, browser
	//  closure, etc. This function will not be called if OnShowPermissionPrompt
	//  returned false(0) for |prompt_id|.
	//  This event will be called on the browser process CEF UI thread.
	//  <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h(cef_permission_handler_t)</a>
	SetOnDismissPermissionPrompt(fn TOnDismissPermissionPrompt) // property event
}

// TChromiumCore Parent: TComponent
//
//	Parent class of TChromium and TFMXChromium that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type TChromiumCore struct {
	TComponent
	textResultAvailablePtr                  uintptr
	pdfPrintFinishedPtr                     uintptr
	prefsAvailablePtr                       uintptr
	prefsUpdatedPtr                         uintptr
	cookiesDeletedPtr                       uintptr
	resolvedHostAvailablePtr                uintptr
	navigationVisitorResultAvailablePtr     uintptr
	downloadImageFinishedPtr                uintptr
	cookiesFlushedPtr                       uintptr
	certificateExceptionsClearedPtr         uintptr
	httpAuthCredentialsClearedPtr           uintptr
	allConnectionsClosedPtr                 uintptr
	executeTaskOnCefThreadPtr               uintptr
	cookiesVisitedPtr                       uintptr
	cookieVisitorDestroyedPtr               uintptr
	cookieSetPtr                            uintptr
	zoomPctAvailablePtr                     uintptr
	mediaRouteCreateFinishedPtr             uintptr
	mediaSinkDeviceInfoPtr                  uintptr
	canFocusPtr                             uintptr
	browserCompMsgPtr                       uintptr
	widgetCompMsgPtr                        uintptr
	renderCompMsgPtr                        uintptr
	processMessageReceivedPtr               uintptr
	loadStartPtr                            uintptr
	loadEndPtr                              uintptr
	loadErrorPtr                            uintptr
	loadingStateChangePtr                   uintptr
	takeFocusPtr                            uintptr
	setFocusPtr                             uintptr
	gotFocusPtr                             uintptr
	beforeContextMenuPtr                    uintptr
	runContextMenuPtr                       uintptr
	contextMenuCommandPtr                   uintptr
	contextMenuDismissedPtr                 uintptr
	runQuickMenuPtr                         uintptr
	quickMenuCommandPtr                     uintptr
	quickMenuDismissedPtr                   uintptr
	preKeyEventPtr                          uintptr
	keyEventPtr                             uintptr
	addressChangePtr                        uintptr
	titleChangePtr                          uintptr
	favIconUrlChangePtr                     uintptr
	fullScreenModeChangePtr                 uintptr
	tooltipPtr                              uintptr
	statusMessagePtr                        uintptr
	consoleMessagePtr                       uintptr
	autoResizePtr                           uintptr
	loadingProgressChangePtr                uintptr
	cursorChangePtr                         uintptr
	mediaAccessChangePtr                    uintptr
	canDownloadPtr                          uintptr
	beforeDownloadPtr                       uintptr
	downloadUpdatedPtr                      uintptr
	jsdialogPtr                             uintptr
	beforeUnloadDialogPtr                   uintptr
	resetDialogStatePtr                     uintptr
	dialogClosedPtr                         uintptr
	beforePopupPtr                          uintptr
	afterCreatedPtr                         uintptr
	beforeClosePtr                          uintptr
	closePtr                                uintptr
	beforeBrowsePtr                         uintptr
	openUrlFromTabPtr                       uintptr
	getAuthCredentialsPtr                   uintptr
	certificateErrorPtr                     uintptr
	selectClientCertificatePtr              uintptr
	renderViewReadyPtr                      uintptr
	renderProcessTerminatedPtr              uintptr
	getResourceRequestHandler_ReqHdlrPtr    uintptr
	documentAvailableInMainFramePtr         uintptr
	beforeResourceLoadPtr                   uintptr
	getResourceHandlerPtr                   uintptr
	resourceRedirectPtr                     uintptr
	resourceResponsePtr                     uintptr
	getResourceResponseFilterPtr            uintptr
	resourceLoadCompletePtr                 uintptr
	protocolExecutionPtr                    uintptr
	canSendCookiePtr                        uintptr
	canSaveCookiePtr                        uintptr
	fileDialogPtr                           uintptr
	getAccessibilityHandlerPtr              uintptr
	getRootScreenRectPtr                    uintptr
	getViewRectPtr                          uintptr
	getScreenPointPtr                       uintptr
	getScreenInfoPtr                        uintptr
	popupShowPtr                            uintptr
	popupSizePtr                            uintptr
	paintPtr                                uintptr
	acceleratedPaintPtr                     uintptr
	getTouchHandleSizePtr                   uintptr
	touchHandleStateChangedPtr              uintptr
	startDraggingPtr                        uintptr
	updateDragCursorPtr                     uintptr
	scrollOffsetChangedPtr                  uintptr
	iMECompositionRangeChangedPtr           uintptr
	textSelectionChangedPtr                 uintptr
	virtualKeyboardRequestedPtr             uintptr
	dragEnterPtr                            uintptr
	draggableRegionsChangedPtr              uintptr
	findResultPtr                           uintptr
	requestContextInitializedPtr            uintptr
	getResourceRequestHandler_ReqCtxHdlrPtr uintptr
	sinksPtr                                uintptr
	routesPtr                               uintptr
	routeStateChangedPtr                    uintptr
	routeMessageReceivedPtr                 uintptr
	getAudioParametersPtr                   uintptr
	audioStreamStartedPtr                   uintptr
	audioStreamPacketPtr                    uintptr
	audioStreamStoppedPtr                   uintptr
	audioStreamErrorPtr                     uintptr
	devToolsMessagePtr                      uintptr
	devToolsRawMessagePtr                   uintptr
	devToolsMethodResultPtr                 uintptr
	devToolsMethodRawResultPtr              uintptr
	devToolsEventPtr                        uintptr
	devToolsRawEventPtr                     uintptr
	devToolsAgentAttachedPtr                uintptr
	devToolsAgentDetachedPtr                uintptr
	extensionLoadFailedPtr                  uintptr
	extensionLoadedPtr                      uintptr
	extensionUnloadedPtr                    uintptr
	extensionBeforeBackgroundBrowserPtr     uintptr
	extensionBeforeBrowserPtr               uintptr
	extensionGetActiveBrowserPtr            uintptr
	extensionCanAccessBrowserPtr            uintptr
	extensionGetExtensionResourcePtr        uintptr
	printStartPtr                           uintptr
	printSettingsPtr                        uintptr
	printDialogPtr                          uintptr
	printJobPtr                             uintptr
	printResetPtr                           uintptr
	getPDFPaperSizePtr                      uintptr
	frameCreatedPtr                         uintptr
	frameAttachedPtr                        uintptr
	frameDetachedPtr                        uintptr
	mainFrameChangedPtr                     uintptr
	chromeCommandPtr                        uintptr
	isChromeAppMenuItemVisiblePtr           uintptr
	isChromeAppMenuItemEnabledPtr           uintptr
	isChromePageActionIconVisiblePtr        uintptr
	isChromeToolbarButtonVisiblePtr         uintptr
	requestMediaAccessPermissionPtr         uintptr
	showPermissionPromptPtr                 uintptr
	dismissPermissionPromptPtr              uintptr
}

func NewChromiumCore(aOwner IComponent) IChromiumCore {
	r1 := CEF().SysCallN(1748, GetObjectUintptr(aOwner))
	return AsChromiumCore(r1)
}

func (m *TChromiumCore) DefaultUrl() string {
	r1 := CEF().SysCallN(1762, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetDefaultUrl(AValue string) {
	CEF().SysCallN(1762, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) Options() IChromiumOptions {
	var resultChromiumOptions uintptr
	CEF().SysCallN(1851, 0, m.Instance(), 0, uintptr(unsafePointer(&resultChromiumOptions)))
	return AsChromiumOptions(resultChromiumOptions)
}

func (m *TChromiumCore) SetOptions(AValue IChromiumOptions) {
	CEF().SysCallN(1851, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TChromiumCore) FontOptions() IChromiumFontOptions {
	r1 := CEF().SysCallN(1789, 0, m.Instance(), 0)
	return AsChromiumFontOptions(r1)
}

func (m *TChromiumCore) SetFontOptions(AValue IChromiumFontOptions) {
	CEF().SysCallN(1789, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TChromiumCore) PDFPrintOptions() IPDFPrintOptions {
	r1 := CEF().SysCallN(1852, 0, m.Instance(), 0)
	return AsPDFPrintOptions(r1)
}

func (m *TChromiumCore) SetPDFPrintOptions(AValue IPDFPrintOptions) {
	CEF().SysCallN(1852, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TChromiumCore) BrowserId() int32 {
	r1 := CEF().SysCallN(1722, m.Instance())
	return int32(r1)
}

func (m *TChromiumCore) Browser() ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(1718, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TChromiumCore) BrowserById(id int32) ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(1719, m.Instance(), uintptr(id), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TChromiumCore) BrowserCount() int32 {
	r1 := CEF().SysCallN(1720, m.Instance())
	return int32(r1)
}

func (m *TChromiumCore) BrowserIdByIndex(i int32) int32 {
	r1 := CEF().SysCallN(1723, m.Instance(), uintptr(i))
	return int32(r1)
}

func (m *TChromiumCore) CefClient() ICefClient {
	var resultCefClient uintptr
	CEF().SysCallN(1730, m.Instance(), uintptr(unsafePointer(&resultCefClient)))
	return AsCefClient(resultCefClient)
}

func (m *TChromiumCore) ReqContextHandler() ICefRequestContextHandler {
	var resultCefRequestContextHandler uintptr
	CEF().SysCallN(1870, m.Instance(), uintptr(unsafePointer(&resultCefRequestContextHandler)))
	return AsCefRequestContextHandler(resultCefRequestContextHandler)
}

func (m *TChromiumCore) ResourceRequestHandler() ICefResourceRequestHandler {
	var resultCefResourceRequestHandler uintptr
	CEF().SysCallN(1879, m.Instance(), uintptr(unsafePointer(&resultCefResourceRequestHandler)))
	return AsCefResourceRequestHandler(resultCefResourceRequestHandler)
}

func (m *TChromiumCore) CefWindowInfo() (resultCefWindowInfo TCefWindowInfo) {
	r1 := CEF().SysCallN(1731, m.Instance())
	return *(*TCefWindowInfo)(unsafePointer(r1))
}

func (m *TChromiumCore) VisibleNavigationEntry() ICefNavigationEntry {
	var resultCefNavigationEntry uintptr
	CEF().SysCallN(2068, m.Instance(), uintptr(unsafePointer(&resultCefNavigationEntry)))
	return AsCefNavigationEntry(resultCefNavigationEntry)
}

func (m *TChromiumCore) RequestContext() ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(1871, m.Instance(), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *TChromiumCore) MediaRouter() ICefMediaRouter {
	var resultCefMediaRouter uintptr
	CEF().SysCallN(1842, m.Instance(), uintptr(unsafePointer(&resultCefMediaRouter)))
	return AsCefMediaRouter(resultCefMediaRouter)
}

func (m *TChromiumCore) MediaObserver() ICefMediaObserver {
	var resultCefMediaObserver uintptr
	CEF().SysCallN(1840, m.Instance(), uintptr(unsafePointer(&resultCefMediaObserver)))
	return AsCefMediaObserver(resultCefMediaObserver)
}

func (m *TChromiumCore) MediaObserverReg() ICefRegistration {
	var resultCefRegistration uintptr
	CEF().SysCallN(1841, m.Instance(), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TChromiumCore) DevToolsMsgObserver() ICefDevToolsMessageObserver {
	var resultCefDevToolsMessageObserver uintptr
	CEF().SysCallN(1766, m.Instance(), uintptr(unsafePointer(&resultCefDevToolsMessageObserver)))
	return AsCefDevToolsMessageObserver(resultCefDevToolsMessageObserver)
}

func (m *TChromiumCore) DevToolsMsgObserverReg() ICefRegistration {
	var resultCefRegistration uintptr
	CEF().SysCallN(1767, m.Instance(), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TChromiumCore) ExtensionHandler() ICefExtensionHandler {
	var resultCefExtensionHandler uintptr
	CEF().SysCallN(1786, m.Instance(), uintptr(unsafePointer(&resultCefExtensionHandler)))
	return AsCefExtensionHandler(resultCefExtensionHandler)
}

func (m *TChromiumCore) MultithreadApp() bool {
	r1 := CEF().SysCallN(1844, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) IsLoading() bool {
	r1 := CEF().SysCallN(1823, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) HasDocument() bool {
	r1 := CEF().SysCallN(1807, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) HasView() bool {
	r1 := CEF().SysCallN(1810, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) HasDevTools() bool {
	r1 := CEF().SysCallN(1806, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) HasClientHandler() bool {
	r1 := CEF().SysCallN(1805, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) HasBrowser() bool {
	r1 := CEF().SysCallN(1804, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) CanGoBack() bool {
	r1 := CEF().SysCallN(1726, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) CanGoForward() bool {
	r1 := CEF().SysCallN(1727, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) IsPopUp() bool {
	r1 := CEF().SysCallN(1824, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) WindowHandle() TCefWindowHandle {
	r1 := CEF().SysCallN(2077, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TChromiumCore) OpenerWindowHandle() TCefWindowHandle {
	r1 := CEF().SysCallN(1850, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TChromiumCore) BrowserHandle() THandle {
	r1 := CEF().SysCallN(1721, m.Instance())
	return THandle(r1)
}

func (m *TChromiumCore) WidgetHandle() THandle {
	r1 := CEF().SysCallN(2076, m.Instance())
	return THandle(r1)
}

func (m *TChromiumCore) RenderHandle() THandle {
	r1 := CEF().SysCallN(1868, m.Instance())
	return THandle(r1)
}

func (m *TChromiumCore) FrameIsFocused() bool {
	r1 := CEF().SysCallN(1791, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) Initialized() bool {
	r1 := CEF().SysCallN(1821, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) RequestContextCache() string {
	r1 := CEF().SysCallN(1872, m.Instance())
	return GoStr(r1)
}

func (m *TChromiumCore) RequestContextIsGlobal() bool {
	r1 := CEF().SysCallN(1873, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) DocumentURL() string {
	r1 := CEF().SysCallN(1770, m.Instance())
	return GoStr(r1)
}

func (m *TChromiumCore) ZoomLevel() (resultFloat64 float64) {
	CEF().SysCallN(2080, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TChromiumCore) SetZoomLevel(AValue float64) {
	CEF().SysCallN(2080, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TChromiumCore) DefaultZoomLevel() (resultFloat64 float64) {
	CEF().SysCallN(1764, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TChromiumCore) CanIncZoom() bool {
	r1 := CEF().SysCallN(1728, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) CanDecZoom() bool {
	r1 := CEF().SysCallN(1724, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) CanResetZoom() bool {
	r1 := CEF().SysCallN(1729, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) ZoomPct() (resultFloat64 float64) {
	CEF().SysCallN(2081, 0, m.Instance(), uintptr(unsafePointer(&resultFloat64)), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TChromiumCore) SetZoomPct(AValue float64) {
	CEF().SysCallN(2081, 1, m.Instance(), uintptr(unsafePointer(&AValue)), uintptr(unsafePointer(&AValue)))
}

func (m *TChromiumCore) ZoomStep() byte {
	r1 := CEF().SysCallN(2082, 0, m.Instance(), 0)
	return byte(r1)
}

func (m *TChromiumCore) SetZoomStep(AValue byte) {
	CEF().SysCallN(2082, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) WindowlessFrameRate() int32 {
	r1 := CEF().SysCallN(2078, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumCore) SetWindowlessFrameRate(AValue int32) {
	CEF().SysCallN(2078, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) CustomHeaderName() string {
	r1 := CEF().SysCallN(1757, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetCustomHeaderName(AValue string) {
	CEF().SysCallN(1757, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) CustomHeaderValue() string {
	r1 := CEF().SysCallN(1758, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetCustomHeaderValue(AValue string) {
	CEF().SysCallN(1758, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) DoNotTrack() bool {
	r1 := CEF().SysCallN(1769, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetDoNotTrack(AValue bool) {
	CEF().SysCallN(1769, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) SendReferrer() bool {
	r1 := CEF().SysCallN(1900, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetSendReferrer(AValue bool) {
	CEF().SysCallN(1900, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) HyperlinkAuditing() bool {
	r1 := CEF().SysCallN(1812, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetHyperlinkAuditing(AValue bool) {
	CEF().SysCallN(1812, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) AllowOutdatedPlugins() bool {
	r1 := CEF().SysCallN(1712, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetAllowOutdatedPlugins(AValue bool) {
	CEF().SysCallN(1712, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) AlwaysAuthorizePlugins() bool {
	r1 := CEF().SysCallN(1713, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetAlwaysAuthorizePlugins(AValue bool) {
	CEF().SysCallN(1713, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) AlwaysOpenPDFExternally() bool {
	r1 := CEF().SysCallN(1714, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetAlwaysOpenPDFExternally(AValue bool) {
	CEF().SysCallN(1714, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) SpellChecking() bool {
	r1 := CEF().SysCallN(2059, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetSpellChecking(AValue bool) {
	CEF().SysCallN(2059, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) SpellCheckerDicts() string {
	r1 := CEF().SysCallN(2058, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetSpellCheckerDicts(AValue string) {
	CEF().SysCallN(2058, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) HasValidMainFrame() bool {
	r1 := CEF().SysCallN(1809, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) FrameCount() NativeUInt {
	r1 := CEF().SysCallN(1790, m.Instance())
	return NativeUInt(r1)
}

func (m *TChromiumCore) DragOperations() TCefDragOperations {
	r1 := CEF().SysCallN(1772, 0, m.Instance(), 0)
	return TCefDragOperations(r1)
}

func (m *TChromiumCore) SetDragOperations(AValue TCefDragOperations) {
	CEF().SysCallN(1772, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) AudioMuted() bool {
	r1 := CEF().SysCallN(1715, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetAudioMuted(AValue bool) {
	CEF().SysCallN(1715, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) Fullscreen() bool {
	r1 := CEF().SysCallN(1792, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) SafeSearch() bool {
	r1 := CEF().SysCallN(1886, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetSafeSearch(AValue bool) {
	CEF().SysCallN(1886, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) YouTubeRestrict() int32 {
	r1 := CEF().SysCallN(2079, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumCore) SetYouTubeRestrict(AValue int32) {
	CEF().SysCallN(2079, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) PrintingEnabled() bool {
	r1 := CEF().SysCallN(1855, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetPrintingEnabled(AValue bool) {
	CEF().SysCallN(1855, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) AcceptLanguageList() string {
	r1 := CEF().SysCallN(1708, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetAcceptLanguageList(AValue string) {
	CEF().SysCallN(1708, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) AcceptCookies() TCefCookiePref {
	r1 := CEF().SysCallN(1707, 0, m.Instance(), 0)
	return TCefCookiePref(r1)
}

func (m *TChromiumCore) SetAcceptCookies(AValue TCefCookiePref) {
	CEF().SysCallN(1707, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) Block3rdPartyCookies() bool {
	r1 := CEF().SysCallN(1717, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetBlock3rdPartyCookies(AValue bool) {
	CEF().SysCallN(1717, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) MultiBrowserMode() bool {
	r1 := CEF().SysCallN(1843, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetMultiBrowserMode(AValue bool) {
	CEF().SysCallN(1843, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) DefaultWindowInfoExStyle() DWORD {
	r1 := CEF().SysCallN(1763, 0, m.Instance(), 0)
	return DWORD(r1)
}

func (m *TChromiumCore) SetDefaultWindowInfoExStyle(AValue DWORD) {
	CEF().SysCallN(1763, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) Offline() bool {
	r1 := CEF().SysCallN(1849, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetOffline(AValue bool) {
	CEF().SysCallN(1849, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) QuicAllowed() bool {
	r1 := CEF().SysCallN(1864, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetQuicAllowed(AValue bool) {
	CEF().SysCallN(1864, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) JavascriptEnabled() bool {
	r1 := CEF().SysCallN(1826, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetJavascriptEnabled(AValue bool) {
	CEF().SysCallN(1826, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) LoadImagesAutomatically() bool {
	r1 := CEF().SysCallN(1828, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetLoadImagesAutomatically(AValue bool) {
	CEF().SysCallN(1828, 1, m.Instance(), PascalBool(AValue))
}

func (m *TChromiumCore) BatterySaverModeState() TCefBatterySaverModeState {
	r1 := CEF().SysCallN(1716, 0, m.Instance(), 0)
	return TCefBatterySaverModeState(r1)
}

func (m *TChromiumCore) SetBatterySaverModeState(AValue TCefBatterySaverModeState) {
	CEF().SysCallN(1716, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) HighEfficiencyModeState() TCefHighEfficiencyModeState {
	r1 := CEF().SysCallN(1811, 0, m.Instance(), 0)
	return TCefHighEfficiencyModeState(r1)
}

func (m *TChromiumCore) SetHighEfficiencyModeState(AValue TCefHighEfficiencyModeState) {
	CEF().SysCallN(1811, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) CanFocus() bool {
	r1 := CEF().SysCallN(1725, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) EnableFocusDelayMs() uint32 {
	r1 := CEF().SysCallN(1779, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TChromiumCore) SetEnableFocusDelayMs(AValue uint32) {
	CEF().SysCallN(1779, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) WebRTCIPHandlingPolicy() TCefWebRTCHandlingPolicy {
	r1 := CEF().SysCallN(2073, 0, m.Instance(), 0)
	return TCefWebRTCHandlingPolicy(r1)
}

func (m *TChromiumCore) SetWebRTCIPHandlingPolicy(AValue TCefWebRTCHandlingPolicy) {
	CEF().SysCallN(2073, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) WebRTCMultipleRoutes() TCefState {
	r1 := CEF().SysCallN(2074, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumCore) SetWebRTCMultipleRoutes(AValue TCefState) {
	CEF().SysCallN(2074, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) WebRTCNonproxiedUDP() TCefState {
	r1 := CEF().SysCallN(2075, 0, m.Instance(), 0)
	return TCefState(r1)
}

func (m *TChromiumCore) SetWebRTCNonproxiedUDP(AValue TCefState) {
	CEF().SysCallN(2075, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) ProxyType() int32 {
	r1 := CEF().SysCallN(1862, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumCore) SetProxyType(AValue int32) {
	CEF().SysCallN(1862, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) ProxyScheme() TCefProxyScheme {
	r1 := CEF().SysCallN(1859, 0, m.Instance(), 0)
	return TCefProxyScheme(r1)
}

func (m *TChromiumCore) SetProxyScheme(AValue TCefProxyScheme) {
	CEF().SysCallN(1859, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) ProxyServer() string {
	r1 := CEF().SysCallN(1861, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetProxyServer(AValue string) {
	CEF().SysCallN(1861, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) ProxyPort() int32 {
	r1 := CEF().SysCallN(1858, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumCore) SetProxyPort(AValue int32) {
	CEF().SysCallN(1858, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) ProxyUsername() string {
	r1 := CEF().SysCallN(1863, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetProxyUsername(AValue string) {
	CEF().SysCallN(1863, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) ProxyPassword() string {
	r1 := CEF().SysCallN(1857, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetProxyPassword(AValue string) {
	CEF().SysCallN(1857, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) ProxyScriptURL() string {
	r1 := CEF().SysCallN(1860, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetProxyScriptURL(AValue string) {
	CEF().SysCallN(1860, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) ProxyByPassList() string {
	r1 := CEF().SysCallN(1856, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TChromiumCore) SetProxyByPassList(AValue string) {
	CEF().SysCallN(1856, 1, m.Instance(), PascalStr(AValue))
}

func (m *TChromiumCore) MaxConnectionsPerProxy() int32 {
	r1 := CEF().SysCallN(1839, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TChromiumCore) SetMaxConnectionsPerProxy(AValue int32) {
	CEF().SysCallN(1839, 1, m.Instance(), uintptr(AValue))
}

func (m *TChromiumCore) CreateClientHandler(aIsOSR bool) bool {
	r1 := CEF().SysCallN(1751, m.Instance(), PascalBool(aIsOSR))
	return GoBool(r1)
}

func (m *TChromiumCore) CreateClientHandler1(aClient *ICefClient, aIsOSR bool) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(1752, m.Instance(), uintptr(unsafePointer(&result0)), PascalBool(aIsOSR))
	*aClient = AsCefClient(result0)
	return GoBool(r1)
}

func (m *TChromiumCore) TryCloseBrowser() bool {
	r1 := CEF().SysCallN(2064, m.Instance())
	return GoBool(r1)
}

func (m *TChromiumCore) SelectBrowser(aID int32) bool {
	r1 := CEF().SysCallN(1889, m.Instance(), uintptr(aID))
	return GoBool(r1)
}

func (m *TChromiumCore) IndexOfBrowserID(aID int32) int32 {
	r1 := CEF().SysCallN(1819, m.Instance(), uintptr(aID))
	return int32(r1)
}

func (m *TChromiumCore) ShareRequestContext(aContext *ICefRequestContext, aHandler ICefRequestContextHandler) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(2054, m.Instance(), uintptr(unsafePointer(&result0)), GetObjectUintptr(aHandler))
	*aContext = AsCefRequestContext(result0)
	return GoBool(r1)
}

func (m *TChromiumCore) SetNewBrowserParent(aNewParentHwnd HWND) bool {
	r1 := CEF().SysCallN(1905, m.Instance(), uintptr(aNewParentHwnd))
	return GoBool(r1)
}

func (m *TChromiumCore) CreateBrowser(aParentHandle TCefWindowHandle, aParentRect *TRect, aWindowName string, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue, aForceAsPopup bool) bool {
	r1 := CEF().SysCallN(1749, m.Instance(), uintptr(aParentHandle), uintptr(unsafePointer(aParentRect)), PascalStr(aWindowName), GetObjectUintptr(aContext), GetObjectUintptr(aExtraInfo), PascalBool(aForceAsPopup))
	return GoBool(r1)
}

func (m *TChromiumCore) CreateBrowser1(aURL string, aBrowserViewComp ICEFBrowserViewComponent, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(1750, m.Instance(), PascalStr(aURL), GetObjectUintptr(aBrowserViewComp), GetObjectUintptr(aContext), GetObjectUintptr(aExtraInfo))
	return GoBool(r1)
}

func (m *TChromiumCore) ClearCertificateExceptions(aClearImmediately bool) bool {
	r1 := CEF().SysCallN(1734, m.Instance(), PascalBool(aClearImmediately))
	return GoBool(r1)
}

func (m *TChromiumCore) ClearHttpAuthCredentials(aClearImmediately bool) bool {
	r1 := CEF().SysCallN(1736, m.Instance(), PascalBool(aClearImmediately))
	return GoBool(r1)
}

func (m *TChromiumCore) CloseAllConnections(aCloseImmediately bool) bool {
	r1 := CEF().SysCallN(1744, m.Instance(), PascalBool(aCloseImmediately))
	return GoBool(r1)
}

func (m *TChromiumCore) GetFrameNames(aFrameNames *IStrings) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(1798, m.Instance(), uintptr(unsafePointer(&result0)))
	*aFrameNames = AsStrings(result0)
	return GoBool(r1)
}

func (m *TChromiumCore) GetFrameIdentifiers(aFrameCount *NativeUInt, aFrameIdentifierArray *ICefFrameIdentifierArray) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := CEF().SysCallN(1797, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*aFrameCount = NativeUInt(result0)
	*aFrameIdentifierArray = FrameIdentifierArrayRef.New(int(*aFrameCount), result1)
	return GoBool(r1)
}

func (m *TChromiumCore) IsSameBrowser(aBrowser ICefBrowser) bool {
	r1 := CEF().SysCallN(1825, m.Instance(), GetObjectUintptr(aBrowser))
	return GoBool(r1)
}

func (m *TChromiumCore) ExecuteTaskOnCefThread(aCefThreadId TCefThreadId, aTaskID uint32, aDelayMs int64) bool {
	r1 := CEF().SysCallN(1784, m.Instance(), uintptr(aCefThreadId), uintptr(aTaskID), uintptr(unsafePointer(&aDelayMs)))
	return GoBool(r1)
}

func (m *TChromiumCore) DeleteCookies(url string, cookieName string, aDeleteImmediately bool) bool {
	r1 := CEF().SysCallN(1765, m.Instance(), PascalStr(url), PascalStr(cookieName), PascalBool(aDeleteImmediately))
	return GoBool(r1)
}

func (m *TChromiumCore) VisitAllCookies(aID int32) bool {
	r1 := CEF().SysCallN(2069, m.Instance(), uintptr(aID))
	return GoBool(r1)
}

func (m *TChromiumCore) VisitURLCookies(url string, includeHttpOnly bool, aID int32) bool {
	r1 := CEF().SysCallN(2070, m.Instance(), PascalStr(url), PascalBool(includeHttpOnly), uintptr(aID))
	return GoBool(r1)
}

func (m *TChromiumCore) FlushCookieStore(aFlushImmediately bool) bool {
	r1 := CEF().SysCallN(1788, m.Instance(), PascalBool(aFlushImmediately))
	return GoBool(r1)
}

func (m *TChromiumCore) SendDevToolsMessage(message string) bool {
	r1 := CEF().SysCallN(1891, m.Instance(), PascalStr(message))
	return GoBool(r1)
}

func (m *TChromiumCore) ExecuteDevToolsMethod(messageid int32, method string, params ICefDictionaryValue) int32 {
	r1 := CEF().SysCallN(1780, m.Instance(), uintptr(messageid), PascalStr(method), GetObjectUintptr(params))
	return int32(r1)
}

func (m *TChromiumCore) AddDevToolsMessageObserver(observer ICefDevToolsMessageObserver) ICefRegistration {
	var resultCefRegistration uintptr
	CEF().SysCallN(1709, m.Instance(), GetObjectUintptr(observer), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TChromiumCore) CreateUrlRequest(request ICefRequest, client ICefUrlRequestClient, aFrameName string) ICefUrlRequest {
	var resultCefUrlRequest uintptr
	CEF().SysCallN(1754, m.Instance(), GetObjectUintptr(request), GetObjectUintptr(client), PascalStr(aFrameName), uintptr(unsafePointer(&resultCefUrlRequest)))
	return AsCefUrlRequest(resultCefUrlRequest)
}

func (m *TChromiumCore) CreateUrlRequest1(request ICefRequest, client ICefUrlRequestClient, aFrame ICefFrame) ICefUrlRequest {
	var resultCefUrlRequest uintptr
	CEF().SysCallN(1755, m.Instance(), GetObjectUintptr(request), GetObjectUintptr(client), GetObjectUintptr(aFrame), uintptr(unsafePointer(&resultCefUrlRequest)))
	return AsCefUrlRequest(resultCefUrlRequest)
}

func (m *TChromiumCore) CreateUrlRequest2(request ICefRequest, client ICefUrlRequestClient, aFrameIdentifier int64) ICefUrlRequest {
	var resultCefUrlRequest uintptr
	CEF().SysCallN(1756, m.Instance(), GetObjectUintptr(request), GetObjectUintptr(client), uintptr(unsafePointer(&aFrameIdentifier)), uintptr(unsafePointer(&resultCefUrlRequest)))
	return AsCefUrlRequest(resultCefUrlRequest)
}

func (m *TChromiumCore) AddObserver(observer ICefMediaObserver) ICefRegistration {
	var resultCefRegistration uintptr
	CEF().SysCallN(1710, m.Instance(), GetObjectUintptr(observer), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TChromiumCore) GetSource(urn string) ICefMediaSource {
	var resultCefMediaSource uintptr
	CEF().SysCallN(1800, m.Instance(), PascalStr(urn), uintptr(unsafePointer(&resultCefMediaSource)))
	return AsCefMediaSource(resultCefMediaSource)
}

func (m *TChromiumCore) LoadExtension(rootdirectory string, manifest ICefDictionaryValue, handler ICefExtensionHandler, requestContext ICefRequestContext) bool {
	r1 := CEF().SysCallN(1827, m.Instance(), PascalStr(rootdirectory), GetObjectUintptr(manifest), GetObjectUintptr(handler), GetObjectUintptr(requestContext))
	return GoBool(r1)
}

func (m *TChromiumCore) DidLoadExtension(extensionid string) bool {
	r1 := CEF().SysCallN(1768, m.Instance(), PascalStr(extensionid))
	return GoBool(r1)
}

func (m *TChromiumCore) HasExtension(extensionid string) bool {
	r1 := CEF().SysCallN(1808, m.Instance(), PascalStr(extensionid))
	return GoBool(r1)
}

func (m *TChromiumCore) GetExtensions(extensionids IStringList) bool {
	r1 := CEF().SysCallN(1796, m.Instance(), GetObjectUintptr(extensionids))
	return GoBool(r1)
}

func (m *TChromiumCore) GetExtension(extensionid string) ICefExtension {
	var resultCefExtension uintptr
	CEF().SysCallN(1795, m.Instance(), PascalStr(extensionid), uintptr(unsafePointer(&resultCefExtension)))
	return AsCefExtension(resultCefExtension)
}

func (m *TChromiumCore) GetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) ICefValue {
	var resultCefValue uintptr
	CEF().SysCallN(1801, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype), uintptr(unsafePointer(&resultCefValue)))
	return AsCefValue(resultCefValue)
}

func (m *TChromiumCore) GetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes) TCefContentSettingValues {
	r1 := CEF().SysCallN(1793, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype))
	return TCefContentSettingValues(r1)
}

func ChromiumCoreClass() TClass {
	ret := CEF().SysCallN(1732)
	return TClass(ret)
}

func (m *TChromiumCore) CloseBrowser(aForceClose bool) {
	CEF().SysCallN(1745, m.Instance(), PascalBool(aForceClose))
}

func (m *TChromiumCore) CloseAllBrowsers() {
	CEF().SysCallN(1743, m.Instance())
}

func (m *TChromiumCore) InitializeDragAndDrop(aDropTargetWnd HWND) {
	CEF().SysCallN(1820, m.Instance(), uintptr(aDropTargetWnd))
}

func (m *TChromiumCore) ShutdownDragAndDrop() {
	CEF().SysCallN(2056, m.Instance())
}

func (m *TChromiumCore) LoadURL(aURL string, aFrameName string) {
	CEF().SysCallN(1836, m.Instance(), PascalStr(aURL), PascalStr(aFrameName))
}

func (m *TChromiumCore) LoadURL1(aURL string, aFrame ICefFrame) {
	CEF().SysCallN(1837, m.Instance(), PascalStr(aURL), GetObjectUintptr(aFrame))
}

func (m *TChromiumCore) LoadURL2(aURL string, aFrameIdentifier int64) {
	CEF().SysCallN(1838, m.Instance(), PascalStr(aURL), uintptr(unsafePointer(&aFrameIdentifier)))
}

func (m *TChromiumCore) LoadString(aHTML string, aFrameName string) {
	CEF().SysCallN(1833, m.Instance(), PascalStr(aHTML), PascalStr(aFrameName))
}

func (m *TChromiumCore) LoadString1(aHTML string, aFrame ICefFrame) {
	CEF().SysCallN(1834, m.Instance(), PascalStr(aHTML), GetObjectUintptr(aFrame))
}

func (m *TChromiumCore) LoadString2(aHTML string, aFrameIdentifier int64) {
	CEF().SysCallN(1835, m.Instance(), PascalStr(aHTML), uintptr(unsafePointer(&aFrameIdentifier)))
}

func (m *TChromiumCore) LoadResource(aStream ICustomMemoryStream, aMimeType, aCharset string, aFrameName string) {
	CEF().SysCallN(1830, m.Instance(), GetObjectUintptr(aStream), PascalStr(aMimeType), PascalStr(aCharset), PascalStr(aFrameName))
}

func (m *TChromiumCore) LoadResource1(aStream ICustomMemoryStream, aMimeType, aCharset string, aFrame ICefFrame) {
	CEF().SysCallN(1831, m.Instance(), GetObjectUintptr(aStream), PascalStr(aMimeType), PascalStr(aCharset), GetObjectUintptr(aFrame))
}

func (m *TChromiumCore) LoadResource2(aStream ICustomMemoryStream, aMimeType, aCharset string, aFrameIdentifier int64) {
	CEF().SysCallN(1832, m.Instance(), GetObjectUintptr(aStream), PascalStr(aMimeType), PascalStr(aCharset), uintptr(unsafePointer(&aFrameIdentifier)))
}

func (m *TChromiumCore) LoadRequest(aRequest ICefRequest) {
	CEF().SysCallN(1829, m.Instance(), GetObjectUintptr(aRequest))
}

func (m *TChromiumCore) GoBack() {
	CEF().SysCallN(1802, m.Instance())
}

func (m *TChromiumCore) GoForward() {
	CEF().SysCallN(1803, m.Instance())
}

func (m *TChromiumCore) Reload() {
	CEF().SysCallN(1866, m.Instance())
}

func (m *TChromiumCore) ReloadIgnoreCache() {
	CEF().SysCallN(1867, m.Instance())
}

func (m *TChromiumCore) StopLoad() {
	CEF().SysCallN(2062, m.Instance())
}

func (m *TChromiumCore) StartDownload(aURL string) {
	CEF().SysCallN(2060, m.Instance(), PascalStr(aURL))
}

func (m *TChromiumCore) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool) {
	CEF().SysCallN(1771, m.Instance(), PascalStr(imageUrl), PascalBool(isFavicon), uintptr(maxImageSize), PascalBool(bypassCache))
}

func (m *TChromiumCore) SimulateMouseWheel(aDeltaX, aDeltaY int32) {
	CEF().SysCallN(2057, m.Instance(), uintptr(aDeltaX), uintptr(aDeltaY))
}

func (m *TChromiumCore) RetrieveHTML(aFrameName string) {
	CEF().SysCallN(1880, m.Instance(), PascalStr(aFrameName))
}

func (m *TChromiumCore) RetrieveHTML1(aFrame ICefFrame) {
	CEF().SysCallN(1881, m.Instance(), GetObjectUintptr(aFrame))
}

func (m *TChromiumCore) RetrieveHTML2(aFrameIdentifier int64) {
	CEF().SysCallN(1882, m.Instance(), uintptr(unsafePointer(&aFrameIdentifier)))
}

func (m *TChromiumCore) RetrieveText(aFrameName string) {
	CEF().SysCallN(1883, m.Instance(), PascalStr(aFrameName))
}

func (m *TChromiumCore) RetrieveText1(aFrame ICefFrame) {
	CEF().SysCallN(1884, m.Instance(), GetObjectUintptr(aFrame))
}

func (m *TChromiumCore) RetrieveText2(aFrameIdentifier int64) {
	CEF().SysCallN(1885, m.Instance(), uintptr(unsafePointer(&aFrameIdentifier)))
}

func (m *TChromiumCore) GetNavigationEntries(currentOnly bool) {
	CEF().SysCallN(1799, m.Instance(), PascalBool(currentOnly))
}

func (m *TChromiumCore) ExecuteJavaScript(aCode, aScriptURL string, aFrameName string, aStartLine int32) {
	CEF().SysCallN(1781, m.Instance(), PascalStr(aCode), PascalStr(aScriptURL), PascalStr(aFrameName), uintptr(aStartLine))
}

func (m *TChromiumCore) ExecuteJavaScript1(aCode, aScriptURL string, aFrame ICefFrame, aStartLine int32) {
	CEF().SysCallN(1782, m.Instance(), PascalStr(aCode), PascalStr(aScriptURL), GetObjectUintptr(aFrame), uintptr(aStartLine))
}

func (m *TChromiumCore) ExecuteJavaScript2(aCode, aScriptURL string, aFrameIdentifier int64, aStartLine int32) {
	CEF().SysCallN(1783, m.Instance(), PascalStr(aCode), PascalStr(aScriptURL), uintptr(unsafePointer(&aFrameIdentifier)), uintptr(aStartLine))
}

func (m *TChromiumCore) UpdatePreferences() {
	CEF().SysCallN(2066, m.Instance())
}

func (m *TChromiumCore) SavePreferences(aFileName string) {
	CEF().SysCallN(1887, m.Instance(), PascalStr(aFileName))
}

func (m *TChromiumCore) ResolveHost(aURL string) {
	CEF().SysCallN(1878, m.Instance(), PascalStr(aURL))
}

func (m *TChromiumCore) SetUserAgentOverride(aUserAgent string, aAcceptLanguage string, aPlatform string) {
	CEF().SysCallN(2052, m.Instance(), PascalStr(aUserAgent), PascalStr(aAcceptLanguage), PascalStr(aPlatform))
}

func (m *TChromiumCore) ClearDataForOrigin(aOrigin string, aStorageTypes TCefClearDataStorageTypes) {
	CEF().SysCallN(1735, m.Instance(), PascalStr(aOrigin), uintptr(aStorageTypes))
}

func (m *TChromiumCore) ClearCache() {
	CEF().SysCallN(1733, m.Instance())
}

func (m *TChromiumCore) ToggleAudioMuted() {
	CEF().SysCallN(2063, m.Instance())
}

func (m *TChromiumCore) ShowDevTools(inspectElementAt *TPoint, aWindowInfo *TCefWindowInfo) {
	inArgs1 := aWindowInfo.Pointer()
	CEF().SysCallN(2055, m.Instance(), uintptr(unsafePointer(inspectElementAt)), uintptr(unsafePointer(inArgs1)))
}

func (m *TChromiumCore) CloseDevTools() {
	CEF().SysCallN(1746, m.Instance())
}

func (m *TChromiumCore) CloseDevTools1(aDevToolsWnd TCefWindowHandle) {
	CEF().SysCallN(1747, m.Instance(), uintptr(aDevToolsWnd))
}

func (m *TChromiumCore) Find(aSearchText string, aForward, aMatchCase, aFindNext bool) {
	CEF().SysCallN(1787, m.Instance(), PascalStr(aSearchText), PascalBool(aForward), PascalBool(aMatchCase), PascalBool(aFindNext))
}

func (m *TChromiumCore) StopFinding(aClearSelection bool) {
	CEF().SysCallN(2061, m.Instance(), PascalBool(aClearSelection))
}

func (m *TChromiumCore) Print() {
	CEF().SysCallN(1853, m.Instance())
}

func (m *TChromiumCore) PrintToPDF(aFilePath string) {
	CEF().SysCallN(1854, m.Instance(), PascalStr(aFilePath))
}

func (m *TChromiumCore) ClipboardCopy() {
	CEF().SysCallN(1737, m.Instance())
}

func (m *TChromiumCore) ClipboardPaste() {
	CEF().SysCallN(1740, m.Instance())
}

func (m *TChromiumCore) ClipboardCut() {
	CEF().SysCallN(1738, m.Instance())
}

func (m *TChromiumCore) ClipboardUndo() {
	CEF().SysCallN(1742, m.Instance())
}

func (m *TChromiumCore) ClipboardRedo() {
	CEF().SysCallN(1741, m.Instance())
}

func (m *TChromiumCore) ClipboardDel() {
	CEF().SysCallN(1739, m.Instance())
}

func (m *TChromiumCore) SelectAll() {
	CEF().SysCallN(1888, m.Instance())
}

func (m *TChromiumCore) IncZoomStep() {
	CEF().SysCallN(1818, m.Instance())
}

func (m *TChromiumCore) DecZoomStep() {
	CEF().SysCallN(1761, m.Instance())
}

func (m *TChromiumCore) IncZoomPct() {
	CEF().SysCallN(1817, m.Instance())
}

func (m *TChromiumCore) DecZoomPct() {
	CEF().SysCallN(1760, m.Instance())
}

func (m *TChromiumCore) ResetZoomStep() {
	CEF().SysCallN(1877, m.Instance())
}

func (m *TChromiumCore) ResetZoomLevel() {
	CEF().SysCallN(1875, m.Instance())
}

func (m *TChromiumCore) ResetZoomPct() {
	CEF().SysCallN(1876, m.Instance())
}

func (m *TChromiumCore) ReadZoom() {
	CEF().SysCallN(1865, m.Instance())
}

func (m *TChromiumCore) IncZoomCommand() {
	CEF().SysCallN(1816, m.Instance())
}

func (m *TChromiumCore) DecZoomCommand() {
	CEF().SysCallN(1759, m.Instance())
}

func (m *TChromiumCore) ResetZoomCommand() {
	CEF().SysCallN(1874, m.Instance())
}

func (m *TChromiumCore) WasResized() {
	CEF().SysCallN(2072, m.Instance())
}

func (m *TChromiumCore) WasHidden(hidden bool) {
	CEF().SysCallN(2071, m.Instance(), PascalBool(hidden))
}

func (m *TChromiumCore) NotifyScreenInfoChanged() {
	CEF().SysCallN(1848, m.Instance())
}

func (m *TChromiumCore) NotifyMoveOrResizeStarted() {
	CEF().SysCallN(1847, m.Instance())
}

func (m *TChromiumCore) Invalidate(type_ TCefPaintElementType) {
	CEF().SysCallN(1822, m.Instance(), uintptr(type_))
}

func (m *TChromiumCore) ExitFullscreen(willcauseresize bool) {
	CEF().SysCallN(1785, m.Instance(), PascalBool(willcauseresize))
}

func (m *TChromiumCore) SendExternalBeginFrame() {
	CEF().SysCallN(1892, m.Instance())
}

func (m *TChromiumCore) SendKeyEvent(event *TCefKeyEvent) {
	CEF().SysCallN(1893, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TChromiumCore) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	CEF().SysCallN(1894, m.Instance(), uintptr(unsafePointer(event)), uintptr(type_), PascalBool(mouseUp), uintptr(clickCount))
}

func (m *TChromiumCore) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	CEF().SysCallN(1895, m.Instance(), uintptr(unsafePointer(event)), PascalBool(mouseLeave))
}

func (m *TChromiumCore) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	CEF().SysCallN(1896, m.Instance(), uintptr(unsafePointer(event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TChromiumCore) SendTouchEvent(event *TCefTouchEvent) {
	CEF().SysCallN(1901, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TChromiumCore) SendCaptureLostEvent() {
	CEF().SysCallN(1890, m.Instance())
}

func (m *TChromiumCore) SendProcessMessage(targetProcess TCefProcessId, procMessage ICefProcessMessage, aFrameName string) {
	CEF().SysCallN(1897, m.Instance(), uintptr(targetProcess), GetObjectUintptr(procMessage), PascalStr(aFrameName))
}

func (m *TChromiumCore) SendProcessMessage1(targetProcess TCefProcessId, procMessage ICefProcessMessage, aFrame ICefFrame) {
	CEF().SysCallN(1898, m.Instance(), uintptr(targetProcess), GetObjectUintptr(procMessage), GetObjectUintptr(aFrame))
}

func (m *TChromiumCore) SendProcessMessage2(targetProcess TCefProcessId, procMessage ICefProcessMessage, aFrameIdentifier int64) {
	CEF().SysCallN(1899, m.Instance(), uintptr(targetProcess), GetObjectUintptr(procMessage), uintptr(unsafePointer(&aFrameIdentifier)))
}

func (m *TChromiumCore) SetFocus(focus bool) {
	CEF().SysCallN(1904, m.Instance(), PascalBool(focus))
}

func (m *TChromiumCore) SetAccessibilityState(accessibilityState TCefState) {
	CEF().SysCallN(1902, m.Instance(), uintptr(accessibilityState))
}

func (m *TChromiumCore) DragTargetDragEnter(dragData ICefDragData, event *TCefMouseEvent, allowedOps TCefDragOperations) {
	CEF().SysCallN(1775, m.Instance(), GetObjectUintptr(dragData), uintptr(unsafePointer(event)), uintptr(allowedOps))
}

func (m *TChromiumCore) DragTargetDragOver(event *TCefMouseEvent, allowedOps TCefDragOperations) {
	CEF().SysCallN(1777, m.Instance(), uintptr(unsafePointer(event)), uintptr(allowedOps))
}

func (m *TChromiumCore) DragTargetDragLeave() {
	CEF().SysCallN(1776, m.Instance())
}

func (m *TChromiumCore) DragTargetDrop(event *TCefMouseEvent) {
	CEF().SysCallN(1778, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TChromiumCore) DragSourceEndedAt(x, y int32, op TCefDragOperation) {
	CEF().SysCallN(1773, m.Instance(), uintptr(x), uintptr(y), uintptr(op))
}

func (m *TChromiumCore) DragSourceSystemDragEnded() {
	CEF().SysCallN(1774, m.Instance())
}

func (m *TChromiumCore) IMECommitText(text string, replacementrange *TCefRange, relativecursorpos int32) {
	CEF().SysCallN(1814, m.Instance(), PascalStr(text), uintptr(unsafePointer(replacementrange)), uintptr(relativecursorpos))
}

func (m *TChromiumCore) IMEFinishComposingText(keepselection bool) {
	CEF().SysCallN(1815, m.Instance(), PascalBool(keepselection))
}

func (m *TChromiumCore) IMECancelComposition() {
	CEF().SysCallN(1813, m.Instance())
}

func (m *TChromiumCore) ReplaceMisspelling(aWord string) {
	CEF().SysCallN(1869, m.Instance(), PascalStr(aWord))
}

func (m *TChromiumCore) AddWordToDictionary(aWord string) {
	CEF().SysCallN(1711, m.Instance(), PascalStr(aWord))
}

func (m *TChromiumCore) UpdateBrowserSize(aLeft, aTop, aWidth, aHeight int32) {
	CEF().SysCallN(2065, m.Instance(), uintptr(aLeft), uintptr(aTop), uintptr(aWidth), uintptr(aHeight))
}

func (m *TChromiumCore) UpdateXWindowVisibility(aVisible bool) {
	CEF().SysCallN(2067, m.Instance(), PascalBool(aVisible))
}

func (m *TChromiumCore) NotifyCurrentSinks() {
	CEF().SysCallN(1846, m.Instance())
}

func (m *TChromiumCore) NotifyCurrentRoutes() {
	CEF().SysCallN(1845, m.Instance())
}

func (m *TChromiumCore) CreateRoute(source ICefMediaSource, sink ICefMediaSink) {
	CEF().SysCallN(1753, m.Instance(), GetObjectUintptr(source), GetObjectUintptr(sink))
}

func (m *TChromiumCore) GetDeviceInfo(aMediaSink ICefMediaSink) {
	CEF().SysCallN(1794, m.Instance(), GetObjectUintptr(aMediaSink))
}

func (m *TChromiumCore) SetWebsiteSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value ICefValue) {
	CEF().SysCallN(2053, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype), GetObjectUintptr(value))
}

func (m *TChromiumCore) SetContentSetting(requestingurl, toplevelurl string, contenttype TCefContentSettingTypes, value TCefContentSettingValues) {
	CEF().SysCallN(1903, m.Instance(), PascalStr(requestingurl), PascalStr(toplevelurl), uintptr(contenttype), uintptr(value))
}

func (m *TChromiumCore) SetOnTextResultAvailable(fn TOnTextResultAvailable) {
	if m.textResultAvailablePtr != 0 {
		RemoveEventElement(m.textResultAvailablePtr)
	}
	m.textResultAvailablePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2043, m.Instance(), m.textResultAvailablePtr)
}

func (m *TChromiumCore) SetOnPdfPrintFinished(fn TOnPdfPrintFinished) {
	if m.pdfPrintFinishedPtr != 0 {
		RemoveEventElement(m.pdfPrintFinishedPtr)
	}
	m.pdfPrintFinishedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2005, m.Instance(), m.pdfPrintFinishedPtr)
}

func (m *TChromiumCore) SetOnPrefsAvailable(fn TOnPrefsAvailable) {
	if m.prefsAvailablePtr != 0 {
		RemoveEventElement(m.prefsAvailablePtr)
	}
	m.prefsAvailablePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2009, m.Instance(), m.prefsAvailablePtr)
}

func (m *TChromiumCore) SetOnPrefsUpdated(fn TNotify) {
	if m.prefsUpdatedPtr != 0 {
		RemoveEventElement(m.prefsUpdatedPtr)
	}
	m.prefsUpdatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2010, m.Instance(), m.prefsUpdatedPtr)
}

func (m *TChromiumCore) SetOnCookiesDeleted(fn TOnCookiesDeleted) {
	if m.cookiesDeletedPtr != 0 {
		RemoveEventElement(m.cookiesDeletedPtr)
	}
	m.cookiesDeletedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1936, m.Instance(), m.cookiesDeletedPtr)
}

func (m *TChromiumCore) SetOnResolvedHostAvailable(fn TOnResolvedIPsAvailable) {
	if m.resolvedHostAvailablePtr != 0 {
		RemoveEventElement(m.resolvedHostAvailablePtr)
	}
	m.resolvedHostAvailablePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2026, m.Instance(), m.resolvedHostAvailablePtr)
}

func (m *TChromiumCore) SetOnNavigationVisitorResultAvailable(fn TOnNavigationVisitorResultAvailable) {
	if m.navigationVisitorResultAvailablePtr != 0 {
		RemoveEventElement(m.navigationVisitorResultAvailablePtr)
	}
	m.navigationVisitorResultAvailablePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2002, m.Instance(), m.navigationVisitorResultAvailablePtr)
}

func (m *TChromiumCore) SetOnDownloadImageFinished(fn TOnDownloadImageFinished) {
	if m.downloadImageFinishedPtr != 0 {
		RemoveEventElement(m.downloadImageFinishedPtr)
	}
	m.downloadImageFinishedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1951, m.Instance(), m.downloadImageFinishedPtr)
}

func (m *TChromiumCore) SetOnCookiesFlushed(fn TNotify) {
	if m.cookiesFlushedPtr != 0 {
		RemoveEventElement(m.cookiesFlushedPtr)
	}
	m.cookiesFlushedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1937, m.Instance(), m.cookiesFlushedPtr)
}

func (m *TChromiumCore) SetOnCertificateExceptionsCleared(fn TNotify) {
	if m.certificateExceptionsClearedPtr != 0 {
		RemoveEventElement(m.certificateExceptionsClearedPtr)
	}
	m.certificateExceptionsClearedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1928, m.Instance(), m.certificateExceptionsClearedPtr)
}

func (m *TChromiumCore) SetOnHttpAuthCredentialsCleared(fn TNotify) {
	if m.httpAuthCredentialsClearedPtr != 0 {
		RemoveEventElement(m.httpAuthCredentialsClearedPtr)
	}
	m.httpAuthCredentialsClearedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1985, m.Instance(), m.httpAuthCredentialsClearedPtr)
}

func (m *TChromiumCore) SetOnAllConnectionsClosed(fn TNotify) {
	if m.allConnectionsClosedPtr != 0 {
		RemoveEventElement(m.allConnectionsClosedPtr)
	}
	m.allConnectionsClosedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1909, m.Instance(), m.allConnectionsClosedPtr)
}

func (m *TChromiumCore) SetOnExecuteTaskOnCefThread(fn TOnExecuteTaskOnCefThread) {
	if m.executeTaskOnCefThreadPtr != 0 {
		RemoveEventElement(m.executeTaskOnCefThreadPtr)
	}
	m.executeTaskOnCefThreadPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1955, m.Instance(), m.executeTaskOnCefThreadPtr)
}

func (m *TChromiumCore) SetOnCookiesVisited(fn TOnCookiesVisited) {
	if m.cookiesVisitedPtr != 0 {
		RemoveEventElement(m.cookiesVisitedPtr)
	}
	m.cookiesVisitedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1938, m.Instance(), m.cookiesVisitedPtr)
}

func (m *TChromiumCore) SetOnCookieVisitorDestroyed(fn TOnCookieVisitorDestroyed) {
	if m.cookieVisitorDestroyedPtr != 0 {
		RemoveEventElement(m.cookieVisitorDestroyedPtr)
	}
	m.cookieVisitorDestroyedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1935, m.Instance(), m.cookieVisitorDestroyedPtr)
}

func (m *TChromiumCore) SetOnCookieSet(fn TOnCookieSet) {
	if m.cookieSetPtr != 0 {
		RemoveEventElement(m.cookieSetPtr)
	}
	m.cookieSetPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1934, m.Instance(), m.cookieSetPtr)
}

func (m *TChromiumCore) SetOnZoomPctAvailable(fn TOnZoomPctAvailable) {
	if m.zoomPctAvailablePtr != 0 {
		RemoveEventElement(m.zoomPctAvailablePtr)
	}
	m.zoomPctAvailablePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2051, m.Instance(), m.zoomPctAvailablePtr)
}

func (m *TChromiumCore) SetOnMediaRouteCreateFinished(fn TOnMediaRouteCreateFinished) {
	if m.mediaRouteCreateFinishedPtr != 0 {
		RemoveEventElement(m.mediaRouteCreateFinishedPtr)
	}
	m.mediaRouteCreateFinishedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2000, m.Instance(), m.mediaRouteCreateFinishedPtr)
}

func (m *TChromiumCore) SetOnMediaSinkDeviceInfo(fn TOnMediaSinkDeviceInfo) {
	if m.mediaSinkDeviceInfoPtr != 0 {
		RemoveEventElement(m.mediaSinkDeviceInfoPtr)
	}
	m.mediaSinkDeviceInfoPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2001, m.Instance(), m.mediaSinkDeviceInfoPtr)
}

func (m *TChromiumCore) SetOnCanFocus(fn TNotify) {
	if m.canFocusPtr != 0 {
		RemoveEventElement(m.canFocusPtr)
	}
	m.canFocusPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1924, m.Instance(), m.canFocusPtr)
}

func (m *TChromiumCore) SetOnBrowserCompMsg(fn TOnCompMsg) {
	if m.browserCompMsgPtr != 0 {
		RemoveEventElement(m.browserCompMsgPtr)
	}
	m.browserCompMsgPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1922, m.Instance(), m.browserCompMsgPtr)
}

func (m *TChromiumCore) SetOnWidgetCompMsg(fn TOnCompMsg) {
	if m.widgetCompMsgPtr != 0 {
		RemoveEventElement(m.widgetCompMsgPtr)
	}
	m.widgetCompMsgPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2050, m.Instance(), m.widgetCompMsgPtr)
}

func (m *TChromiumCore) SetOnRenderCompMsg(fn TOnCompMsg) {
	if m.renderCompMsgPtr != 0 {
		RemoveEventElement(m.renderCompMsgPtr)
	}
	m.renderCompMsgPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2020, m.Instance(), m.renderCompMsgPtr)
}

func (m *TChromiumCore) SetOnProcessMessageReceived(fn TOnProcessMessageReceived) {
	if m.processMessageReceivedPtr != 0 {
		RemoveEventElement(m.processMessageReceivedPtr)
	}
	m.processMessageReceivedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2016, m.Instance(), m.processMessageReceivedPtr)
}

func (m *TChromiumCore) SetOnLoadStart(fn TOnLoadStart) {
	if m.loadStartPtr != 0 {
		RemoveEventElement(m.loadStartPtr)
	}
	m.loadStartPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1995, m.Instance(), m.loadStartPtr)
}

func (m *TChromiumCore) SetOnLoadEnd(fn TOnLoadEnd) {
	if m.loadEndPtr != 0 {
		RemoveEventElement(m.loadEndPtr)
	}
	m.loadEndPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1993, m.Instance(), m.loadEndPtr)
}

func (m *TChromiumCore) SetOnLoadError(fn TOnLoadError) {
	if m.loadErrorPtr != 0 {
		RemoveEventElement(m.loadErrorPtr)
	}
	m.loadErrorPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1994, m.Instance(), m.loadErrorPtr)
}

func (m *TChromiumCore) SetOnLoadingStateChange(fn TOnLoadingStateChange) {
	if m.loadingStateChangePtr != 0 {
		RemoveEventElement(m.loadingStateChangePtr)
	}
	m.loadingStateChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1997, m.Instance(), m.loadingStateChangePtr)
}

func (m *TChromiumCore) SetOnTakeFocus(fn TOnTakeFocus) {
	if m.takeFocusPtr != 0 {
		RemoveEventElement(m.takeFocusPtr)
	}
	m.takeFocusPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2042, m.Instance(), m.takeFocusPtr)
}

func (m *TChromiumCore) SetOnSetFocus(fn TOnSetFocus) {
	if m.setFocusPtr != 0 {
		RemoveEventElement(m.setFocusPtr)
	}
	m.setFocusPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2037, m.Instance(), m.setFocusPtr)
}

func (m *TChromiumCore) SetOnGotFocus(fn TOnGotFocus) {
	if m.gotFocusPtr != 0 {
		RemoveEventElement(m.gotFocusPtr)
	}
	m.gotFocusPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1984, m.Instance(), m.gotFocusPtr)
}

func (m *TChromiumCore) SetOnBeforeContextMenu(fn TOnBeforeContextMenu) {
	if m.beforeContextMenuPtr != 0 {
		RemoveEventElement(m.beforeContextMenuPtr)
	}
	m.beforeContextMenuPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1917, m.Instance(), m.beforeContextMenuPtr)
}

func (m *TChromiumCore) SetOnRunContextMenu(fn TOnRunContextMenu) {
	if m.runContextMenuPtr != 0 {
		RemoveEventElement(m.runContextMenuPtr)
	}
	m.runContextMenuPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2033, m.Instance(), m.runContextMenuPtr)
}

func (m *TChromiumCore) SetOnContextMenuCommand(fn TOnContextMenuCommand) {
	if m.contextMenuCommandPtr != 0 {
		RemoveEventElement(m.contextMenuCommandPtr)
	}
	m.contextMenuCommandPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1932, m.Instance(), m.contextMenuCommandPtr)
}

func (m *TChromiumCore) SetOnContextMenuDismissed(fn TOnContextMenuDismissed) {
	if m.contextMenuDismissedPtr != 0 {
		RemoveEventElement(m.contextMenuDismissedPtr)
	}
	m.contextMenuDismissedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1933, m.Instance(), m.contextMenuDismissedPtr)
}

func (m *TChromiumCore) SetOnRunQuickMenu(fn TOnRunQuickMenu) {
	if m.runQuickMenuPtr != 0 {
		RemoveEventElement(m.runQuickMenuPtr)
	}
	m.runQuickMenuPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2034, m.Instance(), m.runQuickMenuPtr)
}

func (m *TChromiumCore) SetOnQuickMenuCommand(fn TOnQuickMenuCommand) {
	if m.quickMenuCommandPtr != 0 {
		RemoveEventElement(m.quickMenuCommandPtr)
	}
	m.quickMenuCommandPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2018, m.Instance(), m.quickMenuCommandPtr)
}

func (m *TChromiumCore) SetOnQuickMenuDismissed(fn TOnQuickMenuDismissed) {
	if m.quickMenuDismissedPtr != 0 {
		RemoveEventElement(m.quickMenuDismissedPtr)
	}
	m.quickMenuDismissedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2019, m.Instance(), m.quickMenuDismissedPtr)
}

func (m *TChromiumCore) SetOnPreKeyEvent(fn TOnPreKey) {
	if m.preKeyEventPtr != 0 {
		RemoveEventElement(m.preKeyEventPtr)
	}
	m.preKeyEventPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2008, m.Instance(), m.preKeyEventPtr)
}

func (m *TChromiumCore) SetOnKeyEvent(fn TOnKey) {
	if m.keyEventPtr != 0 {
		RemoveEventElement(m.keyEventPtr)
	}
	m.keyEventPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1992, m.Instance(), m.keyEventPtr)
}

func (m *TChromiumCore) SetOnAddressChange(fn TOnAddressChange) {
	if m.addressChangePtr != 0 {
		RemoveEventElement(m.addressChangePtr)
	}
	m.addressChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1907, m.Instance(), m.addressChangePtr)
}

func (m *TChromiumCore) SetOnTitleChange(fn TOnTitleChange) {
	if m.titleChangePtr != 0 {
		RemoveEventElement(m.titleChangePtr)
	}
	m.titleChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2045, m.Instance(), m.titleChangePtr)
}

func (m *TChromiumCore) SetOnFavIconUrlChange(fn TOnFavIconUrlChange) {
	if m.favIconUrlChangePtr != 0 {
		RemoveEventElement(m.favIconUrlChangePtr)
	}
	m.favIconUrlChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1964, m.Instance(), m.favIconUrlChangePtr)
}

func (m *TChromiumCore) SetOnFullScreenModeChange(fn TOnFullScreenModeChange) {
	if m.fullScreenModeChangePtr != 0 {
		RemoveEventElement(m.fullScreenModeChangePtr)
	}
	m.fullScreenModeChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1970, m.Instance(), m.fullScreenModeChangePtr)
}

func (m *TChromiumCore) SetOnTooltip(fn TOnTooltip) {
	if m.tooltipPtr != 0 {
		RemoveEventElement(m.tooltipPtr)
	}
	m.tooltipPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2046, m.Instance(), m.tooltipPtr)
}

func (m *TChromiumCore) SetOnStatusMessage(fn TOnStatusMessage) {
	if m.statusMessagePtr != 0 {
		RemoveEventElement(m.statusMessagePtr)
	}
	m.statusMessagePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2041, m.Instance(), m.statusMessagePtr)
}

func (m *TChromiumCore) SetOnConsoleMessage(fn TOnConsoleMessage) {
	if m.consoleMessagePtr != 0 {
		RemoveEventElement(m.consoleMessagePtr)
	}
	m.consoleMessagePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1931, m.Instance(), m.consoleMessagePtr)
}

func (m *TChromiumCore) SetOnAutoResize(fn TOnAutoResize) {
	if m.autoResizePtr != 0 {
		RemoveEventElement(m.autoResizePtr)
	}
	m.autoResizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1914, m.Instance(), m.autoResizePtr)
}

func (m *TChromiumCore) SetOnLoadingProgressChange(fn TOnLoadingProgressChange) {
	if m.loadingProgressChangePtr != 0 {
		RemoveEventElement(m.loadingProgressChangePtr)
	}
	m.loadingProgressChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1996, m.Instance(), m.loadingProgressChangePtr)
}

func (m *TChromiumCore) SetOnCursorChange(fn TOnCursorChange) {
	if m.cursorChangePtr != 0 {
		RemoveEventElement(m.cursorChangePtr)
	}
	m.cursorChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1939, m.Instance(), m.cursorChangePtr)
}

func (m *TChromiumCore) SetOnMediaAccessChange(fn TOnMediaAccessChange) {
	if m.mediaAccessChangePtr != 0 {
		RemoveEventElement(m.mediaAccessChangePtr)
	}
	m.mediaAccessChangePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1999, m.Instance(), m.mediaAccessChangePtr)
}

func (m *TChromiumCore) SetOnCanDownload(fn TOnCanDownload) {
	if m.canDownloadPtr != 0 {
		RemoveEventElement(m.canDownloadPtr)
	}
	m.canDownloadPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1923, m.Instance(), m.canDownloadPtr)
}

func (m *TChromiumCore) SetOnBeforeDownload(fn TOnBeforeDownload) {
	if m.beforeDownloadPtr != 0 {
		RemoveEventElement(m.beforeDownloadPtr)
	}
	m.beforeDownloadPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1918, m.Instance(), m.beforeDownloadPtr)
}

func (m *TChromiumCore) SetOnDownloadUpdated(fn TOnDownloadUpdated) {
	if m.downloadUpdatedPtr != 0 {
		RemoveEventElement(m.downloadUpdatedPtr)
	}
	m.downloadUpdatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1952, m.Instance(), m.downloadUpdatedPtr)
}

func (m *TChromiumCore) SetOnJsdialog(fn TOnJsdialog) {
	if m.jsdialogPtr != 0 {
		RemoveEventElement(m.jsdialogPtr)
	}
	m.jsdialogPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1991, m.Instance(), m.jsdialogPtr)
}

func (m *TChromiumCore) SetOnBeforeUnloadDialog(fn TOnBeforeUnloadDialog) {
	if m.beforeUnloadDialogPtr != 0 {
		RemoveEventElement(m.beforeUnloadDialogPtr)
	}
	m.beforeUnloadDialogPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1921, m.Instance(), m.beforeUnloadDialogPtr)
}

func (m *TChromiumCore) SetOnResetDialogState(fn TOnResetDialogState) {
	if m.resetDialogStatePtr != 0 {
		RemoveEventElement(m.resetDialogStatePtr)
	}
	m.resetDialogStatePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2025, m.Instance(), m.resetDialogStatePtr)
}

func (m *TChromiumCore) SetOnDialogClosed(fn TOnDialogClosed) {
	if m.dialogClosedPtr != 0 {
		RemoveEventElement(m.dialogClosedPtr)
	}
	m.dialogClosedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1948, m.Instance(), m.dialogClosedPtr)
}

func (m *TChromiumCore) SetOnBeforePopup(fn TOnBeforePopup) {
	if m.beforePopupPtr != 0 {
		RemoveEventElement(m.beforePopupPtr)
	}
	m.beforePopupPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1919, m.Instance(), m.beforePopupPtr)
}

func (m *TChromiumCore) SetOnAfterCreated(fn TOnAfterCreated) {
	if m.afterCreatedPtr != 0 {
		RemoveEventElement(m.afterCreatedPtr)
	}
	m.afterCreatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1908, m.Instance(), m.afterCreatedPtr)
}

func (m *TChromiumCore) SetOnBeforeClose(fn TOnBeforeClose) {
	if m.beforeClosePtr != 0 {
		RemoveEventElement(m.beforeClosePtr)
	}
	m.beforeClosePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1916, m.Instance(), m.beforeClosePtr)
}

func (m *TChromiumCore) SetOnClose(fn TOnClose) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1930, m.Instance(), m.closePtr)
}

func (m *TChromiumCore) SetOnBeforeBrowse(fn TOnBeforeBrowse) {
	if m.beforeBrowsePtr != 0 {
		RemoveEventElement(m.beforeBrowsePtr)
	}
	m.beforeBrowsePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1915, m.Instance(), m.beforeBrowsePtr)
}

func (m *TChromiumCore) SetOnOpenUrlFromTab(fn TOnOpenUrlFromTab) {
	if m.openUrlFromTabPtr != 0 {
		RemoveEventElement(m.openUrlFromTabPtr)
	}
	m.openUrlFromTabPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2003, m.Instance(), m.openUrlFromTabPtr)
}

func (m *TChromiumCore) SetOnGetAuthCredentials(fn TOnGetAuthCredentials) {
	if m.getAuthCredentialsPtr != 0 {
		RemoveEventElement(m.getAuthCredentialsPtr)
	}
	m.getAuthCredentialsPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1973, m.Instance(), m.getAuthCredentialsPtr)
}

func (m *TChromiumCore) SetOnCertificateError(fn TOnCertificateError) {
	if m.certificateErrorPtr != 0 {
		RemoveEventElement(m.certificateErrorPtr)
	}
	m.certificateErrorPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1927, m.Instance(), m.certificateErrorPtr)
}

func (m *TChromiumCore) SetOnSelectClientCertificate(fn TOnSelectClientCertificate) {
	if m.selectClientCertificatePtr != 0 {
		RemoveEventElement(m.selectClientCertificatePtr)
	}
	m.selectClientCertificatePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2036, m.Instance(), m.selectClientCertificatePtr)
}

func (m *TChromiumCore) SetOnRenderViewReady(fn TOnRenderViewReady) {
	if m.renderViewReadyPtr != 0 {
		RemoveEventElement(m.renderViewReadyPtr)
	}
	m.renderViewReadyPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2022, m.Instance(), m.renderViewReadyPtr)
}

func (m *TChromiumCore) SetOnRenderProcessTerminated(fn TOnRenderProcessTerminated) {
	if m.renderProcessTerminatedPtr != 0 {
		RemoveEventElement(m.renderProcessTerminatedPtr)
	}
	m.renderProcessTerminatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2021, m.Instance(), m.renderProcessTerminatedPtr)
}

func (m *TChromiumCore) SetOnGetResourceRequestHandler_ReqHdlr(fn TOnGetResourceRequestHandler) {
	if m.getResourceRequestHandler_ReqHdlrPtr != 0 {
		RemoveEventElement(m.getResourceRequestHandler_ReqHdlrPtr)
	}
	m.getResourceRequestHandler_ReqHdlrPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1977, m.Instance(), m.getResourceRequestHandler_ReqHdlrPtr)
}

func (m *TChromiumCore) SetOnDocumentAvailableInMainFrame(fn TOnDocumentAvailableInMainFrame) {
	if m.documentAvailableInMainFramePtr != 0 {
		RemoveEventElement(m.documentAvailableInMainFramePtr)
	}
	m.documentAvailableInMainFramePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1950, m.Instance(), m.documentAvailableInMainFramePtr)
}

func (m *TChromiumCore) SetOnBeforeResourceLoad(fn TOnBeforeResourceLoad) {
	if m.beforeResourceLoadPtr != 0 {
		RemoveEventElement(m.beforeResourceLoadPtr)
	}
	m.beforeResourceLoadPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1920, m.Instance(), m.beforeResourceLoadPtr)
}

func (m *TChromiumCore) SetOnGetResourceHandler(fn TOnGetResourceHandler) {
	if m.getResourceHandlerPtr != 0 {
		RemoveEventElement(m.getResourceHandlerPtr)
	}
	m.getResourceHandlerPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1975, m.Instance(), m.getResourceHandlerPtr)
}

func (m *TChromiumCore) SetOnResourceRedirect(fn TOnResourceRedirect) {
	if m.resourceRedirectPtr != 0 {
		RemoveEventElement(m.resourceRedirectPtr)
	}
	m.resourceRedirectPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2028, m.Instance(), m.resourceRedirectPtr)
}

func (m *TChromiumCore) SetOnResourceResponse(fn TOnResourceResponse) {
	if m.resourceResponsePtr != 0 {
		RemoveEventElement(m.resourceResponsePtr)
	}
	m.resourceResponsePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2029, m.Instance(), m.resourceResponsePtr)
}

func (m *TChromiumCore) SetOnGetResourceResponseFilter(fn TOnGetResourceResponseFilter) {
	if m.getResourceResponseFilterPtr != 0 {
		RemoveEventElement(m.getResourceResponseFilterPtr)
	}
	m.getResourceResponseFilterPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1978, m.Instance(), m.getResourceResponseFilterPtr)
}

func (m *TChromiumCore) SetOnResourceLoadComplete(fn TOnResourceLoadComplete) {
	if m.resourceLoadCompletePtr != 0 {
		RemoveEventElement(m.resourceLoadCompletePtr)
	}
	m.resourceLoadCompletePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2027, m.Instance(), m.resourceLoadCompletePtr)
}

func (m *TChromiumCore) SetOnProtocolExecution(fn TOnProtocolExecution) {
	if m.protocolExecutionPtr != 0 {
		RemoveEventElement(m.protocolExecutionPtr)
	}
	m.protocolExecutionPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2017, m.Instance(), m.protocolExecutionPtr)
}

func (m *TChromiumCore) SetOnCanSendCookie(fn TOnCanSendCookie) {
	if m.canSendCookiePtr != 0 {
		RemoveEventElement(m.canSendCookiePtr)
	}
	m.canSendCookiePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1926, m.Instance(), m.canSendCookiePtr)
}

func (m *TChromiumCore) SetOnCanSaveCookie(fn TOnCanSaveCookie) {
	if m.canSaveCookiePtr != 0 {
		RemoveEventElement(m.canSaveCookiePtr)
	}
	m.canSaveCookiePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1925, m.Instance(), m.canSaveCookiePtr)
}

func (m *TChromiumCore) SetOnFileDialog(fn TOnFileDialog) {
	if m.fileDialogPtr != 0 {
		RemoveEventElement(m.fileDialogPtr)
	}
	m.fileDialogPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1965, m.Instance(), m.fileDialogPtr)
}

func (m *TChromiumCore) SetOnGetAccessibilityHandler(fn TOnGetAccessibilityHandler) {
	if m.getAccessibilityHandlerPtr != 0 {
		RemoveEventElement(m.getAccessibilityHandlerPtr)
	}
	m.getAccessibilityHandlerPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1971, m.Instance(), m.getAccessibilityHandlerPtr)
}

func (m *TChromiumCore) SetOnGetRootScreenRect(fn TOnGetRootScreenRect) {
	if m.getRootScreenRectPtr != 0 {
		RemoveEventElement(m.getRootScreenRectPtr)
	}
	m.getRootScreenRectPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1979, m.Instance(), m.getRootScreenRectPtr)
}

func (m *TChromiumCore) SetOnGetViewRect(fn TOnGetViewRect) {
	if m.getViewRectPtr != 0 {
		RemoveEventElement(m.getViewRectPtr)
	}
	m.getViewRectPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1983, m.Instance(), m.getViewRectPtr)
}

func (m *TChromiumCore) SetOnGetScreenPoint(fn TOnGetScreenPoint) {
	if m.getScreenPointPtr != 0 {
		RemoveEventElement(m.getScreenPointPtr)
	}
	m.getScreenPointPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1981, m.Instance(), m.getScreenPointPtr)
}

func (m *TChromiumCore) SetOnGetScreenInfo(fn TOnGetScreenInfo) {
	if m.getScreenInfoPtr != 0 {
		RemoveEventElement(m.getScreenInfoPtr)
	}
	m.getScreenInfoPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1980, m.Instance(), m.getScreenInfoPtr)
}

func (m *TChromiumCore) SetOnPopupShow(fn TOnPopupShow) {
	if m.popupShowPtr != 0 {
		RemoveEventElement(m.popupShowPtr)
	}
	m.popupShowPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2006, m.Instance(), m.popupShowPtr)
}

func (m *TChromiumCore) SetOnPopupSize(fn TOnPopupSize) {
	if m.popupSizePtr != 0 {
		RemoveEventElement(m.popupSizePtr)
	}
	m.popupSizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2007, m.Instance(), m.popupSizePtr)
}

func (m *TChromiumCore) SetOnPaint(fn TOnPaint) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2004, m.Instance(), m.paintPtr)
}

func (m *TChromiumCore) SetOnAcceleratedPaint(fn TOnAcceleratedPaint) {
	if m.acceleratedPaintPtr != 0 {
		RemoveEventElement(m.acceleratedPaintPtr)
	}
	m.acceleratedPaintPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1906, m.Instance(), m.acceleratedPaintPtr)
}

func (m *TChromiumCore) SetOnGetTouchHandleSize(fn TOnGetTouchHandleSize) {
	if m.getTouchHandleSizePtr != 0 {
		RemoveEventElement(m.getTouchHandleSizePtr)
	}
	m.getTouchHandleSizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1982, m.Instance(), m.getTouchHandleSizePtr)
}

func (m *TChromiumCore) SetOnTouchHandleStateChanged(fn TOnTouchHandleStateChanged) {
	if m.touchHandleStateChangedPtr != 0 {
		RemoveEventElement(m.touchHandleStateChangedPtr)
	}
	m.touchHandleStateChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2047, m.Instance(), m.touchHandleStateChangedPtr)
}

func (m *TChromiumCore) SetOnStartDragging(fn TOnStartDragging) {
	if m.startDraggingPtr != 0 {
		RemoveEventElement(m.startDraggingPtr)
	}
	m.startDraggingPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2040, m.Instance(), m.startDraggingPtr)
}

func (m *TChromiumCore) SetOnUpdateDragCursor(fn TOnUpdateDragCursor) {
	if m.updateDragCursorPtr != 0 {
		RemoveEventElement(m.updateDragCursorPtr)
	}
	m.updateDragCursorPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2048, m.Instance(), m.updateDragCursorPtr)
}

func (m *TChromiumCore) SetOnScrollOffsetChanged(fn TOnScrollOffsetChanged) {
	if m.scrollOffsetChangedPtr != 0 {
		RemoveEventElement(m.scrollOffsetChangedPtr)
	}
	m.scrollOffsetChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2035, m.Instance(), m.scrollOffsetChangedPtr)
}

func (m *TChromiumCore) SetOnIMECompositionRangeChanged(fn TOnIMECompositionRangeChanged) {
	if m.iMECompositionRangeChangedPtr != 0 {
		RemoveEventElement(m.iMECompositionRangeChangedPtr)
	}
	m.iMECompositionRangeChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1986, m.Instance(), m.iMECompositionRangeChangedPtr)
}

func (m *TChromiumCore) SetOnTextSelectionChanged(fn TOnTextSelectionChanged) {
	if m.textSelectionChangedPtr != 0 {
		RemoveEventElement(m.textSelectionChangedPtr)
	}
	m.textSelectionChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2044, m.Instance(), m.textSelectionChangedPtr)
}

func (m *TChromiumCore) SetOnVirtualKeyboardRequested(fn TOnVirtualKeyboardRequested) {
	if m.virtualKeyboardRequestedPtr != 0 {
		RemoveEventElement(m.virtualKeyboardRequestedPtr)
	}
	m.virtualKeyboardRequestedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2049, m.Instance(), m.virtualKeyboardRequestedPtr)
}

func (m *TChromiumCore) SetOnDragEnter(fn TOnDragEnter) {
	if m.dragEnterPtr != 0 {
		RemoveEventElement(m.dragEnterPtr)
	}
	m.dragEnterPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1953, m.Instance(), m.dragEnterPtr)
}

func (m *TChromiumCore) SetOnDraggableRegionsChanged(fn TOnDraggableRegionsChanged) {
	if m.draggableRegionsChangedPtr != 0 {
		RemoveEventElement(m.draggableRegionsChangedPtr)
	}
	m.draggableRegionsChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1954, m.Instance(), m.draggableRegionsChangedPtr)
}

func (m *TChromiumCore) SetOnFindResult(fn TOnFindResult) {
	if m.findResultPtr != 0 {
		RemoveEventElement(m.findResultPtr)
	}
	m.findResultPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1966, m.Instance(), m.findResultPtr)
}

func (m *TChromiumCore) SetOnRequestContextInitialized(fn TOnRequestContextInitialized) {
	if m.requestContextInitializedPtr != 0 {
		RemoveEventElement(m.requestContextInitializedPtr)
	}
	m.requestContextInitializedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2023, m.Instance(), m.requestContextInitializedPtr)
}

func (m *TChromiumCore) SetOnGetResourceRequestHandler_ReqCtxHdlr(fn TOnGetResourceRequestHandler) {
	if m.getResourceRequestHandler_ReqCtxHdlrPtr != 0 {
		RemoveEventElement(m.getResourceRequestHandler_ReqCtxHdlrPtr)
	}
	m.getResourceRequestHandler_ReqCtxHdlrPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1976, m.Instance(), m.getResourceRequestHandler_ReqCtxHdlrPtr)
}

func (m *TChromiumCore) SetOnSinks(fn TOnSinks) {
	if m.sinksPtr != 0 {
		RemoveEventElement(m.sinksPtr)
	}
	m.sinksPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2039, m.Instance(), m.sinksPtr)
}

func (m *TChromiumCore) SetOnRoutes(fn TOnRoutes) {
	if m.routesPtr != 0 {
		RemoveEventElement(m.routesPtr)
	}
	m.routesPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2032, m.Instance(), m.routesPtr)
}

func (m *TChromiumCore) SetOnRouteStateChanged(fn TOnRouteStateChanged) {
	if m.routeStateChangedPtr != 0 {
		RemoveEventElement(m.routeStateChangedPtr)
	}
	m.routeStateChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2031, m.Instance(), m.routeStateChangedPtr)
}

func (m *TChromiumCore) SetOnRouteMessageReceived(fn TOnRouteMessageReceived) {
	if m.routeMessageReceivedPtr != 0 {
		RemoveEventElement(m.routeMessageReceivedPtr)
	}
	m.routeMessageReceivedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2030, m.Instance(), m.routeMessageReceivedPtr)
}

func (m *TChromiumCore) SetOnGetAudioParameters(fn TOnGetAudioParameters) {
	if m.getAudioParametersPtr != 0 {
		RemoveEventElement(m.getAudioParametersPtr)
	}
	m.getAudioParametersPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1972, m.Instance(), m.getAudioParametersPtr)
}

func (m *TChromiumCore) SetOnAudioStreamStarted(fn TOnAudioStreamStarted) {
	if m.audioStreamStartedPtr != 0 {
		RemoveEventElement(m.audioStreamStartedPtr)
	}
	m.audioStreamStartedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1912, m.Instance(), m.audioStreamStartedPtr)
}

func (m *TChromiumCore) SetOnAudioStreamPacket(fn TOnAudioStreamPacket) {
	if m.audioStreamPacketPtr != 0 {
		RemoveEventElement(m.audioStreamPacketPtr)
	}
	m.audioStreamPacketPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1911, m.Instance(), m.audioStreamPacketPtr)
}

func (m *TChromiumCore) SetOnAudioStreamStopped(fn TOnAudioStreamStopped) {
	if m.audioStreamStoppedPtr != 0 {
		RemoveEventElement(m.audioStreamStoppedPtr)
	}
	m.audioStreamStoppedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1913, m.Instance(), m.audioStreamStoppedPtr)
}

func (m *TChromiumCore) SetOnAudioStreamError(fn TOnAudioStreamError) {
	if m.audioStreamErrorPtr != 0 {
		RemoveEventElement(m.audioStreamErrorPtr)
	}
	m.audioStreamErrorPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1910, m.Instance(), m.audioStreamErrorPtr)
}

func (m *TChromiumCore) SetOnDevToolsMessage(fn TOnDevToolsMessage) {
	if m.devToolsMessagePtr != 0 {
		RemoveEventElement(m.devToolsMessagePtr)
	}
	m.devToolsMessagePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1943, m.Instance(), m.devToolsMessagePtr)
}

func (m *TChromiumCore) SetOnDevToolsRawMessage(fn TOnDevToolsRawMessage) {
	if m.devToolsRawMessagePtr != 0 {
		RemoveEventElement(m.devToolsRawMessagePtr)
	}
	m.devToolsRawMessagePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1947, m.Instance(), m.devToolsRawMessagePtr)
}

func (m *TChromiumCore) SetOnDevToolsMethodResult(fn TOnDevToolsMethodResult) {
	if m.devToolsMethodResultPtr != 0 {
		RemoveEventElement(m.devToolsMethodResultPtr)
	}
	m.devToolsMethodResultPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1945, m.Instance(), m.devToolsMethodResultPtr)
}

func (m *TChromiumCore) SetOnDevToolsMethodRawResult(fn TOnDevToolsMethodRawResult) {
	if m.devToolsMethodRawResultPtr != 0 {
		RemoveEventElement(m.devToolsMethodRawResultPtr)
	}
	m.devToolsMethodRawResultPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1944, m.Instance(), m.devToolsMethodRawResultPtr)
}

func (m *TChromiumCore) SetOnDevToolsEvent(fn TOnDevToolsEvent) {
	if m.devToolsEventPtr != 0 {
		RemoveEventElement(m.devToolsEventPtr)
	}
	m.devToolsEventPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1942, m.Instance(), m.devToolsEventPtr)
}

func (m *TChromiumCore) SetOnDevToolsRawEvent(fn TOnDevToolsRawEvent) {
	if m.devToolsRawEventPtr != 0 {
		RemoveEventElement(m.devToolsRawEventPtr)
	}
	m.devToolsRawEventPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1946, m.Instance(), m.devToolsRawEventPtr)
}

func (m *TChromiumCore) SetOnDevToolsAgentAttached(fn TOnDevToolsAgentAttached) {
	if m.devToolsAgentAttachedPtr != 0 {
		RemoveEventElement(m.devToolsAgentAttachedPtr)
	}
	m.devToolsAgentAttachedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1940, m.Instance(), m.devToolsAgentAttachedPtr)
}

func (m *TChromiumCore) SetOnDevToolsAgentDetached(fn TOnDevToolsAgentDetached) {
	if m.devToolsAgentDetachedPtr != 0 {
		RemoveEventElement(m.devToolsAgentDetachedPtr)
	}
	m.devToolsAgentDetachedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1941, m.Instance(), m.devToolsAgentDetachedPtr)
}

func (m *TChromiumCore) SetOnExtensionLoadFailed(fn TOnExtensionLoadFailed) {
	if m.extensionLoadFailedPtr != 0 {
		RemoveEventElement(m.extensionLoadFailedPtr)
	}
	m.extensionLoadFailedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1961, m.Instance(), m.extensionLoadFailedPtr)
}

func (m *TChromiumCore) SetOnExtensionLoaded(fn TOnExtensionLoaded) {
	if m.extensionLoadedPtr != 0 {
		RemoveEventElement(m.extensionLoadedPtr)
	}
	m.extensionLoadedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1962, m.Instance(), m.extensionLoadedPtr)
}

func (m *TChromiumCore) SetOnExtensionUnloaded(fn TOnExtensionUnloaded) {
	if m.extensionUnloadedPtr != 0 {
		RemoveEventElement(m.extensionUnloadedPtr)
	}
	m.extensionUnloadedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1963, m.Instance(), m.extensionUnloadedPtr)
}

func (m *TChromiumCore) SetOnExtensionBeforeBackgroundBrowser(fn TOnBeforeBackgroundBrowser) {
	if m.extensionBeforeBackgroundBrowserPtr != 0 {
		RemoveEventElement(m.extensionBeforeBackgroundBrowserPtr)
	}
	m.extensionBeforeBackgroundBrowserPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1956, m.Instance(), m.extensionBeforeBackgroundBrowserPtr)
}

func (m *TChromiumCore) SetOnExtensionBeforeBrowser(fn TOnBeforeBrowser) {
	if m.extensionBeforeBrowserPtr != 0 {
		RemoveEventElement(m.extensionBeforeBrowserPtr)
	}
	m.extensionBeforeBrowserPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1957, m.Instance(), m.extensionBeforeBrowserPtr)
}

func (m *TChromiumCore) SetOnExtensionGetActiveBrowser(fn TOnGetActiveBrowser) {
	if m.extensionGetActiveBrowserPtr != 0 {
		RemoveEventElement(m.extensionGetActiveBrowserPtr)
	}
	m.extensionGetActiveBrowserPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1959, m.Instance(), m.extensionGetActiveBrowserPtr)
}

func (m *TChromiumCore) SetOnExtensionCanAccessBrowser(fn TOnCanAccessBrowser) {
	if m.extensionCanAccessBrowserPtr != 0 {
		RemoveEventElement(m.extensionCanAccessBrowserPtr)
	}
	m.extensionCanAccessBrowserPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1958, m.Instance(), m.extensionCanAccessBrowserPtr)
}

func (m *TChromiumCore) SetOnExtensionGetExtensionResource(fn TOnGetExtensionResource) {
	if m.extensionGetExtensionResourcePtr != 0 {
		RemoveEventElement(m.extensionGetExtensionResourcePtr)
	}
	m.extensionGetExtensionResourcePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1960, m.Instance(), m.extensionGetExtensionResourcePtr)
}

func (m *TChromiumCore) SetOnPrintStart(fn TOnPrintStart) {
	if m.printStartPtr != 0 {
		RemoveEventElement(m.printStartPtr)
	}
	m.printStartPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2015, m.Instance(), m.printStartPtr)
}

func (m *TChromiumCore) SetOnPrintSettings(fn TOnPrintSettings) {
	if m.printSettingsPtr != 0 {
		RemoveEventElement(m.printSettingsPtr)
	}
	m.printSettingsPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2014, m.Instance(), m.printSettingsPtr)
}

func (m *TChromiumCore) SetOnPrintDialog(fn TOnPrintDialog) {
	if m.printDialogPtr != 0 {
		RemoveEventElement(m.printDialogPtr)
	}
	m.printDialogPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2011, m.Instance(), m.printDialogPtr)
}

func (m *TChromiumCore) SetOnPrintJob(fn TOnPrintJob) {
	if m.printJobPtr != 0 {
		RemoveEventElement(m.printJobPtr)
	}
	m.printJobPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2012, m.Instance(), m.printJobPtr)
}

func (m *TChromiumCore) SetOnPrintReset(fn TOnPrintReset) {
	if m.printResetPtr != 0 {
		RemoveEventElement(m.printResetPtr)
	}
	m.printResetPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2013, m.Instance(), m.printResetPtr)
}

func (m *TChromiumCore) SetOnGetPDFPaperSize(fn TOnGetPDFPaperSize) {
	if m.getPDFPaperSizePtr != 0 {
		RemoveEventElement(m.getPDFPaperSizePtr)
	}
	m.getPDFPaperSizePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1974, m.Instance(), m.getPDFPaperSizePtr)
}

func (m *TChromiumCore) SetOnFrameCreated(fn TOnFrameCreated) {
	if m.frameCreatedPtr != 0 {
		RemoveEventElement(m.frameCreatedPtr)
	}
	m.frameCreatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1968, m.Instance(), m.frameCreatedPtr)
}

func (m *TChromiumCore) SetOnFrameAttached(fn TOnFrameAttached) {
	if m.frameAttachedPtr != 0 {
		RemoveEventElement(m.frameAttachedPtr)
	}
	m.frameAttachedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1967, m.Instance(), m.frameAttachedPtr)
}

func (m *TChromiumCore) SetOnFrameDetached(fn TOnFrameDetached) {
	if m.frameDetachedPtr != 0 {
		RemoveEventElement(m.frameDetachedPtr)
	}
	m.frameDetachedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1969, m.Instance(), m.frameDetachedPtr)
}

func (m *TChromiumCore) SetOnMainFrameChanged(fn TOnMainFrameChanged) {
	if m.mainFrameChangedPtr != 0 {
		RemoveEventElement(m.mainFrameChangedPtr)
	}
	m.mainFrameChangedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1998, m.Instance(), m.mainFrameChangedPtr)
}

func (m *TChromiumCore) SetOnChromeCommand(fn TOnChromeCommand) {
	if m.chromeCommandPtr != 0 {
		RemoveEventElement(m.chromeCommandPtr)
	}
	m.chromeCommandPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1929, m.Instance(), m.chromeCommandPtr)
}

func (m *TChromiumCore) SetOnIsChromeAppMenuItemVisible(fn TOnIsChromeAppMenuItemVisible) {
	if m.isChromeAppMenuItemVisiblePtr != 0 {
		RemoveEventElement(m.isChromeAppMenuItemVisiblePtr)
	}
	m.isChromeAppMenuItemVisiblePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1988, m.Instance(), m.isChromeAppMenuItemVisiblePtr)
}

func (m *TChromiumCore) SetOnIsChromeAppMenuItemEnabled(fn TOnIsChromeAppMenuItemEnabled) {
	if m.isChromeAppMenuItemEnabledPtr != 0 {
		RemoveEventElement(m.isChromeAppMenuItemEnabledPtr)
	}
	m.isChromeAppMenuItemEnabledPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1987, m.Instance(), m.isChromeAppMenuItemEnabledPtr)
}

func (m *TChromiumCore) SetOnIsChromePageActionIconVisible(fn TOnIsChromePageActionIconVisible) {
	if m.isChromePageActionIconVisiblePtr != 0 {
		RemoveEventElement(m.isChromePageActionIconVisiblePtr)
	}
	m.isChromePageActionIconVisiblePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1989, m.Instance(), m.isChromePageActionIconVisiblePtr)
}

func (m *TChromiumCore) SetOnIsChromeToolbarButtonVisible(fn TOnIsChromeToolbarButtonVisible) {
	if m.isChromeToolbarButtonVisiblePtr != 0 {
		RemoveEventElement(m.isChromeToolbarButtonVisiblePtr)
	}
	m.isChromeToolbarButtonVisiblePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1990, m.Instance(), m.isChromeToolbarButtonVisiblePtr)
}

func (m *TChromiumCore) SetOnRequestMediaAccessPermission(fn TOnRequestMediaAccessPermission) {
	if m.requestMediaAccessPermissionPtr != 0 {
		RemoveEventElement(m.requestMediaAccessPermissionPtr)
	}
	m.requestMediaAccessPermissionPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2024, m.Instance(), m.requestMediaAccessPermissionPtr)
}

func (m *TChromiumCore) SetOnShowPermissionPrompt(fn TOnShowPermissionPrompt) {
	if m.showPermissionPromptPtr != 0 {
		RemoveEventElement(m.showPermissionPromptPtr)
	}
	m.showPermissionPromptPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2038, m.Instance(), m.showPermissionPromptPtr)
}

func (m *TChromiumCore) SetOnDismissPermissionPrompt(fn TOnDismissPermissionPrompt) {
	if m.dismissPermissionPromptPtr != 0 {
		RemoveEventElement(m.dismissPermissionPromptPtr)
	}
	m.dismissPermissionPromptPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(1949, m.Instance(), m.dismissPermissionPromptPtr)
}
