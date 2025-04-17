//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/cyber-xxm/energy/v2/pkgs/channel"
	"github.com/energye/golcl/lcl"
)

func main() {
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	cefApp := cef.NewApplication()
	// 强制切换到VF窗口组件, 需要同时设置以下两个配置项为 false
	cefApp.EnableVFWindow(true)
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsViewsFramework() {
			trayDemo(window)
		}
	})
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	cef.Run(cefApp)
}

// LCL + [VF] 托盘 只适用 windows 基于html 和 ipc 实现功能
//
// VF组件托盘，无法使用LCL相关组件
func trayDemo(browserWindow cef.IBrowserWindow) {
	vfBw := browserWindow.AsViewsFrameworkBrowserWindow().BrowserWindow()
	var url = "http://localhost:22022/tray_lcl_vf.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	vfTray := tray.AsViewsFrameTray()
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIconFS("resources/icon.ico")
	tray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	ipc.On("tray-show-balloon", func(context channel.IIPCContext) {
		fmt.Println("tray-show-balloon")
		// 窗口控制在UI线程控制
		cef.RunOnMainThread(func() {
			vfTray.Notice("气泡标题", "气泡内容", 2000)
			vfTray.Hide()
			fmt.Println("tray-show-balloon end")
		})
	})
	var vfBwVisible = true
	ipc.On("tray-show-main-window", func(context channel.IIPCContext) {
		// 窗口控制在UI线程控制
		cef.RunOnMainThread(func() {
			if vfBwVisible {
				vfBw.Hide()
				vfBwVisible = false
			} else {
				vfBw.Show()
				vfBwVisible = true
			}
			vfTray.Hide()
		})
	})
	ipc.On("tray-close-main-window", func(context channel.IIPCContext) {
		// 窗口控制在UI线程控制
		cef.RunOnMainThread(func() {
			browserWindow.CloseBrowserWindow()
		})
	})
	ipc.On("tray-show-message-box", func(context channel.IIPCContext) {
		cef.RunOnMainThread(func() {
			vfTray.Hide()
			lcl.ShowMessage("tray-show-message-box 提示消息")
		})
	})
	//托盘 end
}
