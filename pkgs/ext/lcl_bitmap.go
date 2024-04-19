//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package ext

import (
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
)

type TBitmap struct {
	*lcl.TBitmap
}

func NewBitmap() *TBitmap {
	m := &TBitmap{
		TBitmap: lcl.NewBitmap(),
	}
	return m
}

func (m *TBitmap) LoadFromIntfImage(intfImage *TLazIntfImage) {
	if !m.IsValid() || !intfImage.IsValid() {
		return
	}
	imports.LibLCLExt().Proc(Bitmap_LoadFromIntfImage).Call(m.Instance(), intfImage.Instance())
}
