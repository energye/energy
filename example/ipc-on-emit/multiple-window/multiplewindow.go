package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/callback"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed assets
var assets embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &assets)
	//创建应用
	app := cef.NewApplication()
	app.SetExternalMessagePump(false)
	app.SetMultiThreadedMessageLoop(false)
	cef.BrowserWindow.Config.Title = "Energy - ipc multiple-window"
	//cef.BrowserWindow.Config.EnableMainWindow = false
	//本地资源加载
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "assets",
		FS:         &assets,
		Home:       "multiple-window.html",
	}.Build())

	// 多窗口接收消息
	ipc.On("sendMessage", func(channel callback.IChannel) {
		// 获得所有窗口
		infos := cef.BrowserWindow.GetWindowInfos()
		fmt.Println("windows-count:", len(infos))
		for _, info := range infos {
			// 将消息发送到目标窗口, 类型为JS接收
			iTarget := info.Browser().Target(target.TgJs)
			ipc.EmitTarget("receiveMessage", iTarget, time.Now().String())
		}
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			popupWindow.SetSize(800, 600)
			return false
		})
	})
	//运行应用
	cef.Run(app)
}
