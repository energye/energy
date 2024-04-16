package main

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"os"
	"path/filepath"
)

var (
	chromium cef.IChromium
)

func main() {
	cef.GlobalInit(nil, nil)
	rootCache := filepath.Join(consts.CurrentExecuteDir, "rootcache")
	app := cef.CreateApplication()
	app.SetFrameworkDirPath(os.Getenv("ENERGY_HOME"))
	app.SetMultiThreadedMessageLoop(false)
	app.SetExternalMessagePump(false)
	app.SetChromeRuntime(true) // 启用此行可启用“ChromeRuntime”模式。它处于实验状态。
	app.SetDisablePopupBlocking(true)
	app.SetRootCache(rootCache)
	app.SetCache(filepath.Join(rootCache, "cache"))
	app.SetLocale(consts.LANGUAGE_zh_CN)
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
		chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
			return true
		})
		chromium.SetOnOpenUrlFromTab(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool {
			return true
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
