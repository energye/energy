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
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/example/dev-test/traydemo"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"os"
	"path"
)

//go:embed resources
var resources embed.FS

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication()
	cefApp.SetMultiThreadedMessageLoop(false)
	cefApp.SetExternalMessagePump(false)
	fmt.Println("TotalSystemMemory", cefApp.TotalSystemMemory(), cefApp.UsedMemory())
	cefApp.SetRemoteDebuggingPort(33333)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.EnableWebkitAppRegion = true
	cef.BrowserWindow.Config.EnableHideCaption = true
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		return
		//window.DisableResize()
		window.SetCenterWindow(true)
		window.SetTitle("这里改变了窗口标题")
		window.SetSize(1024, 900)
		fmt.Println("cef.BrowserWindow.SetViewFrameBrowserInit", window)
		fmt.Println("LCL", window.AsLCLBrowserWindow(), "VF", window.AsViewsFrameworkBrowserWindow())
		window.AsViewsFrameworkBrowserWindow().SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
			fmt.Println("WindowCreated.window", window.WindowAppIcon().Width(), window.WindowAppIcon().Height())
			image := cef.ImageRef.New()
			cw, _ := os.Getwd()
			cw = path.Join(cw, "example", "dev-test", "vf-browser", "resources", "icon.png")
			byt, err := os.ReadFile(cw)
			fmt.Println("image icon.png", len(byt), err)
			image.AddPng(1.2, byt)
			fmt.Println("image", image.Width(), image.Height())
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("IsViewsFramework:", popupWindow.IsViewsFramework())
			popupWindow.SetTitle("修改了标题: " + beforePopupInfo.TargetUrl)
			popupWindow.EnableResize()
			popupWindow.DisableMaximize()
			popupWindow.DisableResize()
			popupWindow.DisableMinimize()
			popupWindow.SetSize(800, 600)
			browserWindow := popupWindow.AsViewsFrameworkBrowserWindow()
			browserWindow.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
				fmt.Println("popupWindow.SetOnWindowCreated", window.WindowAppIcon())
			})
			//browserWindow.SetOnGetInitialBounds(func(sender lcl.IObject, window *cef.ICefWindow, aResult *cef.TCefRect) {
			//	fmt.Println("popupWindow.SetOnGetInitialBounds", *aResult)
			//})
			//browserWindow.BrowserWindow().CreateTopLevelWindow()
			//browserWindow.BrowserWindow().HideTitle()
			fmt.Println("browserWindow:", browserWindow, browserWindow.WindowComponent().WindowHandle())
			return false
		})
		window.AsViewsFrameworkBrowserWindow().WindowComponent().SetOnWindowActivationChanged(func(sender lcl.IObject, window *cef.ICefWindow, active bool) {
			fmt.Println("SetOnWindowActivationChanged", active)
		})
		window.AsViewsFrameworkBrowserWindow().BrowserViewComponent().SetOnBrowserCreated(func(sender lcl.IObject, browserView *cef.ICefBrowserView, browser *cef.ICefBrowser) {
			fmt.Println("BrowserViewComponent OnBrowserCreated Instance", browser.Identifier(), "Instance Identifier", browserView.Browser().Identifier())
		})
		window.Show()
		fmt.Println("SetBrowserInit 结束")
		if common.IsLinux() || common.IsDarwin() {
			//在VF窗口组件中, 推荐linux和macosx中使用
			traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用,
		} else {
			//不支持windows VF窗口组件中无法创建或使用LCL组件
			//traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
			//traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
			//支持windows
			traydemo.LCLVFTrayDemo(window) //对于LCL+VF web端技术托盘实现
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
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
