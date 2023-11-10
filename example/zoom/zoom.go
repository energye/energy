package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/assetserve"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	app := cef.NewApplication()
	// 强制使用VF窗口
	//app.SetExternalMessagePump(false)
	//app.SetMultiThreadedMessageLoop(false)
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	ipc.On("zoom-inc", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().BrowserZoom(consts.ZOOM_INC)
		fmt.Println("zoom-inc")
	})
	ipc.On("zoom-dec", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().BrowserZoom(consts.ZOOM_DEC)
		fmt.Println("zoom-dec")
	})
	ipc.On("zoom-reset", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().BrowserZoom(consts.ZOOM_RESET)
		fmt.Println("zoom-reset")
	})
	//运行应用
	cef.Run(app)
}
