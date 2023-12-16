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
			wp.RevertCustomAnchors() // 恢复默认根据窗口定位, energy默认是把当前browser铺满整个窗口
			// 重新设置位置大小
			wp.SetLeft(50)
			wp.SetTop(50)
			wp.SetWidth(400)
			wp.SetHeight(600)
			// 重新设置根据窗口位置大小自动调整browser窗口
			wp.SetAnchors(types.NewSet(types.AkTop, types.AkLeft, types.AkRight, types.AkBottom))

			// 创建第二个ChromiumBrowser, 也可创建更多个
			chromium := cef.NewChromiumBrowser(bw, nil)
			rect2 := chromium.WindowParent().BoundsRect()
			rect2.Left = 460                                                                             // 当前这个browser相对当前窗口的左边距
			rect2.Top = 50                                                                               // 当前这个browser相对当前窗口的上边距
			rect2.SetSize(400, 600)                                                                      // browser大小
			chromium.WindowParent().SetBoundsRect(rect2)                                                 // 设置browser位置和大小
			chromium.WindowParent().SetAnchors(types.NewSet(types.AkTop, types.AkRight, types.AkBottom)) // 根据窗口大小browser自动调整相对窗口位置
			chromium.Chromium().SetDefaultURL("https://www.baidu.com")                                   // 设置加载的页面地址
			// 通过chromium获取回调事件
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
