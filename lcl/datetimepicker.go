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

// IDateTimePicker Parent: ICustomDateTimePicker
type IDateTimePicker interface {
	ICustomDateTimePicker
	DateTime() TDateTime                           // property
	SetDateTime(AValue TDateTime)                  // property
	DroppedDown() bool                             // property
	ArrowShape() TArrowShape                       // property
	SetArrowShape(AValue TArrowShape)              // property
	ShowCheckBox() bool                            // property
	SetShowCheckBox(AValue bool)                   // property
	Checked() bool                                 // property
	SetChecked(AValue bool)                        // property
	CenturyFrom() Word                             // property
	SetCenturyFrom(AValue Word)                    // property
	DateDisplayOrder() TDateDisplayOrder           // property
	SetDateDisplayOrder(AValue TDateDisplayOrder)  // property
	MaxDate() TDate                                // property
	SetMaxDate(AValue TDate)                       // property
	MinDate() TDate                                // property
	SetMinDate(AValue TDate)                       // property
	ReadOnly() bool                                // property
	SetReadOnly(AValue bool)                       // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	DateSeparator() string                         // property
	SetDateSeparator(AValue string)                // property
	TrailingSeparator() bool                       // property
	SetTrailingSeparator(AValue bool)              // property
	TextForNullDate() string                       // property
	SetTextForNullDate(AValue string)              // property
	LeadingZeros() bool                            // property
	SetLeadingZeros(AValue bool)                   // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	NullInputAllowed() bool                        // property
	SetNullInputAllowed(AValue bool)               // property
	Kind() TDateTimeKind                           // property
	SetKind(AValue TDateTimeKind)                  // property
	TimeSeparator() string                         // property
	SetTimeSeparator(AValue string)                // property
	DecimalSeparator() string                      // property
	SetDecimalSeparator(AValue string)             // property
	TimeFormat() TTimeFormat                       // property
	SetTimeFormat(AValue TTimeFormat)              // property
	TimeDisplay() TTimeDisplay                     // property
	SetTimeDisplay(AValue TTimeDisplay)            // property
	DateMode() TDTDateMode                         // property
	SetDateMode(AValue TDTDateMode)                // property
	Date() TDate                                   // property
	SetDate(AValue TDate)                          // property
	Time() TTime                                   // property
	SetTime(AValue TTime)                          // property
	UseDefaultSeparators() bool                    // property
	SetUseDefaultSeparators(AValue bool)           // property
	Cascade() bool                                 // property
	SetCascade(AValue bool)                        // property
	AutoButtonSize() bool                          // property
	SetAutoButtonSize(AValue bool)                 // property
	AutoAdvance() bool                             // property
	SetAutoAdvance(AValue bool)                    // property
	HideDateTimeParts() TDateTimeParts             // property
	SetHideDateTimeParts(AValue TDateTimeParts)    // property
	MonthDisplay() TMonthDisplay                   // property
	SetMonthDisplay(AValue TMonthDisplay)          // property
	CustomMonthNames() IStrings                    // property
	SetCustomMonthNames(AValue IStrings)           // property
	ShowMonthNames() bool                          // property
	SetShowMonthNames(AValue bool)                 // property
	CalAlignment() TDTCalAlignment                 // property
	SetCalAlignment(AValue TDTCalAlignment)        // property
	Alignment() TAlignment                         // property
	SetAlignment(AValue TAlignment)                // property
	Options() TDateTimePickerOptions               // property
	SetOptions(AValue TDateTimePickerOptions)      // property
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCheckBoxChange(fn TNotifyEvent)           // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
}

// TDateTimePicker Parent: TCustomDateTimePicker
type TDateTimePicker struct {
	TCustomDateTimePicker
	changePtr         uintptr
	checkBoxChangePtr uintptr
	dropDownPtr       uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	editingDonePtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
}

func NewDateTimePicker(AOwner IComponent) IDateTimePicker {
	r1 := LCL().SysCallN(2522, GetObjectUintptr(AOwner))
	return AsDateTimePicker(r1)
}

func (m *TDateTimePicker) DateTime() TDateTime {
	r1 := LCL().SysCallN(2528, 0, m.Instance(), 0)
	return TDateTime(r1)
}

func (m *TDateTimePicker) SetDateTime(AValue TDateTime) {
	LCL().SysCallN(2528, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) DroppedDown() bool {
	r1 := LCL().SysCallN(2530, m.Instance())
	return GoBool(r1)
}

func (m *TDateTimePicker) ArrowShape() TArrowShape {
	r1 := LCL().SysCallN(2514, 0, m.Instance(), 0)
	return TArrowShape(r1)
}

func (m *TDateTimePicker) SetArrowShape(AValue TArrowShape) {
	LCL().SysCallN(2514, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) ShowCheckBox() bool {
	r1 := LCL().SysCallN(2558, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetShowCheckBox(AValue bool) {
	LCL().SysCallN(2558, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) Checked() bool {
	r1 := LCL().SysCallN(2520, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetChecked(AValue bool) {
	LCL().SysCallN(2520, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) CenturyFrom() Word {
	r1 := LCL().SysCallN(2519, 0, m.Instance(), 0)
	return Word(r1)
}

func (m *TDateTimePicker) SetCenturyFrom(AValue Word) {
	LCL().SysCallN(2519, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) DateDisplayOrder() TDateDisplayOrder {
	r1 := LCL().SysCallN(2525, 0, m.Instance(), 0)
	return TDateDisplayOrder(r1)
}

func (m *TDateTimePicker) SetDateDisplayOrder(AValue TDateDisplayOrder) {
	LCL().SysCallN(2525, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) MaxDate() TDate {
	r1 := LCL().SysCallN(2534, 0, m.Instance(), 0)
	return TDate(r1)
}

func (m *TDateTimePicker) SetMaxDate(AValue TDate) {
	LCL().SysCallN(2534, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) MinDate() TDate {
	r1 := LCL().SysCallN(2535, 0, m.Instance(), 0)
	return TDate(r1)
}

func (m *TDateTimePicker) SetMinDate(AValue TDate) {
	LCL().SysCallN(2535, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) ReadOnly() bool {
	r1 := LCL().SysCallN(2542, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetReadOnly(AValue bool) {
	LCL().SysCallN(2542, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) ParentFont() bool {
	r1 := LCL().SysCallN(2540, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetParentFont(AValue bool) {
	LCL().SysCallN(2540, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) ParentColor() bool {
	r1 := LCL().SysCallN(2539, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetParentColor(AValue bool) {
	LCL().SysCallN(2539, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) DateSeparator() string {
	r1 := LCL().SysCallN(2527, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetDateSeparator(AValue string) {
	LCL().SysCallN(2527, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) TrailingSeparator() bool {
	r1 := LCL().SysCallN(2565, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetTrailingSeparator(AValue bool) {
	LCL().SysCallN(2565, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) TextForNullDate() string {
	r1 := LCL().SysCallN(2560, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetTextForNullDate(AValue string) {
	LCL().SysCallN(2560, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) LeadingZeros() bool {
	r1 := LCL().SysCallN(2533, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetLeadingZeros(AValue bool) {
	LCL().SysCallN(2533, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) ParentShowHint() bool {
	r1 := LCL().SysCallN(2541, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2541, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) NullInputAllowed() bool {
	r1 := LCL().SysCallN(2537, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetNullInputAllowed(AValue bool) {
	LCL().SysCallN(2537, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) Kind() TDateTimeKind {
	r1 := LCL().SysCallN(2532, 0, m.Instance(), 0)
	return TDateTimeKind(r1)
}

func (m *TDateTimePicker) SetKind(AValue TDateTimeKind) {
	LCL().SysCallN(2532, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) TimeSeparator() string {
	r1 := LCL().SysCallN(2564, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetTimeSeparator(AValue string) {
	LCL().SysCallN(2564, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) DecimalSeparator() string {
	r1 := LCL().SysCallN(2529, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TDateTimePicker) SetDecimalSeparator(AValue string) {
	LCL().SysCallN(2529, 1, m.Instance(), PascalStr(AValue))
}

func (m *TDateTimePicker) TimeFormat() TTimeFormat {
	r1 := LCL().SysCallN(2563, 0, m.Instance(), 0)
	return TTimeFormat(r1)
}

func (m *TDateTimePicker) SetTimeFormat(AValue TTimeFormat) {
	LCL().SysCallN(2563, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) TimeDisplay() TTimeDisplay {
	r1 := LCL().SysCallN(2562, 0, m.Instance(), 0)
	return TTimeDisplay(r1)
}

func (m *TDateTimePicker) SetTimeDisplay(AValue TTimeDisplay) {
	LCL().SysCallN(2562, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) DateMode() TDTDateMode {
	r1 := LCL().SysCallN(2526, 0, m.Instance(), 0)
	return TDTDateMode(r1)
}

func (m *TDateTimePicker) SetDateMode(AValue TDTDateMode) {
	LCL().SysCallN(2526, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Date() TDate {
	r1 := LCL().SysCallN(2524, 0, m.Instance(), 0)
	return TDate(r1)
}

func (m *TDateTimePicker) SetDate(AValue TDate) {
	LCL().SysCallN(2524, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Time() TTime {
	r1 := LCL().SysCallN(2561, 0, m.Instance(), 0)
	return TTime(r1)
}

func (m *TDateTimePicker) SetTime(AValue TTime) {
	LCL().SysCallN(2561, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) UseDefaultSeparators() bool {
	r1 := LCL().SysCallN(2566, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetUseDefaultSeparators(AValue bool) {
	LCL().SysCallN(2566, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) Cascade() bool {
	r1 := LCL().SysCallN(2518, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetCascade(AValue bool) {
	LCL().SysCallN(2518, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) AutoButtonSize() bool {
	r1 := LCL().SysCallN(2516, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetAutoButtonSize(AValue bool) {
	LCL().SysCallN(2516, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) AutoAdvance() bool {
	r1 := LCL().SysCallN(2515, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetAutoAdvance(AValue bool) {
	LCL().SysCallN(2515, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) HideDateTimeParts() TDateTimeParts {
	r1 := LCL().SysCallN(2531, 0, m.Instance(), 0)
	return TDateTimeParts(r1)
}

func (m *TDateTimePicker) SetHideDateTimeParts(AValue TDateTimeParts) {
	LCL().SysCallN(2531, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) MonthDisplay() TMonthDisplay {
	r1 := LCL().SysCallN(2536, 0, m.Instance(), 0)
	return TMonthDisplay(r1)
}

func (m *TDateTimePicker) SetMonthDisplay(AValue TMonthDisplay) {
	LCL().SysCallN(2536, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) CustomMonthNames() IStrings {
	r1 := LCL().SysCallN(2523, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TDateTimePicker) SetCustomMonthNames(AValue IStrings) {
	LCL().SysCallN(2523, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDateTimePicker) ShowMonthNames() bool {
	r1 := LCL().SysCallN(2559, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDateTimePicker) SetShowMonthNames(AValue bool) {
	LCL().SysCallN(2559, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDateTimePicker) CalAlignment() TDTCalAlignment {
	r1 := LCL().SysCallN(2517, 0, m.Instance(), 0)
	return TDTCalAlignment(r1)
}

func (m *TDateTimePicker) SetCalAlignment(AValue TDTCalAlignment) {
	LCL().SysCallN(2517, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Alignment() TAlignment {
	r1 := LCL().SysCallN(2513, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TDateTimePicker) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(2513, 1, m.Instance(), uintptr(AValue))
}

func (m *TDateTimePicker) Options() TDateTimePickerOptions {
	r1 := LCL().SysCallN(2538, 0, m.Instance(), 0)
	return TDateTimePickerOptions(r1)
}

func (m *TDateTimePicker) SetOptions(AValue TDateTimePickerOptions) {
	LCL().SysCallN(2538, 1, m.Instance(), uintptr(AValue))
}

func DateTimePickerClass() TClass {
	ret := LCL().SysCallN(2521)
	return TClass(ret)
}

func (m *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2543, m.Instance(), m.changePtr)
}

func (m *TDateTimePicker) SetOnCheckBoxChange(fn TNotifyEvent) {
	if m.checkBoxChangePtr != 0 {
		RemoveEventElement(m.checkBoxChangePtr)
	}
	m.checkBoxChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2544, m.Instance(), m.checkBoxChangePtr)
}

func (m *TDateTimePicker) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2548, m.Instance(), m.dropDownPtr)
}

func (m *TDateTimePicker) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2545, m.Instance(), m.closeUpPtr)
}

func (m *TDateTimePicker) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2546, m.Instance(), m.contextPopupPtr)
}

func (m *TDateTimePicker) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2547, m.Instance(), m.dblClickPtr)
}

func (m *TDateTimePicker) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2549, m.Instance(), m.editingDonePtr)
}

func (m *TDateTimePicker) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2550, m.Instance(), m.mouseDownPtr)
}

func (m *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2551, m.Instance(), m.mouseEnterPtr)
}

func (m *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2552, m.Instance(), m.mouseLeavePtr)
}

func (m *TDateTimePicker) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2553, m.Instance(), m.mouseMovePtr)
}

func (m *TDateTimePicker) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2554, m.Instance(), m.mouseUpPtr)
}

func (m *TDateTimePicker) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2555, m.Instance(), m.mouseWheelPtr)
}

func (m *TDateTimePicker) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2556, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TDateTimePicker) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2557, m.Instance(), m.mouseWheelUpPtr)
}
