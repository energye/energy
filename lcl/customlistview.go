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

// ICustomListView Parent: IWinControl
type ICustomListView interface {
	IWinControl
	BoundingRect() (resultRect TRect)                                                                    // property
	BorderStyle() TBorderStyle                                                                           // property
	SetBorderStyle(AValue TBorderStyle)                                                                  // property
	Canvas() ICanvas                                                                                     // property
	Checkboxes() bool                                                                                    // property
	SetCheckboxes(AValue bool)                                                                           // property
	Column(AIndex int32) IListColumn                                                                     // property
	ColumnCount() int32                                                                                  // property
	DropTarget() IListItem                                                                               // property
	SetDropTarget(AValue IListItem)                                                                      // property
	FlatScrollBars() bool                                                                                // property
	SetFlatScrollBars(AValue bool)                                                                       // property
	FullDrag() bool                                                                                      // property
	SetFullDrag(AValue bool)                                                                             // property
	GridLines() bool                                                                                     // property
	SetGridLines(AValue bool)                                                                            // property
	HotTrack() bool                                                                                      // property
	SetHotTrack(AValue bool)                                                                             // property
	HotTrackStyles() TListHotTrackStyles                                                                 // property
	SetHotTrackStyles(AValue TListHotTrackStyles)                                                        // property
	IconOptions() IIconOptions                                                                           // property
	SetIconOptions(AValue IIconOptions)                                                                  // property
	ItemFocused() IListItem                                                                              // property
	SetItemFocused(AValue IListItem)                                                                     // property
	ItemIndex() int32                                                                                    // property
	SetItemIndex(AValue int32)                                                                           // property
	Items() IListItems                                                                                   // property
	SetItems(AValue IListItems)                                                                          // property
	MultiSelect() bool                                                                                   // property
	SetMultiSelect(AValue bool)                                                                          // property
	OwnerData() bool                                                                                     // property
	SetOwnerData(AValue bool)                                                                            // property
	ReadOnly() bool                                                                                      // property
	SetReadOnly(AValue bool)                                                                             // property
	RowSelect() bool                                                                                     // property
	SetRowSelect(AValue bool)                                                                            // property
	SelCount() int32                                                                                     // property
	Selected() IListItem                                                                                 // property
	SetSelected(AValue IListItem)                                                                        // property
	LastSelected() IListItem                                                                             // property
	TopItem() IListItem                                                                                  // property
	ViewOrigin() (resultPoint TPoint)                                                                    // property
	SetViewOrigin(AValue *TPoint)                                                                        // property
	VisibleRowCount() int32                                                                              // property
	AlphaSort() bool                                                                                     // function
	CustomSort(fn TLVCompare, AOptionalParam uint32) bool                                                // function
	FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem // function
	FindData(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem                            // function
	GetHitTestInfoAt(X, Y int32) THitTests                                                               // function
	GetItemAt(x, y int32) IListItem                                                                      // function
	GetNearestItem(APoint *TPoint, Direction TSearchDirection) IListItem                                 // function
	GetNextItem(StartItem IListItem, Direction TSearchDirection, States TListItemStates) IListItem       // function
	IsEditing() bool                                                                                     // function
	AddItem(Item string, AObject IObject)                                                                // procedure
	Sort()                                                                                               // procedure
	BeginUpdate()                                                                                        // procedure
	Clear()                                                                                              // procedure
	EndUpdate()                                                                                          // procedure
	ClearSelection()                                                                                     // procedure
	SelectAll()                                                                                          // procedure
}

// TCustomListView Parent: TWinControl
type TCustomListView struct {
	TWinControl
	customSortPtr uintptr
}

func NewCustomListView(AOwner IComponent) ICustomListView {
	r1 := LCL().SysCallN(2054, GetObjectUintptr(AOwner))
	return AsCustomListView(r1)
}

func (m *TCustomListView) BoundingRect() (resultRect TRect) {
	LCL().SysCallN(2046, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TCustomListView) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(2045, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomListView) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(2045, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) Canvas() ICanvas {
	r1 := LCL().SysCallN(2047, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomListView) Checkboxes() bool {
	r1 := LCL().SysCallN(2048, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetCheckboxes(AValue bool) {
	LCL().SysCallN(2048, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) Column(AIndex int32) IListColumn {
	r1 := LCL().SysCallN(2052, m.Instance(), uintptr(AIndex))
	return AsListColumn(r1)
}

func (m *TCustomListView) ColumnCount() int32 {
	r1 := LCL().SysCallN(2053, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) DropTarget() IListItem {
	r1 := LCL().SysCallN(2056, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetDropTarget(AValue IListItem) {
	LCL().SysCallN(2056, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) FlatScrollBars() bool {
	r1 := LCL().SysCallN(2060, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetFlatScrollBars(AValue bool) {
	LCL().SysCallN(2060, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) FullDrag() bool {
	r1 := LCL().SysCallN(2061, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetFullDrag(AValue bool) {
	LCL().SysCallN(2061, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) GridLines() bool {
	r1 := LCL().SysCallN(2066, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetGridLines(AValue bool) {
	LCL().SysCallN(2066, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) HotTrack() bool {
	r1 := LCL().SysCallN(2067, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetHotTrack(AValue bool) {
	LCL().SysCallN(2067, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) HotTrackStyles() TListHotTrackStyles {
	r1 := LCL().SysCallN(2068, 0, m.Instance(), 0)
	return TListHotTrackStyles(r1)
}

func (m *TCustomListView) SetHotTrackStyles(AValue TListHotTrackStyles) {
	LCL().SysCallN(2068, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) IconOptions() IIconOptions {
	r1 := LCL().SysCallN(2069, 0, m.Instance(), 0)
	return AsIconOptions(r1)
}

func (m *TCustomListView) SetIconOptions(AValue IIconOptions) {
	LCL().SysCallN(2069, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) ItemFocused() IListItem {
	r1 := LCL().SysCallN(2071, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetItemFocused(AValue IListItem) {
	LCL().SysCallN(2071, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) ItemIndex() int32 {
	r1 := LCL().SysCallN(2072, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListView) SetItemIndex(AValue int32) {
	LCL().SysCallN(2072, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) Items() IListItems {
	r1 := LCL().SysCallN(2073, 0, m.Instance(), 0)
	return AsListItems(r1)
}

func (m *TCustomListView) SetItems(AValue IListItems) {
	LCL().SysCallN(2073, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) MultiSelect() bool {
	r1 := LCL().SysCallN(2075, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetMultiSelect(AValue bool) {
	LCL().SysCallN(2075, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) OwnerData() bool {
	r1 := LCL().SysCallN(2076, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetOwnerData(AValue bool) {
	LCL().SysCallN(2076, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) ReadOnly() bool {
	r1 := LCL().SysCallN(2077, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetReadOnly(AValue bool) {
	LCL().SysCallN(2077, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) RowSelect() bool {
	r1 := LCL().SysCallN(2078, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetRowSelect(AValue bool) {
	LCL().SysCallN(2078, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) SelCount() int32 {
	r1 := LCL().SysCallN(2079, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) Selected() IListItem {
	r1 := LCL().SysCallN(2081, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetSelected(AValue IListItem) {
	LCL().SysCallN(2081, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) LastSelected() IListItem {
	r1 := LCL().SysCallN(2074, m.Instance())
	return AsListItem(r1)
}

func (m *TCustomListView) TopItem() IListItem {
	r1 := LCL().SysCallN(2083, m.Instance())
	return AsListItem(r1)
}

func (m *TCustomListView) ViewOrigin() (resultPoint TPoint) {
	LCL().SysCallN(2084, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TCustomListView) SetViewOrigin(AValue *TPoint) {
	LCL().SysCallN(2084, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TCustomListView) VisibleRowCount() int32 {
	r1 := LCL().SysCallN(2085, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) AlphaSort() bool {
	r1 := LCL().SysCallN(2043, m.Instance())
	return GoBool(r1)
}

func (m *TCustomListView) CustomSort(fn TLVCompare, AOptionalParam uint32) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := LCL().SysCallN(2055, m.Instance(), m.customSortPtr, uintptr(AOptionalParam))
	return GoBool(r1)
}

func (m *TCustomListView) FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem {
	r1 := LCL().SysCallN(2058, m.Instance(), uintptr(StartIndex), PascalStr(Value), PascalBool(Partial), PascalBool(Inclusive), PascalBool(Wrap), PascalBool(PartStart))
	return AsListItem(r1)
}

func (m *TCustomListView) FindData(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem {
	r1 := LCL().SysCallN(2059, m.Instance(), uintptr(StartIndex), uintptr(Value), PascalBool(Inclusive), PascalBool(Wrap))
	return AsListItem(r1)
}

func (m *TCustomListView) GetHitTestInfoAt(X, Y int32) THitTests {
	r1 := LCL().SysCallN(2062, m.Instance(), uintptr(X), uintptr(Y))
	return THitTests(r1)
}

func (m *TCustomListView) GetItemAt(x, y int32) IListItem {
	r1 := LCL().SysCallN(2063, m.Instance(), uintptr(x), uintptr(y))
	return AsListItem(r1)
}

func (m *TCustomListView) GetNearestItem(APoint *TPoint, Direction TSearchDirection) IListItem {
	r1 := LCL().SysCallN(2064, m.Instance(), uintptr(unsafePointer(APoint)), uintptr(Direction))
	return AsListItem(r1)
}

func (m *TCustomListView) GetNextItem(StartItem IListItem, Direction TSearchDirection, States TListItemStates) IListItem {
	r1 := LCL().SysCallN(2065, m.Instance(), GetObjectUintptr(StartItem), uintptr(Direction), uintptr(States))
	return AsListItem(r1)
}

func (m *TCustomListView) IsEditing() bool {
	r1 := LCL().SysCallN(2070, m.Instance())
	return GoBool(r1)
}

func CustomListViewClass() TClass {
	ret := LCL().SysCallN(2049)
	return TClass(ret)
}

func (m *TCustomListView) AddItem(Item string, AObject IObject) {
	LCL().SysCallN(2042, m.Instance(), PascalStr(Item), GetObjectUintptr(AObject))
}

func (m *TCustomListView) Sort() {
	LCL().SysCallN(2082, m.Instance())
}

func (m *TCustomListView) BeginUpdate() {
	LCL().SysCallN(2044, m.Instance())
}

func (m *TCustomListView) Clear() {
	LCL().SysCallN(2050, m.Instance())
}

func (m *TCustomListView) EndUpdate() {
	LCL().SysCallN(2057, m.Instance())
}

func (m *TCustomListView) ClearSelection() {
	LCL().SysCallN(2051, m.Instance())
}

func (m *TCustomListView) SelectAll() {
	LCL().SysCallN(2080, m.Instance())
}
