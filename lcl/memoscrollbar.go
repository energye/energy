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

// IMemoScrollBar Parent: IControlScrollBar
type IMemoScrollBar interface {
	IControlScrollBar
}

// TMemoScrollBar Parent: TControlScrollBar
type TMemoScrollBar struct {
	TControlScrollBar
}

func NewMemoScrollBar(AControl IWinControl, AKind TScrollBarKind) IMemoScrollBar {
	r1 := LCL().SysCallN(4204, GetObjectUintptr(AControl), uintptr(AKind))
	return AsMemoScrollBar(r1)
}

func MemoScrollBarClass() TClass {
	ret := LCL().SysCallN(4203)
	return TClass(ret)
}
