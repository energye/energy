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
	r1 := LCL().SysCallN(2996, GetObjectUintptr(TheOwner))
	return AsFileDialog(r1)
}

func (m *TFileDialog) Files() IStrings {
	r1 := LCL().SysCallN(3000, m.Instance())
	return AsStrings(r1)
}

func (m *TFileDialog) HistoryList() IStrings {
	r1 := LCL().SysCallN(3003, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TFileDialog) SetHistoryList(AValue IStrings) {
	LCL().SysCallN(3003, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFileDialog) DefaultExt() string {
	r1 := LCL().SysCallN(2997, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetDefaultExt(AValue string) {
	LCL().SysCallN(2997, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) FileName() string {
	r1 := LCL().SysCallN(2999, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetFileName(AValue string) {
	LCL().SysCallN(2999, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) Filter() string {
	r1 := LCL().SysCallN(3001, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetFilter(AValue string) {
	LCL().SysCallN(3001, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFileDialog) FilterIndex() int32 {
	r1 := LCL().SysCallN(3002, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFileDialog) SetFilterIndex(AValue int32) {
	LCL().SysCallN(3002, 1, m.Instance(), uintptr(AValue))
}

func (m *TFileDialog) InitialDir() string {
	r1 := LCL().SysCallN(3004, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFileDialog) SetInitialDir(AValue string) {
	LCL().SysCallN(3004, 1, m.Instance(), PascalStr(AValue))
}

func FileDialogClass() TClass {
	ret := LCL().SysCallN(2995)
	return TClass(ret)
}

func (m *TFileDialog) DoTypeChange() {
	LCL().SysCallN(2998, m.Instance())
}

func (m *TFileDialog) IntfFileTypeChanged(NewFilterIndex int32) {
	LCL().SysCallN(3005, m.Instance(), uintptr(NewFilterIndex))
}

func (m *TFileDialog) SetOnHelpClicked(fn TNotifyEvent) {
	if m.helpClickedPtr != 0 {
		RemoveEventElement(m.helpClickedPtr)
	}
	m.helpClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3006, m.Instance(), m.helpClickedPtr)
}

func (m *TFileDialog) SetOnTypeChange(fn TNotifyEvent) {
	if m.typeChangePtr != 0 {
		RemoveEventElement(m.typeChangePtr)
	}
	m.typeChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3007, m.Instance(), m.typeChangePtr)
}
