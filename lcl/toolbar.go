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

// IToolBar Parent: IToolWindow
type IToolBar interface {
	IToolWindow
	ButtonCount() int32                                    // property
	Buttons(Index int32) IToolButton                       // property
	ButtonList() IList                                     // property
	RowCount() int32                                       // property
	ButtonDropWidth() int32                                // property
	ButtonHeight() int32                                   // property
	SetButtonHeight(AValue int32)                          // property
	ButtonWidth() int32                                    // property
	SetButtonWidth(AValue int32)                           // property
	DisabledImages() ICustomImageList                      // property
	SetDisabledImages(AValue ICustomImageList)             // property
	DragCursor() TCursor                                   // property
	SetDragCursor(AValue TCursor)                          // property
	DragKind() TDragKind                                   // property
	SetDragKind(AValue TDragKind)                          // property
	DragMode() TDragMode                                   // property
	SetDragMode(AValue TDragMode)                          // property
	DropDownWidth() int32                                  // property
	SetDropDownWidth(AValue int32)                         // property
	Flat() bool                                            // property
	SetFlat(AValue bool)                                   // property
	HotImages() ICustomImageList                           // property
	SetHotImages(AValue ICustomImageList)                  // property
	Images() ICustomImageList                              // property
	SetImages(AValue ICustomImageList)                     // property
	ImagesWidth() int32                                    // property
	SetImagesWidth(AValue int32)                           // property
	Indent() int32                                         // property
	SetIndent(AValue int32)                                // property
	List() bool                                            // property
	SetList(AValue bool)                                   // property
	ParentColor() bool                                     // property
	SetParentColor(AValue bool)                            // property
	ParentFont() bool                                      // property
	SetParentFont(AValue bool)                             // property
	ParentShowHint() bool                                  // property
	SetParentShowHint(AValue bool)                         // property
	ShowCaptions() bool                                    // property
	SetShowCaptions(AValue bool)                           // property
	Transparent() bool                                     // property
	SetTransparent(AValue bool)                            // property
	Wrapable() bool                                        // property
	SetWrapable(AValue bool)                               // property
	GetEnumeratorForToolBarEnumerator() IToolBarEnumerator // function
	SetButtonSize(NewButtonWidth, NewButtonHeight int32)   // procedure
	SetOnContextPopup(fn TContextPopupEvent)               // property event
	SetOnDblClick(fn TNotifyEvent)                         // property event
	SetOnDragDrop(fn TDragDropEvent)                       // property event
	SetOnDragOver(fn TDragOverEvent)                       // property event
	SetOnPaintButton(fn TToolBarOnPaintButton)             // property event
	SetOnEndDrag(fn TEndDragEvent)                         // property event
	SetOnMouseDown(fn TMouseEvent)                         // property event
	SetOnMouseEnter(fn TNotifyEvent)                       // property event
	SetOnMouseLeave(fn TNotifyEvent)                       // property event
	SetOnMouseMove(fn TMouseMoveEvent)                     // property event
	SetOnMouseUp(fn TMouseEvent)                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                   // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)         // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)           // property event
	SetOnStartDrag(fn TStartDragEvent)                     // property event
}

// TToolBar Parent: TToolWindow
type TToolBar struct {
	TToolWindow
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	paintButtonPtr    uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewToolBar(TheOwner IComponent) IToolBar {
	r1 := LCL().SysCallN(5465, GetObjectUintptr(TheOwner))
	return AsToolBar(r1)
}

func (m *TToolBar) ButtonCount() int32 {
	r1 := LCL().SysCallN(5458, m.Instance())
	return int32(r1)
}

func (m *TToolBar) Buttons(Index int32) IToolButton {
	r1 := LCL().SysCallN(5463, m.Instance(), uintptr(Index))
	return AsToolButton(r1)
}

func (m *TToolBar) ButtonList() IList {
	r1 := LCL().SysCallN(5461, m.Instance())
	return AsList(r1)
}

func (m *TToolBar) RowCount() int32 {
	r1 := LCL().SysCallN(5481, m.Instance())
	return int32(r1)
}

func (m *TToolBar) ButtonDropWidth() int32 {
	r1 := LCL().SysCallN(5459, m.Instance())
	return int32(r1)
}

func (m *TToolBar) ButtonHeight() int32 {
	r1 := LCL().SysCallN(5460, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetButtonHeight(AValue int32) {
	LCL().SysCallN(5460, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) ButtonWidth() int32 {
	r1 := LCL().SysCallN(5462, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetButtonWidth(AValue int32) {
	LCL().SysCallN(5462, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DisabledImages() ICustomImageList {
	r1 := LCL().SysCallN(5466, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TToolBar) SetDisabledImages(AValue ICustomImageList) {
	LCL().SysCallN(5466, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(5467, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TToolBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5467, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(5468, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TToolBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5468, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(5469, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TToolBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5469, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) DropDownWidth() int32 {
	r1 := LCL().SysCallN(5470, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetDropDownWidth(AValue int32) {
	LCL().SysCallN(5470, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) Flat() bool {
	r1 := LCL().SysCallN(5471, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetFlat(AValue bool) {
	LCL().SysCallN(5471, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) HotImages() ICustomImageList {
	r1 := LCL().SysCallN(5473, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TToolBar) SetHotImages(AValue ICustomImageList) {
	LCL().SysCallN(5473, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolBar) Images() ICustomImageList {
	r1 := LCL().SysCallN(5474, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TToolBar) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(5474, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TToolBar) ImagesWidth() int32 {
	r1 := LCL().SysCallN(5475, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetImagesWidth(AValue int32) {
	LCL().SysCallN(5475, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) Indent() int32 {
	r1 := LCL().SysCallN(5476, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TToolBar) SetIndent(AValue int32) {
	LCL().SysCallN(5476, 1, m.Instance(), uintptr(AValue))
}

func (m *TToolBar) List() bool {
	r1 := LCL().SysCallN(5477, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetList(AValue bool) {
	LCL().SysCallN(5477, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ParentColor() bool {
	r1 := LCL().SysCallN(5478, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetParentColor(AValue bool) {
	LCL().SysCallN(5478, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ParentFont() bool {
	r1 := LCL().SysCallN(5479, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetParentFont(AValue bool) {
	LCL().SysCallN(5479, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(5480, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5480, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) ShowCaptions() bool {
	r1 := LCL().SysCallN(5498, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetShowCaptions(AValue bool) {
	LCL().SysCallN(5498, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) Transparent() bool {
	r1 := LCL().SysCallN(5499, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetTransparent(AValue bool) {
	LCL().SysCallN(5499, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) Wrapable() bool {
	r1 := LCL().SysCallN(5500, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToolBar) SetWrapable(AValue bool) {
	LCL().SysCallN(5500, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToolBar) GetEnumeratorForToolBarEnumerator() IToolBarEnumerator {
	r1 := LCL().SysCallN(5472, m.Instance())
	return AsToolBarEnumerator(r1)
}

func ToolBarClass() TClass {
	ret := LCL().SysCallN(5464)
	return TClass(ret)
}

func (m *TToolBar) SetButtonSize(NewButtonWidth, NewButtonHeight int32) {
	LCL().SysCallN(5482, m.Instance(), uintptr(NewButtonWidth), uintptr(NewButtonHeight))
}

func (m *TToolBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5483, m.Instance(), m.contextPopupPtr)
}

func (m *TToolBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5484, m.Instance(), m.dblClickPtr)
}

func (m *TToolBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5485, m.Instance(), m.dragDropPtr)
}

func (m *TToolBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5486, m.Instance(), m.dragOverPtr)
}

func (m *TToolBar) SetOnPaintButton(fn TToolBarOnPaintButton) {
	if m.paintButtonPtr != 0 {
		RemoveEventElement(m.paintButtonPtr)
	}
	m.paintButtonPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5496, m.Instance(), m.paintButtonPtr)
}

func (m *TToolBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5487, m.Instance(), m.endDragPtr)
}

func (m *TToolBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5488, m.Instance(), m.mouseDownPtr)
}

func (m *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5489, m.Instance(), m.mouseEnterPtr)
}

func (m *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5490, m.Instance(), m.mouseLeavePtr)
}

func (m *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5491, m.Instance(), m.mouseMovePtr)
}

func (m *TToolBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5492, m.Instance(), m.mouseUpPtr)
}

func (m *TToolBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5493, m.Instance(), m.mouseWheelPtr)
}

func (m *TToolBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5494, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TToolBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5495, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TToolBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5497, m.Instance(), m.startDragPtr)
}
