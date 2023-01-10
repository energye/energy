package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	sys_tray "github.com/energye/energy/example/dev-test/sys-tray"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//event.SetOnWidgetCompMsg(func(sender lcl.IObject, message types.TMessage, aHandled bool) {
		//	fmt.Println("SetOnWidgetCompMsg:", message)
		//})
		browserWindow := window.AsLCLBrowserWindow().BrowserWindow()
		event.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, regions *cef.TCefDraggableRegions) {
			fmt.Println("RegionsCount:", regions.RegionsCount(), regions.Regions())
			for i := 0; i < regions.RegionsCount(); i++ {
				fmt.Println("region:", i, regions.Regions()[i])
			}
			newButton := lcl.NewButton(browserWindow)
			newButton.SetParent(browserWindow)
			newButton.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
				fmt.Println("SetOnMouseDown")
			})
			button := lcl.NewImageButton(browserWindow)
			button.SetParent(browserWindow)
			button.SetImageCount(0)
			button.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
				fmt.Println("SetOnMouseDown")
			})
			button.SetBounds(0, 0, 100, 100)
			//
			panel := lcl.NewPanel(browserWindow)
			panel.SetParent(browserWindow)
			handle := panel.Handle()
			panel.SetCaption("adsfasdfsadfdsaf " + fmt.Sprintf("%d", handle))
			//win.SetWindowLong(browserWindow.Handle(), win.GWL_EXSTYLE, uintptr(win.GetWindowLong(handle, win.GWL_EXSTYLE)|win.WS_EX_LAYERED))
			//win.SetLayeredWindowAttributes(browserWindow.Handle(), 0, 100, win.LWA_ALPHA)
		})
	})
	cef.BrowserWindow.SetBrowserInitAfter(func(window cef.IBrowserWindow) {
		sys_tray.TrayMain()
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
