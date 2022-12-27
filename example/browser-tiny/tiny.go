package main

import (
	"fmt"
	"github.com/energye/energy/cef"
	"github.com/energye/energy/common"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
)

func main() {
	inits.Init(nil, nil)
	fmt.Println("main", common.Args.ProcessType())

	config := cef.NewApplicationConfig()
	config.SetMultiThreadedMessageLoop(false)
	config.SetExternalMessagePump(false)
	//config.SetChromeRuntime(true)
	application := cef.NewCEFApplication(config)
	application.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized()")
		component := lcl.NewComponent(nil)
		chromiumConfig := cef.NewChromiumConfig()
		chromium := cef.NewChromium(component, chromiumConfig)
		chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
			fmt.Println("OnBeforeClose")
			application.QuitMessageLoop()
		})
		chromium.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
			fmt.Println("OnTitleChange", title)
		})
		browserViewComponent := cef.NewBrowserViewComponent(component)
		windowComponent := cef.NewWindowComponent(component)
		windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
			fmt.Println("OnWindowCreated")
			b := chromium.CreateBrowserByBrowserViewComponent("https://www.baidu.com", browserViewComponent)
			fmt.Println("\tCreateBrowserByBrowserViewComponent", b)
			windowComponent.AddChildView(browserViewComponent)

			windowComponent.CenterWindow(cef.NewCefSize(1024, 768))
			browserViewComponent.RequestFocus()
			windowComponent.Show()
		})
		windowComponent.SetOnCanClose(func(sender lcl.IObject, window *cef.ICefWindow, aResult *bool) {
			fmt.Println("OnCanClose")
			chromium.TryCloseBrowser()
		})

		windowComponent.CreateTopLevelWindow()
	})
	application.SetOnGetDefaultClient(func(client *cef.ICefClient) {
		fmt.Println("OnGetDefaultClient")
	})
	process := application.StartMainProcess()
	fmt.Println("application.StartMainProcess()", process)
	if process {
		fmt.Println("application.RunMessageLoop()")
		application.RunMessageLoop()
	}
}
