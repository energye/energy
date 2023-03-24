package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/cef/ipc"
	"github.com/energye/energy/common/assetserve"
	"os"
	"path"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	cef.BrowserWindow.Config.Title = "Energy 打印PFD预览"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	wd, _ := os.Getwd()
	//监听事件
	ipc.On("print-pdf", func(context ipc.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		savePath := path.Join(wd, "example", "browser-print-pdf", "test.pdf")
		fmt.Println("当前页面保存为PDF", savePath)
		bw.Chromium().PrintToPDF(savePath)
	})
	//运行应用
	cef.Run(cefApp)
}
