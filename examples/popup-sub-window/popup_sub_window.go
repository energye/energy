package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/cyber-xxm/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	lclTypes "github.com/energye/golcl/lcl/types"
	"strings"
)

//go:embed resources
var resources embed.FS

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.LDebug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	if common.IsDarwin() {
		app.SetUseMockKeyChain(true)
	}
	//cefApp.EnableVFWindow(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	if common.IsLinux() && api.WidgetUI().IsGTK3() {
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
					cef.QueueAsyncCall(func(id int) {
						newForm = cef.NewLCLWindow(cef.NewWindowProperty(), nil)
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
						newForm.Show()
					})
				}
			})
			ipc.On("openBrowserWindow", func(url, title string, width, height int32) {
				cef.QueueAsyncCall(func(id int) {
					if browserWindow == nil || browserWindow.IsClosing() {
						wp := cef.NewWindowProperty()
						wp.Url = url
						wp.Title = title
						wp.EnableHideCaption = true
						wp.ShowInTaskBar = lclTypes.StNever
						browserWindow = cef.NewLCLBrowserWindow(nil, wp, nil)
						browserWindow.SetWidth(width)
						browserWindow.SetHeight(height)
						// 必须设置为子窗口
						browserWindow.SetWindowType(consts.WT_POPUP_SUB_BROWSER)
						//browserWindow.HideTitle()
						browserWindow.EnableAllDefaultEvent()
						// 置顶控制
						//browserWindow.SetFormStyle(lclTypes.FsSystemStayOnTop)
						//browserWindow.SetFormStyle(lclTypes.FsStayOnTop)
						browserWindow.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
							fmt.Println("SetOnTitleChange", wp.Title, title)
						})
					}
					browserWindow.Show() // 欌态窗口，打开开发者工具后，关闭窗口有问题。
				})
			})
		}

		var elliptic = func(window cef.IBrowserWindow) {
			hRegion := winapi.CreateEllipticRgn(0, 0, 550, 550)
			winapi.SetWindowRgn(types.HWND(window.Handle()), hRegion, true)
			winapi.DeleteObject(hRegion)
		}
		var transparent = func(window cef.IBrowserWindow) {
			if window.IsLCL() {
				browserWindow := window.AsLCLBrowserWindow().BrowserWindow()
				browserWindow.EnableTransparent(100) //窗口透明
			}
		}
		//弹出子窗口
		event.SetOnBeforePopup(func(sender lcl.IObject, popupWindow cef.IBrowserWindow, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo,
			popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			fmt.Println("beforePopupInfo-TargetUrl", beforePopupInfo.TargetUrl, strings.Index(beforePopupInfo.TargetUrl, "popup_1"), strings.Index(beforePopupInfo.TargetUrl, "popup_2"))
			fmt.Printf("beforePopupInfo: %+v\n", beforePopupInfo)
			fmt.Printf("popupFeatures: %+v\n", popupFeatures)
			fmt.Printf("windowInfo: %+v\n", windowInfo)
			fmt.Printf("settings: %+v\n", settings)
			if strings.Index(beforePopupInfo.TargetUrl, "popup_1") > 0 {
				popupWindow.SetSize(800, 600)
				popupWindow.HideTitle()
				popupWindow.WindowProperty().EnableResize = false
				popupWindow.WindowProperty().ShowInTaskBar = lclTypes.StNever
			} else if strings.Index(beforePopupInfo.TargetUrl, "popup_2") > 0 {
				popupWindow.SetSize(300, 300)
				popupWindow.HideTitle()
			} else if strings.Index(beforePopupInfo.TargetUrl, "elliptic") > 0 {
				// windows和linux GTK2下有效果
				popupWindow.WindowProperty().EnableWebkitAppRegionDClk = false
				window.RunOnMainThread(func() {
					popupWindow.SetSize(550, 550)
					//popupWindow.HideTitle()
					if popupWindow.IsLCL() {
						popupWindow.AsLCLBrowserWindow().BrowserWindow().SetBorderStyle(lclTypes.BsNone)
					}
					//popupWindow.Chromium().Config().SetEnableMenu(false)
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
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
