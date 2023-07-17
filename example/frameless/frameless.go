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
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl/rtl/version"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication()
	// TODO 下面2个配置项用来切换使用VF或LCL窗口组件
	// VF = (ExternalMessagePump = false && MultiThreadedMessageLoop = false)
	//cefApp.SetExternalMessagePump(false)
	//cefApp.SetMultiThreadedMessageLoop(false)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	if common.IsLinux() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	if !common.IsDarwin() {
		// LCL macos 隐藏标题栏后，不能调整大小
		cef.BrowserWindow.Config.EnableHideCaption = true
		// LCL macos 隐藏标题栏后，该选项不生效
		cef.BrowserWindow.Config.EnableResize = true
	}
	cef.BrowserWindow.Config.Title = "Energy Vue + ElementUI 示例"
	cef.BrowserWindow.Config.Width = 1200
	chromiumConfig := cef.BrowserWindow.Config.ChromiumConfig()
	chromiumConfig.SetEnableMenu(false) //禁用右键菜单

	//监听窗口状态事件
	ipc.On("window-state", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		state := context.ArgumentList().GetIntByIndex(0)
		if state == 0 {
			fmt.Println("窗口最小化")
			bw.Minimize()
		} else if state == 1 {
			fmt.Println("窗口最大化/还原")
			bw.Maximize()
		}
	})
	//监听窗口关闭事件
	ipc.On("window-close", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.CloseBrowserWindow()
	})
	ipc.On("os-info", func(context context.IContext) {
		fmt.Println("系统信息", version.OSVersion.ToString())
		context.Result(version.OSVersion.ToString())
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//window.AsLCLBrowserWindow().FramelessForLine()
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
