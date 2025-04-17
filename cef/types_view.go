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
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

/*
*********************************
************* Views *************
*********************************

(*) Has CEF creation function
(d) Has delegate

----------------          ----------------------
| TCefView (d) | -------> | TCefTextfield (*d) |
----------------    |     ----------------------
					|
					|     ----------------------
					|---> | TCefScrollView (*) |
					|     ----------------------
					|
					|     ------------------          -------------------
					|---> | TCefPanel (*d) | -------> | TCefWindow (*d) |
					|     ------------------          -------------------
					|
					|     ------------------------
					|---> | TCefBrowserView (*d) |
					|     ------------------------
					|
					|     ------------------          -----------------------          -----------------------
					|---> | TCefButton (d) | -------> | TCefLabelButton (*) | -------> | TCefMenuButton (*d) |
						  ------------------          -----------------------          -----------------------


--------------          -----------------
| TCefLayout | -------> | TCefBoxLayout |
--------------    |     -----------------
				  |
				  |     ------------------
				  |---> | TCefFillLayout |
				  		------------------
*/

// ICefView
// include/capi/views/cef_view_capi.h (cef_view_t)
type ICefView struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// Instance 实例
func (m *ICefView) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefView) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefView) AsView() *ICefView {
	return m
}

func (m *ICefView) AsBrowserView() *ICefBrowserView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_AsBrowserView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefBrowserView{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefView) AsButton() *ICefButton {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_AsButton).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefButton{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefView) AsPanel() *ICefPanel {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_AsPanel).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefPanel{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefView) AsScrollView() *ICefScrollView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_AsScrollView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefScrollView{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefView) AsTextfield() *ICefTextfield {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_AsTextfield).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefTextfield{&ICefView{instance: getInstance(result)}}
	}
	return nil
}

func (m *ICefView) GetTypeString() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.View_GetTypeString).Call(m.Instance())
	return api.GoStr(r1)
}

func (m *ICefView) ToStringEx(includeChildren bool) string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.View_ToStringEx).Call(m.Instance(), api.PascalBool(includeChildren))
	return api.GoStr(r1)
}

func (m *ICefView) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsValid).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) IsAttached() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsAttached).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) IsSame(that *ICefView) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsSame).Call(m.Instance(), that.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) GetDelegate() *ICefViewDelegate {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_GetDelegate).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefViewDelegate{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefView) GetWindow() *ICefWindow {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_GetWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

func (m *ICefView) GetID() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.View_GetID).Call(m.Instance())
	return int32(r1)
}

func (m *ICefView) SetID(id int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetID).Call(m.Instance(), uintptr(id))
}

func (m *ICefView) GetGroupID() int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.View_GetGroupID).Call(m.Instance())
	return int32(r1)
}

func (m *ICefView) SetGroupID(groupId int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetGroupID).Call(m.Instance(), uintptr(groupId))
}

func (m *ICefView) GetParentView() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_GetParentView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefView) GetViewForID(id int32) *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.View_GetViewForID).Call(m.Instance(), uintptr(id), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: getInstance(result)}
	}
	return nil
}

func (m *ICefView) SetBounds(bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
}

func (m *ICefView) GetBounds() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

func (m *ICefView) GetBoundsInScreen() (bounds TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetBoundsInScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&bounds)))
	return
}

func (m *ICefView) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

func (m *ICefView) GetSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

func (m *ICefView) SetPosition(position TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&position)))
}

func (m *ICefView) GetPosition() (point TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return
}

func (m *ICefView) SetInsets(insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetInsets).Call(m.Instance(), uintptr(unsafe.Pointer(&insets)))
}

func (m *ICefView) GetInsets() (insets TCefInsets) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetInsets).Call(m.Instance(), uintptr(unsafe.Pointer(&insets)))
	return
}

func (m *ICefView) GetPreferredSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetPreferredSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

func (m *ICefView) SizeToPreferredSize() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SizeToPreferredSize).Call(m.Instance())
}

func (m *ICefView) GetMinimumSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetMinimumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

func (m *ICefView) GetMaximumSize() (size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_GetMaximumSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
	return
}

func (m *ICefView) GetHeightForWidth(width int32) int32 {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.View_GetHeightForWidth).Call(m.Instance(), uintptr(width))
	return int32(r1)
}

func (m *ICefView) InvalidateLayout() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_InvalidateLayout).Call(m.Instance())
}

func (m *ICefView) SetVisible(visible bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetVisible).Call(m.Instance(), api.PascalBool(visible))
}

func (m *ICefView) IsVisible() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsVisible).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) IsDrawn() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsDrawn).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) SetEnabled(enabled bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetEnabled).Call(m.Instance(), api.PascalBool(enabled))
}

func (m *ICefView) IsEnabled() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsEnabled).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) SetFocusable(focusable bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetFocusable).Call(m.Instance(), api.PascalBool(focusable))
}

func (m *ICefView) IsFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsFocusable).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) IsAccessibilityFocusable() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_IsAccessibilityFocusable).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *ICefView) RequestFocus() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_RequestFocus).Call(m.Instance())
}

func (m *ICefView) SetBackgroundColor(color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.View_SetBackgroundColor).Call(m.Instance(), uintptr(color))
}

func (m *ICefView) GetBackgroundColor() (color types.TCefColor) {
	if !m.IsValid() {
		return 0
	}
	imports.Proc(def.View_GetBackgroundColor).Call(m.Instance(), uintptr(unsafe.Pointer(&color)))
	return
}

func (m *ICefView) ConvertPointToScreen(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_ConvertPointToScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *ICefView) ConvertPointFromScreen(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_ConvertPointFromScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *ICefView) ConvertPointToWindow(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_ConvertPointToWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *ICefView) ConvertPointFromWindow(point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_ConvertPointFromWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *ICefView) ConvertPointToView(view *ICefView, point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_ConvertPointToView).Call(m.Instance(), view.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}

func (m *ICefView) ConvertPointFromView(view *ICefView, point TCefPoint) bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.View_ConvertPointFromView).Call(m.Instance(), view.Instance(), uintptr(unsafe.Pointer(&point)))
	return api.GoBool(r1)
}
