package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/view_source_index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	cef.BrowserWindow.Config.Title = "Energy 查看网页源码"
	config := cef.BrowserWindow.Config.ChromiumConfig()
	config.SetEnableViewSource(true)
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
