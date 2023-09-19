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
	"embed"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	if common.IsLinux() && cefApp.IsUIGtk3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	// 默认隐藏窗口，CEF初始化时是在显示窗口时创建
	// 1.默认不居中
	// 2.窗口显示在屏幕之外，数值要大于窗口宽高
	cef.BrowserWindow.Config.EnableCenterWindow = false
	cef.BrowserWindow.Config.X = -1200
	cef.BrowserWindow.Config.Y = -800
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// chromium 创建完成之后再隐藏掉窗口
		event.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser, win cef.IBrowserWindow) bool {
			window.RunOnMainThread(func() { // 在这UI线程执行
				window.Hide()
			})
			println("hide window")
			// 5秒后显示窗口
			go func() {
				println("5秒后显示窗口")
				time.Sleep(time.Second * 5)
				window.RunOnMainThread(func() {
					//window.SetPoint(300, 300)
					window.SetCenterWindow(true)
					window.Show()
					println("show window")
				})
			}()
			return false
		})
	})
	//运行应用
	cef.Run(cefApp)
}
