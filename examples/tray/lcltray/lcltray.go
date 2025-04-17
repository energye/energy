package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/examples/common"
	"github.com/cyber-xxm/energy/v2/examples/common/tray"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
)

func main() {
	cef.GlobalInit(nil, common.ResourcesFS())
	cefApp := cef.NewApplication()
	cefApp.SetUseMockKeyChain(true)
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			tray.LCLTray(window)
		}
	})
	cef.Run(cefApp)
}
