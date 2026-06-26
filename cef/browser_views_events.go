//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
)

func (m *TViewsBrowser) initViewsBrowserDefaultEvent() {
	logger.Debug("Browser.initViewsBrowserDefaultEvent")

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)

	// close browser
	m.chromium.SetOnBeforeClose(m.chromiumOnBeforeClose)
	m.chromium.SetOnClose(m.chromiumOnClose)
}

func (m *TViewsBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnAfterCreated", browser.GetIdentifier())
	if m.browserId == 0 {
		m.browserId = uint32(browser.GetIdentifier())
		// ipc
		ipc.RegisterProcessMessage(m)
		// pre-creates a window
		if m.options.AutoPopupWindow {

		}
	}
	if m.onBrowserAfterCreated != nil {
		m.onBrowserAfterCreated(sender)
	}
}

func (m *TViewsBrowser) chromiumOnClose(sender lcl.IObject, browser cef.ICefBrowser, action *cefTypes.TCefCloseBrowserAction) {
	logger.Debug("Chromium.OnClose", browser.GetIdentifier())

}

func (m *TViewsBrowser) chromiumOnBeforeClose(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnBeforeClose", "Current-BrowserID:", m.browserId, "Target-BrowserID:", browser.GetIdentifier())
	if m.browserId != uint32(browser.GetIdentifier()) {
		logger.Debug("Chromium.OnBeforeClose Non-current user browser")
		return
	}
}
