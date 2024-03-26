//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

// IWVBrowserEvents
//
//	{4E06D91F-1213-46C1-ABB8-D41D8CC19E81}
//	Proxy Object Event Interface, Event callback from: IWVBrowser IWVBrowserBase
type IWVBrowserEvents interface {
	IComponent
	// Instance
	//  return instance uintptr
	Instance() uintptr
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

// NewWVBrowserDelegateEvents
//
//	IWVBrowserEvents 代理对象，设置事件接收对象
//	该对象和默认的浏览器对象不同，它被创建后是单独的，默认不会执行浏览器事件
//	只有在用户设置事件接收对象时有效，且会触发对应的事件
func NewWVBrowserDelegateEvents() IWVBrowserEvents {
	return NewWVBrowser(nil)
}
