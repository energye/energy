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

// ITreeView Parent: ICustomTreeView
type ITreeView interface {
	ICustomTreeView
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	AutoExpand() bool                                              // property
	SetAutoExpand(AValue bool)                                     // property
	DisabledFontColor() TColor                                     // property
	SetDisabledFontColor(AValue TColor)                            // property
	DragKind() TDragKind                                           // property
	SetDragKind(AValue TDragKind)                                  // property
	DragCursor() TCursor                                           // property
	SetDragCursor(AValue TCursor)                                  // property
	DragMode() TDragMode                                           // property
	SetDragMode(AValue TDragMode)                                  // property
	HideSelection() bool                                           // property
	SetHideSelection(AValue bool)                                  // property
	HotTrack() bool                                                // property
	SetHotTrack(AValue bool)                                       // property
	HotTrackColor() TColor                                         // property
	SetHotTrackColor(AValue TColor)                                // property
	Indent() int32                                                 // property
	SetIndent(AValue int32)                                        // property
	MultiSelect() bool                                             // property
	SetMultiSelect(AValue bool)                                    // property
	ParentColor() bool                                             // property
	SetParentColor(AValue bool)                                    // property
	ParentFont() bool                                              // property
	SetParentFont(AValue bool)                                     // property
	ParentShowHint() bool                                          // property
	SetParentShowHint(AValue bool)                                 // property
	ReadOnly() bool                                                // property
	SetReadOnly(AValue bool)                                       // property
	RightClickSelect() bool                                        // property
	SetRightClickSelect(AValue bool)                               // property
	RowSelect() bool                                               // property
	SetRowSelect(AValue bool)                                      // property
	ShowButtons() bool                                             // property
	SetShowButtons(AValue bool)                                    // property
	ShowLines() bool                                               // property
	SetShowLines(AValue bool)                                      // property
	ShowRoot() bool                                                // property
	SetShowRoot(AValue bool)                                       // property
	ShowSeparators() bool                                          // property
	SetShowSeparators(AValue bool)                                 // property
	SortType() TSortType                                           // property
	SetSortType(AValue TSortType)                                  // property
	ToolTips() bool                                                // property
	SetToolTips(AValue bool)                                       // property
	SetOnAddition(fn TTVExpandedEvent)                             // property event
	SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent)         // property event
	SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) // property event
	SetOnChange(fn TTVChangedEvent)                                // property event
	SetOnChanging(fn TTVChangingEvent)                             // property event
	SetOnCollapsed(fn TTVExpandedEvent)                            // property event
	SetOnCollapsing(fn TTVCollapsingEvent)                         // property event
	SetOnCompare(fn TTVCompareEvent)                               // property event
	SetOnContextPopup(fn TContextPopupEvent)                       // property event
	SetOnCreateNodeClass(fn TTVCreateNodeClassEvent)               // property event
	SetOnCustomCreateItem(fn TTVCustomCreateNodeEvent)             // property event
	SetOnCustomDraw(fn TTVCustomDrawEvent)                         // property event
	SetOnCustomDrawItem(fn TTVCustomDrawItemEvent)                 // property event
	SetOnCustomDrawArrow(fn TTVCustomDrawArrowEvent)               // property event
	SetOnDblClick(fn TNotifyEvent)                                 // property event
	SetOnDeletion(fn TTVExpandedEvent)                             // property event
	SetOnDragDrop(fn TDragDropEvent)                               // property event
	SetOnDragOver(fn TDragOverEvent)                               // property event
	SetOnEdited(fn TTVEditedEvent)                                 // property event
	SetOnEditing(fn TTVEditingEvent)                               // property event
	SetOnEditingEnd(fn TTVEditingEndEvent)                         // property event
	SetOnEndDrag(fn TEndDragEvent)                                 // property event
	SetOnExpanded(fn TTVExpandedEvent)                             // property event
	SetOnExpanding(fn TTVExpandingEvent)                           // property event
	SetOnGetImageIndex(fn TTVExpandedEvent)                        // property event
	SetOnGetSelectedIndex(fn TTVExpandedEvent)                     // property event
	SetOnHasChildren(fn TTVHasChildrenEvent)                       // property event
	SetOnMouseDown(fn TMouseEvent)                                 // property event
	SetOnMouseEnter(fn TNotifyEvent)                               // property event
	SetOnMouseLeave(fn TNotifyEvent)                               // property event
	SetOnMouseMove(fn TMouseMoveEvent)                             // property event
	SetOnMouseUp(fn TMouseEvent)                                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)                 // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)                   // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)                       // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)                 // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)                // property event
	SetOnNodeChanged(fn TTVNodeChangedEvent)                       // property event
	SetOnSelectionChanged(fn TNotifyEvent)                         // property event
	SetOnStartDrag(fn TStartDragEvent)                             // property event
}

// TTreeView Parent: TCustomTreeView
type TTreeView struct {
	TCustomTreeView
	additionPtr               uintptr
	advancedCustomDrawPtr     uintptr
	advancedCustomDrawItemPtr uintptr
	changePtr                 uintptr
	changingPtr               uintptr
	collapsedPtr              uintptr
	collapsingPtr             uintptr
	comparePtr                uintptr
	contextPopupPtr           uintptr
	createNodeClassPtr        uintptr
	customCreateItemPtr       uintptr
	customDrawPtr             uintptr
	customDrawItemPtr         uintptr
	customDrawArrowPtr        uintptr
	dblClickPtr               uintptr
	deletionPtr               uintptr
	dragDropPtr               uintptr
	dragOverPtr               uintptr
	editedPtr                 uintptr
	editingPtr                uintptr
	editingEndPtr             uintptr
	endDragPtr                uintptr
	expandedPtr               uintptr
	expandingPtr              uintptr
	getImageIndexPtr          uintptr
	getSelectedIndexPtr       uintptr
	hasChildrenPtr            uintptr
	mouseDownPtr              uintptr
	mouseEnterPtr             uintptr
	mouseLeavePtr             uintptr
	mouseMovePtr              uintptr
	mouseUpPtr                uintptr
	mouseWheelPtr             uintptr
	mouseWheelDownPtr         uintptr
	mouseWheelUpPtr           uintptr
	mouseWheelHorzPtr         uintptr
	mouseWheelLeftPtr         uintptr
	mouseWheelRightPtr        uintptr
	nodeChangedPtr            uintptr
	selectionChangedPtr       uintptr
	startDragPtr              uintptr
}

func NewTreeView(AnOwner IComponent) ITreeView {
	r1 := LCL().SysCallN(5059, GetObjectUintptr(AnOwner))
	return AsTreeView(r1)
}

func (m *TTreeView) AutoExpand() bool {
	r1 := LCL().SysCallN(5057, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetAutoExpand(AValue bool) {
	LCL().SysCallN(5057, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) DisabledFontColor() TColor {
	r1 := LCL().SysCallN(5060, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTreeView) SetDisabledFontColor(AValue TColor) {
	LCL().SysCallN(5060, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) DragKind() TDragKind {
	r1 := LCL().SysCallN(5062, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TTreeView) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5062, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) DragCursor() TCursor {
	r1 := LCL().SysCallN(5061, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TTreeView) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5061, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) DragMode() TDragMode {
	r1 := LCL().SysCallN(5063, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TTreeView) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5063, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) HideSelection() bool {
	r1 := LCL().SysCallN(5064, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetHideSelection(AValue bool) {
	LCL().SysCallN(5064, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) HotTrack() bool {
	r1 := LCL().SysCallN(5065, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetHotTrack(AValue bool) {
	LCL().SysCallN(5065, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) HotTrackColor() TColor {
	r1 := LCL().SysCallN(5066, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TTreeView) SetHotTrackColor(AValue TColor) {
	LCL().SysCallN(5066, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) Indent() int32 {
	r1 := LCL().SysCallN(5067, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeView) SetIndent(AValue int32) {
	LCL().SysCallN(5067, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) MultiSelect() bool {
	r1 := LCL().SysCallN(5068, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetMultiSelect(AValue bool) {
	LCL().SysCallN(5068, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ParentColor() bool {
	r1 := LCL().SysCallN(5069, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetParentColor(AValue bool) {
	LCL().SysCallN(5069, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ParentFont() bool {
	r1 := LCL().SysCallN(5070, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetParentFont(AValue bool) {
	LCL().SysCallN(5070, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ParentShowHint() bool {
	r1 := LCL().SysCallN(5071, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5071, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ReadOnly() bool {
	r1 := LCL().SysCallN(5072, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetReadOnly(AValue bool) {
	LCL().SysCallN(5072, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) RightClickSelect() bool {
	r1 := LCL().SysCallN(5073, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetRightClickSelect(AValue bool) {
	LCL().SysCallN(5073, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) RowSelect() bool {
	r1 := LCL().SysCallN(5074, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetRowSelect(AValue bool) {
	LCL().SysCallN(5074, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowButtons() bool {
	r1 := LCL().SysCallN(5116, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowButtons(AValue bool) {
	LCL().SysCallN(5116, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowLines() bool {
	r1 := LCL().SysCallN(5117, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowLines(AValue bool) {
	LCL().SysCallN(5117, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowRoot() bool {
	r1 := LCL().SysCallN(5118, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowRoot(AValue bool) {
	LCL().SysCallN(5118, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) ShowSeparators() bool {
	r1 := LCL().SysCallN(5119, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetShowSeparators(AValue bool) {
	LCL().SysCallN(5119, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeView) SortType() TSortType {
	r1 := LCL().SysCallN(5120, 0, m.Instance(), 0)
	return TSortType(r1)
}

func (m *TTreeView) SetSortType(AValue TSortType) {
	LCL().SysCallN(5120, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeView) ToolTips() bool {
	r1 := LCL().SysCallN(5121, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeView) SetToolTips(AValue bool) {
	LCL().SysCallN(5121, 1, m.Instance(), PascalBool(AValue))
}

func TreeViewClass() TClass {
	ret := LCL().SysCallN(5058)
	return TClass(ret)
}

func (m *TTreeView) SetOnAddition(fn TTVExpandedEvent) {
	if m.additionPtr != 0 {
		RemoveEventElement(m.additionPtr)
	}
	m.additionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5075, m.Instance(), m.additionPtr)
}

func (m *TTreeView) SetOnAdvancedCustomDraw(fn TTVAdvancedCustomDrawEvent) {
	if m.advancedCustomDrawPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawPtr)
	}
	m.advancedCustomDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5076, m.Instance(), m.advancedCustomDrawPtr)
}

func (m *TTreeView) SetOnAdvancedCustomDrawItem(fn TTVAdvancedCustomDrawItemEvent) {
	if m.advancedCustomDrawItemPtr != 0 {
		RemoveEventElement(m.advancedCustomDrawItemPtr)
	}
	m.advancedCustomDrawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5077, m.Instance(), m.advancedCustomDrawItemPtr)
}

func (m *TTreeView) SetOnChange(fn TTVChangedEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5078, m.Instance(), m.changePtr)
}

func (m *TTreeView) SetOnChanging(fn TTVChangingEvent) {
	if m.changingPtr != 0 {
		RemoveEventElement(m.changingPtr)
	}
	m.changingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5079, m.Instance(), m.changingPtr)
}

func (m *TTreeView) SetOnCollapsed(fn TTVExpandedEvent) {
	if m.collapsedPtr != 0 {
		RemoveEventElement(m.collapsedPtr)
	}
	m.collapsedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5080, m.Instance(), m.collapsedPtr)
}

func (m *TTreeView) SetOnCollapsing(fn TTVCollapsingEvent) {
	if m.collapsingPtr != 0 {
		RemoveEventElement(m.collapsingPtr)
	}
	m.collapsingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5081, m.Instance(), m.collapsingPtr)
}

func (m *TTreeView) SetOnCompare(fn TTVCompareEvent) {
	if m.comparePtr != 0 {
		RemoveEventElement(m.comparePtr)
	}
	m.comparePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5082, m.Instance(), m.comparePtr)
}

func (m *TTreeView) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5083, m.Instance(), m.contextPopupPtr)
}

func (m *TTreeView) SetOnCreateNodeClass(fn TTVCreateNodeClassEvent) {
	if m.createNodeClassPtr != 0 {
		RemoveEventElement(m.createNodeClassPtr)
	}
	m.createNodeClassPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5084, m.Instance(), m.createNodeClassPtr)
}

func (m *TTreeView) SetOnCustomCreateItem(fn TTVCustomCreateNodeEvent) {
	if m.customCreateItemPtr != 0 {
		RemoveEventElement(m.customCreateItemPtr)
	}
	m.customCreateItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5085, m.Instance(), m.customCreateItemPtr)
}

func (m *TTreeView) SetOnCustomDraw(fn TTVCustomDrawEvent) {
	if m.customDrawPtr != 0 {
		RemoveEventElement(m.customDrawPtr)
	}
	m.customDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5086, m.Instance(), m.customDrawPtr)
}

func (m *TTreeView) SetOnCustomDrawItem(fn TTVCustomDrawItemEvent) {
	if m.customDrawItemPtr != 0 {
		RemoveEventElement(m.customDrawItemPtr)
	}
	m.customDrawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5088, m.Instance(), m.customDrawItemPtr)
}

func (m *TTreeView) SetOnCustomDrawArrow(fn TTVCustomDrawArrowEvent) {
	if m.customDrawArrowPtr != 0 {
		RemoveEventElement(m.customDrawArrowPtr)
	}
	m.customDrawArrowPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5087, m.Instance(), m.customDrawArrowPtr)
}

func (m *TTreeView) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5089, m.Instance(), m.dblClickPtr)
}

func (m *TTreeView) SetOnDeletion(fn TTVExpandedEvent) {
	if m.deletionPtr != 0 {
		RemoveEventElement(m.deletionPtr)
	}
	m.deletionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5090, m.Instance(), m.deletionPtr)
}

func (m *TTreeView) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5091, m.Instance(), m.dragDropPtr)
}

func (m *TTreeView) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5092, m.Instance(), m.dragOverPtr)
}

func (m *TTreeView) SetOnEdited(fn TTVEditedEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5093, m.Instance(), m.editedPtr)
}

func (m *TTreeView) SetOnEditing(fn TTVEditingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5094, m.Instance(), m.editingPtr)
}

func (m *TTreeView) SetOnEditingEnd(fn TTVEditingEndEvent) {
	if m.editingEndPtr != 0 {
		RemoveEventElement(m.editingEndPtr)
	}
	m.editingEndPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5095, m.Instance(), m.editingEndPtr)
}

func (m *TTreeView) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5096, m.Instance(), m.endDragPtr)
}

func (m *TTreeView) SetOnExpanded(fn TTVExpandedEvent) {
	if m.expandedPtr != 0 {
		RemoveEventElement(m.expandedPtr)
	}
	m.expandedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5097, m.Instance(), m.expandedPtr)
}

func (m *TTreeView) SetOnExpanding(fn TTVExpandingEvent) {
	if m.expandingPtr != 0 {
		RemoveEventElement(m.expandingPtr)
	}
	m.expandingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5098, m.Instance(), m.expandingPtr)
}

func (m *TTreeView) SetOnGetImageIndex(fn TTVExpandedEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5099, m.Instance(), m.getImageIndexPtr)
}

func (m *TTreeView) SetOnGetSelectedIndex(fn TTVExpandedEvent) {
	if m.getSelectedIndexPtr != 0 {
		RemoveEventElement(m.getSelectedIndexPtr)
	}
	m.getSelectedIndexPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5100, m.Instance(), m.getSelectedIndexPtr)
}

func (m *TTreeView) SetOnHasChildren(fn TTVHasChildrenEvent) {
	if m.hasChildrenPtr != 0 {
		RemoveEventElement(m.hasChildrenPtr)
	}
	m.hasChildrenPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5101, m.Instance(), m.hasChildrenPtr)
}

func (m *TTreeView) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5102, m.Instance(), m.mouseDownPtr)
}

func (m *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5103, m.Instance(), m.mouseEnterPtr)
}

func (m *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5104, m.Instance(), m.mouseLeavePtr)
}

func (m *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5105, m.Instance(), m.mouseMovePtr)
}

func (m *TTreeView) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5106, m.Instance(), m.mouseUpPtr)
}

func (m *TTreeView) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5107, m.Instance(), m.mouseWheelPtr)
}

func (m *TTreeView) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5108, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TTreeView) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5112, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TTreeView) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5109, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TTreeView) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5110, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TTreeView) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5111, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TTreeView) SetOnNodeChanged(fn TTVNodeChangedEvent) {
	if m.nodeChangedPtr != 0 {
		RemoveEventElement(m.nodeChangedPtr)
	}
	m.nodeChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5113, m.Instance(), m.nodeChangedPtr)
}

func (m *TTreeView) SetOnSelectionChanged(fn TNotifyEvent) {
	if m.selectionChangedPtr != 0 {
		RemoveEventElement(m.selectionChangedPtr)
	}
	m.selectionChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5114, m.Instance(), m.selectionChangedPtr)
}

func (m *TTreeView) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5115, m.Instance(), m.startDragPtr)
}
