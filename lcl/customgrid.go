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

// ICustomGrid Parent: ICustomControl
type ICustomGrid interface {
	ICustomControl
	CursorState() TGridCursorState                              // property
	SelectedRange(AIndex int32) (resultGridRect TGridRect)      // property
	SelectedRangeCount() int32                                  // property
	SortOrder() TSortOrder                                      // property
	SetSortOrder(AValue TSortOrder)                             // property
	SortColumn() int32                                          // property
	CellRect(ACol, ARow int32) (resultRect TRect)               // function
	CellToGridZone(aCol, aRow int32) TGridZone                  // function
	ClearCols() bool                                            // function
	ClearRows() bool                                            // function
	EditorByStyle(Style TColumnButtonStyle) IWinControl         // function
	HasMultiSelection() bool                                    // function
	IsCellVisible(aCol, aRow int32) bool                        // function
	IsFixedCellVisible(aCol, aRow int32) bool                   // function
	MouseCoord(X, Y int32) (resultGridCoord TGridCoord)         // function
	MouseToCell(Mouse *TPoint) (resultPoint TPoint)             // function
	MouseToLogcell(Mouse *TPoint) (resultPoint TPoint)          // function
	MouseToGridZone(X, Y int32) TGridZone                       // function
	AdjustInnerCellRect(ARect *TRect)                           // procedure
	AutoAdjustColumns()                                         // procedure
	BeginUpdate()                                               // procedure
	CheckPosition()                                             // procedure
	Clear()                                                     // procedure
	ClearSelections()                                           // procedure
	EditorKeyDown(Sender IObject, Key *Word, Shift TShiftState) // procedure
	EditorKeyPress(Sender IObject, Key *Char)                   // procedure
	EditorUTF8KeyPress(Sender IObject, UTF8Key *TUTF8Char)      // procedure
	EditorKeyUp(Sender IObject, key *Word, shift TShiftState)   // procedure
	EditorTextChanged(aCol, aRow int32, aText string)           // procedure
	EndUpdate(aRefresh bool)                                    // procedure
	HideSortArrow()                                             // procedure
	InvalidateCell(aCol, aRow int32)                            // procedure
	InvalidateCol(ACol int32)                                   // procedure
	InvalidateRange(aRange *TRect)                              // procedure
	InvalidateRow(ARow int32)                                   // procedure
	LoadFromFile(FileName string)                               // procedure
	LoadFromStream(AStream IStream)                             // procedure
	MouseToCell1(X, Y int32, OutCol, OutRow *int32)             // procedure
	SaveToFile(FileName string)                                 // procedure
	SaveToStream(AStream IStream)                               // procedure
}

// TCustomGrid Parent: TCustomControl
type TCustomGrid struct {
	TCustomControl
}

func NewCustomGrid(AOwner IComponent) ICustomGrid {
	r1 := LCL().SysCallN(1748, GetObjectUintptr(AOwner))
	return AsCustomGrid(r1)
}

func (m *TCustomGrid) CursorState() TGridCursorState {
	r1 := LCL().SysCallN(1749, m.Instance())
	return TGridCursorState(r1)
}

func (m *TCustomGrid) SelectedRange(AIndex int32) (resultGridRect TGridRect) {
	LCL().SysCallN(1774, m.Instance(), uintptr(AIndex), uintptr(unsafePointer(&resultGridRect)))
	return
}

func (m *TCustomGrid) SelectedRangeCount() int32 {
	r1 := LCL().SysCallN(1775, m.Instance())
	return int32(r1)
}

func (m *TCustomGrid) SortOrder() TSortOrder {
	r1 := LCL().SysCallN(1777, 0, m.Instance(), 0)
	return TSortOrder(r1)
}

func (m *TCustomGrid) SetSortOrder(AValue TSortOrder) {
	LCL().SysCallN(1777, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomGrid) SortColumn() int32 {
	r1 := LCL().SysCallN(1776, m.Instance())
	return int32(r1)
}

func (m *TCustomGrid) CellRect(ACol, ARow int32) (resultRect TRect) {
	LCL().SysCallN(1740, m.Instance(), uintptr(ACol), uintptr(ARow), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomGrid) CellToGridZone(aCol, aRow int32) TGridZone {
	r1 := LCL().SysCallN(1741, m.Instance(), uintptr(aCol), uintptr(aRow))
	return TGridZone(r1)
}

func (m *TCustomGrid) ClearCols() bool {
	r1 := LCL().SysCallN(1745, m.Instance())
	return GoBool(r1)
}

func (m *TCustomGrid) ClearRows() bool {
	r1 := LCL().SysCallN(1746, m.Instance())
	return GoBool(r1)
}

func (m *TCustomGrid) EditorByStyle(Style TColumnButtonStyle) IWinControl {
	r1 := LCL().SysCallN(1750, m.Instance(), uintptr(Style))
	return AsWinControl(r1)
}

func (m *TCustomGrid) HasMultiSelection() bool {
	r1 := LCL().SysCallN(1757, m.Instance())
	return GoBool(r1)
}

func (m *TCustomGrid) IsCellVisible(aCol, aRow int32) bool {
	r1 := LCL().SysCallN(1763, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomGrid) IsFixedCellVisible(aCol, aRow int32) bool {
	r1 := LCL().SysCallN(1764, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomGrid) MouseCoord(X, Y int32) (resultGridCoord TGridCoord) {
	LCL().SysCallN(1767, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&resultGridCoord)))
	return
}

func (m *TCustomGrid) MouseToCell(Mouse *TPoint) (resultPoint TPoint) {
	LCL().SysCallN(1768, m.Instance(), uintptr(unsafePointer(Mouse)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomGrid) MouseToLogcell(Mouse *TPoint) (resultPoint TPoint) {
	LCL().SysCallN(1771, m.Instance(), uintptr(unsafePointer(Mouse)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomGrid) MouseToGridZone(X, Y int32) TGridZone {
	r1 := LCL().SysCallN(1770, m.Instance(), uintptr(X), uintptr(Y))
	return TGridZone(r1)
}

func CustomGridClass() TClass {
	ret := LCL().SysCallN(1743)
	return TClass(ret)
}

func (m *TCustomGrid) AdjustInnerCellRect(ARect *TRect) {
	var result0 uintptr
	LCL().SysCallN(1737, m.Instance(), uintptr(unsafePointer(&result0)))
	*ARect = *(*TRect)(getPointer(result0))
}

func (m *TCustomGrid) AutoAdjustColumns() {
	LCL().SysCallN(1738, m.Instance())
}

func (m *TCustomGrid) BeginUpdate() {
	LCL().SysCallN(1739, m.Instance())
}

func (m *TCustomGrid) CheckPosition() {
	LCL().SysCallN(1742, m.Instance())
}

func (m *TCustomGrid) Clear() {
	LCL().SysCallN(1744, m.Instance())
}

func (m *TCustomGrid) ClearSelections() {
	LCL().SysCallN(1747, m.Instance())
}

func (m *TCustomGrid) EditorKeyDown(Sender IObject, Key *Word, Shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(1751, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(Shift))
	*Key = Word(result1)
}

func (m *TCustomGrid) EditorKeyPress(Sender IObject, Key *Char) {
	var result1 uintptr
	LCL().SysCallN(1752, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*Key = Char(result1)
}

func (m *TCustomGrid) EditorUTF8KeyPress(Sender IObject, UTF8Key *TUTF8Char) {
	var result1 uintptr
	LCL().SysCallN(1755, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)))
	*UTF8Key = *(*TUTF8Char)(getPointer(result1))
}

func (m *TCustomGrid) EditorKeyUp(Sender IObject, key *Word, shift TShiftState) {
	var result1 uintptr
	LCL().SysCallN(1753, m.Instance(), GetObjectUintptr(Sender), uintptr(unsafePointer(&result1)), uintptr(shift))
	*key = Word(result1)
}

func (m *TCustomGrid) EditorTextChanged(aCol, aRow int32, aText string) {
	LCL().SysCallN(1754, m.Instance(), uintptr(aCol), uintptr(aRow), PascalStr(aText))
}

func (m *TCustomGrid) EndUpdate(aRefresh bool) {
	LCL().SysCallN(1756, m.Instance(), PascalBool(aRefresh))
}

func (m *TCustomGrid) HideSortArrow() {
	LCL().SysCallN(1758, m.Instance())
}

func (m *TCustomGrid) InvalidateCell(aCol, aRow int32) {
	LCL().SysCallN(1759, m.Instance(), uintptr(aCol), uintptr(aRow))
}

func (m *TCustomGrid) InvalidateCol(ACol int32) {
	LCL().SysCallN(1760, m.Instance(), uintptr(ACol))
}

func (m *TCustomGrid) InvalidateRange(aRange *TRect) {
	LCL().SysCallN(1761, m.Instance(), uintptr(unsafePointer(aRange)))
}

func (m *TCustomGrid) InvalidateRow(ARow int32) {
	LCL().SysCallN(1762, m.Instance(), uintptr(ARow))
}

func (m *TCustomGrid) LoadFromFile(FileName string) {
	LCL().SysCallN(1765, m.Instance(), PascalStr(FileName))
}

func (m *TCustomGrid) LoadFromStream(AStream IStream) {
	LCL().SysCallN(1766, m.Instance(), GetObjectUintptr(AStream))
}

func (m *TCustomGrid) MouseToCell1(X, Y int32, OutCol, OutRow *int32) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(1769, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutCol = int32(result1)
	*OutRow = int32(result2)
}

func (m *TCustomGrid) SaveToFile(FileName string) {
	LCL().SysCallN(1772, m.Instance(), PascalStr(FileName))
}

func (m *TCustomGrid) SaveToStream(AStream IStream) {
	LCL().SysCallN(1773, m.Instance(), GetObjectUintptr(AStream))
}
