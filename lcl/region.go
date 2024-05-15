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

// IRegion Parent: IGraphicsObject
type IRegion interface {
	IGraphicsObject
	ClipRect() (resultRect TRect)      // property
	SetClipRect(AValue *TRect)         // property
	AddRectangle(X1, Y1, X2, Y2 int32) // procedure
}

// TRegion Parent: TGraphicsObject
type TRegion struct {
	TGraphicsObject
}

func NewRegion() IRegion {
	r1 := LCL().SysCallN(4740)
	return AsRegion(r1)
}

func (m *TRegion) ClipRect() (resultRect TRect) {
	LCL().SysCallN(4739, 0, m.Instance(), uintptr(unsafePointer(&resultRect)), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TRegion) SetClipRect(AValue *TRect) {
	LCL().SysCallN(4739, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func RegionClass() TClass {
	ret := LCL().SysCallN(4738)
	return TClass(ret)
}

func (m *TRegion) AddRectangle(X1, Y1, X2, Y2 int32) {
	LCL().SysCallN(4737, m.Instance(), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
}
