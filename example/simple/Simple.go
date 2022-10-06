package main

import (
	"fmt"
	"github.com/energye/energy/cef"
)

//这是一个简单的窗口创建示例
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, nil)
	//可选的应用配置
	cfg := cef.NewApplicationConfig()
	//指定chromium的二进制包框架根目录, 不指定为当前程序执行目录
	cfg.SetFrameworkDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64")
	//创建应用
	cefApp := cef.NewApplication(cfg)
	//对主窗口的配置
	cef.BrowserWindow.Config.DefaultUrl = "https://www.baidu.com"
	cef.BrowserWindow.Config.Title = "energy - 这是一个简单的窗口示例"
	cef.BrowserWindow.Config.Width = 1024
	cef.BrowserWindow.Config.Height = 768
	cef.BrowserWindow.Config.SetChromiumConfig(cef.NewChromiumConfig())
	//通过创建窗口时的回调函数 对浏览器事件设置，和窗口属性组件等创建和修改
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		//设置应用图标
		//lcl.Application.Icon().LoadFromFSFile("")
		fmt.Println("SetBrowserInit")
	})
	//通过创建窗口之后对对主窗口的属性、组件或子创建的创建
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
	//运行应用
	cef.Run(cefApp)
}
