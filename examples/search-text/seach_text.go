package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/common"
	demoCommon "github.com/energye/energy/v2/examples/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/seach_text.html"
	cef.BrowserWindow.Config.Title = "Energy 搜索页面中文本"
	if common.IsLinux() && cefApp.IsUIGtk3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//搜索的结果在这个函数中返回
		event.SetOnFindResult(func(sender lcl.IObject, browser *cef.ICefBrowser, identifier, count int32, selectionRect *cef.TCefRect, activeMatchOrdinal int32, finalUpdate bool) {
			fmt.Println("OnFindResult:", identifier, count, selectionRect, activeMatchOrdinal, finalUpdate)
		})
	})
	//监听事件
	ipc.On("search-text", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		fmt.Println("搜索文本", bw)
		text := context.ArgumentList().GetStringByIndex(0)
		fmt.Println("搜索内容", text)
		bw.Browser().Find(text, false, false, true)
	})
	//运行应用
	cef.Run(cefApp)
}
