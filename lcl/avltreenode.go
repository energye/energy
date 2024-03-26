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

// IAVLTreeNode Parent: IObject
type IAVLTreeNode interface {
	IObject
	Successor() IAVLTreeNode        // function
	Precessor() IAVLTreeNode        // function
	TreeDepth() int32               // function
	GetCount() SizeInt              // function
	Clear()                         // procedure
	ConsistencyCheck(Tree IAVLTree) // procedure
}

// TAVLTreeNode Parent: TObject
type TAVLTreeNode struct {
	TObject
}

func NewAVLTreeNode() IAVLTreeNode {
	r1 := LCL().SysCallN(34)
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNode) Successor() IAVLTreeNode {
	r1 := LCL().SysCallN(37, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNode) Precessor() IAVLTreeNode {
	r1 := LCL().SysCallN(36, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTreeNode) TreeDepth() int32 {
	r1 := LCL().SysCallN(38, m.Instance())
	return int32(r1)
}

func (m *TAVLTreeNode) GetCount() SizeInt {
	r1 := LCL().SysCallN(35, m.Instance())
	return SizeInt(r1)
}

func AVLTreeNodeClass() TClass {
	ret := LCL().SysCallN(31)
	return TClass(ret)
}

func (m *TAVLTreeNode) Clear() {
	LCL().SysCallN(32, m.Instance())
}

func (m *TAVLTreeNode) ConsistencyCheck(Tree IAVLTree) {
	LCL().SysCallN(33, m.Instance(), GetObjectUintptr(Tree))
}
