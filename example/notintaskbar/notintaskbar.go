package main

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/golcl/lcl"
)

func main() {
	// 全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 创建应用
	cefApp := cef.NewApplication()
	// 设置不在任务栏中显示
	cef.BrowserWindow.Config.MainFormOnTaskBar = false
	// 不显示在任务栏后 窗口状态失效, 所以这里设置不能最大化
	cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		bw := window.AsLCLBrowserWindow().BrowserWindow()
		bw.Constraints().SetOnChange(func(sender lcl.IObject) {
			println("window change", bw.WindowState())
		})
	})
	// 运行应用
	cef.Run(cefApp)
}
