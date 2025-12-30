// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package wv

import (
	"fmt"
	"github.com/energye/energy/v3/application"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool/exec"
	wv "github.com/energye/wv/windows"
	"path/filepath"
)

// application
var (
	gApplication    *Application
	gGlobalWVLoader wv.IWVLoader
)

const _WebView2Loader = "WebView2Loader.dll"

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

type Application struct {
	wv.IWVLoader
	application.Application
	onGetCustomSchemes wv.TLoaderGetCustomSchemesEvent
}

func NewWebviewApplication() *Application {
	if gApplication == nil {
		gApplication = &Application{
			IWVLoader: NewWVLoader(),
		}
		gApplication.SetUserDataFolder(filepath.Join(exec.AppDir(), "energyCache"))
		dir, _ := filepath.Split(libname.LibName)
		wv2Loader := filepath.Join(dir, _WebView2Loader)
		gApplication.SetLoaderDllPath(wv2Loader)
		gApplication.initDefaultEvent()
		application.GApplication = &gApplication.Application
	}
	return gApplication
}

func (m *Application) SetOnGetCustomSchemes(fn wv.TLoaderGetCustomSchemesEvent) {
	m.onGetCustomSchemes = fn
}

func (m *Application) initDefaultEvent() {
	m.IWVLoader.SetOnGetCustomSchemes(func(sender lcl.IObject, customSchemes *wv.IWVCustomSchemeInfoArrayWrap) {
		if m.onGetCustomSchemes != nil {
			m.onGetCustomSchemes(sender, customSchemes)
		}
		if m.LocalLoad != nil {
			//*customSchemes = wv.NewCustomSchemeInfoArrayWrapWithInt(1)
			fmt.Println("size:", (*customSchemes).Size())
			(*customSchemes).SetValue(-1, wv.TWVCustomSchemeInfo{
				SchemeName:            m.LocalLoad.Scheme,
				TreatAsSecure:         1,
				HasAuthorityComponent: 1,
			})
			fmt.Println("size:", (*customSchemes).Size())
		}
	})
}

func (m *Application) Start() {
	m.StartWebView2()
}
