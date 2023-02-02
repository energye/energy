package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication(nil)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//运行应用
	cef.Run(cefApp)
}
