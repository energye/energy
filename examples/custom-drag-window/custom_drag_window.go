package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, resources)
	app := cef.NewApplication()
	app.SetUseMockKeyChain(true) // macos dev:test
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.Title = "ENERGY - Custom Drag Window"
	cef.BrowserWindow.Config.Width = 800
	cef.BrowserWindow.Config.Height = 600
	cef.BrowserWindow.Config.EnableHideCaption = true
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = resources
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnBeforePopup(func(sender lcl.IObject, popupWindow cef.IBrowserWindow, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			popupWindow.WindowProperty().Title = "test1111"
			return false
		})
		event.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32, window cef.IBrowserWindow) {

			rgn := winapi.CreateRectRgn(0, 0, 150, 150)
			fmt.Println(winapi.PtInRegion(rgn, 100, 100))
		})
	})
	cef.Run(app)
}
