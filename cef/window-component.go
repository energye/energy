//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// energy VF窗口组件
package cef

import (
	"errors"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/types"
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
	lcl.IComponent
	instance unsafe.Pointer
}

// NewWindowComponent 创建一个Window组件
func NewWindowComponent(AOwner lcl.IComponent) *TCEFWindowComponent {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_Create).Call(lcl.CheckPtr(AOwner))
	return &TCEFWindowComponent{
		instance: unsafe.Pointer(r1),
	}
}

// Instance 实例
func (m *TCEFWindowComponent) Instance() uintptr {
	return uintptr(m.instance)
}

// CreateTopLevelWindow 创建顶层窗口
func (m *TCEFWindowComponent) CreateTopLevelWindow() {
	imports.Proc(internale_CEFWindowComponent_CreateTopLevelWindow).Call(m.Instance())
}

// Show 显示窗口
func (m *TCEFWindowComponent) Show() {
	imports.Proc(internale_CEFWindowComponent_Show).Call(m.Instance())
}

// Hide 显示窗口
func (m *TCEFWindowComponent) Hide() {
	imports.Proc(internale_CEFWindowComponent_Hide).Call(m.Instance())
}

// CenterWindow 根据大小窗口居中
func (m *TCEFWindowComponent) CenterWindow(size *TCefSize) {
	imports.Proc(internale_CEFWindowComponent_CenterWindow).Call(m.Instance(), uintptr(unsafe.Pointer(size)))
}

// Close 关闭窗口， 主窗口调用
func (m *TCEFWindowComponent) Close() {
	imports.Proc(internale_CEFWindowComponent_Close).Call(m.Instance())
}

// Activate 激活窗口
func (m *TCEFWindowComponent) Activate() {
	imports.Proc(internale_CEFWindowComponent_Activate).Call(m.Instance())
}

// Deactivate 停止激活窗口
func (m *TCEFWindowComponent) Deactivate() {
	imports.Proc(internale_CEFWindowComponent_Deactivate).Call(m.Instance())
}

// BringToTop 将窗口移至最上层
func (m *TCEFWindowComponent) BringToTop() {
	imports.Proc(internale_CEFWindowComponent_BringToTop).Call(m.Instance())
}

// Maximize 最大化窗口
func (m *TCEFWindowComponent) Maximize() {
	imports.Proc(internale_CEFWindowComponent_Maximize).Call(m.Instance())
}

// Minimize 最小化窗口
func (m *TCEFWindowComponent) Minimize() {
	imports.Proc(internale_CEFWindowComponent_Minimize).Call(m.Instance())
}

// Restore 窗口还原
func (m *TCEFWindowComponent) Restore() {
	imports.Proc(internale_CEFWindowComponent_Restore).Call(m.Instance())
}

//func (m *TCEFWindowComponent) AddOverlayView() {
//imports.Proc(internale_CEFWindowComponent_AddOverlayView).Call(m.Instance())
//}

// ShowMenu 显示菜单
func (m *TCEFWindowComponent) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	imports.Proc(internale_CEFWindowComponent_ShowMenu).Call(m.Instance(), uintptr(menuModel.instance), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

// CancelMenu 取消菜单
func (m *TCEFWindowComponent) CancelMenu() {
	imports.Proc(internale_CEFWindowComponent_CancelMenu).Call(m.Instance())
}

// SetDraggableRegions 设置拖拽区域
func (m *TCEFWindowComponent) SetDraggableRegions(regions []TCefDraggableRegion) {
	imports.Proc(internale_CEFWindowComponent_SetDraggableRegions).Call(m.Instance(), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

// SendKeyPress 发送键盘事件
func (m *TCEFWindowComponent) SendKeyPress(keyCode int32, eventFlags uint32) {
	imports.Proc(internale_CEFWindowComponent_SendKeyPress).Call(m.Instance(), uintptr(keyCode), uintptr(eventFlags))
}

// SendMouseMove 发送鼠标移动事件
func (m *TCEFWindowComponent) SendMouseMove(screenX, screenY int32) {
	imports.Proc(internale_CEFWindowComponent_SendMouseMove).Call(m.Instance(), uintptr(screenX), uintptr(screenY))
}

// SendMouseEvents 发送鼠标事件
func (m *TCEFWindowComponent) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	imports.Proc(internale_CEFWindowComponent_SendMouseEvents).Call(m.Instance(), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

// SetAccelerator 设置快捷键
func (m *TCEFWindowComponent) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) {
	imports.Proc(internale_CEFWindowComponent_SetAccelerator).Call(m.Instance(), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
}

// RemoveAccelerator 移除指定快捷键
func (m *TCEFWindowComponent) RemoveAccelerator(commandId int32) {
	imports.Proc(internale_CEFWindowComponent_RemoveAccelerator).Call(m.Instance(), uintptr(commandId))
}

// RemoveAllAccelerators 移除所有快捷键
func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	imports.Proc(internale_CEFWindowComponent_RemoveAllAccelerators).Call(m.Instance())
}

// SetAlwaysOnTop 设置窗口是否置顶
func (m *TCEFWindowComponent) SetAlwaysOnTop(onTop bool) {
	imports.Proc(internale_CEFWindowComponent_SetAlwaysOnTop).Call(m.Instance(), api.PascalBool(onTop))
}

// SetFullscreen 设置窗口全屏
func (m *TCEFWindowComponent) SetFullscreen(fullscreen bool) {
	imports.Proc(internale_CEFWindowComponent_SetFullscreen).Call(m.Instance(), api.PascalBool(fullscreen))
}

// SetBackgroundColor 设置背景色
func (m *TCEFWindowComponent) SetBackgroundColor(rect *types.TCefColor) {
	imports.Proc(internale_CEFWindowComponent_SetBackgroundColor).Call(m.Instance(), rect.ToPtr())
}

// SetBounds 获取窗口边界
func (m *TCEFWindowComponent) Bounds() (result *TCefRect) {
	imports.Proc(internale_CEFWindowComponent_Bounds).Call(m.Instance(), uintptr(unsafe.Pointer(result)))
	return
}

// SetSize 获取窗口宽高
func (m *TCEFWindowComponent) Size() (result *TCefSize) {
	imports.Proc(internale_CEFWindowComponent_Size).Call(m.Instance(), uintptr(unsafe.Pointer(result)))
	return
}

// Position 获取窗口位置
func (m *TCEFWindowComponent) Position() (result *TCefPoint) {
	imports.Proc(internale_CEFWindowComponent_Position).Call(m.Instance(), uintptr(unsafe.Pointer(result)))
	return
}

// SetBounds 设置窗口边界
func (m *TCEFWindowComponent) SetBounds(rect *TCefRect) {
	imports.Proc(internale_CEFWindowComponent_SetBounds).Call(m.Instance(), uintptr(unsafe.Pointer(rect)))
}

// SetSize 设置窗口宽高
func (m *TCEFWindowComponent) SetSize(size *TCefSize) {
	imports.Proc(internale_CEFWindowComponent_SetSize).Call(m.Instance(), uintptr(unsafe.Pointer(size)))
}

// SetPosition 设置窗口位置
func (m *TCEFWindowComponent) SetPosition(point *TCefPoint) {
	imports.Proc(internale_CEFWindowComponent_SetPosition).Call(m.Instance(), uintptr(unsafe.Pointer(point)))
}

// SetTitle 设置窗口标题
func (m *TCEFWindowComponent) SetTitle(title string) {
	imports.Proc(internale_CEFWindowComponent_SetTitle).Call(m.Instance(), api.PascalStr(title))
}

// GetTitle 获取窗口标题
func (m *TCEFWindowComponent) Title() string {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_Title).Call(m.Instance())
	return api.GoStr(r1)
}

// WindowIcon 获取窗口图标
func (m *TCEFWindowComponent) WindowIcon() *ICefImage {
	var result uintptr
	imports.Proc(internale_CEFWindowComponent_WindowIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

// WindowAppIcon 获取窗口应用图标
func (m *TCEFWindowComponent) WindowAppIcon() *ICefImage {
	var result uintptr
	imports.Proc(internale_CEFWindowComponent_WindowAppIcon).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefImage{
		instance: unsafe.Pointer(result),
	}
}

func (m *TCEFWindowComponent) SetWindowIcon(icon *ICefImage) {
	imports.Proc(internale_CEFWindowComponent_SetWindowIcon).Call(m.Instance(), icon.Instance())
}

func (m *TCEFWindowComponent) checkICON(filename string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	if ".png" != ext && ".jpeg" != ext {
		return "", errors.New("only png and jpeg image formats are supported")
	}
	ext = ext[1:]
	return ext, nil
}

// SetWindowIconByFile 设置窗口图标
func (m *TCEFWindowComponent) SetWindowIconByFile(scaleFactor float32, filename string) error {
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
	imports.Proc(internale_CEFWindowComponent_SetWindowAppIcon).Call(m.Instance(), icon.Instance())
}

// SetWindowAppIconByFile 设置窗口应用图标
func (m *TCEFWindowComponent) SetWindowAppIconByFile(scaleFactor float32, filename string) error {
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
	var result uintptr
	imports.Proc(internale_CEFWindowComponent_Display).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefDisplay{
		instance: unsafe.Pointer(result),
	}
}

// ClientAreaBoundsInScreen 获取客户端所在指定屏幕位置
func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (result TCefRect) {
	imports.Proc(internale_CEFWindowComponent_ClientAreaBoundsInScreen).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// WindowHandle 获取窗口句柄
func (m *TCEFWindowComponent) WindowHandle() consts.TCefWindowHandle {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_WindowHandle).Call(m.Instance())
	return consts.TCefWindowHandle(r1)
}

// IsClosed 是否关闭
func (m *TCEFWindowComponent) IsClosed() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsClosed).Call(m.Instance())
	return api.GoBool(r1)
}

// IsActive 是否激活
func (m *TCEFWindowComponent) IsActive() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsActive).Call(m.Instance())
	return api.GoBool(r1)
}

// IsAlwaysOnTop 窗口是否置顶
func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsAlwaysOnTop).Call(m.Instance())
	return api.GoBool(r1)
}

// IsFullscreen 是否全屏
func (m *TCEFWindowComponent) IsFullscreen() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsFullscreen).Call(m.Instance())
	return api.GoBool(r1)
}

// IsMaximized 是否最大化
func (m *TCEFWindowComponent) IsMaximized() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsMaximized).Call(m.Instance())
	return api.GoBool(r1)
}

// IsMinimized 是否最小化
func (m *TCEFWindowComponent) IsMinimized() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsMinimized).Call(m.Instance())
	return api.GoBool(r1)
}

// AddChildView 添加浏览器显示组件
func (m *TCEFWindowComponent) AddChildView(browserViewComponent *TCEFBrowserViewComponent) {
	imports.Proc(internale_CEFWindowComponent_AddChildView).Call(m.Instance(), browserViewComponent.Instance())
}

// SetOnWindowCreated 窗口创建回调事件
func (m *TCEFWindowComponent) SetOnWindowCreated(fn WindowComponentOnWindowCreated) {
	imports.Proc(internale_CEFWindowComponent_SetOnWindowCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnWindowDestroyed 窗口销毁回调事件
func (m *TCEFWindowComponent) SetOnWindowDestroyed(fn WindowComponentOnWindowDestroyed) {
	imports.Proc(internale_CEFWindowComponent_SetOnWindowDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnWindowActivationChanged 窗口激活改变回调事件
func (m *TCEFWindowComponent) SetOnWindowActivationChanged(fn WindowComponentOnWindowActivationChanged) {
	imports.Proc(internale_CEFWindowComponent_SetOnWindowActivationChanged).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetParentWindow 获取父组件回调事件
func (m *TCEFWindowComponent) SetOnGetParentWindow(fn WindowComponentOnGetParentWindow) {
	imports.Proc(internale_CEFWindowComponent_SetOnGetParentWindow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetInitialBounds 窗口初始窗口边界回调事件
func (m *TCEFWindowComponent) SetOnGetInitialBounds(fn WindowComponentOnGetInitialBounds) {
	imports.Proc(internale_CEFWindowComponent_SetOnGetInitialBounds).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetInitialShowState 窗口初始显示状态回调事件
func (m *TCEFWindowComponent) SetOnGetInitialShowState(fn WindowComponentOnGetInitialShowState) {
	imports.Proc(internale_CEFWindowComponent_SetOnGetInitialShowState).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnIsFrameless 窗口无边框回调事件
func (m *TCEFWindowComponent) SetOnIsFrameless(fn WindowComponentOnIsFrameless) {
	imports.Proc(internale_CEFWindowComponent_SetOnIsFrameless).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnCanResize 设置窗口是否允许调整大小回调事件
func (m *TCEFWindowComponent) SetOnCanResize(fn WindowComponentOnCanResize) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanResize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnCanMaximize 设置窗口是否允许最大化回调事件
func (m *TCEFWindowComponent) SetOnCanMaximize(fn WindowComponentOnCanMaximize) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanMaximize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnCanMinimize 设置窗口是否允许最小化回调事件
func (m *TCEFWindowComponent) SetOnCanMinimize(fn WindowComponentOnCanMinimize) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanMinimize).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnCanClose 设置窗口是否允许关闭回调事件
func (m *TCEFWindowComponent) SetOnCanClose(fn WindowComponentOnCanClose) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanClose).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnAccelerator 设置快捷键回调事件
func (m *TCEFWindowComponent) SetOnAccelerator(fn WindowComponentOnAccelerator) {
	imports.Proc(internale_CEFWindowComponent_SetOnAccelerator).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnKeyEvent 设置键盘事件回调事件
func (m *TCEFWindowComponent) SetOnKeyEvent(fn WindowComponentOnKeyEvent) {
	imports.Proc(internale_CEFWindowComponent_SetOnKeyEvent).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case WindowComponentOnWindowCreated:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnWindowCreated)(lcl.AsObject(sender), &ICefWindow{instance: window})
		case WindowComponentOnWindowDestroyed:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnWindowDestroyed)(lcl.AsObject(sender), &ICefWindow{instance: window})
		case WindowComponentOnWindowActivationChanged:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnWindowActivationChanged)(lcl.AsObject(sender), &ICefWindow{instance: window}, api.GoBool(getVal(2)))
		case WindowComponentOnGetParentWindow:
			sender := getPtr(0)
			window := getPtr(1)
			resultWindowPtr := (*uintptr)(getPtr(4))
			resultWindow := &ICefWindow{}
			fn.(WindowComponentOnGetParentWindow)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)), (*bool)(getPtr(3)), resultWindow)
			*resultWindowPtr = uintptr(resultWindow.instance)
		case WindowComponentOnGetInitialBounds:
			sender := getPtr(0)
			window := getPtr(1)
			resultRectPtr := (*TCefRect)(getPtr(2))
			resultRect := new(TCefRect)
			resultRect.X = 0
			resultRect.Y = 0
			resultRect.Width = 800
			resultRect.Height = 600
			fn.(WindowComponentOnGetInitialBounds)(lcl.AsObject(sender), &ICefWindow{instance: window}, resultRect)
			*resultRectPtr = *resultRect
		case WindowComponentOnGetInitialShowState:
			sender := getPtr(0)
			window := getPtr(1)
			resultShowState := (*consts.TCefShowState)(getPtr(2))
			fn.(WindowComponentOnGetInitialShowState)(lcl.AsObject(sender), &ICefWindow{instance: window}, resultShowState)
		case WindowComponentOnIsFrameless:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnIsFrameless)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanResize:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanResize)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanMaximize:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanMaximize)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanMinimize:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanMinimize)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanClose:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanClose)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnAccelerator:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnAccelerator)(lcl.AsObject(sender), &ICefWindow{instance: window}, int32(getVal(2)), (*bool)(getPtr(3)))
		case WindowComponentOnKeyEvent:
			sender := getPtr(0)
			window := getPtr(1)
			keyEvent := (*TCefKeyEvent)(getPtr(2))
			fn.(WindowComponentOnKeyEvent)(lcl.AsObject(sender), &ICefWindow{instance: window}, keyEvent, (*bool)(getPtr(3)))
		default:
			return false
		}
		return true
	})
}
