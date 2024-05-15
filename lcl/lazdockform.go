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

// ILazDockForm Parent: ICustomForm
type ILazDockForm interface {
	ICustomForm
	MainControl() IControl                                       // property
	SetMainControl(AValue IControl)                              // property
	FindMainControlCandidate() IControl                          // function
	FindHeader(x, y int32, OutPart *TLazDockHeaderPart) IControl // function
	IsDockedControl(Control IControl) bool                       // function
	ControlHasTitle(Control IControl) bool                       // function
	GetTitleRect(Control IControl) (resultRect TRect)            // function
	GetTitleOrientation(Control IControl) TDockOrientation       // function
	UpdateCaption()                                              // procedure
}

// TLazDockForm Parent: TCustomForm
type TLazDockForm struct {
	TCustomForm
}

func NewLazDockForm(AOwner IComponent) ILazDockForm {
	r1 := LCL().SysCallN(3517, GetObjectUintptr(AOwner))
	return AsLazDockForm(r1)
}

func (m *TLazDockForm) MainControl() IControl {
	r1 := LCL().SysCallN(3523, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TLazDockForm) SetMainControl(AValue IControl) {
	LCL().SysCallN(3523, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazDockForm) FindMainControlCandidate() IControl {
	r1 := LCL().SysCallN(3519, m.Instance())
	return AsControl(r1)
}

func (m *TLazDockForm) FindHeader(x, y int32, OutPart *TLazDockHeaderPart) IControl {
	var result1 uintptr
	r1 := LCL().SysCallN(3518, m.Instance(), uintptr(x), uintptr(y), uintptr(unsafePointer(&result1)))
	*OutPart = TLazDockHeaderPart(result1)
	return AsControl(r1)
}

func (m *TLazDockForm) IsDockedControl(Control IControl) bool {
	r1 := LCL().SysCallN(3522, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TLazDockForm) ControlHasTitle(Control IControl) bool {
	r1 := LCL().SysCallN(3516, m.Instance(), GetObjectUintptr(Control))
	return GoBool(r1)
}

func (m *TLazDockForm) GetTitleRect(Control IControl) (resultRect TRect) {
	LCL().SysCallN(3521, m.Instance(), GetObjectUintptr(Control), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TLazDockForm) GetTitleOrientation(Control IControl) TDockOrientation {
	r1 := LCL().SysCallN(3520, m.Instance(), GetObjectUintptr(Control))
	return TDockOrientation(r1)
}

func LazDockFormClass() TClass {
	ret := LCL().SysCallN(3515)
	return TClass(ret)
}

func (m *TLazDockForm) UpdateCaption() {
	LCL().SysCallN(3524, m.Instance())
}
