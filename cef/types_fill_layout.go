//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"unsafe"
)

// ICefFillLayout
// include/capi/views/cef_fill_layout_capi.h (cef_fill_layout_t)
type ICefFillLayout struct {
	*ICefLayout
}

// FillLayoutRef -> ICefFillLayout
var FillLayoutRef fillLayout

type fillLayout uintptr

func (*fillLayout) UnWrap(data *ICefFillLayout) *ICefFillLayout {
	var result uintptr
	imports.Proc(def.FillLayoutRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return nil
}
