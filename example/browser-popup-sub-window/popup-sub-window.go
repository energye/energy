package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"strings"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	if common.IsLinux() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		var elliptic = func(window cef.IBrowserWindow) {
			hRegion := winapi.WinCreateEllipticRgn(1, 1, 200, 200)
			winapi.WinSetWindowRgn(types.HWND(window.Handle()), hRegion, true)
		}
		var transparent = func(window cef.IBrowserWindow) {
			if window.IsLCL() {
				browserWindow := window.AsLCLBrowserWindow().BrowserWindow()
				browserWindow.EnableTransparent(100) //窗口透明
			}
		}
		//弹出子窗口
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("beforePopupInfo-TargetUrl", beforePopupInfo.TargetUrl, strings.Index(beforePopupInfo.TargetUrl, "popup_1"), strings.Index(beforePopupInfo.TargetUrl, "popup_2"))
			if strings.Index(beforePopupInfo.TargetUrl, "popup_1") > 0 {
				popupWindow.SetSize(800, 600)
				popupWindow.HideTitle()
			} else if strings.Index(beforePopupInfo.TargetUrl, "popup_2") > 0 {
				popupWindow.SetSize(300, 300)
				popupWindow.HideTitle()
			} else if strings.Index(beforePopupInfo.TargetUrl, "elliptic") > 0 {
				popupWindow.SetSize(200, 200)
				popupWindow.HideTitle()
				cef.QueueAsyncCall(func(id int) {
					// 如果使用winapi方式改变窗口，需要在主线程中运行
					elliptic(popupWindow)
				})
			} else if strings.Index(beforePopupInfo.TargetUrl, "transparent") > 0 {
				popupWindow.SetSize(200, 200)
				popupWindow.HideTitle()
				transparent(popupWindow)
			} else if strings.Index(beforePopupInfo.TargetUrl, "model_window") > 0 {
				popupWindow.SetSize(200, 200)
				popupWindow.HideTitle()
				//popupWindow.AsLCLBrowserWindow().BrowserWindow().EnableTransparent(100)
				popupWindow.WindowProperty().IsShowModel = true
			}
			return false
		})
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
