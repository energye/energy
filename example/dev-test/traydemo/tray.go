package traydemo

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

// LCL + [CEF|VF] 托盘 只适用 windows 基于html 和 ipc 实现功能
func LCLCefTrayDemo(browserWindow cef.IBrowserWindow) {
	var lclbw *cef.LCLBrowserWindow
	var vfbw *cef.ViewsFrameworkBrowserWindow
	if browserWindow.IsLCL() {
		lclbw = browserWindow.AsLCLBrowserWindow().BrowserWindow()
	} else if browserWindow.IsViewsFramework() {
		vfbw = browserWindow.AsViewsFrameworkBrowserWindow().BrowserWindow()
	}
	var url = "http://localhost:22022/tray-lcl-vf.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	asCEFTray := tray.AsCEFTray()
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIconFS("resources/icon.ico")
	tray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	ipc.IPC.Browser().On("tray-show-balloon", func(context ipc.IIPCContext) {
		fmt.Println("tray-show-balloon")
		asCEFTray.Notice("气泡标题", "气泡内容", 2000)
		asCEFTray.Hide()
	})
	ipc.IPC.Browser().On("tray-show-main-window", func(context ipc.IIPCContext) {
		if lclbw != nil {
			vb := !lclbw.Visible()
			lclbw.SetVisible(vb)
			if vb {
				if lclbw.WindowState() == types.WsMinimized {
					lclbw.SetWindowState(types.WsNormal)
				}
				lclbw.Focused()
			}
		} else if vfbw != nil {
		}
		asCEFTray.Hide()
	})
	ipc.IPC.Browser().On("tray-close-main-window", func(context ipc.IIPCContext) {
		browserWindow.CloseBrowserWindow()
	})
	ipc.IPC.Browser().On("tray-show-message-box", func(context ipc.IIPCContext) {
		cef.QueueAsyncCall(func(id int) {
			lcl.ShowMessage("tray-show-message-box 提示消息")
		})
		asCEFTray.Hide()
	})
	//托盘 end
}

// LCL组件托盘, 适用windows和macosx, 不支持linux因GTK2和GTK3共存问题,
func LCLTrayDemo(browserWindow cef.IBrowserWindow) {
	window := browserWindow.AsLCLBrowserWindow().BrowserWindow()
	//托盘 windows linux macos 系统托盘
	newTray := window.NewTray()
	newTray.SetTitle("任务管理器里显示的标题")
	newTray.SetHint("这里是文字\n文字啊")
	newTray.SetIconFS("resources/icon.ico")
	tray := newTray.AsLCLTray()
	menu1 := tray.AddMenuItem("父菜单", nil)
	menu1.Add(tray.NewMenuItem("子菜单", func() {
		lcl.ShowMessage("子菜单点击 提示消息")
	}))
	tray.AddMenuItem("显示气泡", func() {
		//linux下有些问题
		tray.Notice("气泡标题", "气泡内容", 2000)
	})
	tray.AddMenuItem("显示/隐藏", func() {
		vis := window.Visible()
		cef.BrowserWindow.GetWindowInfo(1)
		window.SetVisible(!vis)
	})
	tray.AddMenuItem("退出", func() {
		browserWindow.CloseBrowserWindow()
	})
	//托盘 end
	tray.Show()
}

//系统托盘 和LCL组件差不多,但不如LCL组件的好用，适用 windows,linux,macosx
func SysTrayDemo(browserWindow cef.IBrowserWindow) {
	sysTray := browserWindow.NewSysTray()
	sysTray.SetIconFS("resources/icon.ico")
	sysTray.SetHint("中文hint\n换行中文")
	sysTray.SetOnClick(func() {
		fmt.Println("SetOnClick")
	})
	tray := sysTray.AsSysTray()
	check := tray.AddMenuItem("check")
	check.Check()
	not := tray.AddMenuItem("通知")
	not.Click(func() {
		tray.Notice("标题", "内notice 是一个跨平台的系统通知库\nnotice 是一个跨平台的系统通知库", 1000)
	})
	enable := tray.AddMenuItem("启用/禁用")
	enable.Click(func() {
		fmt.Println("启用/禁用 点击")
	})
	tray.AddSeparator()
	menuItem := tray.AddMenuItem("1级菜单1", func() {
		fmt.Println("1级菜单1")
	})
	menuItem.SetIconFS("resources/icon.ico")
	tray.AddSeparator()
	item := tray.AddMenuItem("1级菜单2")
	item.AddSubMenu("2级子菜单1")
	sub2Menu := item.AddSubMenu("2级子菜单2")
	sub2Menu.AddSubMenu("3级子菜单1")
	tray.AddSeparator()
	tray.AddMenuItem("退出", func() {
		fmt.Println("退出")
		browserWindow.CloseBrowserWindow()
	})

	sysTray.Show()
	//测试图标切换
	go func() {
		var b bool
		for {
			time.Sleep(time.Second * 2)
			b = !b
			if b {
				sysTray.SetHint(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				sysTray.SetIconFS("resources/icon_1.ico")
				menuItem.SetIconFS("resources/icon_1.ico")
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Enable()
				check.Check()
			} else {
				sysTray.SetHint(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				sysTray.SetIconFS("resources/icon.ico")
				menuItem.SetIconFS("resources/icon.ico")
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Disable()
				check.Uncheck()
			}
		}
	}()
}
