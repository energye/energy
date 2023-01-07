package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//创建应用
	config := cef.NewApplicationConfig()
	config.SetMultiThreadedMessageLoop(false)
	config.SetExternalMessagePump(false)
	cefApp := cef.NewApplication(config)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.CanResize = false
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		window.Show()
		fmt.Println("cef.BrowserWindow.SetViewFrameBrowserInit", window)
		fmt.Println("LCL", window.AsLCLBrowserWindow(), "VF", window.AsViewsFrameworkBrowserWindow())
		event.SetOnBeforeContextMenu(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, model *cef.ICefMenuModel) {
			model.AddCheckItem(model.CefMis.NextCommandId(), "测试")
		})
		event.SetOnContextMenuCommand(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, params *cef.ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool) {
			fmt.Println("SetOnContextMenuCommand", commandId)
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("IsViewsFramework:", popupWindow.IsViewsFramework())
			popupWindow.SetTitle("修改了标题: " + beforePopupInfo.TargetUrl)
			popupWindow.SetCenterWindow(true)
			browserWindow := popupWindow.AsViewsFrameworkBrowserWindow()
			//browserWindow.BrowserWindow().CreateTopLevelWindow()
			//browserWindow.BrowserWindow().HideTitle()
			fmt.Println("browserWindow:", browserWindow, browserWindow.WindowComponent().WindowHandle())
			return false
		})
		handle := window.AsViewsFrameworkBrowserWindow().WindowComponent().WindowHandle()
		fmt.Println("handle", handle, handle.ToPtr())
		window.AsViewsFrameworkBrowserWindow().WindowComponent().SetOnWindowActivationChanged(func(sender lcl.IObject, window *cef.ICefWindow, active bool) {
			fmt.Println("SetOnWindowActivationChanged", active)
		})

		//设置隐藏窗口标题
		//window.HideTitle()
		cefTray(window)
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

// 托盘 只适用 windows 的系统托盘, 基于html 和 ipc 实现功能
func cefTray(browserWindow cef.IBrowserWindow) {
	var url = "http://localhost:22022/min-browser-tray.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIcon("resources/icon.ico")
	tray.SetOnClick(func(sender lcl.IObject) {
		//browserWindow.SetVisible(!browserWindow.Visible())
	})
	tray.SetBalloon("气泡标题", "气泡内容", 2000)
	ipc.IPC.Browser().On("tray-show-balloon", func(context ipc.IIPCContext) {
		fmt.Println("tray-show-balloon")
		tray.ShowBalloon()
		tray.Hide()
	})
	ipc.IPC.Browser().On("tray-show-main-window", func(context ipc.IIPCContext) {
		//vb := !browserWindow.Visible()
		//browserWindow.SetVisible(vb)
		//if vb {
		//	if browserWindow.WindowState() == types.WsMinimized {
		//		browserWindow.SetWindowState(types.WsNormal)
		//	}
		//	browserWindow.Focused()
		//}
		tray.Hide()
	})
	ipc.IPC.Browser().On("tray-close-main-window", func(context ipc.IIPCContext) {
		browserWindow.CloseBrowserWindow()
	})
	ipc.IPC.Browser().On("tray-show-message-box", func(context ipc.IIPCContext) {
		cef.QueueAsyncCall(func(id int) {
			lcl.ShowMessage("tray-show-message-box 提示消息")
		})
		tray.Hide()
	})
	//托盘 end
}

// 托盘 系统原生 windows linux macos
func tray(browserWindow cef.IBrowserWindow) {
	//托盘 windows linux macos 系统托盘
	newTray := browserWindow.NewTray()
	tray := newTray.Tray()
	tray.SetIcon("resources/icon.ico")
	tray.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) bool {
		fmt.Println("SetOnMouseUp", button, shift, x, y)
		return false
	})
	menu1 := tray.AddMenuItem("父菜单", nil)
	menu1.Add(tray.NewMenuItem("子菜单", func(object lcl.IObject) {
		lcl.ShowMessage("子菜单点击 提示消息")
	}))
	tray.AddMenuItem("显示气泡", func(object lcl.IObject) {
		newTray.ShowBalloon()
	})
	tray.AddMenuItem("显示/隐藏", func(object lcl.IObject) {
	})
	tray.AddMenuItem("退出", func(object lcl.IObject) {
		browserWindow.CloseBrowserWindow()
	})
	//linux下有些问题
	tray.SetBalloon("气泡标题", "气泡内容", 2000)
	//托盘 end
}
