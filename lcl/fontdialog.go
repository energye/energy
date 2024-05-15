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
	r1 := LCL().SysCallN(3073, GetObjectUintptr(AOwner))
	return AsFontDialog(r1)
}

func (m *TFontDialog) Font() IFont {
	r1 := LCL().SysCallN(3074, 0, m.Instance(), 0)
	return AsFont(r1)
}

func (m *TFontDialog) SetFont(AValue IFont) {
	LCL().SysCallN(3074, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TFontDialog) MinFontSize() int32 {
	r1 := LCL().SysCallN(3076, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFontDialog) SetMinFontSize(AValue int32) {
	LCL().SysCallN(3076, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) MaxFontSize() int32 {
	r1 := LCL().SysCallN(3075, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TFontDialog) SetMaxFontSize(AValue int32) {
	LCL().SysCallN(3075, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) Options() TFontDialogOptions {
	r1 := LCL().SysCallN(3077, 0, m.Instance(), 0)
	return TFontDialogOptions(r1)
}

func (m *TFontDialog) SetOptions(AValue TFontDialogOptions) {
	LCL().SysCallN(3077, 1, m.Instance(), uintptr(AValue))
}

func (m *TFontDialog) PreviewText() string {
	r1 := LCL().SysCallN(3078, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFontDialog) SetPreviewText(AValue string) {
	LCL().SysCallN(3078, 1, m.Instance(), PascalStr(AValue))
}

func FontDialogClass() TClass {
	ret := LCL().SysCallN(3072)
	return TClass(ret)
}

func (m *TFontDialog) ApplyClicked() {
	LCL().SysCallN(3071, m.Instance())
}

func (m *TFontDialog) SetOnApplyClicked(fn TNotifyEvent) {
	if m.applyClickedPtr != 0 {
		RemoveEventElement(m.applyClickedPtr)
	}
	m.applyClickedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3079, m.Instance(), m.applyClickedPtr)
}
