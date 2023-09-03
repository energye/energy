package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/pkgs/assetserve"
)

//go:embed resources
var resources embed.FS

func main() {
	//Global initialization must be called
	cef.GlobalInit(nil, &resources)
	//Create an application
	cefApp := cef.NewApplication()
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
	//run app
	cef.Run(cefApp)
}
