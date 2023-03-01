package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
)

//go:embed resources
var resources embed.FS
var cefApp *cef.TCEFApplication

func main() {
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/ipc-event.html"
	cef.BrowserWindow.Config.Title = "Energy - ipc-event"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnBrowseProcessMessageReceived(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *cef.ICefProcessMessage) bool {
			fmt.Println("browser 进程接收消息", message.Name())
			return false
		})
	})
	cefApp.SetOnProcessMessageReceived(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, sourceProcess consts.CefProcessId, message *cef.ICefProcessMessage) bool {
		fmt.Println("render 进程接收消息", message.Name())
		return false
	})
	//运行应用
	cef.Run(cefApp)
}
