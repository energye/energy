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
	"github.com/energye/energy/v2/cef"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

//go:embed resources
var resources embed.FS

func main() {
	// 打开 GIF 文件
	//gifPlay := NewGIFPlay()
	//gifPlay.LoadFile("E:\\SWT\\gopath\\src\\github.com\\energye\\energy\\examples\\window\\masking\\resources\\autumn.gif")
	//return
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	app := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "https://energy.yanghy.cn"

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		if window.IsLCL() {
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			// 页面呈现组件，默认让它隐藏, 也可不隐藏
			part := bw.WindowParent()
			part.SetVisible(false)

			// 创建一个 form 或 panel、或其它任意 IWinControl 组件模拟遮罩
			// form 可以设置透明度，但也没啥用
			mask := lcl.NewForm(bw)           // form它是有窗口句柄的，所以也可以使用 api 调整它
			mask.SetParent(bw)                // 显示在主窗口里的一个子窗口
			mask.SetBorderStyle(types.BsNone) // 因为是窗口所以要去掉标签栏，效果和panel差不多了
			mask.SetAlign(types.AlClient)     //铺满整个主窗口
			// 这是透明设置，只能form, 但是它的子组件也会跟着半透明
			//mask.SetAlphaBlend(true)
			//mask.SetAlphaBlendValue(80)
			// 创建一个gif播放组件
			gifPlay := cef.NewGIFPlay(mask)
			gifPlay.SetParent(mask)
			gifPlay.LoadFSFile("resources/autumn.gif") // gif图片可能有点要求吧？？
			// 这里是根据 gif 图片实际大小
			gifPlay.SetWidth(639)
			gifPlay.SetHeight(426)
			// 居中显示
			gifPlay.SetLeft(bw.Width()/2 - gifPlay.Width()/2)
			gifPlay.SetTop(bw.Height()/2 - gifPlay.Height()/2)
			// 窗口显示时开始播放gif动画
			mask.SetOnShow(func(sender lcl.IObject) {
				gifPlay.Start()
			})
			// 遮罩窗口关闭时，停止并释放gif
			mask.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
				gifPlay.Free()
			})
			// 显示form
			mask.Show()
			bw.SetOnResize(func(sender lcl.IObject) bool {
				if gifPlay.IsValid() {
					gifPlay.SetLeft(bw.Width()/2 - gifPlay.Width()/2)
					gifPlay.SetTop(bw.Height()/2 - gifPlay.Height()/2)
				}
				return false
			})
			// 页面加载进度
			bw.Chromium().SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
				v := progress * 100
				fmt.Println("OnLoadingProgressChange:", v)
				if v == 100 {
					cef.RunOnMainThread(func() {
						// 页面加载进度100%后关闭遮罩层
						if mask != nil {
							mask.Close()
							mask.Free()
							mask = nil
							// 最后显示浏览器组件
							part.SetVisible(true)
						}
					})
				}
			})
		}
	})
	//运行应用
	cef.Run(app)
}
