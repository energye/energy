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

// IOpenPictureDialog Parent: IPreviewFileDialog
type IOpenPictureDialog interface {
	IPreviewFileDialog
	DefaultFilter() string // property
	GetFilterExt() string  // function
}

// TOpenPictureDialog Parent: TPreviewFileDialog
type TOpenPictureDialog struct {
	TPreviewFileDialog
}

func NewOpenPictureDialog(TheOwner IComponent) IOpenPictureDialog {
	r1 := LCL().SysCallN(3736, GetObjectUintptr(TheOwner))
	return AsOpenPictureDialog(r1)
}

func (m *TOpenPictureDialog) DefaultFilter() string {
	r1 := LCL().SysCallN(3737, m.Instance())
	return GoStr(r1)
}

func (m *TOpenPictureDialog) GetFilterExt() string {
	r1 := LCL().SysCallN(3738, m.Instance())
	return GoStr(r1)
}

func OpenPictureDialogClass() TClass {
	ret := LCL().SysCallN(3735)
	return TClass(ret)
}
