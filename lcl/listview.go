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

// IListView Parent: ICustomListView
type IListView interface {
	ICustomListView
	DeleteSelected()
	AllocBy() int32                                                      // property
	SetAllocBy(AValue int32)                                             // property
	AutoSort() bool                                                      // property
	SetAutoSort(AValue bool)                                             // property
	AutoSortIndicator() bool                                             // property
	SetAutoSortIndicator(AValue bool)                                    // property
	AutoWidthLastColumn() bool                                           // property
	SetAutoWidthLastColumn(AValue bool)                                  // property
	Columns() IListColumns                                               // property
	SetColumns(AValue IListColumns)                                      // property
	ColumnClick() bool                                                   // property
	SetColumnClick(AValue bool)                                          // property
	DragCursor() TCursor                                                 // property
	SetDragCursor(AValue TCursor)                                        // property
	DragKind() TDragKind                                                 // property
	SetDragKind(AValue TDragKind)                                        // property
	DragMode() TDragMode                                                 // property
	SetDragMode(AValue TDragMode)                                        // property
	HideSelection() bool                                                 // property
	SetHideSelection(AValue bool)                                        // property
	LargeImages() ICustomImageList                                       // property
	SetLargeImages(AValue ICustomImageList)                              // property
	LargeImagesWidth() int32                                             // property
	SetLargeImagesWidth(AValue int32)                                    // property
	OwnerDraw() bool                                                     // property
	SetOwnerDraw(AValue bool)                                            // property
	ParentColor() bool                                                   // property
	SetParentColor(AValue bool)                                          // property
	ParentFont() bool                                                    // property
	SetParentFont(AValue bool)                                           // property
	ParentShowHint() bool                                                // property
	SetParentShowHint(AValue bool)                                       // property
	ScrollBars() TScrollStyle                                            // property
	SetScrollBars(AValue TScrollStyle)                                   // property
	ShowColumnHeaders() bool                                             // property
	SetShowColumnHeaders(AValue bool)                                    // property
	SmallImages() ICustomImageList                                       // property
	SetSmallImages(AValue ICustomImageList)                              // property
	SmallImagesWidth() int32                                             // property
	SetSmallImagesWidth(AValue int32)                                    // property
	SortColumn() int32                                                   // property
	SetSortColumn(AValue int32)                                          // property
	SortDirection() TSortDirection                                       // property
	SetSortDirection(AValue TSortDirection)                              // property
	SortType() TSortType                                                 // property
	SetSortType(AValue TSortType)                                        // property
	StateImages() ICustomImageList                                       // property
	SetStateImages(AValue ICustomImageList)                              // property
	StateImagesWidth() int32                                             // property
	SetStateImagesWidth(AValue int32)                                    // property
	ToolTips() bool                                                      // property
	SetToolTips(AValue bool)                                             // property
	ViewStyle() TViewStyle                                               // property
	SetViewStyle(AValue TViewStyle)                                      // property
	SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent)               // property event
	SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent)       // property event
	SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) // property event
	SetOnChange(fn TLVChangeEvent)                                       // property event
	SetOnColumnClick(fn TLVColumnClickEvent)                             // property event
	SetOnCompare(fn TLVCompareEvent)                                     // property event
	SetOnContextPopup(fn TContextPopupEvent)                             // property event
	SetOnCreateItemClass(fn TLVCreateItemClassEvent)                     // property event
	SetOnCustomDraw(fn TLVCustomDrawEvent)                               // property event
	SetOnCustomDrawItem(fn TLVCustomDrawItemEvent)                       // property event
	SetOnCustomDrawSubItem(fn TLVCustomDrawSubItemEvent)                 // property event
	SetOnData(fn TLVDataEvent)                                           // property event
	SetOnDataFind(fn TLVDataFindEvent)                                   // property event
	SetOnDataHint(fn TLVDataHintEvent)                                   // property event
	SetOnDataStateChange(fn TLVDataStateChangeEvent)                     // property event
	SetOnDblClick(fn TNotifyEvent)                                       // property event
	SetOnDeletion(fn TLVDeletedEvent)                                    // property event
	SetOnDragDrop(fn TDragDropEvent)                                     // property event
	SetOnDragOver(fn TDragOverEvent)                                     // property event
	SetOnDrawItem(fn TLVDrawItemEvent)                                   // property event
	SetOnEdited(fn TLVEditedEvent)                                       // property event
	SetOnEditing(fn TLVEditingEvent)                                     // property event
	SetOnEndDock(fn TEndDragEvent)                                       // property event
	SetOnEndDrag(fn TEndDragEvent)                                       // property event
	SetOnInsert(fn TLVInsertEvent)                                       // property event
	SetOnItemChecked(fn TLVCheckedItemEvent)                             // property event
	SetOnMouseDown(fn TMouseEvent)                                       // property event
	SetOnMouseEnter(fn TNotifyEvent)                                     // property event
	SetOnMouseLeave(fn TNotifyEvent)                                     // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                   // property event
	SetOnMouseUp(fn TMouseEvent)                                         // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                 // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                       // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                         // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)                             // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)                       // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)                      // property event
	SetOnSelectItem(fn TLVSelectItemEvent)                               // property event
	SetOnStartDock(fn TStartDockEvent)                                   // property event
	SetOnStartDrag(fn TStartDragEvent)                                   // property event
}

// TListView Parent: TCustomListView
type TListView struct {
	TCustomListView
	advancedCustomDrawPtr        uintptr
	advancedCustomDrawItemPtr    uintptr
	advancedCustomDrawSubItemPtr uintptr
	changePtr                    uintptr
	columnClickPtr               uintptr
	comparePtr                   uintptr
	contextPopupPtr              uintptr
	createItemClassPtr           uintptr
	customDrawPtr                uintptr
	customDrawItemPtr            uintptr
	customDrawSubItemPtr         uintptr
	dataPtr                      uintptr
	dataFindPtr                  uintptr
	dataHintPtr                  uintptr
	dataStateChangePtr           uintptr
	dblClickPtr                  uintptr
	deletionPtr                  uintptr
	dragDropPtr                  uintptr
	dragOverPtr                  uintptr
	drawItemPtr                  uintptr
	editedPtr                    uintptr
	editingPtr                   uintptr
	endDockPtr                   uintptr
	endDragPtr                   uintptr
	insertPtr                    uintptr
	itemCheckedPtr               uintptr
	mouseDownPtr                 uintptr
	mouseEnterPtr                uintptr
	mouseLeavePtr                uintptr
	mouseMovePtr                 uintptr
	mouseUpPtr                   uintptr
	mouseWheelPtr                uintptr
	mouseWheelDownPtr            uintptr
	mouseWheelUpPtr              uintptr
	mouseWheelHorzPtr            uintptr
	mouseWheelLeftPtr            uintptr
	mouseWheelRightPtr           uintptr
	selectItemPtr                uintptr
	startDockPtr                 uintptr
	startDragPtr                 uintptr
}

func NewListView(AOwner IComponent) IListView {
	r1 := LCL().SysCallN(4080, GetObjectUintptr(AOwner))
	return AsListView(r1)
}

func (m *TListView) AllocBy() int32 {
	r1 := LCL().SysCallN(4073, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetAllocBy(AValue int32) {
	LCL().SysCallN(4073, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) AutoSort() bool {
	r1 := LCL().SysCallN(4074, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetAutoSort(AValue bool) {
	LCL().SysCallN(4074, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) AutoSortIndicator() bool {
	r1 := LCL().SysCallN(4075, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetAutoSortIndicator(AValue bool) {
	LCL().SysCallN(4075, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) AutoWidthLastColumn() bool {
	r1 := LCL().SysCallN(4076, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetAutoWidthLastColumn(AValue bool) {
	LCL().SysCallN(4076, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) Columns() IListColumns {
	r1 := LCL().SysCallN(4079, 0, m.Instance(), 0)
	return AsListColumns(r1)
}

func (m *TListView) SetColumns(AValue IListColumns) {
	LCL().SysCallN(4079, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) ColumnClick() bool {
	r1 := LCL().SysCallN(4078, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetColumnClick(AValue bool) {
	LCL().SysCallN(4078, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) DragCursor() TCursor {
	r1 := LCL().SysCallN(4081, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TListView) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4081, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) DragKind() TDragKind {
	r1 := LCL().SysCallN(4082, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TListView) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4082, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) DragMode() TDragMode {
	r1 := LCL().SysCallN(4083, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TListView) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4083, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) HideSelection() bool {
	r1 := LCL().SysCallN(4084, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetHideSelection(AValue bool) {
	LCL().SysCallN(4084, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) LargeImages() ICustomImageList {
	r1 := LCL().SysCallN(4085, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TListView) SetLargeImages(AValue ICustomImageList) {
	LCL().SysCallN(4085, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) LargeImagesWidth() int32 {
	r1 := LCL().SysCallN(4086, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetLargeImagesWidth(AValue int32) {
	LCL().SysCallN(4086, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) OwnerDraw() bool {
	r1 := LCL().SysCallN(4087, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetOwnerDraw(AValue bool) {
	LCL().SysCallN(4087, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ParentColor() bool {
	r1 := LCL().SysCallN(4088, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetParentColor(AValue bool) {
	LCL().SysCallN(4088, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ParentFont() bool {
	r1 := LCL().SysCallN(4089, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetParentFont(AValue bool) {
	LCL().SysCallN(4089, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ParentShowHint() bool {
	r1 := LCL().SysCallN(4090, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4090, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ScrollBars() TScrollStyle {
	r1 := LCL().SysCallN(4091, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TListView) SetScrollBars(AValue TScrollStyle) {
	LCL().SysCallN(4091, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) ShowColumnHeaders() bool {
	r1 := LCL().SysCallN(4132, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetShowColumnHeaders(AValue bool) {
	LCL().SysCallN(4132, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) SmallImages() ICustomImageList {
	r1 := LCL().SysCallN(4133, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TListView) SetSmallImages(AValue ICustomImageList) {
	LCL().SysCallN(4133, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) SmallImagesWidth() int32 {
	r1 := LCL().SysCallN(4134, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetSmallImagesWidth(AValue int32) {
	LCL().SysCallN(4134, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) SortColumn() int32 {
	r1 := LCL().SysCallN(4135, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetSortColumn(AValue int32) {
	LCL().SysCallN(4135, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) SortDirection() TSortDirection {
	r1 := LCL().SysCallN(4136, 0, m.Instance(), 0)
	return TSortDirection(r1)
}

func (m *TListView) SetSortDirection(AValue TSortDirection) {
	LCL().SysCallN(4136, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) SortType() TSortType {
	r1 := LCL().SysCallN(4137, 0, m.Instance(), 0)
	return TSortType(r1)
}

func (m *TListView) SetSortType(AValue TSortType) {
	LCL().SysCallN(4137, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) StateImages() ICustomImageList {
	r1 := LCL().SysCallN(4138, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TListView) SetStateImages(AValue ICustomImageList) {
	LCL().SysCallN(4138, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListView) StateImagesWidth() int32 {
	r1 := LCL().SysCallN(4139, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListView) SetStateImagesWidth(AValue int32) {
	LCL().SysCallN(4139, 1, m.Instance(), uintptr(AValue))
}

func (m *TListView) ToolTips() bool {
	r1 := LCL().SysCallN(4140, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListView) SetToolTips(AValue bool) {
	LCL().SysCallN(4140, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListView) ViewStyle() TViewStyle {
	r1 := LCL().SysCallN(4141, 0, m.Instance(), 0)
	return TViewStyle(r1)
}

func (m *TListView) SetViewStyle(AValue TViewStyle) {
	LCL().SysCallN(4141, 1, m.Instance(), uintptr(AValue))
}

func ListViewClass() TClass {
	ret := LCL().SysCallN(4077)
	return TClass(ret)
}

func (m *TListView) SetOnAdvancedCustomDraw(fn TLVAdvancedCustomDrawEvent) {
	if m.advancedCustomDrawPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawPtr)
	}
	m.advancedCustomDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4092, m.Instance(), m.advancedCustomDrawPtr)
}

func (m *TListView) SetOnAdvancedCustomDrawItem(fn TLVAdvancedCustomDrawItemEvent) {
	if m.advancedCustomDrawItemPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawItemPtr)
	}
	m.advancedCustomDrawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4093, m.Instance(), m.advancedCustomDrawItemPtr)
}

func (m *TListView) SetOnAdvancedCustomDrawSubItem(fn TLVAdvancedCustomDrawSubItemEvent) {
	if m.advancedCustomDrawSubItemPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawSubItemPtr)
	}
	m.advancedCustomDrawSubItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4094, m.Instance(), m.advancedCustomDrawSubItemPtr)
}

func (m *TListView) SetOnChange(fn TLVChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4095, m.Instance(), m.changePtr)
}

func (m *TListView) SetOnColumnClick(fn TLVColumnClickEvent) {
	if m.columnClickPtr != 0 {
		RemoveEventElement(m.columnClickPtr)
	}
	m.columnClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4096, m.Instance(), m.columnClickPtr)
}

func (m *TListView) SetOnCompare(fn TLVCompareEvent) {
	if m.comparePtr != 0 {
		RemoveEventElement(m.comparePtr)
	}
	m.comparePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4097, m.Instance(), m.comparePtr)
}

func (m *TListView) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4098, m.Instance(), m.contextPopupPtr)
}

func (m *TListView) SetOnCreateItemClass(fn TLVCreateItemClassEvent) {
	if m.createItemClassPtr != 0 {
		RemoveEventElement(m.createItemClassPtr)
	}
	m.createItemClassPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4099, m.Instance(), m.createItemClassPtr)
}

func (m *TListView) SetOnCustomDraw(fn TLVCustomDrawEvent) {
	if m.customDrawPtr != 0 {
		RemoveEventElement(m.customDrawPtr)
	}
	m.customDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4100, m.Instance(), m.customDrawPtr)
}

func (m *TListView) SetOnCustomDrawItem(fn TLVCustomDrawItemEvent) {
	if m.customDrawItemPtr != 0 {
		RemoveEventElement(m.customDrawItemPtr)
	}
	m.customDrawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4101, m.Instance(), m.customDrawItemPtr)
}

func (m *TListView) SetOnCustomDrawSubItem(fn TLVCustomDrawSubItemEvent) {
	if m.customDrawSubItemPtr != 0 {
		RemoveEventElement(m.customDrawSubItemPtr)
	}
	m.customDrawSubItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4102, m.Instance(), m.customDrawSubItemPtr)
}

func (m *TListView) SetOnData(fn TLVDataEvent) {
	if m.dataPtr != 0 {
		RemoveEventElement(m.dataPtr)
	}
	m.dataPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4103, m.Instance(), m.dataPtr)
}

func (m *TListView) SetOnDataFind(fn TLVDataFindEvent) {
	if m.dataFindPtr != 0 {
		RemoveEventElement(m.dataFindPtr)
	}
	m.dataFindPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4104, m.Instance(), m.dataFindPtr)
}

func (m *TListView) SetOnDataHint(fn TLVDataHintEvent) {
	if m.dataHintPtr != 0 {
		RemoveEventElement(m.dataHintPtr)
	}
	m.dataHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4105, m.Instance(), m.dataHintPtr)
}

func (m *TListView) SetOnDataStateChange(fn TLVDataStateChangeEvent) {
	if m.dataStateChangePtr != 0 {
		RemoveEventElement(m.dataStateChangePtr)
	}
	m.dataStateChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4106, m.Instance(), m.dataStateChangePtr)
}

func (m *TListView) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4107, m.Instance(), m.dblClickPtr)
}

func (m *TListView) SetOnDeletion(fn TLVDeletedEvent) {
	if m.deletionPtr != 0 {
		RemoveEventElement(m.deletionPtr)
	}
	m.deletionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4108, m.Instance(), m.deletionPtr)
}

func (m *TListView) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4109, m.Instance(), m.dragDropPtr)
}

func (m *TListView) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4110, m.Instance(), m.dragOverPtr)
}

func (m *TListView) SetOnDrawItem(fn TLVDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4111, m.Instance(), m.drawItemPtr)
}

func (m *TListView) SetOnEdited(fn TLVEditedEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4112, m.Instance(), m.editedPtr)
}

func (m *TListView) SetOnEditing(fn TLVEditingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4113, m.Instance(), m.editingPtr)
}

func (m *TListView) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4114, m.Instance(), m.endDockPtr)
}

func (m *TListView) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4115, m.Instance(), m.endDragPtr)
}

func (m *TListView) SetOnInsert(fn TLVInsertEvent) {
	if m.insertPtr != 0 {
		RemoveEventElement(m.insertPtr)
	}
	m.insertPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4116, m.Instance(), m.insertPtr)
}

func (m *TListView) SetOnItemChecked(fn TLVCheckedItemEvent) {
	if m.itemCheckedPtr != 0 {
		RemoveEventElement(m.itemCheckedPtr)
	}
	m.itemCheckedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4117, m.Instance(), m.itemCheckedPtr)
}

func (m *TListView) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4118, m.Instance(), m.mouseDownPtr)
}

func (m *TListView) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4119, m.Instance(), m.mouseEnterPtr)
}

func (m *TListView) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4120, m.Instance(), m.mouseLeavePtr)
}

func (m *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4121, m.Instance(), m.mouseMovePtr)
}

func (m *TListView) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4122, m.Instance(), m.mouseUpPtr)
}

func (m *TListView) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4123, m.Instance(), m.mouseWheelPtr)
}

func (m *TListView) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4124, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TListView) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4128, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TListView) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4125, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TListView) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4126, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TListView) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4127, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TListView) SetOnSelectItem(fn TLVSelectItemEvent) {
	if m.selectItemPtr != 0 {
		RemoveEventElement(m.selectItemPtr)
	}
	m.selectItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4129, m.Instance(), m.selectItemPtr)
}

func (m *TListView) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4130, m.Instance(), m.startDockPtr)
}

func (m *TListView) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4131, m.Instance(), m.startDragPtr)
}
