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

// IForm Parent: ICustomForm
type IForm interface {
	ICustomForm
	EnabledMaximize(v bool)
	EnabledMinimize(v bool)
	EnabledSystemMenu(v bool)
	ScreenCenter()
	WorkAreaCenter()
	InheritedWndProc(heMessage *TMessage)
	SetOnWndProc(fn TWndProcEvent)
	SetGoPtr(ptr uintptr)
	ScaleSelf()
	ScaleForPPI(newPPI int32)
	ScaleForCurrentDpi()
	ClientHandle() HWND                                // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	SessionProperties() string                         // property
	SetSessionProperties(AValue string)                // property
	LCLVersion() string                                // property
	SetLCLVersion(AValue string)                       // property
	Cascade()                                          // procedure
	Next()                                             // procedure
	Previous()                                         // procedure
	Tile()                                             // procedure
	ArrangeIcons()                                     // procedure
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)             // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
}

// TForm Parent: TCustomForm
type TForm struct {
	TCustomForm
	constrainedResizePtr uintptr
	contextPopupPtr      uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	endDockPtr           uintptr
	getSiteInfoPtr       uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
	mouseWheelHorzPtr    uintptr
	mouseWheelLeftPtr    uintptr
	mouseWheelRightPtr   uintptr
	startDockPtr         uintptr
}

func NewForm(TheOwner IComponent) IForm {
	r1 := LCL().SysCallN(2859, GetObjectUintptr(TheOwner))
	return AsForm(r1)
}

func (m *TForm) ClientHandle() HWND {
	r1 := LCL().SysCallN(2858, m.Instance())
	return HWND(r1)
}

func (m *TForm) DragKind() TDragKind {
	r1 := LCL().SysCallN(2860, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TForm) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(2860, 1, m.Instance(), uintptr(AValue))
}

func (m *TForm) DragMode() TDragMode {
	r1 := LCL().SysCallN(2861, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TForm) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(2861, 1, m.Instance(), uintptr(AValue))
}

func (m *TForm) SessionProperties() string {
	r1 := LCL().SysCallN(2865, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TForm) SetSessionProperties(AValue string) {
	LCL().SysCallN(2865, 1, m.Instance(), PascalStr(AValue))
}

func (m *TForm) LCLVersion() string {
	r1 := LCL().SysCallN(2862, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TForm) SetLCLVersion(AValue string) {
	LCL().SysCallN(2862, 1, m.Instance(), PascalStr(AValue))
}

func FormClass() TClass {
	ret := LCL().SysCallN(2857)
	return TClass(ret)
}

func (m *TForm) Cascade() {
	LCL().SysCallN(2856, m.Instance())
}

func (m *TForm) Next() {
	LCL().SysCallN(2863, m.Instance())
}

func (m *TForm) Previous() {
	LCL().SysCallN(2864, m.Instance())
}

func (m *TForm) Tile() {
	LCL().SysCallN(2885, m.Instance())
}

func (m *TForm) ArrangeIcons() {
	LCL().SysCallN(2855, m.Instance())
}

func (m *TForm) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2866, m.Instance(), m.constrainedResizePtr)
}

func (m *TForm) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2867, m.Instance(), m.contextPopupPtr)
}

func (m *TForm) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2868, m.Instance(), m.dblClickPtr)
}

func (m *TForm) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2869, m.Instance(), m.dragDropPtr)
}

func (m *TForm) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2870, m.Instance(), m.dragOverPtr)
}

func (m *TForm) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2871, m.Instance(), m.endDockPtr)
}

func (m *TForm) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2872, m.Instance(), m.getSiteInfoPtr)
}

func (m *TForm) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2873, m.Instance(), m.mouseDownPtr)
}

func (m *TForm) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2874, m.Instance(), m.mouseEnterPtr)
}

func (m *TForm) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2875, m.Instance(), m.mouseLeavePtr)
}

func (m *TForm) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2876, m.Instance(), m.mouseMovePtr)
}

func (m *TForm) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2877, m.Instance(), m.mouseUpPtr)
}

func (m *TForm) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2878, m.Instance(), m.mouseWheelPtr)
}

func (m *TForm) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2879, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TForm) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2883, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TForm) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2880, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TForm) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2881, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TForm) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2882, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TForm) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2884, m.Instance(), m.startDockPtr)
}
