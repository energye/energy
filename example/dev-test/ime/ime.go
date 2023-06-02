package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
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
	cef.BrowserWindow.Config.Title = "Energy - IME"
	cef.BrowserWindow.Config.IconFS = "resources/icon.ico"

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		event.ChromiumEvent().SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
			fmt.Println("OnLoadEnd")
			var (
				underlines                       []*cef.TCefCompositionUnderline
				replacementRange, selectionRange *cef.TCefRange
			)
			underlines = make([]*cef.TCefCompositionUnderline, 2)
			underlines[0] = &cef.TCefCompositionUnderline{
				Range:           cef.TCefRange{From: 100, To: 100},
				Color:           244,
				BackgroundColor: 122,
				Thick:           400,
				Style:           consts.CEF_CUS_DOT,
			}
			underlines[1] = &cef.TCefCompositionUnderline{
				Range:           cef.TCefRange{From: 200, To: 200},
				Color:           233,
				BackgroundColor: 100,
				Thick:           500,
				Style:           consts.CEF_CUS_DASH,
			}
			replacementRange = &cef.TCefRange{From: 50, To: 50}
			selectionRange = &cef.TCefRange{From: 150, To: 150}
			window.Chromium().IMESetComposition("文本", underlines, replacementRange, selectionRange)
			//window.Chromium().IMECommitText()
		})
	})
	//运行应用
	cef.Run(app)
}
