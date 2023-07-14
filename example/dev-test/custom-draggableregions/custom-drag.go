package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/pkgs/assetserve"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, nil)
	cefApp := cef.NewApplication()
	//cefApp.SetExternalMessagePump(false)
	//cefApp.SetMultiThreadedMessageLoop(false)
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	//cef.BrowserWindow.Config.EnableHideCaption = true
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			//window.AsLCLBrowserWindow().BrowserWindow().FramelessForDefault() //.FramelessForLine()
		}
	})
	cef.Run(cefApp)
}
