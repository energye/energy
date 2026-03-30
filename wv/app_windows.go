//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows

package wv

import (
	"fmt"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/application/pack"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool/exec"
	wv "github.com/energye/wv/windows"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// application
var (
	gApplication    *Application
	gGlobalWVLoader wv.IWVLoader
	gWebView2Loader = "WebView2Loader-%s.dll"
)

func init() {
	gWebView2Loader = fmt.Sprintf(gWebView2Loader, runtime.GOARCH)
}

// Init 全局初始化, 需手动调用的函数
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) *Application {
	lcl.Init(libs, resources)
	wv.Init()
	return NewApplication()
}

// NewWVLoader 创建并返回一个WebView2加载器实例
func NewWVLoader() wv.IWVLoader {
	if gGlobalWVLoader == nil {
		if gGlobalWVLoader = wv.GetGlobalWebView2Loader(); gGlobalWVLoader != nil {
			return gGlobalWVLoader
		} else {
			gGlobalWVLoader = wv.NewLoader(nil)
			wv.SetGlobalWebView2Loader(gGlobalWVLoader)
		}
	}
	return gGlobalWVLoader
}

func appLocalAppData() string {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData != "" {
		return localAppData
	}
	localAppData = os.Getenv("USERPROFILE")
	if localAppData != "" {
		return filepath.Join(localAppData, "AppData", "Local")
	}
	return os.TempDir()
}

// NewApplication 创建并返回单例Application实例
// 如果全局Application实例尚未初始化，则进行初始化设置
func NewApplication() *Application {
	if gApplication == nil {
		gApplication = &Application{
			IWVLoader: NewWVLoader(),
		}
		localAppData := appLocalAppData()
		if pack.Identity != "" {
			localAppData = filepath.Join(localAppData, pack.Identity)
		} else {
			fileName := strings.TrimSuffix(exec.Name, filepath.Ext(exec.Name))
			localAppData = filepath.Join(localAppData, "com.energy."+fileName)
		}
		gApplication.SetUserDataFolder(filepath.Join(localAppData, "WebView2"))
		dir, _ := filepath.Split(libname.LibName)
		wv2Loader := filepath.Join(dir, gWebView2Loader)
		gApplication.SetLoaderDllPath(wv2Loader)
		gApplication.initDefaultEvent()
		application.GApplication = &gApplication.Application
	}
	return gApplication
}

func DestroyGlobalLoader() {
	wv.DestroyGlobalWebView2Loader()
}

type Application struct {
	wv.IWVLoader
	application.Application
	onCustomSchemes TApplicationOnCustomSchemesEvent
}

func (m *Application) SetOnCustomSchemes(fn TApplicationOnCustomSchemesEvent) {
	m.onCustomSchemes = fn
}

func (m *Application) initDefaultEvent() {
	m.IWVLoader.SetOnGetCustomSchemes(func(sender lcl.IObject, customSchemeArray *wv.IWVCustomSchemeInfoArrayWrap) {
		if m.onCustomSchemes != nil {
			customSchemes := &TCustomSchemes{}
			m.onCustomSchemes(customSchemes)
			for _, scheme := range customSchemes.schemes {
				(*customSchemeArray).AddValue(wv.TWVCustomSchemeInfo{
					SchemeName:            scheme.Scheme,
					TreatAsSecure:         1,
					HasAuthorityComponent: 1,
				})
			}
		}
		if m.LocalLoad != nil {
			(*customSchemeArray).AddValue(wv.TWVCustomSchemeInfo{
				SchemeName:            m.LocalLoad.Scheme,
				TreatAsSecure:         1,
				HasAuthorityComponent: 1,
			})
		}
	})
}

// Start 启动应用程序
// 在所有设置后调用
func (m *Application) Start() bool {
	return m.StartWebView2()
}
