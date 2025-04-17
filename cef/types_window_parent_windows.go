//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// TCEFWindowParent 组件
type TCEFWindowParent struct {
	BaseWinControl
	x, y, w, h int32
}

// NewCEFWindowParent 创建一个新的 TCEFWindowParent 组件
func NewCEFWindowParent(owner lcl.IComponent) *TCEFWindowParent {
	m := new(TCEFWindowParent)
	r1, _, _ := imports.Proc(def.CEFWindow_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	return m
}

func (m *TCEFWindowParent) Instance() uintptr {
	return uintptr(m.instance)
}

// Handle 组件句柄
func (m *TCEFWindowParent) Handle() types.HWND {
	ret, _, _ := imports.Proc(def.CEFWindow_GetHandle).Call(m.Instance())
	return types.HWND(ret)
}

// UpdateSize 更新组件大小
func (m *TCEFWindowParent) UpdateSize() {
	imports.Proc(def.CEFWindow_UpdateSize).Call(m.Instance())
}

func (m *TCEFWindowParent) Repaint() {
	imports.Proc(def.CEFWindow_Repaint).Call(m.Instance())
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
	ret, _, _ := imports.Proc(def.CEFWindow_HandleAllocated).Call(m.Instance())
	return api.GoBool(ret)
}

// CreateHandle 创建句柄
func (m *TCEFWindowParent) CreateHandle() {
	imports.Proc(def.CEFWindow_CreateHandle).Call(m.Instance())
}

// DestroyChildWindow 销毁子窗口
func (m *TCEFWindowParent) DestroyChildWindow() bool {
	ret, _, _ := imports.Proc(def.CEFWindow_DestroyChildWindow).Call(m.Instance())
	return api.GoBool(ret)
}

// SetOnEnter 进入事件
func (m *TCEFWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	imports.Proc(def.CEFWindow_OnEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnExit 退出事件
func (m *TCEFWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	imports.Proc(def.CEFWindow_OnExit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Free 释放
func (m *TCEFWindowParent) Free() {
	if m.IsValid() {
		imports.Proc(def.CEFWindow_Free).Call(m.Instance())
		m.instance = nil
	}
}

// Name 获取组件名称
func (m *TCEFWindowParent) Name() string {
	ret, _, _ := imports.Proc(def.CEFWindow_GetName).Call(m.Instance())
	return api.GoStr(ret)
}

// SetName 设置组件名称
func (m *TCEFWindowParent) SetName(value string) {
	imports.Proc(def.CEFWindow_SetName).Call(m.Instance(), api.PascalStr(value))
}

// SetParent 设置控件父容器
func (m *TCEFWindowParent) SetParent(value lcl.IWinControl) {
	imports.Proc(def.CEFWindow_SetParent).Call(m.Instance(), lcl.CheckPtr(value))
}

// RevertCustomAnchors 恢复到自定义四角锚点定位
func (m *TCEFWindowParent) RevertCustomAnchors() {
	m.SetAlign(types.AlCustom)
	m.SetAnchors(types.NewSet())
}

// DefaultAnchors 恢复到默认四角锚点定位
func (m *TCEFWindowParent) DefaultAnchors() {
	m.SetAlign(types.AlClient)
	m.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
}

// Align 获取控件自动调整
func (m *TCEFWindowParent) Align() types.TAlign {
	ret, _, _ := imports.Proc(def.CEFWindow_GetAlign).Call(m.Instance())
	return types.TAlign(ret)
}

// SetAlign 设置控件自动调整
func (m *TCEFWindowParent) SetAlign(value types.TAlign) {
	imports.Proc(def.CEFWindow_SetAlign).Call(m.Instance(), uintptr(value))
}

// Anchors 获取四个角位置的锚点
func (m *TCEFWindowParent) Anchors() types.TAnchors {
	ret, _, _ := imports.Proc(def.CEFWindow_GetAnchors).Call(m.Instance())
	return types.TAnchors(ret)
}

// SetAnchors 设置四个角位置的锚点
func (m *TCEFWindowParent) SetAnchors(value types.TAnchors) {
	imports.Proc(def.CEFWindow_SetAnchors).Call(m.Instance(), uintptr(value))
}

// Visible 获取控件可视
func (m *TCEFWindowParent) Visible() bool {
	ret, _, _ := imports.Proc(def.CEFWindow_GetVisible).Call(m.Instance())
	return api.GoBool(ret)
}

// SetVisible 设置控件可视
func (m *TCEFWindowParent) SetVisible(value bool) {
	imports.Proc(def.CEFWindow_SetVisible).Call(m.Instance(), api.PascalBool(value))
}

// Enabled 获取是否启用
func (m *TCEFWindowParent) Enabled() bool {
	ret, _, _ := imports.Proc(def.CEFWindow_GetEnabled).Call(m.Instance())
	return api.GoBool(ret)
}

// SetEnabled 设置是否启用
func (m *TCEFWindowParent) SetEnabled(value bool) {
	imports.Proc(def.CEFWindow_SetEnabled).Call(m.Instance(), api.PascalBool(value))
}

// Left 获取左边距
func (m *TCEFWindowParent) Left() int32 {
	ret, _, _ := imports.Proc(def.CEFWindow_GetLeft).Call(m.Instance())
	return int32(ret)
}

// SetLeft 设置左边距
func (m *TCEFWindowParent) SetLeft(value int32) {
	m.x = value
	imports.Proc(def.CEFWindow_SetLeft).Call(m.Instance(), uintptr(value))
}

// Top 获取上边距
func (m *TCEFWindowParent) Top() int32 {
	ret, _, _ := imports.Proc(def.CEFWindow_GetTop).Call(m.Instance())
	return int32(ret)
}

// SetTop 设置上边距
func (m *TCEFWindowParent) SetTop(value int32) {
	m.y = value
	imports.Proc(def.CEFWindow_SetTop).Call(m.Instance(), uintptr(value))
}

// Width 获取宽度
func (m *TCEFWindowParent) Width() int32 {
	ret, _, _ := imports.Proc(def.CEFWindow_GetWidth).Call(m.Instance())
	return int32(ret)
}

// SetWidth 设置宽度
func (m *TCEFWindowParent) SetWidth(value int32) {
	m.w = value
	imports.Proc(def.CEFWindow_SetWidth).Call(m.Instance(), uintptr(value))
}

// Height 获取高度
func (m *TCEFWindowParent) Height() int32 {
	ret, _, _ := imports.Proc(def.CEFWindow_GetHeight).Call(m.Instance())
	return int32(ret)
}

// SetHeight 设置高度
func (m *TCEFWindowParent) SetHeight(value int32) {
	m.h = value
	imports.Proc(def.CEFWindow_SetHeight).Call(m.Instance(), uintptr(value))
}

// BoundsRect 获取坐标位置和宽高
func (m *TCEFWindowParent) BoundsRect() (result types.TRect) {
	imports.Proc(def.CEFWindow_GetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetBoundsRect 设置坐标位置和宽高
func (m *TCEFWindowParent) SetBoundsRect(value types.TRect) {
	m.x = value.Left
	m.y = value.Top
	m.w = value.Width()
	m.h = value.Height()
	imports.Proc(def.CEFWindow_SetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}

func (m *TCEFWindowParent) SetTag(tag int) {
	imports.Proc(def.CEFWindow_SetTag).Call(m.Instance(), uintptr(int32(tag)))
}

func (m *TCEFWindowParent) Tag() int {
	r1, _, _ := imports.Proc(def.CEFWindow_GetTag).Call(m.Instance())
	return int(int32(r1))
}

func (m *TCEFWindowParent) point() (x, y int32) {
	return m.x, m.y
}

func (m *TCEFWindowParent) size() (w, h int32) {
	return m.w, m.h
}

func (m *TCEFWindowParent) SetFocus() {
	imports.Proc(def.CEFWindow_SetFocus).Call(m.Instance())
}

func (m *TCEFWindowParent) CanFocus() bool {
	r1, _, _ := imports.Proc(def.CEFWindow_CanFocus).Call(m.Instance())
	return api.GoBool(r1)
}
