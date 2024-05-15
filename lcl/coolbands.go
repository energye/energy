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

// ICoolBands Parent: ICollection
type ICoolBands interface {
	ICollection
	ItemsForCoolBand(Index int32) ICoolBand            // property
	SetItemsForCoolBand(Index int32, AValue ICoolBand) // property
	AddForCoolBand() ICoolBand                         // function
	FindBand(AControl IControl) ICoolBand              // function
	FindBandIndex(AControl IControl) int32             // function
}

// TCoolBands Parent: TCollection
type TCoolBands struct {
	TCollection
}

func NewCoolBands(ACoolBar ICustomCoolBar) ICoolBands {
	r1 := LCL().SysCallN(1171, GetObjectUintptr(ACoolBar))
	return AsCoolBands(r1)
}

func (m *TCoolBands) ItemsForCoolBand(Index int32) ICoolBand {
	r1 := LCL().SysCallN(1174, 0, m.Instance(), uintptr(Index))
	return AsCoolBand(r1)
}

func (m *TCoolBands) SetItemsForCoolBand(Index int32, AValue ICoolBand) {
	LCL().SysCallN(1174, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TCoolBands) AddForCoolBand() ICoolBand {
	r1 := LCL().SysCallN(1169, m.Instance())
	return AsCoolBand(r1)
}

func (m *TCoolBands) FindBand(AControl IControl) ICoolBand {
	r1 := LCL().SysCallN(1172, m.Instance(), GetObjectUintptr(AControl))
	return AsCoolBand(r1)
}

func (m *TCoolBands) FindBandIndex(AControl IControl) int32 {
	r1 := LCL().SysCallN(1173, m.Instance(), GetObjectUintptr(AControl))
	return int32(r1)
}

func CoolBandsClass() TClass {
	ret := LCL().SysCallN(1170)
	return TClass(ret)
}
