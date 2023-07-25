package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/logger"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"strings"
)

//go:embed resources
var resources embed.FS

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	if common.IsLinux() && cefApp.IsUIGtk3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			var browserWindow *cef.LCLBrowserWindow
			var newForm *cef.LCLBrowserWindow
			ipc.On("openWindow", func() {
				if newForm == nil {
					newForm = cef.NewLCLWindow(cef.NewWindowProperty())
					newForm.SetTitle("新窗口标题")
					newForm.SetWidth(400)
					newForm.SetHeight(200)
					btn := lcl.NewButton(newForm)
					btn.SetParent(newForm)
					btn.SetCaption("点击我有提示")
					btn.SetWidth(100)
					btn.SetLeft(100)
					btn.SetTop(50)
					btn.SetOnClick(func(sender lcl.IObject) {
						lcl.ShowMessage("新窗口的按钮事件提示")
					})
				}
				cef.QueueAsyncCall(func(id int) {
					newForm.ShowModal()
				})
			})
			ipc.On("openBrowserWindow", func(url, title string, width, height int32) {
				if browserWindow == nil || browserWindow.IsClosing() {
					wp := cef.NewWindowProperty()
					wp.Url = url
					wp.Title = title
					browserWindow = cef.NewLCLBrowserWindow(nil, wp)
					browserWindow.SetWidth(width)
					browserWindow.SetHeight(height)
					browserWindow.SetShowInTaskBar()
					browserWindow.EnableDefaultCloseEvent()
					browserWindow.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
						fmt.Println("SetOnTitleChange", wp.Title, title)
					})
				}
				cef.QueueAsyncCall(func(id int) {
					browserWindow.ShowModal() // 欌态窗口，打开开发者工具后，关闭窗口有问题。
				})
			})
		}

		var elliptic = func(window cef.IBrowserWindow) {
			hRegion := winapi.WinCreateEllipticRgn(0, 0, 550, 550)
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
				// windows和linux GTK2下有效果
				popupWindow.WindowProperty().EnableWebkitAppRegionDClk = false
				popupWindow.SetSize(550, 550)
				popupWindow.HideTitle()
				popupWindow.Chromium().Config().SetEnableMenu(false)
				window.RunOnMainThread(func() {
					// 如果使用winapi方式改变窗口，需要在主线程中运行
					elliptic(popupWindow)
				})
			} else if strings.Index(beforePopupInfo.TargetUrl, "transparent") > 0 {
				// windows和linux GTK2下有效果
				popupWindow.SetSize(200, 200)
				popupWindow.HideTitle()
				transparent(popupWindow)
			} else if strings.Index(beforePopupInfo.TargetUrl, "model_window") > 0 {
				// windows和linux GTK2下有效果
				popupWindow.SetSize(200, 200)
				popupWindow.HideTitle()
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
