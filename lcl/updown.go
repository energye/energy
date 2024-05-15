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

// IUpDown Parent: ICustomUpDown
type IUpDown interface {
	ICustomUpDown
	AlignButton() TUDAlignButton                    // property
	SetAlignButton(AValue TUDAlignButton)           // property
	ArrowKeys() bool                                // property
	SetArrowKeys(AValue bool)                       // property
	Associate() IWinControl                         // property
	SetAssociate(AValue IWinControl)                // property
	Increment() int32                               // property
	SetIncrement(AValue int32)                      // property
	Max() SmallInt                                  // property
	SetMax(AValue SmallInt)                         // property
	Min() SmallInt                                  // property
	SetMin(AValue SmallInt)                         // property
	MinRepeatInterval() Byte                        // property
	SetMinRepeatInterval(AValue Byte)               // property
	Orientation() TUDOrientation                    // property
	SetOrientation(AValue TUDOrientation)           // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	Position() SmallInt                             // property
	SetPosition(AValue SmallInt)                    // property
	Thousands() bool                                // property
	SetThousands(AValue bool)                       // property
	Flat() bool                                     // property
	SetFlat(AValue bool)                            // property
	Wrap() bool                                     // property
	SetWrap(AValue bool)                            // property
	SetOnChanging(fn TUDChangingEvent)              // property event
	SetOnChangingEx(fn TUDChangingEventEx)          // property event
	SetOnClickForUDClickEvent(fn TUDClickEvent)     // property event
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

// TUpDown Parent: TCustomUpDown
type TUpDown struct {
	TCustomUpDown
	changingPtr             uintptr
	changingExPtr           uintptr
	clickForUDClickEventPtr uintptr
	contextPopupPtr         uintptr
	mouseDownPtr            uintptr
	mouseEnterPtr           uintptr
	mouseLeavePtr           uintptr
	mouseMovePtr            uintptr
	mouseUpPtr              uintptr
	mouseWheelPtr           uintptr
	mouseWheelDownPtr       uintptr
	mouseWheelUpPtr         uintptr
	mouseWheelHorzPtr       uintptr
	mouseWheelLeftPtr       uintptr
	mouseWheelRightPtr      uintptr
}

func NewUpDown(AOwner IComponent) IUpDown {
	r1 := LCL().SysCallN(5784, GetObjectUintptr(AOwner))
	return AsUpDown(r1)
}

func (m *TUpDown) AlignButton() TUDAlignButton {
	r1 := LCL().SysCallN(5780, 0, m.Instance(), 0)
	return TUDAlignButton(r1)
}

func (m *TUpDown) SetAlignButton(AValue TUDAlignButton) {
	LCL().SysCallN(5780, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) ArrowKeys() bool {
	r1 := LCL().SysCallN(5781, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetArrowKeys(AValue bool) {
	LCL().SysCallN(5781, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Associate() IWinControl {
	r1 := LCL().SysCallN(5782, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TUpDown) SetAssociate(AValue IWinControl) {
	LCL().SysCallN(5782, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TUpDown) Increment() int32 {
	r1 := LCL().SysCallN(5786, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TUpDown) SetIncrement(AValue int32) {
	LCL().SysCallN(5786, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Max() SmallInt {
	r1 := LCL().SysCallN(5787, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TUpDown) SetMax(AValue SmallInt) {
	LCL().SysCallN(5787, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Min() SmallInt {
	r1 := LCL().SysCallN(5788, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TUpDown) SetMin(AValue SmallInt) {
	LCL().SysCallN(5788, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) MinRepeatInterval() Byte {
	r1 := LCL().SysCallN(5789, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TUpDown) SetMinRepeatInterval(AValue Byte) {
	LCL().SysCallN(5789, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Orientation() TUDOrientation {
	r1 := LCL().SysCallN(5790, 0, m.Instance(), 0)
	return TUDOrientation(r1)
}

func (m *TUpDown) SetOrientation(AValue TUDOrientation) {
	LCL().SysCallN(5790, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) ParentColor() bool {
	r1 := LCL().SysCallN(5791, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetParentColor(AValue bool) {
	LCL().SysCallN(5791, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) ParentShowHint() bool {
	r1 := LCL().SysCallN(5792, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5792, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Position() SmallInt {
	r1 := LCL().SysCallN(5793, 0, m.Instance(), 0)
	return SmallInt(r1)
}

func (m *TUpDown) SetPosition(AValue SmallInt) {
	LCL().SysCallN(5793, 1, m.Instance(), uintptr(AValue))
}

func (m *TUpDown) Thousands() bool {
	r1 := LCL().SysCallN(5809, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetThousands(AValue bool) {
	LCL().SysCallN(5809, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Flat() bool {
	r1 := LCL().SysCallN(5785, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetFlat(AValue bool) {
	LCL().SysCallN(5785, 1, m.Instance(), PascalBool(AValue))
}

func (m *TUpDown) Wrap() bool {
	r1 := LCL().SysCallN(5810, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TUpDown) SetWrap(AValue bool) {
	LCL().SysCallN(5810, 1, m.Instance(), PascalBool(AValue))
}

func UpDownClass() TClass {
	ret := LCL().SysCallN(5783)
	return TClass(ret)
}

func (m *TUpDown) SetOnChanging(fn TUDChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5794, m.Instance(), m.changingPtr)
}

func (m *TUpDown) SetOnChangingEx(fn TUDChangingEventEx) {
	if m.changingExPtr != 0 {
		RemoveEventElement(m.changingExPtr)
	}
	m.changingExPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5795, m.Instance(), m.changingExPtr)
}

func (m *TUpDown) SetOnClickForUDClickEvent(fn TUDClickEvent) {
	if m.clickForUDClickEventPtr != 0 {
		RemoveEventElement(m.clickForUDClickEventPtr)
	}
	m.clickForUDClickEventPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5796, m.Instance(), m.clickForUDClickEventPtr)
}

func (m *TUpDown) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5797, m.Instance(), m.contextPopupPtr)
}

func (m *TUpDown) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5798, m.Instance(), m.mouseDownPtr)
}

func (m *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5799, m.Instance(), m.mouseEnterPtr)
}

func (m *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5800, m.Instance(), m.mouseLeavePtr)
}

func (m *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5801, m.Instance(), m.mouseMovePtr)
}

func (m *TUpDown) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5802, m.Instance(), m.mouseUpPtr)
}

func (m *TUpDown) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5803, m.Instance(), m.mouseWheelPtr)
}

func (m *TUpDown) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5804, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TUpDown) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5808, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TUpDown) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5805, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TUpDown) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5806, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TUpDown) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5807, m.Instance(), m.mouseWheelRightPtr)
}
