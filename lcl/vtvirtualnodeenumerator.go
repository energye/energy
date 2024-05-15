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

// IVTVirtualNodeEnumerator Parent: IObject
type IVTVirtualNodeEnumerator interface {
	IObject
	Current() IVirtualNode // property
	MoveNext() bool        // function
}

// TVTVirtualNodeEnumerator Parent: TObject
type TVTVirtualNodeEnumerator struct {
	TObject
}

func NewVTVirtualNodeEnumerator() IVTVirtualNodeEnumerator {
	r1 := LCL().SysCallN(5891)
	return AsVTVirtualNodeEnumerator(r1)
}

func (m *TVTVirtualNodeEnumerator) Current() IVirtualNode {
	r1 := LCL().SysCallN(5892, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TVTVirtualNodeEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(5893, m.Instance())
	return GoBool(r1)
}

func VTVirtualNodeEnumeratorClass() TClass {
	ret := LCL().SysCallN(5890)
	return TClass(ret)
}
