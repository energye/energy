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

// ICustomPrintDialog Parent: ICommonDialog
type ICustomPrintDialog interface {
	ICommonDialog
	Collate() bool                         // property
	SetCollate(AValue bool)                // property
	Copies() int32                         // property
	SetCopies(AValue int32)                // property
	FromPage() int32                       // property
	SetFromPage(AValue int32)              // property
	MinPage() int32                        // property
	SetMinPage(AValue int32)               // property
	MaxPage() int32                        // property
	SetMaxPage(AValue int32)               // property
	Options() TPrintDialogOptions          // property
	SetOptions(AValue TPrintDialogOptions) // property
	PrintToFile() bool                     // property
	SetPrintToFile(AValue bool)            // property
	PrintRange() TPrintRange               // property
	SetPrintRange(AValue TPrintRange)      // property
	ToPage() int32                         // property
	SetToPage(AValue int32)                // property
}

// TCustomPrintDialog Parent: TCommonDialog
type TCustomPrintDialog struct {
	TCommonDialog
}

func NewCustomPrintDialog(TheOwner IComponent) ICustomPrintDialog {
	r1 := LCL().SysCallN(2129, GetObjectUintptr(TheOwner))
	return AsCustomPrintDialog(r1)
}

func (m *TCustomPrintDialog) Collate() bool {
	r1 := LCL().SysCallN(2127, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPrintDialog) SetCollate(AValue bool) {
	LCL().SysCallN(2127, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPrintDialog) Copies() int32 {
	r1 := LCL().SysCallN(2128, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetCopies(AValue int32) {
	LCL().SysCallN(2128, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) FromPage() int32 {
	r1 := LCL().SysCallN(2130, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetFromPage(AValue int32) {
	LCL().SysCallN(2130, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) MinPage() int32 {
	r1 := LCL().SysCallN(2132, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetMinPage(AValue int32) {
	LCL().SysCallN(2132, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) MaxPage() int32 {
	r1 := LCL().SysCallN(2131, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetMaxPage(AValue int32) {
	LCL().SysCallN(2131, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) Options() TPrintDialogOptions {
	r1 := LCL().SysCallN(2133, 0, m.Instance(), 0)
	return TPrintDialogOptions(r1)
}

func (m *TCustomPrintDialog) SetOptions(AValue TPrintDialogOptions) {
	LCL().SysCallN(2133, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) PrintToFile() bool {
	r1 := LCL().SysCallN(2135, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomPrintDialog) SetPrintToFile(AValue bool) {
	LCL().SysCallN(2135, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomPrintDialog) PrintRange() TPrintRange {
	r1 := LCL().SysCallN(2134, 0, m.Instance(), 0)
	return TPrintRange(r1)
}

func (m *TCustomPrintDialog) SetPrintRange(AValue TPrintRange) {
	LCL().SysCallN(2134, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomPrintDialog) ToPage() int32 {
	r1 := LCL().SysCallN(2136, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomPrintDialog) SetToPage(AValue int32) {
	LCL().SysCallN(2136, 1, m.Instance(), uintptr(AValue))
}

func CustomPrintDialogClass() TClass {
	ret := LCL().SysCallN(2126)
	return TClass(ret)
}
