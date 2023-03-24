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
// MacOSX, Linux
package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// TCEFLinkedWindowParent 组件
type TCEFLinkedWindowParent struct {
	BaseWinControl
}

// NewCEFLinkedWindowParent 创建一个新的WindowParent组件
func NewCEFLinkedWindowParent(owner lcl.IComponent) *TCEFLinkedWindowParent {
	m := new(TCEFLinkedWindowParent)
	r1, _, _ := imports.Proc(internale_CEFLinkedWindow_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	return m
}

// Handle 组件句柄
func (m *TCEFLinkedWindowParent) Handle() types.HWND {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetHandle).Call(m.Instance())
	return types.HWND(ret)
}

// UpdateSize 更新组件大小
func (m *TCEFLinkedWindowParent) UpdateSize() {
	imports.Proc(internale_CEFLinkedWindow_UpdateSize).Call(m.Instance())
}

// Type 组件类型, 这里返回 TCEFLinkedWindowParent 类型
func (m *TCEFLinkedWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_LinkedWindowParent
}

// SetChromium 设置 IChromium, 只 TCEFLinkedWindowParent 有效
func (m *TCEFLinkedWindowParent) SetChromium(chromium IChromium, tag int32) {
	imports.Proc(internale_CEFLinkedWindow_SetChromium).Call(uintptr(m.instance), chromium.Instance(), uintptr(tag))
}

// HandleAllocated 处理所有
func (m *TCEFLinkedWindowParent) HandleAllocated() bool {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_HandleAllocated).Call(m.Instance())
	return api.GoBool(ret)
}

// CreateHandle 创建句柄
func (m *TCEFLinkedWindowParent) CreateHandle() {
	imports.Proc(internale_CEFLinkedWindow_CreateHandle).Call(m.Instance())
}

// DestroyChildWindow 销毁子窗口
func (m *TCEFLinkedWindowParent) DestroyChildWindow() bool {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_DestroyChildWindow).Call(m.Instance())
	return api.GoBool(ret)
}

// SetOnEnter 进入事件
func (m *TCEFLinkedWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	imports.Proc(internale_CEFLinkedWindow_OnEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnExit 退出事件
func (m *TCEFLinkedWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	imports.Proc(internale_CEFLinkedWindow_OnExit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Free 释放
func (m *TCEFLinkedWindowParent) Free() {
	if m.IsValid() {
		imports.Proc(internale_CEFLinkedWindow_Free).Call(m.Instance())
		m.instance = nil
	}
}

// Name 获取组件名称
func (m *TCEFLinkedWindowParent) Name() string {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetName).Call(m.Instance())
	return api.GoStr(ret)
}

// SetName 设置组件名称
func (m *TCEFLinkedWindowParent) SetName(value string) {
	imports.Proc(internale_CEFLinkedWindow_SetName).Call(m.Instance(), api.PascalStr(value))
}

// SetParent 设置控件父容器
func (m *TCEFLinkedWindowParent) SetParent(value lcl.IWinControl) {
	imports.Proc(internale_CEFLinkedWindow_SetParent).Call(m.Instance(), lcl.CheckPtr(value))
}

// Align Align 获取控件自动调整
func (m *TCEFLinkedWindowParent) Align() types.TAlign {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetAlign).Call(m.Instance())
	return types.TAlign(ret)
}

// SetAlign 设置控件自动调整
func (m *TCEFLinkedWindowParent) SetAlign(value types.TAlign) {
	imports.Proc(internale_CEFLinkedWindow_SetAlign).Call(m.Instance(), uintptr(value))
}

// Anchors 获取四个角位置的锚点
func (m *TCEFLinkedWindowParent) Anchors() types.TAnchors {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetAnchors).Call(m.Instance())
	return types.TAnchors(ret)
}

// SetAnchors 设置四个角位置的锚点
func (m *TCEFLinkedWindowParent) SetAnchors(value types.TAnchors) {
	imports.Proc(internale_CEFLinkedWindow_SetAnchors).Call(m.Instance(), uintptr(value))
}

// Visible 获取控件可视
func (m *TCEFLinkedWindowParent) Visible() bool {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetVisible).Call(m.Instance())
	return api.GoBool(ret)
}

// SetVisible 设置控件可视
func (m *TCEFLinkedWindowParent) SetVisible(value bool) {
	imports.Proc(internale_CEFLinkedWindow_SetVisible).Call(m.Instance(), api.PascalBool(value))
}

// Enabled 获取是否启用
func (m *TCEFLinkedWindowParent) Enabled() bool {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetEnabled).Call(m.Instance())
	return api.GoBool(ret)
}

// SetEnabled 设置是否启用
func (m *TCEFLinkedWindowParent) SetEnabled(value bool) {
	imports.Proc(internale_CEFLinkedWindow_SetEnabled).Call(m.Instance(), api.PascalBool(value))
}

// Left 获取左边距
func (m *TCEFLinkedWindowParent) Left() int32 {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetLeft).Call(m.Instance())
	return int32(ret)
}

// SetLeft 设置左边距
func (m *TCEFLinkedWindowParent) SetLeft(value int32) {
	imports.Proc(internale_CEFLinkedWindow_SetLeft).Call(m.Instance(), uintptr(value))
}

// Top 获取上边距
func (m *TCEFLinkedWindowParent) Top() int32 {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetTop).Call(m.Instance())
	return int32(ret)
}

// SetTop 设置上边距
func (m *TCEFLinkedWindowParent) SetTop(value int32) {
	imports.Proc(internale_CEFLinkedWindow_SetTop).Call(m.Instance(), uintptr(value))
}

// Width 获取宽度
func (m *TCEFLinkedWindowParent) Width() int32 {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetWidth).Call(m.Instance())
	return int32(ret)
}

// SetWidth 设置宽度
func (m *TCEFLinkedWindowParent) SetWidth(value int32) {
	imports.Proc(internale_CEFLinkedWindow_SetWidth).Call(m.Instance(), uintptr(value))
}

// Height 获取高度
func (m *TCEFLinkedWindowParent) Height() int32 {
	ret, _, _ := imports.Proc(internale_CEFLinkedWindow_GetHeight).Call(m.Instance())
	return int32(ret)
}

// SetHeight 设置高度
func (m *TCEFLinkedWindowParent) SetHeight(value int32) {
	imports.Proc(internale_CEFLinkedWindow_SetHeight).Call(m.Instance(), uintptr(value))
}

// BoundsRect 获取矩形边界
func (m *TCEFLinkedWindowParent) BoundsRect() (result types.TRect) {
	imports.Proc(internale_CEFLinkedWindow_GetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetBoundsRect 设置矩形边界
func (m *TCEFLinkedWindowParent) SetBoundsRect(value types.TRect) {
	imports.Proc(internale_CEFLinkedWindow_SetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}
