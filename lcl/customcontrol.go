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
	r1 := LCL().SysCallN(1473, GetObjectUintptr(AOwner))
	return AsCustomControl(r1)
}

func (m *TCustomControl) Canvas() ICanvas {
	r1 := LCL().SysCallN(1471, 0, m.Instance(), 0)
	return AsCanvas(r1)
}

func (m *TCustomControl) SetCanvas(AValue ICanvas) {
	LCL().SysCallN(1471, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomControl) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1470, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomControl) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1470, 1, m.Instance(), uintptr(AValue))
}

func CustomControlClass() TClass {
	ret := LCL().SysCallN(1472)
	return TClass(ret)
}

func (m *TCustomControl) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1474, m.Instance(), m.paintPtr)
}
