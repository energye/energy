//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl/rtl/version"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	config := cef.NewApplicationConfig()
	//config.SetMultiThreadedMessageLoop(false)
	//config.SetExternalMessagePump(false)
	//config.SetRemoteDebuggingPort(9999)
	cefApp := cef.NewApplication(config)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.EnableHideCaption = true
	cef.BrowserWindow.Config.Title = "Energy Vue + ElementUI 示例"
	cef.BrowserWindow.Config.Width = 1200
	chromiumConfig := cef.BrowserWindow.Config.ChromiumConfig()
	chromiumConfig.SetEnableMenu(false) //禁用右键菜单

	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		//监听窗口状态事件
		event.On("window-state", func(context ipc.IIPCContext) {
			bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			state := context.Arguments().GetInt32(0)
			if state == 0 {
				fmt.Println("窗口最小化")
				bw.Minimize()
			} else if state == 1 {
				fmt.Println("窗口最大化/还原")
				bw.Maximize()
			}
		})
		//监听窗口关闭事件
		event.On("window-close", func(context ipc.IIPCContext) {
			bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
			bw.CloseBrowserWindow()
		})
		event.On("os-info", func(context ipc.IIPCContext) {
			fmt.Println("系统信息", version.OSVersion.ToString())
			context.Result().SetString(version.OSVersion.ToString())
		})
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
