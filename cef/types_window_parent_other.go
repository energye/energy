//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows
// +build !windows

// CEFWindowParent组件
// MacOSX, Linux

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

// TCEFLinkedWindowParent 组件
type TCEFLinkedWindowParent struct {
	BaseWinControl
	x, y, w, h int32
}

// NewCEFWindowParent 创建一个新的 TCEFLinkedWindowParent 组件
func NewCEFWindowParent(owner lcl.IComponent) *TCEFLinkedWindowParent {
	m := new(TCEFLinkedWindowParent)
	r1, _, _ := imports.Proc(def.CEFLinkedWindow_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	return m
}

func (m *TCEFLinkedWindowParent) Instance() uintptr {
	return uintptr(m.instance)
}

// Handle 组件句柄
func (m *TCEFLinkedWindowParent) Handle() types.HWND {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetHandle).Call(m.Instance())
	return types.HWND(ret)
}

// UpdateSize 更新组件大小
func (m *TCEFLinkedWindowParent) UpdateSize() {
	imports.Proc(def.CEFLinkedWindow_UpdateSize).Call(m.Instance())
}

func (m *TCEFLinkedWindowParent) Repaint() {
	imports.Proc(def.CEFLinkedWindow_Repaint).Call(m.Instance())
}

// Type 组件类型, 这里返回 TCEFLinkedWindowParent 类型
func (m *TCEFLinkedWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_LinkedWindowParent
}

// SetChromium 设置 IChromium, 只 TCEFLinkedWindowParent 有效
func (m *TCEFLinkedWindowParent) SetChromium(chromium IChromium, tag int32) {
	imports.Proc(def.CEFLinkedWindow_SetChromium).Call(uintptr(m.instance), chromium.Instance(), uintptr(tag))
}

// HandleAllocated 处理所有
func (m *TCEFLinkedWindowParent) HandleAllocated() bool {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_HandleAllocated).Call(m.Instance())
	return api.GoBool(ret)
}

// CreateHandle 创建句柄
func (m *TCEFLinkedWindowParent) CreateHandle() {
	imports.Proc(def.CEFLinkedWindow_CreateHandle).Call(m.Instance())
}

// DestroyChildWindow 销毁子窗口
func (m *TCEFLinkedWindowParent) DestroyChildWindow() bool {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_DestroyChildWindow).Call(m.Instance())
	return api.GoBool(ret)
}

// SetOnEnter 进入事件
func (m *TCEFLinkedWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	imports.Proc(def.CEFLinkedWindow_OnEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnExit 退出事件
func (m *TCEFLinkedWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	imports.Proc(def.CEFLinkedWindow_OnExit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Free 释放
func (m *TCEFLinkedWindowParent) Free() {
	if m.IsValid() {
		imports.Proc(def.CEFLinkedWindow_Free).Call(m.Instance())
		m.instance = nil
	}
}

// Name 获取组件名称
func (m *TCEFLinkedWindowParent) Name() string {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetName).Call(m.Instance())
	return api.GoStr(ret)
}

// SetName 设置组件名称
func (m *TCEFLinkedWindowParent) SetName(value string) {
	imports.Proc(def.CEFLinkedWindow_SetName).Call(m.Instance(), api.PascalStr(value))
}

// SetParent 设置控件父容器
func (m *TCEFLinkedWindowParent) SetParent(value lcl.IWinControl) {
	imports.Proc(def.CEFLinkedWindow_SetParent).Call(m.Instance(), lcl.CheckPtr(value))
}

// RevertCustomAnchors 恢复到自定义四角锚点定位
func (m *TCEFLinkedWindowParent) RevertCustomAnchors() {
	m.SetAlign(types.AlCustom)
	m.SetAnchors(types.NewSet())
}

// DefaultAnchors 恢复到默认四角锚点定位
func (m *TCEFLinkedWindowParent) DefaultAnchors() {
	m.SetAlign(types.AlClient)
	m.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))
}

// Align Align 获取控件自动调整
func (m *TCEFLinkedWindowParent) Align() types.TAlign {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetAlign).Call(m.Instance())
	return types.TAlign(ret)
}

// SetAlign 设置控件自动调整
func (m *TCEFLinkedWindowParent) SetAlign(value types.TAlign) {
	imports.Proc(def.CEFLinkedWindow_SetAlign).Call(m.Instance(), uintptr(value))
}

// Anchors 获取四个角位置的锚点
func (m *TCEFLinkedWindowParent) Anchors() types.TAnchors {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetAnchors).Call(m.Instance())
	return types.TAnchors(ret)
}

// SetAnchors 设置四个角位置的锚点
func (m *TCEFLinkedWindowParent) SetAnchors(value types.TAnchors) {
	imports.Proc(def.CEFLinkedWindow_SetAnchors).Call(m.Instance(), uintptr(value))
}

// Visible 获取控件可视
func (m *TCEFLinkedWindowParent) Visible() bool {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetVisible).Call(m.Instance())
	return api.GoBool(ret)
}

// SetVisible 设置控件可视
func (m *TCEFLinkedWindowParent) SetVisible(value bool) {
	imports.Proc(def.CEFLinkedWindow_SetVisible).Call(m.Instance(), api.PascalBool(value))
}

// Enabled 获取是否启用
func (m *TCEFLinkedWindowParent) Enabled() bool {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetEnabled).Call(m.Instance())
	return api.GoBool(ret)
}

// SetEnabled 设置是否启用
func (m *TCEFLinkedWindowParent) SetEnabled(value bool) {
	imports.Proc(def.CEFLinkedWindow_SetEnabled).Call(m.Instance(), api.PascalBool(value))
}

// Left 获取左边距
func (m *TCEFLinkedWindowParent) Left() int32 {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetLeft).Call(m.Instance())
	return int32(ret)
}

// SetLeft 设置左边距
func (m *TCEFLinkedWindowParent) SetLeft(value int32) {
	m.x = value
	imports.Proc(def.CEFLinkedWindow_SetLeft).Call(m.Instance(), uintptr(value))
}

// Top 获取上边距
func (m *TCEFLinkedWindowParent) Top() int32 {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetTop).Call(m.Instance())
	return int32(ret)
}

// SetTop 设置上边距
func (m *TCEFLinkedWindowParent) SetTop(value int32) {
	m.y = value
	imports.Proc(def.CEFLinkedWindow_SetTop).Call(m.Instance(), uintptr(value))
}

// Width 获取宽度
func (m *TCEFLinkedWindowParent) Width() int32 {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetWidth).Call(m.Instance())
	return int32(ret)
}

// SetWidth 设置宽度
func (m *TCEFLinkedWindowParent) SetWidth(value int32) {
	m.w = value
	imports.Proc(def.CEFLinkedWindow_SetWidth).Call(m.Instance(), uintptr(value))
}

// Height 获取高度
func (m *TCEFLinkedWindowParent) Height() int32 {
	ret, _, _ := imports.Proc(def.CEFLinkedWindow_GetHeight).Call(m.Instance())
	return int32(ret)
}

// SetHeight 设置高度
func (m *TCEFLinkedWindowParent) SetHeight(value int32) {
	m.h = value
	imports.Proc(def.CEFLinkedWindow_SetHeight).Call(m.Instance(), uintptr(value))
}

// BoundsRect 获取矩形边界
func (m *TCEFLinkedWindowParent) BoundsRect() (result types.TRect) {
	imports.Proc(def.CEFLinkedWindow_GetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

// SetBoundsRect 设置矩形边界
func (m *TCEFLinkedWindowParent) SetBoundsRect(value types.TRect) {
	m.x = value.Left
	m.y = value.Top
	m.w = value.Width()
	m.h = value.Height()
	imports.Proc(def.CEFLinkedWindow_SetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}

func (m *TCEFLinkedWindowParent) SetTag(tag int) {
	imports.Proc(def.CEFLinkedWindow_SetTag).Call(m.Instance(), uintptr(int32(tag)))
}

func (m *TCEFLinkedWindowParent) Tag() int {
	r1, _, _ := imports.Proc(def.CEFLinkedWindow_GetTag).Call(m.Instance())
	return int(int32(r1))
}

func (m *TCEFLinkedWindowParent) point() (x, y int32) {
	return m.x, m.y
}

func (m *TCEFLinkedWindowParent) size() (w, h int32) {
	return m.w, m.h
}

func (m *TCEFLinkedWindowParent) SetFocus() {
	imports.Proc(def.CEFLinkedWindow_SetFocus).Call(m.Instance())
}

func (m *TCEFLinkedWindowParent) CanFocus() bool {
	r1, _, _ := imports.Proc(def.CEFLinkedWindow_CanFocus).Call(m.Instance())
	return api.GoBool(r1)
}
