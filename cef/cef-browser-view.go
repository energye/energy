package cef

import "unsafe"

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
