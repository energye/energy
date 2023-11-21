package main

import (
	"github.com/energye/energy/v2/cef"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://www.runoob.com/"
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			// 在这里改变主窗口的默认行为
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			wp := bw.WindowParent()
			wp.RevertCustomAnchors() // 恢复默认根据窗口定位,
			wp.SetLeft(50)
			wp.SetTop(50)
			wp.SetWidth(400)
			wp.SetHeight(600)
			// 设置根据窗口位置大小自动调整browser窗口
			wp.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))

			//
			chromium := cef.NewChromiumBrowser(bw, nil)
			rect2 := chromium.WindowParent().BoundsRect()
			rect2.Left = 460
			rect2.Top = 50
			rect2.SetSize(400, 600)
			chromium.WindowParent().SetBoundsRect(rect2)
			chromium.WindowParent().SetAnchors(types.NewSet(types.AkTop, types.AkRight, types.AkBottom))
			chromium.Chromium().SetDefaultURL("https://www.baidu.com")
			chromium.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
				// 在这里更新 WindowParent 大小, 以保证渲染到窗口中
				chromium.WindowParent().UpdateSize()
				return false
			})
		}
	})
	//运行应用
	cef.Run(app)
}
