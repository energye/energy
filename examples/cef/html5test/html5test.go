package main

import (
	"github.com/energye/energy/v2/cef"
	_ "github.com/energye/energy/v2/examples/syso"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://html5test.opensuse.org"
	//运行应用
	cef.Run(app)
}
