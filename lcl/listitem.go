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

// IListItem Parent: IPersistent
type IListItem interface {
	IPersistent
	Caption() string                                                        // property
	SetCaption(AValue string)                                               // property
	Checked() bool                                                          // property
	SetChecked(AValue bool)                                                 // property
	Cut() bool                                                              // property
	SetCut(AValue bool)                                                     // property
	Data() uintptr                                                          // property
	SetData(AValue uintptr)                                                 // property
	DropTarget() bool                                                       // property
	SetDropTarget(AValue bool)                                              // property
	Focused() bool                                                          // property
	SetFocused(AValue bool)                                                 // property
	Index() int32                                                           // property
	ImageIndex() TImageIndex                                                // property
	SetImageIndex(AValue TImageIndex)                                       // property
	Left() int32                                                            // property
	SetLeft(AValue int32)                                                   // property
	ListView() ICustomListView                                              // property
	Owner() IListItems                                                      // property
	Position() (resultPoint TPoint)                                         // property
	SetPosition(AValue *TPoint)                                             // property
	Selected() bool                                                         // property
	SetSelected(AValue bool)                                                // property
	StateIndex() TImageIndex                                                // property
	SetStateIndex(AValue TImageIndex)                                       // property
	SubItems() IStrings                                                     // property
	SetSubItems(AValue IStrings)                                            // property
	SubItemImages(AIndex int32) int32                                       // property
	SetSubItemImages(AIndex int32, AValue int32)                            // property
	Top() int32                                                             // property
	SetTop(AValue int32)                                                    // property
	DisplayRect(Code TDisplayCode) (resultRect TRect)                       // function
	DisplayRectSubItem(subItem int32, Code TDisplayCode) (resultRect TRect) // function
	EditCaption() bool                                                      // function
	GetStates() TListItemStates                                             // function
	Delete()                                                                // procedure
	MakeVisible(PartialOK bool)                                             // procedure
}

// TListItem Parent: TPersistent
type TListItem struct {
	TPersistent
}

func NewListItem(AOwner IListItems) IListItem {
	r1 := LCL().SysCallN(3384, GetObjectUintptr(AOwner))
	return AsListItem(r1)
}

func (m *TListItem) Caption() string {
	r1 := LCL().SysCallN(3381, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TListItem) SetCaption(AValue string) {
	LCL().SysCallN(3381, 1, m.Instance(), PascalStr(AValue))
}

func (m *TListItem) Checked() bool {
	r1 := LCL().SysCallN(3382, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetChecked(AValue bool) {
	LCL().SysCallN(3382, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Cut() bool {
	r1 := LCL().SysCallN(3385, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetCut(AValue bool) {
	LCL().SysCallN(3385, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Data() uintptr {
	r1 := LCL().SysCallN(3386, 0, m.Instance(), 0)
	return uintptr(r1)
}

func (m *TListItem) SetData(AValue uintptr) {
	LCL().SysCallN(3386, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) DropTarget() bool {
	r1 := LCL().SysCallN(3390, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetDropTarget(AValue bool) {
	LCL().SysCallN(3390, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Focused() bool {
	r1 := LCL().SysCallN(3392, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetFocused(AValue bool) {
	LCL().SysCallN(3392, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) Index() int32 {
	r1 := LCL().SysCallN(3395, m.Instance())
	return int32(r1)
}

func (m *TListItem) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(3394, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListItem) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(3394, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) Left() int32 {
	r1 := LCL().SysCallN(3396, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItem) SetLeft(AValue int32) {
	LCL().SysCallN(3396, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) ListView() ICustomListView {
	r1 := LCL().SysCallN(3397, m.Instance())
	return AsCustomListView(r1)
}

func (m *TListItem) Owner() IListItems {
	r1 := LCL().SysCallN(3399, m.Instance())
	return AsListItems(r1)
}

func (m *TListItem) Position() (resultPoint TPoint) {
	LCL().SysCallN(3400, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TListItem) SetPosition(AValue *TPoint) {
	LCL().SysCallN(3400, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TListItem) Selected() bool {
	r1 := LCL().SysCallN(3401, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TListItem) SetSelected(AValue bool) {
	LCL().SysCallN(3401, 1, m.Instance(), PascalBool(AValue))
}

func (m *TListItem) StateIndex() TImageIndex {
	r1 := LCL().SysCallN(3402, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TListItem) SetStateIndex(AValue TImageIndex) {
	LCL().SysCallN(3402, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) SubItems() IStrings {
	r1 := LCL().SysCallN(3404, 0, m.Instance(), 0)
	return AsStrings(r1)
}

func (m *TListItem) SetSubItems(AValue IStrings) {
	LCL().SysCallN(3404, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TListItem) SubItemImages(AIndex int32) int32 {
	r1 := LCL().SysCallN(3403, 0, m.Instance(), uintptr(AIndex))
	return int32(r1)
}

func (m *TListItem) SetSubItemImages(AIndex int32, AValue int32) {
	LCL().SysCallN(3403, 1, m.Instance(), uintptr(AIndex), uintptr(AValue))
}

func (m *TListItem) Top() int32 {
	r1 := LCL().SysCallN(3405, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TListItem) SetTop(AValue int32) {
	LCL().SysCallN(3405, 1, m.Instance(), uintptr(AValue))
}

func (m *TListItem) DisplayRect(Code TDisplayCode) (resultRect TRect) {
	LCL().SysCallN(3388, m.Instance(), uintptr(Code), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TListItem) DisplayRectSubItem(subItem int32, Code TDisplayCode) (resultRect TRect) {
	LCL().SysCallN(3389, m.Instance(), uintptr(subItem), uintptr(Code), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TListItem) EditCaption() bool {
	r1 := LCL().SysCallN(3391, m.Instance())
	return GoBool(r1)
}

func (m *TListItem) GetStates() TListItemStates {
	r1 := LCL().SysCallN(3393, m.Instance())
	return TListItemStates(r1)
}

func ListItemClass() TClass {
	ret := LCL().SysCallN(3383)
	return TClass(ret)
}

func (m *TListItem) Delete() {
	LCL().SysCallN(3387, m.Instance())
}

func (m *TListItem) MakeVisible(PartialOK bool) {
	LCL().SysCallN(3398, m.Instance(), PascalBool(PartialOK))
}
