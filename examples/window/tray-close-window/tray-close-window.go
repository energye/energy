package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

func main() {
	logger.SetEnable(true)
	logger.SetLevel(logger.LDebug)
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	app := cef.NewApplication()
	if common.IsDarwin() {
		app.SetUseMockKeyChain(true)
	}
	cef.BrowserWindow.Config.Url = "https://www.baidu.com/"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//默认隐藏主窗口
		lcl.Application.SetShowMainForm(false)

		lclWindow := window.AsLCLBrowserWindow().BrowserWindow()
		// 托盘关闭事件中=true
		var isTrayClose bool
		// 创建一个子窗口，默认显示它
		wp := cef.NewWindowProperty()
		wp.Url = "https://www.baidu.com/"
		wp.Title = "测试子窗口"
		browserWindow := cef.NewLCLBrowserWindow(nil, wp, nil)
		browserWindow.SetWidth(600)
		browserWindow.SetHeight(400)
		browserWindow.SetShowInTaskBar()
		browserWindow.EnableAllDefaultEvent()
		browserWindow.SetWindowType(consts.WT_POPUP_SUB_BROWSER)
		browserWindow.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
			fmt.Println("SetOnTitleChange", wp.Title, title)
		})
		newTray := lclWindow.NewTray()
		newTray.SetTitle("任务管理器里显示的标题")
		newTray.SetHint("这里是文字\n文字啊")
		newTray.SetIconFS("resources/icon.png")
		newTray.SetOnDblClick(func() {
			isShow := lclWindow.Visible()
			lclWindow.SetVisible(!isShow)
		})
		tray := newTray.AsLCLTray()
		tray.AddMenuItem("显示", func(sender lcl.IObject) {
			window.Show()
		})
		tray.AddMenuItem("退出", func(sender lcl.IObject) {
			isTrayClose = true
			window.CloseBrowserWindow()
		})

		lclWindow.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) bool {
			fmt.Println("关闭主窗口")
			if isTrayClose {
				fmt.Println("从托盘关闭窗口")
				// 托盘点击退出，关闭主窗口，退出应用
				return false
			} else {
				fmt.Println("其它关闭方式都给窗口隐藏掉")
				// 点击窗口关闭按钮，不关闭它，把主窗口隐藏掉
				*canClose = false
				window.Hide()
				return true
			}
		})

		browserWindow.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
			fmt.Println("关闭子窗口")
			window.Show()  // 子窗口关闭 显示主窗口
			newTray.Show() //显示托盘
			return false   //执行默认逻辑
		})
		browserWindow.Show() //显示子窗口
	})
	//运行应用
	cef.Run(app)
}
