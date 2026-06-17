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

// Init Webkit2Gtk Global initialization, invoked at application startup in main
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

// NewWKLoader creates and returns a WebKit2 loader instance
func NewWKLoader() wv.IWkLoader {
	if gGlobalWkLoader == nil {
		gGlobalWkLoader = wv.NewLoader(nil)
		// Dynamically select webkit2gtk 4.0 or 4.1 via webkit2gtk.Webkit2Ver()
		// When not using nocgo, attempt to use 4.1 first with priority
		// For cgo build, use build tag webkit2_4_1 to enable 4.1; default to 4.0
		if webkit2gtk.Webkit2Ver() == wvTypes.Wkv4_1 {
			gGlobalWkLoader.SetLoaderWebKit2DllPath(platformLinux.Libwebkit2gtk4_1_0)
			gGlobalWkLoader.SetLoaderJavascriptCoreDllPath(platformLinux.Libjavascriptcoregtk4_1_0)
			gGlobalWkLoader.SetLoaderSoupDllPath(platformLinux.Libsoup3_0_0)
			gGlobalWkLoader.SetWebkit2Version(wvTypes.Wkv4_1)
		}
	}
	return gGlobalWkLoader
}

// NewApplication creates and returns an Application instance
// If the global Application instance is not initialized yet, it will perform initialization setup
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

// Start starts the application
func (m *Application) Start() bool {
	v := m.StartWebKit2()
	return v
}
