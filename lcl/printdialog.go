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

// IPrintDialog Parent: ICustomPrintDialog
type IPrintDialog interface {
	ICustomPrintDialog
	AttachTo() ICustomForm                   // property
	SetAttachTo(AValue ICustomForm)          // property
	SetOnDialogResult(fn TDialogResultEvent) // property event
}

// TPrintDialog Parent: TCustomPrintDialog
type TPrintDialog struct {
	TCustomPrintDialog
	dialogResultPtr uintptr
}

func NewPrintDialog(TheOwner IComponent) IPrintDialog {
	r1 := LCL().SysCallN(4580, GetObjectUintptr(TheOwner))
	return AsPrintDialog(r1)
}

func (m *TPrintDialog) AttachTo() ICustomForm {
	r1 := LCL().SysCallN(4578, 0, m.Instance(), 0)
	return AsCustomForm(r1)
}

func (m *TPrintDialog) SetAttachTo(AValue ICustomForm) {
	LCL().SysCallN(4578, 1, m.Instance(), GetObjectUintptr(AValue))
}

func PrintDialogClass() TClass {
	ret := LCL().SysCallN(4579)
	return TClass(ret)
}

func (m *TPrintDialog) SetOnDialogResult(fn TDialogResultEvent) {
	if m.dialogResultPtr != 0 {
		RemoveEventElement(m.dialogResultPtr)
	}
	m.dialogResultPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4581, m.Instance(), m.dialogResultPtr)
}
