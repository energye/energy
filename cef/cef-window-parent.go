//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TCEFWindowParent struct {
	BaseWinControl
}

func NewCEFWindowParent(owner lcl.IComponent) *TCEFWindowParent {
	m := new(TCEFWindowParent)
	r1, _, _ := Proc(internale_CEFWindow_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	return m
}

func (m *TCEFWindowParent) Handle() types.HWND {
	ret, _, _ := Proc(internale_CEFWindow_GetHandle).Call(m.Instance())
	return types.HWND(ret)
}

func (m *TCEFWindowParent) UpdateSize() {
	Proc(internale_CEFWindow_UpdateSize).Call(m.Instance())
}

func (m *TCEFWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_WindowParent
}

func (m *TCEFWindowParent) SetChromium(chromium IChromium, tag int32) {
}

func (m *TCEFWindowParent) HandleAllocated() bool {
	ret, _, _ := Proc(internale_CEFWindow_HandleAllocated).Call(m.Instance())
	return api.GoBool(ret)
}

func (m *TCEFWindowParent) CreateHandle() {
	Proc(internale_CEFWindow_CreateHandle).Call(m.Instance())
}

func (m *TCEFWindowParent) DestroyChildWindow() bool {
	ret, _, _ := Proc(internale_CEFWindow_DestroyChildWindow).Call(m.Instance())
	return api.GoBool(ret)
}

func (m *TCEFWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	Proc(internale_CEFWindow_OnEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	Proc(internale_CEFWindow_OnExit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowParent) Free() {
	if m.IsValid() {
		Proc(internale_CEFWindow_Free).Call(m.Instance())
		m.instance = nullptr
	}
}

// 获取组件名称。
func (m *TCEFWindowParent) Name() string {
	ret, _, _ := Proc(internale_CEFWindow_GetName).Call(m.Instance())
	return api.GoStr(ret)
}

// 设置组件名称。
func (m *TCEFWindowParent) SetName(value string) {
	Proc(internale_CEFWindow_SetName).Call(m.Instance(), api.PascalStr(value))
}

// 设置控件父容器。
func (m *TCEFWindowParent) SetParent(value lcl.IWinControl) {
	Proc(internale_CEFWindow_SetParent).Call(m.Instance(), lcl.CheckPtr(value))
}

//Align 获取控件自动调整。
func (m *TCEFWindowParent) Align() types.TAlign {
	ret, _, _ := Proc(internale_CEFWindow_GetAlign).Call(m.Instance())
	return types.TAlign(ret)
}

//SetAlign 设置控件自动调整。
func (m *TCEFWindowParent) SetAlign(value types.TAlign) {
	Proc(internale_CEFWindow_SetAlign).Call(m.Instance(), uintptr(value))
}

//Anchors 获取四个角位置的锚点。
func (m *TCEFWindowParent) Anchors() types.TAnchors {
	ret, _, _ := Proc(internale_CEFWindow_GetAnchors).Call(m.Instance())
	return types.TAnchors(ret)
}

//SetAnchors 设置四个角位置的锚点。
func (m *TCEFWindowParent) SetAnchors(value types.TAnchors) {
	Proc(internale_CEFWindow_SetAnchors).Call(m.Instance(), uintptr(value))
}

//Visible 获取控件可视。
func (m *TCEFWindowParent) Visible() bool {
	ret, _, _ := Proc(internale_CEFWindow_GetVisible).Call(m.Instance())
	return api.GoBool(ret)
}

//SetVisible 设置控件可视。
func (m *TCEFWindowParent) SetVisible(value bool) {
	Proc(internale_CEFWindow_SetVisible).Call(m.Instance(), api.PascalBool(value))
}

//Enabled 获取是否启用
func (m *TCEFWindowParent) Enabled() bool {
	ret, _, _ := Proc(internale_CEFWindow_GetEnabled).Call(m.Instance())
	return api.GoBool(ret)
}

//SetEnabled 设置是否启用
func (m *TCEFWindowParent) SetEnabled(value bool) {
	Proc(internale_CEFWindow_SetEnabled).Call(m.Instance(), api.PascalBool(value))
}

//Left 获取左边距
func (m *TCEFWindowParent) Left() int32 {
	ret, _, _ := Proc(internale_CEFWindow_GetLeft).Call(m.Instance())
	return int32(ret)
}

//SetLeft 设置左边距
func (m *TCEFWindowParent) SetLeft(value int32) {
	Proc(internale_CEFWindow_SetLeft).Call(m.Instance(), uintptr(value))
}

//Top 获取上边距
func (m *TCEFWindowParent) Top() int32 {
	ret, _, _ := Proc(internale_CEFWindow_GetTop).Call(m.Instance())
	return int32(ret)
}

//SetTop 设置上边距
func (m *TCEFWindowParent) SetTop(value int32) {
	Proc(internale_CEFWindow_SetTop).Call(m.Instance(), uintptr(value))
}

//Width 获取宽度
func (m *TCEFWindowParent) Width() int32 {
	ret, _, _ := Proc(internale_CEFWindow_GetWidth).Call(m.Instance())
	return int32(ret)
}

//SetWidth 设置宽度
func (m *TCEFWindowParent) SetWidth(value int32) {
	Proc(internale_CEFWindow_SetWidth).Call(m.Instance(), uintptr(value))
}

//Height 获取高度
func (m *TCEFWindowParent) Height() int32 {
	ret, _, _ := Proc(internale_CEFWindow_GetHeight).Call(m.Instance())
	return int32(ret)
}

//SetHeight 设置高度
func (m *TCEFWindowParent) SetHeight(value int32) {
	Proc(internale_CEFWindow_SetHeight).Call(m.Instance(), uintptr(value))
}

func (m *TCEFWindowParent) BoundsRect() (result types.TRect) {
	Proc(internale_CEFWindow_GetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *TCEFWindowParent) SetBoundsRect(value types.TRect) {
	Proc(internale_CEFWindow_SetBoundsRect).Call(m.Instance(), uintptr(unsafe.Pointer(&value)))
}
