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
	"github.com/cyber-xxm/energy/v2/cef"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"time"
)

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, nil)
	//创建应用
	cefApp := cef.NewApplication()

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"
	cef.BrowserWindow.Config.ChromiumConfig().SetEnableMenu(false)
	cef.BrowserWindow.Config.EnableClose = false
	//这个示例演示了两种窗口组件的使用, LCL和VF

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			//LCL 窗口是我们创建的，需要我们自己管理窗口
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			bw.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
				*action = types.CaMinimize //最小化窗口
				// 5秒后还原窗口
				go func() {
					println("LCL 最小化窗口, 5秒后还原.")
					time.Sleep(time.Second * 5)
					window.RunOnMainThread(func() {
						window.Restore() //还原窗口
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
			bw.SetOnCloseQuery(func(win *cef.ICefWindow, window cef.IBrowserWindow, canClose *bool) bool {
				*canClose = false // 取消关闭 , 如果想关闭窗口 true
				window.Minimize() //最小化窗口
				// 5秒后显示窗口
				go func() {
					println("VF 最小化窗口, 5秒后还原.")
					time.Sleep(time.Second * 5)
					window.RunOnMainThread(func() {
						window.Restore() //还原窗口
					})
				}()
				return true
			})
		}
	})
	//运行应用
	cef.Run(cefApp)
}
