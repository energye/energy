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

// IReplaceDialog Parent: IFindDialog
type IReplaceDialog interface {
	IFindDialog
	ReplaceText() string          // property
	SetReplaceText(AValue string) // property
	SetOnReplace(fn TNotifyEvent) // property event
}

// TReplaceDialog Parent: TFindDialog
type TReplaceDialog struct {
	TFindDialog
	replacePtr uintptr
}

func NewReplaceDialog(AOwner IComponent) IReplaceDialog {
	r1 := LCL().SysCallN(4161, GetObjectUintptr(AOwner))
	return AsReplaceDialog(r1)
}

func (m *TReplaceDialog) ReplaceText() string {
	r1 := LCL().SysCallN(4162, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TReplaceDialog) SetReplaceText(AValue string) {
	LCL().SysCallN(4162, 1, m.Instance(), PascalStr(AValue))
}

func ReplaceDialogClass() TClass {
	ret := LCL().SysCallN(4160)
	return TClass(ret)
}

func (m *TReplaceDialog) SetOnReplace(fn TNotifyEvent) {
	if m.replacePtr != 0 {
		RemoveEventElement(m.replacePtr)
	}
	m.replacePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4163, m.Instance(), m.replacePtr)
}
