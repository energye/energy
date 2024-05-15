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

// IBaseVirtualTree Parent: ICustomControl
// ----- TBaseVirtualTree
type IBaseVirtualTree interface {
	ICustomControl
	BottomNode() IVirtualNode                                                                                                             // property
	SetBottomNode(AValue IVirtualNode)                                                                                                    // property
	CheckedCount() int32                                                                                                                  // property
	CheckImages() ICustomImageList                                                                                                        // property
	CheckState(Node IVirtualNode) TCheckState                                                                                             // property
	SetCheckState(Node IVirtualNode, AValue TCheckState)                                                                                  // property
	CheckType(Node IVirtualNode) TCheckType                                                                                               // property
	SetCheckType(Node IVirtualNode, AValue TCheckType)                                                                                    // property
	ChildCount(Node IVirtualNode) uint32                                                                                                  // property
	SetChildCount(Node IVirtualNode, AValue uint32)                                                                                       // property
	ChildrenInitialized(Node IVirtualNode) bool                                                                                           // property
	CutCopyCount() int32                                                                                                                  // property
	DragImage() IVTDragImage                                                                                                              // property
	DropTargetNode() IVirtualNode                                                                                                         // property
	SetDropTargetNode(AValue IVirtualNode)                                                                                                // property
	EmptyListMessage() string                                                                                                             // property
	SetEmptyListMessage(AValue string)                                                                                                    // property
	Expanded(Node IVirtualNode) bool                                                                                                      // property
	SetExpanded(Node IVirtualNode, AValue bool)                                                                                           // property
	FocusedColumn() TColumnIndex                                                                                                          // property
	SetFocusedColumn(AValue TColumnIndex)                                                                                                 // property
	FocusedNode() IVirtualNode                                                                                                            // property
	SetFocusedNode(AValue IVirtualNode)                                                                                                   // property
	FullyVisible(Node IVirtualNode) bool                                                                                                  // property
	SetFullyVisible(Node IVirtualNode, AValue bool)                                                                                       // property
	HasChildren(Node IVirtualNode) bool                                                                                                   // property
	SetHasChildren(Node IVirtualNode, AValue bool)                                                                                        // property
	HotNode() IVirtualNode                                                                                                                // property
	IsDisabled(Node IVirtualNode) bool                                                                                                    // property
	SetIsDisabled(Node IVirtualNode, AValue bool)                                                                                         // property
	IsEffectivelyFiltered(Node IVirtualNode) bool                                                                                         // property
	IsEffectivelyVisible(Node IVirtualNode) bool                                                                                          // property
	IsFiltered(Node IVirtualNode) bool                                                                                                    // property
	SetIsFiltered(Node IVirtualNode, AValue bool)                                                                                         // property
	NodeIsVisible(Node IVirtualNode) bool                                                                                                 // property
	SetNodeIsVisible(Node IVirtualNode, AValue bool)                                                                                      // property
	MultiLine(Node IVirtualNode) bool                                                                                                     // property
	SetMultiLine(Node IVirtualNode, AValue bool)                                                                                          // property
	NodeHeight(Node IVirtualNode) uint32                                                                                                  // property
	SetNodeHeight(Node IVirtualNode, AValue uint32)                                                                                       // property
	NodeParent(Node IVirtualNode) IVirtualNode                                                                                            // property
	SetNodeParent(Node IVirtualNode, AValue IVirtualNode)                                                                                 // property
	OffsetX() int32                                                                                                                       // property
	SetOffsetX(AValue int32)                                                                                                              // property
	OffsetXY() (resultPoint TPoint)                                                                                                       // property
	SetOffsetXY(AValue *TPoint)                                                                                                           // property
	OffsetY() int32                                                                                                                       // property
	SetOffsetY(AValue int32)                                                                                                              // property
	OperationCount() uint32                                                                                                               // property
	RootNode() IVirtualNode                                                                                                               // property
	SearchBuffer() string                                                                                                                 // property
	Selected(Node IVirtualNode) bool                                                                                                      // property
	SetSelected(Node IVirtualNode, AValue bool)                                                                                           // property
	SelectionLocked() bool                                                                                                                // property
	SetSelectionLocked(AValue bool)                                                                                                       // property
	TotalCount() uint32                                                                                                                   // property
	TreeStates() TVirtualTreeStates                                                                                                       // property
	SetTreeStates(AValue TVirtualTreeStates)                                                                                              // property
	SelectedCount() int32                                                                                                                 // property
	TopNode() IVirtualNode                                                                                                                // property
	SetTopNode(AValue IVirtualNode)                                                                                                       // property
	VerticalAlignment(Node IVirtualNode) Byte                                                                                             // property
	SetVerticalAlignment(Node IVirtualNode, AValue Byte)                                                                                  // property
	VisibleCount() uint32                                                                                                                 // property
	VisiblePath(Node IVirtualNode) bool                                                                                                   // property
	SetVisiblePath(Node IVirtualNode, AValue bool)                                                                                        // property
	UpdateCount() uint32                                                                                                                  // property
	AbsoluteIndex(Node IVirtualNode) uint32                                                                                               // function
	AddChild(Parent IVirtualNode, UserData uintptr) IVirtualNode                                                                          // function
	CancelEditNode() bool                                                                                                                 // function
	CanEdit(Node IVirtualNode, Column TColumnIndex) bool                                                                                  // function
	CopyTo(Source IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode                            // function
	CopyTo1(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode                                          // function
	EditNode(Node IVirtualNode, Column TColumnIndex) bool                                                                                 // function
	EndEditNode() bool                                                                                                                    // function
	GetDisplayRect(Node IVirtualNode, Column TColumnIndex, TextOnly bool, Unclipped bool, ApplyCellContentMargin bool) (resultRect TRect) // function
	GetEffectivelyFiltered(Node IVirtualNode) bool                                                                                        // function
	GetEffectivelyVisible(Node IVirtualNode) bool                                                                                         // function
	GetFirst(ConsiderChildrenAbove bool) IVirtualNode                                                                                     // function
	GetFirstChecked(State TCheckState, ConsiderChildrenAbove bool) IVirtualNode                                                           // function
	GetFirstChild(Node IVirtualNode) IVirtualNode                                                                                         // function
	GetFirstChildNoInit(Node IVirtualNode) IVirtualNode                                                                                   // function
	GetFirstCutCopy(ConsiderChildrenAbove bool) IVirtualNode                                                                              // function
	GetFirstInitialized(ConsiderChildrenAbove bool) IVirtualNode                                                                          // function
	GetFirstLeaf() IVirtualNode                                                                                                           // function
	GetFirstLevel(NodeLevel uint32) IVirtualNode                                                                                          // function
	GetFirstNoInit(ConsiderChildrenAbove bool) IVirtualNode                                                                               // function
	GetFirstSelected(ConsiderChildrenAbove bool) IVirtualNode                                                                             // function
	GetFirstVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                                     // function
	GetFirstVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                            // function
	GetFirstVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                      // function
	GetFirstVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                               // function
	GetLast(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                                   // function
	GetLastInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetLastNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                             // function
	GetLastChild(Node IVirtualNode) IVirtualNode                                                                                          // function
	GetLastChildNoInit(Node IVirtualNode) IVirtualNode                                                                                    // function
	GetLastVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                                      // function
	GetLastVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                             // function
	GetLastVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                       // function
	GetLastVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode                                // function
	GetMaxColumnWidth(Column TColumnIndex, UseSmartColumnWidth bool) int32                                                                // function
	GetNext(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                                   // function
	GetNextChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode                                         // function
	GetNextChecked1(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                           // function
	GetNextCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                            // function
	GetNextInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetNextLeaf(Node IVirtualNode) IVirtualNode                                                                                           // function
	GetNextLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode                                                                        // function
	GetNextNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                             // function
	GetNextSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                           // function
	GetNextSibling(Node IVirtualNode) IVirtualNode                                                                                        // function
	GetNextSiblingNoInit(Node IVirtualNode) IVirtualNode                                                                                  // function
	GetNextVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                            // function
	GetNextVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                      // function
	GetNextVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                           // function
	GetNextVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                     // function
	GetNodeAt(P *TPoint) IVirtualNode                                                                                                     // function
	GetNodeAt1(X, Y int32) IVirtualNode                                                                                                   // function
	GetNodeAt2(X, Y int32, Relative bool, NodeTop *int32) IVirtualNode                                                                    // function
	GetNodeData(Node IVirtualNode) uintptr                                                                                                // function
	GetNodeLevel(Node IVirtualNode) uint32                                                                                                // function
	GetPrevious(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                               // function
	GetPreviousChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode                                     // function
	GetPreviousCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetPreviousInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                    // function
	GetPreviousLeaf(Node IVirtualNode) IVirtualNode                                                                                       // function
	GetPreviousLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode                                                                    // function
	GetPreviousNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                         // function
	GetPreviousSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                       // function
	GetPreviousSibling(Node IVirtualNode) IVirtualNode                                                                                    // function
	GetPreviousSiblingNoInit(Node IVirtualNode) IVirtualNode                                                                              // function
	GetPreviousVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                        // function
	GetPreviousVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode                                                  // function
	GetPreviousVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                       // function
	GetPreviousVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                 // function
	GetSortedCutCopySet(Resolve bool) TNodeArray                                                                                          // function
	GetSortedSelection(Resolve bool) TNodeArray                                                                                           // function
	GetTreeRect() (resultRect TRect)                                                                                                      // function
	GetVisibleParent(Node IVirtualNode, IncludeFiltered bool) IVirtualNode                                                                // function
	HasAsParent(Node, PotentialParent IVirtualNode) bool                                                                                  // function
	InsertNode(Node IVirtualNode, Mode TVTNodeAttachMode, UserData uintptr) IVirtualNode                                                  // function
	InvalidateNode(Node IVirtualNode) (resultRect TRect)                                                                                  // function
	IsEditing() bool                                                                                                                      // function
	IsMouseSelecting() bool                                                                                                               // function
	IsEmpty() bool                                                                                                                        // function
	PasteFromClipboard() bool                                                                                                             // function
	ScrollIntoView(Node IVirtualNode, Center bool, Horizontally bool) bool                                                                // function
	ScrollIntoView1(Column TColumnIndex, Center bool) bool                                                                                // function
	Nodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                           // function
	CheckedNodes(State TCheckState, ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                 // function
	ChildNodes(Node IVirtualNode) IVTVirtualNodeEnumeration                                                                               // function
	CutCopyNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                    // function
	InitializedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                // function
	LeafNodes() IVTVirtualNodeEnumeration                                                                                                 // function
	LevelNodes(NodeLevel uint32) IVTVirtualNodeEnumeration                                                                                // function
	NoInitNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                     // function
	SelectedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration                                                                   // function
	VisibleNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration                           // function
	VisibleChildNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration                                                  // function
	VisibleChildNoInitNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration                                            // function
	VisibleNoInitNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration                     // function
	AddFromStream(Stream IStream, TargetNode IVirtualNode)                                                                                // procedure
	BeginSynch()                                                                                                                          // procedure
	BeginUpdate()                                                                                                                         // procedure
	CancelCutOrCopy()                                                                                                                     // procedure
	CancelOperation()                                                                                                                     // procedure
	Clear()                                                                                                                               // procedure
	ClearChecked()                                                                                                                        // procedure
	ClearSelection()                                                                                                                      // procedure
	CopyToClipboard()                                                                                                                     // procedure
	CutToClipboard()                                                                                                                      // procedure
	DeleteChildren(Node IVirtualNode, ResetHasChildren bool)                                                                              // procedure
	DeleteNode(Node IVirtualNode, Reindex bool)                                                                                           // procedure
	DeleteSelectedNodes()                                                                                                                 // procedure
	EndSynch()                                                                                                                            // procedure
	EndUpdate()                                                                                                                           // procedure
	EnsureNodeSelected()                                                                                                                  // procedure
	FinishCutOrCopy()                                                                                                                     // procedure
	FlushClipboard()                                                                                                                      // procedure
	FullCollapse(Node IVirtualNode)                                                                                                       // procedure
	FullExpand(Node IVirtualNode)                                                                                                         // procedure
	GetTextInfo(Node IVirtualNode, Column TColumnIndex, AFont IFont, R *TRect, OutText *string)                                           // procedure
	InvalidateChildren(Node IVirtualNode, Recursive bool)                                                                                 // procedure
	InvalidateColumn(Column TColumnIndex)                                                                                                 // procedure
	InvalidateToBottom(Node IVirtualNode)                                                                                                 // procedure
	InvertSelection(VisibleOnly bool)                                                                                                     // procedure
	LoadFromFile(FileName string)                                                                                                         // procedure
	LoadFromStream(Stream IStream)                                                                                                        // procedure
	MeasureItemHeight(Canvas ICanvas, Node IVirtualNode)                                                                                  // procedure
	MoveTo(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool)                                                        // procedure
	MoveTo1(Node IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool)                                          // procedure
	PaintTree(TargetCanvas ICanvas, Window *TRect, Target *TPoint, PaintOptions TVTInternalPaintOptions, PixelFormat TPixelFormat)        // procedure
	RepaintNode(Node IVirtualNode)                                                                                                        // procedure
	ReinitChildren(Node IVirtualNode, Recursive bool)                                                                                     // procedure
	ReinitNode(Node IVirtualNode, Recursive bool)                                                                                         // procedure
	ResetNode(Node IVirtualNode)                                                                                                          // procedure
	SaveToFile(FileName string)                                                                                                           // procedure
	SaveToStream(Stream IStream, Node IVirtualNode)                                                                                       // procedure
	SelectAll(VisibleOnly bool)                                                                                                           // procedure
	Sort(Node IVirtualNode, Column TColumnIndex, Direction TSortDirection, DoInit bool)                                                   // procedure
	SortTree(Column TColumnIndex, Direction TSortDirection, DoInit bool)                                                                  // procedure
	ToggleNode(Node IVirtualNode)                                                                                                         // procedure
	UpdateHorizontalRange()                                                                                                               // procedure
	UpdateHorizontalScrollBar(DoRepaint bool)                                                                                             // procedure
	UpdateRanges()                                                                                                                        // procedure
	UpdateScrollBars(DoRepaint bool)                                                                                                      // procedure
	UpdateVerticalRange()                                                                                                                 // procedure
	UpdateVerticalScrollBar(DoRepaint bool)                                                                                               // procedure
	ValidateChildren(Node IVirtualNode, Recursive bool)                                                                                   // procedure
	ValidateNode(Node IVirtualNode, Recursive bool)                                                                                       // procedure
}

// TBaseVirtualTree Parent: TCustomControl
// ----- TBaseVirtualTree
type TBaseVirtualTree struct {
	TCustomControl
}

func NewBaseVirtualTree(AOwner IComponent) IBaseVirtualTree {
	r1 := LCL().SysCallN(229, GetObjectUintptr(AOwner))
	return AsBaseVirtualTree(r1)
}

func (m *TBaseVirtualTree) BottomNode() IVirtualNode {
	r1 := LCL().SysCallN(209, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetBottomNode(AValue IVirtualNode) {
	LCL().SysCallN(209, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) CheckedCount() int32 {
	r1 := LCL().SysCallN(217, m.Instance())
	return int32(r1)
}

func (m *TBaseVirtualTree) CheckImages() ICustomImageList {
	r1 := LCL().SysCallN(214, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TBaseVirtualTree) CheckState(Node IVirtualNode) TCheckState {
	r1 := LCL().SysCallN(215, 0, m.Instance(), GetObjectUintptr(Node))
	return TCheckState(r1)
}

func (m *TBaseVirtualTree) SetCheckState(Node IVirtualNode, AValue TCheckState) {
	LCL().SysCallN(215, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) CheckType(Node IVirtualNode) TCheckType {
	r1 := LCL().SysCallN(216, 0, m.Instance(), GetObjectUintptr(Node))
	return TCheckType(r1)
}

func (m *TBaseVirtualTree) SetCheckType(Node IVirtualNode, AValue TCheckType) {
	LCL().SysCallN(216, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) ChildCount(Node IVirtualNode) uint32 {
	r1 := LCL().SysCallN(219, 0, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) SetChildCount(Node IVirtualNode, AValue uint32) {
	LCL().SysCallN(219, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) ChildrenInitialized(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(221, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) CutCopyCount() int32 {
	r1 := LCL().SysCallN(230, m.Instance())
	return int32(r1)
}

func (m *TBaseVirtualTree) DragImage() IVTDragImage {
	r1 := LCL().SysCallN(236, m.Instance())
	return AsVTDragImage(r1)
}

func (m *TBaseVirtualTree) DropTargetNode() IVirtualNode {
	r1 := LCL().SysCallN(237, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetDropTargetNode(AValue IVirtualNode) {
	LCL().SysCallN(237, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) EmptyListMessage() string {
	r1 := LCL().SysCallN(239, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TBaseVirtualTree) SetEmptyListMessage(AValue string) {
	LCL().SysCallN(239, 1, m.Instance(), PascalStr(AValue))
}

func (m *TBaseVirtualTree) Expanded(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(244, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetExpanded(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(244, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) FocusedColumn() TColumnIndex {
	r1 := LCL().SysCallN(247, 0, m.Instance(), 0)
	return TColumnIndex(r1)
}

func (m *TBaseVirtualTree) SetFocusedColumn(AValue TColumnIndex) {
	LCL().SysCallN(247, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) FocusedNode() IVirtualNode {
	r1 := LCL().SysCallN(248, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetFocusedNode(AValue IVirtualNode) {
	LCL().SysCallN(248, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) FullyVisible(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(251, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetFullyVisible(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(251, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) HasChildren(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(319, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetHasChildren(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(319, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) HotNode() IVirtualNode {
	r1 := LCL().SysCallN(320, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) IsDisabled(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(328, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetIsDisabled(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(328, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) IsEffectivelyFiltered(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(330, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsEffectivelyVisible(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(331, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsFiltered(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(333, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetIsFiltered(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(333, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) NodeIsVisible(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(345, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetNodeIsVisible(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(345, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) MultiLine(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(342, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetMultiLine(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(342, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) NodeHeight(Node IVirtualNode) uint32 {
	r1 := LCL().SysCallN(344, 0, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) SetNodeHeight(Node IVirtualNode, AValue uint32) {
	LCL().SysCallN(344, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) NodeParent(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(346, 0, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetNodeParent(Node IVirtualNode, AValue IVirtualNode) {
	LCL().SysCallN(346, 1, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) OffsetX() int32 {
	r1 := LCL().SysCallN(348, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TBaseVirtualTree) SetOffsetX(AValue int32) {
	LCL().SysCallN(348, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) OffsetXY() (resultPoint TPoint) {
	LCL().SysCallN(349, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TBaseVirtualTree) SetOffsetXY(AValue *TPoint) {
	LCL().SysCallN(349, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TBaseVirtualTree) OffsetY() int32 {
	r1 := LCL().SysCallN(350, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TBaseVirtualTree) SetOffsetY(AValue int32) {
	LCL().SysCallN(350, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) OperationCount() uint32 {
	r1 := LCL().SysCallN(351, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) RootNode() IVirtualNode {
	r1 := LCL().SysCallN(358, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SearchBuffer() string {
	r1 := LCL().SysCallN(363, m.Instance())
	return GoStr(r1)
}

func (m *TBaseVirtualTree) Selected(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(365, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetSelected(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(365, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) SelectionLocked() bool {
	r1 := LCL().SysCallN(368, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetSelectionLocked(AValue bool) {
	LCL().SysCallN(368, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBaseVirtualTree) TotalCount() uint32 {
	r1 := LCL().SysCallN(373, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) TreeStates() TVirtualTreeStates {
	r1 := LCL().SysCallN(374, 0, m.Instance(), 0)
	return TVirtualTreeStates(r1)
}

func (m *TBaseVirtualTree) SetTreeStates(AValue TVirtualTreeStates) {
	LCL().SysCallN(374, 1, m.Instance(), uintptr(AValue))
}

func (m *TBaseVirtualTree) SelectedCount() int32 {
	r1 := LCL().SysCallN(366, m.Instance())
	return int32(r1)
}

func (m *TBaseVirtualTree) TopNode() IVirtualNode {
	r1 := LCL().SysCallN(372, 0, m.Instance(), 0)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) SetTopNode(AValue IVirtualNode) {
	LCL().SysCallN(372, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBaseVirtualTree) VerticalAlignment(Node IVirtualNode) Byte {
	r1 := LCL().SysCallN(384, 0, m.Instance(), GetObjectUintptr(Node))
	return Byte(r1)
}

func (m *TBaseVirtualTree) SetVerticalAlignment(Node IVirtualNode, AValue Byte) {
	LCL().SysCallN(384, 1, m.Instance(), GetObjectUintptr(Node), uintptr(AValue))
}

func (m *TBaseVirtualTree) VisibleCount() uint32 {
	r1 := LCL().SysCallN(387, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) VisiblePath(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(390, 0, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) SetVisiblePath(Node IVirtualNode, AValue bool) {
	LCL().SysCallN(390, 1, m.Instance(), GetObjectUintptr(Node), PascalBool(AValue))
}

func (m *TBaseVirtualTree) UpdateCount() uint32 {
	r1 := LCL().SysCallN(375, m.Instance())
	return uint32(r1)
}

func (m *TBaseVirtualTree) AbsoluteIndex(Node IVirtualNode) uint32 {
	r1 := LCL().SysCallN(204, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) AddChild(Parent IVirtualNode, UserData uintptr) IVirtualNode {
	r1 := LCL().SysCallN(205, m.Instance(), GetObjectUintptr(Parent), uintptr(UserData))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) CancelEditNode() bool {
	r1 := LCL().SysCallN(212, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) CanEdit(Node IVirtualNode, Column TColumnIndex) bool {
	r1 := LCL().SysCallN(210, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) CopyTo(Source IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode {
	r1 := LCL().SysCallN(226, m.Instance(), GetObjectUintptr(Source), GetObjectUintptr(Tree), uintptr(Mode), PascalBool(ChildrenOnly))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) CopyTo1(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool) IVirtualNode {
	r1 := LCL().SysCallN(227, m.Instance(), GetObjectUintptr(Source), GetObjectUintptr(Target), uintptr(Mode), PascalBool(ChildrenOnly))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) EditNode(Node IVirtualNode, Column TColumnIndex) bool {
	r1 := LCL().SysCallN(238, m.Instance(), GetObjectUintptr(Node), uintptr(Column))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) EndEditNode() bool {
	r1 := LCL().SysCallN(240, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) GetDisplayRect(Node IVirtualNode, Column TColumnIndex, TextOnly bool, Unclipped bool, ApplyCellContentMargin bool) (resultRect TRect) {
	LCL().SysCallN(252, m.Instance(), GetObjectUintptr(Node), uintptr(Column), PascalBool(TextOnly), PascalBool(Unclipped), PascalBool(ApplyCellContentMargin), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TBaseVirtualTree) GetEffectivelyFiltered(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(253, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) GetEffectivelyVisible(Node IVirtualNode) bool {
	r1 := LCL().SysCallN(254, m.Instance(), GetObjectUintptr(Node))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) GetFirst(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(255, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstChecked(State TCheckState, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(256, m.Instance(), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstChild(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(257, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstChildNoInit(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(258, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstCutCopy(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(259, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstInitialized(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(260, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstLeaf() IVirtualNode {
	r1 := LCL().SysCallN(261, m.Instance())
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstLevel(NodeLevel uint32) IVirtualNode {
	r1 := LCL().SysCallN(262, m.Instance(), uintptr(NodeLevel))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstNoInit(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(263, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstSelected(ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(264, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(265, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(266, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(267, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetFirstVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(268, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLast(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(269, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(272, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(273, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastChild(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(270, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastChildNoInit(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(271, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisible(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(274, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisibleChild(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(275, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisibleChildNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(276, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetLastVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(277, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetMaxColumnWidth(Column TColumnIndex, UseSmartColumnWidth bool) int32 {
	r1 := LCL().SysCallN(278, m.Instance(), uintptr(Column), PascalBool(UseSmartColumnWidth))
	return int32(r1)
}

func (m *TBaseVirtualTree) GetNext(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(279, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(280, m.Instance(), GetObjectUintptr(Node), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextChecked1(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(281, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(282, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(283, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextLeaf(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(284, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode {
	r1 := LCL().SysCallN(285, m.Instance(), GetObjectUintptr(Node), uintptr(NodeLevel))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(286, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(287, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextSibling(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(288, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextSiblingNoInit(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(289, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(290, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(291, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(292, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNextVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(293, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeAt(P *TPoint) IVirtualNode {
	r1 := LCL().SysCallN(294, m.Instance(), uintptr(unsafePointer(P)))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeAt1(X, Y int32) IVirtualNode {
	r1 := LCL().SysCallN(295, m.Instance(), uintptr(X), uintptr(Y))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeAt2(X, Y int32, Relative bool, NodeTop *int32) IVirtualNode {
	var result2 uintptr
	r1 := LCL().SysCallN(296, m.Instance(), uintptr(X), uintptr(Y), PascalBool(Relative), uintptr(unsafePointer(&result2)))
	*NodeTop = int32(result2)
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetNodeData(Node IVirtualNode) uintptr {
	r1 := LCL().SysCallN(297, m.Instance(), GetObjectUintptr(Node))
	return uintptr(r1)
}

func (m *TBaseVirtualTree) GetNodeLevel(Node IVirtualNode) uint32 {
	r1 := LCL().SysCallN(298, m.Instance(), GetObjectUintptr(Node))
	return uint32(r1)
}

func (m *TBaseVirtualTree) GetPrevious(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(299, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousChecked(Node IVirtualNode, State TCheckState, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(300, m.Instance(), GetObjectUintptr(Node), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousCutCopy(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(301, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousInitialized(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(302, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousLeaf(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(303, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousLevel(Node IVirtualNode, NodeLevel uint32) IVirtualNode {
	r1 := LCL().SysCallN(304, m.Instance(), GetObjectUintptr(Node), uintptr(NodeLevel))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(305, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousSelected(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(306, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousSibling(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(307, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousSiblingNoInit(Node IVirtualNode) IVirtualNode {
	r1 := LCL().SysCallN(308, m.Instance(), GetObjectUintptr(Node))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisible(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(309, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisibleNoInit(Node IVirtualNode, ConsiderChildrenAbove bool) IVirtualNode {
	r1 := LCL().SysCallN(310, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisibleSibling(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(311, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetPreviousVisibleSiblingNoInit(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(312, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) GetSortedCutCopySet(Resolve bool) TNodeArray {
	r1 := LCL().SysCallN(313, m.Instance(), PascalBool(Resolve))
	return TNodeArray(r1)
}

func (m *TBaseVirtualTree) GetSortedSelection(Resolve bool) TNodeArray {
	r1 := LCL().SysCallN(314, m.Instance(), PascalBool(Resolve))
	return TNodeArray(r1)
}

func (m *TBaseVirtualTree) GetTreeRect() (resultRect TRect) {
	LCL().SysCallN(316, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TBaseVirtualTree) GetVisibleParent(Node IVirtualNode, IncludeFiltered bool) IVirtualNode {
	r1 := LCL().SysCallN(317, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) HasAsParent(Node, PotentialParent IVirtualNode) bool {
	r1 := LCL().SysCallN(318, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(PotentialParent))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) InsertNode(Node IVirtualNode, Mode TVTNodeAttachMode, UserData uintptr) IVirtualNode {
	r1 := LCL().SysCallN(322, m.Instance(), GetObjectUintptr(Node), uintptr(Mode), uintptr(UserData))
	return AsVirtualNode(r1)
}

func (m *TBaseVirtualTree) InvalidateNode(Node IVirtualNode) (resultRect TRect) {
	LCL().SysCallN(325, m.Instance(), GetObjectUintptr(Node), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TBaseVirtualTree) IsEditing() bool {
	r1 := LCL().SysCallN(329, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsMouseSelecting() bool {
	r1 := LCL().SysCallN(334, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) IsEmpty() bool {
	r1 := LCL().SysCallN(332, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) PasteFromClipboard() bool {
	r1 := LCL().SysCallN(353, m.Instance())
	return GoBool(r1)
}

func (m *TBaseVirtualTree) ScrollIntoView(Node IVirtualNode, Center bool, Horizontally bool) bool {
	r1 := LCL().SysCallN(361, m.Instance(), GetObjectUintptr(Node), PascalBool(Center), PascalBool(Horizontally))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) ScrollIntoView1(Column TColumnIndex, Center bool) bool {
	r1 := LCL().SysCallN(362, m.Instance(), uintptr(Column), PascalBool(Center))
	return GoBool(r1)
}

func (m *TBaseVirtualTree) Nodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(347, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) CheckedNodes(State TCheckState, ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(218, m.Instance(), uintptr(State), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) ChildNodes(Node IVirtualNode) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(220, m.Instance(), GetObjectUintptr(Node))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) CutCopyNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(231, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) InitializedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(321, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) LeafNodes() IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(335, m.Instance())
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) LevelNodes(NodeLevel uint32) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(336, m.Instance(), uintptr(NodeLevel))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) NoInitNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(343, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) SelectedNodes(ConsiderChildrenAbove bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(367, m.Instance(), PascalBool(ConsiderChildrenAbove))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(389, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleChildNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(386, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleChildNoInitNodes(Node IVirtualNode, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(385, m.Instance(), GetObjectUintptr(Node), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func (m *TBaseVirtualTree) VisibleNoInitNodes(Node IVirtualNode, ConsiderChildrenAbove bool, IncludeFiltered bool) IVTVirtualNodeEnumeration {
	r1 := LCL().SysCallN(388, m.Instance(), GetObjectUintptr(Node), PascalBool(ConsiderChildrenAbove), PascalBool(IncludeFiltered))
	return AsVTVirtualNodeEnumeration(r1)
}

func BaseVirtualTreeClass() TClass {
	ret := LCL().SysCallN(222)
	return TClass(ret)
}

func (m *TBaseVirtualTree) AddFromStream(Stream IStream, TargetNode IVirtualNode) {
	LCL().SysCallN(206, m.Instance(), GetObjectUintptr(Stream), GetObjectUintptr(TargetNode))
}

func (m *TBaseVirtualTree) BeginSynch() {
	LCL().SysCallN(207, m.Instance())
}

func (m *TBaseVirtualTree) BeginUpdate() {
	LCL().SysCallN(208, m.Instance())
}

func (m *TBaseVirtualTree) CancelCutOrCopy() {
	LCL().SysCallN(211, m.Instance())
}

func (m *TBaseVirtualTree) CancelOperation() {
	LCL().SysCallN(213, m.Instance())
}

func (m *TBaseVirtualTree) Clear() {
	LCL().SysCallN(223, m.Instance())
}

func (m *TBaseVirtualTree) ClearChecked() {
	LCL().SysCallN(224, m.Instance())
}

func (m *TBaseVirtualTree) ClearSelection() {
	LCL().SysCallN(225, m.Instance())
}

func (m *TBaseVirtualTree) CopyToClipboard() {
	LCL().SysCallN(228, m.Instance())
}

func (m *TBaseVirtualTree) CutToClipboard() {
	LCL().SysCallN(232, m.Instance())
}

func (m *TBaseVirtualTree) DeleteChildren(Node IVirtualNode, ResetHasChildren bool) {
	LCL().SysCallN(233, m.Instance(), GetObjectUintptr(Node), PascalBool(ResetHasChildren))
}

func (m *TBaseVirtualTree) DeleteNode(Node IVirtualNode, Reindex bool) {
	LCL().SysCallN(234, m.Instance(), GetObjectUintptr(Node), PascalBool(Reindex))
}

func (m *TBaseVirtualTree) DeleteSelectedNodes() {
	LCL().SysCallN(235, m.Instance())
}

func (m *TBaseVirtualTree) EndSynch() {
	LCL().SysCallN(241, m.Instance())
}

func (m *TBaseVirtualTree) EndUpdate() {
	LCL().SysCallN(242, m.Instance())
}

func (m *TBaseVirtualTree) EnsureNodeSelected() {
	LCL().SysCallN(243, m.Instance())
}

func (m *TBaseVirtualTree) FinishCutOrCopy() {
	LCL().SysCallN(245, m.Instance())
}

func (m *TBaseVirtualTree) FlushClipboard() {
	LCL().SysCallN(246, m.Instance())
}

func (m *TBaseVirtualTree) FullCollapse(Node IVirtualNode) {
	LCL().SysCallN(249, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) FullExpand(Node IVirtualNode) {
	LCL().SysCallN(250, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) GetTextInfo(Node IVirtualNode, Column TColumnIndex, AFont IFont, R *TRect, OutText *string) {
	var result3 uintptr
	var result4 uintptr
	LCL().SysCallN(315, m.Instance(), GetObjectUintptr(Node), uintptr(Column), GetObjectUintptr(AFont), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)))
	*R = *(*TRect)(getPointer(result3))
	*OutText = GoStr(result4)
}

func (m *TBaseVirtualTree) InvalidateChildren(Node IVirtualNode, Recursive bool) {
	LCL().SysCallN(323, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) InvalidateColumn(Column TColumnIndex) {
	LCL().SysCallN(324, m.Instance(), uintptr(Column))
}

func (m *TBaseVirtualTree) InvalidateToBottom(Node IVirtualNode) {
	LCL().SysCallN(326, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) InvertSelection(VisibleOnly bool) {
	LCL().SysCallN(327, m.Instance(), PascalBool(VisibleOnly))
}

func (m *TBaseVirtualTree) LoadFromFile(FileName string) {
	LCL().SysCallN(337, m.Instance(), PascalStr(FileName))
}

func (m *TBaseVirtualTree) LoadFromStream(Stream IStream) {
	LCL().SysCallN(338, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TBaseVirtualTree) MeasureItemHeight(Canvas ICanvas, Node IVirtualNode) {
	LCL().SysCallN(339, m.Instance(), GetObjectUintptr(Canvas), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) MoveTo(Source, Target IVirtualNode, Mode TVTNodeAttachMode, ChildrenOnly bool) {
	LCL().SysCallN(340, m.Instance(), GetObjectUintptr(Source), GetObjectUintptr(Target), uintptr(Mode), PascalBool(ChildrenOnly))
}

func (m *TBaseVirtualTree) MoveTo1(Node IVirtualNode, Tree IBaseVirtualTree, Mode TVTNodeAttachMode, ChildrenOnly bool) {
	LCL().SysCallN(341, m.Instance(), GetObjectUintptr(Node), GetObjectUintptr(Tree), uintptr(Mode), PascalBool(ChildrenOnly))
}

func (m *TBaseVirtualTree) PaintTree(TargetCanvas ICanvas, Window *TRect, Target *TPoint, PaintOptions TVTInternalPaintOptions, PixelFormat TPixelFormat) {
	LCL().SysCallN(352, m.Instance(), GetObjectUintptr(TargetCanvas), uintptr(unsafePointer(Window)), uintptr(unsafePointer(Target)), uintptr(PaintOptions), uintptr(PixelFormat))
}

func (m *TBaseVirtualTree) RepaintNode(Node IVirtualNode) {
	LCL().SysCallN(356, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) ReinitChildren(Node IVirtualNode, Recursive bool) {
	LCL().SysCallN(354, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) ReinitNode(Node IVirtualNode, Recursive bool) {
	LCL().SysCallN(355, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) ResetNode(Node IVirtualNode) {
	LCL().SysCallN(357, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) SaveToFile(FileName string) {
	LCL().SysCallN(359, m.Instance(), PascalStr(FileName))
}

func (m *TBaseVirtualTree) SaveToStream(Stream IStream, Node IVirtualNode) {
	LCL().SysCallN(360, m.Instance(), GetObjectUintptr(Stream), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) SelectAll(VisibleOnly bool) {
	LCL().SysCallN(364, m.Instance(), PascalBool(VisibleOnly))
}

func (m *TBaseVirtualTree) Sort(Node IVirtualNode, Column TColumnIndex, Direction TSortDirection, DoInit bool) {
	LCL().SysCallN(369, m.Instance(), GetObjectUintptr(Node), uintptr(Column), uintptr(Direction), PascalBool(DoInit))
}

func (m *TBaseVirtualTree) SortTree(Column TColumnIndex, Direction TSortDirection, DoInit bool) {
	LCL().SysCallN(370, m.Instance(), uintptr(Column), uintptr(Direction), PascalBool(DoInit))
}

func (m *TBaseVirtualTree) ToggleNode(Node IVirtualNode) {
	LCL().SysCallN(371, m.Instance(), GetObjectUintptr(Node))
}

func (m *TBaseVirtualTree) UpdateHorizontalRange() {
	LCL().SysCallN(376, m.Instance())
}

func (m *TBaseVirtualTree) UpdateHorizontalScrollBar(DoRepaint bool) {
	LCL().SysCallN(377, m.Instance(), PascalBool(DoRepaint))
}

func (m *TBaseVirtualTree) UpdateRanges() {
	LCL().SysCallN(378, m.Instance())
}

func (m *TBaseVirtualTree) UpdateScrollBars(DoRepaint bool) {
	LCL().SysCallN(379, m.Instance(), PascalBool(DoRepaint))
}

func (m *TBaseVirtualTree) UpdateVerticalRange() {
	LCL().SysCallN(380, m.Instance())
}

func (m *TBaseVirtualTree) UpdateVerticalScrollBar(DoRepaint bool) {
	LCL().SysCallN(381, m.Instance(), PascalBool(DoRepaint))
}

func (m *TBaseVirtualTree) ValidateChildren(Node IVirtualNode, Recursive bool) {
	LCL().SysCallN(382, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}

func (m *TBaseVirtualTree) ValidateNode(Node IVirtualNode, Recursive bool) {
	LCL().SysCallN(383, m.Instance(), GetObjectUintptr(Node), PascalBool(Recursive))
}
