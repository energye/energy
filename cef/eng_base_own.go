//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// ICEFAccessibilityHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to receive accessibility notification when accessibility events have been registered. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_accessibility_handler_capi.h">CEF source file: /include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t))</a>
type ICEFAccessibilityHandler interface {
	ICefBaseRefCountedOwn
}

// TCEFAccessibilityHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to receive accessibility notification when accessibility events have been registered. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_accessibility_handler_capi.h">CEF source file: /include/capi/cef_accessibility_handler_capi.h (cef_accessibility_handler_t))</a>
type TCEFAccessibilityHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDevToolsMessageObserver Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.AddDevToolsMessageObserver. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h (cef_dev_tools_message_observer_t))</a>
type ICefDevToolsMessageObserver interface {
	ICefBaseRefCountedOwn
}

// TCefDevToolsMessageObserver Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.AddDevToolsMessageObserver. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_devtools_message_observer_capi.h">CEF source file: /include/capi/cef_devtools_message_observer_capi.h (cef_dev_tools_message_observer_t))</a>
type TCefDevToolsMessageObserver struct {
	TCefBaseRefCountedOwn
}

// ICEFServerHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle HTTP server requests. A new thread will be created for each ICefServer.CreateServer call (the "dedicated server thread"), and the functions of this interface will be called on that thread. It is therefore recommended to use a different ICefServerHandler instance for each ICefServer.CreateServer call to avoid thread safety issues in the ICefServerHandler implementation.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h (cef_server_handler_t))</a>
type ICEFServerHandler interface {
	ICefBaseRefCountedOwn
}

// TCEFServerHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle HTTP server requests. A new thread will be created for each ICefServer.CreateServer call (the "dedicated server thread"), and the functions of this interface will be called on that thread. It is therefore recommended to use a different ICefServerHandler instance for each ICefServer.CreateServer call to avoid thread safety issues in the ICefServerHandler implementation.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_server_capi.h">CEF source file: /include/capi/cef_server_capi.h (cef_server_handler_t))</a>
type TCEFServerHandler struct {
	TCefBaseRefCountedOwn
}

// ICefAudioHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle audio events.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h (cef_audio_handler_t))</a>
type ICefAudioHandler interface {
	ICefBaseRefCountedOwn
}

// TCefAudioHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle audio events.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_audio_handler_capi.h">CEF source file: /include/capi/cef_audio_handler_capi.h (cef_audio_handler_t))</a>
type TCefAudioHandler struct {
	TCefBaseRefCountedOwn
}

// ICefBaseRefCountedOwn Parent: IObject
type ICefBaseRefCountedOwn interface {
	IObject
}

// TCefBaseRefCountedOwn Parent: TObject
type TCefBaseRefCountedOwn struct {
	TObject
}

// ICefBrowserViewDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle BrowserView events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_browser_view_delegate_capi.h">CEF source file: /include/capi/views/cef_browser_view_delegate_capi.h (cef_browser_view_delegate_t)</a>
type ICefBrowserViewDelegate interface {
	ICefViewDelegate
}

// TCefBrowserViewDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle BrowserView events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_browser_view_delegate_capi.h">CEF source file: /include/capi/views/cef_browser_view_delegate_capi.h (cef_browser_view_delegate_t)</a>
type TCefBrowserViewDelegate struct {
	TCefViewDelegate
}

// ICefButtonDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle Button events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_delegate_capi.h">CEF source file: /include/capi/views/cef_button_delegate_capi.h (cef_button_delegate_t)</a>
type ICefButtonDelegate interface {
	ICefViewDelegate
}

// TCefButtonDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle Button events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_button_delegate_capi.h">CEF source file: /include/capi/views/cef_button_delegate_capi.h (cef_button_delegate_t)</a>
type TCefButtonDelegate struct {
	TCefViewDelegate
}

// ICefClient Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to provide handler implementations.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_client_capi.h">CEF source file: /include/capi/cef_client_capi.h (cef_client_t))</a>
type ICefClient interface {
	ICefBaseRefCountedOwn
}

// TCefClient Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to provide handler implementations.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_client_capi.h">CEF source file: /include/capi/cef_client_capi.h (cef_client_t))</a>
type TCefClient struct {
	TCefBaseRefCountedOwn
}

// ICefCommandHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to commands. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h (cef_command_handler_t))</a>
type ICefCommandHandler interface {
	ICefBaseRefCountedOwn
}

// TCefCommandHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to commands. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_command_handler_capi.h">CEF source file: /include/capi/cef_command_handler_capi.h (cef_command_handler_t))</a>
type TCefCommandHandler struct {
	TCefBaseRefCountedOwn
}

// ICefCompletionCallback Parent: ICefBaseRefCountedOwn
//
//	Generic callback interface used for asynchronous completion.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_callback_capi.h">CEF source file: /include/capi/cef_callback_capi.h (cef_completion_callback_t))</a>
type ICefCompletionCallback interface {
	ICefBaseRefCountedOwn
}

// TCefCompletionCallback Parent: TCefBaseRefCountedOwn
//
//	Generic callback interface used for asynchronous completion.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_callback_capi.h">CEF source file: /include/capi/cef_callback_capi.h (cef_completion_callback_t))</a>
type TCefCompletionCallback struct {
	TCefBaseRefCountedOwn
}

// ICefContextMenuHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle context menu events. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_context_menu_handler_t))</a>
type ICefContextMenuHandler interface {
	ICefBaseRefCountedOwn
}

// TCefContextMenuHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle context menu events. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_context_menu_handler_capi.h">CEF source file: /include/capi/cef_context_menu_handler_capi.h (cef_context_menu_handler_t))</a>
type TCefContextMenuHandler struct {
	TCefBaseRefCountedOwn
}

// ICefCookieVisitor Parent: ICefBaseRefCountedOwn
//
//	Interface to implement for visiting cookie values. The functions of this interface will always be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_cookie_visitor_t))</a>
type ICefCookieVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefCookieVisitor Parent: TCefBaseRefCountedOwn
//
//	Interface to implement for visiting cookie values. The functions of this interface will always be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_cookie_visitor_t))</a>
type TCefCookieVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefDeleteCookiesCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Interface to implement to be notified of asynchronous completion via ICefCookieManager.DeleteCookies.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_delete_cookies_callback_t))</a>
type ICefDeleteCookiesCallback interface {
	ICefBaseRefCountedOwn
}

// TCefDeleteCookiesCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Interface to implement to be notified of asynchronous completion via ICefCookieManager.DeleteCookies.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_delete_cookies_callback_t))</a>
type TCefDeleteCookiesCallback struct {
	TCefBaseRefCountedOwn
}

// ICefDialogHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle dialog events. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h (cef_dialog_handler_t))</a>
type ICefDialogHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDialogHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle dialog events. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dialog_handler_capi.h">CEF source file: /include/capi/cef_dialog_handler_capi.h (cef_dialog_handler_t))</a>
type TCefDialogHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDisplayHandler Parent: ICefBaseRefCountedOwn
//
//	Event handler related to browser display state.
//	The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h (cef_display_handler_t)</a>
type ICefDisplayHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDisplayHandler Parent: TCefBaseRefCountedOwn
//
//	Event handler related to browser display state.
//	The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_display_handler_capi.h">CEF source file: /include/capi/cef_display_handler_capi.h (cef_display_handler_t)</a>
type TCefDisplayHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDomVisitor Parent: ICefBaseRefCountedOwn
//
//	Interface to implement for visiting the DOM. The functions of this interface will be called on the render process main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domvisitor_t))</a>
type ICefDomVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefDomVisitor Parent: TCefBaseRefCountedOwn
//
//	Interface to implement for visiting the DOM. The functions of this interface will be called on the render process main thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domvisitor_t))</a>
type TCefDomVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefDownloadHandler Parent: ICefBaseRefCountedOwn
//
//	Interface used to handle file downloads. The functions of this interface will called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_download_handler_t))</a>
type ICefDownloadHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDownloadHandler Parent: TCefBaseRefCountedOwn
//
//	Interface used to handle file downloads. The functions of this interface will called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_download_handler_capi.h">CEF source file: /include/capi/cef_download_handler_capi.h (cef_download_handler_t))</a>
type TCefDownloadHandler struct {
	TCefBaseRefCountedOwn
}

// ICefDownloadImageCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.DownloadImage. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_download_image_callback_t))</a>
type ICefDownloadImageCallback interface {
	ICefBaseRefCountedOwn
}

// TCefDownloadImageCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.DownloadImage. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_download_image_callback_t))</a>
type TCefDownloadImageCallback struct {
	TCefBaseRefCountedOwn
}

// ICefDragHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to dragging. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_handler_capi.h">CEF source file: /include/capi/cef_drag_handler_capi.h (cef_drag_handler_t))</a>
type ICefDragHandler interface {
	ICefBaseRefCountedOwn
}

// TCefDragHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to dragging. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_handler_capi.h">CEF source file: /include/capi/cef_drag_handler_capi.h (cef_drag_handler_t))</a>
type TCefDragHandler struct {
	TCefBaseRefCountedOwn
}

// ICefExtensionHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser extensions. The functions of this interface will be called on the UI thread. See ICefRequestContext.LoadExtension for information about extension loading.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h (cef_extension_handler_t))</a>
type ICefExtensionHandler interface {
	ICefBaseRefCountedOwn
}

// TCefExtensionHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser extensions. The functions of this interface will be called on the UI thread. See ICefRequestContext.LoadExtension for information about extension loading.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_extension_handler_capi.h">CEF source file: /include/capi/cef_extension_handler_capi.h (cef_extension_handler_t))</a>
type TCefExtensionHandler struct {
	TCefBaseRefCountedOwn
}

// ICefFindHandler Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to find results. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_find_handler_capi.h">CEF source file: /include/capi/cef_find_handler_capi.h (cef_find_handler_t))</a>
type ICefFindHandler interface {
	ICefBaseRefCountedOwn
}

// TCefFindHandler Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to find results. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_find_handler_capi.h">CEF source file: /include/capi/cef_find_handler_capi.h (cef_find_handler_t))</a>
type TCefFindHandler struct {
	TCefBaseRefCountedOwn
}

// ICefFocusHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to focus. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_focus_handler_capi.h">CEF source file: /include/capi/cef_focus_handler_capi.h (cef_focus_handler_t))</a>
type ICefFocusHandler interface {
	ICefBaseRefCountedOwn
}

// TCefFocusHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to focus. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_focus_handler_capi.h">CEF source file: /include/capi/cef_focus_handler_capi.h (cef_focus_handler_t))</a>
type TCefFocusHandler struct {
	TCefBaseRefCountedOwn
}

// ICefFrameHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to ICefFrame life span. The order of callbacks is: (1) During initial ICefBrowserHost creation and navigation of the main frame: - ICefFrameHandler.OnFrameCreated => The initial main frame object has been created. Any commands will be queued until the frame is attached. - ICefFrameHandler.OnMainFrameChanged => The initial main frame object has been assigned to the browser. - ICefLifeSpanHandler.OnAfterCreated => The browser is now valid and can be used. - ICefFrameHandler.OnFrameAttached => The initial main frame object is now connected to its peer in the renderer process. Commands can be routed. (2) During further ICefBrowserHost navigation/loading of the main frame and/or sub-frames: - ICefFrameHandler.OnFrameCreated => A new main frame or sub-frame object has been created. Any commands will be queued until the frame is attached. - ICefFrameHandler.OnFrameAttached => A new main frame or sub-frame object is now connected to its peer in the renderer process. Commands can be routed. - ICefFrameHandler.OnFrameDetached => An existing main frame or sub- frame object has lost its connection to the renderer process. If multiple objects are detached at the same time then notifications will be sent for any sub-frame objects before the main frame object. Commands can no longer be routed and will be discarded. - ICefFrameHandler.OnMainFrameChanged => A new main frame object has been assigned to the browser. This will only occur with cross-origin navigation or re-navigation after renderer process termination (due to crashes, etc). (3) During final ICefBrowserHost destruction of the main frame: - ICefFrameHandler.OnFrameDetached => Any sub-frame objects have lost their connection to the renderer process. Commands can no longer be routed and will be discarded. - ICefLifeSpanHandler.OnBeforeClose => The browser has been destroyed. - ICefFrameHandler.OnFrameDetached => The main frame object have lost its connection to the renderer process. Notifications will be sent for any sub-frame objects before the main frame object. Commands can no longer be routed and will be discarded. - ICefFrameHandler.OnMainFrameChanged => The final main frame object has been removed from the browser. Cross-origin navigation and/or loading receives special handling. When the main frame navigates to a different origin the OnMainFrameChanged callback (2) will be executed with the old and new main frame objects. When a new sub-frame is loaded in, or an existing sub-frame is navigated to, a different origin from the parent frame, a temporary sub-frame object will first be created in the parent's renderer process. That temporary sub-frame will then be discarded after the real cross-origin sub-frame is created in the new/target renderer process. The client will receive cross-origin navigation callbacks (2) for the transition from the temporary sub-frame to the real sub-frame. The temporary sub-frame will not recieve or execute commands during this transitional period (any sent commands will be discarded). When a new popup browser is created in a different origin from the parent browser, a temporary main frame object for the popup will first be created in the parent's renderer process. That temporary main frame will then be discarded after the real cross-origin main frame is created in the new/target renderer process. The client will recieve creation and initial navigation callbacks (1) for the temporary main frame, followed by cross- origin navigation callbacks (2) for the transition from the temporary main frame to the real main frame. The temporary main frame may receive and execute commands during this transitional period (any sent commands may be executed, but the behavior is potentially undesirable since they execute in the parent browser's renderer process and not the new/target renderer process). Callbacks will not be executed for placeholders that may be created during pre-commit navigation for sub-frames that do not yet exist in the renderer process. Placeholders will have ICefFrame.GetIdentifier() == -4. The functions of this interface will be called on the UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_handler_capi.h">CEF source file: /include/capi/cef_frame_handler_capi.h (cef_frame_handler_t))</a>
type ICefFrameHandler interface {
	ICefBaseRefCountedOwn
}

// TCefFrameHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to ICefFrame life span. The order of callbacks is: (1) During initial ICefBrowserHost creation and navigation of the main frame: - ICefFrameHandler.OnFrameCreated => The initial main frame object has been created. Any commands will be queued until the frame is attached. - ICefFrameHandler.OnMainFrameChanged => The initial main frame object has been assigned to the browser. - ICefLifeSpanHandler.OnAfterCreated => The browser is now valid and can be used. - ICefFrameHandler.OnFrameAttached => The initial main frame object is now connected to its peer in the renderer process. Commands can be routed. (2) During further ICefBrowserHost navigation/loading of the main frame and/or sub-frames: - ICefFrameHandler.OnFrameCreated => A new main frame or sub-frame object has been created. Any commands will be queued until the frame is attached. - ICefFrameHandler.OnFrameAttached => A new main frame or sub-frame object is now connected to its peer in the renderer process. Commands can be routed. - ICefFrameHandler.OnFrameDetached => An existing main frame or sub- frame object has lost its connection to the renderer process. If multiple objects are detached at the same time then notifications will be sent for any sub-frame objects before the main frame object. Commands can no longer be routed and will be discarded. - ICefFrameHandler.OnMainFrameChanged => A new main frame object has been assigned to the browser. This will only occur with cross-origin navigation or re-navigation after renderer process termination (due to crashes, etc). (3) During final ICefBrowserHost destruction of the main frame: - ICefFrameHandler.OnFrameDetached => Any sub-frame objects have lost their connection to the renderer process. Commands can no longer be routed and will be discarded. - ICefLifeSpanHandler.OnBeforeClose => The browser has been destroyed. - ICefFrameHandler.OnFrameDetached => The main frame object have lost its connection to the renderer process. Notifications will be sent for any sub-frame objects before the main frame object. Commands can no longer be routed and will be discarded. - ICefFrameHandler.OnMainFrameChanged => The final main frame object has been removed from the browser. Cross-origin navigation and/or loading receives special handling. When the main frame navigates to a different origin the OnMainFrameChanged callback (2) will be executed with the old and new main frame objects. When a new sub-frame is loaded in, or an existing sub-frame is navigated to, a different origin from the parent frame, a temporary sub-frame object will first be created in the parent's renderer process. That temporary sub-frame will then be discarded after the real cross-origin sub-frame is created in the new/target renderer process. The client will receive cross-origin navigation callbacks (2) for the transition from the temporary sub-frame to the real sub-frame. The temporary sub-frame will not recieve or execute commands during this transitional period (any sent commands will be discarded). When a new popup browser is created in a different origin from the parent browser, a temporary main frame object for the popup will first be created in the parent's renderer process. That temporary main frame will then be discarded after the real cross-origin main frame is created in the new/target renderer process. The client will recieve creation and initial navigation callbacks (1) for the temporary main frame, followed by cross- origin navigation callbacks (2) for the transition from the temporary main frame to the real main frame. The temporary main frame may receive and execute commands during this transitional period (any sent commands may be executed, but the behavior is potentially undesirable since they execute in the parent browser's renderer process and not the new/target renderer process). Callbacks will not be executed for placeholders that may be created during pre-commit navigation for sub-frames that do not yet exist in the renderer process. Placeholders will have ICefFrame.GetIdentifier() == -4. The functions of this interface will be called on the UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_frame_handler_capi.h">CEF source file: /include/capi/cef_frame_handler_capi.h (cef_frame_handler_t))</a>
type TCefFrameHandler struct {
	TCefBaseRefCountedOwn
}

// ICefJsDialogHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to JavaScript dialogs. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h (cef_jsdialog_handler_t))</a>
type ICefJsDialogHandler interface {
	ICefBaseRefCountedOwn
}

// TCefJsDialogHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to JavaScript dialogs. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_jsdialog_handler_capi.h">CEF source file: /include/capi/cef_jsdialog_handler_capi.h (cef_jsdialog_handler_t))</a>
type TCefJsDialogHandler struct {
	TCefBaseRefCountedOwn
}

// ICefKeyboardHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to keyboard input. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_keyboard_handler_capi.h">CEF source file: /include/capi/cef_keyboard_handler_capi.h (cef_keyboard_handler_t))</a>
type ICefKeyboardHandler interface {
	ICefBaseRefCountedOwn
}

// TCefKeyboardHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to keyboard input. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_keyboard_handler_capi.h">CEF source file: /include/capi/cef_keyboard_handler_capi.h (cef_keyboard_handler_t))</a>
type TCefKeyboardHandler struct {
	TCefBaseRefCountedOwn
}

// ICefLifeSpanHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser life span. The functions of this interface will be called on the UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_life_span_handler_capi.h">CEF source file: /include/capi/cef_life_span_handler_capi.h (cef_life_span_handler_t))</a>
type ICefLifeSpanHandler interface {
	ICefBaseRefCountedOwn
}

// TCefLifeSpanHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser life span. The functions of this interface will be called on the UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_life_span_handler_capi.h">CEF source file: /include/capi/cef_life_span_handler_capi.h (cef_life_span_handler_t))</a>
type TCefLifeSpanHandler struct {
	TCefBaseRefCountedOwn
}

// ICefLoadHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser load status. The functions of this interface will be called on the browser process UI thread or render process main thread (TID_RENDERER).
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h (cef_load_handler_t))</a>
type ICefLoadHandler interface {
	ICefBaseRefCountedOwn
}

// TCefLoadHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser load status. The functions of this interface will be called on the browser process UI thread or render process main thread (TID_RENDERER).
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_load_handler_capi.h">CEF source file: /include/capi/cef_load_handler_capi.h (cef_load_handler_t))</a>
type TCefLoadHandler struct {
	TCefBaseRefCountedOwn
}

// ICefMediaObserver Parent: ICefBaseRefCountedOwn
//
//	Implemented by the client to observe MediaRouter events and registered via ICefMediaRouter.AddObserver. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_observer_t))</a>
type ICefMediaObserver interface {
	ICefBaseRefCountedOwn
}

// TCefMediaObserver Parent: TCefBaseRefCountedOwn
//
//	Implemented by the client to observe MediaRouter events and registered via ICefMediaRouter.AddObserver. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_observer_t))</a>
type TCefMediaObserver struct {
	TCefBaseRefCountedOwn
}

// ICefMediaRouteCreateCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefMediaRouter.CreateRoute. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_create_callback_t))</a>
type ICefMediaRouteCreateCallback interface {
	ICefBaseRefCountedOwn
}

// TCefMediaRouteCreateCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefMediaRouter.CreateRoute. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_route_create_callback_t))</a>
type TCefMediaRouteCreateCallback struct {
	TCefBaseRefCountedOwn
}

// ICefMediaSinkDeviceInfoCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefMediaSink.GetDeviceInfo. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_device_info_callback_t))</a>
type ICefMediaSinkDeviceInfoCallback interface {
	ICefBaseRefCountedOwn
}

// TCefMediaSinkDeviceInfoCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefMediaSink.GetDeviceInfo. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_media_router_capi.h">CEF source file: /include/capi/cef_media_router_capi.h (cef_media_sink_device_info_callback_t))</a>
type TCefMediaSinkDeviceInfoCallback struct {
	TCefBaseRefCountedOwn
}

// ICefMenuButtonDelegate Parent: ICefButtonDelegate
//
//	Implement this interface to handle MenuButton events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_delegate_t)</a>
type ICefMenuButtonDelegate interface {
	ICefButtonDelegate
}

// TCefMenuButtonDelegate Parent: TCefButtonDelegate
//
//	Implement this interface to handle MenuButton events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_menu_button_delegate_capi.h">CEF source file: /include/capi/views/cef_menu_button_delegate_capi.h (cef_menu_button_delegate_t)</a>
type TCefMenuButtonDelegate struct {
	TCefButtonDelegate
}

// ICefMenuModelDelegate Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle menu model events. The functions of this interface will be called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_delegate_capi.h">CEF source file: /include/capi/cef_menu_model_delegate_capi.h (cef_menu_model_delegate_t))</a>
type ICefMenuModelDelegate interface {
	ICefBaseRefCountedOwn
}

// TCefMenuModelDelegate Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle menu model events. The functions of this interface will be called on the browser process UI thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_delegate_capi.h">CEF source file: /include/capi/cef_menu_model_delegate_capi.h (cef_menu_model_delegate_t))</a>
type TCefMenuModelDelegate struct {
	TCefBaseRefCountedOwn
}

// ICefNavigationEntryVisitor Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.GetNavigationEntries. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_navigation_entry_visitor_t))</a>
type ICefNavigationEntryVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefNavigationEntryVisitor Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.GetNavigationEntries. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_navigation_entry_visitor_t))</a>
type TCefNavigationEntryVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefPanelDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle Panel events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_panel_delegate_capi.h">CEF source file: /include/capi/views/cef_panel_delegate_capi.h (cef_panel_delegate_t)</a>
type ICefPanelDelegate interface {
	ICefViewDelegate
}

// TCefPanelDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle Panel events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_panel_delegate_capi.h">CEF source file: /include/capi/views/cef_panel_delegate_capi.h (cef_panel_delegate_t)</a>
type TCefPanelDelegate struct {
	TCefViewDelegate
}

// ICefPdfPrintCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.PrintToPDF. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_pdf_print_callback_t))</a>
type ICefPdfPrintCallback interface {
	ICefBaseRefCountedOwn
}

// TCefPdfPrintCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.PrintToPDF. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_pdf_print_callback_t))</a>
type TCefPdfPrintCallback struct {
	TCefBaseRefCountedOwn
}

// ICefPermissionHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to permission requests. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_permission_handler_t))</a>
type ICefPermissionHandler interface {
	ICefBaseRefCountedOwn
}

// TCefPermissionHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to permission requests. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_permission_handler_capi.h">CEF source file: /include/capi/cef_permission_handler_capi.h (cef_permission_handler_t))</a>
type TCefPermissionHandler struct {
	TCefBaseRefCountedOwn
}

// ICefPrintHandler Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle printing on Linux. Each browser will have only one print job in progress at a time. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_handler_t))</a>
type ICefPrintHandler interface {
	ICefBaseRefCountedOwn
}

// TCefPrintHandler Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle printing on Linux. Each browser will have only one print job in progress at a time. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_print_handler_capi.h">CEF source file: /include/capi/cef_print_handler_capi.h (cef_print_handler_t))</a>
type TCefPrintHandler struct {
	TCefBaseRefCountedOwn
}

// ICefRenderHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events when window rendering is disabled. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h (cef_render_handler_t))</a>
type ICefRenderHandler interface {
	ICefBaseRefCountedOwn
}

// TCefRenderHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events when window rendering is disabled. The functions of this interface will be called on the UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_handler_capi.h">CEF source file: /include/capi/cef_render_handler_capi.h (cef_render_handler_t))</a>
type TCefRenderHandler struct {
	TCefBaseRefCountedOwn
}

// ICefRenderProcessHandler Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Interface used to implement render process callbacks. The functions of this interface will be called on the render process main thread (TID_RENDERER) unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h (cef_render_process_handler_t))</a>
type ICefRenderProcessHandler interface {
	ICefBaseRefCountedOwn
}

// TCefRenderProcessHandler Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Interface used to implement render process callbacks. The functions of this interface will be called on the render process main thread (TID_RENDERER) unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_render_process_handler_capi.h">CEF source file: /include/capi/cef_render_process_handler_capi.h (cef_render_process_handler_t))</a>
type TCefRenderProcessHandler struct {
	TCefBaseRefCountedOwn
}

// ICefRequestHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser requests. The functions of this interface will be called on the thread indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h (cef_request_handler_t))</a>
type ICefRequestHandler interface {
	ICefBaseRefCountedOwn
}

// TCefRequestHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser requests. The functions of this interface will be called on the thread indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_handler_capi.h">CEF source file: /include/capi/cef_request_handler_capi.h (cef_request_handler_t))</a>
type TCefRequestHandler struct {
	TCefBaseRefCountedOwn
}

// ICefResolveCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefRequestContext.ResolveHost.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_context_capi.h">CEF source file: /include/capi/cef_request_context_capi.h (cef_resolve_callback_t))</a>
type ICefResolveCallback interface {
	ICefBaseRefCountedOwn
}

// TCefResolveCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefRequestContext.ResolveHost.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_request_context_capi.h">CEF source file: /include/capi/cef_request_context_capi.h (cef_resolve_callback_t))</a>
type TCefResolveCallback struct {
	TCefBaseRefCountedOwn
}

// ICefResourceHandler Parent: ICefBaseRefCountedOwn
//
//	Interface used to implement a custom request handler interface. The functions of this interface will be called on the IO thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_handler_capi.h">CEF source file: /include/capi/cef_resource_handler_capi.h (cef_resource_handler_t))</a>
type ICefResourceHandler interface {
	ICefBaseRefCountedOwn
}

// TCefResourceHandler Parent: TCefBaseRefCountedOwn
//
//	Interface used to implement a custom request handler interface. The functions of this interface will be called on the IO thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_handler_capi.h">CEF source file: /include/capi/cef_resource_handler_capi.h (cef_resource_handler_t))</a>
type TCefResourceHandler struct {
	TCefBaseRefCountedOwn
}

// ICefResourceRequestHandler Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser requests. The functions of this interface will be called on the IO thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h (cef_resource_request_handler_t))</a>
type ICefResourceRequestHandler interface {
	ICefBaseRefCountedOwn
}

// TCefResourceRequestHandler Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle events related to browser requests. The functions of this interface will be called on the IO thread unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_resource_request_handler_capi.h">CEF source file: /include/capi/cef_resource_request_handler_capi.h (cef_resource_request_handler_t))</a>
type TCefResourceRequestHandler struct {
	TCefBaseRefCountedOwn
}

// ICefResponseFilter Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to filter resource response content. The functions of this interface will be called on the browser process IO thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_response_filter_capi.h">CEF source file: /include/capi/cef_response_filter_capi.h (cef_response_filter_t))</a>
type ICefResponseFilter interface {
	ICefBaseRefCountedOwn
}

// TCefResponseFilter Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to filter resource response content. The functions of this interface will be called on the browser process IO thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_response_filter_capi.h">CEF source file: /include/capi/cef_response_filter_capi.h (cef_response_filter_t))</a>
type TCefResponseFilter struct {
	TCefBaseRefCountedOwn
}

// ICefRunFileDialogCallback Parent: ICefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.RunFileDialog. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t))</a>
type ICefRunFileDialogCallback interface {
	ICefBaseRefCountedOwn
}

// TCefRunFileDialogCallback Parent: TCefBaseRefCountedOwn
//
//	Callback interface for ICefBrowserHost.RunFileDialog. The functions of this interface will be called on the browser process UI thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_browser_capi.h">CEF source file: /include/capi/cef_browser_capi.h (cef_run_file_dialog_callback_t))</a>
type TCefRunFileDialogCallback struct {
	TCefBaseRefCountedOwn
}

// ICefSchemeHandlerFactory Parent: ICefBaseRefCountedOwn
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type ICefSchemeHandlerFactory interface {
	ICefBaseRefCountedOwn
}

// TCefSchemeHandlerFactory Parent: TCefBaseRefCountedOwn
//
//	Class that creates ICefResourceHandler instances for handling scheme
//	requests.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_scheme_capi.h">CEF source file: /include/capi/cef_scheme_capi.h (cef_scheme_handler_factory_t)</a>
type TCefSchemeHandlerFactory struct {
	TCefBaseRefCountedOwn
}

// ICefSetCookieCallback Is Abstract Class Parent: ICefBaseRefCountedOwn
//
//	Interface to implement to be notified of asynchronous completion via ICefCookieManager.SetCookie.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_set_cookie_callback_t))</a>
type ICefSetCookieCallback interface {
	ICefBaseRefCountedOwn
}

// TCefSetCookieCallback Is Abstract Class Parent: TCefBaseRefCountedOwn
//
//	Interface to implement to be notified of asynchronous completion via ICefCookieManager.SetCookie.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_cookie_capi.h">CEF source file: /include/capi/cef_cookie_capi.h (cef_set_cookie_callback_t))</a>
type TCefSetCookieCallback struct {
	TCefBaseRefCountedOwn
}

// ICefStringVisitor Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to receive string values asynchronously.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_string_visitor_capi.h">CEF source file: /include/capi/cef_string_visitor_capi.h (cef_string_visitor_t))</a>
type ICefStringVisitor interface {
	ICefBaseRefCountedOwn
}

// TCefStringVisitor Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to receive string values asynchronously.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_string_visitor_capi.h">CEF source file: /include/capi/cef_string_visitor_capi.h (cef_string_visitor_t))</a>
type TCefStringVisitor struct {
	TCefBaseRefCountedOwn
}

// ICefTask Parent: ICefBaseRefCountedOwn
//
//	Implement this interface for asynchronous task execution. If the task is posted successfully and if the associated message loop is still running then the execute() function will be called on the target thread. If the task fails to post then the task object may be destroyed on the source thread instead of the target thread. For this reason be cautious when performing work in the task object destructor.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_t))</a>
type ICefTask interface {
	ICefBaseRefCountedOwn
}

// TCefTask Parent: TCefBaseRefCountedOwn
//
//	Implement this interface for asynchronous task execution. If the task is posted successfully and if the associated message loop is still running then the execute() function will be called on the target thread. If the task fails to post then the task object may be destroyed on the source thread instead of the target thread. For this reason be cautious when performing work in the task object destructor.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_task_capi.h">CEF source file: /include/capi/cef_task_capi.h (cef_task_t))</a>
type TCefTask struct {
	TCefBaseRefCountedOwn
}

// ICefTextfieldDelegate Parent: ICefViewDelegate
//
//	Implement this interface to handle Textfield events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_textfield_delegate_capi.h">CEF source file: /include/capi/views/cef_textfield_delegate_capi.h (cef_textfield_delegate_t)</a>
type ICefTextfieldDelegate interface {
	ICefViewDelegate
}

// TCefTextfieldDelegate Parent: TCefViewDelegate
//
//	Implement this interface to handle Textfield events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_textfield_delegate_capi.h">CEF source file: /include/capi/views/cef_textfield_delegate_capi.h (cef_textfield_delegate_t)</a>
type TCefTextfieldDelegate struct {
	TCefViewDelegate
}

// ICefUrlRequestClient Parent: ICefBaseRefCountedOwn
//
//	Interface that should be implemented by the ICefUrlRequest client. The functions of this interface will be called on the same thread that created the request unless otherwise documented.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h (cef_urlrequest_client_t))</a>
type ICefUrlRequestClient interface {
	ICefBaseRefCountedOwn
}

// TCefUrlRequestClient Parent: TCefBaseRefCountedOwn
//
//	Interface that should be implemented by the ICefUrlRequest client. The functions of this interface will be called on the same thread that created the request unless otherwise documented.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_urlrequest_capi.h">CEF source file: /include/capi/cef_urlrequest_capi.h (cef_urlrequest_client_t))</a>
type TCefUrlRequestClient struct {
	TCefBaseRefCountedOwn
}

// ICefV8Accessor Parent: ICefBaseRefCountedOwn
//
//	Interface that should be implemented to handle V8 accessor calls. Accessor identifiers are registered by calling ICefV8value.SetValue(). The functions of this interface will be called on the thread associated with the V8 accessor.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8accessor_t))</a>
type ICefV8Accessor interface {
	ICefBaseRefCountedOwn
}

// TCefV8Accessor Parent: TCefBaseRefCountedOwn
//
//	Interface that should be implemented to handle V8 accessor calls. Accessor identifiers are registered by calling ICefV8value.SetValue(). The functions of this interface will be called on the thread associated with the V8 accessor.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8accessor_t))</a>
type TCefV8Accessor struct {
	TCefBaseRefCountedOwn
}

// ICefV8Interceptor Parent: ICefBaseRefCountedOwn
//
//	Interface that should be implemented to handle V8 interceptor calls. The functions of this interface will be called on the thread associated with the V8 interceptor. Interceptor's named property handlers (with first argument of type CefString) are called when object is indexed by string. Indexed property handlers (with first argument of type int) are called when object is indexed by integer.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8interceptor_t))</a>
type ICefV8Interceptor interface {
	ICefBaseRefCountedOwn
}

// TCefV8Interceptor Parent: TCefBaseRefCountedOwn
//
//	Interface that should be implemented to handle V8 interceptor calls. The functions of this interface will be called on the thread associated with the V8 interceptor. Interceptor's named property handlers (with first argument of type CefString) are called when object is indexed by string. Indexed property handlers (with first argument of type int) are called when object is indexed by integer.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8interceptor_t))</a>
type TCefV8Interceptor struct {
	TCefBaseRefCountedOwn
}

// ICefViewDelegate Parent: ICefBaseRefCountedOwn
//
//	Implement this interface to handle view events. All size and position values
//	are in density independent pixels (DIP) unless otherwise indicated. The
//	functions of this interface will be called on the browser process UI thread
//	unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_delegate_capi.h">CEF source file: /include/capi/views/cef_view_delegate_capi.h (cef_view_delegate_t)</a>
type ICefViewDelegate interface {
	ICefBaseRefCountedOwn
}

// TCefViewDelegate Parent: TCefBaseRefCountedOwn
//
//	Implement this interface to handle view events. All size and position values
//	are in density independent pixels (DIP) unless otherwise indicated. The
//	functions of this interface will be called on the browser process UI thread
//	unless otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_view_delegate_capi.h">CEF source file: /include/capi/views/cef_view_delegate_capi.h (cef_view_delegate_t)</a>
type TCefViewDelegate struct {
	TCefBaseRefCountedOwn
}

// ICefWindowDelegate Parent: ICefPanelDelegate
//
//	Implement this interface to handle window events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_window_delegate_capi.h">CEF source file: /include/capi/views/cef_window_delegate_capi.h (cef_window_delegate_t)</a>
type ICefWindowDelegate interface {
	ICefPanelDelegate
}

// TCefWindowDelegate Parent: TCefPanelDelegate
//
//	Implement this interface to handle window events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise
//	indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_window_delegate_capi.h">CEF source file: /include/capi/views/cef_window_delegate_capi.h (cef_window_delegate_t)</a>
type TCefWindowDelegate struct {
	TCefPanelDelegate
}

// ICefWriteHandler Parent: ICefBaseRefCountedOwn
//
//	Interface the client can implement to provide a custom stream writer. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_write_handler_t))</a>
type ICefWriteHandler interface {
	ICefBaseRefCountedOwn
}

// TCefWriteHandler Parent: TCefBaseRefCountedOwn
//
//	Interface the client can implement to provide a custom stream writer. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_stream_capi.h">CEF source file: /include/capi/cef_stream_capi.h (cef_write_handler_t))</a>
type TCefWriteHandler struct {
	TCefBaseRefCountedOwn
}

// ICefv8ArrayBufferReleaseCallback Parent: ICefBaseRefCountedOwn
//
//	Callback interface that is passed to ICefV8value.CreateArrayBuffer.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8array_buffer_release_callback_t))</a>
type ICefv8ArrayBufferReleaseCallback interface {
	ICefBaseRefCountedOwn
}

// TCefv8ArrayBufferReleaseCallback Parent: TCefBaseRefCountedOwn
//
//	Callback interface that is passed to ICefV8value.CreateArrayBuffer.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8array_buffer_release_callback_t))</a>
type TCefv8ArrayBufferReleaseCallback struct {
	TCefBaseRefCountedOwn
}

// ICefv8Handler Parent: ICefBaseRefCountedOwn
//
//	Interface that should be implemented to handle V8 function calls. The functions of this interface will be called on the thread associated with the V8 function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8handler_t))</a>
type ICefv8Handler interface {
	ICefBaseRefCountedOwn
}

// TCefv8Handler Parent: TCefBaseRefCountedOwn
//
//	Interface that should be implemented to handle V8 function calls. The functions of this interface will be called on the thread associated with the V8 function.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_v8_capi.h">CEF source file: /include/capi/cef_v8_capi.h (cef_v8handler_t))</a>
type TCefv8Handler struct {
	TCefBaseRefCountedOwn
}
