//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICustomCoolBar Parent: IToolWindow
type ICustomCoolBar interface {
	IToolWindow
	BandBorderStyle() TBorderStyle                               // property
	SetBandBorderStyle(AValue TBorderStyle)                      // property
	BandMaximize() TCoolBandMaximize                             // property
	SetBandMaximize(AValue TCoolBandMaximize)                    // property
	Bands() ICoolBands                                           // property
	SetBands(AValue ICoolBands)                                  // property
	Bitmap() IBitmap                                             // property
	SetBitmap(AValue IBitmap)                                    // property
	FixedSize() bool                                             // property
	SetFixedSize(AValue bool)                                    // property
	FixedOrder() bool                                            // property
	SetFixedOrder(AValue bool)                                   // property
	GrabStyle() TGrabStyle                                       // property
	SetGrabStyle(AValue TGrabStyle)                              // property
	GrabWidth() int32                                            // property
	SetGrabWidth(AValue int32)                                   // property
	HorizontalSpacing() int32                                    // property
	SetHorizontalSpacing(AValue int32)                           // property
	Images() ICustomImageList                                    // property
	SetImages(AValue ICustomImageList)                           // property
	ImagesWidth() int32                                          // property
	SetImagesWidth(AValue int32)                                 // property
	ShowText() bool                                              // property
	SetShowText(AValue bool)                                     // property
	Themed() bool                                                // property
	SetThemed(AValue bool)                                       // property
	Vertical() bool                                              // property
	SetVertical(AValue bool)                                     // property
	VerticalSpacing() int32                                      // property
	SetVerticalSpacing(AValue int32)                             // property
	AutosizeBands()                                              // procedure
	MouseToBandPos(X, Y int32, OutBand *int32, OutGrabber *bool) // procedure
	SetOnChange(fn TNotifyEvent)                                 // property event
}

// TCustomCoolBar Parent: TToolWindow
type TCustomCoolBar struct {
	TToolWindow
	changePtr uintptr
}

func NewCustomCoolBar(AOwner IComponent) ICustomCoolBar {
	r1 := LCL().SysCallN(1481, GetObjectUintptr(AOwner))
	return AsCustomCoolBar(r1)
}

func (m *TCustomCoolBar) BandBorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1476, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomCoolBar) SetBandBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1476, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) BandMaximize() TCoolBandMaximize {
	r1 := LCL().SysCallN(1477, 0, m.Instance(), 0)
	return TCoolBandMaximize(r1)
}

func (m *TCustomCoolBar) SetBandMaximize(AValue TCoolBandMaximize) {
	LCL().SysCallN(1477, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) Bands() ICoolBands {
	r1 := LCL().SysCallN(1478, 0, m.Instance(), 0)
	return AsCoolBands(r1)
}

func (m *TCustomCoolBar) SetBands(AValue ICoolBands) {
	LCL().SysCallN(1478, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCoolBar) Bitmap() IBitmap {
	r1 := LCL().SysCallN(1479, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TCustomCoolBar) SetBitmap(AValue IBitmap) {
	LCL().SysCallN(1479, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCoolBar) FixedSize() bool {
	r1 := LCL().SysCallN(1483, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetFixedSize(AValue bool) {
	LCL().SysCallN(1483, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) FixedOrder() bool {
	r1 := LCL().SysCallN(1482, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetFixedOrder(AValue bool) {
	LCL().SysCallN(1482, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) GrabStyle() TGrabStyle {
	r1 := LCL().SysCallN(1484, 0, m.Instance(), 0)
	return TGrabStyle(r1)
}

func (m *TCustomCoolBar) SetGrabStyle(AValue TGrabStyle) {
	LCL().SysCallN(1484, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) GrabWidth() int32 {
	r1 := LCL().SysCallN(1485, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetGrabWidth(AValue int32) {
	LCL().SysCallN(1485, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) HorizontalSpacing() int32 {
	r1 := LCL().SysCallN(1486, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetHorizontalSpacing(AValue int32) {
	LCL().SysCallN(1486, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) Images() ICustomImageList {
	r1 := LCL().SysCallN(1487, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomCoolBar) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1487, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomCoolBar) ImagesWidth() int32 {
	r1 := LCL().SysCallN(1488, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetImagesWidth(AValue int32) {
	LCL().SysCallN(1488, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCoolBar) ShowText() bool {
	r1 := LCL().SysCallN(1491, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetShowText(AValue bool) {
	LCL().SysCallN(1491, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) Themed() bool {
	r1 := LCL().SysCallN(1492, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetThemed(AValue bool) {
	LCL().SysCallN(1492, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) Vertical() bool {
	r1 := LCL().SysCallN(1493, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomCoolBar) SetVertical(AValue bool) {
	LCL().SysCallN(1493, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomCoolBar) VerticalSpacing() int32 {
	r1 := LCL().SysCallN(1494, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomCoolBar) SetVerticalSpacing(AValue int32) {
	LCL().SysCallN(1494, 1, m.Instance(), uintptr(AValue))
}

func CustomCoolBarClass() TClass {
	ret := LCL().SysCallN(1480)
	return TClass(ret)
}

func (m *TCustomCoolBar) AutosizeBands() {
	LCL().SysCallN(1475, m.Instance())
}

func (m *TCustomCoolBar) MouseToBandPos(X, Y int32, OutBand *int32, OutGrabber *bool) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(1489, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutBand = int32(result1)
	*OutGrabber = GoBool(result2)
}

func (m *TCustomCoolBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1490, m.Instance(), m.changePtr)
}
