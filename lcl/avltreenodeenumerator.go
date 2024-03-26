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

// IAVLTreeNodeEnumerator Parent: IObject
type IAVLTreeNodeEnumerator interface {
	IObject
	Current() IAVLTreeNode                 // property
	LowToHigh() bool                       // property
	GetEnumerator() IAVLTreeNodeEnumerator // function
	MoveNext() bool                        // function
}

// TAVLTreeNodeEnumerator Parent: TObject
type TAVLTreeNodeEnumerator struct {
	TObject
}

func NewAVLTreeNodeEnumerator(Tree IAVLTree, aLowToHigh bool) IAVLTreeNodeEnumerator {
	r1 := LCL().SysCallN(26, GetObjectUintptr(Tree), PascalBool(aLowToHigh))
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTreeNodeEnumerator) Current() IAVLTreeNode {
	r1 := LCL().SysCallN(27, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNodeEnumerator) LowToHigh() bool {
	r1 := LCL().SysCallN(29, m.Instance())
	return GoBool(r1)
}

func (m *TAVLTreeNodeEnumerator) GetEnumerator() IAVLTreeNodeEnumerator {
	r1 := LCL().SysCallN(28, m.Instance())
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTreeNodeEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(30, m.Instance())
	return GoBool(r1)
}

func AVLTreeNodeEnumeratorClass() TClass {
	ret := LCL().SysCallN(25)
	return TClass(ret)
}
