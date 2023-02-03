package traydemo

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/golcl/lcl"
	"time"
)

// LCL组件托盘, 适用windows和macosx, 不支持linux因GTK2和GTK3共存问题,
func LCLTrayDemo(browserWindow cef.IBrowserWindow) {
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
	iconItem.Bitmap().SetSize(32, 32)      //图标情况调整大小
	iconItem.Bitmap().SetTransparent(true) //透明
	icon := lcl.NewIcon()
	icon.LoadFromFSFile("resources/icon_1.ico")
	iconItem.Bitmap().Canvas().Draw(0, 0, icon) //画上去
	tray.TrayMenu().Items().Add(iconItem)

	menu1.Add(tray.NewMenuItem("子菜单", func() {
		lcl.ShowMessage("子菜单点击 提示消息")
	}))
	tray.AddMenuItem("显示气泡", func() {
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
//
//推荐linux中使用
func SysTrayDemo(browserWindow cef.IBrowserWindow) {
	sysTray := browserWindow.NewSysTray()
	if common.IsLinux() {
		sysTray.SetIconFS("resources/icon.png")
	} else {
		sysTray.SetIconFS("resources/icon.ico")
	}
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
