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
	"energye/notice"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//cef.BrowserWindow.Config.EnableWebkitAppRegion = false
	//cef.BrowserWindow.Config.EnableResize = false
	//cef.BrowserWindow.Config.EnableWebkitAppRegionDClk = false
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//event.SetOnWidgetCompMsg(func(sender lcl.IObject, message types.TMessage, aHandled bool) {
		//	fmt.Println("SetOnWidgetCompMsg:", message)
		//})
		//browserWindow := window.AsLCLBrowserWindow().BrowserWindow()
		//browserWindow.Constraints().SetMinWidth(300)
		//browserWindow.Constraints().SetMinHeight(300)
		window.HideTitle()

		//window.AsLCLBrowserWindow().WindowParent().SetBoundsRect(types.Rect(100, 100, 800, 500))
		//window.DisableResize()

		//browserWindow.BorderIcons().Exclude(types.BiHelp, types.BiMinimize, types.BiMaximize, types.BiSystemMenu)
		//browserWindow.SetBorderStyle(types.BsSizeable)
		event.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, regions *cef.TCefDraggableRegions) {
			fmt.Println("RegionsCount:", regions.RegionsCount(), regions.Regions())
			for i := 0; i < regions.RegionsCount(); i++ {
				fmt.Printf("i: %+v region: %+v\n", i, regions.Regions()[i])
			}
			//win.SetWindowLong(window.Handle(), win.GWL_EXSTYLE, uintptr(win.GetWindowLong(handle, win.GWL_EXSTYLE)|win.WS_EX_LAYERED))
			//win.SetLayeredWindowAttributes(window.Handle(), 0, 100, win.LWA_ALPHA)
			//win.UpdateLayeredWindow
		})
		// show or hide, caption bar
		//go func() {
		//	var b = true
		//	for {
		//		time.Sleep(time.Second)
		//		cef.QueueAsyncCall(func(id int) {
		//			b = !b
		//			if b {
		//				window.ShowTitle()
		//			} else {
		//				window.HideTitle()
		//			}
		//		})
		//	}
		//}()
	})
	cef.BrowserWindow.SetBrowserInitAfter(func(window cef.IBrowserWindow) {
		sysTray(window)
		//sys_tray.TrayMain()
		return
	})
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
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
		notice.SendNotification(notice.NewNotification("标题", "内容内容内容内容内容内容内容内容内容内容内容内容内容内容\n内容内容内容内容内容内容内容内容内容内容内容内容内容内"))
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
	//return
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
				sysTray.SetIconFS("resources/icon.png")
				menuItem.SetIconFS("resources/icon.ico")
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Disable()
				check.Uncheck()
			}
		}
	}()
}
