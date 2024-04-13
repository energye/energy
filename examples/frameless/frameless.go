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
	"github.com/energye/energy/v2/cef/ipc"
	"github.com/energye/energy/v2/cef/ipc/context"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/examples/common/tray"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/pkgs/assetserve"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl/rtl/version"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

//go:embed resources
var resources embed.FS

//用于Go版本低于1.16
//go:generate energy bindata --fs --o=assets/assets.go --pkg=assets --paths=./resources/...

// go build -ldflags "-s -w"
func main() {
	//命令: go generate 生成内置资源
	//resources := assets.AssetFile()
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	app := cef.NewApplication()
	// 强制使用VF窗口模式
	//app.EnableVFWindow(true)

	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/index.html"
	cef.BrowserWindow.Config.EnableHideCaption = true
	// LCL macos 隐藏标题栏后，该选项不生效
	//cef.BrowserWindow.Config.EnableResize = true
	cef.BrowserWindow.Config.Title = "Energy Vue + ElementUI 示例"
	cef.BrowserWindow.Config.Width = 1366
	chromiumConfig := cef.BrowserWindow.Config.ChromiumConfig()
	chromiumConfig.SetEnableMenu(false)        //禁用右键菜单
	chromiumConfig.SetEnableWindowPopup(false) //禁用弹

	//监听窗口状态事件
	ipc.On("window-state", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		state := context.ArgumentList().GetIntByIndex(0)
		// 当前 ipc 在非UI线程中执行，窗口控制需要在UI线程执行
		cef.RunOnMainThread(func() {
			if state == 0 {
				fmt.Println("窗口最小化")
				bw.Minimize()
			} else if state == 1 {
				fmt.Println("窗口最大化/还原")
				bw.Maximize()
			} else if state == 3 {
				fmt.Println("全屏/退出全屏")
				if bw.IsFullScreen() {
					bw.ExitFullScreen()
				} else {
					bw.FullScreen()
				}
			}
		})
	})
	//监听窗口关闭事件
	ipc.On("window-close", func(context context.IContext) {
		bw := cef.BrowserWindow.GetWindowInfo(context.BrowserId())
		bw.CloseBrowserWindow()
	})
	ipc.On("os-info", func(context context.IContext) {
		fmt.Println("系统信息", version.OSVersion.ToString())
		context.Result(version.OSVersion.ToString())
	})

	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		//window.AsLCLBrowserWindow().FramelessForLine()
		if window.IsLCL() && common.IsWindows() {
			// 边框圆角, 仅LCL
			//window.AsLCLBrowserWindow().SetRoundRectRgn(10) // WindowParent 未铺满窗口会有严重的闪烁
			var (
				borderSpace    int32 = 2
				titleBarHeight int32 = 30
			)
			bw := window.AsLCLBrowserWindow().BrowserWindow()
			bw.SetColor(colors.ClBrown)             //给窗口随便设置一个颜色
			bw.WindowParent().RevertCustomAnchors() // 恢复到自定义定位，在下面代码中重新设置布局
			bw.WindowParent().SetTop(borderSpace + titleBarHeight)
			bw.WindowParent().SetLeft(borderSpace)
			bw.WindowParent().SetWidth(bw.Width() - borderSpace*2)
			bw.WindowParent().SetHeight(bw.Height() - borderSpace*2 - titleBarHeight)
			bw.WindowParent().SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))

			// 禁止窗口边框指定方向拖拽调整大小
			bda := bw.BroderDirectionAdjustments()
			bda = bda.Exclude(et.BdaTopLeft, et.BdaTop, et.BdaTopRight) //禁用 左上, 上, 右上
			bw.SetBroderDirectionAdjustments(bda)

			// 创建一个自定义chromium
			chromium := cef.NewChromiumBrowser(bw, window.Chromium().Config())
			bda = chromium.BroderDirectionAdjustments()
			bda = bda.Exclude(et.BdaBottomLeft, et.BdaBottom, et.BdaBottomRight) //禁用 左下, 下, 右下
			chromium.SetBroderDirectionAdjustments(bda)

			// 设置在窗口显示的位置
			part := chromium.WindowParent()
			part.SetHeight(titleBarHeight)
			part.SetTop(borderSpace)
			part.SetLeft(borderSpace)
			part.SetWidth(bw.Width() - borderSpace*2)
			part.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
			//chromium.SetSelfWindow(bw)// 如果 NewChromiumBrowser 参数 owner 非 IBrowserWindow 类型时需要手动指定
			chromium.Chromium().SetDefaultURL("http://localhost:22022/titlebar.html")

			chromium.RegisterDefaultEvent()      // 注册默认事件
			chromium.RegisterDefaultPopupEvent() // 注册默认弹出窗口事件
			chromium.CreateBrowser()             //最后创建浏览器

			//browserWindow := cef.NewBrowserWindow()
			//browserWindow.EnableAllDefaultEvent()
		}
		if window.IsLCL() {
			tray.LCLTray(window)
		}
	})
	//在主进程启动成功之后执行
	//在这里启动内置http服务
	//内置http服务需要使用 go:embed resources 内置资源到执行程序中
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022               //服务端口号
		server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
		server.Assets = resources
		go server.StartHttpServer()
	})
	//运行应用
	cef.Run(app)
}
