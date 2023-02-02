//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"energye/notice"
	"energye/systray"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"sync"
)

//TMouseEvent 鼠标事件
type TMouseEvent func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) bool

//TrayICONClick 托盘图标鼠标事件
type TrayICONClick func()

//ITray 托盘接口
//
//实现4种系统托盘 1: LCLTray LCL组件, 2: CEFTray CEF基于LCL组件+html, 3: ViewsFrameTray VF(views framework)组件+html, 4: SysTray 系统原生
//
//1. LCLTray 对Windows、MacOSX支持较好，linux由于gtk2与gtk3原因目前无法正常使用
//
//2. CEFTray Windows
//
//3. ViewsFrameTray Windows
//
//4. SysTray 对Windows、MacOSX和Linux支持较好
type ITray interface {
	SetTitle(title string)                       //SetTitle 设置标题
	Show()                                       //Show 显示/启动 托盘
	close()                                      //
	SetOnClick(fn TrayICONClick)                 //SetOnClick 单击事件
	SetOnDblClick(fn TrayICONClick)              //SetOnDblClick 双击事件
	SetIconFS(iconResourcePath string)           //SetIconFS 设置托盘图标
	SetIcon(iconResourcePath string)             //SetIcon 设置托盘图标
	SetHint(value string)                        //SetHint 设置托盘hint(鼠标移动到托盘图标显示的文字)
	AsSysTray() *SysTray                         //AsSysTray 尝试转换为 SysTray 组件托盘，如果创建的是其它类型托盘返回nil
	AsViewsFrameTray() *ViewsFrameTray           //AsViewsFrameTray 尝试转换为 views framework 组件托盘, 如果创建的是其它类型托盘返回nil
	AsCEFTray() *CEFTray                         //AsCEFTray 尝试转换为 LCL+CEF 组件托盘, 如果创建的是其它类型托盘返回nil
	AsLCLTray() *LCLTray                         //AsLCLTray 尝试转换为 LCL 组件托盘, 如果创建的是其它类型托盘返回nil
	Notice(title, content string, timeout int32) //Notice 托盘系统通知
}

//LCLTray LCL组件 托盘
type LCLTray struct {
	owner     lcl.IComponent
	trayIcon  *lcl.TTrayIcon
	popupMenu *lcl.TPopupMenu
}

//ViewsFrameTray VF(views framework)组件+html 托盘
type ViewsFrameTray struct {
	trayWindow *ViewsFrameworkBrowserWindow
	trayIcon   *lcl.TTrayIcon
	visible    bool
	x, y, w, h int32
	mouseUp    TMouseEvent
	isClosing  bool
}

//CEFTray CEF基于LCL组件+html 托盘
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

//SysTray 系统原生
type SysTray struct {
	once           sync.Once
	menu           *SysMenu
	icon           []byte
	title, tooltip string
	click          TrayICONClick
	dClick         TrayICONClick
	rClick         func(menu systray.IMenu)
	start, stop    func()
}

func notification(tray lcl.IComponent, title, content string, timeout int32) {
	var lclTrayNotice *lcl.TTrayIcon
	if tray != nil {
		lclTrayNotice = tray.(*lcl.TTrayIcon)
	}
	var lclNotice = func() {
		lclTrayNotice.SetBalloonTitle(title)
		lclTrayNotice.SetBalloonHint(content)
		lclTrayNotice.SetBalloonTimeout(timeout)
		lclTrayNotice.ShowBalloonHint()
	}
	var sysNotice = func() {
		notify := notice.NewNotification(title, content)
		notify.SetTimeout(timeout)
		notice.SendNotification(notify)
	}
	if lclTrayNotice != nil {
		lclNotice()
	} else {
		sysNotice()
	}
}
