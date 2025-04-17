package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"path/filepath"
)

//go:embed resources
var resources embed.FS

var (
	app    *cef.TCEFApplication
	vf     *ViewsFramework
	width  int32 = 1280
	height int32 = 900
)

type ViewsFramework struct {
	*lcl.TComponent
	chromium             cef.IChromium
	windowComponent      *cef.TCEFWindowComponent
	browserViewComponent *cef.TCEFBrowserViewComponent
	homePage             string
}

func (m *ViewsFramework) Create() {
	m.TComponent = lcl.NewComponent(nil)
	m.chromium = cef.NewChromium(m, nil)
	m.chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("OnBeforeClose")
		app.QuitMessageLoop()
	})
	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		fmt.Println("OnBeforePopup")
		// 禁止弹出窗口
		return true
	})
	m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
		fmt.Println("OnTitleChange:", title)
		m.windowComponent.SetTitle(title)
	})
	m.browserViewComponent = cef.BrowserViewComponentRef.New(m)
	m.windowComponent = cef.WindowComponentRef.New(m)
	m.windowComponent.SetOnWindowCreated(func(window *cef.ICefWindow) {
		fmt.Println("OnWindowCreated")
		if m.chromium.CreateBrowserByBrowserViewComponent(m.homePage, m.browserViewComponent, nil, nil) {
			m.windowComponent.AddChildView(m.browserViewComponent.BrowserView().AsView())
			// 窗口居中
			display := m.windowComponent.Display()
			if display.IsValid() {
				workArea := display.WorkArea()
				position := cef.TCefPoint{
					X: ((workArea.Width - width) / 2) + workArea.X,
					Y: ((workArea.Height - height) / 2) + workArea.Y,
				}
				m.windowComponent.SetPosition(position)
			}
			// 应用图标
			data, _ := resources.ReadFile("resources/icon.png")
			icon := cef.ImageRef.New()
			icon.AddPng(1, data)
			m.windowComponent.SetWindowAppIcon(icon)
			m.browserViewComponent.RequestFocus()
			m.windowComponent.Show()
		}
	})
	m.windowComponent.SetOnCanClose(func(window *cef.ICefWindow, result *bool) {
		fmt.Println("OnCanClose:", *result)
		*result = m.chromium.TryCloseBrowser()
	})
	m.windowComponent.SetOnGetInitialBounds(func(window *cef.ICefWindow, result *cef.TCefRect) {
		fmt.Println("OnGetInitialBounds")
		result.Width = width
		result.Height = height
	})
}

func (m *ViewsFramework) Client() *cef.ICefClient {
	if m.chromium.IsValid() {
		return m.chromium.Client()
	}
	return nil
}

func (m *ViewsFramework) CreateTopLevelWindow() {
	fmt.Println("CreateTopLevelWindow")
	if m.windowComponent != nil {
		m.windowComponent.CreateTopLevelWindow()
	}
}

func main() {
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	app = cef.CreateApplication()
	app.SetFrameworkDirPath(config.Get().FrameworkPath())
	app.SetMultiThreadedMessageLoop(false)
	app.SetExternalMessagePump(false)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetLocale(consts.LANGUAGE_zh_CN)
	if common.IsLinux() {
		app.SetDisableZygote(true)
	}
	app.SetEnableGPU(true)
	app.SetOnBeforeChildProcessLaunch(func(commandLine *cef.ICefCommandLine) {
		commandLine.AppendSwitch("enable-unsafe-swiftshader")
	})
	// 创建 ViewsFramework
	app.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized")
		fmt.Println("  GetScreenDPI:", cef.GetScreenDPI(), "GetDeviceScaleFactor:", cef.GetDeviceScaleFactor())
		vf = new(ViewsFramework)
		vf.homePage = "https://www.baidu.com"
		vf.Create()
		vf.CreateTopLevelWindow()
	})
	// 此事件仅用于“ChromeRuntime”模式
	app.SetOnGetDefaultClient(func(client *cef.ICefClient) {
		fmt.Println("OnGetDefaultClient", vf != nil)
		if vf != nil {
			*client = *vf.Client()
		}
	})
	if app.StartMainProcess() {
		fmt.Println("StartMainProcess Success")
		app.RunMessageLoop()
	}
}
