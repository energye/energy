//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https//www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefScrollView
// include/capi/views/cef_scroll_view_capi.h (cef_scroll_view_t)
type ICefScrollView struct {
	*ICefView
}

// ScrollViewRef -> ICefScrollView
var ScrollViewRef scrollView

type scrollView uintptr

func (*scrollView) New(delegate *ICefViewDelegate) *ICefScrollView {
	var result uintptr
	imports.Proc(def.ScrollViewRef_NewCreateScrollView).Call(delegate.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefScrollView{&ICefView{
			instance: unsafe.Pointer(result),
		}}
	}
	return nil
}

func (m *ICefScrollView) SetContentView(view *ICefView) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ScrollView_SetContentView).Call(m.Instance(), view.Instance())
}

func (m *ICefScrollView) GetContentView() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ScrollView_GetContentView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{
			instance: unsafe.Pointer(result),
		}
	}
	return nil
}

func (m *ICefScrollView) GetVisibleContentRect() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ScrollView_GetVisibleContentRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefScrollView) HasHorizontalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ScrollView_HasHorizontalScrollbar).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefScrollView) GetHorizontalScrollbarHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ScrollView_GetHorizontalScrollbarHeight).Call(m.Instance())
	return int32(r1)
}

func (m *ICefScrollView) HasVerticalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ScrollView_HasVerticalScrollbar).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefScrollView) GetVerticalScrollbarWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ScrollView_GetVerticalScrollbarWidth).Call(m.Instance())
	return int32(r1)
}
