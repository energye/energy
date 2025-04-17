package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()
	//可选配置: VF组件窗口需要指定该配置
	cefApp.SetRemoteDebuggingPort(8888) //远程端口方式, 需自定义端口号
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.csdn.net"
	//chromium 配置
	config := cef.BrowserWindow.Config.ChromiumConfig()
	config.SetEnableMenu(true)     //启用右键菜单
	config.SetEnableDevTools(true) //启用开发者工具
	//运行应用
	cef.Run(cefApp)
}
