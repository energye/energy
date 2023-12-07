//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/callback"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl/rtl/version"
	"time"
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
	cef.BrowserWindow.Config.Title = "Energy Vue"
	cef.BrowserWindow.Config.Width = 1200

	ipc.On("ipcGetBrowserId", func(context context.IContext) {
		fmt.Println("ipcGetBrowserId", context.BrowserId())
		context.Result(context.BrowserId())
	})
	ipc.On("ipcGetChannelId", func(channel callback.IChannel) int64 {
		fmt.Println("ipcGetChannelId", channel.ChannelId())
		return channel.ChannelId()
	})
	ipc.On("os-info", func(context context.IContext) {
		fmt.Println("os-info", version.OSVersion.ToString())
		context.Result(version.OSVersion.ToString())
	})
	ipc.On("ipcGetDateTime", func() string {
		now := time.Now()
		return now.Format("2006-01-02 15:04:05")
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
		// 在主进程中同步时间到JS事件
		go func() {
			for {
				now := time.Now()
				ipc.Emit("dateTime", now.Format("2006-01-02 15:04:05"))
				time.Sleep(time.Second)
			}
		}()
	})
	//运行应用
	cef.Run(cefApp)
}
