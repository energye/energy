package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/exception"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/callback"
	"github.com/energye/energy/v2/cef/ipc/target"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed assets
var assets embed.FS

func main() {
	//logger.SetEnable(true)
	//logger.SetLevel(logger.CefLog_Debug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &assets)
	exception.SetOnException(func(message string) {
		fmt.Println("Exception message", message)
	})
	//创建应用
	app := cef.NewApplication()
	// 强制使用VF窗口
	//app.SetExternalMessagePump(false)
	//app.SetMultiThreadedMessageLoop(false)
	cef.BrowserWindow.Config.Title = "Energy - ipc multiple-window"
	// 关闭主窗口配置, 默认开启, 关闭后如果有多个窗口同时存在, 在关闭主窗口时应用进程不会结束，直到最后一个窗口关闭才结束应用进程
	cef.BrowserWindow.Config.EnableMainWindow = false

	//本地资源加载
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "assets",
		FS:         &assets,
		Home:       "multiple-window.html",
	}.Build())

	// 多窗口接收消息
	ipc.On("sendMessage", func(channel callback.IChannel, type_ int) {
		// 获得所有窗口
		infos := cef.BrowserWindow.GetWindowInfos()
		fmt.Println("windows-count:", len(infos))
		for _, info := range infos {
			// 将消息发送到目标窗口, 类型为JS接收
			iTarget := info.Target(target.TgJs)
			ipc.EmitTarget("receiveMessage", iTarget, time.Now().String())
			ipc.EmitTargetAndCallback("receiveMessage", iTarget, []any{"带有callback的触发事件: " + time.Now().String()}, func() {
				fmt.Println("target callback")
			})
		}
		// 关闭主窗口时 energy 会取最小的窗口ID设置为一个新的临时主窗口做为IPC通信发送
		ipc.Emit("receiveMessage", "测试当前新主窗口接收")
		ipc.EmitAndCallback("receiveMessage", []any{"带有callback的触发事件"}, func() {
			fmt.Println("callback")
		})
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("BeforePopup", browser.Identifier())
			popupWindow.SetSize(800, 600)
			return false
		})
	})
	//运行应用
	cef.Run(app)
}
