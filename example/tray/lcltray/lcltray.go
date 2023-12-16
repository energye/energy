package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/example/common/tray"
)

//go:embed resources
var resources embed.FS

func main() {
	cef.GlobalInit(nil, &resources)
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
