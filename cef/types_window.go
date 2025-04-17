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
	"errors"
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl/api"
	t "github.com/energye/golcl/lcl/types"
	"io/ioutil"
	"unsafe"
)

// ICefWindow
// include/capi/views/cef_window_capi.h (cef_window_t)
type ICefWindow struct {
	*ICefPanel
}

// CefWindowRef -> ICefWindow
var CefWindowRef cefWindow

type cefWindow uintptr

func (*cefWindow) New(windowComponent *TCEFWindowComponent) *ICefWindow {
	if !windowComponent.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ICEFWindowRef_CreateTopLevel).Call(windowComponent.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefWindow{&ICefPanel{&ICefView{instance: getInstance(result)}}}
	}
	return nil
}

func (m *ICefWindow) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefWindow) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefWindow) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return true
}

// Show 显示窗口
func (m *ICefWindow) Show() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Show).Call(m.Instance())
}

// ShowAsBrowserModalDialog 显示窗口 浏览器模式对话框
func (m *ICefWindow) ShowAsBrowserModalDialog(browserView *ICefBrowserView) {
	if !m.IsValid() || !browserView.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_ShowAsBrowserModalDialog).Call(m.Instance(), browserView.Instance())
}

// Hide 显示窗口
func (m *ICefWindow) Hide() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Hide).Call(m.Instance())
}

// CenterWindow 根据大小窗口居中
func (m *ICefWindow) CenterWindow(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_CenterWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

// Close 关闭窗口， 主窗口调用
func (m *ICefWindow) Close() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Close).Call(m.Instance())
}

// IsClosed 是否关闭
func (m *ICefWindow) IsClosed() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_IsClosed).Call(m.Instance())
	return api.GoBool(r1)
}

// Activate 激活窗口
func (m *ICefWindow) Activate() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Activate).Call(m.Instance())
}

// Deactivate 停止激活窗口
func (m *ICefWindow) Deactivate() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Deactivate).Call(m.Instance())
}

// IsActive 是否激活
func (m *ICefWindow) IsActive() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_IsActive).Call(m.Instance())
	return api.GoBool(r1)
}

// BringToTop 将窗口移至最上层
func (m *ICefWindow) BringToTop() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_BringToTop).Call(m.Instance())
}

// SetAlwaysOnTop 设置窗口是否置顶
func (m *ICefWindow) SetAlwaysOnTop(onTop bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetAlwaysOnTop).Call(m.Instance(), api.PascalBool(onTop))
}

// IsAlwaysOnTop 窗口是否置顶
func (m *ICefWindow) IsAlwaysOnTop() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_IsAlwaysOnTop).Call(m.Instance())
	return api.GoBool(r1)
}

// WindowState 返回窗口最小化、最大化、全屏状态
func (m *ICefWindow) WindowState() t.TWindowState {
	if !m.IsValid() {
		return -1
	}
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
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Maximize).Call(m.Instance())
}

// Minimize 最小化窗口
func (m *ICefWindow) Minimize() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Minimize).Call(m.Instance())
}

// Restore 窗口还原
func (m *ICefWindow) Restore() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_Restore).Call(m.Instance())
}

// SetFullscreen 设置窗口全屏
func (m *ICefWindow) SetFullscreen(fullscreen bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetFullscreen).Call(m.Instance(), api.PascalBool(fullscreen))
}

// SetBackgroundColor 设置背景色
func (m *ICefWindow) SetBackgroundColor(rect types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetBackgroundColor).Call(m.Instance(), rect.ToPtr())
}

// SetBounds 设置窗口边界
func (m *ICefWindow) SetBounds(rect TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&rect)))
}

// SetSize 设置窗口宽高
func (m *ICefWindow) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

// SetPosition 设置窗口位置
func (m *ICefWindow) SetPosition(point TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
}

// IsMaximized 是否最大化
func (m *ICefWindow) IsMaximized() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_IsMaximized).Call(m.Instance())
	return api.GoBool(r1)
}

// IsMinimized 是否最小化
func (m *ICefWindow) IsMinimized() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_IsMinimized).Call(m.Instance())
	return api.GoBool(r1)
}

// IsFullscreen 是否全屏
func (m *ICefWindow) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_IsFullscreen).Call(m.Instance())
	return api.GoBool(r1)
}

// SetTitle 设置窗口标题
func (m *ICefWindow) SetTitle(title string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetTitle).Call(m.Instance(), api.PascalStr(title))
}

// Title 获取窗口标题
func (m *ICefWindow) Title() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_GetTitle).Call(m.Instance())
	return api.GoStr(r1)
}

// SetWindowIcon 设置窗口图标
func (m *ICefWindow) SetWindowIcon(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("invalid")
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(def.ICEFWindow_SetWindowIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// SetWindowIconFS 设置窗口图标
func (m *ICefWindow) SetWindowIconFS(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("invalid")
	}
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(def.ICEFWindow_SetWindowIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// WindowIcon 获取窗口图标
func (m *ICefWindow) WindowIcon() *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ICEFWindow_GetWindowIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// WindowAppIcon 获取窗口应用图标
func (m *ICefWindow) WindowAppIcon() *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ICEFWindow_GetWindowAppIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// SetWindowAppIcon 设置窗口应用图标
func (m *ICefWindow) SetWindowAppIcon(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("invalid")
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(def.ICEFWindow_SetWindowAppIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// SetWindowAppIconFS 设置窗口应用图标
func (m *ICefWindow) SetWindowAppIconFS(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("invalid")
	}
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(def.ICEFWindow_SetWindowAppIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

// AddOverlayView
func (m *ICefWindow) AddOverlayView(view *ICefView, dockingMode consts.TCefDockingMode, canActivate bool) *ICefOverlayController {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.ICEFWindow_AddOverlayView).Call(m.Instance(), view.Instance(), uintptr(dockingMode), api.PascalBool(canActivate), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefOverlayController{instance: getInstance(result)}
	}
	return nil
}

// ShowMenu 显示菜单
func (m *ICefWindow) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_ShowMenu).Call(m.Instance(), menuModel.Instance(), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

// CancelMenu 取消菜单
func (m *ICefWindow) CancelMenu() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_CancelMenu).Call(m.Instance())
}

// Display
func (m *ICefWindow) Display() *ICefDisplay {
	if !m.IsValid() {
		return nil
	}
	var ret uintptr
	imports.Proc(def.ICEFWindow_GetDisplay).Call(m.Instance(), uintptr(unsafe.Pointer(&ret)))
	return &ICefDisplay{
		instance: unsafe.Pointer(ret),
	}
}

// ClientAreaBoundsInScreen 获取客户端所在指定屏幕位置
func (m *ICefWindow) ClientAreaBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_GetClientAreaBoundsInScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// SetDraggableRegions 设置拖拽区域
func (m *ICefWindow) SetDraggableRegions(regions []TCefDraggableRegion) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetDraggableRegions).Call(m.Instance(), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

// WindowHandle 获取窗口句柄
func (m *ICefWindow) WindowHandle() consts.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_GetWindowHandle).Call(m.Instance())
	return consts.TCefWindowHandle(r1)
}

// SendKeyPress 发送键盘事件
func (m *ICefWindow) SendKeyPress(keyCode int32, eventFlags uint32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SendKeyPress).Call(m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

// SendMouseMove 发送鼠标移动事件
func (m *ICefWindow) SendMouseMove(screenX, screenY int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SendMouseMove).Call(m.Instance(), uintptr(screenX), uintptr(screenY))
}

// SendMouseEvents 发送鼠标事件
func (m *ICefWindow) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SendMouseEvents).Call(m.Instance(), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

// SetAccelerator 设置快捷键
func (m *ICefWindow) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed, highPriority bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetAccelerator).Call(m.Instance(), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed), api.PascalBool(highPriority))
}

// RemoveAccelerator 移除指定快捷键
func (m *ICefWindow) RemoveAccelerator(commandId int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_RemoveAccelerator).Call(m.Instance(), uintptr(commandId))
}

// RemoveAllAccelerators 移除所有快捷键
func (m *ICefWindow) RemoveAllAccelerators() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_RemoveAllAccelerators).Call(m.Instance())
}

// SetThemeColor
//
//	Override a standard theme color or add a custom color associated with
//	|color_id|. See cef_color_ids.h for standard ID values. Recommended usage
//	is as follows:
//	<code>
//	1. Customize the default native/OS theme by calling SetThemeColor before
//	   showing the first Window. When done setting colors call
//	   ICefWindow.ThemeChanged to trigger ICefViewDelegate.OnThemeChanged
//	   notifications.
//	2. Customize the current native/OS or Chrome theme after it changes by
//	   calling SetThemeColor from the ICefWindowDelegate.OnThemeColorsChanged
//	   callback. ICefViewDelegate.OnThemeChanged notifications will then be
//	   triggered automatically.
//	</code>
//	<para>The configured color will be available immediately via
//	ICefView.GetThemeColor and will be applied to each View in this
//	Window's component hierarchy when ICefViewDelegate.OnThemeChanged is
//	called. See OnThemeColorsChanged documentation for additional details.
//	Clients wishing to add custom colors should use |color_id| values >= CEF_ChromeColorsEnd.
func (m *ICefWindow) SetThemeColor(colorId int32, color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_SetThemeColor).Call(m.Instance(), uintptr(colorId), uintptr(color))
}

// ThemeChanged
//
//	Trigger ICefViewDelegate.OnThemeChanged callbacks for each View in
//	this Window's component hierarchy. Unlike a native/OS or Chrome theme
//	change this function does not reset theme colors to standard values and
//	does not result in a call to ICefWindowDelegate.OnThemeColorsChanged.
//	<para>Do not call this function from ICefWindowDelegate.OnThemeColorsChanged
//	or ICefViewDelegate.OnThemeChanged.
func (m *ICefWindow) ThemeChanged() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.ICEFWindow_ThemeChanged).Call(m.Instance())
}

// RuntimeStyle
// Returns the runtime style for this Window (ALLOY or CHROME). See
// TCefRuntimeStyle documentation for details.
func (m *ICefWindow) RuntimeStyle() consts.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.ICEFWindow_RuntimeStyle).Call(m.Instance())
	return consts.TCefRuntimeStyle(r1)
}

func (m *ICefWindow) SetWindow(window *ICefWindow) {
	m.instance = window.instance
}
