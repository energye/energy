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

// ITaskDialogBaseButtonItem Parent: ICollectionItem
type ITaskDialogBaseButtonItem interface {
	ICollectionItem
	ModalResult() TModalResult          // property
	SetModalResult(AValue TModalResult) // property
	Caption() string                    // property
	SetCaption(AValue string)           // property
	Default() bool                      // property
	SetDefault(AValue bool)             // property
}

// TTaskDialogBaseButtonItem Parent: TCollectionItem
type TTaskDialogBaseButtonItem struct {
	TCollectionItem
}

func NewTaskDialogBaseButtonItem(ACollection ICollection) ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(4688, GetObjectUintptr(ACollection))
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogBaseButtonItem) ModalResult() TModalResult {
	r1 := LCL().SysCallN(4690, 0, m.Instance(), 0)
	return TModalResult(r1)
}

func (m *TTaskDialogBaseButtonItem) SetModalResult(AValue TModalResult) {
	LCL().SysCallN(4690, 1, m.Instance(), uintptr(AValue))
}

func (m *TTaskDialogBaseButtonItem) Caption() string {
	r1 := LCL().SysCallN(4686, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTaskDialogBaseButtonItem) SetCaption(AValue string) {
	LCL().SysCallN(4686, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTaskDialogBaseButtonItem) Default() bool {
	r1 := LCL().SysCallN(4689, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTaskDialogBaseButtonItem) SetDefault(AValue bool) {
	LCL().SysCallN(4689, 1, m.Instance(), PascalBool(AValue))
}

func TaskDialogBaseButtonItemClass() TClass {
	ret := LCL().SysCallN(4687)
	return TClass(ret)
}
