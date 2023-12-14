package tray

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/golcl/lcl"
)

// LCLTray LCL组件托盘, 适用windows和macosx
// 目前 不支持linux因GTK2和GTK3共存问题, 以后解决GTK3问题将支持
// 或linux cef 106以前的版本使用GTK2的的也可使用该托盘
func LCLTray(browserWindow cef.IBrowserWindow) {
	window := browserWindow.AsLCLBrowserWindow().BrowserWindow()
	//托盘 windows linux macos 系统托盘
	newTray := window.NewTray()
	newTray.SetTitle("任务管理器里显示的标题")
	newTray.SetHint("这里是文字\n文字啊")
	if common.IsLinux() {
		newTray.SetIconFS("resources/icon.png")
	} else {
		newTray.SetIconFS("resources/icon.ico")
	}
	tray := newTray.AsLCLTray()
	menu1 := tray.AddMenuItem("父菜单", nil)
	//带图标的菜单
	iconItem := tray.NewMenuItem("带个图标", nil)
	iconItem.Bitmap().SetSize(16, 16)      //图标情况调整大小
	iconItem.Bitmap().SetTransparent(true) //透明
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/icon.ico")
	iconItem.Bitmap().Canvas().Draw(0, 0, icon) //画上去
	tray.TrayMenu().Items().Add(iconItem)

	menu1.Add(tray.NewMenuItem("子菜单", func() {
		lcl.ShowMessage("子菜单点击 提示消息")
	}))
	tray.AddMenuItem("显示气泡", func() {
		tray.Notice("气泡标题", "气泡内容", 2000)
	})
	tray.AddMenuItem("显示/隐藏", func() {
		// 所有窗口
		for _, info := range cef.BrowserWindow.GetWindowInfos() {
			window := info.AsLCLBrowserWindow().BrowserWindow()
			window.SetVisible(!window.Visible())
			if window.Visible() { //之后的显示状态
				//window.SetFocus()
			}
		}
	})
	tray.AddMenuItem("退出", func() {
		// 关闭所有窗口
		for _, info := range cef.BrowserWindow.GetWindowInfos() {
			info.CloseBrowserWindow()
		}
	})
	//托盘 end
	tray.Show()
}
