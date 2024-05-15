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

// IColorButton Parent: ICustomSpeedButton
type IColorButton interface {
	ICustomSpeedButton
	BorderWidth() int32                            // property
	SetBorderWidth(AValue int32)                   // property
	ButtonColorAutoSize() bool                     // property
	SetButtonColorAutoSize(AValue bool)            // property
	ButtonColorSize() int32                        // property
	SetButtonColorSize(AValue int32)               // property
	ButtonColor() TColor                           // property
	SetButtonColor(AValue TColor)                  // property
	ColorDialog() IColorDialog                     // property
	SetColorDialog(AValue IColorDialog)            // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnColorChanged(fn TNotifyEvent)             // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
}

// TColorButton Parent: TCustomSpeedButton
type TColorButton struct {
	TCustomSpeedButton
	colorChangedPtr   uintptr
	dblClickPtr       uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	paintPtr          uintptr
}

func NewColorButton(AnOwner IComponent) IColorButton {
	r1 := LCL().SysCallN(749, GetObjectUintptr(AnOwner))
	return AsColorButton(r1)
}

func (m *TColorButton) BorderWidth() int32 {
	r1 := LCL().SysCallN(743, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorButton) SetBorderWidth(AValue int32) {
	LCL().SysCallN(743, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorButton) ButtonColorAutoSize() bool {
	r1 := LCL().SysCallN(745, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorButton) SetButtonColorAutoSize(AValue bool) {
	LCL().SysCallN(745, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorButton) ButtonColorSize() int32 {
	r1 := LCL().SysCallN(746, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorButton) SetButtonColorSize(AValue int32) {
	LCL().SysCallN(746, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorButton) ButtonColor() TColor {
	r1 := LCL().SysCallN(744, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TColorButton) SetButtonColor(AValue TColor) {
	LCL().SysCallN(744, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorButton) ColorDialog() IColorDialog {
	r1 := LCL().SysCallN(748, 0, m.Instance(), 0)
	return AsColorDialog(r1)
}

func (m *TColorButton) SetColorDialog(AValue IColorDialog) {
	LCL().SysCallN(748, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TColorButton) ParentFont() bool {
	r1 := LCL().SysCallN(750, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorButton) SetParentFont(AValue bool) {
	LCL().SysCallN(750, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(751, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(751, 1, m.Instance(), PascalBool(AValue))
}

func ColorButtonClass() TClass {
	ret := LCL().SysCallN(747)
	return TClass(ret)
}

func (m *TColorButton) SetOnColorChanged(fn TNotifyEvent) {
	if m.colorChangedPtr != 0 {
		RemoveEventElement(m.colorChangedPtr)
	}
	m.colorChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(752, m.Instance(), m.colorChangedPtr)
}

func (m *TColorButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(753, m.Instance(), m.dblClickPtr)
}

func (m *TColorButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(754, m.Instance(), m.mouseDownPtr)
}

func (m *TColorButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(755, m.Instance(), m.mouseEnterPtr)
}

func (m *TColorButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(756, m.Instance(), m.mouseLeavePtr)
}

func (m *TColorButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(757, m.Instance(), m.mouseMovePtr)
}

func (m *TColorButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(758, m.Instance(), m.mouseUpPtr)
}

func (m *TColorButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(759, m.Instance(), m.mouseWheelPtr)
}

func (m *TColorButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(760, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TColorButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(761, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TColorButton) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(762, m.Instance(), m.paintPtr)
}
