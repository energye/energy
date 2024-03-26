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

// IChromium Parent: IChromiumCore
//
//	VCL and LCL version of TChromiumCore that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type IChromium interface {
	IChromiumCore
	// CreateBrowserForBool
	//  Used to create the browser after the global request context has been
	//  initialized. You need to set all properties and events before calling
	//  this function because it will only create the internal handlers needed
	//  for those events and the property values will be used in the browser
	//  initialization.
	//  The browser will be fully initialized when the TChromiumCore.OnAfterCreated
	//  event is triggered.
	CreateBrowserForBool(aBrowserParent IWinControl, aWindowName string, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue) bool // function
	// SaveAsBitmapStream
	//  Copy the DC to a bitmap stream. Only works on Windows with browsers without GPU acceleration.
	//  It's recommended to use the "Page.captureScreenshot" DevTools method instead.
	SaveAsBitmapStream(aStream IStream) bool // function
	// TakeSnapshot
	//  Copy the DC to a TBitmap. Only works on Windows with browsers without GPU acceleration.
	//  It's recommended to use the "Page.captureScreenshot" DevTools method instead.
	TakeSnapshot(aBitmap *IBitmap) bool // function
	// InitializeDragAndDropForIWinControl
	//  Used with browsers in OSR mode to initialize drag and drop in Windows.
	InitializeDragAndDropForIWinControl(aDropTargetCtrl IWinControl) // procedure
	// ShowDevToolsForIWinControl
	//  Open developer tools(DevTools) in its own browser. If inspectElementAt has a valid point
	//  with coordinates different than low(integer) then the element at the specified location
	//  will be inspected. If the DevTools browser is already open then it will be focused.
	ShowDevToolsForIWinControl(inspectElementAt *TPoint, aDevTools IWinControl) // procedure
	// CloseDevToolsForIWinControl
	//  Close the developer tools.
	CloseDevToolsForIWinControl(aDevTools IWinControl) // procedure
	// MoveFormTo
	//  Move the parent form to the x and y coordinates.
	MoveFormTo(x, y int32) // procedure
	// MoveFormBy
	//  Move the parent form adding x and y to the coordinates.
	MoveFormBy(x, y int32) // procedure
	// ResizeFormWidthTo
	//  Add x to the parent form width.
	ResizeFormWidthTo(x int32) // procedure
	// ResizeFormHeightTo
	//  Add y to the parent form height.
	ResizeFormHeightTo(y int32) // procedure
	// SetFormLeftTo
	//  Set the parent form left property to x.
	SetFormLeftTo(x int32) // procedure
	// SetFormTopTo
	//  Set the parent form top property to y.
	SetFormTopTo(y int32) // procedure
}

// TChromium Parent: TChromiumCore
//
//	VCL and LCL version of TChromiumCore that puts together all browser procedures, functions, properties and events in one place.
//	It has all you need to create, modify and destroy a web browser.
type TChromium struct {
	TChromiumCore
}

func NewChromium(aOwner IComponent) IChromium {
	r1 := CEF().SysCallN(2115, GetObjectUintptr(aOwner))
	return AsChromium(r1)
}

func (m *TChromium) CreateBrowserForBool(aBrowserParent IWinControl, aWindowName string, aContext ICefRequestContext, aExtraInfo ICefDictionaryValue) bool {
	r1 := CEF().SysCallN(2116, m.Instance(), GetObjectUintptr(aBrowserParent), PascalStr(aWindowName), GetObjectUintptr(aContext), GetObjectUintptr(aExtraInfo))
	return GoBool(r1)
}

func (m *TChromium) SaveAsBitmapStream(aStream IStream) bool {
	r1 := CEF().SysCallN(2122, m.Instance(), GetObjectUintptr(aStream))
	return GoBool(r1)
}

func (m *TChromium) TakeSnapshot(aBitmap *IBitmap) bool {
	var result0 uintptr
	r1 := CEF().SysCallN(2126, m.Instance(), uintptr(unsafePointer(&result0)))
	*aBitmap = AsBitmap(result0)
	return GoBool(r1)
}

func ChromiumClass() TClass {
	ret := CEF().SysCallN(2113)
	return TClass(ret)
}

func (m *TChromium) InitializeDragAndDropForIWinControl(aDropTargetCtrl IWinControl) {
	CEF().SysCallN(2117, m.Instance(), GetObjectUintptr(aDropTargetCtrl))
}

func (m *TChromium) ShowDevToolsForIWinControl(inspectElementAt *TPoint, aDevTools IWinControl) {
	CEF().SysCallN(2125, m.Instance(), uintptr(unsafePointer(inspectElementAt)), GetObjectUintptr(aDevTools))
}

func (m *TChromium) CloseDevToolsForIWinControl(aDevTools IWinControl) {
	CEF().SysCallN(2114, m.Instance(), GetObjectUintptr(aDevTools))
}

func (m *TChromium) MoveFormTo(x, y int32) {
	CEF().SysCallN(2119, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TChromium) MoveFormBy(x, y int32) {
	CEF().SysCallN(2118, m.Instance(), uintptr(x), uintptr(y))
}

func (m *TChromium) ResizeFormWidthTo(x int32) {
	CEF().SysCallN(2121, m.Instance(), uintptr(x))
}

func (m *TChromium) ResizeFormHeightTo(y int32) {
	CEF().SysCallN(2120, m.Instance(), uintptr(y))
}

func (m *TChromium) SetFormLeftTo(x int32) {
	CEF().SysCallN(2123, m.Instance(), uintptr(x))
}

func (m *TChromium) SetFormTopTo(y int32) {
	CEF().SysCallN(2124, m.Instance(), uintptr(y))
}
