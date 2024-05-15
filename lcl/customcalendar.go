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

// ICustomCalendar Parent: IWinControl
type ICustomCalendar interface {
	IWinControl
	Date() string                               // property
	SetDate(AValue string)                      // property
	DateTime() TDateTime                        // property
	SetDateTime(AValue TDateTime)               // property
	DisplaySettings() TDisplaySettings          // property
	SetDisplaySettings(AValue TDisplaySettings) // property
	FirstDayOfWeek() TCalDayOfWeek              // property
	SetFirstDayOfWeek(AValue TCalDayOfWeek)     // property
	MaxDate() TDateTime                         // property
	SetMaxDate(AValue TDateTime)                // property
	MinDate() TDateTime                         // property
	SetMinDate(AValue TDateTime)                // property
	HitTest(APoint *TPoint) TCalendarPart       // function
	GetCalendarView() TCalendarView             // function
	SetOnChange(fn TNotifyEvent)                // property event
	SetOnDayChanged(fn TNotifyEvent)            // property event
	SetOnMonthChanged(fn TNotifyEvent)          // property event
	SetOnYearChanged(fn TNotifyEvent)           // property event
}

// TCustomCalendar Parent: TWinControl
type TCustomCalendar struct {
	TWinControl
	changePtr       uintptr
	dayChangedPtr   uintptr
	monthChangedPtr uintptr
	yearChangedPtr  uintptr
}

func NewCustomCalendar(AOwner IComponent) ICustomCalendar {
	r1 := LCL().SysCallN(1343, GetObjectUintptr(AOwner))
	return AsCustomCalendar(r1)
}

func (m *TCustomCalendar) Date() string {
	r1 := LCL().SysCallN(1344, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCustomCalendar) SetDate(AValue string) {
	LCL().SysCallN(1344, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCustomCalendar) DateTime() TDateTime {
	r1 := LCL().SysCallN(1345, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCustomCalendar) SetDateTime(AValue TDateTime) {
	LCL().SysCallN(1345, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) DisplaySettings() TDisplaySettings {
	r1 := LCL().SysCallN(1346, 0, m.Instance(), 0)
	return TDisplaySettings(r1)
}

func (m *TCustomCalendar) SetDisplaySettings(AValue TDisplaySettings) {
	LCL().SysCallN(1346, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) FirstDayOfWeek() TCalDayOfWeek {
	r1 := LCL().SysCallN(1347, 0, m.Instance(), 0)
	return TCalDayOfWeek(r1)
}

func (m *TCustomCalendar) SetFirstDayOfWeek(AValue TCalDayOfWeek) {
	LCL().SysCallN(1347, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) MaxDate() TDateTime {
	r1 := LCL().SysCallN(1350, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCustomCalendar) SetMaxDate(AValue TDateTime) {
	LCL().SysCallN(1350, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) MinDate() TDateTime {
	r1 := LCL().SysCallN(1351, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TCustomCalendar) SetMinDate(AValue TDateTime) {
	LCL().SysCallN(1351, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomCalendar) HitTest(APoint *TPoint) TCalendarPart {
	r1 := LCL().SysCallN(1349, m.Instance(), uintptr(unsafePointer(APoint)))
	return TCalendarPart(r1)
}

func (m *TCustomCalendar) GetCalendarView() TCalendarView {
	r1 := LCL().SysCallN(1348, m.Instance())
	return TCalendarView(r1)
}

func CustomCalendarClass() TClass {
	ret := LCL().SysCallN(1342)
	return TClass(ret)
}

func (m *TCustomCalendar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1352, m.Instance(), m.changePtr)
}

func (m *TCustomCalendar) SetOnDayChanged(fn TNotifyEvent) {
	if m.dayChangedPtr != 0 {
		RemoveEventElement(m.dayChangedPtr)
	}
	m.dayChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1353, m.Instance(), m.dayChangedPtr)
}

func (m *TCustomCalendar) SetOnMonthChanged(fn TNotifyEvent) {
	if m.monthChangedPtr != 0 {
		RemoveEventElement(m.monthChangedPtr)
	}
	m.monthChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1354, m.Instance(), m.monthChangedPtr)
}

func (m *TCustomCalendar) SetOnYearChanged(fn TNotifyEvent) {
	if m.yearChangedPtr != 0 {
		RemoveEventElement(m.yearChangedPtr)
	}
	m.yearChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1355, m.Instance(), m.yearChangedPtr)
}
