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

// IRichMemoInline Parent: IObject
type IRichMemoInline interface {
	IObject
	Owner() ICustomRichMemo            // property
	Draw(Canvas ICanvas, ASize *TSize) // procedure
	SetVisible(AVisible bool)          // procedure
	Invalidate()                       // procedure
}

// TRichMemoInline Parent: TObject
type TRichMemoInline struct {
	TObject
}

func NewRichMemoInline() IRichMemoInline {
	r1 := LCL().SysCallN(4174)
	return AsRichMemoInline(r1)
}

func (m *TRichMemoInline) Owner() ICustomRichMemo {
	r1 := LCL().SysCallN(4177, m.Instance())
	return AsCustomRichMemo(r1)
}

func RichMemoInlineClass() TClass {
	ret := LCL().SysCallN(4173)
	return TClass(ret)
}

func (m *TRichMemoInline) Draw(Canvas ICanvas, ASize *TSize) {
	LCL().SysCallN(4175, m.Instance(), GetObjectUintptr(Canvas), uintptr(unsafe.Pointer(ASize)))
}

func (m *TRichMemoInline) SetVisible(AVisible bool) {
	LCL().SysCallN(4178, m.Instance(), PascalBool(AVisible))
}

func (m *TRichMemoInline) Invalidate() {
	LCL().SysCallN(4176, m.Instance())
}
