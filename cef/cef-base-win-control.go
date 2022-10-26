//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type BaseWinControl struct {
	lcl.IWinControl
	procName string
	instance uintptr
	ptr      unsafe.Pointer
}

func (m *BaseWinControl) IsValid() bool {
	return m.instance != 0
}

// 获取组件名称。
func (m *BaseWinControl) Name() string {
	return GetName(m.procName, m.instance)
}

// 设置组件名称。
func (m *BaseWinControl) SetName(value string) {
	SetName(m.procName, m.instance, value)
}

func (m *BaseWinControl) Free() {
	if m.instance != 0 {
		Free(m.procName, m.instance)
		m.instance, m.ptr = 0, nullptr
	}
}

//Instance 当前实例地址
func (m *BaseWinControl) Instance() uintptr {
	return m.instance
}

// 设置控件父容器。
func (m *BaseWinControl) SetParent(value lcl.IWinControl) {
	SetParent(m.procName, m.instance, lcl.CheckPtr(value))
}

//Align 获取控件自动调整。
func (m *BaseWinControl) Align() types.TAlign {
	return Align(m.procName, m.instance)
}

//SetAlign 设置控件自动调整。
func (m *BaseWinControl) SetAlign(value types.TAlign) {
	SetAlign(m.procName, m.instance, value)
}

//Anchors 获取四个角位置的锚点。
func (m *BaseWinControl) Anchors() types.TAnchors {
	return Anchors(m.procName, m.instance)
}

//SetAnchors 设置四个角位置的锚点。
func (m *BaseWinControl) SetAnchors(value types.TAnchors) {
	SetAnchors(m.procName, m.instance, value)
}

//Visible 获取控件可视。
func (m *BaseWinControl) Visible() bool {
	return GetVisible(m.procName, m.instance)
}

//SetVisible 设置控件可视。
func (m *BaseWinControl) SetVisible(value bool) {
	SetVisible(m.procName, m.instance, value)
}

//Enabled 获取是否启用
func (m *BaseWinControl) Enabled() bool {
	return GetEnabled(m.procName, m.instance)
}

//SetEnabled 设置是否启用
func (m *BaseWinControl) SetEnabled(value bool) {
	SetEnabled(m.procName, m.instance, value)
}

//Left 获取左边距
func (m *BaseWinControl) Left() int32 {
	return GetLeft(m.procName, m.instance)
}

//SetLeft 设置左边距
func (m *BaseWinControl) SetLeft(value int32) {
	SetLeft(m.procName, m.instance, value)
}

//Top 获取上边距
func (m *BaseWinControl) Top() int32 {
	return GetTop(m.procName, m.instance)
}

//SetTop 设置上边距
func (m *BaseWinControl) SetTop(value int32) {
	SetTop(m.procName, m.instance, value)
}

//Width 获取宽度
func (m *BaseWinControl) Width() int32 {
	return GetWidth(m.procName, m.instance)
}

//SetWidth 设置宽度
func (m *BaseWinControl) SetWidth(value int32) {
	SetWidth(m.procName, m.instance, value)
}

//Height 获取高度
func (m *BaseWinControl) Height() int32 {
	return GetHeight(m.procName, m.instance)
}

//SetHeight 设置高度
func (m *BaseWinControl) SetHeight(value int32) {
	SetHeight(m.procName, m.instance, value)
}

func (m *BaseWinControl) BoundsRect() types.TRect {
	return GetBoundsRect(m.procName, m.instance)
}

func (m *BaseWinControl) SetBoundsRect(value types.TRect) {
	SetBoundsRect(m.procName, m.instance, value)
}
