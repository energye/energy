package main

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/config"
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"path/filepath"
)

var (
	chromium cef.IChromium
)

func main() {
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	app := cef.CreateApplication()
	app.SetFrameworkDirPath(config.Get().FrameworkPath())
	app.SetMultiThreadedMessageLoop(false)
	app.SetExternalMessagePump(false)
	app.SetDisablePopupBlocking(true)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetLocale(consts.LANGUAGE_zh_CN)
	app.SetFrameworkDirPath(config.Get().FrameworkPath())
	if common.IsLinux() {
		app.SetDisableZygote(true)
	}
	app.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized")
		fmt.Println("  GetScreenDPI:", cef.GetScreenDPI(), "GetDeviceScaleFactor:", cef.GetDeviceScaleFactor())
		handle := cef.InitializeWindowHandle()
		rect := types.TRect{}
		chromium = cef.NewChromium(nil, nil)
		chromium.SetDefaultURL("https://www.baidu.com")
		chromium.SetOnBeforeClose(func(sender lcl.IObject, browser *cef.ICefBrowser) {
			app.QuitMessageLoop()
		})
		var tabURL string
		chromium.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
			fmt.Println("OnLoadStart", browser.Identifier())
			if tabURL != "" {
				frame.LoadUrl(tabURL)
				tabURL = ""
			}
		})
		chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			browser.ExecuteChromeCommand(consts.IDC_NEW_TAB, consts.CEF_WOD_CURRENT_TAB)
			tabURL = beforePopupInfo.TargetUrl
			fmt.Println("OnBeforePopup", tabURL)
			return true
		})
		chromium.SetOnOpenUrlFromTab(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool {
			fmt.Println("OpenUrlFromTab", tabURL)
			return false
		})
		chromium.CreateBrowserByWindowHandle(handle, rect, "tiny browser", nil, nil, true)
	})
	app.SetOnGetDefaultClient(func(client *cef.ICefClient) {
		fmt.Println("OnGetDefaultClient:", chromium)
		if chromium != nil {
			*client = *chromium.Client()
		}
	})
	if app.StartMainProcess() {
		fmt.Println("StartMainProcess Success")
		app.RunMessageLoop()
	}
}
