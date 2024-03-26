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

// ICoreWebView2 Parent: IObject
type ICoreWebView2 interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2 // property
	// Settings
	//  The `ICoreWebView2Settings` object contains various modifiable settings
	//  for the running WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_settings">See the ICoreWebView2 article.</a>
	Settings() ICoreWebView2Settings // property
	// BrowserProcessID
	//  The process ID of the browser process that hosts the WebView.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#get_browserprocessid">See the ICoreWebView2 article.</a>
	BrowserProcessID() DWORD // property
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
	// Environment
	//  Exposes the CoreWebView2Environment used to create this CoreWebView2.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_2#get_environment">See the ICoreWebView2_2 article.</a>
	Environment() ICoreWebView2Environment // property
	// IsSuspended
	//  Whether WebView is suspended.
	//  `TRUE` when WebView is suspended, from the time when TrySuspend has completed
	//  successfully until WebView is resumed.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_3#get_issuspended">See the ICoreWebView2_3 article.</a>
	IsSuspended() bool // property
	// IsMuted
	//  Indicates whether all audio output from this CoreWebView2 is muted or not.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_ismuted">See the ICoreWebView2_8 article.</a>
	IsMuted() bool // property
	// SetIsMuted Set IsMuted
	SetIsMuted(AValue bool) // property
	// IsDocumentPlayingAudio
	//  Indicates whether any audio output from this CoreWebView2 is playing.
	//  This property will be true if audio is playing even if IsMuted is true.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_8#get_isdocumentplayingaudio">See the ICoreWebView2_8 article.</a>
	IsDocumentPlayingAudio() bool // property
	// IsDefaultDownloadDialogOpen
	//  `TRUE` if the default download dialog is currently open. The value of this
	//  property changes only when the default download dialog is explicitly
	//  opened or closed. Hiding the WebView implicitly hides the dialog, but does
	//  not change the value of this property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_9#get_isdefaultdownloaddialogopen">See the ICoreWebView2_9 article.</a>
	IsDefaultDownloadDialogOpen() bool // property
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
	// StatusBarText
	//  The status message text.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_12#get_statusbartext">See the ICoreWebView2_12 article.</a>
	StatusBarText() string // property
	// Profile
	//  The associated `ICoreWebView2Profile` object. If this CoreWebView2 was created with a
	//  CoreWebView2ControllerOptions, the CoreWebView2Profile will match those specified options.
	//  Otherwise if this CoreWebView2 was created without a CoreWebView2ControllerOptions, then
	//  this will be the default CoreWebView2Profile for the corresponding CoreWebView2Environment.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_13#get_profile">See the ICoreWebView2_13 article.</a>
	Profile() ICoreWebView2Profile // property
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
	// FrameId
	//  The unique identifier of the main frame. It's the same kind of ID as
	//  with the `FrameId` in `ICoreWebView2Frame` and via `ICoreWebView2FrameInfo`.
	//  Note that `FrameId` may not be valid if `ICoreWebView2` has not done
	//  any navigation. It's safe to get this value during or after the first
	//  `ContentLoading` event. Otherwise, it could return the invalid frame Id 0.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2_20#get_frameid">See the ICoreWebView2_20 article.</a>
	FrameId() uint32 // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// SubscribeToDevToolsProtocolEvent
	//  Subscribe to a DevTools protocol event. The TWVBrowserBase.OnDevToolsProtocolEventReceived
	//  event will be triggered on each DevTools event.
	//  <param name="aEventName">The DevTools protocol event name.</param>
	//  <param name="aEventID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	SubscribeToDevToolsProtocolEvent(aEventName string, aEventID int32, aBrowserComponent IComponent) bool // function
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
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	CapturePreview(aImageFormat TWVCapturePreviewImageFormat, aImageStream IStream, aBrowserComponent IComponent) bool // function
	// ExecuteScript
	//  Run JavaScript code from the JavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The TWVBrowserBase.OnExecuteScriptCompleted event is triggered
	//  when it finishes.
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
	//  <param name="JavaScript">The JavaScript code.</param>
	//  <param name="aExecutionID">A custom event ID that will be passed as a parameter in the TWVBrowserBase event.</param>
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	ExecuteScript(JavaScript string, aExecutionID int32, aBrowserComponent IComponent) bool // function
	// GoBack
	//  Navigates the WebView to the previous page in the navigation history.
	GoBack() bool // function
	// GoForward
	//  Navigates the WebView to the next page in the navigation history.
	GoForward() bool // function
	// Navigate
	//  Cause a navigation of the top-level document to run to the specified URI.
	Navigate(aURI string) bool // function
	// NavigateToString
	//  Initiates a navigation to aHTMLContent as source HTML of a new document.
	//  The `aHTMLContent` parameter may not be larger than 2 MB(2 * 1024 * 1024 bytes) in total size.
	//  The origin of the new page is `about:blank`.
	NavigateToString(aHTMLContent string) bool // function
	// NavigateWithWebResourceRequest
	//  Navigates using a constructed ICoreWebView2WebResourceRequest object. This lets you
	//  provide post data or additional request headers during navigation.
	//  The headers in aRequest override headers added by WebView2 runtime except for Cookie headers.
	//  Method can only be either "GET" or "POST". Provided post data will only
	//  be sent only if the method is "POST" and the uri scheme is HTTP(S).
	NavigateWithWebResourceRequest(aRequest ICoreWebView2WebResourceRequestRef) bool // function
	// Reload
	//  Reload the current page. This is similar to navigating to the URI of
	//  current top level document including all navigation events firing and
	//  respecting any entries in the HTTP cache. But, the back or forward
	//  history are not modified.
	Reload() bool // function
	// Stop
	//  Stop all navigations and pending resource fetches. Does not stop scripts.
	Stop() bool // function
	// TrySuspend
	//  An app may call the `TrySuspend` API to have the WebView2 consume less memory.
	//  This is useful when a Win32 app becomes invisible, or when a Universal Windows
	//  Platform app is being suspended, during the suspended event handler before completing
	//  the suspended event.
	//  The CoreWebView2Controller's IsVisible property must be false when the API is called.
	//  Otherwise, the API fails with `HRESULT_FROM_WIN32(ERROR_INVALID_STATE)`.
	//  Suspending is similar to putting a tab to sleep in the Edge browser. Suspending pauses
	//  WebView script timers and animations, minimizes CPU usage for the associated
	//  browser renderer process and allows the operating system to reuse the memory that was
	//  used by the renderer process for other processes.
	//  Note that Suspend is best effort and considered completed successfully once the request
	//  is sent to browser renderer process. If there is a running script, the script will continue
	//  to run and the renderer process will be suspended after that script is done.
	//  See [Sleeping Tabs FAQ](https://techcommunity.microsoft.com/t5/articles/sleeping-tabs-faq/m-p/1705434)
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
	TrySuspend(aHandler ICoreWebView2TrySuspendCompletedHandler) bool // function
	// Resume
	//  Resumes the WebView so that it resumes activities on the web page.
	//  This API can be called while the WebView2 controller is invisible.
	//  The app can interact with the WebView immediately after `Resume`.
	//  WebView will be automatically resumed when it becomes visible.
	Resume() bool // function
	// SetVirtualHostNameToFolderMapping
	//  Sets a mapping between a virtual host name and a folder path to make available to web sites
	//  via that host name.
	//  After setting the mapping, documents loaded in the WebView can use HTTP or HTTPS URLs at
	//  the specified host name specified by hostName to access files in the local folder specified
	//  by folderPath.
	//  This mapping applies to both top-level document and iframe navigations as well as subresource
	//  references from a document. This also applies to web workers including dedicated/shared worker
	//  and service worker, for loading either worker scripts or subresources
	// (importScripts(), fetch(), XHR, etc.) issued from within a worker.
	//  For virtual host mapping to work with service worker, please keep the virtual host name
	//  mappings consistent among all WebViews sharing the same browser instance. As service worker
	//  works independently of WebViews, we merge mappings from all WebViews when resolving virtual
	//  host name, inconsistent mappings between WebViews would lead unexpected behavior.
	//  Due to a current implementation limitation, media files accessed using virtual host name can be
	//  very slow to load.
	//  As the resource loaders for the current page might have already been created and running,
	//  changes to the mapping might not be applied to the current page and a reload of the page is
	//  needed to apply the new mapping.
	//  Both absolute and relative paths are supported for folderPath. Relative paths are interpreted
	//  as relative to the folder where the exe of the app is in.
	//  Note that the folderPath length must not exceed the Windows MAX_PATH limit.
	//  accessKind specifies the level of access to resources under the virtual host from other sites.
	//  For example, after calling
	//  <code>
	//  ```cpp
	//  SetVirtualHostNameToFolderMapping(
	//  L"appassets.example", L"assets",
	//  COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY);
	//  ```
	//  </code>
	//  navigating to `https://appassets.example/my-local-file.html` will
	//  show the content from my-local-file.html in the assets subfolder located on disk under
	//  the same path as the app's executable file.
	//  DOM elements that want to reference local files will have their host reference virtual host in the source.
	//  If there are multiple folders being used, define one unique virtual host per folder.
	//  For example, you can embed a local image like this
	//  <code>
	//  ```cpp
	//  WCHAR c_navString[] = L"<img src=\"http://appassets.example/wv2.png\"/>";
	//  m_webView->NavigateToString(c_navString);
	//  ```
	//  </code>
	//  The example above shows the image wv2.png by resolving the folder mapping above.
	//  You should typically choose virtual host names that are never used by real sites.
	//  If you own a domain such as example.com, another option is to use a subdomain reserved for
	//  the app(like my-app.example.com).
	//  [RFC 6761](https://tools.ietf.org/html/rfc6761) has reserved several special-use domain
	//  names that are guaranteed to not be used by real sites(for example, .example, .test, and
	//  .invalid.)
	//  Note that using `.local` as the top-level domain name will work but can cause a delay
	//  during navigations. You should avoid using `.local` if you can.
	//  Apps should use distinct domain names when mapping folder from different sources that
	//  should be isolated from each other. For instance, the app might use app-file.example for
	//  files that ship as part of the app, and book1.example might be used for files containing
	//  books from a less trusted source that were previously downloaded and saved to the disk by
	//  the app.
	//  The host name used in the APIs is canonicalized using Chromium's host name parsing logic
	//  before being used internally. For more information see [HTML5 2.6 URLs](https://dev.w3.org/html5/spec-LC/urls.html).
	//  All host names that are canonicalized to the same string are considered identical.
	//  For example, `EXAMPLE.COM` and `example.com` are treated as the same host name.
	//  An international host name and its Punycode-encoded host name are considered the same host
	//  name. There is no DNS resolution for host name and the trailing '.' is not normalized as
	//  part of canonicalization.
	//  Therefore `example.com` and `example.com.` are treated as different host names. Similarly,
	//  `virtual-host-name` and `virtual-host-name.example.com` are treated as different host names
	//  even if the machine has a DNS suffix of `example.com`.
	//  Specify the minimal cross-origin access necessary to run the app. If there is not a need to
	//  access local resources from other origins, use COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY.
	SetVirtualHostNameToFolderMapping(aHostName, aFolderPath string, aAccessKind TWVHostResourceAcccessKind) bool // function
	// ClearVirtualHostNameToFolderMapping
	//  Clears a host name mapping for local folder that was added by `SetVirtualHostNameToFolderMapping`.
	ClearVirtualHostNameToFolderMapping(aHostName string) bool // function
	// OpenTaskManagerWindow
	//  Opens the Browser Task Manager view as a new window in the foreground.
	//  If the Browser Task Manager is already open, this will bring it into
	//  the foreground. WebView2 currently blocks the Shift+Esc shortcut for
	//  opening the task manager. An end user can open the browser task manager
	//  manually via the `Browser task manager` entry of the DevTools window's
	//  title bar's context menu.
	OpenTaskManagerWindow() bool // function
	// PrintToPdf
	//  Print the current page to PDF asynchronously with the provided settings.
	//  See `ICoreWebView2PrintSettings` for description of settings. Passing
	//  nullptr for `printSettings` results in default print settings used.
	//  Use `resultFilePath` to specify the path to the PDF file. The host should
	//  provide an absolute path, including file name. If the path
	//  points to an existing file, the file will be overwritten. If the path is
	//  not valid, the method fails with `E_INVALIDARG`.
	//  The async `PrintToPdf` operation completes when the data has been written
	//  to the PDF file. At this time the
	//  `ICoreWebView2PrintToPdfCompletedHandler` is invoked. If the
	//  application exits before printing is complete, the file is not saved.
	//  Only one `Printing` operation can be in progress at a time. If
	//  `PrintToPdf` is called while a `PrintToPdf` or `PrintToPdfStream` or `Print` or
	//  `ShowPrintUI` job is in progress, the completed handler is immediately invoked
	//  with `isSuccessful` set to FALSE.
	PrintToPdf(aResultFilePath string, aPrintSettings ICoreWebView2PrintSettings, aHandler ICoreWebView2PrintToPdfCompletedHandler) bool // function
	// OpenDevToolsWindow
	//  Opens the DevTools window for the current document in the WebView. Does
	//  nothing if run when the DevTools window is already open.
	OpenDevToolsWindow() bool // function
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
	//  `ICoreWebView2Settings::IsWebMessageEnabled` setting must be `TRUE` or
	//  the web message will not be sent. The `data` property of the event
	//  arg is the `webMessage` string parameter parsed as a JSON string into a
	//  JavaScript object. The `source` property of the event arg is a reference
	//  to the `window.chrome.webview` object. For information about sending
	//  messages from the HTML document in the WebView to the host, navigate to
	//  [add_WebMessageReceived](/microsoft-edge/webview2/reference/win32/icorewebview2#add_webmessagereceived).
	//  The message is delivered asynchronously. If a navigation occurs before
	//  the message is posted to the page, the message is discarded.
	PostWebMessageAsJson(aWebMessageAsJson string) bool // function
	// PostWebMessageAsString
	//  Posts a message that is a simple string rather than a JSON string
	//  representation of a JavaScript object. This behaves in exactly the same
	//  manner as `PostWebMessageAsJson`, but the `data` property of the event
	//  arg of the `window.chrome.webview` message is a string with the same
	//  value as `webMessageAsString`. Use this instead of
	//  `PostWebMessageAsJson` if you want to communicate using simple strings
	//  rather than JSON objects.
	PostWebMessageAsString(aWebMessageAsString string) bool // function
	// CallDevToolsProtocolMethod
	//  Runs an asynchronous `DevToolsProtocol` method. For more information
	//  about available methods, navigate to
	//  [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/tot).
	//  The `methodName` parameter is the full name of the method in the
	//  `{domain}.{method}` format. The `parametersAsJson` parameter is a JSON
	//  formatted string containing the parameters for the corresponding method.
	//  The `Invoke` method of the `handler` is run when the method
	//  asynchronously completes. `Invoke` is run with the return object of the
	//  method as a JSON string. This function returns E_INVALIDARG if the `methodName` is
	//  unknown or the `parametersAsJson` has an error. In the case of such an error, the
	//  `returnObjectAsJson` parameter of the handler will include information
	//  about the error.
	//  Note even though WebView2 dispatches the CDP messages in the order called,
	//  CDP method calls may be processed out of order.
	//  If you require CDP methods to run in a particular order, you should wait
	//  for the previous method's completed handler to run before calling the
	//  next method.
	CallDevToolsProtocolMethod(aMethodName, aParametersAsJson string, aExecutionID int32, aBrowserComponent IComponent) bool // function
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
	//  For more information about targets and sessions, navigate to
	//  [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/tot/Target).
	//  For more information about available methods, navigate to
	//  [DevTools Protocol Viewer](https://chromedevtools.github.io/devtools-protocol/tot)
	//  The `sessionId` parameter is the sessionId for an attached target.
	//  nullptr or empty string is treated as the session for the default target for the top page.
	//  The `methodName` parameter is the full name of the method in the
	//  `{domain}.{method}` format. The `parametersAsJson` parameter is a JSON
	//  formatted string containing the parameters for the corresponding method.
	//  The `Invoke` method of the `handler` is run when the method
	//  asynchronously completes. `Invoke` is run with the return object of the
	//  method as a JSON string. This function returns E_INVALIDARG if the `methodName` is
	//  unknown or the `parametersAsJson` has an error. In the case of such an error, the
	//  `returnObjectAsJson` parameter of the handler will include information
	//  about the error.
	CallDevToolsProtocolMethodForSession(aSessionId, aMethodName, aParametersAsJson string, aExecutionID int32, aBrowserComponent IComponent) bool // function
	// AddWebResourceRequestedFilter
	//  Adds a URI and resource context filter for the `WebResourceRequested`
	//  event. A web resource request with a resource context that matches this
	//  filter's resource context and a URI that matches this filter's URI
	//  wildcard string will be raised via the `WebResourceRequested` event.
	//  The `uri` parameter value is a wildcard string matched against the URI
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
	//  Specifying a `nullptr` for the uri is equivalent to an empty string which
	//  matches no URIs.
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
	AddWebResourceRequestedFilter(URI string, ResourceContext TWVWebResourceContext) bool // function
	// RemoveWebResourceRequestedFilter
	//  Removes a matching WebResource filter that was previously added for the
	//  `WebResourceRequested` event. If the same filter was added multiple
	//  times, then it must be removed as many times as it was added for the
	//  removal to be effective. Returns `E_INVALIDARG` for a filter that was
	//  never added.
	RemoveWebResourceRequestedFilter(URI string, ResourceContext TWVWebResourceContext) bool // function
	// RemoveHostObjectFromScript
	//  Remove the host object specified by the name so that it is no longer
	//  accessible from JavaScript code in the WebView. While new access
	//  attempts are denied, if the object is already obtained by JavaScript code
	//  in the WebView, the JavaScript code continues to have access to that
	//  object. Run this method for a name that is already removed or never
	//  added fails.
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
	//  If the method is run in add_NewWindowRequested handler it should be called
	//  before the new window is set. If called after setting the NewWindow property, the initial script
	//  may or may not apply to the initial navigation and may only apply to the subsequent navigation.
	//  For more details see `ICoreWebView2NewWindowRequestedEventArgs::put_NewWindow`.
	//  NOTE: If an HTML document is running in a sandbox of some kind using
	//  [sandbox](https://developer.mozilla.org/docs/Web/HTML/Element/iframe#attr-sandbox)
	//  properties or the
	//  [Content-Security-Policy](https://developer.mozilla.org/docs/Web/HTTP/Headers/Content-Security-Policy)
	//  HTTP header affects the script that runs. For example, if the
	//  `allow-modals` keyword is not set then requests to run the `alert`
	//  function are ignored.
	AddScriptToExecuteOnDocumentCreated(JavaScript string, aBrowserComponent IComponent) bool // function
	// RemoveScriptToExecuteOnDocumentCreated
	//  Remove the corresponding JavaScript added using
	//  `AddScriptToExecuteOnDocumentCreated` with the specified script ID. The
	//  script ID should be the one returned by the `AddScriptToExecuteOnDocumentCreated`.
	//  Both use `AddScriptToExecuteOnDocumentCreated` and this method in `NewWindowRequested`
	//  event handler at the same time sometimes causes trouble. Since invalid scripts will
	//  be ignored, the script IDs you got may not be valid anymore.
	RemoveScriptToExecuteOnDocumentCreated(aID string) bool // function
	// OpenDefaultDownloadDialog
	//  Open the default download dialog. If the dialog is opened before there
	//  are recent downloads, the dialog shows all past downloads for the
	//  current profile. Otherwise, the dialog shows only the recent downloads
	//  with a "See more" button for past downloads. Calling this method raises
	//  the `IsDefaultDownloadDialogOpenChanged` event if the dialog was closed.
	//  No effect if the dialog is already open.
	OpenDefaultDownloadDialog() bool // function
	// CloseDefaultDownloadDialog
	//  Close the default download dialog. Calling this method raises the
	//  `IsDefaultDownloadDialogOpenChanged` event if the dialog was open. No
	//  effect if the dialog is already closed.
	CloseDefaultDownloadDialog() bool // function
	// ClearServerCertificateErrorActions
	//  Clears all cached decisions to proceed with TLS certificate errors from the
	//  ServerCertificateErrorDetected event for all WebView2's sharing the same session.
	ClearServerCertificateErrorActions(aBrowserComponent IComponent) bool // function
	// GetFavicon
	//  Async function for getting the actual image data of the favicon.
	//  The image is copied to the `imageStream` object in `ICoreWebView2GetFaviconCompletedHandler`.
	//  If there is no image then no data would be copied into the imageStream.
	//  The `format` is the file format to return the image stream.
	//  `completedHandler` is executed at the end of the operation.
	GetFavicon(aFormat TWVFaviconImageFormat, aBrowserComponent IComponent) bool // function
	// Print
	//  Print the current web page asynchronously to the specified printer with the provided settings.
	//  See `ICoreWebView2PrintSettings` for description of settings. Passing
	//  nullptr for `printSettings` results in default print settings used.
	//  The handler will return `errorCode` as `S_OK` and `printStatus` as COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE
	//  if `printerName` doesn't match with the name of any installed printers on the user OS. The handler
	//  will return `errorCode` as `E_INVALIDARG` and `printStatus` as COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR
	//  if the caller provides invalid settings for a given printer.
	//  The async `Print` operation completes when it finishes printing to the printer.
	//  At this time the `ICoreWebView2PrintCompletedHandler` is invoked.
	//  Only one `Printing` operation can be in progress at a time. If `Print` is called while a `Print` or `PrintToPdf`
	//  or `PrintToPdfStream` or `ShowPrintUI` job is in progress, the completed handler is immediately invoked
	//  with `E_ABORT` and `printStatus` is COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR.
	//  This is only for printing operation on one webview.
	//  <code>
	//  | errorCode | printStatus | Notes |
	//  | --- | --- | --- |
	//  | S_OK | COREWEBVIEW2_PRINT_STATUS_SUCCEEDED | Print operation succeeded. |
	//  | S_OK | COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE | If specified printer is not found or printer status is not available, offline or error state. |
	//  | S_OK | COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR | Print operation is failed. |
	//  | E_INVALIDARG | COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR | If the caller provides invalid settings for the specified printer. |
	//  | E_ABORT | COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR | Print operation is failed as printing job already in progress. |
	//  </code>
	Print(aPrintSettings ICoreWebView2PrintSettings, aHandler ICoreWebView2PrintCompletedHandler) bool // function
	// ShowPrintUI
	//  Opens the print dialog to print the current web page. See `COREWEBVIEW2_PRINT_DIALOG_KIND`
	//  for descriptions of print dialog kinds.
	//  Invoking browser or system print dialog doesn't open new print dialog if
	//  it is already open.
	ShowPrintUI(aPrintDialogKind TWVPrintDialogKind) bool // function
	// PrintToPdfStream
	//  Provides the Pdf data of current web page asynchronously for the provided settings.
	//  Stream will be rewound to the start of the pdf data.
	//  See `ICoreWebView2PrintSettings` for description of settings. Passing
	//  nullptr for `printSettings` results in default print settings used.
	//  The async `PrintToPdfStream` operation completes when it finishes
	//  writing to the stream. At this time the `ICoreWebView2PrintToPdfStreamCompletedHandler`
	//  is invoked. Only one `Printing` operation can be in progress at a time. If
	//  `PrintToPdfStream` is called while a `PrintToPdfStream` or `PrintToPdf` or `Print`
	//  or `ShowPrintUI` job is in progress, the completed handler is immediately invoked with `E_ABORT`.
	//  This is only for printing operation on one webview.
	PrintToPdfStream(aPrintSettings ICoreWebView2PrintSettings, aHandler ICoreWebView2PrintToPdfStreamCompletedHandler) bool // function
	// PostSharedBufferToScript
	//  Share a shared buffer object with script of the main frame in the WebView.
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
	PostSharedBufferToScript(aSharedBuffer ICoreWebView2SharedBuffer, aAccess TWVSharedBufferAccess, aAdditionalDataAsJson string) bool // function
	// ExecuteScriptWithResult
	//  Run JavaScript code from the JavaScript parameter in the current
	//  top-level document rendered in the WebView.
	//  The result of the execution is returned asynchronously in the CoreWebView2ExecuteScriptResult object
	//  which has methods and properties to obtain the successful result of script execution as well as any
	//  unhandled JavaScript exceptions.
	//  If this method is
	//  run after the NavigationStarting event during a navigation, the script
	//  runs in the new document when loading it, around the time
	//  ContentLoading is run. This operation executes the script even if
	//  ICoreWebView2Settings::IsScriptEnabled is set to FALSE.
	//  \snippet ScriptComponent.cpp ExecuteScriptWithResult
	ExecuteScriptWithResult(JavaScript string, aExecutionID int32, aBrowserComponent IComponent) bool // function
}

// TCoreWebView2 Parent: TObject
type TCoreWebView2 struct {
	TObject
}

func NewCoreWebView2(aBaseIntf ICoreWebView2) ICoreWebView2 {
	r1 := WV().SysCallN(736, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2(r1)
}

func (m *TCoreWebView2) Initialized() bool {
	r1 := WV().SysCallN(748, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) BaseIntf() ICoreWebView2 {
	var resultCoreWebView2 uintptr
	WV().SysCallN(723, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2)))
	return AsCoreWebView2(resultCoreWebView2)
}

func (m *TCoreWebView2) Settings() ICoreWebView2Settings {
	var resultCoreWebView2Settings uintptr
	WV().SysCallN(773, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Settings)))
	return AsCoreWebView2Settings(resultCoreWebView2Settings)
}

func (m *TCoreWebView2) BrowserProcessID() DWORD {
	r1 := WV().SysCallN(724, m.Instance())
	return DWORD(r1)
}

func (m *TCoreWebView2) CanGoBack() bool {
	r1 := WV().SysCallN(727, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) CanGoForward() bool {
	r1 := WV().SysCallN(728, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) ContainsFullScreenElement() bool {
	r1 := WV().SysCallN(734, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) DocumentTitle() string {
	r1 := WV().SysCallN(739, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2) Source() string {
	r1 := WV().SysCallN(775, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2) CookieManager() ICoreWebView2CookieManager {
	var resultCoreWebView2CookieManager uintptr
	WV().SysCallN(735, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CookieManager)))
	return AsCoreWebView2CookieManager(resultCoreWebView2CookieManager)
}

func (m *TCoreWebView2) Environment() ICoreWebView2Environment {
	var resultCoreWebView2Environment uintptr
	WV().SysCallN(740, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Environment)))
	return AsCoreWebView2Environment(resultCoreWebView2Environment)
}

func (m *TCoreWebView2) IsSuspended() bool {
	r1 := WV().SysCallN(752, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) IsMuted() bool {
	r1 := WV().SysCallN(751, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2) SetIsMuted(AValue bool) {
	WV().SysCallN(751, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2) IsDocumentPlayingAudio() bool {
	r1 := WV().SysCallN(750, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) IsDefaultDownloadDialogOpen() bool {
	r1 := WV().SysCallN(749, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) DefaultDownloadDialogCornerAlignment() TWVDefaultDownloadDialogCornerAlignment {
	r1 := WV().SysCallN(737, 0, m.Instance(), 0)
	return TWVDefaultDownloadDialogCornerAlignment(r1)
}

func (m *TCoreWebView2) SetDefaultDownloadDialogCornerAlignment(AValue TWVDefaultDownloadDialogCornerAlignment) {
	WV().SysCallN(737, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2) DefaultDownloadDialogMargin() (resultPoint TPoint) {
	WV().SysCallN(738, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2) SetDefaultDownloadDialogMargin(AValue *TPoint) {
	WV().SysCallN(738, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2) StatusBarText() string {
	r1 := WV().SysCallN(776, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2) Profile() ICoreWebView2Profile {
	var resultCoreWebView2Profile uintptr
	WV().SysCallN(766, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Profile)))
	return AsCoreWebView2Profile(resultCoreWebView2Profile)
}

func (m *TCoreWebView2) FaviconURI() string {
	r1 := WV().SysCallN(743, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2) MemoryUsageTargetLevel() TWVMemoryUsageTargetLevel {
	r1 := WV().SysCallN(753, 0, m.Instance(), 0)
	return TWVMemoryUsageTargetLevel(r1)
}

func (m *TCoreWebView2) SetMemoryUsageTargetLevel(AValue TWVMemoryUsageTargetLevel) {
	WV().SysCallN(753, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2) FrameId() uint32 {
	r1 := WV().SysCallN(744, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(720, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) SubscribeToDevToolsProtocolEvent(aEventName string, aEventID int32, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(778, m.Instance(), PascalStr(aEventName), uintptr(aEventID), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) CapturePreview(aImageFormat TWVCapturePreviewImageFormat, aImageStream IStream, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(729, m.Instance(), uintptr(aImageFormat), GetObjectUintptr(aImageStream), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) ExecuteScript(JavaScript string, aExecutionID int32, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(741, m.Instance(), PascalStr(JavaScript), uintptr(aExecutionID), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) GoBack() bool {
	r1 := WV().SysCallN(746, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) GoForward() bool {
	r1 := WV().SysCallN(747, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) Navigate(aURI string) bool {
	r1 := WV().SysCallN(754, m.Instance(), PascalStr(aURI))
	return GoBool(r1)
}

func (m *TCoreWebView2) NavigateToString(aHTMLContent string) bool {
	r1 := WV().SysCallN(755, m.Instance(), PascalStr(aHTMLContent))
	return GoBool(r1)
}

func (m *TCoreWebView2) NavigateWithWebResourceRequest(aRequest ICoreWebView2WebResourceRequestRef) bool {
	r1 := WV().SysCallN(756, m.Instance(), GetObjectUintptr(aRequest))
	return GoBool(r1)
}

func (m *TCoreWebView2) Reload() bool {
	r1 := WV().SysCallN(767, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) Stop() bool {
	r1 := WV().SysCallN(777, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) TrySuspend(aHandler ICoreWebView2TrySuspendCompletedHandler) bool {
	r1 := WV().SysCallN(779, m.Instance(), GetObjectUintptr(aHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2) Resume() bool {
	r1 := WV().SysCallN(771, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) SetVirtualHostNameToFolderMapping(aHostName, aFolderPath string, aAccessKind TWVHostResourceAcccessKind) bool {
	r1 := WV().SysCallN(772, m.Instance(), PascalStr(aHostName), PascalStr(aFolderPath), uintptr(aAccessKind))
	return GoBool(r1)
}

func (m *TCoreWebView2) ClearVirtualHostNameToFolderMapping(aHostName string) bool {
	r1 := WV().SysCallN(732, m.Instance(), PascalStr(aHostName))
	return GoBool(r1)
}

func (m *TCoreWebView2) OpenTaskManagerWindow() bool {
	r1 := WV().SysCallN(759, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) PrintToPdf(aResultFilePath string, aPrintSettings ICoreWebView2PrintSettings, aHandler ICoreWebView2PrintToPdfCompletedHandler) bool {
	r1 := WV().SysCallN(764, m.Instance(), PascalStr(aResultFilePath), GetObjectUintptr(aPrintSettings), GetObjectUintptr(aHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2) OpenDevToolsWindow() bool {
	r1 := WV().SysCallN(758, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) PostWebMessageAsJson(aWebMessageAsJson string) bool {
	r1 := WV().SysCallN(761, m.Instance(), PascalStr(aWebMessageAsJson))
	return GoBool(r1)
}

func (m *TCoreWebView2) PostWebMessageAsString(aWebMessageAsString string) bool {
	r1 := WV().SysCallN(762, m.Instance(), PascalStr(aWebMessageAsString))
	return GoBool(r1)
}

func (m *TCoreWebView2) CallDevToolsProtocolMethod(aMethodName, aParametersAsJson string, aExecutionID int32, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(725, m.Instance(), PascalStr(aMethodName), PascalStr(aParametersAsJson), uintptr(aExecutionID), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) CallDevToolsProtocolMethodForSession(aSessionId, aMethodName, aParametersAsJson string, aExecutionID int32, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(726, m.Instance(), PascalStr(aSessionId), PascalStr(aMethodName), PascalStr(aParametersAsJson), uintptr(aExecutionID), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) AddWebResourceRequestedFilter(URI string, ResourceContext TWVWebResourceContext) bool {
	r1 := WV().SysCallN(722, m.Instance(), PascalStr(URI), uintptr(ResourceContext))
	return GoBool(r1)
}

func (m *TCoreWebView2) RemoveWebResourceRequestedFilter(URI string, ResourceContext TWVWebResourceContext) bool {
	r1 := WV().SysCallN(770, m.Instance(), PascalStr(URI), uintptr(ResourceContext))
	return GoBool(r1)
}

func (m *TCoreWebView2) RemoveHostObjectFromScript(aName string) bool {
	r1 := WV().SysCallN(768, m.Instance(), PascalStr(aName))
	return GoBool(r1)
}

func (m *TCoreWebView2) AddScriptToExecuteOnDocumentCreated(JavaScript string, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(721, m.Instance(), PascalStr(JavaScript), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) RemoveScriptToExecuteOnDocumentCreated(aID string) bool {
	r1 := WV().SysCallN(769, m.Instance(), PascalStr(aID))
	return GoBool(r1)
}

func (m *TCoreWebView2) OpenDefaultDownloadDialog() bool {
	r1 := WV().SysCallN(757, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) CloseDefaultDownloadDialog() bool {
	r1 := WV().SysCallN(733, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2) ClearServerCertificateErrorActions(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(731, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) GetFavicon(aFormat TWVFaviconImageFormat, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(745, m.Instance(), uintptr(aFormat), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2) Print(aPrintSettings ICoreWebView2PrintSettings, aHandler ICoreWebView2PrintCompletedHandler) bool {
	r1 := WV().SysCallN(763, m.Instance(), GetObjectUintptr(aPrintSettings), GetObjectUintptr(aHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2) ShowPrintUI(aPrintDialogKind TWVPrintDialogKind) bool {
	r1 := WV().SysCallN(774, m.Instance(), uintptr(aPrintDialogKind))
	return GoBool(r1)
}

func (m *TCoreWebView2) PrintToPdfStream(aPrintSettings ICoreWebView2PrintSettings, aHandler ICoreWebView2PrintToPdfStreamCompletedHandler) bool {
	r1 := WV().SysCallN(765, m.Instance(), GetObjectUintptr(aPrintSettings), GetObjectUintptr(aHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2) PostSharedBufferToScript(aSharedBuffer ICoreWebView2SharedBuffer, aAccess TWVSharedBufferAccess, aAdditionalDataAsJson string) bool {
	r1 := WV().SysCallN(760, m.Instance(), GetObjectUintptr(aSharedBuffer), uintptr(aAccess), PascalStr(aAdditionalDataAsJson))
	return GoBool(r1)
}

func (m *TCoreWebView2) ExecuteScriptWithResult(JavaScript string, aExecutionID int32, aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(742, m.Instance(), PascalStr(JavaScript), uintptr(aExecutionID), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func CoreWebView2Class() TClass {
	ret := WV().SysCallN(730)
	return TClass(ret)
}
