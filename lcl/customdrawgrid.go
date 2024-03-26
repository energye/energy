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
	r1 := LCL().SysCallN(1329, GetObjectUintptr(AOwner))
	return AsCustomDrawGrid(r1)
}

func (m *TCustomDrawGrid) AllowOutboundEvents() bool {
	r1 := LCL().SysCallN(1318, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAllowOutboundEvents(AValue bool) {
	LCL().SysCallN(1318, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) BorderColor() TColor {
	r1 := LCL().SysCallN(1322, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetBorderColor(AValue TColor) {
	LCL().SysCallN(1322, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Col() int32 {
	r1 := LCL().SysCallN(1324, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetCol(AValue int32) {
	LCL().SysCallN(1324, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ColWidths(aCol int32) int32 {
	r1 := LCL().SysCallN(1327, 0, m.Instance(), uintptr(aCol))
	return int32(r1)
}

func (m *TCustomDrawGrid) SetColWidths(aCol int32, AValue int32) {
	LCL().SysCallN(1327, 1, m.Instance(), uintptr(aCol), uintptr(AValue))
}

func (m *TCustomDrawGrid) ColRow() (resultPoint TPoint) {
	LCL().SysCallN(1326, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TCustomDrawGrid) SetColRow(AValue *TPoint) {
	LCL().SysCallN(1326, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TCustomDrawGrid) DisabledFontColor() TColor {
	r1 := LCL().SysCallN(1337, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetDisabledFontColor(AValue TColor) {
	LCL().SysCallN(1337, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Editor() IWinControl {
	r1 := LCL().SysCallN(1338, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TCustomDrawGrid) SetEditor(AValue IWinControl) {
	LCL().SysCallN(1338, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomDrawGrid) EditorBorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1339, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomDrawGrid) SetEditorBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1339, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) EditorMode() bool {
	r1 := LCL().SysCallN(1340, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetEditorMode(AValue bool) {
	LCL().SysCallN(1340, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) ExtendedColSizing() bool {
	r1 := LCL().SysCallN(1342, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetExtendedColSizing(AValue bool) {
	LCL().SysCallN(1342, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) AltColorStartNormal() bool {
	r1 := LCL().SysCallN(1319, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAltColorStartNormal(AValue bool) {
	LCL().SysCallN(1319, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FastEditing() bool {
	r1 := LCL().SysCallN(1344, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFastEditing(AValue bool) {
	LCL().SysCallN(1344, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FixedGridLineColor() TColor {
	r1 := LCL().SysCallN(1347, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedGridLineColor(AValue TColor) {
	LCL().SysCallN(1347, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FocusColor() TColor {
	r1 := LCL().SysCallN(1351, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFocusColor(AValue TColor) {
	LCL().SysCallN(1351, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FocusRectVisible() bool {
	r1 := LCL().SysCallN(1352, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFocusRectVisible(AValue bool) {
	LCL().SysCallN(1352, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) GridHeight() int32 {
	r1 := LCL().SysCallN(1353, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) GridWidth() int32 {
	r1 := LCL().SysCallN(1357, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) IsCellSelected(aCol, aRow int32) bool {
	r1 := LCL().SysCallN(1359, m.Instance(), uintptr(aCol), uintptr(aRow))
	return GoBool(r1)
}

func (m *TCustomDrawGrid) LeftCol() int32 {
	r1 := LCL().SysCallN(1360, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetLeftCol(AValue int32) {
	LCL().SysCallN(1360, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Row() int32 {
	r1 := LCL().SysCallN(1365, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRow(AValue int32) {
	LCL().SysCallN(1365, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) RowHeights(aRow int32) int32 {
	r1 := LCL().SysCallN(1367, 0, m.Instance(), uintptr(aRow))
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRowHeights(aRow int32, AValue int32) {
	LCL().SysCallN(1367, 1, m.Instance(), uintptr(aRow), uintptr(AValue))
}

func (m *TCustomDrawGrid) SaveOptions() TSaveOptions {
	r1 := LCL().SysCallN(1368, 0, m.Instance(), 0)
	return TSaveOptions(r1)
}

func (m *TCustomDrawGrid) SetSaveOptions(AValue TSaveOptions) {
	LCL().SysCallN(1368, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) SelectedColor() TColor {
	r1 := LCL().SysCallN(1370, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetSelectedColor(AValue TColor) {
	LCL().SysCallN(1370, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) SelectedColumn() IGridColumn {
	r1 := LCL().SysCallN(1371, m.Instance())
	return AsGridColumn(r1)
}

func (m *TCustomDrawGrid) Selection() (resultGridRect TGridRect) {
	LCL().SysCallN(1372, 0, m.Instance(), uintptr(unsafe.Pointer(&resultGridRect)), uintptr(unsafe.Pointer(&resultGridRect)))
	return
}

func (m *TCustomDrawGrid) SetSelection(AValue *TGridRect) {
	LCL().SysCallN(1372, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TCustomDrawGrid) StrictSort() bool {
	r1 := LCL().SysCallN(1413, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetStrictSort(AValue bool) {
	LCL().SysCallN(1413, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) TopRow() int32 {
	r1 := LCL().SysCallN(1415, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetTopRow(AValue int32) {
	LCL().SysCallN(1415, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) UseXORFeatures() bool {
	r1 := LCL().SysCallN(1416, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetUseXORFeatures(AValue bool) {
	LCL().SysCallN(1416, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) AutoAdvance() TAutoAdvance {
	r1 := LCL().SysCallN(1320, 0, m.Instance(), 0)
	return TAutoAdvance(r1)
}

func (m *TCustomDrawGrid) SetAutoAdvance(AValue TAutoAdvance) {
	LCL().SysCallN(1320, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) AutoFillColumns() bool {
	r1 := LCL().SysCallN(1321, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetAutoFillColumns(AValue bool) {
	LCL().SysCallN(1321, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) ColCount() int32 {
	r1 := LCL().SysCallN(1325, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetColCount(AValue int32) {
	LCL().SysCallN(1325, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Columns() IGridColumns {
	r1 := LCL().SysCallN(1328, 0, m.Instance(), 0)
	return AsGridColumns(r1)
}

func (m *TCustomDrawGrid) SetColumns(AValue IGridColumns) {
	LCL().SysCallN(1328, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomDrawGrid) DefaultColWidth() int32 {
	r1 := LCL().SysCallN(1330, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetDefaultColWidth(AValue int32) {
	LCL().SysCallN(1330, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) DefaultDrawing() bool {
	r1 := LCL().SysCallN(1332, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetDefaultDrawing(AValue bool) {
	LCL().SysCallN(1332, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) DefaultRowHeight() int32 {
	r1 := LCL().SysCallN(1333, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetDefaultRowHeight(AValue int32) {
	LCL().SysCallN(1333, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FadeUnfocusedSelection() bool {
	r1 := LCL().SysCallN(1343, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFadeUnfocusedSelection(AValue bool) {
	LCL().SysCallN(1343, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) FixedColor() TColor {
	r1 := LCL().SysCallN(1345, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedColor(AValue TColor) {
	LCL().SysCallN(1345, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedCols() int32 {
	r1 := LCL().SysCallN(1346, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetFixedCols(AValue int32) {
	LCL().SysCallN(1346, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedHotColor() TColor {
	r1 := LCL().SysCallN(1348, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetFixedHotColor(AValue TColor) {
	LCL().SysCallN(1348, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) FixedRows() int32 {
	r1 := LCL().SysCallN(1349, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetFixedRows(AValue int32) {
	LCL().SysCallN(1349, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Flat() bool {
	r1 := LCL().SysCallN(1350, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetFlat(AValue bool) {
	LCL().SysCallN(1350, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) GridLineColor() TColor {
	r1 := LCL().SysCallN(1354, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomDrawGrid) SetGridLineColor(AValue TColor) {
	LCL().SysCallN(1354, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) GridLineStyle() TPenStyle {
	r1 := LCL().SysCallN(1355, 0, m.Instance(), 0)
	return TPenStyle(r1)
}

func (m *TCustomDrawGrid) SetGridLineStyle(AValue TPenStyle) {
	LCL().SysCallN(1355, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) GridLineWidth() int32 {
	r1 := LCL().SysCallN(1356, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetGridLineWidth(AValue int32) {
	LCL().SysCallN(1356, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Options() TGridOptions {
	r1 := LCL().SysCallN(1362, 0, m.Instance(), 0)
	return TGridOptions(r1)
}

func (m *TCustomDrawGrid) SetOptions(AValue TGridOptions) {
	LCL().SysCallN(1362, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) Options2() TGridOptions2 {
	r1 := LCL().SysCallN(1363, 0, m.Instance(), 0)
	return TGridOptions2(r1)
}

func (m *TCustomDrawGrid) SetOptions2(AValue TGridOptions2) {
	LCL().SysCallN(1363, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ParentShowHint() bool {
	r1 := LCL().SysCallN(1364, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomDrawGrid) SetParentShowHint(AValue bool) {
	LCL().SysCallN(1364, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomDrawGrid) RowCount() int32 {
	r1 := LCL().SysCallN(1366, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomDrawGrid) SetRowCount(AValue int32) {
	LCL().SysCallN(1366, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) ScrollBars() TScrollStyle {
	r1 := LCL().SysCallN(1369, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomDrawGrid) SetScrollBars(AValue TScrollStyle) {
	LCL().SysCallN(1369, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) TabAdvance() TAutoAdvance {
	r1 := LCL().SysCallN(1414, 0, m.Instance(), 0)
	return TAutoAdvance(r1)
}

func (m *TCustomDrawGrid) SetTabAdvance(AValue TAutoAdvance) {
	LCL().SysCallN(1414, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomDrawGrid) VisibleColCount() int32 {
	r1 := LCL().SysCallN(1417, m.Instance())
	return int32(r1)
}

func (m *TCustomDrawGrid) VisibleRowCount() int32 {
	r1 := LCL().SysCallN(1418, m.Instance())
	return int32(r1)
}

func CustomDrawGridClass() TClass {
	ret := LCL().SysCallN(1323)
	return TClass(ret)
}

func (m *TCustomDrawGrid) DeleteColRow(IsColumn bool, index int32) {
	LCL().SysCallN(1335, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) DeleteCol(Index int32) {
	LCL().SysCallN(1334, m.Instance(), uintptr(Index))
}

func (m *TCustomDrawGrid) DeleteRow(Index int32) {
	LCL().SysCallN(1336, m.Instance(), uintptr(Index))
}

func (m *TCustomDrawGrid) ExchangeColRow(IsColumn bool, index, WithIndex int32) {
	LCL().SysCallN(1341, m.Instance(), PascalBool(IsColumn), uintptr(index), uintptr(WithIndex))
}

func (m *TCustomDrawGrid) InsertColRow(IsColumn bool, index int32) {
	LCL().SysCallN(1358, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) MoveColRow(IsColumn bool, FromIndex, ToIndex int32) {
	LCL().SysCallN(1361, m.Instance(), PascalBool(IsColumn), uintptr(FromIndex), uintptr(ToIndex))
}

func (m *TCustomDrawGrid) SortColRow(IsColumn bool, index int32) {
	LCL().SysCallN(1411, m.Instance(), PascalBool(IsColumn), uintptr(index))
}

func (m *TCustomDrawGrid) SortColRow1(IsColumn bool, Index, FromIndex, ToIndex int32) {
	LCL().SysCallN(1412, m.Instance(), PascalBool(IsColumn), uintptr(Index), uintptr(FromIndex), uintptr(ToIndex))
}

func (m *TCustomDrawGrid) DefaultDrawCell(aCol, aRow int32, aRect *TRect, aState TGridDrawState) {
	var result1 uintptr
	LCL().SysCallN(1331, m.Instance(), uintptr(aCol), uintptr(aRow), uintptr(unsafe.Pointer(&result1)), uintptr(aState))
	*aRect = *(*TRect)(getPointer(result1))
}

func (m *TCustomDrawGrid) SetOnAfterSelection(fn TOnSelectEvent) {
	if m.afterSelectionPtr != 0 {
		RemoveEventElement(m.afterSelectionPtr)
	}
	m.afterSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1373, m.Instance(), m.afterSelectionPtr)
}

func (m *TCustomDrawGrid) SetOnBeforeSelection(fn TOnSelectEvent) {
	if m.beforeSelectionPtr != 0 {
		RemoveEventElement(m.beforeSelectionPtr)
	}
	m.beforeSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1374, m.Instance(), m.beforeSelectionPtr)
}

func (m *TCustomDrawGrid) SetOnColRowDeleted(fn TGridOperationEvent) {
	if m.colRowDeletedPtr != 0 {
		RemoveEventElement(m.colRowDeletedPtr)
	}
	m.colRowDeletedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1376, m.Instance(), m.colRowDeletedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowExchanged(fn TGridOperationEvent) {
	if m.colRowExchangedPtr != 0 {
		RemoveEventElement(m.colRowExchangedPtr)
	}
	m.colRowExchangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1377, m.Instance(), m.colRowExchangedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowInserted(fn TGridOperationEvent) {
	if m.colRowInsertedPtr != 0 {
		RemoveEventElement(m.colRowInsertedPtr)
	}
	m.colRowInsertedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1378, m.Instance(), m.colRowInsertedPtr)
}

func (m *TCustomDrawGrid) SetOnColRowMoved(fn TGridOperationEvent) {
	if m.colRowMovedPtr != 0 {
		RemoveEventElement(m.colRowMovedPtr)
	}
	m.colRowMovedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1379, m.Instance(), m.colRowMovedPtr)
}

func (m *TCustomDrawGrid) SetOnCompareCells(fn TOnCompareCells) {
	if m.compareCellsPtr != 0 {
		RemoveEventElement(m.compareCellsPtr)
	}
	m.compareCellsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1380, m.Instance(), m.compareCellsPtr)
}

func (m *TCustomDrawGrid) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1381, m.Instance(), m.contextPopupPtr)
}

func (m *TCustomDrawGrid) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1382, m.Instance(), m.dblClickPtr)
}

func (m *TCustomDrawGrid) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1383, m.Instance(), m.dragDropPtr)
}

func (m *TCustomDrawGrid) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1384, m.Instance(), m.dragOverPtr)
}

func (m *TCustomDrawGrid) SetOnDrawCell(fn TOnDrawCell) {
	if m.drawCellPtr != 0 {
		RemoveEventElement(m.drawCellPtr)
	}
	m.drawCellPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1385, m.Instance(), m.drawCellPtr)
}

func (m *TCustomDrawGrid) SetOnButtonClick(fn TOnSelectEvent) {
	if m.buttonClickPtr != 0 {
		RemoveEventElement(m.buttonClickPtr)
	}
	m.buttonClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1375, m.Instance(), m.buttonClickPtr)
}

func (m *TCustomDrawGrid) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1386, m.Instance(), m.endDockPtr)
}

func (m *TCustomDrawGrid) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1387, m.Instance(), m.endDragPtr)
}

func (m *TCustomDrawGrid) SetOnGetEditMask(fn TGetEditEvent) {
	if m.getEditMaskPtr != 0 {
		RemoveEventElement(m.getEditMaskPtr)
	}
	m.getEditMaskPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1388, m.Instance(), m.getEditMaskPtr)
}

func (m *TCustomDrawGrid) SetOnGetEditText(fn TGetEditEvent) {
	if m.getEditTextPtr != 0 {
		RemoveEventElement(m.getEditTextPtr)
	}
	m.getEditTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1389, m.Instance(), m.getEditTextPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderClick(fn THdrEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1390, m.Instance(), m.headerClickPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderSized(fn THdrEvent) {
	if m.headerSizedPtr != 0 {
		RemoveEventElement(m.headerSizedPtr)
	}
	m.headerSizedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1391, m.Instance(), m.headerSizedPtr)
}

func (m *TCustomDrawGrid) SetOnHeaderSizing(fn THeaderSizingEvent) {
	if m.headerSizingPtr != 0 {
		RemoveEventElement(m.headerSizingPtr)
	}
	m.headerSizingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1392, m.Instance(), m.headerSizingPtr)
}

func (m *TCustomDrawGrid) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1393, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomDrawGrid) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1394, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomDrawGrid) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1395, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomDrawGrid) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1396, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomDrawGrid) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1397, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1398, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1399, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomDrawGrid) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1400, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomDrawGrid) SetOnPickListSelect(fn TNotifyEvent) {
	if m.pickListSelectPtr != 0 {
		RemoveEventElement(m.pickListSelectPtr)
	}
	m.pickListSelectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1401, m.Instance(), m.pickListSelectPtr)
}

func (m *TCustomDrawGrid) SetOnPrepareCanvas(fn TOnPrepareCanvasEvent) {
	if m.prepareCanvasPtr != 0 {
		RemoveEventElement(m.prepareCanvasPtr)
	}
	m.prepareCanvasPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1402, m.Instance(), m.prepareCanvasPtr)
}

func (m *TCustomDrawGrid) SetOnSelectEditor(fn TSelectEditorEvent) {
	if m.selectEditorPtr != 0 {
		RemoveEventElement(m.selectEditorPtr)
	}
	m.selectEditorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1404, m.Instance(), m.selectEditorPtr)
}

func (m *TCustomDrawGrid) SetOnSelection(fn TOnSelectEvent) {
	if m.selectionPtr != 0 {
		RemoveEventElement(m.selectionPtr)
	}
	m.selectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1405, m.Instance(), m.selectionPtr)
}

func (m *TCustomDrawGrid) SetOnSelectCell(fn TOnSelectCellEvent) {
	if m.selectCellPtr != 0 {
		RemoveEventElement(m.selectCellPtr)
	}
	m.selectCellPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1403, m.Instance(), m.selectCellPtr)
}

func (m *TCustomDrawGrid) SetOnSetEditText(fn TSetEditEvent) {
	if m.setEditTextPtr != 0 {
		RemoveEventElement(m.setEditTextPtr)
	}
	m.setEditTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1406, m.Instance(), m.setEditTextPtr)
}

func (m *TCustomDrawGrid) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1407, m.Instance(), m.startDockPtr)
}

func (m *TCustomDrawGrid) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1408, m.Instance(), m.startDragPtr)
}

func (m *TCustomDrawGrid) SetOnTopleftChanged(fn TNotifyEvent) {
	if m.topleftChangedPtr != 0 {
		RemoveEventElement(m.topleftChangedPtr)
	}
	m.topleftChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1409, m.Instance(), m.topleftChangedPtr)
}

func (m *TCustomDrawGrid) SetOnValidateEntry(fn TValidateEntryEvent) {
	if m.validateEntryPtr != 0 {
		RemoveEventElement(m.validateEntryPtr)
	}
	m.validateEntryPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1410, m.Instance(), m.validateEntryPtr)
}
