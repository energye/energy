package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/win"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	var app = cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/msgbox.html"
	cef.BrowserWindow.Config.Title = "Energy - msgbox"
	if common.IsLinux() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}

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
		// 系统消息提示框目前仅能在LCL窗口组件下使用
		// LCL 各种系统组件需要在UI线程中执行, 但ipc.on非UI线程
		// 所以需要使用 QueueAsyncCall 包裹在UI线程中执行
		if window.IsLCL() {
			ipc.On("showmsgbox", func() {
				cef.QueueAsyncCall(func(id int) {
					lcl.ShowMessage("消息提示框")
				})
			})
			ipc.On("showmsgbox-confirm", func() {
				// lcl 各种系统组件需要在UI线程中执行, 但ipc.on非UI线程
				// 所以需要使用 QueueAsyncCall 包裹在UI线程中执行
				cef.QueueAsyncCall(func(id int) {
					if lcl.Application.MessageBox("消息", "标题", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == types.IdOK {
						lcl.ShowMessage("你点击了“是")
					}
				})
			})
			// 使用窗口模拟一个消息提示框

		}
	})

	//运行应用
	cef.Run(app)
}
