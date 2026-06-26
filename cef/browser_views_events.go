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
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
)

func (m *TViewsBrowser) initViewsBrowserDefaultEvent() {
	logger.Debug("Browser.initViewsBrowserDefaultEvent")

	m.chromium.SetOnAfterCreated(m.chromiumOnAfterCreated)
}

func (m *TViewsBrowser) chromiumOnAfterCreated(sender lcl.IObject, browser cef.ICefBrowser) {
	logger.Debug("Chromium.OnAfterCreated", browser.GetIdentifier())
	if m.browserId == 0 {
		m.browserId = uint32(browser.GetIdentifier())
		// ipc
		ipc.RegisterProcessMessage(m)
		// local load
		//m.schemeHandlerFactory = createSchemeHandlerFactory(browser)
		// pre-creates a window
		if m.options.AutoPopupWindow {

		}
	}
	if m.onBrowserAfterCreated != nil {
		m.onBrowserAfterCreated(sender)
	}
}
