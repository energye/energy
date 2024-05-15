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

// ILCLReferenceComponent Is Abstract Class Parent: ILCLComponent
// A base class for all components having a handle
type ILCLReferenceComponent interface {
	ILCLComponent
	HandleAllocated() bool    // property
	ReferenceAllocated() bool // property
}

// TLCLReferenceComponent Is Abstract Class Parent: TLCLComponent
// A base class for all components having a handle
type TLCLReferenceComponent struct {
	TLCLComponent
}

func (m *TLCLReferenceComponent) HandleAllocated() bool {
	r1 := LCL().SysCallN(3428, m.Instance())
	return GoBool(r1)
}

func (m *TLCLReferenceComponent) ReferenceAllocated() bool {
	r1 := LCL().SysCallN(3429, m.Instance())
	return GoBool(r1)
}

func LCLReferenceComponentClass() TClass {
	ret := LCL().SysCallN(3427)
	return TClass(ret)
}
