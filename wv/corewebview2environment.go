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

// ICoreWebView2Environment Parent: IObject
//
//	Represents the WebView2 Environment.  WebViews created from an environment
//	run on the browser process specified with environment parameters and
//	objects created from an environment should be used in the same
//	environment.  Using it in different environments are not guaranteed to be
//	 compatible and may fail.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment">See the ICoreWebView2Environment article.</a>
type ICoreWebView2Environment interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Environment // property
	// BrowserVersionInfo
	//  The browser version info of the current `ICoreWebView2Environment`,
	//  including channel name if it is not the WebView2 Runtime. It matches the
	//  format of the `GetAvailableCoreWebView2BrowserVersionString` API.
	//  Channel names are `beta`, `dev`, and `canary`.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#get_browserversionstring">See the ICoreWebView2Environment article.</a>
	BrowserVersionInfo() string // property
	// SupportsCompositionController
	//  Returns true if the current WebView2 runtime version supports Composition Controllers.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3">See the ICoreWebView2Environment3 article.</a>
	SupportsCompositionController() bool // property
	// SupportsControllerOptions
	//  Returns true if the current WebView2 runtime version supports Controller Options.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10">See the ICoreWebView2Environment10 article.</a>
	SupportsControllerOptions() bool // property
	// UserDataFolder
	//  Returns the user data folder that all CoreWebView2's created from this
	//  environment are using.
	//  This could be either the value passed in by the developer when creating
	//  the environment object or the calculated one for default handling. It
	//  will always be an absolute path.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment7#get_userdatafolder">See the ICoreWebView2Environment7 article.</a>
	UserDataFolder() string // property
	// ProcessInfos
	//  Returns the `ICoreWebView2ProcessInfoCollection`. Provide a list of all
	//  process using same user data folder except for crashpad process.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment8#getprocessinfos">See the ICoreWebView2Environment8 article.</a>
	ProcessInfos() ICoreWebView2ProcessInfoCollection // property
	// FailureReportFolderPath
	//  `FailureReportFolderPath` returns the path of the folder where minidump files are written.
	//  Whenever a WebView2 process crashes, a crash dump file will be created in the crash dump folder.
	//  The crash dump format is minidump files. Please see
	//  [Minidump Files documentation](/windows/win32/debug/minidump-files) for detailed information.
	//  Normally when a single child process fails, a minidump will be generated and written to disk,
	//  then the `ProcessFailed` event is raised. But for unexpected crashes, a minidump file might not be generated
	//  at all, despite whether `ProcessFailed` event is raised. If there are multiple
	//  process failures at once, multiple minidump files could be generated. Thus `FailureReportFolderPath`
	//  could contain old minidump files that are not associated with a specific `ProcessFailed` event.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment11#get_failurereportfolderpath">See the ICoreWebView2Environment11 article.</a>
	FailureReportFolderPath() string // property
	// AddAllLoaderEvents
	//  Adds all the events of this class to an existing TWVLoader instance.
	//  <param name="aLoaderComponent">The TWVLoader instance.</param>
	AddAllLoaderEvents(aLoaderComponent IComponent) bool // function
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// CreateCoreWebView2Controller
	//  Asynchronously create a new WebView.
	//  <param name="aParentWindow">Handle of the control in which the WebView should be displayed.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#createcorewebview2controller">See the ICoreWebView2Environment article.</a>
	CreateCoreWebView2Controller(aParentWindow THandle, aBrowserEvents IWVBrowserEvents, aResult *int32) bool // function
	// CreateWebResourceResponse
	//  Create a new web resource response object.
	//  <param name="aContent">HTTP response content as stream.</param>
	//  <param name="aStatusCode">The HTTP response status code.</param>
	//  <param name="aReasonPhrase">The HTTP response reason phrase.</param>
	//  <param name="aHeaders">Overridden HTTP response headers.</param>
	//  <param name="aResponse">The new ICoreWebView2WebResourceResponse instance.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment#createcorewebview2controller">See the ICoreWebView2Environment article.</a>
	CreateWebResourceResponse(aContent IStream, aStatusCode int32, aReasonPhrase, aHeaders string, aResponse *ICoreWebView2WebResourceResponse) bool // function
	// CreateWebResourceRequest
	//  Create a new web resource request object.
	//  URI parameter must be absolute URI.
	//  The headers string is the raw request header string delimited by CRLF
	// (optional in last header).
	//  It's also possible to create this object with null headers string
	//  and then use the ICoreWebView2HttpRequestHeaders to construct the headers
	//  line by line.
	//  <param name="aURI">The request URI.</param>
	//  <param name="aMethod">The HTTP request method.</param>
	//  <param name="aPostData">The HTTP request message body as stream.</param>
	//  <param name="aHeaders">The mutable HTTP request headers.</param>
	//  <param name="aRequest">The new ICoreWebView2WebResourceRequest instance.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment2#createwebresourcerequest">See the ICoreWebView2Environment2 article.</a>
	CreateWebResourceRequest(aURI, aMethod string, aPostData IStream, aHeaders string, aRequest *ICoreWebView2WebResourceRequestRef) bool // function
	// CreateCoreWebView2CompositionController
	//  Asynchronously create a new WebView for use with visual hosting.
	//  <param name="aParentWindow">Handle of the control in which the app will connect the visual tree of the WebView.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3#createcorewebview2compositioncontroller">See the ICoreWebView2Environment3 article.</a>
	CreateCoreWebView2CompositionController(aParentWindow THandle, aBrowserEvents IWVBrowserEvents, aResult *int32) bool // function
	// CreateCoreWebView2PointerInfo
	//  Create an empty ICoreWebView2PointerInfo. The returned
	//  ICoreWebView2PointerInfo needs to be populated with all of the relevant
	//  info before calling SendPointerInput.
	//  <param name="aPointerInfo">The new ICoreWebView2PointerInfo instance.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment3#createcorewebview2pointerinfo">See the ICoreWebView2Environment3 article.</a>
	CreateCoreWebView2PointerInfo(aPointerInfo *ICoreWebView2PointerInfo) bool // function
	// GetAutomationProviderForWindow
	//  Returns the Automation Provider for the WebView that matches the provided
	//  window. Host apps are expected to implement
	//  IRawElementProviderHwndOverride. When GetOverrideProviderForHwnd is
	//  called, the app can pass the HWND to GetAutomationProviderForWindow to
	//  find the matching WebView Automation Provider.
	//  <param name="aHandle">Handle used to find the matching WebView Automation Provider.</param>
	//  <param name="aProvider">The Automation Provider for the WebView that matches the provided window.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment4#getautomationproviderforwindow">See the ICoreWebView2Environment4 article.</a>
	GetAutomationProviderForWindow(aHandle THandle, aProvider *IUnknown) bool // function
	// CreatePrintSettings
	//  Creates the `ICoreWebView2PrintSettings` used by the `PrintToPdf` method.
	//  <param name="aPrintSettings">The new ICoreWebView2PrintSettings instance.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment6#createprintsettings">See the ICoreWebView2Environment6 article.</a>
	CreatePrintSettings(aPrintSettings *ICoreWebView2PrintSettings) bool // function
	// CreateContextMenuItem
	//  Create a custom `ContextMenuItem` object to insert into the WebView context menu.
	//  CoreWebView2 will rewind the icon stream before decoding.
	//  There is a limit of 1000 active custom context menu items at a given time.
	//  Attempting to create more before deleting existing ones will fail with
	//  ERROR_NOT_ENOUGH_QUOTA.
	//  It is recommended to reuse ContextMenuItems across ContextMenuRequested events
	//  for performance.
	//  The returned ContextMenuItem object's `IsEnabled` property will default to `TRUE`
	//  and `IsChecked` property will default to `FALSE`. A `CommandId` will be assigned
	//  to the ContextMenuItem object that's unique across active custom context menu items,
	//  but command ID values of deleted ContextMenuItems can be reassigned.
	//  <param name="aLabel">Context menu item label.</param>
	//  <param name="aIconStream">Context menu item icon as stream.</param>
	//  <param name="aKind">Context menu item kind.</param>
	//  <param name="aMenuItem">The new ICoreWebView2ContextMenuItem instance.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment9#createcontextmenuitem">See the ICoreWebView2Environment9 article.</a>
	CreateContextMenuItem(aLabel string, aIconStream IStream, aKind TWVMenuItemKind, aMenuItem *ICoreWebView2ContextMenuItem) bool // function
	// CreateCoreWebView2ControllerOptions
	//  Create a new ICoreWebView2ControllerOptions to be passed as a parameter of
	//  CreateCoreWebView2ControllerWithOptions and CreateCoreWebView2CompositionControllerWithOptions.
	//  The 'options' is settable and in it the default value for profile name is the empty string,
	//  and the default value for IsInPrivateModeEnabled is false.
	//  Also the profile name can be reused.
	//  <param name="aOptions">The new ICoreWebView2ControllerOptions instance.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10#createcorewebview2controlleroptions">See the ICoreWebView2Environment10 article.</a>
	CreateCoreWebView2ControllerOptions(aOptions *ICoreWebView2ControllerOptions, aResult *int32) bool // function
	// CreateCoreWebView2ControllerWithOptions
	//  Create a new WebView with options.
	//  <param name="aParentWindow">Handle of the control in which the WebView should be displayed.</param>
	//  <param name="aOptions">The ICoreWebView2ControllerOptions instance created with CreateCoreWebView2ControllerOptions.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10#createcorewebview2controllerwithoptions">See the ICoreWebView2Environment10 article.</a>
	CreateCoreWebView2ControllerWithOptions(aParentWindow HWND, aOptions ICoreWebView2ControllerOptions, aBrowserEvents IWVBrowserEvents, aResult *int32) bool // function
	// CreateCoreWebView2CompositionControllerWithOptions
	//  Create a new WebView in visual hosting mode with options.
	//  <param name="aParentWindow">Handle of the control in which the app will connect the visual tree of the WebView.</param>
	//  <param name="aOptions">The ICoreWebView2ControllerOptions instance created with CreateCoreWebView2ControllerOptions.</param>
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <param name="aResult">Result code.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment10#createcorewebview2compositioncontrollerwithoptions">See the ICoreWebView2Environment10 article.</a>
	CreateCoreWebView2CompositionControllerWithOptions(aParentWindow HWND, aOptions ICoreWebView2ControllerOptions, aBrowserEvents IWVBrowserEvents, aResult *int32) bool // function
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
	//  <param name="aSize">Buffer size in bytes.</param>
	//  <param name="aSharedBuffer">The new ICoreWebView2SharedBuffer instance.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment12#createsharedbuffer">See the ICoreWebView2Environment12 article.</a>
	CreateSharedBuffer(aSize int64, aSharedBuffer *ICoreWebView2SharedBuffer) bool // function
	// GetProcessExtendedInfos
	//  Gets a snapshot collection of `ProcessExtendedInfo`s corresponding to all
	//  currently running processes associated with this `ICoreWebView2Environment`
	//  excludes crashpad process.
	//  This provides the same list of `ProcessInfo`s as what's provided in
	//  `GetProcessInfos`, but additionally provides a list of associated `FrameInfo`s
	//  which are actively running(showing or hiding UI elements) in the renderer
	//  process. See `AssociatedFrameInfos` for more information.
	//  <param name="aBrowserEvents">The TWVBrowserBase instance that will receive all the events.</param>
	//  <returns>True if successfull.</return>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment13#getprocessextendedinfos">See the ICoreWebView2Environment13 article.</a>
	GetProcessExtendedInfos(aBrowserEvents IWVBrowserEvents) bool // function
}

// TCoreWebView2Environment Parent: TObject
//
//	Represents the WebView2 Environment.  WebViews created from an environment
//	run on the browser process specified with environment parameters and
//	objects created from an environment should be used in the same
//	environment.  Using it in different environments are not guaranteed to be
//	 compatible and may fail.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2environment">See the ICoreWebView2Environment article.</a>
type TCoreWebView2Environment struct {
	TObject
}

func NewCoreWebView2Environment(aBaseIntf ICoreWebView2Environment) ICoreWebView2Environment {
	r1 := WV().SysCallN(284, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Environment(r1)
}

func (m *TCoreWebView2Environment) Initialized() bool {
	r1 := WV().SysCallN(299, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) BaseIntf() ICoreWebView2Environment {
	var resultCoreWebView2Environment uintptr
	WV().SysCallN(281, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2Environment)))
	return AsCoreWebView2Environment(resultCoreWebView2Environment)
}

func (m *TCoreWebView2Environment) BrowserVersionInfo() string {
	r1 := WV().SysCallN(282, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Environment) SupportsCompositionController() bool {
	r1 := WV().SysCallN(301, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) SupportsControllerOptions() bool {
	r1 := WV().SysCallN(302, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) UserDataFolder() string {
	r1 := WV().SysCallN(303, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Environment) ProcessInfos() ICoreWebView2ProcessInfoCollection {
	var resultCoreWebView2ProcessInfoCollection uintptr
	WV().SysCallN(300, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ProcessInfoCollection)))
	return AsCoreWebView2ProcessInfoCollection(resultCoreWebView2ProcessInfoCollection)
}

func (m *TCoreWebView2Environment) FailureReportFolderPath() string {
	r1 := WV().SysCallN(296, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Environment) AddAllLoaderEvents(aLoaderComponent IComponent) bool {
	r1 := WV().SysCallN(280, m.Instance(), GetObjectUintptr(aLoaderComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(279, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2Controller(aParentWindow THandle, aBrowserEvents IWVBrowserEvents, aResult *int32) bool {
	var result2 uintptr
	r1 := WV().SysCallN(288, m.Instance(), uintptr(aParentWindow), GetObjectUintptr(aBrowserEvents), uintptr(unsafePointer(&result2)))
	*aResult = int32(result2)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateWebResourceResponse(aContent IStream, aStatusCode int32, aReasonPhrase, aHeaders string, aResponse *ICoreWebView2WebResourceResponse) bool {
	var result3 uintptr
	r1 := WV().SysCallN(295, m.Instance(), GetObjectUintptr(aContent), uintptr(aStatusCode), PascalStr(aReasonPhrase), PascalStr(aHeaders), uintptr(unsafePointer(&result3)))
	*aResponse = AsCoreWebView2WebResourceResponse(result3)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateWebResourceRequest(aURI, aMethod string, aPostData IStream, aHeaders string, aRequest *ICoreWebView2WebResourceRequestRef) bool {
	var result3 uintptr
	r1 := WV().SysCallN(294, m.Instance(), PascalStr(aURI), PascalStr(aMethod), GetObjectUintptr(aPostData), PascalStr(aHeaders), uintptr(unsafePointer(&result3)))
	*aRequest = AsCoreWebView2WebResourceRequestRef(result3)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2CompositionController(aParentWindow THandle, aBrowserEvents IWVBrowserEvents, aResult *int32) bool {
	var result2 uintptr
	r1 := WV().SysCallN(286, m.Instance(), uintptr(aParentWindow), GetObjectUintptr(aBrowserEvents), uintptr(unsafePointer(&result2)))
	*aResult = int32(result2)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2PointerInfo(aPointerInfo *ICoreWebView2PointerInfo) bool {
	var result0 uintptr
	r1 := WV().SysCallN(291, m.Instance(), uintptr(unsafePointer(&result0)))
	*aPointerInfo = AsCoreWebView2PointerInfo(result0)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) GetAutomationProviderForWindow(aHandle THandle, aProvider *IUnknown) bool {
	var result1 uintptr
	r1 := WV().SysCallN(297, m.Instance(), uintptr(aHandle), uintptr(unsafePointer(&result1)))
	*aProvider = AsUnknown(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreatePrintSettings(aPrintSettings *ICoreWebView2PrintSettings) bool {
	var result0 uintptr
	r1 := WV().SysCallN(292, m.Instance(), uintptr(unsafePointer(&result0)))
	*aPrintSettings = AsCoreWebView2PrintSettings(result0)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateContextMenuItem(aLabel string, aIconStream IStream, aKind TWVMenuItemKind, aMenuItem *ICoreWebView2ContextMenuItem) bool {
	var result3 uintptr
	r1 := WV().SysCallN(285, m.Instance(), PascalStr(aLabel), GetObjectUintptr(aIconStream), uintptr(aKind), uintptr(unsafePointer(&result3)))
	*aMenuItem = AsCoreWebView2ContextMenuItem(result3)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2ControllerOptions(aOptions *ICoreWebView2ControllerOptions, aResult *int32) bool {
	var result0 uintptr
	var result1 uintptr
	r1 := WV().SysCallN(289, m.Instance(), uintptr(unsafePointer(&result0)), uintptr(unsafePointer(&result1)))
	*aOptions = AsCoreWebView2ControllerOptions(result0)
	*aResult = int32(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2ControllerWithOptions(aParentWindow HWND, aOptions ICoreWebView2ControllerOptions, aBrowserEvents IWVBrowserEvents, aResult *int32) bool {
	var result3 uintptr
	r1 := WV().SysCallN(290, m.Instance(), uintptr(aParentWindow), GetObjectUintptr(aOptions), GetObjectUintptr(aBrowserEvents), uintptr(unsafePointer(&result3)))
	*aResult = int32(result3)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateCoreWebView2CompositionControllerWithOptions(aParentWindow HWND, aOptions ICoreWebView2ControllerOptions, aBrowserEvents IWVBrowserEvents, aResult *int32) bool {
	var result3 uintptr
	r1 := WV().SysCallN(287, m.Instance(), uintptr(aParentWindow), GetObjectUintptr(aOptions), GetObjectUintptr(aBrowserEvents), uintptr(unsafePointer(&result3)))
	*aResult = int32(result3)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) CreateSharedBuffer(aSize int64, aSharedBuffer *ICoreWebView2SharedBuffer) bool {
	var result1 uintptr
	r1 := WV().SysCallN(293, m.Instance(), uintptr(unsafePointer(&aSize)), uintptr(unsafePointer(&result1)))
	*aSharedBuffer = AsCoreWebView2SharedBuffer(result1)
	return GoBool(r1)
}

func (m *TCoreWebView2Environment) GetProcessExtendedInfos(aBrowserEvents IWVBrowserEvents) bool {
	r1 := WV().SysCallN(298, m.Instance(), GetObjectUintptr(aBrowserEvents))
	return GoBool(r1)
}

func CoreWebView2EnvironmentClass() TClass {
	ret := WV().SysCallN(283)
	return TClass(ret)
}
