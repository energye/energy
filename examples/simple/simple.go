package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	app.SetUseMockKeyChain(true)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energye.github.io"
	//运行应用
	cef.Run(app)
}
