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

// ILCLComponent Parent: IComponent
type ILCLComponent interface {
	IComponent
	LCLRefCount() int32                         // property
	RemoveAllHandlersOfObject(AnObject IObject) // procedure
	IncLCLRefCount()                            // procedure
	DecLCLRefCount()                            // procedure
}

// TLCLComponent Parent: TComponent
type TLCLComponent struct {
	TComponent
}

func NewLCLComponent(TheOwner IComponent) ILCLComponent {
	r1 := LCL().SysCallN(3179, GetObjectUintptr(TheOwner))
	return AsLCLComponent(r1)
}

func (m *TLCLComponent) LCLRefCount() int32 {
	r1 := LCL().SysCallN(3182, m.Instance())
	return int32(r1)
}

func LCLComponentClass() TClass {
	ret := LCL().SysCallN(3178)
	return TClass(ret)
}

func (m *TLCLComponent) RemoveAllHandlersOfObject(AnObject IObject) {
	LCL().SysCallN(3183, m.Instance(), GetObjectUintptr(AnObject))
}

func (m *TLCLComponent) IncLCLRefCount() {
	LCL().SysCallN(3181, m.Instance())
}

func (m *TLCLComponent) DecLCLRefCount() {
	LCL().SysCallN(3180, m.Instance())
}
