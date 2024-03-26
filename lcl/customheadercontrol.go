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

// ICustomHeaderControl Parent: ICustomControl
type ICustomHeaderControl interface {
	ICustomControl
	SectionFromOriginalIndex(OriginalIndex int32) IHeaderSection // property
	DragReorder() bool                                           // property
	SetDragReorder(AValue bool)                                  // property
	Images() ICustomImageList                                    // property
	SetImages(AValue ICustomImageList)                           // property
	ImagesWidth() int32                                          // property
	SetImagesWidth(AValue int32)                                 // property
	Sections() IHeaderSections                                   // property
	SetSections(AValue IHeaderSections)                          // property
	GetSectionAt(P *TPoint) int32                                // function
	Click()                                                      // procedure
	DblClick()                                                   // procedure
	Paint()                                                      // procedure
	PaintSection(Index int32)                                    // procedure
	ChangeScale(M, D int32)                                      // procedure
	SetOnSectionDrag(fn TSectionDragEvent)                       // property event
	SetOnSectionEndDrag(fn TNotifyEvent)                         // property event
	SetOnSectionClick(fn TCustomSectionNotifyEvent)              // property event
	SetOnSectionResize(fn TCustomSectionNotifyEvent)             // property event
	SetOnSectionTrack(fn TCustomSectionTrackEvent)               // property event
	SetOnSectionSeparatorDblClick(fn TCustomSectionNotifyEvent)  // property event
	SetOnCreateSectionClass(fn TCustomHCCreateSectionClassEvent) // property event
}

// TCustomHeaderControl Parent: TCustomControl
type TCustomHeaderControl struct {
	TCustomControl
	sectionDragPtr              uintptr
	sectionEndDragPtr           uintptr
	sectionClickPtr             uintptr
	sectionResizePtr            uintptr
	sectionTrackPtr             uintptr
	sectionSeparatorDblClickPtr uintptr
	createSectionClassPtr       uintptr
}

func NewCustomHeaderControl(AOwner IComponent) ICustomHeaderControl {
	r1 := LCL().SysCallN(1594, GetObjectUintptr(AOwner))
	return AsCustomHeaderControl(r1)
}

func (m *TCustomHeaderControl) SectionFromOriginalIndex(OriginalIndex int32) IHeaderSection {
	r1 := LCL().SysCallN(1602, m.Instance(), uintptr(OriginalIndex))
	return AsHeaderSection(r1)
}

func (m *TCustomHeaderControl) DragReorder() bool {
	r1 := LCL().SysCallN(1596, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomHeaderControl) SetDragReorder(AValue bool) {
	LCL().SysCallN(1596, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomHeaderControl) Images() ICustomImageList {
	r1 := LCL().SysCallN(1598, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomHeaderControl) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1598, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomHeaderControl) ImagesWidth() int32 {
	r1 := LCL().SysCallN(1599, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomHeaderControl) SetImagesWidth(AValue int32) {
	LCL().SysCallN(1599, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomHeaderControl) Sections() IHeaderSections {
	r1 := LCL().SysCallN(1603, 0, m.Instance(), 0)
	return AsHeaderSections(r1)
}

func (m *TCustomHeaderControl) SetSections(AValue IHeaderSections) {
	LCL().SysCallN(1603, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomHeaderControl) GetSectionAt(P *TPoint) int32 {
	r1 := LCL().SysCallN(1597, m.Instance(), uintptr(unsafe.Pointer(P)))
	return int32(r1)
}

func CustomHeaderControlClass() TClass {
	ret := LCL().SysCallN(1592)
	return TClass(ret)
}

func (m *TCustomHeaderControl) Click() {
	LCL().SysCallN(1593, m.Instance())
}

func (m *TCustomHeaderControl) DblClick() {
	LCL().SysCallN(1595, m.Instance())
}

func (m *TCustomHeaderControl) Paint() {
	LCL().SysCallN(1600, m.Instance())
}

func (m *TCustomHeaderControl) PaintSection(Index int32) {
	LCL().SysCallN(1601, m.Instance(), uintptr(Index))
}

func (m *TCustomHeaderControl) ChangeScale(M, D int32) {
	LCL().SysCallN(1591, m.Instance(), uintptr(M), uintptr(D))
}

func (m *TCustomHeaderControl) SetOnSectionDrag(fn TSectionDragEvent) {
	if m.sectionDragPtr != 0 {
		RemoveEventElement(m.sectionDragPtr)
	}
	m.sectionDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1606, m.Instance(), m.sectionDragPtr)
}

func (m *TCustomHeaderControl) SetOnSectionEndDrag(fn TNotifyEvent) {
	if m.sectionEndDragPtr != 0 {
		RemoveEventElement(m.sectionEndDragPtr)
	}
	m.sectionEndDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1607, m.Instance(), m.sectionEndDragPtr)
}

func (m *TCustomHeaderControl) SetOnSectionClick(fn TCustomSectionNotifyEvent) {
	if m.sectionClickPtr != 0 {
		RemoveEventElement(m.sectionClickPtr)
	}
	m.sectionClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1605, m.Instance(), m.sectionClickPtr)
}

func (m *TCustomHeaderControl) SetOnSectionResize(fn TCustomSectionNotifyEvent) {
	if m.sectionResizePtr != 0 {
		RemoveEventElement(m.sectionResizePtr)
	}
	m.sectionResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1608, m.Instance(), m.sectionResizePtr)
}

func (m *TCustomHeaderControl) SetOnSectionTrack(fn TCustomSectionTrackEvent) {
	if m.sectionTrackPtr != 0 {
		RemoveEventElement(m.sectionTrackPtr)
	}
	m.sectionTrackPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1610, m.Instance(), m.sectionTrackPtr)
}

func (m *TCustomHeaderControl) SetOnSectionSeparatorDblClick(fn TCustomSectionNotifyEvent) {
	if m.sectionSeparatorDblClickPtr != 0 {
		RemoveEventElement(m.sectionSeparatorDblClickPtr)
	}
	m.sectionSeparatorDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1609, m.Instance(), m.sectionSeparatorDblClickPtr)
}

func (m *TCustomHeaderControl) SetOnCreateSectionClass(fn TCustomHCCreateSectionClassEvent) {
	if m.createSectionClassPtr != 0 {
		RemoveEventElement(m.createSectionClassPtr)
	}
	m.createSectionClassPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1604, m.Instance(), m.createSectionClassPtr)
}
