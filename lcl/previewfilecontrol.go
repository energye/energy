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

// IPreviewFileControl Parent: IWinControl
type IPreviewFileControl interface {
	IWinControl
	PreviewFileDialog() IPreviewFileDialog          // property
	SetPreviewFileDialog(AValue IPreviewFileDialog) // property
}

// TPreviewFileControl Parent: TWinControl
type TPreviewFileControl struct {
	TWinControl
}

func NewPreviewFileControl(TheOwner IComponent) IPreviewFileControl {
	r1 := LCL().SysCallN(3931, GetObjectUintptr(TheOwner))
	return AsPreviewFileControl(r1)
}

func (m *TPreviewFileControl) PreviewFileDialog() IPreviewFileDialog {
	r1 := LCL().SysCallN(3932, 0, m.Instance(), 0)
	return AsPreviewFileDialog(r1)
}

func (m *TPreviewFileControl) SetPreviewFileDialog(AValue IPreviewFileDialog) {
	LCL().SysCallN(3932, 1, m.Instance(), GetObjectUintptr(AValue))
}

func PreviewFileControlClass() TClass {
	ret := LCL().SysCallN(3930)
	return TClass(ret)
}
