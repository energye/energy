package main

import (
	"github.com/energye/energy/v2/cef"
)

func main() {
	//Global initialization must be called
	cef.GlobalInit(nil, nil)
	//Create an application
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	//运行应用
	cef.Run(app)
}
