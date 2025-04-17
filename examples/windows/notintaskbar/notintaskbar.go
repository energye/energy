//go:build windows
// +build windows

package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
)

func main() {
	// 全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 创建应用
	cefApp := cef.NewApplication()
	// 禁用主窗口
	cef.BrowserWindow.Config.EnableMainWindow = false
	// 不显示在任务栏后 窗口状态失效, 所以这里设置不能最大化
	cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			//取到 Window
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			//设置不显示在任务栏
			bw.SetNotInTaskBar()
		}
	})
	// 运行应用
	cef.Run(cefApp)
}
