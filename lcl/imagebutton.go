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

// IImageButton Parent: IGraphicControl
type IImageButton interface {
	IGraphicControl
	Caption() string                         // property
	SetCaption(AValue string)                // property
	DragCursor() TCursor                     // property
	SetDragCursor(AValue TCursor)            // property
	DragKind() TDragKind                     // property
	SetDragKind(AValue TDragKind)            // property
	DragMode() TDragMode                     // property
	SetDragMode(AValue TDragMode)            // property
	ImageCount() int32                       // property
	SetImageCount(AValue int32)              // property
	Orientation() TImageOrientation          // property
	SetOrientation(AValue TImageOrientation) // property
	ModalResult() TModalResult               // property
	SetModalResult(AValue TModalResult)      // property
	ParentShowHint() bool                    // property
	SetParentShowHint(AValue bool)           // property
	ParentFont() bool                        // property
	SetParentFont(AValue bool)               // property
	Picture() IPicture                       // property
	SetPicture(AValue IPicture)              // property
	ShowCaption() bool                       // property
	SetShowCaption(AValue bool)              // property
	Wordwarp() bool                          // property
	SetWordwarp(AValue bool)                 // property
	Click()                                  // procedure
	SetOnContextPopup(fn TContextPopupEvent) // property event
	SetOnDblClick(fn TNotifyEvent)           // property event
	SetOnDragDrop(fn TDragDropEvent)         // property event
	SetOnDragOver(fn TDragOverEvent)         // property event
	SetOnEndDock(fn TEndDragEvent)           // property event
	SetOnEndDrag(fn TEndDragEvent)           // property event
	SetOnMouseDown(fn TMouseEvent)           // property event
	SetOnMouseEnter(fn TNotifyEvent)         // property event
	SetOnMouseLeave(fn TNotifyEvent)         // property event
	SetOnMouseMove(fn TMouseMoveEvent)       // property event
	SetOnMouseUp(fn TMouseEvent)             // property event
}

// TImageButton Parent: TGraphicControl
type TImageButton struct {
	TGraphicControl
	contextPopupPtr uintptr
	dblClickPtr     uintptr
	dragDropPtr     uintptr
	dragOverPtr     uintptr
	endDockPtr      uintptr
	endDragPtr      uintptr
	mouseDownPtr    uintptr
	mouseEnterPtr   uintptr
	mouseLeavePtr   uintptr
	mouseMovePtr    uintptr
	mouseUpPtr      uintptr
}

func NewImageButton(AOwner IComponent) IImageButton {
	r1 := LCL().SysCallN(3355, GetObjectUintptr(AOwner))
	return AsImageButton(r1)
}

func (m *TImageButton) Caption() string {
	r1 := LCL().SysCallN(3352, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TImageButton) SetCaption(AValue string) {
	LCL().SysCallN(3352, 1, m.Instance(), PascalStr(AValue))
}

func (m *TImageButton) DragCursor() TCursor {
	r1 := LCL().SysCallN(3356, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TImageButton) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3356, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) DragKind() TDragKind {
	r1 := LCL().SysCallN(3357, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TImageButton) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3357, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) DragMode() TDragMode {
	r1 := LCL().SysCallN(3358, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TImageButton) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3358, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) ImageCount() int32 {
	r1 := LCL().SysCallN(3359, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TImageButton) SetImageCount(AValue int32) {
	LCL().SysCallN(3359, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) Orientation() TImageOrientation {
	r1 := LCL().SysCallN(3361, 0, m.Instance(), 0)
	return TImageOrientation(r1)
}

func (m *TImageButton) SetOrientation(AValue TImageOrientation) {
	LCL().SysCallN(3361, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) ModalResult() TModalResult {
	r1 := LCL().SysCallN(3360, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TImageButton) SetModalResult(AValue TModalResult) {
	LCL().SysCallN(3360, 1, m.Instance(), uintptr(AValue))
}

func (m *TImageButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(3363, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3363, 1, m.Instance(), PascalBool(AValue))
}

func (m *TImageButton) ParentFont() bool {
	r1 := LCL().SysCallN(3362, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetParentFont(AValue bool) {
	LCL().SysCallN(3362, 1, m.Instance(), PascalBool(AValue))
}

func (m *TImageButton) Picture() IPicture {
	r1 := LCL().SysCallN(3364, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TImageButton) SetPicture(AValue IPicture) {
	LCL().SysCallN(3364, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TImageButton) ShowCaption() bool {
	r1 := LCL().SysCallN(3376, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetShowCaption(AValue bool) {
	LCL().SysCallN(3376, 1, m.Instance(), PascalBool(AValue))
}

func (m *TImageButton) Wordwarp() bool {
	r1 := LCL().SysCallN(3377, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImageButton) SetWordwarp(AValue bool) {
	LCL().SysCallN(3377, 1, m.Instance(), PascalBool(AValue))
}

func ImageButtonClass() TClass {
	ret := LCL().SysCallN(3353)
	return TClass(ret)
}

func (m *TImageButton) Click() {
	LCL().SysCallN(3354, m.Instance())
}

func (m *TImageButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3365, m.Instance(), m.contextPopupPtr)
}

func (m *TImageButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3366, m.Instance(), m.dblClickPtr)
}

func (m *TImageButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3367, m.Instance(), m.dragDropPtr)
}

func (m *TImageButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3368, m.Instance(), m.dragOverPtr)
}

func (m *TImageButton) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3369, m.Instance(), m.endDockPtr)
}

func (m *TImageButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3370, m.Instance(), m.endDragPtr)
}

func (m *TImageButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3371, m.Instance(), m.mouseDownPtr)
}

func (m *TImageButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3372, m.Instance(), m.mouseEnterPtr)
}

func (m *TImageButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3373, m.Instance(), m.mouseLeavePtr)
}

func (m *TImageButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3374, m.Instance(), m.mouseMovePtr)
}

func (m *TImageButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3375, m.Instance(), m.mouseUpPtr)
}
