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

//TMouseEvent 鼠标事件
type TMouseEvent func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) bool

//TrayICONClick 托盘图标鼠标事件
type TrayICONClick func()

//ITray 托盘接口
type ITray interface {
	SetTitle(title string)                                 //SetTitle 设置标题
	SetVisible(v bool)                                     //SetVisible 设置显示和隐藏托盘图标
	Visible() bool                                         //Visible 托盘的显示和隐藏状态
	Show()                                                 //Show 显示托盘菜单窗口 windows有效
	Hide()                                                 //Hide 隐藏托盘菜单窗口 windows有效
	close()                                                //close 关闭托盘菜单窗口 windows有效
	SetOnClick(fn TrayICONClick)                           //SetOnClick 单击事件
	SetOnDblClick(fn TrayICONClick)                        //SetOnDblClick 双击事件 linux 和 macos 可能不启作用
	SetIconFS(iconResourcePath string)                     //SetIconFS 设置托盘图标
	SetIcon(iconResourcePath string)                       //SetIcon 设置托盘图标
	SetHint(value string)                                  //SetHint 设置托盘hint(鼠标移动到托盘图标显示的文字)
	ShowBalloon()                                          //ShowBalloon 显示托盘气泡
	SetBalloon(title, content string, timeout int32) ITray //SetBalloon 设置托盘气泡内容
	AsSysTray() *SysTray                                   //AsSysTray 尝试转换为 SysTray 组件托盘，如果创建的是其它类型托盘返回nil
	AsViewsFrameTray() *ViewsFrameTray                     //AsViewsFrameTray 尝试转换为 views framework 组件托盘, 如果创建的是其它类型托盘返回nil
	AsCEFTray() *CEFTray                                   //AsCEFTray 尝试转换为 LCL+CEF 组件托盘, 如果创建的是其它类型托盘返回nil
	AsLCLTray() *LCLTray                                   //AsLCLTray 尝试转换为 LCL 组件托盘, 如果创建的是其它类型托盘返回nil
}

//LCLTray LCL 托盘
type LCLTray struct {
	owner     lcl.IComponent
	trayIcon  *lcl.TTrayIcon
	popupMenu *lcl.TPopupMenu
}

//ViewsFrameTray CEF views framework 托盘
type ViewsFrameTray struct {
	trayWindow *ViewsFrameworkBrowserWindow
	trayIcon   *lcl.TTrayIcon
	x, y, w, h int32
	mouseUp    TMouseEvent
	isClosing  bool
}

//CEFTray CEF + LCL 托盘
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

//SysTray 系统原生托盘
type SysTray struct {
	menu                *SysMenu
	trayStart, trayStop func()
}
