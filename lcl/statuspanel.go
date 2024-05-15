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

// IStatusPanel Parent: ICollectionItem
type IStatusPanel interface {
	ICollectionItem
	Alignment() TAlignment             // property
	SetAlignment(AValue TAlignment)    // property
	Bevel() TStatusPanelBevel          // property
	SetBevel(AValue TStatusPanelBevel) // property
	BidiMode() TBiDiMode               // property
	SetBidiMode(AValue TBiDiMode)      // property
	ParentBiDiMode() bool              // property
	SetParentBiDiMode(AValue bool)     // property
	Style() TStatusPanelStyle          // property
	SetStyle(AValue TStatusPanelStyle) // property
	Text() string                      // property
	SetText(AValue string)             // property
	Width() int32                      // property
	SetWidth(AValue int32)             // property
	StatusBar() IStatusBar             // function
}

// TStatusPanel Parent: TCollectionItem
type TStatusPanel struct {
	TCollectionItem
}

func NewStatusPanel(ACollection ICollection) IStatusPanel {
	r1 := LCL().SysCallN(5146, GetObjectUintptr(ACollection))
	return AsStatusPanel(r1)
}

func (m *TStatusPanel) Alignment() TAlignment {
	r1 := LCL().SysCallN(5142, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TStatusPanel) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(5142, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) Bevel() TStatusPanelBevel {
	r1 := LCL().SysCallN(5143, 0, m.Instance(), 0)
	return TStatusPanelBevel(r1)
}

func (m *TStatusPanel) SetBevel(AValue TStatusPanelBevel) {
	LCL().SysCallN(5143, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) BidiMode() TBiDiMode {
	r1 := LCL().SysCallN(5144, 0, m.Instance(), 0)
	return TBiDiMode(r1)
}

func (m *TStatusPanel) SetBidiMode(AValue TBiDiMode) {
	LCL().SysCallN(5144, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) ParentBiDiMode() bool {
	r1 := LCL().SysCallN(5147, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusPanel) SetParentBiDiMode(AValue bool) {
	LCL().SysCallN(5147, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusPanel) Style() TStatusPanelStyle {
	r1 := LCL().SysCallN(5149, 0, m.Instance(), 0)
	return TStatusPanelStyle(r1)
}

func (m *TStatusPanel) SetStyle(AValue TStatusPanelStyle) {
	LCL().SysCallN(5149, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) Text() string {
	r1 := LCL().SysCallN(5150, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStatusPanel) SetText(AValue string) {
	LCL().SysCallN(5150, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStatusPanel) Width() int32 {
	r1 := LCL().SysCallN(5151, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TStatusPanel) SetWidth(AValue int32) {
	LCL().SysCallN(5151, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusPanel) StatusBar() IStatusBar {
	r1 := LCL().SysCallN(5148, m.Instance())
	return AsStatusBar(r1)
}

func StatusPanelClass() TClass {
	ret := LCL().SysCallN(5145)
	return TClass(ret)
}
