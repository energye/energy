//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//辅助工具
type auxTools struct {
	devToolsWindow   *LCLBrowserWindow //devTools
	devToolsX        int32             //上次改变的窗体位置，宽度
	devToolsY        int32             //
	devToolsWidth    int32             //
	devToolsHeight   int32             //
	viewSourceWindow *LCLBrowserWindow //viewSource
	viewSourceUrl    string            //
	viewSourceX      int32             //上次改变的窗体位置，宽度
	viewSourceY      int32             //
	viewSourceWidth  int32             //
	viewSourceHeight int32             //
}

//窗口属性
type WindowProperty struct {
	IsShowModel  bool               //是否以模态窗口显示
	WindowState  types.TWindowState //窗口 状态
	Title        string             //窗口 标题
	Url          string             //默认打开URL
	Icon         string             //窗口图标 加载本地图标
	IconFS       string             //窗口图标 加载emfs内置图标
	CanMinimize  bool               //窗口 是否启用最小化功能
	CanMaximize  bool               //窗口 是否启用最大化功能
	CanResize    bool               //窗口 是否允许调整窗口大小
	CanClose     bool               //窗口 关闭时是否关闭窗口
	CenterWindow bool               //窗口 是否居中显示
	AlwaysOnTop  bool               //窗口 置顶
	X            int32              //窗口 CenterWindow=false X坐标
	Y            int32              //窗口 CenterWindow=false Y坐标
	Width        int32              //窗口 宽
	Height       int32              //窗口 高
}

type IBrowserWindow interface {
	Id() int32
	Show()
	Hide()
	Maximize()
	Minimize()
	CloseBrowserWindow()
	WindowType() consts.WINDOW_TYPE
	SetWindowType(windowType consts.WINDOW_TYPE)
	Browser() *ICefBrowser
	Chromium() IChromium
	DisableMaximize()
	DisableMinimize()
	EnableMaximize()
	EnableMinimize()
	IsClosing() bool
	AsViewsFrameworkBrowserWindow() IViewsFrameworkBrowserWindow
	AsLCLBrowserWindow() ILCLBrowserWindow
	Frames() TCEFFrame
	addFrame(frame *ICefFrame)
	setBrowser(browser *ICefBrowser)
	createAuxTools()
	getAuxTools() *auxTools
	SetTitle(title string)
	IsViewsFramework() bool
	IsLCL() bool
	WindowProperty() *WindowProperty
	SetWidth(value int32)
	SetHeight(value int32)
	SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32)
	SetCenterWindow(value bool)
	NewCefTray(width, height int32, url string) ITray
	NewTray() ITray
}

type ILCLBrowserWindow interface {
	IBrowserWindow
	BrowserWindow() *LCLBrowserWindow
	EnableDefaultCloseEvent()
	EnableAllDefaultEvent()
	SetDefaultInTaskBar()
	SetShowInTaskBar()
	SetNotInTaskBar()
	WindowParent() ITCefWindowParent
	DisableTransparent()
	EnableTransparent(value uint8)
	DisableSystemMenu()
	DisableHelp()
	EnableSystemMenu()
	EnableHelp()
}

type IViewsFrameworkBrowserWindow interface {
	IBrowserWindow
	BrowserWindow() *ViewsFrameworkBrowserWindow
	CreateTopLevelWindow()
	CenterWindow(size *TCefSize)
	Component() lcl.IComponent
	WindowComponent() *TCEFWindowComponent
	BrowserViewComponent() *TCEFBrowserViewComponent
}

//创建一个 窗口默认属性
func NewWindowProperty() *WindowProperty {
	return &WindowProperty{
		Title:        "Energy",
		Url:          "about:blank",
		CanMinimize:  true,
		CanMaximize:  true,
		CanResize:    true,
		CanClose:     true,
		CenterWindow: true,
		X:            100,
		Y:            100,
		Width:        1024,
		Height:       768,
	}
}
