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
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, &resources)
	//创建应用
	cefApp := cef.NewApplication()
	//强制使用 VF 窗口
	//cefApp.SetExternalMessagePump(false)
	//cefApp.SetMultiThreadedMessageLoop(false)

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	if common.IsLinux() && cefApp.IsUIGtk3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	//这个示例演示了两种窗口组件的使用, LCL和VF

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			//LCL 窗口是我们创建的，需要我们自己管理窗口
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			bw.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
				*action = types.CaHide //隐藏窗口
				// 5秒后显示窗口
				go func() {
					println("LCL 窗口隐藏, 5秒后显示.")
					time.Sleep(time.Second * 5)
					window.RunOnMainThread(func() {
						bw.Show()
						bw.SetFocus()
					})
				}()
				return true //跳过默认事件, 如果想关闭窗口，这里返回false
			})
			bw.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) bool {
				return true //跳过默认事件, 如果想关闭窗口，这里返回false
			})
		} else if window.IsViewsFramework() {
			//VF 窗口是CEF自己创建的，这里我们只管Chromium的Close事件即可
			bw := window.AsViewsFrameworkBrowserWindow().BrowserWindow()
			bw.Chromium().SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
				*aAction = consts.CbaCancel //取消关闭 , 如果想关闭窗口， *aAction = consts.CbaClose
				window.Hide()               //隐藏窗口
				// 5秒后显示窗口
				go func() {
					println("VF 窗口隐藏, 5秒后显示.")
					time.Sleep(time.Second * 5)
					window.RunOnMainThread(func() {
						window.Show() //显示窗口
					})
				}()
			})
		}
	})
	//运行应用
	cef.Run(cefApp)
}
