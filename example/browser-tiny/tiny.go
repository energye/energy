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
		component := lcl.NewComponent(nil)
		fmt.Println("OnContextInitialized(", component)

	})
	application.SetOnGetDefaultClient(func(client *cef.ICefClient) {
		fmt.Println("OnGetDefaultClient")
	})
	//chromiumConfig := cef.NewChromiumConfig()
	//chromium := cef.NewChromium(component, chromiumConfig)
	//chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
	//	fmt.Println("OnBeforeClose")
	//})
	//
	//windowComponent := cef.NewWindowComponent(component)
	//windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *cef.ICefWindow) {
	//	fmt.Println("OnWindowCreated")
	//})
	//windowComponent.SetOnGetInitialBounds(func(sender lcl.IObject, window *cef.ICefWindow, aResult *cef.TCefRect) {
	//	fmt.Println("OnGetInitialBounds")
	//})
	process := application.StartMainProcess()
	fmt.Println("application.StartMainProcess()", process)
	if process {
		fmt.Println("application.RunMessageLoop()")
		application.RunMessageLoop()
	}
	fmt.Println("end")
}
