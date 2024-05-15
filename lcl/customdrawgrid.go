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

// ICustomDrawGrid Parent: ICustomGrid
type ICustomDrawGrid interface {
	ICustomGrid
	AllowOutboundEvents() bool                                             // property
	SetAllowOutboundEvents(AValue bool)                                    // property
	BorderColor() TColor                                                   // property
	SetBorderColor(AValue TColor)                                          // property
	Col() int32                                                            // property
	SetCol(AValue int32)                                                   // property
	ColWidths(aCol int32) int32                                            // property
	SetColWidths(aCol int32, AValue int32)                                 // property
	ColRow() (resultPoint TPoint)                                          // property
	SetColRow(AValue *TPoint)                                              // property
	DisabledFontColor() TColor                                             // property
	SetDisabledFontColor(AValue TColor)                                    // property
	Editor() IWinControl                                                   // property
	SetEditor(AValue IWinControl)                                          // property
	EditorBorderStyle() TBorderStyle                                       // property
	SetEditorBorderStyle(AValue TBorderStyle)                              // property
	EditorMode() bool                                                      // property
	SetEditorMode(AValue bool)                                             // property
	ExtendedColSizing() bool                                               // property
	SetExtendedColSizing(AValue bool)                                      // property
	AltColorStartNormal() bool                                             // property
	SetAltColorStartNormal(AValue bool)                                    // property
	FastEditing() bool                                                     // property
	SetFastEditing(AValue bool)                                            // property
	FixedGridLineColor() TColor                                            // property
	SetFixedGridLineColor(AValue TColor)                                   // property
	FocusColor() TColor                                                    // property
	SetFocusColor(AValue TColor)                                           // property
	FocusRectVisible() bool                                                // property
	SetFocusRectVisible(AValue bool)                                       // property
	GridHeight() int32                                                     // property
	GridWidth() int32                                                      // property
	IsCellSelected(aCol, aRow int32) bool                                  // property
	LeftCol() int32                                                        // property
	SetLeftCol(AValue int32)                                               // property
	Row() int32                                                            // property
	SetRow(AValue int32)                                                   // property
	RowHeights(aRow int32) int32                                           // property
	SetRowHeights(aRow int32, AValue int32)                                // property
	SaveOptions() TSaveOptions                                             // property
	SetSaveOptions(AValue TSaveOptions)                                    // property
	SelectedColor() TColor                                                 // property
	SetSelectedColor(AValue TColor)                                        // property
	SelectedColumn() IGridColumn                                           // property
	Selection() (resultGridRect TGridRect)                                 // property
	SetSelection(AValue *TGridRect)                                        // property
	StrictSort() bool                                                      // property
	SetStrictSort(AValue bool)                                             // property
	TopRow() int32                                                         // property
	SetTopRow(AValue int32)                                                // property
	UseXORFeatures() bool                                                  // property
	SetUseXORFeatures(AValue bool)                                         // property
	AutoAdvance() TAutoAdvance                                             // property
	SetAutoAdvance(AValue TAutoAdvance)                                    // property
	AutoFillColumns() bool                                                 // property
	SetAutoFillColumns(AValue bool)                                        // property
	ColCount() int32                                                       // property
	SetColCount(AValue int32)                                              // property
	Columns() IGridColumns                                                 // property
	SetColumns(AValue IGridColumns)                                        // property
	DefaultColWidth() int32                                                // property
	SetDefaultColWidth(AValue int32)                                       // property
	DefaultDrawing() bool                                                  // property
	SetDefaultDrawing(AValue bool)                                         // property
	DefaultRowHeight() int32                                               // property
	SetDefaultRowHeight(AValue int32)                                      // property
	FadeUnfocusedSelection() bool                                          // property
	SetFadeUnfocusedSelection(AValue bool)                                 // property
	FixedColor() TColor                                                    // property
	SetFixedColor(AValue TColor)                                           // property
	FixedCols() int32                                                      // property
	SetFixedCols(AValue int32)                                             // property
	FixedHotColor() TColor                                                 // property
	SetFixedHotColor(AValue TColor)                                        // property
	FixedRows() int32                                                      // property
	SetFixedRows(AValue int32)                                             // property
	Flat() bool                                                            // property
	SetFlat(AValue bool)                                                   // property
	GridLineColor() TColor                                                 // property
	SetGridLineColor(AValue TColor)                                        // property
	GridLineStyle() TPenStyle                                              // property
	SetGridLineStyle(AValue TPenStyle)                                     // property
	GridLineWidth() int32                                                  // property
	SetGridLineWidth(AValue int32)                                         // property
	Options() TGridOptions                                                 // property
	SetOptions(AValue TGridOptions)                                        // property
	Options2() TGridOptions2                                               // property
	SetOptions2(AValue TGridOptions2)                                      // property
	ParentShowHint() bool                                                  // property
	SetParentShowHint(AValue bool)                                         // property
	RowCount() int32                                                       // property
	SetRowCount(AValue int32)                                              // property
	ScrollBars() TScrollStyle                                              // property
	SetScrollBars(AValue TScrollStyle)                                     // property
	TabAdvance() TAutoAdvance                                              // property
	SetTabAdvance(AValue TAutoAdvance)                                     // property
	VisibleColCount() int32                                                // property
	VisibleRowCount() int32                                                // property
	DeleteColRow(IsColumn bool, index int32)                               // procedure
	DeleteCol(Index int32)                                                 // procedure
	DeleteRow(Index int32)                                                 // procedure
	ExchangeColRow(IsColumn bool, index, WithIndex int32)                  // procedure
	InsertColRow(IsColumn bool, index int32)                               // procedure
	MoveColRow(IsColumn bool, FromIndex, ToIndex int32)                    // procedure
	SortColRow(IsColumn bool, index int32)                                 // procedure
	SortColRow1(IsColumn bool, Index, FromIndex, ToIndex int32)            // procedure
	DefaultDrawCell(aCol, aRow int32, aRect *TRect, aState TGridDrawState) // procedure
	SetOnAfterSelection(fn TOnSelectEvent)                                 // property event
	SetOnBeforeSelection(fn TOnSelectEvent)                                // property event
	SetOnColRowDeleted(fn TGridOperationEvent)                             // property event
	SetOnColRowExchanged(fn TGridOperationEvent)                           // property event
	SetOnColRowInserted(fn TGridOperationEvent)                            // property event
	SetOnColRowMoved(fn TGridOperationEvent)                               // property event
	SetOnCompareCells(fn TOnCompareCells)                                  // property event
	SetOnContextPopup(fn TContextPopupEvent)                               // property event
	SetOnDblClick(fn TNotifyEvent)                                         // property event
	SetOnDragDrop(fn TDragDropEvent)                                       // property event
	SetOnDragOver(fn TDragOverEvent)                                       // property event
	SetOnDrawCell(fn TOnDrawCell)                                          // property event
	SetOnButtonClick(fn TOnSelectEvent)                                    // property event
	SetOnEndDock(fn TEndDragEvent)                                         // property event
	SetOnEndDrag(fn TEndDragEvent)                                         // property event
	SetOnGetEditMask(fn TGetEditEvent)                                     // property event
	SetOnGetEditText(fn TGetEditEvent)                                     // property event
	SetOnHeaderClick(fn THdrEvent)                                         // property event
	SetOnHeaderSized(fn THdrEvent)                                         // property event
	SetOnHeaderSizing(fn THeaderSizingEvent)                               // property event
	SetOnMouseDown(fn TMouseEvent)                                         // property event
	SetOnMouseEnter(fn TNotifyEvent)                                       // property event
	SetOnMouseLeave(fn TNotifyEvent)                                       // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                     // property event
	SetOnMouseUp(fn TMouseEvent)                                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                   // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                         // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                           // property event
	SetOnPickListSelect(fn TNotifyEvent)                                   // property event
	SetOnPrepareCanvas(fn TOnPrepareCanvasEvent)                           // property event
	SetOnSelectEditor(fn TSelectEditorEvent)                               // property event
	SetOnSelection(fn TOnSelectEvent)                                      // property event
	SetOnSelectCell(fn TOnSelectCellEvent)                                 // property event
	SetOnSetEditText(fn TSetEditEvent)                                     // property event
	SetOnStartDock(fn TStartDockEvent)                                     // property event
	SetOnStartDrag(fn TStartDragEvent)                                     // property event
	SetOnTopleftChanged(fn TNotifyEvent)                                   // property event
	SetOnValidateEntry(fn TValidateEntryEvent)                             // property event
}

// TCustomDrawGrid Parent: TCustomGrid
type TCustomDrawGrid struct {
	TCustomGrid
	afterSelectionPtr  uintptr
	beforeSelectionPtr uintptr
	colRowDeletedPtr   uintptr
	colRowExchangedPtr uintptr
	colRowInsertedPtr  uintptr
	colRowMovedPtr     uintptr
	compareCellsPtr    uintptr
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	drawCellPtr        uintptr
	buttonClickPtr     uintptr
	endDockPtr         uintptr
	endDragPtr         uintptr
	getEditMaskPtr     uintptr
	getEditTextPtr     uintptr
	headerClickPtr     uintptr
	headerSizedPtr     uintptr
	headerSizingPtr    uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	pickListSelectPtr  uintptr
	prepareCanvasPtr   uintptr
	selectEditorPtr    uintptr
	selectionPtr       uintptr
	selectCellPtr      uintptr
	setEditTextPtr     uintptr
	startDockPtr       uintptr
	startDragPtr       uintptr
	topleftChangedPtr  uintptr
	validateEntryPtr   uintptr
}

func NewCustomDrawGrid(AOwner IComponent) ICustomDrawGrid {
	r1 := LCL().SysCallN(1519, GetObjectUintptr(AOwner))
	return AsCustomDrawGrid(r1)
}

func (m *TCustomDrawGrid) AllowOutboundEvents() bool {
	r1 := LCL().SysCallN(1508, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAllowOutboundEvents(AValue bool) {
	LCL().SysCallN(1508, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) BorderColor() TColor {
	r1 := LCL().SysCallN(1512, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetBorderColor(AValue TColor) {
	LCL().SysCallN(1512, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Col() int32 {
	r1 := LCL().SysCallN(1514, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetCol(AValue int32) {
	LCL().SysCallN(1514, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ColWidths(aCol int32) int32 {
	r1 := LCL().SysCallN(1517, 0, m.Instance(), uintptr(aCol))
	return int32(r1)
}

func (m *TCustomDrawGrid) SetColWidths(aCol int32, AValue int32) {
	LCL().SysCallN(1517, 1, m.Instance(), uintptr(aCol), uintptr(AValue))
}

func (m *TCustomDrawGrid) ColRow() (resultPoint TPoint) {
	LCL().SysCallN(1516, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomDrawGrid) SetColRow(AValue *TPoint) {
	LCL().SysCallN(1516, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomDrawGrid) DisabledFontColor() TColor {
	r1 := LCL().SysCallN(1527, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetDisabledFontColor(AValue TColor) {
	LCL().SysCallN(1527, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Editor() IWinControl {
	r1 := LCL().SysCallN(1528, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomDrawGrid) SetEditor(AValue IWinControl) {
	LCL().SysCallN(1528, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomDrawGrid) EditorBorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1529, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomDrawGrid) SetEditorBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1529, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) EditorMode() bool {
	r1 := LCL().SysCallN(1530, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetEditorMode(AValue bool) {
	LCL().SysCallN(1530, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) ExtendedColSizing() bool {
	r1 := LCL().SysCallN(1532, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetExtendedColSizing(AValue bool) {
	LCL().SysCallN(1532, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) AltColorStartNormal() bool {
	r1 := LCL().SysCallN(1509, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAltColorStartNormal(AValue bool) {
	LCL().SysCallN(1509, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FastEditing() bool {
	r1 := LCL().SysCallN(1534, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFastEditing(AValue bool) {
	LCL().SysCallN(1534, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FixedGridLineColor() TColor {
	r1 := LCL().SysCallN(1537, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedGridLineColor(AValue TColor) {
	LCL().SysCallN(1537, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FocusColor() TColor {
	r1 := LCL().SysCallN(1541, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFocusColor(AValue TColor) {
	LCL().SysCallN(1541, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FocusRectVisible() bool {
	r1 := LCL().SysCallN(1542, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFocusRectVisible(AValue bool) {
	LCL().SysCallN(1542, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) GridHeight() int32 {
	r1 := LCL().SysCallN(1543, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) GridWidth() int32 {
	r1 := LCL().SysCallN(1547, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) IsCellSelected(aCol, aRow int32) bool {
	r1 := LCL().SysCallN(1549, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomDrawGrid) LeftCol() int32 {
	r1 := LCL().SysCallN(1550, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetLeftCol(AValue int32) {
	LCL().SysCallN(1550, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Row() int32 {
	r1 := LCL().SysCallN(1555, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRow(AValue int32) {
	LCL().SysCallN(1555, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) RowHeights(aRow int32) int32 {
	r1 := LCL().SysCallN(1557, 0, m.Instance(), uintptr(aRow))
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRowHeights(aRow int32, AValue int32) {
	LCL().SysCallN(1557, 1, m.Instance(), uintptr(aRow), uintptr(AValue))
}

func (m *TCustomDrawGrid) SaveOptions() TSaveOptions {
	r1 := LCL().SysCallN(1558, 0, m.Instance(), 0)
	return TSaveOptions(r1)
}

func (m *TCustomDrawGrid) SetSaveOptions(AValue TSaveOptions) {
	LCL().SysCallN(1558, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) SelectedColor() TColor {
	r1 := LCL().SysCallN(1560, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetSelectedColor(AValue TColor) {
	LCL().SysCallN(1560, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) SelectedColumn() IGridColumn {
	r1 := LCL().SysCallN(1561, m.Instance())
	return AsGridColumn(r1)
}

func (m *TCustomDrawGrid) Selection() (resultGridRect TGridRect) {
	LCL().SysCallN(1562, 0, m.Instance(), uintptr(unsafePointer(&resultGridRect)), uintptr(unsafePointer(&resultGridRect)))
	return
}

func (m *TCustomDrawGrid) SetSelection(AValue *TGridRect) {
	LCL().SysCallN(1562, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomDrawGrid) StrictSort() bool {
	r1 := LCL().SysCallN(1603, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetStrictSort(AValue bool) {
	LCL().SysCallN(1603, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) TopRow() int32 {
	r1 := LCL().SysCallN(1605, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetTopRow(AValue int32) {
	LCL().SysCallN(1605, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) UseXORFeatures() bool {
	r1 := LCL().SysCallN(1606, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetUseXORFeatures(AValue bool) {
	LCL().SysCallN(1606, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) AutoAdvance() TAutoAdvance {
	r1 := LCL().SysCallN(1510, 0, m.Instance(), 0)
	return TAutoAdvance(r1)
}

func (m *TCustomDrawGrid) SetAutoAdvance(AValue TAutoAdvance) {
	LCL().SysCallN(1510, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) AutoFillColumns() bool {
	r1 := LCL().SysCallN(1511, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAutoFillColumns(AValue bool) {
	LCL().SysCallN(1511, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) ColCount() int32 {
	r1 := LCL().SysCallN(1515, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetColCount(AValue int32) {
	LCL().SysCallN(1515, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Columns() IGridColumns {
	r1 := LCL().SysCallN(1518, 0, m.Instance(), 0)
	return AsGridColumns(r1)
}

func (m *TCustomDrawGrid) SetColumns(AValue IGridColumns) {
	LCL().SysCallN(1518, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomDrawGrid) DefaultColWidth() int32 {
	r1 := LCL().SysCallN(1520, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetDefaultColWidth(AValue int32) {
	LCL().SysCallN(1520, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) DefaultDrawing() bool {
	r1 := LCL().SysCallN(1522, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetDefaultDrawing(AValue bool) {
	LCL().SysCallN(1522, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) DefaultRowHeight() int32 {
	r1 := LCL().SysCallN(1523, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetDefaultRowHeight(AValue int32) {
	LCL().SysCallN(1523, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FadeUnfocusedSelection() bool {
	r1 := LCL().SysCallN(1533, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFadeUnfocusedSelection(AValue bool) {
	LCL().SysCallN(1533, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FixedColor() TColor {
	r1 := LCL().SysCallN(1535, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedColor(AValue TColor) {
	LCL().SysCallN(1535, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedCols() int32 {
	r1 := LCL().SysCallN(1536, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetFixedCols(AValue int32) {
	LCL().SysCallN(1536, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedHotColor() TColor {
	r1 := LCL().SysCallN(1538, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedHotColor(AValue TColor) {
	LCL().SysCallN(1538, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedRows() int32 {
	r1 := LCL().SysCallN(1539, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetFixedRows(AValue int32) {
	LCL().SysCallN(1539, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Flat() bool {
	r1 := LCL().SysCallN(1540, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFlat(AValue bool) {
	LCL().SysCallN(1540, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) GridLineColor() TColor {
	r1 := LCL().SysCallN(1544, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetGridLineColor(AValue TColor) {
	LCL().SysCallN(1544, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) GridLineStyle() TPenStyle {
	r1 := LCL().SysCallN(1545, 0, m.Instance(), 0)
	return TPenStyle(r1)
}

func (m *TCustomDrawGrid) SetGridLineStyle(AValue TPenStyle) {
	LCL().SysCallN(1545, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) GridLineWidth() int32 {
	r1 := LCL().SysCallN(1546, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetGridLineWidth(AValue int32) {
	LCL().SysCallN(1546, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Options() TGridOptions {
	r1 := LCL().SysCallN(1552, 0, m.Instance(), 0)
	return TGridOptions(r1)
}

func (m *TCustomDrawGrid) SetOptions(AValue TGridOptions) {
	LCL().SysCallN(1552, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Options2() TGridOptions2 {
	r1 := LCL().SysCallN(1553, 0, m.Instance(), 0)
	return TGridOptions2(r1)
}

func (m *TCustomDrawGrid) SetOptions2(AValue TGridOptions2) {
	LCL().SysCallN(1553, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ParentShowHint() bool {
	r1 := LCL().SysCallN(1554, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetParentShowHint(AValue bool) {
	LCL().SysCallN(1554, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) RowCount() int32 {
	r1 := LCL().SysCallN(1556, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRowCount(AValue int32) {
	LCL().SysCallN(1556, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ScrollBars() TScrollStyle {
	r1 := LCL().SysCallN(1559, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomDrawGrid) SetScrollBars(AValue TScrollStyle) {
	LCL().SysCallN(1559, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) TabAdvance() TAutoAdvance {
	r1 := LCL().SysCallN(1604, 0, m.Instance(), 0)
	return TAutoAdvance(r1)
}

func (m *TCustomDrawGrid) SetTabAdvance(AValue TAutoAdvance) {
	LCL().SysCallN(1604, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) VisibleColCount() int32 {
	r1 := LCL().SysCallN(1607, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) VisibleRowCount() int32 {
	r1 := LCL().SysCallN(1608, m.Instance())
	return int32(r1)
}

func CustomDrawGridClass() TClass {
	ret := LCL().SysCallN(1513)
	return TClass(ret)
}

func (m *TCustomDrawGrid) DeleteColRow(IsColumn bool, index int32) {
	LCL().SysCallN(1525, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) DeleteCol(Index int32) {
	LCL().SysCallN(1524, m.Instance(), uintptr(Index))
}

func (m *TCustomDrawGrid) DeleteRow(Index int32) {
	LCL().SysCallN(1526, m.Instance(), uintptr(Index))
}

func (m *TCustomDrawGrid) ExchangeColRow(IsColumn bool, index, WithIndex int32) {
	LCL().SysCallN(1531, m.Instance(), PascalBool(IsColumn), uintptr(index), uintptr(WithIndex))
}

func (m *TCustomDrawGrid) InsertColRow(IsColumn bool, index int32) {
	LCL().SysCallN(1548, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) MoveColRow(IsColumn bool, FromIndex, ToIndex int32) {
	LCL().SysCallN(1551, m.Instance(), PascalBool(IsColumn), uintptr(FromIndex), uintptr(ToIndex))
}

func (m *TCustomDrawGrid) SortColRow(IsColumn bool, index int32) {
	LCL().SysCallN(1601, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) SortColRow1(IsColumn bool, Index, FromIndex, ToIndex int32) {
	LCL().SysCallN(1602, m.Instance(), PascalBool(IsColumn), uintptr(Index), uintptr(FromIndex), uintptr(ToIndex))
}

func (m *TCustomDrawGrid) DefaultDrawCell(aCol, aRow int32, aRect *TRect, aState TGridDrawState) {
	var result1 uintptr
	LCL().SysCallN(1521, m.Instance(), uintptr(aCol), uintptr(aRow), uintptr(unsafePointer(&result1)), uintptr(aState))
	*aRect = *(*TRect)(getPointer(result1))
}

func (m *TCustomDrawGrid) SetOnAfterSelection(fn TOnSelectEvent) {
	if m.afterSelectionPtr != 0 {
		RemoveEventElement(m.afterSelectionPtr)
	}
	m.afterSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1563, m.Instance(), m.afterSelectionPtr)
}

func (m *TCustomDrawGrid) SetOnBeforeSelection(fn TOnSelectEvent) {
	if m.beforeSelectionPtr != 0 {
		RemoveEventElement(m.beforeSelectionPtr)
	}
	m.beforeSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1564, m.Instance(), m.beforeSelectionPtr)
}

func (m *TCustomDrawGrid) SetOnColRowDeleted(fn TGridOperationEvent) {
	if m.colRowDeletedPtr != 0 {
		RemoveEventElement(m.colRowDeletedPtr)
	}
	m.colRowDeletedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1566, m.Instance(), m.colRowDeletedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowExchanged(fn TGridOperationEvent) {
	if m.colRowExchangedPtr != 0 {
		RemoveEventElement(m.colRowExchangedPtr)
	}
	m.colRowExchangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1567, m.Instance(), m.colRowExchangedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowInserted(fn TGridOperationEvent) {
	if m.colRowInsertedPtr != 0 {
		RemoveEventElement(m.colRowInsertedPtr)
	}
	m.colRowInsertedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1568, m.Instance(), m.colRowInsertedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowMoved(fn TGridOperationEvent) {
	if m.colRowMovedPtr != 0 {
		RemoveEventElement(m.colRowMovedPtr)
	}
	m.colRowMovedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1569, m.Instance(), m.colRowMovedPtr)
}

func (m *TCustomDrawGrid) SetOnCompareCells(fn TOnCompareCells) {
	if m.compareCellsPtr != 0 {
		RemoveEventElement(m.compareCellsPtr)
	}
	m.compareCellsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1570, m.Instance(), m.compareCellsPtr)
}

func (m *TCustomDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1571, m.Instance(), m.contextPopupPtr)
}

func (m *TCustomDrawGrid) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1572, m.Instance(), m.dblClickPtr)
}

func (m *TCustomDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1573, m.Instance(), m.dragDropPtr)
}

func (m *TCustomDrawGrid) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1574, m.Instance(), m.dragOverPtr)
}

func (m *TCustomDrawGrid) SetOnDrawCell(fn TOnDrawCell) {
	if m.drawCellPtr != 0 {
		RemoveEventElement(m.drawCellPtr)
	}
	m.drawCellPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1575, m.Instance(), m.drawCellPtr)
}

func (m *TCustomDrawGrid) SetOnButtonClick(fn TOnSelectEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1565, m.Instance(), m.buttonClickPtr)
}

func (m *TCustomDrawGrid) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1576, m.Instance(), m.endDockPtr)
}

func (m *TCustomDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1577, m.Instance(), m.endDragPtr)
}

func (m *TCustomDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
	if m.getEditMaskPtr != 0 {
		RemoveEventElement(m.getEditMaskPtr)
	}
	m.getEditMaskPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1578, m.Instance(), m.getEditMaskPtr)
}

func (m *TCustomDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
	if m.getEditTextPtr != 0 {
		RemoveEventElement(m.getEditTextPtr)
	}
	m.getEditTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1579, m.Instance(), m.getEditTextPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderClick(fn THdrEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1580, m.Instance(), m.headerClickPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderSized(fn THdrEvent) {
	if m.headerSizedPtr != 0 {
		RemoveEventElement(m.headerSizedPtr)
	}
	m.headerSizedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1581, m.Instance(), m.headerSizedPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderSizing(fn THeaderSizingEvent) {
	if m.headerSizingPtr != 0 {
		RemoveEventElement(m.headerSizingPtr)
	}
	m.headerSizingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1582, m.Instance(), m.headerSizingPtr)
}

func (m *TCustomDrawGrid) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1583, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1584, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1585, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1586, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomDrawGrid) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1587, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1588, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1589, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1590, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomDrawGrid) SetOnPickListSelect(fn TNotifyEvent) {
	if m.pickListSelectPtr != 0 {
		RemoveEventElement(m.pickListSelectPtr)
	}
	m.pickListSelectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1591, m.Instance(), m.pickListSelectPtr)
}

func (m *TCustomDrawGrid) SetOnPrepareCanvas(fn TOnPrepareCanvasEvent) {
	if m.prepareCanvasPtr != 0 {
		RemoveEventElement(m.prepareCanvasPtr)
	}
	m.prepareCanvasPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1592, m.Instance(), m.prepareCanvasPtr)
}

func (m *TCustomDrawGrid) SetOnSelectEditor(fn TSelectEditorEvent) {
	if m.selectEditorPtr != 0 {
		RemoveEventElement(m.selectEditorPtr)
	}
	m.selectEditorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1594, m.Instance(), m.selectEditorPtr)
}

func (m *TCustomDrawGrid) SetOnSelection(fn TOnSelectEvent) {
	if m.selectionPtr != 0 {
		RemoveEventElement(m.selectionPtr)
	}
	m.selectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1595, m.Instance(), m.selectionPtr)
}

func (m *TCustomDrawGrid) SetOnSelectCell(fn TOnSelectCellEvent) {
	if m.selectCellPtr != 0 {
		RemoveEventElement(m.selectCellPtr)
	}
	m.selectCellPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1593, m.Instance(), m.selectCellPtr)
}

func (m *TCustomDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
	if m.setEditTextPtr != 0 {
		RemoveEventElement(m.setEditTextPtr)
	}
	m.setEditTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1596, m.Instance(), m.setEditTextPtr)
}

func (m *TCustomDrawGrid) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1597, m.Instance(), m.startDockPtr)
}

func (m *TCustomDrawGrid) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1598, m.Instance(), m.startDragPtr)
}

func (m *TCustomDrawGrid) SetOnTopleftChanged(fn TNotifyEvent) {
	if m.topleftChangedPtr != 0 {
		RemoveEventElement(m.topleftChangedPtr)
	}
	m.topleftChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1599, m.Instance(), m.topleftChangedPtr)
}

func (m *TCustomDrawGrid) SetOnValidateEntry(fn TValidateEntryEvent) {
	if m.validateEntryPtr != 0 {
		RemoveEventElement(m.validateEntryPtr)
	}
	m.validateEntryPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1600, m.Instance(), m.validateEntryPtr)
}
