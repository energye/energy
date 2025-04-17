package tray

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"time"
)

// SYSTray 系统托盘 和LCL组件差不多,但不如LCL组件的好用，适用 windows,linux, macosx
//
// 主要给linux提供的，推荐linux中使用, 非linux 使用 lcl 实现的它更好用
func SYSTray(browserWindow cef.IBrowserWindow) {
	sysTray := browserWindow.NewSysTray()
	var (
		icon1 string
		icon2 string
	)
	if common.IsLinux() {
		icon1 = "resources/icon.png"
		icon2 = "resources/icon_red.png"
	} else {
		icon1 = "resources/icon.ico"
		icon2 = "resources/icon_red.ico"
	}
	sysTray.SetIconFS(icon1)
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
	menuItem.SetIconFS(icon1)
	tray.AddSeparator()
	item := tray.AddMenuItem("1级菜单2")
	item.AddSubMenu("2级子菜单1")
	sub2Menu := item.AddSubMenu("2级子菜单2")
	sub2Menu.AddSubMenu("3级子菜单1")
	tray.AddSeparator()
	tray.AddMenuItem("退出", func() {
		fmt.Println("退出")
		// 关闭所有窗口
		for _, info := range cef.BrowserWindow.GetWindowInfos() {
			info.CloseBrowserWindow()
		}
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
				sysTray.SetIconFS(icon1)
				menuItem.SetIconFS(icon1)
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Enable()
				check.Check()
			} else {
				sysTray.SetHint(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				sysTray.SetIconFS(icon2)
				menuItem.SetIconFS(icon2)
				enable.SetLabel(fmt.Sprintf("%d\n%v", time.Now().Second(), b))
				enable.Disable()
				check.Uncheck()
			}
		}
	}()
}
