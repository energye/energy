package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/exception"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/cef/ipc/callback"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	"github.com/cyber-xxm/energy/v2/examples/common/tray"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/energye/golcl/lcl"
	"time"
)

// go build -ldflags "-s -w"
func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.LDebug)
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	exception.SetOnException(func(message string) {
		fmt.Println("Global Exception message", message)
	})
	//创建应用
	app := cef.NewApplication()
	// 强制使用VF窗口
	//app.SetExternalMessagePump(false)
	//app.SetMultiThreadedMessageLoop(false)
	cef.BrowserWindow.Config.Title = "Energy - ipc multiple-window"
	// 禁用主窗口配置, 该属性默认开启, 禁用后如果有多个窗口同时存在, 在关闭主窗口时应用进程不会结束，直到最后一个窗口关闭才结束应用进程
	cef.BrowserWindow.Config.EnableMainWindow = false

	//本地资源加载
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         demoCommon.ResourcesFS(),
		Home:       "multiple-window.html",
	}.Build())

	// 多窗口发送/接收消息
	ipc.On("sendMessage", func(channel callback.IChannel, type_ int) {
		var ok bool
		// 获得所有窗口
		infos := cef.BrowserWindow.GetWindowInfos()
		fmt.Println("window-count:", len(infos))
		for _, info := range infos {
			// 将消息发送到目标窗口, 类型为JS接收
			iTarget := info.Target()
			ok = ipc.EmitTarget("receiveMessage", iTarget, time.Now().String())
			println("ipc.EmitTarget", ok, iTarget.BrowserId())
			ok = ipc.EmitTargetAndCallback("receiveMessage", iTarget, []interface{}{"带有callback的触发事件: " + time.Now().String()}, func() {
				fmt.Println("target callback")
			})
			println("ipc.EmitTargetAndCallback", ok, iTarget.BrowserId())
		}

		// 主窗口接收, 主窗口被关闭后发送无效
		ok = ipc.Emit("receiveMessage", "测试当前新主窗口接收")
		println("ipc.Emit", ok)
		ok = ipc.EmitAndCallback("receiveMessage", []interface{}{"带有callback的触发事件"}, func() {
			fmt.Println("callback")
		})
		println("ipc.EmitAndCallback", ok)
	})
	// 刷新所有窗口
	ipc.On("refresh", func() {
		infos := cef.BrowserWindow.GetWindowInfos()
		for _, info := range infos {
			// 将消息发送到目标窗口, 类型为JS接收
			iTarget := info.Target()
			var ok = ipc.EmitTarget("refresh", iTarget)
			println("ipc.EmitTarget-refresh", ok)
		}
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		// 弹出窗口处理
		event.SetOnBeforePopup(func(sender lcl.IObject, popupWindow cef.IBrowserWindow, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			fmt.Println("BeforePopup", browser.Identifier())
			popupWindow.RunOnMainThread(func() {
				popupWindow.SetSize(800, 600)
			})
			return false
		})
		// 多窗口消息发送
		go func() {
			var count = 1
			for true {
				// 一秒一次, 计数循环
				time.Sleep(time.Second)
				// 获得当前所有窗口信息
				infos := cef.BrowserWindow.GetWindowInfos()
				for _, info := range infos {
					iTarget := info.Target()
					ipc.EmitTarget("count", iTarget, count)
				}
				count++
			}
		}()
		// 托盘，如果关闭主窗口模式，在关闭托盘所创建的窗口时
		// 我们实际应该将这个窗口隐藏掉，因此目前还有点问题
		if window.IsLCL() {
			// LCL窗口 Window, MacOS
			tray.LCLTray(window)
		} else {
			// VF窗口 默认Linux
			tray.SYSTray(window)
		}
	})
	//运行应用
	cef.Run(app)
}
