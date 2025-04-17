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
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefLayout
// include/capi/views/cef_layout_capi.h (cef_layout_t)
type ICefLayout struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// LayoutRef -> ICefLayout
var LayoutRef layout

type layout uintptr

func (*layout) UnWrap(data *ICefLayout) *ICefLayout {
	var result uintptr
	imports.Proc(def.LayoutRef_UnWrap).Call(data.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	data.base.Free(data.Instance())
	data.instance = getInstance(result)
	return nil
}

func (m *ICefLayout) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefLayout) AsBoxLayout() *ICefBoxLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.Layout_AsBoxLayout).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	return &ICefBoxLayout{
		&ICefLayout{
			instance: getInstance(result),
		},
	}
}

func (m *ICefLayout) AsFillLayout() *ICefFillLayout {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.Layout_AsFillLayout).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result == 0 {
		return nil
	}
	return &ICefFillLayout{
		&ICefLayout{
			instance: getInstance(result),
		},
	}
}

func (m *ICefLayout) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.Layout_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}
