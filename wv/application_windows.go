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
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool/exec"
	wv "github.com/energye/wv/windows"
	"path/filepath"
	"runtime"
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

func NewApplication() *Application {
	if gApplication == nil {
		gApplication = &Application{
			IWVLoader: NewWVLoader(),
		}
		gApplication.SetUserDataFolder(filepath.Join(exec.AppDir(), "energyCache"))
		dir, _ := filepath.Split(libname.LibName)
		wv2Loader := filepath.Join(dir, gWebView2Loader)
		gApplication.SetLoaderDllPath(wv2Loader)
		gApplication.initDefaultEvent()
		application.GApplication = &gApplication.Application
	}
	return gApplication
}

type Application struct {
	wv.IWVLoader
	application.Application
	onCustomSchemes wv.TLoaderGetCustomSchemesEvent
}

func (m *Application) SetOnCustomSchemes(fn wv.TLoaderGetCustomSchemesEvent) {
	m.onCustomSchemes = fn
}

func (m *Application) initDefaultEvent() {
	m.IWVLoader.SetOnGetCustomSchemes(func(sender lcl.IObject, customSchemes *wv.IWVCustomSchemeInfoArrayWrap) {
		if m.onCustomSchemes != nil {
			m.onCustomSchemes(sender, customSchemes)
		}
		if m.LocalLoad != nil {
			(*customSchemes).AddValue(wv.TWVCustomSchemeInfo{
				SchemeName:            m.LocalLoad.Scheme,
				TreatAsSecure:         1,
				HasAuthorityComponent: 1,
			})
		}
	})
}

func (m *Application) Start() {
	m.StartWebView2()
}
