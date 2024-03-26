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

// ILabel Parent: ICustomLabel
type ILabel interface {
	ICustomLabel
	Alignment() TAlignment                          // property
	SetAlignment(AValue TAlignment)                 // property
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	FocusControl() IWinControl                      // property
	SetFocusControl(AValue IWinControl)             // property
	Layout() TTextLayout                            // property
	SetLayout(AValue TTextLayout)                   // property
	OptimalFill() bool                              // property
	SetOptimalFill(AValue bool)                     // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	ShowAccelChar() bool                            // property
	SetShowAccelChar(AValue bool)                   // property
	Transparent() bool                              // property
	SetTransparent(AValue bool)                     // property
	WordWrap() bool                                 // property
	SetWordWrap(AValue bool)                        // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TLabel Parent: TCustomLabel
type TLabel struct {
	TCustomLabel
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewLabel(TheOwner IComponent) ILabel {
	r1 := LCL().SysCallN(3189, GetObjectUintptr(TheOwner))
	return AsLabel(r1)
}

func (m *TLabel) Alignment() TAlignment {
	r1 := LCL().SysCallN(3187, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLabel) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3187, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) DragCursor() TCursor {
	r1 := LCL().SysCallN(3190, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLabel) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3190, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) DragKind() TDragKind {
	r1 := LCL().SysCallN(3191, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLabel) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3191, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) DragMode() TDragMode {
	r1 := LCL().SysCallN(3192, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLabel) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3192, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) FocusControl() IWinControl {
	r1 := LCL().SysCallN(3193, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TLabel) SetFocusControl(AValue IWinControl) {
	LCL().SysCallN(3193, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLabel) Layout() TTextLayout {
	r1 := LCL().SysCallN(3194, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TLabel) SetLayout(AValue TTextLayout) {
	LCL().SysCallN(3194, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabel) OptimalFill() bool {
	r1 := LCL().SysCallN(3195, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetOptimalFill(AValue bool) {
	LCL().SysCallN(3195, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ParentColor() bool {
	r1 := LCL().SysCallN(3196, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetParentColor(AValue bool) {
	LCL().SysCallN(3196, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ParentFont() bool {
	r1 := LCL().SysCallN(3197, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetParentFont(AValue bool) {
	LCL().SysCallN(3197, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ParentShowHint() bool {
	r1 := LCL().SysCallN(3198, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3198, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) ShowAccelChar() bool {
	r1 := LCL().SysCallN(3216, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetShowAccelChar(AValue bool) {
	LCL().SysCallN(3216, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) Transparent() bool {
	r1 := LCL().SysCallN(3217, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetTransparent(AValue bool) {
	LCL().SysCallN(3217, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabel) WordWrap() bool {
	r1 := LCL().SysCallN(3218, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabel) SetWordWrap(AValue bool) {
	LCL().SysCallN(3218, 1, m.Instance(), PascalBool(AValue))
}

func LabelClass() TClass {
	ret := LCL().SysCallN(3188)
	return TClass(ret)
}

func (m *TLabel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3199, m.Instance(), m.contextPopupPtr)
}

func (m *TLabel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3200, m.Instance(), m.dblClickPtr)
}

func (m *TLabel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3201, m.Instance(), m.dragDropPtr)
}

func (m *TLabel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3202, m.Instance(), m.dragOverPtr)
}

func (m *TLabel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3203, m.Instance(), m.endDragPtr)
}

func (m *TLabel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3204, m.Instance(), m.mouseDownPtr)
}

func (m *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3205, m.Instance(), m.mouseEnterPtr)
}

func (m *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3206, m.Instance(), m.mouseLeavePtr)
}

func (m *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3207, m.Instance(), m.mouseMovePtr)
}

func (m *TLabel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3208, m.Instance(), m.mouseUpPtr)
}

func (m *TLabel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3209, m.Instance(), m.mouseWheelPtr)
}

func (m *TLabel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3210, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TLabel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3214, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TLabel) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3211, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TLabel) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3212, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TLabel) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3213, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TLabel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3215, m.Instance(), m.startDragPtr)
}
