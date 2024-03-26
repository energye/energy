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

// IPaperSize Parent: IObject
type IPaperSize interface {
	IObject
	DefaultPapers() bool                                   // property
	Width() int32                                          // property
	Height() int32                                         // property
	PaperName() string                                     // property
	SetPaperName(AValue string)                            // property
	DefaultPaperName() string                              // property
	PaperRect() (resultPaperRect TPaperRect)               // property
	SetPaperRect(AValue *TPaperRect)                       // property
	SupportedPapers() IStrings                             // property
	PaperRectOf(aName string) (resultPaperRect TPaperRect) // property
}

// TPaperSize Parent: TObject
type TPaperSize struct {
	TObject
}

func NewPaperSize(aOwner IPrinter) IPaperSize {
	r1 := LCL().SysCallN(3852, GetObjectUintptr(aOwner))
	return AsPaperSize(r1)
}

func (m *TPaperSize) DefaultPapers() bool {
	r1 := LCL().SysCallN(3854, m.Instance())
	return GoBool(r1)
}

func (m *TPaperSize) Width() int32 {
	r1 := LCL().SysCallN(3860, m.Instance())
	return int32(r1)
}

func (m *TPaperSize) Height() int32 {
	r1 := LCL().SysCallN(3855, m.Instance())
	return int32(r1)
}

func (m *TPaperSize) PaperName() string {
	r1 := LCL().SysCallN(3856, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TPaperSize) SetPaperName(AValue string) {
	LCL().SysCallN(3856, 1, m.Instance(), PascalStr(AValue))
}

func (m *TPaperSize) DefaultPaperName() string {
	r1 := LCL().SysCallN(3853, m.Instance())
	return GoStr(r1)
}

func (m *TPaperSize) PaperRect() (resultPaperRect TPaperRect) {
	r1 := LCL().SysCallN(3857, 0, m.Instance(), 0)
	return *(*TPaperRect)(getPointer(r1))
}

func (m *TPaperSize) SetPaperRect(AValue *TPaperRect) {
	LCL().SysCallN(3857, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)))
}

func (m *TPaperSize) SupportedPapers() IStrings {
	r1 := LCL().SysCallN(3859, m.Instance())
	return AsStrings(r1)
}

func (m *TPaperSize) PaperRectOf(aName string) (resultPaperRect TPaperRect) {
	r1 := LCL().SysCallN(3858, m.Instance(), PascalStr(aName))
	return *(*TPaperRect)(getPointer(r1))
}

func PaperSizeClass() TClass {
	ret := LCL().SysCallN(3851)
	return TClass(ret)
}
