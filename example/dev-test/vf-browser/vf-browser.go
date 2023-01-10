package main

import (
	"embed"
	"energye/systray"
	"energye/systray/icon"
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common/assetserve"
	"github.com/energye/energy/ipc"
	"github.com/energye/golcl/lcl"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalCEFInit(nil, &resources)
	//创建应用
	config := cef.NewApplicationConfig()
	config.SetMultiThreadedMessageLoop(false)
	config.SetExternalMessagePump(false)
	config.SetRemoteDebuggingPort(33333)
	cefApp := cef.NewApplication(config)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.CanDragFile = true
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		window.DisableResize()
		window.SetTitle("这里改变了窗口标题")
		window.SetSize(1600, 900)
		fmt.Println("cef.BrowserWindow.SetViewFrameBrowserInit", window)
		fmt.Println("LCL", window.AsLCLBrowserWindow(), "VF", window.AsViewsFrameworkBrowserWindow())
		event.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, regions *cef.TCefDraggableRegions) {
			fmt.Println("SetOnDraggableRegionsChanged", regions.RegionsCount(), "frame:", frame.Id, frame.Url)
		})
		event.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupWindow cef.IBrowserWindow, noJavascriptAccess *bool) bool {
			fmt.Println("IsViewsFramework:", popupWindow.IsViewsFramework())
			popupWindow.SetTitle("修改了标题: " + beforePopupInfo.TargetUrl)
			popupWindow.SetCenterWindow(true)
			popupWindow.EnableResize()
			popupWindow.SetSize(1600, 1600)
			browserWindow := popupWindow.AsViewsFrameworkBrowserWindow()
			browserWindow.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
				fmt.Println("popupWindow.SetOnWindowCreated", window)
			})
			//browserWindow.BrowserWindow().CreateTopLevelWindow()
			//browserWindow.BrowserWindow().HideTitle()
			fmt.Println("browserWindow:", browserWindow, browserWindow.WindowComponent().WindowHandle())
			return false
		})
		window.AsViewsFrameworkBrowserWindow().WindowComponent().SetOnWindowActivationChanged(func(sender lcl.IObject, window *cef.ICefWindow, active bool) {
			fmt.Println("SetOnWindowActivationChanged", active)
		})
		//设置隐藏窗口标题
		//window.HideTitle()
		window.Show()
		fmt.Println("SetBrowserInit 结束")
	})
	cef.BrowserWindow.SetBrowserInitAfter(func(window cef.IBrowserWindow) {
		bw := window.AsViewsFrameworkBrowserWindow().BrowserWindow()
		fmt.Println("handle", bw.WindowComponent().WindowHandle().ToPtr())
		//cefTray(window)
		trayMain()
		fmt.Println("SetBrowserInitAfter 结束")
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = &resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(cefApp)
}

// 托盘 只适用 windows 的系统托盘, 基于html 和 ipc 实现功能
func cefTray(browserWindow cef.IBrowserWindow) {
	var url = "http://localhost:22022/min-browser-tray.html"
	tray := browserWindow.NewCefTray(250, 300, url)
	if tray == nil {
		return
	}
	tray.SetTitle("任务管理器里显示的标题")
	tray.SetHint("这里是文字\n文字啊")
	tray.SetIconFS("resources/icon.ico")
	var s = false
	tray.SetOnClick(func(sender lcl.IObject) {
		s = !s
		if s {
			browserWindow.HideTitle()
		} else {
			browserWindow.ShowTitle()
		}
		browserWindow.Show()
	})
	tray.SetBalloon("气泡标题", "气泡内容", 2000)
	ipc.IPC.Browser().On("tray-show-balloon", func(context ipc.IIPCContext) {
		fmt.Println("tray-show-balloon")
		tray.ShowBalloon()
		tray.Hide()
	})
	ipc.IPC.Browser().On("tray-show-main-window", func(context ipc.IIPCContext) {
		browserWindow.Hide()
		tray.Hide()
	})
	ipc.IPC.Browser().On("tray-close-main-window", func(context ipc.IIPCContext) {
		browserWindow.CloseBrowserWindow()
	})
	ipc.IPC.Browser().On("tray-show-message-box", func(context ipc.IIPCContext) {
		//无法使用lcl组件
		//lcl.ShowMessage("提示?") //直接异常退出
		tray.Hide()
	})
	//托盘 end
}

func trayMain() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}
	go systray.Run(onReady, onExit)
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	addQuitItem()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Pretty awesome棒棒嗒")
		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItemCheckbox("Checked", "Check Me", true)
		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		// Sets the icon of a menu item. Only available on Mac.
		mEnabled.SetTemplateIcon(icon.Data, icon.Data)

		systray.AddMenuItem("Ignored", "Ignored")

		subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
		subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
		subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
		subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")

		systray.AddSeparator()
		mToggle := systray.AddMenuItem("Toggle", "Toggle some menu items")
		shown := true
		toggle := func() {
			if shown {
				subMenuBottom.Check()
				subMenuBottom2.Hide()
				mEnabled.Hide()
				shown = false
			} else {
				subMenuBottom.Uncheck()
				subMenuBottom2.Show()
				mEnabled.Show()
				shown = true
			}
		}
		mReset := systray.AddMenuItem("Reset", "Reset all items")

		for {
			select {
			case <-mChange.ClickedCh:
				mChange.SetTitle("I've Changed")
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("Unchecked")
				} else {
					mChecked.Check()
					mChecked.SetTitle("Checked")
				}
			case <-mEnabled.ClickedCh:
				mEnabled.SetTitle("Disabled")
				mEnabled.Disable()
			case <-subMenuBottom2.ClickedCh:
				panic("panic button pressed")
			case <-subMenuBottom.ClickedCh:
				toggle()
			case <-mReset.ClickedCh:
				systray.ResetMenu()
				addQuitItem()
			case <-mToggle.ClickedCh:
				toggle()
			}
		}
	}()
}
