//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package opengl

import (
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	. "github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
)

// TOpenGLControl Parent: TCustomOpenGLControl
type TOpenGLControl struct {
	*TCustomOpenGLControl
	constrainedResizePtr uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
}

func NewOpenGLControl(TheOwner lcl.IComponent) *TOpenGLControl {
	r1 := openGLDllTableAPI().SysCallN(0, TheOwner.Instance())
	r := &TOpenGLControl{TCustomOpenGLControl: &TCustomOpenGLControl{TWinControl: lcl.AsWinControl(r1)}}
	return r
}

func (m *TOpenGLControl) SetOnConstrainedResize(fn lcl.TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(1, m.Instance(), m.constrainedResizePtr)
}

func (m *TOpenGLControl) SetOnDblClick(fn lcl.TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(2, m.Instance(), m.dblClickPtr)
}

func (m *TOpenGLControl) SetOnDragDrop(fn lcl.TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(3, m.Instance(), m.dragDropPtr)
}

func (m *TOpenGLControl) SetOnDragOver(fn lcl.TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(4, m.Instance(), m.dragOverPtr)
}

func (m *TOpenGLControl) SetOnMouseDown(fn lcl.TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(5, m.Instance(), m.mouseDownPtr)
}

func (m *TOpenGLControl) SetOnMouseEnter(fn lcl.TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(6, m.Instance(), m.mouseEnterPtr)
}

func (m *TOpenGLControl) SetOnMouseLeave(fn lcl.TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(7, m.Instance(), m.mouseLeavePtr)
}

func (m *TOpenGLControl) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(8, m.Instance(), m.mouseMovePtr)
}

func (m *TOpenGLControl) SetOnMouseUp(fn lcl.TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(9, m.Instance(), m.mouseUpPtr)
}

func (m *TOpenGLControl) SetOnMouseWheel(fn lcl.TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(10, m.Instance(), m.mouseWheelPtr)
}

func (m *TOpenGLControl) SetOnMouseWheelDown(fn lcl.TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(11, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TOpenGLControl) SetOnMouseWheelUp(fn lcl.TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	openGLDllTableAPI().SysCallN(12, m.Instance(), m.mouseWheelUpPtr)
}

var (
	openGLDllTable *imports.DllTable
)

func openGLDllTableAPI() *imports.DllTable {
	if openGLDllTable == nil {
		var importDefs = []*dllimports.ImportTable{
			dllimports.NewEnergyImport("OpenGLControl_Create", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnConstrainedResize", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnDblClick", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnDragDrop", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnDragOver", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseDown", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseEnter", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseLeave", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseMove", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseUp", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseWheel", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseWheelDown", 0),
			dllimports.NewEnergyImport("OpenGLControl_SetOnMouseWheelUp", 0),
		}
		openGLDllTable = new(imports.DllTable)
		openGLDllTable.SetOk(true)
		openGLDllTable.SetDll(imports.LibLCLExt().Dll())
		openGLDllTable.SetImportTable(importDefs)
	}
	return openGLDllTable
}
