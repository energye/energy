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

// IDrawGrid Parent: ICustomDrawGrid
type IDrawGrid interface {
	ICustomDrawGrid
	InplaceEditor() IWinControl                          // property
	AlternateColor() TColor                              // property
	SetAlternateColor(AValue TColor)                     // property
	AutoEdit() bool                                      // property
	SetAutoEdit(AValue bool)                             // property
	ColRowDraggingCursor() TCursor                       // property
	SetColRowDraggingCursor(AValue TCursor)              // property
	ColRowDragIndicatorColor() TColor                    // property
	SetColRowDragIndicatorColor(AValue TColor)           // property
	ColSizingCursor() TCursor                            // property
	SetColSizingCursor(AValue TCursor)                   // property
	ColumnClickSorts() bool                              // property
	SetColumnClickSorts(AValue bool)                     // property
	DragCursor() TCursor                                 // property
	SetDragCursor(AValue TCursor)                        // property
	DragKind() TDragKind                                 // property
	SetDragKind(AValue TDragKind)                        // property
	DragMode() TDragMode                                 // property
	SetDragMode(AValue TDragMode)                        // property
	ExtendedSelect() bool                                // property
	SetExtendedSelect(AValue bool)                       // property
	HeaderHotZones() TGridZoneSet                        // property
	SetHeaderHotZones(AValue TGridZoneSet)               // property
	HeaderPushZones() TGridZoneSet                       // property
	SetHeaderPushZones(AValue TGridZoneSet)              // property
	ImageIndexSortAsc() TImageIndex                      // property
	SetImageIndexSortAsc(AValue TImageIndex)             // property
	ImageIndexSortDesc() TImageIndex                     // property
	SetImageIndexSortDesc(AValue TImageIndex)            // property
	MouseWheelOption() TMouseWheelOption                 // property
	SetMouseWheelOption(AValue TMouseWheelOption)        // property
	ParentColor() bool                                   // property
	SetParentColor(AValue bool)                          // property
	ParentFont() bool                                    // property
	SetParentFont(AValue bool)                           // property
	RangeSelectMode() TRangeSelectMode                   // property
	SetRangeSelectMode(AValue TRangeSelectMode)          // property
	RowSizingCursor() TCursor                            // property
	SetRowSizingCursor(AValue TCursor)                   // property
	TitleFont() IFont                                    // property
	SetTitleFont(AValue IFont)                           // property
	TitleImageList() IImageList                          // property
	SetTitleImageList(AValue IImageList)                 // property
	TitleImageListWidth() int32                          // property
	SetTitleImageListWidth(AValue int32)                 // property
	TitleStyle() TTitleStyle                             // property
	SetTitleStyle(AValue TTitleStyle)                    // property
	SetOnCheckboxToggled(fn TToggledCheckboxEvent)       // property event
	SetOnEditingDone(fn TNotifyEvent)                    // property event
	SetOnGetCellHint(fn TGetCellHintEvent)               // property event
	SetOnGetCheckboxState(fn TGetCheckboxStateEvent)     // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)      // property event
	SetOnSetCheckboxState(fn TSetCheckboxStateEvent)     // property event
	SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) // property event
	SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent)   // property event
}

// TDrawGrid Parent: TCustomDrawGrid
type TDrawGrid struct {
	TCustomDrawGrid
	checkboxToggledPtr    uintptr
	editingDonePtr        uintptr
	getCellHintPtr        uintptr
	getCheckboxStatePtr   uintptr
	mouseWheelHorzPtr     uintptr
	mouseWheelLeftPtr     uintptr
	mouseWheelRightPtr    uintptr
	setCheckboxStatePtr   uintptr
	userCheckboxBitmapPtr uintptr
	userCheckboxImagePtr  uintptr
}

func NewDrawGrid(AOwner IComponent) IDrawGrid {
	r1 := LCL().SysCallN(2492, GetObjectUintptr(AOwner))
	return AsDrawGrid(r1)
}

func (m *TDrawGrid) InplaceEditor() IWinControl {
	r1 := LCL().SysCallN(2501, m.Instance())
	return AsWinControl(r1)
}

func (m *TDrawGrid) AlternateColor() TColor {
	r1 := LCL().SysCallN(2485, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TDrawGrid) SetAlternateColor(AValue TColor) {
	LCL().SysCallN(2485, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) AutoEdit() bool {
	r1 := LCL().SysCallN(2486, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetAutoEdit(AValue bool) {
	LCL().SysCallN(2486, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) ColRowDraggingCursor() TCursor {
	r1 := LCL().SysCallN(2489, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetColRowDraggingCursor(AValue TCursor) {
	LCL().SysCallN(2489, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ColRowDragIndicatorColor() TColor {
	r1 := LCL().SysCallN(2488, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TDrawGrid) SetColRowDragIndicatorColor(AValue TColor) {
	LCL().SysCallN(2488, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ColSizingCursor() TCursor {
	r1 := LCL().SysCallN(2490, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetColSizingCursor(AValue TCursor) {
	LCL().SysCallN(2490, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ColumnClickSorts() bool {
	r1 := LCL().SysCallN(2491, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetColumnClickSorts(AValue bool) {
	LCL().SysCallN(2491, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) DragCursor() TCursor {
	r1 := LCL().SysCallN(2493, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2493, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) DragKind() TDragKind {
	r1 := LCL().SysCallN(2494, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TDrawGrid) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(2494, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) DragMode() TDragMode {
	r1 := LCL().SysCallN(2495, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDrawGrid) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2495, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ExtendedSelect() bool {
	r1 := LCL().SysCallN(2496, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetExtendedSelect(AValue bool) {
	LCL().SysCallN(2496, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) HeaderHotZones() TGridZoneSet {
	r1 := LCL().SysCallN(2497, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TDrawGrid) SetHeaderHotZones(AValue TGridZoneSet) {
	LCL().SysCallN(2497, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) HeaderPushZones() TGridZoneSet {
	r1 := LCL().SysCallN(2498, 0, m.Instance(), 0)
	return TGridZoneSet(r1)
}

func (m *TDrawGrid) SetHeaderPushZones(AValue TGridZoneSet) {
	LCL().SysCallN(2498, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ImageIndexSortAsc() TImageIndex {
	r1 := LCL().SysCallN(2499, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDrawGrid) SetImageIndexSortAsc(AValue TImageIndex) {
	LCL().SysCallN(2499, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ImageIndexSortDesc() TImageIndex {
	r1 := LCL().SysCallN(2500, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TDrawGrid) SetImageIndexSortDesc(AValue TImageIndex) {
	LCL().SysCallN(2500, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) MouseWheelOption() TMouseWheelOption {
	r1 := LCL().SysCallN(2502, 0, m.Instance(), 0)
	return TMouseWheelOption(r1)
}

func (m *TDrawGrid) SetMouseWheelOption(AValue TMouseWheelOption) {
	LCL().SysCallN(2502, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) ParentColor() bool {
	r1 := LCL().SysCallN(2503, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetParentColor(AValue bool) {
	LCL().SysCallN(2503, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) ParentFont() bool {
	r1 := LCL().SysCallN(2504, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDrawGrid) SetParentFont(AValue bool) {
	LCL().SysCallN(2504, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDrawGrid) RangeSelectMode() TRangeSelectMode {
	r1 := LCL().SysCallN(2505, 0, m.Instance(), 0)
	return TRangeSelectMode(r1)
}

func (m *TDrawGrid) SetRangeSelectMode(AValue TRangeSelectMode) {
	LCL().SysCallN(2505, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) RowSizingCursor() TCursor {
	r1 := LCL().SysCallN(2506, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDrawGrid) SetRowSizingCursor(AValue TCursor) {
	LCL().SysCallN(2506, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) TitleFont() IFont {
	r1 := LCL().SysCallN(2517, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TDrawGrid) SetTitleFont(AValue IFont) {
	LCL().SysCallN(2517, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDrawGrid) TitleImageList() IImageList {
	r1 := LCL().SysCallN(2518, 0, m.Instance(), 0)
	return AsImageList(r1)
}

func (m *TDrawGrid) SetTitleImageList(AValue IImageList) {
	LCL().SysCallN(2518, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TDrawGrid) TitleImageListWidth() int32 {
	r1 := LCL().SysCallN(2519, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDrawGrid) SetTitleImageListWidth(AValue int32) {
	LCL().SysCallN(2519, 1, m.Instance(), uintptr(AValue))
}

func (m *TDrawGrid) TitleStyle() TTitleStyle {
	r1 := LCL().SysCallN(2520, 0, m.Instance(), 0)
	return TTitleStyle(r1)
}

func (m *TDrawGrid) SetTitleStyle(AValue TTitleStyle) {
	LCL().SysCallN(2520, 1, m.Instance(), uintptr(AValue))
}

func DrawGridClass() TClass {
	ret := LCL().SysCallN(2487)
	return TClass(ret)
}

func (m *TDrawGrid) SetOnCheckboxToggled(fn TToggledCheckboxEvent) {
	if m.checkboxToggledPtr != 0 {
		RemoveEventElement(m.checkboxToggledPtr)
	}
	m.checkboxToggledPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2507, m.Instance(), m.checkboxToggledPtr)
}

func (m *TDrawGrid) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2508, m.Instance(), m.editingDonePtr)
}

func (m *TDrawGrid) SetOnGetCellHint(fn TGetCellHintEvent) {
	if m.getCellHintPtr != 0 {
		RemoveEventElement(m.getCellHintPtr)
	}
	m.getCellHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2509, m.Instance(), m.getCellHintPtr)
}

func (m *TDrawGrid) SetOnGetCheckboxState(fn TGetCheckboxStateEvent) {
	if m.getCheckboxStatePtr != 0 {
		RemoveEventElement(m.getCheckboxStatePtr)
	}
	m.getCheckboxStatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2510, m.Instance(), m.getCheckboxStatePtr)
}

func (m *TDrawGrid) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2511, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TDrawGrid) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2512, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TDrawGrid) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2513, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TDrawGrid) SetOnSetCheckboxState(fn TSetCheckboxStateEvent) {
	if m.setCheckboxStatePtr != 0 {
		RemoveEventElement(m.setCheckboxStatePtr)
	}
	m.setCheckboxStatePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2514, m.Instance(), m.setCheckboxStatePtr)
}

func (m *TDrawGrid) SetOnUserCheckboxBitmap(fn TUserCheckBoxBitmapEvent) {
	if m.userCheckboxBitmapPtr != 0 {
		RemoveEventElement(m.userCheckboxBitmapPtr)
	}
	m.userCheckboxBitmapPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2515, m.Instance(), m.userCheckboxBitmapPtr)
}

func (m *TDrawGrid) SetOnUserCheckboxImage(fn TUserCheckBoxImageEvent) {
	if m.userCheckboxImagePtr != 0 {
		RemoveEventElement(m.userCheckboxImagePtr)
	}
	m.userCheckboxImagePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2516, m.Instance(), m.userCheckboxImagePtr)
}
