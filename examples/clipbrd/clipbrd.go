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
	"github.com/cyber-xxm/energy/v2/cef/exception"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"path/filepath"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/clipbrd.html"
	if common.IsLinux() && api.WidgetUI().IsGTK3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	//主进程启动成功之后回调
	cef.SetBrowserProcessStartAfterCallback(func(b bool) {
		fmt.Println("主进程启动 创建一个内置http服务")
		//通过内置http服务加载资源
		server := assetserve.NewAssetsHttpServer()
		server.PORT = 22022
		server.AssetsFSName = "resources" //必须设置目录名
		server.Assets = resources
		go server.StartHttpServer()
	})
	cef.BrowserWindow.SetBrowserInit(func(event *cef.BrowserEvent, window cef.IBrowserWindow) {
		exception.SetOnException(func(message string) {
			fmt.Println(message)
		})

		ipc.On("setText", func(text string) {
			lcl.Clipboard.SetAsText(text)
			println("set-text", text)
		})
		ipc.On("getText", func() string {
			text := lcl.Clipboard.AsText()
			println("get-text", text)
			return text
		})
		ipc.On("setPic", func() string {
			println("setPic")
			mem := lcl.NewMemoryStream()
			defer mem.Free()
			mem.LoadFromFSFile("resources/bg.bmp")
			mem.SetPosition(0)
			// 预定义格式
			format := lcl.PredefinedClipboardFormat(types.PcfBitmap)
			fmt.Println("format:", format)
			if lcl.Clipboard.SetFormat(format, mem) {
				return "设置成功"
			} else {
				return "设置失败"
			}
		})
		ipc.On("getPic", func() string {
			println("getPic")
			if !lcl.Clipboard.HasPictureFormat() {
				return "false"
			}
			bmpFormat := lcl.Clipboard.FindPictureFormatID()
			mem := lcl.NewMemoryStream()
			defer mem.Free()
			if lcl.Clipboard.GetFormat(bmpFormat, mem) {
				mem.SetPosition(0)
				wd := consts.CurrentExecuteDir
				path := filepath.Join(wd, "bg.bmp")
				mem.SaveToFile(path)
				return path
			}
			return "获取失败"
		})
	})
	//运行应用
	cef.Run(cefApp)
}
