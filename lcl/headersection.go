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

// IHeaderSection Parent: ICollectionItem
type IHeaderSection interface {
	ICollectionItem
	Left() int32                         // property
	Right() int32                        // property
	State() THeaderSectionState          // property
	SetState(AValue THeaderSectionState) // property
	Alignment() TAlignment               // property
	SetAlignment(AValue TAlignment)      // property
	ImageIndex() TImageIndex             // property
	SetImageIndex(AValue TImageIndex)    // property
	MaxWidth() int32                     // property
	SetMaxWidth(AValue int32)            // property
	MinWidth() int32                     // property
	SetMinWidth(AValue int32)            // property
	Text() string                        // property
	SetText(AValue string)               // property
	Width() int32                        // property
	SetWidth(AValue int32)               // property
	Visible() bool                       // property
	SetVisible(AValue bool)              // property
	OriginalIndex() int32                // property
}

// THeaderSection Parent: TCollectionItem
type THeaderSection struct {
	TCollectionItem
}

func NewHeaderSection(ACollection ICollection) IHeaderSection {
	r1 := LCL().SysCallN(3070, GetObjectUintptr(ACollection))
	return AsHeaderSection(r1)
}

func (m *THeaderSection) Left() int32 {
	r1 := LCL().SysCallN(3072, m.Instance())
	return int32(r1)
}

func (m *THeaderSection) Right() int32 {
	r1 := LCL().SysCallN(3076, m.Instance())
	return int32(r1)
}

func (m *THeaderSection) State() THeaderSectionState {
	r1 := LCL().SysCallN(3077, 0, m.Instance(), 0)
	return THeaderSectionState(r1)
}

func (m *THeaderSection) SetState(AValue THeaderSectionState) {
	LCL().SysCallN(3077, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) Alignment() TAlignment {
	r1 := LCL().SysCallN(3068, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *THeaderSection) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3068, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(3071, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *THeaderSection) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(3071, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) MaxWidth() int32 {
	r1 := LCL().SysCallN(3073, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *THeaderSection) SetMaxWidth(AValue int32) {
	LCL().SysCallN(3073, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) MinWidth() int32 {
	r1 := LCL().SysCallN(3074, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *THeaderSection) SetMinWidth(AValue int32) {
	LCL().SysCallN(3074, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) Text() string {
	r1 := LCL().SysCallN(3078, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *THeaderSection) SetText(AValue string) {
	LCL().SysCallN(3078, 1, m.Instance(), PascalStr(AValue))
}

func (m *THeaderSection) Width() int32 {
	r1 := LCL().SysCallN(3080, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *THeaderSection) SetWidth(AValue int32) {
	LCL().SysCallN(3080, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderSection) Visible() bool {
	r1 := LCL().SysCallN(3079, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *THeaderSection) SetVisible(AValue bool) {
	LCL().SysCallN(3079, 1, m.Instance(), PascalBool(AValue))
}

func (m *THeaderSection) OriginalIndex() int32 {
	r1 := LCL().SysCallN(3075, m.Instance())
	return int32(r1)
}

func HeaderSectionClass() TClass {
	ret := LCL().SysCallN(3069)
	return TClass(ret)
}
