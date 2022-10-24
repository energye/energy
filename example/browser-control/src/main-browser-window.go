package src

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/commons"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"os"
)

func MainBrowserWindow() {
	//只有启动主进程才会继续执行
	if !commons.Args.IsMain() {
		return
	}
	fmt.Println("os.Args", os.Args)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\example\\browser-control\\resources\\index.html"
	//窗口的标题
	cef.BrowserWindow.Config.Title = "Energy - 浏览器控制"
	//窗口宽高
	cef.BrowserWindow.Config.Width = 1024
	cef.BrowserWindow.Config.Height = 768
	//chromium配置
	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)
	config.SetEnableDevTools(true)
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//创建窗口时的回调函数 对浏览器事件设置，和窗口属性组件等创建和修改
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		//设置应用图标 这里加载的图标是内置到执行程序里的资源文件
		lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
		//页面加载处理进度
		event.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
			//参数-进度
			args := ipc.NewArgumentList()
			args.SetFloat64(0, progress)
			//触发html页面监听的事件，使用ipc通信发送给指定的浏览器页面
			browserWindow.Chromium().Emit("onLoadingProgressChange", args, browserWindow.Browser)
		})
		//页面加载状态，根据状态判断是否加载完成，和是否可前进后退
		event.SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
			//参数-加载状态，是否可前进，后退
			args := ipc.NewArgumentList()
			args.SetBool(0, isLoading)
			args.SetBool(1, canGoBack)
			args.SetBool(2, canGoForward)
			//触发html页面监听的事件，使用ipc通信发送给指定的浏览器页面
			browserWindow.Chromium().Emit("onLoadingStateChange", args, browserWindow.Browser)
		})
	})
	//创建窗口之后对对主窗口的属性、组件或子创建的创建
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
	//页面按钮事件 跳转指定URL
	ipc.IPC.Browser().On("goUrl", func(context ipc.IIPCContext) {
		//参数
		args := context.Arguments()
		//URL参数
		fmt.Println("goUrl", args.GetString(0))
		//获得当前浏览器的window信息
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		//得到页面上所有iframe,默认有一个主Frame
		for _, frame := range info.Frames {
			//页面内容是使用iframe实现的，这个iframe不是主要的
			if !frame.IsMain() {
				frame.LoadUrl(args.GetString(0))
			}
		}
	})
	//页面按钮事件 返回
	ipc.IPC.Browser().On("goBack", func(context ipc.IIPCContext) {
		fmt.Println("goBack")
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().GoBack()
	})
	//页面按钮事件 前进
	ipc.IPC.Browser().On("goForward", func(context ipc.IIPCContext) {
		fmt.Println("goForward")
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().GoForward()
	})
	//页面按钮事件 刷新
	ipc.IPC.Browser().On("refresh", func(context ipc.IIPCContext) {
		fmt.Println("refresh")
		info := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		info.Chromium().Reload()
	})
}
