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
	"github.com/energye/energy/common"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	fmt.Println("ARGS", common.Args.ProcessType())
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	config := cef.NewApplicationConfig()
	config.SetMultiThreadedMessageLoop(false)
	config.SetExternalMessagePump(false)
	config.SetRemoteDebuggingPort(33333)
	cefApp := cef.NewApplication(config)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.EnableWebkitAppRegion = true
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//window.DisableResize()
		window.SetTitle("这里改变了窗口标题")
		window.SetSize(1600, 900)
		fmt.Println("cef.BrowserWindow.SetViewFrameBrowserInit", window)
		fmt.Println("LCL", window.AsLCLBrowserWindow(), "VF", window.AsViewsFrameworkBrowserWindow())
		event.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, regions *cef.TCefDraggableRegions) {
			fmt.Println("SetOnDraggableRegionsChanged", regions.RegionsCount(), "frame:", frame.Id, frame.Url)
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("IsViewsFramework:", popupWindow.IsViewsFramework())
			popupWindow.SetTitle("修改了标题: " + beforePopupInfo.TargetUrl)
			popupWindow.SetCenterWindow(true)
			popupWindow.EnableResize()
			popupWindow.SetSize(1600, 1600)
			browserWindow := popupWindow.AsViewsFrameworkBrowserWindow()
			browserWindow.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
				fmt.Println("popupWindow.SetOnWindowCreated", window)
			})
			//browserWindow.BrowserWindow().CreateTopLevelWindow()
			//browserWindow.BrowserWindow().HideTitle()
			fmt.Println("browserWindow:", browserWindow, browserWindow.WindowComponent().WindowHandle())
			return false
		})
		window.AsViewsFrameworkBrowserWindow().WindowComponent().SetOnWindowActivationChanged(func(sender lcl.IObject, window *cef.ICefWindow, active bool) {
			fmt.Println("SetOnWindowActivationChanged", active)
		})
		window.Show()
		fmt.Println("SetBrowserInit 结束")
	})
	cef.BrowserWindow.SetBrowserInitAfter(func(window cef.IBrowserWindow) {
		bw := window.AsViewsFrameworkBrowserWindow().BrowserWindow()
		fmt.Println("handle", bw.WindowComponent().WindowHandle().ToPtr())
		//设置隐藏窗口标题
		//window.HideTitle()
		fmt.Println("SetBrowserInitAfter 结束")
		sysTray(window)
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

func sysTray(browserWindow cef.IBrowserWindow) {
	sysTray := browserWindow.NewSysTray()
	sysTray.SetIconFS("resources/icon.ico")
	sysTray.SetHint("中文hint\n换行中文")
	sysTray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	tray := sysTray.AsSysTray()
	check := tray.AddMenuItem("check")
	check.Check()
	not := tray.AddMenuItem("通知")
	not.Click(func() {
		tray.Notice("标题", "内notice 是一个跨平台的系统通知库\nnotice 是一个跨平台的系统通知库", 1000)
	})
	enable := tray.AddMenuItem("启用/禁用")
	enable.Click(func() {
		fmt.Println("启用/禁用 点击")
	})
	tray.AddSeparator()
	menuItem := tray.AddMenuItem("1级菜单1", func() {
		fmt.Println("1级菜单1")
	})
	menuItem.SetIconFS("resources/icon.ico")
	tray.AddSeparator()
	item := tray.AddMenuItem("1级菜单2")
	item.AddSubMenu("2级子菜单1")
	sub2Menu := item.AddSubMenu("2级子菜单2")
	sub2Menu.AddSubMenu("3级子菜单1")
	tray.AddSeparator()
	tray.AddMenuItem("退出", func() {
		fmt.Println("退出")
		browserWindow.CloseBrowserWindow()
	})

	sysTray.Show()
	return
	//测试图标切换
	go func() {
		var b bool
		for {
			time.Sleep(time.Second * 2)
			b = !b
			if b {
				sysTray.SetHint(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				sysTray.SetIconFS("resources/icon_1.ico")
				menuItem.SetIconFS("resources/icon_1.ico")
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Enable()
				check.Check()
			} else {
				sysTray.SetHint(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				sysTray.SetIconFS("resources/icon.ico")
				menuItem.SetIconFS("resources/icon.ico")
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Disable()
				check.Uncheck()
			}
		}
	}()
}
