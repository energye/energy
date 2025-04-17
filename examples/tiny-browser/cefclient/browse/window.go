package browse

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/tiny-browser/cefclient/views_style"
	"github.com/energye/golcl/lcl"
)

var (
	window *ViewsFramework
	app    *cef.TCEFApplication
	width  int32 = 800
	height int32 = 600
)

type ViewsFramework struct {
	*lcl.TComponent
	chromium    cef.IChromium
	window      *cef.TCEFWindowComponent
	browserView *cef.TCEFBrowserViewComponent
	homePage    string
	titleBar    *TitleBar
	menuBar     *MenuBar
	toolBar     *ToolBar
}

func Application() *cef.TCEFApplication {
	if app == nil {
		app = cef.CreateApplication()
	}
	return app
}

func MainWindow() *ViewsFramework {
	if window == nil {
		window = new(ViewsFramework)
		window.homePage = "http://localhost:22022"
		window.Create()
		window.window.CreateTopLevelWindow()
		app.RunMessageLoop()
	}
	return window
}

func (m *ViewsFramework) Create() {
	m.TComponent = lcl.NewComponent(nil)
	m.chromium = cef.NewChromium(m, nil)
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("OnBeforeClose")
		app.QuitMessageLoop()
	})
	var minimumWindowSize = cef.TCefSize{Width: 100, Height: 100}

	m.browserView = cef.BrowserViewComponentRef.New(m)
	m.browserView.SetID(ID_BROWSER_VIEW)
	//m.browserView.SetPreferAccelerators(true)
	m.window = cef.WindowComponentRef.New(m)
	m.window.SetID(ID_WINDOW)
	m.window.SetTitle("energy cefclient")
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
	})
	m.chromium.SetOnAutoResize(func(sender lcl.IObject, browser *cef.ICefBrowser, newSize *cef.TCefSize) bool {
		fmt.Println("OnAutoResize", newSize)
		return true
	})
	m.chromium.SetOnLoadingStateChange(func(sender lcl.IObject, browser *cef.ICefBrowser, isLoading, canGoBack, canGoForward bool) {
		fmt.Println("OnLoadingStateChange:", isLoading, canGoBack, canGoForward)
		if m.toolBar != nil {
			m.toolBar.UpdateBrowserStatus(isLoading, canGoBack, canGoForward)
			m.toolBar.locationBar.SetText(browser.MainFrame().Url())
		}
	})
	m.chromium.SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
		fmt.Println("OnLoadEnd httpStatusCode:", httpStatusCode)
	})
	m.chromium.SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, regions *cef.TCefDraggableRegions) {
		fmt.Println("OnDraggableRegionsChanged RegionsCount:", regions.RegionsCount())
		//for i := 0; i < regions.RegionsCount(); i++ {
		//	fmt.Println("Region:", regions.Region(i))
		//	regions.Region(i).Bounds.X = 0
		//	regions.Region(i).Bounds.Y = 0
		//	regions.Region(i).Bounds.Height = 30
		//}
		//m.window.SetDraggableRegions(regions.Regions())
	})
	m.window.SetOnGetMinimumSize(func(view *cef.ICefView, result *cef.TCefSize) {
		if view.GetID() == ID_WINDOW {
			//fmt.Println("OnGetMinimumSize", result)
			*result = minimumWindowSize
		}
	})
	m.window.SetOnGetLinuxWindowProperties(func(window *cef.ICefWindow, properties *cef.TLinuxWindowProperties, result *bool) {
		fmt.Println("OnGetLinuxWindowProperties", *result)
		properties.WmClassName = "energy-cefclient"
		properties.WmClassClass = "energy-cefclient"
		properties.WmRoleName = "energy-cefclient"
		*result = true
	})
	m.window.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
		//m.window.SetBackgroundColor(cef.CefColorSetARGB(255, 33, 34, 38))
		fmt.Println("OnGetPreferredSize")
	})
	m.window.SetOnGetTitleBarHeight(func(window *cef.ICefWindow, titleBarHeight *float32, result *bool) {
		fmt.Println("OnGetTitleBarHeight:", *titleBarHeight, *result)
		*titleBarHeight = 30
		*result = true
	})
	m.window.SetOnWindowBoundsChanged(func(window *cef.ICefWindow, newBounds cef.TCefRect) {
		//fmt.Println("OnWindowBoundsChanged", newBounds)
		//m.window.SizeToPreferredSize()
		//m.browserView.SetSize(cef.TCefSize{Width: newBounds.Width - minimumWindowSize.Height, Height: newBounds.Height})
	})
	m.window.SetOnCanClose(func(window *cef.ICefWindow, result *bool) {
		fmt.Println("OnCanClose:", *result)
		*result = m.chromium.TryCloseBrowser()
	})
	m.window.SetOnIsFrameless(func(window *cef.ICefWindow, result *bool) {
		*result = true
	})
	m.window.SetOnGetInitialBounds(func(window *cef.ICefWindow, result *cef.TCefRect) {
		fmt.Println("OnGetInitialBounds")
		result.Width = width
		result.Height = height
	})
	m.window.SetOnWindowActivationChanged(func(window *cef.ICefWindow, active bool) {
		fmt.Println("OnWindowActivationChanged")
	})
	m.window.SetOnGetInitialBounds(func(window *cef.ICefWindow, result *cef.TCefRect) {
		result.Width = 800
		result.Height = 600
		fmt.Println("OnGetInitialBounds")
	})
	m.window.SetOnThemeColorsChanged(func(window *cef.ICefWindow, chromeTheme int32) {
		views_style.WindowApplyTo(window)
	})
	const (
		cmdIdCtrlA int32 = iota + 32800
		cmdIdCtrlC
		cmdIdCtrlX
		cmdIdCtrlV
	)
	var addAccelerator = func() {
		m.window.SetAccelerator(cmdIdCtrlA, 'A', false, true, false, false)
		m.window.SetAccelerator(cmdIdCtrlC, 'C', false, true, false, false)
		m.window.SetAccelerator(cmdIdCtrlX, 'X', false, true, false, false)
		m.window.SetAccelerator(cmdIdCtrlV, 'V', false, true, false, false)
	}
	m.window.SetOnAccelerator(func(window *cef.ICefWindow, commandId int32, result *bool) {
		fmt.Println("OnAccelerator commandId:", commandId)
		if ID_QUIT == commandId {
			//表示已处理，否则还会执行多次
			m.chromium.CloseBrowser(true)
			*result = true
		} else if cmdIdCtrlA == commandId {
			m.chromium.Browser().MainFrame().SelectAll()
			*result = true
		} else if cmdIdCtrlC == commandId {
			m.chromium.Browser().MainFrame().Copy()
			*result = true
		} else if cmdIdCtrlX == commandId {
			m.chromium.Browser().MainFrame().Cut()
			*result = true
		} else if cmdIdCtrlV == commandId {
			m.chromium.Browser().MainFrame().Paste()
			*result = true
		}
	})
	m.window.SetOnWindowChanged(func(view *cef.ICefView, added bool) {
		fmt.Println("OnWindowChanged added:", added)
		if added {
		}
	})
	m.window.SetOnGetWindowRuntimeStyle(func(result *consts.TCefRuntimeStyle) {
		fmt.Println("OnGetWindowRuntimeStyle RuntimeStyle:", *result)
		//*result = consts.CEF_RUNTIME_STYLE_CHROME
		//*result = consts.CEF_RUNTIME_STYLE_ALLOY
	})
	//m.window.SetOnKeyEvent(func(window *cef.ICefWindow, event *cef.TCefKeyEvent, result *bool) {
	//	fmt.Println("OnKeyEvent")
	//})
	//m.window.SetOnWithStandardWindowButtons()
	m.window.SetOnWindowCreated(func(window *cef.ICefWindow) {
		fmt.Println("OnWindowCreated")
		m.window.SetID(ID_WINDOW)
		//window.SetThemeColor()
		m.window.ThemeChanged()
		m.window.SetWindowIcon(LoadImage("app-icon.png"))
		m.window.SetWindowAppIcon(LoadImage("app-icon.png"))
		m.window.SetTitle("Go ENERGY Client")
		m.window.SetLinuxWindowProperties("energy-cefclient", "energy-cefclient")

		m.titleBar = NewTitleBar(m.window) // 顶部标题栏
		m.menuBar = NewMenuBar(m.window)   // 顶部菜单栏
		m.toolBar = NewToolBar(m.window)   // 顶部工具栏

		if m.chromium.CreateBrowserByBrowserViewComponent(m.homePage, m.browserView, nil, nil) {
			//regions := []cef.TCefDraggableRegion{}
			//regions = append(regions, cef.TCefDraggableRegion{Bounds: cef.TCefRect{X: 130, Y: 0, Width: 100, Height: 30}})
			//m.window.SetDraggableRegions(regions)
			fmt.Println("ChromeToolbar:", m.browserView.GetChromeToolbar().IsValid())
			fmt.Println("WindowHandle:", m.window.WindowHandle())
			// 允许|browser_view_|增长并填充任何剩余空间。
			windowLayout := m.window.SetToBoxLayout(cef.TCefBoxLayoutSettings{
				BetweenChildSpacing: 5,
				CrossAxisAlignment:  consts.CEF_AXIS_ALIGNMENT_STRETCH,
			})
			// 菜单栏, 创建菜单，并添加到菜单栏中
			m.menuBar.CreateFileMenuItems()
			m.menuBar.CreateTestMenuItems()

			// 工具栏, 创建工具组件，并添加到工具栏中
			m.toolBar.CreateToolComponent()

			m.titleBar.CreateTitleBrowser(m)
			//var minWidth int32 = toolBar.AllButtonWidth()
			//var minHeight int32 = toolBar.EnsureToolPanel().GetBounds().Height + 100
			//minimumWindowSize = cef.TCefSize{Width: minWidth, Height: minHeight}
			//fmt.Println("minWidth:", minWidth, "minHeight:", minHeight)

			// 标题栏添加到窗口
			m.window.AddChildView(m.titleBar.EnsureTitlePanel().AsView())

			// 菜单栏添加到窗口
			m.window.AddChildView(m.menuBar.EnsureMenuPanel().AsView())

			// 工具栏添加到窗口
			m.window.AddChildView(m.toolBar.EnsureToolPanel().AsView())

			// 浏览器view添加到窗口
			m.window.AddChildView(m.browserView.AsView())
			windowLayout.SetFlexForView(m.browserView.AsView(), 1)

			m.window.Layout()

			addAccelerator()
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
