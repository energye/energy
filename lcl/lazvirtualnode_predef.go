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
	"github.com/energye/energy/v2/types"
)

type IVirtualNode interface {
	IObject
	Index() types.Cardinal
	SetIndex(value types.Cardinal)
	ChildCount() types.Cardinal
	SetChildCount(value types.Cardinal)
	NodeHeight() types.Word
	SetNodeHeight(value types.Word)
	States() types.TVirtualNodeStates
	SetStates(value types.TVirtualNodeStates)
	Align() types.Byte
	SetAlign(value types.Byte)
	CheckState() types.TCheckState
	SetCheckState(value types.TCheckState)
	CheckType() types.TCheckType
	SetCheckType(value types.TCheckType)
	Dummy() types.Byte
	SetDummy(value types.Byte)
	TotalCount() types.Cardinal
	SetTotalCount(value types.Cardinal)
	TotalHeight() types.Cardinal
	SetTotalHeight(value types.Cardinal)
	Parent() IVirtualNode
	SetParent(value IVirtualNode)
	PrevSibling() IVirtualNode
	SetPrevSibling(value IVirtualNode)
	NextSibling() IVirtualNode
	SetNextSibling(value IVirtualNode)
	FirstChild() IVirtualNode
	SetFirstChild(value IVirtualNode)
	LastChild() IVirtualNode
	SetLastChild(value IVirtualNode)
	Data() uintptr
	SetData(value uintptr, totalInternalDataSize types.Cardinal)
}

// TVirtualNode = ^TVirtualNode
// type PVirtualNode uintptr
type TVirtualNode struct {
	TObject
}

// 应该不能自己创建
//func NewVirtualNode() IVirtualNode {
//	var result uintptr
//	LCLPreDef().SysCallN(VirtualNodeCreate(), uintptr(unsafePointer(&result)))
//	return AsVirtualNode(result)
//}

// Index index of node with regard to its parent
func (m *TVirtualNode) Index() types.Cardinal {
	r1 := LCLPreDef().SysCallN(VirtualNodeIndex(), 0, m.Instance(), 0)
	return types.Cardinal(r1)
}

func (m *TVirtualNode) SetIndex(value types.Cardinal) {
	LCLPreDef().SysCallN(VirtualNodeIndex(), 1, m.Instance(), uintptr(value))
}

// ChildCount number of child nodes
func (m *TVirtualNode) ChildCount() types.Cardinal {
	r1 := LCLPreDef().SysCallN(VirtualNodeChildCount(), 0, m.Instance(), 0)
	return types.Cardinal(r1)
}

func (m *TVirtualNode) SetChildCount(value types.Cardinal) {
	LCLPreDef().SysCallN(VirtualNodeChildCount(), 1, m.Instance(), uintptr(value))
}

// NodeHeight height in pixels
func (m *TVirtualNode) NodeHeight() types.Word {
	r1 := LCLPreDef().SysCallN(VirtualNodeNodeHeight(), 0, m.Instance(), 0)
	return types.Word(r1)
}

func (m *TVirtualNode) SetNodeHeight(value types.Word) {
	LCLPreDef().SysCallN(VirtualNodeNodeHeight(), 1, m.Instance(), uintptr(value))
}

// States states describing various properties of the node (expanded, initialized etc.)
func (m *TVirtualNode) States() types.TVirtualNodeStates {
	r1 := LCLPreDef().SysCallN(VirtualNodeStates(), 0, m.Instance(), 0)
	return types.TVirtualNodeStates(r1)
}

func (m *TVirtualNode) SetStates(value types.TVirtualNodeStates) {
	LCLPreDef().SysCallN(VirtualNodeStates(), 1, m.Instance(), uintptr(value))
}

// Align line/button alignment
func (m *TVirtualNode) Align() types.Byte {
	r1 := LCLPreDef().SysCallN(VirtualNodeAlign(), 0, m.Instance(), 0)
	return types.Byte(r1)
}

func (m *TVirtualNode) SetAlign(value types.Byte) {
	LCLPreDef().SysCallN(VirtualNodeAlign(), 1, m.Instance(), uintptr(value))
}

// CheckState indicates the current check state (e.g. checked, pressed etc.)
func (m *TVirtualNode) CheckState() types.TCheckState {
	r1 := LCLPreDef().SysCallN(VirtualNodeCheckState(), 0, m.Instance(), 0)
	return types.TCheckState(r1)
}

func (m *TVirtualNode) SetCheckState(value types.TCheckState) {
	LCLPreDef().SysCallN(VirtualNodeCheckState(), 1, m.Instance(), uintptr(value))
}

// CheckType indicates which check type shall be used for this node
func (m *TVirtualNode) CheckType() types.TCheckType {
	r1 := LCLPreDef().SysCallN(VirtualNodeCheckType(), 0, m.Instance(), 0)
	return types.TCheckState(r1)
}

func (m *TVirtualNode) SetCheckType(value types.TCheckType) {
	LCLPreDef().SysCallN(VirtualNodeCheckType(), 1, m.Instance(), uintptr(value))
}

// Dummy dummy value to fill DWORD boundary
func (m *TVirtualNode) Dummy() types.Byte {
	r1 := LCLPreDef().SysCallN(VirtualNodeDummy(), 0, m.Instance(), 0)
	return types.Byte(r1)
}

func (m *TVirtualNode) SetDummy(value types.Byte) {
	LCLPreDef().SysCallN(VirtualNodeDummy(), 1, m.Instance(), uintptr(value))
}

// TotalCount sum of this node, all of its child nodes and their child nodes etc.
func (m *TVirtualNode) TotalCount() types.Cardinal {
	r1 := LCLPreDef().SysCallN(VirtualNodeTotalCount(), 0, m.Instance(), 0)
	return types.Cardinal(r1)
}

func (m *TVirtualNode) SetTotalCount(value types.Cardinal) {
	LCLPreDef().SysCallN(VirtualNodeTotalCount(), 1, m.Instance(), uintptr(value))
}

// TotalHeight height in pixels this node covers on screen including the height of all of its children
func (m *TVirtualNode) TotalHeight() types.Cardinal {
	r1 := LCLPreDef().SysCallN(VirtualNodeTotalHeight(), 0, m.Instance(), 0)
	return types.Cardinal(r1)
}

func (m *TVirtualNode) SetTotalHeight(value types.Cardinal) {
	LCLPreDef().SysCallN(VirtualNodeTotalHeight(), 1, m.Instance(), uintptr(value))
}

//	Note: Some copy routines require that all pointers (as well as the data area) in a node are
//	      located at the end of the node! Hence if you want to add new member fields (except pointers to internal
//	      data) then put them before field Parent.

// Parent reference to the node's parent (for the root this contains the treeview)
func (m *TVirtualNode) Parent() IVirtualNode {
	var result uintptr
	LCLPreDef().SysCallN(VirtualNodeParent(), 0, m.Instance(), 0, uintptr(unsafePointer(&result)))
	return AsVirtualNode(result)
}

func (m *TVirtualNode) SetParent(value IVirtualNode) {
	LCLPreDef().SysCallN(VirtualNodeParent(), 1, m.Instance(), value.Instance(), 0)
}

// PrevSibling link to the node's previous sibling or nil if it is the first node
func (m *TVirtualNode) PrevSibling() IVirtualNode {
	var result uintptr
	LCLPreDef().SysCallN(VirtualNodePrevSibling(), 0, m.Instance(), 0, uintptr(unsafePointer(&result)))
	return AsVirtualNode(result)
}

func (m *TVirtualNode) SetPrevSibling(value IVirtualNode) {
	LCLPreDef().SysCallN(VirtualNodePrevSibling(), 1, m.Instance(), value.Instance(), 0)
}

// NextSibling link to the node's next sibling or nil if it is the last node
func (m *TVirtualNode) NextSibling() IVirtualNode {
	var result uintptr
	LCLPreDef().SysCallN(VirtualNodeNextSibling(), 0, m.Instance(), 0, uintptr(unsafePointer(&result)))
	return AsVirtualNode(result)
}

func (m *TVirtualNode) SetNextSibling(value IVirtualNode) {
	LCLPreDef().SysCallN(VirtualNodeNextSibling(), 1, m.Instance(), value.Instance(), 0)
}

// FirstChild link to the node's first child...
func (m *TVirtualNode) FirstChild() IVirtualNode {
	var result uintptr
	LCLPreDef().SysCallN(VirtualNodeFirstChild(), 0, m.Instance(), 0, uintptr(unsafePointer(&result)))
	return AsVirtualNode(result)
}

func (m *TVirtualNode) SetFirstChild(value IVirtualNode) {
	LCLPreDef().SysCallN(VirtualNodeFirstChild(), 1, m.Instance(), value.Instance(), 0)
}

// LastChild link to the node's last child...
func (m *TVirtualNode) LastChild() IVirtualNode {
	var result uintptr
	LCLPreDef().SysCallN(VirtualNodeLastChild(), 0, m.Instance(), 0, uintptr(unsafePointer(&result)))
	return AsVirtualNode(result)
}

func (m *TVirtualNode) SetLastChild(value IVirtualNode) {
	LCLPreDef().SysCallN(VirtualNodeLastChild(), 1, m.Instance(), value.Instance(), 0)
}

// Data record end;  this is a placeholder, each node gets extra data determined by NodeDataSize
func (m *TVirtualNode) Data() uintptr {
	var result uintptr
	LCLPreDef().SysCallN(VirtualNodeData(), 0, m.Instance(), 0, 0, uintptr(unsafePointer(&result)))
	return result
}

// SetData totalInternalDataSize default: 0
func (m *TVirtualNode) SetData(value uintptr, totalInternalDataSize types.Cardinal) {
	LCLPreDef().SysCallN(VirtualNodeData(), 1, m.Instance(), value, uintptr(totalInternalDataSize), 0)
}
