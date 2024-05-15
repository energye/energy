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

// IDividerBevel Parent: IGraphicControl
type IDividerBevel interface {
	IGraphicControl
	BevelStyle() TBevelStyle                    // property
	SetBevelStyle(AValue TBevelStyle)           // property
	BevelWidth() int32                          // property
	SetBevelWidth(AValue int32)                 // property
	CaptionSpacing() int32                      // property
	SetCaptionSpacing(AValue int32)             // property
	DragCursor() TCursor                        // property
	SetDragCursor(AValue TCursor)               // property
	DragKind() TDragKind                        // property
	SetDragKind(AValue TDragKind)               // property
	DragMode() TDragMode                        // property
	SetDragMode(AValue TDragMode)               // property
	LeftIndent() int32                          // property
	SetLeftIndent(AValue int32)                 // property
	Orientation() TTrackBarOrientation          // property
	SetOrientation(AValue TTrackBarOrientation) // property
	ParentColor() bool                          // property
	SetParentColor(AValue bool)                 // property
	ParentFont() bool                           // property
	SetParentFont(AValue bool)                  // property
	ParentShowHint() bool                       // property
	SetParentShowHint(AValue bool)              // property
	Style() TGrabStyle                          // property
	SetStyle(AValue TGrabStyle)                 // property
	Transparent() bool                          // property
	SetTransparent(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)    // property event
	SetOnDblClick(fn TNotifyEvent)              // property event
	SetOnDragDrop(fn TDragDropEvent)            // property event
	SetOnDragOver(fn TDragOverEvent)            // property event
	SetOnEndDrag(fn TEndDragEvent)              // property event
	SetOnMouseDown(fn TMouseEvent)              // property event
	SetOnMouseEnter(fn TNotifyEvent)            // property event
	SetOnMouseLeave(fn TNotifyEvent)            // property event
	SetOnMouseMove(fn TMouseMoveEvent)          // property event
	SetOnMouseUp(fn TMouseEvent)                // property event
	SetOnStartDrag(fn TStartDragEvent)          // property event
}

// TDividerBevel Parent: TGraphicControl
type TDividerBevel struct {
	TGraphicControl
	contextPopupPtr uintptr
	dblClickPtr     uintptr
	dragDropPtr     uintptr
	dragOverPtr     uintptr
	endDragPtr      uintptr
	mouseDownPtr    uintptr
	mouseEnterPtr   uintptr
	mouseLeavePtr   uintptr
	mouseMovePtr    uintptr
	mouseUpPtr      uintptr
	startDragPtr    uintptr
}

func NewDividerBevel(AOwner IComponent) IDividerBevel {
	r1 := LCL().SysCallN(2602, GetObjectUintptr(AOwner))
	return AsDividerBevel(r1)
}

func (m *TDividerBevel) BevelStyle() TBevelStyle {
	r1 := LCL().SysCallN(2598, 0, m.Instance(), 0)
	return TBevelStyle(r1)
}

func (m *TDividerBevel) SetBevelStyle(AValue TBevelStyle) {
	LCL().SysCallN(2598, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) BevelWidth() int32 {
	r1 := LCL().SysCallN(2599, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDividerBevel) SetBevelWidth(AValue int32) {
	LCL().SysCallN(2599, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) CaptionSpacing() int32 {
	r1 := LCL().SysCallN(2600, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDividerBevel) SetCaptionSpacing(AValue int32) {
	LCL().SysCallN(2600, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) DragCursor() TCursor {
	r1 := LCL().SysCallN(2603, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TDividerBevel) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(2603, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) DragKind() TDragKind {
	r1 := LCL().SysCallN(2604, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TDividerBevel) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(2604, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) DragMode() TDragMode {
	r1 := LCL().SysCallN(2605, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TDividerBevel) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2605, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) LeftIndent() int32 {
	r1 := LCL().SysCallN(2606, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDividerBevel) SetLeftIndent(AValue int32) {
	LCL().SysCallN(2606, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) Orientation() TTrackBarOrientation {
	r1 := LCL().SysCallN(2607, 0, m.Instance(), 0)
	return TTrackBarOrientation(r1)
}

func (m *TDividerBevel) SetOrientation(AValue TTrackBarOrientation) {
	LCL().SysCallN(2607, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) ParentColor() bool {
	r1 := LCL().SysCallN(2608, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetParentColor(AValue bool) {
	LCL().SysCallN(2608, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDividerBevel) ParentFont() bool {
	r1 := LCL().SysCallN(2609, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetParentFont(AValue bool) {
	LCL().SysCallN(2609, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDividerBevel) ParentShowHint() bool {
	r1 := LCL().SysCallN(2610, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(2610, 1, m.Instance(), PascalBool(AValue))
}

func (m *TDividerBevel) Style() TGrabStyle {
	r1 := LCL().SysCallN(2622, 0, m.Instance(), 0)
	return TGrabStyle(r1)
}

func (m *TDividerBevel) SetStyle(AValue TGrabStyle) {
	LCL().SysCallN(2622, 1, m.Instance(), uintptr(AValue))
}

func (m *TDividerBevel) Transparent() bool {
	r1 := LCL().SysCallN(2623, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDividerBevel) SetTransparent(AValue bool) {
	LCL().SysCallN(2623, 1, m.Instance(), PascalBool(AValue))
}

func DividerBevelClass() TClass {
	ret := LCL().SysCallN(2601)
	return TClass(ret)
}

func (m *TDividerBevel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2611, m.Instance(), m.contextPopupPtr)
}

func (m *TDividerBevel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2612, m.Instance(), m.dblClickPtr)
}

func (m *TDividerBevel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2613, m.Instance(), m.dragDropPtr)
}

func (m *TDividerBevel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2614, m.Instance(), m.dragOverPtr)
}

func (m *TDividerBevel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2615, m.Instance(), m.endDragPtr)
}

func (m *TDividerBevel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2616, m.Instance(), m.mouseDownPtr)
}

func (m *TDividerBevel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2617, m.Instance(), m.mouseEnterPtr)
}

func (m *TDividerBevel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2618, m.Instance(), m.mouseLeavePtr)
}

func (m *TDividerBevel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2619, m.Instance(), m.mouseMovePtr)
}

func (m *TDividerBevel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2620, m.Instance(), m.mouseUpPtr)
}

func (m *TDividerBevel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2621, m.Instance(), m.startDragPtr)
}
