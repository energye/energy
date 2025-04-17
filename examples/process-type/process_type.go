package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/process"
	"github.com/cyber-xxm/energy/v2/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl/api"
)

func main() {
	// 进程类型获取
	// 通过命令行参数 type 区分
	process.Args.Print()
	if process.Args.ProcessType() == process.PT_MAIN {
		// 主进程
		println("main process")
	} else if process.Args.ProcessType() == process.PT_GPU {
		println("gpu process")
	} else if process.Args.ProcessType() == process.PT_UTILITY {
		println("utility process")
	} else if process.Args.ProcessType() == process.PT_RENDERER {
		// 渲染进程
		println("renderer process")
	}
	process.FrameId()
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	if common.IsLinux() && api.WidgetUI().IsGTK3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	//运行应用
	cef.Run(cefApp)
}
