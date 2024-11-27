package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/process"
	"github.com/energye/energy/v2/consts"
	. "github.com/energye/energy/v2/examples/tiny-browser/cefclient/browse"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/pkgs/libname"
	"os"
	"path/filepath"
)

var (
	app    *cef.TCEFApplication
	vf     *ViewsFramework
	width  int32 = 800
	height int32 = 600
)

type ViewsFramework struct {
	*lcl.TComponent
	chromium    cef.IChromium
	window      *cef.TCEFWindowComponent
	browserView *cef.TCEFBrowserViewComponent
	homePage    string
}

func main() {
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	app = cef.CreateApplication()
	app.SetFrameworkDirPath(os.Getenv("ENERGY_HOME"))
	app.SetMultiThreadedMessageLoop(false)
	app.SetExternalMessagePump(false)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetLocale(consts.LANGUAGE_zh_CN)
	app.SetTouchEvents(consts.STATE_ENABLED)
	app.SetDisableZygote(true)
	fmt.Println("libname.LibName:", libname.LibName)
	fmt.Println("WidgetUI:", api.WidgetUI())
	fmt.Println("ChromeVersion:", app.ChromeVersion())
	fmt.Println("LibCefVersion:", app.LibCefVersion())

	kPrefWindowRestore := "cefclient.window_restore"
	app.SetOnRegisterCustomPreferences(func(type_ consts.TCefPreferencesType, registrar *cef.TCefPreferenceRegistrarRef) {
		fmt.Println("OnRegisterCustomPreferences ProcessType:", process.Args.ProcessType())
		if type_ == consts.CEF_PREFERENCES_TYPE_GLOBAL {
			dict := cef.DictionaryValueRef.New()
			dict.SetInt(kPrefWindowRestore, int32(consts.CEF_SHOW_STATE_NORMAL))
			value := cef.ValueRef.New()
			value.SetDictionary(dict)
			registrar.AddPreference(kPrefWindowRestore, value)
		}
	})
	app.SetOnAlreadyRunningAppRelaunch(func(commandLine *cef.ICefCommandLine, currentDirectory string) bool {
		fmt.Println("OnAlreadyRunningAppRelaunch ProcessType:", process.Args.ProcessType())
		// 在此处创建一个新窗口

		// 重新启动处理好了
		return true
	})
	app.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized ProcessType:", process.Args.ProcessType())
		fmt.Println("  GetScreenDPI:", cef.GetScreenDPI(), "GetDeviceScaleFactor:", cef.GetDeviceScaleFactor())
	})
	if app.StartMainProcess() {
		fmt.Println("StartMainProcess Success")
		StartServer()
		// 创建窗口
		vf = new(ViewsFramework)
		vf.homePage = "http://localhost:22022"
		vf.Create()
		vf.window.CreateTopLevelWindow()
		app.RunMessageLoop()
	}
}

func (m *ViewsFramework) Create() {
	m.TComponent = lcl.NewComponent(nil)
	m.chromium = cef.NewChromium(m, nil)
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("OnBeforeClose")
		app.QuitMessageLoop()
	})
	m.browserView = cef.BrowserViewComponentRef.New(m)
	//m.browserView.SetPreferAccelerators(true)
	m.window = cef.WindowComponentRef.New(m)
	m.chromium.SetOnAutoResize(func(sender lcl.IObject, browser *cef.ICefBrowser, newSize *cef.TCefSize) bool {
		fmt.Println("OnAutoResize", newSize)
		return true
	})

	m.window.SetOnCanClose(func(window *cef.ICefWindow, result *bool) {
		fmt.Println("OnCanClose:", *result)
		*result = m.chromium.TryCloseBrowser()
	})
	m.window.SetOnGetInitialBounds(func(window *cef.ICefWindow, result *cef.TCefRect) {
		fmt.Println("OnGetInitialBounds")
		result.Width = width
		result.Height = height
	})

	m.window.SetOnGetInitialBounds(func(window *cef.ICefWindow, result *cef.TCefRect) {
		result.Width = 800
		result.Height = 600
		fmt.Println("OnGetInitialBounds")
	})
	m.window.SetOnAccelerator(func(window *cef.ICefWindow, commandId int32, result *bool) {
		fmt.Println("OnAccelerator commandId:", commandId)
		//表示已处理，否则还会执行多次
		app.QuitMessageLoop()
		*result = true
	})
	m.window.SetOnWindowCreated(func(window *cef.ICefWindow) {
		fmt.Println("OnWindowCreated")
		m.window.SetID(ID_WINDOW)
		//window.SetThemeColor()
		m.window.ThemeChanged()

		menuBar := NewMenuBar(m.window) // 顶部菜单栏
		toolBar := NewToolBar(m.window) // 顶部工具栏

		if m.chromium.CreateBrowserByBrowserViewComponent(m.homePage, m.browserView, nil, nil) {
			fmt.Println("ChromeToolbar:", m.browserView.ChromeToolbar().IsValid())
			windowLayout := m.window.SetToBoxLayout(cef.TCefBoxLayoutSettings{Horizontal: 0, BetweenChildSpacing: 2})
			fmt.Println("windowLayout:", windowLayout.IsValid())
			windowLayout.SetFlexForView(m.browserView.BrowserView().AsView(), 1)

			// 菜单栏, 创建菜单，并添加到菜单栏中
			menuBar.CreateFileMenuItems()
			menuBar.CreateTestMenuItems()

			// 工具栏, 创建工具组件，并添加到工具栏中
			toolBar.CreateToolComponent()

			// 菜单栏添加到窗口
			m.window.AddChildView(menuBar.EnsureMenuPanel().AsView())
			// 工具栏添加到窗口
			m.window.AddChildView(toolBar.EnsureToolPanel().AsView())

			// 浏览器view添加到窗口
			//m.window.AddChildView(m.browserView.BrowserView().AsView())

			m.window.Layout()
			// 窗口居中
			display := m.window.Display()
			if display.IsValid() {
				workArea := display.WorkArea()
				position := &cef.TCefPoint{
					X: ((workArea.Width - width) / 2) + workArea.X,
					Y: ((workArea.Height - height) / 2) + workArea.Y,
				}
				m.window.SetBounds(cef.NewCefRect(position.X, position.Y, width, height))
			}
			m.window.Show()
			//m.window.ShowAsBrowserModalDialog(m.browserView.BrowserView())
			m.browserView.RequestFocus()
		}
	})
}
