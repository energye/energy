package main

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/example/common/tray"
)

func main() {
	cef.GlobalInit(nil, nil)
	cefApp := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			tray.LCLTray(window)
		}
	})
	cef.Run(cefApp)
}
