package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
)

func main() {
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	app := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "http://chrome.360.cn/html5_labs/demos/dnd/"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.Title = "ENERGY - Drag File"
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = demoCommon.ResourcesFS()
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnDragEnter(func(sender lcl.IObject, browser *cef.ICefBrowser, dragData *cef.ICefDragData, mask consts.TCefDragOperations, window cef.IBrowserWindow, result *bool) {
			if mask&consts.DRAG_OPERATION_LINK == consts.DRAG_OPERATION_LINK {
				fmt.Println("SetOnDragEnter", mask&consts.DRAG_OPERATION_LINK, dragData.IsLink(), dragData.IsFile(), "GetFileName:", dragData.GetFileName(), "GetFileNames:", dragData.GetFileNames())
				*result = false
			} else {
				*result = true
			}
		})
	})
	cef.Run(app)
}
