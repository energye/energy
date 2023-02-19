//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF Browser View
package cef

import "unsafe"

// ICefBrowserView TODO 还未实现
type ICefBrowserView struct {
	instance unsafe.Pointer
}

func (m *ICefBrowserView) GetBrowser() *ICefBrowser {
	return nil
}
func (m *ICefBrowserView) GetChromeToolbar() *ICefView {
	return nil
}
func (m *ICefBrowserView) SetPreferAccelerators(preferAccelerators bool) {
}
