package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	"github.com/cyber-xxm/energy/v2/examples/common/tray"
)

func main() {
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	cefApp := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		tray.SYSTray(window)
	})
	cef.Run(cefApp)
}
