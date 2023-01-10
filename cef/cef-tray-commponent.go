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
)

type TMouseEvent func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) bool

type ITray interface {
	SetTitle(title string)                                 //设置标题
	SetVisible(v bool)                                     //显示和隐藏托盘图标
	Visible() bool                                         //
	Show()                                                 //显示托盘菜单窗口 windows有效
	Hide()                                                 //隐藏托盘菜单窗口 windows有效
	close()                                                //关闭托盘菜单窗口 windows有效
	SetOnDblClick(fn lcl.TNotifyEvent)                     //双击事件 linux 和 macos 可能不启作用
	SetOnClick(fn lcl.TNotifyEvent)                        //单击事件
	SetOnMouseUp(fn TMouseEvent)                           //up事件 linux 和 macos 可能不启作用
	SetOnMouseDown(fn lcl.TMouseEvent)                     //down事件 linux 和 macos 可能不启作用
	SetOnMouseMove(fn lcl.TMouseMoveEvent)                 //move事件 linux 和 macos 可能不启作用
	SetIconFS(iconResourcePath string)                     //设置托盘图标
	SetIcon(iconResourcePath string)                       //设置托盘图标
	SetHint(value string)                                  //设置托盘hint(鼠标移动到托盘图标显示的文字)
	ShowBalloon()                                          //显示托盘气泡
	SetBalloon(title, content string, timeout int32) ITray //设置托盘气泡内容
	Tray() ITray                                           //获得 LCLTray, ( CefTray 返回 nil )
	AsViewsFrameTray() *ViewsFrameTray
	AsCEFTray() *CEFTray
	AsLCLTray() *LCLTray
}

//LCL 系统托盘
type LCLTray struct {
	owner     lcl.IComponent
	trayIcon  *lcl.TTrayIcon
	popupMenu *lcl.TPopupMenu
}

//CEF views framework 系统托盘
type ViewsFrameTray struct {
	trayWindow *ViewsFrameworkBrowserWindow
	trayIcon   *lcl.TTrayIcon
	x, y, w, h int32
	mouseUp    TMouseEvent
	isClosing  bool
}

//CEF + LCL 托盘
type CEFTray struct {
	*lcl.TForm
	owner        lcl.IComponent
	trayIcon     *lcl.TTrayIcon
	chromium     IChromium
	windowParent ITCefWindowParent
	x, y, w, h   int32
	mouseUp      TMouseEvent
	isClosing    bool
	url          string
}

type SysTray struct {
}
