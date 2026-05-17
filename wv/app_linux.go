//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux

package wv

import (
	"github.com/energye/energy/v3/application"
	platformLinux "github.com/energye/energy/v3/platform/linux"
	"github.com/energye/energy/v3/platform/linux/webkit2gtk"
	"github.com/energye/lcl/lcl"
	wv "github.com/energye/wv/linux"
	wvTypes "github.com/energye/wv/types/linux"
)

var (
	gApplication    *Application
	gGlobalWkLoader wv.IWkLoader
	gWk2Context     wv.IWkWebContext
)

// Init 全局初始化, 需手动调用的函数
func Init() *Application {
	lcl.Init()
	wv.Init()
	return NewApplication()
}

type Application struct {
	wv.IWkLoader
	application.Application
	onCustomSchemes TApplicationOnCustomSchemesEvent
}

func (m *Application) SetOnCustomSchemes(fn TApplicationOnCustomSchemesEvent) {
	m.onCustomSchemes = fn
}

// NewWKLoader 创建并返回一个Webkit2加载器实例
func NewWKLoader() wv.IWkLoader {
	if gGlobalWkLoader == nil {
		gGlobalWkLoader = wv.NewLoader(nil)
		// 通过 webkit2gtk.Webkit2Ver() 动态控制使用 webkit2gtk 4.0 或 4.1
		// 当使用非 nocgo 优先尝试从 4.1 开始尝试
		// 如果使用 cgo 需要使用条件编译 webkit2_4_1 选择 4.1, 默认 4.0
		if webkit2gtk.Webkit2Ver() == wvTypes.Wkv4_1 {
			gGlobalWkLoader.SetLoaderWebKit2DllPath(platformLinux.Libwebkit2gtk4_1_0)
			gGlobalWkLoader.SetLoaderJavascriptCoreDllPath(platformLinux.Libjavascriptcoregtk4_1_0)
			gGlobalWkLoader.SetLoaderSoupDllPath(platformLinux.Libsoup3_0_0)
			gGlobalWkLoader.SetWebkit2Version(wvTypes.Wkv4_1)
		}
	}
	return gGlobalWkLoader
}

// NewApplication 创建并返回单例Application实例
// 如果全局Application实例尚未初始化，则进行初始化设置
func NewApplication() *Application {
	if gApplication == nil {
		gApplication = &Application{
			IWkLoader: NewWKLoader(),
		}
		application.GApplication = &gApplication.Application
	}
	return gApplication
}

func DestroyGlobalLoader() {
}

// Start 启动应用程序
// 在所有设置后调用
func (m *Application) Start() bool {
	v := m.StartWebKit2()
	return v
}
