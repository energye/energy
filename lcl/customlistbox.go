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

// ICustomListBox Parent: IWinControl
type ICustomListBox interface {
	IWinControl
	BorderStyle() TBorderStyle                     // property
	SetBorderStyle(AValue TBorderStyle)            // property
	Canvas() ICanvas                               // property
	ClickOnSelChange() bool                        // property
	SetClickOnSelChange(AValue bool)               // property
	Columns() int32                                // property
	SetColumns(AValue int32)                       // property
	Count() int32                                  // property
	ExtendedSelect() bool                          // property
	SetExtendedSelect(AValue bool)                 // property
	IntegralHeight() bool                          // property
	SetIntegralHeight(AValue bool)                 // property
	ItemHeight() int32                             // property
	SetItemHeight(AValue int32)                    // property
	ItemIndex() int32                              // property
	SetItemIndex(AValue int32)                     // property
	Items() IStrings                               // property
	SetItems(AValue IStrings)                      // property
	MultiSelect() bool                             // property
	SetMultiSelect(AValue bool)                    // property
	Options() TListBoxOptions                      // property
	SetOptions(AValue TListBoxOptions)             // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	ScrollWidth() int32                            // property
	SetScrollWidth(AValue int32)                   // property
	SelCount() int32                               // property
	Selected(Index int32) bool                     // property
	SetSelected(Index int32, AValue bool)          // property
	Sorted() bool                                  // property
	SetSorted(AValue bool)                         // property
	Style() TListBoxStyle                          // property
	SetStyle(AValue TListBoxStyle)                 // property
	TopIndex() int32                               // property
	SetTopIndex(AValue int32)                      // property
	GetIndexAtXY(X, Y int32) int32                 // function
	GetIndexAtY(Y int32) int32                     // function
	GetSelectedText() string                       // function
	ItemAtPos(Pos *TPoint, Existing bool) int32    // function
	ItemRect(Index int32) (resultRect TRect)       // function
	ItemVisible(Index int32) bool                  // function
	ItemFullyVisible(Index int32) bool             // function
	AddItem(Item string, AnObject IObject)         // procedure
	Click()                                        // procedure
	Clear()                                        // procedure
	ClearSelection()                               // procedure
	LockSelectionChange()                          // procedure
	MakeCurrentVisible()                           // procedure
	MeasureItem(Index int32, TheHeight *int32)     // procedure
	SelectAll()                                    // procedure
	SelectRange(ALow, AHigh int32, ASelected bool) // procedure
	DeleteSelected()                               // procedure
	UnlockSelectionChange()                        // procedure
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDrawItem(fn TDrawItemEvent)               // property event
	SetOnMeasureItem(fn TMeasureItemEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelectionChange(fn TSelectionChangeEvent) // property event
}

// TCustomListBox Parent: TWinControl
type TCustomListBox struct {
	TWinControl
	dblClickPtr        uintptr
	drawItemPtr        uintptr
	measureItemPtr     uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	selectionChangePtr uintptr
}

func NewCustomListBox(TheOwner IComponent) ICustomListBox {
	r1 := LCL().SysCallN(1999, GetObjectUintptr(TheOwner))
	return AsCustomListBox(r1)
}

func (m *TCustomListBox) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1990, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomListBox) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1990, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Canvas() ICanvas {
	r1 := LCL().SysCallN(1991, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomListBox) ClickOnSelChange() bool {
	r1 := LCL().SysCallN(1996, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetClickOnSelChange(AValue bool) {
	LCL().SysCallN(1996, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Columns() int32 {
	r1 := LCL().SysCallN(1997, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetColumns(AValue int32) {
	LCL().SysCallN(1997, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Count() int32 {
	r1 := LCL().SysCallN(1998, m.Instance())
	return int32(r1)
}

func (m *TCustomListBox) ExtendedSelect() bool {
	r1 := LCL().SysCallN(2001, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetExtendedSelect(AValue bool) {
	LCL().SysCallN(2001, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) IntegralHeight() bool {
	r1 := LCL().SysCallN(2005, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetIntegralHeight(AValue bool) {
	LCL().SysCallN(2005, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ItemHeight() int32 {
	r1 := LCL().SysCallN(2008, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetItemHeight(AValue int32) {
	LCL().SysCallN(2008, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) ItemIndex() int32 {
	r1 := LCL().SysCallN(2009, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetItemIndex(AValue int32) {
	LCL().SysCallN(2009, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Items() IStrings {
	r1 := LCL().SysCallN(2012, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomListBox) SetItems(AValue IStrings) {
	LCL().SysCallN(2012, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListBox) MultiSelect() bool {
	r1 := LCL().SysCallN(2016, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetMultiSelect(AValue bool) {
	LCL().SysCallN(2016, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Options() TListBoxOptions {
	r1 := LCL().SysCallN(2017, 0, m.Instance(), 0)
	return TListBoxOptions(r1)
}

func (m *TCustomListBox) SetOptions(AValue TListBoxOptions) {
	LCL().SysCallN(2017, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) ParentColor() bool {
	r1 := LCL().SysCallN(2018, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentColor(AValue bool) {
	LCL().SysCallN(2018, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ParentFont() bool {
	r1 := LCL().SysCallN(2019, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentFont(AValue bool) {
	LCL().SysCallN(2019, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(2020, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2020, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ScrollWidth() int32 {
	r1 := LCL().SysCallN(2021, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetScrollWidth(AValue int32) {
	LCL().SysCallN(2021, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) SelCount() int32 {
	r1 := LCL().SysCallN(2022, m.Instance())
	return int32(r1)
}

func (m *TCustomListBox) Selected(Index int32) bool {
	r1 := LCL().SysCallN(2025, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomListBox) SetSelected(Index int32, AValue bool) {
	LCL().SysCallN(2025, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomListBox) Sorted() bool {
	r1 := LCL().SysCallN(2038, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetSorted(AValue bool) {
	LCL().SysCallN(2038, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Style() TListBoxStyle {
	r1 := LCL().SysCallN(2039, 0, m.Instance(), 0)
	return TListBoxStyle(r1)
}

func (m *TCustomListBox) SetStyle(AValue TListBoxStyle) {
	LCL().SysCallN(2039, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) TopIndex() int32 {
	r1 := LCL().SysCallN(2040, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetTopIndex(AValue int32) {
	LCL().SysCallN(2040, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) GetIndexAtXY(X, Y int32) int32 {
	r1 := LCL().SysCallN(2002, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TCustomListBox) GetIndexAtY(Y int32) int32 {
	r1 := LCL().SysCallN(2003, m.Instance(), uintptr(Y))
	return int32(r1)
}

func (m *TCustomListBox) GetSelectedText() string {
	r1 := LCL().SysCallN(2004, m.Instance())
	return GoStr(r1)
}

func (m *TCustomListBox) ItemAtPos(Pos *TPoint, Existing bool) int32 {
	r1 := LCL().SysCallN(2006, m.Instance(), uintptr(unsafePointer(Pos)), PascalBool(Existing))
	return int32(r1)
}

func (m *TCustomListBox) ItemRect(Index int32) (resultRect TRect) {
	LCL().SysCallN(2010, m.Instance(), uintptr(Index), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomListBox) ItemVisible(Index int32) bool {
	r1 := LCL().SysCallN(2011, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomListBox) ItemFullyVisible(Index int32) bool {
	r1 := LCL().SysCallN(2007, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func CustomListBoxClass() TClass {
	ret := LCL().SysCallN(1992)
	return TClass(ret)
}

func (m *TCustomListBox) AddItem(Item string, AnObject IObject) {
	LCL().SysCallN(1989, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject))
}

func (m *TCustomListBox) Click() {
	LCL().SysCallN(1995, m.Instance())
}

func (m *TCustomListBox) Clear() {
	LCL().SysCallN(1993, m.Instance())
}

func (m *TCustomListBox) ClearSelection() {
	LCL().SysCallN(1994, m.Instance())
}

func (m *TCustomListBox) LockSelectionChange() {
	LCL().SysCallN(2013, m.Instance())
}

func (m *TCustomListBox) MakeCurrentVisible() {
	LCL().SysCallN(2014, m.Instance())
}

func (m *TCustomListBox) MeasureItem(Index int32, TheHeight *int32) {
	var result1 uintptr
	LCL().SysCallN(2015, m.Instance(), uintptr(Index), uintptr(unsafePointer(&result1)))
	*TheHeight = int32(result1)
}

func (m *TCustomListBox) SelectAll() {
	LCL().SysCallN(2023, m.Instance())
}

func (m *TCustomListBox) SelectRange(ALow, AHigh int32, ASelected bool) {
	LCL().SysCallN(2024, m.Instance(), uintptr(ALow), uintptr(AHigh), PascalBool(ASelected))
}

func (m *TCustomListBox) DeleteSelected() {
	LCL().SysCallN(2000, m.Instance())
}

func (m *TCustomListBox) UnlockSelectionChange() {
	LCL().SysCallN(2041, m.Instance())
}

func (m *TCustomListBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2026, m.Instance(), m.dblClickPtr)
}

func (m *TCustomListBox) SetOnDrawItem(fn TDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2027, m.Instance(), m.drawItemPtr)
}

func (m *TCustomListBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2028, m.Instance(), m.measureItemPtr)
}

func (m *TCustomListBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2029, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomListBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2030, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomListBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2031, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomListBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2032, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomListBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2033, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomListBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2034, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomListBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2035, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomListBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2036, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomListBox) SetOnSelectionChange(fn TSelectionChangeEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2037, m.Instance(), m.selectionChangePtr)
}
