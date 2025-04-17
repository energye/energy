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

// ICefBoxLayout
// include/capi/views/cef_box_layout_capi.h (cef_box_layout_t)
type ICefBoxLayout struct {
	*ICefLayout
}

// BoxLayoutRef -> ICefBoxLayout
var BoxLayoutRef boxLayout

type boxLayout uintptr

func (*boxLayout) UnWrap(data *ICefBoxLayout) *ICefBoxLayout {
	var result uintptr
	imports.Proc(def.BoxLayoutRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return nil
}

// SetFlexForView
//
//	Set the flex weight for the given |view|. Using the preferred size as the
//	basis, free space along the main axis is distributed to views in the ratio
//	of their flex weights. Similarly, if the views will overflow the parent,
//	space is subtracted in these ratios. A flex of 0 means this view is not
//	resized. Flex values must not be negative.
func (m *ICefBoxLayout) SetFlexForView(view *ICefView, flex int32) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.BoxLayout_SetFlexForView).Call(m.Instance(), view.Instance(), uintptr(flex))
}

// ClearFlexForView
//
//	Clears the flex for the given |view|, causing it to use the default flex
//	specified via TCefBoxLayoutSettings.default_flex.
func (m *ICefBoxLayout) ClearFlexForView(view *ICefView) {
	if !m.IsValid() || !view.IsValid() {
		return
	}
	imports.Proc(def.BoxLayout_ClearFlexForView).Call(m.Instance(), view.Instance())
}
