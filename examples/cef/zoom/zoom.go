package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	demoCommon "github.com/energye/energy/v2/examples/cef/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/energy/v2/types"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//创建应用
	app := cef.NewApplication()
	//主窗口的配置
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/zoom_index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	ipc.On("zoom-inc", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().BrowserZoom(types.ZOOM_INC)
		fmt.Println("zoom-inc")
	})
	ipc.On("zoom-dec", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().BrowserZoom(types.ZOOM_DEC)
		fmt.Println("zoom-dec")
	})
	ipc.On("zoom-reset", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.Chromium().BrowserZoom(types.ZOOM_RESET)
		fmt.Println("zoom-reset")
	})
	//运行应用
	cef.Run(app)
}
