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

// IRichEdit Parent: IRichMemo
type IRichEdit interface {
	IRichMemo
	Zoom() int32                                                                             // property
	SetZoom(AValue int32)                                                                    // property
	HideScrollBars() bool                                                                    // property
	SetHideScrollBars(AValue bool)                                                           // property
	PlainText() bool                                                                         // property
	SetPlainText(AValue bool)                                                                // property
	DefAttributes() ITextAttributes                                                          // property
	SetDefAttributes(AValue ITextAttributes)                                                 // property
	SelAttributes() ITextAttributes                                                          // property
	SetSelAttributes(AValue ITextAttributes)                                                 // property
	Paragraph() IParaAttributes                                                              // property
	SetParagraph(AValue IParaAttributes)                                                     // property
	FindText(ASearchStr string, AStartPos int32, ALength int32, AOptions TSearchTypes) int32 // function
}

// TRichEdit Parent: TRichMemo
type TRichEdit struct {
	TRichMemo
}

func NewRichEdit(AOnwer IComponent) IRichEdit {
	r1 := LCL().SysCallN(4807, GetObjectUintptr(AOnwer))
	return AsRichEdit(r1)
}

func (m *TRichEdit) Zoom() int32 {
	r1 := LCL().SysCallN(4814, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TRichEdit) SetZoom(AValue int32) {
	LCL().SysCallN(4814, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichEdit) HideScrollBars() bool {
	r1 := LCL().SysCallN(4810, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichEdit) SetHideScrollBars(AValue bool) {
	LCL().SysCallN(4810, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichEdit) PlainText() bool {
	r1 := LCL().SysCallN(4812, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichEdit) SetPlainText(AValue bool) {
	LCL().SysCallN(4812, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichEdit) DefAttributes() ITextAttributes {
	r1 := LCL().SysCallN(4808, 0, m.Instance(), 0)
	return AsTextAttributes(r1)
}

func (m *TRichEdit) SetDefAttributes(AValue ITextAttributes) {
	LCL().SysCallN(4808, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TRichEdit) SelAttributes() ITextAttributes {
	r1 := LCL().SysCallN(4813, 0, m.Instance(), 0)
	return AsTextAttributes(r1)
}

func (m *TRichEdit) SetSelAttributes(AValue ITextAttributes) {
	LCL().SysCallN(4813, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TRichEdit) Paragraph() IParaAttributes {
	r1 := LCL().SysCallN(4811, 0, m.Instance(), 0)
	return AsParaAttributes(r1)
}

func (m *TRichEdit) SetParagraph(AValue IParaAttributes) {
	LCL().SysCallN(4811, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TRichEdit) FindText(ASearchStr string, AStartPos int32, ALength int32, AOptions TSearchTypes) int32 {
	r1 := LCL().SysCallN(4809, m.Instance(), PascalStr(ASearchStr), uintptr(AStartPos), uintptr(ALength), uintptr(AOptions))
	return int32(r1)
}

func RichEditClass() TClass {
	ret := LCL().SysCallN(4806)
	return TClass(ret)
}
