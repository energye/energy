package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/commons"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//主进程浏览器初始化
func MainBrowserInit() {
	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableViewSource(true)
	config.SetEnableDevTools(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//默认加载的URL
	if commons.IsWindows() {
		cef.BrowserWindow.Config.DefaultUrl = "E:\\SWT\\gopath\\src\\swt-lazarus\\demo17-dll-load\\demo-golang-dll-01-chromium\\demos\\demo-sub-process\\resources\\demo-misc.html"
	} else if commons.IsLinux() {
		cef.BrowserWindow.Config.DefaultUrl = "/home/sxm/app/swt/gopath/src/github.com/energye/energy/demos/demo-sub-process/resources/demo-misc.html"
	} else if commons.IsDarwin() {
		cef.BrowserWindow.Config.DefaultUrl = "/Users/zhangli/go/src/github.com/energye/energy/demos/demo-sub-process/resources/demo-misc.html"
	}
	//主进程 IPC事件
	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		fmt.Println("主进程IPC事件注册")
	})

	//主进程初始化回调函数
	//在这个函数里，主进程浏览初始化之前创建窗口之后
	//在这里可以设置窗口的属性和事件监听
	//SetOnBeforePopup是子进程弹出窗口时触发，可以改变子进程窗口属性
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		lcl.Application.SetOnMinimize(func(sender lcl.IObject) {
			fmt.Println("minimize")
		})
		fmt.Println("主进程初始化回调函数")
		browserWindow.Window.SetCaption("这里设置应用标题")
		browserWindow.Window.SetPosition(types.PoScreenCenter) //窗口局中显示
		browserWindow.Window.SetWidth(1024)
		browserWindow.Window.SetHeight(768)

		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow *cef.TCefWindowInfo, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup: " + beforePopupInfo.TargetUrl)
			popupWindow.Window.SetCaption("改变了标题 - " + beforePopupInfo.TargetUrl)
			popupWindow.Window.SetWidth(800)
			popupWindow.Window.SetHeight(600)
			return false
		})
	})
}
