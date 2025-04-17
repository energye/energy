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
	"github.com/energye/golcl/lcl/types"
)

func main() {
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	cefApp := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
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

// LCL + CEF 托盘只适用 windows 基于html 和 ipc 实现
func trayDemo(browserWindow cef.IBrowserWindow) {
	lclBw := browserWindow.AsLCLBrowserWindow().BrowserWindow()
	var url = "http://localhost:22022/tray_lcl_vf.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	cefTray := tray.AsCEFTray()
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIconFS("resources/icon.ico")
	tray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	ipc.On("tray-show-balloon", func(context channel.IIPCContext) {
		fmt.Println("tray-show-balloon")
		cefTray.Notice("气泡标题", "气泡内容", 2000)
		cefTray.Hide()
		fmt.Println("tray-show-balloon end")
	})
	ipc.On("tray-show-main-window", func(context channel.IIPCContext) {
		vb := !lclBw.Visible()
		lclBw.SetVisible(vb)
		if vb {
			if lclBw.WindowState() == types.WsMinimized {
				lclBw.SetWindowState(types.WsNormal)
			}
			lclBw.Focused()
		}
		cefTray.Hide()
	})
	ipc.On("tray-close-main-window", func(context channel.IIPCContext) {
		browserWindow.CloseBrowserWindow()
	})
	ipc.On("tray-show-message-box", func(context channel.IIPCContext) {
		cef.QueueAsyncCall(func(id int) {
			lcl.ShowMessage("tray-show-message-box 提示消息")
		})
		cefTray.Hide()
	})
	//托盘 end
}
