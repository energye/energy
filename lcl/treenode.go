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
	"unsafe"
)

// ITreeNode Parent: IPersistent
type ITreeNode interface {
	IPersistent
	AbsoluteIndex() int32                                          // property
	Count() int32                                                  // property
	Cut() bool                                                     // property
	SetCut(AValue bool)                                            // property
	Data() uintptr                                                 // property
	SetData(AValue uintptr)                                        // property
	Deleting() bool                                                // property
	DropTarget() bool                                              // property
	SetDropTarget(AValue bool)                                     // property
	Expanded() bool                                                // property
	SetExpanded(AValue bool)                                       // property
	Focused() bool                                                 // property
	SetFocused(AValue bool)                                        // property
	Handle() THandle                                               // property
	HasChildren() bool                                             // property
	SetHasChildren(AValue bool)                                    // property
	Height() int32                                                 // property
	SetHeight(AValue int32)                                        // property
	ImageIndex() TImageIndex                                       // property
	SetImageIndex(AValue TImageIndex)                              // property
	Index() int32                                                  // property
	SetIndex(AValue int32)                                         // property
	IsFullHeightVisible() bool                                     // property
	IsVisible() bool                                               // property
	Items(ItemIndex int32) ITreeNode                               // property
	SetItems(ItemIndex int32, AValue ITreeNode)                    // property
	Level() int32                                                  // property
	MultiSelected() bool                                           // property
	SetMultiSelected(AValue bool)                                  // property
	NodeEffect() TGraphicsDrawEffect                               // property
	SetNodeEffect(AValue TGraphicsDrawEffect)                      // property
	OverlayIndex() int32                                           // property
	SetOverlayIndex(AValue int32)                                  // property
	Owner() ITreeNodes                                             // property
	Parent() ITreeNode                                             // property
	Selected() bool                                                // property
	SetSelected(AValue bool)                                       // property
	SelectedIndex() int32                                          // property
	SetSelectedIndex(AValue int32)                                 // property
	StateIndex() int32                                             // property
	SetStateIndex(AValue int32)                                    // property
	States() TNodeStates                                           // property
	SubTreeCount() int32                                           // property
	Text() string                                                  // property
	SetText(AValue string)                                         // property
	Top() int32                                                    // property
	TreeNodes() ITreeNodes                                         // property
	TreeView() ICustomTreeView                                     // property
	Visible() bool                                                 // property
	SetVisible(AValue bool)                                        // property
	Enabled() bool                                                 // property
	SetEnabled(AValue bool)                                        // property
	AlphaSort() bool                                               // function
	Bottom() int32                                                 // function
	BottomExpanded() int32                                         // function
	CustomSort(fn TTreeNodeCompare) bool                           // function
	DefaultTreeViewSort(Node1, Node2 ITreeNode) int32              // function
	DisplayExpandSignLeft() int32                                  // function
	DisplayExpandSignRect() (resultRect TRect)                     // function
	DisplayExpandSignRight() int32                                 // function
	DisplayIconLeft() int32                                        // function
	DisplayRect(TextOnly bool) (resultRect TRect)                  // function
	DisplayStateIconLeft() int32                                   // function
	DisplayTextLeft() int32                                        // function
	DisplayTextRight() int32                                       // function
	EditText() bool                                                // function
	FindNode(NodeText string) ITreeNode                            // function
	GetFirstChild() ITreeNode                                      // function
	GetFirstSibling() ITreeNode                                    // function
	GetFirstVisibleChild(aEnabledOnly bool) ITreeNode              // function
	GetHandle() THandle                                            // function
	GetLastChild() ITreeNode                                       // function
	GetLastSibling() ITreeNode                                     // function
	GetLastSubChild() ITreeNode                                    // function
	GetLastVisibleChild(aEnabledOnly bool) ITreeNode               // function
	GetNext() ITreeNode                                            // function
	GetNextChild(AValue ITreeNode) ITreeNode                       // function
	GetNextExpanded(aEnabledOnly bool) ITreeNode                   // function
	GetNextMultiSelected() ITreeNode                               // function
	GetNextSibling() ITreeNode                                     // function
	GetNextSkipChildren() ITreeNode                                // function
	GetNextVisible(aEnabledOnly bool) ITreeNode                    // function
	GetNextVisibleSibling(aEnabledOnly bool) ITreeNode             // function
	GetParentNodeOfAbsoluteLevel(TheAbsoluteLevel int32) ITreeNode // function
	GetPrev() ITreeNode                                            // function
	GetPrevChild(AValue ITreeNode) ITreeNode                       // function
	GetPrevExpanded(aEnabledOnly bool) ITreeNode                   // function
	GetPrevMultiSelected() ITreeNode                               // function
	GetPrevSibling() ITreeNode                                     // function
	GetPrevVisible(aEnabledOnly bool) ITreeNode                    // function
	GetPrevVisibleSibling(aEnabledOnly bool) ITreeNode             // function
	GetTextPath() string                                           // function
	HasAsParent(AValue ITreeNode) bool                             // function
	IndexOf(AValue ITreeNode) int32                                // function
	IndexOfText(NodeText string) int32                             // function
	Collapse(Recurse bool)                                         // procedure
	ConsistencyCheck()                                             // procedure
	Delete()                                                       // procedure
	DeleteChildren()                                               // procedure
	EndEdit(Cancel bool)                                           // procedure
	Expand(Recurse bool)                                           // procedure
	ExpandParents()                                                // procedure
	FreeAllNodeData()                                              // procedure
	MakeVisible()                                                  // procedure
	MoveTo(Destination ITreeNode, Mode TNodeAttachMode)            // procedure
	MultiSelectGroup()                                             // procedure
	Update()                                                       // procedure
	WriteDebugReport(Prefix string, Recurse bool)                  // procedure
}

// TTreeNode Parent: TPersistent
type TTreeNode struct {
	TPersistent
	customSortPtr uintptr
}

func NewTreeNode(AnOwner ITreeNodes) ITreeNode {
	r1 := LCL().SysCallN(4922, GetObjectUintptr(AnOwner))
	return AsTreeNode(r1)
}

func (m *TTreeNode) AbsoluteIndex() int32 {
	r1 := LCL().SysCallN(4914, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) Count() int32 {
	r1 := LCL().SysCallN(4921, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) Cut() bool {
	r1 := LCL().SysCallN(4924, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetCut(AValue bool) {
	LCL().SysCallN(4924, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Data() uintptr {
	r1 := LCL().SysCallN(4925, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TTreeNode) SetData(AValue uintptr) {
	LCL().SysCallN(4925, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) Deleting() bool {
	r1 := LCL().SysCallN(4929, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) DropTarget() bool {
	r1 := LCL().SysCallN(4938, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetDropTarget(AValue bool) {
	LCL().SysCallN(4938, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Expanded() bool {
	r1 := LCL().SysCallN(4944, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetExpanded(AValue bool) {
	LCL().SysCallN(4944, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Focused() bool {
	r1 := LCL().SysCallN(4946, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetFocused(AValue bool) {
	LCL().SysCallN(4946, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Handle() THandle {
	r1 := LCL().SysCallN(4973, m.Instance())
	return THandle(r1)
}

func (m *TTreeNode) HasChildren() bool {
	r1 := LCL().SysCallN(4975, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetHasChildren(AValue bool) {
	LCL().SysCallN(4975, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Height() int32 {
	r1 := LCL().SysCallN(4976, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetHeight(AValue int32) {
	LCL().SysCallN(4976, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(4977, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TTreeNode) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(4977, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) Index() int32 {
	r1 := LCL().SysCallN(4978, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetIndex(AValue int32) {
	LCL().SysCallN(4978, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) IsFullHeightVisible() bool {
	r1 := LCL().SysCallN(4981, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) IsVisible() bool {
	r1 := LCL().SysCallN(4982, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) Items(ItemIndex int32) ITreeNode {
	r1 := LCL().SysCallN(4983, 0, m.Instance(), uintptr(ItemIndex))
	return AsTreeNode(r1)
}

func (m *TTreeNode) SetItems(ItemIndex int32, AValue ITreeNode) {
	LCL().SysCallN(4983, 1, m.Instance(), uintptr(ItemIndex), GetObjectUintptr(AValue))
}

func (m *TTreeNode) Level() int32 {
	r1 := LCL().SysCallN(4984, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) MultiSelected() bool {
	r1 := LCL().SysCallN(4988, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetMultiSelected(AValue bool) {
	LCL().SysCallN(4988, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) NodeEffect() TGraphicsDrawEffect {
	r1 := LCL().SysCallN(4989, 0, m.Instance(), 0)
	return TGraphicsDrawEffect(r1)
}

func (m *TTreeNode) SetNodeEffect(AValue TGraphicsDrawEffect) {
	LCL().SysCallN(4989, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) OverlayIndex() int32 {
	r1 := LCL().SysCallN(4990, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetOverlayIndex(AValue int32) {
	LCL().SysCallN(4990, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) Owner() ITreeNodes {
	r1 := LCL().SysCallN(4991, m.Instance())
	return AsTreeNodes(r1)
}

func (m *TTreeNode) Parent() ITreeNode {
	r1 := LCL().SysCallN(4992, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) Selected() bool {
	r1 := LCL().SysCallN(4993, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetSelected(AValue bool) {
	LCL().SysCallN(4993, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) SelectedIndex() int32 {
	r1 := LCL().SysCallN(4994, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetSelectedIndex(AValue int32) {
	LCL().SysCallN(4994, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) StateIndex() int32 {
	r1 := LCL().SysCallN(4995, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TTreeNode) SetStateIndex(AValue int32) {
	LCL().SysCallN(4995, 1, m.Instance(), uintptr(AValue))
}

func (m *TTreeNode) States() TNodeStates {
	r1 := LCL().SysCallN(4996, m.Instance())
	return TNodeStates(r1)
}

func (m *TTreeNode) SubTreeCount() int32 {
	r1 := LCL().SysCallN(4997, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) Text() string {
	r1 := LCL().SysCallN(4998, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TTreeNode) SetText(AValue string) {
	LCL().SysCallN(4998, 1, m.Instance(), PascalStr(AValue))
}

func (m *TTreeNode) Top() int32 {
	r1 := LCL().SysCallN(4999, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) TreeNodes() ITreeNodes {
	r1 := LCL().SysCallN(5000, m.Instance())
	return AsTreeNodes(r1)
}

func (m *TTreeNode) TreeView() ICustomTreeView {
	r1 := LCL().SysCallN(5001, m.Instance())
	return AsCustomTreeView(r1)
}

func (m *TTreeNode) Visible() bool {
	r1 := LCL().SysCallN(5003, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetVisible(AValue bool) {
	LCL().SysCallN(5003, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) Enabled() bool {
	r1 := LCL().SysCallN(4940, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTreeNode) SetEnabled(AValue bool) {
	LCL().SysCallN(4940, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTreeNode) AlphaSort() bool {
	r1 := LCL().SysCallN(4915, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) Bottom() int32 {
	r1 := LCL().SysCallN(4916, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) BottomExpanded() int32 {
	r1 := LCL().SysCallN(4917, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) CustomSort(fn TTreeNodeCompare) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := LCL().SysCallN(4923, m.Instance(), m.customSortPtr)
	return GoBool(r1)
}

func (m *TTreeNode) DefaultTreeViewSort(Node1, Node2 ITreeNode) int32 {
	r1 := LCL().SysCallN(4926, m.Instance(), GetObjectUintptr(Node1), GetObjectUintptr(Node2))
	return int32(r1)
}

func (m *TTreeNode) DisplayExpandSignLeft() int32 {
	r1 := LCL().SysCallN(4930, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayExpandSignRect() (resultRect TRect) {
	LCL().SysCallN(4931, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TTreeNode) DisplayExpandSignRight() int32 {
	r1 := LCL().SysCallN(4932, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayIconLeft() int32 {
	r1 := LCL().SysCallN(4933, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayRect(TextOnly bool) (resultRect TRect) {
	LCL().SysCallN(4934, m.Instance(), PascalBool(TextOnly), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TTreeNode) DisplayStateIconLeft() int32 {
	r1 := LCL().SysCallN(4935, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayTextLeft() int32 {
	r1 := LCL().SysCallN(4936, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) DisplayTextRight() int32 {
	r1 := LCL().SysCallN(4937, m.Instance())
	return int32(r1)
}

func (m *TTreeNode) EditText() bool {
	r1 := LCL().SysCallN(4939, m.Instance())
	return GoBool(r1)
}

func (m *TTreeNode) FindNode(NodeText string) ITreeNode {
	r1 := LCL().SysCallN(4945, m.Instance(), PascalStr(NodeText))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetFirstChild() ITreeNode {
	r1 := LCL().SysCallN(4948, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetFirstSibling() ITreeNode {
	r1 := LCL().SysCallN(4949, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetFirstVisibleChild(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4950, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetHandle() THandle {
	r1 := LCL().SysCallN(4951, m.Instance())
	return THandle(r1)
}

func (m *TTreeNode) GetLastChild() ITreeNode {
	r1 := LCL().SysCallN(4952, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetLastSibling() ITreeNode {
	r1 := LCL().SysCallN(4953, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetLastSubChild() ITreeNode {
	r1 := LCL().SysCallN(4954, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetLastVisibleChild(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4955, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNext() ITreeNode {
	r1 := LCL().SysCallN(4956, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextChild(AValue ITreeNode) ITreeNode {
	r1 := LCL().SysCallN(4957, m.Instance(), GetObjectUintptr(AValue))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextExpanded(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4958, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextMultiSelected() ITreeNode {
	r1 := LCL().SysCallN(4959, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextSibling() ITreeNode {
	r1 := LCL().SysCallN(4960, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextSkipChildren() ITreeNode {
	r1 := LCL().SysCallN(4961, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextVisible(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4962, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetNextVisibleSibling(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4963, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetParentNodeOfAbsoluteLevel(TheAbsoluteLevel int32) ITreeNode {
	r1 := LCL().SysCallN(4964, m.Instance(), uintptr(TheAbsoluteLevel))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrev() ITreeNode {
	r1 := LCL().SysCallN(4965, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevChild(AValue ITreeNode) ITreeNode {
	r1 := LCL().SysCallN(4966, m.Instance(), GetObjectUintptr(AValue))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevExpanded(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4967, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevMultiSelected() ITreeNode {
	r1 := LCL().SysCallN(4968, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevSibling() ITreeNode {
	r1 := LCL().SysCallN(4969, m.Instance())
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevVisible(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4970, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetPrevVisibleSibling(aEnabledOnly bool) ITreeNode {
	r1 := LCL().SysCallN(4971, m.Instance(), PascalBool(aEnabledOnly))
	return AsTreeNode(r1)
}

func (m *TTreeNode) GetTextPath() string {
	r1 := LCL().SysCallN(4972, m.Instance())
	return GoStr(r1)
}

func (m *TTreeNode) HasAsParent(AValue ITreeNode) bool {
	r1 := LCL().SysCallN(4974, m.Instance(), GetObjectUintptr(AValue))
	return GoBool(r1)
}

func (m *TTreeNode) IndexOf(AValue ITreeNode) int32 {
	r1 := LCL().SysCallN(4979, m.Instance(), GetObjectUintptr(AValue))
	return int32(r1)
}

func (m *TTreeNode) IndexOfText(NodeText string) int32 {
	r1 := LCL().SysCallN(4980, m.Instance(), PascalStr(NodeText))
	return int32(r1)
}

func TreeNodeClass() TClass {
	ret := LCL().SysCallN(4918)
	return TClass(ret)
}

func (m *TTreeNode) Collapse(Recurse bool) {
	LCL().SysCallN(4919, m.Instance(), PascalBool(Recurse))
}

func (m *TTreeNode) ConsistencyCheck() {
	LCL().SysCallN(4920, m.Instance())
}

func (m *TTreeNode) Delete() {
	LCL().SysCallN(4927, m.Instance())
}

func (m *TTreeNode) DeleteChildren() {
	LCL().SysCallN(4928, m.Instance())
}

func (m *TTreeNode) EndEdit(Cancel bool) {
	LCL().SysCallN(4941, m.Instance(), PascalBool(Cancel))
}

func (m *TTreeNode) Expand(Recurse bool) {
	LCL().SysCallN(4942, m.Instance(), PascalBool(Recurse))
}

func (m *TTreeNode) ExpandParents() {
	LCL().SysCallN(4943, m.Instance())
}

func (m *TTreeNode) FreeAllNodeData() {
	LCL().SysCallN(4947, m.Instance())
}

func (m *TTreeNode) MakeVisible() {
	LCL().SysCallN(4985, m.Instance())
}

func (m *TTreeNode) MoveTo(Destination ITreeNode, Mode TNodeAttachMode) {
	LCL().SysCallN(4986, m.Instance(), GetObjectUintptr(Destination), uintptr(Mode))
}

func (m *TTreeNode) MultiSelectGroup() {
	LCL().SysCallN(4987, m.Instance())
}

func (m *TTreeNode) Update() {
	LCL().SysCallN(5002, m.Instance())
}

func (m *TTreeNode) WriteDebugReport(Prefix string, Recurse bool) {
	LCL().SysCallN(5004, m.Instance(), PascalStr(Prefix), PascalBool(Recurse))
}
