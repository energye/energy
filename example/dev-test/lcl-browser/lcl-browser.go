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
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
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
		event.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, regions *cef.TCefDraggableRegions, window cef.IBrowserWindow) {
			fmt.Println("RegionsCount:", regions.RegionsCount(), regions.Regions())
			for i := 0; i < regions.RegionsCount(); i++ {
				fmt.Printf("i: %+v region: %+v\n", i, regions.Regions()[i])
			}
			//win.SetWindowLong(window.Handle(), win.GWL_EXSTYLE, uintptr(win.GetWindowLong(handle, win.GWL_EXSTYLE)|win.WS_EX_LAYERED))
			//win.SetLayeredWindowAttributes(window.Handle(), 0, 100, win.LWA_ALPHA)
			//win.UpdateLayeredWindow
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			popupWindow.SetSize(800, 600)
			return false
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
		//linux系统，默认使用 VF
		if window.IsLCL() {
			if common.IsWindows() {
				//支持 windows
				//traydemo.LCLCefTrayDemo(window) //对于LCL+CEF web端技术托盘实现无法在VF中使用
				traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
				//traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用, 推荐linux中使用
			} else {
				traydemo.SysTrayDemo(window) //系统原生托盘，在windows下不如lcl组件的好用, 推荐linux中使用
				//LCL窗口中,托盘组件支持 windows or macosx
				//traydemo.LCLTrayDemo(window) //LCL托盘, VF窗口组件中无法创建或使用LCL组件
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

//WindowTransparent 窗口透明组件不透明设置
//func WindowTransparent(hWnd types.HWND) {
//	exStyle := win.GetWindowLong(hWnd, win.GWL_EXSTYLE)
//	exStyle = exStyle | win.WS_EX_LAYERED&^win.WS_EX_TRANSPARENT // or WS_EX_TRANSPARENT;
//	win.SetWindowLong(hWnd, win.GWL_EXSTYLE, uintptr(exStyle))
//	win.SetLayeredWindowAttributes(hWnd, //指定分层窗口句柄
//		colors.ClNavy,                  //crKey指定需要透明的背景颜色值，可用RGB()宏  0-255
//		255,                            //bAlpha设置透明度，0表示完全透明，255表示不透明
//		win.LWA_ALPHA|win.LWA_COLORKEY) //LWA_ALPHA: crKey无效，bAlpha有效；
//	//LWA_COLORKEY：窗体中的所有颜色为crKey的地方全透明，bAlpha无效。
//	//LWA_ALPHA | LWA_COLORKEY：crKey的地方全透明，其它地方根据bAlpha确定透明度
//}

//WindowAngle 窗口四圆角设置
//
// nLeftRect 指定了x坐标的左上角区域逻辑单位
// nTopRect 指定了y坐标的左上角区域逻辑单位
// nRightRect 指定了x坐标的右下角区域逻辑单位
// nBottomRect 指定了y坐标的右下角区域逻辑单位
// nWidthEllipse 指定创建圆角的宽度逻辑单位
// nHeightEllipse 指定创建圆角的高度逻辑单位
//func WindowAngle(hWnd types.HWND, nLeftRect, nTopRect, nRightRect, nBottomRect, nWidthEllipse, nHeightEllipse int32) {
//	hr := cef.WinCreateRoundRectRgn(t.LongInt(nLeftRect), t.LongInt(nTopRect), t.LongInt(nRightRect), t.LongInt(nBottomRect), t.LongInt(nWidthEllipse), t.LongInt(nHeightEllipse))
//	cef.WinSetWindowRgn(hWnd, hr, true)
//	cef.WinDeleteObject(hr)
//}
