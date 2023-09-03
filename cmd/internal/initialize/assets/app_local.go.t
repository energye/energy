package main

import (
	"embed"
	"github.com/energye/energy/v2/cef"
)

//go:embed resources
var resources embed.FS

func main() {
	//Global initialization must be called
	cef.GlobalInit(nil, nil)
	//Create an application
	app := cef.NewApplication()
	//Local load resources
	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         &resources,
	}.Build())
	//run app
	cef.Run(app)
}
