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

// IAnchorSide Parent: IPersistent
type IAnchorSide interface {
	IPersistent
	Owner() IControl                                                                                                                                                     // property
	Kind() TAnchorKind                                                                                                                                                   // property
	Control() IControl                                                                                                                                                   // property
	SetControl(AValue IControl)                                                                                                                                          // property
	Side() TAnchorSideReference                                                                                                                                          // property
	SetSide(AValue TAnchorSideReference)                                                                                                                                 // property
	CheckSidePosition(NewControl IControl, NewSide TAnchorSideReference, OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32) bool // function
	IsAnchoredToParent(ParentSide TAnchorKind) bool                                                                                                                      // function
	GetSidePosition(OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32)                                                           // procedure
	FixCenterAnchoring()                                                                                                                                                 // procedure
}

// TAnchorSide Parent: TPersistent
type TAnchorSide struct {
	TPersistent
}

func NewAnchorSide(TheOwner IControl, TheKind TAnchorKind) IAnchorSide {
	r1 := LCL().SysCallN(90, GetObjectUintptr(TheOwner), uintptr(TheKind))
	return AsAnchorSide(r1)
}

func (m *TAnchorSide) Owner() IControl {
	r1 := LCL().SysCallN(95, m.Instance())
	return AsControl(r1)
}

func (m *TAnchorSide) Kind() TAnchorKind {
	r1 := LCL().SysCallN(94, m.Instance())
	return TAnchorKind(r1)
}

func (m *TAnchorSide) Control() IControl {
	r1 := LCL().SysCallN(89, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TAnchorSide) SetControl(AValue IControl) {
	LCL().SysCallN(89, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TAnchorSide) Side() TAnchorSideReference {
	r1 := LCL().SysCallN(96, 0, m.Instance(), 0)
	return TAnchorSideReference(r1)
}

func (m *TAnchorSide) SetSide(AValue TAnchorSideReference) {
	LCL().SysCallN(96, 1, m.Instance(), uintptr(AValue))
}

func (m *TAnchorSide) CheckSidePosition(NewControl IControl, NewSide TAnchorSideReference, OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32) bool {
	var result2 uintptr
	var result3 uintptr
	var result4 uintptr
	r1 := LCL().SysCallN(87, m.Instance(), GetObjectUintptr(NewControl), uintptr(NewSide), uintptr(unsafe.Pointer(&result2)), uintptr(unsafe.Pointer(&result3)), uintptr(unsafe.Pointer(&result4)))
	*OutReferenceControl = AsControl(result2)
	*OutReferenceSide = TAnchorSideReference(result3)
	*OutPosition = int32(result4)
	return GoBool(r1)
}

func (m *TAnchorSide) IsAnchoredToParent(ParentSide TAnchorKind) bool {
	r1 := LCL().SysCallN(93, m.Instance(), uintptr(ParentSide))
	return GoBool(r1)
}

func AnchorSideClass() TClass {
	ret := LCL().SysCallN(88)
	return TClass(ret)
}

func (m *TAnchorSide) GetSidePosition(OutReferenceControl *IControl, OutReferenceSide *TAnchorSideReference, OutPosition *int32) {
	var result0 uintptr
	var result1 uintptr
	var result2 uintptr
	LCL().SysCallN(92, m.Instance(), uintptr(unsafe.Pointer(&result0)), uintptr(unsafe.Pointer(&result1)), uintptr(unsafe.Pointer(&result2)))
	*OutReferenceControl = AsControl(result0)
	*OutReferenceSide = TAnchorSideReference(result1)
	*OutPosition = int32(result2)
}

func (m *TAnchorSide) FixCenterAnchoring() {
	LCL().SysCallN(91, m.Instance())
}
