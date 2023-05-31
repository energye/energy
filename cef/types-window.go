//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// cef window
//
// VF窗口组件
package cef

import (
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl/api"
	t "github.com/energye/golcl/lcl/types"
	"io/ioutil"
	"unsafe"
)

func (m *ICefWindow) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *ICefWindow) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

// Show 显示窗口
func (m *ICefWindow) Show() {
	imports.Proc(internale_ICEFWindow_Show).Call(uintptr(m.instance))
}

// Hide 显示窗口
func (m *ICefWindow) Hide() {
	imports.Proc(internale_ICEFWindow_Hide).Call(uintptr(m.instance))
}

// CenterWindow 根据大小窗口居中
func (m *ICefWindow) CenterWindow(size *TCefSize) {
	imports.Proc(internale_ICEFWindow_CenterWindow).Call(uintptr(m.instance), uintptr(unsafe.Pointer(size)))
}

// Close 关闭窗口， 主窗口调用
func (m *ICefWindow) Close() {
	imports.Proc(internale_ICEFWindow_Close).Call(uintptr(m.instance))
}

// IsClosed 是否关闭
func (m *ICefWindow) IsClosed() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsClosed).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

// Activate 激活窗口
func (m *ICefWindow) Activate() {
	imports.Proc(internale_ICEFWindow_Activate).Call(uintptr(m.instance))
}

// Deactivate 停止激活窗口
func (m *ICefWindow) Deactivate() {
	imports.Proc(internale_ICEFWindow_Deactivate).Call(uintptr(m.instance))
}

// IsActive 是否激活
func (m *ICefWindow) IsActive() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsActive).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

// BringToTop 将窗口移至最上层
func (m *ICefWindow) BringToTop() {
	imports.Proc(internale_ICEFWindow_BringToTop).Call(uintptr(m.instance))
}

// SetAlwaysOnTop 设置窗口是否置顶
func (m *ICefWindow) SetAlwaysOnTop(onTop bool) {
	imports.Proc(internale_ICEFWindow_SetAlwaysOnTop).Call(uintptr(m.instance), api.PascalBool(onTop))
}

// IsAlwaysOnTop 窗口是否置顶
func (m *ICefWindow) IsAlwaysOnTop() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsAlwaysOnTop).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

// WindowState 返回窗口最小化、最大化、全屏状态
func (m *ICefWindow) WindowState() t.TWindowState {
	if m.IsMinimized() {
		return t.WsMinimized
	} else if m.IsMaximized() {
		return t.WsMaximized
	} else if m.IsFullscreen() {
		return t.WsFullScreen
	}
	return t.WsNormal
}

// Maximize 最大化窗口
func (m *ICefWindow) Maximize() {
	imports.Proc(internale_ICEFWindow_Maximize).Call(uintptr(m.instance))
}

// Minimize 最小化窗口
func (m *ICefWindow) Minimize() {
	imports.Proc(internale_ICEFWindow_Minimize).Call(uintptr(m.instance))
}

// Restore 窗口还原
func (m *ICefWindow) Restore() {
	imports.Proc(internale_ICEFWindow_Restore).Call(uintptr(m.instance))
}

// SetFullscreen 设置窗口全屏
func (m *ICefWindow) SetFullscreen(fullscreen bool) {
	imports.Proc(internale_ICEFWindow_SetFullscreen).Call(uintptr(m.instance), api.PascalBool(fullscreen))
}

// SetBackgroundColor 设置背景色
func (m *ICefWindow) SetBackgroundColor(rect *types.TCefColor) {
	imports.Proc(internale_ICEFWindow_SetBackgroundColor).Call(uintptr(m.instance), rect.ToPtr())
}

// SetBounds 设置窗口边界
func (m *ICefWindow) SetBounds(rect *TCefRect) {
	imports.Proc(internale_ICEFWindow_SetBounds).Call(uintptr(m.instance), uintptr(unsafe.Pointer(rect)))
}

// SetSize 设置窗口宽高
func (m *ICefWindow) SetSize(size *TCefSize) {
	imports.Proc(internale_ICEFWindow_SetSize).Call(uintptr(m.instance), uintptr(unsafe.Pointer(size)))
}

// SetPosition 设置窗口位置
func (m *ICefWindow) SetPosition(point *TCefPoint) {
	imports.Proc(internale_ICEFWindow_SetPosition).Call(uintptr(m.instance), uintptr(unsafe.Pointer(point)))
}

// IsMaximized 是否最大化
func (m *ICefWindow) IsMaximized() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsMaximized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

// IsMinimized 是否最小化
func (m *ICefWindow) IsMinimized() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsMinimized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

// IsFullscreen 是否全屏
func (m *ICefWindow) IsFullscreen() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsFullscreen).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

// SetTitle 设置窗口标题
func (m *ICefWindow) SetTitle(title string) {
	imports.Proc(internale_ICEFWindow_SetTitle).Call(uintptr(m.instance), api.PascalStr(title))
}

// Title 获取窗口标题
func (m *ICefWindow) Title() string {
	r1, _, _ := imports.Proc(internale_ICEFWindow_GetTitle).Call(uintptr(m.instance))
	return api.GoStr(r1)
}

// SetWindowIcon 设置窗口图标
func (m *ICefWindow) SetWindowIcon(scaleFactor float32, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// SetWindowIconFS 设置窗口图标
func (m *ICefWindow) SetWindowIconFS(scaleFactor float32, filename string) error {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// WindowIcon 获取窗口图标
func (m *ICefWindow) WindowIcon() *ICefImage {
	var result uintptr
	imports.Proc(internale_ICEFWindow_GetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// WindowAppIcon 获取窗口应用图标
func (m *ICefWindow) WindowAppIcon() *ICefImage {
	var result uintptr
	imports.Proc(internale_ICEFWindow_GetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// SetWindowAppIcon 设置窗口应用图标
func (m *ICefWindow) SetWindowAppIcon(scaleFactor float32, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// SetWindowAppIconFS 设置窗口应用图标
func (m *ICefWindow) SetWindowAppIconFS(scaleFactor float32, filename string) error {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// AddOverlayView TODO 未实现
func (m *ICefWindow) AddOverlayView() {
	//do not implement
	//imports.Proc(internale_ICEFWindow_AddOverlayView).Call(uintptr(m.instance))
}

// ShowMenu 显示菜单
func (m *ICefWindow) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	imports.Proc(internale_ICEFWindow_ShowMenu).Call(uintptr(m.instance), uintptr(menuModel.instance), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

// CancelMenu 取消菜单
func (m *ICefWindow) CancelMenu() {
	imports.Proc(internale_ICEFWindow_CancelMenu).Call(uintptr(m.instance))
}

// Display
func (m *ICefWindow) Display() *ICefDisplay {
	var ret uintptr
	imports.Proc(internale_ICEFWindow_GetDisplay).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefDisplay{
		instance: unsafe.Pointer(ret),
	}
}

// ClientAreaBoundsInScreen 获取客户端所在指定屏幕位置
func (m *ICefWindow) ClientAreaBoundsInScreen() (result TCefRect) {
	imports.Proc(internale_ICEFWindow_GetClientAreaBoundsInScreen).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

// SetDraggableRegions 设置拖拽区域
func (m *ICefWindow) SetDraggableRegions(regions []TCefDraggableRegion) {
	imports.Proc(internale_ICEFWindow_SetDraggableRegions).Call(uintptr(m.instance), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

// WindowHandle 获取窗口句柄
func (m *ICefWindow) WindowHandle() consts.TCefWindowHandle {
	r1, _, _ := imports.Proc(internale_ICEFWindow_GetWindowHandle).Call(uintptr(m.instance))
	return consts.TCefWindowHandle(r1)
}

// SendKeyPress 发送键盘事件
func (m *ICefWindow) SendKeyPress(keyCode int32, eventFlags uint32) {
	imports.Proc(internale_ICEFWindow_SendKeyPress).Call(uintptr(m.instance), uintptr(keyCode), uintptr(eventFlags))
}

// SendMouseMove 发送鼠标移动事件
func (m *ICefWindow) SendMouseMove(screenX, screenY int32) {
	imports.Proc(internale_ICEFWindow_SendMouseMove).Call(uintptr(m.instance), uintptr(screenX), uintptr(screenY))
}

// SendMouseEvents 发送鼠标事件
func (m *ICefWindow) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	imports.Proc(internale_ICEFWindow_SendMouseEvents).Call(uintptr(m.instance), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

// SetAccelerator 设置快捷键
func (m *ICefWindow) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) {
	imports.Proc(internale_ICEFWindow_SetAccelerator).Call(uintptr(m.instance), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
}

// RemoveAccelerator 移除指定快捷键
func (m *ICefWindow) RemoveAccelerator(commandId int32) {
	imports.Proc(internale_ICEFWindow_RemoveAccelerator).Call(uintptr(m.instance), uintptr(commandId))
}

// RemoveAllAccelerators 移除所有快捷键
func (m *ICefWindow) RemoveAllAccelerators() {
	imports.Proc(internale_ICEFWindow_RemoveAllAccelerators).Call(uintptr(m.instance))
}

func (m *ICefWindow) SetWindow(window *ICefWindow) {
	m.instance = window.instance
}
