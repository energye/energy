//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/examples/window/masking/mask"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	//cef.BrowserWindow.Config.Url = "https://github.com/energye/energy"
	cef.BrowserWindow.Config.Url = "https://gitee.com/energye/energy"

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			part := bw.WindowParent()
			if common.IsLinux() { // Linux 只能在Gtk2使用，还有些问题。不能像 windows 和 MacOS 那样
				part.RevertCustomAnchors()
				part.SetWidth(1)
				part.SetHeight(1)
			}
			maskForm := mask.Create(bw)
			// 页面加载进度, 控制何时关闭遮罩
			bw.Chromium().SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
				v := int(progress * 100)
				fmt.Println("OnLoadingProgressChange:", v)
				// 更新UI组件或状态必须在UI线程中运行
				cef.RunOnMainThread(func() {
					if maskForm.IsValid() {
						maskForm.Progress(v)
						if v == 100 {
							maskForm.Hide()
							maskForm.Stop()

							if common.IsLinux() {
								part.SetWidth(bw.Width())
								part.SetHeight(bw.Height())
								part.SetAlign(types.AlClient)
								part.UpdateSize()
								bw.Chromium().NotifyMoveOrResizeStarted()
							}
						}
					}
				})
			})
			bw.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
				cef.RunOnMainThread(func() {
					maskForm.Show()
					maskForm.Start()
				})
				return false
			})
		}
	})
	//运行应用
	cef.Run(app)
}
