package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/golcl/lcl"
	//_ "net/http/pprof"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	app := cef.NewApplication()
	cef.BrowserWindow.Config.Url = "https://www.baidu.com"
	cef.BrowserWindow.Config.Title = "Energy - Print-Setting"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.ChromiumEvent().SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			settings := cef.PrintSettingsRef.New()
			var physicalSizeDeviceUnits *cef.TCefSize = &cef.TCefSize{Width: 100, Height: 200}
			var printableAreaDeviceUnits *cef.TCefRect = &cef.TCefRect{Width: 200, Height: 300, X: 90, Y: 80}
			settings.SetPrinterPrintableArea(physicalSizeDeviceUnits, printableAreaDeviceUnits, true)
			ranges := make([]cef.TCefRange, 2)
			ranges[0] = cef.TCefRange{From: 100, To: 200}
			ranges[1] = cef.TCefRange{From: 200, To: 400}
			settings.SetPageRanges(ranges)
			pageRanges := settings.GetPageRanges()
			fmt.Println("pageRanges", pageRanges, "count", settings.GetPageRangesCount())
		})
	})
	//运行应用
	cef.Run(app)
}
