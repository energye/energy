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
	"github.com/energye/energy/v2/consts"
)

// Instance 实例
func (m *ICefBrowserViewDelegate) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefBrowserViewDelegate) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefBrowserViewDelegate) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

type onBrowserCreated func(browserView *ICefBrowserView, browser *ICefBrowser)
type onBrowserDestroyed func(browserView *ICefBrowserView, browser *ICefBrowser)
type onGetDelegateForPopupBrowserView func(browserView *ICefBrowserView, browserSettings *TCefBrowserSettings, client *ICefClient, isDevtools bool, aResult *ICefBrowserViewDelegate)
type onPopupBrowserViewCreated func(browserView, popupBrowserView *ICefBrowserView, isDevtools bool, aResult *bool)
type onGetChromeToolbarType func(aResult *consts.TCefChromeToolbarType)
