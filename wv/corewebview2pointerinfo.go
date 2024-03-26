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

// ICoreWebView2PointerInfo Parent: IObject
//
//	This mostly represents a combined win32
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO object. It takes fields
//	from all three and excludes some win32 specific data types like HWND and
//	HANDLE. Note, sourceDevice is taken out but we expect the PointerDeviceRect
//	and DisplayRect to cover the existing use cases of sourceDevice.
//	Another big difference is that any of the point or rect locations are
//	expected to be in WebView physical coordinates. That is, coordinates
//	relative to the WebView and no DPI scaling applied.
//	The PointerId, PointerFlags, ButtonChangeKind, PenFlags, PenMask, TouchFlags,
//	and TouchMask are all #defined flags or enums in the
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO structure. We define those
//	properties here as UINT32 or INT32 and expect the developer to know how to
//	populate those values based on the Windows definitions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo">See the ICoreWebView2PointerInfo article.</a>
type ICoreWebView2PointerInfo interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2PointerInfo // property
	// PointerKind
	//  Get the PointerKind of the pointer event. This corresponds to the
	//  pointerKind property of the POINTER_INFO struct. The values are defined by
	//  the POINTER_INPUT_KIND enum in the Windows SDK(winuser.h). Supports
	//  PT_PEN and PT_TOUCH.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerkind">See the ICoreWebView2PointerInfo article.</a>
	PointerKind() uint32 // property
	// SetPointerKind Set PointerKind
	SetPointerKind(AValue uint32) // property
	// PointerId
	//  Get the PointerId of the pointer event. This corresponds to the pointerId
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerid">See the ICoreWebView2PointerInfo article.</a>
	PointerId() uint32 // property
	// SetPointerId Set PointerId
	SetPointerId(AValue uint32) // property
	// FrameId
	//  Get the FrameID of the pointer event. This corresponds to the frameId
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_frameid">See the ICoreWebView2PointerInfo article.</a>
	FrameId() uint32 // property
	// SetFrameId Set FrameId
	SetFrameId(AValue uint32) // property
	// PointerFlags
	//  Get the PointerFlags of the pointer event. This corresponds to the
	//  pointerFlags property of the POINTER_INFO struct. The values are defined
	//  by the POINTER_FLAGS constants in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerflags">See the ICoreWebView2PointerInfo article.</a>
	PointerFlags() uint32 // property
	// SetPointerFlags Set PointerFlags
	SetPointerFlags(AValue uint32) // property
	// PointerDeviceRect
	//  Get the PointerDeviceRect of the sourceDevice property of the
	//  POINTER_INFO struct as defined in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pointerdevicerect">See the ICoreWebView2PointerInfo article.</a>
	PointerDeviceRect() (resultRect TRect) // property
	// SetPointerDeviceRect Set PointerDeviceRect
	SetPointerDeviceRect(AValue *TRect) // property
	// DisplayRect
	//  Get the DisplayRect of the sourceDevice property of the POINTER_INFO
	//  struct as defined in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_displayrect">See the ICoreWebView2PointerInfo article.</a>
	DisplayRect() (resultRect TRect) // property
	// SetDisplayRect Set DisplayRect
	SetDisplayRect(AValue *TRect) // property
	// PixelLocation
	//  Get the PixelLocation of the pointer event. This corresponds to the
	//  ptPixelLocation property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pixellocation">See the ICoreWebView2PointerInfo article.</a>
	PixelLocation() (resultPoint TPoint) // property
	// SetPixelLocation Set PixelLocation
	SetPixelLocation(AValue *TPoint) // property
	// HimetricLocation
	//  Get the HimetricLocation of the pointer event. This corresponds to the
	//  ptHimetricLocation property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_himetriclocation">See the ICoreWebView2PointerInfo article.</a>
	HimetricLocation() (resultPoint TPoint) // property
	// SetHimetricLocation Set HimetricLocation
	SetHimetricLocation(AValue *TPoint) // property
	// PixelLocationRaw
	//  Get the PixelLocationRaw of the pointer event. This corresponds to the
	//  ptPixelLocationRaw property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pixellocationraw">See the ICoreWebView2PointerInfo article.</a>
	PixelLocationRaw() (resultPoint TPoint) // property
	// SetPixelLocationRaw Set PixelLocationRaw
	SetPixelLocationRaw(AValue *TPoint) // property
	// HimetricLocationRaw
	//  Get the HimetricLocationRaw of the pointer event. This corresponds to the
	//  ptHimetricLocationRaw property of the POINTER_INFO struct as defined in
	//  the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_himetriclocationraw">See the ICoreWebView2PointerInfo article.</a>
	HimetricLocationRaw() (resultPoint TPoint) // property
	// SetHimetricLocationRaw Set HimetricLocationRaw
	SetHimetricLocationRaw(AValue *TPoint) // property
	// Time
	//  Get the Time of the pointer event. This corresponds to the dwTime property
	//  of the POINTER_INFO struct as defined in the Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_time">See the ICoreWebView2PointerInfo article.</a>
	Time() uint32 // property
	// SetTime Set Time
	SetTime(AValue uint32) // property
	// HistoryCount
	//  Get the HistoryCount of the pointer event. This corresponds to the
	//  historyCount property of the POINTER_INFO struct as defined in the Windows
	//  SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_historycount">See the ICoreWebView2PointerInfo article.</a>
	HistoryCount() uint32 // property
	// SetHistoryCount Set HistoryCount
	SetHistoryCount(AValue uint32) // property
	// InputData
	//  Get the InputData of the pointer event. This corresponds to the InputData
	//  property of the POINTER_INFO struct as defined in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_inputdata">See the ICoreWebView2PointerInfo article.</a>
	InputData() int32 // property
	// SetInputData Set InputData
	SetInputData(AValue int32) // property
	// KeyStates
	//  Get the KeyStates of the pointer event. This corresponds to the
	//  dwKeyStates property of the POINTER_INFO struct as defined in the Windows
	//  SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_keystates">See the ICoreWebView2PointerInfo article.</a>
	KeyStates() uint32 // property
	// SetKeyStates Set KeyStates
	SetKeyStates(AValue uint32) // property
	// PerformanceCount
	//  Get the PerformanceCount of the pointer event. This corresponds to the
	//  PerformanceCount property of the POINTER_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_performancecount">See the ICoreWebView2PointerInfo article.</a>
	PerformanceCount() uint64 // property
	// SetPerformanceCount Set PerformanceCount
	SetPerformanceCount(AValue uint64) // property
	// ButtonChangeKind
	//  Get the ButtonChangeKind of the pointer event. This corresponds to the
	//  ButtonChangeKind property of the POINTER_INFO struct. The values are
	//  defined by the POINTER_BUTTON_CHANGE_KIND enum in the Windows SDK
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_buttonchangekind">See the ICoreWebView2PointerInfo article.</a>
	ButtonChangeKind() int32 // property
	// SetButtonChangeKind Set ButtonChangeKind
	SetButtonChangeKind(AValue int32) // property
	// PenFlags
	//  Get the PenFlags of the pointer event. This corresponds to the penFlags
	//  property of the POINTER_PEN_INFO struct. The values are defined by the
	//  PEN_FLAGS constants in the Windows SDK(winuser.h).
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penflags">See the ICoreWebView2PointerInfo article.</a>
	PenFlags() uint32 // property
	// SetPenFlags Set PenFlags
	SetPenFlags(AValue uint32) // property
	// PenMask
	//  Get the PenMask of the pointer event. This corresponds to the penMask
	//  property of the POINTER_PEN_INFO struct. The values are defined by the
	//  PEN_MASK constants in the Windows SDK(winuser.h).
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penmask">See the ICoreWebView2PointerInfo article.</a>
	PenMask() uint32 // property
	// SetPenMask Set PenMask
	SetPenMask(AValue uint32) // property
	// PenPressure
	//  Get the PenPressure of the pointer event. This corresponds to the pressure
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penpressure">See the ICoreWebView2PointerInfo article.</a>
	PenPressure() uint32 // property
	// SetPenPressure Set PenPressure
	SetPenPressure(AValue uint32) // property
	// PenRotation
	//  Get the PenRotation of the pointer event. This corresponds to the rotation
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_penrotation">See the ICoreWebView2PointerInfo article.</a>
	PenRotation() uint32 // property
	// SetPenRotation Set PenRotation
	SetPenRotation(AValue uint32) // property
	// PenTiltX
	//  Get the PenTiltX of the pointer event. This corresponds to the tiltX
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pentiltx">See the ICoreWebView2PointerInfo article.</a>
	PenTiltX() int32 // property
	// SetPenTiltX Set PenTiltX
	SetPenTiltX(AValue int32) // property
	// PenTiltY
	//  Get the PenTiltY of the pointer event. This corresponds to the tiltY
	//  property of the POINTER_PEN_INFO struct as defined in the Windows SDK
	//  This is a Pen specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_pentilty">See the ICoreWebView2PointerInfo article.</a>
	PenTiltY() int32 // property
	// SetPenTiltY Set PenTiltY
	SetPenTiltY(AValue int32) // property
	// TouchFlags
	//  Get the TouchFlags of the pointer event. This corresponds to the
	//  touchFlags property of the POINTER_TOUCH_INFO struct. The values are
	//  defined by the TOUCH_FLAGS constants in the Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchflags">See the ICoreWebView2PointerInfo article.</a>
	TouchFlags() uint32 // property
	// SetTouchFlags Set TouchFlags
	SetTouchFlags(AValue uint32) // property
	// TouchMask
	//  Get the TouchMask of the pointer event. This corresponds to the
	//  touchMask property of the POINTER_TOUCH_INFO struct. The values are
	//  defined by the TOUCH_MASK constants in the Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchmask">See the ICoreWebView2PointerInfo article.</a>
	TouchMask() uint32 // property
	// SetTouchMask Set TouchMask
	SetTouchMask(AValue uint32) // property
	// TouchContact
	//  Get the TouchContact of the pointer event. This corresponds to the
	//  rcContact property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchcontact">See the ICoreWebView2PointerInfo article.</a>
	TouchContact() (resultRect TRect) // property
	// SetTouchContact Set TouchContact
	SetTouchContact(AValue *TRect) // property
	// TouchContactRaw
	//  Get the TouchContactRaw of the pointer event. This corresponds to the
	//  rcContactRaw property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchcontactraw">See the ICoreWebView2PointerInfo article.</a>
	TouchContactRaw() (resultRect TRect) // property
	// SetTouchContactRaw Set TouchContactRaw
	SetTouchContactRaw(AValue *TRect) // property
	// TouchOrientation
	//  Get the TouchOrientation of the pointer event. This corresponds to the
	//  orientation property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchorientation">See the ICoreWebView2PointerInfo article.</a>
	TouchOrientation() uint32 // property
	// SetTouchOrientation Set TouchOrientation
	SetTouchOrientation(AValue uint32) // property
	// TouchPressure
	//  Get the TouchPressure of the pointer event. This corresponds to the
	//  pressure property of the POINTER_TOUCH_INFO struct as defined in the
	//  Windows SDK(winuser.h).
	//  This is a Touch specific attribute.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo#get_touchpressure">See the ICoreWebView2PointerInfo article.</a>
	TouchPressure() uint32 // property
	// SetTouchPressure Set TouchPressure
	SetTouchPressure(AValue uint32) // property
}

// TCoreWebView2PointerInfo Parent: TObject
//
//	This mostly represents a combined win32
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO object. It takes fields
//	from all three and excludes some win32 specific data types like HWND and
//	HANDLE. Note, sourceDevice is taken out but we expect the PointerDeviceRect
//	and DisplayRect to cover the existing use cases of sourceDevice.
//	Another big difference is that any of the point or rect locations are
//	expected to be in WebView physical coordinates. That is, coordinates
//	relative to the WebView and no DPI scaling applied.
//	The PointerId, PointerFlags, ButtonChangeKind, PenFlags, PenMask, TouchFlags,
//	and TouchMask are all #defined flags or enums in the
//	POINTER_INFO/POINTER_TOUCH_INFO/POINTER_PEN_INFO structure. We define those
//	properties here as UINT32 or INT32 and expect the developer to know how to
//	populate those values based on the Windows definitions.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2pointerinfo">See the ICoreWebView2PointerInfo article.</a>
type TCoreWebView2PointerInfo struct {
	TObject
}

func NewCoreWebView2PointerInfo(aBaseIntf ICoreWebView2PointerInfo) ICoreWebView2PointerInfo {
	r1 := WV().SysCallN(462, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2PointerInfo(r1)
}

func (m *TCoreWebView2PointerInfo) Initialized() bool {
	r1 := WV().SysCallN(468, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2PointerInfo) BaseIntf() ICoreWebView2PointerInfo {
	var resultCoreWebView2PointerInfo uintptr
	WV().SysCallN(459, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2PointerInfo)))
	return AsCoreWebView2PointerInfo(resultCoreWebView2PointerInfo)
}

func (m *TCoreWebView2PointerInfo) PointerKind() uint32 {
	r1 := WV().SysCallN(483, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPointerKind(AValue uint32) {
	WV().SysCallN(483, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PointerId() uint32 {
	r1 := WV().SysCallN(482, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPointerId(AValue uint32) {
	WV().SysCallN(482, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) FrameId() uint32 {
	r1 := WV().SysCallN(464, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetFrameId(AValue uint32) {
	WV().SysCallN(464, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PointerFlags() uint32 {
	r1 := WV().SysCallN(481, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPointerFlags(AValue uint32) {
	WV().SysCallN(481, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PointerDeviceRect() (resultRect TRect) {
	WV().SysCallN(480, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPointerDeviceRect(AValue *TRect) {
	WV().SysCallN(480, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) DisplayRect() (resultRect TRect) {
	WV().SysCallN(463, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetDisplayRect(AValue *TRect) {
	WV().SysCallN(463, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) PixelLocation() (resultPoint TPoint) {
	WV().SysCallN(478, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPixelLocation(AValue *TPoint) {
	WV().SysCallN(478, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) HimetricLocation() (resultPoint TPoint) {
	WV().SysCallN(465, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetHimetricLocation(AValue *TPoint) {
	WV().SysCallN(465, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) PixelLocationRaw() (resultPoint TPoint) {
	WV().SysCallN(479, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetPixelLocationRaw(AValue *TPoint) {
	WV().SysCallN(479, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) HimetricLocationRaw() (resultPoint TPoint) {
	WV().SysCallN(466, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCoreWebView2PointerInfo) SetHimetricLocationRaw(AValue *TPoint) {
	WV().SysCallN(466, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) Time() uint32 {
	r1 := WV().SysCallN(484, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTime(AValue uint32) {
	WV().SysCallN(484, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) HistoryCount() uint32 {
	r1 := WV().SysCallN(467, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetHistoryCount(AValue uint32) {
	WV().SysCallN(467, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) InputData() int32 {
	r1 := WV().SysCallN(469, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetInputData(AValue int32) {
	WV().SysCallN(469, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) KeyStates() uint32 {
	r1 := WV().SysCallN(470, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetKeyStates(AValue uint32) {
	WV().SysCallN(470, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PerformanceCount() uint64 {
	r1 := WV().SysCallN(477, 0, m.Instance(), 0)
	return uint64(r1)
}

func (m *TCoreWebView2PointerInfo) SetPerformanceCount(AValue uint64) {
	WV().SysCallN(477, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) ButtonChangeKind() int32 {
	r1 := WV().SysCallN(460, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetButtonChangeKind(AValue int32) {
	WV().SysCallN(460, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenFlags() uint32 {
	r1 := WV().SysCallN(471, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenFlags(AValue uint32) {
	WV().SysCallN(471, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenMask() uint32 {
	r1 := WV().SysCallN(472, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenMask(AValue uint32) {
	WV().SysCallN(472, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenPressure() uint32 {
	r1 := WV().SysCallN(473, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenPressure(AValue uint32) {
	WV().SysCallN(473, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenRotation() uint32 {
	r1 := WV().SysCallN(474, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenRotation(AValue uint32) {
	WV().SysCallN(474, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenTiltX() int32 {
	r1 := WV().SysCallN(475, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenTiltX(AValue int32) {
	WV().SysCallN(475, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) PenTiltY() int32 {
	r1 := WV().SysCallN(476, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCoreWebView2PointerInfo) SetPenTiltY(AValue int32) {
	WV().SysCallN(476, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchFlags() uint32 {
	r1 := WV().SysCallN(487, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchFlags(AValue uint32) {
	WV().SysCallN(487, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchMask() uint32 {
	r1 := WV().SysCallN(488, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchMask(AValue uint32) {
	WV().SysCallN(488, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchContact() (resultRect TRect) {
	WV().SysCallN(485, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetTouchContact(AValue *TRect) {
	WV().SysCallN(485, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) TouchContactRaw() (resultRect TRect) {
	WV().SysCallN(486, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCoreWebView2PointerInfo) SetTouchContactRaw(AValue *TRect) {
	WV().SysCallN(486, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCoreWebView2PointerInfo) TouchOrientation() uint32 {
	r1 := WV().SysCallN(489, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchOrientation(AValue uint32) {
	WV().SysCallN(489, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2PointerInfo) TouchPressure() uint32 {
	r1 := WV().SysCallN(490, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCoreWebView2PointerInfo) SetTouchPressure(AValue uint32) {
	WV().SysCallN(490, 1, m.Instance(), uintptr(AValue))
}

func CoreWebView2PointerInfoClass() TClass {
	ret := WV().SysCallN(461)
	return TClass(ret)
}
