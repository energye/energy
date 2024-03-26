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

// ICoreWebView2CompositionController Parent: IObject
//
//	This interface is an extension of the ICoreWebView2Controller interface to
//	support visual hosting. An object implementing the
//	ICoreWebView2CompositionController interface will also implement
//	ICoreWebView2Controller. Callers are expected to use
//	ICoreWebView2Controller for resizing, visibility, focus, and so on, and
//	then use ICoreWebView2CompositionController to connect to a composition
//	tree and provide input meant for the WebView.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller">See the ICoreWebView2CompositionController article.</a>
type ICoreWebView2CompositionController interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2CompositionController // property
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
	// Cursor
	//  The current cursor that WebView thinks it should be. The cursor should be
	//  set in WM_SETCURSOR through \::SetCursor or set on the corresponding
	//  parent/ancestor HWND of the WebView through \::SetClassLongPtr. The HCURSOR
	//  can be freed so CopyCursor/DestroyCursor is recommended to keep your own
	//  copy if you are doing more than immediately setting the cursor.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller#get_cursor">See the ICoreWebView2CompositionController article.</a>
	Cursor() HCURSOR // property
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
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// SendMouseInput
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
	SendMouseInput(aEventKind TWVMouseEventKind, aVirtualKeys TWVMouseEventVirtualKeys, aMouseData uint32, aPoint *TPoint) bool // function
	// SendPointerInput
	//  SendPointerInput accepts touch or pen pointer input of types defined in
	//  COREWEBVIEW2_POINTER_EVENT_KIND. Any pointer input from the system must be
	//  converted into an ICoreWebView2PointerInfo first.
	SendPointerInput(aEventKind TWVPointerEventKind, aPointerInfo ICoreWebView2PointerInfo) bool // function
	// DragEnter
	//  This function corresponds to [IDropTarget::DragEnter](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragenter).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragEnter calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates(Similar to how SendMouseInput works).
	DragEnter(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 // function
	// DragLeave
	//  This function corresponds to [IDropTarget::DragLeave](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragleave).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragLeave calls to this function.
	DragLeave() int32 // function
	// DragOver
	//  This function corresponds to [IDropTarget::DragOver](/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragover).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward DragOver calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates(Similar to how SendMouseInput works).
	DragOver(keyState uint32, point *TPoint, OutEffect *uint32) int32 // function
	// Drop
	//  This function corresponds to [IDropTarget::Drop](/windows/win32/api/oleidl/nf-oleidl-idroptarget-drop).
	//  This function has a dependency on AllowExternalDrop property of
	//  CoreWebView2Controller and return E_FAIL to callers to indicate this
	//  operation is not allowed if AllowExternalDrop property is set to false.
	//  The hosting application must register as an IDropTarget and implement
	//  and forward Drop calls to this function.
	//  point parameter must be modified to include the WebView's offset and be in
	//  the WebView's client coordinates(Similar to how SendMouseInput works).
	Drop(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 // function
}

// TCoreWebView2CompositionController Parent: TObject
//
//	This interface is an extension of the ICoreWebView2Controller interface to
//	support visual hosting. An object implementing the
//	ICoreWebView2CompositionController interface will also implement
//	ICoreWebView2Controller. Callers are expected to use
//	ICoreWebView2Controller for resizing, visibility, focus, and so on, and
//	then use ICoreWebView2CompositionController to connect to a composition
//	tree and provide input meant for the WebView.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2compositioncontroller">See the ICoreWebView2CompositionController article.</a>
type TCoreWebView2CompositionController struct {
	TObject
}

func NewCoreWebView2CompositionController(aBaseIntf ICoreWebView2CompositionController) ICoreWebView2CompositionController {
	r1 := WV().SysCallN(105, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2CompositionController(r1)
}

func (m *TCoreWebView2CompositionController) Initialized() bool {
	r1 := WV().SysCallN(111, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2CompositionController) BaseIntf() ICoreWebView2CompositionController {
	var resultCoreWebView2CompositionController uintptr
	WV().SysCallN(103, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CompositionController)))
	return AsCoreWebView2CompositionController(resultCoreWebView2CompositionController)
}

func (m *TCoreWebView2CompositionController) RootVisualTarget() IUnknown {
	var resultUnknown uintptr
	WV().SysCallN(112, 0, m.Instance(), 0, uintptr(unsafePointer(&resultUnknown)))
	return AsUnknown(resultUnknown)
}

func (m *TCoreWebView2CompositionController) SetRootVisualTarget(AValue IUnknown) {
	WV().SysCallN(112, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2CompositionController) Cursor() HCURSOR {
	r1 := WV().SysCallN(106, m.Instance())
	return HCURSOR(r1)
}

func (m *TCoreWebView2CompositionController) SystemCursorID() uint32 {
	r1 := WV().SysCallN(115, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2CompositionController) AutomationProvider() IUnknown {
	var resultUnknown uintptr
	WV().SysCallN(102, m.Instance(), uintptr(unsafePointer(&resultUnknown)))
	return AsUnknown(resultUnknown)
}

func (m *TCoreWebView2CompositionController) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(101, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2CompositionController) SendMouseInput(aEventKind TWVMouseEventKind, aVirtualKeys TWVMouseEventVirtualKeys, aMouseData uint32, aPoint *TPoint) bool {
	r1 := WV().SysCallN(113, m.Instance(), uintptr(aEventKind), uintptr(aVirtualKeys), uintptr(aMouseData), uintptr(unsafePointer(aPoint)))
	return GoBool(r1)
}

func (m *TCoreWebView2CompositionController) SendPointerInput(aEventKind TWVPointerEventKind, aPointerInfo ICoreWebView2PointerInfo) bool {
	r1 := WV().SysCallN(114, m.Instance(), uintptr(aEventKind), GetObjectUintptr(aPointerInfo))
	return GoBool(r1)
}

func (m *TCoreWebView2CompositionController) DragEnter(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 {
	var result3 uintptr
	r1 := WV().SysCallN(107, m.Instance(), GetObjectUintptr(dataObject), uintptr(keyState), uintptr(unsafePointer(point)), uintptr(unsafePointer(&result3)))
	*OutEffect = uint32(result3)
	return int32(r1)
}

func (m *TCoreWebView2CompositionController) DragLeave() int32 {
	r1 := WV().SysCallN(108, m.Instance())
	return int32(r1)
}

func (m *TCoreWebView2CompositionController) DragOver(keyState uint32, point *TPoint, OutEffect *uint32) int32 {
	var result2 uintptr
	r1 := WV().SysCallN(109, m.Instance(), uintptr(keyState), uintptr(unsafePointer(point)), uintptr(unsafePointer(&result2)))
	*OutEffect = uint32(result2)
	return int32(r1)
}

func (m *TCoreWebView2CompositionController) Drop(dataObject IDataObject, keyState uint32, point *TPoint, OutEffect *uint32) int32 {
	var result3 uintptr
	r1 := WV().SysCallN(110, m.Instance(), GetObjectUintptr(dataObject), uintptr(keyState), uintptr(unsafePointer(point)), uintptr(unsafePointer(&result3)))
	*OutEffect = uint32(result3)
	return int32(r1)
}

func CoreWebView2CompositionControllerClass() TClass {
	ret := WV().SysCallN(104)
	return TClass(ret)
}
