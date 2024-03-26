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
	"unsafe"
)

// IATGauge Parent: IGraphicControl
type IATGauge interface {
	IGraphicControl
	PercentDone() int32                            // property
	Theme() (resultPATFlatTheme TATFlatTheme)      // property
	SetTheme(AValue *TATFlatTheme)                 // property
	BorderStyle() TBorderStyle                     // property
	SetBorderStyle(AValue TBorderStyle)            // property
	DoubleBuffered() bool                          // property
	SetDoubleBuffered(AValue bool)                 // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	Kind() TATGaugeKind                            // property
	SetKind(AValue TATGaugeKind)                   // property
	Progress() int32                               // property
	SetProgress(AValue int32)                      // property
	MinValue() int32                               // property
	SetMinValue(AValue int32)                      // property
	MaxValue() int32                               // property
	SetMaxValue(AValue int32)                      // property
	ShowText() bool                                // property
	SetShowText(AValue bool)                       // property
	ShowTextInverted() bool                        // property
	SetShowTextInverted(AValue bool)               // property
	AddProgress(AValue int32)                      // procedure
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

// TATGauge Parent: TGraphicControl
type TATGauge struct {
	TGraphicControl
	dblClickPtr       uintptr
	contextPopupPtr   uintptr
	mouseDownPtr      uintptr
	mouseUpPtr        uintptr
	mouseMovePtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
}

func NewATGauge(AOwner IComponent) IATGauge {
	r1 := LCL().SysCallN(3, GetObjectUintptr(AOwner))
	return AsATGauge(r1)
}

func (m *TATGauge) PercentDone() int32 {
	r1 := LCL().SysCallN(10, m.Instance())
	return int32(r1)
}

func (m *TATGauge) Theme() (resultPATFlatTheme TATFlatTheme) {
	r1 := LCL().SysCallN(24, 0, m.Instance(), 0)
	return *(*TATFlatTheme)(getPointer(r1))
}

func (m *TATGauge) SetTheme(AValue *TATFlatTheme) {
	LCL().SysCallN(24, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)))
}

func (m *TATGauge) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TATGauge) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) DoubleBuffered() bool {
	r1 := LCL().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetDoubleBuffered(AValue bool) {
	LCL().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) ParentColor() bool {
	r1 := LCL().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetParentColor(AValue bool) {
	LCL().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) ParentShowHint() bool {
	r1 := LCL().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetParentShowHint(AValue bool) {
	LCL().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) Kind() TATGaugeKind {
	r1 := LCL().SysCallN(5, 0, m.Instance(), 0)
	return TATGaugeKind(r1)
}

func (m *TATGauge) SetKind(AValue TATGaugeKind) {
	LCL().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) Progress() int32 {
	r1 := LCL().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TATGauge) SetProgress(AValue int32) {
	LCL().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) MinValue() int32 {
	r1 := LCL().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TATGauge) SetMinValue(AValue int32) {
	LCL().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) MaxValue() int32 {
	r1 := LCL().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TATGauge) SetMaxValue(AValue int32) {
	LCL().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) ShowText() bool {
	r1 := LCL().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetShowText(AValue bool) {
	LCL().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TATGauge) ShowTextInverted() bool {
	r1 := LCL().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TATGauge) SetShowTextInverted(AValue bool) {
	LCL().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func ATGaugeClass() TClass {
	ret := LCL().SysCallN(2)
	return TClass(ret)
}

func (m *TATGauge) AddProgress(AValue int32) {
	LCL().SysCallN(0, m.Instance(), uintptr(AValue))
}

func (m *TATGauge) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(13, m.Instance(), m.dblClickPtr)
}

func (m *TATGauge) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(12, m.Instance(), m.contextPopupPtr)
}

func (m *TATGauge) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(14, m.Instance(), m.mouseDownPtr)
}

func (m *TATGauge) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(18, m.Instance(), m.mouseUpPtr)
}

func (m *TATGauge) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(17, m.Instance(), m.mouseMovePtr)
}

func (m *TATGauge) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(15, m.Instance(), m.mouseEnterPtr)
}

func (m *TATGauge) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(16, m.Instance(), m.mouseLeavePtr)
}

func (m *TATGauge) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(19, m.Instance(), m.mouseWheelPtr)
}

func (m *TATGauge) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(20, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TATGauge) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(21, m.Instance(), m.mouseWheelUpPtr)
}
