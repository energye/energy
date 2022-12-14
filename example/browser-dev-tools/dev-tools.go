package main

import (
	"github.com/energye/energy/cef"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, nil)
	applicationConfig := cef.NewApplicationConfig()
	applicationConfig.SetRemoteDebuggingPort(8888) //远程端口方式, 需自定义端口号
	//创建应用
	cefApp := cef.NewApplication(applicationConfig)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.DefaultUrl = "https://www.csdn.net"
	//chromium 配置
	config := cef.NewChromiumConfig()
	config.SetEnableMenu(true)     //启用右键菜单
	config.SetEnableDevTools(true) //启用开发者工具
	cef.BrowserWindow.Config.SetChromiumConfig(config)
	//运行应用
	cef.Run(cefApp)
}
