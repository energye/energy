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

// ICustomColorListBox Parent: ICustomListBox
type ICustomColorListBox interface {
	ICustomListBox
	ColorRectWidth() int32                          // property
	SetColorRectWidth(AValue int32)                 // property
	ColorRectOffset() int32                         // property
	SetColorRectOffset(AValue int32)                // property
	StyleForColorBoxStyle() TColorBoxStyle          // property
	SetStyleForColorBoxStyle(AValue TColorBoxStyle) // property
	Colors(Index int32) TColor                      // property
	SetColors(Index int32, AValue TColor)           // property
	ColorNames(Index int32) string                  // property
	SelectedForColor() TColor                       // property
	SetSelectedForColor(AValue TColor)              // property
	DefaultColorColor() TColor                      // property
	SetDefaultColorColor(AValue TColor)             // property
	NoneColorColor() TColor                         // property
	SetNoneColorColor(AValue TColor)                // property
	ColorDialog() IColorDialog                      // property
	SetColorDialog(AValue IColorDialog)             // property
	SetOnGetColors(fn TLBGetColorsEvent)            // property event
}

// TCustomColorListBox Parent: TCustomListBox
type TCustomColorListBox struct {
	TCustomListBox
	getColorsPtr uintptr
}

func NewCustomColorListBox(AOwner IComponent) ICustomColorListBox {
	r1 := LCL().SysCallN(1419, GetObjectUintptr(AOwner))
	return AsCustomColorListBox(r1)
}

func (m *TCustomColorListBox) ColorRectWidth() int32 {
	r1 := LCL().SysCallN(1417, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorListBox) SetColorRectWidth(AValue int32) {
	LCL().SysCallN(1417, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) ColorRectOffset() int32 {
	r1 := LCL().SysCallN(1416, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomColorListBox) SetColorRectOffset(AValue int32) {
	LCL().SysCallN(1416, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) StyleForColorBoxStyle() TColorBoxStyle {
	r1 := LCL().SysCallN(1424, 0, m.Instance(), 0)
	return TColorBoxStyle(r1)
}

func (m *TCustomColorListBox) SetStyleForColorBoxStyle(AValue TColorBoxStyle) {
	LCL().SysCallN(1424, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) Colors(Index int32) TColor {
	r1 := LCL().SysCallN(1418, 0, m.Instance(), uintptr(Index))
	return TColor(r1)
}

func (m *TCustomColorListBox) SetColors(Index int32, AValue TColor) {
	LCL().SysCallN(1418, 1, m.Instance(), uintptr(Index), uintptr(AValue))
}

func (m *TCustomColorListBox) ColorNames(Index int32) string {
	r1 := LCL().SysCallN(1415, m.Instance(), uintptr(Index))
	return GoStr(r1)
}

func (m *TCustomColorListBox) SelectedForColor() TColor {
	r1 := LCL().SysCallN(1422, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorListBox) SetSelectedForColor(AValue TColor) {
	LCL().SysCallN(1422, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) DefaultColorColor() TColor {
	r1 := LCL().SysCallN(1420, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorListBox) SetDefaultColorColor(AValue TColor) {
	LCL().SysCallN(1420, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) NoneColorColor() TColor {
	r1 := LCL().SysCallN(1421, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomColorListBox) SetNoneColorColor(AValue TColor) {
	LCL().SysCallN(1421, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomColorListBox) ColorDialog() IColorDialog {
	r1 := LCL().SysCallN(1414, 0, m.Instance(), 0)
	return AsColorDialog(r1)
}

func (m *TCustomColorListBox) SetColorDialog(AValue IColorDialog) {
	LCL().SysCallN(1414, 1, m.Instance(), GetObjectUintptr(AValue))
}

func CustomColorListBoxClass() TClass {
	ret := LCL().SysCallN(1413)
	return TClass(ret)
}

func (m *TCustomColorListBox) SetOnGetColors(fn TLBGetColorsEvent) {
	if m.getColorsPtr != 0 {
		RemoveEventElement(m.getColorsPtr)
	}
	m.getColorsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1423, m.Instance(), m.getColorsPtr)
}
