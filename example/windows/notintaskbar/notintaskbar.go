//go:build windows

package main

import (
	"github.com/energye/energy/v2/cef"
)

func main() {
	// 全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	// 创建应用
	cefApp := cef.NewApplication()
	// 设置主窗口不在任务栏中显示 仅windows
	cef.BrowserWindow.Config.MainFormOnTaskBar = false
	// 不显示在任务栏后 窗口状态失效, 所以这里设置不能最大化
	cef.BrowserWindow.Config.EnableResize = false
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	// 运行应用
	cef.Run(cefApp)
}
