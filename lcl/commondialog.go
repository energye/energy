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

// ICommonDialog Parent: ILCLComponent
type ICommonDialog interface {
	ILCLComponent
	Handle() THandle                    // property
	SetHandle(AValue THandle)           // property
	UserChoice() int32                  // property
	SetUserChoice(AValue int32)         // property
	Width() int32                       // property
	SetWidth(AValue int32)              // property
	Height() int32                      // property
	SetHeight(AValue int32)             // property
	HelpContext() THelpContext          // property
	SetHelpContext(AValue THelpContext) // property
	Title() string                      // property
	SetTitle(AValue string)             // property
	Execute() bool                      // function
	HandleAllocated() bool              // function
	Close()                             // procedure
	SetOnClose(fn TNotifyEvent)         // property event
	SetOnCanClose(fn TCloseQueryEvent)  // property event
	SetOnShow(fn TNotifyEvent)          // property event
}

// TCommonDialog Parent: TLCLComponent
type TCommonDialog struct {
	TLCLComponent
	closePtr    uintptr
	canClosePtr uintptr
	showPtr     uintptr
}

func NewCommonDialog(TheOwner IComponent) ICommonDialog {
	r1 := LCL().SysCallN(862, GetObjectUintptr(TheOwner))
	return AsCommonDialog(r1)
}

func (m *TCommonDialog) Handle() THandle {
	r1 := LCL().SysCallN(864, 0, m.Instance(), 0)
	return THandle(r1)
}

func (m *TCommonDialog) SetHandle(AValue THandle) {
	LCL().SysCallN(864, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) UserChoice() int32 {
	r1 := LCL().SysCallN(872, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCommonDialog) SetUserChoice(AValue int32) {
	LCL().SysCallN(872, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) Width() int32 {
	r1 := LCL().SysCallN(873, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCommonDialog) SetWidth(AValue int32) {
	LCL().SysCallN(873, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) Height() int32 {
	r1 := LCL().SysCallN(866, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCommonDialog) SetHeight(AValue int32) {
	LCL().SysCallN(866, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) HelpContext() THelpContext {
	r1 := LCL().SysCallN(867, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TCommonDialog) SetHelpContext(AValue THelpContext) {
	LCL().SysCallN(867, 1, m.Instance(), uintptr(AValue))
}

func (m *TCommonDialog) Title() string {
	r1 := LCL().SysCallN(871, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCommonDialog) SetTitle(AValue string) {
	LCL().SysCallN(871, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCommonDialog) Execute() bool {
	r1 := LCL().SysCallN(863, m.Instance())
	return GoBool(r1)
}

func (m *TCommonDialog) HandleAllocated() bool {
	r1 := LCL().SysCallN(865, m.Instance())
	return GoBool(r1)
}

func CommonDialogClass() TClass {
	ret := LCL().SysCallN(860)
	return TClass(ret)
}

func (m *TCommonDialog) Close() {
	LCL().SysCallN(861, m.Instance())
}

func (m *TCommonDialog) SetOnClose(fn TNotifyEvent) {
	if m.closePtr != 0 {
		RemoveEventElement(m.closePtr)
	}
	m.closePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(869, m.Instance(), m.closePtr)
}

func (m *TCommonDialog) SetOnCanClose(fn TCloseQueryEvent) {
	if m.canClosePtr != 0 {
		RemoveEventElement(m.canClosePtr)
	}
	m.canClosePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(868, m.Instance(), m.canClosePtr)
}

func (m *TCommonDialog) SetOnShow(fn TNotifyEvent) {
	if m.showPtr != 0 {
		RemoveEventElement(m.showPtr)
	}
	m.showPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(870, m.Instance(), m.showPtr)
}
