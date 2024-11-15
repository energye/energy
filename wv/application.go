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
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/lcl"
	"github.com/energye/wv/windows"
)

// application
var application *Application

type Application struct {
	wv.IWVLoader
	mainWindow         *MainWindow
	onGetCustomSchemes wv.TOnLoaderGetCustomSchemesEvent
	localLoad          *LocalLoad
}

func NewApplication() *Application {
	if application == nil {
		application = &Application{
			IWVLoader:  wv.GlobalWebView2Loader(),
			mainWindow: &MainWindow{},
		}
		application.SetLoaderDllPath(libname.GetLibPath("WebView2Loader.dll"))
		application.defaultEvent()
	}
	return application
}

func (m *Application) Run() {
	if m.StartWebView2() {
		lcl.Application.Initialize()
		lcl.Application.SetMainFormOnTaskBar(true)
		lcl.Application.CreateForm(m.mainWindow)
		lcl.Application.Run()
	}
}

func (m *Application) SetOptions(options Options) {
	m.mainWindow.options = options
}

func (m *Application) SetLocalLoad(localLoad LocalLoad) {
	m.localLoad = &localLoad
}

func (m *Application) SetOnWindowCreate(fn OnCreate) {
	m.mainWindow.onWindowCreate = fn
}

func (m *Application) SetOnWindowAfterCreate(fn OnCreate) {
	m.mainWindow.onWindowAfterCreate = fn
}

func (m *Application) SetOnGetCustomSchemes(fn wv.TOnLoaderGetCustomSchemesEvent) {
	m.onGetCustomSchemes = fn
}

func (m *Application) defaultEvent() {
	m.IWVLoader.SetOnGetCustomSchemes(func(sender wv.IObject, customSchemes *wv.TWVCustomSchemeInfoArray) {
		if m.onGetCustomSchemes != nil {
			m.onGetCustomSchemes(sender, customSchemes)
		}
		if m.localLoad != nil {
			*customSchemes = append(*customSchemes, &wv.TWVCustomSchemeInfo{
				SchemeName:            m.localLoad.Scheme,
				TreatAsSecure:         true,
				HasAuthorityComponent: true,
			})
		}
	})
}
