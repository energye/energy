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
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		tray.SYSTray(window)
	})
	cef.Run(cefApp)
}
