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

// ICustomColorBox Parent: ICustomComboBox
type ICustomColorBox interface {
	ICustomComboBox
	ColorRectWidth() int32                          // property
	SetColorRectWidth(AValue int32)                 // property
	ColorRectOffset() int32                         // property
	SetColorRectOffset(AValue int32)                // property
	StyleForColorBoxStyle() TColorBoxStyle          // property
	SetStyleForColorBoxStyle(AValue TColorBoxStyle) // property
	Colors(Index int32) TColor                      // property
	ColorNames(Index int32) string                  // property
	Selected() TColor                               // property
	SetSelected(AValue TColor)                      // property
	DefaultColorColor() TColor                      // property
	SetDefaultColorColor(AValue TColor)             // property
	NoneColorColor() TColor                         // property
	SetNoneColorColor(AValue TColor)                // property
	ColorDialog() IColorDialog                      // property
	SetColorDialog(AValue IColorDialog)             // property
	SetOnGetColors(fn TGetColorsEvent)              // property event
}

// TCustomColorBox Parent: TCustomComboBox
type TCustomColorBox struct {
	TCustomComboBox
	getColorsPtr uintptr
}

func NewCustomColorBox(AOwner IComponent) ICustomColorBox {
	r1 := LCL().SysCallN(1407, GetObjectUintptr(AOwner))
	return AsCustomColorBox(r1)
}

func (m *TCustomColorBox) ColorRectWidth() int32 {
	r1 := LCL().SysCallN(1405, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorBox) SetColorRectWidth(AValue int32) {
	LCL().SysCallN(1405, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) ColorRectOffset() int32 {
	r1 := LCL().SysCallN(1404, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorBox) SetColorRectOffset(AValue int32) {
	LCL().SysCallN(1404, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) StyleForColorBoxStyle() TColorBoxStyle {
	r1 := LCL().SysCallN(1412, 0, m.Instance(), 0)
	return TColorBoxStyle(r1)
}

func (m *TCustomColorBox) SetStyleForColorBoxStyle(AValue TColorBoxStyle) {
	LCL().SysCallN(1412, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) Colors(Index int32) TColor {
	r1 := LCL().SysCallN(1406, m.Instance(), uintptr(Index))
	return TColor(r1)
}

func (m *TCustomColorBox) ColorNames(Index int32) string {
	r1 := LCL().SysCallN(1403, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TCustomColorBox) Selected() TColor {
	r1 := LCL().SysCallN(1410, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorBox) SetSelected(AValue TColor) {
	LCL().SysCallN(1410, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) DefaultColorColor() TColor {
	r1 := LCL().SysCallN(1408, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorBox) SetDefaultColorColor(AValue TColor) {
	LCL().SysCallN(1408, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) NoneColorColor() TColor {
	r1 := LCL().SysCallN(1409, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorBox) SetNoneColorColor(AValue TColor) {
	LCL().SysCallN(1409, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorBox) ColorDialog() IColorDialog {
	r1 := LCL().SysCallN(1402, 0, m.Instance(), 0)
	return AsColorDialog(r1)
}

func (m *TCustomColorBox) SetColorDialog(AValue IColorDialog) {
	LCL().SysCallN(1402, 1, m.Instance(), GetObjectUintptr(AValue))
}

func CustomColorBoxClass() TClass {
	ret := LCL().SysCallN(1401)
	return TClass(ret)
}

func (m *TCustomColorBox) SetOnGetColors(fn TGetColorsEvent) {
	if m.getColorsPtr != 0 {
		RemoveEventElement(m.getColorsPtr)
	}
	m.getColorsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1411, m.Instance(), m.getColorsPtr)
}
