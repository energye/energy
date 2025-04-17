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
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	app := cef.NewApplication()

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "fs://energy"
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         resources,
	}.Build())
	//cef.BrowserWindow.Config.EnableClose = false
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//浏览器窗口之后回调，在这里获取创建的浏览器ID
		event.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser, window cef.IBrowserWindow) bool {
			// 创建完之后再拿浏览器id
			fmt.Println("on-create-window-ok", browser.Identifier(), window.Id())
			ipc.Emit("on-create-window-ok", browser.Identifier(), window.Id())
			return false // 什么都不做
		})
		//浏览器窗口关闭回调, 在这里触发ipc事件通知主窗口
		event.SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction, window cef.IBrowserWindow) bool {
			ipc.Emit("on-close-window", window.Id())
			return false
		})
		//---- ipc 监听事件
		// 监听事件, 创建新窗口
		ipc.On("create-window", func(name string) {
			println("create-window", name)
			// 创建窗口常规属性对象
			wp := cef.NewWindowProperty()
			wp.Url = "fs://energy/new-window.html"
			wp.Title = name
			wp.EnableHideCaption = true // 无标题窗口
			// 创建浏览器窗口
			newWindow := cef.NewBrowserWindow(nil, wp, nil)
			newWindow.SetWidth(800)
			newWindow.SetHeight(600)
			// EnableAllDefaultEvent 启用所有默认实现事件
			// 如果未启用，所有事件需要你自己控制和实现, 大部分功能无法使用
			newWindow.EnableAllDefaultEvent()
			window.RunOnMainThread(func() {
				println("create-window show", name)
				// 在主线程中
				newWindow.Show()
			})
		})
		// 改变当前窗口大小
		ipc.On("resize", func(_type int, channel callback.IChannel) {
			println("resize type", _type, "channel", channel.ChannelId(), channel.BrowserId())
			win := cef.BrowserWindow.GetWindowInfo(channel.BrowserId())
			if win == nil {
				return
			}
			window.RunOnMainThread(func() {
				switch _type {
				case 1:
					win.SetSize(400, 200)
				case 2:
					win.SetSize(600, 400)
				case 3:
					win.SetSize(1024, 768)
				}
				win.SetCenterWindow(true)
			})
		})

	})
	//运行应用
	cef.Run(app)
}
