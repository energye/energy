package wv

import (
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tools/exec"
	"github.com/energye/wv/wv"
	"path/filepath"
)

// application
var application *Application

type Application struct {
	wv.IWVLoader
	mainWindow         *BrowserWindow
	onGetCustomSchemes wv.TOnLoaderGetCustomSchemesEvent
}

func NewApplication() *Application {
	if application == nil {
		application = &Application{
			IWVLoader: wv.GlobalWebView2Loader(),
			mainWindow: &BrowserWindow{
				options: Options{
					Width:  800,
					Height: 600,
				},
			},
		}
		webView2Loader, _ := filepath.Split(libname.LibName)
		webView2Loader = filepath.Join(webView2Loader, "WebView2Loader.dll")
		application.SetUserDataFolder(filepath.Join(exec.CurrentDir, "EnergyCache"))
		application.SetLoaderDllPath(webView2Loader)
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
	if options.Width == 0 {
		options.Width = 800
	}
	if options.Height == 0 {
		options.Height = 600
	}
	m.mainWindow.options = options
}

func (m *Application) SetOnWindowCreate(fn OnWindowCreate) {
	m.mainWindow.onWindowCreate = fn
}

func (m *Application) SetOnGetCustomSchemes(fn wv.TOnLoaderGetCustomSchemesEvent) {
	m.onGetCustomSchemes = fn
}

func (m *Application) defaultEvent() {
	m.IWVLoader.SetOnGetCustomSchemes(func(sender wv.IObject, customSchemes *wv.TWVCustomSchemeInfoArray) {
		if m.onGetCustomSchemes != nil {
			m.onGetCustomSchemes(sender, customSchemes)
		}
		*customSchemes = append(*customSchemes, &wv.TWVCustomSchemeInfo{
			SchemeName:            "fs",
			TreatAsSecure:         true,
			HasAuthorityComponent: true,
		})
	})
}
