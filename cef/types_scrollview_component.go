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
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCEFScrollViewComponent
type TCEFScrollViewComponent struct {
	*TCEFViewComponent
}

// ScrollViewComponentRef -> TCEFScrollViewComponent
var ScrollViewComponentRef scrollViewComponent

type scrollViewComponent uintptr

func (*scrollViewComponent) New(AOwner lcl.IComponent) *TCEFScrollViewComponent {
	var result uintptr
	imports.Proc(def.ScrollViewComponent_Create).Call(lcl.CheckPtr(AOwner), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &TCEFScrollViewComponent{&TCEFViewComponent{instance: getInstance(result)}}
	}
	return nil
}

func (m *TCEFScrollViewComponent) CreateScrollView() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ScrollViewComponent_CreateScrollView).Call(m.Instance())
}

func (m *TCEFScrollViewComponent) SetContentView(view *ICefView) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ScrollViewComponent_SetContentView).Call(m.Instance(), view.Instance())
}

func (m *TCEFScrollViewComponent) GetContentView() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ScrollViewComponent_GetContentView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{
			instance: unsafe.Pointer(result),
		}
	}
	return nil
}

func (m *TCEFScrollViewComponent) VisibleContentRect() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ScrollViewComponent_VisibleContentRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFScrollViewComponent) HasHorizontalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ScrollViewComponent_HasHorizontalScrollbar).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFScrollViewComponent) HorizontalScrollbarHeight() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ScrollViewComponent_HorizontalScrollbarHeight).Call(m.Instance())
	return int32(r1)
}

func (m *TCEFScrollViewComponent) HasVerticalScrollbar() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ScrollViewComponent_HasVerticalScrollbar).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TCEFScrollViewComponent) VerticalScrollbarWidth() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ScrollViewComponent_VerticalScrollbarWidth).Call(m.Instance())
	return int32(r1)
}
