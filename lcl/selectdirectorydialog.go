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

// ISelectDirectoryDialog Parent: IOpenDialog
type ISelectDirectoryDialog interface {
	IOpenDialog
}

// TSelectDirectoryDialog Parent: TOpenDialog
type TSelectDirectoryDialog struct {
	TOpenDialog
}

func NewSelectDirectoryDialog(AOwner IComponent) ISelectDirectoryDialog {
	r1 := LCL().SysCallN(4321, GetObjectUintptr(AOwner))
	return AsSelectDirectoryDialog(r1)
}

func SelectDirectoryDialogClass() TClass {
	ret := LCL().SysCallN(4320)
	return TClass(ret)
}
