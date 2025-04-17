package tray

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// LCLTray LCL组件托盘, 适用windows和macosx
// 目前 不支持linux因GTK2和GTK3共存问题, 以后解决GTK3问题将支持
// 或linux cef 106以前的版本使用GTK2的的也可使用该托盘
func LCLTray(browserWindow cef.IBrowserWindow) {
	window := browserWindow.AsLCLBrowserWindow().BrowserWindow()
	//托盘 windows linux macos 系统托盘
	var (
		icon1 string
		icon2 string
	)
	if common.IsLinux() || common.IsDarwin() {
		icon1 = "resources/icon.png"
		icon2 = "resources/icon_red.png"
	} else {
		icon1 = "resources/icon.ico"
		icon2 = "resources/icon_red.ico"
	}

	newTray := window.NewTray()
	newTray.SetTitle("任务管理器里显示的标题")
	newTray.SetHint("这里是文字\n文字啊")
	newTray.SetIconFS(icon1)
	// 托盘图标事件
	newTray.SetOnClick(func() {
		fmt.Println("click")
	})
	newTray.SetOnDblClick(func() {
		fmt.Println("dclick")
	})
	tray := newTray.AsLCLTray()
	// image
	imageMenu := tray.NewMenuItem("", nil)
	imageMenu.SetOnMeasureItem(func(sender lcl.IObject, aCanvas *lcl.TCanvas, width, height *int32) {
		*height = 60
		*width = 270
	})
	imageMenuPng := lcl.NewIcon()
	imageMenuPng.LoadFromFSFile("resources/icon_red.ico")
	imageMenu.Bitmap().Canvas().Draw(0, 0, imageMenuPng) //画上去
	imageMenuPng.Free()
	tray.TrayMenu().Items().Add(imageMenu)
	//--
	tray.AddMenuItem("-", nil)
	menu1 := tray.AddMenuItem("父菜单", nil)
	//带图标的菜单
	iconItem := tray.NewMenuItem("带个图标", nil)
	icon := lcl.NewPngImage()
	icon.LoadFromFSFile("resources/icon.png")
	iconItem.Bitmap().Assign(icon) //.Canvas().Draw(0, 0, icon) //画上去
	tray.TrayMenu().Items().Add(iconItem)

	menu1.Add(tray.NewMenuItem("子菜单", func(s lcl.IObject) {
		lcl.ShowMessage("子菜单点击 提示消息")
	}))
	tray.AddMenuItem("显示气泡", func(s lcl.IObject) {
		tray.Notice("气泡标题", "气泡内容", 2000)
	})
	tray.AddMenuItem("显示/隐藏", func(s lcl.IObject) {
		// 所有窗口
		for _, info := range cef.BrowserWindow.GetWindowInfos() {
			window := info.AsLCLBrowserWindow().BrowserWindow()
			window.SetVisible(!window.Visible())
			if window.Visible() { //之后的显示状态
				//window.SetFocus()
			}
		}
	})
	// --
	tray.AddMenuItem("-", nil)
	var check *lcl.TMenuItem
	check = tray.NewMenuItem("Check", func(s lcl.IObject) {
		check.SetChecked(!check.Checked())
	})
	check.SetChecked(true)
	tray.TrayMenu().Items().Add(check)
	// --
	tray.AddMenuItem("-", nil)
	var (
		radio1     *lcl.TMenuItem
		radio2     *lcl.TMenuItem
		radio3     *lcl.TMenuItem
		groupIndex uint8 = 1
	)
	radio1 = tray.NewMenuItem("Radio1", func(s lcl.IObject) {
		fmt.Println("Radio1")
		lcl.AsMenuItem(s).SetChecked(true)
	})
	radio1.SetGroupIndex(groupIndex)
	radio1.SetRadioItem(true)
	tray.TrayMenu().Items().Add(radio1)
	radio2 = tray.NewMenuItem("Radio2", func(s lcl.IObject) {
		fmt.Println("Radio2")
		lcl.AsMenuItem(s).SetChecked(true)
	})
	radio2.SetGroupIndex(groupIndex)
	radio2.SetRadioItem(true)
	tray.TrayMenu().Items().Add(radio2)
	radio3 = tray.NewMenuItem("Radio3", func(s lcl.IObject) {
		fmt.Println("Radio3")
		lcl.AsMenuItem(s).SetChecked(true)
	})
	radio3.SetGroupIndex(groupIndex)
	radio3.SetRadioItem(true)
	tray.TrayMenu().Items().Add(radio3)
	// --
	tray.AddMenuItem("-", nil)
	var (
		showMenu *lcl.TMenuItem
		wotMenu  *lcl.TMenuItem
	)
	var wot = func(sender lcl.IObject) {
		wotMenu.SetVisible(!wotMenu.Visible())
		if wotMenu.Visible() {
			showMenu.SetCaption("Hide WOT")
		} else {
			showMenu.SetCaption("Show WOT")
		}
		fmt.Println(wotMenu.Visible())
	}
	showMenu = tray.AddMenuItem("Hide WOT", wot)
	wotMenu = tray.AddMenuItem("Me WOT", wot)
	// --
	tray.AddMenuItem("-", nil)

	tray.AddMenuItem("退出", func(s lcl.IObject) {
		// 关闭所有窗口
		for _, info := range cef.BrowserWindow.GetWindowInfos() {
			info.CloseBrowserWindow()
		}
	})
	//托盘 end
	tray.Show()
	// 图标切换 定时器
	var trayICON bool
	timer := lcl.NewTimer(window)
	timer.SetInterval(1000)
	timer.SetOnTimer(func(sender lcl.IObject) {
		if window.IsClosing() {
			// 如果窗口关闭或关闭中
			return
		}
		if trayICON {
			newTray.SetIconFS(icon1)
		} else {
			newTray.SetIconFS(icon2)
		}
		trayICON = !trayICON
		fmt.Println("timer", trayICON)
	})
	timer.SetEnabled(true)
	// 窗口关闭，关闭定时器
	window.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
		timer.SetEnabled(false)
		return false
	})
}
