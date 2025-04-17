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
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/erod/examples/devtools/crawling"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	lcl.DEBUG = true // 一些底层调用时错误会输出
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()

	//远程调试开启
	//app.SetRemoteDebuggingPort(22222)
	//app.SetRemoteAllowOrigins("*")// 指定 http://ip:port 或  *

	cef.BrowserWindow.Config.Width = 800
	cef.BrowserWindow.Config.Height = 600
	cef.BrowserWindow.Config.X = 300
	cef.BrowserWindow.Config.Y = 300
	cef.BrowserWindow.Config.EnableCenterWindow = false
	cef.BrowserWindow.Config.EnableMaximize = false
	cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Url = "http://localhost:22023/index.html"

	/*
		此示例仅限学习
		这个并不是一个几行代码的示例
		它是在 rod 基础上增加的一个扩展, 完全复用 rod 的功能(对于一些窗口状态管理, Chromium操作还是需要在energy API控制)
		devtools-protocol 和 rod 的发送消息和处理上，完全不同
			rod: WebSocket
			energy: CEF API

		使用方式:
			1. 创建一个energy扩展rod的window(rod.NewWindow)或chromium(rod.NewChromium)
			2. 创建浏览器或显示浏览器窗口 CreateBrowser
			3. 获取 rod Page 对象开始搞事情

		rod使用: https://go-rod.github.io
		devtools: https://chromedevtools.github.io/devtools-protocol

		除此之外你也可以自己定义 devtools 的使用
	*/
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 返回 ids
		ipc.On("window-infos", func() []*crawling.Info {
			return crawling.WindowIds()
		})
		// 创建一个窗口
		ipc.On("create", func(url string, testType int) int {
			return crawling.Create(url, testType)
		})
		// 显示这个窗口
		ipc.On("show", func(windowId int, url string) {
			fmt.Println("open-url:", url)
			crawling.Show(windowId, url)
		})
		// 关闭窗口
		ipc.On("close-window", func(windowId int) bool {
			fmt.Println("close-windowId:", windowId)
			return crawling.Close(windowId)
		})
		// 主窗口的控制台消息
		chromium := window.Chromium()
		chromium.SetOnConsoleMessage(func(sender lcl.IObject, browser *cef.ICefBrowser, level consts.TCefLogSeverity, message, source string, line int32) bool {
			fmt.Println("javascript-console.log:", message)
			return false
		})
		// 仅测试区分测试的功能类型
		const (
			testTypeDefault  = 0
			testTypeUpload   = 1
			testTypeDownload = 2
		)
		// 抓取
		ipc.On("crawling", func(windowId, testType int) {
			fmt.Println("crawling windowId:", windowId)
			// 在ipc里使用rod当前需要开启协程，或使用异步IPC监听选项配置
			if testType == testTypeDefault {
				go crawling.Crawling(windowId)
			} else if testType == testTypeUpload {
				go crawling.Upload(windowId)
			} else if testType == testTypeDownload {
				go crawling.Download(windowId)
			}
		})
		// 异步IPC监听选项配置
		//ipc.On("crawling", func(windowId, testType int) {
		//
		//}, types.OnOptions{Mode: types.MAsync})

		// 测试上传
		var url string
		ipc.On("upload-start-server", func(typ int) (string, int) {
			if url == "" {
				url = crawling.UploadServer()
			}
			fmt.Println("启动上传文件测试服务", url)
			windowId := crawling.Create(url, typ)
			return url, windowId
		})
		// 测试下载文件
		ipc.On("download-file", func(typ int) (string, int) {
			url := "http://localhost:22023/download.html"
			windowId := crawling.Create(url, typ)
			return url, windowId
		})
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22023               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
