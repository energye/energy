//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"fmt"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
)

type ViewsFrameworkBrowserWindow struct {
}

func (m *browserWindow) appContextInitialized(app *TCEFApplication) {
	app.SetOnContextInitialized(func() {
		fmt.Println("OnContextInitialized()")
		component := lcl.NewComponent(nil)
		if BrowserWindow.Config.chromiumConfig == nil {
			BrowserWindow.Config.chromiumConfig = NewChromiumConfig()
			BrowserWindow.Config.chromiumConfig.SetEnableMenu(true)
			BrowserWindow.Config.chromiumConfig.SetEnableDevTools(true)
			BrowserWindow.Config.chromiumConfig.SetEnableOpenUrlTab(true)
			BrowserWindow.Config.chromiumConfig.SetEnableWindowPopup(true)
		}
		m.chromium = NewChromium(component, BrowserWindow.Config.chromiumConfig)
		m.chromium.SetEnableMultiBrowserMode(true)
		m.browserViewComponent = NewBrowserViewComponent(component)
		m.windowComponent = NewWindowComponent(component)

		m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool {
			fmt.Println("OnBeforePopup TargetUrl:", beforePopupInfo.TargetUrl)

			return false
		})
		m.windowComponent.SetOnWindowCreated(func(sender lcl.IObject, window *ICefWindow) {
			if m.chromium.CreateBrowserByBrowserViewComponent(BrowserWindow.Config.DefaultUrl, m.browserViewComponent) {
				m.windowComponent.AddChildView(m.browserViewComponent)
				m.windowComponent.SetTitle(BrowserWindow.Config.Title)
				window.CenterWindow(NewCefSize(BrowserWindow.Config.Width, BrowserWindow.Config.Height))
				m.browserViewComponent.RequestFocus()
				if BrowserWindow.Config.Icon != "" {
					window.SetWindowAppIconFS(1, BrowserWindow.Config.Icon)
				}
				if BrowserWindow.Config.browserWindowOnEventCallback != nil {
					BrowserWindow.browserEvent.chromium = m.chromium
					BrowserWindow.Config.browserWindowOnEventCallback(BrowserWindow.browserEvent, m.windowInfo)
				}
				window.Show()
			}
		})
		m.windowComponent.SetOnCanClose(func(sender lcl.IObject, window *ICefWindow, aResult *bool) {
			fmt.Println("OnCanClose")
			app.QuitMessageLoop()
			*aResult = true
		})
		m.windowComponent.SetOnGetInitialBounds(func(sender lcl.IObject, window *ICefWindow, aResult *TCefRect) {
			fmt.Println("OnGetInitialBounds")
		})
		m.windowComponent.SetOnGetInitialShowState(func(sender lcl.IObject, window *ICefWindow, aResult *consts.TCefShowState) {
			fmt.Println("OnGetInitialShowState", *aResult)
		})
		m.windowComponent.CreateTopLevelWindow()
	})
}
