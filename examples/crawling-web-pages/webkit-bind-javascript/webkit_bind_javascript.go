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
	"github.com/cyber-xxm/energy/v2/examples/crawling-web-pages/webkit-bind-javascript/implant"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
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
	cef.BrowserWindow.Config.Width = 800
	cef.BrowserWindow.Config.Height = 600
	cef.BrowserWindow.Config.X = 300
	cef.BrowserWindow.Config.Y = 300
	cef.BrowserWindow.Config.EnableCenterWindow = false
	cef.BrowserWindow.Config.EnableMaximize = false
	cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"

	// 注入本地js
	app.SetOnWebKitInitialized(func() {
		// 在处理器中接收注入 js 的native本地函数，不支持dom操作
		// 下面示例接收自定义js解析dom后返回支持的类型到处理器
		v8Handler := cef.V8HandlerRef.New()
		v8Handler.Execute(func(name string, object *cef.ICefV8Value, arguments *cef.TCefV8ValueArray, retVal *cef.ResultV8Value, exception *cef.ResultString) bool {
			fmt.Println("Execute javascript native method-name:", name)
			htmlValue := arguments.Get(0)
			if htmlValue.IsString() {
				fmt.Println("dom-html-type string")
				html := htmlValue.GetStringValue()
				fmt.Println("dom-html:", html)
				return true
			} else if htmlValue.IsArray() {
				// 多dom获取，例如css的 class 或 tag 选择器
				fmt.Println("dom-html-type array", htmlValue.GetArrayLength())
				for i := 0; i < htmlValue.GetArrayLength(); i++ {
					args := htmlValue.GetValueByIndex(i)
					fmt.Println("dom-html:", args.GetStringValue())
				}
				return true
			} else if htmlValue.IsObject() {
				fmt.Println("dom-html-type object")
				return true
			}
			return false
		})
		jsCode := implant.HelperJS()
		// 注册JS, dom 是自定义注入的js全局变量名或常量名
		cef.RegisterExtension("v8/dom", jsCode, v8Handler)
	})
	ipc.On("implantName", func(data string) {
		fmt.Println("implantName", data)
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// ipc
		ipc.On("execute", func() {
			// 所有窗口都执行JS
			// cef.BrowserWindow.GetWindowInfo() // 或使用 browserId 获取指定窗口并执行js
			// 不具备js上下文功能
			for _, w := range cef.BrowserWindow.GetWindowInfos() {
				w.Chromium().ExecuteJavaScript(`dom.element("body")`, "", w.Chromium().Browser().MainFrame(), 0)
				w.Chromium().ExecuteJavaScript(`dom.elements("div")`, "", w.Chromium().Browser().MainFrame(), 0)
				w.Chromium().ExecuteJavaScript(`dom.elementX("//div")`, "", w.Chromium().Browser().MainFrame(), 0)
			}
		})

		// 主窗口的控制台消息
		chromium := window.Chromium()
		chromium.SetOnConsoleMessage(func(sender lcl.IObject, browser *cef.ICefBrowser, level consts.TCefLogSeverity, message, source string, line int32) bool {
			// 这里只能取出字符串
			fmt.Println("javascript-console.log:", message)
			return false
		})
		chromium.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			chromium.ExecuteJavaScript(`dom.element("body")`, "", frame, 0)
		})
	})
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
