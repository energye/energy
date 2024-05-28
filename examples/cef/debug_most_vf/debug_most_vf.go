package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/common"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"os"
	"path/filepath"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	app := cef.NewCefApplication()
	cef.SetGlobalCEFApp(app)
	// CEF message loop
	app.SetExternalMessagePump(false)
	app.SetMultiThreadedMessageLoop(false)
	if common.IsDarwin() {
		app.SetUseMockKeyChain(true)
		app.InitLibLocationFromArgs()
		// MacOS使用扩展消息泵
		cef.AddCrDelegate()
		cef.GlobalWorkSchedulerCreate(nil)
		app.SetOnScheduleMessagePumpWork(nil)
		if !process.Args.IsMain() {
			// MacOS 多进程时，需要调用StartSubProcess来启动子进程
			subStart := app.StartSubProcess()
			fmt.Println("subStart:", subStart, process.Args.ProcessType())
			app.Free()
			return
		}
	} else { // MacOS不需要设置CEF框架目录，它是一个固定的目录结构
		// 非MacOS需要指定CEF框架目录，执行文件在CEF目录则不需要设置
		// 指定 CEF Framework
		frameworkDir := os.Getenv("ENERGY_HOME")
		app.SetFrameworkDirPath(frameworkDir)
		app.SetResourcesDirPath(frameworkDir)
		app.SetLocalesDirPath(filepath.Join(frameworkDir, "locales"))
	}
	if common.IsLinux() {
		// 这是一个解决“GPU不可用错误”问题的方法 linux
		// https://bitbucket.org/chromiumembedded/cef/issues/2964/gpu-is-not-usable-error-during-cef
		app.SetDisableZygote(true)
	}
	app.SetOnContextInitialized(func() {
		fmt.Println("SetOnContextInitialized")
		component := lcl.NewComponent(nil)
		chromium := cef.NewChromium(component)
		windowComponent := cef.NewCEFWindowComponent(component)
		viewComponent := cef.NewCEFBrowserViewComponent(component)
		url := "https://gitee.com/energye/energy"
		windowComponent.SetOnWindowCreated(func(sender cef.IObject, window cef.ICefWindow) {
			ok := chromium.CreateBrowserByBrowserViewComponent(url, viewComponent, nil, nil)
			fmt.Println("SetOnWindowCreated CreateBrowserByBrowserViewComponent:", true)
			if ok {
				windowComponent.AddChildView(viewComponent.BrowserView())
				viewComponent.RequestFocus()
				windowComponent.Show()
			}
		})
		chromium.SetOnBeforeClose(func(sender cef.IObject, browser cef.ICefBrowser) {
			app.QuitMessageLoop()
		})
		windowComponent.CreateTopLevelWindow()
	})
	mainStart := app.StartMainProcess()
	fmt.Println("mainStart:", mainStart, process.Args.ProcessType())
	if mainStart {
		// 结束应用后释放资源
		app.RunMessageLoop()
		fmt.Println("app free")
		app.Free()
	}
	///usr/share/lazarus/3.2.0/lcl/graphics.pp
}
