package main

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/golcl/lcl/types"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.Config.WindowInitState = types.WsFullScreen
	cef.BrowserWindow.Config.EnableHideCaption = true
	//运行应用
	cef.Run(app)
}
