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
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod"
	"github.com/energye/energy/v2/pkgs/assetserve"
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
	app.SetDevToolsProtocolLogFile("E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\devtoolsprotocol.log")
	//app.SetRemoteDebuggingPort(8777) // http://localhost:8777/json/protocol
	app.SetOnFocusedNodeChanged(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, node *cef.ICefDomNode) {
		//fmt.Println("OnFocusedNodeChanged", node.GetElementTagName(), node.GetElementAttributes())
	})
	cef.BrowserWindow.Config.Width = 800
	cef.BrowserWindow.Config.Height = 600
	cef.BrowserWindow.Config.X = 300
	cef.BrowserWindow.Config.Y = 300
	cef.BrowserWindow.Config.EnableCenterWindow = false
	cef.BrowserWindow.Config.EnableMaximize = false
	cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			// 打开新窗口或url
			var rodWindow *rod.Chromium
			ipc.On("open-url-window", func(url string, windowId int) int {
				fmt.Println("open-url:", url)
				wp := cef.NewWindowProperty()
				wp.Url = url
				rodWindow = rod.NewWindow(nil, wp, nil)
				window.RunOnMainThread(func() {
					rodWindow.CreateBrowser()
				})
				return 0
			})
			// 关闭窗口
			ipc.On("close-window", func(windowId int) {
				fmt.Println("close-window:", windowId)
				//if tmpWindow, ok := windows[windowId]; ok {
				//	tmpWindow.Window.CloseBrowserWindow()
				//}
			})
			// 主窗口的控制台消息
			chromium := window.Chromium()
			chromium.SetOnConsoleMessage(func(sender lcl.IObject, browser *cef.ICefBrowser, level consts.TCefLogSeverity, message, source string, line int32) bool {
				fmt.Println("ConsoleMessage:", message)
				return false
			})

			// 抓取
			//var page *rod.Page
			ipc.On("crawling", func(windowId int) string {
				// 以下所有操作都需要在线程里，否则UI线程被锁死
				fmt.Println("crawling")
				if rodWindow == nil {
					return "还未打开窗口"
				}
				page := rodWindow.Page()
				page.MustElement("#kw").MustSelectAllText().MustInput("") //清空文本框
				page.MustElement("#kw").MustInput("go energy")            //输入内容
				page.MustElement("#su").MustClick()                       //点击按钮
				wrapper := page.MustElement("#wrapper_wrapper")           //根据id获取标签
				containers := wrapper.MustElements(".c-container")        //根据class样式获取所有标签
				for len(containers) == 0 {                                // 返回0个继续获取
					containers = wrapper.MustElements(".c-container")
				}
				fmt.Println("containers:", len(containers))
				for _, container := range containers {
					a := container.MustElement("a")
					fmt.Println("a:", a.MustText())
				}
				fmt.Println(page.MustHTML()[0:100])
				return page.MustHTML()[0:100]
			})
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
