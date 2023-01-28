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
	"github.com/energye/energy/example/dev-test/traydemo"
	"github.com/energye/golcl/lcl"
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
		//linux系统，默认使用 VF
		if window.IsLCL() {
			if common.IsWindows() {
				//支持 windows
				traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
			} else {
				//traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用, 推荐linux中使用
				//LCL窗口中,托盘组件支持 windows or macosx
				traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
			}
		} else if window.IsViewsFramework() {
			if common.IsLinux() || common.IsDarwin() {
				//在VF窗口组件中, 推荐linux和macosx中使用
				traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用,
			} else {
				//不支持windows VF窗口组件中无法创建或使用LCL组件
				//traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
				//traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
				//支持windows
				//traydemo.LCLVFTrayDemo(window) //对于LCL+VF web端技术托盘实现
			}
		}
		println("browser init after end")
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
