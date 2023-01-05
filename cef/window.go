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

//创建一个新window窗口
func NewWindow(windowProperty *WindowProperty, owner ...lcl.IComponent) *LCLBrowserWindow {
	if windowProperty == nil {
		windowProperty = NewWindowProperty()
	}
	var window = &LCLBrowserWindow{}
	window.windowProperty = windowProperty
	if len(owner) > 0 {
		window.TForm = lcl.NewForm(owner[0])
	} else {
		lcl.Application.CreateForm(&window)
	}
	window.ParentDoubleBuffered()
	window.FormCreate()
	window.SetShowInTaskBar()
	window.defaultWindowEvent()
	window.SetCaption(windowProperty.Title)
	if windowProperty.CenterWindow {
		window.SetWidth(windowProperty.Width)
		window.SetHeight(windowProperty.Height)
		window.SetPosition(types.PoDesktopCenter)
	} else {
		window.SetPosition(types.PoDesigned)
		window.SetBounds(windowProperty.X, windowProperty.Y, windowProperty.Width, windowProperty.Height)
	}
	if windowProperty.IconFS != "" {
		_ = window.Icon().LoadFromFSFile(windowProperty.IconFS)
	} else if windowProperty.Icon != "" {
		window.Icon().LoadFromFile(windowProperty.Icon)
	}
	if windowProperty.AlwaysOnTop {
		window.SetFormStyle(types.FsSystemStayOnTop)
	}
	window.EnabledMinimize(windowProperty.CanMinimize)
	window.EnabledMaximize(windowProperty.CanMaximize)
	if !windowProperty.CanResize {
		window.SetBorderStyle(types.BsSingle)
	}
	return window
}

type IBrowserWindow interface {
	Show()
	Hide()
	Close()
}

type ILCLBrowserWindow interface {
	IBrowserWindow
	Id() int32
}

type IViewsFrameworkBrowserWindow interface {
	IBrowserWindow
	CreateTopLevelWindow()
	CenterWindow(size *TCefSize)
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
