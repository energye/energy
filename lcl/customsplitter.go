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
	r1 := LCL().SysCallN(2062, GetObjectUintptr(TheOwner))
	return AsCustomSplitter(r1)
}

func (m *TCustomSplitter) ResizeControl() IControl {
	r1 := LCL().SysCallN(2068, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomSplitter) SetResizeControl(AValue IControl) {
	LCL().SysCallN(2068, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSplitter) AutoSnap() bool {
	r1 := LCL().SysCallN(2059, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetAutoSnap(AValue bool) {
	LCL().SysCallN(2059, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) Beveled() bool {
	r1 := LCL().SysCallN(2060, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetBeveled(AValue bool) {
	LCL().SysCallN(2060, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) MinSize() int32 {
	r1 := LCL().SysCallN(2065, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSplitter) SetMinSize(AValue int32) {
	LCL().SysCallN(2065, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeAnchor() TAnchorKind {
	r1 := LCL().SysCallN(2067, 0, m.Instance(), 0)
	return TAnchorKind(r1)
}

func (m *TCustomSplitter) SetResizeAnchor(AValue TAnchorKind) {
	LCL().SysCallN(2067, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeStyle() TResizeStyle {
	r1 := LCL().SysCallN(2069, 0, m.Instance(), 0)
	return TResizeStyle(r1)
}

func (m *TCustomSplitter) SetResizeStyle(AValue TResizeStyle) {
	LCL().SysCallN(2069, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) GetOtherResizeControl() IControl {
	r1 := LCL().SysCallN(2063, m.Instance())
	return AsControl(r1)
}

func (m *TCustomSplitter) GetSplitterPosition() int32 {
	r1 := LCL().SysCallN(2064, m.Instance())
	return int32(r1)
}

func CustomSplitterClass() TClass {
	ret := LCL().SysCallN(2061)
	return TClass(ret)
}

func (m *TCustomSplitter) AnchorSplitter(Kind TAnchorKind, AControl IControl) {
	LCL().SysCallN(2058, m.Instance(), uintptr(Kind), GetObjectUintptr(AControl))
}

func (m *TCustomSplitter) MoveSplitter(Offset int32) {
	LCL().SysCallN(2066, m.Instance(), uintptr(Offset))
}

func (m *TCustomSplitter) SetSplitterPosition(NewPosition int32) {
	LCL().SysCallN(2073, m.Instance(), uintptr(NewPosition))
}

func (m *TCustomSplitter) SetOnCanOffset(fn TCanOffsetEvent) {
	if m.canOffsetPtr != 0 {
		RemoveEventElement(m.canOffsetPtr)
	}
	m.canOffsetPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2070, m.Instance(), m.canOffsetPtr)
}

func (m *TCustomSplitter) SetOnCanResize(fn TCanResizeEvent) {
	if m.canResizePtr != 0 {
		RemoveEventElement(m.canResizePtr)
	}
	m.canResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2071, m.Instance(), m.canResizePtr)
}

func (m *TCustomSplitter) SetOnMoved(fn TNotifyEvent) {
	if m.movedPtr != 0 {
		RemoveEventElement(m.movedPtr)
	}
	m.movedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2072, m.Instance(), m.movedPtr)
}
