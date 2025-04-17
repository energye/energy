package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	"github.com/cyber-xxm/energy/v2/cef/ipc/target"
	ipcTypes "github.com/cyber-xxm/energy/v2/cef/ipc/types"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.Config.Title = "Energy - go on event - go emit event"

	/*
		主进程和子进程IPC事件在Go交互
			事件监听
				1. 在主进程(browser)中监听事件,用于接收子进程消息
				2. 在子进程(render)中监听事件,用于接收主进程回复消息
			事件触发
				1. 子进程触发主进程监听的事件
				2. 主进程回复消息给子进程
	*/

	// 事件监听
	//2. 渲染进程
	ipc.On("renderOn", func(data string) {
		fmt.Println("renderOn data:", data)
	},
		// 仅子进程监听
		ipcTypes.OnOptions{OnType: ipcTypes.OtSub})

	// 在渲染进程页面加载时触发主进程 browserOn 事件
	app.SetOnRenderLoadStart(func(browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
		url := frame.Url()
		arguments := []interface{}{url}
		//触发主进程事件, 并接收返回结果
		ipc.EmitTargetAndCallback("browserOn", target.NewTargetMain(), arguments, func(result string) {
			// 接收主进程直接返回结果
			fmt.Println("render result:", result)
		})
	})

	// 主进程窗口初始化
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 事件监听
		//1. 主进程
		ipc.On("browserOn", func(channel callback.IChannel, url string) string {
			fmt.Println("browserOn channel:", channel.ChannelId(), "url:", url)
			// 多种方式返回结果到子进程
			// 1. 使用ipc事件返回结果到子进程
			ipc.EmitTarget("renderOn", target.NewTarget(nil, 0, channel.ChannelId()), "reply result")
			// 2. 直接返回结果到子进程
			return "reply result"
		})
	})

	//运行应用
	cef.Run(app)
}
