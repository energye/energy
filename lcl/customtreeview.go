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

// ICustomTreeView Parent: ICustomControl
type ICustomTreeView interface {
	ICustomControl
	AccessibilityOn() bool                                                                                  // property
	SetAccessibilityOn(AValue bool)                                                                         // property
	BackgroundColor() TColor                                                                                // property
	SetBackgroundColor(AValue TColor)                                                                       // property
	BottomItem() ITreeNode                                                                                  // property
	SetBottomItem(AValue ITreeNode)                                                                         // property
	DefaultItemHeight() int32                                                                               // property
	SetDefaultItemHeight(AValue int32)                                                                      // property
	DropTarget() ITreeNode                                                                                  // property
	SetDropTarget(AValue ITreeNode)                                                                         // property
	ExpandSignColor() TColor                                                                                // property
	SetExpandSignColor(AValue TColor)                                                                       // property
	ExpandSignSize() int32                                                                                  // property
	SetExpandSignSize(AValue int32)                                                                         // property
	ExpandSignWidth() int32                                                                                 // property
	SetExpandSignWidth(AValue int32)                                                                        // property
	ExpandSignType() TTreeViewExpandSignType                                                                // property
	SetExpandSignType(AValue TTreeViewExpandSignType)                                                       // property
	Images() ICustomImageList                                                                               // property
	SetImages(AValue ICustomImageList)                                                                      // property
	ImagesWidth() int32                                                                                     // property
	SetImagesWidth(AValue int32)                                                                            // property
	InsertMarkNode() ITreeNode                                                                              // property
	SetInsertMarkNode(AValue ITreeNode)                                                                     // property
	InsertMarkType() TTreeViewInsertMarkType                                                                // property
	SetInsertMarkType(AValue TTreeViewInsertMarkType)                                                       // property
	Items() ITreeNodes                                                                                      // property
	SetItems(AValue ITreeNodes)                                                                             // property
	KeepCollapsedNodes() bool                                                                               // property
	SetKeepCollapsedNodes(AValue bool)                                                                      // property
	MultiSelectStyle() TMultiSelectStyle                                                                    // property
	SetMultiSelectStyle(AValue TMultiSelectStyle)                                                           // property
	Options() TTreeViewOptions                                                                              // property
	SetOptions(AValue TTreeViewOptions)                                                                     // property
	ScrollBars() TScrollStyle                                                                               // property
	SetScrollBars(AValue TScrollStyle)                                                                      // property
	Selected() ITreeNode                                                                                    // property
	SetSelected(AValue ITreeNode)                                                                           // property
	SelectionColor() TColor                                                                                 // property
	SetSelectionColor(AValue TColor)                                                                        // property
	SelectionCount() uint32                                                                                 // property
	SelectionFontColor() TColor                                                                             // property
	SetSelectionFontColor(AValue TColor)                                                                    // property
	SelectionFontColorUsed() bool                                                                           // property
	SetSelectionFontColorUsed(AValue bool)                                                                  // property
	Selections(AIndex int32) ITreeNode                                                                      // property
	SeparatorColor() TColor                                                                                 // property
	SetSeparatorColor(AValue TColor)                                                                        // property
	StateImages() ICustomImageList                                                                          // property
	SetStateImages(AValue ICustomImageList)                                                                 // property
	StateImagesWidth() int32                                                                                // property
	SetStateImagesWidth(AValue int32)                                                                       // property
	TopItem() ITreeNode                                                                                     // property
	SetTopItem(AValue ITreeNode)                                                                            // property
	TreeLineColor() TColor                                                                                  // property
	SetTreeLineColor(AValue TColor)                                                                         // property
	TreeLinePenStyle() TPenStyle                                                                            // property
	SetTreeLinePenStyle(AValue TPenStyle)                                                                   // property
	AlphaSort() bool                                                                                        // function
	CustomSort(fn TTreeNodeCompare) bool                                                                    // function
	DefaultTreeViewSort(Node1, Node2 ITreeNode) int32                                                       // function
	GetHitTestInfoAt(X, Y int32) THitTests                                                                  // function
	GetNodeAt(X, Y int32) ITreeNode                                                                         // function
	GetNodeWithExpandSignAt(X, Y int32) ITreeNode                                                           // function
	IsEditing() bool                                                                                        // function
	GetFirstMultiSelected() ITreeNode                                                                       // function
	GetLastMultiSelected() ITreeNode                                                                        // function
	SelectionVisible() bool                                                                                 // function
	StoreCurrentSelection() IStringList                                                                     // function
	ClearSelection(KeepPrimary bool)                                                                        // procedure
	ConsistencyCheck()                                                                                      // procedure
	GetInsertMarkAt(X, Y int32, OutNInsertMarkNode *ITreeNode, OutNInsertMarkType *TTreeViewInsertMarkType) // procedure
	SetInsertMark(AnInsertMarkNode ITreeNode, AnInsertMarkType TTreeViewInsertMarkType)                     // procedure
	SetInsertMarkAt(X, Y int32)                                                                             // procedure
	BeginUpdate()                                                                                           // procedure
	EndUpdate()                                                                                             // procedure
	FullCollapse()                                                                                          // procedure
	FullExpand()                                                                                            // procedure
	LoadFromFile(FileName string)                                                                           // procedure
	LoadFromStream(Stream IStream)                                                                          // procedure
	SaveToFile(FileName string)                                                                             // procedure
	SaveToStream(Stream IStream)                                                                            // procedure
	WriteDebugReport(Prefix string, AllNodes bool)                                                          // procedure
	LockSelectionChangeEvent()                                                                              // procedure
	UnlockSelectionChangeEvent()                                                                            // procedure
	Select(Node ITreeNode, ShiftState TShiftState)                                                          // procedure
	Select1(Nodes IList)                                                                                    // procedure
	MakeSelectionVisible()                                                                                  // procedure
	ClearInvisibleSelection()                                                                               // procedure
	ApplyStoredSelection(ASelection IStringList, FreeList bool)                                             // procedure
	MoveToNextNode(ASelect bool)                                                                            // procedure
	MoveToPrevNode(ASelect bool)                                                                            // procedure
	MovePageDown(ASelect bool)                                                                              // procedure
	MovePageUp(ASelect bool)                                                                                // procedure
	MoveLeft(ASelect bool)                                                                                  // procedure
	MoveRight(ASelect bool)                                                                                 // procedure
	MoveExpand(ASelect bool)                                                                                // procedure
	MoveCollapse(ASelect bool)                                                                              // procedure
	MoveHome(ASelect bool)                                                                                  // procedure
	MoveEnd(ASelect bool)                                                                                   // procedure
}

// TCustomTreeView Parent: TCustomControl
type TCustomTreeView struct {
	TCustomControl
	customSortPtr uintptr
}

func NewCustomTreeView(AnOwner IComponent) ICustomTreeView {
	r1 := LCL().SysCallN(2412, GetObjectUintptr(AnOwner))
	return AsCustomTreeView(r1)
}

func (m *TCustomTreeView) AccessibilityOn() bool {
	r1 := LCL().SysCallN(2402, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTreeView) SetAccessibilityOn(AValue bool) {
	LCL().SysCallN(2402, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTreeView) BackgroundColor() TColor {
	r1 := LCL().SysCallN(2405, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetBackgroundColor(AValue TColor) {
	LCL().SysCallN(2405, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) BottomItem() ITreeNode {
	r1 := LCL().SysCallN(2407, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetBottomItem(AValue ITreeNode) {
	LCL().SysCallN(2407, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) DefaultItemHeight() int32 {
	r1 := LCL().SysCallN(2414, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetDefaultItemHeight(AValue int32) {
	LCL().SysCallN(2414, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) DropTarget() ITreeNode {
	r1 := LCL().SysCallN(2416, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetDropTarget(AValue ITreeNode) {
	LCL().SysCallN(2416, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignColor() TColor {
	r1 := LCL().SysCallN(2418, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetExpandSignColor(AValue TColor) {
	LCL().SysCallN(2418, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignSize() int32 {
	r1 := LCL().SysCallN(2419, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetExpandSignSize(AValue int32) {
	LCL().SysCallN(2419, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignWidth() int32 {
	r1 := LCL().SysCallN(2421, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetExpandSignWidth(AValue int32) {
	LCL().SysCallN(2421, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ExpandSignType() TTreeViewExpandSignType {
	r1 := LCL().SysCallN(2420, 0, m.Instance(), 0)
	return TTreeViewExpandSignType(r1)
}

func (m *TCustomTreeView) SetExpandSignType(AValue TTreeViewExpandSignType) {
	LCL().SysCallN(2420, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Images() ICustomImageList {
	r1 := LCL().SysCallN(2430, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTreeView) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(2430, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) ImagesWidth() int32 {
	r1 := LCL().SysCallN(2431, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetImagesWidth(AValue int32) {
	LCL().SysCallN(2431, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) InsertMarkNode() ITreeNode {
	r1 := LCL().SysCallN(2432, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetInsertMarkNode(AValue ITreeNode) {
	LCL().SysCallN(2432, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) InsertMarkType() TTreeViewInsertMarkType {
	r1 := LCL().SysCallN(2433, 0, m.Instance(), 0)
	return TTreeViewInsertMarkType(r1)
}

func (m *TCustomTreeView) SetInsertMarkType(AValue TTreeViewInsertMarkType) {
	LCL().SysCallN(2433, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Items() ITreeNodes {
	r1 := LCL().SysCallN(2435, 0, m.Instance(), 0)
	return AsTreeNodes(r1)
}

func (m *TCustomTreeView) SetItems(AValue ITreeNodes) {
	LCL().SysCallN(2435, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) KeepCollapsedNodes() bool {
	r1 := LCL().SysCallN(2436, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTreeView) SetKeepCollapsedNodes(AValue bool) {
	LCL().SysCallN(2436, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTreeView) MultiSelectStyle() TMultiSelectStyle {
	r1 := LCL().SysCallN(2451, 0, m.Instance(), 0)
	return TMultiSelectStyle(r1)
}

func (m *TCustomTreeView) SetMultiSelectStyle(AValue TMultiSelectStyle) {
	LCL().SysCallN(2451, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Options() TTreeViewOptions {
	r1 := LCL().SysCallN(2452, 0, m.Instance(), 0)
	return TTreeViewOptions(r1)
}

func (m *TCustomTreeView) SetOptions(AValue TTreeViewOptions) {
	LCL().SysCallN(2452, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) ScrollBars() TScrollStyle {
	r1 := LCL().SysCallN(2455, 0, m.Instance(), 0)
	return TScrollStyle(r1)
}

func (m *TCustomTreeView) SetScrollBars(AValue TScrollStyle) {
	LCL().SysCallN(2455, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) Selected() ITreeNode {
	r1 := LCL().SysCallN(2458, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetSelected(AValue ITreeNode) {
	LCL().SysCallN(2458, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) SelectionColor() TColor {
	r1 := LCL().SysCallN(2459, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetSelectionColor(AValue TColor) {
	LCL().SysCallN(2459, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) SelectionCount() uint32 {
	r1 := LCL().SysCallN(2460, m.Instance())
	return uint32(r1)
}

func (m *TCustomTreeView) SelectionFontColor() TColor {
	r1 := LCL().SysCallN(2461, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetSelectionFontColor(AValue TColor) {
	LCL().SysCallN(2461, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) SelectionFontColorUsed() bool {
	r1 := LCL().SysCallN(2462, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomTreeView) SetSelectionFontColorUsed(AValue bool) {
	LCL().SysCallN(2462, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomTreeView) Selections(AIndex int32) ITreeNode {
	r1 := LCL().SysCallN(2464, m.Instance(), uintptr(AIndex))
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SeparatorColor() TColor {
	r1 := LCL().SysCallN(2465, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetSeparatorColor(AValue TColor) {
	LCL().SysCallN(2465, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) StateImages() ICustomImageList {
	r1 := LCL().SysCallN(2468, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomTreeView) SetStateImages(AValue ICustomImageList) {
	LCL().SysCallN(2468, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) StateImagesWidth() int32 {
	r1 := LCL().SysCallN(2469, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomTreeView) SetStateImagesWidth(AValue int32) {
	LCL().SysCallN(2469, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) TopItem() ITreeNode {
	r1 := LCL().SysCallN(2471, 0, m.Instance(), 0)
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SetTopItem(AValue ITreeNode) {
	LCL().SysCallN(2471, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomTreeView) TreeLineColor() TColor {
	r1 := LCL().SysCallN(2472, 0, m.Instance(), 0)
	return TColor(r1)
}

func (m *TCustomTreeView) SetTreeLineColor(AValue TColor) {
	LCL().SysCallN(2472, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) TreeLinePenStyle() TPenStyle {
	r1 := LCL().SysCallN(2473, 0, m.Instance(), 0)
	return TPenStyle(r1)
}

func (m *TCustomTreeView) SetTreeLinePenStyle(AValue TPenStyle) {
	LCL().SysCallN(2473, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomTreeView) AlphaSort() bool {
	r1 := LCL().SysCallN(2403, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTreeView) CustomSort(fn TTreeNodeCompare) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := LCL().SysCallN(2413, m.Instance(), m.customSortPtr)
	return GoBool(r1)
}

func (m *TCustomTreeView) DefaultTreeViewSort(Node1, Node2 ITreeNode) int32 {
	r1 := LCL().SysCallN(2415, m.Instance(), GetObjectUintptr(Node1), GetObjectUintptr(Node2))
	return int32(r1)
}

func (m *TCustomTreeView) GetHitTestInfoAt(X, Y int32) THitTests {
	r1 := LCL().SysCallN(2425, m.Instance(), uintptr(X), uintptr(Y))
	return THitTests(r1)
}

func (m *TCustomTreeView) GetNodeAt(X, Y int32) ITreeNode {
	r1 := LCL().SysCallN(2428, m.Instance(), uintptr(X), uintptr(Y))
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) GetNodeWithExpandSignAt(X, Y int32) ITreeNode {
	r1 := LCL().SysCallN(2429, m.Instance(), uintptr(X), uintptr(Y))
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) IsEditing() bool {
	r1 := LCL().SysCallN(2434, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTreeView) GetFirstMultiSelected() ITreeNode {
	r1 := LCL().SysCallN(2424, m.Instance())
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) GetLastMultiSelected() ITreeNode {
	r1 := LCL().SysCallN(2427, m.Instance())
	return AsTreeNode(r1)
}

func (m *TCustomTreeView) SelectionVisible() bool {
	r1 := LCL().SysCallN(2463, m.Instance())
	return GoBool(r1)
}

func (m *TCustomTreeView) StoreCurrentSelection() IStringList {
	r1 := LCL().SysCallN(2470, m.Instance())
	return AsStringList(r1)
}

func CustomTreeViewClass() TClass {
	ret := LCL().SysCallN(2408)
	return TClass(ret)
}

func (m *TCustomTreeView) ClearSelection(KeepPrimary bool) {
	LCL().SysCallN(2410, m.Instance(), PascalBool(KeepPrimary))
}

func (m *TCustomTreeView) ConsistencyCheck() {
	LCL().SysCallN(2411, m.Instance())
}

func (m *TCustomTreeView) GetInsertMarkAt(X, Y int32, OutNInsertMarkNode *ITreeNode, OutNInsertMarkType *TTreeViewInsertMarkType) {
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(2426, m.Instance(), uintptr(X), uintptr(Y), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)))
	*OutNInsertMarkNode = AsTreeNode(result1)
	*OutNInsertMarkType = TTreeViewInsertMarkType(result2)
}

func (m *TCustomTreeView) SetInsertMark(AnInsertMarkNode ITreeNode, AnInsertMarkType TTreeViewInsertMarkType) {
	LCL().SysCallN(2466, m.Instance(), GetObjectUintptr(AnInsertMarkNode), uintptr(AnInsertMarkType))
}

func (m *TCustomTreeView) SetInsertMarkAt(X, Y int32) {
	LCL().SysCallN(2467, m.Instance(), uintptr(X), uintptr(Y))
}

func (m *TCustomTreeView) BeginUpdate() {
	LCL().SysCallN(2406, m.Instance())
}

func (m *TCustomTreeView) EndUpdate() {
	LCL().SysCallN(2417, m.Instance())
}

func (m *TCustomTreeView) FullCollapse() {
	LCL().SysCallN(2422, m.Instance())
}

func (m *TCustomTreeView) FullExpand() {
	LCL().SysCallN(2423, m.Instance())
}

func (m *TCustomTreeView) LoadFromFile(FileName string) {
	LCL().SysCallN(2437, m.Instance(), PascalStr(FileName))
}

func (m *TCustomTreeView) LoadFromStream(Stream IStream) {
	LCL().SysCallN(2438, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TCustomTreeView) SaveToFile(FileName string) {
	LCL().SysCallN(2453, m.Instance(), PascalStr(FileName))
}

func (m *TCustomTreeView) SaveToStream(Stream IStream) {
	LCL().SysCallN(2454, m.Instance(), GetObjectUintptr(Stream))
}

func (m *TCustomTreeView) WriteDebugReport(Prefix string, AllNodes bool) {
	LCL().SysCallN(2475, m.Instance(), PascalStr(Prefix), PascalBool(AllNodes))
}

func (m *TCustomTreeView) LockSelectionChangeEvent() {
	LCL().SysCallN(2439, m.Instance())
}

func (m *TCustomTreeView) UnlockSelectionChangeEvent() {
	LCL().SysCallN(2474, m.Instance())
}

func (m *TCustomTreeView) Select(Node ITreeNode, ShiftState TShiftState) {
	LCL().SysCallN(2456, m.Instance(), GetObjectUintptr(Node), uintptr(ShiftState))
}

func (m *TCustomTreeView) Select1(Nodes IList) {
	LCL().SysCallN(2457, m.Instance(), GetObjectUintptr(Nodes))
}

func (m *TCustomTreeView) MakeSelectionVisible() {
	LCL().SysCallN(2440, m.Instance())
}

func (m *TCustomTreeView) ClearInvisibleSelection() {
	LCL().SysCallN(2409, m.Instance())
}

func (m *TCustomTreeView) ApplyStoredSelection(ASelection IStringList, FreeList bool) {
	LCL().SysCallN(2404, m.Instance(), GetObjectUintptr(ASelection), PascalBool(FreeList))
}

func (m *TCustomTreeView) MoveToNextNode(ASelect bool) {
	LCL().SysCallN(2449, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveToPrevNode(ASelect bool) {
	LCL().SysCallN(2450, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MovePageDown(ASelect bool) {
	LCL().SysCallN(2446, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MovePageUp(ASelect bool) {
	LCL().SysCallN(2447, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveLeft(ASelect bool) {
	LCL().SysCallN(2445, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveRight(ASelect bool) {
	LCL().SysCallN(2448, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveExpand(ASelect bool) {
	LCL().SysCallN(2443, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveCollapse(ASelect bool) {
	LCL().SysCallN(2441, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveHome(ASelect bool) {
	LCL().SysCallN(2444, m.Instance(), PascalBool(ASelect))
}

func (m *TCustomTreeView) MoveEnd(ASelect bool) {
	LCL().SysCallN(2442, m.Instance(), PascalBool(ASelect))
}
