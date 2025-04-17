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
	"github.com/cyber-xxm/energy/v2/cef/internal/platform"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"io/ioutil"
	"path/filepath"
	"strings"
	"unsafe"
)

// TCEFWindowComponent 窗口组件
type TCEFWindowComponent struct {
	*TCEFPanelComponent
}

// WindowComponentRef -> TCEFWindowComponent
var WindowComponentRef windowComponent

type windowComponent uintptr

// New 创建一个Window组件
func (*windowComponent) New(AOwner lcl.IComponent) *TCEFWindowComponent {
	r1, _, _ := imports.Proc(def.CEFWindowComponent_Create).Call(lcl.CheckPtr(AOwner))
	return &TCEFWindowComponent{&TCEFPanelComponent{&TCEFViewComponent{
		instance: unsafe.Pointer(r1),
	}}}
}

// CreateTopLevelWindow 创建顶层窗口
func (m *TCEFWindowComponent) CreateTopLevelWindow() {
	imports.Proc(def.CEFWindowComponent_CreateTopLevelWindow).Call(m.Instance())
}

// Show 显示窗口
func (m *TCEFWindowComponent) Show() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Show).Call(m.Instance())
}

// ShowAsBrowserModalDialog
// Show the Window as a browser modal dialog relative to |browser_view|. A
// parent Window must be returned via
// ICefWindowDelegate.OnGetParentWindow and |browser_view| must belong
// to that parent Window. While this Window is visible, |browser_view| will
// be disabled while other controls in the parent Window remain enabled.
// Navigating or destroying the |browser_view| will close this Window
// automatically. Alternately, use show() and return true (1) from
// ICefWindowDelegate.OnIsWindowModalDialog for a window modal dialog
// where all controls in the parent Window are disabled.
func (m *TCEFWindowComponent) ShowAsBrowserModalDialog(browserView *ICefBrowserView) {
	if !m.IsValid() || !browserView.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_ShowAsBrowserModalDialog).Call(m.Instance(), browserView.Instance())
}

// Hide 显示窗口
func (m *TCEFWindowComponent) Hide() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Hide).Call(m.Instance())
}

// CenterWindow 根据大小窗口居中
func (m *TCEFWindowComponent) CenterWindow(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_CenterWindow).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

// Close 关闭窗口， 主窗口调用
func (m *TCEFWindowComponent) Close() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Close).Call(m.Instance())
}

// Activate 激活窗口
func (m *TCEFWindowComponent) Activate() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Activate).Call(m.Instance())
}

// Deactivate 停止激活窗口
func (m *TCEFWindowComponent) Deactivate() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Deactivate).Call(m.Instance())
}

// BringToTop 将窗口移至最上层
func (m *TCEFWindowComponent) BringToTop() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_BringToTop).Call(m.Instance())
}

// Maximize 最大化窗口
func (m *TCEFWindowComponent) Maximize() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Maximize).Call(m.Instance())
}

// Minimize 最小化窗口
func (m *TCEFWindowComponent) Minimize() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Minimize).Call(m.Instance())
}

// Restore 窗口还原
func (m *TCEFWindowComponent) Restore() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Restore).Call(m.Instance())
}

func (m *TCEFWindowComponent) AddOverlayView(view *ICefView, dockingMode consts.TCefDockingMode, canActivate bool) *ICefOverlayController {
	var result uintptr
	imports.Proc(def.CEFWindowComponent_AddOverlayView).Call(m.Instance(), view.Instance(), uintptr(dockingMode), api.PascalBool(canActivate), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefOverlayController{instance: getInstance(result)}
	}
	return nil
}

// ShowMenu 显示菜单
func (m *TCEFWindowComponent) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_ShowMenu).Call(m.Instance(), uintptr(menuModel.instance), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

// CancelMenu 取消菜单
func (m *TCEFWindowComponent) CancelMenu() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_CancelMenu).Call(m.Instance())
}

// SetDraggableRegions 设置拖拽区域
func (m *TCEFWindowComponent) SetDraggableRegions(regions []TCefDraggableRegion) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetDraggableRegions).Call(m.Instance(), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

// SendKeyPress 发送键盘事件
func (m *TCEFWindowComponent) SendKeyPress(keyCode int32, eventFlags uint32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SendKeyPress).Call(m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

// SendMouseMove 发送鼠标移动事件
func (m *TCEFWindowComponent) SendMouseMove(screenX, screenY int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SendMouseMove).Call(m.Instance(), uintptr(screenX), uintptr(screenY))
}

// SendMouseEvents 发送鼠标事件
func (m *TCEFWindowComponent) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SendMouseEvents).Call(m.Instance(), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

// SetAccelerator 设置快捷键
func (m *TCEFWindowComponent) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed, highPriority bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetAccelerator).Call(m.Instance(), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed), api.PascalBool(highPriority))
}

// RemoveAccelerator 移除指定快捷键
func (m *TCEFWindowComponent) RemoveAccelerator(commandId int32) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_RemoveAccelerator).Call(m.Instance(), uintptr(commandId))
}

// RemoveAllAccelerators 移除所有快捷键
func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_RemoveAllAccelerators).Call(m.Instance())
}

// SetThemeColor
// Override a standard theme color or add a custom color associated with
// |color_id|. See cef_color_ids.h for standard ID values. Recommended usage
// is as follows:</para>
// <code>
//  1. Customize the default native/OS theme by calling SetThemeColor before
//     showing the first Window. When done setting colors call
//     ICefWindow.ThemeChanged to trigger ICefViewDelegate.OnThemeChanged
//     notifications.
//  2. Customize the current native/OS or Chrome theme after it changes by
//     calling SetThemeColor from the ICefWindowDelegate.OnThemeColorsChanged
//     callback. ICefViewDelegate.OnThemeChanged notifications will then be
//     triggered automatically.
//
// </code>
// <para>The configured color will be available immediately via
// ICefView.GetThemeColor and will be applied to each View in this
// Window's component hierarchy when ICefViewDelegate.OnThemeChanged is
// called. See OnThemeColorsChanged documentation for additional details.</para>
// <para>Clients wishing to add custom colors should use |color_id| values >=
// CEF_ChromeColorsEnd.
func (m *TCEFWindowComponent) SetThemeColor(colorId int32, color types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetThemeColor).Call(m.Instance(), uintptr(colorId), uintptr(color))
}

// ThemeChanged
// Trigger ICefViewDelegate.OnThemeChanged callbacks for each View in
// this Window's component hierarchy. Unlike a native/OS or Chrome theme
// change this function does not reset theme colors to standard values and
// does not result in a call to ICefWindowDelegate.OnThemeColorsChanged.
// Do not call this function from ICefWindowDelegate.OnThemeColorsChanged
// or ICefViewDelegate.OnThemeChanged.
func (m *TCEFWindowComponent) ThemeChanged() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_ThemeChanged).Call(m.Instance())
}

// SetAlwaysOnTop 设置窗口是否置顶
func (m *TCEFWindowComponent) SetAlwaysOnTop(onTop bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetAlwaysOnTop).Call(m.Instance(), api.PascalBool(onTop))
}

// SetFullscreen 设置窗口全屏
func (m *TCEFWindowComponent) SetFullscreen(fullscreen bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetFullscreen).Call(m.Instance(), api.PascalBool(fullscreen))
}

// SetBackgroundColor 设置背景色
func (m *TCEFWindowComponent) SetBackgroundColor(rect types.TCefColor) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetBackgroundColor).Call(m.Instance(), uintptr(rect))
}

// Bounds 获取窗口边界
func (m *TCEFWindowComponent) Bounds() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Bounds).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// Size 获取窗口宽高
func (m *TCEFWindowComponent) Size() (result TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Size).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// Position 获取窗口位置
func (m *TCEFWindowComponent) Position() (result TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_Position).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// SetBounds 设置窗口边界
func (m *TCEFWindowComponent) SetBounds(rect TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(&rect)))
}

// SetSize 设置窗口宽高
func (m *TCEFWindowComponent) SetSize(size TCefSize) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetSize).Call(m.Instance(), uintptr(unsafe.Pointer(&size)))
}

// SetPosition 设置窗口位置
func (m *TCEFWindowComponent) SetPosition(point TCefPoint) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(&point)))
}

// SetTitle 设置窗口标题
func (m *TCEFWindowComponent) SetTitle(title string) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetTitle).Call(m.Instance(), api.PascalStr(title))
}

// Title 获取窗口标题
func (m *TCEFWindowComponent) Title() string {
	if !m.IsValid() {
		return ""
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_Title).Call(m.Instance())
	return api.GoStr(r1)
}

// WindowIcon 获取窗口图标
func (m *TCEFWindowComponent) WindowIcon() *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFWindowComponent_WindowIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// WindowAppIcon 获取窗口应用图标
func (m *TCEFWindowComponent) WindowAppIcon() *ICefImage {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFWindowComponent_WindowAppIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

func (m *TCEFWindowComponent) SetWindowIcon(icon *ICefImage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetWindowIcon).Call(m.Instance(), icon.Instance())
}

func (m *TCEFWindowComponent) checkICON(filename string) (string, error) {
	if !m.IsValid() {
		return "", errors.New("window component is nil")
	}
	ext := strings.ToLower(filepath.Ext(filename))
	if ".png" != ext && ".jpeg" != ext {
		return "", errors.New("only png and jpeg image formats are supported")
	}
	return ext[1:], nil
}

// SetWindowIconByFile 设置窗口图标
func (m *TCEFWindowComponent) SetWindowIconByFile(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("window component is nil")
	}
	var (
		ext string
		err error
	)
	if ext, err = m.checkICON(filename); err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	icon := ImageRef.New()
	if "png" == ext {
		icon.AddPng(scaleFactor, bytes)
	} else if "jpeg" == ext {
		icon.AddJpeg(scaleFactor, bytes)
	}
	m.SetWindowIcon(icon)
	return nil
}

// SetWindowIconByFSFile 设置窗口图标
func (m *TCEFWindowComponent) SetWindowIconByFSFile(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("window component is nil")
	}
	var (
		ext string
		err error
	)
	if ext, err = m.checkICON(filename); err != nil {
		return err
	}
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	icon := ImageRef.New()
	if "png" == ext {
		icon.AddPng(scaleFactor, bytes)
	} else if "jpeg" == ext {
		icon.AddJpeg(scaleFactor, bytes)
	}
	m.SetWindowIcon(icon)
	return nil
}

func (m *TCEFWindowComponent) SetWindowAppIcon(icon *ICefImage) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_SetWindowAppIcon).Call(m.Instance(), icon.Instance())
}

// SetWindowAppIconByFile 设置窗口应用图标
func (m *TCEFWindowComponent) SetWindowAppIconByFile(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("window component is nil")
	}
	var (
		ext string
		err error
	)
	if ext, err = m.checkICON(filename); err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	icon := ImageRef.New()
	if "png" == ext {
		icon.AddPng(scaleFactor, bytes)
	} else if "jpeg" == ext {
		icon.AddJpeg(scaleFactor, bytes)
	}
	m.SetWindowAppIcon(icon)
	return nil
}

// SetWindowAppIconByFSFile 设置窗口应用图标
func (m *TCEFWindowComponent) SetWindowAppIconByFSFile(scaleFactor float32, filename string) error {
	if !m.IsValid() {
		return errors.New("window component is nil")
	}
	var (
		ext string
		err error
	)
	if ext, err = m.checkICON(filename); err != nil {
		return err
	}
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	icon := ImageRef.New()
	if "png" == ext {
		icon.AddPng(scaleFactor, bytes)
	} else if "jpeg" == ext {
		icon.AddJpeg(scaleFactor, bytes)
	}
	m.SetWindowAppIcon(icon)
	return nil
}

// Display
func (m *TCEFWindowComponent) Display() *ICefDisplay {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFWindowComponent_Display).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefDisplay{
		instance: unsafe.Pointer(result),
	}
}

// ClientAreaBoundsInScreen 获取客户端所在指定屏幕位置
func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (result TCefRect) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_ClientAreaBoundsInScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// WindowHandle 获取窗口句柄
func (m *TCEFWindowComponent) WindowHandle() consts.TCefWindowHandle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_WindowHandle).Call(m.Instance())
	return consts.TCefWindowHandle(r1)
}

// IsClosed 是否关闭
func (m *TCEFWindowComponent) IsClosed() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_IsClosed).Call(m.Instance())
	return api.GoBool(r1)
}

// IsActive 是否激活
func (m *TCEFWindowComponent) IsActive() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_IsActive).Call(m.Instance())
	return api.GoBool(r1)
}

// IsAlwaysOnTop 窗口是否置顶
func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_IsAlwaysOnTop).Call(m.Instance())
	return api.GoBool(r1)
}

// IsFullscreen 是否全屏
func (m *TCEFWindowComponent) IsFullscreen() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_IsFullscreen).Call(m.Instance())
	return api.GoBool(r1)
}

// IsMaximized 是否最大化
func (m *TCEFWindowComponent) IsMaximized() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_IsMaximized).Call(m.Instance())
	return api.GoBool(r1)
}

// IsMinimized 是否最小化
func (m *TCEFWindowComponent) IsMinimized() bool {
	if !m.IsValid() {
		return false
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_IsMinimized).Call(m.Instance())
	return api.GoBool(r1)
}

// AddChildView 添加浏览器显示组件
func (m *TCEFWindowComponent) AddChildView(view *ICefView) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFWindowComponent_AddChildView).Call(m.Instance(), view.Instance())
}

// RuntimeStyle
// Returns the runtime style for this Window (ALLOY or CHROME). See
// TCefRuntimeStyle documentation for details.
func (m *TCEFWindowComponent) RuntimeStyle() consts.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFWindowComponent_RuntimeStyle).Call(m.Instance())
	return consts.TCefRuntimeStyle(r1)
}

// SetLinuxWindowProperties Linux-specific window properties for correctly handling by window managers.
// Main window's WM_CLASS_CLASS in X11
// Main window's WM_CLASS_NAME in X11
func (m *TCEFWindowComponent) SetLinuxWindowProperties(wmClassName, wmClassClass string) {
	handle := uintptr(m.WindowHandle())
	if handle != 0 {
		platform.SetWMClass(wmClassName, wmClassClass, handle)
	}
}

// Called when |window| is created.
func (m *TCEFWindowComponent) SetOnWindowCreated(fn windowOnWindowCreated) {
	imports.Proc(def.CEFWindowComponent_SetOnWindowCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |window| is closing.
func (m *TCEFWindowComponent) SetOnWindowClosing(fn windowOnWindowClosing) {
	imports.Proc(def.CEFWindowComponent_SetOnWindowClosing).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |window| is destroyed. Release all references to |window| and
// do not attempt to execute any functions on |window| after this callbackreturns.
func (m *TCEFWindowComponent) SetOnWindowDestroyed(fn windowOnWindowDestroyed) {
	imports.Proc(def.CEFWindowComponent_SetOnWindowDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |window| is activated or deactivated.
func (m *TCEFWindowComponent) SetOnWindowActivationChanged(fn windowOnWindowActivationChanged) {
	imports.Proc(def.CEFWindowComponent_SetOnWindowActivationChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |window| bounds have changed. |new_bounds| will be in DIP
// screen coordinates.
func (m *TCEFWindowComponent) SetOnWindowBoundsChanged(fn windowOnWindowBoundsChanged) {
	imports.Proc(def.CEFWindowComponent_SetOnWindowBoundsChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return the parent for |window| or NULL if the |window| does not have a
// parent. Windows with parents will not get a taskbar button. Set |is_menu|
// to true (1) if |window| will be displayed as a menu, in which case it will
// not be clipped to the parent window bounds. Set |can_activate_menu| to
// false (0) if |is_menu| is true (1) and |window| should not be activated
// (given keyboard focus) when displayed.
func (m *TCEFWindowComponent) SetOnGetParentWindow(fn windowOnGetParentWindow) {
	imports.Proc(def.CEFWindowComponent_SetOnGetParentWindow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| should be created as a window modal dialog.
// Only called when a Window is returned via get_parent_window() with
// |is_menu| set to false (0). All controls in the parent Window will be
// disabled while |window| is visible. This functionality is not supported by
// all Linux window managers. Alternately, use
// ICefWindow.ShowAsBrowserModalDialog() for a browser modal dialog
// that works on all platforms.
func (m *TCEFWindowComponent) SetOnIsWindowModalDialog(fn windowOnIsWindowModalDialog) {
	imports.Proc(def.CEFWindowComponent_SetOnIsWindowModalDialog).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return the initial bounds for |window| in density independent pixel (DIP)
// coordinates. If this function returns an NULL CefRect then
// GetPreferredSize() will be called to retrieve the size, and the window
// will be placed on the screen with origin (0,0). This function can be used
// in combination with ICefView.GetBoundsInScreen() to restore the
// previous window bounds.
func (m *TCEFWindowComponent) SetOnGetInitialBounds(fn windowOnGetInitialBounds) {
	imports.Proc(def.CEFWindowComponent_SetOnGetInitialBounds).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return the initial show state for |window|.
func (m *TCEFWindowComponent) SetOnGetInitialShowState(fn windowOnGetInitialShowState) {
	imports.Proc(def.CEFWindowComponent_SetOnGetInitialShowState).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| should be created without a frame or title
// bar. The window will be resizable if can_resize() returns true (1). Use
// ICefWindow.SetDraggableRegions() to specify draggable regions.
func (m *TCEFWindowComponent) SetOnIsFrameless(fn windowOnIsFrameless) {
	imports.Proc(def.CEFWindowComponent_SetOnIsFrameless).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| should be created with standard window buttons
// like close, minimize and zoom. This function is only supported on macOS.
func (m *TCEFWindowComponent) SetOnWithStandardWindowButtons(fn windowOnWithStandardWindowButtons) {
	imports.Proc(def.CEFWindowComponent_SetOnWithStandardWindowButtons).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return whether the titlebar height should be overridden, and sets the
// height of the titlebar in |titlebar_height|. On macOS, it can also be used
// to adjust the vertical position of the traffic light buttons in frameless
// windows. The buttons will be positioned halfway down the titlebar at a
// height of |titlebar_height| / 2.
func (m *TCEFWindowComponent) SetOnGetTitleBarHeight(fn windowOnGetTitleBarHeight) {
	imports.Proc(def.CEFWindowComponent_SetOnGetTitlebarHeight).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// <para>Return whether the view should accept the initial mouse-down event,
// allowing it to respond to click-through behavior. If STATE_ENABLED is
// returned, the view will be sent a mouseDown: message for an initial mouse-
// down event, activating the view with one click, instead of clicking first
// to make the window active and then clicking the view.</para>
// <para>This function is only supported on macOS. For more details, refer to the
// documentation of acceptsFirstMouse.</para>
func (m *TCEFWindowComponent) SetOnAcceptsFirstMouse(fn windowOnAcceptsFirstMouse) {
	imports.Proc(def.CEFWindowComponent_SetOnAcceptsFirstMouse).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| can be resized.
func (m *TCEFWindowComponent) SetOnCanResize(fn windowOnCanResize) {
	imports.Proc(def.CEFWindowComponent_SetOnCanResize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| can be maximized.
func (m *TCEFWindowComponent) SetOnCanMaximize(fn windowOnCanMaximize) {
	imports.Proc(def.CEFWindowComponent_SetOnCanMaximize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| can be minimized.
func (m *TCEFWindowComponent) SetOnCanMinimize(fn windowOnCanMinimize) {
	imports.Proc(def.CEFWindowComponent_SetOnCanMinimize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return true (1) if |window| can be closed. This will be called for user-
// initiated window close actions and when ICefWindow.close() is called.
func (m *TCEFWindowComponent) SetOnCanClose(fn windowOnCanClose) {
	imports.Proc(def.CEFWindowComponent_SetOnCanClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when a keyboard accelerator registered with
// ICefWindow.SetAccelerator is triggered. Return true (1) if the
// accelerator was handled or false (0) otherwise.
func (m *TCEFWindowComponent) SetOnAccelerator(fn windowOnAccelerator) {
	imports.Proc(def.CEFWindowComponent_SetOnAccelerator).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called after all other controls in the window have had a chance to handle
// the event. |event| contains information about the keyboard event. Return
// true (1) if the keyboard event was handled or false (0) otherwise.
func (m *TCEFWindowComponent) SetOnKeyEvent(fn windowOnKey) {
	imports.Proc(def.CEFWindowComponent_SetOnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when |window| is transitioning to or from fullscreen mode. On MacOS
// the transition occurs asynchronously with |is_competed| set to false (0)
// when the transition starts and true (1) after the transition completes. On
// other platforms the transition occurs synchronously with |is_completed|
// set to true (1) after the transition completes. With Alloy style you must
// also implement ICefDisplayHandler.OnFullscreenModeChange to handle
// fullscreen transitions initiated by browser content.
func (m *TCEFWindowComponent) SetOnWindowFullscreenTransition(fn windowOnWindowFullscreenTransition) {
	imports.Proc(def.CEFWindowComponent_SetOnWindowFullscreenTransition).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// <para>Called after the native/OS or Chrome theme for |window| has changed.
// |chrome_theme| will be true (1) if the notification is for a Chrome theme.</para>
// <para>Native/OS theme colors are configured globally and do not need to be
// customized for each Window individually. An example of a native/OS theme
// change that triggers this callback is when the user switches between dark
// and light mode during application lifespan. Native/OS theme changes can be
// disabled by passing the `--force-dark-mode` or `--force-light-mode`
// command-line flag.</para>
// <para>Chrome theme colors will be applied and this callback will be triggered
// if/when a BrowserView is added to the Window's component hierarchy. Chrome
// theme colors can be configured on a per-RequestContext basis using
// ICefRequestContext.SetChromeColorScheme or (Chrome style only) by
// visiting chrome://settings/manageProfile. Any theme changes using those
// mechanisms will also trigger this callback. Chrome theme colors will be
// persisted and restored from disk cache.</para>
// <para>This callback is not triggered on Window creation so clients that wish to
// customize the initial native/OS theme must call
// ICefWindow.SetThemeColor and ICefWindow.ThemeChanged before showing
// the first Window.</para>
// <para>Theme colors will be reset to standard values before this callback is
// called for the first affected Window. Call ICefWindow.SetThemeColor
// from inside this callback to override a standard color or add a custom
// color. ICefViewDelegate.OnThemeChanged will be called after this
// callback for the complete |window| component hierarchy.</para>
func (m *TCEFWindowComponent) SetOnThemeColorsChanged(fn windowOnThemeColorsChanged) {
	imports.Proc(def.CEFWindowComponent_SetOnThemeColorsChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Optionally change the runtime style for this Window. See
// TCefRuntimeStyle documentation for details.
func (m *TCEFWindowComponent) SetOnGetWindowRuntimeStyle(fn windowOnGetWindowRuntimeStyle) {
	imports.Proc(def.CEFWindowComponent_SetOnGetWindowRuntimeStyle).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Return Linux-specific window properties for correctly handling by window  managers.
func (m *TCEFWindowComponent) SetOnGetLinuxWindowProperties(fn windowOnGetLinuxWindowProperties) {
	imports.Proc(def.CEFWindowComponent_SetOnGetLinuxWindowProperties).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
