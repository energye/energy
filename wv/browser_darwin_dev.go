//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin && dev

package wv

import wv "github.com/energye/wv/darwin"

// _EnableDevTools 启用开发者工具
func (m *TWebview) _EnableDevTools(preference wv.IWkPreferences) {
	if !gApplication.Options.DisableDevTools {
		preference.SetBoolValueForKey(true, "developerExtrasEnabled")
	}
}
