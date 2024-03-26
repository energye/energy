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

// IStatusBar Parent: IWinControl
type IStatusBar interface {
	IWinControl
	Canvas() ICanvas                                          // property
	AutoHint() bool                                           // property
	SetAutoHint(AValue bool)                                  // property
	DragCursor() TCursor                                      // property
	SetDragCursor(AValue TCursor)                             // property
	DragKind() TDragKind                                      // property
	SetDragKind(AValue TDragKind)                             // property
	DragMode() TDragMode                                      // property
	SetDragMode(AValue TDragMode)                             // property
	Panels() IStatusPanels                                    // property
	SetPanels(AValue IStatusPanels)                           // property
	ParentColor() bool                                        // property
	SetParentColor(AValue bool)                               // property
	ParentFont() bool                                         // property
	SetParentFont(AValue bool)                                // property
	ParentShowHint() bool                                     // property
	SetParentShowHint(AValue bool)                            // property
	SimpleText() string                                       // property
	SetSimpleText(AValue string)                              // property
	SimplePanel() bool                                        // property
	SetSimplePanel(AValue bool)                               // property
	SizeGrip() bool                                           // property
	SetSizeGrip(AValue bool)                                  // property
	UseSystemFont() bool                                      // property
	SetUseSystemFont(AValue bool)                             // property
	GetPanelIndexAt(X, Y int32) int32                         // function
	SizeGripEnabled() bool                                    // function
	UpdatingStatusBar() bool                                  // function
	InvalidatePanel(PanelIndex int32, PanelParts TPanelParts) // procedure
	BeginUpdate()                                             // procedure
	EndUpdate()                                               // procedure
	SetOnContextPopup(fn TContextPopupEvent)                  // property event
	SetOnCreatePanelClass(fn TSBCreatePanelClassEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                            // property event
	SetOnDragDrop(fn TDragDropEvent)                          // property event
	SetOnDragOver(fn TDragOverEvent)                          // property event
	SetOnDrawPanel(fn TDrawPanelEvent)                        // property event
	SetOnEndDock(fn TEndDragEvent)                            // property event
	SetOnEndDrag(fn TEndDragEvent)                            // property event
	SetOnHint(fn TNotifyEvent)                                // property event
	SetOnMouseDown(fn TMouseEvent)                            // property event
	SetOnMouseEnter(fn TNotifyEvent)                          // property event
	SetOnMouseLeave(fn TNotifyEvent)                          // property event
	SetOnMouseMove(fn TMouseMoveEvent)                        // property event
	SetOnMouseUp(fn TMouseEvent)                              // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                      // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)            // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)              // property event
	SetOnStartDock(fn TStartDockEvent)                        // property event
	SetOnStartDrag(fn TStartDragEvent)                        // property event
}

// TStatusBar Parent: TWinControl
type TStatusBar struct {
	TWinControl
	contextPopupPtr     uintptr
	createPanelClassPtr uintptr
	dblClickPtr         uintptr
	dragDropPtr         uintptr
	dragOverPtr         uintptr
	drawPanelPtr        uintptr
	endDockPtr          uintptr
	endDragPtr          uintptr
	hintPtr             uintptr
	mouseDownPtr        uintptr
	mouseEnterPtr       uintptr
	mouseLeavePtr       uintptr
	mouseMovePtr        uintptr
	mouseUpPtr          uintptr
	mouseWheelPtr       uintptr
	mouseWheelDownPtr   uintptr
	mouseWheelUpPtr     uintptr
	startDockPtr        uintptr
	startDragPtr        uintptr
}

func NewStatusBar(TheOwner IComponent) IStatusBar {
	r1 := LCL().SysCallN(4457, GetObjectUintptr(TheOwner))
	return AsStatusBar(r1)
}

func (m *TStatusBar) Canvas() ICanvas {
	r1 := LCL().SysCallN(4455, m.Instance())
	return AsCanvas(r1)
}

func (m *TStatusBar) AutoHint() bool {
	r1 := LCL().SysCallN(4453, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetAutoHint(AValue bool) {
	LCL().SysCallN(4453, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(4458, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStatusBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4458, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(4459, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStatusBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4459, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(4460, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStatusBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4460, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) Panels() IStatusPanels {
	r1 := LCL().SysCallN(4464, 0, m.Instance(), 0)
	return AsStatusPanels(r1)
}

func (m *TStatusBar) SetPanels(AValue IStatusPanels) {
	LCL().SysCallN(4464, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStatusBar) ParentColor() bool {
	r1 := LCL().SysCallN(4465, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentColor(AValue bool) {
	LCL().SysCallN(4465, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentFont() bool {
	r1 := LCL().SysCallN(4466, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentFont(AValue bool) {
	LCL().SysCallN(4466, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(4467, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4467, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SimpleText() string {
	r1 := LCL().SysCallN(4488, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStatusBar) SetSimpleText(AValue string) {
	LCL().SysCallN(4488, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStatusBar) SimplePanel() bool {
	r1 := LCL().SysCallN(4487, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSimplePanel(AValue bool) {
	LCL().SysCallN(4487, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SizeGrip() bool {
	r1 := LCL().SysCallN(4489, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSizeGrip(AValue bool) {
	LCL().SysCallN(4489, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) UseSystemFont() bool {
	r1 := LCL().SysCallN(4492, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetUseSystemFont(AValue bool) {
	LCL().SysCallN(4492, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) GetPanelIndexAt(X, Y int32) int32 {
	r1 := LCL().SysCallN(4462, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TStatusBar) SizeGripEnabled() bool {
	r1 := LCL().SysCallN(4490, m.Instance())
	return GoBool(r1)
}

func (m *TStatusBar) UpdatingStatusBar() bool {
	r1 := LCL().SysCallN(4491, m.Instance())
	return GoBool(r1)
}

func StatusBarClass() TClass {
	ret := LCL().SysCallN(4456)
	return TClass(ret)
}

func (m *TStatusBar) InvalidatePanel(PanelIndex int32, PanelParts TPanelParts) {
	LCL().SysCallN(4463, m.Instance(), uintptr(PanelIndex), uintptr(PanelParts))
}

func (m *TStatusBar) BeginUpdate() {
	LCL().SysCallN(4454, m.Instance())
}

func (m *TStatusBar) EndUpdate() {
	LCL().SysCallN(4461, m.Instance())
}

func (m *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4468, m.Instance(), m.contextPopupPtr)
}

func (m *TStatusBar) SetOnCreatePanelClass(fn TSBCreatePanelClassEvent) {
	if m.createPanelClassPtr != 0 {
		RemoveEventElement(m.createPanelClassPtr)
	}
	m.createPanelClassPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4469, m.Instance(), m.createPanelClassPtr)
}

func (m *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4470, m.Instance(), m.dblClickPtr)
}

func (m *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4471, m.Instance(), m.dragDropPtr)
}

func (m *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4472, m.Instance(), m.dragOverPtr)
}

func (m *TStatusBar) SetOnDrawPanel(fn TDrawPanelEvent) {
	if m.drawPanelPtr != 0 {
		RemoveEventElement(m.drawPanelPtr)
	}
	m.drawPanelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4473, m.Instance(), m.drawPanelPtr)
}

func (m *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4474, m.Instance(), m.endDockPtr)
}

func (m *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4475, m.Instance(), m.endDragPtr)
}

func (m *TStatusBar) SetOnHint(fn TNotifyEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4476, m.Instance(), m.hintPtr)
}

func (m *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4477, m.Instance(), m.mouseDownPtr)
}

func (m *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4478, m.Instance(), m.mouseEnterPtr)
}

func (m *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4479, m.Instance(), m.mouseLeavePtr)
}

func (m *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4480, m.Instance(), m.mouseMovePtr)
}

func (m *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4481, m.Instance(), m.mouseUpPtr)
}

func (m *TStatusBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4482, m.Instance(), m.mouseWheelPtr)
}

func (m *TStatusBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4483, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TStatusBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4484, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4485, m.Instance(), m.startDockPtr)
}

func (m *TStatusBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4486, m.Instance(), m.startDragPtr)
}
