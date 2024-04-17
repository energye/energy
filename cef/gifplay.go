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
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// TGIFPlay GIF 图片播放组件
type TGIFPlay struct {
	instance unsafe.Pointer
}

// NewGIFPlay
//
// 创建一个新的对象。
func NewGIFPlay(owner lcl.IComponent) *TGIFPlay {
	m := new(TGIFPlay)
	r1, _, _ := imports.Proc(def.GIFPlay_Create).Call(owner.Instance())
	m.instance = unsafe.Pointer(r1)
	return m
}

// SetParent
//
// 设置控件父容器。
func (m *TGIFPlay) SetParent(value lcl.IWinControl) {
	imports.Proc(def.GIFPlay_SetParent).Call(m.Instance(), value.Instance())
}

// Free
//
// 释放对象。
func (m *TGIFPlay) Free() {
	if m.instance != nil {
		imports.Proc(def.GIFPlay_Free).Call(m.Instance())
		m.instance = nil
	}
}

func (m *TGIFPlay) SetColor(value types.TColor) {
	imports.Proc(def.GIFPlay_SetColor).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) Dragging() bool {
	r1, _, _ := imports.Proc(def.GIFPlay_Dragging).Call(m.Instance())
	return api.GoBool(r1)
}

func (m *TGIFPlay) Animate(value bool) {
	imports.Proc(def.GIFPlay_Animate).Call(m.Instance(), api.PascalBool(value))
}

func (m *TGIFPlay) Start() {
	m.Animate(true)
}

func (m *TGIFPlay) Stop() {
	m.Animate(false)
}

func (m *TGIFPlay) NextFrame() {
	imports.Proc(def.GIFPlay_NextFrame).Call(m.Instance())
}

func (m *TGIFPlay) PriorFrame() {
	imports.Proc(def.GIFPlay_PriorFrame).Call(m.Instance())
}

func (m *TGIFPlay) Empty() {
	imports.Proc(def.GIFPlay_Empty).Call(m.Instance())
}

func (m *TGIFPlay) CurrentImageIndex() int32 {
	r1, _, _ := imports.Proc(def.GIFPlay_CurrentImageIndex).Call(m.Instance())
	return int32(r1)
}

func (m *TGIFPlay) LoadFromFile(filePath string) {
	imports.Proc(def.GIFPlay_LoadFromFile).Call(m.Instance(), api.PascalStr(filePath))
}

func (m *TGIFPlay) LoadFromStream(stream lcl.IStream) {
	imports.Proc(def.GIFPlay_LoadFromStream).Call(m.Instance(), stream.Instance())
}

func (m *TGIFPlay) Left() int32 {
	r1, _, _ := imports.Proc(def.GIFPlay_GetLeft).Call(m.Instance())
	return int32(r1)
}

func (m *TGIFPlay) SetLeft(value int32) {
	imports.Proc(def.GIFPlay_SetLeft).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) Top() int32 {
	r1, _, _ := imports.Proc(def.GIFPlay_GetTop).Call(m.Instance())
	return int32(r1)
}

func (m *TGIFPlay) SetTop(value int32) {
	imports.Proc(def.GIFPlay_SetTop).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) Width() int32 {
	r1, _, _ := imports.Proc(def.GIFPlay_GetWidth).Call(m.Instance())
	return int32(r1)
}

func (m *TGIFPlay) SetWidth(value int32) {
	imports.Proc(def.GIFPlay_SetWidth).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) Height() int32 {
	r1, _, _ := imports.Proc(def.GIFPlay_GetHeight).Call(m.Instance())
	return int32(r1)
}

func (m *TGIFPlay) SetHeight(value int32) {
	imports.Proc(def.GIFPlay_SetHeight).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	imports.Proc(def.GIFPlay_SetBounds).Call(m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func (m *TGIFPlay) Update() {
	imports.Proc(def.GIFPlay_Update).Call(m.Instance())
}

func (m *TGIFPlay) SetAlign(value types.TAlign) {
	imports.Proc(def.GIFPlay_SetAlign).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) SetAnchors(value types.TAnchors) {
	imports.Proc(def.GIFPlay_SetAnchors).Call(m.Instance(), uintptr(value))
}

func (m *TGIFPlay) SetAutoSize(value bool) {
	imports.Proc(def.GIFPlay_SetAutoSize).Call(m.Instance(), api.PascalBool(value))
}

func (m *TGIFPlay) SetVisible(value bool) {
	imports.Proc(def.GIFPlay_SetVisible).Call(m.Instance(), api.PascalBool(value))
}

func (m *TGIFPlay) SetOnClick(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnClick).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnFrameChanged(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnFrameChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnDblClick(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnDblClick).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnMouseDown(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnMouseDown).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnMouseEnter(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnMouseEnter).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnMouseLeave(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnMouseLeave).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	imports.Proc(def.GIFPlay_SetOnMouseMove).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnMouseUp(fn lcl.TMouseEvent) {
	imports.Proc(def.GIFPlay_SetOnMouseUp).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnStart(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnStart).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TGIFPlay) SetOnStop(fn lcl.TNotifyEvent) {
	imports.Proc(def.GIFPlay_SetOnStop).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Instance
//
// 返回对象实例指针。
func (m *TGIFPlay) Instance() uintptr {
	return uintptr(m.instance)
}

// UnsafeAddr
//
// 获取一个不安全的地址。
func (m *TGIFPlay) UnsafeAddr() unsafe.Pointer {
	return m.instance
}

// IsValid
//
// 检测地址是否为空。
func (m *TGIFPlay) IsValid() bool {
	return m.instance != nil
}
