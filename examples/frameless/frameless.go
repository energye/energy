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
	"github.com/cyber-xxm/energy/v2/cef/ipc/context"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/common/tray"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl/rtl/version"
	"path/filepath"
)

//go:embed resources
var resources embed.FS

//用于Go版本低于1.16
//go:generate energy bindata --fs --o=assets/assets.go --pkg=assets --paths=./resources/...

// go build -ldflags "-s -w"
func main() {
	//命令: go generate 生成内置资源
	//resources := assets.AssetFile()
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache", "frameless")
	//创建应用
	app := cef.NewApplication()
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetUseMockKeyChain(true) // MacOS dev:test mock key

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.EnableHideCaption = true
	//cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Title = "Energy Vue + ElementUI 示例"
	cef.BrowserWindow.Config.Width = 1200
	cef.BrowserWindow.Config.Height = 600
	chromiumConfig := cef.BrowserWindow.Config.ChromiumConfig()
	chromiumConfig.SetEnableMenu(false)        //禁用右键菜单
	chromiumConfig.SetEnableWindowPopup(false) //禁用弹
	//监听窗口状态事件
	ipc.On("window-state", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		state := context.ArgumentList().GetIntByIndex(0)
		// 当前 ipc 在非UI线程中执行，窗口控制需要在UI线程执行
		cef.RunOnMainThread(func() {
			if state == 0 {
				fmt.Println("窗口最小化")
				bw.Minimize()
			} else if state == 1 {
				fmt.Println("窗口最大化/还原")
				bw.Maximize()
			} else if state == 3 {
				fmt.Println("全屏/退出全屏", bw.IsFullScreen(), bw.WindowState())
				if bw.IsFullScreen() {
					bw.ExitFullScreen()
				} else {
					bw.FullScreen()
				}
			}
		})
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
		if window.IsLCL() {
			tray.LCLTray(window)
		}
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
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
