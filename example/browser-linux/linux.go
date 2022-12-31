package main

import (
	"embed"
	"github.com/energye/energy/cef"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, nil)
	//创建应用
	config := cef.NewApplicationConfig()
	config.SetMultiThreadedMessageLoop(false)
	config.SetExternalMessagePump(false)
	cefApp := cef.NewApplication(config)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "https://www.baidu.com"
	//运行应用
	cef.Run(cefApp)
}
