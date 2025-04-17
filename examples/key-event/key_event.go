package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/key-event.html"
	cef.BrowserWindow.Config.Title = "Energy - Key Event"
	if common.IsLinux() && api.WidgetUI().IsGTK3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	//在主窗口初始化回调函数里设置浏览器事件
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, browserWindow cef.IBrowserWindow) {
		event.SetOnKeyEvent(func(sender lcl.IObject, browser *cef.ICefBrowser, event *cef.TCefKeyEvent, osEvent consts.TCefEventHandle, window cef.IBrowserWindow, result *bool) {
			fmt.Printf("%s  KeyEvent:%+v osEvent:%+v\n", string(rune(event.Character)), event, osEvent)
		})
		browserWindow.Chromium().SetOnPreKeyEvent(func(sender lcl.IObject, browser *cef.ICefBrowser, event *cef.TCefKeyEvent, osEvent consts.TCefEventHandle) (isKeyboardShortcut, result bool) {
			fmt.Printf("%s  PreKeyEvent:%+v osEvent:%+v\n", string(rune(event.Character)), event, osEvent)
			return
		})
	})
	//内置http服务链接安全配置
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}
