package main

import "github.com/energye/energy/v2/cef"

// windows xp 使用 CEF 49, 是最后一个支持 windows xp 的CEF版本, 使用-与比较新的CEF版本没太多区别，但注意的是不支持新的CEF API 使用时需注意, 虽然energy提供了。
// 仅在 windows xp, 以及其它 windows 系统中运行, 不支持 Linux, MacOS

func main() {

	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	//运行应用
	cef.Run(app)
}
