package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
)

//主进程浏览器初始化
func MainBrowserInit() {
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//主进程 IPC事件
	ipc.IPC.Browser().SetOnEvent(func(event ipc.IEventOn) {
		fmt.Println("主进程IPC事件注册")
	})
	//主窗口初始化回调函数
	//在这个函数里，主进程浏览初始化之前创建窗口之后
	//在这里可以设置窗口的属性和事件监听
	//SetOnBeforePopup是子进程弹出窗口时触发，可以改变子进程窗口属性
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		fmt.Println("主窗口初始化回调函数")
		window.SetTitle("这里设置应用标题")
		window.SetCenterWindow(true)

		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup: " + beforePopupInfo.TargetUrl)
			popupWindow.SetTitle("改变了标题 - " + beforePopupInfo.TargetUrl)
			return false
		})
	})

}
