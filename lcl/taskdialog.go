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

// ITaskDialog Parent: ICustomTaskDialog
type ITaskDialog interface {
	ICustomTaskDialog
}

// TTaskDialog Parent: TCustomTaskDialog
type TTaskDialog struct {
	TCustomTaskDialog
}

func NewTaskDialog(AOwner IComponent) ITaskDialog {
	r1 := LCL().SysCallN(5366, GetObjectUintptr(AOwner))
	return AsTaskDialog(r1)
}

func TaskDialogClass() TClass {
	ret := LCL().SysCallN(5365)
	return TClass(ret)
}
