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

// ICefDisplay Parent: ICefBaseRefCounted
//
//	This class typically, but not always, corresponds to a physical display
//	connected to the system. A fake Display may exist on a headless system, or a
//	Display may correspond to a remote, virtual display. All size and position
//	values are in density independent pixel (DIP) coordinates unless otherwise
//	indicated. Methods must be called on the browser process UI thread unless
//	otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_display_capi.h">CEF source file: /include/capi/views/cef_display_capi.h (cef_display_t)</a>
type ICefDisplay interface {
	ICefBaseRefCounted
	// GetID
	//  Returns the unique identifier for this Display.
	GetID() (resultInt64 int64) // function
	// GetDeviceScaleFactor
	//  Returns this Display's device pixel scale factor. This specifies how much
	//  the UI should be scaled when the actual output has more pixels than
	//  standard displays(which is around 100~120dpi). The potential return
	//  values differ by platform.
	GetDeviceScaleFactor() (resultFloat32 float32) // function
	// GetBounds
	//  Returns this Display's bounds in DIP screen coordinates. This is the full
	//  size of the display.
	GetBounds() (resultCefRect TCefRect) // function
	// GetWorkArea
	//  Returns this Display's work area in DIP screen coordinates. This excludes
	//  areas of the display that are occupied with window manager toolbars, etc.
	GetWorkArea() (resultCefRect TCefRect) // function
	// GetRotation
	//  Returns this Display's rotation in degrees.
	GetRotation() int32 // function
	// ConvertPointToPixels
	//  Convert |point| from DIP coordinates to pixel coordinates using this
	//  Display's device scale factor.
	ConvertPointToPixels(point *TCefPoint) // procedure
	// ConvertPointFromPixels
	//  Convert |point| from pixel coordinates to DIP coordinates using this
	//  Display's device scale factor.
	ConvertPointFromPixels(point *TCefPoint) // procedure
}

// TCefDisplay Parent: TCefBaseRefCounted
//
//	This class typically, but not always, corresponds to a physical display
//	connected to the system. A fake Display may exist on a headless system, or a
//	Display may correspond to a remote, virtual display. All size and position
//	values are in density independent pixel (DIP) coordinates unless otherwise
//	indicated. Methods must be called on the browser process UI thread unless
//	otherwise indicated.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/views/cef_display_capi.h">CEF source file: /include/capi/views/cef_display_capi.h (cef_display_t)</a>
type TCefDisplay struct {
	TCefBaseRefCounted
}

// DisplayRef -> ICefDisplay
var DisplayRef display

// display TCefDisplay Ref
type display uintptr

// UnWrap
//
//	Returns a ICefDisplay instance using a PCefDisplay data pointer.
func (m *display) UnWrap(data uintptr) ICefDisplay {
	var resultCefDisplay uintptr
	CEF().SysCallN(835, uintptr(data), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// Primary
//
//	Returns the primary Display.
func (m *display) Primary() ICefDisplay {
	var resultCefDisplay uintptr
	CEF().SysCallN(830, uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// NearestPoint
//
//	Returns the Display nearest |point|. Set |input_pixel_coords| to true(1) if
//	|point| is in pixel screen coordinates instead of DIP screen coordinates.
func (m *display) NearestPoint(point *TCefPoint, inputpixelcoords bool) ICefDisplay {
	var resultCefDisplay uintptr
	CEF().SysCallN(829, uintptr(unsafePointer(point)), PascalBool(inputpixelcoords), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// MatchingBounds
//
//	Returns the Display that most closely intersects |bounds|. Set
//	|input_pixel_coords| to true(1) if |bounds| is in pixel screen coordinates
//	instead of DIP screen coordinates.
func (m *display) MatchingBounds(bounds *TCefRect, inputpixelcoords bool) ICefDisplay {
	var resultCefDisplay uintptr
	CEF().SysCallN(828, uintptr(unsafePointer(bounds)), PascalBool(inputpixelcoords), uintptr(unsafePointer(&resultCefDisplay)))
	return AsCefDisplay(resultCefDisplay)
}

// GetCount
//
//	Returns the total number of Displays. Mirrored displays are excluded; this
//	function is intended to return the number of distinct, usable displays.
func (m *display) GetCount() NativeUInt {
	r1 := CEF().SysCallN(823)
	return NativeUInt(r1)
}

// ScreenPointToPixels
//
//	Convert |point| from DIP screen coordinates to pixel screen coordinates.
//	This function is only used on Windows.
func (m *display) ScreenPointToPixels(aScreenPoint *TPoint) (resultPoint TPoint) {
	CEF().SysCallN(832, uintptr(unsafePointer(aScreenPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

// ScreenPointFromPixels
//
//	Convert |point| from pixel screen coordinates to DIP screen coordinates.
//	This function is only used on Windows.
func (m *display) ScreenPointFromPixels(aPixelsPoint *TPoint) (resultPoint TPoint) {
	CEF().SysCallN(831, uintptr(unsafePointer(aPixelsPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

// ScreenRectToPixels
//
//	Convert |rect| from DIP screen coordinates to pixel screen coordinates. This
//	function is only used on Windows.
func (m *display) ScreenRectToPixels(aScreenRect *TRect) (resultRect TRect) {
	CEF().SysCallN(834, uintptr(unsafePointer(aScreenRect)), uintptr(unsafePointer(&resultRect)))
	return
}

// ScreenRectFromPixels
//
//	Convert |rect| from pixel screen coordinates to DIP screen coordinates. This
//	function is only used on Windows.
func (m *display) ScreenRectFromPixels(aPixelsRect *TRect) (resultRect TRect) {
	CEF().SysCallN(833, uintptr(unsafePointer(aPixelsRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCefDisplay) GetID() (resultInt64 int64) {
	CEF().SysCallN(825, m.Instance(), uintptr(unsafePointer(&resultInt64)))
	return
}

func (m *TCefDisplay) GetDeviceScaleFactor() (resultFloat32 float32) {
	CEF().SysCallN(824, m.Instance(), uintptr(unsafePointer(&resultFloat32)))
	return
}

func (m *TCefDisplay) GetBounds() (resultCefRect TCefRect) {
	CEF().SysCallN(822, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefDisplay) GetWorkArea() (resultCefRect TCefRect) {
	CEF().SysCallN(827, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefDisplay) GetRotation() int32 {
	r1 := CEF().SysCallN(826, m.Instance())
	return int32(r1)
}

func (m *TCefDisplay) ConvertPointToPixels(point *TCefPoint) {
	var result0 uintptr
	CEF().SysCallN(821, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
}

func (m *TCefDisplay) ConvertPointFromPixels(point *TCefPoint) {
	var result0 uintptr
	CEF().SysCallN(820, m.Instance(), uintptr(unsafePointer(&result0)))
	*point = *(*TCefPoint)(unsafePointer(result0))
}
