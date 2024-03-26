//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/energy/v2/api"
)

// IInterfacedObject Parent: IObject
type IInterfacedObject interface {
	IObject
	RefCount() int32 // property
}

// TInterfacedObject Parent: TObject
type TInterfacedObject struct {
	TObject
}

func NewInterfacedObject() IInterfacedObject {
	r1 := api.LCLPreDef().SysCallN(api.InterfacedObjectCreate())
	return AsInterfacedObject(r1)
}

func (m *TInterfacedObject) RefCount() int32 {
	r1 := api.LCLPreDef().SysCallN(api.InterfacedObjectRefCount(), m.Instance())
	return int32(r1)
}
