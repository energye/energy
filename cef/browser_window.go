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
	"github.com/energye/energy/v3/logger"
	"github.com/energye/energy/v3/window"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
)

var (
	gPrePopupWindow *TPopupWindow
)

type TPopupWindow struct {
	window.TWindow
	browser *TBrowser
}

func NewPopupWindow() *TPopupWindow {
	m := &TPopupWindow{}
	lcl.Application.NewForm(m)
	return m
}

func (m *TPopupWindow) FormCreate(sender lcl.IObject) {
	logger.Debug("PopupWindow.FormCreate")
	m.InternalBeforeFormCreate()

	m.SetCaption("ENERGY - CEF Simple 测试示例")

	m.browser = NewBrowser(m)
	m.browser.SetAlign(types.AlClient)
	m.browser.SetParent(m)
	m.browser.SetWindow(m)
	m.browser.Chromium().SetDefaultUrl("fs://energy/index-home.html")

	m.TWindow.FormCreate(sender)
}

func (m *TPopupWindow) OnShow(sender lcl.IObject) {
	logger.Debug("PopupWindow.OnShow")
	m.TWindow.OnShow(sender)
}
