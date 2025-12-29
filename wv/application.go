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
	"github.com/energye/lcl/lcl"
	wv "github.com/energye/wv/windows"
	"path/filepath"
)

// application
var (
	application    *Application
	globalWVLoader wv.IWVLoader
)

func NewWVLoader() wv.IWVLoader {
	if globalWVLoader == nil {
		if globalWVLoader = wv.GetGlobalWebView2Loader(); globalWVLoader != nil {
			return globalWVLoader
		} else {
			globalWVLoader = wv.NewLoader(nil)
			wv.SetGlobalWebView2Loader(globalWVLoader)
		}
	}
	return globalWVLoader
}

type Application struct {
	wv.IWVLoader
	mainWindow         *MainWindow
	onGetCustomSchemes wv.TLoaderGetCustomSchemesEvent
	options            Options
	localLoad          *LocalLoadResource
}

func NewApplication() *Application {
	if application == nil {
		application = &Application{
			IWVLoader:  NewWVLoader(),
			mainWindow: &MainWindow{},
		}
		webview2Home, wv2Loader := wv2load.Wv2Load()
		application.SetUserDataFolder(filepath.Join(webview2Home, "webview2Cache"))
		application.SetLoaderDllPath(wv2Loader)
		application.initDefaultEvent()
	}
	return application
}

func (m *Application) Run() {
	if m.StartWebView2() {
		lcl.Application.Initialize()
		lcl.Application.SetMainFormOnTaskBar(true)
		lcl.Application.NewForm(m.mainWindow)
		lcl.Application.Run()
	}
}

func (m *Application) SetOptions(options Options) {
	m.options = options
}

func (m *Application) SetLocalLoad(localLoad LocalLoad) {
	m.localLoad = NewLocalLoadResource(&localLoad)
	m.localLoad.LocalLoad = &localLoad
}

func (m *Application) SetOnGetCustomSchemes(fn wv.TLoaderGetCustomSchemesEvent) {
	m.onGetCustomSchemes = fn
}

func (m *Application) initDefaultEvent() {
	m.SetOnGetCustomSchemes(func(sender lcl.IObject, customSchemes *wv.IWVCustomSchemeInfoArrayWrap) {
		if m.onGetCustomSchemes != nil {
			m.onGetCustomSchemes(sender, customSchemes)
		}
		if m.localLoad != nil {
			if customSchemes == nil {
				*customSchemes = wv.NewCustomSchemeInfoArrayWrapWithInt(1)
				(*customSchemes).SetValue((*customSchemes).Size()-1, wv.TWVCustomSchemeInfo{
					SchemeName:            m.localLoad.Scheme,
					TreatAsSecure:         1,
					HasAuthorityComponent: 1,
				})
			}
		}
	})
}
