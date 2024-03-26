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
	r1 := LCL().SysCallN(5300, GetObjectUintptr(AOwner))
	return AsXButton(r1)
}

func (m *TXButton) Caption() string {
	r1 := LCL().SysCallN(5298, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TXButton) SetCaption(AValue string) {
	LCL().SysCallN(5298, 1, m.Instance(), PascalStr(AValue))
}

func (m *TXButton) ShowCaption() bool {
	r1 := LCL().SysCallN(5318, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetShowCaption(AValue bool) {
	LCL().SysCallN(5318, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) BackColor() TColor {
	r1 := LCL().SysCallN(5295, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBackColor(AValue TColor) {
	LCL().SysCallN(5295, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverColor() TColor {
	r1 := LCL().SysCallN(5304, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverColor(AValue TColor) {
	LCL().SysCallN(5304, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownColor() TColor {
	r1 := LCL().SysCallN(5301, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownColor(AValue TColor) {
	LCL().SysCallN(5301, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderWidth() int32 {
	r1 := LCL().SysCallN(5297, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TXButton) SetBorderWidth(AValue int32) {
	LCL().SysCallN(5297, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) BorderColor() TColor {
	r1 := LCL().SysCallN(5296, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetBorderColor(AValue TColor) {
	LCL().SysCallN(5296, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) Picture() IPicture {
	r1 := LCL().SysCallN(5310, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TXButton) SetPicture(AValue IPicture) {
	LCL().SysCallN(5310, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TXButton) DrawMode() TDrawImageMode {
	r1 := LCL().SysCallN(5303, 0, m.Instance(), 0)
	return TDrawImageMode(r1)
}

func (m *TXButton) SetDrawMode(AValue TDrawImageMode) {
	LCL().SysCallN(5303, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) NormalFontColor() TColor {
	r1 := LCL().SysCallN(5306, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetNormalFontColor(AValue TColor) {
	LCL().SysCallN(5306, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) DownFontColor() TColor {
	r1 := LCL().SysCallN(5302, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetDownFontColor(AValue TColor) {
	LCL().SysCallN(5302, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) HoverFontColor() TColor {
	r1 := LCL().SysCallN(5305, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TXButton) SetHoverFontColor(AValue TColor) {
	LCL().SysCallN(5305, 1, m.Instance(), uintptr(AValue))
}

func (m *TXButton) ParentFont() bool {
	r1 := LCL().SysCallN(5308, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentFont(AValue bool) {
	LCL().SysCallN(5308, 1, m.Instance(), PascalBool(AValue))
}

func (m *TXButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(5309, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TXButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5309, 1, m.Instance(), PascalBool(AValue))
}

func XButtonClass() TClass {
	ret := LCL().SysCallN(5299)
	return TClass(ret)
}

func (m *TXButton) Paint() {
	LCL().SysCallN(5307, m.Instance())
}

func (m *TXButton) Resize() {
	LCL().SysCallN(5311, m.Instance())
}

func (m *TXButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5312, m.Instance(), m.dblClickPtr)
}

func (m *TXButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5313, m.Instance(), m.mouseDownPtr)
}

func (m *TXButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5314, m.Instance(), m.mouseEnterPtr)
}

func (m *TXButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5315, m.Instance(), m.mouseLeavePtr)
}

func (m *TXButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5316, m.Instance(), m.mouseMovePtr)
}

func (m *TXButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5317, m.Instance(), m.mouseUpPtr)
}
