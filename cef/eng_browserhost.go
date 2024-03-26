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

// ICefBrowserHost Parent: ICefBaseRefCounted
//
//	Interface used to represent the browser process aspects of a browser. The functions of this interface can only be called in the browser process. They may be called on any thread in that process unless otherwise indicated in the comments.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_browser_host_t))
type ICefBrowserHost interface {
	ICefBaseRefCounted
	IMESetComposition(text string, underlines TCefCompositionUnderlineDynArray, replacementrange, selectionrange *TCefRange)
	// GetBrowser
	//  Returns the hosted browser object.
	GetBrowser() ICefBrowser // function
	// TryCloseBrowser
	//  Helper for closing a browser. Call this function from the top-level window close handler (if any). Internally this calls CloseBrowser(false (0)) if the close has not yet been initiated. This function returns false (0) while the close is pending and true (1) after the close has completed. See CloseBrowser() and ICefLifeSpanHandler.DoClose() documentation for additional usage information. This function must be called on the browser process UI thread.
	TryCloseBrowser() bool // function
	// GetWindowHandle
	//  Retrieve the window handle (if any) for this browser. If this browser is wrapped in a ICefBrowserView this function should be called on the browser process UI thread and it will return the handle for the top-level native window.
	GetWindowHandle() TCefWindowHandle // function
	// GetOpenerWindowHandle
	//  Retrieve the window handle (if any) of the browser that opened this browser. Will return NULL for non-popup browsers or if this browser is wrapped in a ICefBrowserView. This function can be used in combination with custom handling of modal windows.
	GetOpenerWindowHandle() TCefWindowHandle // function
	// HasView
	//  Returns true (1) if this browser is wrapped in a ICefBrowserView.
	HasView() bool // function
	// GetClient
	//  Returns the client for this browser.
	GetClient() ICefClient // function
	// GetRequestContext
	//  Returns the request context for this browser.
	GetRequestContext() ICefRequestContext // function
	// CanZoom
	//  Returns true (1) if this browser can execute the specified zoom command. This function can only be called on the UI thread.
	CanZoom(command TCefZoomCommand) bool // function
	// GetDefaultZoomLevel
	//  Get the default zoom level. This value will be 0.0 by default but can be configured with the Chrome runtime. This function can only be called on the UI thread.
	GetDefaultZoomLevel() (resultFloat64 float64) // function
	// GetZoomLevel
	//  Get the current zoom level. This function can only be called on the UI thread.
	GetZoomLevel() (resultFloat64 float64) // function
	// HasDevTools
	//  Returns true (1) if this browser currently has an associated DevTools browser. Must be called on the browser process UI thread.
	HasDevTools() bool // function
	// SendDevToolsMessage
	//  Send a function call message over the DevTools protocol. |message| must be a UTF8-encoded JSON dictionary that contains "id" (int), "function" (string) and "params" (dictionary, optional) values. See the DevTools protocol documentation at https://chromedevtools.github.io/devtools- protocol/ for details of supported functions and the expected "params" dictionary contents. |message| will be copied if necessary. This function will return true (1) if called on the UI thread and the message was successfully submitted for validation, otherwise false (0). Validation will be applied asynchronously and any messages that fail due to formatting errors or missing parameters may be discarded without notification. Prefer ExecuteDevToolsMethod if a more structured approach to message formatting is desired. Every valid function call will result in an asynchronous function result or error message that references the sent message "id". Event messages are received while notifications are enabled (for example, between function calls for "Page.enable" and "Page.disable"). All received messages will be delivered to the observer(s) registered with AddDevToolsMessageObserver. See ICefDevToolsMessageObserver.OnDevToolsMessage documentation for details of received message contents. Usage of the SendDevToolsMessage, ExecuteDevToolsMethod and AddDevToolsMessageObserver functions does not require an active DevTools front-end or remote-debugging session. Other active DevTools sessions will continue to function independently. However, any modification of global browser state by one session may not be reflected in the UI of other sessions. Communication with the DevTools front-end (when displayed) can be logged for development purposes by passing the `--devtools-protocol-log- file=<path>` command-line flag.
	SendDevToolsMessage(message string) bool // function
	// ExecuteDevToolsMethod
	//  Execute a function call over the DevTools protocol. This is a more structured version of SendDevToolsMessage. |message_id| is an incremental number that uniquely identifies the message (pass 0 to have the next number assigned automatically based on previous values). |function| is the function name. |params| are the function parameters, which may be NULL. See the DevTools protocol documentation (linked above) for details of supported functions and the expected |params| dictionary contents. This function will return the assigned message ID if called on the UI thread and the message was successfully submitted for validation, otherwise 0. See the SendDevToolsMessage documentation for additional usage information.
	ExecuteDevToolsMethod(messageid int32, method string, params ICefDictionaryValue) int32 // function
	// AddDevToolsMessageObserver
	//  Add an observer for DevTools protocol messages (function results and events). The observer will remain registered until the returned Registration object is destroyed. See the SendDevToolsMessage documentation for additional usage information.
	AddDevToolsMessageObserver(observer ICefDevToolsMessageObserver) ICefRegistration // function
	// IsWindowRenderingDisabled
	//  Returns true (1) if window rendering is disabled.
	IsWindowRenderingDisabled() bool // function
	// GetWindowlessFrameRate
	//  Returns the maximum rate in frames per second (fps) that ICefRenderHandler.OnPaint will be called for a windowless browser. The actual fps may be lower if the browser cannot generate frames at the requested rate. The minimum value is 1 and the maximum value is 60 (default 30). This function can only be called on the UI thread.
	GetWindowlessFrameRate() int32 // function
	// GetVisibleNavigationEntry
	//  Returns the current visible navigation entry for this browser. This function can only be called on the UI thread.
	GetVisibleNavigationEntry() ICefNavigationEntry // function
	// GetExtension
	//  Returns the extension hosted in this browser or NULL if no extension is hosted. See ICefRequestContext.LoadExtension for details.
	GetExtension() ICefExtension // function
	// IsBackgroundHost
	//  Returns true (1) if this browser is hosting an extension background script. Background hosts do not have a window and are not displayable. See ICefRequestContext.LoadExtension for details.
	IsBackgroundHost() bool // function
	// IsAudioMuted
	//  Returns true (1) if the browser's audio is muted. This function can only be called on the UI thread.
	IsAudioMuted() bool // function
	// IsFullscreen
	//  Returns true (1) if the renderer is currently in browser fullscreen. This differs from window fullscreen in that browser fullscreen is entered using the JavaScript Fullscreen API and modifies CSS attributes such as the ::backdrop pseudo-element and :fullscreen pseudo-structure. This function can only be called on the UI thread.
	IsFullscreen() bool // function
	// CloseBrowser
	//  Request that the browser close. The JavaScript 'onbeforeunload' event will be fired. If |force_close| is false (0) the event handler, if any, will be allowed to prompt the user and the user can optionally cancel the close. If |force_close| is true (1) the prompt will not be displayed and the close will proceed. Results in a call to ICefLifeSpanHandler.DoClose() if the event handler allows the close or if |force_close| is true (1). See ICefLifeSpanHandler.DoClose() documentation for additional usage information.
	CloseBrowser(forceClose bool) // procedure
	// SetFocus
	//  Set whether the browser is focused.
	SetFocus(focus bool) // procedure
	// Zoom
	//  Execute a zoom command in this browser. If called on the UI thread the change will be applied immediately. Otherwise, the change will be applied asynchronously on the UI thread.
	Zoom(command TCefZoomCommand) // procedure
	// SetZoomLevel
	//  Change the zoom level to the specified value. Specify 0.0 to reset the zoom level to the default. If called on the UI thread the change will be applied immediately. Otherwise, the change will be applied asynchronously on the UI thread.
	SetZoomLevel(zoomLevel float64) // procedure
	// RunFileDialog
	//  Call to run a file chooser dialog. Only a single file chooser dialog may be pending at any given time. |mode| represents the type of dialog to display. |title| to the title to be used for the dialog and may be NULL to show the default title ("Open" or "Save" depending on the mode). |default_file_path| is the path with optional directory and/or file name component that will be initially selected in the dialog. |accept_filters| are used to restrict the selectable file types and may any combination of (a) valid lower-cased MIME types (e.g. "text/*" or "image/*"), (b) individual file extensions (e.g. ".txt" or ".png"), or (c) combined description and file extension delimited using "|" and ";" (e.g. "Image Types|.png;.gif;.jpg"). |callback| will be executed after the dialog is dismissed or immediately if another dialog is already pending. The dialog will be initiated asynchronously on the UI thread.
	RunFileDialog(mode TCefFileDialogMode, title, defaultFilePath string, acceptFilters IStrings, callback ICefRunFileDialogCallback) // procedure
	// StartDownload
	//  Download the file at |url| using ICefDownloadHandler.
	StartDownload(url string) // procedure
	// DownloadImage
	//  Download |image_url| and execute |callback| on completion with the images received from the renderer. If |is_favicon| is true (1) then cookies are not sent and not accepted during download. Images with density independent pixel (DIP) sizes larger than |max_image_size| are filtered out from the image results. Versions of the image at different scale factors may be downloaded up to the maximum scale factor supported by the system. If there are no image results <= |max_image_size| then the smallest image is resized to |max_image_size| and is the only result. A |max_image_size| of 0 means unlimited. If |bypass_cache| is true (1) then |image_url| is requested from the server even if it is present in the browser cache.
	DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback ICefDownloadImageCallback) // procedure
	// Print
	//  Print the current browser contents.
	Print() // procedure
	// PrintToPdf
	//  Print the current browser contents to the PDF file specified by |path| and execute |callback| on completion. The caller is responsible for deleting |path| when done. For PDF printing to work on Linux you must implement the ICefPrintHandler.GetPdfPaperSize function.
	PrintToPdf(path string, settings *TCefPdfPrintSettings, callback ICefPdfPrintCallback) // procedure
	// Find
	//  Search for |searchText|. |forward| indicates whether to search forward or backward within the page. |matchCase| indicates whether the search should be case-sensitive. |findNext| indicates whether this is the first request or a follow-up. The search will be restarted if |searchText| or |matchCase| change. The search will be stopped if |searchText| is NULL. The ICefFindHandler instance, if any, returned via ICefClient.GetFindHandler will be called to report find results.
	Find(searchText string, forward, matchCase, findNext bool) // procedure
	// StopFinding
	//  Cancel all searches that are currently going on.
	StopFinding(clearSelection bool) // procedure
	// ShowDevTools
	//  Open developer tools (DevTools) in its own browser. The DevTools browser will remain associated with this browser. If the DevTools browser is already open then it will be focused, in which case the |windowInfo|, |client| and |settings| parameters will be ignored. If |inspectElementAt| is non-NULL then the element at the specified (x,y) location will be inspected. The |windowInfo| parameter will be ignored if this browser is wrapped in a ICefBrowserView.
	ShowDevTools(windowInfo *TCefWindowInfo, client ICefClient, settings *TCefBrowserSettings, inspectElementAt *TCefPoint) // procedure
	// CloseDevTools
	//  Explicitly close the associated DevTools browser, if any.
	CloseDevTools() // procedure
	// GetNavigationEntries
	//  Retrieve a snapshot of current navigation entries as values sent to the specified visitor. If |current_only| is true (1) only the current navigation entry will be sent, otherwise all navigation entries will be sent.
	GetNavigationEntries(visitor ICefNavigationEntryVisitor, currentOnly bool) // procedure
	// ReplaceMisspelling
	//  If a misspelled word is currently selected in an editable node calling this function will replace it with the specified |word|.
	ReplaceMisspelling(word string) // procedure
	// AddWordToDictionary
	//  Add the specified |word| to the spelling dictionary.
	AddWordToDictionary(word string) // procedure
	// WasResized
	//  Notify the browser that the widget has been resized. The browser will first call ICefRenderHandler.GetViewRect to get the new size and then call ICefRenderHandler.OnPaint asynchronously with the updated regions. This function is only used when window rendering is disabled.
	WasResized() // procedure
	// NotifyScreenInfoChanged
	//  Send a notification to the browser that the screen info has changed. The browser will then call ICefRenderHandler.GetScreenInfo to update the screen information with the new values. This simulates moving the webview window from one display to another, or changing the properties of the current display. This function is only used when window rendering is disabled.
	NotifyScreenInfoChanged() // procedure
	// WasHidden
	//  Notify the browser that it has been hidden or shown. Layouting and ICefRenderHandler.OnPaint notification will stop when the browser is hidden. This function is only used when window rendering is disabled.
	WasHidden(hidden bool) // procedure
	// Invalidate
	//  Invalidate the view. The browser will call ICefRenderHandler.OnPaint asynchronously. This function is only used when window rendering is disabled.
	Invalidate(type_ TCefPaintElementType) // procedure
	// SendExternalBeginFrame
	//  Issue a BeginFrame request to Chromium. Only valid when TCefWindowInfo.external_begin_frame_enabled is set to true (1).
	SendExternalBeginFrame() // procedure
	// SendKeyEvent
	//  Send a key event to the browser.
	SendKeyEvent(event *TCefKeyEvent) // procedure
	// SendMouseClickEvent
	//  Send a mouse click event to the browser. The |x| and |y| coordinates are relative to the upper-left corner of the view.
	SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) // procedure
	// SendMouseMoveEvent
	//  Send a mouse move event to the browser. The |x| and |y| coordinates are relative to the upper-left corner of the view.
	SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) // procedure
	// SendMouseWheelEvent
	//  Send a mouse wheel event to the browser. The |x| and |y| coordinates are relative to the upper-left corner of the view. The |deltaX| and |deltaY| values represent the movement delta in the X and Y directions respectively. In order to scroll inside select popups with window rendering disabled ICefRenderHandler.GetScreenPoint should be implemented properly.
	SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) // procedure
	// SendTouchEvent
	//  Send a touch event to the browser for a windowless browser.
	SendTouchEvent(event *TCefTouchEvent) // procedure
	// SendCaptureLostEvent
	//  Send a capture lost event to the browser.
	SendCaptureLostEvent() // procedure
	// NotifyMoveOrResizeStarted
	//  Notify the browser that the window hosting it is about to be moved or resized. This function is only used on Windows and Linux.
	NotifyMoveOrResizeStarted() // procedure
	// SetWindowlessFrameRate
	//  Set the maximum rate in frames per second (fps) that ICefRenderHandler.OnPaint will be called for a windowless browser. The actual fps may be lower if the browser cannot generate frames at the requested rate. The minimum value is 1 and the maximum value is 60 (default 30). Can also be set at browser creation via TCefBrowserSettings.windowless_frame_rate.
	SetWindowlessFrameRate(frameRate int32) // procedure
	// IMECommitText
	//  Completes the existing composition by optionally inserting the specified |text| into the composition node. |replacement_range| is an optional range of the existing text that will be replaced. |relative_cursor_pos| is where the cursor will be positioned relative to the current cursor position. See comments on ImeSetComposition for usage. The |replacement_range| and |relative_cursor_pos| values are only used on OS X. This function is only used when window rendering is disabled.
	IMECommitText(text string, replacementrange *TCefRange, relativecursorpos int32) // procedure
	// IMEFinishComposingText
	//  Completes the existing composition by applying the current composition node contents. If |keep_selection| is false (0) the current selection, if any, will be discarded. See comments on ImeSetComposition for usage. This function is only used when window rendering is disabled.
	IMEFinishComposingText(keepselection bool) // procedure
	// IMECancelComposition
	//  Cancels the existing composition and discards the composition node contents without applying them. See comments on ImeSetComposition for usage. This function is only used when window rendering is disabled.
	IMECancelComposition() // procedure
	// DragTargetDragEnter
	//  Call this function when the user drags the mouse into the web view (before calling DragTargetDragOver/DragTargetLeave/DragTargetDrop). |drag_data| should not contain file contents as this type of data is not allowed to be dragged into the web view. File contents can be removed using ICefDragData.ResetFileContents (for example, if |drag_data| comes from ICefRenderHandler.StartDragging). This function is only used when window rendering is disabled.
	DragTargetDragEnter(dragData ICefDragData, event *TCefMouseEvent, allowedOps TCefDragOperations) // procedure
	// DragTargetDragOver
	//  Call this function each time the mouse is moved across the web view during a drag operation (after calling DragTargetDragEnter and before calling DragTargetDragLeave/DragTargetDrop). This function is only used when window rendering is disabled.
	DragTargetDragOver(event *TCefMouseEvent, allowedOps TCefDragOperations) // procedure
	// DragTargetDragLeave
	//  Call this function when the user drags the mouse out of the web view (after calling DragTargetDragEnter). This function is only used when window rendering is disabled.
	DragTargetDragLeave() // procedure
	// DragTargetDrop
	//  Call this function when the user completes the drag operation by dropping the object onto the web view (after calling DragTargetDragEnter). The object being dropped is |drag_data|, given as an argument to the previous DragTargetDragEnter call. This function is only used when window rendering is disabled.
	DragTargetDrop(event *TCefMouseEvent) // procedure
	// DragSourceEndedAt
	//  Call this function when the drag operation started by a ICefRenderHandler.StartDragging call has ended either in a drop or by being cancelled. |x| and |y| are mouse coordinates relative to the upper- left corner of the view. If the web view is both the drag source and the drag target then all DragTarget* functions should be called before DragSource* mthods. This function is only used when window rendering is disabled.
	DragSourceEndedAt(x, y int32, op TCefDragOperation) // procedure
	// DragSourceSystemDragEnded
	//  Call this function when the drag operation started by a ICefRenderHandler.StartDragging call has completed. This function may be called immediately without first calling DragSourceEndedAt to cancel a drag operation. If the web view is both the drag source and the drag target then all DragTarget* functions should be called before DragSource* mthods. This function is only used when window rendering is disabled.
	DragSourceSystemDragEnded() // procedure
	// SetAccessibilityState
	//  Set accessibility state for all frames. |accessibility_state| may be default, enabled or disabled. If |accessibility_state| is STATE_DEFAULT then accessibility will be disabled by default and the state may be further controlled with the "force-renderer-accessibility" and "disable- renderer-accessibility" command-line switches. If |accessibility_state| is STATE_ENABLED then accessibility will be enabled. If |accessibility_state| is STATE_DISABLED then accessibility will be completely disabled. For windowed browsers accessibility will be enabled in Complete mode (which corresponds to kAccessibilityModeComplete in Chromium). In this mode all platform accessibility objects will be created and managed by Chromium's internal implementation. The client needs only to detect the screen reader and call this function appropriately. For example, on macOS the client can handle the @"AXEnhancedUserStructure" accessibility attribute to detect VoiceOver state changes and on Windows the client can handle WM_GETOBJECT with OBJID_CLIENT to detect accessibility readers. For windowless browsers accessibility will be enabled in TreeOnly mode (which corresponds to kAccessibilityModeWebContentsOnly in Chromium). In this mode renderer accessibility is enabled, the full tree is computed, and events are passed to CefAccessibiltyHandler, but platform accessibility objects are not created. The client may implement platform accessibility objects using CefAccessibiltyHandler callbacks if desired.
	SetAccessibilityState(accessibilityState TCefState) // procedure
	// SetAutoResizeEnabled
	//  Enable notifications of auto resize via ICefDisplayHandler.OnAutoResize. Notifications are disabled by default. |min_size| and |max_size| define the range of allowed sizes.
	SetAutoResizeEnabled(enabled bool, minsize, maxsize *TCefSize) // procedure
	// SetAudioMuted
	//  Set whether the browser's audio is muted.
	SetAudioMuted(mute bool) // procedure
	// ExitFullscreen
	//  Requests the renderer to exit browser fullscreen. In most cases exiting window fullscreen should also exit browser fullscreen. With the Alloy runtime this function should be called in response to a user action such as clicking the green traffic light button on MacOS (ICefWindowDelegate.OnWindowFullscreenTransition callback) or pressing the "ESC" key (ICefKeyboardHandler.OnPreKeyEvent callback). With the Chrome runtime these standard exit actions are handled internally but new/additional user actions can use this function. Set |will_cause_resize| to true (1) if exiting browser fullscreen will cause a view resize.
	ExitFullscreen(willcauseresize bool) // procedure
}

// TCefBrowserHost Parent: TCefBaseRefCounted
//
//	Interface used to represent the browser process aspects of a browser. The functions of this interface can only be called in the browser process. They may be called on any thread in that process unless otherwise indicated in the comments.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_browser_host_t))
type TCefBrowserHost struct {
	TCefBaseRefCounted
}

// BrowserHostRef -> ICefBrowserHost
var BrowserHostRef browserHost

// browserHost TCefBrowserHost Ref
type browserHost uintptr

func (m *browserHost) UnWrap(data uintptr) ICefBrowserHost {
	var resultCefBrowserHost uintptr
	CEF().SysCallN(669, uintptr(data), uintptr(unsafePointer(&resultCefBrowserHost)))
	return AsCefBrowserHost(resultCefBrowserHost)
}

func (m *TCefBrowserHost) GetBrowser() ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(624, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCefBrowserHost) TryCloseBrowser() bool {
	r1 := CEF().SysCallN(668, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetWindowHandle() TCefWindowHandle {
	r1 := CEF().SysCallN(632, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCefBrowserHost) GetOpenerWindowHandle() TCefWindowHandle {
	r1 := CEF().SysCallN(629, m.Instance())
	return TCefWindowHandle(r1)
}

func (m *TCefBrowserHost) HasView() bool {
	r1 := CEF().SysCallN(636, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetClient() ICefClient {
	var resultCefClient uintptr
	CEF().SysCallN(625, m.Instance(), uintptr(unsafePointer(&resultCefClient)))
	return AsCefClient(resultCefClient)
}

func (m *TCefBrowserHost) GetRequestContext() ICefRequestContext {
	var resultCefRequestContext uintptr
	CEF().SysCallN(630, m.Instance(), uintptr(unsafePointer(&resultCefRequestContext)))
	return AsCefRequestContext(resultCefRequestContext)
}

func (m *TCefBrowserHost) CanZoom(command TCefZoomCommand) bool {
	r1 := CEF().SysCallN(611, m.Instance(), uintptr(command))
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetDefaultZoomLevel() (resultFloat64 float64) {
	CEF().SysCallN(626, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefBrowserHost) GetZoomLevel() (resultFloat64 float64) {
	CEF().SysCallN(634, m.Instance(), uintptr(unsafePointer(&resultFloat64)))
	return
}

func (m *TCefBrowserHost) HasDevTools() bool {
	r1 := CEF().SysCallN(635, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) SendDevToolsMessage(message string) bool {
	r1 := CEF().SysCallN(652, m.Instance(), PascalStr(message))
	return GoBool(r1)
}

func (m *TCefBrowserHost) ExecuteDevToolsMethod(messageid int32, method string, params ICefDictionaryValue) int32 {
	r1 := CEF().SysCallN(621, m.Instance(), uintptr(messageid), PascalStr(method), GetObjectUintptr(params))
	return int32(r1)
}

func (m *TCefBrowserHost) AddDevToolsMessageObserver(observer ICefDevToolsMessageObserver) ICefRegistration {
	var resultCefRegistration uintptr
	CEF().SysCallN(609, m.Instance(), GetObjectUintptr(observer), uintptr(unsafePointer(&resultCefRegistration)))
	return AsCefRegistration(resultCefRegistration)
}

func (m *TCefBrowserHost) IsWindowRenderingDisabled() bool {
	r1 := CEF().SysCallN(644, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) GetWindowlessFrameRate() int32 {
	r1 := CEF().SysCallN(633, m.Instance())
	return int32(r1)
}

func (m *TCefBrowserHost) GetVisibleNavigationEntry() ICefNavigationEntry {
	var resultCefNavigationEntry uintptr
	CEF().SysCallN(631, m.Instance(), uintptr(unsafePointer(&resultCefNavigationEntry)))
	return AsCefNavigationEntry(resultCefNavigationEntry)
}

func (m *TCefBrowserHost) GetExtension() ICefExtension {
	var resultCefExtension uintptr
	CEF().SysCallN(627, m.Instance(), uintptr(unsafePointer(&resultCefExtension)))
	return AsCefExtension(resultCefExtension)
}

func (m *TCefBrowserHost) IsBackgroundHost() bool {
	r1 := CEF().SysCallN(642, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) IsAudioMuted() bool {
	r1 := CEF().SysCallN(641, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) IsFullscreen() bool {
	r1 := CEF().SysCallN(643, m.Instance())
	return GoBool(r1)
}

func (m *TCefBrowserHost) CloseBrowser(forceClose bool) {
	CEF().SysCallN(612, m.Instance(), PascalBool(forceClose))
}

func (m *TCefBrowserHost) SetFocus(focus bool) {
	CEF().SysCallN(662, m.Instance(), PascalBool(focus))
}

func (m *TCefBrowserHost) Zoom(command TCefZoomCommand) {
	CEF().SysCallN(672, m.Instance(), uintptr(command))
}

func (m *TCefBrowserHost) SetZoomLevel(zoomLevel float64) {
	CEF().SysCallN(664, m.Instance(), uintptr(unsafePointer(&zoomLevel)))
}

func (m *TCefBrowserHost) RunFileDialog(mode TCefFileDialogMode, title, defaultFilePath string, acceptFilters IStrings, callback ICefRunFileDialogCallback) {
	CEF().SysCallN(650, m.Instance(), uintptr(mode), PascalStr(title), PascalStr(defaultFilePath), GetObjectUintptr(acceptFilters), GetObjectUintptr(callback))
}

func (m *TCefBrowserHost) StartDownload(url string) {
	CEF().SysCallN(666, m.Instance(), PascalStr(url))
}

func (m *TCefBrowserHost) DownloadImage(imageUrl string, isFavicon bool, maxImageSize uint32, bypassCache bool, callback ICefDownloadImageCallback) {
	CEF().SysCallN(614, m.Instance(), PascalStr(imageUrl), PascalBool(isFavicon), uintptr(maxImageSize), PascalBool(bypassCache), GetObjectUintptr(callback))
}

func (m *TCefBrowserHost) Print() {
	CEF().SysCallN(647, m.Instance())
}

func (m *TCefBrowserHost) PrintToPdf(path string, settings *TCefPdfPrintSettings, callback ICefPdfPrintCallback) {
	inArgs1 := settings.Pointer()
	CEF().SysCallN(648, m.Instance(), PascalStr(path), uintptr(unsafePointer(inArgs1)), GetObjectUintptr(callback))
}

func (m *TCefBrowserHost) Find(searchText string, forward, matchCase, findNext bool) {
	CEF().SysCallN(623, m.Instance(), PascalStr(searchText), PascalBool(forward), PascalBool(matchCase), PascalBool(findNext))
}

func (m *TCefBrowserHost) StopFinding(clearSelection bool) {
	CEF().SysCallN(667, m.Instance(), PascalBool(clearSelection))
}

func (m *TCefBrowserHost) ShowDevTools(windowInfo *TCefWindowInfo, client ICefClient, settings *TCefBrowserSettings, inspectElementAt *TCefPoint) {
	inArgs0 := windowInfo.Pointer()
	inArgs2 := settings.Pointer()
	CEF().SysCallN(665, m.Instance(), uintptr(unsafePointer(inArgs0)), GetObjectUintptr(client), uintptr(unsafePointer(inArgs2)), uintptr(unsafePointer(inspectElementAt)))
}

func (m *TCefBrowserHost) CloseDevTools() {
	CEF().SysCallN(613, m.Instance())
}

func (m *TCefBrowserHost) GetNavigationEntries(visitor ICefNavigationEntryVisitor, currentOnly bool) {
	CEF().SysCallN(628, m.Instance(), GetObjectUintptr(visitor), PascalBool(currentOnly))
}

func (m *TCefBrowserHost) ReplaceMisspelling(word string) {
	CEF().SysCallN(649, m.Instance(), PascalStr(word))
}

func (m *TCefBrowserHost) AddWordToDictionary(word string) {
	CEF().SysCallN(610, m.Instance(), PascalStr(word))
}

func (m *TCefBrowserHost) WasResized() {
	CEF().SysCallN(671, m.Instance())
}

func (m *TCefBrowserHost) NotifyScreenInfoChanged() {
	CEF().SysCallN(646, m.Instance())
}

func (m *TCefBrowserHost) WasHidden(hidden bool) {
	CEF().SysCallN(670, m.Instance(), PascalBool(hidden))
}

func (m *TCefBrowserHost) Invalidate(type_ TCefPaintElementType) {
	CEF().SysCallN(640, m.Instance(), uintptr(type_))
}

func (m *TCefBrowserHost) SendExternalBeginFrame() {
	CEF().SysCallN(653, m.Instance())
}

func (m *TCefBrowserHost) SendKeyEvent(event *TCefKeyEvent) {
	CEF().SysCallN(654, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TCefBrowserHost) SendMouseClickEvent(event *TCefMouseEvent, type_ TCefMouseButtonType, mouseUp bool, clickCount int32) {
	CEF().SysCallN(655, m.Instance(), uintptr(unsafePointer(event)), uintptr(type_), PascalBool(mouseUp), uintptr(clickCount))
}

func (m *TCefBrowserHost) SendMouseMoveEvent(event *TCefMouseEvent, mouseLeave bool) {
	CEF().SysCallN(656, m.Instance(), uintptr(unsafePointer(event)), PascalBool(mouseLeave))
}

func (m *TCefBrowserHost) SendMouseWheelEvent(event *TCefMouseEvent, deltaX, deltaY int32) {
	CEF().SysCallN(657, m.Instance(), uintptr(unsafePointer(event)), uintptr(deltaX), uintptr(deltaY))
}

func (m *TCefBrowserHost) SendTouchEvent(event *TCefTouchEvent) {
	CEF().SysCallN(658, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TCefBrowserHost) SendCaptureLostEvent() {
	CEF().SysCallN(651, m.Instance())
}

func (m *TCefBrowserHost) NotifyMoveOrResizeStarted() {
	CEF().SysCallN(645, m.Instance())
}

func (m *TCefBrowserHost) SetWindowlessFrameRate(frameRate int32) {
	CEF().SysCallN(663, m.Instance(), uintptr(frameRate))
}

func (m *TCefBrowserHost) IMECommitText(text string, replacementrange *TCefRange, relativecursorpos int32) {
	CEF().SysCallN(638, m.Instance(), PascalStr(text), uintptr(unsafePointer(replacementrange)), uintptr(relativecursorpos))
}

func (m *TCefBrowserHost) IMEFinishComposingText(keepselection bool) {
	CEF().SysCallN(639, m.Instance(), PascalBool(keepselection))
}

func (m *TCefBrowserHost) IMECancelComposition() {
	CEF().SysCallN(637, m.Instance())
}

func (m *TCefBrowserHost) DragTargetDragEnter(dragData ICefDragData, event *TCefMouseEvent, allowedOps TCefDragOperations) {
	CEF().SysCallN(617, m.Instance(), GetObjectUintptr(dragData), uintptr(unsafePointer(event)), uintptr(allowedOps))
}

func (m *TCefBrowserHost) DragTargetDragOver(event *TCefMouseEvent, allowedOps TCefDragOperations) {
	CEF().SysCallN(619, m.Instance(), uintptr(unsafePointer(event)), uintptr(allowedOps))
}

func (m *TCefBrowserHost) DragTargetDragLeave() {
	CEF().SysCallN(618, m.Instance())
}

func (m *TCefBrowserHost) DragTargetDrop(event *TCefMouseEvent) {
	CEF().SysCallN(620, m.Instance(), uintptr(unsafePointer(event)))
}

func (m *TCefBrowserHost) DragSourceEndedAt(x, y int32, op TCefDragOperation) {
	CEF().SysCallN(615, m.Instance(), uintptr(x), uintptr(y), uintptr(op))
}

func (m *TCefBrowserHost) DragSourceSystemDragEnded() {
	CEF().SysCallN(616, m.Instance())
}

func (m *TCefBrowserHost) SetAccessibilityState(accessibilityState TCefState) {
	CEF().SysCallN(659, m.Instance(), uintptr(accessibilityState))
}

func (m *TCefBrowserHost) SetAutoResizeEnabled(enabled bool, minsize, maxsize *TCefSize) {
	CEF().SysCallN(661, m.Instance(), PascalBool(enabled), uintptr(unsafePointer(minsize)), uintptr(unsafePointer(maxsize)))
}

func (m *TCefBrowserHost) SetAudioMuted(mute bool) {
	CEF().SysCallN(660, m.Instance(), PascalBool(mute))
}

func (m *TCefBrowserHost) ExitFullscreen(willcauseresize bool) {
	CEF().SysCallN(622, m.Instance(), PascalBool(willcauseresize))
}
