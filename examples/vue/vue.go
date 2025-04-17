package main

import (
	"embed"
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/examples/vue/src"
	"github.com/cyber-xxm/energy/v2/examples/vue/src/env"
)

//go:embed resources
var resources embed.FS

// go build -ldflags="-s -w" -tags="prod"
// go build -ldflags="-s -w"
func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)

	app := cef.NewApplication()
	cef.BrowserWindow.Config.Title = "Energy Vue"
	cef.BrowserWindow.Config.EnableHideCaption = true
	if env.ENV == "DEV" {
		println("DEV")
		cef.BrowserWindow.Config.Url = "http://localhost:5173"
		app.SetUseMockKeyChain(true) // MacOS
	} else {
		// go build -ldflags="-s -w" -tags="prod"
		println("PRODUCE")
		cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
			ResRootDir: "resources",
			FS:         resources,
		}.Build())
		cfg := cef.BrowserWindow.Config.ChromiumConfig()
		cfg.SetEnableMenu(false)
		cfg.SetEnableDevTools(false)
		cfg.SetEnableViewSource(false)
	}

	cef.SetBrowserProcessStartAfterCallback(src.BrowserProcessStart)
	cef.BrowserWindow.SetBrowserInit(src.BrowserInit)

	//运行应用
	cef.Run(app)
}
