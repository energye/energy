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

// IAVLTree Parent: IObject
type IAVLTree interface {
	IObject
	Root() IAVLTreeNode                                                                              // property
	Count() SizeInt                                                                                  // property
	NewNode() IAVLTreeNode                                                                           // function
	Add(Data uintptr) IAVLTreeNode                                                                   // function
	AddAscendingSequence(Data uintptr, LastAdded IAVLTreeNode, Successor *IAVLTreeNode) IAVLTreeNode // function
	Remove(Data uintptr) bool                                                                        // function
	RemovePointer(Data uintptr) bool                                                                 // function
	IsEqual(aTree IAVLTree, CheckDataPointer bool) bool                                              // function
	Compare(Data1, Data2 uintptr) int32                                                              // function
	Find(Data uintptr) IAVLTreeNode                                                                  // function
	FindSuccessor(ANode IAVLTreeNode) IAVLTreeNode                                                   // function
	FindPrecessor(ANode IAVLTreeNode) IAVLTreeNode                                                   // function
	FindLowest() IAVLTreeNode                                                                        // function
	FindHighest() IAVLTreeNode                                                                       // function
	FindNearest(Data uintptr) IAVLTreeNode                                                           // function
	FindPointer(Data uintptr) IAVLTreeNode                                                           // function
	FindLeftMost(Data uintptr) IAVLTreeNode                                                          // function
	FindRightMost(Data uintptr) IAVLTreeNode                                                         // function
	FindLeftMostSameKey(ANode IAVLTreeNode) IAVLTreeNode                                             // function
	FindRightMostSameKey(ANode IAVLTreeNode) IAVLTreeNode                                            // function
	GetEnumerator() IAVLTreeNodeEnumerator                                                           // function
	GetEnumeratorHighToLow() IAVLTreeNodeEnumerator                                                  // function
	NodeToReportStr(aNode IAVLTreeNode) string                                                       // function
	ReportAsString() string                                                                          // function
	DisposeNode(ANode IAVLTreeNode)                                                                  // procedure
	Add1(ANode IAVLTreeNode)                                                                         // procedure
	Delete(ANode IAVLTreeNode)                                                                       // procedure
	MoveDataLeftMost(ANode *IAVLTreeNode)                                                            // procedure
	MoveDataRightMost(ANode *IAVLTreeNode)                                                           // procedure
	Clear()                                                                                          // procedure
	FreeAndClear()                                                                                   // procedure
	FreeAndDelete(ANode IAVLTreeNode)                                                                // procedure
	Assign(aTree IAVLTree)                                                                           // procedure
	ConsistencyCheck()                                                                               // procedure
	WriteReportToStream(s IStream)                                                                   // procedure
}

// TAVLTree Parent: TObject
type TAVLTree struct {
	TObject
}

func NewAVLTree() IAVLTree {
	r1 := LCL().SysCallN(48)
	return AsAVLTree(r1)
}

func (m *TAVLTree) Root() IAVLTreeNode {
	r1 := LCL().SysCallN(74, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) Count() SizeInt {
	r1 := LCL().SysCallN(47, m.Instance())
	return SizeInt(r1)
}

func (m *TAVLTree) NewNode() IAVLTreeNode {
	r1 := LCL().SysCallN(69, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) Add(Data uintptr) IAVLTreeNode {
	r1 := LCL().SysCallN(39, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) AddAscendingSequence(Data uintptr, LastAdded IAVLTreeNode, Successor *IAVLTreeNode) IAVLTreeNode {
	var result2 uintptr
	r1 := LCL().SysCallN(41, m.Instance(), uintptr(Data), GetObjectUintptr(LastAdded), uintptr(unsafePointer(&result2)))
	*Successor = AsAVLTreeNode(result2)
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) Remove(Data uintptr) bool {
	r1 := LCL().SysCallN(71, m.Instance(), uintptr(Data))
	return GoBool(r1)
}

func (m *TAVLTree) RemovePointer(Data uintptr) bool {
	r1 := LCL().SysCallN(72, m.Instance(), uintptr(Data))
	return GoBool(r1)
}

func (m *TAVLTree) IsEqual(aTree IAVLTree, CheckDataPointer bool) bool {
	r1 := LCL().SysCallN(66, m.Instance(), GetObjectUintptr(aTree), PascalBool(CheckDataPointer))
	return GoBool(r1)
}

func (m *TAVLTree) Compare(Data1, Data2 uintptr) int32 {
	r1 := LCL().SysCallN(45, m.Instance(), uintptr(Data1), uintptr(Data2))
	return int32(r1)
}

func (m *TAVLTree) Find(Data uintptr) IAVLTreeNode {
	r1 := LCL().SysCallN(51, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindSuccessor(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := LCL().SysCallN(61, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindPrecessor(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := LCL().SysCallN(58, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindLowest() IAVLTreeNode {
	r1 := LCL().SysCallN(55, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindHighest() IAVLTreeNode {
	r1 := LCL().SysCallN(52, m.Instance())
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindNearest(Data uintptr) IAVLTreeNode {
	r1 := LCL().SysCallN(56, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindPointer(Data uintptr) IAVLTreeNode {
	r1 := LCL().SysCallN(57, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindLeftMost(Data uintptr) IAVLTreeNode {
	r1 := LCL().SysCallN(53, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindRightMost(Data uintptr) IAVLTreeNode {
	r1 := LCL().SysCallN(59, m.Instance(), uintptr(Data))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindLeftMostSameKey(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := LCL().SysCallN(54, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) FindRightMostSameKey(ANode IAVLTreeNode) IAVLTreeNode {
	r1 := LCL().SysCallN(60, m.Instance(), GetObjectUintptr(ANode))
	return AsAVLTreeNode(r1)
}

func (m *TAVLTree) GetEnumerator() IAVLTreeNodeEnumerator {
	r1 := LCL().SysCallN(64, m.Instance())
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTree) GetEnumeratorHighToLow() IAVLTreeNodeEnumerator {
	r1 := LCL().SysCallN(65, m.Instance())
	return AsAVLTreeNodeEnumerator(r1)
}

func (m *TAVLTree) NodeToReportStr(aNode IAVLTreeNode) string {
	r1 := LCL().SysCallN(70, m.Instance(), GetObjectUintptr(aNode))
	return GoStr(r1)
}

func (m *TAVLTree) ReportAsString() string {
	r1 := LCL().SysCallN(73, m.Instance())
	return GoStr(r1)
}

func AVLTreeClass() TClass {
	ret := LCL().SysCallN(43)
	return TClass(ret)
}

func (m *TAVLTree) DisposeNode(ANode IAVLTreeNode) {
	LCL().SysCallN(50, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) Add1(ANode IAVLTreeNode) {
	LCL().SysCallN(40, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) Delete(ANode IAVLTreeNode) {
	LCL().SysCallN(49, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) MoveDataLeftMost(ANode *IAVLTreeNode) {
	var result0 uintptr
	LCL().SysCallN(67, m.Instance(), uintptr(unsafePointer(&result0)))
	*ANode = AsAVLTreeNode(result0)
}

func (m *TAVLTree) MoveDataRightMost(ANode *IAVLTreeNode) {
	var result0 uintptr
	LCL().SysCallN(68, m.Instance(), uintptr(unsafePointer(&result0)))
	*ANode = AsAVLTreeNode(result0)
}

func (m *TAVLTree) Clear() {
	LCL().SysCallN(44, m.Instance())
}

func (m *TAVLTree) FreeAndClear() {
	LCL().SysCallN(62, m.Instance())
}

func (m *TAVLTree) FreeAndDelete(ANode IAVLTreeNode) {
	LCL().SysCallN(63, m.Instance(), GetObjectUintptr(ANode))
}

func (m *TAVLTree) Assign(aTree IAVLTree) {
	LCL().SysCallN(42, m.Instance(), GetObjectUintptr(aTree))
}

func (m *TAVLTree) ConsistencyCheck() {
	LCL().SysCallN(46, m.Instance())
}

func (m *TAVLTree) WriteReportToStream(s IStream) {
	LCL().SysCallN(75, m.Instance(), GetObjectUintptr(s))
}
