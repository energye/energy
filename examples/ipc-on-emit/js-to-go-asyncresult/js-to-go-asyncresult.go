package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/exception"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	exception.SetOnException(func(message string) {
		fmt.Println("Global Exception message", message)
	})
	app := cef.NewApplication()
	app.SetUseMockKeyChain(true)
	cef.BrowserWindow.Config.Title = "Energy - ipc async result"
	cef.BrowserWindow.Config.Url = "http://localhost:22111/index.html"

	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22111
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		ipc.On("gotest", func() {
			alreadyRunCount := 0
			infos := cef.BrowserWindow.GetWindowInfos()
			next := func() {
				fmt.Println("next task")
			}
			for _, info := range infos {
				ipc.EmitTargetAndCallback("jstest", info.Target(), []interface{}{"testdata"}, func(data string, channel callback.IChannel) {
					fmt.Println("data:", data, "windowId:", channel.BrowserId(), "channelId:", channel.ChannelId())
					alreadyRunCount++
					total := len(cef.BrowserWindow.GetWindowInfos())
					if total == alreadyRunCount {
						next()
					}
				})
			}
			fmt.Println("gotest end")
		})
	})
	//运行应用
	cef.Run(app)
}
