//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"time"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	// 默认隐藏窗口，CEF初始化时是在显示窗口时创建
	// 1.默认不居中
	// 2.窗口显示在屏幕之外，数值要大于窗口宽高
	cef.BrowserWindow.Config.EnableCenterWindow = false
	cef.BrowserWindow.Config.X = 0
	cef.BrowserWindow.Config.Y = 0
	cef.BrowserWindow.Config.Width = 1
	cef.BrowserWindow.Config.Height = 1
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// chromium 创建完成之后再隐藏掉窗口
		event.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType, window cef.IBrowserWindow) {
			window.Hide()
			println("hide window")
			// 5秒后显示窗口
			go func() {
				println("5秒后显示窗口")
				time.Sleep(time.Second * 5)
				window.RunOnMainThread(func() {
					//window.SetPoint(300, 300)
					window.SetSize(800, 600)
					window.SetCenterWindow(true)
					window.Show()
					println("show window")
				})
			}()
		})
	})
	//运行应用
	cef.Run(cefApp)
}
