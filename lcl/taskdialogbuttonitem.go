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

// ITaskDialogButtonItem Parent: ITaskDialogBaseButtonItem
type ITaskDialogButtonItem interface {
	ITaskDialogBaseButtonItem
}

// TTaskDialogButtonItem Parent: TTaskDialogBaseButtonItem
type TTaskDialogButtonItem struct {
	TTaskDialogBaseButtonItem
}

func NewTaskDialogButtonItem(ACollection ICollection) ITaskDialogButtonItem {
	r1 := LCL().SysCallN(5350, GetObjectUintptr(ACollection))
	return AsTaskDialogButtonItem(r1)
}

func TaskDialogButtonItemClass() TClass {
	ret := LCL().SysCallN(5349)
	return TClass(ret)
}
