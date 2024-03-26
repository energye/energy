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
	r1 := LCL().SysCallN(1809, GetObjectUintptr(TheOwner))
	return AsCustomListBox(r1)
}

func (m *TCustomListBox) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1800, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomListBox) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1800, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Canvas() ICanvas {
	r1 := LCL().SysCallN(1801, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomListBox) ClickOnSelChange() bool {
	r1 := LCL().SysCallN(1806, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetClickOnSelChange(AValue bool) {
	LCL().SysCallN(1806, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Columns() int32 {
	r1 := LCL().SysCallN(1807, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetColumns(AValue int32) {
	LCL().SysCallN(1807, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Count() int32 {
	r1 := LCL().SysCallN(1808, m.Instance())
	return int32(r1)
}

func (m *TCustomListBox) ExtendedSelect() bool {
	r1 := LCL().SysCallN(1811, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetExtendedSelect(AValue bool) {
	LCL().SysCallN(1811, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) IntegralHeight() bool {
	r1 := LCL().SysCallN(1815, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetIntegralHeight(AValue bool) {
	LCL().SysCallN(1815, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ItemHeight() int32 {
	r1 := LCL().SysCallN(1818, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetItemHeight(AValue int32) {
	LCL().SysCallN(1818, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) ItemIndex() int32 {
	r1 := LCL().SysCallN(1819, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetItemIndex(AValue int32) {
	LCL().SysCallN(1819, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) Items() IStrings {
	r1 := LCL().SysCallN(1822, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TCustomListBox) SetItems(AValue IStrings) {
	LCL().SysCallN(1822, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListBox) MultiSelect() bool {
	r1 := LCL().SysCallN(1826, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetMultiSelect(AValue bool) {
	LCL().SysCallN(1826, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Options() TListBoxOptions {
	r1 := LCL().SysCallN(1827, 0, m.Instance(), 0)
	return TListBoxOptions(r1)
}

func (m *TCustomListBox) SetOptions(AValue TListBoxOptions) {
	LCL().SysCallN(1827, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) ParentColor() bool {
	r1 := LCL().SysCallN(1828, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentColor(AValue bool) {
	LCL().SysCallN(1828, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ParentFont() bool {
	r1 := LCL().SysCallN(1829, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentFont(AValue bool) {
	LCL().SysCallN(1829, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(1830, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(1830, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) ScrollWidth() int32 {
	r1 := LCL().SysCallN(1831, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetScrollWidth(AValue int32) {
	LCL().SysCallN(1831, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) SelCount() int32 {
	r1 := LCL().SysCallN(1832, m.Instance())
	return int32(r1)
}

func (m *TCustomListBox) Selected(Index int32) bool {
	r1 := LCL().SysCallN(1835, 0, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomListBox) SetSelected(Index int32, AValue bool) {
	LCL().SysCallN(1835, 1, m.Instance(), uintptr(Index), PascalBool(AValue))
}

func (m *TCustomListBox) Sorted() bool {
	r1 := LCL().SysCallN(1848, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListBox) SetSorted(AValue bool) {
	LCL().SysCallN(1848, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListBox) Style() TListBoxStyle {
	r1 := LCL().SysCallN(1849, 0, m.Instance(), 0)
	return TListBoxStyle(r1)
}

func (m *TCustomListBox) SetStyle(AValue TListBoxStyle) {
	LCL().SysCallN(1849, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) TopIndex() int32 {
	r1 := LCL().SysCallN(1850, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListBox) SetTopIndex(AValue int32) {
	LCL().SysCallN(1850, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListBox) GetIndexAtXY(X, Y int32) int32 {
	r1 := LCL().SysCallN(1812, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TCustomListBox) GetIndexAtY(Y int32) int32 {
	r1 := LCL().SysCallN(1813, m.Instance(), uintptr(Y))
	return int32(r1)
}

func (m *TCustomListBox) GetSelectedText() string {
	r1 := LCL().SysCallN(1814, m.Instance())
	return GoStr(r1)
}

func (m *TCustomListBox) ItemAtPos(Pos *TPoint, Existing bool) int32 {
	r1 := LCL().SysCallN(1816, m.Instance(), uintptr(unsafe.Pointer(Pos)), PascalBool(Existing))
	return int32(r1)
}

func (m *TCustomListBox) ItemRect(Index int32) (resultRect TRect) {
	LCL().SysCallN(1820, m.Instance(), uintptr(Index), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TCustomListBox) ItemVisible(Index int32) bool {
	r1 := LCL().SysCallN(1821, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func (m *TCustomListBox) ItemFullyVisible(Index int32) bool {
	r1 := LCL().SysCallN(1817, m.Instance(), uintptr(Index))
	return GoBool(r1)
}

func CustomListBoxClass() TClass {
	ret := LCL().SysCallN(1802)
	return TClass(ret)
}

func (m *TCustomListBox) AddItem(Item string, AnObject IObject) {
	LCL().SysCallN(1799, m.Instance(), PascalStr(Item), GetObjectUintptr(AnObject))
}

func (m *TCustomListBox) Click() {
	LCL().SysCallN(1805, m.Instance())
}

func (m *TCustomListBox) Clear() {
	LCL().SysCallN(1803, m.Instance())
}

func (m *TCustomListBox) ClearSelection() {
	LCL().SysCallN(1804, m.Instance())
}

func (m *TCustomListBox) LockSelectionChange() {
	LCL().SysCallN(1823, m.Instance())
}

func (m *TCustomListBox) MakeCurrentVisible() {
	LCL().SysCallN(1824, m.Instance())
}

func (m *TCustomListBox) MeasureItem(Index int32, TheHeight *int32) {
	var result1 uintptr
	LCL().SysCallN(1825, m.Instance(), uintptr(Index), uintptr(unsafe.Pointer(&result1)))
	*TheHeight = int32(result1)
}

func (m *TCustomListBox) SelectAll() {
	LCL().SysCallN(1833, m.Instance())
}

func (m *TCustomListBox) SelectRange(ALow, AHigh int32, ASelected bool) {
	LCL().SysCallN(1834, m.Instance(), uintptr(ALow), uintptr(AHigh), PascalBool(ASelected))
}

func (m *TCustomListBox) DeleteSelected() {
	LCL().SysCallN(1810, m.Instance())
}

func (m *TCustomListBox) UnlockSelectionChange() {
	LCL().SysCallN(1851, m.Instance())
}

func (m *TCustomListBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1836, m.Instance(), m.dblClickPtr)
}

func (m *TCustomListBox) SetOnDrawItem(fn TDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1837, m.Instance(), m.drawItemPtr)
}

func (m *TCustomListBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1838, m.Instance(), m.measureItemPtr)
}

func (m *TCustomListBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1839, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomListBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1840, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomListBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1841, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomListBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1842, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomListBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1843, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomListBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1844, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomListBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1845, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomListBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1846, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomListBox) SetOnSelectionChange(fn TSelectionChangeEvent) {
	if m.selectionChangePtr != 0 {
		RemoveEventElement(m.selectionChangePtr)
	}
	m.selectionChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1847, m.Instance(), m.selectionChangePtr)
}
