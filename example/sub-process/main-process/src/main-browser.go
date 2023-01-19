package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
)

//主进程浏览器初始化
func MainBrowserInit() {
	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableViewSource(true)
	config.SetEnableDevTools(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//默认加载的URL
	if common.IsWindows() {
		cef.BrowserWindow.Config.Url = "E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium\\demos\\demo-sub-process\\resources\\demo-misc.html"
	} else if common.IsLinux() {
		cef.BrowserWindow.Config.Url = "/home/sxm/app/swt/gopath/src/github.com/energye/energy/demos/demo-sub-process/resources/demo-misc.html"
	} else if common.IsDarwin() {
		cef.BrowserWindow.Config.Url = "/Users/zhangli/go/src/github.com/energye/energy/demos/demo-sub-process/resources/demo-misc.html"
	}
	//主进程 IPC事件
	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		fmt.Println("主进程IPC事件注册")
	})

	//主窗口初始化回调函数
	//在这个函数里，主进程浏览初始化之前创建窗口之后
	//在这里可以设置窗口的属性和事件监听
	//SetOnBeforePopup是子进程弹出窗口时触发，可以改变子进程窗口属性
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		lcl.Application.SetOnMinimize(func(sender lcl.IObject) {
			fmt.Println("SetBrowserInit minimize")
		})
		fmt.Println("主窗口初始化回调函数")
		window.SetTitle("这里设置应用标题")
		window.SetCenterWindow(true)
		window.SetWidth(1024)
		window.SetHeight(768)

		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup: " + beforePopupInfo.TargetUrl)
			popupWindow.SetTitle("改变了标题 - " + beforePopupInfo.TargetUrl)
			popupWindow.SetWidth(800)
			popupWindow.SetHeight(600)
			return false
		})
	})
}
