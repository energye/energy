package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	var app = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/msgbox.html"
	cef.BrowserWindow.Config.Title = "Energy - msgbox"

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

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 系统消息提示框目前仅能在LCL窗口组件下使用
		// LCL 各种系统组件需要在UI线程中执行, 但ipc.on非UI线程
		// 所以需要使用 QueueAsyncCall 包裹在UI线程中执行
		ipc.On("showmsgbox", func() {
			fmt.Println("showmsgbox")
			window.RunOnMainThread(func() {
				fmt.Println("消息提示框")
				lcl.ShowMessage("消息提示框")
			})
		})
		ipc.On("showmsgbox-confirm", func() {
			// lcl 各种系统组件需要在UI线程中执行, 但ipc.on非UI线程
			// 所以需要使用 QueueAsyncCall 包裹在UI线程中执行
			window.RunOnMainThread(func() {
				if lcl.Application.MessageBox("消息", "标题", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == types.IdOK {
					lcl.ShowMessage("你点击了“是")
				}
			})
		})
	})

	//运行应用
	cef.Run(app)
}
