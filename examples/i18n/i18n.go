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
	"github.com/cyber-xxm/energy/v2/cef/i18n"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/pkgs/assetserve"
	"github.com/energye/golcl/lcl/api"
)

//go:embed resources
var resources embed.FS

func main() {
	//全局初始化 每个应用都必须调用的
	cef.GlobalInit(nil, resources)
	//创建应用
	cefApp := cef.NewApplication()
	// 在这里设置i18n资源加载方式, 这里使用内置到exe方式
	i18n.SetLocalFS(resources, "resources")
	// 默认使用中文
	i18n.Switch(consts.LANGUAGE_zh_CN)
	//指定一个URL地址，或本地html文件目录
	cef.BrowserWindow.Config.Url = "http://localhost:22022/i18n.html"
	if common.IsLinux() && api.WidgetUI().IsGTK3() {
		cef.BrowserWindow.Config.IconFS = "resources/icon.png"
	} else {
		cef.BrowserWindow.Config.IconFS = "resources/icon.ico"
	}
	cef.BrowserWindow.Config.Title = "Energy i18n"

	//监听语言切换
	// t: 1 中文, 2 英文
	ipc.On("switch", func(t int) string {
		// 在页面使用按钮点击事件触发语言切换监听
		// 然后我们观察页面上的鼠标右键菜单 前进 后退
		if t == 1 {
			i18n.Switch(consts.LANGUAGE_zh_CN)
			return string(consts.LANGUAGE_zh_CN)
		} else if t == 2 {
			i18n.Switch(consts.LANGUAGE_en_US)
			return string(consts.LANGUAGE_en_US)
		}
		return ""
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
	cef.Run(cefApp)
}
