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

// ICustomPrinterSetupDialog Parent: ICommonDialog
type ICustomPrinterSetupDialog interface {
	ICommonDialog
}

// TCustomPrinterSetupDialog Parent: TCommonDialog
type TCustomPrinterSetupDialog struct {
	TCommonDialog
}

func NewCustomPrinterSetupDialog(TheOwner IComponent) ICustomPrinterSetupDialog {
	r1 := LCL().SysCallN(1948, GetObjectUintptr(TheOwner))
	return AsCustomPrinterSetupDialog(r1)
}

func CustomPrinterSetupDialogClass() TClass {
	ret := LCL().SysCallN(1947)
	return TClass(ret)
}
