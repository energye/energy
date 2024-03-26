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
	r1 := LCL().SysCallN(1864, GetObjectUintptr(AOwner))
	return AsCustomListView(r1)
}

func (m *TCustomListView) BoundingRect() (resultRect TRect) {
	LCL().SysCallN(1856, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TCustomListView) BorderStyle() TBorderStyle {
	r1 := LCL().SysCallN(1855, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCustomListView) SetBorderStyle(AValue TBorderStyle) {
	LCL().SysCallN(1855, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) Canvas() ICanvas {
	r1 := LCL().SysCallN(1857, m.Instance())
	return AsCanvas(r1)
}

func (m *TCustomListView) Checkboxes() bool {
	r1 := LCL().SysCallN(1858, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetCheckboxes(AValue bool) {
	LCL().SysCallN(1858, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) Column(AIndex int32) IListColumn {
	r1 := LCL().SysCallN(1862, m.Instance(), uintptr(AIndex))
	return AsListColumn(r1)
}

func (m *TCustomListView) ColumnCount() int32 {
	r1 := LCL().SysCallN(1863, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) DropTarget() IListItem {
	r1 := LCL().SysCallN(1866, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetDropTarget(AValue IListItem) {
	LCL().SysCallN(1866, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) FlatScrollBars() bool {
	r1 := LCL().SysCallN(1870, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetFlatScrollBars(AValue bool) {
	LCL().SysCallN(1870, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) FullDrag() bool {
	r1 := LCL().SysCallN(1871, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetFullDrag(AValue bool) {
	LCL().SysCallN(1871, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) GridLines() bool {
	r1 := LCL().SysCallN(1876, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetGridLines(AValue bool) {
	LCL().SysCallN(1876, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) HotTrack() bool {
	r1 := LCL().SysCallN(1877, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetHotTrack(AValue bool) {
	LCL().SysCallN(1877, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) HotTrackStyles() TListHotTrackStyles {
	r1 := LCL().SysCallN(1878, 0, m.Instance(), 0)
	return TListHotTrackStyles(r1)
}

func (m *TCustomListView) SetHotTrackStyles(AValue TListHotTrackStyles) {
	LCL().SysCallN(1878, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) IconOptions() IIconOptions {
	r1 := LCL().SysCallN(1879, 0, m.Instance(), 0)
	return AsIconOptions(r1)
}

func (m *TCustomListView) SetIconOptions(AValue IIconOptions) {
	LCL().SysCallN(1879, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) ItemFocused() IListItem {
	r1 := LCL().SysCallN(1881, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetItemFocused(AValue IListItem) {
	LCL().SysCallN(1881, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) ItemIndex() int32 {
	r1 := LCL().SysCallN(1882, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomListView) SetItemIndex(AValue int32) {
	LCL().SysCallN(1882, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomListView) Items() IListItems {
	r1 := LCL().SysCallN(1883, 0, m.Instance(), 0)
	return AsListItems(r1)
}

func (m *TCustomListView) SetItems(AValue IListItems) {
	LCL().SysCallN(1883, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) MultiSelect() bool {
	r1 := LCL().SysCallN(1885, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetMultiSelect(AValue bool) {
	LCL().SysCallN(1885, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) OwnerData() bool {
	r1 := LCL().SysCallN(1886, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetOwnerData(AValue bool) {
	LCL().SysCallN(1886, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) ReadOnly() bool {
	r1 := LCL().SysCallN(1887, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetReadOnly(AValue bool) {
	LCL().SysCallN(1887, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) RowSelect() bool {
	r1 := LCL().SysCallN(1888, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomListView) SetRowSelect(AValue bool) {
	LCL().SysCallN(1888, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomListView) SelCount() int32 {
	r1 := LCL().SysCallN(1889, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) Selected() IListItem {
	r1 := LCL().SysCallN(1891, 0, m.Instance(), 0)
	return AsListItem(r1)
}

func (m *TCustomListView) SetSelected(AValue IListItem) {
	LCL().SysCallN(1891, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomListView) LastSelected() IListItem {
	r1 := LCL().SysCallN(1884, m.Instance())
	return AsListItem(r1)
}

func (m *TCustomListView) TopItem() IListItem {
	r1 := LCL().SysCallN(1893, m.Instance())
	return AsListItem(r1)
}

func (m *TCustomListView) ViewOrigin() (resultPoint TPoint) {
	LCL().SysCallN(1894, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TCustomListView) SetViewOrigin(AValue *TPoint) {
	LCL().SysCallN(1894, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TCustomListView) VisibleRowCount() int32 {
	r1 := LCL().SysCallN(1895, m.Instance())
	return int32(r1)
}

func (m *TCustomListView) AlphaSort() bool {
	r1 := LCL().SysCallN(1853, m.Instance())
	return GoBool(r1)
}

func (m *TCustomListView) CustomSort(fn TLVCompare, AOptionalParam uint32) bool {
	if m.customSortPtr != 0 {
		RemoveEventElement(m.customSortPtr)
	}
	m.customSortPtr = MakeEventDataPtr(fn)
	r1 := LCL().SysCallN(1865, m.Instance(), m.customSortPtr, uintptr(AOptionalParam))
	return GoBool(r1)
}

func (m *TCustomListView) FindCaption(StartIndex int32, Value string, Partial, Inclusive, Wrap bool, PartStart bool) IListItem {
	r1 := LCL().SysCallN(1868, m.Instance(), uintptr(StartIndex), PascalStr(Value), PascalBool(Partial), PascalBool(Inclusive), PascalBool(Wrap), PascalBool(PartStart))
	return AsListItem(r1)
}

func (m *TCustomListView) FindData(StartIndex int32, Value uintptr, Inclusive, Wrap bool) IListItem {
	r1 := LCL().SysCallN(1869, m.Instance(), uintptr(StartIndex), uintptr(Value), PascalBool(Inclusive), PascalBool(Wrap))
	return AsListItem(r1)
}

func (m *TCustomListView) GetHitTestInfoAt(X, Y int32) THitTests {
	r1 := LCL().SysCallN(1872, m.Instance(), uintptr(X), uintptr(Y))
	return THitTests(r1)
}

func (m *TCustomListView) GetItemAt(x, y int32) IListItem {
	r1 := LCL().SysCallN(1873, m.Instance(), uintptr(x), uintptr(y))
	return AsListItem(r1)
}

func (m *TCustomListView) GetNearestItem(APoint *TPoint, Direction TSearchDirection) IListItem {
	r1 := LCL().SysCallN(1874, m.Instance(), uintptr(unsafe.Pointer(APoint)), uintptr(Direction))
	return AsListItem(r1)
}

func (m *TCustomListView) GetNextItem(StartItem IListItem, Direction TSearchDirection, States TListItemStates) IListItem {
	r1 := LCL().SysCallN(1875, m.Instance(), GetObjectUintptr(StartItem), uintptr(Direction), uintptr(States))
	return AsListItem(r1)
}

func (m *TCustomListView) IsEditing() bool {
	r1 := LCL().SysCallN(1880, m.Instance())
	return GoBool(r1)
}

func CustomListViewClass() TClass {
	ret := LCL().SysCallN(1859)
	return TClass(ret)
}

func (m *TCustomListView) AddItem(Item string, AObject IObject) {
	LCL().SysCallN(1852, m.Instance(), PascalStr(Item), GetObjectUintptr(AObject))
}

func (m *TCustomListView) Sort() {
	LCL().SysCallN(1892, m.Instance())
}

func (m *TCustomListView) BeginUpdate() {
	LCL().SysCallN(1854, m.Instance())
}

func (m *TCustomListView) Clear() {
	LCL().SysCallN(1860, m.Instance())
}

func (m *TCustomListView) EndUpdate() {
	LCL().SysCallN(1867, m.Instance())
}

func (m *TCustomListView) ClearSelection() {
	LCL().SysCallN(1861, m.Instance())
}

func (m *TCustomListView) SelectAll() {
	LCL().SysCallN(1890, m.Instance())
}
