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

// ICustomSplitter Parent: ICustomControl
type ICustomSplitter interface {
	ICustomControl
	ResizeControl() IControl                            // property
	SetResizeControl(AValue IControl)                   // property
	AutoSnap() bool                                     // property
	SetAutoSnap(AValue bool)                            // property
	Beveled() bool                                      // property
	SetBeveled(AValue bool)                             // property
	MinSize() int32                                     // property
	SetMinSize(AValue int32)                            // property
	ResizeAnchor() TAnchorKind                          // property
	SetResizeAnchor(AValue TAnchorKind)                 // property
	ResizeStyle() TResizeStyle                          // property
	SetResizeStyle(AValue TResizeStyle)                 // property
	GetOtherResizeControl() IControl                    // function
	GetSplitterPosition() int32                         // function
	AnchorSplitter(Kind TAnchorKind, AControl IControl) // procedure
	MoveSplitter(Offset int32)                          // procedure
	SetSplitterPosition(NewPosition int32)              // procedure
	SetOnCanOffset(fn TCanOffsetEvent)                  // property event
	SetOnCanResize(fn TCanResizeEvent)                  // property event
	SetOnMoved(fn TNotifyEvent)                         // property event
}

// TCustomSplitter Parent: TCustomControl
type TCustomSplitter struct {
	TCustomControl
	canOffsetPtr uintptr
	canResizePtr uintptr
	movedPtr     uintptr
}

func NewCustomSplitter(TheOwner IComponent) ICustomSplitter {
	r1 := LCL().SysCallN(2252, GetObjectUintptr(TheOwner))
	return AsCustomSplitter(r1)
}

func (m *TCustomSplitter) ResizeControl() IControl {
	r1 := LCL().SysCallN(2258, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomSplitter) SetResizeControl(AValue IControl) {
	LCL().SysCallN(2258, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSplitter) AutoSnap() bool {
	r1 := LCL().SysCallN(2249, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetAutoSnap(AValue bool) {
	LCL().SysCallN(2249, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) Beveled() bool {
	r1 := LCL().SysCallN(2250, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetBeveled(AValue bool) {
	LCL().SysCallN(2250, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) MinSize() int32 {
	r1 := LCL().SysCallN(2255, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSplitter) SetMinSize(AValue int32) {
	LCL().SysCallN(2255, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeAnchor() TAnchorKind {
	r1 := LCL().SysCallN(2257, 0, m.Instance(), 0)
	return TAnchorKind(r1)
}

func (m *TCustomSplitter) SetResizeAnchor(AValue TAnchorKind) {
	LCL().SysCallN(2257, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeStyle() TResizeStyle {
	r1 := LCL().SysCallN(2259, 0, m.Instance(), 0)
	return TResizeStyle(r1)
}

func (m *TCustomSplitter) SetResizeStyle(AValue TResizeStyle) {
	LCL().SysCallN(2259, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) GetOtherResizeControl() IControl {
	r1 := LCL().SysCallN(2253, m.Instance())
	return AsControl(r1)
}

func (m *TCustomSplitter) GetSplitterPosition() int32 {
	r1 := LCL().SysCallN(2254, m.Instance())
	return int32(r1)
}

func CustomSplitterClass() TClass {
	ret := LCL().SysCallN(2251)
	return TClass(ret)
}

func (m *TCustomSplitter) AnchorSplitter(Kind TAnchorKind, AControl IControl) {
	LCL().SysCallN(2248, m.Instance(), uintptr(Kind), GetObjectUintptr(AControl))
}

func (m *TCustomSplitter) MoveSplitter(Offset int32) {
	LCL().SysCallN(2256, m.Instance(), uintptr(Offset))
}

func (m *TCustomSplitter) SetSplitterPosition(NewPosition int32) {
	LCL().SysCallN(2263, m.Instance(), uintptr(NewPosition))
}

func (m *TCustomSplitter) SetOnCanOffset(fn TCanOffsetEvent) {
	if m.canOffsetPtr != 0 {
		RemoveEventElement(m.canOffsetPtr)
	}
	m.canOffsetPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2260, m.Instance(), m.canOffsetPtr)
}

func (m *TCustomSplitter) SetOnCanResize(fn TCanResizeEvent) {
	if m.canResizePtr != 0 {
		RemoveEventElement(m.canResizePtr)
	}
	m.canResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2261, m.Instance(), m.canResizePtr)
}

func (m *TCustomSplitter) SetOnMoved(fn TNotifyEvent) {
	if m.movedPtr != 0 {
		RemoveEventElement(m.movedPtr)
	}
	m.movedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2262, m.Instance(), m.movedPtr)
}
