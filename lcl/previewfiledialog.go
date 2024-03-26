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

// IPreviewFileDialog Parent: IOpenDialog
type IPreviewFileDialog interface {
	IOpenDialog
	PreviewFileControl() IPreviewFileControl // property
}

// TPreviewFileDialog Parent: TOpenDialog
type TPreviewFileDialog struct {
	TOpenDialog
}

func NewPreviewFileDialog(TheOwner IComponent) IPreviewFileDialog {
	r1 := LCL().SysCallN(3934, GetObjectUintptr(TheOwner))
	return AsPreviewFileDialog(r1)
}

func (m *TPreviewFileDialog) PreviewFileControl() IPreviewFileControl {
	r1 := LCL().SysCallN(3935, m.Instance())
	return AsPreviewFileControl(r1)
}

func PreviewFileDialogClass() TClass {
	ret := LCL().SysCallN(3933)
	return TClass(ret)
}
