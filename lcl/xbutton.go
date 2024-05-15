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

// IXButton Parent: IGraphicControl
type IXButton interface {
	IGraphicControl
	Caption() string                   // property
	SetCaption(AValue string)          // property
	ShowCaption() bool                 // property
	SetShowCaption(AValue bool)        // property
	BackColor() TColor                 // property
	SetBackColor(AValue TColor)        // property
	HoverColor() TColor                // property
	SetHoverColor(AValue TColor)       // property
	DownColor() TColor                 // property
	SetDownColor(AValue TColor)        // property
	BorderWidth() int32                // property
	SetBorderWidth(AValue int32)       // property
	BorderColor() TColor               // property
	SetBorderColor(AValue TColor)      // property
	Picture() IPicture                 // property
	SetPicture(AValue IPicture)        // property
	DrawMode() TDrawImageMode          // property
	SetDrawMode(AValue TDrawImageMode) // property
	NormalFontColor() TColor           // property
	SetNormalFontColor(AValue TColor)  // property
	DownFontColor() TColor             // property
	SetDownFontColor(AValue TColor)    // property
	HoverFontColor() TColor            // property
	SetHoverFontColor(AValue TColor)   // property
	ParentFont() bool                  // property
	SetParentFont(AValue bool)         // property
	ParentShowHint() bool              // property
	SetParentShowHint(AValue bool)     // property
	Paint()                            // procedure
	Resize()                           // procedure
	SetOnDblClick(fn TNotifyEvent)     // property event
	SetOnMouseDown(fn TMouseEvent)     // property event
	SetOnMouseEnter(fn TNotifyEvent)   // property event
	SetOnMouseLeave(fn TNotifyEvent)   // property event
	SetOnMouseMove(fn TMouseMoveEvent) // property event
	SetOnMouseUp(fn TMouseEvent)       // property event
}

// TXButton Parent: TGraphicControl
type TXButton struct {
	TGraphicControl
	dblClickPtr   uintptr
	mouseDownPtr  uintptr
	mouseEnterPtr uintptr
	mouseLeavePtr uintptr
	mouseMovePtr  uintptr
	mouseUpPtr    uintptr
}

func NewXButton(AOwner IComponent) IXButton {
	r1 := LCL().SysCallN(6109, GetObjectUintptr(AOwner))
	return AsXButton(r1)
}

func (m *TXButton) Caption() string {
	r1 := LCL().SysCallN(6107, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TXButton) SetCaption(AValue string) {
	LCL().SysCallN(6107, 1, m.Instance(), PascalStr(AValue))
}

func (m *TXButton) ShowCaption() bool {
	r1 := LCL().SysCallN(6127, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetShowCaption(AValue bool) {
	LCL().SysCallN(6127, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) BackColor() TColor {
	r1 := LCL().SysCallN(6104, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBackColor(AValue TColor) {
	LCL().SysCallN(6104, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverColor() TColor {
	r1 := LCL().SysCallN(6113, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverColor(AValue TColor) {
	LCL().SysCallN(6113, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownColor() TColor {
	r1 := LCL().SysCallN(6110, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownColor(AValue TColor) {
	LCL().SysCallN(6110, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderWidth() int32 {
	r1 := LCL().SysCallN(6106, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TXButton) SetBorderWidth(AValue int32) {
	LCL().SysCallN(6106, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderColor() TColor {
	r1 := LCL().SysCallN(6105, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBorderColor(AValue TColor) {
	LCL().SysCallN(6105, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) Picture() IPicture {
	r1 := LCL().SysCallN(6119, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TXButton) SetPicture(AValue IPicture) {
	LCL().SysCallN(6119, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TXButton) DrawMode() TDrawImageMode {
	r1 := LCL().SysCallN(6112, 0, m.Instance(), 0)
	return TDrawImageMode(r1)
}

func (m *TXButton) SetDrawMode(AValue TDrawImageMode) {
	LCL().SysCallN(6112, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) NormalFontColor() TColor {
	r1 := LCL().SysCallN(6115, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetNormalFontColor(AValue TColor) {
	LCL().SysCallN(6115, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownFontColor() TColor {
	r1 := LCL().SysCallN(6111, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownFontColor(AValue TColor) {
	LCL().SysCallN(6111, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverFontColor() TColor {
	r1 := LCL().SysCallN(6114, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverFontColor(AValue TColor) {
	LCL().SysCallN(6114, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) ParentFont() bool {
	r1 := LCL().SysCallN(6117, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentFont(AValue bool) {
	LCL().SysCallN(6117, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(6118, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(6118, 1, m.Instance(), PascalBool(AValue))
}

func XButtonClass() TClass {
	ret := LCL().SysCallN(6108)
	return TClass(ret)
}

func (m *TXButton) Paint() {
	LCL().SysCallN(6116, m.Instance())
}

func (m *TXButton) Resize() {
	LCL().SysCallN(6120, m.Instance())
}

func (m *TXButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6121, m.Instance(), m.dblClickPtr)
}

func (m *TXButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6122, m.Instance(), m.mouseDownPtr)
}

func (m *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6123, m.Instance(), m.mouseEnterPtr)
}

func (m *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6124, m.Instance(), m.mouseLeavePtr)
}

func (m *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6125, m.Instance(), m.mouseMovePtr)
}

func (m *TXButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(6126, m.Instance(), m.mouseUpPtr)
}
