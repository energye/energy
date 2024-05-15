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

// ILazAccessibleObjectEnumerator Parent: IAVLTreeNodeEnumerator
type ILazAccessibleObjectEnumerator interface {
	IAVLTreeNodeEnumerator
	CurrentForLazAccessibleObject() ILazAccessibleObject // property
}

// TLazAccessibleObjectEnumerator Parent: TAVLTreeNodeEnumerator
type TLazAccessibleObjectEnumerator struct {
	TAVLTreeNodeEnumerator
}

func NewLazAccessibleObjectEnumerator(Tree IAVLTree, aLowToHigh bool) ILazAccessibleObjectEnumerator {
	r1 := LCL().SysCallN(3485, GetObjectUintptr(Tree), PascalBool(aLowToHigh))
	return AsLazAccessibleObjectEnumerator(r1)
}

func (m *TLazAccessibleObjectEnumerator) CurrentForLazAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(3486, m.Instance())
	return AsLazAccessibleObject(r1)
}

func LazAccessibleObjectEnumeratorClass() TClass {
	ret := LCL().SysCallN(3484)
	return TClass(ret)
}
