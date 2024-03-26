package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	demoCommon "github.com/energye/energy/v2/examples/cef/common"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
)

func main() {
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	app := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "http://chrome.360.cn/html5_labs/demos/dnd/"
	cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	cef.BrowserWindow.Config.Title = "ENERGY - Drag File"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.SetOnDragEnter(func(sender lcl.IObject, browser *cef.ICefBrowser, dragData *cef.ICefDragData, mask types.TCefDragOperations, window cef.IBrowserWindow, result *bool) {
			if mask&types.DRAG_OPERATION_LINK == types.DRAG_OPERATION_LINK {
				fmt.Println("SetOnDragEnter", mask&types.DRAG_OPERATION_LINK, dragData.IsLink(), dragData.IsFile(), "GetFileName:", dragData.GetFileName(), "GetFileNames:", dragData.GetFileNames())
				*result = false
			} else {
				*result = true
			}
		})
	})
	cef.Run(app)
}
