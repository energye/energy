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

// ICustomFrame Parent: ICustomDesignControl
type ICustomFrame interface {
	ICustomDesignControl
	ParentBackground() bool          // property
	SetParentBackground(AValue bool) // property
}

// TCustomFrame Parent: TCustomDesignControl
type TCustomFrame struct {
	TCustomDesignControl
}

func NewCustomFrame(AOwner IComponent) ICustomFrame {
	r1 := LCL().SysCallN(1545, GetObjectUintptr(AOwner))
	return AsCustomFrame(r1)
}

func (m *TCustomFrame) ParentBackground() bool {
	r1 := LCL().SysCallN(1546, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomFrame) SetParentBackground(AValue bool) {
	LCL().SysCallN(1546, 1, m.Instance(), PascalBool(AValue))
}

func CustomFrameClass() TClass {
	ret := LCL().SysCallN(1544)
	return TClass(ret)
}
