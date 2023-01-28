package traydemo

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 仅适用windows
//
// LCL + [CEF] 托盘 只适用 windows 基于html 和 ipc 实现功能
//
//推荐在windows或macosx中使用
func LCLCefTrayDemo(browserWindow cef.IBrowserWindow) {
	lclBw := browserWindow.AsLCLBrowserWindow().BrowserWindow()
	var url = "http://localhost:22022/tray-lcl-vf.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	cefTray := tray.AsCEFTray()
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIconFS("resources/icon.ico")
	tray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	ipc.IPC.Browser().On("tray-show-balloon", func(context ipc.IIPCContext) {
		fmt.Println("tray-show-balloon")
		cefTray.Notice("气泡标题", "气泡内容", 2000)
		cefTray.Hide()
		fmt.Println("tray-show-balloon end")
	})
	ipc.IPC.Browser().On("tray-show-main-window", func(context ipc.IIPCContext) {
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
	ipc.IPC.Browser().On("tray-close-main-window", func(context ipc.IIPCContext) {
		browserWindow.CloseBrowserWindow()
	})
	ipc.IPC.Browser().On("tray-show-message-box", func(context ipc.IIPCContext) {
		cef.QueueAsyncCall(func(id int) {
			lcl.ShowMessage("tray-show-message-box 提示消息")
		})
		cefTray.Hide()
	})
	//托盘 end
}

// 仅适用windows
//
// LCL + [VF] 托盘 只适用 windows 基于html 和 ipc 实现功能
//
// VF组件托盘，无法使用LCL相关组件
func LCLVFTrayDemo(browserWindow cef.IBrowserWindow) {
	vfBw := browserWindow.AsViewsFrameworkBrowserWindow().BrowserWindow()
	var url = "http://localhost:22022/tray-lcl-vf.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	vfTray := tray.AsViewsFrameTray()
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIconFS("resources/icon.ico")
	tray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	ipc.IPC.Browser().On("tray-show-balloon", func(context ipc.IIPCContext) {
		fmt.Println("tray-show-balloon")
		vfTray.Notice("气泡标题", "气泡内容", 2000)
		vfTray.Hide()
		fmt.Println("tray-show-balloon end")
	})
	var vfBwVisible = true
	ipc.IPC.Browser().On("tray-show-main-window", func(context ipc.IIPCContext) {
		if vfBwVisible {
			vfBw.Hide()
			vfBwVisible = false
		} else {
			vfBw.Show()
			vfBwVisible = true
		}

		vfTray.Hide()
	})
	ipc.IPC.Browser().On("tray-close-main-window", func(context ipc.IIPCContext) {
		browserWindow.CloseBrowserWindow()
	})
	ipc.IPC.Browser().On("tray-show-message-box", func(context ipc.IIPCContext) {
		//在VF窗口组件中无法使用LCL组件
		//cef.QueueAsyncCall(func(id int) {
		//	lcl.ShowMessage("tray-show-message-box 提示消息")
		//})
		vfTray.Hide()
	})
	//托盘 end
}
