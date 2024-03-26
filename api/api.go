//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/energy/v2/api/imports"
	"github.com/energye/energy/v2/api/internal"
	"github.com/energye/energy/v2/api/internal/cef"
	"github.com/energye/energy/v2/api/internal/lcl"
	"github.com/energye/energy/v2/api/internal/winapi"
	"github.com/energye/energy/v2/api/internal/wv"
)

var (
	releaseCallback    func()
	canWidgetSetInit   bool
	customWidgetImport *imports.Imports // 自定义组件初始化导入
	winapiImport       *imports.Imports // winapi 导入
	liblclImport       *imports.Imports // LCL 导入
	liblclPreDefImport *imports.Imports // LCL 预定义导入
	cefImport          *imports.Imports // CEF 导入
	cefPreDefImport    *imports.Imports // CEF 预定义导入
	wvImport           *imports.Imports // WV 导入
	wvPreDefImport     *imports.Imports // WV 预定义导入
)

func init() {
	customWidgetImport = new(imports.Imports) // 自定义组件初始化导入
	winapiImport = new(imports.Imports)       // winapi 导入
	liblclImport = new(imports.Imports)       // LCL 导入
	liblclPreDefImport = new(imports.Imports) // LCL 预定义导入
	cefImport = new(imports.Imports)          // CEF 导入
	cefPreDefImport = new(imports.Imports)    // CEF 预定义导入
	wvImport = new(imports.Imports)           // WV 导入
	wvPreDefImport = new(imports.Imports)     // WV 预定义导入
}

// APIInit API初始化
func APIInit() {
	// 加载动态链接库
	uiLib = loadUILib()

	// 自定义组件初始化导入
	customWidgetImport.SetDll(uiLib)
	internal.InitCustomWidgetImport(customWidgetImport)

	// winapi 导入
	winapiImport.SetDll(uiLib)
	winapi.InitWinAPIPreDefsImport(winapiImport)

	// LCL 导入
	liblclImport.SetDll(uiLib)
	internal.InitLCLAutoGenImport(liblclImport)
	// LCL 预定义导入
	liblclPreDefImport.SetDll(uiLib)
	lcl.InitPreDefsImport(liblclPreDefImport)

	// CEF 导入
	cefImport.SetDll(uiLib)
	internal.InitCEFAutoGenImport(cefImport)
	// CEF 预定义导入
	cefPreDefImport.SetDll(uiLib)
	cef.InitCEFPreDefsImport(cefPreDefImport)

	// WV 导入
	wvImport.SetDll(uiLib)
	internal.InitWVAutoGenImport(wvImport)
	// WV 预定义导入
	wvPreDefImport.SetDll(uiLib)
	wv.InitWVPreDefsImport(wvPreDefImport)
}

// LibRelease 在energy中释放
func LibRelease() {
	if releaseCallback != nil {
		releaseCallback()
	}
	CustomWidgetSetFinalization()
	// 开启了finalizerOn选项后，以防止关闭库后GC还没开始调用。
	callGC()
	// 运行结束后就结束close掉lib，不然他不会关掉的
	closeLib()
}

// SetReleaseCallback 应用运行结束后释放资源之前执行
func SetReleaseCallback(fn func()) {
	if releaseCallback == nil {
		releaseCallback = fn
	}
}

// LCL 导入表
func LCL() imports.CallImport {
	return liblclImport
}

// LCLPreDef 预定义LcL导入表
func LCLPreDef() imports.CallImport {
	return liblclPreDefImport
}

// CEF 导入表
func CEF() imports.CallImport {
	return cefImport
}

// CEFPreDef 导入表
func CEFPreDef() imports.CallImport {
	return cefPreDefImport
}

// WinAPI 导入表
func WinAPI() imports.CallImport {
	return winapiImport
}

// WV webview2 导入表
func WV() imports.CallImport {
	return wvImport
}

// WVPreDef 预定义webview2导入表
func WVPreDef() imports.CallImport {
	return wvPreDefImport
}
