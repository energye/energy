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

// ITaskDialogRadioButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogRadioButtonItem interface {
	ITaskDialogBaseButtonItem
}

// TTaskDialogRadioButtonItem Parent: TTaskDialogBaseButtonItem
type TTaskDialogRadioButtonItem struct {
	TTaskDialogBaseButtonItem
}

func NewTaskDialogRadioButtonItem(ACollection ICollection) ITaskDialogRadioButtonItem {
	r1 := LCL().SysCallN(4706, GetObjectUintptr(ACollection))
	return AsTaskDialogRadioButtonItem(r1)
}

func TaskDialogRadioButtonItemClass() TClass {
	ret := LCL().SysCallN(4705)
	return TClass(ret)
}
