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

// IFileDialog Parent: ICommonDialog
type IFileDialog interface {
	ICommonDialog
	Files() IStrings                          // property
	HistoryList() IStrings                    // property
	SetHistoryList(AValue IStrings)           // property
	DefaultExt() string                       // property
	SetDefaultExt(AValue string)              // property
	FileName() string                         // property
	SetFileName(AValue string)                // property
	Filter() string                           // property
	SetFilter(AValue string)                  // property
	FilterIndex() int32                       // property
	SetFilterIndex(AValue int32)              // property
	InitialDir() string                       // property
	SetInitialDir(AValue string)              // property
	DoTypeChange()                            // procedure
	IntfFileTypeChanged(NewFilterIndex int32) // procedure
	SetOnHelpClicked(fn TNotifyEvent)         // property event
	SetOnTypeChange(fn TNotifyEvent)          // property event
}

// TFileDialog Parent: TCommonDialog
type TFileDialog struct {
	TCommonDialog
	helpClickedPtr uintptr
	typeChangePtr  uintptr
}

func NewFileDialog(TheOwner IComponent) IFileDialog {
	r1 := LCL().SysCallN(2753, GetObjectUintptr(TheOwner))
	return AsFileDialog(r1)
}

func (m *TFileDialog) Files() IStrings {
	r1 := LCL().SysCallN(2757, m.Instance())
	return AsStrings(r1)
}

func (m *TFileDialog) HistoryList() IStrings {
	r1 := LCL().SysCallN(2760, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TFileDialog) SetHistoryList(AValue IStrings) {
	LCL().SysCallN(2760, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFileDialog) DefaultExt() string {
	r1 := LCL().SysCallN(2754, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetDefaultExt(AValue string) {
	LCL().SysCallN(2754, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) FileName() string {
	r1 := LCL().SysCallN(2756, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetFileName(AValue string) {
	LCL().SysCallN(2756, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) Filter() string {
	r1 := LCL().SysCallN(2758, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetFilter(AValue string) {
	LCL().SysCallN(2758, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) FilterIndex() int32 {
	r1 := LCL().SysCallN(2759, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFileDialog) SetFilterIndex(AValue int32) {
	LCL().SysCallN(2759, 1, m.Instance(), uintptr(AValue))
}

func (m *TFileDialog) InitialDir() string {
	r1 := LCL().SysCallN(2761, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetInitialDir(AValue string) {
	LCL().SysCallN(2761, 1, m.Instance(), PascalStr(AValue))
}

func FileDialogClass() TClass {
	ret := LCL().SysCallN(2752)
	return TClass(ret)
}

func (m *TFileDialog) DoTypeChange() {
	LCL().SysCallN(2755, m.Instance())
}

func (m *TFileDialog) IntfFileTypeChanged(NewFilterIndex int32) {
	LCL().SysCallN(2762, m.Instance(), uintptr(NewFilterIndex))
}

func (m *TFileDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if m.helpClickedPtr != 0 {
		RemoveEventElement(m.helpClickedPtr)
	}
	m.helpClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2763, m.Instance(), m.helpClickedPtr)
}

func (m *TFileDialog) SetOnTypeChange(fn TNotifyEvent) {
	if m.typeChangePtr != 0 {
		RemoveEventElement(m.typeChangePtr)
	}
	m.typeChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2764, m.Instance(), m.typeChangePtr)
}
