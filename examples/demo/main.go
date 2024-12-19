package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl/version"
)

//go:embed resources
var resources embed.FS

func main() {
	//Global initialization must be called
	cef.GlobalInit(nil, &resources)
	//Create an application
	app := cef.NewApplication()
	//http's url
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	//Security key and value settings for built-in static resource services
	assetserve.AssetsServerHeaderKeyName = "energy"
	assetserve.AssetsServerHeaderKeyValue = "energy"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		server := assetserve.NewAssetsHttpServer() //Built in HTTP service
		server.PORT = 22022                        //Service Port Number
		server.AssetsFSName = "resources"          //Resource folder with the same name
		server.Assets = &resources                 //Assets resources
		go server.StartHttpServer()
	})
	// run main process and main thread
	cef.BrowserWindow.SetBrowserInit(browserInit)
	//run app
	cef.Run(app)
}

// run main process and main thread
func browserInit(event *cef.BrowserEvent, window cef.IBrowserWindow) {
	// index.html ipc.emit("count", [count++])
	ipc.On("count", func(value int) {
		println("count", value)
	})
	// page load end
	event.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32, window cef.IBrowserWindow) {
		// index.html, ipc.on("osInfo", function(){...})
		println("osInfo", version.OSVersion.ToString())
		ipc.Emit("osInfo", version.OSVersion.ToString())
		var windowType string
		if window.IsLCL() {
			windowType = "LCL"
		} else {
			windowType = "VF"
		}
		// index.html, ipc.on("windowType", function(){...});
		ipc.Emit("windowType", windowType)
	})
}
