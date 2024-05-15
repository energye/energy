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

// ITaskDialogButtons Parent: IOwnedCollection
type ITaskDialogButtons interface {
	IOwnedCollection
	DefaultButton() ITaskDialogBaseButtonItem                                          // property
	SetDefaultButton(AValue ITaskDialogBaseButtonItem)                                 // property
	ItemsForTaskDialogBaseButtonItem(Index int32) ITaskDialogBaseButtonItem            // property
	SetItemsForTaskDialogBaseButtonItem(Index int32, AValue ITaskDialogBaseButtonItem) // property
	AddForTaskDialogBaseButtonItem() ITaskDialogBaseButtonItem                         // function
	FindButton(AModalResult TModalResult) ITaskDialogBaseButtonItem                    // function
	GetEnumeratorForTaskDialogButtonsEnumerator() ITaskDialogButtonsEnumerator         // function
}

// TTaskDialogButtons Parent: TOwnedCollection
type TTaskDialogButtons struct {
	TOwnedCollection
}

func NewTaskDialogButtons(AOwner IPersistent, AItemClass TCollectionItemClass) ITaskDialogButtons {
	r1 := LCL().SysCallN(5358, GetObjectUintptr(AOwner), uintptr(AItemClass))
	return AsTaskDialogButtons(r1)
}

func (m *TTaskDialogButtons) DefaultButton() ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(5359, 0, m.Instance(), 0)
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) SetDefaultButton(AValue ITaskDialogBaseButtonItem) {
	LCL().SysCallN(5359, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TTaskDialogButtons) ItemsForTaskDialogBaseButtonItem(Index int32) ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(5362, 0, m.Instance(), uintptr(Index))
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) SetItemsForTaskDialogBaseButtonItem(Index int32, AValue ITaskDialogBaseButtonItem) {
	LCL().SysCallN(5362, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TTaskDialogButtons) AddForTaskDialogBaseButtonItem() ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(5356, m.Instance())
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) FindButton(AModalResult TModalResult) ITaskDialogBaseButtonItem {
	r1 := LCL().SysCallN(5360, m.Instance(), uintptr(AModalResult))
	return AsTaskDialogBaseButtonItem(r1)
}

func (m *TTaskDialogButtons) GetEnumeratorForTaskDialogButtonsEnumerator() ITaskDialogButtonsEnumerator {
	r1 := LCL().SysCallN(5361, m.Instance())
	return AsTaskDialogButtonsEnumerator(r1)
}

func TaskDialogButtonsClass() TClass {
	ret := LCL().SysCallN(5357)
	return TClass(ret)
}
