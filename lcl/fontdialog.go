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

// IFontDialog Parent: ICommonDialog
type IFontDialog interface {
	ICommonDialog
	Font() IFont                          // property
	SetFont(AValue IFont)                 // property
	MinFontSize() int32                   // property
	SetMinFontSize(AValue int32)          // property
	MaxFontSize() int32                   // property
	SetMaxFontSize(AValue int32)          // property
	Options() TFontDialogOptions          // property
	SetOptions(AValue TFontDialogOptions) // property
	PreviewText() string                  // property
	SetPreviewText(AValue string)         // property
	ApplyClicked()                        // procedure
	SetOnApplyClicked(fn TNotifyEvent)    // property event
}

// TFontDialog Parent: TCommonDialog
type TFontDialog struct {
	TCommonDialog
	applyClickedPtr uintptr
}

func NewFontDialog(AOwner IComponent) IFontDialog {
	r1 := LCL().SysCallN(2830, GetObjectUintptr(AOwner))
	return AsFontDialog(r1)
}

func (m *TFontDialog) Font() IFont {
	r1 := LCL().SysCallN(2831, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TFontDialog) SetFont(AValue IFont) {
	LCL().SysCallN(2831, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFontDialog) MinFontSize() int32 {
	r1 := LCL().SysCallN(2833, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFontDialog) SetMinFontSize(AValue int32) {
	LCL().SysCallN(2833, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) MaxFontSize() int32 {
	r1 := LCL().SysCallN(2832, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFontDialog) SetMaxFontSize(AValue int32) {
	LCL().SysCallN(2832, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) Options() TFontDialogOptions {
	r1 := LCL().SysCallN(2834, 0, m.Instance(), 0)
	return TFontDialogOptions(r1)
}

func (m *TFontDialog) SetOptions(AValue TFontDialogOptions) {
	LCL().SysCallN(2834, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) PreviewText() string {
	r1 := LCL().SysCallN(2835, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFontDialog) SetPreviewText(AValue string) {
	LCL().SysCallN(2835, 1, m.Instance(), PascalStr(AValue))
}

func FontDialogClass() TClass {
	ret := LCL().SysCallN(2829)
	return TClass(ret)
}

func (m *TFontDialog) ApplyClicked() {
	LCL().SysCallN(2828, m.Instance())
}

func (m *TFontDialog) SetOnApplyClicked(fn TNotifyEvent) {
	if m.applyClickedPtr != 0 {
		RemoveEventElement(m.applyClickedPtr)
	}
	m.applyClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2836, m.Instance(), m.applyClickedPtr)
}
