//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEFWindowParent组件
// Windows
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// TCEFWindowParent 组件
type TCEFWindowParent struct {
	BaseWinControl
}

// NewCEFWindowParent 创建一个新的WindowParent组件
func NewCEFWindowParent(owner lcl.IComponent) *TCEFWindowParent {
	m := new(TCEFWindowParent)
	r1, _, _ := imports.Proc(internale_CEFWindow_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	return m
}

// Handle 组件句柄
func (m *TCEFWindowParent) Handle() types.HWND {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetHandle).Call(m.Instance())
	return types.HWND(ret)
}

// UpdateSize 更新组件大小
func (m *TCEFWindowParent) UpdateSize() {
	imports.Proc(internale_CEFWindow_UpdateSize).Call(m.Instance())
}

// Type 组件类型, 这里返回 TCEFWindowParent 类型
func (m *TCEFWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_WindowParent
}

// SetChromium 设置 IChromium, 只 TCEFLinkedWindowParent 有效
func (m *TCEFWindowParent) SetChromium(chromium IChromium, tag int32) {
}

// HandleAllocated 处理所有
func (m *TCEFWindowParent) HandleAllocated() bool {
	ret, _, _ := imports.Proc(internale_CEFWindow_HandleAllocated).Call(m.Instance())
	return api.GoBool(ret)
}

// CreateHandle 创建句柄
func (m *TCEFWindowParent) CreateHandle() {
	imports.Proc(internale_CEFWindow_CreateHandle).Call(m.Instance())
}

// DestroyChildWindow 销毁子窗口
func (m *TCEFWindowParent) DestroyChildWindow() bool {
	ret, _, _ := imports.Proc(internale_CEFWindow_DestroyChildWindow).Call(m.Instance())
	return api.GoBool(ret)
}

// SetOnEnter 进入事件
func (m *TCEFWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	imports.Proc(internale_CEFWindow_OnEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnExit 退出事件
func (m *TCEFWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	imports.Proc(internale_CEFWindow_OnExit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Free 释放
func (m *TCEFWindowParent) Free() {
	if m.IsValid() {
		imports.Proc(internale_CEFWindow_Free).Call(m.Instance())
		m.instance = nullptr
	}
}

// Name 获取组件名称
func (m *TCEFWindowParent) Name() string {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetName).Call(m.Instance())
	return api.GoStr(ret)
}

// SetName 设置组件名称
func (m *TCEFWindowParent) SetName(value string) {
	imports.Proc(internale_CEFWindow_SetName).Call(m.Instance(), api.PascalStr(value))
}

// SetParent 设置控件父容器
func (m *TCEFWindowParent) SetParent(value lcl.IWinControl) {
	imports.Proc(internale_CEFWindow_SetParent).Call(m.Instance(), lcl.CheckPtr(value))
}

// Align 获取控件自动调整
func (m *TCEFWindowParent) Align() types.TAlign {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetAlign).Call(m.Instance())
	return types.TAlign(ret)
}

// SetAlign 设置控件自动调整
func (m *TCEFWindowParent) SetAlign(value types.TAlign) {
	imports.Proc(internale_CEFWindow_SetAlign).Call(m.Instance(), uintptr(value))
}

// Anchors 获取四个角位置的锚点
func (m *TCEFWindowParent) Anchors() types.TAnchors {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetAnchors).Call(m.Instance())
	return types.TAnchors(ret)
}

// SetAnchors 设置四个角位置的锚点
func (m *TCEFWindowParent) SetAnchors(value types.TAnchors) {
	imports.Proc(internale_CEFWindow_SetAnchors).Call(m.Instance(), uintptr(value))
}

// Visible 获取控件可视
func (m *TCEFWindowParent) Visible() bool {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetVisible).Call(m.Instance())
	return api.GoBool(ret)
}

// SetVisible 设置控件可视
func (m *TCEFWindowParent) SetVisible(value bool) {
	imports.Proc(internale_CEFWindow_SetVisible).Call(m.Instance(), api.PascalBool(value))
}

// Enabled 获取是否启用
func (m *TCEFWindowParent) Enabled() bool {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetEnabled).Call(m.Instance())
	return api.GoBool(ret)
}

// SetEnabled 设置是否启用
func (m *TCEFWindowParent) SetEnabled(value bool) {
	imports.Proc(internale_CEFWindow_SetEnabled).Call(m.Instance(), api.PascalBool(value))
}

// Left 获取左边距
func (m *TCEFWindowParent) Left() int32 {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetLeft).Call(m.Instance())
	return int32(ret)
}

// SetLeft 设置左边距
func (m *TCEFWindowParent) SetLeft(value int32) {
	imports.Proc(internale_CEFWindow_SetLeft).Call(m.Instance(), uintptr(value))
}

// Top 获取上边距
func (m *TCEFWindowParent) Top() int32 {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetTop).Call(m.Instance())
	return int32(ret)
}

// SetTop 设置上边距
func (m *TCEFWindowParent) SetTop(value int32) {
	imports.Proc(internale_CEFWindow_SetTop).Call(m.Instance(), uintptr(value))
}

// Width 获取宽度
func (m *TCEFWindowParent) Width() int32 {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetWidth).Call(m.Instance())
	return int32(ret)
}

// SetWidth 设置宽度
func (m *TCEFWindowParent) SetWidth(value int32) {
	imports.Proc(internale_CEFWindow_SetWidth).Call(m.Instance(), uintptr(value))
}

// Height 获取高度
func (m *TCEFWindowParent) Height() int32 {
	ret, _, _ := imports.Proc(internale_CEFWindow_GetHeight).Call(m.Instance())
	return int32(ret)
}

// SetHeight 设置高度
func (m *TCEFWindowParent) SetHeight(value int32) {
	imports.Proc(internale_CEFWindow_SetHeight).Call(m.Instance(), uintptr(value))
}

// BoundsRect 获取矩形边界
func (m *TCEFWindowParent) BoundsRect() (result types.TRect) {
	imports.Proc(internale_CEFWindow_GetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetBoundsRect 设置矩形边界
func (m *TCEFWindowParent) SetBoundsRect(value types.TRect) {
	imports.Proc(internale_CEFWindow_SetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}
