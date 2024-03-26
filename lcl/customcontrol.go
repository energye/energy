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

// ICustomControl Parent: IWinControl
type ICustomControl interface {
	IWinControl
	Canvas() ICanvas                    // property
	SetCanvas(AValue ICanvas)           // property
	BorderStyle() TBorderStyle          // property
	SetBorderStyle(AValue TBorderStyle) // property
	SetOnPaint(fn TNotifyEvent)         // property event
}

// TCustomControl Parent: TWinControl
type TCustomControl struct {
	TWinControl
	paintPtr uintptr
}

func NewCustomControl(AOwner IComponent) ICustomControl {
	r1 := LCL().SysCallN(1283, GetObjectUintptr(AOwner))
	return AsCustomControl(r1)
}

func (m *TCustomControl) Canvas() ICanvas {
	r1 := LCL().SysCallN(1281, 0, m.Instance(), 0)
	return AsCanvas(r1)
}

func (m *TCustomControl) SetCanvas(AValue ICanvas) {
	LCL().SysCallN(1281, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomControl) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1280, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomControl) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1280, 1, m.Instance(), uintptr(AValue))
}

func CustomControlClass() TClass {
	ret := LCL().SysCallN(1282)
	return TClass(ret)
}

func (m *TCustomControl) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1284, m.Instance(), m.paintPtr)
}
