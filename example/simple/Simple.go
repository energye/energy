package main

import (
	"fmt"
	"github.com/energye/energy/cef"
)

func main() {
	cef.GlobalCEFInit(nil, nil)
	cfg := cef.NewApplicationConfig()
	cfg.SetFrameworkDirPath("E:\\SWT\\CEF4Delphi-Libs-105.3.39\\chromium-64")
	cefApp := cef.NewApplication(cfg)
	cef.BrowserWindow.Config.DefaultUrl = "https://www.baidu.com"
	cef.BrowserWindow.Config.Title = "这是一个简单示例"
	cef.BrowserWindow.Config.Width = 1024
	cef.BrowserWindow.Config.Height = 768
	cef.BrowserWindow.Config.SetChromiumConfig(cef.NewChromiumConfig())
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInit")
	})
	cef.BrowserWindow.SetBrowserInitAfter(func(browserWindow *cef.TCefWindowInfo) {
		fmt.Println("SetBrowserInitAfter")
	})
	cef.Run(cefApp)
}
